package listeners

import (
	"log"

	"github.com/disgoorg/disgo/events"

	"utils/internal/core"
)

func init() {
	core.AddListener(func(ctx *core.Context, event *events.Ready) {
		log.Printf("Logged in as %s#%s", event.User.Username, event.User.Discriminator)

		_, err := ctx.Client.Rest.SetGlobalCommands(event.User.ID, core.GetCommands())
		if err != nil {
			log.Printf("failed to register commands: %v", err)
			return
		}

		log.Print("Registered commands")
	})
}
