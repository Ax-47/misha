package setup

import (
	"misha/extensions"

	"github.com/bwmarrin/discordgo"
)

func create_embed_index(c *extensions.Ex, lang, ID string) *discordgo.MessageEmbed {
	return &discordgo.MessageEmbed{
		Color:       0xff4700,
		Title:       c.Lang(lang).Help.Pages.Index.Title,
		Description: c.Lang(lang).Help.Pages.Index.Description,
		Image: &discordgo.MessageEmbedImage{
			URL: c.Lang(lang).Help.Pages.Index.Image,
		},
		Footer: &discordgo.MessageEmbedFooter{
			Text: ID,
		},
	}
}
func create_embed_setup(c *extensions.Ex, lang, ID string) *discordgo.MessageEmbed {
	return &discordgo.MessageEmbed{
		Color:       0xff4700,
		Title:       c.Lang(lang).Help.Pages.Setup.Title,
		Description: c.Lang(lang).Help.Pages.Setup.Description,
		Footer: &discordgo.MessageEmbedFooter{
			Text: ID,
		},
	}
}
func create_embed_music(c *extensions.Ex, lang, ID string) *discordgo.MessageEmbed {
	return &discordgo.MessageEmbed{
		Color:       0xff4700,
		Title:       c.Lang(lang).Help.Pages.Music.Title,
		Description: c.Lang(lang).Help.Pages.Music.Description,
		Footer: &discordgo.MessageEmbedFooter{
			Text: ID,
		},
	}
}
func create_Components(c *extensions.Ex, lang string) []discordgo.MessageComponent {
	return []discordgo.MessageComponent{
		discordgo.SelectMenu{
			// Select menu, as other components, must have a customID, so we set it to this value.
			CustomID:    "select",
			Placeholder: c.Lang(lang).Help.Components.Placeholder,
			Options: []discordgo.SelectMenuOption{
				{
					Label: "Index",
					// As with components, this things must have their own unique "id" to identify which is which.
					// In this case such id is Value field.
					Value: "index",
					Emoji: discordgo.ComponentEmoji{
						Name: "🏠",
					},
					// You can also make it a default option, but in this case we won't.
					Default:     false,
					Description: c.Lang(lang).Help.Components.Options.Index,
				},
				{
					Label: "Setup",
					Value: "setup",
					Emoji: discordgo.ComponentEmoji{
						Name: "⚙️",
					},
					Description: c.Lang(lang).Help.Components.Options.Setup,
				},
				{
					Label: "Music",
					Value: "music",
					Emoji: discordgo.ComponentEmoji{
						Name: "🎵",
					},
					Description: c.Lang(lang).Help.Components.Options.Music,
				},
			},
		}}
}
func Help(c *extensions.Ex, s *discordgo.Session, i *discordgo.InteractionCreate) {

	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Embeds: []*discordgo.MessageEmbed{create_embed_index(c, i.Locale.String(), i.Member.User.ID)},
			Components: []discordgo.MessageComponent{
				discordgo.ActionsRow{
					Components: create_Components(c, i.Locale.String()),
				},
			}}})
}

func Help_Component(c *extensions.Ex, s *discordgo.Session, i *discordgo.InteractionCreate) {
	var Embeds *discordgo.MessageEmbed

	data := i.MessageComponentData()
	//requester
	//select
	if i.Member.User.ID != i.Message.Embeds[0].Footer.Text {
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,

			Data: &discordgo.InteractionResponseData{
				Flags:   discordgo.MessageFlagsEphemeral,
				Content: c.Lang(i.Locale.String()).Help.Error.SelecterIsntAuthor,
			},
		})
		return
	}
	switch data.Values[0] {
	case "index":

		Embeds = create_embed_index(c, i.Locale.String(), i.Member.User.ID)
	case "setup":
		Embeds = create_embed_setup(c, i.Locale.String(), i.Member.User.ID)
	case "music":
		Embeds = create_embed_music(c, i.Locale.String(), i.Member.User.ID)
	default:
		Embeds = create_embed_index(c, i.Locale.String(), i.Member.User.ID)
	}

	err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseUpdateMessage,
		Data: &discordgo.InteractionResponseData{
			Embeds: []*discordgo.MessageEmbed{Embeds},
			Components: []discordgo.MessageComponent{
				discordgo.ActionsRow{
					Components: create_Components(c, i.Locale.String()),
				},
			},
		},
	})
	if err != nil {
		panic(err)
	}

}
