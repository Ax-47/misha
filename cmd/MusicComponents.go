package cmd

import (
	"misha/music"

	"github.com/bwmarrin/discordgo"
)

func (c *Cmd) HandlerComponentsQueuePrevious(s *discordgo.Session, i *discordgo.InteractionCreate) {
	music.HandlerComponentsQueuePrevious(c.Ex, s, i)
}
func (c *Cmd) HandlerComponentsQueueNext(s *discordgo.Session, i *discordgo.InteractionCreate) {
	music.HandlerComponentsQueueNext(c.Ex, s, i)
}
