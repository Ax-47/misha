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
	MusicCommands struct {
		Errors struct {
			PlayerNotFound        string
			NotfoundTrack         string
			TracksInQueueNotFound string
			NotFoundTrackPlaying  string
			UserWasNotJoin        string
			QueueLessThanTwo      string
			UserNotInTheRoom      string
		}
		PlayTrack    string
		PlayPlaylist string
		Queue        string
		Skip         string
		Stop         string
		Shuffle      string
		Pause        string
		ClearQueue   string
		Loop         string
		Seek         string
		Swap         string
		Remove       string
		AutoPlay     string
	}
	FilterCommands struct {
		Errors struct {
			TimescaleErrorInput string
		}
	}
}
