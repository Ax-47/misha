package music

import (
	"context"
	"misha/extensions"

	"github.com/bwmarrin/discordgo"
	"github.com/disgoorg/disgolink/v2/lavalink"
	"github.com/disgoorg/snowflake/v2"
)

func Bassboost(c *extensions.Ex, s *discordgo.Session, i *discordgo.InteractionCreate) {
	player := c.Bot.Lavalink.ExistingPlayer(snowflake.MustParse(i.GuildID))
	identifier := i.ApplicationCommandData().Options[0].StringValue()

	if player == nil {
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Embeds: []*discordgo.MessageEmbed{embedPlayerNotFound(c.Lang(i.Locale.String()))},
			},
		})
		return
	}
	var eq lavalink.Equalizer
	filter := player.Filters()
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
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Embeds: []*discordgo.MessageEmbed{embedError(err)},
			},
		})
		return
	}
	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Embeds: []*discordgo.MessageEmbed{embedEqualizer(identifier)},
		},
	})
}
func Timescale(c *extensions.Ex, s *discordgo.Session, i *discordgo.InteractionCreate) {
	player := c.Bot.Lavalink.ExistingPlayer(snowflake.MustParse(i.GuildID))

	if player == nil {
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Embeds: []*discordgo.MessageEmbed{embedPlayerNotFound(c.Lang(i.Locale.String()))},
			},
		})
		return
	}

	filter := player.Filters()
	filter.Timescale = &lavalink.Timescale{}
	filter.Timescale.Speed = float64(i.ApplicationCommandData().Options[0].IntValue()) * 0.1
	if filter.Timescale.Speed <= 0 {
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Embeds: []*discordgo.MessageEmbed{embedTimescaleErrorInput(c.Lang(i.Locale.String()), "speed")},
			},
		})
		return
	}
	filter.Timescale.Pitch = float64(i.ApplicationCommandData().Options[1].IntValue()) * 0.1
	if filter.Timescale.Pitch <= 0 {
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Embeds: []*discordgo.MessageEmbed{embedTimescaleErrorInput(c.Lang(i.Locale.String()), "pitch")},
			},
		})
		return
	}
	filter.Timescale.Rate = float64(i.ApplicationCommandData().Options[2].IntValue()) * 0.1
	if filter.Timescale.Rate <= 0 {
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Embeds: []*discordgo.MessageEmbed{embedTimescaleErrorInput(c.Lang(i.Locale.String()), "rate")},
			},
		})
		return
	}
	if err := player.Update(context.TODO(), lavalink.WithFilters(filter)); err != nil {
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Embeds: []*discordgo.MessageEmbed{embedError(err)},
			},
		})
		return
	}
	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Embeds: []*discordgo.MessageEmbed{embedTimescale(filter.Timescale.Speed, filter.Timescale.Pitch, filter.Timescale.Rate)},
		},
	})
}
func Tremolo(c *extensions.Ex, s *discordgo.Session, i *discordgo.InteractionCreate) {
	player := c.Bot.Lavalink.ExistingPlayer(snowflake.MustParse(i.GuildID))
	if player == nil {
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Embeds: []*discordgo.MessageEmbed{embedPlayerNotFound(c.Lang(i.Locale.String()))},
			},
		})
		return
	}
	var t lavalink.Tremolo
	filter := player.Filters()
	filter.Tremolo = &lavalink.Tremolo{}
	filter.Tremolo.Frequency = float32(i.ApplicationCommandData().Options[0].IntValue()) * 0.1
	if t.Frequency < 0 {
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Embeds: []*discordgo.MessageEmbed{embedTimescaleErrorInput(c.Lang(i.Locale.String()), "Frequency")},
			},
		})
		return
	}
	filter.Tremolo.Depth = float32(i.ApplicationCommandData().Options[1].IntValue()) * 0.1
	if t.Depth < 0 || t.Depth >= 1 {
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Embeds: []*discordgo.MessageEmbed{embedTimescaleErrorInput(c.Lang(i.Locale.String()), "Depth")},
			},
		})
		return
	}
	if err := player.Update(context.TODO(), lavalink.WithFilters(filter)); err != nil {
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Embeds: []*discordgo.MessageEmbed{embedError(err)},
			},
		})
		return
	}
	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Embeds: []*discordgo.MessageEmbed{embedTremolo(filter.Tremolo.Frequency, filter.Tremolo.Depth)},
		},
	})
}
func Volume(c *extensions.Ex, s *discordgo.Session, i *discordgo.InteractionCreate) {

	player := c.Bot.Lavalink.ExistingPlayer(snowflake.MustParse(i.GuildID))
	if player == nil {
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Embeds: []*discordgo.MessageEmbed{embedPlayerNotFound(c.Lang(i.Locale.String()))},
			},
		})
		return
	}

	volume := int(i.ApplicationCommandData().Options[0].IntValue())
	if volume <= 0 || volume > 500 {
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Embeds: []*discordgo.MessageEmbed{},
			},
		})
		return
	}
	if err := player.Update(context.TODO(), lavalink.WithVolume(volume)); err != nil {
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Embeds: []*discordgo.MessageEmbed{embedError(err)},
			},
		})
		return
	}
	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Embeds: []*discordgo.MessageEmbed{embedVolume(volume)},
		},
	})

}

func Filter(c *extensions.Ex, s *discordgo.Session, i *discordgo.InteractionCreate) {
	player := c.Bot.Lavalink.ExistingPlayer(snowflake.MustParse(i.GuildID))
	filter := player.Filters()

	f := i.ApplicationCommandData().Options[0].StringValue()
	switch f {
	case "karaoke":
		filter.Karaoke = &lavalink.Karaoke{Level: 1.0,
			MonoLevel:   1.0,
			FilterBand:  220.0,
			FilterWidth: 100.0,
		}
	case "8d":
		filter.Rotation = &lavalink.Rotation{
			RotationHz: 0.2,
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

	if player == nil {
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Embeds: []*discordgo.MessageEmbed{embedPlayerNotFound(c.Lang(i.Locale.String()))},
			},
		})
		return
	}
	if err := player.Update(context.TODO(), lavalink.WithFilters(filter)); err != nil {
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Embeds: []*discordgo.MessageEmbed{embedError(err)},
			},
		})
		return
	}
	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Embeds: []*discordgo.MessageEmbed{embedFilters(f)},
		},
	})
}
