// SPDX-FileCopyrightText: The RamenDR authors
// SPDX-License-Identifier: Apache-2.0

package api

import "k8s.io/client-go/kubernetes"

type DRTopology string

const (
	RegionalDR = DRTopology("regional-dr")
	MetroDR    = DRTopology("metro-dr")
)

type ClusterConfig struct {
	Kubeconfig string `json:"kubeconfig"`
}

type ClusterSetConfig struct {
	Name     string                    `json:"name"`
	Topology DRTopology                `json:"topology"`
	Clusters map[string]*ClusterConfig `json:"clusters"`
}

type Cluster struct {
	name   string
	client *kubernetes.Interface
}

type ClusterSet struct {
	name     string
	topology DRTopology
	hub      *Cluster
	cluster1 *Cluster
	cluster2 *Cluster
}
