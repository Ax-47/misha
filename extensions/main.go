package extensions

import (
	"misha/config"
	d "misha/database"
	l "misha/lang"
	"misha/lava"

	"github.com/zmb3/spotify/v2"
)

type Ex struct {
	languages map[string]l.Lang
	DB        *d.Database
	Bot       lava.Bot
	Spotify   *spotify.Client
}

func (c *Ex) Init(configed *config.Config) error {
	var (
		err error
		con chan error
	)
	con = make(chan error)
	db := d.Database{}
	c.DB = &db
	go func(con chan error) {
		err = c.DB.Init(configed.Database.Url, configed.Database.Database, configed.Database.Collection)
		if err != nil {
			con <- err
			return
		}
		con <- nil

	}(con)
	c.languages = make(map[string]l.Lang, 4)
	c.Spotify = Auth(configed.Spotify.Client, configed.Spotify.Secret)
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
