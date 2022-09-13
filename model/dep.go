package model

import "e.coding.net/itdesk/weixin/golib/utils"

// 部门 department
type Dep struct {
	Id         uint32 `gorm:"column:id"`
	Name       string `gorm:"column:name"`
	ManagerId  string `gorm:"column:manager_id"`
	FatherId   uint32 `gorm:"column:father_id"`
	CompanyId  uint32 `gorm:"column:company_id"`
	CreateDate string `gorm:"column:CreateDate"`
	UpdateDate string `gorm:"column:UpdateDate"`
	OrderNo    uint32 `gorm:"column:OrderNo"`
	State      uint32 `gorm:"column:state"`
}

// 查询详细信息
func GetDepDetail(id uint32) (*Dep, error) {
	var res Dep
	err := utils.MysqlIns().Table("ti_department").Where("id=?", id).First(&res).Error
	if err != nil {
		return nil, err
	}
	return &res, nil
}
