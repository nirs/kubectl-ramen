// SPDX-FileCopyrightText: The RamenDR authors
// SPDX-License-Identifier: Apache-2.0

package core_test

import (
	"io/fs"
	"os"
	"path/filepath"
	"reflect"
	"sort"
	"testing"

	"github.com/nirs/kubectl-ramen/core"
)

func mkdir(t *testing.T, path string, perm fs.FileMode) {
	err := os.Mkdir(path, perm)
	if err != nil {
		t.Fatal(err)
	}
}

func mkfile(t *testing.T, path string, data []byte, perm fs.FileMode) {
	err := os.WriteFile(path, data, perm)
	if err != nil {
		t.Fatal(err)
	}
}

func TestConfigNoConfigDir(t *testing.T) {
	// ~/.config/kubectl-ramen/ does not exist.
	configDir := filepath.Join(t.TempDir(), "missing")
	checkEmptyConfig(t, configDir)
}

func TestConfigNoClustersetsDir(t *testing.T) {
	// ~/.config/kubectl-ramen/clustersets does not exist.
	configDir := t.TempDir()
	checkEmptyConfig(t, configDir)
}

func TestConfigNoClusterSet(t *testing.T) {
	// ~/.config/kubectl-ramen/clustersets is empty.
	configDir := t.TempDir()
	clustersetsDir := filepath.Join(configDir, "clustersets")
	mkdir(t, clustersetsDir, 0700)
	checkEmptyConfig(t, configDir)
}

func TestConfigSomeInvalidClusterSets(t *testing.T) {
	// ~/.config/kubectl-ramen/clustersets contains some (invalid) clustersets.
	configDir := t.TempDir()
	clustersetsDir := filepath.Join(configDir, "clustersets")
	mkdir(t, clustersetsDir, 0700)

	dirs := []string{"cs1", "cs2", "cs3"}

	// Directories are considred as clusterset.
	for _, name := range dirs {
		clusterset := filepath.Join(clustersetsDir, name)
		mkdir(t, clusterset, 0700)
	}

	// Anything else is ignored.
	mkfile(t, filepath.Join(clustersetsDir, "file"), []byte("ignored"), 0600)

	s := core.NewConfigStorage(configDir)
	clustersets, err := s.ListClusterSets()
	if err != nil {
		t.Fatal(err)
	}

	sort.Strings(clustersets)

	if !reflect.DeepEqual(dirs, clustersets) {
		t.Fatalf("Expected %v, got %v", dirs, clustersets)
	}
}

func TestConfigDefault(t *testing.T) {
	// We don't know if we have some config, but we can test that we get a
	// config that we can query without errors.
	s := core.DefaultConfigStorage()
	_, err := s.ListClusterSets()
	if err != nil {
		t.Fatalf("Error listing default config store: %s", err)
	}
}

func checkEmptyConfig(t *testing.T, path string) {
	s := core.NewConfigStorage(path)
	clustersets, err := s.ListClusterSets()
	if err != nil {
		t.Fatal(err)
	}
	if len(clustersets) != 0 {
		t.Fatalf("Expected empty clustersets, got %v", clustersets)
	}
}
