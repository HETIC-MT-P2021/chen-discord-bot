package command

import "github.com/HETIC-MT-P2021/chen-discord-bot/discord"

func InitRouter(router *discord.Router) {
	router.RegisterCmd(&discord.Command{
		Name:        "card",
		Description: "Return data from the Pokémon associated to the current user or a specific Pokémon with ID",
		Usage:       "card <name, id or nothing>",
		Example:     "card <bulbasaur or 1>",
		IgnoreCase:  true,
		Handler:     CardCommandHandler(),
	})

	router.RegisterCmd(&discord.Command{
		Name:        "list",
		Description: "Returns the list of users and their associated Pokémon",
		Usage:       "list",
		Example:     "list",
		IgnoreCase:  true,
		Handler:     ListCommandHandler(),
	})

	router.RegisterCmd(&discord.Command{
		Name:        "claim",
		Description: "Associate a random Pokémon to the current user or replace it",
		Usage:       "claim",
		Example:     "claim",
		IgnoreCase:  true,
		Handler:     ClaimCommandHandler(),
	})
}
