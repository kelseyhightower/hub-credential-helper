// Copyright 2017 Google Inc. All Rights Reserved.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.

package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v2"
)

type configs map[string][]config

type config struct {
	OAuthToken string `yaml:"oauth_token"`
	Protocol   string `yaml:"protocol"`
	User       string `yaml:"user"`
}

func main() {
	log.SetFlags(0)

	if len(os.Args) != 2 {
		log.Fatalf("Usage: %s <operation>", os.Args[0])
	}

	operation := os.Args[1]

	// Silently ignore non-get operations.
	if operation != "get" {
		os.Exit(0)
	}

	var host string
	var protocol string

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "host=") {
			s := strings.SplitN(line, "=", 2)
			host = s[1]
			continue
		}

		if strings.HasPrefix(line, "protocol=") {
			s := strings.SplitN(line, "=", 2)
			protocol = s[1]
		}
	}

	configFile := os.Getenv("HUB_CONFIG")
	if configFile == "" {
		configFile = filepath.Join(os.Getenv("HOME"), ".config", "hub")
	}

	data, err := ioutil.ReadFile(configFile)
	if err != nil {
		log.Fatalf("error loading the hub configuration file: %s", err)
	}

	var c configs
	err = yaml.Unmarshal(data, &c)
	if err != nil {
		log.Fatalf("error loading the hub configuration file: %s", err)
	}

	configs, ok := c[host]
	if !ok {
		log.Fatalf("host %s configuration not found", host)
	}

	for _, c := range configs {
		if c.Protocol == protocol {
			fmt.Printf("username=%s\n", c.OAuthToken)
			fmt.Printf("password=%s\n", "x-oauth-basic")
			os.Exit(0)
		}
	}

	log.Fatalf("protocol %s configuration not found", protocol)
}
