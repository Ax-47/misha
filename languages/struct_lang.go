package languages

type Lang struct {
	Help struct {
		Pages struct {
			Index struct {
				Title       string `json:"title"`
				Description string `json:"description"`
				Image       string `json:"Image"`
			}
			Setup struct {
				Title       string `json:"title"`
				Description string `json:"description"`
			}
			Music struct {
				Title       string `json:"title"`
				Description string `json:"description"`
			}
		}
		Components struct {
			Placeholder string `json:"placeholder"`
			Options     struct {
				Index  string `json:"index"`
				Setup  string `json:"setup"`
				Music  string `json:"music"`
				Filter string `json:"filter"`
			}
		}
		Error struct {
			SelecterIsntAuthor string `json:"selecter_isnt_author"`
		}
	}
}
