package model

import "e.coding.net/itdesk/weixin/golib/utils"

type Company struct {
	Id         uint32 `gorm:"column:id"`
	GroupId    uint32 `gorm:"column:group_id"`
	Type       uint32 `gorm:"column:type"`
	Name       string `gorm:"column:name"`
	FullName   string `gorm:"column:fullname"`
	ManagerId  uint32 `gorm:"column:manager_id"`
	CreateDate string `gorm:"column:CreateDate"`
	UpdateDate string `gorm:"column:UpdateDate"`
	OrderNo    uint32 `gorm:"column:OrderNo"`
	State      uint32 `gorm:"column:state"`
}

// 查询详细信息
func GetCompanyDetail(id uint32) (*Company, error) {
	var res Company
	err := utils.MysqlIns().Table("ti_company").Where("id=?", id).First(&res).Error
	if err != nil {
		return nil, err
	}
	return &res, nil
}
