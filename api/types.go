// SPDX-FileCopyrightText: The RamenDR authors
// SPDX-License-Identifier: Apache-2.0

package api

type DRTopology string

const (
	RegionalDR = DRTopology("regional-dr")
	MetroDR    = DRTopology("metro-dr")
)

type Cluster struct {
	Name       string `json:"name"`
	Kubeconfig string `json:"kubeconfig"`
}

type ClusterSet struct {
	Name     string     `json:"name"`
	Topology DRTopology `json:"topology"`
	Hub      *Cluster   `json:"hub"`
	Cluster1 *Cluster   `json:"cluster1"`
	Cluster2 *Cluster   `json:"cluster2"`
}
