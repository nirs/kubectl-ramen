// SPDX-FileCopyrightText: The RamenDR authors
// SPDX-License-Identifier: Apache-2.0

package drenv_test

import (
	"path/filepath"
	"reflect"
	"testing"

	"github.com/nirs/kubectl-ramen/config/drenv"
)

func TestConfigDrenvLoadMissing(t *testing.T) {
	missing := filepath.Join(t.TempDir(), "no-such-env.yaml")
	_, err := drenv.Load(missing, drenv.Options{})
	if err == nil {
		t.Fatal("Loading missing envfile did not fail")
	}
}

func TestConfigDrenvLoadE2E(t *testing.T) {
	path := filepath.Join("testdata", "e2e.yaml")
	env, err := drenv.Load(path, drenv.Options{})
	if err != nil {
		t.Fatal(err)
	}

	if env.Name != "e2e" {
		t.Errorf("Expected name %q, got %q", "e2e", env.Name)
	}

	if env.Ramen.Hub != "hub" {
		t.Errorf("Expected hub %q, got %q", "hub", env.Ramen.Hub)
	}
	clusters := []string{"dr1", "dr2"}
	if !reflect.DeepEqual(env.Ramen.Clusters, clusters) {
		t.Errorf("Expected clusters %q, got %q", clusters, env.Ramen.Clusters)
	}
}

func TestConfigDrenvLoadNamePrefix(t *testing.T) {
	const namePrefix = "prefix-"
	path := filepath.Join("testdata", "e2e.yaml")
	env, err := drenv.Load(path, drenv.Options{NamePrefix: namePrefix})
	if err != nil {
		t.Fatal(err)
	}

	if env.Name != namePrefix+"e2e" {
		t.Errorf("Expected name %q, got %q", namePrefix+"e2e", env.Name)
	}

	if env.Ramen.Hub != namePrefix+"hub" {
		t.Errorf("Expected hub %q, got %q", "hub", env.Ramen.Hub)
	}
	clusters := []string{namePrefix + "dr1", namePrefix + "dr2"}
	if !reflect.DeepEqual(env.Ramen.Clusters, clusters) {
		t.Errorf("Expected clusters %q, got %q", clusters, env.Ramen.Clusters)
	}
}

func TestConfigDrenvLoadHubless(t *testing.T) {
	path := filepath.Join("testdata", "hubless.yaml")
	env, err := drenv.Load(path, drenv.Options{})
	if err != nil {
		t.Fatal(err)
	}

	if env.Name != "hubless" {
		t.Errorf("Expected name %q, got %q", "e2e", env.Name)
	}

	if env.Ramen.Hub != "" {
		t.Errorf("Expected hub %q, got %q", "hub", env.Ramen.Hub)
	}
	clusters := []string{"dr1", "dr2"}
	if !reflect.DeepEqual(env.Ramen.Clusters, clusters) {
		t.Errorf("Expected clusters %q, got %q", clusters, env.Ramen.Clusters)
	}
}
