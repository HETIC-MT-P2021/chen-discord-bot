package discord

import (
	"github.com/bwmarrin/discordgo"
	"sort"
	"strings"
)

// Router represents a DiscordGo discord router
type Router struct {
	Prefixes         []string
	IgnorePrefixCase bool
	BotsAllowed      bool
	Commands         []*Command
}

func Create(router *Router) *Router {
	return router
}

// RegisterCmd registers a new discord
func (router *Router) RegisterCmd(command *Command) {
	router.Commands = append(router.Commands, command)
}

// GetCmd returns the command with the given name if it exists
func (router *Router) GetCmd(name string) *Command {
	// Sort the commands slice using the length of the command name
	sort.Slice(router.Commands, func(i, j int) bool {
		return len(router.Commands[i].Name) > len(router.Commands[j].Name)
	})

	// Loop through all commands to find the correct one
	for _, command := range router.Commands {
		// Define the slice to check
		toCheck := buildCheckPrefixes(command)

		// Check the prefix of the string
		if stringArrayContains(toCheck, name, command.IgnoreCase) {
			return command
		}
	}
	return nil
}

// Initialize initializes the message event listener
func (router *Router) Initialize(session *discordgo.Session) {
	session.AddHandler(router.Handler())
}

// Handler provides the discordgo handler for the given router
func (router *Router) Handler() func(*discordgo.Session, *discordgo.MessageCreate) {
	return func(session *discordgo.Session, message *discordgo.MessageCreate) {
		// Check if the message was sent by a bot
		if message.Author.Bot && !router.BotsAllowed {
			return
		}

		// Check if the message starts with one of the defined prefixes
		hasPrefix, content := stringHasPrefix(message.Content, router.Prefixes, router.IgnorePrefixCase)
		if !hasPrefix {
			return
		}

		// Check if the message is empty after the prefix processing
		args := strings.Split(message.Content, " ")[1:]
		if len(args) == 0 {
			session.ChannelMessageSend(message.ChannelID, ErrorMessage("Command missing", "For a list of discord type !poke help"))
			return
		}

		// Check if the message starts with a discord name
		for _, command := range router.Commands {
			toCheck := buildCheckPrefixes(command)

			// Check if the content is the current discord
			isCommand, content := stringHasPrefix(content, toCheck, command.IgnoreCase)
			if !isCommand {
				continue
			}

			// Check if the remaining string is empty or starts with a space or newline
			isValid, content := stringHasPrefix(content, []string{" ", "\n"}, false)
			if content == "" || isValid {
				command.Handler(&Context{
					Session:       session,
					Event:         message,
					Arguments:     args,
					Router:        router,
					Command:       command,
				})
			}
		}
	}
}

func buildCheckPrefixes(command *Command) []string {
	// Define an array containing the discord name and the aliases
	toCheck := make([]string, len(command.Aliases)+1)
	toCheck = append(toCheck, command.Name)
	toCheck = append(toCheck, command.Aliases...)

	sort.Slice(toCheck, func(i, j int) bool {
		return len(toCheck[i]) > len(toCheck[j])
	})

	return toCheck
}