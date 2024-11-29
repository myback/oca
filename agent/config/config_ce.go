// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build !consulent
// +build !consulent

package config

import (
	"github.com/myback/oca/acl"
	"github.com/myback/oca/agent/structs"
)

// EnterpriseMeta stub
type EnterpriseMeta struct{}

func (_ *EnterpriseMeta) ToStructs() acl.EnterpriseMeta {
	return *structs.DefaultEnterpriseMetaInDefaultPartition()
}
