package cmd

import (
	"log"
	"misha/extensions"
)

type Cmd struct {
	Ex *extensions.Ex
}

func (c *Cmd) Init() {
	var err error
	c.Ex = &extensions.Ex{}
	c.Ex.Init()
	if err != nil {
		log.Fatalf("Invalid Language Handlers: %v", err)
	}
}
