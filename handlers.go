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
					Name:        "query",
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
			Name:        "skip",
			Description: "skip a song",
		},
		{
			Name:        "seek",
			Description: "Plays a song",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Type:        discordgo.ApplicationCommandOptionInteger,
					Name:        "sec",
					Description: "Seek",
					Required:    true,
				},
			},
		},
		{
			Name:        "remove",
			Description: "Plays a song",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Type:        discordgo.ApplicationCommandOptionInteger,
					Name:        "adress",
					Description: "Seek",
					Required:    true,
				},
			},
		},
		{
			Name:        "swap",
			Description: "Plays a song",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Type:        discordgo.ApplicationCommandOptionInteger,
					Name:        "from",
					Description: "adress1",
					Required:    true,
				},
				{
					Type:        discordgo.ApplicationCommandOptionInteger,
					Name:        "to",
					Description: "adress2",
					Required:    true,
				},
			},
		},
		{
			Name:        "auto-play",
			Description: "auto play",
		},
		{
			Name:        "loop",
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
							Value: "repeat_track",
						},
						{
							Name:  "repeat-queue",
							Value: "repeat_queue",
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
		"re":     c.HandlerComponentsQueuePrevious,
		"next":   c.HandlerComponentsQueueNext,
	}
}
func CommandsHandlers_init(c cmd.Cmd) map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate) {
	return map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
		"help":   c.Help,
		"choose": c.SetupChoose,

		//music commands
		"play":        c.Play,
		"pause":       c.Pause,
		"now-playing": c.NowPlaying,
		"stop":        c.Stop,
		"queue":       c.Queue,
		"clear-queue": c.ClearQueue,
		"loop":        c.QueueType,
		"shuffle":     c.Shuffle,
		"skip":        c.Skip,
		"seek":        c.Seek,
		"swap":        c.Swap,

		"remove":    c.Remove,
		"auto-play": c.Autoplay,
		//in the future
		//
		//filter
	}
}

func Handlers(s *discordgo.Session, i *discordgo.InteractionCreate) {
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
