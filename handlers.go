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
		{
			Name:        "play",
			Description: "Plays a song",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Type:        discordgo.ApplicationCommandOptionString,
					Name:        "identifier",
					Description: "The song link or search query",
					Required:    true,
				},
			},
		},
		{
			Name:        "pause",
			Description: "Pauses the current song",
		},
		{
			Name:        "now-playing",
			Description: "Shows the current playing song",
		},
		{
			Name:        "stop",
			Description: "Stops the current song and stops the player",
		},
		{
			Name:        "players",
			Description: "Shows all active players",
		},
		{
			Name:        "shuffle",
			Description: "Shuffles the current queue",
		},
		{
			Name:        "queue",
			Description: "Shows the current queue",
		},
		{
			Name:        "clear-queue",
			Description: "Clears the current queue",
		},
		{
			Name:        "queue-type",
			Description: "Sets the queue type",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Type:        discordgo.ApplicationCommandOptionString,
					Name:        "type",
					Description: "The queue type",
					Required:    true,
					Choices: []*discordgo.ApplicationCommandOptionChoice{
						{
							Name:  "default",
							Value: "default",
						},
						{
							Name:  "repeat-track",
							Value: "repeat-track",
						},
						{
							Name:  "repeat-queue",
							Value: "repeat-queue",
						},
					},
				},
			},
		},
	}
)

func ComponentsHandlers_init(c cmd.Cmd) map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate) {
	return map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
		"select": c.HelpComponent,
	}
}
func CommandsHandlers_init(c cmd.Cmd) map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate) {
	return map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
		"help":        c.Help,
		"choose":      c.SetupChoose,
		"play":        c.Play,
		"pause":       c.Pause,
		"now-playing": c.NowPlaying,
		"stop":        c.Stop,
		"queue":       c.Queue,
		"clear-queue": c.ClearQueue,
		"queue-type":  c.QueueType,
		"shuffle":     c.Shuffle,
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
