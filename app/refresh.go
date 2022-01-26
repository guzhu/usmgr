package app

import (
	"time"

	"github.com/rs/zerolog/log"
)

func (app *App) refresh() {
	log.Debug().Msg("app: starting refresh loop")
	tick := time.NewTicker(refreshInterval)
	for {
		select {
		case <-tick.C:
			app.initHead()
			app.Application.Draw()
		}
	}
}
