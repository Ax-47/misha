package lava

import (
	"context"
	"fmt"
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
	if auto && len(queue.Tracks) == 0 {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		cache := b.Queues.Cache[event.GuildID().String()]
		c := b.Lavalink.BestNode()
		if res, _ := c.LoadTracks(ctx, fmt.Sprintf("https://www.youtube.com/watch?v=%v&list=RD%v", cache, cache)); res.LoadType == lavalink.LoadTypePlaylistLoaded {
			queue.Add(res.Tracks[1])

		} else {
			cache = "gykWYPrArbY"
			res, _ := c.LoadTracks(ctx, fmt.Sprintf("https://www.youtube.com/watch?v=%v&list=RD%v", cache, cache))
			queue.Add(res.Tracks[1])

		}
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
