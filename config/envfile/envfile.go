// SPDX-FileCopyrightText: The RamenDR authors
// SPDX-License-Identifier: Apache-2.0

package envfile

import (
	"os"

	"github.com/nirs/kubectl-ramen/api"
	"sigs.k8s.io/yaml"
)

// EnvInfo describes the ramen environment. The hub and clusters names are
// context names stored in the home kubeconfig (~/.kube/config).
type EnvInfo struct {
	Hub      string         `json:"hub"`
	Clusters []string       `json:"clusters"`
	Topology api.DRTopology `json:"topology"`
}

// EnvFile is ramen testing environment file.
type EnvFile struct {
	Name  string   `json:"name"`
	Ramen *EnvInfo `json:"ramen"`
}

type Options struct {
	NamePrefix string
}

func Load(path string, options Options) (*EnvFile, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	env := &EnvFile{}
	err = yaml.Unmarshal(data, env)
	if err != nil {
		return nil, err
	}

	if options.NamePrefix != "" {
		prefixNames(env, options.NamePrefix)
	}

	return env, nil
}

func prefixNames(env *EnvFile, prefix string) {
	env.Name = prefix + env.Name
	if env.Ramen.Hub != "" {
		env.Ramen.Hub = prefix + env.Ramen.Hub
	}
	for i, cluster := range env.Ramen.Clusters {
		env.Ramen.Clusters[i] = prefix + cluster
	}
}
