package core

import (
	"database/sql"

	"github.com/disgoorg/disgo/bot"
)

type Context struct {
	Client *bot.Client
	DB     *sql.DB
}
