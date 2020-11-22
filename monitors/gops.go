package monitors

import (
	"github.com/google/gops/agent"
	"github.com/c3sr/config"
)

func init() {
	config.AfterInit(func() {
		if !memberQ(Config.Monitors, "gops") {
			return
		}
		if err := agent.Listen(agent.Options{}); err != nil {
			log.WithError(err).Error("failed to listen for gops")
		}
	})
}
