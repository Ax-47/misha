package cmd

import (
	"misha/setup"

	"github.com/bwmarrin/discordgo"
)

func (c *Cmd) Help(s *discordgo.Session, i *discordgo.InteractionCreate) {
	setup.Help(c.Ex, s, i)
}
func (c *Cmd) SetupChoose(s *discordgo.Session, i *discordgo.InteractionCreate) {
	setup.SetupChoose(c.Ex, s, i)
}
func (c *Cmd) HelpComponent(s *discordgo.Session, i *discordgo.InteractionCreate) {
	setup.Help_Component(c.Ex, s, i)
}
