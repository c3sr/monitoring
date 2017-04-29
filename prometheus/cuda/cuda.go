package cuda

import (
	"errors"
	"sync"

	"github.com/prometheus/client_golang/prometheus"
	nvidiasmi "github.com/rai-project/nvidia-smi"
)

// see https://github.com/tankbusta/nvidia_exporter/blob/master/nvidia_exporter.go

type Exporter struct {
	mutex sync.RWMutex

	up     prometheus.Gauge
	gauges map[string]*prometheus.GaugeVec

	devices []nvidiasmi.GPU
}

// VecInfo stores the prometheus help and labels to
type VecInfo struct {
	help   string
	labels []string
}

var (
	DefaultNamespace = "cuda"

	gaugeMetrics = map[string]*VecInfo{
		"power_watts": &VecInfo{
			help:   "Power Usage of an NVIDIA GPU in Watts",
			labels: []string{"device_uuid", "device_name"},
		},
		"gpu_percent": &VecInfo{
			help:   "Percent of GPU Utilized",
			labels: []string{"device_uuid", "device_name"},
		},
		"memory_free": &VecInfo{
			help:   "Number of bytes free in the GPU Memory",
			labels: []string{"device_uuid", "device_name"},
		},
		"memory_total": &VecInfo{
			help:   "Total bytes of the GPU's memory",
			labels: []string{"device_uuid", "device_name"},
		},
		"memory_used": &VecInfo{
			help:   "Total number of bytes used in the GPU Memory",
			labels: []string{"device_uuid", "device_name"},
		},
		"memory_percent": &VecInfo{
			help:   "Percent of GPU Memory Utilized",
			labels: []string{"device_uuid", "device_name"},
		},
		"temperature_fahrenheit": &VecInfo{
			help:   "GPU Temperature in Fahrenheit",
			labels: []string{"device_uuid", "device_name"},
		},
		"temperature_celsius": &VecInfo{
			help:   "GPU Temperature in Celsius",
			labels: []string{"device_uuid", "device_name"},
		},
	}
)

func New() (*Exporter, error) {
	if !nvidiasmi.HasGPU {
		return nil, errors.New("no gpus were detected on the system")
	}
	exp := &Exporter{
		gauges: make(map[string]*prometheus.GaugeVec, len(gaugeMetrics)),
		up: prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace: DefaultNamespace,
			Name:      "up",
			Help:      "Were the CUDA queries successful?",
		}),
	}

	return exp, nil
}
