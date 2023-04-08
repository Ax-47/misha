package languages

import (
	"encoding/json"
	"io/ioutil"
)

var (
	files = map[string]string{
		"Thai":        "languages/TH.json",
		"English, US": "languages/EN_USA.json",
		"Russian":     "languages/RU.json",
		"Japon":       "languages/JP.json",
	}
)

type L struct {
}

func Lang_init() (map[string]Lang, error) {
	languages := make(map[string]Lang)
	for language, file := range files {

		buf, err := ioutil.ReadFile(file)
		if err != nil {
			return nil, err
		}

		lang := &Lang{}

		err = json.Unmarshal(buf, lang)
		if err != nil {
			return nil, err
		}
		languages[language] = *lang
	}
	return languages, nil
}
