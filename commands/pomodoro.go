package commands

import (
	"github.com/bwmarrin/discordgo"
	"time"
)

type Pomodoro struct {
	channel                       string
	TimeWorkM, TimeBreakM, Rounds int64
	currentTimer                  *time.Timer
	workCounter                   int64
}

type pomodoroMap = map[string]Pomodoro

var pomodoros pomodoroMap

func _ensureMade() {
	if pomodoros == nil {
		pomodoros = make(pomodoroMap)
	}
}

func _clearPomodoro(pomodoro *Pomodoro) {
	pomodoro.currentTimer.Stop()
	delete(pomodoros, pomodoro.channel)
}

func pomodoroStartHandler(session *discordgo.Session, i *discordgo.InteractionCreate, pomodoro Pomodoro) {
	_ensureMade()

	pomodoro.channel = i.ChannelID

	if _, ok := pomodoros[pomodoro.channel]; ok {
		session.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "A pomodoro is already active in this channel!",
			},
		})
		return
	}

	var onWorkEnd, onBreakEnd func()

	onWorkEnd = func() {
		pomodoro.workCounter++
		if pomodoro.workCounter >= pomodoro.Rounds {
			session.ChannelMessageSend(pomodoro.channel, "✨ Pomodoro Finished ✨")
			_clearPomodoro(&pomodoro)
			return
		} else {
			if cat, _ := queryCat(); cat != nil {
				session.ChannelMessageSend(pomodoro.channel, cat.URL)
			}
			session.ChannelMessageSend(pomodoro.channel, "Break Time!")
		}
		pomodoro.currentTimer = time.AfterFunc(time.Duration(pomodoro.TimeBreakM)*time.Minute, onBreakEnd)
	}

	onBreakEnd = func() {
		_, err := session.ChannelMessageSend(pomodoro.channel, "🍅 Work Time 🍅")
		if err != nil {
			panic(err)
		}
		pomodoro.currentTimer = time.AfterFunc(time.Duration(pomodoro.TimeWorkM)*time.Minute, onWorkEnd)
	}

	pomodoro.currentTimer = time.AfterFunc(time.Duration(pomodoro.TimeWorkM)*time.Minute, onWorkEnd)

	session.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "🍅",
		},
	})

	pomodoros[i.ChannelID] = pomodoro
}

func pomodoroEndHandler(session *discordgo.Session, i *discordgo.InteractionCreate) {
	_ensureMade()

	if pomodoro, ok := pomodoros[i.ChannelID]; ok {
		_clearPomodoro(&pomodoro)
		session.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "Pomodoro cleared!",
			},
		})
	} else {
		session.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "No current pomodoro!",
			},
		})
	}
}
