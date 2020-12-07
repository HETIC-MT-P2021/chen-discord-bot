package pokeapi

import (
	"testing"

	"github.com/mtslzr/pokeapi-go"
	"github.com/stretchr/testify/assert"
)

func TestFindPokemon(t *testing.T) {
	id := "1"
	p, _ := pokeapi.Pokemon(id)
	ps, _ := pokeapi.PokemonSpecies(p.Name)
	pokemon, _ := FindPokemon(id, "fr")

	assert.Equal(t, p, pokemon.Pokemon)
	assert.Equal(t, ps, pokemon.Species)

	assert.Equal(t, "fr", pokemon.Language)
}

func TestPokemon_ID(t *testing.T) {
	id := "1"
	pokemon, _ := FindPokemon(id, "fr")
	assert.Equal(t, id, pokemon.ID())
}

func TestPokemon_Name(t *testing.T) {
	pokemon, _ := FindPokemon("1", "fr")
	assert.Equal(t, "Bulbizarre", pokemon.Name())
}

func TestPokemon_Title(t *testing.T) {
	pokemon, _ := FindPokemon("1", "fr")
	assert.Equal(t, "Bulbizarre #1", pokemon.Title())
}

func TestPokemon_Description(t *testing.T) {
	pokemon, _ := FindPokemon("1", "fr")
	assert.Equal(t, "Quand il est jeune, il absorbe les nutriments conservés dans son dos pour grandir et se développer.", pokemon.Description())
}

func TestPokemon_Category(t *testing.T) {
	pokemon, _ := FindPokemon("1", "fr")
	assert.Equal(t, "Pokémon Graine", pokemon.Category())
}

func TestPokemon_Types(t *testing.T) {
	pokemon, _ := FindPokemon("1", "fr")
	assert.Equal(t, "Plante, Poison", pokemon.Types())
}

func TestPokemon_Image(t *testing.T) {
	pokemon, _ := FindPokemon("1", "fr")
	assert.Equal(t, "https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/1.png", pokemon.Image())
}

func TestPokemon_Link(t *testing.T) {
	pokemon, _ := FindPokemon("1", "fr")
	assert.Equal(t, "https://www.pokemon.com/us/pokedex/bulbasaur", pokemon.Link())
}

func TestPokemon_Stats(t *testing.T) {
	pokemon, _ := FindPokemon("1", "fr")
	stats := []PokemonInfo{
		{Name: "PV", Value: 45},
		{Name: "Attaque", Value: 49},
		{Name: "Défense", Value: 49},
		{Name: "Attaque Spéciale", Value: 65},
		{Name: "Défense Spéciale", Value: 65},
		{Name: "Vitesse", Value: 45},
	}
	assert.Equal(t, stats, pokemon.Stats())
}
