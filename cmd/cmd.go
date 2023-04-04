package cmd

import (
	"log"
	l "misha/languages"
)

type Cmd struct {
	languages map[string]l.Lang
}

func (c *Cmd) Init() {
	var err error
	c.languages, err = l.Lang_init()
	if err != nil {
		log.Fatalf("Invalid Language Handlers: %v", err)
	}
}
