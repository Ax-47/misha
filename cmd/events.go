package cmd

import (
	"misha/events"

	"github.com/bwmarrin/discordgo"
)

func (c *Cmd) SetupEventGuildCreate(s *discordgo.Session, e *discordgo.Event) {
	events.GuildCreate(c.Ex, s, e)
}
func (c *Cmd) SetupEventGuildDelete(s *discordgo.Session, e *discordgo.Event) {
	events.GuildDelete(s, e)
}
