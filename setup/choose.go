package setup

import (
	"misha/extensions"

	"github.com/bwmarrin/discordgo"
)

func SetupChoose(c *extensions.Ex, s *discordgo.Session, i *discordgo.InteractionCreate) {

	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Embeds: []*discordgo.MessageEmbed{create_embed_index(c, i.Locale.String(), i.Member.User.ID)},
		}})
}
