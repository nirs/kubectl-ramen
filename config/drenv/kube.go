// SPDX-FileCopyrightText: The RamenDR authors
// SPDX-License-Identifier: Apache-2.0

package drenv

import (
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/tools/clientcmd/api"
)

// LoadKubeConfigFor loads config for specified context form minikube generated
// kubeconfig file with multiple contexts.
//
// The config refers to minikube certificates and keys on the same host, so it
// cannot be used from another host.
//
// Like `kubectl config view --minify --context context`.  Based on
// kubectl/pkg/cmd/config/view.go
func LoadKubeConfigFor(context string, path string) (*api.Config, error) {
	config, err := clientcmd.LoadFromFile(path)
	if err != nil {
		return nil, err
	}

	config.CurrentContext = context

	if err := api.MinifyConfig(config); err != nil {
		return nil, err
	}

	return config, nil
}
