package main

import (
	"misha/extensions"
	"misha/music"
	"misha/setup"

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
			Name:        "volume",
			Description: "Plays a song",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Type:        discordgo.ApplicationCommandOptionInteger,
					Name:        "vol",
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
			Name:        "timescale",
			Description: "Plays a song",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Type:        discordgo.ApplicationCommandOptionInteger,
					Name:        "speed",
					Description: "adress1",
					Required:    true,
				},
				{
					Type:        discordgo.ApplicationCommandOptionInteger,
					Name:        "pitch",
					Description: "adress2",
					Required:    true,
				},
				{
					Type:        discordgo.ApplicationCommandOptionInteger,
					Name:        "rate",
					Description: "adress2",
					Required:    true,
				},
			},
		},
		{
			Name:        "tremolo",
			Description: "Plays a song",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Type:        discordgo.ApplicationCommandOptionInteger,
					Name:        "frequency",
					Description: "adress1",
					Required:    true,
				},
				{
					Type:        discordgo.ApplicationCommandOptionInteger,
					Name:        "depth",
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
			Name:        "filter",
			Description: "auto play",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Type:        discordgo.ApplicationCommandOptionString,
					Name:        "type",
					Description: "The queue type",
					Required:    true,
					Choices: []*discordgo.ApplicationCommandOptionChoice{
						{
							Name:  "Karaoke",
							Value: "karaoke",
						},
						{
							Name:  "8d",
							Value: "8d",
						},
						{
							Name:  "Smoothing",
							Value: "smoothing",
						},
						{
							Name:  "Nightcore",
							Value: "nightcore",
						},

						{
							Name:  "LoveNightcore",
							Value: "lovenightcore",
						},
						{
							Name:  "Superfast",
							Value: "superfast",
						},
						{
							Name:  "Vaporewave",
							Value: "vaporewave",
						},
						{
							Name:  "Reset",
							Value: "reset",
						},
					},
				},
			},
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
		}, {
			Name:        "equalizer",
			Description: "equalizer",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Type:        discordgo.ApplicationCommandOptionString,
					Name:        "type",
					Description: "The queue type",
					Required:    true,
					Choices: []*discordgo.ApplicationCommandOptionChoice{
						{
							Name:  "radio",
							Value: "radio",
						},
						{
							Name:  "electronic",
							Value: "electronic",
						},
						{
							Name:  "gaming",
							Value: "gaming",
						},
						{
							Name:  "classical",
							Value: "classical",
						},
						{
							Name:  "pop",
							Value: "pop",
						},
						{
							Name:  "rock",
							Value: "rock",
						}, {
							Name:  "bassboost",
							Value: "bassboost",
						}, {
							Name:  "bass",
							Value: "bass",
						}, {
							Name:  "bassboosthigh",
							Value: "bassboosthigh",
						}, {
							Name:  "highfull",
							Value: "highfull",
						}, {
							Name:  "treblebass",
							Value: "treblebass",
						}, {
							Name:  "clean",
							Value: "clean",
						}, {
							Name:  "wtf",
							Value: "wtf",
						},
					},
				},
			},
		},
	}
)

func ComponentsHandlers_init() map[string]func(c *extensions.Ex, s *discordgo.Session, i *discordgo.InteractionCreate) {
	return map[string]func(c *extensions.Ex, s *discordgo.Session, i *discordgo.InteractionCreate){
		"select": setup.HelpComponent,
		"re":     music.HandlerComponentsQueuePrevious,
		"next":   music.HandlerComponentsQueueNext,
	}
}
func CommandsHandlers_init() map[string]func(c *extensions.Ex, s *discordgo.Session, i *discordgo.InteractionCreate) {
	return map[string]func(c *extensions.Ex, s *discordgo.Session, i *discordgo.InteractionCreate){
		"help": setup.Help,
		//music commands
		"play":        music.Play,
		"pause":       music.Pause,
		"now-playing": music.NowPlaying,
		"stop":        music.Stop,
		"queue":       music.Queue,
		"clear-queue": music.ClearQueue,
		"loop":        music.QueueType,
		"shuffle":     music.Shuffle,
		"skip":        music.Skip,
		"seek":        music.Seek,
		"swap":        music.Swap,
		"remove":      music.Remove,
		"auto-play":   music.Autoplay,
		"equalizer":   music.Bassboost,
		"timescale":   music.Timescale,
		"tremolo":     music.Tremolo,
		"filter":      music.Filter,
		"volume":      music.Volume,
	}
}

func Handlers(s *discordgo.Session, i *discordgo.InteractionCreate) {
	switch i.Type {
	case discordgo.InteractionApplicationCommand:
		if h, ok := CommandsHandlers_init()[i.ApplicationCommandData().Name]; ok {
			h(Ex, s, i)
		}
	case discordgo.InteractionMessageComponent:

		if h, ok := ComponentsHandlers_init()[i.MessageComponentData().CustomID]; ok {
			h(Ex, s, i)
		}
	}
}
