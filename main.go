package main

import (
	"context"
	"fmt"
	"log"
	"misha/config"
	"misha/extensions"
	"misha/lava"
	"os"
	"os/signal"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/disgoorg/disgolink/v2/disgolink"
	"github.com/disgoorg/snowflake/v2"
	"github.com/servusdei2018/shards"
)

var (
	s   *shards.Manager
	con *config.Config
	Ex  *extensions.Ex
)

func init() {
	fmt.Print("\033[32m")
}
func init() {
	var err error
	con, err = config.Config_init()
	if err != nil {
		log.Fatalf("Invalid Config: %v", err)
	}
}

func init() {
	var err error
	s, err = shards.New("Bot " + con.Discord.Token)
	if err != nil {
		log.Fatalf("Invalid bot parameters: %v", err)
	}
}

func init() {
	Ex = &extensions.Ex{}
	Ex.Init(con)
	s.AddHandler(Handlers)
	s.AddHandler(EventHandler)
}

func main() {

	s.AddHandler(func(s *discordgo.Session, r *discordgo.Ready) {
		fmt.Print("\033[32m")
		log.Printf("Logged in as: %v#%v", s.State.User.Username, s.State.User.Discriminator)
		log.Printf("Author: %v", con.Info.Author)
		log.Printf("Version: %v", con.Info.Version)

		s.State.TrackVoice = true

		Ex.Bot.Queues = &lava.QueueManager{
			Queues:   make(map[string]*lava.Queue),
			Autoplay: make(map[string]bool),
			Cache:    make(map[string]string),
		}
		Ex.Bot.Lavalink = disgolink.New(snowflake.MustParse(s.State.User.ID),

			disgolink.WithListenerFunc(Ex.Bot.OnTrackEnd),
			disgolink.WithListenerFunc(Ex.Bot.OnTrackException),
		)
		Ex.Bot.S = s
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		nc := disgolink.NodeConfig{
			Name:     con.Lavalink.Name,
			Address:  con.Lavalink.Address,
			Password: con.Lavalink.Password,
			Secure:   con.Lavalink.Https,
		}
		node, err := Ex.Bot.Lavalink.AddNode(ctx, nc)
		if err != nil {
			log.Fatal(err)
		}
		version, err := node.Version(ctx)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("node version: %s", version)
	})

	s.AddHandler(Ex.Bot.OnVoiceStateUpdate)
	s.AddHandler(Ex.Bot.OnVoiceServerUpdate)

	log.Println("Adding commands...")
	registeredCommands := make([]*discordgo.ApplicationCommand, len(Commands))
	for i, v := range Commands {
		err := s.ApplicationCommandCreate(con.Discord.Guild, v)
		if err != nil {
			log.Panicf("Cannot create '%v' command: %v", v.Name, err)
		}
		registeredCommands[i] = v
	}
	s.RegisterIntent(discordgo.IntentsAll)
	err := s.Start()
	if err != nil {
		log.Println(err)
	}

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	log.Println("Press Ctrl+C to exit")
	<-stop
	log.Println("Gracefully shutting down.")
}
