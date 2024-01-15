package embed

import "github.com/disgoorg/disgo/discord"

func HelpIndex() discord.Embed {

	return discord.Embed{
		Title:       "Ax47 | Help",
		Description: "จ้า! mishaมาแล้วสหาย \nอยากรู้คำสั่งอะไรเลือกด้านล่างเลยจ้า donateให้เจ้าของbotได้ที่ [tipme](https://tipme.in.th/ax-47) \nปล.เจ้าของbotไปเซ็นร้านค้าเป็นร้อยแล้วววววว",
		Image: &discord.EmbedResource{
			URL: "https://cdn.discordapp.com/attachments/1092016880136503296/1092777250920861726/no_money.png",
		},
		Color: 0xff4700,
	}
}
func HelpSetting() discord.Embed {

	return discord.Embed{
		Title: "Ax47 | Setting Server",
		Color: 0xff4700,
	}
}
func HelpMusic() discord.Embed {

	return discord.Embed{
		Title: "Ax47 | Music",
		Color: 0xff4700,
	}
}

func HelpComponent() discord.ActionRowComponent {

	return discord.NewActionRow(discord.NewSuccessButton("🏠", "index"), discord.NewSuccessButton("⚙️", "setting"), discord.NewSuccessButton("🎵", "music"))
}
