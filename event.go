package main

import (
	"misha/events"
	"misha/extensions"

	"github.com/bwmarrin/discordgo"
)

func Events() map[string]func(c *extensions.Ex, s *discordgo.Session, e *discordgo.Event) {
	return map[string]func(c *extensions.Ex, s *discordgo.Session, e *discordgo.Event){
		"GUILD_CREATE": events.GuildCreate,
	}
}
func EventHandler(s *discordgo.Session, e *discordgo.Event) {

	if h, ok := Events()[e.Type]; ok {
		h(Ex, s, e)
	}
}
