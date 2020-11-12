package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/HETIC-MT-P2021/chen-discord-bot/database"
	"github.com/HETIC-MT-P2021/chen-discord-bot/discord"

	"github.com/bwmarrin/discordgo"
)

// Variables used for command line parameters
var (
	Token string
)

func init() {

	flag.StringVar(&Token, "t", "", "Bot Token")
	flag.Parse()
}

func main() {

	// Create a new SQLite connexion.
	database.Connect()

	// Create a new Discord session using the provided bot token.
	dg, err := discordgo.New("Bot " + Token)
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}

	// Register the messageCreate func as a callback for MessageCreate events.
	dg.AddHandler(discord.CommandsHandler)

	// In this example, we only care about receiving message events.
	dg.Identify.Intents = discordgo.MakeIntent(discordgo.IntentsGuildMessages)

	// Open a websocket connection to Discord and begin listening.
	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Cleanly close down the Discord session.
	dg.Close()
}

// This function will be called (due to AddHandler above) every time a new
// message is created on any channel that the authenticated bot has access to.
func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

	fmt.Println(m.Content)
	fmt.Println(m.Mentions)
	fmt.Println(m.Author, m.Author.ID)

	// Ignore all messages created by the bot itself
	// This isn't required in this specific example but it's a good practice.
	if m.Author.ID == s.State.User.ID {
		return
	}
	// If the message is "ping" reply with "Pong!"
	if m.Content == "!pokedex bulbasaur" {
		msg := &discordgo.MessageEmbed{
			Author: &discordgo.MessageEmbedAuthor{
				Name:    "Pokédex",
				IconURL: "https://icon-library.com/images/pokedex-icon/pokedex-icon-20.jpg",
			},
			Thumbnail: &discordgo.MessageEmbedThumbnail{
				URL: "https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/1.png",
			},
			URL:         "https://www.pokemon.com/us/pokedex/bulbasaur",
			Title:       "Bulbasaur #1",
			Color:       15158332,
			Description: "A strange seed was planted on its back at birth. The plant sprouts and grows with this Pokémon.",
			Fields: []*discordgo.MessageEmbedField{
				{
					Name:   "Type",
					Value:  "grass, poison",
					Inline: true,
				},
				{
					Name:   "Height",
					Value:  "7",
					Inline: false,
				},
				{
					Name:  "Weight",
					Value: "69",
				},
			},
		}
		s.ChannelMessageSendEmbed(m.ChannelID, msg)
	}

	// If the message is "pong" reply with "Ping!"
	if m.Content == "pong" {
		s.ChannelMessageSend(m.ChannelID, "Ping!")
	}
}
