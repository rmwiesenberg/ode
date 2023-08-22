package commands

import (
	"github.com/bwmarrin/discordgo"
)

var roundsMinValue = 1.0
var roundsMaxValue = 10.0
var timeMinValue = 1.0
var timeMaxValue = 120.0

var (
	Commands = []*discordgo.ApplicationCommand{
		{
			Name:        "cat",
			Description: "Get yo-self a cat",
		},
		{
			Name:        "dog",
			Description: "Get yo-self a dog",
		},
		{
			Name:        "scatter",
			Description: "Fly! You fools!",
		},
		{
			Name:        "moms-home",
			Description: "HIDE! MOM'S HOME!",
		},
		{
			Name:        "pomodoro",
			Description: "Control pomodoro sessions",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Name:        "start",
					Description: "Start Pomodoro in this Channel",
					Type:        discordgo.ApplicationCommandOptionSubCommand,
					Options: []*discordgo.ApplicationCommandOption{
						{
							Type:        discordgo.ApplicationCommandOptionInteger,
							Name:        "time-work",
							Description: "Minutes to work for",
							MinValue:    &timeMinValue,
							MaxValue:    timeMaxValue,
							Required:    true,
						},
						{
							Type:        discordgo.ApplicationCommandOptionInteger,
							Name:        "time-break",
							Description: "Minutes to break in-between",
							MinValue:    &timeMinValue,
							MaxValue:    timeMaxValue,
							Required:    true,
						},
						{
							Type:        discordgo.ApplicationCommandOptionInteger,
							Name:        "rounds",
							Description: "Number of rounds",
							MinValue:    &roundsMinValue,
							MaxValue:    roundsMaxValue,
							Required:    true,
						},
					},
				},
				{
					Name:        "end",
					Description: "End Active Pomodoro in this Channel",
					Type:        discordgo.ApplicationCommandOptionSubCommand,
				},
			},
		},
	}
	CommandHandlers = map[string]func(session *discordgo.Session, i *discordgo.InteractionCreate){
		"cat":       catsHandler,
		"dog":       dogsHandler,
		"scatter":   scatterHandler,
		"moms-home": momsHandler,
		"pomodoro": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			subcommand := i.ApplicationCommandData().Options[0]

			switch subcommand.Name {
			case "start":
				options := subcommand.Options

				optionMap := make(map[string]*discordgo.ApplicationCommandInteractionDataOption, len(options))
				for _, opt := range options {
					optionMap[opt.Name] = opt
				}

				pomodoro := Pomodoro{
					TimeWorkM:  optionMap["time-work"].IntValue(),
					TimeBreakM: optionMap["time-break"].IntValue(),
					Rounds:     optionMap["rounds"].IntValue(),
				}
				pomodoroStartHandler(s, i, pomodoro)
			case "end":
				pomodoroEndHandler(s, i)
			}
		},
	}
)
