package embed

import (
	"fmt"

	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/disgolink/v3/lavalink"
)

func Track(title, uri, author, artworkurl, duration string) discord.Embed {
	return discord.Embed{
		Title: "เพิ่มเพลงแล้ว",
		Fields: []discord.EmbedField{
			{Name: "เพลง", Value: fmt.Sprintf("[%s](%s)", title, uri)},
			{Name: "ศิลปิน", Value: author},
			{Name: "ระยะเวลา", Value: duration},
		},
		Thumbnail: &discord.EmbedResource{
			URL: artworkurl,
		},
		Color: 0xff4700,
	}
}
func Playlist(name, len_playlist, artworkurl string) discord.Embed {
	return discord.Embed{
		Title: "เพิ่มเพลงจากPlaylistแล้ว",
		Fields: []discord.EmbedField{
			{Name: "Playlist", Value: name},
			{Name: "Tracks", Value: len_playlist},
		},
		Thumbnail: &discord.EmbedResource{
			URL: artworkurl,
		},
		Color: 0xff4700,
	}
}
func NotFound() discord.Embed {
	return discord.Embed{
		Title: "404 หาไม่เจอ",
		Color: 0x0,
	}
}

func NowPlaying(title, uri, author, artworkurl, duration string) discord.Embed {
	return discord.Embed{
		Title: "Now Play",
		Fields: []discord.EmbedField{
			{Name: "เพลง", Value: fmt.Sprintf("[%s](%s)", title, uri)},
			{Name: "ศิลปิน", Value: author},
			{Name: "ระยะเวลา", Value: duration},
		},
		Thumbnail: &discord.EmbedResource{
			URL: artworkurl,
		},
		Color: 0xff4700,
	}
}
func Disconnect() discord.Embed {
	return discord.Embed{
		Title: "ขออนุญาตบิด",
		Color: 0xff4700,
	}
}
func Stop() discord.Embed {
	return discord.Embed{
		Title: "shhhhhhh",
		Color: 0xff4700,
	}
}
func Pause(state bool) discord.Embed {
	status := "เล่น"
	if state {
		status = "หยุด"
	}
	return discord.Embed{
		Title: fmt.Sprintf("%sเพลงให้ละ", status),
		Color: 0xff4700,
	}
}
func Queue(id, queuetype string, tracks []lavalink.Track, index int) discord.Embed {
	var tracks_ string

	lengthtracks := len(tracks)

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
	for i, track := range tracks[index*10 : end] {
		if i >= 10 {
			break
		}
		tracks_ += fmt.Sprintf("%d : [`%s`](<%s>)\n", (index*10)+i+1, track.Info.Title, *track.Info.URI)

	}
	return discord.Embed{
		Title:       fmt.Sprintf("Queue : `%s`", queuetype),
		Description: tracks_,
		Footer: &discord.EmbedFooter{
			Text: fmt.Sprintf("page %d/%d | %s", index+1, pages, id),
		},
		Color: 0xff4700,
	}
}
func Loop(queuetype string) discord.Embed {
	return discord.Embed{
		Title: fmt.Sprintf("loop :%s", queuetype),
		Color: 0xff4700,
	}
}
func Clear() discord.Embed {
	return discord.Embed{
		Title: "Clear ละ",
		Color: 0xff4700,
	}
}
func Skip() discord.Embed {
	return discord.Embed{
		Title: "Skip ละ",
		Color: 0xff4700,
	}
}
func BassBoost(state bool) discord.Embed {
	status := "เปิด"
	if state {
		status = "ปิด"
	}
	return discord.Embed{
		Title: fmt.Sprintf("%sbassboostให้ละ", status),
		Color: 0xff4700,
	}
}
func Seek(position string) discord.Embed {
	return discord.Embed{
		Title: fmt.Sprintf("กอไปที่ %sให้ละ", position),
		Color: 0xff4700,
	}
}
func Shuffle() discord.Embed {
	return discord.Embed{
		Title: "สุ่มเพลงละ",
		Color: 0xff4700,
	}
}
func Autoplay(state bool) discord.Embed {
	status := "เปิด"
	if !state {
		status = "ปิด"
	}
	return discord.Embed{
		Title: fmt.Sprintf("`%s` autoplayละ", status),
		Color: 0xff4700,
	}
}
func QueueButtons() discord.ActionRowComponent {

	return discord.NewActionRow(discord.NewSuccessButton("◀", "previous"), discord.NewSuccessButton("▶", "next"))
}
