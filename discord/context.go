package discord

import (
	"database/sql"
	"github.com/bwmarrin/discordgo"
)

type Context struct {
	Session   *discordgo.Session
	Event     *discordgo.MessageCreate
	Arguments []string
	Router    *Router
	Command   *Command
	Sql       *sql.DB
}

// ExecutionHandler represents a handler for a context execution
type ExecutionHandler func(*Context)

// Send responds with the given text message
func (ctx *Context) Send(text string) error {
	_, err := ctx.Session.ChannelMessageSend(ctx.Event.ChannelID, text)
	return err
}

// SendEmbed responds with the given embed message
func (ctx *Context) SendEmbed(embed *discordgo.MessageEmbed) error {
	_, err := ctx.Session.ChannelMessageSendEmbed(ctx.Event.ChannelID, embed)
	return err
}
