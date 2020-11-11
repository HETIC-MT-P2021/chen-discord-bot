package command

import "github.com/HETIC-MT-P2021/chen-discord-bot/discord"

func InitRouter(router *discord.Router) {
	router.RegisterCmd(&discord.Command{
		Name:        "card",
		Description: "Return data from a specific Pokémon",
		Usage:       "card <name or id>",
		Example:     "card <bulbasaur or 1>",
		IgnoreCase:  true,
		Handler:     CardCommandHandler(),
	})
}
