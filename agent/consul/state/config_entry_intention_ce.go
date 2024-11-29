// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build !consulent
// +build !consulent

package state

import (
	"github.com/myback/oca/acl"
	"github.com/myback/oca/agent/structs"
)

func getIntentionPrecedenceMatchServiceNames(serviceName string, entMeta *acl.EnterpriseMeta) []structs.ServiceName {
	if serviceName == structs.WildcardSpecifier {
		return []structs.ServiceName{
			structs.NewServiceName(structs.WildcardSpecifier, entMeta),
		}
	}

	return []structs.ServiceName{
		structs.NewServiceName(serviceName, entMeta),
		structs.NewServiceName(structs.WildcardSpecifier, entMeta),
	}
}
