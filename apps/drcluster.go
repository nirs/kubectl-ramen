// SPDX-FileCopyrightText: The RamenDR authors
// SPDX-License-Identifier: Apache-2.0

package apps

import "github.com/nirs/kubectl-ramen/api"

type DRClusterWatcher struct {
	Cluster   *api.Cluster
	Namespace string
}
type VRGStatus struct {
	Name       string            `json:"name"`
	State      string            `json:"state"`
	Conditions map[string]string `json:"conditions"`
}

type VRStatus struct {
	Name       string            `json:"name"`
	State      string            `json:"state"`
	Conditions map[string]string `json:"conditions"`
}

type DRClusterStatus struct {
	Name string      `json:"name"`
	VRG  *VRGStatus  `json:"vrg,omitempty"`
	VRs  []*VRStatus `json:"vrs,omitempty"`
}

func newDRClusterWatcher(cluster *api.Cluster, namespace string) *DRClusterWatcher {
	return &DRClusterWatcher{Cluster: cluster, Namespace: namespace}
}

func (w *DRClusterWatcher) Status() *DRClusterStatus {
	// XXX Fake status
	if w.Cluster.Name == "dr1" {
		return &DRClusterStatus{
			Name: w.Cluster.Name,
			VRG: &VRGStatus{
				Name:  "busybox-drpc",
				State: "Primary",
				Conditions: map[string]string{
					"DataReady":            "True",
					"DataProtected":        "False",
					"ClusterDataReady":     "True",
					"ClusterDataProtected": "True",
				},
			},
			VRs: []*VRStatus{
				{
					Name:  "busybox-pvc",
					State: "Primary",
					Conditions: map[string]string{
						"Completed": "True",
						"Degrated":  "False",
						"Resyncing": "False",
					},
				},
			},
		}
	} else {
		return &DRClusterStatus{
			Name: w.Cluster.Name,
		}
	}
}
