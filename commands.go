package main

import (
	"log/slog"

	"github.com/disgoorg/disgo/bot"
	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/disgolink/v3/lavalink"
	"github.com/disgoorg/json"
	"github.com/disgoorg/snowflake/v2"
)

var commands = []discord.ApplicationCommandCreate{
	discord.SlashCommandCreate{
		Name:        "play",
		Description: "Plays a song",
		Options: []discord.ApplicationCommandOption{
			discord.ApplicationCommandOptionString{
				Name:        "identifier",
				Description: "The song link or search query",
				Required:    true,
			},
			discord.ApplicationCommandOptionString{
				Name:        "source",
				Description: "The source to search on",
				Required:    false,
				Choices: []discord.ApplicationCommandOptionChoiceString{
					{
						Name:  "YouTube",
						Value: string(lavalink.SearchTypeYouTube),
					},
					{
						Name:  "YouTube Music",
						Value: string(lavalink.SearchTypeYouTubeMusic),
					},
					{
						Name:  "SoundCloud",
						Value: string(lavalink.SearchTypeSoundCloud),
					},
					{
						Name:  "Deezer",
						Value: "dzsearch",
					},
					{
						Name:  "Deezer ISRC",
						Value: "dzisrc",
					},
					{
						Name:  "Spotify",
						Value: "spsearch",
					},
					{
						Name:  "AppleMusic",
						Value: "amsearch",
					},
				},
			},
		},
	},
	discord.SlashCommandCreate{
		Name:        "pause",
		Description: "Pauses the current song",
	},
	discord.SlashCommandCreate{
		Name:        "now-playing",
		Description: "Shows the current playing song",
	},
	discord.SlashCommandCreate{
		Name:        "stop",
		Description: "Stops the current song and stops the player",
	},
	discord.SlashCommandCreate{
		Name:        "disconnect",
		Description: "Disconnects the player",
	},
	discord.SlashCommandCreate{
		Name:        "bass-boost",
		Description: "Enables or disables bass boost",
		Options: []discord.ApplicationCommandOption{
			discord.ApplicationCommandOptionBool{
				Name:        "enabled",
				Description: "Whether bass boost should be enabled or disabled",
				Required:    true,
			},
		},
	},
	discord.SlashCommandCreate{
		Name:        "players",
		Description: "Shows all active players",
	},
	discord.SlashCommandCreate{
		Name:        "skip",
		Description: "Skips the current song",
		Options: []discord.ApplicationCommandOption{
			discord.ApplicationCommandOptionInt{
				Name:        "amount",
				Description: "The amount of songs to skip",
				Required:    false,
			},
		},
	},
	discord.SlashCommandCreate{
		Name:        "volume",
		Description: "Sets the volume of the player",
		Options: []discord.ApplicationCommandOption{
			discord.ApplicationCommandOptionInt{
				Name:        "volume",
				Description: "The volume to set",
				Required:    true,
				MaxValue:    json.Ptr(1000),
				MinValue:    json.Ptr(0),
			},
		},
	},
	discord.SlashCommandCreate{
		Name:        "seek",
		Description: "Seeks to a specific position in the current song",
		Options: []discord.ApplicationCommandOption{
			discord.ApplicationCommandOptionInt{
				Name:        "position",
				Description: "The position to seek to",
				Required:    true,
			},
			discord.ApplicationCommandOptionInt{
				Name:        "unit",
				Description: "The unit of the position",
				Required:    false,
				Choices: []discord.ApplicationCommandOptionChoiceInt{
					{
						Name:  "Milliseconds",
						Value: int(lavalink.Millisecond),
					},
					{
						Name:  "Seconds",
						Value: int(lavalink.Second),
					},
					{
						Name:  "Minutes",
						Value: int(lavalink.Minute),
					},
					{
						Name:  "Hours",
						Value: int(lavalink.Hour),
					},
				},
			},
		},
	},
	discord.SlashCommandCreate{
		Name:        "shuffle",
		Description: "Shuffles the current queue",
	},
	discord.SlashCommandCreate{
		Name:        "queue",
		Description: "Shows the current queue",
	},
	discord.SlashCommandCreate{
		Name:        "clear-queue",
		Description: "clear queue",
	},
	discord.SlashCommandCreate{
		Name:        "loop",
		Description: "loop",
		Options: []discord.ApplicationCommandOption{
			discord.ApplicationCommandOptionString{
				Name:        "type",
				Description: "loop type",
				Required:    true,
				Choices: []discord.ApplicationCommandOptionChoiceString{
					{
						Name:  "Normal",
						Value: "normal",
					},
					{
						Name:  "Repeat Track",
						Value: "repeat_track",
					},
					{
						Name:  "Repeat Queue",
						Value: "repeat_queue",
					},
				},
			},
		},
	},
	discord.SlashCommandCreate{
		Name:        "autoplay",
		Description: "add the next song",
	},
	discord.SlashCommandCreate{
		Name:        "timescale",
		Description: "Plays a song",
		Options: []discord.ApplicationCommandOption{
			discord.ApplicationCommandOptionInt{
				Name:        "speed",
				Description: "speed",
				Required:    true,
			},
			discord.ApplicationCommandOptionInt{
				Name:        "pitch",
				Description: "pitch",
				Required:    true,
			},
			discord.ApplicationCommandOptionInt{
				Name:        "rate",
				Description: "rate",
				Required:    true,
			},
		},
	},
	discord.SlashCommandCreate{
		Name:        "tremolo",
		Description: "Plays a song",
		Options: []discord.ApplicationCommandOption{
			discord.ApplicationCommandOptionInt{
				Name:        "frequency",
				Description: "adress1",
				Required:    true,
			},
			discord.ApplicationCommandOptionInt{
				Name:        "depth",
				Description: "adress2",
				Required:    true,
			},
		},
	},
	discord.SlashCommandCreate{
		Name:        "filter",
		Description: "auto play",
		Options: []discord.ApplicationCommandOption{
			discord.ApplicationCommandOptionString{
				Name:        "filter",
				Description: "The queue type",
				Required:    true,
				Choices: []discord.ApplicationCommandOptionChoiceString{
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
						Name:  "Reset",
						Value: "reset",
					},
				},
			},
		},
	}, discord.SlashCommandCreate{
		Name:        "equalizer",
		Description: "equalizer",
		Options: []discord.ApplicationCommandOption{
			discord.ApplicationCommandOptionString{
				Name:        "type",
				Description: "The queue type",
				Required:    true,
				Choices: []discord.ApplicationCommandOptionChoiceString{
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
	discord.SlashCommandCreate{
		Name:        "help",
		Description: "help",
	},
}

func registerCommands(client bot.Client) {
	if _, err := client.Rest().SetGuildCommands(client.ApplicationID(), snowflake.GetEnv("GUILD_ID"), commands); err != nil {
		slog.Error("error while registering commands ", slog.Any("err", err))
	}
}
