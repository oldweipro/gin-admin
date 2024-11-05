package example

import (
	"github.com/oldweipro/gin-admin/model/common/request"
	"github.com/oldweipro/gin-admin/model/example"
	"github.com/oldweipro/gin-admin/model/system"
	"github.com/oldweipro/gin-admin/pkg/app"
	systemService "github.com/oldweipro/gin-admin/service/system"
)

type CustomerService struct{}

var CustomerServiceApp = new(CustomerService)

//@author: [oldweipro](https://github.com/oldweipro)
//@function: CreateExaCustomer
//@description: 创建客户
//@param: e model.ExaCustomer
//@return: err error

func (exa *CustomerService) CreateExaCustomer(e example.ExaCustomer) (err error) {
	err = app.DBClient.Create(&e).Error
	return err
}

//@author: [oldweipro](https://github.com/oldweipro)
//@function: DeleteFileChunk
//@description: 删除客户
//@param: e model.ExaCustomer
//@return: err error

func (exa *CustomerService) DeleteExaCustomer(e example.ExaCustomer) (err error) {
	err = app.DBClient.Delete(&e).Error
	return err
}

//@author: [oldweipro](https://github.com/oldweipro)
//@function: UpdateExaCustomer
//@description: 更新客户
//@param: e *model.ExaCustomer
//@return: err error

func (exa *CustomerService) UpdateExaCustomer(e *example.ExaCustomer) (err error) {
	err = app.DBClient.Save(e).Error
	return err
}

//@author: [oldweipro](https://github.com/oldweipro)
//@function: GetExaCustomer
//@description: 获取客户信息
//@param: id uint
//@return: customer model.ExaCustomer, err error

func (exa *CustomerService) GetExaCustomer(id uint) (customer example.ExaCustomer, err error) {
	err = app.DBClient.Where("id = ?", id).First(&customer).Error
	return
}

//@author: [oldweipro](https://github.com/oldweipro)
//@function: GetCustomerInfoList
//@description: 分页获取客户列表
//@param: sysUserAuthorityID string, info request.PageInfo
//@return: list interface{}, total int64, err error

func (exa *CustomerService) GetCustomerInfoList(sysUserAuthorityID uint, info request.PageInfo) (list interface{}, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := app.DBClient.Model(&example.ExaCustomer{})
	var a system.SysAuthority
	a.AuthorityId = sysUserAuthorityID
	auth, err := systemService.AuthorityServiceApp.GetAuthorityInfo(a)
	if err != nil {
		return
	}
	var dataId []uint
	for _, v := range auth.DataAuthorityId {
		dataId = append(dataId, v.AuthorityId)
	}
	var CustomerList []example.ExaCustomer
	err = db.Where("sys_user_authority_id in ?", dataId).Count(&total).Error
	if err != nil {
		return CustomerList, total, err
	} else {
		err = db.Limit(limit).Offset(offset).Preload("SysUser").Where("sys_user_authority_id in ?", dataId).Find(&CustomerList).Error
	}
	return CustomerList, total, err
}
