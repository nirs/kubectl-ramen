// SPDX-FileCopyrightText: The RamenDR authors
// SPDX-License-Identifier: Apache-2.0

package apps

import (
	"github.com/nirs/kubectl-ramen/api"
)

type ApplicationWatcher struct {
	Clusterset *api.ClusterSet
	Namespace  string
	Hub        *HubWatcher
	Cluster1   *DRClusterWatcher
	Cluster2   *DRClusterWatcher
}

type ApplicationStatus struct {
	Namespace string           `json:"namespace"`
	Hub       *HubStatus       `json:"hub"`
	Cluster1  *DRClusterStatus `json:"cluster1"`
	Cluster2  *DRClusterStatus `json:"cluster2"`
}

func NewApplicationWatcher(clusterset *api.ClusterSet, namespace string) (*ApplicationWatcher, error) {
	watcher := &ApplicationWatcher{
		Clusterset: clusterset,
		Namespace:  namespace,
		Hub:        newHubWatcher(clusterset.Hub, namespace),
		Cluster1:   newDRClusterWatcher(clusterset.Cluster1, namespace),
		Cluster2:   newDRClusterWatcher(clusterset.Cluster2, namespace),
	}

	// XXX Wait until all watchers are ready.

	return watcher, nil
}

func (w *ApplicationWatcher) Status() (*ApplicationStatus, error) {
	return &ApplicationStatus{
		Namespace: w.Namespace,
		Hub:       w.Hub.Status(),
		Cluster1:  w.Cluster1.Status(),
		Cluster2:  w.Cluster2.Status(),
	}, nil
}
