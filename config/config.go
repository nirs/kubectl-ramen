// SPDX-FileCopyrightText: The RamenDR authors
// SPDX-License-Identifier: Apache-2.0

package config

import (
	"os"
	"path/filepath"

	"k8s.io/client-go/util/homedir"
)

func configDir() string {
	return filepath.Join(homedir.HomeDir(), ".config", "kubectl-ramen")
}

// Store stores and loads configurations.
type Store struct {
	path string
}

// DefaultStore return the deault configuration storage.
func DefaultStore() *Store {
	return NewStore(configDir())
}

// NewStore return a new configuration storage using the specified path.
func NewStore(path string) *Store {
	return &Store{path: path}
}

// ListClusterSets return slice of clusterset names. The result may contain
// invalid configurations.
func (s *Store) ListClusterSets() ([]string, error) {
	var clustersets []string

	entries, err := os.ReadDir(s.clustersetsDir())
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
func (s *Store) clustersetsDir() string {
	return filepath.Join(s.path, "clustersets")
}

