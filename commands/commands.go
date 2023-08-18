package commands

import (
	"github.com/bwmarrin/discordgo"
)

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
	}
	CommandHandlers = map[string]func(session *discordgo.Session, i *discordgo.InteractionCreate){
		"cat":       catsHandler,
		"dog":       dogsHandler,
		"scatter":   scatterHandler,
		"moms-home": momsHandler,
	}
)
