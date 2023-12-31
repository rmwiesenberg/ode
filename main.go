package main

import (
	"flag"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/rmwiesenberg/ode/commands"
	"github.com/rmwiesenberg/ode/mentions"
	"log"
	"os"
	"os/signal"
	"strings"
)

// Bot parameters
var (
	BotID   = flag.String("id", "1140718447018377417", "BotID")
	GuildID = flag.String("guild", "", "Test guild ID. If not passed - bot registers tools globally")
)

func messageHandler(session *discordgo.Session, message *discordgo.MessageCreate) {
	if message.Author.ID == *BotID || message.Author.Bot {
		// Do nothing because the bot is talking
		return
	} else if strings.HasPrefix(message.Content, fmt.Sprintf("<@!%s> ", *BotID)) || strings.HasPrefix(message.Content, fmt.Sprintf("<@%s> ", *BotID)) {
		mentions.MentionHandler(session, message)
	}
}

func main() {
	flag.Parse()

	session, err := discordgo.New("Bot " + os.Getenv("DISCORD_BOT_TOKEN"))
	if err != nil {
		panic(err)
	}

	session.AddHandler(messageHandler)
	session.AddHandler(func(session *discordgo.Session, i *discordgo.InteractionCreate) {
		if handler, ok := commands.CommandHandlers[i.ApplicationCommandData().Name]; ok {
			handler(session, i)
		}
	})

	session.AddHandler(func(session *discordgo.Session, r *discordgo.Ready) {
		log.Printf("Ode has started on %d servers", len(session.State.Guilds))
	})
	err = session.Open()
	if err != nil {
		panic(err)
	}
	if err != nil {
		log.Fatalf("Cannot open the session: %v", err)
	}

	for _, v := range commands.Commands {
		_, err := session.ApplicationCommandCreate(session.State.User.ID, *GuildID, v)
		if err != nil {
			log.Panicf("Cannot create '%v' command: %v", v.Name, err)
		}
	}

	defer func(session *discordgo.Session) {
		err := session.Close()
		if err != nil {
			log.Fatalf("Cannot close the session: %v", err)
		}
	}(session)

	stop := make(chan os.Signal)
	signal.Notify(stop, os.Interrupt)
	<-stop
	log.Println("Gracefully shutdowning")
}
