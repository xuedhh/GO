package enterprise

import (
	"context"
	"user/internal/model/enterprise"
)

type IEnterprise interface {
	CreateEnterprise(ctx context.Context, input enterprise.CreateEnterpriseInput) (output enterprise.CreateEnterpriseOutput, err error)
	DeleteEnterprise(ctx context.Context, input enterprise.DeleteEnterpriseInput) (output enterprise.DeleteEnterpriseOutput, err error)
	UpdateEnterprise(ctx context.Context, input enterprise.UpdateEnterpriseInput) (output enterprise.UpdateEnterpriseOutput, err error)
	SelectEnterprise(ctx context.Context, input enterprise.SelectEnterpriseInput) (output enterprise.SelectEnterpriseOutput, err error)
	SelectEnterpriseAll(ctx context.Context, input enterprise.SelectEnterpriseAllInput) (output enterprise.SelectEnterpriseAllOutput, err error)
	SelectEnterprisePage(ctx context.Context, input enterprise.SelectEnterprisePageInput) (output enterprise.SelectEnterprisePageOutput, err error)
}

var enterpriseService IEnterprise

func Enterprise() IEnterprise {
	if enterpriseService == nil {
		panic("enterpriseService not init")
	}
	return enterpriseService
}

func RegisterEnterpriseService(service IEnterprise) {
	enterpriseService = service
}
