package command

import (
	"github.com/HETIC-MT-P2021/chen-discord-bot/discord"
	"github.com/HETIC-MT-P2021/chen-discord-bot/pokeapi"
	"github.com/bwmarrin/discordgo"
)

// Sends a specific pokemon card
func sendCardPokemon(ctx *discord.Context) {
	var err error
	name := ctx.Arguments[1]

	p, err := pokeapi.FindPokemon(name, "fr")
	if err != nil {
		ctx.Send(discord.ErrorMessage("Not Fount", "Pokemon \""+name+"\" was not found."))
		return
	}

	err = ctx.SendEmbed(cardEmbed(*p))

	if err != nil {
		ctx.Send(discord.ErrorMessage("Bot error", "Error sending the message."))
		return
	}
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
			ctx.Send(discord.ErrorMessage("Invalid Command", "Ensure that you have selected a pokemon in the format of: `!poke card <name or ID>`."))
			return
		}

		sendCardPokemon(ctx)
	}
}
