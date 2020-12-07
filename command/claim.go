package command

import (
	"math/rand"

	"github.com/HETIC-MT-P2021/chen-discord-bot/discord"
	"github.com/bwmarrin/discordgo"
)

// Associate Pokemon to a user
func AssociatePokemonToUser(ctx *discord.Context) {

	var randomPokemonID = rand.Intn(898)

	const query = "REPLACE INTO user_pokemon VALUES ($1, $2)"
	tx, err := ctx.Sql.Begin()
	if err != nil {
		ctx.Send(discord.ErrorMessage("Bot error", "Database error."))
		return
	}
	_, err = tx.Exec(query, ctx.Event.Author.ID, randomPokemonID)
	if err != nil {
		tx.Rollback()
		ctx.Send(discord.ErrorMessage("Bot error", "Database error."))
		return
	}
	tx.Commit()
	userPokemonCard(ctx)
}

// Handles all list related discord
func ClaimCommandHandler() discord.ExecutionHandler {
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

		if len(ctx.Arguments) > 1 {
			ctx.Send(discord.ErrorMessage("Invalid Command", "No parameter is required for this command, just use: `!poke claim`."))
			return
		}

		AssociatePokemonToUser(ctx)
	}
}
