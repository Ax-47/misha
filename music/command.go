package music

import (
	"context"
	"fmt"
	"misha/extensions"
	"misha/lava"
	"regexp"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/disgoorg/disgolink/v2/disgolink"
	"github.com/disgoorg/disgolink/v2/lavalink"
	"github.com/disgoorg/json"
	"github.com/disgoorg/snowflake/v2"
)

var (
	urlPattern    = regexp.MustCompile("^https?://[-a-zA-Z0-9+&@#/%?=~_|!:,.;]*[-a-zA-Z0-9+&@#/%=~_|]?")
	searchPattern = regexp.MustCompile(`^(.{2})search:(.+)`)
)

func Shuffle(c *extensions.Ex, s *discordgo.Session, event *discordgo.InteractionCreate, data discordgo.ApplicationCommandInteractionData) error {
	queue := c.Bot.Queues.Get(event.GuildID)
	if queue == nil {
		return s.InteractionRespond(event.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "No player found",
			},
		})
	}

	queue.Shuffle()
	return s.InteractionRespond(event.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "Queue shuffled",
		},
	})
}

func QueueType(c *extensions.Ex, s *discordgo.Session, event *discordgo.InteractionCreate, data discordgo.ApplicationCommandInteractionData) error {
	queue := c.Bot.Queues.Get(event.GuildID)
	if queue == nil {
		return s.InteractionRespond(event.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "No player found",
			},
		})
	}

	queue.Type = lava.QueueType(data.Options[0].Value.(string))
	return s.InteractionRespond(event.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: fmt.Sprintf("Queue type set to `%s`", queue.Type),
		},
	})
}

func ClearQueue(c *extensions.Ex, s *discordgo.Session, event *discordgo.InteractionCreate, data discordgo.ApplicationCommandInteractionData) error {
	queue := c.Bot.Queues.Get(event.GuildID)
	if queue == nil {
		return s.InteractionRespond(event.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "No player found",
			},
		})
	}

	queue.Clear()
	return s.InteractionRespond(event.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "Queue cleared",
		},
	})
}

func Queue(c *extensions.Ex, s *discordgo.Session, event *discordgo.InteractionCreate, data discordgo.ApplicationCommandInteractionData) error {
	queue := c.Bot.Queues.Get(event.GuildID)
	if queue == nil {
		return s.InteractionRespond(event.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "No player found",
			},
		})
	}

	if len(queue.Tracks) == 0 {
		return s.InteractionRespond(event.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "No tracks in queue",
			},
		})
	}

	var tracks string
	for i, track := range queue.Tracks {
		tracks += fmt.Sprintf("%d. [`%s`](<%s>)\n", i+1, track.Info.Title, *track.Info.URI)
	}

	return s.InteractionRespond(event.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: fmt.Sprintf("Queue `%s`:\n%s", queue.Type, tracks),
		},
	})
}

func Pause(c *extensions.Ex, s *discordgo.Session, event *discordgo.InteractionCreate, data discordgo.ApplicationCommandInteractionData) error {
	player := c.Bot.Lavalink.ExistingPlayer(snowflake.MustParse(event.GuildID))
	if player == nil {
		return s.InteractionRespond(event.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "No player found",
			},
		})
	}

	if err := player.Update(context.TODO(), lavalink.WithPaused(!player.Paused())); err != nil {
		return s.InteractionRespond(event.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: fmt.Sprintf("Error while pausing: `%s`", err),
			},
		})
	}

	status := "playing"
	if player.Paused() {
		status = "paused"
	}

	return s.InteractionRespond(event.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: fmt.Sprintf("Player is now %s", status),
		},
	})
}

func Stop(c *extensions.Ex, s *discordgo.Session, event *discordgo.InteractionCreate, data discordgo.ApplicationCommandInteractionData) error {
	player := c.Bot.Lavalink.ExistingPlayer(snowflake.MustParse(event.GuildID))
	if player == nil {
		return s.InteractionRespond(event.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "No player found",
			},
		})
	}

	if err := s.ChannelVoiceJoinManual(event.GuildID, "", false, false); err != nil {
		return s.InteractionRespond(event.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: fmt.Sprintf("Error while disconnecting: `%s`", err),
			},
		})
	}

	return s.InteractionRespond(event.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "Player stopped",
		},
	})
}

func NowPlaying(c *extensions.Ex, s *discordgo.Session, event *discordgo.InteractionCreate, data discordgo.ApplicationCommandInteractionData) error {
	player := c.Bot.Lavalink.ExistingPlayer(snowflake.MustParse(event.GuildID))
	if player == nil {
		return s.InteractionRespond(event.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "No player found",
			},
		})
	}

	track := player.Track()
	if track == nil {
		return s.InteractionRespond(event.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "No track found",
			},
		})
	}

	return s.InteractionRespond(event.Interaction, &discordgo.InteractionResponse{
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

func Play(c *extensions.Ex, s *discordgo.Session, event *discordgo.InteractionCreate, data discordgo.ApplicationCommandInteractionData) error {
	identifier := data.Options[0].StringValue()
	if !urlPattern.MatchString(identifier) && !searchPattern.MatchString(identifier) {
		identifier = lavalink.SearchTypeYoutube.Apply(identifier)
	}

	voiceState, err := s.State.VoiceState(event.GuildID, event.Member.User.ID)
	if err != nil {
		return s.InteractionRespond(event.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: fmt.Sprintf("Error while getting voice state: `%s`", err),
			},
		})
	}

	if err := s.InteractionRespond(event.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseDeferredChannelMessageWithSource,
	}); err != nil {
		return err
	}

	player := c.Bot.Lavalink.Player(snowflake.MustParse(event.GuildID))
	queue := c.Bot.Queues.Get(event.GuildID)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var toPlay *lavalink.Track
	c.Bot.Lavalink.BestNode().LoadTracksHandler(ctx, identifier, disgolink.NewResultHandler(
		func(track lavalink.Track) {
			_, _ = s.InteractionResponseEdit(event.Interaction, &discordgo.WebhookEdit{
				Content: json.Ptr(fmt.Sprintf("Loading track: [`%s`](<%s>)", track.Info.Title, *track.Info.URI)),
			})
			if player.Track() == nil {
				toPlay = &track
			} else {
				queue.Add(track)
			}
		},
		func(playlist lavalink.Playlist) {
			_, _ = s.InteractionResponseEdit(event.Interaction, &discordgo.WebhookEdit{
				Content: json.Ptr(fmt.Sprintf("Loaded playlist: `%s` with `%d` tracks", playlist.Info.Name, len(playlist.Tracks))),
			})
			if player.Track() == nil {
				toPlay = &playlist.Tracks[0]
				queue.Add(playlist.Tracks[1:]...)
			} else {
				queue.Add(playlist.Tracks...)
			}
		},
		func(tracks []lavalink.Track) {
			_, _ = s.InteractionResponseEdit(event.Interaction, &discordgo.WebhookEdit{
				Content: json.Ptr(fmt.Sprintf("Loaded search result: [`%s`](<%s>)", tracks[0].Info.Title, *tracks[0].Info.URI)),
			})
			if player.Track() == nil {
				toPlay = &tracks[0]
			} else {
				queue.Add(tracks[0])
			}
		},
		func() {
			_, _ = s.InteractionResponseEdit(event.Interaction, &discordgo.WebhookEdit{
				Content: json.Ptr(fmt.Sprintf("Nothing found for: `%s`", identifier)),
			})
		},
		func(err error) {
			_, _ = s.InteractionResponseEdit(event.Interaction, &discordgo.WebhookEdit{
				Content: json.Ptr(fmt.Sprintf("Error while looking up query: `%s`", err)),
			})
		},
	))
	if toPlay == nil {
		return nil
	}

	if err := s.ChannelVoiceJoinManual(event.GuildID, voiceState.ChannelID, false, false); err != nil {
		return err
	}

	return player.Update(context.TODO(), lavalink.WithTrack(*toPlay))
}
