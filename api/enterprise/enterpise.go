package enterprise

import (
	"context"
	v1 "user/api/enterprise/v1"
)

type IEnterprise interface {
	CreateEnterprise(ctx context.Context, req *v1.CreateEnterpriseUserReq) (res *v1.CreateEnterpriseUserRes, err error)
	DeleteEnterprise(ctx context.Context, req *v1.DeleteEnterpriseUserReq) (res *v1.DeleteEnterpriseUserRes, err error)
	UpdateEnterprise(ctx context.Context, req *v1.UpdateEnterpriseUserReq) (res *v1.UpdateEnterpriseUserRes, err error)
	SelectEnterprise(ctx context.Context, req *v1.SelectEnterpriseUserReq) (res *v1.SelectEnterpriseUserRes, err error)
	SelectEnterpriseAll(ctx context.Context, req *v1.SelectEnterpriseUserAllReq) (res *v1.SelectEnterpriseUserAllRes, err error)
	SelectEnterprisePage(ctx context.Context, req *v1.SelectEnterpriseUserPageReq) (res *v1.SelectEnterpriseUserPageRes, err error)
}
