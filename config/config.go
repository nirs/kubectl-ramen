// SPDX-FileCopyrightText: The RamenDR authors
// SPDX-License-Identifier: Apache-2.0

package config

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/nirs/kubectl-ramen/api"
	"github.com/nirs/kubectl-ramen/config/envfile"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/tools/clientcmd/api/latest"
	"k8s.io/client-go/util/homedir"
	"sigs.k8s.io/yaml"
)

func configDir() string {
	return filepath.Join(homedir.HomeDir(), ".config", "kubectl-ramen")
}

// Store stores and loads configurations.
type Store struct {
	path       string
	kubeconfig string
}

// DefaultStore return the default configuration storage.
func DefaultStore() *Store {
	return NewStore(configDir(), clientcmd.RecommendedHomeFile)
}

// NewStore return a new configuration storage using the specified path.
func NewStore(path string, kubeconfig string) *Store {
	return &Store{path: path, kubeconfig: kubeconfig}
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

// AddClusterSetFromEnvFile add a clusterset from ramen environment file.
func (s *Store) AddClusterSetFromEnvFile(name string, env *envfile.EnvFile) error {
	if !s.isValidName(name) {
		return fmt.Errorf("invalid clusterset name: %q", name)
	}

	dir := filepath.Join(s.clustersetsDir(), name)
	clusterset := s.newClusterSetFromEnv(name, env, dir)

	err := s.createClusterSetDir(dir)
	if err != nil {
		if os.IsExist(err) {
			return fmt.Errorf("clusterset %q already exist", name)
		}
		return err
	}

	err = s.writeClusterSet(clusterset, dir)
	if err != nil {
		os.RemoveAll(dir)
		return err
	}

	return nil
}

func (s *Store) newClusterSetFromEnv(name string, env *envfile.EnvFile, dir string) *api.ClusterSet {
	clusterset := &api.ClusterSet{
		Name:     name,
		Topology: env.Ramen.Topology,
		Cluster1: &api.Cluster{
			Name:       env.Ramen.Clusters[0],
			Kubeconfig: filepath.Join(dir, env.Ramen.Clusters[0]+".kubeconfig"),
		},
		Cluster2: &api.Cluster{
			Name:       env.Ramen.Clusters[1],
			Kubeconfig: filepath.Join(dir, env.Ramen.Clusters[1]+".kubeconfig"),
		},
	}

	if env.Ramen.Hub != "" {
		clusterset.Hub = &api.Cluster{
			Name:       env.Ramen.Hub,
			Kubeconfig: filepath.Join(dir, env.Ramen.Hub+".kubeconfig"),
		}
	}

	return clusterset
}

func (s *Store) createClusterSetDir(path string) error {
	err := os.MkdirAll(s.clustersetsDir(), 0700)
	if err != nil {
		return err
	}

	err = os.Mkdir(path, 0700)
	if err != nil {
		return err
	}

	return nil
}

func (s *Store) writeClusterSet(clusterset *api.ClusterSet, dir string) error {
	if clusterset.Hub != nil {
		err := s.copyKubeConfigFor(clusterset.Hub)
		if err != nil {
			return err
		}
	}

	err := s.copyKubeConfigFor(clusterset.Cluster1)
	if err != nil {
		return err
	}

	err = s.copyKubeConfigFor(clusterset.Cluster2)
	if err != nil {
		return err
	}

	data, err := yaml.Marshal(clusterset)
	if err != nil {
		return err
	}

	config := filepath.Join(dir, "config.yaml")
	err = os.WriteFile(config, data, 0600)
	if err != nil {
		return err
	}

	return nil
}

func (s *Store) copyKubeConfigFor(cluster *api.Cluster) error {
	kubeconfig, err := envfile.LoadKubeConfigFor(cluster.Name, s.kubeconfig)
	if err != nil {
		return err
	}

	// NOTE: Loading the config fails without the conversion. Not documented
	// but kubectl does this.
	converted, err := latest.Scheme.ConvertToVersion(kubeconfig, latest.ExternalVersion)
	if err != nil {
		return err
	}

	data, err := yaml.Marshal(converted)
	if err != nil {
		return err
	}

	err = os.WriteFile(cluster.Kubeconfig, data, 0600)
	if err != nil {
		return err
	}

	return nil
}

func (s *Store) GetClusterSet(name string) (*api.ClusterSet, error) {
	if !s.isValidName(name) {
		return nil, fmt.Errorf("invalid clusterset name: %q", name)
	}

	dir := filepath.Join(s.clustersetsDir(), name)
	_, err := os.Stat(dir)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, fmt.Errorf("clusterset %q does not exist", name)
		}
		return nil, err
	}

	config := filepath.Join(dir, "config.yaml")
	data, err := os.ReadFile(config)
	if err != nil {
		return nil, err
	}

	clusterset := api.ClusterSet{}
	err = yaml.Unmarshal(data, &clusterset)
	if err != nil {
		return nil, err
	}

	return &clusterset, nil
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
