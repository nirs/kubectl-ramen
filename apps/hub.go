// SPDX-FileCopyrightText: The RamenDR authors
// SPDX-License-Identifier: Apache-2.0

package apps

import "github.com/nirs/kubectl-ramen/api"

type HubWatcher struct {
	Cluster   *api.Cluster
	Namespace string
}

type DRPCStatus struct {
	Name              string            `json:"name"`
	Phase             string            `json:"phase"`
	Progression       string            `json:"progression"`
	LastGroupSyncTime string            `json:"lastGroupSyncTime,omitempty"`
	Conditions        map[string]string `json:"conditions"`
}

type HubStatus struct {
	Name string      `json:"name"`
	DRPC *DRPCStatus `json:"drpc,omitempty"`
}

func newHubWatcher(cluster *api.Cluster, namespace string) *HubWatcher {
	return &HubWatcher{Cluster: cluster, Namespace: namespace}
}

func (w *HubWatcher) Status() *HubStatus {
	// XXX Fake status
	return &HubStatus{
		Name: w.Cluster.Name,
		DRPC: &DRPCStatus{
			Name:        "busybox-drpc",
			Phase:       "Relocate",
			Progression: "Completed",
			Conditions: map[string]string{
				"Available": "True",
				"PeerReady": "True",
			},
		},
	}
}
