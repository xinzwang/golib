package model

import "e.coding.net/itdesk/weixin/golib/utils"

type RepairType struct {
	Id         uint32 `gorm:"column:id"`
	Name       string `gorm:"column:name"`
	FatherId   uint32 `gorm:"column:father_id"`
	OrderNo    uint32 `gorm:"column:OrderNo"`
	CreateDate string `gorm:"column:createdate"`
	UpdateDate string `gorm:"column:Updatedate"`
	UserId     uint32 `gorm:"column:user_id"`
}

func GetRepairTypeDetail(id uint32) (*RepairType, error) {
	var res RepairType
	if err := utils.MysqlIns().Table("tf_task_type").Where("id=?", id).First(&res).Error; err != nil {
		return nil, err
	} else {
		return &res, nil
	}
}
