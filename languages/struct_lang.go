package languages

type Lang struct {
	Help struct {
		Pages struct {
			Index struct {
				Title       string `json:"title"`
				Description string `json:"description"`
				Image       string `json:"Image"`
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
	}
}

type Music struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}
type Setup struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}
