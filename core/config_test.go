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

func tmpdir(t *testing.T) string {
	path, err := os.MkdirTemp("", "kubectl-ramen.tmp")
	if err != nil {
		t.Fatal(err)
	}

	t.Cleanup(func() {
		err := os.RemoveAll(path)
		if err != nil {
			// Nothing we can do here expect logging. The test runner will have
			// to clean up manually.
			t.Logf("Cannot remove temporary directory %q: %s", path, err)
		}
	})

	return path
}

func mkdir(t *testing.T, path string, perm fs.FileMode) {
	err := os.Mkdir(path, perm)
	if err != nil {
		t.Fatal(err)
	}
}

func TestConfigNoConfigDir(t *testing.T) {
	// ~/.config/kubectl-ramen/ does not exist.
	configDir := filepath.Join(tmpdir(t), "missing")
	checkEmptyConfig(t, configDir)
}

func TestConfigNoClustersetsDir(t *testing.T) {
	// ~/.config/kubectl-ramen/clustersets does not exist.
	configDir := tmpdir(t)
	checkEmptyConfig(t, configDir)
}

func TestConfigNoClusterSet(t *testing.T) {
	// ~/.config/kubectl-ramen/clustersets is empty.
	configDir := tmpdir(t)
	clustersetsDir := filepath.Join(configDir, "clustersets")
	mkdir(t, clustersetsDir, 0777)
	checkEmptyConfig(t, configDir)
}

func TestConfigSomeInvalidClusterSets(t *testing.T) {
	// ~/.config/kubectl-ramen/clustersets contains some (invalid) clustersets.
	configDir := tmpdir(t)
	clustersetsDir := filepath.Join(configDir, "clustersets")
	mkdir(t, clustersetsDir, 0777)

	expected := []string{"cs1", "cs2", "cs3"}

	for _, name := range expected {
		clusterset := filepath.Join(clustersetsDir, name)
		mkdir(t, clusterset, 0777)
	}

	s := core.NewConfigStorage(configDir)
	actual, err := s.ListClusterSets()
	if err != nil {
		t.Fatal(err)
	}

	sort.Strings(expected)
	sort.Strings(actual)

	if !reflect.DeepEqual(expected, actual) {
		t.Fatalf("Expected %v, got %v", expected, actual)
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
