package mentions

import (
	"fmt"
	"github.com/rmwiesenberg/ode/tools"
	"log"
	"strings"

	"github.com/bwmarrin/discordgo"
)

// MentionHandler is the primary dispatcher for @ode mentions.
func MentionHandler(discord *discordgo.Session, message *discordgo.MessageCreate) {
	splitMessage := strings.Split(message.Content, " ")
	verbed := tools.Verber(strings.Join(splitMessage[1:], " "))

	msg := "wat"
	if verbed != "" {
		msg = fmt.Sprintf("_%s_", verbed)
	}

	_, err := discord.ChannelMessageSend(message.ChannelID, msg)
	if err != nil {
		log.Panicf("Cannot send msg %s", msg)
	}
}
