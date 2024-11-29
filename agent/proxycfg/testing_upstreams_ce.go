// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build !consulent
// +build !consulent

package proxycfg

import (
	"github.com/mitchellh/go-testing-interface"

	"github.com/myback/oca/acl"
	"github.com/myback/oca/agent/structs"
	"github.com/myback/oca/proto/private/pbpeering"
)

func extraDiscoChainConfig(t testing.T, variation string, entMeta acl.EnterpriseMeta) ([]structs.ConfigEntry, []*pbpeering.Peering) {
	t.Fatalf("unexpected variation: %q", variation)
	return nil, nil
}

func extraUpdateEvents(t testing.T, variation string, dbUID UpstreamID) []UpdateEvent {
	t.Fatalf("unexpected variation: %q", variation)
	return nil
}
