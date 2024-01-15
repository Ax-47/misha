package bot

import (
	"0x47/misha/embed"
	"strconv"
	"strings"

	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/disgo/events"
)

func (b *Bot) ButtonsEvent(event *events.ComponentInteractionCreate) {

	switch event.ButtonInteractionData().CustomID() {
	case "previous":
		previous(b, event)
	case "next":
		next(b, event)
	case "index":
		index(b, event)
	case "setting":
		setting(b, event)
	case "music":
		music(b, event)
	}

}

func previous(b *Bot, event *events.ComponentInteractionCreate) error {
	queue := b.Queues.Get(*event.GuildID())
	item := strings.Split(event.Message.Embeds[0].Footer.Text, " | ")
	if event.Member().User.ID.String() != item[1] {
		return event.CreateMessage(discord.MessageCreate{
			Embeds: []discord.Embed{
				embed.ErrorItIsNotYours(),
			},
			Flags: discord.MessageFlagEphemeral,
		})
	}
	var (
		page, max_page int
	)
	it := strings.Split(strings.TrimPrefix(item[0], "page "), "/")
	page, _ = strconv.Atoi(it[0])
	max_page, _ = strconv.Atoi(it[1])
	if page == 1 {
		page = max_page
	} else {
		page -= 1
	}
	return event.UpdateMessage(discord.MessageUpdate{
		Embeds: &[]discord.Embed{
			embed.Queue(item[1], queue.Type.String(), queue.Tracks, page),
		},
		Components: &[]discord.ContainerComponent{embed.QueueButtons()},
	})
}
func next(b *Bot, event *events.ComponentInteractionCreate) error {
	queue := b.Queues.Get(*event.GuildID())
	item := strings.Split(event.Message.Embeds[0].Footer.Text, " | ")
	if event.Member().User.ID.String() != item[1] {
		return event.CreateMessage(discord.MessageCreate{
			Embeds: []discord.Embed{
				embed.ErrorItIsNotYours(),
			},
			Flags: discord.MessageFlagEphemeral,
		})
	}
	var (
		page, max_page int
	)
	it := strings.Split(strings.TrimPrefix(item[0], "page "), "/")
	page, _ = strconv.Atoi(it[0])
	max_page, _ = strconv.Atoi(it[1])
	if page == max_page {
		page = 1
	} else {
		page += 1
	}
	return event.UpdateMessage(discord.MessageUpdate{
		Embeds: &[]discord.Embed{
			embed.Queue(item[1], queue.Type.String(), queue.Tracks, page),
		},
		Components: &[]discord.ContainerComponent{embed.QueueButtons()},
	})
}
func index(b *Bot, event *events.ComponentInteractionCreate) error {

	return event.CreateMessage(discord.MessageCreate{
		Embeds: []discord.Embed{
			embed.HelpIndex(),
		},
		Components: []discord.ContainerComponent{embed.HelpComponent()},
	})
}
func setting(b *Bot, event *events.ComponentInteractionCreate) error {

	return event.CreateMessage(discord.MessageCreate{
		Embeds: []discord.Embed{
			embed.HelpSetting(),
		},
		Components: []discord.ContainerComponent{embed.HelpComponent()},
	})
}
func music(b *Bot, event *events.ComponentInteractionCreate) error {

	return event.CreateMessage(discord.MessageCreate{
		Embeds: []discord.Embed{
			embed.HelpMusic(),
		},
		Components: []discord.ContainerComponent{embed.HelpComponent()},
	})
}
