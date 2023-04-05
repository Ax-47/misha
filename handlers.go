package main

import (
	"misha/cmd"

	"github.com/bwmarrin/discordgo"
)

var (
	Commands = []*discordgo.ApplicationCommand{
		{
			Name:        "help",
			Description: "Basic command",
			Type:        discordgo.ChatApplicationCommand,
		},
		{
			Name:        "choose",
			Description: "Basic command",
			Type:        discordgo.ChatApplicationCommand,
		},
	}
)

func ComponentsHandlers_init(c cmd.Cmd) map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate) {
	return map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
		"select": c.Help_Component,
	}
}
func CommandsHandlers_init(c cmd.Cmd) map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate) {
	return map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
		"help":   c.Help,
		"choose": c.SetupChoose,
	}
}
func CommandsHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	switch i.Type {
	case discordgo.InteractionApplicationCommand:
		if h, ok := CommandsHandlers_init(c)[i.ApplicationCommandData().Name]; ok {
			h(s, i)
		}
	case discordgo.InteractionMessageComponent:

		if h, ok := ComponentsHandlers_init(c)[i.MessageComponentData().CustomID]; ok {
			h(s, i)
		}
	}
}
