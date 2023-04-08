package lava

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"github.com/disgoorg/log"

	"github.com/disgoorg/disgolink/v2/disgolink"
	"github.com/disgoorg/disgolink/v2/lavalink"
)

func (b *Bot) OnPlayerPause(player disgolink.Player, event lavalink.PlayerPauseEvent) {
	fmt.Printf("onPlayerPause: %v\n", event)
}

func (b *Bot) OnPlayerResume(player disgolink.Player, event lavalink.PlayerResumeEvent) {
	fmt.Printf("onPlayerResume: %v\n", event)
}

func (b *Bot) OnTrackStart(player disgolink.Player, event lavalink.TrackStartEvent) {

	fmt.Printf("onTrackStart: %v\n", b.Queues.Cache[event.GuildID().String()])
}

func (b *Bot) OnTrackEnd(player disgolink.Player, event lavalink.TrackEndEvent) {
	if !event.Reason.MayStartNext() {
		return
	}
	fmt.Printf("onTrackEnd: %v\n", event)
	queue := b.Queues.Get(event.GuildID().String())
	auto := b.Queues.GetAuto(event.GuildID().String())
	var (
		nextTrack lavalink.Track
		ok        bool
	)
	if auto {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		cache := b.Queues.Cache[event.GuildID().String()]
		url := fmt.Sprintf("https://www.youtube.com/watch?v=%v&list=RD%v", cache, cache)
		b.Lavalink.BestNode().LoadTracksHandler(ctx,
			url,
			disgolink.NewResultHandler(func(track lavalink.Track) {
			}, func(playlist lavalink.Playlist) {
				queue.Add(playlist.Tracks[rand.Intn(25)])
			}, func(tracks []lavalink.Track) {
			}, func() {},
				func(err error) {
					fmt.Println(err)
				}))
	}
	switch queue.Type {
	case QueueTypeNormal:
		nextTrack, ok = queue.Next()
		if ok {
			b.Queues.Cache[event.GuildID().String()] = nextTrack.Info.Identifier
		}
	case QueueTypeRepeatTrack:
		nextTrack = *player.Track()

	case QueueTypeRepeatQueue:
		lastTrack, _ := b.Lavalink.BestNode().DecodeTrack(context.TODO(), event.EncodedTrack)
		queue.Add(*lastTrack)
		nextTrack, ok = queue.Next()
	}

	if !ok && !auto {
		b.S.ChannelVoiceJoinManual(event.GuildID_.String(), "", false, false)
		return
	}

	if err := player.Update(context.TODO(), lavalink.WithTrack(nextTrack)); err != nil {
		log.Error("Failed to play next track: ", err)
	}

}

func (b *Bot) OnTrackException(player disgolink.Player, event lavalink.TrackExceptionEvent) {
	fmt.Printf("onTrackException: %v\n", event)
}

func (b *Bot) OnTrackStuck(player disgolink.Player, event lavalink.TrackStuckEvent) {
	fmt.Printf("onTrackStuck: %v\n", event)
}

func (b *Bot) OnWebSocketClosed(player disgolink.Player, event lavalink.WebSocketClosedEvent) {
	fmt.Printf("onWebSocketClosed: %v\n", event)
}
