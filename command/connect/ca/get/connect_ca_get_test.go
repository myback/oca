// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package get

import (
	"strings"
	"testing"

	"github.com/myback/oca/testrpc"

	"github.com/mitchellh/cli"
	"github.com/myback/oca/agent"
)

func TestConnectCAGetConfigCommand_noTabs(t *testing.T) {
	t.Parallel()
	if strings.ContainsRune(New(cli.NewMockUi()).Help(), '\t') {
		t.Fatal("help has tabs")
	}
}

func TestConnectCAGetConfigCommand(t *testing.T) {
	if testing.Short() {
		t.Skip("too slow for testing.Short")
	}

	t.Parallel()
	a := agent.NewTestAgent(t, ``)
	defer a.Shutdown()
	testrpc.WaitForTestAgent(t, a.RPC, "dc1")

	ui := cli.NewMockUi()
	c := New(ui)
	args := []string{"-http-addr=" + a.HTTPAddr()}

	code := c.Run(args)
	if code != 0 {
		t.Fatalf("bad: %d. %#v", code, ui.ErrorWriter.String())
	}
	output := strings.TrimSpace(ui.OutputWriter.String())
	if !strings.Contains(output, `"Provider": "consul"`) {
		t.Fatalf("bad: %s", output)
	}
}
