package command

import (
	"github.com/HETIC-MT-P2021/chen-discord-bot/discord"
	"github.com/HETIC-MT-P2021/chen-discord-bot/pokeapi"
	"github.com/bwmarrin/discordgo"
)

func getAuthorPokedex() *discordgo.MessageEmbedAuthor {
	return &discordgo.MessageEmbedAuthor{
		Name:    "Pokédex",
		IconURL: "https://icon-library.com/images/pokedex-icon/pokedex-icon-20.jpg",
	}
}

func sendCard(ctx *discord.Context, pokemon string) {
	var err error
	p, err := pokeapi.FindPokemon(pokemon, "fr")
	if err != nil {
		ctx.Send(discord.ErrorMessage("Not Found", "Pokemon \""+pokemon+"\" was not found."))
		return
	}

	err = ctx.SendEmbed(cardEmbed(*p))

	if err != nil {
		ctx.Send(discord.ErrorMessage("Bot error", "Error sending the message."))
		return
	}
}
