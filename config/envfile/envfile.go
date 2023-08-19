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

func Load(path string) (*EnvFile, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	env := &EnvFile{}
	err = yaml.Unmarshal(data, env)
	if err != nil {
		return nil, err
	}

	return env, nil
}
