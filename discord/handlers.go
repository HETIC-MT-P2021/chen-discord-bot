package discord

import (
	"strings"

	"github.com/bwmarrin/discordgo"
)

// Handles all checkers commands
func CommandsHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Ignore all messages created by the bot itself
	if m.Author.ID == s.State.User.ID {
		return
	}

	// Ignore all messages that don't have the !poke prefix
	if !strings.HasPrefix(m.Content, "!poke") {
		return
	}

	// Get the arguments
	args := strings.Split(m.Content, " ")[1:]
	// Ensure valid command
	if len(args) == 0 {
		s.ChannelMessageSend(m.ChannelID, errorMessage("Command missing", "For a list of commands type !poke help"))
		return
	}

	// Call the corresponding handler
	switch args[0] {
	case "ping":
		s.ChannelMessageSend(m.ChannelID, "Pong!")
	case "card":
		cardCommandHandler(s, m, args)
	default:
		s.ChannelMessageSend(m.ChannelID, errorMessage("Invalid command", "For a list of commands type !poke help"))
	}
}
