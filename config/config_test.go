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

	"github.com/nirs/kubectl-ramen/api"
	"github.com/nirs/kubectl-ramen/config"
	"github.com/nirs/kubectl-ramen/config/drenv"
	"github.com/nirs/kubectl-ramen/config/kube"
	"k8s.io/client-go/tools/clientcmd"
)

var kubeconfig = filepath.Join("drenv", "testdata", "minikube.kubeconfig")

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

	s := config.NewStore(configDir, kubeconfig)
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

	s := config.NewStore(configDir, kubeconfig)
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
	s := config.NewStore(path, kubeconfig)
	clustersets, err := s.ListClusterSets()
	if err != nil {
		t.Fatal(err)
	}
	if len(clustersets) != 0 {
		t.Fatalf("Expected empty clustersets, got %v", clustersets)
	}
}

// Adding clusterset from env file

func TestConfigAddClusterSetFromEnvFile(t *testing.T) {
	env := drenv.Environment{
		Name: "e2e",
		Ramen: &drenv.RamenInfo{
			Hub:      "hub",
			Clusters: []string{"dr1", "dr2"},
			Topology: api.RegionalDR,
		},
	}

	s := config.NewStore(t.TempDir(), kubeconfig)

	err := s.AddClusterSetFromEnv(env.Name, &env)
	if err != nil {
		t.Fatal(err)
	}

	clusterset, err := s.GetClusterSet(env.Name)
	if err != nil {
		t.Fatal(err)
	}

	checkClusterSet(t, clusterset, env.Name, api.RegionalDR)
	checkCluster(t, clusterset.Hub, env.Ramen.Hub)
	checkCluster(t, clusterset.Cluster1, env.Ramen.Clusters[0])
	checkCluster(t, clusterset.Cluster2, env.Ramen.Clusters[1])
}

func TestConfigAddClusterSetFromEnvFileHubless(t *testing.T) {
	env := drenv.Environment{
		Name: "hubless",
		Ramen: &drenv.RamenInfo{
			Hub:      "",
			Clusters: []string{"dr1", "dr2"},
			Topology: api.RegionalDR,
		},
	}

	s := config.NewStore(t.TempDir(), kubeconfig)

	err := s.AddClusterSetFromEnv(env.Name, &env)
	if err != nil {
		t.Fatal(err)
	}

	clusterset, err := s.GetClusterSet(env.Name)
	if err != nil {
		t.Fatal(err)
	}

	checkClusterSet(t, clusterset, env.Name, api.RegionalDR)
	if clusterset.Hub != nil {
		t.Fatalf("Expected nil Hub, got %v", clusterset.Hub)
	}
	checkCluster(t, clusterset.Cluster1, "dr1")
	checkCluster(t, clusterset.Cluster2, "dr2")
}

// Adding clusterset from kubeconfigs.

func TestConfigAddClusterSetFromConfigs(t *testing.T) {
	configs, err := kube.NewConfigSet(
		filepath.Join("kube", "testdata", "perf1.kubeconfig"),
		filepath.Join("kube", "testdata", "perf2.kubeconfig"),
		filepath.Join("kube", "testdata", "perf3.kubeconfig"),
	)
	if err != nil {
		t.Fatal(err)
	}

	s := config.NewStore(t.TempDir(), kubeconfig)

	err = s.AddClusterSetFromConfigs("perf", configs)
	if err != nil {
		t.Fatal(err)
	}

	clusterset, err := s.GetClusterSet("perf")
	if err != nil {
		t.Fatal(err)
	}

	checkClusterSet(t, clusterset, "perf", "")
	checkCluster(t, clusterset.Hub, "perf1")
	checkCluster(t, clusterset.Cluster1, "perf2")
	checkCluster(t, clusterset.Cluster2, "perf3")
}

// Getting clustersets

func TestConfigGetClusterSetNoClusterSetsDir(t *testing.T) {
	s := config.NewStore(t.TempDir(), kubeconfig)
	_, err := s.GetClusterSet("name")
	if err == nil {
		t.Fatal("Get non existing clusterset succeeded")
	}
}

func TestConfigGetClusterSetMissingClusterSet(t *testing.T) {
	env := drenv.Environment{
		Name: "e2e",
		Ramen: &drenv.RamenInfo{
			Hub:      "hub",
			Clusters: []string{"dr1", "dr2"},
		},
	}
	s := config.NewStore(t.TempDir(), kubeconfig)
	err := s.AddClusterSetFromEnv(env.Name, &env)
	if err != nil {
		t.Fatal(err)
	}
	_, err = s.GetClusterSet("name")
	if err == nil {
		t.Fatal("Get non existing clusterset succeeded")
	}
}

// Verifying clustersets

func checkClusterSet(t *testing.T, clusterset *api.ClusterSet, name string, topology api.DRTopology) {
	if clusterset.Name != name {
		t.Fatalf("Expected clusterset name %q, got %q", name, clusterset.Name)
	}
	if clusterset.Topology != topology {
		t.Fatalf("Expected clusterset topology %q, got %q", topology, clusterset.Topology)
	}
}

func checkCluster(t *testing.T, cluster *api.Cluster, name string) {
	if cluster.Name != name {
		t.Fatalf("Expected cluster name %q, got %q", name, cluster.Name)
	}
	_, err := clientcmd.LoadFromFile(cluster.Kubeconfig)
	if err != nil {
		t.Fatalf("Error loading cluster %q kubeconfig: %s", name, err)
	}
}

// Removing clustersets

func TestConfigRemoveNoClustersetsDir(t *testing.T) {
	configDir := t.TempDir()
	s := config.NewStore(configDir, kubeconfig)
	err := s.RemoveClusterSet("missing")
	if err != nil {
		t.Fatal("Error removing with missing clustersets directory")
	}
}

func TestConfigRemoveMissingClusterSet(t *testing.T) {
	configDir := mkclustersets(t)
	s := config.NewStore(configDir, kubeconfig)
	err := s.RemoveClusterSet("missing")
	if err != nil {
		t.Fatal("Error removing missing clusterset")
	}
}

func TestConfigRemoveExisting(t *testing.T) {
	configDir := mkclustersets(t, "c1", "c2", "c3")

	s := config.NewStore(configDir, kubeconfig)
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
	s := config.NewStore(configDir, kubeconfig)

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

func TestConfigAddClusterSetFromEnvFileInvalid(t *testing.T) {
	configDir := mkclustersets(t)

	s := config.NewStore(configDir, kubeconfig)

	for _, n := range invalidNames {
		t.Run(n.description, func(t *testing.T) {
			err := s.AddClusterSetFromEnv(n.value, &drenv.Environment{})
			if err == nil {
				t.Fatalf("Adding invalid name %q did not fail", n.value)
			} else {
				t.Logf("expected error: %s", err)
			}
		})
	}
}

func TestConfigAdddClusterSetFromConfigsInvalid(t *testing.T) {
	configs, err := kube.NewConfigSet(
		filepath.Join("kube", "testdata", "perf1.kubeconfig"),
		filepath.Join("kube", "testdata", "perf2.kubeconfig"),
		filepath.Join("kube", "testdata", "perf3.kubeconfig"),
	)
	if err != nil {
		t.Fatal(err)
	}

	configDir := mkclustersets(t)
	s := config.NewStore(configDir, kubeconfig)

	for _, n := range invalidNames {
		t.Run(n.description, func(t *testing.T) {
			err := s.AddClusterSetFromConfigs(n.value, configs)
			if err == nil {
				t.Fatalf("Adding invalid name %q did not fail", n.value)
			} else {
				t.Logf("expected error: %s", err)
			}
		})
	}
}
func TestConfiGetdClusterSetInvalid(t *testing.T) {
	configDir := mkclustersets(t)
	s := config.NewStore(configDir, kubeconfig)

	for _, n := range invalidNames {
		t.Run(n.description, func(t *testing.T) {
			_, err := s.GetClusterSet(n.value)
			if err == nil {
				t.Fatalf("Adding invalid name %q did not fail", n.value)
			} else {
				t.Logf("expected error: %s", err)
			}
		})
	}
}
