// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package raft

import (
	"github.com/mitchellh/cli"
	"github.com/myback/oca/command/flags"
)

func New() *cmd {
	return &cmd{}
}

type cmd struct{}

func (c *cmd) Run(args []string) int {
	return cli.RunResultHelp
}

func (c *cmd) Synopsis() string {
	return synopsis
}

func (c *cmd) Help() string {
	return flags.Usage(help, nil)
}

const synopsis = "Provides cluster-level tools for Consul operators"
const help = `
Usage: consul operator raft <subcommand> [options]

The Raft operator command is used to interact with Consul's Raft subsystem. The
command can be used to verify Raft peers or in rare cases to recover quorum by
removing invalid peers.
`
