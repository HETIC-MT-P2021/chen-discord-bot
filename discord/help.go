package discord

import (
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
)

// RegisterDefaultHelpCommand registers the default help command
func (router *Router) RegisterDefaultHelpCommand() {
	// Register the default help command
	router.RegisterCmd(&Command{
		Name:        "help",
		Description: "Lists all the available commands or displays some information about a specific command",
		Usage:       "help [command name]",
		Example:     "help yourCommand",
		IgnoreCase:  true,
		Handler:     generalHelpCommand,
	})
}

// generalHelpCommand handles the general help command
func generalHelpCommand(ctx *Context) {
	// Check if the user provided an argument
	if len(ctx.Arguments) > 1 {
		specificHelpCommand(ctx)
		return
	}

	// Send the general help embed
	err := ctx.SendEmbed(generalHelpEmbed(ctx.Router))
	if err != nil {
		ctx.Send(ErrorMessage("Bot error", "Error sending the message."))
		return
	}
}

// generalHelpEmbed renders the general help embed on the given page
func generalHelpEmbed(router *Router) *discordgo.MessageEmbed {
	// Define useful variables
	commands := router.Commands
	prefix := router.Prefixes[0]

	// Prepare the fields for the embed
	fields := make([]*discordgo.MessageEmbedField, len(commands))
	for index, command := range commands {
		fields[index] = &discordgo.MessageEmbedField{
			Name:   command.Name,
			Value:  "`" + command.Description + "`",
			Inline: false,
		}
	}

	// Return the embed and the new page
	return &discordgo.MessageEmbed{
		Type:        "rich",
		Title:       "Command List",
		Description: "These are all the available commands. Type `" + prefix + " help <command name>` to find out more about a specific command.",
		Timestamp:   time.Now().Format(time.RFC3339),
		Color:       0xffff00,
		Fields:      fields,
	}
}

// specificHelpCommand handles the specific help command
func specificHelpCommand(ctx *Context) {
	// Define the command
	command := ctx.Router.GetCmd(ctx.Arguments[1])

	// Send the help embed
	ctx.SendEmbed(specificHelpEmbed(ctx, command))
}

// specificHelpEmbed renders the specific help embed of the given command
func specificHelpEmbed(ctx *Context, command *Command) *discordgo.MessageEmbed {
	// Define useful variables
	prefix := ctx.Router.Prefixes[0]

	// Check if the command is invalid
	if command == nil {
		return &discordgo.MessageEmbed{
			Type:      "rich",
			Title:     "Error",
			Timestamp: time.Now().Format(time.RFC3339),
			Color:     0xff0000,
			Fields: []*discordgo.MessageEmbedField{
				{
					Name:   "Message",
					Value:  "```The given command doesn't exist. Type `" + prefix + " help` for a list of available commands.```",
					Inline: false,
				},
			},
		}
	}

	// Define the aliases string
	aliases := "No aliases"
	if len(command.Aliases) > 0 {
		aliases = "`" + strings.Join(command.Aliases, "`, `") + "`"
	}

	// Return the embed
	return &discordgo.MessageEmbed{
		Type:        "rich",
		Title:       "Command Information",
		Description: "Displaying the information for the `" + command.Name + "` command.",
		Timestamp:   time.Now().Format(time.RFC3339),
		Color:       0xffff00,
		Fields: []*discordgo.MessageEmbedField{
			{
				Name:   "Name",
				Value:  "`" + command.Name + "`",
				Inline: false,
			},
			{
				Name:   "Aliases",
				Value:  aliases,
				Inline: false,
			},
			{
				Name:   "Description",
				Value:  "```" + command.Description + "```",
				Inline: false,
			},
			{
				Name:   "Usage",
				Value:  "```" + prefix + " " + command.Usage + "```",
				Inline: false,
			},
			{
				Name:   "Example",
				Value:  "```" + prefix + " " + command.Example + "```",
				Inline: false,
			},
		},
	}
}
