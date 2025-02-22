package telemetry

import (
	"github.com/DataDog/KubeHound/pkg/config"
	"github.com/DataDog/KubeHound/pkg/telemetry/log"
	"github.com/DataDog/KubeHound/pkg/telemetry/profiler"
	"github.com/DataDog/KubeHound/pkg/telemetry/statsd"
	"github.com/DataDog/KubeHound/pkg/telemetry/tracer"
)

type TelemetryState struct {
	Enabled bool
}

// Initialize all telemetry required
// return client to enable clean shutdown
func Initialize(cfg *config.KubehoundConfig) (*TelemetryState, error) {
	if !cfg.Telemetry.Enabled {
		return &TelemetryState{}, nil
	}

	// profiling
	profiler.Initialize(cfg)
	//Tracing
	tracer.Initialize(cfg)
	// Metrics
	err := statsd.Setup(cfg.Telemetry.Statsd.URL)
	if err != nil {
		return &TelemetryState{
			Enabled: true,
		}, err
	}

	return &TelemetryState{
		Enabled: true,
	}, nil
}

func Shutdown(ts *TelemetryState) {
	if !ts.Enabled {
		return
	}

	//Profiling
	profiler.Shutdown()

	//Tracing
	tracer.Shutdown()
	// Metrics
	err := statsd.Flush()
	if err != nil {
		log.I.Warnf("Failed to flush statsd client: %v", err)
	}

	err = statsd.Close()
	if err != nil {
		log.I.Warnf("Failed to close statsd client: %v", err)
	}
}
