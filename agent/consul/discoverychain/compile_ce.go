// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build !consulent
// +build !consulent

package discoverychain

import (
	"github.com/myback/oca/acl"
	"github.com/myback/oca/agent/structs"
)

func (c *compiler) GetEnterpriseMeta() *acl.EnterpriseMeta {
	return structs.DefaultEnterpriseMetaInDefaultPartition()
}
