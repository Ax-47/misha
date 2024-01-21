package embed

import (
	"fmt"

	"github.com/disgoorg/disgo/discord"
)

func Filter(filter string) discord.Embed {
	return discord.Embed{
		Title: fmt.Sprintf("filter : %s", filter),
		Color: 0xff4700,
	}
}

func Tremolo(frequency, depth int) discord.Embed {
	return discord.Embed{
		Title: fmt.Sprintf("frequency : %d ,depth : %d", frequency, depth),
		Color: 0xff4700,
	}
}
func Timescale(speed, pitch, rate int) discord.Embed {
	return discord.Embed{
		Title: fmt.Sprintf("speed : %d ,pitch : %d ,rate : %d", speed, pitch, rate),
		Color: 0xff4700,
	}
}

func Eq(eq string) discord.Embed {
	return discord.Embed{
		Title: fmt.Sprintf("BassBoost : %s", eq),
		Color: 0xff4700,
	}
}
