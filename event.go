package main

import (
	"misha/cmd"

	"github.com/bwmarrin/discordgo"
)

func Event(c cmd.Cmd) map[string]func(s *discordgo.Session, e *discordgo.Event) {
	return map[string]func(s *discordgo.Session, i *discordgo.Event){}
}
func EventHandler(s *discordgo.Session, e *discordgo.Event) {

}
