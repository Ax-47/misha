package music

import (
	"fmt"
	"misha/extensions"
	"strconv"
	"strings"

	"github.com/bwmarrin/discordgo"
)

func HandlerComponentsQueuePrevious(c *extensions.Ex, s *discordgo.Session, i *discordgo.InteractionCreate) {
	item := strings.Split(i.Message.Embeds[0].Footer.Text, " | ")
	if i.Member.User.ID != item[1] {
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,

			Data: &discordgo.InteractionResponseData{
				Flags:   discordgo.MessageFlagsEphemeral,
				Content: c.Lang(i.Locale.String()).Help.Error.SelecterIsntAuthor,
			},
		})
		return
	}
	var (
		page int
		err  error
	)
	it := strings.Split(strings.TrimPrefix(item[0], "page "), "/")
	page, err = strconv.Atoi(it[0])
	if err != nil {
		fmt.Println("Error during conversion")
		return
	}

	err = s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseUpdateMessage,
		Data: &discordgo.InteractionResponseData{
			Embeds: []*discordgo.MessageEmbed{embedQueue(page-1, c.Bot.Queues.Get(i.GuildID), i.Member.User.ID)},
			Components: []discordgo.MessageComponent{
				discordgo.ActionsRow{
					Components: []discordgo.MessageComponent{
						discordgo.Button{
							Label: "◀",
							Style: discordgo.SuccessButton,

							CustomID: "re",
						},
						discordgo.Button{
							Label: "▶",
							Style: discordgo.SuccessButton,

							CustomID: "next",
						},
						discordgo.Button{
							Label: "🗙",
							Style: discordgo.DangerButton,

							CustomID: "close",
						},
					},
				},
			},
		},
	})
	if err != nil {
		panic(err)
	}

}
func HandlerComponentsQueueNext(c *extensions.Ex, s *discordgo.Session, i *discordgo.InteractionCreate) {
	item := strings.Split(i.Message.Embeds[0].Footer.Text, " | ")
	if i.Member.User.ID != item[1] {
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,

			Data: &discordgo.InteractionResponseData{
				Flags:   discordgo.MessageFlagsEphemeral,
				Content: c.Lang(i.Locale.String()).Help.Error.SelecterIsntAuthor,
			},
		})
		return
	}
	var (
		page int
		err  error
	)

	it := strings.Split(strings.TrimPrefix(item[0], "page "), "/")
	page, err = strconv.Atoi(it[0])
	if err != nil {
		fmt.Println("Error during conversion")
		return
	}
	err = s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseUpdateMessage,
		Data: &discordgo.InteractionResponseData{
			Embeds: []*discordgo.MessageEmbed{embedQueue(page+1, c.Bot.Queues.Get(i.GuildID), i.Member.User.ID)},
			Components: []discordgo.MessageComponent{
				discordgo.ActionsRow{
					Components: []discordgo.MessageComponent{
						discordgo.Button{
							Label: "◀",
							Style: discordgo.SuccessButton,

							CustomID: "re",
						},
						discordgo.Button{
							Label: "▶",
							Style: discordgo.SuccessButton,

							CustomID: "next",
						},
						discordgo.Button{
							Label: "🗙",
							Style: discordgo.DangerButton,

							CustomID: "close",
						},
					},
				},
			},
		},
	})
	if err != nil {
		panic(err)
	}

}
