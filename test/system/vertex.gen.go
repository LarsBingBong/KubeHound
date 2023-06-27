// PLEASE DO NOT EDIT
// THIS HAS BEEN GENERATED AUTOMATICALLY on 2023-06-26 18:29
//
// Generate it with "go generate ./..."
//
// currently support only:
// - nodes
// - pods
// - containers
// - volumes
//
// TODO: roles, rolebinding, clusterrole, clusterrolebindings

package system

import (
	"github.com/DataDog/KubeHound/pkg/kubehound/models/graph"
	"github.com/DataDog/KubeHound/pkg/kubehound/models/shared"
)

var expectedPods = map[string]graph.Pod{
	"impersonate-pod": {
		StoreID:                "",
		Name:                   "impersonate-pod",
		IsNamespaced:           true,
		Namespace:              "default",
		Compromised:            shared.CompromiseNone,
		ServiceAccount:         "impersonate-sa",
		shareProcessNamespace: false,
		Critical:               false,
	},
	"modload-pod": {
		StoreID:                "",
		Name:                   "modload-pod",
		IsNamespaced:           true,
		Namespace:              "default",
		Compromised:            shared.CompromiseNone,
		ServiceAccount:         "default",
		shareProcessNamespace: false,
		Critical:               false,
	},
	"netadmin-pod": {
		StoreID:                "",
		Name:                   "netadmin-pod",
		IsNamespaced:           true,
		Namespace:              "default",
		Compromised:            shared.CompromiseNone,
		ServiceAccount:         "default",
		shareProcessNamespace: false,
		Critical:               false,
	},
	"nsenter-pod": {
		StoreID:                "",
		Name:                   "nsenter-pod",
		IsNamespaced:           true,
		Namespace:              "default",
		Compromised:            shared.CompromiseNone,
		ServiceAccount:         "default",
		shareProcessNamespace: false,
		Critical:               false,
	},
	"pod-create-pod": {
		StoreID:                "",
		Name:                   "pod-create-pod",
		IsNamespaced:           true,
		Namespace:              "default",
		Compromised:            shared.CompromiseNone,
		ServiceAccount:         "pod-create-sa",
		shareProcessNamespace: false,
		Critical:               false,
	},
	"pod-patch-pod": {
		StoreID:                "",
		Name:                   "pod-patch-pod",
		IsNamespaced:           true,
		Namespace:              "default",
		Compromised:            shared.CompromiseNone,
		ServiceAccount:         "pod-patch-sa",
		shareProcessNamespace: false,
		Critical:               false,
	},
	"priv-pod": {
		StoreID:                "",
		Name:                   "priv-pod",
		IsNamespaced:           true,
		Namespace:              "default",
		Compromised:            shared.CompromiseNone,
		ServiceAccount:         "default",
		shareProcessNamespace: false,
		Critical:               false,
	},
	"rolebind-pod": {
		StoreID:                "",
		Name:                   "rolebind-pod",
		IsNamespaced:           true,
		Namespace:              "default",
		Compromised:            shared.CompromiseNone,
		ServiceAccount:         "rolebind-sa",
		shareProcessNamespace: false,
		Critical:               false,
	},
	"sharedps-pod": {
		StoreID:                "",
		Name:                   "sharedps-pod",
		IsNamespaced:           true,
		Namespace:              "default",
		Compromised:            shared.CompromiseNone,
		ServiceAccount:         "default",
		shareProcessNamespace: true,
		Critical:               false,
	},
	"tokenget-pod": {
		StoreID:                "",
		Name:                   "tokenget-pod",
		IsNamespaced:           true,
		Namespace:              "default",
		Compromised:            shared.CompromiseNone,
		ServiceAccount:         "tokenget-sa",
		shareProcessNamespace: false,
		Critical:               false,
	},
	"tokenlist-pod": {
		StoreID:                "",
		Name:                   "tokenlist-pod",
		IsNamespaced:           true,
		Namespace:              "default",
		Compromised:            shared.CompromiseNone,
		ServiceAccount:         "tokenlist-sa",
		shareProcessNamespace: false,
		Critical:               false,
	},
	"umh-core-pod": {
		StoreID:                "",
		Name:                   "umh-core-pod",
		IsNamespaced:           true,
		Namespace:              "default",
		Compromised:            shared.CompromiseNone,
		ServiceAccount:         "default",
		shareProcessNamespace: false,
		Critical:               false,
	},
	"varlog-pod": {
		StoreID:                "",
		Name:                   "varlog-pod",
		IsNamespaced:           true,
		Namespace:              "default",
		Compromised:            shared.CompromiseNone,
		ServiceAccount:         "default",
		shareProcessNamespace: false,
		Critical:               false,
	},
}

var expectedNodes = map[string]graph.Node{
	"kubehound.test.local-control-plane": {
		StoreID:      "",
		Name:         "kubehound.test.local-control-plane",
		IsNamespaced: false,
		Namespace:    "",
		Compromised:  shared.CompromiseNone,
		Critical:     false,
	},
	"kubehound.test.local-worker": {
		StoreID:      "",
		Name:         "kubehound.test.local-worker",
		IsNamespaced: false,
		Namespace:    "",
		Compromised:  shared.CompromiseNone,
		Critical:     false,
	},
	"kubehound.test.local-worker2": {
		StoreID:      "",
		Name:         "kubehound.test.local-worker2",
		IsNamespaced: false,
		Namespace:    "",
		Compromised:  shared.CompromiseNone,
		Critical:     false,
	},
}

var expectedVolumes = map[string]graph.Volume{
	"nodelog": {
		StoreID: "",
		Name:    "nodelog",
		Type:    "HostPath",
		Path:    "/var/log",
	},
	"nodeproc": {
		StoreID: "",
		Name:    "nodeproc",
		Type:    "HostPath",
		Path:    "/proc/sys/kernel",
	},
}

var expectedContainers = map[string]graph.Container{
	"impersonate-pod": {
		StoreID:      "",
		Name:         "impersonate-pod",
		Image:        "ubuntu",
		Command:      []string{},
		Args:         []string{},
		Capabilities: []string{},
		Privileged:   false,
		PrivEsc:      false,
		HostPID:      false,
		HostPath:     false,
		HostIPC:      false,
		HostNetwork:  false,
		RunAsUser:    0,
		Ports:        []string{},
		Pod:          "impersonate-pod",
		// Node:         "",
		Compromised: 0,
		Critical:    false,
	},
	"modload-pod": {
		StoreID:      "",
		Name:         "modload-pod",
		Image:        "ubuntu",
		Command:      []string{},
		Args:         []string{},
		Capabilities: []string{},
		Privileged:   false,
		PrivEsc:      false,
		HostPID:      false,
		HostPath:     false,
		HostIPC:      false,
		HostNetwork:  false,
		RunAsUser:    0,
		Ports:        []string{},
		Pod:          "modload-pod",
		// Node:         "",
		Compromised: 0,
		Critical:    false,
	},
	"netadmin-pod": {
		StoreID:      "",
		Name:         "netadmin-pod",
		Image:        "ubuntu",
		Command:      []string{},
		Args:         []string{},
		Capabilities: []string{},
		Privileged:   false,
		PrivEsc:      false,
		HostPID:      false,
		HostPath:     false,
		HostIPC:      false,
		HostNetwork:  false,
		RunAsUser:    0,
		Ports:        []string{},
		Pod:          "netadmin-pod",
		// Node:         "",
		Compromised: 0,
		Critical:    false,
	},
	"nsenter-pod": {
		StoreID:      "",
		Name:         "nsenter-pod",
		Image:        "ubuntu",
		Command:      []string{},
		Args:         []string{},
		Capabilities: []string{},
		Privileged:   true,
		PrivEsc:      false,
		HostPID:      false,
		HostPath:     false,
		HostIPC:      false,
		HostNetwork:  false,
		RunAsUser:    0,
		Ports:        []string{},
		Pod:          "nsenter-pod",
		// Node:         "",
		Compromised: 0,
		Critical:    false,
	},
	"pod-create-pod": {
		StoreID:      "",
		Name:         "pod-create-pod",
		Image:        "ubuntu",
		Command:      []string{},
		Args:         []string{},
		Capabilities: []string{},
		Privileged:   false,
		PrivEsc:      false,
		HostPID:      false,
		HostPath:     false,
		HostIPC:      false,
		HostNetwork:  false,
		RunAsUser:    0,
		Ports:        []string{},
		Pod:          "pod-create-pod",
		// Node:         "",
		Compromised: 0,
		Critical:    false,
	},
	"pod-patch-pod": {
		StoreID:      "",
		Name:         "pod-patch-pod",
		Image:        "ubuntu",
		Command:      []string{},
		Args:         []string{},
		Capabilities: []string{},
		Privileged:   false,
		PrivEsc:      false,
		HostPID:      false,
		HostPath:     false,
		HostIPC:      false,
		HostNetwork:  false,
		RunAsUser:    0,
		Ports:        []string{},
		Pod:          "pod-patch-pod",
		// Node:         "",
		Compromised: 0,
		Critical:    false,
	},
	"priv-pod": {
		StoreID:      "",
		Name:         "priv-pod",
		Image:        "ubuntu",
		Command:      []string{},
		Args:         []string{},
		Capabilities: []string{},
		Privileged:   true,
		PrivEsc:      false,
		HostPID:      false,
		HostPath:     false,
		HostIPC:      false,
		HostNetwork:  false,
		RunAsUser:    0,
		Ports:        []string{},
		Pod:          "priv-pod",
		// Node:         "",
		Compromised: 0,
		Critical:    false,
	},
	"rolebind-pod": {
		StoreID:      "",
		Name:         "rolebind-pod",
		Image:        "ubuntu",
		Command:      []string{},
		Args:         []string{},
		Capabilities: []string{},
		Privileged:   false,
		PrivEsc:      false,
		HostPID:      false,
		HostPath:     false,
		HostIPC:      false,
		HostNetwork:  false,
		RunAsUser:    0,
		Ports:        []string{},
		Pod:          "rolebind-pod",
		// Node:         "",
		Compromised: 0,
		Critical:    false,
	},
	"sharedps-pod": {
		StoreID:      "",
		Name:         "sharedps-pod",
		Image:        "ubuntu",
		Command:      []string{},
		Args:         []string{},
		Capabilities: []string{},
		Privileged:   false,
		PrivEsc:      false,
		HostPID:      false,
		HostPath:     false,
		HostIPC:      false,
		HostNetwork:  false,
		RunAsUser:    0,
		Ports:        []string{},
		Pod:          "sharedps-pod",
		// Node:         "",
		Compromised: 0,
		Critical:    false,
	},
	"tokenget-pod": {
		StoreID:      "",
		Name:         "tokenget-pod",
		Image:        "ubuntu",
		Command:      []string{},
		Args:         []string{},
		Capabilities: []string{},
		Privileged:   false,
		PrivEsc:      false,
		HostPID:      false,
		HostPath:     false,
		HostIPC:      false,
		HostNetwork:  false,
		RunAsUser:    0,
		Ports:        []string{},
		Pod:          "tokenget-pod",
		// Node:         "",
		Compromised: 0,
		Critical:    false,
	},
	"tokenlist-pod": {
		StoreID:      "",
		Name:         "tokenlist-pod",
		Image:        "ubuntu",
		Command:      []string{},
		Args:         []string{},
		Capabilities: []string{},
		Privileged:   false,
		PrivEsc:      false,
		HostPID:      false,
		HostPath:     false,
		HostIPC:      false,
		HostNetwork:  false,
		RunAsUser:    0,
		Ports:        []string{},
		Pod:          "tokenlist-pod",
		// Node:         "",
		Compromised: 0,
		Critical:    false,
	},
	"umh-core-pod": {
		StoreID:      "",
		Name:         "umh-core-pod",
		Image:        "ubuntu",
		Command:      []string{},
		Args:         []string{},
		Capabilities: []string{},
		Privileged:   false,
		PrivEsc:      false,
		HostPID:      false,
		HostPath:     false,
		HostIPC:      false,
		HostNetwork:  false,
		RunAsUser:    0,
		Ports:        []string{},
		Pod:          "umh-core-pod",
		// Node:         "",
		Compromised: 0,
		Critical:    false,
	},
	"varlog-pod": {
		StoreID:      "",
		Name:         "varlog-pod",
		Image:        "ubuntu",
		Command:      []string{},
		Args:         []string{},
		Capabilities: []string{},
		Privileged:   false,
		PrivEsc:      false,
		HostPID:      false,
		HostPath:     false,
		HostIPC:      false,
		HostNetwork:  false,
		RunAsUser:    0,
		Ports:        []string{},
		Pod:          "varlog-pod",
		// Node:         "",
		Compromised: 0,
		Critical:    false,
	},
}
