package extensions

import (
	d "misha/database"
	l "misha/languages"
)

type Ex struct {
	languages map[string]l.Lang
	DB        d.Database
}

func (c *Ex) Init(url, database string, colls []string) error {
	var err error
	c.DB = d.Database{}
	err = c.DB.Init(url, database, colls)
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
