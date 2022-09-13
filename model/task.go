package model

import (
	"time"

	"e.coding.net/itdesk/weixin/golib/utils"
)

type Task struct {
	Id         int     `gorm:"column:id"`
	TypeId     int     `gorm:"column:type_id"`
	UserId     int     `gorm:"column:user_id"`
	Text       *string `gorm:"column:text"`
	Urgent     int     `gorm:"column:urgent"`
	QrCode     *string `gorm:"column:qr_code"`
	Pic1       *string `gorm:"column:pic1"`
	Pic2       *string `gorm:"column:pic2"`
	Pic3       *string `gorm:"column:pic3"`
	Pic4       *string `gorm:"column:pic4"`
	Pic5       *string `gorm:"column:pic5"`
	Voice      *string `gorm:"column:voice"`
	CreateDate *string `gorm:"column:CreateDate"`
	CreateMode int     `gorm:"column:create_mode"`
}

func PassTask(taskId uint32, engineerId uint32, reason string) error {
	timeNow := time.Now().Format("2006-01-02 15:04:05")
	ins := utils.MysqlIns().Begin()
	var log TaskLog
	if err := ins.Table("tf_task_log").Where("task_id=?", taskId).Order("out DESC").Find(&log).Error; err != nil {
		ins.Rollback()
		return err
	}
	// 关闭旧流程
	if err := ins.Table("tf_task_log").Where("id=?", log.Id).Update(map[string]interface{}{
		"FinishDate":   timeNow,
		"state":        3, //转出
		"out_reason":   reason,
		"process_time": utils.TimeSub(timeNow, *log.CreateDate),
	}).Error; err != nil {
		ins.Rollback()
		return err
	}
	// 创建新流程
	if err := CreateTaskLog(&TaskLog{
		Id:          0,
		TaskId:      taskId,
		EngineerId:  engineerId,
		State:       0,
		Mode:        2, //转交
		CreateDate:  &timeNow,
		Out:         log.Out + 1,
		ProcessTime: 0,
	}); err != nil {
		ins.Rollback()
		return err
	}
	return nil
}
