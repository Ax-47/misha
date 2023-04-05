package models

type GuildDoc struct {
	GuildID string
	Choose  struct {
		MassageID string
		Selection map[string]string
	}
}
