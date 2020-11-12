package command

import "github.com/bwmarrin/discordgo"

func getAuthorPokedex() *discordgo.MessageEmbedAuthor {
	return &discordgo.MessageEmbedAuthor{
		Name:    "Pokédex",
		IconURL: "https://icon-library.com/images/pokedex-icon/pokedex-icon-20.jpg",
	}
}
