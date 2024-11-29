// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package controllers

import (
	"github.com/myback/oca/internal/catalog/internal/controllers/endpoints"
	"github.com/myback/oca/internal/catalog/internal/controllers/nodehealth"
	"github.com/myback/oca/internal/catalog/internal/controllers/workloadhealth"
	"github.com/myback/oca/internal/controller"
)

type Dependencies struct {
	WorkloadHealthNodeMapper workloadhealth.NodeMapper
	EndpointsWorkloadMapper  endpoints.WorkloadMapper
}

func Register(mgr *controller.Manager, deps Dependencies) {
	mgr.Register(nodehealth.NodeHealthController())
	mgr.Register(workloadhealth.WorkloadHealthController(deps.WorkloadHealthNodeMapper))
	mgr.Register(endpoints.ServiceEndpointsController(deps.EndpointsWorkloadMapper))
}
