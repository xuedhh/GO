package enterprise

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"reflect"
	"user/internal/dao"
	"user/internal/model/do"
	enterprise2 "user/internal/model/enterprise"
	"user/internal/model/entity"
	enterprise1 "user/internal/service/enterprise"
)

type sEnterprise struct{}

func init() {
	enterprise1.RegisterEnterpriseService(New())
}

func New() enterprise1.IEnterprise {
	return &sEnterprise{}
}

//func (s *sEnterprise) SelectEnterprisePage(ctx context.Context, input enterprise2.SelectEnterprisePageInput) (items []entity.Enterprise, total int64, err error) {
//	//TODO implement me
//	//dao.Enterprise.Ctx(ctx).Limit(input.Limit, input.Offset).ScanAndCount(&items, &(int)total, true)
//	//panic("implement me")
//	var enterprise []*entity.Enterprise
//	err = dao.Enterprise.Ctx(ctx).OmitEmpty().
//		Limit(input.Offset, input.Limit).
//		ScanAndCount(&items, &total, true)
//	return
//}

func (s *sEnterprise) SelectEnterprisePage(ctx context.Context, input enterprise2.SelectEnterprisePageInput) (output enterprise2.SelectEnterprisePageOutput, err error) {
	//TODO implement me
	var enterprise []*entity.Enterprise

	//all, totalCount, err := dao.Enterprise.Ctx(ctx).Limit(input.Limit, input.Offset).AllAndCount(true)
	orm := dao.Enterprise.Ctx(ctx).Limit(input.Offset, input.Limit)
	err = orm.Scan(&enterprise)
	if err != nil {
		return output, err
	}
	count, err := orm.Count()

	output.Total = count
	output.Limit = input.Limit
	output.Offset = input.Offset
	output.Enterprise = enterprise
	return output, nil
}

func (s *sEnterprise) SelectEnterpriseAll(ctx context.Context, input enterprise2.SelectEnterpriseAllInput) (output enterprise2.SelectEnterpriseAllOutput, err error) {
	//TODO implement me
	var enterprise []*entity.Enterprise
	err = dao.Enterprise.Ctx(ctx).OrderDesc("id").Scan(&enterprise)
	if err != nil {
		return output, err
	}
	if len(enterprise) == 0 {
		return output, nil
		//return output, gerror.Newf("没有数据了")
	}
	output.Enterprise = enterprise
	return output, nil
}

func (s *sEnterprise) SelectEnterprise(ctx context.Context, input enterprise2.SelectEnterpriseInput) (output enterprise2.SelectEnterpriseOutput, err error) {
	//TODO implement me

	one, err := dao.Enterprise.Ctx(ctx).Where("id", input.Id).One()
	if err != nil {
		return output, err
	}
	if len(one) == 0 {
		return output, nil
	}
	//TODO 方式一 将从数据库中查询出来的数据，转换成map ，然后再去使用 - 开始
	//m := one.Map()
	//fmt.Println(m)
	//id := one[dao.Enterprise.Columns().Id]
	//desc := one[dao.Enterprise.Columns().EnterpriseDesc]
	//name := one[dao.Enterprise.Columns().EnterpriseName]
	//fmt.Println(id, desc, name)
	//TODO 方式一 将从数据库中查询出来的数据，转换成map ，然后再去使用 - 结束

	// TODO 方式二 将从数据库中查询的数据，转换成struct ，方便使用 - 开始
	//var enterprise entity.Enterprise
	//if err := gconv.Struct(one, &enterprise); err != nil {
	//	panic(err)
	//}
	//fmt.Println(enterprise.EnterpriseDesc, enterprise.EnterpriseName, enterprise.Id)
	// TODO 方式二 将从数据库中查询的数据，转换成struct ，方便使用 - 结束

	// TODO 方式三 查询的时候直接将结构体传入进去进行查询，出来结构体就能够使用  - 推荐使用这种方式 - 开始
	var enterprise *entity.Enterprise
	//one, err := dao.Enterprise.Ctx(ctx).Where(do.Enterprise{Id: input.Id}).One()
	err = dao.Enterprise.Ctx(ctx).Where(do.Enterprise{Id: input.Id}).Scan(&enterprise)
	//count, err := dao.Enterprise.Ctx(ctx).Where(do.Enterprise{Id: input.Id}).Count()
	//fmt.Println("查询条数", count)
	if err != nil {
		return output, err
	}
	//if len(one) == 0 {
	//	return output, nil
	//}
	//if one.IsEmpty() {
	//	return output, nil
	//}
	if enterprise == nil {
		return output, gerror.Newf("数据为空，未找到")
	}
	output.EnterpriseDesc = enterprise.EnterpriseDesc
	output.EnterpriseName = enterprise.EnterpriseName
	output.Id = int(enterprise.Id)
	output.CreateAt = enterprise.CreateAt
	output.UpdateAt = enterprise.UpdateAt
	// TODO 方式三 查询的时候直接将结构体传入进去进行查询，出来结构体就能够使用  - 推荐使用这种方式 - 结束
	return output, nil

}

func (s *sEnterprise) UpdateEnterprise(ctx context.Context, input enterprise2.UpdateEnterpriseInput) (output enterprise2.UpdateEnterpriseOutput, err error) {
	//TODO implement me
	// 根据用户ID查询用户的数据是否存在
	one, err := dao.Enterprise.Ctx(ctx).Where(do.Enterprise{Id: input.Id}).One()
	if err != nil {
		output.Id = 1
		return output, gerror.Newf("未找到记录")
	}

	EnterpriseName := one[dao.Enterprise.Columns().EnterpriseName]
	name := reflect.ValueOf(EnterpriseName)
	if name.String() == input.Name {
		return output, gerror.Newf("名称已经被占用")
	}
	fmt.Println(one)
	output.Id = 1
	// 进行修改
	_, err = dao.Enterprise.Ctx(ctx).Where("id", input.Id).Update(do.Enterprise{
		EnterpriseName: input.Name,
	})
	if err != nil {
		return output, gerror.Newf("修改失败")
	}
	output.Id = input.Id
	return output, nil
	//EnterpriseDesc := one["enterpriseName"]
	//fmt.Println(reflect.TypeOf(EnterpriseName))
}

func (s *sEnterprise) DeleteEnterprise(ctx context.Context, input enterprise2.DeleteEnterpriseInput) (output enterprise2.DeleteEnterpriseOutput, err error) {
	//TODO implement me
	//dao.Enterprise.Ctx(ctx).Where("id", input.Id).Delete()
	_, err = dao.Enterprise.Ctx(ctx).Delete(g.Map{"id": input.Id})
	//fmt.Println("删除的时候错误1", nil)
	//fmt.Println("删除的时候错误2", err)
	if err != nil {
		output.OK = false
		return output, gerror.Newf("删除失败，记录不存在")
	}
	//fmt.Println("结果", result)
	//id, err := result.LastInsertId()
	//fmt.Println("删除的id:", id)
	if err != nil {
		output.OK = false
		return output, gerror.Newf("删除失败,无法获取ID")
	}
	output.Id = input.Id
	output.OK = true
	return output, nil
}

func (s *sEnterprise) CreateEnterprise(ctx context.Context, input enterprise2.CreateEnterpriseInput) (output enterprise2.CreateEnterpriseOutput, err error) {
	enterprise := g.Model("enterprise")
	one, err := enterprise.Where(do.Enterprise{EnterpriseName: input.Name}).One()
	fmt.Println(one)
	if one != nil {
		return enterprise2.CreateEnterpriseOutput{}, gerror.Newf("企业名称被占用")
	}
	// 创建记录
	insert, err := g.Model("enterprise").Data(
		do.Enterprise{
			EnterpriseName: input.Name,
			EnterpriseDesc: input.Desc,
		},
	).Insert()
	if insert == nil {
		return enterprise2.CreateEnterpriseOutput{}, gerror.Newf("创建失败")
	}
	id, err := insert.LastInsertId()
	//TODO implement me
	output = enterprise2.CreateEnterpriseOutput{
		Id:   int(id),
		Name: input.Name,
		Desc: input.Desc,
	}
	return output, nil
}
