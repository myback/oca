// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build !consulent
// +build !consulent

package pbservice

import (
	fuzz "github.com/google/gofuzz"

	"github.com/myback/oca/acl"
)

func randEnterpriseMeta(_ *acl.EnterpriseMeta, _ fuzz.Continue) {
}
