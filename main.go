package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/HETIC-MT-P2021/chen-discord-bot/command"
	"github.com/HETIC-MT-P2021/chen-discord-bot/discord"

	"github.com/HETIC-MT-P2021/chen-discord-bot/database"

	"github.com/bwmarrin/discordgo"
)

// Token : Variables used for discord line parameters
var Token string

func init() {
	flag.StringVar(&Token, "t", "", "Bot Token")
	flag.Parse()
}

func main() {

	// Create a new SQLite connexion.
	database.Connect()

	// Create a new Discord session using the provided bot token.
	session, err := discordgo.New("Bot " + Token)
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}

	// Create a new discord router
	router := discord.Create(&discord.Router{
		Prefixes: []string{"!poke"},
	})

	// Register a simple ping discord
	command.InitRouter(router)

	// Register the default help discord
	router.RegisterDefaultHelpCommand()

	// Initialize the router
	router.Initialize(session)

	// In this example, we only care about receiving message events.
	session.Identify.Intents = discordgo.MakeIntent(discordgo.IntentsGuildMessages)

	// Open a websocket connection to Discord and begin listening.
	err = session.Open()
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
	session.Close()
}
