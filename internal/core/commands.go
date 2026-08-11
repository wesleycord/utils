package core

import (
	"fmt"

	"github.com/disgoorg/disgo/bot"
	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/disgo/events"
)

type Command struct {
	Data   discord.ApplicationCommandCreate
	Handle func(ctx *Context, event *events.ApplicationCommandInteractionCreate) error
}

var commands = map[string]Command{}

func key(t discord.ApplicationCommandType, name string) string {
	return fmt.Sprintf("%d:%s", t, name)
}

func AddCommand(c Command) {
	commands[key(c.Data.Type(), c.Data.CommandName())] = c
}

func GetCommands() []discord.ApplicationCommandCreate {
	cmds := make([]discord.ApplicationCommandCreate, 0, len(commands))
	for _, c := range commands {
		cmds = append(cmds, c.Data)
	}
	return cmds
}

func OnCommandInteractionCreate(ctx *Context) bot.EventListener {
	return bot.NewListenerFunc(
		func(e *events.ApplicationCommandInteractionCreate) {
			data := e.Data

			c, ok := commands[key(data.Type(), data.CommandName())]
			if !ok {
				_ = e.CreateMessage(discord.MessageCreate{
					Content: "An error occurred - Command not found",
					Flags:   discord.MessageFlagEphemeral,
				})
				return
			}

			_ = c.Handle(ctx, e)
		},
	)
}
