package extensions

import (
	d "misha/database"
	l "misha/lang"

	"misha/lava"

	"github.com/bwmarrin/discordgo"
)

type Ex struct {
	languages map[string]l.Lang
	DB        *d.Database
	Bot       lava.Bot
}

func (c *Ex) Init(url, database string, colls []string, s *discordgo.Session, name, address, password string, https bool) error {
	var (
		err error
		con chan error
	)
	con = make(chan error)
	db := d.Database{}
	c.DB = &db
	go func(con chan error) {
		err = c.DB.Init(url, database, colls)
		if err != nil {
			con <- err
			return
		}
		con <- nil

	}(con)
	c.languages = make(map[string]l.Lang, 4)

	c.languages, err = l.Lang_init()
	if err != nil {
		return err
	} else if err = <-con; err != nil {
		return err
	}

	return nil
}
func (c *Ex) Lang(lang string) l.Lang {
	switch lang {
	case "Thai", "English, US", "Russian", "Japon":
	default:
		lang = "English, US"
	}

	return c.languages[lang]
}
