package setup

import (
	"context"
	"misha/extensions"

	"github.com/bwmarrin/discordgo"
	"go.mongodb.org/mongo-driver/bson"
)

func createEmbedChooseSetup(c *extensions.Ex, lang string) *discordgo.MessageEmbed {
	return &discordgo.MessageEmbed{
		Color: 0xff4700,
		Title: c.Lang(lang).Help.Pages.Index.Title,
	}
}
func SetupChoose(c *extensions.Ex, s *discordgo.Session, i *discordgo.InteractionCreate) {

	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Flags:   discordgo.MessageFlagsEphemeral,
			Content: "อย่าลืมแก้นะ นี้เว็บ",
		}})
	M, err := s.FollowupMessageCreate(i.Interaction, true, &discordgo.WebhookParams{
		Embeds: []*discordgo.MessageEmbed{createEmbedChooseSetup(c, i.Locale.String())},
	})
	if err != nil {
		panic(err)
	}
	c.DB.Colls["guilds"].UpdateOne(context.TODO(), bson.M{"GuildID": i.GuildID}, bson.M{"$set": bson.M{"Choose.MassageID": M.ID}})
}
