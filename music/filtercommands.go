package music

import (
	"context"
	"misha/extensions"

	"github.com/bwmarrin/discordgo"
	"github.com/disgoorg/disgolink/v2/lavalink"
	"github.com/disgoorg/snowflake/v2"
)

func Bassbost(c *extensions.Ex, s *discordgo.Session, i *discordgo.InteractionCreate) error {
	player := c.Bot.Lavalink.ExistingPlayer(snowflake.MustParse(i.GuildID))
	identifier := i.ApplicationCommandData().Options[0].StringValue()

	if player == nil {
		return s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Embeds: []*discordgo.MessageEmbed{embedPlayerNotFound(c.Lang(i.Locale.String()))},
			},
		})
	}
	var eq lavalink.Equalizer

	switch identifier {
	case "radio":
		eq[0] = 0.65
		eq[1] = 0.45
		eq[2] = -0.45
		eq[3] = -0.65
		eq[4] = -0.35
		eq[5] = 0.45
		eq[6] = 0.55
		eq[7] = 0.6
		eq[8] = 0.6
		eq[9] = 0.6
		eq[10] = 0
		eq[11] = 0
		eq[12] = 0
		eq[13] = 0
		eq[14] = 0
	case "electronic":
		eq[0] = 0.375
		eq[1] = 0.350
		eq[2] = 0.125
		eq[3] = 0
		eq[4] = 0
		eq[5] = -0.125
		eq[6] = -0.125
		eq[7] = 0
		eq[8] = 0.25
		eq[9] = 0.125
		eq[10] = 0.15
		eq[11] = 0.2
		eq[12] = 0.250
		eq[13] = 0.350
		eq[14] = 0.400

	case "gaming":
		eq[0] = 0.350
		eq[1] = 0.300
		eq[2] = 0.250
		eq[3] = 0.200
		eq[4] = 0.150
		eq[5] = 0.100
		eq[6] = 0.050
		eq[7] = 0.0
		eq[8] = -0.050
		eq[9] = -0.100
		eq[10] = -0.150
		eq[11] = -0.200
		eq[12] = -0.250
		eq[13] = -0.300
		eq[14] = -0.350

	case "classical":
		eq[0] = 0.375
		eq[1] = 0.350
		eq[2] = 0.125
		eq[3] = 0
		eq[4] = 0
		eq[5] = 0.125
		eq[6] = 0.550
		eq[7] = 0.050
		eq[8] = 0.125
		eq[9] = 0.250
		eq[10] = 0.200
		eq[11] = 0.250
		eq[12] = 0.300
		eq[13] = 0.250
		eq[14] = 0.300
	case "pop":
		eq[0] = -0.25
		eq[1] = 0.48
		eq[2] = 0.59
		eq[3] = 0.72
		eq[4] = 0.56
		eq[5] = 0.15
		eq[6] = -0.24
		eq[7] = -0.24
		eq[8] = -0.16
		eq[9] = -0.16
		eq[10] = 0
		eq[11] = 0
		eq[12] = 0
		eq[13] = 0
		eq[14] = 0

	case "rock":
		eq[0] = 0.300
		eq[1] = 0.250
		eq[2] = 0.200
		eq[3] = 0.100
		eq[4] = -0.050
		eq[5] = -0.050
		eq[6] = -0.150
		eq[7] = -0.200
		eq[8] = -0.100
		eq[9] = -0.050
		eq[10] = 0.050
		eq[11] = 0.100
		eq[12] = 0.200
		eq[13] = 0.250
		eq[14] = 0.300

	case "bassboost":
		eq[0] = 0.6
		eq[1] = 0.67
		eq[2] = 0.67
		eq[3] = 0
		eq[4] = -0.5
		eq[5] = 0.15
		eq[6] = -0.45
		eq[7] = 0.23
		eq[8] = 0.35
		eq[9] = 0.45
		eq[10] = 0.55
		eq[11] = 0.6
		eq[12] = 0.55
		eq[13] = 0
	case "bass":
		eq[0] = 0.6
		eq[1] = 0.7
		eq[2] = 0.8
		eq[3] = 0.55
		eq[4] = -0.25
		eq[5] = 0
		eq[6] = 0.25
		eq[7] = -0.45
		eq[8] = -0.55
		eq[9] = -0.7
		eq[10] = -0.3
		eq[11] = -0.25
		eq[12] = 0
		eq[13] = 0
		eq[14] = 0
	case "bassboosthigh":
		eq[0] = 0.1875
		eq[1] = 0.375
		eq[2] = -0.375
		eq[3] = -0.1875
		eq[4] = 0
		eq[5] = -0.0125
		eq[6] = -0.025
		eq[7] = -0.0175
		eq[8] = 0
		eq[9] = 0
		eq[10] = 0.0125
		eq[11] = 0.025
		eq[12] = 0.375
		eq[13] = 0.125
		eq[14] = 0.125
	case "highfull":
		eq[0] = 0.25 + 0.375
		eq[1] = 0.25 + 0.025
		eq[2] = 0.25 + 0.0125
		eq[3] = 0.25
		eq[4] = 0.25
		eq[5] = 0.25 + -0.0125
		eq[6] = 0.25 + -0.025
		eq[7] = 0.25 + 0.0175
		eq[8] = 0.25
		eq[9] = 0.25
		eq[10] = 0.25 + 0.0125
		eq[11] = 0.25 + 0.025
		eq[12] = 0.25 + 0.375
		eq[13] = 0.25 + 0.125
		eq[14] = 0.25 + 0.125
	case "treblebass":
		eq[0] = 0.6
		eq[1] = 0.67
		eq[2] = 0.67
		eq[3] = 0
		eq[4] = -0.5
		eq[5] = 0.15
		eq[6] = -0.45
		eq[7] = 0.23
		eq[8] = 0.35
		eq[9] = 0.45
		eq[10] = 0.55
		eq[11] = 0.6
		eq[12] = 0.55
		eq[13] = 0
	case "wtf":
		eq[0] = 1
		eq[1] = 1
		eq[2] = 1
		eq[3] = 1
		eq[4] = 1
		eq[5] = 1
		eq[6] = 1
		eq[7] = 1
		eq[8] = 1
		eq[9] = 1
		eq[10] = 1
		eq[11] = 1
		eq[12] = 1
		eq[13] = 1
		eq[14] = 1
	case "clean":
	}
	if err := player.Update(context.TODO(), lavalink.WithFilters(lavalink.Filters{Equalizer: &eq})); err != nil {
		return s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Embeds: []*discordgo.MessageEmbed{embedError(err)},
			},
		})
	}
	return s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Embeds: []*discordgo.MessageEmbed{embedEqualizer(identifier)},
		},
	})
}
func Timescale(c *extensions.Ex, s *discordgo.Session, i *discordgo.InteractionCreate) error {
	player := c.Bot.Lavalink.ExistingPlayer(snowflake.MustParse(i.GuildID))

	if player == nil {
		return s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Embeds: []*discordgo.MessageEmbed{embedPlayerNotFound(c.Lang(i.Locale.String()))},
			},
		})
	}
	var t lavalink.Timescale
	t.Speed = i.ApplicationCommandData().Options[0].FloatValue()
	if t.Speed <= 0 {
		return s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Embeds: []*discordgo.MessageEmbed{embedTimescaleErrorInput(c.Lang(i.Locale.String()), "speed")},
			},
		})
	}
	t.Pitch = i.ApplicationCommandData().Options[1].FloatValue()
	if t.Pitch <= 0 {
		return s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Embeds: []*discordgo.MessageEmbed{embedTimescaleErrorInput(c.Lang(i.Locale.String()), "pitch")},
			},
		})
	}
	t.Rate = i.ApplicationCommandData().Options[2].FloatValue()
	if t.Rate <= 0 {
		return s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Embeds: []*discordgo.MessageEmbed{embedTimescaleErrorInput(c.Lang(i.Locale.String()), "rate")},
			},
		})
	}
	if err := player.Update(context.TODO(), lavalink.WithFilters(lavalink.Filters{})); err != nil {
		return s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Embeds: []*discordgo.MessageEmbed{embedError(err)},
			},
		})
	}
	return s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Embeds: []*discordgo.MessageEmbed{},
		},
	})
}
func Karaoke(c *extensions.Ex, s *discordgo.Session, i *discordgo.InteractionCreate) error {
	player := c.Bot.Lavalink.ExistingPlayer(snowflake.MustParse(i.GuildID))

	if player == nil {
		return s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Embeds: []*discordgo.MessageEmbed{embedPlayerNotFound(c.Lang(i.Locale.String()))},
			},
		})
	}
	k := &lavalink.Karaoke{}
	if player.Filters().Karaoke.Level != 1 ||
		player.Filters().Karaoke.MonoLevel != 1 ||
		player.Filters().Karaoke.FilterBand != 220.0 ||
		player.Filters().Karaoke.FilterWidth != 100.0 {
		k = &lavalink.Karaoke{Level: 1.0,
			MonoLevel:   1.0,
			FilterBand:  220.0,
			FilterWidth: 100.0,
		}
	}
	if err := player.Update(context.TODO(), lavalink.WithFilters(lavalink.Filters{Karaoke: k})); err != nil {
		return s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Embeds: []*discordgo.MessageEmbed{embedError(err)},
			},
		})
	}
	return s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Embeds: []*discordgo.MessageEmbed{},
		},
	})
}
func Tremolo(c *extensions.Ex, s *discordgo.Session, i *discordgo.InteractionCreate) error {
	player := c.Bot.Lavalink.ExistingPlayer(snowflake.MustParse(i.GuildID))
	if player == nil {
		return s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Embeds: []*discordgo.MessageEmbed{embedPlayerNotFound(c.Lang(i.Locale.String()))},
			},
		})
	}
	var t lavalink.Tremolo
	t.Frequency = float32(i.ApplicationCommandData().Options[0].FloatValue())
	if t.Frequency < 0 {
		return s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Embeds: []*discordgo.MessageEmbed{embedTimescaleErrorInput(c.Lang(i.Locale.String()), "Tremolo")},
			},
		})
	}
	t.Depth = float32(i.ApplicationCommandData().Options[1].FloatValue())
	if t.Depth < 0 || t.Depth >= 1 {
		return s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Embeds: []*discordgo.MessageEmbed{embedTimescaleErrorInput(c.Lang(i.Locale.String()), "Depth")},
			},
		})
	}
	if err := player.Update(context.TODO(), lavalink.WithFilters(lavalink.Filters{})); err != nil {
		return s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Embeds: []*discordgo.MessageEmbed{embedError(err)},
			},
		})
	}
	return s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Embeds: []*discordgo.MessageEmbed{},
		},
	})
}
func Audio8D(c *extensions.Ex, s *discordgo.Session, i *discordgo.InteractionCreate) error {
	player := c.Bot.Lavalink.ExistingPlayer(snowflake.MustParse(i.GuildID))

	if player == nil {
		return s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Embeds: []*discordgo.MessageEmbed{embedPlayerNotFound(c.Lang(i.Locale.String()))},
			},
		})
	}
	k := &lavalink.Rotation{}
	if player.Filters().Rotation.RotationHz != 2 {
		k.RotationHz = 2
	}
	if err := player.Update(context.TODO(), lavalink.WithFilters(lavalink.Filters{Rotation: k})); err != nil {
		return s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Embeds: []*discordgo.MessageEmbed{embedError(err)},
			},
		})
	}
	return s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Embeds: []*discordgo.MessageEmbed{},
		},
	})
}
func Smoothing(c *extensions.Ex, s *discordgo.Session, i *discordgo.InteractionCreate) error {
	player := c.Bot.Lavalink.ExistingPlayer(snowflake.MustParse(i.GuildID))

	if player == nil {
		return s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Embeds: []*discordgo.MessageEmbed{embedPlayerNotFound(c.Lang(i.Locale.String()))},
			},
		})
	}
	k := &lavalink.LowPass{}
	if player.Filters().LowPass.Smoothing != 20.0 {
		k.Smoothing = 20.0
	}
	if err := player.Update(context.TODO(), lavalink.WithFilters(lavalink.Filters{LowPass: k})); err != nil {
		return s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Embeds: []*discordgo.MessageEmbed{embedError(err)},
			},
		})
	}
	return s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Embeds: []*discordgo.MessageEmbed{},
		},
	})
}
func Nightcore(c *extensions.Ex, s *discordgo.Session, i *discordgo.InteractionCreate) error {
	player := c.Bot.Lavalink.ExistingPlayer(snowflake.MustParse(i.GuildID))

	if player == nil {
		return s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Embeds: []*discordgo.MessageEmbed{embedPlayerNotFound(c.Lang(i.Locale.String()))},
			},
		})
	}
	k := &lavalink.Timescale{}
	if player.Filters().Timescale.Speed != 1.3 ||
		player.Filters().Timescale.Pitch != 1.3 ||
		player.Filters().Timescale.Rate != 1.0 {
		k = &lavalink.Timescale{
			Speed: 1.3,
			Pitch: 1.3,
			Rate:  1.0,
		}
	}
	if err := player.Update(context.TODO(), lavalink.WithFilters(lavalink.Filters{Timescale: k})); err != nil {
		return s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Embeds: []*discordgo.MessageEmbed{embedError(err)},
			},
		})
	}
	return s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Embeds: []*discordgo.MessageEmbed{},
		},
	})
}
func LoveNightcore(c *extensions.Ex, s *discordgo.Session, i *discordgo.InteractionCreate) error {
	player := c.Bot.Lavalink.ExistingPlayer(snowflake.MustParse(i.GuildID))

	if player == nil {
		return s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Embeds: []*discordgo.MessageEmbed{embedPlayerNotFound(c.Lang(i.Locale.String()))},
			},
		})
	}
	k := &lavalink.Timescale{}
	if player.Filters().Timescale.Speed != 1.1 ||
		player.Filters().Timescale.Pitch != 1.2 ||
		player.Filters().Timescale.Rate != 1.0 {
		k = &lavalink.Timescale{
			Speed: 1.1,
			Pitch: 1.2,
			Rate:  1.0,
		}
	}
	if err := player.Update(context.TODO(), lavalink.WithFilters(lavalink.Filters{Timescale: k})); err != nil {
		return s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Embeds: []*discordgo.MessageEmbed{embedError(err)},
			},
		})
	}
	return s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Embeds: []*discordgo.MessageEmbed{},
		},
	})
}
func SuperFast(c *extensions.Ex, s *discordgo.Session, i *discordgo.InteractionCreate) error {
	player := c.Bot.Lavalink.ExistingPlayer(snowflake.MustParse(i.GuildID))

	if player == nil {
		return s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Embeds: []*discordgo.MessageEmbed{embedPlayerNotFound(c.Lang(i.Locale.String()))},
			},
		})
	}
	k := &lavalink.Timescale{}
	if player.Filters().Timescale.Speed != 1.5010 ||
		player.Filters().Timescale.Pitch != 1.2450 ||
		player.Filters().Timescale.Rate != 1.9210 {
		k = &lavalink.Timescale{
			Speed: 1.5010,
			Pitch: 1.2450,
			Rate:  1.9210,
		}
	}
	if err := player.Update(context.TODO(), lavalink.WithFilters(lavalink.Filters{Timescale: k})); err != nil {
		return s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Embeds: []*discordgo.MessageEmbed{embedError(err)},
			},
		})
	}
	return s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Embeds: []*discordgo.MessageEmbed{},
		},
	})
}
func ResetFilters(c *extensions.Ex, s *discordgo.Session, i *discordgo.InteractionCreate) error {
	player := c.Bot.Lavalink.ExistingPlayer(snowflake.MustParse(i.GuildID))

	if player == nil {
		return s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Embeds: []*discordgo.MessageEmbed{embedPlayerNotFound(c.Lang(i.Locale.String()))},
			},
		})
	}
	if err := player.Update(context.TODO(), lavalink.WithFilters(lavalink.Filters{})); err != nil {
		return s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Embeds: []*discordgo.MessageEmbed{embedError(err)},
			},
		})
	}
	return s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Embeds: []*discordgo.MessageEmbed{},
		},
	})
}
func Vaporewave(c *extensions.Ex, s *discordgo.Session, i *discordgo.InteractionCreate) error {
	player := c.Bot.Lavalink.ExistingPlayer(snowflake.MustParse(i.GuildID))

	if player == nil {
		return s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Embeds: []*discordgo.MessageEmbed{embedPlayerNotFound(c.Lang(i.Locale.String()))},
			},
		})
	}
	k := lavalink.Filters{}
	var eq lavalink.Equalizer
	eq[0] = 0.3
	eq[1] = 0.3
	if player.Filters().Timescale.Pitch != 0.5 ||
		player.Filters().Tremolo.Depth != 0.3 ||
		player.Filters().Tremolo.Frequency != 14 ||
		player.Filters().Equalizer[1] != 0.3 ||
		player.Filters().Equalizer[0] != 0.3 {
		k = lavalink.Filters{
			Equalizer: &eq,
			Timescale: &lavalink.Timescale{Pitch: 0.5},
			Tremolo:   &lavalink.Tremolo{Depth: 0.3, Frequency: 14},
		}
	}
	if err := player.Update(context.TODO(), lavalink.WithFilters(k)); err != nil {
		return s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Embeds: []*discordgo.MessageEmbed{embedError(err)},
			},
		})
	}
	return s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Embeds: []*discordgo.MessageEmbed{},
		},
	})
}
