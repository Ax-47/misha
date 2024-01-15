package bot

import (
	"0x47/misha/embed"

	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/disgo/events"
)

func (b *Bot) Help(event *events.ApplicationCommandInteractionCreate, data discord.SlashCommandInteractionData) error {

	return event.CreateMessage(discord.MessageCreate{
		Embeds: []discord.Embed{
			embed.HelpIndex(),
		},
		Components: []discord.ContainerComponent{embed.HelpComponent()},
	})
}
