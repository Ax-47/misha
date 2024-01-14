package bot

import (
	"context"
	"fmt"
	"regexp"
	"strconv"
	"time"

	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/disgo/events"

	"0x47/misha/embed"

	"github.com/disgoorg/disgolink/v3/disgolink"
	"github.com/disgoorg/disgolink/v3/lavalink"
)

var (
	urlPattern    = regexp.MustCompile("^https?://[-a-zA-Z0-9+&@#/%?=~_|!:,.;]*[-a-zA-Z0-9+&@#/%=~_|]?")
	searchPattern = regexp.MustCompile(`^(.{2})search:(.+)`)
)
var bassBoost = &lavalink.Equalizer{
	0:  0.2,
	1:  0.15,
	2:  0.1,
	3:  0.05,
	4:  0.0,
	5:  -0.05,
	6:  -0.1,
	7:  -0.1,
	8:  -0.1,
	9:  -0.1,
	10: -0.1,
	11: -0.1,
	12: -0.1,
	13: -0.1,
	14: -0.1,
}

func (b *Bot) Shuffle(event *events.ApplicationCommandInteractionCreate, data discord.SlashCommandInteractionData) error {
	queue := b.Queues.Get(*event.GuildID())
	if queue == nil {
		return event.CreateMessage(discord.MessageCreate{
			Embeds: []discord.Embed{
				embed.ErrorNotFoundPlayer(),
			},
		})
	}

	queue.Shuffle()
	return event.CreateMessage(discord.MessageCreate{
		Embeds: []discord.Embed{
			embed.Shuffle(),
		},
	})
}

func (b *Bot) Volume(event *events.ApplicationCommandInteractionCreate, data discord.SlashCommandInteractionData) error {
	player := b.Lavalink.ExistingPlayer(*event.GuildID())
	if player == nil {
		return event.CreateMessage(discord.MessageCreate{
			Embeds: []discord.Embed{
				embed.ErrorNotFoundPlayer(),
			},
		})
	}

	volume := data.Int("volume")
	if volume > 1000 {
		return event.CreateMessage(discord.MessageCreate{
			Embeds: []discord.Embed{
				embed.ErrorOverLimitVoice(),
			},
		})
	}
	if volume < 0 {
		return event.CreateMessage(discord.MessageCreate{
			Embeds: []discord.Embed{
				embed.ErrorUnderZeroVoice(),
			},
		})
	}
	if err := player.Update(context.TODO(), lavalink.WithVolume(volume)); err != nil {
		return event.CreateMessage(discord.MessageCreate{
			Embeds: []discord.Embed{
				embed.ErrorOther(),
			},
		})
	}

	return event.CreateMessage(discord.MessageCreate{
		Embeds: []discord.Embed{
			embed.Volume(strconv.Itoa(volume)),
		},
	})
}

func (b *Bot) Seek(event *events.ApplicationCommandInteractionCreate, data discord.SlashCommandInteractionData) error {
	player := b.Lavalink.ExistingPlayer(*event.GuildID())
	if player == nil {
		return event.CreateMessage(discord.MessageCreate{
			Embeds: []discord.Embed{
				embed.ErrorNotFoundPlayer(),
			},
		})
	}

	position := data.Int("position")
	unit, ok := data.OptInt("unit")
	if !ok || unit == 0 {
		unit = 1
	}
	finalPosition := lavalink.Duration(position * unit)
	if err := player.Update(context.TODO(), lavalink.WithPosition(finalPosition)); err != nil {
		return event.CreateMessage(discord.MessageCreate{
			Embeds: []discord.Embed{
				embed.ErrorOther(),
			},
		})
	}

	return event.CreateMessage(discord.MessageCreate{
		Embeds: []discord.Embed{
			embed.Seek(FormatPosition(finalPosition)),
		},
	})
}

func (b *Bot) BassBoost(event *events.ApplicationCommandInteractionCreate, data discord.SlashCommandInteractionData) error {
	player := b.Lavalink.ExistingPlayer(*event.GuildID())
	if player == nil {
		return event.CreateMessage(discord.MessageCreate{
			Embeds: []discord.Embed{
				embed.ErrorNotFoundPlayer(),
			},
		})
	}

	enabled := data.Bool("enabled")
	filters := player.Filters()
	if enabled {
		filters.Equalizer = bassBoost
	} else {
		filters.Equalizer = nil
	}

	if err := player.Update(context.TODO(), lavalink.WithFilters(filters)); err != nil {
		return event.CreateMessage(discord.MessageCreate{
			Embeds: []discord.Embed{
				embed.ErrorOther(),
			},
		})
	}

	return event.CreateMessage(discord.MessageCreate{
		Embeds: []discord.Embed{
			embed.BassBoost(enabled),
		},
	})
}

func (b *Bot) Skip(event *events.ApplicationCommandInteractionCreate, data discord.SlashCommandInteractionData) error {
	player := b.Lavalink.ExistingPlayer(*event.GuildID())
	queue := b.Queues.Get(*event.GuildID())
	if player == nil || queue == nil {
		return event.CreateMessage(discord.MessageCreate{
			Embeds: []discord.Embed{
				embed.ErrorNotFoundPlayer(),
			},
		})
	}

	amount, ok := data.OptInt("amount")
	if !ok {
		amount = 1
	}

	track, ok := queue.Skip(amount)

	if !ok && !queue.Autoplay {
		if err := player.Update(context.TODO(), lavalink.WithNullTrack()); err != nil {
			return event.CreateMessage(discord.MessageCreate{
				Embeds: []discord.Embed{
					embed.ErrorOther(),
				},
			})
		}

		return event.CreateMessage(discord.MessageCreate{
			Embeds: []discord.Embed{
				embed.ErrorQueueIsNil(),
			},
		})
	} else if !ok && queue.Autoplay {

		track = b.findtrack(player.Node(), player.Track().Info.Identifier)
	}

	if err := player.Update(context.TODO(), lavalink.WithTrack(track)); err != nil {
		return event.CreateMessage(discord.MessageCreate{
			Embeds: []discord.Embed{
				embed.ErrorOther(),
			},
		})
	}

	return event.CreateMessage(discord.MessageCreate{
		Embeds: []discord.Embed{
			embed.Skip(),
		},
	})
}

func (b *Bot) Loop(event *events.ApplicationCommandInteractionCreate, data discord.SlashCommandInteractionData) error {
	queue := b.Queues.Get(*event.GuildID())
	if queue == nil {
		return event.CreateMessage(discord.MessageCreate{
			Embeds: []discord.Embed{
				embed.ErrorNotFoundPlayer(),
			},
		})
	}

	queue.Type = QueueType(data.String("type"))
	return event.CreateMessage(discord.MessageCreate{
		Embeds: []discord.Embed{
			embed.Loop(queue.Type.String()),
		},
	})
}

func (b *Bot) ClearQueue(event *events.ApplicationCommandInteractionCreate, data discord.SlashCommandInteractionData) error {
	queue := b.Queues.Get(*event.GuildID())
	if queue == nil {
		return event.CreateMessage(discord.MessageCreate{
			Embeds: []discord.Embed{
				embed.ErrorNotFoundPlayer(),
			},
		})
	}

	queue.Clear()
	return event.CreateMessage(discord.MessageCreate{
		Embeds: []discord.Embed{
			embed.Clear(),
		},
	})
}

func (b *Bot) Queue(event *events.ApplicationCommandInteractionCreate, data discord.SlashCommandInteractionData) error {
	queue := b.Queues.Get(*event.GuildID())
	if queue == nil {
		return event.CreateMessage(discord.MessageCreate{
			Embeds: []discord.Embed{
				embed.ErrorNotFoundPlayer(),
			},
		})
	}

	if len(queue.Tracks) == 0 {
		return event.CreateMessage(discord.MessageCreate{
			Embeds: []discord.Embed{
				embed.ErrorQueueIsNil(),
			},
		})
	}
	return event.CreateMessage(discord.MessageCreate{
		Embeds: []discord.Embed{
			embed.Queue(event.Member().User.ID.String(), queue.Type.String(), queue.Tracks, 1),
		},
		Components: []discord.ContainerComponent{embed.QueueButtons()},
	})
}

func (b *Bot) Players(event *events.ApplicationCommandInteractionCreate, data discord.SlashCommandInteractionData) error {
	var description string
	b.Lavalink.ForPlayers(func(player disgolink.Player) {
		description += fmt.Sprintf("GuildID: `%s`\n", player.GuildID())
	})

	return event.CreateMessage(discord.MessageCreate{
		Content: fmt.Sprintf("Players:\n%s", description),
	})
}

func (b *Bot) Pause(event *events.ApplicationCommandInteractionCreate, data discord.SlashCommandInteractionData) error {
	player := b.Lavalink.ExistingPlayer(*event.GuildID())
	if player == nil {
		return event.CreateMessage(discord.MessageCreate{
			Embeds: []discord.Embed{
				embed.ErrorNotFoundPlayer(),
			},
		})
	}

	if err := player.Update(context.TODO(), lavalink.WithPaused(!player.Paused())); err != nil {
		return event.CreateMessage(discord.MessageCreate{
			Embeds: []discord.Embed{
				embed.ErrorOther(),
			},
		})
	}
	return event.CreateMessage(discord.MessageCreate{
		Embeds: []discord.Embed{
			embed.Pause(player.Paused()),
		},
	})
}

func (b *Bot) Stop(event *events.ApplicationCommandInteractionCreate, data discord.SlashCommandInteractionData) error {
	player := b.Lavalink.ExistingPlayer(*event.GuildID())
	if player == nil {
		return event.CreateMessage(discord.MessageCreate{
			Embeds: []discord.Embed{
				embed.ErrorNotFoundPlayer(),
			},
		})
	}

	if err := player.Update(context.TODO(), lavalink.WithNullTrack()); err != nil {
		return event.CreateMessage(discord.MessageCreate{
			Embeds: []discord.Embed{
				embed.ErrorOther(),
			},
		})
	}

	return event.CreateMessage(discord.MessageCreate{
		Embeds: []discord.Embed{
			embed.Stop(),
		},
	})
}

func (b *Bot) Disconnect(event *events.ApplicationCommandInteractionCreate, data discord.SlashCommandInteractionData) error {
	player := b.Lavalink.ExistingPlayer(*event.GuildID())
	if player == nil {
		return event.CreateMessage(discord.MessageCreate{
			Embeds: []discord.Embed{
				embed.ErrorNotFoundPlayer(),
			},
		})
	}
	if err := player.Update(context.TODO(), lavalink.WithNullTrack()); err != nil {
		return event.CreateMessage(discord.MessageCreate{
			Embeds: []discord.Embed{
				embed.ErrorOther(),
			},
		})
	}
	if err := b.Client.UpdateVoiceState(context.TODO(), *event.GuildID(), nil, false, false); err != nil {
		return event.CreateMessage(discord.MessageCreate{
			Embeds: []discord.Embed{
				embed.ErrorOther(),
			},
		})
	}

	return event.CreateMessage(discord.MessageCreate{
		Embeds: []discord.Embed{
			embed.Disconnect(),
		},
	})
}

func (b *Bot) NowPlaying(event *events.ApplicationCommandInteractionCreate, data discord.SlashCommandInteractionData) error {
	player := b.Lavalink.ExistingPlayer(*event.GuildID())
	if player == nil {
		return event.CreateMessage(discord.MessageCreate{
			Embeds: []discord.Embed{
				embed.ErrorNotFoundPlayer(),
			},
		})
	}

	track := player.Track()
	if track == nil {
		return event.CreateMessage(discord.MessageCreate{
			Embeds: []discord.Embed{
				embed.ErrorYouHaveNotPlay(),
			},
		})
	}

	return event.CreateMessage(discord.MessageCreate{
		Embeds: []discord.Embed{
			embed.NowPlaying(track.Info.Title, *track.Info.URI, track.Info.Author, *track.Info.ArtworkURL, FormatPosition(track.Info.Length)),
		},
	})
}

func FormatPosition(position lavalink.Duration) string {
	if position == 0 {
		return "0:00"
	}
	return fmt.Sprintf("%d:%02d", position.Minutes(), position.SecondsPart())
}

func (b *Bot) Play(event *events.ApplicationCommandInteractionCreate, data discord.SlashCommandInteractionData) error {
	identifier := data.String("identifier")
	if source, ok := data.OptString("source"); ok {
		identifier = lavalink.SearchType(source).Apply(identifier)
	} else if !urlPattern.MatchString(identifier) && !searchPattern.MatchString(identifier) {
		identifier = lavalink.SearchTypeYouTube.Apply(identifier)
	}

	var guildID = event.GuildID()
	voiceState, ok := b.Client.Caches().VoiceState(*guildID, event.User().ID)
	if !ok {
		return event.CreateMessage(discord.MessageCreate{
			Embeds: []discord.Embed{
				embed.ErrorYouAreNotVC(),
			},
		})
	}

	if err := event.DeferCreateMessage(false); err != nil {
		return event.CreateMessage(discord.MessageCreate{
			Embeds: []discord.Embed{
				embed.ErrorOther(),
			},
		})
	}
	queue := b.Queues.Get(*guildID)
	player := b.Lavalink.Player(*guildID)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var toPlay *lavalink.Track
	b.Lavalink.BestNode().LoadTracksHandler(ctx, identifier, disgolink.NewResultHandler(
		func(track lavalink.Track) {
			_, _ = b.Client.Rest().UpdateInteractionResponse(event.ApplicationID(), event.Token(), discord.MessageUpdate{
				Embeds: &[]discord.Embed{
					embed.Track(track.Info.Title, *track.Info.URI, track.Info.Author, *track.Info.ArtworkURL, FormatPosition(track.Info.Length)),
				},
			})
			if player.Track() == nil {
				toPlay = &track
			} else {

				queue.Add(track)
			}
		},
		func(playlist lavalink.Playlist) {
			_, _ = b.Client.Rest().UpdateInteractionResponse(event.ApplicationID(), event.Token(), discord.MessageUpdate{
				Embeds: &[]discord.Embed{
					embed.Playlist(playlist.Info.Name, strconv.Itoa(len(playlist.Tracks)), *playlist.Tracks[0].Info.ArtworkURL),
				},
			})
			if player.Track() == nil {
				toPlay = &playlist.Tracks[0]
				queue.Add(playlist.Tracks[1:]...)
			} else {
				queue.Add(playlist.Tracks...)
			}
		},
		func(tracks []lavalink.Track) {
			_, _ = b.Client.Rest().UpdateInteractionResponse(event.ApplicationID(), event.Token(), discord.MessageUpdate{
				Embeds: &[]discord.Embed{
					embed.Track(tracks[0].Info.Title, *tracks[0].Info.URI, tracks[0].Info.Author, *tracks[0].Info.ArtworkURL, FormatPosition(tracks[0].Info.Length)),
				},
			})
			if player.Track() == nil {
				toPlay = &tracks[0]
			} else {
				queue.Add(tracks[0])
			}
		},
		func() {
			_, _ = b.Client.Rest().UpdateInteractionResponse(event.ApplicationID(), event.Token(), discord.MessageUpdate{
				Embeds: &[]discord.Embed{
					embed.NotFound(),
				},
			})
		},
		func(err error) {
			_, _ = b.Client.Rest().UpdateInteractionResponse(event.ApplicationID(), event.Token(), discord.MessageUpdate{
				Embeds: &[]discord.Embed{
					embed.ErrorOther(),
				},
			})
		},
	))
	if err := b.settingCache(player, guildID.String(), event); err != nil {
		_, _ = b.Client.Rest().UpdateInteractionResponse(event.ApplicationID(), event.Token(), discord.MessageUpdate{
			Embeds: &[]discord.Embed{
				embed.ErrorOther(),
			},
		})
	}
	if err := b.Client.UpdateVoiceState(context.TODO(), *guildID, voiceState.ChannelID, false, true); err != nil {
		return event.CreateMessage(discord.MessageCreate{
			Embeds: []discord.Embed{
				embed.ErrorCanNotConnectVC(),
			},
		})
	}
	if toPlay == nil {
		return nil
	}
	return player.Update(context.TODO(), lavalink.WithTrack(*toPlay))
}
func (b *Bot) Autoplay(event *events.ApplicationCommandInteractionCreate, data discord.SlashCommandInteractionData) error {
	queue := b.Queues.Get(*event.GuildID())
	if queue == nil {
		return event.CreateMessage(discord.MessageCreate{
			Embeds: []discord.Embed{
				embed.ErrorNotFoundPlayer(),
			},
		})
	}

	queue.Autoplay = !queue.Autoplay
	return event.CreateMessage(discord.MessageCreate{
		Embeds: []discord.Embed{
			embed.Autoplay(queue.Autoplay),
		},
	})
}
