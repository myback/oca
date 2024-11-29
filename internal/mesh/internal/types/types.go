// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package types

import (
	"github.com/myback/oca/internal/resource"
)

const (
	GroupName       = "mesh"
	VersionV1Alpha1 = "v1alpha1"
	CurrentVersion  = VersionV1Alpha1
)

func Register(r resource.Registry) {
	RegisterProxyConfiguration(r)
	RegisterUpstreams(r)
}
