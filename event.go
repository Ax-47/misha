package main

import (
	"misha/cmd"

	"github.com/bwmarrin/discordgo"
)

func Events(c cmd.Cmd) map[string]func(s *discordgo.Session, e *discordgo.Event) {
	return map[string]func(s *discordgo.Session, e *discordgo.Event){
		"GUILD_CREATE": c.SetupEventGuildCreate,
	}
}
func EventHandler(s *discordgo.Session, e *discordgo.Event) {

	if h, ok := Events(c)[e.Type]; ok {
		h(s, e)
	}
}
