package model

import "e.coding.net/itdesk/weixin/golib/utils"

type Task struct {
	Id         uint32  `gorm:"column:id"`
	TypeId     uint32  `gorm:"column:type_id"`
	UserId     uint32  `gorm:"column:user_id"`
	Text       *string `gorm:"column:text"`
	Urgent     bool    `gorm:"column:urgent"`
	QrCode     *string `gorm:"column:qr_code"`
	Pic1       *string `gorm:"column:pic1"`
	Pic2       *string `gorm:"column:pic2"`
	Pic3       *string `gorm:"column:pic3"`
	Pic4       *string `gorm:"column:pic4"`
	Pic5       *string `gorm:"column:pic5"`
	Voice      *string `gorm:"column:voice"`
	CreateDate *string `gorm:"column:CreateDate"`
	CreateMode uint32  `gorm:"column:create_mode"`
}

func GetTaskDetail(id uint32) (*Task, error) {
	var res Task
	if err := utils.MysqlIns().Table("tf_task").Where("id=?", id).First(&res).Error; err != nil {
		return nil, err
	} else {
		return &res, nil
	}
}

func (a *Task) CollectPic() []string {
	var res []string
	if a.Pic1 != nil {
		res = append(res, *a.Pic1)
	}
	if a.Pic2 != nil {
		res = append(res, *a.Pic2)
	}
	if a.Pic3 != nil {
		res = append(res, *a.Pic3)
	}
	if a.Pic4 != nil {
		res = append(res, *a.Pic4)
	}
	if a.Pic5 != nil {
		res = append(res, *a.Pic5)
	}
	return res
}
