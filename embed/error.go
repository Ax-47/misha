package embed

import (
	"github.com/disgoorg/disgo/discord"
)

func Error(err, description string) discord.Embed {
	return discord.Embed{
		Title:       err,
		Description: description,
		Color:       0x0,
	}
}
func ErrorNotFoundPlayer() discord.Embed {
	return discord.Embed{
		Title:       "ไม่มีplayer",
		Description: "คุณเน่คุณเปิดเพลงก่อนสิ",
		Color:       0x0,
	}
}
func ErrorCanNotConnectVC() discord.Embed {
	return discord.Embed{
		Title:       "mishaไม่สามารถเข้าvoice chat",
		Description: "mishaไม่มีสิทธ์เข้าห้อง gimme permission",
		Color:       0x0,
	}
}
func ErrorOther() discord.Embed {
	return discord.Embed{
		Title:       "Error.EXE!",
		Description: "mishaติดbug /report ให้sr. mishaทราบปัญหา",
		Color:       0x0,
	}
}
func ErrorYouAreNotVC() discord.Embed {
	return discord.Embed{
		Title:       "คุณไม่ได้อยู่ในvoice chat",
		Description: "คุณจำเป็นต้องอยู่ในvoice chatถึงสามารถใช้คำสั่งนี้ได้",
		Color:       0x0,
	}
}
func ErrorYouHaveNotPlay() discord.Embed {
	return discord.Embed{
		Title:       "คุณยังไม่ได้เล่นเพลงเลยนะ",
		Description: "เล่นเพลงก่อนสิ",
		Color:       0x0,
	}
}
func ErrorQueueIsNil() discord.Embed {
	return discord.Embed{
		Title:       "ไม่มีQueueต่อไปจ้า",
		Description: "เล่นเพลงต่อสิ",
		Color:       0x0,
	}
}
func ErrorOverLimitVoice() discord.Embed {
	return discord.Embed{
		Title:       "คุณเร่งดังเกิน1000",
		Description: "อย่าเร่งเกิน1000",
		Color:       0x0,
	}
}
func ErrorUnderZeroVoice() discord.Embed {
	return discord.Embed{
		Title:       "คุณปรับเสียงตำกว่า0",
		Description: "อย่าปรับเสียงตำกว่า0",
		Color:       0x0,
	}
}
func ErrorItIsNotYours() discord.Embed {
	return discord.Embed{
		Title:       "ใช้ของคุณ",
		Description: "`/queue`",
		Color:       0x0,
	}
}
