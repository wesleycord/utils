package main

import (
	"context"
	"database/sql"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/disgoorg/disgo"
	"github.com/disgoorg/disgo/bot"
	"github.com/disgoorg/disgo/gateway"

	"github.com/joho/godotenv"
	_ "modernc.org/sqlite"

	_ "utils/commands"
	_ "utils/listeners"

	"utils/internal/core"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file was found")
	}

	db, err := sql.Open("sqlite", "utils.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	client, err := disgo.New(os.Getenv("DISCORD_TOKEN"),
		bot.WithGatewayConfigOpts(
			gateway.WithIntents(
				gateway.IntentsDefault,
			),
		),
	)
	if err != nil {
		panic(err)
	}

	ctx := &core.Context{
		Client: client,
		DB:     db,
	}

	client.AddEventListeners(core.OnCommandInteractionCreate(ctx))
	client.AddEventListeners(core.BuildListeners(ctx)...)

	defer client.Close(context.Background())

	if err = client.OpenGateway(context.Background()); err != nil {
		panic(err)
	}

	s := make(chan os.Signal, 1)
	signal.Notify(s, syscall.SIGINT, syscall.SIGTERM)
	<-s
}
