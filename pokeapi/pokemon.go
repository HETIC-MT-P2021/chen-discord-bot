package pokeapi

import (
	"strconv"
	"strings"

	"github.com/mtslzr/pokeapi-go"
	"github.com/mtslzr/pokeapi-go/structs"
)

//go:generate mockgen -destination=mocks/mock_pokeapi.go -package=mocks . PokemonInterface
type PokemonInterface interface {
	ID() string
	Name() string
	Title() string
	Description() string
	Category() string
	Types() string
	Image() string
	Link() string
	Stats() []PokemonInfo
}

type Pokemon struct {
	Pokemon  structs.Pokemon
	Species  structs.PokemonSpecies
	Language string
}

type PokemonInfo struct {
	Name  string
	Value int
}

func FindPokemon(name string, lang string) (*Pokemon, error) {
	var err error

	p := new(Pokemon)

	p.Pokemon, err = pokeapi.Pokemon(name)
	if err != nil {
		return nil, err
	}

	p.Species, _ = pokeapi.PokemonSpecies(p.Pokemon.Name)
	p.Language = lang

	return p, nil
}

func (p Pokemon) ID() string {
	return strconv.Itoa(p.Pokemon.ID)
}

func (p Pokemon) Name() string {
	name := ""
	for _, n := range p.Species.Names {
		if n.Language.Name == p.Language {
			name = n.Name
		}
	}

	if name == "" {
		name = p.Pokemon.Name
	}

	return strings.Title(name)
}

func (p Pokemon) Title() string {
	return p.Name() + " #" + p.ID()
}

func (p Pokemon) Description() string {
	description := ""
	for _, v := range p.Species.FlavorTextEntries {
		if v.Language.Name == p.Language {
			description = strings.ReplaceAll(v.FlavorText, "\n", " ")
		}
	}
	return description
}

func (p Pokemon) Category() string {
	category := ""
	for _, c := range p.Species.Genera {
		if c.Language.Name == p.Language {
			category = c.Genus
		}
	}
	return category
}

func (p Pokemon) Types() string {
	var typeName []string
	for _, i := range p.Pokemon.Types {
		t, _ := pokeapi.Type(i.Type.Name)
		for _, j := range t.Names {
			if j.Language.Name == p.Language {
				typeName = append(typeName, j.Name)
			}
		}
	}
	return strings.Join(typeName, ", ")
}

func (p Pokemon) Image() string {
	return p.Pokemon.Sprites.FrontDefault
}

func (p Pokemon) Link() string {
	return "https://www.pokemon.com/us/pokedex/" + strings.ToLower(p.Pokemon.Name)
}

func (p Pokemon) Stats() []PokemonInfo {
	var stats []PokemonInfo
	for _, i := range p.Pokemon.Stats {
		s, _ := pokeapi.Stat(i.Stat.Name)
		for _, j := range s.Names {
			if j.Language.Name == p.Language {
				stats = append(stats, PokemonInfo{
					Name:  j.Name,
					Value: i.BaseStat,
				})
			}
		}
	}
	return stats
}
