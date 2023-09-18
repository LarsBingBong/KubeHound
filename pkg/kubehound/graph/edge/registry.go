package edge

import (
	"fmt"
	"sync"

	"github.com/DataDog/KubeHound/pkg/telemetry/log"
)

type RegistrationFlag uint8

const (
	RegisterDefault         RegistrationFlag = 1 << iota // Default edge
	RegisterGraphMutation                                // Edge can mutate the graph
	RegisterGraphDependency                              // Edge has a dependency on default/mutating edges
)

// Registry holds details of edges (i.e attacks) registered in KubeHound.
type Registry struct {
	mutating  map[string]Builder
	simple    map[string]Builder
	dependent map[string]DependentBuilder
}

// newRegistry creates a new registry instance. This should not be called directly.
func newRegistry() *Registry {
	r := &Registry{
		mutating:  make(map[string]Builder),
		simple:    make(map[string]Builder),
		dependent: make(map[string]DependentBuilder),
	}

	return r
}

// Registry singleton support
var registryInstance *Registry
var erOnce sync.Once

// Registered returns the edge registry singleton.
func Registered() *Registry {
	erOnce.Do(func() {
		registryInstance = newRegistry()
	})

	return registryInstance
}

// Mutating returns the map of registered mutating edge builders.
func (r *Registry) Mutating() map[string]Builder {
	return r.mutating
}

// Simple returns the map of registered edge builders.
func (r *Registry) Simple() map[string]Builder {
	return r.simple
}

// Dependent returns the map of registered edge builders with default edge dependencies.
func (r *Registry) Dependent() map[string]DependentBuilder {
	return r.dependent
}

// Verify verifies the integrity and consistency of the registry.
// Function should only be called once all edges have been registered via init() calls.
func (r *Registry) Verify() error {
	// Ensure all dependent edges have dependencies registered in mutating or default collections
	for name, builder := range r.dependent {
		for _, d := range builder.Dependencies() {
			_, depSimple := r.simple[d]
			_, depMutating := r.mutating[d]

			if !depSimple && !depMutating {
				return fmt.Errorf("unregistered dependency (%s) for dependent edge %s", d, name)
			}
		}
	}

	return nil
}

// Register loads the provided edge into the registry.
func Register(edge Builder, flags RegistrationFlag) {
	registry := Registered()
	switch {
	case flags&RegisterGraphMutation != 0:
		log.I.Debugf("Registering mutating edge builder %s -> %s", edge.Name(), edge.Label())
		if _, ok := registry.mutating[edge.Name()]; ok {
			log.I.Fatalf("edge name collision: %s", edge.Name())
		}

		registry.mutating[edge.Name()] = edge
	case flags&RegisterGraphDependency != 0:
		log.I.Debugf("Registering dependent edge builder %s -> %s", edge.Name(), edge.Label())
		if _, ok := registry.dependent[edge.Name()]; ok {
			log.I.Fatalf("edge name collision: %s", edge.Name())
		}

		dependent, ok := edge.(DependentBuilder)
		if !ok {
			log.I.Fatalf("dependent edge must implement DependentBuilder: %s", edge.Name())
		}

		registry.dependent[edge.Name()] = dependent
	default:
		log.I.Debugf("Registering default edge builder %s -> %s", edge.Name(), edge.Label())
		if _, ok := registry.simple[edge.Name()]; ok {
			log.I.Fatalf("edge name collision: %s", edge.Name())
		}

		registry.simple[edge.Name()] = edge
	}
}
