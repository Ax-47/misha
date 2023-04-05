package cmd

import (
	"log"
	"misha/extensions"

	"github.com/bwmarrin/discordgo"
)

type Cmd struct {
	Ex *extensions.Ex
}

func (c *Cmd) Init(url, database string, colls []string, s *discordgo.Session, name, address, password string, https bool) {
	var err error
	c.Ex = &extensions.Ex{}
	c.Ex.Init(url, database, colls, s, name, address, password, https)
	if err != nil {
		log.Fatalf("Invalid Language Handlers: %v", err)
	}
}
