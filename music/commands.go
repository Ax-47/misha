package music

import (
	"context"
	"fmt"
	"math/rand"
	"misha/extensions"
	"misha/lava"
	"regexp"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/disgoorg/disgolink/v2/disgolink"
	"github.com/disgoorg/disgolink/v2/lavalink"
	"github.com/disgoorg/snowflake/v2"
)

var (
	urlPattern = regexp.MustCompile("^https?://[-a-zA-Z0-9+&@#/%?=~_|!:,.;]*[-a-zA-Z0-9+&@#/%=~_|]?")
)

func Shuffle(c *extensions.Ex, s *discordgo.Session, i *discordgo.InteractionCreate) {
	queue := c.Bot.Queues.Get(i.GuildID)
	if queue == nil {
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Embeds: []*discordgo.MessageEmbed{embedPlayerNotFound(c.Lang(i.Locale.String()))},
			},
		})
		return
	}

	queue.Shuffle()
	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Embeds: []*discordgo.MessageEmbed{embedShifle(c.Lang(i.Locale.String()))},
		},
	})

}

func QueueType(c *extensions.Ex, s *discordgo.Session, i *discordgo.InteractionCreate) {
	data := i.ApplicationCommandData()
	queue := c.Bot.Queues.Get(i.GuildID)
	if queue == nil {
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Embeds: []*discordgo.MessageEmbed{embedPlayerNotFound(c.Lang(i.Locale.String()))},
			},
		})
		return
	}

	queue.Type = lava.QueueType(data.Options[0].Value.(string))
	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Embeds: []*discordgo.MessageEmbed{embedLoop(c.Lang(i.Locale.String()), queue.Type.String())},
		},
	})

}

func ClearQueue(c *extensions.Ex, s *discordgo.Session, i *discordgo.InteractionCreate) {
	queue := c.Bot.Queues.Get(i.GuildID)
	if queue == nil {
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Embeds: []*discordgo.MessageEmbed{embedPlayerNotFound(c.Lang(i.Locale.String()))},
			},
		})
		return
	}

	queue.Clear()
	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Embeds: []*discordgo.MessageEmbed{embedClearQueue(c.Lang(i.Locale.String()))},
		},
	})

}

func Queue(c *extensions.Ex, s *discordgo.Session, i *discordgo.InteractionCreate) {
	queue := c.Bot.Queues.Get(i.GuildID)
	if queue == nil {
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Embeds: []*discordgo.MessageEmbed{embedPlayerNotFound(c.Lang(i.Locale.String()))},
			},
		})
		return
	}

	if len(queue.Tracks) == 0 {
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Embeds: []*discordgo.MessageEmbed{embedTracksInQueueNotFound(c.Lang(i.Locale.String()))},
			},
		})
		return
	}

	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Embeds: []*discordgo.MessageEmbed{embedQueue(c.Lang(i.Locale.String()), 1, queue, i.Member.User.ID)},
			Components: []discordgo.MessageComponent{
				discordgo.ActionsRow{
					Components: []discordgo.MessageComponent{
						discordgo.Button{
							Label: "◀",
							Style: discordgo.SuccessButton,

							CustomID: "re",
						},
						discordgo.Button{
							Label: "▶",
							Style: discordgo.SuccessButton,

							CustomID: "next",
						},
					},
				},
			},
		},
	})
}

func Pause(c *extensions.Ex, s *discordgo.Session, i *discordgo.InteractionCreate) {
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

	if err := player.Update(context.TODO(), lavalink.WithPaused(!player.Paused())); err != nil {
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Embeds: []*discordgo.MessageEmbed{embedError(err)},
			},
		})
		return
	}

	status := "playing"
	if player.Paused() {
		status = "paused"
	}

	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Embeds: []*discordgo.MessageEmbed{embedPause(c.Lang(i.Locale.String()), status)},
		},
	})

}

func Stop(c *extensions.Ex, s *discordgo.Session, i *discordgo.InteractionCreate) {
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

	if err := s.ChannelVoiceJoinManual(i.GuildID, "", false, false); err != nil {
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
			Embeds: []*discordgo.MessageEmbed{embedStop(c.Lang(i.Locale.String()))},
		},
	})

}

func NowPlaying(c *extensions.Ex, s *discordgo.Session, i *discordgo.InteractionCreate) {
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

	track := player.Track()
	if track == nil {
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Embeds: []*discordgo.MessageEmbed{embedNotFoundTrackPlaying(c.Lang(i.Locale.String()))},
			},
		})
		return
	}

	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Embeds: []*discordgo.MessageEmbed{embedNow(c.Lang(i.Locale.String()), *track, player)},
		},
	})

}

func FormatPosition(position lavalink.Duration) string {
	if position == 0 {
		return "0:00"
	}
	return fmt.Sprintf("%d:%02d", position.Minutes(), position.SecondsPart())
}

func Play(c *extensions.Ex, s *discordgo.Session, i *discordgo.InteractionCreate) {
	data := i.ApplicationCommandData()
	identifier := data.Options[0].StringValue()
	if !urlPattern.MatchString(identifier) {
		identifier = lavalink.SearchTypeYoutube.Apply(identifier)
	}
	bot, _ := s.State.VoiceState(i.GuildID, s.State.User.ID)
	voiceState, err := s.State.VoiceState(i.GuildID, i.Member.User.ID)
	if bot != nil {
		if bot != voiceState {
			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Embeds: []*discordgo.MessageEmbed{embedUser(c.Lang(i.Locale.String()))},
				},
			})
			return
		}
	}
	switch err {
	default:
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Embeds: []*discordgo.MessageEmbed{embedError(err)},
			},
		})
		return
	case discordgo.ErrStateNotFound:
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Embeds: []*discordgo.MessageEmbed{embedJoin(c.Lang(i.Locale.String()))},
			},
		})
		return

	case nil:
	}

	if err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseDeferredChannelMessageWithSource,
	}); err != nil {
		return
	}
	if c.Bot.Lavalink == nil {
		s.InteractionResponseEdit(i.Interaction, &discordgo.WebhookEdit{
			Embeds: &[]*discordgo.MessageEmbed{embedErrorLavalink()},
		})
	}
	player := c.Bot.Lavalink.Player(snowflake.MustParse(i.GuildID))
	queue := c.Bot.Queues.Get(i.GuildID)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var toPlay *lavalink.Track
	c.Bot.Lavalink.BestNode().LoadTracksHandler(ctx, identifier, disgolink.NewResultHandler(
		func(track lavalink.Track) {
			go s.InteractionResponseEdit(i.Interaction, &discordgo.WebhookEdit{
				Embeds: &[]*discordgo.MessageEmbed{embedPlayFoundTrack(c.Lang(i.Locale.String()), track)},
			})
			c.Bot.Queues.Cache[i.GuildID] = track.Info.Identifier
			if player.Track() == nil {
				toPlay = &track
			} else {

				queue.Add(track)
			}
		},
		func(playlist lavalink.Playlist) {
			go s.InteractionResponseEdit(i.Interaction, &discordgo.WebhookEdit{
				Embeds: &[]*discordgo.MessageEmbed{embedPlayFoundPlaylist(c.Lang(i.Locale.String()), playlist, identifier)},
			})
			if player.Track() == nil {
				toPlay = &playlist.Tracks[0]
				queue.Add(playlist.Tracks[1:]...)
			} else {
				queue.Add(playlist.Tracks...)
			}
		},
		func(tracks []lavalink.Track) {
			go s.InteractionResponseEdit(i.Interaction, &discordgo.WebhookEdit{
				Embeds: &[]*discordgo.MessageEmbed{embedPlayFoundTrack(c.Lang(i.Locale.String()), tracks[0])},
			})
			c.Bot.Queues.Cache[i.GuildID] = tracks[0].Info.Identifier
			if player.Track() == nil {
				toPlay = &tracks[0]
			} else {
				queue.Add(tracks[0])

			}
		},
		func() {
			go s.InteractionResponseEdit(i.Interaction, &discordgo.WebhookEdit{
				Embeds: &[]*discordgo.MessageEmbed{embedNotFoundTrack(c.Lang(i.Locale.String()))},
			})
		},
		func(err error) {
			go s.InteractionResponseEdit(i.Interaction, &discordgo.WebhookEdit{
				Embeds: &[]*discordgo.MessageEmbed{embedError(err)},
			})
		},
	))
	if toPlay == nil {
		return
	}

	if err := s.ChannelVoiceJoinManual(i.GuildID, voiceState.ChannelID, false, true); err != nil {
		s.InteractionResponseEdit(i.Interaction, &discordgo.WebhookEdit{
			Embeds: &[]*discordgo.MessageEmbed{embedError(err)},
		})
		return
	}

	player.Update(context.TODO(), lavalink.WithTrack(*toPlay))

}
func Skip(c *extensions.Ex, s *discordgo.Session, i *discordgo.InteractionCreate) {
	player := c.Bot.Lavalink.ExistingPlayer(snowflake.MustParse(i.GuildID))
	queue := c.Bot.Queues.Get(i.GuildID)
	if player == nil || queue == nil {
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Embeds: []*discordgo.MessageEmbed{embedPlayerNotFound(c.Lang(i.Locale.String()))},
			},
		})
		return
	}
	if c.Bot.Queues.GetAuto(i.GuildID) {
		r := rand.Intn(25)
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		cache := c.Bot.Queues.Cache[i.GuildID]

		c.Bot.Lavalink.BestNode().LoadTracksHandler(ctx,
			fmt.Sprintf("https://www.youtube.com/watch?v=%v&list=RD%v", cache, cache),
			disgolink.NewResultHandler(func(track lavalink.Track) {
			}, func(playlist lavalink.Playlist) {
				queue.Add(playlist.Tracks[r])
				c.Bot.Queues.Cache[i.GuildID] = playlist.Tracks[r].Info.Identifier
			}, func(tracks []lavalink.Track) {
			}, func() {},
				func(err error) {}))
	}
	track, ok := queue.Next()
	if !ok {
		if err := s.ChannelVoiceJoinManual(i.GuildID, "", false, false); err != nil {
			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Embeds: []*discordgo.MessageEmbed{embedError(err)},
				},
			})
			return
		}

	}

	if err := player.Update(context.TODO(), lavalink.WithTrack(track)); err != nil {
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
			Embeds: []*discordgo.MessageEmbed{embedSkip(c.Lang(i.Locale.String()))},
		},
	})

}
func Seek(c *extensions.Ex, s *discordgo.Session, i *discordgo.InteractionCreate) {
	data := i.ApplicationCommandData()
	player := c.Bot.Lavalink.ExistingPlayer(snowflake.MustParse(i.GuildID))
	identifier := data.Options[0].IntValue()

	if player == nil {
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Embeds: []*discordgo.MessageEmbed{embedPlayerNotFound(c.Lang(i.Locale.String()))},
			},
		})
		return
	}
	duration := lavalink.Duration(lavalink.Duration(identifier).Seconds())
	if err := player.Update(context.TODO(), lavalink.WithPosition(duration)); err != nil {
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
			Embeds: []*discordgo.MessageEmbed{embedSeek(c.Lang(i.Locale.String()), FormatPosition(duration))},
		},
	})

}
func Remove(c *extensions.Ex, s *discordgo.Session, i *discordgo.InteractionCreate) {
	data := i.ApplicationCommandData()
	player := c.Bot.Lavalink.ExistingPlayer(snowflake.MustParse(i.GuildID))
	identifier := data.Options[0].IntValue()
	queue := c.Bot.Queues.Get(i.GuildID)
	if player == nil || queue == nil {
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Embeds: []*discordgo.MessageEmbed{embedPlayerNotFound(c.Lang(i.Locale.String()))},
			},
		})
		return
	}
	if len(queue.Tracks) == 0 {
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Embeds: []*discordgo.MessageEmbed{embedTracksInQueueNotFound(c.Lang(i.Locale.String()))},
			},
		})
		return

	}
	song := queue.Tracks[identifier-1].Info.Title
	queue.Delete(int(identifier))
	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Embeds: []*discordgo.MessageEmbed{embedRemove(c.Lang(i.Locale.String()), song)},
		},
	})

}
func Swap(c *extensions.Ex, s *discordgo.Session, i *discordgo.InteractionCreate) {
	data := i.ApplicationCommandData()
	player := c.Bot.Lavalink.ExistingPlayer(snowflake.MustParse(i.GuildID))
	adress1 := data.Options[0].IntValue()
	adress2 := data.Options[1].IntValue()
	queue := c.Bot.Queues.Get(i.GuildID)
	if player == nil || queue == nil {
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Embeds: []*discordgo.MessageEmbed{embedPlayerNotFound(c.Lang(i.Locale.String()))},
			},
		})
		return
	}
	if len(queue.Tracks) == 0 {
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Embeds: []*discordgo.MessageEmbed{embedTracksInQueueNotFound(c.Lang(i.Locale.String()))},
			},
		})
		return

	}
	if len(queue.Tracks) < 2 {
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Embeds: []*discordgo.MessageEmbed{embedQueueLessThanTwo(c.Lang(i.Locale.String()))},
			},
		})
		return

	}
	song1 := queue.Tracks[adress1-1].Info.Title
	song2 := queue.Tracks[adress2-1].Info.Title
	queue.Swap(int(adress1), int(adress2))
	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Embeds: []*discordgo.MessageEmbed{embedSwap(c.Lang(i.Locale.String()), song1, song2)},
		},
	})

}
func Autoplay(c *extensions.Ex, s *discordgo.Session, i *discordgo.InteractionCreate) {
	player := c.Bot.Lavalink.ExistingPlayer(snowflake.MustParse(i.GuildID))
	autoplay := c.Bot.Queues.GetAuto(i.GuildID)
	if player == nil {
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Embeds: []*discordgo.MessageEmbed{embedPlayerNotFound(c.Lang(i.Locale.String()))},
			},
		})
		return
	}
	c.Bot.Queues.Autoplay[i.GuildID] = !autoplay
	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Embeds: []*discordgo.MessageEmbed{embedAutoPlay(c.Lang(i.Locale.String()), !autoplay)},
		},
	})

}
