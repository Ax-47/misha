package music

import (
	"fmt"
	"misha/lava"

	"github.com/bwmarrin/discordgo"
	"github.com/disgoorg/disgolink/v2/lavalink"
)

func embedPlayFoundTrack(track lavalink.Track) *discordgo.MessageEmbed {
	return &discordgo.MessageEmbed{
		Color: 0xff4700,
		Title: "🎵 added track",
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
func embedPlayFoundPlaylist(playlist lavalink.Playlist, link string) *discordgo.MessageEmbed {
	return &discordgo.MessageEmbed{
		Color: 0xff4700,
		Title: "🎵 added the playlist",
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
func embedPlayNotFound() *discordgo.MessageEmbed {
	return &discordgo.MessageEmbed{
		Color: 0xff4700,
		Title: ">_ 404 Not Found ",
	}
}
func embedError(err error) *discordgo.MessageEmbed {
	return &discordgo.MessageEmbed{
		Color:       0xff4700,
		Title:       "have error",
		Description: err.Error(),
	}
}

func embedQueue(index int, queue *lava.Queue, id string) *discordgo.MessageEmbed {
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
		Title:       fmt.Sprintf("%d Tracks in Queue", lengthtracks),
		Description: tracks,

		Footer: &discordgo.MessageEmbedFooter{
			Text: fmt.Sprintf("page %d/%d | %s", index+1, pages, id),
		},
	}
}
