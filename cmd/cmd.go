package cmd

import (
	"log"
	"misha/extensions"
)

type Cmd struct {
	Ex *extensions.Ex
}

func (c *Cmd) Init(url, database string, colls []string) {
	var err error
	c.Ex = &extensions.Ex{}
	c.Ex.Init(url, database, colls)
	if err != nil {
		log.Fatalf("Invalid Language Handlers: %v", err)
	}
}
