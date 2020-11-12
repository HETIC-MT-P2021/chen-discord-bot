package command

import (
	"fmt"
	"github.com/HETIC-MT-P2021/chen-discord-bot/discord"
	"github.com/HETIC-MT-P2021/chen-discord-bot/pokeapi"
	"github.com/bwmarrin/discordgo"
	"strings"
	"time"
)

// Sends the list of users and their associated Pokémon
func listAllUsersPoekemon(ctx *discord.Context) {
	var err error
	var userList []string

	rows, _ := ctx.Sql.Query("SELECT user_id, pokemon_id FROM user_pokemon ORDER BY user_id")

	var userId string
	var pokemonId string
	for rows.Next() {
		rows.Scan(&userId, &pokemonId)
		pokemonName, _ := pokeapi.FindPokemon(pokemonId, "fr")

		userName, _ := ctx.Session.User(userId)
		userList = append(userList, fmt.Sprintf("%s • %s", userName.Username, pokemonName.Title()))
	}

	err = ctx.SendEmbed(&discordgo.MessageEmbed{
		Title:       "List users",
		Description: strings.Join(userList, "\n"),
		Color:       discord.C_RED,
		Timestamp:   time.Now().Format(time.RFC3339),
	})
	if err != nil {
		ctx.Send(discord.ErrorMessage("Bot error", "Error sending the message."))
		return
	}
}

// Handles all list related discord
func ListCommandHandler() discord.ExecutionHandler {
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
			ctx.Send(discord.ErrorMessage("Invalid Command", "No parameter is required for this command, just use: `!poke list`."))
			return
		}

		listAllUsersPoekemon(ctx)
	}
}
