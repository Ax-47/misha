package bot

import (
	"context"
	"fmt"
	"log/slog"
	"strconv"
	"time"

	"github.com/disgoorg/disgolink/v3/disgolink"
	"github.com/disgoorg/disgolink/v3/lavalink"
)

func (b *Bot) OnPlayerPause(player disgolink.Player, event lavalink.PlayerPauseEvent) {

}

func (b *Bot) OnPlayerResume(player disgolink.Player, event lavalink.PlayerResumeEvent) {

}

func (b *Bot) OnTrackStart(player disgolink.Player, event lavalink.TrackStartEvent) {
}

func (b *Bot) OnTrackEnd(player disgolink.Player, event lavalink.TrackEndEvent) {
	guilid := event.GuildID()
	queue := b.Queues.Get(guilid)
	var (
		nextTrack lavalink.Track
		ok        bool
	)
	switch queue.Type {
	case QueueTypeNormal:
		nextTrack, ok = queue.Next()

	case QueueTypeRepeatTrack:
		nextTrack = event.Track

	case QueueTypeRepeatQueue:
		queue.Add(event.Track)
		nextTrack, ok = queue.Next()
	}
	if queue.Autoplay && !ok && player.Track() == nil {
		nextTrack = b.findtrack(player.Node(), event.Track.Info.Identifier)
	} else if !queue.Autoplay && !ok && player.Track() == nil {
		b.Cache.SetCache(guilid.String(), strconv.Itoa(player.Volume()))
		b.Client.UpdateVoiceState(context.TODO(), guilid, nil, false, false)
		return
	}
	if err := player.Update(context.TODO(), lavalink.WithTrack(nextTrack)); err != nil {
		slog.Error("Failed to play next track", slog.Any("err", err))
	}
}

func (b *Bot) OnTrackException(player disgolink.Player, event lavalink.TrackExceptionEvent) {

}

func (b *Bot) OnTrackStuck(player disgolink.Player, event lavalink.TrackStuckEvent) {

}

func (b *Bot) OnWebSocketClosed(player disgolink.Player, event lavalink.WebSocketClosedEvent) {

}

func (b *Bot) OnUnknownEvent(p disgolink.Player, e lavalink.UnknownEvent) {
	slog.Info("unknown event", slog.Any("event", e.Type()), slog.String("data", string(e.Data)))
}

func (b *Bot) findtrack(node disgolink.Node, identifier string) lavalink.Track {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if res, _ := node.LoadTracks(ctx, fmt.Sprintf("https://www.youtube.com/watch?v=%v&list=RD%v", identifier, identifier)); res.LoadType == lavalink.LoadTypePlaylist {
		tracks := res.Data.(lavalink.Playlist).Tracks
		fmt.Println(tracks)
		return tracks[1]

	}
	identifier = "gykWYPrArbY"
	res, _ := node.LoadTracks(ctx, fmt.Sprintf("https://www.youtube.com/watch?v=%v&list=RD%v", identifier, identifier))

	tracks := res.Data.(lavalink.Playlist).Tracks
	return tracks[1]

}
