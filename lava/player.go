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

func (b *Bot) OnTrackEnd(player disgolink.Player, event lavalink.TrackEndEvent) {
	if !event.Reason.MayStartNext() {
		return
	}
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
				cache = "gykWYPrArbY"
			}, func() {
				cache = "gykWYPrArbY"
			},
				func(err error) {
					cache = "gykWYPrArbY"
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
