package monitors

import (
	"github.com/c3sr/config"
	"github.com/c3sr/logger"
)

var (
	log = logger.New().WithField("pkg", "dlframework/predictor/cmd/monitors")
)

func init() {
	config.AfterInit(func() {
		log = logger.New().WithField("pkg", "dlframework/predictor/cmd/monitors")
	})
}
