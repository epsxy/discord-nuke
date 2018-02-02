package main

import (
	"fmt"
	"math"

	"github.com/bwmarrin/discordgo"
)

func percent(a, b float64) float64 {
	return a / b * 100
}

func formatFloatAsPercent(a float64) string {
	return fmt.Sprintf("%.2f%%", a)
}

func equal(a, b float64) bool {
	return math.Nextafter(a, b) == b
}

// these functions are used to quickly get structs from their IDs
// i'm only using these because i'm a terrible lazy programmer

func guild(guildID string) *discordgo.Guild {
	g, _ := session.Guild(guildID)
	return g
}

func channel(channelID string) *discordgo.Channel {
	c, _ := session.Channel(channelID)
	return c
}

func user(userID string) *discordgo.User {
	u, _ := session.User(userID)
	return u
}
