package commands

import (
	"encoding/json"
	"errors"
	"fmt"
	"math/rand"
	"net/http"
	"os"

	"github.com/bwmarrin/discordgo"
)

type (
	// Cat is the data model for thecatapi return
	Cat struct {
		URL string `json:"URL"`
	}

	// Cats is a collection of Cat structs
	Cats []Cat
)

const CatBaseURL = "https://api.thecatapi.com/v1/images/search?api_key=%s"

func queryCat() (*string, error) {
	var cats Cats

	query := fmt.Sprintf(CatBaseURL, os.Getenv("CATS_API_KEY"))
	res, err := http.Get(query)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode == http.StatusOK {
		err := json.NewDecoder(res.Body).Decode(&cats)
		return &cats[0].URL, err
	}
	err = errors.New("Recieved: " + res.Status)
	return nil, err
}

// catsHandler accepts the cat request from discord
func catsHandler(session *discordgo.Session, i *discordgo.InteractionCreate) {
	cat, err := queryCat()
	if err != nil {
		panic(err)
	}

	session.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: *cat,
		},
	})
}

type (
	// Dog is the data model for dog.ceo return
	Dog struct {
		Message string `json:"Message"`
	}

	// Dogs is a collection of Dog structs
	Dogs []Dog
)

func queryDog() (*string, error) {
	var dog Dog

	baseURL := "https://dog.ceo/api/breeds/image/random"
	res, err := http.Get(baseURL)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode == http.StatusOK {
		err := json.NewDecoder(res.Body).Decode(&dog)
		return &dog.Message, err
	}
	err = errors.New("Recieved: " + res.Status)
	return nil, err
}

// dogsHandler accepts the dog request from discord
func dogsHandler(session *discordgo.Session, i *discordgo.InteractionCreate) {
	dog, err := queryDog()
	if err != nil {
		panic(err)
	}

	session.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: *dog,
		},
	})
}

func randomAnimal() (*string, error) {
	switch rand.Intn(2) {
	case 0:
		return queryCat()
	case 1:
		return queryDog()
	default:
		return nil, errors.New("Random out of range")
	}
}

func animalsHandler(session *discordgo.Session, i *discordgo.InteractionCreate) {
	animal, err := randomAnimal()
	if err != nil {
		panic(err)
	}

	session.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: *animal,
		},
	})
}

func sendRandomAnimal(session *discordgo.Session, channelID string) error {
	animal, err := randomAnimal()
	if animal != nil {
		session.ChannelMessageSend(channelID, *animal)
	}
	return err
}
