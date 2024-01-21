package bot

import (
	"0x47/misha/embed"
	"context"

	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/disgo/events"
	"github.com/disgoorg/disgolink/v3/lavalink"
)

var radio = &lavalink.Equalizer{
	0:  0.65,
	1:  0.45,
	2:  -0.45,
	3:  -0.65,
	4:  -0.35,
	5:  0.45,
	6:  0.55,
	7:  0.6,
	8:  0.6,
	9:  0.6,
	10: 0,
	11: 0,
	12: 0,
	13: 0,
	14: 0,
}
var electronic = &lavalink.Equalizer{
	0:  0.375,
	1:  0.350,
	2:  0.125,
	3:  0,
	4:  0,
	5:  -0.125,
	6:  -0.125,
	7:  0,
	8:  0.25,
	9:  0.125,
	10: 0.15,
	11: 0.2,
	12: 0.250,
	13: 0.350,
	14: 0.400,
}
var gaming = &lavalink.Equalizer{
	0:  0.350,
	1:  0.300,
	2:  0.250,
	3:  0.200,
	4:  0.150,
	5:  0.100,
	6:  0.050,
	7:  0.0,
	8:  -0.050,
	9:  -0.100,
	10: -0.150,
	11: -0.200,
	12: -0.250,
	13: -0.300,
	14: -0.350,
}
var classical = &lavalink.Equalizer{
	0:  0.375,
	1:  0.350,
	2:  0.125,
	3:  0,
	4:  0,
	5:  0.125,
	6:  0.550,
	7:  0.050,
	8:  0.125,
	9:  0.250,
	10: 0.200,
	11: 0.250,
	12: 0.300,
	13: 0.250,
	14: 0.300,
}
var pop = &lavalink.Equalizer{
	0:  -0.25,
	1:  0.48,
	2:  0.59,
	3:  0.72,
	4:  0.56,
	5:  0.15,
	6:  -0.24,
	7:  -0.24,
	8:  -0.16,
	9:  -0.16,
	10: 0,
	11: 0,
	12: 0,
	13: 0,
	14: 0,
}
var rock = &lavalink.Equalizer{
	0:  0.300,
	1:  0.250,
	2:  0.200,
	3:  0.100,
	4:  -0.050,
	5:  -0.050,
	6:  -0.150,
	7:  -0.200,
	8:  -0.100,
	9:  -0.050,
	10: 0.050,
	11: 0.100,
	12: 0.200,
	13: 0.250,
	14: 0.300,
}
var bassBoost = &lavalink.Equalizer{
	0:  0.6,
	1:  0.67,
	2:  0.67,
	3:  0,
	4:  -0.5,
	5:  0.15,
	6:  -0.45,
	7:  0.23,
	8:  0.35,
	9:  0.45,
	10: 0.55,
	11: 0.6,
	12: 0.200,
	13: 0.55,
	14: 0,
}

var bass = &lavalink.Equalizer{
	0:  0.6,
	1:  0.7,
	2:  0.8,
	3:  0.55,
	4:  -0.25,
	5:  0,
	6:  0.25,
	7:  -0.45,
	8:  -0.55,
	9:  -0.7,
	10: -0.3,
	11: -0.25,
	12: 0,
	13: 0,
	14: 0,
}
var bassboosthigh = &lavalink.Equalizer{
	0:  0.1875,
	1:  0.375,
	2:  -0.375,
	3:  -0.1875,
	4:  0,
	5:  -0.0125,
	6:  -0.025,
	7:  -0.0175,
	8:  0,
	9:  0,
	10: 0.0125,
	11: 0.025,
	12: 0.375,
	13: 0.125,
	14: 0.125,
}
var highfull = &lavalink.Equalizer{
	0:  0.25 + 0.375,
	1:  0.25 + 0.025,
	2:  0.25 + 0.0125,
	3:  0.25,
	4:  0.25,
	5:  0.25 + -0.0125,
	6:  0.25 + -0.025,
	7:  0.25 + 0.0175,
	8:  0.25,
	9:  0.25,
	10: 0.25 + 0.0125,
	11: 0.25 + 0.025,
	12: 0.25 + 0.375,
	13: 0.25 + 0.125,
	14: 0.25 + 0.125,
}
var treblebass = &lavalink.Equalizer{
	0:  0.6,
	1:  0.67,
	2:  0.67,
	3:  0,
	4:  -0.5,
	5:  0.15,
	6:  -0.45,
	7:  0.23,
	8:  0.35,
	9:  0.45,
	10: 0.55,
	11: 0.6,
	12: 0.55,
	13: 0,
}

func (b *Bot) BassBoost(event *events.ApplicationCommandInteractionCreate, data discord.SlashCommandInteractionData) error {

	bassboost, ok := data.OptString("filter")
	if !ok {
		return event.CreateMessage(discord.MessageCreate{
			Embeds: []discord.Embed{
				embed.ErrorNotFoundPlayer(),
			},
		})
	}
	player := b.Lavalink.ExistingPlayer(*event.GuildID())
	if player == nil {
		return event.CreateMessage(discord.MessageCreate{
			Embeds: []discord.Embed{
				embed.ErrorNotFoundPlayer(),
			},
		})
	}
	var eq lavalink.Equalizer
	filter := player.Filters()
	switch bassboost {
	case "radio":
		eq = *radio
	case "electronic":
		eq = *electronic

	case "gaming":
		eq = *gaming

	case "classical":
		eq = *classical
	case "pop":
		eq = *pop

	case "rock":
		eq = *rock
	case "bassboost":
		eq = *bassBoost
	case "bass":
		eq = *bass
	case "bassboosthigh":
		eq = *bassboosthigh
	case "highfull":
		eq = *highfull
	case "treblebass":
		eq = *treblebass
	case "wtf":
		for i := 0; i < 15; i++ {
			eq[i] = 1
		}
	case "clean":
		for i := 0; i < 15; i++ {
			eq[i] = 0
		}
	}
	filter.Equalizer = &eq
	if err := player.Update(context.TODO(), lavalink.WithFilters(filter)); err != nil {
		return event.CreateMessage(discord.MessageCreate{
			Embeds: []discord.Embed{
				embed.ErrorOther(),
			},
		})
	}

	return event.CreateMessage(discord.MessageCreate{
		Embeds: []discord.Embed{
			embed.Eq(bassboost),
		},
	})
}
func (b *Bot) Timescale(event *events.ApplicationCommandInteractionCreate, data discord.SlashCommandInteractionData) error {
	player := b.Lavalink.ExistingPlayer(*event.GuildID())
	if player == nil {
		return event.CreateMessage(discord.MessageCreate{
			Embeds: []discord.Embed{
				embed.ErrorNotFoundPlayer(),
			},
		})
	}
	speed, ok := data.OptInt("speed")
	if !ok {
		speed = 1
	}
	pitch, ok := data.OptInt("pitch")
	if !ok {
		pitch = 1
	}
	rate, ok := data.OptInt("rate")
	if !ok {
		rate = 1
	}
	filter := player.Filters()
	filter.Timescale = &lavalink.Timescale{}
	filter.Timescale.Speed = float64(speed) * 0.1
	filter.Timescale.Pitch = float64(pitch) * 0.1
	filter.Timescale.Rate = float64(rate) * 0.1
	if filter.Timescale.Speed <= 0 || filter.Timescale.Pitch <= 0 || filter.Timescale.Rate <= 0 {
		return event.CreateMessage(discord.MessageCreate{
			Embeds: []discord.Embed{
				embed.ErrorTimeScale(),
			},
		})
	}
	if err := player.Update(context.TODO(), lavalink.WithFilters(filter)); err != nil {
		return event.CreateMessage(discord.MessageCreate{
			Embeds: []discord.Embed{
				embed.ErrorOther(),
			},
		})
	}
	return event.CreateMessage(discord.MessageCreate{
		Embeds: []discord.Embed{
			embed.Timescale(speed, pitch, rate),
		},
	})
}
func (b *Bot) Tremolo(event *events.ApplicationCommandInteractionCreate, data discord.SlashCommandInteractionData) error {
	player := b.Lavalink.ExistingPlayer(*event.GuildID())
	if player == nil {
		return event.CreateMessage(discord.MessageCreate{
			Embeds: []discord.Embed{
				embed.ErrorNotFoundPlayer(),
			},
		})
	}

	frequency, ok := data.OptInt("frequency")
	if !ok {
		frequency = int(player.Filters().Tremolo.Frequency * 10)
	}
	depth, ok := data.OptInt("depth")
	if !ok {
		depth = int(player.Filters().Tremolo.Depth * 10)
	}
	var t lavalink.Tremolo
	filter := player.Filters()
	filter.Tremolo = &lavalink.Tremolo{}
	filter.Tremolo.Frequency = float32(frequency) * 0.1
	filter.Tremolo.Depth = float32(depth) * 0.1
	if t.Frequency < 0 {
		return event.CreateMessage(discord.MessageCreate{
			Embeds: []discord.Embed{
				embed.ErrorFrequency(),
			},
		})
	}

	if t.Depth < 0 || t.Depth >= 1 {
		return event.CreateMessage(discord.MessageCreate{
			Embeds: []discord.Embed{
				embed.ErrorDepth(),
			},
		})
	}
	if err := player.Update(context.TODO(), lavalink.WithFilters(filter)); err != nil {
		return event.CreateMessage(discord.MessageCreate{
			Embeds: []discord.Embed{
				embed.ErrorOther(),
			},
		})
	}
	return event.CreateMessage(discord.MessageCreate{
		Embeds: []discord.Embed{
			embed.Tremolo(frequency, depth),
		},
	})
}

func (b *Bot) Filter(event *events.ApplicationCommandInteractionCreate, data discord.SlashCommandInteractionData) error {
	player := b.Lavalink.ExistingPlayer(*event.GuildID())
	if player == nil {
		return event.CreateMessage(discord.MessageCreate{
			Embeds: []discord.Embed{
				embed.ErrorNotFoundPlayer(),
			},
		})
	}
	filter := player.Filters()

	f, ok := data.OptString("filter")
	if !ok {
		return event.CreateMessage(discord.MessageCreate{
			Embeds: []discord.Embed{
				embed.ErrorNotFoundPlayer(),
			},
		})
	}
	switch f {
	case "karaoke":
		filter.Karaoke = &lavalink.Karaoke{
			Level:       1.0,
			MonoLevel:   1.0,
			FilterBand:  220.0,
			FilterWidth: 100.0,
		}
	case "8d":
		filter.Rotation = &lavalink.Rotation{
			RotationHz: 1,
		}
	case "smoothing":
		filter.LowPass = &lavalink.LowPass{
			Smoothing: 20,
		}
	case "nightcore":
		filter.Timescale = &lavalink.Timescale{
			Speed: 1.3,
			Pitch: 1.3,
			Rate:  1.0,
		}
	case "lovenightcore":
		filter.Timescale = &lavalink.Timescale{
			Speed: 1.1,
			Pitch: 1.2,
			Rate:  1.0,
		}
	case "superfast":
		filter.Timescale = &lavalink.Timescale{
			Speed: 1.5010,
			Pitch: 1.2450,
			Rate:  1.9210,
		}
	case "reset":
		filter = lavalink.Filters{}
	}

	if err := player.Update(context.TODO(), lavalink.WithFilters(filter)); err != nil {
		return event.CreateMessage(discord.MessageCreate{
			Embeds: []discord.Embed{
				embed.ErrorOther(),
			},
		})
	}
	return event.CreateMessage(discord.MessageCreate{
		Embeds: []discord.Embed{
			embed.Filter(f),
		},
	})
}
