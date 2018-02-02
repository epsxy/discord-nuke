package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/briandowns/spinner"
	"github.com/bwmarrin/discordgo"
	"github.com/jessevdk/go-flags"
)

var (
	session                        *discordgo.Session
	limitChannelMessages           = 100
	limitChannelMessagesBulkDelete = 100

	opts struct {
		Email string `short:"e" long:"email" description:"Discord account email"`
		Pass  string `short:"p" long:"pass" description:"Discord account password"`
	}
)

func init() {
	_, err := flags.Parse(&opts)
	if err != nil {
		log.Fatalln(err)
	}

	if opts.Email != "" && opts.Pass != "" {
		session, err = discordgo.New(opts.Email, opts.Pass)
		if err != nil {
			log.Fatalln(err)
		}
	} else {
		log.Fatalln("error: email and password must be provided.")
	}
}

func main() {
	s := spinner.New([]string{"|", "/", "-", "\\"}, time.Millisecond*100)

	fmt.Printf("starting discord-nuke...")

	err := session.Open()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	defer session.Close()

	fmt.Println("session open.")

	fmt.Printf("discord-nuke started for %s.\n", session.State.User.String())

	var guilds []*discordgo.Guild
	var channels []*discordgo.Channel
	var messages []*discordgo.Message

	fmt.Println("getting guilds...")
	s.Start()

	guilds = session.State.Guilds

	s.Stop()
	fmt.Println("total guilds:", len(guilds))

	fmt.Println("getting channels...")
	s.Start()

	for _, guild := range guilds {
		c, err := session.GuildChannels(guild.ID)
		if err != nil {
			log.Printf("error: could not get channels from guild '%s' (id: %s): %s\n", guild.Name, guild.ID, err.Error())
			continue
		}

		channels = append(channels, c...)
	}

	s.Stop()
	fmt.Println("total channels:", len(channels))

loopChannels:
	for _, channel := range channels {

		fmt.Printf("getting messages for channel '%s'...", channel.Name)
		s.Start()

		if channel.Type != discordgo.ChannelTypeGuildText {
			fmt.Println("skipping (not a text channel)")
			continue
		}

		before := ""
		for {
			ms, err := session.ChannelMessages(channel.ID, limitChannelMessages, before, "", "")
			if err != nil {
				log.Printf("error: could not get messages from channel '%s' (id: %s): %s\n", channel.Name, channel.ID, err.Error())
				continue loopChannels
			}

			if len(ms) < 1 {
				break
			}

			before = ms[len(ms)-1].ID

			for _, m := range ms {
				if m.Author.ID == session.State.User.ID {
					// message is from self
					messages = append(messages, m)
				}
			}
		}

		s.Stop()
		fmt.Println("done.")
	}

	msgTotal := len(messages)
	msgDeleted := 0

	fmt.Println("total messages:", msgTotal)

	if msgTotal < 1 {
		fmt.Println("there are no messages to delete.")
		os.Exit(0)
	}

	for i, message := range messages {
		fmt.Printf("deleting message %d...", i)
		err := session.ChannelMessageDelete(message.ChannelID, message.ID)
		if err != nil {
			fmt.Printf("error: could not delete message (id:%s): %s\n", message.ID, err.Error())
		} else {
			fmt.Println("done.")
			msgDeleted++
		}
	}

	msgPercentDeleted := percent(float64(msgDeleted), float64(msgTotal))

	fmt.Printf("total messages: %d, deleted messages: %d\n", msgTotal, msgDeleted)

	if equal(msgPercentDeleted, 100.0) {
		fmt.Println("all messages have been deleted.")
	} else {
		fmt.Printf("%s of messages deleted.\n", formatFloatAsPercent(msgPercentDeleted))
	}
}
