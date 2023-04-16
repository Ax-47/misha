package music

import (
	"fmt"
	languages "misha/lang"
	"misha/lava"

	"github.com/bwmarrin/discordgo"
	"github.com/disgoorg/disgolink/v2/disgolink"
	"github.com/disgoorg/disgolink/v2/lavalink"
)

func embedPlayFoundTrack(l languages.Lang, track lavalink.Track) *discordgo.MessageEmbed {
	return &discordgo.MessageEmbed{
		Color: 0xff4700,
		Title: l.MusicCommands.PlayTrack,
		Fields: []*discordgo.MessageEmbedField{
			{
				Name:   "song",
				Value:  fmt.Sprintf("[%s](%s)", track.Info.Title, *track.Info.URI),
				Inline: false,
			},
			{
				Name:   "author",
				Value:  track.Info.Author,
				Inline: true,
			},
			{
				Name:   "source",
				Value:  track.Info.SourceName,
				Inline: true,
			},
		},
		Image: &discordgo.MessageEmbedImage{
			URL:    fmt.Sprintf("https://i.ytimg.com/vi/%s/maxresdefault.jpg", track.Info.Identifier),
			Width:  1280,
			Height: 720,
		},
	}
}
func embedPlayFoundPlaylist(l languages.Lang, playlist lavalink.Playlist, link string) *discordgo.MessageEmbed {
	return &discordgo.MessageEmbed{
		Color: 0xff4700,
		Title: l.MusicCommands.PlayPlaylist,
		Fields: []*discordgo.MessageEmbedField{
			{
				Name:   "playlist",
				Value:  fmt.Sprintf("[%s](%s)", playlist.Info.Name, link),
				Inline: false,
			},

			{
				Name:   "tracks",
				Value:  fmt.Sprintf("%d", len(playlist.Tracks)),
				Inline: true,
			},
		},
		Thumbnail: &discordgo.MessageEmbedThumbnail{
			URL: fmt.Sprintf("https://i.ytimg.com/vi/%s/maxresdefault.jpg", playlist.Tracks[0].Info.Identifier),
		},
	}
}
func embedPlayerNotFound(l languages.Lang) *discordgo.MessageEmbed {
	return &discordgo.MessageEmbed{
		Color: 0xff4700,
		Title: l.MusicCommands.Errors.PlayerNotFound,
	}
}
func embedTracksInQueueNotFound(l languages.Lang) *discordgo.MessageEmbed {
	return &discordgo.MessageEmbed{
		Color: 0xff4700,
		Title: l.MusicCommands.Errors.TracksInQueueNotFound,
	}
}
func embedQueueLessThanTwo(l languages.Lang) *discordgo.MessageEmbed {
	return &discordgo.MessageEmbed{
		Color: 0xff4700,
		Title: l.MusicCommands.Errors.QueueLessThanTwo,
	}
}
func embedNotFoundTrackPlaying(l languages.Lang) *discordgo.MessageEmbed {
	return &discordgo.MessageEmbed{
		Color: 0xff4700,
		Title: l.MusicCommands.Errors.NotFoundTrackPlaying,
	}
}
func embedNotFoundTrack(l languages.Lang) *discordgo.MessageEmbed {
	return &discordgo.MessageEmbed{
		Color: 0x000,
		Title: l.MusicCommands.Errors.NotfoundTrack,
	}
}
func embedJoin(l languages.Lang) *discordgo.MessageEmbed {
	return &discordgo.MessageEmbed{
		Color: 0x000,
		Title: l.MusicCommands.Errors.UserWasNotJoin,
	}
}

func embedError(err error) *discordgo.MessageEmbed {
	return &discordgo.MessageEmbed{
		Color:       0xff4700,
		Title:       "have error",
		Description: err.Error(),
	}
}
func embedUser(l languages.Lang) *discordgo.MessageEmbed {
	return &discordgo.MessageEmbed{
		Color: 0xff4700,
		Title: l.MusicCommands.Errors.UserNotInTheRoom,
	}
}
func embedErrorLavalink() *discordgo.MessageEmbed {
	return &discordgo.MessageEmbed{
		Color: 0xff4700,
		Title: "error",
	}
}
func embedQueue(l languages.Lang, index int, queue *lava.Queue, id string) *discordgo.MessageEmbed {
	var tracks string
	lengthtracks := len(queue.Tracks)

	index = index - 1

	pages := lengthtracks / 10
	if lengthtracks%10 != 0 {
		pages += 1
	}
	if index == -1 {
		index = pages - 1

	} else if index >= pages {
		index = 0

	}
	end := index*10 + 10
	if lengthtracks <= end {
		end = lengthtracks
	}
	for i, track := range queue.Tracks[index*10 : end] {
		if i >= 10 {
			break
		}
		tracks += fmt.Sprintf("%d : [`%s`](<%s>)\n", (index*10)+i+1, track.Info.Title, *track.Info.URI)

	}

	return &discordgo.MessageEmbed{
		Color:       0xff4700,
		Title:       fmt.Sprintf(l.MusicCommands.Queue, lengthtracks),
		Description: tracks,

		Footer: &discordgo.MessageEmbedFooter{
			Text: fmt.Sprintf("page %d/%d | %s", index+1, pages, id),
		},
	}
}
func embedShifle(l languages.Lang) *discordgo.MessageEmbed {
	return &discordgo.MessageEmbed{
		Color: 0xff4700,
		Title: l.MusicCommands.Shuffle,
	}
}
func embedSkip(l languages.Lang) *discordgo.MessageEmbed {
	return &discordgo.MessageEmbed{
		Color: 0xff4700,
		Title: l.MusicCommands.Skip,
	}
}
func embedStop(l languages.Lang) *discordgo.MessageEmbed {
	return &discordgo.MessageEmbed{
		Color: 0xff4700,
		Title: l.MusicCommands.Stop,
	}
}
func embedPause(l languages.Lang, status string) *discordgo.MessageEmbed {
	return &discordgo.MessageEmbed{
		Color: 0xff4700,
		Title: fmt.Sprintf(l.MusicCommands.Pause, status),
	}
}
func embedClearQueue(l languages.Lang) *discordgo.MessageEmbed {
	return &discordgo.MessageEmbed{
		Color: 0xff4700,
		Title: l.MusicCommands.ClearQueue,
	}
}
func embedLoop(l languages.Lang, typeloop string) *discordgo.MessageEmbed {
	return &discordgo.MessageEmbed{
		Color: 0xff4700,
		Title: fmt.Sprintf(l.MusicCommands.Loop, typeloop),
	}
}
func embedSeek(l languages.Lang, duration string) *discordgo.MessageEmbed {
	return &discordgo.MessageEmbed{
		Color: 0xff4700,
		Title: fmt.Sprintf(l.MusicCommands.Seek, duration),
	}
}
func embedSwap(l languages.Lang, song1, song2 string) *discordgo.MessageEmbed {
	return &discordgo.MessageEmbed{
		Color: 0xff4700,
		Title: fmt.Sprintf(l.MusicCommands.Swap, song1, song2),
	}
}
func embedRemove(l languages.Lang, title string) *discordgo.MessageEmbed {
	return &discordgo.MessageEmbed{
		Color: 0xff4700,
		Title: fmt.Sprintf(l.MusicCommands.Remove, title),
	}
}
func embedAutoPlay(l languages.Lang, on bool) *discordgo.MessageEmbed {
	var status string
	status = "on"
	if !on {
		status = "off"
	}
	return &discordgo.MessageEmbed{
		Color: 0xff4700,
		Title: fmt.Sprintf(l.MusicCommands.AutoPlay, status),
	}
}
func embedNow(l languages.Lang, track lavalink.Track, player disgolink.Player) *discordgo.MessageEmbed {
	duration := fmt.Sprintf("`%s-%s`", FormatPosition(player.Position()), FormatPosition(track.Info.Length))
	if player.Track().Info.IsStream {
		duration = fmt.Sprintf("`%s🔴`", FormatPosition(player.Position()))
	}
	return &discordgo.MessageEmbed{
		Color: 0xff4700,
		Title: track.Info.Title,
		URL:   fmt.Sprintf("https://youtu.be/%s", track.Info.Identifier),
		Thumbnail: &discordgo.MessageEmbedThumbnail{
			URL: fmt.Sprintf("https://i.ytimg.com/vi/%s/maxresdefault.jpg", track.Info.Identifier),
		},
		Fields: []*discordgo.MessageEmbedField{
			{
				Name:  "duration",
				Value: duration,
			},
		},
	}
}
func embedEqualizer(typeEq string) *discordgo.MessageEmbed {

	return &discordgo.MessageEmbed{
		Color: 0xff4700,
		Title: fmt.Sprintf("Equalizer: `%s`", typeEq),
	}
}
func embedTimescaleErrorInput(l languages.Lang, err string) *discordgo.MessageEmbed {

	return &discordgo.MessageEmbed{
		Color: 0xff4700,
		Title: fmt.Sprintf(l.FilterCommands.Errors.TimescaleErrorInput, err),
	}
}
func embedFilters(filter string) *discordgo.MessageEmbed {

	return &discordgo.MessageEmbed{
		Color: 0xff4700,
		Title: fmt.Sprintf("Filter: `%s`", filter),
	}
}
func embedVolume(Volume int) *discordgo.MessageEmbed {

	return &discordgo.MessageEmbed{
		Color: 0xff4700,
		Title: fmt.Sprintf("Volume: `%d`", Volume),
	}
}
func embedTremolo(Frequency, Depth float32) *discordgo.MessageEmbed {

	return &discordgo.MessageEmbed{
		Color: 0xff4700,
		Title: fmt.Sprintf("Frequency:`%.1f`,Depth: `%.1f`", Depth, Frequency),
	}
}
func embedTimescale(Speed, Pitch, Rate float64) *discordgo.MessageEmbed {

	return &discordgo.MessageEmbed{
		Color: 0xff4700,
		Title: fmt.Sprintf("Speed:`%.1f`,Pitch: `%.1f`,Rate: `%.1f`", Speed, Pitch, Rate),
	}
}
