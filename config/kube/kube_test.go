// SPDX-FileCopyrightText: The RamenDR authors
// SPDX-License-Identifier: Apache-2.0

package kube_test

import (
	"path/filepath"
	"testing"

	"github.com/nirs/kubectl-ramen/config/kube"
)

func TestConfigKubeNewConfigSet(t *testing.T) {
	type cluster struct {
		filename string
		context  string
	}
	hub := cluster{
		filename: "perf1.kubeconfig",
		context:  "default/api-perf1-examle-com:1234/kube:admin",
	}
	cluster1 := cluster{
		filename: "perf2.kubeconfig",
		context:  "default/api-perf2-examle-com:1234/kube:admin",
	}
	cluster2 := cluster{
		filename: "perf3.kubeconfig",
		context:  "default/api-perf3-examle-com:1234/kube:admin",
	}

	c, err := kube.NewConfigSet(
		filepath.Join("testdata", hub.filename),
		filepath.Join("testdata", cluster1.filename),
		filepath.Join("testdata", cluster2.filename),
	)
	if err != nil {
		t.Fatal(err)
	}

	if c.Hub.Filename != hub.filename {
		t.Fatalf("Expected Hub filename %q, got %q", hub.filename, c.Hub.Filename)
	}
	if c.Hub.Config.CurrentContext != hub.context {
		t.Fatalf("Expected hub context %q, got %q", hub.context, c.Hub.Config.CurrentContext)
	}

	if c.Cluster1.Filename != cluster1.filename {
		t.Fatalf("Expected Hub filename %q, got %q", cluster1.filename, c.Cluster1.Filename)
	}
	if c.Cluster1.Config.CurrentContext != cluster1.context {
		t.Fatalf("Expected hub context %q, got %q", cluster1.context, c.Cluster1.Config.CurrentContext)
	}

	if c.Cluster2.Filename != cluster2.filename {
		t.Fatalf("Expected Hub filename %q, got %q", cluster2.filename, c.Cluster2.Filename)
	}
	if c.Cluster2.Config.CurrentContext != cluster2.context {
		t.Fatalf("Expected hub context %q, got %q", cluster2.context, c.Cluster2.Config.CurrentContext)
	}
}
