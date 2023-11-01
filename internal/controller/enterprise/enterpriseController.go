package enterprise

import (
	"context"
	"user/api/enterprise"
	v1 "user/api/enterprise/v1"
	enterprise2 "user/internal/model/enterprise"
	enterpriseV1 "user/internal/service/enterprise"
)

type EnterpriseController struct{}

func (e *EnterpriseController) SelectEnterprisePage(ctx context.Context, req *v1.SelectEnterpriseUserPageReq) (res *v1.SelectEnterpriseUserPageRes, err error) {
	//TODO implement me
	page, err := enterpriseV1.Enterprise().SelectEnterprisePage(ctx, enterprise2.SelectEnterprisePageInput{
		Offset: req.Offset,
		Limit:  req.Limit,
	})
	if err != nil {
		return res, err
	}
	res = &v1.SelectEnterpriseUserPageRes{SelectEnterprisePageOutput: page}
	return res, nil
}

func (e *EnterpriseController) SelectEnterpriseAll(ctx context.Context, req *v1.SelectEnterpriseUserAllReq) (res *v1.SelectEnterpriseUserAllRes, err error) {
	//TODO implement me
	all, err := enterpriseV1.Enterprise().SelectEnterpriseAll(ctx, enterprise2.SelectEnterpriseAllInput{})
	if err != nil {
		return nil, err
	}
	allRes := &v1.SelectEnterpriseUserAllRes{
		all,
	}
	return allRes, nil
}

func (e *EnterpriseController) SelectEnterprise(ctx context.Context, req *v1.SelectEnterpriseUserReq) (res *v1.SelectEnterpriseUserRes, err error) {
	//TODO implement me
	selectEnterprise, err := enterpriseV1.Enterprise().SelectEnterprise(ctx, enterprise2.SelectEnterpriseInput{
		Id: req.Id,
	})
	if err != nil {
		return nil, err
	}
	res = &v1.SelectEnterpriseUserRes{
		Id:             selectEnterprise.Id,
		EnterpriseName: selectEnterprise.EnterpriseName,
		EnterpriseDesc: selectEnterprise.EnterpriseDesc,
		CreateAt:       selectEnterprise.CreateAt,
		UpdateAt:       selectEnterprise.UpdateAt,
	}
	return res, nil
}

func (e *EnterpriseController) UpdateEnterprise(ctx context.Context, req *v1.UpdateEnterpriseUserReq) (res *v1.UpdateEnterpriseUserRes, err error) {
	//TODO implement me
	updateEnterprise, err := enterpriseV1.Enterprise().UpdateEnterprise(ctx, enterprise2.UpdateEnterpriseInput{
		Id:   req.Id,
		Name: req.EnterpriseName,
	})
	if err != nil {
		return nil, err
	}
	userRes := &v1.UpdateEnterpriseUserRes{Id: updateEnterprise.Id}
	return userRes, nil
}

func (e *EnterpriseController) DeleteEnterprise(ctx context.Context, req *v1.DeleteEnterpriseUserReq) (res *v1.DeleteEnterpriseUserRes, err error) {
	//TODO implement me
	deleteEnterprise, err := enterpriseV1.Enterprise().DeleteEnterprise(ctx, enterprise2.DeleteEnterpriseInput{Id: req.Id})
	if err != nil {
		return nil, err
	}
	userRes := &v1.DeleteEnterpriseUserRes{
		Id:       deleteEnterprise.Id,
		IsDelete: deleteEnterprise.OK,
	}
	return userRes, nil
}

func Enterprise() enterprise.IEnterprise {
	return &EnterpriseController{}
}

func (e *EnterpriseController) CreateEnterprise(ctx context.Context, req *v1.CreateEnterpriseUserReq) (res *v1.CreateEnterpriseUserRes, err error) {
	createEnterprise, err := enterpriseV1.Enterprise().CreateEnterprise(ctx, enterprise2.CreateEnterpriseInput{
		Name: req.EnterpriseName,
		Desc: req.EnterpriseDesc,
	})
	if err != nil {
		return nil, err
	}
	id := createEnterprise.Id
	name := createEnterprise.Name
	desc := createEnterprise.Desc
	res = &v1.CreateEnterpriseUserRes{
		Id:             id,
		EnterpriseName: name,
		EnterpriseDesc: desc,
	}
	return res, nil
}
