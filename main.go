package main

import (
	"context"
	"log"
	"misha/cmd"
	"misha/lava"
	"os"
	"os/signal"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/disgoorg/disgolink/v2/disgolink"
	"github.com/disgoorg/snowflake/v2"
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
	c.Init(con.Database.Url, con.Database.Database, con.Database.Collection, s, con.Lavalink.Name, con.Lavalink.Address, con.Lavalink.Password, con.Lavalink.Https)

	s.AddHandler(Handlers)
}
func init() {
	s.AddHandler(EventHandler)
}

func main() {
	s.AddHandler(func(s *discordgo.Session, r *discordgo.Ready) {
		log.Printf("Logged in as: %v#%v", s.State.User.Username, s.State.User.Discriminator)
	})
	s.State.TrackVoice = true
	s.Identify.Intents = discordgo.IntentsAll
	c.Ex.Bot.Queues = &lava.QueueManager{
		Queues: make(map[string]*lava.Queue),
	}
	c.Ex.Bot.S = s
	s.AddHandler(c.Ex.Bot.OnVoiceStateUpdate)
	s.AddHandler(c.Ex.Bot.OnVoiceServerUpdate)
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
	c.Ex.Bot.Lavalink = disgolink.New(snowflake.MustParse(s.State.User.ID),
		disgolink.WithListenerFunc(c.Ex.Bot.OnPlayerPause),
		disgolink.WithListenerFunc(c.Ex.Bot.OnPlayerResume),
		disgolink.WithListenerFunc(c.Ex.Bot.OnTrackStart),
		disgolink.WithListenerFunc(c.Ex.Bot.OnTrackEnd),
		disgolink.WithListenerFunc(c.Ex.Bot.OnTrackException),
		disgolink.WithListenerFunc(c.Ex.Bot.OnTrackStuck),
		disgolink.WithListenerFunc(c.Ex.Bot.OnWebSocketClosed),
	)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	nc := disgolink.NodeConfig{
		Name:     con.Lavalink.Name,
		Address:  con.Lavalink.Address,
		Password: con.Lavalink.Password,
		Secure:   con.Lavalink.Https,
	}
	node, err := c.Ex.Bot.Lavalink.AddNode(ctx, nc)
	if err != nil {
		log.Fatal(err)
	}
	version, err := node.Version(ctx)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("node version: %s", version)

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
