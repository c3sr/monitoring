package monitors

import (
	"github.com/k0kubun/pp"
	"github.com/c3sr/config"
	"github.com/c3sr/vipertags"
)

type monitorsConfig struct {
	Monitors []string      `json:"monitors" config:"monitors"`
	done     chan struct{} `json:"-" config:"-"`
}

var (
	Config = &monitorsConfig{
		done: make(chan struct{}),
	}
	DefaultMonitors = []string{
		"pprof",
		// "tracing",
		"expvar",
		"gops",
		"memory",
	}
)

func (monitorsConfig) ConfigName() string {
	return "monitors"
}

func (a *monitorsConfig) SetDefaults() {
	vipertags.SetDefaults(a)
}

func (a *monitorsConfig) Read() {
	defer close(a.done)
	vipertags.Fill(a)
	if len(a.Monitors) == 0 {
		a.Monitors = DefaultMonitors
	}
}

func (c monitorsConfig) Wait() {
	<-c.done
}

func (c monitorsConfig) String() string {
	return pp.Sprintln(c)
}

func (c monitorsConfig) Debug() {
	log.Debug("monitors Config = ", c)
}

func init() {
	config.Register(Config)
}
