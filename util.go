package main

import (
	"fmt"
	"math"

	dgo "github.com/bwmarrin/discordgo"
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
// DO NOT USE if you're not 100% certain that an ID is correct
// i'm only using these because i'm a terrible programmer

func guild(guildID string) *dgo.Guild {
	g, _ := session.Guild(guildID)
	return g
}

func channel(channelID string) *dgo.Channel {
	c, _ := session.Channel(channelID)
	return c
}

func user(userID string) *dgo.User {
	u, _ := session.User(userID)
	return u
}
