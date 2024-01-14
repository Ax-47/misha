package bot

import (
	"0x47/misha/embed"
	"context"
	"strconv"

	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/disgo/events"
	"github.com/disgoorg/disgolink/v3/disgolink"
	"github.com/disgoorg/disgolink/v3/lavalink"
)

func (b *Bot) settingCache(player disgolink.Player, guildId string, event *events.ApplicationCommandInteractionCreate) error {
	setting := b.Cache.GetCache(guildId)
	if len(setting) == 0 {
		return nil
	}
	volume, err := strconv.Atoi(setting["volume"])
	if err != nil {
		return err
	}
	if err := player.Update(context.TODO(), lavalink.WithVolume(volume)); err != nil {
		return event.CreateMessage(discord.MessageCreate{
			Embeds: []discord.Embed{
				embed.ErrorOther(),
			},
		})
	}
	return nil
}
