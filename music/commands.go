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
		fmt.Printf("%v %v", bot.ChannelID, voiceState.ChannelID)
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
	typeOfUri := "unknow"

	if spotifyURL.MatchString(identifier) {
		var (
			trackRes    *spotify.FullTrack    = nil
			playlistRes *spotify.FullPlaylist = nil
			albumRes    *spotify.FullAlbum    = nil
			artistRes   []spotify.FullTrack   = nil
		)
		is_spotify = true
		//album|playlist|track|artist
		ctx := context.Background()

		if strings.Contains(identifier, "track") {
			typeOfUri = "track"
			trackRes, err = c.Spotify.GetTrack(ctx, spotify.ID(strings.Split(identifier, "https://open.spotify.com/track/")[1]))
			if err != nil {
				s.InteractionResponseEdit(i.Interaction, &discordgo.WebhookEdit{
					Embeds: &[]*discordgo.MessageEmbed{embedNotFoundTrack(langCode)},
				})
				return
			}
			if trackRes == nil {
				s.InteractionResponseEdit(i.Interaction, &discordgo.WebhookEdit{
					Embeds: &[]*discordgo.MessageEmbed{embedNotFoundTrack(langCode)},
				})
				return
			}
			identifierMap = append(identifierMap, lavalink.SearchTypeYoutube.Apply(trackRes.ExternalIDs["isrc"]))

		} else if strings.Contains(identifier, "playlist") {
			typeOfUri = "playlist"
			playlistRes, err = c.Spotify.GetPlaylist(ctx, spotify.ID(strings.Split(identifier, "https://open.spotify.com/playlist/")[1]))
			if err != nil {
				s.InteractionResponseEdit(i.Interaction, &discordgo.WebhookEdit{
					Embeds: &[]*discordgo.MessageEmbed{embedNotFoundTrack(langCode)},
				})
				return
			}
			if playlistRes == nil {
				s.InteractionResponseEdit(i.Interaction, &discordgo.WebhookEdit{
					Embeds: &[]*discordgo.MessageEmbed{embedNotFoundTrack(langCode)},
				})
				return
			}
			for _, item := range playlistRes.Tracks.Tracks { //o(n)
				identifierMap = append(identifierMap, lavalink.SearchTypeYoutube.Apply(item.Track.ExternalIDs["isrc"]))

			}

		} else if strings.Contains(identifier, "album") {
			typeOfUri = "album"
			id := spotify.ID(strings.Split(identifier, "https://open.spotify.com/album/")[1])
			albumRes, err = c.Spotify.GetAlbum(ctx, id)
			if err != nil {
				s.InteractionResponseEdit(i.Interaction, &discordgo.WebhookEdit{
					Embeds: &[]*discordgo.MessageEmbed{embedNotFoundTrack(langCode)},
				})
				return
			}
			if albumRes == nil {
				s.InteractionResponseEdit(i.Interaction, &discordgo.WebhookEdit{
					Embeds: &[]*discordgo.MessageEmbed{embedNotFoundTrack(langCode)},
				})
				return
			}
			resultsa, _ := c.Spotify.GetAlbumTracks(ctx, id)
			if resultsa == nil {
				s.InteractionResponseEdit(i.Interaction, &discordgo.WebhookEdit{
					Embeds: &[]*discordgo.MessageEmbed{embedNotFoundTrack(langCode)},
				})
				return
			}
			for _, Track := range resultsa.Tracks {
				results, _ := c.Spotify.GetTrack(ctx, spotify.ID(strings.Split(Track.ExternalURLs["spotify"], "https://open.spotify.com/track/")[1]))
				identifierMap = append(identifierMap, lavalink.SearchTypeYoutube.Apply(results.ExternalIDs["isrc"]))
			}

		} else if strings.Contains(identifier, "artist") {
			typeOfUri = "artist"
			id := spotify.ID(strings.Split(identifier, "https://open.spotify.com/artist/")[1])[0:22]
			artistRes, err = c.Spotify.GetArtistsTopTracks(ctx, id, "TH")
			if err != nil {
				s.InteractionResponseEdit(i.Interaction, &discordgo.WebhookEdit{
					Embeds: &[]*discordgo.MessageEmbed{embedNotFoundTrack(langCode)},
				})
				return
			}
			if artistRes == nil {
				s.InteractionResponseEdit(i.Interaction, &discordgo.WebhookEdit{
					Embeds: &[]*discordgo.MessageEmbed{embedNotFoundTrack(langCode)},
				})
				return
			}
			for _, item := range artistRes {
				identifierMap = append(identifierMap, lavalink.SearchTypeYoutube.Apply(item.ExternalIDs["isrc"]))
			}

		}
		node := c.Bot.Lavalink.BestNode()
		for o, k := range identifierMap {
			results, err := node.LoadTracks(ctx, k)
			if err != nil {
				s.InteractionResponseEdit(i.Interaction, &discordgo.WebhookEdit{
					Embeds: &[]*discordgo.MessageEmbed{embedNotFoundTrack(langCode)},
				})
				break
			}
			if results.LoadType == lavalink.LoadTypeNoMatches || results.LoadType == lavalink.LoadTypeLoadFailed {
				if len(identifierMap) == 1 {
					s.InteractionResponseEdit(i.Interaction, &discordgo.WebhookEdit{
						Embeds: &[]*discordgo.MessageEmbed{embedNotFoundTrack(langCode)},
					})
				}
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
			if o == 0 {
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
		}
		if toPlay == nil {
			s.InteractionResponseEdit(i.Interaction, &discordgo.WebhookEdit{
				Embeds: &[]*discordgo.MessageEmbed{embedNotFoundTrack(langCode)},
			})
			return
		}
		switch typeOfUri {
		case "track":
			s.InteractionResponseEdit(i.Interaction, &discordgo.WebhookEdit{
				Embeds: &[]*discordgo.MessageEmbed{embedPlayFoundTrackSpotify(langCode, trackRes)},
			})
		case "playlist":
			s.InteractionResponseEdit(i.Interaction, &discordgo.WebhookEdit{
				Embeds: &[]*discordgo.MessageEmbed{embedPlayFoundspotifyPlaylistSpotify(langCode, playlistRes)},
			})
		case "album":
			s.InteractionResponseEdit(i.Interaction, &discordgo.WebhookEdit{
				Embeds: &[]*discordgo.MessageEmbed{embedPlayFoundspotifyAlbumSpotify(langCode, albumRes)},
			})
		case "artist":
			s.InteractionResponseEdit(i.Interaction, &discordgo.WebhookEdit{
				Embeds: &[]*discordgo.MessageEmbed{embedPlayFoundspotifyArtistSpotify(langCode, artistRes)},
			})
		default:
			s.InteractionResponseEdit(i.Interaction, &discordgo.WebhookEdit{
				Embeds: &[]*discordgo.MessageEmbed{embedNotFoundTrack(langCode)},
			})
		}

	} else {

		c.Bot.Lavalink.BestNode().LoadTracksHandler(ctx, identifier, disgolink.NewResultHandler(
			func(track lavalink.Track) {
				s.InteractionResponseEdit(i.Interaction, &discordgo.WebhookEdit{
					Embeds: &[]*discordgo.MessageEmbed{embedPlayFoundTrack(langCode, track)},
				})
				if track.Info.SourceName == "youtube" {
					c.Bot.Queues.Cache[i.GuildID] = track.Info.Identifier
				} else {
					c.Bot.Queues.Cache[i.GuildID] = "gykWYPrArbY"
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
				} else {
					c.Bot.Queues.Cache[i.GuildID] = "gykWYPrArbY"
				}
				if player.Track() == nil {
					toPlay = &tracks[0]
				} else {
					queue.Add(tracks[0])
				}

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
	if err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{Type: discordgo.InteractionResponseDeferredChannelMessageWithSource}); err != nil {
		return
	}
	if c.Bot.Queues.GetAuto(i.GuildID) && len(queue.Tracks) == 0 {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		r := rand.Intn(3) + 1
		cache := c.Bot.Queues.Cache[i.GuildID]
		node := c.Bot.Lavalink.BestNode()
		if res, _ := node.LoadTracks(ctx, fmt.Sprintf("https://www.youtube.com/watch?v=%v&list=RD%v", cache, cache)); res.LoadType == lavalink.LoadTypePlaylistLoaded {
			queue.Add(res.Tracks[r])
		} else {
			cache = "gykWYPrArbY"
			res, _ := node.LoadTracks(ctx, fmt.Sprintf("https://www.youtube.com/watch?v=%v&list=RD%v", cache, cache))
			queue.Add(res.Tracks[r])
		}

	}
	track, ok := queue.Next()

	if !ok {
		s.ChannelVoiceJoinManual(i.GuildID, " ", false, false)
	} else {

		if err := player.Update(context.TODO(), lavalink.WithTrack(track)); err != nil {
			s.InteractionResponseEdit(i.Interaction, &discordgo.WebhookEdit{
				Embeds: &[]*discordgo.MessageEmbed{embedError(err)},
			})
			return
		}
	}
	s.InteractionResponseEdit(i.Interaction, &discordgo.WebhookEdit{
		Embeds: &[]*discordgo.MessageEmbed{embedSkip(langCode)},
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
