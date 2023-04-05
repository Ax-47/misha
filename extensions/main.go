package extensions

import (
	l "misha/languages"
)

type Ex struct {
	Languages map[string]l.Lang
}

func (c *Ex) Init() error {
	var err error
	c.Languages, err = l.Lang_init()

	return err

}
func (c *Ex) Lang(lang string) l.Lang {
	switch lang {
	case "Thai", "English, US", "Russian", "Japon":
	default:
		lang = "English, US"
	}

	return c.Languages[lang]
}
