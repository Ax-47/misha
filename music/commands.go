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

func Shuffle(c *extensions.Ex, s *discordgo.Session, i *discordgo.InteractionCreate, data discordgo.ApplicationCommandInteractionData) error {
	queue := c.Bot.Queues.Get(i.GuildID)
	if queue == nil {
		return s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "No player found",
			},
		})
	}

	queue.Shuffle()
	return s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "Queue shuffled",
		},
	})
}

func QueueType(c *extensions.Ex, s *discordgo.Session, i *discordgo.InteractionCreate, data discordgo.ApplicationCommandInteractionData) error {
	queue := c.Bot.Queues.Get(i.GuildID)
	if queue == nil {
		return s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "No player found",
			},
		})
	}

	queue.Type = lava.QueueType(data.Options[0].Value.(string))
	return s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: fmt.Sprintf("Queue type set to `%s`", queue.Type),
		},
	})
}

func ClearQueue(c *extensions.Ex, s *discordgo.Session, i *discordgo.InteractionCreate, data discordgo.ApplicationCommandInteractionData) error {
	queue := c.Bot.Queues.Get(i.GuildID)
	if queue == nil {
		return s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "No player found",
			},
		})
	}

	queue.Clear()
	return s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "Queue cleared",
		},
	})
}

func Queue(c *extensions.Ex, s *discordgo.Session, i *discordgo.InteractionCreate, data discordgo.ApplicationCommandInteractionData) {
	queue := c.Bot.Queues.Get(i.GuildID)
	if queue == nil {
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "No player found",
			},
		})
		return
	}

	if len(queue.Tracks) == 0 {
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "No tracks in queue",
			},
		})
		return
	}

	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Embeds: []*discordgo.MessageEmbed{embedQueue(1, queue, i.Member.User.ID)},
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

func Pause(c *extensions.Ex, s *discordgo.Session, i *discordgo.InteractionCreate, data discordgo.ApplicationCommandInteractionData) error {
	player := c.Bot.Lavalink.ExistingPlayer(snowflake.MustParse(i.GuildID))
	if player == nil {
		return s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "No player found",
			},
		})
	}

	if err := player.Update(context.TODO(), lavalink.WithPaused(!player.Paused())); err != nil {
		return s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Embeds: []*discordgo.MessageEmbed{embedError(err)},
			},
		})
	}

	status := "playing"
	if player.Paused() {
		status = "paused"
	}

	return s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: fmt.Sprintf("Player is now %s", status),
		},
	})
}

func Stop(c *extensions.Ex, s *discordgo.Session, i *discordgo.InteractionCreate, data discordgo.ApplicationCommandInteractionData) error {
	player := c.Bot.Lavalink.ExistingPlayer(snowflake.MustParse(i.GuildID))
	if player == nil {
		return s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "No player found",
			},
		})
	}

	if err := s.ChannelVoiceJoinManual(i.GuildID, "", false, false); err != nil {
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
			Content: "Player stopped",
		},
	})
}

func NowPlaying(c *extensions.Ex, s *discordgo.Session, i *discordgo.InteractionCreate, data discordgo.ApplicationCommandInteractionData) error {
	player := c.Bot.Lavalink.ExistingPlayer(snowflake.MustParse(i.GuildID))
	if player == nil {
		return s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "No player found",
			},
		})
	}

	track := player.Track()
	if track == nil {
		return s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "No track found",
			},
		})
	}

	return s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: fmt.Sprintf("Now playing: [`%s`](<%s>)\n\n %s / %s", track.Info.Title, *track.Info.URI, FormatPosition(player.Position()), FormatPosition(track.Info.Length)),
		},
	})
}

func FormatPosition(position lavalink.Duration) string {
	if position == 0 {
		return "0:00"
	}
	return fmt.Sprintf("%d:%02d", position.Minutes(), position.SecondsPart())
}

func Play(c *extensions.Ex, s *discordgo.Session, i *discordgo.InteractionCreate, data discordgo.ApplicationCommandInteractionData) error {
	identifier := data.Options[0].StringValue()
	if !urlPattern.MatchString(identifier) {
		identifier = lavalink.SearchTypeYoutube.Apply(identifier)
	}

	voiceState, err := s.State.VoiceState(i.GuildID, i.Member.User.ID)
	if err != nil {
		return s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Embeds: []*discordgo.MessageEmbed{embedError(err)},
			},
		})
	}
	c.Bot.Queues.Cache[i.GuildID] = "QpiCSmdYPRU"
	if err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseDeferredChannelMessageWithSource,
	}); err != nil {
		return err
	}

	player := c.Bot.Lavalink.Player(snowflake.MustParse(i.GuildID))
	queue := c.Bot.Queues.Get(i.GuildID)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var toPlay *lavalink.Track
	c.Bot.Lavalink.BestNode().LoadTracksHandler(ctx, identifier, disgolink.NewResultHandler(
		func(track lavalink.Track) {
			_, _ = s.InteractionResponseEdit(i.Interaction, &discordgo.WebhookEdit{
				Embeds: &[]*discordgo.MessageEmbed{embedPlayFoundTrack(track)},
			})
			c.Bot.Queues.Cache[i.GuildID] = track.Info.Identifier
			if player.Track() == nil {
				toPlay = &track
			} else {

				queue.Add(track)
			}
		},
		func(playlist lavalink.Playlist) {
			_, _ = s.InteractionResponseEdit(i.Interaction, &discordgo.WebhookEdit{
				Embeds: &[]*discordgo.MessageEmbed{embedPlayFoundPlaylist(playlist, identifier)},
			})
			if player.Track() == nil {
				toPlay = &playlist.Tracks[0]
				queue.Add(playlist.Tracks[1:]...)
			} else {
				queue.Add(playlist.Tracks...)
			}
		},
		func(tracks []lavalink.Track) {
			_, _ = s.InteractionResponseEdit(i.Interaction, &discordgo.WebhookEdit{
				Embeds: &[]*discordgo.MessageEmbed{embedPlayFoundTrack(tracks[0])},
			})
			c.Bot.Queues.Cache[i.GuildID] = tracks[0].Info.Identifier
			if player.Track() == nil {
				toPlay = &tracks[0]
			} else {
				queue.Add(tracks[0])

			}
		},
		func() {
			_, _ = s.InteractionResponseEdit(i.Interaction, &discordgo.WebhookEdit{
				Embeds: &[]*discordgo.MessageEmbed{embedPlayNotFound()},
			})
		},
		func(err error) {
			_, _ = s.InteractionResponseEdit(i.Interaction, &discordgo.WebhookEdit{
				Embeds: &[]*discordgo.MessageEmbed{embedError(err)},
			})
		},
	))
	if toPlay == nil {
		return nil
	}

	if err := s.ChannelVoiceJoinManual(i.GuildID, voiceState.ChannelID, false, true); err != nil {
		s.InteractionResponseEdit(i.Interaction, &discordgo.WebhookEdit{
			Embeds: &[]*discordgo.MessageEmbed{embedError(err)},
		})
		return err
	}

	return player.Update(context.TODO(), lavalink.WithTrack(*toPlay))
}
func Skip(c *extensions.Ex, s *discordgo.Session, i *discordgo.InteractionCreate, data discordgo.ApplicationCommandInteractionData) error {
	player := c.Bot.Lavalink.ExistingPlayer(snowflake.MustParse(i.GuildID))
	queue := c.Bot.Queues.Get(i.GuildID)
	if player == nil || queue == nil {
		return s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "No player found",
			},
		})
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
				func(err error) {
					fmt.Println(err)
				}))
	}
	track, ok := queue.Next()
	if !ok {
		if err := s.ChannelVoiceJoinManual(i.GuildID, "", false, false); err != nil {
			return s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Embeds: []*discordgo.MessageEmbed{embedError(err)},
				},
			})
		}

	}

	if err := player.Update(context.TODO(), lavalink.WithTrack(track)); err != nil {
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
			Content: "Skipped",
		},
	})
}
func Seek(c *extensions.Ex, s *discordgo.Session, i *discordgo.InteractionCreate, data discordgo.ApplicationCommandInteractionData) error {
	player := c.Bot.Lavalink.ExistingPlayer(snowflake.MustParse(i.GuildID))
	identifier := data.Options[0].IntValue()

	if player == nil {
		return s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "No player found",
			},
		})
	}

	if err := player.Update(context.TODO(), lavalink.WithPosition(lavalink.Duration(lavalink.Duration(identifier).Seconds()))); err != nil {
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
			Content: "Seeked",
		},
	})
}
func Remove(c *extensions.Ex, s *discordgo.Session, i *discordgo.InteractionCreate, data discordgo.ApplicationCommandInteractionData) error {
	player := c.Bot.Lavalink.ExistingPlayer(snowflake.MustParse(i.GuildID))
	identifier := data.Options[0].IntValue()
	queue := c.Bot.Queues.Get(i.GuildID)
	if player == nil || queue == nil {
		return s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "No player found",
			},
		})
	}
	queue.Delete(int(identifier))
	return s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "Delete",
		},
	})
}
func Move(c *extensions.Ex, s *discordgo.Session, i *discordgo.InteractionCreate, data discordgo.ApplicationCommandInteractionData) error {
	player := c.Bot.Lavalink.ExistingPlayer(snowflake.MustParse(i.GuildID))
	adress1 := data.Options[0].IntValue()
	adress2 := data.Options[1].IntValue()
	queue := c.Bot.Queues.Get(i.GuildID)
	if player == nil || queue == nil {
		return s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "No player found",
			},
		})
	}
	queue.Swap(int(adress1), int(adress2))
	return s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "Delete",
		},
	})
}
func Autoplay(c *extensions.Ex, s *discordgo.Session, i *discordgo.InteractionCreate, data discordgo.ApplicationCommandInteractionData) error {
	player := c.Bot.Lavalink.ExistingPlayer(snowflake.MustParse(i.GuildID))
	autoplay := c.Bot.Queues.GetAuto(i.GuildID)
	if player == nil {
		return s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "No player found",
			},
		})
	}
	c.Bot.Queues.Autoplay[i.GuildID] = !autoplay
	return s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "Autoplay",
		},
	})
}
