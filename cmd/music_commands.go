package cmd

import (
	"misha/music"

	"github.com/bwmarrin/discordgo"
)

func (c *Cmd) Play(s *discordgo.Session, i *discordgo.InteractionCreate) {
	music.Play(c.Ex, s, i, i.ApplicationCommandData())
}
func (c *Cmd) Pause(s *discordgo.Session, i *discordgo.InteractionCreate) {
	music.Pause(c.Ex, s, i, i.ApplicationCommandData())
}
func (c *Cmd) NowPlaying(s *discordgo.Session, i *discordgo.InteractionCreate) {
	music.NowPlaying(c.Ex, s, i, i.ApplicationCommandData())
}
func (c *Cmd) Stop(s *discordgo.Session, i *discordgo.InteractionCreate) {
	music.Stop(c.Ex, s, i, i.ApplicationCommandData())
}
func (c *Cmd) Queue(s *discordgo.Session, i *discordgo.InteractionCreate) {
	music.Queue(c.Ex, s, i, i.ApplicationCommandData())
}
func (c *Cmd) ClearQueue(s *discordgo.Session, i *discordgo.InteractionCreate) {
	music.ClearQueue(c.Ex, s, i, i.ApplicationCommandData())
}
func (c *Cmd) QueueType(s *discordgo.Session, i *discordgo.InteractionCreate) {
	music.QueueType(c.Ex, s, i, i.ApplicationCommandData())
}
func (c *Cmd) Shuffle(s *discordgo.Session, i *discordgo.InteractionCreate) {
	music.Shuffle(c.Ex, s, i, i.ApplicationCommandData())
}
func (c *Cmd) Skip(s *discordgo.Session, i *discordgo.InteractionCreate) {
	music.Skip(c.Ex, s, i, i.ApplicationCommandData())
}
func (c *Cmd) Seek(s *discordgo.Session, i *discordgo.InteractionCreate) {
	music.Seek(c.Ex, s, i, i.ApplicationCommandData())
}
func (c *Cmd) Remove(s *discordgo.Session, i *discordgo.InteractionCreate) {
	music.Remove(c.Ex, s, i, i.ApplicationCommandData())
}
func (c *Cmd) Move(s *discordgo.Session, i *discordgo.InteractionCreate) {
	music.Move(c.Ex, s, i, i.ApplicationCommandData())
}
func (c *Cmd) Autoplay(s *discordgo.Session, i *discordgo.InteractionCreate) {
	music.Autoplay(c.Ex, s, i, i.ApplicationCommandData())
}
