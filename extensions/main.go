package extensions

import (
	d "misha/database"
	l "misha/lang"

	"misha/lava"

	"github.com/bwmarrin/discordgo"
)

type Ex struct {
	languages map[string]l.Lang
	DB        d.Database
	Bot       lava.Bot
}

func (c *Ex) Init(url, database string, colls []string, s *discordgo.Session, name, address, password string, https bool) error {
	var err error
	c.DB = d.Database{}
	err = c.DB.Init(url, database, colls)
	if err != nil {
		return err
	}

	if err != nil {
		return err
	}

	if err != nil {
		return err
	}
	c.languages, err = l.Lang_init()
	return err

}
func (c *Ex) Lang(lang string) l.Lang {
	switch lang {
	case "Thai", "English, US", "Russian", "Japon":
	default:
		lang = "English, US"
	}

	return c.languages[lang]
}
