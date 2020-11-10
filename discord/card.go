package discord

import (
	"github.com/HETIC-MT-P2021/chen-discord-bot/api"
	"github.com/bwmarrin/discordgo"
)

// Sends a specific pokemon card
func sendCardPokemon(s *discordgo.Session, m *discordgo.MessageCreate, cmd []string) {
	var err error
	name := cmd[1]

	p, err := api.FindPokemon(name, "fr")
	if err != nil {
		s.ChannelMessageSend(m.ChannelID, errorMessage("Not Fount", "Pokemon \""+name+"\" was not found."))
		return
	}

	_, err = s.ChannelMessageSendEmbed(m.ChannelID, cardEmbed(*p))

	if err != nil {
		s.ChannelMessageSend(m.ChannelID, errorMessage("Bot error", "Error sending the message."))
		return
	}
}

// Handles all cards related commands
func cardCommandHandler(s *discordgo.Session, m *discordgo.MessageCreate, cmd []string) {
	c, err := s.Channel(m.ChannelID)
	if err != nil {
		s.ChannelMessageSend(m.ChannelID, errorMessage("Bot error", "Error getting channel."))
		return
	}

	if c.Type == discordgo.ChannelTypeDM {
		s.ChannelMessageSend(m.ChannelID, errorMessage("Invalid channel", "Cannot send invites from a DM"))
		return
	}

	if len(cmd) == 0 {
		s.ChannelMessageSend(m.ChannelID, errorMessage("Invalid Command", "Ensure that you have selected a pokemon in the format of: `!poke card <name>`."))
	}

	sendCardPokemon(s, m, cmd)
}
