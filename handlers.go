package main

import (
	"misha/cmd"
	"misha/languages"

	"github.com/bwmarrin/discordgo"
)

var (
	defaultMemberPermissions int64 = discordgo.PermissionManageServer
	lang                     map[string]languages.Lang
	Commands                 = []*discordgo.ApplicationCommand{
		{
			Name:        "help",
			Description: "Basic command",
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
		"help": c.Help,
	}
}
