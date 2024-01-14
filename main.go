package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"github.com/disgoorg/disgo"
	"github.com/disgoorg/disgo/bot"
	"github.com/disgoorg/disgo/cache"
	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/disgo/events"
	"github.com/disgoorg/disgo/gateway"
	"github.com/disgoorg/disgolink/v3/disgolink"

	"github.com/joho/godotenv"

	sbot "0x47/misha/bot"
)

func init() {
	err := godotenv.Load(".env")

	if err != nil {
		slog.Error("Error loading .env file")
	}
}
func main() {
	slog.Info("starting disgo example...")
	slog.Info("disgo version", slog.String("version", disgo.Version))
	slog.Info("disgolink version: ", slog.String("version", disgolink.Version))

	var (
		Token         = os.Getenv("TOKEN")
		NodeName      = os.Getenv("NODE_NAME")
		NodeAddress   = os.Getenv("NODE_ADDRESS")
		NodePassword  = os.Getenv("NODE_PASSWORD")
		NodeSecure, _ = strconv.ParseBool(os.Getenv("NODE_SECURE"))
		RedisAddress  = os.Getenv("REDIS_ADDRESS")
		RedisPassword = os.Getenv("REDIS_PASSWORD")
	)
	b := sbot.NewBot(RedisAddress, RedisPassword)
	client, err := disgo.New(Token,
		bot.WithGatewayConfigOpts(
			gateway.WithIntents(gateway.IntentGuilds, gateway.IntentGuildVoiceStates),
		),
		bot.WithCacheConfigOpts(
			cache.WithCaches(cache.FlagVoiceStates),
		),
		bot.WithEventListenerFunc(b.OnApplicationCommand),
		bot.WithEventListenerFunc(b.OnVoiceStateUpdate),
		bot.WithEventListenerFunc(b.OnVoiceServerUpdate),
		bot.WithEventListenerFunc(b.ButtonsEvent),
	)
	if err != nil {
		slog.Error("error while building disgo client", slog.Any("err", err))
		os.Exit(1)
	}
	b.Client = client
	registerCommands(client)
	b.Lavalink = disgolink.New(client.ApplicationID(),
		disgolink.WithListenerFunc(b.OnPlayerPause),
		disgolink.WithListenerFunc(b.OnPlayerResume),
		disgolink.WithListenerFunc(b.OnTrackStart),
		disgolink.WithListenerFunc(b.OnTrackEnd),
		disgolink.WithListenerFunc(b.OnTrackException),
		disgolink.WithListenerFunc(b.OnTrackStuck),
		disgolink.WithListenerFunc(b.OnWebSocketClosed),
		disgolink.WithListenerFunc(b.OnUnknownEvent),
	)
	b.Handlers = map[string]func(event *events.ApplicationCommandInteractionCreate, data discord.SlashCommandInteractionData) error{
		"play":        b.Play,
		"pause":       b.Pause,
		"now-playing": b.NowPlaying,
		"stop":        b.Stop,
		"players":     b.Players,
		"queue":       b.Queue,
		"clear-queue": b.ClearQueue,
		"loop":        b.Loop,
		"shuffle":     b.Shuffle,
		"seek":        b.Seek,
		"volume":      b.Volume,
		"skip":        b.Skip,
		"bass-boost":  b.BassBoost,
		"disconnect":  b.Disconnect,
		"autoplay":    b.Autoplay,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err = client.OpenGateway(ctx); err != nil {
		slog.Error("failed to open gateway", slog.Any("err", err))
		os.Exit(1)
	}
	defer client.Close(context.TODO())

	node, err := b.Lavalink.AddNode(ctx, disgolink.NodeConfig{
		Name:     NodeName,
		Address:  NodeAddress,
		Password: NodePassword,
		Secure:   NodeSecure,
	})
	if err != nil {
		slog.Error("failed to add node", slog.Any("err", err))
		os.Exit(1)
	}
	version, err := node.Version(ctx)
	if err != nil {
		slog.Error("failed to get node version", slog.Any("err", err))
		os.Exit(1)
	}

	slog.Info("DisGo example is now running. Press CTRL-C to exit.", slog.String("node_version", version), slog.String("node_session_id", node.SessionID()))
	s := make(chan os.Signal, 1)
	signal.Notify(s, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-s
}
