package events

import (
	"fmt"
	"log/slog"
	"regexp"
	"slices"

	"github.com/disgoorg/disgo/events"
	"github.com/mezotv/discord-flag-remover/config"
)

func OnMessageReactionAdd(event *events.GuildMessageReactionAdd) {
	flagPattern := regexp.MustCompile(`[\x{1F1E6}-\x{1F1FF}]{2}`)

	if len(config.Conf.Settings.ChannelList) != 0 && !slices.Contains(config.Conf.Settings.ChannelList, event.ChannelID.String()) {
		slog.Info("Channel list is not empty and this channel is not in the list, returning...")
		return
	}

	if !flagPattern.MatchString(*event.Emoji.Name) {
		return
	}

	event.Client().Rest().RemoveAllReactionsForEmoji(event.ChannelID, event.MessageID, *event.Emoji.Name)
	slog.Info(fmt.Sprintf("Removed reaction (%s) from message %s send by %s in channel %s",
		*event.Emoji.Name, event.MessageID, event.UserID.String(), event.ChannelID.String()))
}
