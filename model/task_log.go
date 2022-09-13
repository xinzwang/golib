package model

import "e.coding.net/itdesk/weixin/golib/utils"

type TaskLog struct {
	Id          uint32  `gorm:"column:id"`
	TaskId      uint32  `gorm:"column:task_id"`
	EngineerId  uint32  `gorm:"column:engineer_id"`
	State       uint32  `gorm:"column:state"`
	Mode        uint32  `gorm:"column:mode"`
	CreateDate  *string `gorm:"column:CreateDate"`
	StartDate   *string `gorm:"column:StartDate"`
	FinishDate  *string `gorm:"column:FinishDate"`
	Out         uint32  `gorm:"column:out"`
	ProcessTime uint32  `gorm:"column:process_time"`
}

// 创建处理日志
func CreateTaskLog(task *TaskLog) error {
	return utils.MysqlIns().Table("tf_task_log").Create(task).Error
}
