package music

import (
	"context"
	"fmt"
	"math/rand"
	"misha/extensions"
	"misha/lava"
	"regexp"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/disgoorg/disgolink/v2/disgolink"
	"github.com/disgoorg/disgolink/v2/lavalink"
	"github.com/disgoorg/snowflake/v2"
	"github.com/zmb3/spotify/v2"
)

var (
	urlPattern = regexp.MustCompile("^https?://[-a-zA-Z0-9+&@#/%?=~_|!:,.;]*[-a-zA-Z0-9+&@#/%=~_|]?")
	spotifyURL = regexp.MustCompile("https?://open.spotify.com/(?P<type>album|playlist|track|artist)/(?P<id>[a-zA-Z0-9]+)")
)

func Shuffle(c *extensions.Ex, s *discordgo.Session, i *discordgo.InteractionCreate) {
	langCode := c.Lang(i.Locale.String())
	queue := c.Bot.Queues.Get(i.GuildID)
	if queue == nil {
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Embeds: []*discordgo.MessageEmbed{embedPlayerNotFound(langCode)},
			},
		})
		return
	}

	queue.Shuffle()
	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Embeds: []*discordgo.MessageEmbed{embedShifle(langCode)},
		},
	})

}

func QueueType(c *extensions.Ex, s *discordgo.Session, i *discordgo.InteractionCreate) {
	langCode := c.Lang(i.Locale.String())
	data := i.ApplicationCommandData()
	queue := c.Bot.Queues.Get(i.GuildID)
	if queue == nil {
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Embeds: []*discordgo.MessageEmbed{embedPlayerNotFound(langCode)},
			},
		})
		return
	}

	queue.Type = lava.QueueType(data.Options[0].Value.(string))
	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Embeds: []*discordgo.MessageEmbed{embedLoop(langCode, queue.Type.String())},
		},
	})

}

func ClearQueue(c *extensions.Ex, s *discordgo.Session, i *discordgo.InteractionCreate) {

	langCode := c.Lang(i.Locale.String())
	queue := c.Bot.Queues.Get(i.GuildID)
	if queue == nil {
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Embeds: []*discordgo.MessageEmbed{embedPlayerNotFound(langCode)},
			},
		})
		return
	}

	queue.Clear()
	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Embeds: []*discordgo.MessageEmbed{embedClearQueue(langCode)},
		},
	})

}

func Queue(c *extensions.Ex, s *discordgo.Session, i *discordgo.InteractionCreate) {
	langCode := c.Lang(i.Locale.String())
	queue := c.Bot.Queues.Get(i.GuildID)
	if queue == nil {
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Embeds: []*discordgo.MessageEmbed{embedPlayerNotFound(langCode)},
			},
		})
		return
	}

	if len(queue.Tracks) == 0 {
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Embeds: []*discordgo.MessageEmbed{embedTracksInQueueNotFound(langCode)},
			},
		})
		return
	}

	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Embeds: []*discordgo.MessageEmbed{embedQueue(langCode, 1, queue, i.Member.User.ID)},
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
	langCode := c.Lang(i.Locale.String())
	player := c.Bot.Lavalink.ExistingPlayer(snowflake.MustParse(i.GuildID))
	if player == nil {
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Embeds: []*discordgo.MessageEmbed{embedPlayerNotFound(langCode)},
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
			Embeds: []*discordgo.MessageEmbed{embedPause(langCode, status)},
		},
	})

}

func Stop(c *extensions.Ex, s *discordgo.Session, i *discordgo.InteractionCreate) {
	langCode := c.Lang(i.Locale.String())
	player := c.Bot.Lavalink.ExistingPlayer(snowflake.MustParse(i.GuildID))
	if player == nil {
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Embeds: []*discordgo.MessageEmbed{embedPlayerNotFound(langCode)},
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
			Embeds: []*discordgo.MessageEmbed{embedStop(langCode)},
		},
	})

}

func NowPlaying(c *extensions.Ex, s *discordgo.Session, i *discordgo.InteractionCreate) {
	langCode := c.Lang(i.Locale.String())
	player := c.Bot.Lavalink.ExistingPlayer(snowflake.MustParse(i.GuildID))
	if player == nil {
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Embeds: []*discordgo.MessageEmbed{embedPlayerNotFound(langCode)},
			},
		})
		return
	}

	track := player.Track()
	if track == nil {
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Embeds: []*discordgo.MessageEmbed{embedNotFoundTrackPlaying(langCode)},
			},
		})
		return
	}

	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Embeds: []*discordgo.MessageEmbed{embedNow(langCode, *track, player)},
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
	langCode := c.Lang(i.Locale.String())
	data := i.ApplicationCommandData()
	identifier := data.Options[0].StringValue()
	if !urlPattern.MatchString(identifier) {
		identifier = lavalink.SearchTypeYoutube.Apply(identifier)
	}
	var err error

	botr := make(chan *discordgo.VoiceState)
	go func(bots chan *discordgo.VoiceState) {
		sbot, _ := s.State.VoiceState(i.GuildID, s.State.User.ID)
		bots <- sbot
	}(botr)

	voiceState, err := s.State.VoiceState(i.GuildID, i.Member.User.ID)
	if err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{Type: discordgo.InteractionResponseDeferredChannelMessageWithSource}); err != nil {
		return
	}
	if c.Bot.Lavalink == nil {
		s.InteractionResponseEdit(i.Interaction, &discordgo.WebhookEdit{
			Embeds: &[]*discordgo.MessageEmbed{embedErrorLavalink()},
		})
		return
	}
	switch err {
	default:
		s.InteractionResponseEdit(i.Interaction, &discordgo.WebhookEdit{
			Embeds: &[]*discordgo.MessageEmbed{embedError(err)},
		})
		return
	case discordgo.ErrStateNotFound:
		s.InteractionResponseEdit(i.Interaction, &discordgo.WebhookEdit{
			Embeds: &[]*discordgo.MessageEmbed{embedJoin(langCode)},
		})
		return

	case nil:
	}

	bot := <-botr
	if bot != nil {
		if bot.ChannelID != voiceState.ChannelID {
			s.InteractionResponseEdit(i.Interaction, &discordgo.WebhookEdit{
				Embeds: &[]*discordgo.MessageEmbed{embedUser(langCode)},
			},
			)
			return
		}
	}
	player := c.Bot.Lavalink.Player(snowflake.MustParse(i.GuildID))
	queue := c.Bot.Queues.Get(i.GuildID)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var toPlay *lavalink.Track
	identifierMap := []string{}
	is_spotify := false
	if spotifyURL.MatchString(identifier) {
		is_spotify = true
		//album|playlist|track|artist
		ctx := context.Background()

		if strings.Contains(identifier, "track") {
			results, err := c.Spotify.GetTrack(ctx, spotify.ID(strings.Split(identifier, "https://open.spotify.com/track/")[1]))
			if err != nil {
				s.InteractionResponseEdit(i.Interaction, &discordgo.WebhookEdit{
					Embeds: &[]*discordgo.MessageEmbed{embedNotFoundTrack(langCode)},
				})
				return
			}
			identifierMap = append(identifierMap, lavalink.SearchTypeYoutube.Apply(results.ExternalIDs["isrc"]))
			s.InteractionResponseEdit(i.Interaction, &discordgo.WebhookEdit{
				Embeds: &[]*discordgo.MessageEmbed{embedPlayFoundTrackSpotify(langCode, results)},
			})
		} else if strings.Contains(identifier, "playlist") {

			results, err := c.Spotify.GetPlaylist(ctx, spotify.ID(strings.Split(identifier, "https://open.spotify.com/playlist/")[1]))
			if err != nil {
				s.InteractionResponseEdit(i.Interaction, &discordgo.WebhookEdit{
					Embeds: &[]*discordgo.MessageEmbed{embedNotFoundTrack(langCode)},
				})
				return
			}

			for _, item := range results.Tracks.Tracks {
				identifierMap = append(identifierMap, lavalink.SearchTypeYoutube.Apply(item.Track.ExternalIDs["isrc"]))

			}
			s.InteractionResponseEdit(i.Interaction, &discordgo.WebhookEdit{
				Embeds: &[]*discordgo.MessageEmbed{embedPlayFoundspotifyPlaylistSpotify(langCode, results)},
			})
		} else if strings.Contains(identifier, "album") {
			id := spotify.ID(strings.Split(identifier, "https://open.spotify.com/album/")[1])
			results, err := c.Spotify.GetAlbum(ctx, id)
			if err != nil {
				s.InteractionResponseEdit(i.Interaction, &discordgo.WebhookEdit{
					Embeds: &[]*discordgo.MessageEmbed{embedNotFoundTrack(langCode)},
				})
				return
			}

			resultsa, _ := c.Spotify.GetAlbumTracks(ctx, id)

			for _, Track := range resultsa.Tracks {
				results, _ := c.Spotify.GetTrack(ctx, spotify.ID(strings.Split(Track.ExternalURLs["spotify"], "https://open.spotify.com/track/")[1]))
				identifierMap = append(identifierMap, lavalink.SearchTypeYoutube.Apply(results.ExternalIDs["isrc"]))
			}
			s.InteractionResponseEdit(i.Interaction, &discordgo.WebhookEdit{
				Embeds: &[]*discordgo.MessageEmbed{embedPlayFoundspotifyAlbumSpotify(langCode, results)},
			})
		} else if strings.Contains(identifier, "artist") {

			results, err := c.Spotify.GetArtistsTopTracks(ctx, spotify.ID(strings.Split(identifier, "https://open.spotify.com/artist/")[1]), "TH")
			if err != nil {
				s.InteractionResponseEdit(i.Interaction, &discordgo.WebhookEdit{
					Embeds: &[]*discordgo.MessageEmbed{embedNotFoundTrack(langCode)},
				})
				return
			}
			for _, item := range results {
				identifierMap = append(identifierMap, lavalink.SearchTypeYoutube.Apply(item.ExternalIDs["isrc"]))

			}
			fmt.Println(results)
			s.InteractionResponseEdit(i.Interaction, &discordgo.WebhookEdit{
				Embeds: &[]*discordgo.MessageEmbed{embedPlayFoundspotifyArtistSpotify(langCode, results)},
			})
		}
		for _, k := range identifierMap {
			results, err := c.Bot.Lavalink.BestNode().LoadTracks(ctx, k)
			if err != nil {
				s.InteractionResponseEdit(i.Interaction, &discordgo.WebhookEdit{
					Embeds: &[]*discordgo.MessageEmbed{embedNotFoundTrack(langCode)},
				})
				break
			}
			if results.LoadType == lavalink.LoadTypeNoMatches || results.LoadType == lavalink.LoadTypeLoadFailed {
				continue
			}
			if results.Tracks[0].Info.SourceName == "youtube" {
				c.Bot.Queues.Cache[i.GuildID] = results.Tracks[0].Info.Identifier
			}
			if player.Track() == nil && toPlay == nil {
				toPlay = &results.Tracks[0]
			} else {
				queue.Add(results.Tracks...)
			}
		}
	} else {

		c.Bot.Lavalink.BestNode().LoadTracksHandler(ctx, identifier, disgolink.NewResultHandler(
			func(track lavalink.Track) {
				s.InteractionResponseEdit(i.Interaction, &discordgo.WebhookEdit{
					Embeds: &[]*discordgo.MessageEmbed{embedPlayFoundTrack(langCode, track)},
				})
				if track.Info.SourceName == "youtube" {
					c.Bot.Queues.Cache[i.GuildID] = track.Info.Identifier
				}
				if player.Track() == nil {
					toPlay = &track
				} else {

					queue.Add(track)
				}

			},
			func(playlist lavalink.Playlist) {
				s.InteractionResponseEdit(i.Interaction, &discordgo.WebhookEdit{
					Embeds: &[]*discordgo.MessageEmbed{embedPlayFoundPlaylist(langCode, playlist, identifier)},
				})
				if player.Track() == nil {
					toPlay = &playlist.Tracks[0]
					queue.Add(playlist.Tracks[1:]...)
				} else {
					queue.Add(playlist.Tracks...)
				}

			},
			func(tracks []lavalink.Track) {
				if !is_spotify {
					s.InteractionResponseEdit(i.Interaction, &discordgo.WebhookEdit{
						Embeds: &[]*discordgo.MessageEmbed{embedPlayFoundTrack(langCode, tracks[0])},
					})
				}
				if tracks[0].Info.SourceName == "youtube" {
					c.Bot.Queues.Cache[i.GuildID] = tracks[0].Info.Identifier
				}
				if player.Track() == nil {
					toPlay = &tracks[0]
				} else {
					queue.Add(tracks[0])
				}
				fmt.Println(tracks[0].Info.Identifier)
			},
			func() {
				if !is_spotify {
					s.InteractionResponseEdit(i.Interaction, &discordgo.WebhookEdit{
						Embeds: &[]*discordgo.MessageEmbed{embedNotFoundTrack(langCode)},
					})
				}
			},
			func(err error) {
				if !is_spotify {
					s.InteractionResponseEdit(i.Interaction, &discordgo.WebhookEdit{
						Embeds: &[]*discordgo.MessageEmbed{embedError(err)},
					})
				}
			},
		))
	}

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
	langCode := c.Lang(i.Locale.String())
	player := c.Bot.Lavalink.ExistingPlayer(snowflake.MustParse(i.GuildID))
	queue := c.Bot.Queues.Get(i.GuildID)
	if player == nil || queue == nil {
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Embeds: []*discordgo.MessageEmbed{embedPlayerNotFound(langCode)},
			},
		})
		return
	}
	if c.Bot.Queues.GetAuto(i.GuildID) {
		r := rand.Intn(25)
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		cache := c.Bot.Queues.Cache[i.GuildID]
		cha := false

		for {
			if cha {
				break
			}

			c.Bot.Lavalink.BestNode().LoadTracksHandler(ctx,
				fmt.Sprintf("https://www.youtube.com/watch?v=%v&list=RD%v", cache, cache),
				disgolink.NewResultHandler(func(track lavalink.Track) {
					cache = "gykWYPrArbY"

				}, func(playlist lavalink.Playlist) {
					queue.Add(playlist.Tracks[r])
					cache = playlist.Tracks[r].Info.Identifier
					cha = true

				}, func(tracks []lavalink.Track) {
				}, func() {
					cache = "gykWYPrArbY"

				}, func(err error) {
					cache = "gykWYPrArbY"

				}))

		}
		c.Bot.Queues.Cache[i.GuildID] = cache
	}
	track, ok := queue.Next()

	if !ok {
		s.ChannelVoiceJoinManual(i.GuildID, " ", false, false)
	} else {

		if err := player.Update(context.TODO(), lavalink.WithTrack(track)); err != nil {
			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Embeds: []*discordgo.MessageEmbed{embedError(err)},
				},
			})
			return
		}
	}

	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Embeds: []*discordgo.MessageEmbed{embedSkip(langCode)},
		},
	})

}
func Seek(c *extensions.Ex, s *discordgo.Session, i *discordgo.InteractionCreate) {
	langCode := c.Lang(i.Locale.String())
	data := i.ApplicationCommandData()
	player := c.Bot.Lavalink.ExistingPlayer(snowflake.MustParse(i.GuildID))
	identifier := data.Options[0].IntValue()

	if player == nil {
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Embeds: []*discordgo.MessageEmbed{embedPlayerNotFound(langCode)},
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
			Embeds: []*discordgo.MessageEmbed{embedSeek(langCode, FormatPosition(duration))},
		},
	})

}
func Remove(c *extensions.Ex, s *discordgo.Session, i *discordgo.InteractionCreate) {
	langCode := c.Lang(i.Locale.String())
	data := i.ApplicationCommandData()
	player := c.Bot.Lavalink.ExistingPlayer(snowflake.MustParse(i.GuildID))
	identifier := data.Options[0].IntValue()
	queue := c.Bot.Queues.Get(i.GuildID)
	if player == nil || queue == nil {
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Embeds: []*discordgo.MessageEmbed{embedPlayerNotFound(langCode)},
			},
		})
		return
	}
	if len(queue.Tracks) == 0 {
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Embeds: []*discordgo.MessageEmbed{embedTracksInQueueNotFound(langCode)},
			},
		})
		return

	}
	song := queue.Tracks[identifier-1].Info.Title
	queue.Delete(int(identifier - 1))
	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Embeds: []*discordgo.MessageEmbed{embedRemove(langCode, song)},
		},
	})

}
func Swap(c *extensions.Ex, s *discordgo.Session, i *discordgo.InteractionCreate) {
	langCode := c.Lang(i.Locale.String())
	data := i.ApplicationCommandData()
	player := c.Bot.Lavalink.ExistingPlayer(snowflake.MustParse(i.GuildID))
	adress1 := data.Options[0].IntValue()
	adress2 := data.Options[1].IntValue()
	queue := c.Bot.Queues.Get(i.GuildID)
	if player == nil || queue == nil {
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Embeds: []*discordgo.MessageEmbed{embedPlayerNotFound(langCode)},
			},
		})
		return
	}
	if len(queue.Tracks) == 0 {
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Embeds: []*discordgo.MessageEmbed{embedTracksInQueueNotFound(langCode)},
			},
		})
		return

	}
	if len(queue.Tracks) < 2 {
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Embeds: []*discordgo.MessageEmbed{embedQueueLessThanTwo(langCode)},
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
			Embeds: []*discordgo.MessageEmbed{embedSwap(langCode, song1, song2)},
		},
	})

}
func Autoplay(c *extensions.Ex, s *discordgo.Session, i *discordgo.InteractionCreate) {
	langCode := c.Lang(i.Locale.String())
	player := c.Bot.Lavalink.ExistingPlayer(snowflake.MustParse(i.GuildID))
	autoplay := c.Bot.Queues.GetAuto(i.GuildID)
	if player == nil {
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Embeds: []*discordgo.MessageEmbed{embedPlayerNotFound(langCode)},
			},
		})
		return
	}
	c.Bot.Queues.Autoplay[i.GuildID] = !autoplay
	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Embeds: []*discordgo.MessageEmbed{embedAutoPlay(langCode, !autoplay)},
		},
	})

}
