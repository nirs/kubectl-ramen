// SPDX-FileCopyrightText: The RamenDR authors
// SPDX-License-Identifier: Apache-2.0

package kube

import (
	"path"

	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/tools/clientcmd/api"
)

type Config struct {
	Filename string
	Config   *api.Config
}

type ConfigSet struct {
	Hub      *Config
	Cluster1 *Config
	Cluster2 *Config
}

func NewConfigSet(hub string, cluster1 string, cluster2 string) (*ConfigSet, error) {
	hubConfig, err := loadConfig(hub)
	if err != nil {
		return nil, err
	}
	cluster1Config, err := loadConfig(cluster1)
	if err != nil {
		return nil, err
	}
	cluster2Config, err := loadConfig(cluster2)
	if err != nil {
		return nil, err
	}
	return &ConfigSet{
		Hub:      hubConfig,
		Cluster1: cluster1Config,
		Cluster2: cluster2Config,
	}, nil
}

func loadConfig(filename string) (*Config, error) {
	config, err := clientcmd.LoadFromFile(filename)
	if err != nil {
		return nil, err
	}

	err = clientcmd.Validate(*config)
	if err != nil {
		return nil, err
	}

	return &Config{
		Filename: path.Base(filename),
		Config:   config,
	}, nil
}
