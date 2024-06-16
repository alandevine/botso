package bot

import (
	"errors"
	"fmt"
	"log"
	"os"
	"os/signal"
	"sort"
	"strings"

	"adevine/botso/cfg"
	"adevine/botso/footy"

	"github.com/bwmarrin/discordgo"
	"gopkg.in/yaml.v2"
)

var (
	BotToken   string
	FootyToken string
)

type Msg struct {
	Author               string
	AuthorProfilePicture string
	MessageContent       string
}

func formatFootyMessage(currentWinner string, rankings map[string]int) string {
	// Extract keys into a slice
	keys := make([]string, 0, len(rankings))
	for key := range rankings {
		keys = append(keys, key)
	}

	sort.Slice(keys, func(i, j int) bool {
		return rankings[keys[i]] > rankings[keys[j]]
	})

	s := fmt.Sprintf("Current lead: %s\n", currentWinner)
	s += "Current rankings:\n"

	for i, key := range keys {
		s += fmt.Sprintf("%d. %s %d\n", i+1, key, rankings[key])
	}

	return s
}

func msgHandler(discord *discordgo.Session, message *discordgo.MessageCreate) {
	if message.Author.ID == discord.State.User.ID {
		return
	}

	if !strings.HasPrefix(message.Content, ">") {
		return
	}

	if strings.Contains(message.Content, "footy") {
		currentWinner, rankings := footy.DoRankings()
		discord.ChannelMessageSend(message.ChannelID, formatFootyMessage(currentWinner, rankings))
	}
}

func Run() error {
	footyApiKey, ok := os.LookupEnv("FOOTY_TOKEN")

	if !ok {
		log.Fatalln("Failed to load environment varaible `FOOTY_TOKEN`")
	}

	f, err := os.Open("config.yaml")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	var cfg cfg.Config
	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&cfg)
	if err != nil {
		panic(err)
	}

	footy.FootyToken = footyApiKey
	footy.Participents = cfg
	if BotToken == "" {
		return errors.New("BotToken not set")
	}

	log.Println("botso is leaving the little car")

	discord, err := discordgo.New("Bot " + BotToken)
	if err != nil {
		log.Fatalln(err)
	}

	discord.Open()
	defer discord.Close()

	discord.AddHandler(msgHandler)

	log.Println("it's circus time baby")

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c

	log.Println("pack it in boys")

	return nil
}
