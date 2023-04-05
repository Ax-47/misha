package main

import (
	"log"
	"misha/cmd"
	"os"
	"os/signal"

	"github.com/bwmarrin/discordgo"
)

var (
	s   *discordgo.Session
	con *Config
	c   cmd.Cmd
)

func init() {
	var err error
	con, err = Config_init()
	if err != nil {
		log.Fatalf("Invalid Config: %v", err)
	}
}

func init() {
	var err error
	s, err = discordgo.New("Bot " + con.Discord.Token)
	if err != nil {
		log.Fatalf("Invalid bot parameters: %v", err)
	}
}

func init() {
	c.Init(con.Database.Url, con.Database.Database, con.Database.Collection)

	s.AddHandler(CommandsHandler)
}
func init() {
	c.Init(con.Database.Url, con.Database.Database, con.Database.Collection)

	s.AddHandler(EventHandler)
}

func main() {
	s.AddHandler(func(s *discordgo.Session, r *discordgo.Ready) {
		log.Printf("Logged in as: %v#%v", s.State.User.Username, s.State.User.Discriminator)
	})

	err := s.Open()
	if err != nil {
		log.Fatalf("Cannot open the session: %v", err)
	}
	log.Println("Adding commands...")
	registeredCommands := make([]*discordgo.ApplicationCommand, len(Commands))
	for i, v := range Commands {
		cmd, err := s.ApplicationCommandCreate(s.State.User.ID, con.Discord.Guild, v)
		if err != nil {
			log.Panicf("Cannot create '%v' command: %v", v.Name, err)
		}
		registeredCommands[i] = cmd
	}

	defer s.Close()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	log.Println("Press Ctrl+C to exit")
	<-stop

	if con.Discord.Rmcmd {
		log.Println("Removing commands...")
		for _, v := range registeredCommands {
			err := s.ApplicationCommandDelete(s.State.User.ID, con.Discord.Guild, v.ID)
			if err != nil {
				log.Panicf("Cannot delete '%v' command: %v", v.Name, err)
			}
		}
	}

	log.Println("Gracefully shutting down.")
}
