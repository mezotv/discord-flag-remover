package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/disgoorg/disgo"
	"github.com/disgoorg/disgo/bot"
	"github.com/disgoorg/disgo/gateway"
	"github.com/mezotv/discord-flag-remover/config"
	botEvents "github.com/mezotv/discord-flag-remover/events"
)

func main() {
	config.Parse()

	slog.Info("disgo version", slog.String("version", disgo.Version))

	client, err := disgo.New(config.Conf.Token,
		bot.WithGatewayConfigOpts(
			gateway.WithIntents(
				gateway.IntentGuildMessageReactions,
			),
		),
		bot.WithEventListenerFunc(botEvents.OnMessageReactionAdd),
	)
	if err != nil {
		slog.Error("error while building disgo", slog.Any("err", err))
		return
	}

	defer client.Close(context.Background())

	if err = client.OpenGateway(context.Background()); err != nil {
		slog.Error("errors while connecting to gateway", slog.Any("err", err))
		return
	}

	slog.Info("Discord Flag Remover is now running. Press CTRL-C to exit.")
	s := make(chan os.Signal, 1)
	signal.Notify(s, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-s
}
