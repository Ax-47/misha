package setup

import (
	"misha/languages"

	"github.com/bwmarrin/discordgo"
)

func Help(l map[string]languages.Lang, s *discordgo.Session, i *discordgo.InteractionCreate) {

	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Embeds: []*discordgo.MessageEmbed{
				{
					Color:       0xff4700,
					Title:       l[i.Locale.String()].Help.Pages.Index.Title,
					Description: l[i.Locale.String()].Help.Pages.Index.Description,
					Image: &discordgo.MessageEmbedImage{
						URL: l[i.Locale.String()].Help.Pages.Index.Image,
					},
				},
			},
			Components: []discordgo.MessageComponent{
				discordgo.ActionsRow{
					Components: []discordgo.MessageComponent{
						discordgo.SelectMenu{
							// Select menu, as other components, must have a customID, so we set it to this value.
							CustomID:    "select",
							Placeholder: l[i.Locale.String()].Help.Components.Placeholder,
							Options: []discordgo.SelectMenuOption{
								{
									Label: "Index",
									// As with components, this things must have their own unique "id" to identify which is which.
									// In this case such id is Value field.
									Value: "index",
									Emoji: discordgo.ComponentEmoji{
										Name: "🦦",
									},
									// You can also make it a default option, but in this case we won't.
									Default:     false,
									Description: l[i.Locale.String()].Help.Components.Options.Index,
								},
								{
									Label: "Setup",
									Value: "setup",
									Emoji: discordgo.ComponentEmoji{
										Name: "🟨",
									},
									Description: l[i.Locale.String()].Help.Components.Options.Setup,
								},
								{
									Label: "Music",
									Value: "music",
									Emoji: discordgo.ComponentEmoji{
										Name: "🐍",
									},
									Description: l[i.Locale.String()].Help.Components.Options.Music,
								},
							},
						}},
				},
			}}})
}

func Help_Component(l map[string]languages.Lang, s *discordgo.Session, i *discordgo.InteractionCreate) {
	var response *discordgo.InteractionResponse

	data := i.MessageComponentData()

	switch data.Values[0] {
	case "index":
		response = &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "This is the way.",
				Flags:   discordgo.MessageFlagsEphemeral,
			},
		}
	default:
		response = &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "It is not the way to go.",
				Flags:   discordgo.MessageFlagsEphemeral,
			},
		}
	}
	err := s.InteractionRespond(i.Interaction, response)
	if err != nil {
		panic(err)
	}

}
