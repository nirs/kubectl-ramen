// SPDX-FileCopyrightText: The RamenDR authors
// SPDX-License-Identifier: Apache-2.0

package config_test

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"reflect"
	"sort"
	"testing"

	"github.com/nirs/kubectl-ramen/config"
)

// Helpers for creating configurations

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

func chmod(t *testing.T, path string, perm fs.FileMode) {
	err := os.Chmod(path, perm)
	if err != nil {
		t.Fatal(err)
	}
}

func mkclustersets(t *testing.T, names ...string) string {
	configDir := t.TempDir()

	clustersetsDir := filepath.Join(configDir, "clustersets")
	mkdir(t, clustersetsDir, 0700)

	for _, name := range names {
		dir := filepath.Join(clustersetsDir, name)
		mkdir(t, dir, 0700)
		config := filepath.Join(dir, "config.yaml")
		data := fmt.Sprintf("name: %s", name)
		mkfile(t, config, []byte(data), 0600)
	}

	return configDir
}

// Listing clustersets

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
	configDir := mkclustersets(t)
	checkEmptyConfig(t, configDir)
}

func TestConfigSomeInvalidClusterSets(t *testing.T) {
	// ~/.config/kubectl-ramen/clustersets contains some clustersets.
	dirs := []string{"cs1", "cs2", "cs3"}
	configDir := mkclustersets(t, dirs...)

	// Anything else is ignored.
	file := filepath.Join(configDir, "clustersets", "file")
	mkfile(t, file, []byte("ignored"), 0600)

	s := config.NewStore(configDir)
	clustersets, err := s.ListClusterSets()
	if err != nil {
		t.Fatal(err)
	}

	sort.Strings(clustersets)

	if !reflect.DeepEqual(dirs, clustersets) {
		t.Fatalf("Expected %v, got %v", dirs, clustersets)
	}
}

func TestConfigListError(t *testing.T) {
	// ~/.config/kubctl-ramen/clustersets is not readable
	configDir := mkclustersets(t)
	chmod(t, filepath.Join(configDir, "clustersets"), 0)

	s := config.NewStore(configDir)
	_, err := s.ListClusterSets()
	if err == nil {
		t.Fatal("Expected permission error listing unreadable clustersets")
	}
}

func TestConfigDefault(t *testing.T) {
	// We don't know if we have some config, but we can test that we get a
	// config that we can query without errors.
	s := config.DefaultStore()
	_, err := s.ListClusterSets()
	if err != nil {
		t.Fatalf("Error listing default config store: %s", err)
	}
}

func checkEmptyConfig(t *testing.T, path string) {
	s := config.NewStore(path)
	clustersets, err := s.ListClusterSets()
	if err != nil {
		t.Fatal(err)
	}
	if len(clustersets) != 0 {
		t.Fatalf("Expected empty clustersets, got %v", clustersets)
	}
}

// Removing clustersets

func TestConfigRemoveNoClustersetsDir(t *testing.T) {
	configDir := t.TempDir()
	s := config.NewStore(configDir)
	err := s.RemoveClusterSet("missing")
	if err != nil {
		t.Fatal("Error removing with missing clustersets directory")
	}
}

func TestConfigRemoveMissingClusterSet(t *testing.T) {
	configDir := mkclustersets(t)
	s := config.NewStore(configDir)
	err := s.RemoveClusterSet("missing")
	if err != nil {
		t.Fatal("Error removing missing clusterset")
	}
}

func TestConfigRemoveExisting(t *testing.T) {
	configDir := mkclustersets(t, "c1", "c2", "c3")

	s := config.NewStore(configDir)
	err := s.RemoveClusterSet("c2")
	if err != nil {
		t.Fatal("Error removing missing clusterset")
	}

	clustersets, err := s.ListClusterSets()
	if err != nil {
		t.Fatal(err)
	}

	expected := []string{"c1", "c3"}
	sort.Strings(clustersets)

	if !reflect.DeepEqual(expected, clustersets) {
		t.Fatalf("Expected %v, got %v", expected, clustersets)
	}
}

// Validating names

var invalidNames = []struct {
	description string
	value       string
}{
	{"empty", ""},
	{"hidden", "."},
	{"absolute path", "/a"},
	{"multiple path components", "a/b"},
	{"outside clusteersets dir", "../b"},
	{"path manipulations", "a/../b"},
	{"newline", "a\nb"}, // breaks list output in text mode
}

func TestConfigRemoveInvalid(t *testing.T) {
	configDir := mkclustersets(t)
	s := config.NewStore(configDir)

	for _, n := range invalidNames {
		t.Run(n.description, func(t *testing.T) {
			err := s.RemoveClusterSet(n.value)
			if err == nil {
				t.Fatalf("Removing invalid name %q did not fail", n.value)
			} else {
				t.Logf("expected error: %s", err)
			}
		})
	}
}
