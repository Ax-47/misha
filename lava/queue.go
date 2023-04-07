package lava

import (
	"math/rand"

	"github.com/bwmarrin/discordgo"
	"github.com/disgoorg/disgolink/v2/disgolink"
	"github.com/disgoorg/disgolink/v2/lavalink"
)

type Bot struct {
	Lavalink disgolink.Client
	Queues   *QueueManager
	S        *discordgo.Session
}

type QueueType string

const (
	QueueTypeNormal      QueueType = "normal"
	QueueTypeRepeatTrack QueueType = "repeat_track"
	QueueTypeRepeatQueue QueueType = "repeat_queue"
)

func (q QueueType) String() string {
	switch q {
	case QueueTypeNormal:
		return "Normal"
	case QueueTypeRepeatTrack:
		return "Repeat Track"
	case QueueTypeRepeatQueue:
		return "Repeat Queue"
	default:
		return "unknown"
	}
}

type Queue struct {
	Tracks []lavalink.Track
	Type   QueueType
}

func (q *Queue) Shuffle() {
	rand.Shuffle(len(q.Tracks), func(i, j int) {
		q.Tracks[i], q.Tracks[j] = q.Tracks[j], q.Tracks[i]
	})
}

func (q *Queue) Add(track ...lavalink.Track) {
	q.Tracks = append(q.Tracks, track...)
}
func (q *Queue) Delete(track int) {

	q.Tracks = append(q.Tracks[:track], q.Tracks[track+1:]...)

}
func (q *Queue) Swap(track1, track2 int) {
	q.Tracks[track1], q.Tracks[track2] = q.Tracks[track2], q.Tracks[track1]

}
func (q *Queue) Next() (lavalink.Track, bool) {
	if len(q.Tracks) == 0 {
		return lavalink.Track{}, false
	}
	track := q.Tracks[0]
	q.Tracks = q.Tracks[1:]
	return track, true
}

func (q *Queue) Clear() {
	q.Tracks = make([]lavalink.Track, 0)
}

type QueueManager struct {
	Queues   map[string]*Queue
	Autoplay map[string]bool
	Cache    map[string]string
}

func (q *QueueManager) GetAuto(guildID string) bool {
	Autoplay, ok := q.Autoplay[guildID]
	if !ok {
		Autoplay = false
		q.Autoplay[guildID] = Autoplay
	}
	return Autoplay
}

func (q *QueueManager) Get(guildID string) *Queue {
	queue, ok := q.Queues[guildID]
	if !ok {
		queue = &Queue{
			Tracks: make([]lavalink.Track, 0),
			Type:   QueueTypeNormal,
		}
		q.Queues[guildID] = queue
	}
	return queue
}

func (q *QueueManager) Delete(guildID string) {
	delete(q.Queues, guildID)
}
