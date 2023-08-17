// SPDX-FileCopyrightText: The RamenDR authors
// SPDX-License-Identifier: Apache-2.0

package core

import (
	"os"
	"path/filepath"

	"k8s.io/client-go/util/homedir"
)

type ClusterConfig struct {
	Name       string `yaml:"name"`
	Kubeconfig string `yaml:"kubeconfig"`
}

type ClusterSetConfig struct {
	Name     string           `yaml:"name"`
	Topology DRTopology       `yaml:"topology"`
	Clusters []*ClusterConfig `yaml:"clusters"`
}

func configDir() string {
	return filepath.Join(homedir.HomeDir(), ".config", "kubectl-ramen")
}

// ConfigStorage stores and loads configurations.
type ConfigStorage struct {
	path string
}

// DefaultConfigStorage return the deault configuration storage.
func DefaultConfigStorage() *ConfigStorage {
	return NewConfigStorage(configDir())
}

// NewConfigStorage return a new configuration storage using the specified path.
func NewConfigStorage(path string) *ConfigStorage {
	return &ConfigStorage{path: path}
}

// ListClusterSets return slice of clusterset names. The result may contain
// invalid configurations.
func (s *ConfigStorage) ListClusterSets() ([]string, error) {
	clustersetsDir := filepath.Join(s.path, "clustersets")

	var clustersets []string

	entries, err := os.ReadDir(clustersetsDir)
	if err != nil {
		if os.IsNotExist(err) {
			// Expected: nothing was stored yet, or user deleted the directory.
			return clustersets, nil
		}
		return nil, err
	}

	for _, entry := range entries {
		if entry.IsDir() {
			clustersets = append(clustersets, entry.Name())
		}
	}

	return clustersets, nil
}
