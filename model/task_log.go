package model

import "e.coding.net/itdesk/weixin/golib/utils"

type TaskLog struct {
	Id            uint32  `gorm:"column:id"`
	TaskId        uint32  `gorm:"column:task_id"`
	EngineerId    uint32  `gorm:"column:engineer_id"`
	State         uint32  `gorm:"column:state"`
	Mode          uint32  `gorm:"column:mode"`
	CreateDate    *string `gorm:"column:CreateDate"`
	StartDate     *string `gorm:"column:StartDate"`
	FinishDate    *string `gorm:"column:FinishDate"`
	ProcessTime   uint32  `gorm:"column:process_time"`
	Out           uint32  `gorm:"column:out"`
	OutReason     *string `gorm:"column:out_reason"`
	PendingReason *string `gorm:"pending_reason"`
}

// 创建处理日志
func CreateTaskLog(task *TaskLog) error {
	return utils.MysqlIns().Table("tf_task_log").Create(task).Error
}

// 根据id查询日志
func GetTaskLogById(id uint32) (*TaskLog, error) {
	var res TaskLog
	if err := utils.MysqlIns().Table("tf_task_log").Where("id=?", id).First(&res).Error; err != nil {
		return nil, err
	} else {
		return &res, nil
	}
}

// 根据task_id查询最新的一条日志
func GetTaskLogByTaskId(taskId uint32) (*TaskLog, error) {
	var res TaskLog
	if err := utils.MysqlIns().Table("tf_task_log").Where("task_id=?", taskId).Order("out DESC").Find(&res).Error; err != nil {
		return nil, err
	} else {
		return &res, nil
	}
}
