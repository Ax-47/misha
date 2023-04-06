package music

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/disgoorg/disgolink/v2/lavalink"
)

func embedPlayFoundTrack(track lavalink.Track) *discordgo.MessageEmbed {
	return &discordgo.MessageEmbed{
		Color: 0xff4700,
		Title: "🎵 added track",
		Fields: []*discordgo.MessageEmbedField{
			&discordgo.MessageEmbedField{
				Name:   "song",
				Value:  fmt.Sprintf("[%s](%s)", track.Info.Title, *track.Info.URI),
				Inline: false,
			},
			&discordgo.MessageEmbedField{
				Name:   "author",
				Value:  track.Info.Author,
				Inline: true,
			},
			&discordgo.MessageEmbedField{
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
			&discordgo.MessageEmbedField{
				Name:   "playlist",
				Value:  fmt.Sprintf("[%s](%s)", playlist.Info.Name, link),
				Inline: false,
			},

			&discordgo.MessageEmbedField{
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
