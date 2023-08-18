// SPDX-FileCopyrightText: The RamenDR authors
// SPDX-License-Identifier: Apache-2.0

package config

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

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

func (s *Store) RemoveClusterSet(name string) error {
	if !s.isValidName(name) {
		return fmt.Errorf("invalid clusterset name: %q", name)
	}
	clustersetDir := filepath.Join(s.clustersetsDir(), name)
	return os.RemoveAll(clustersetDir)
}

func (s *Store) clustersetsDir() string {
	return filepath.Join(s.path, "clustersets")
}

// forbiddenCharacters cannot be used in clusterset name.
var forbiddenCharacters = string([]rune{os.PathSeparator, '\n'})

// isValidName return true if name can be used for storing items.
func (s *Store) isValidName(name string) bool {
	if name == "" {
		return false
	}
	if strings.ContainsAny(name, forbiddenCharacters) {
		return false
	}
	if name[0] == '.' {
		return false
	}
	return true
}
