package commands

import (
	"fmt"
	"time"
	"utils/internal/core"

	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/disgo/events"
)

func init() {
	core.AddCommand(core.Command{
		Data: discord.SlashCommandCreate{
			Name:        "ping",
			Description: "Check the bot's latency",
		},
		Handle: func(ctx *core.Context, event *events.ApplicationCommandInteractionCreate) error {
			gatewayLatency := ctx.Client.Gateway.Latency()

			apiStart := time.Now()
			if err := event.CreateMessage(discord.MessageCreate{
				Content: "Pinging...",
			}); err != nil {
				return err
			}
			apiLatency := time.Since(apiStart)

			content := fmt.Sprintf(
				"🏓 **Pong!**\n"+
					"**Gateway:** `%dms`\n"+
					"**API:** `%dms`",
				gatewayLatency.Milliseconds(),
				apiLatency.Milliseconds(),
			)

			_, err := ctx.Client.Rest.UpdateInteractionResponse(
				event.ApplicationID(),
				event.Token(),
				discord.MessageUpdate{
					Content: &content,
				},
			)
			return err
		},
	})
}
