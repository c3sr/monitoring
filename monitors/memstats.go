package monitors

import (
	"github.com/gbbr/memstats"
	"github.com/c3sr/config"
)

func init() {
	config.AfterInit(func() {
		if !memberQ(Config.Monitors, "memstats") {
			return
		}
		go memstats.Serve()
	})
}
