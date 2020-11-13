package command

import (
	"github.com/HETIC-MT-P2021/chen-discord-bot/discord"
	"github.com/bwmarrin/discordgo"
)

// Sends a specific pokemon card
func simplePokemonCard(ctx *discord.Context) {
	pokemon := ctx.Arguments[1]

	sendCard(ctx, pokemon)
}

// Sends the user pokemon card
func userPokemonCard(ctx *discord.Context) {
	var pokemonID string
	err := ctx.Sql.QueryRow("SELECT pokemon_id FROM user_pokemon WHERE user_id = ?", ctx.Event.Author.ID).Scan(&pokemonID)
	if err != nil {
		ctx.Send(discord.ErrorMessage("Not Found", "You doesn't have a Pokemon ! Use !poke claim to got one."))
		return
	}
	sendCard(ctx, pokemonID)
}

// Handles all cards related discord
func CardCommandHandler() discord.ExecutionHandler {
	return func(ctx *discord.Context) {
		c, err := ctx.Session.Channel(ctx.Event.ChannelID)
		if err != nil {
			ctx.Send(discord.ErrorMessage("Bot error", "Error getting channel."))
			return
		}

		if c.Type == discordgo.ChannelTypeDM {
			ctx.Send(discord.ErrorMessage("Invalid channel", "Cannot send invites from a DM"))
			return
		}

		if len(ctx.Arguments) == 1 {
			userPokemonCard(ctx)
		} else {
			simplePokemonCard(ctx)
		}
	}
}
