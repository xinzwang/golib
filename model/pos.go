package model

import "e.coding.net/itdesk/weixin/golib/utils"

// 职位 position
type Pos struct {
	Id         uint32 `gorm:"column:id"`
	Name       string `gorm:"column:name"`
	CreateDate string `gorm:"column:CreateDate"`
	UpdateDate string `gorm:"column:UpdateDate"`
	State      uint32 `gorm:"column:state"`
}

// 查询详细信息
func GetPosDetail(id uint32) (*Pos, error) {
	var res Pos
	err := utils.MysqlIns().Table("ti_position").Where("id=?", id).First(&res).Error
	if err != nil {
		return nil, err
	}
	return &res, nil
}

func GetPosList() ([]*Pos, error) {
	var res []*Pos
	err := utils.MysqlIns().Table("ti_position").Find(&res).Error
	if err != nil {
		return nil, err
	}
	return res, nil
}
