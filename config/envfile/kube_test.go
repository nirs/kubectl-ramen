// SPDX-FileCopyrightText: The RamenDR authors
// SPDX-License-Identifier: Apache-2.0

package envfile_test

import (
	"path/filepath"
	"testing"

	"github.com/nirs/kubectl-ramen/config/envfile"
)

func TestConfigLoadKubeConfigFor(t *testing.T) {
	minikubeKubeconfig := filepath.Join("testdata", "minikube.kubeconfig")

	for _, name := range []string{"hub", "dr1", "dr2"} {
		t.Run(name, func(t *testing.T) {
			config, err := envfile.LoadKubeConfigFor(name, minikubeKubeconfig)
			if err != nil {
				t.Fatal(err)
			}

			// Current context must be set to requested context.
			if config.CurrentContext != name {
				t.Fatalf("Expected current context %q, got %q", name, config.CurrentContext)
			}

			// We must thave one context pointing to matching cluster and user.
			if len(config.Contexts) != 1 {
				t.Fatalf("Expected single context, got %v", len(config.Contexts))
			}
			context, ok := config.Contexts[name]
			if !ok {
				t.Fatalf("Context %v missing", name)
			}

			// We must have one cluster matching the context cluster.
			if len(config.Clusters) != 1 {
				t.Fatalf("Expected single cluster, got %v", len(config.Clusters))
			}
			_, ok = config.Clusters[context.Cluster]
			if !ok {
				t.Fatalf("Cluster %v missing", name)
			}

			// We must have one user matching the context user.
			if len(config.AuthInfos) != 1 {
				t.Fatalf("Expected single user, got %v", len(config.AuthInfos))
			}
			_, ok = config.AuthInfos[context.AuthInfo]
			if !ok {
				t.Fatalf("User %v missing", name)
			}
		})
	}
}
