package model

import (
	"fmt"

	"e.coding.net/itdesk/weixin/golib/utils"
	"github.com/jinzhu/gorm"
)

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
	Mode       uint32  `gorm:"column:create_mode"`
}

func GetTaskCount(keyword string, maps map[string]interface{}) (uint32, error) {
	var total uint32
	ins := taskListIns(keyword, maps)
	if err := ins.Count(&total).Error; err != nil {
		return 0, err
	} else {
		return total, nil
	}
}

// 获取任务列表
func GetTaskList(page uint32, limit uint32, sort string, keyword string, maps map[string]interface{}) ([]*Task, error) {
	var tasks []*Task
	ins := taskListIns(keyword, maps)
	// sort
	if len(sort) > 0 {
		sortFlag := "ASC"
		if sort[0] == '-' {
			sortFlag = "DESC"
		}
		sortKey := sort[1:]
		ins = ins.Order(fmt.Sprintf("%s %s", sortKey, sortFlag))
	}
	// page
	if page > 0 && limit > 0 {
		ins = ins.Offset((page - 1) * limit).Limit(limit)
	}
	err := ins.Find(&tasks).Error
	if err != nil {
		return nil, err
	}
	return tasks, err
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

func taskListIns(keyword string, maps map[string]interface{}) *gorm.DB {
	ins := utils.MysqlIns().Table("tf_task")
	if len(keyword) > 0 {
		ins = ins.Where("text LIKE ?", "%"+keyword+"%")
	}
	if maps != nil {
		ins = ins.Where(maps)
	}
	return ins
}
