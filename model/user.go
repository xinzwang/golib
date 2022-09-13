package model

import (
	"fmt"

	"e.coding.net/itdesk/weixin/golib/utils"
	"github.com/jinzhu/gorm"
)

type User struct {
	Id           uint32 `gorm:"column:id" json:"id"`
	OpenId       string `gorm:"column:openid" json:"openid"`
	Username     string `gorm:"column:username" json:"username"`
	IdCard       string `gorm:"column:id_card" json:"id_card"`
	Mobile       string `gorm:"column:mobile" json:"mobile"`
	Photo        string `gorm:"column:photo" json:"photo"`
	WorkPlace    string `gorm:"column:work_place" json:"work_place"`
	ComputerInfo string `gorm:"column:computer_info" json:"computer_info"`
	ExtNum       string `gorm:"column:ext_num" json:"ext_num"`
	Engineer     uint32 `gorm:"column:engineer" json:"engineer"`
	PositionId   uint32 `gorm:"column:position_id" json:"position_id"`
	DepartmentId uint32 `gorm:"column:department_id" json:"department_id"`
	MenuId       string `gorm:"column:menu_id" json:"menu_id"`
	CreateDate   string `gorm:"column:CreateDate" json:"CreateDate"`
	UpdateDate   string `gorm:"column:UpdateDate" json:"UpdateDate"`
	Ip           string `gorm:"column:ip" json:"ip"`
	Email        string `gorm:"column:email" json:"email"`
	State        uint32 `gorm:"column:state" json:"state"`
	Password     string `gorm:"column:password" json:"password"`
}

// 查询单个用户信息
func GetUser(id uint32) (*User, error) {
	var user User
	err := utils.MysqlIns().Table("ti_user").Where("id=?", id).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// 查询用户个数
func GetCount(username string, maps map[string]interface{}) (uint32, error) {
	var total uint32
	ins := listIns(username, maps)
	err := ins.Count(&total).Error
	if err != nil {
		return 0, err
	}
	return total, nil
}

// 查询用户列表
func GetUserList(page uint32, limit uint32, sort string, username string, maps map[string]interface{}) ([]*User, error) {
	var users []*User
	ins := listIns(username, maps)

	if len(sort) > 0 {
		sortFlag := "ASC"
		if sort[0] == '-' {
			sortFlag = "DESC"
		}
		sortKey := sort[1:]
		ins = ins.Order(fmt.Sprintf("%s %s", sortKey, sortFlag))
	}

	err := ins.Offset(page).Limit(limit).Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, err
}

// 新增用户
func AddUser(user *User) error {
	return nil
}

// 删除用户
func DeleteUser(id int) error {
	return nil
}

// 更新用户信息
func UpdateUser(data map[string]interface{}) error {
	return nil
}

func listIns(username string, maps map[string]interface{}) *gorm.DB {
	ins := utils.MysqlIns().Table("ti_user")
	if len(username) != 0 {
		ins = ins.Where("username LIKE ?", "%"+username+"%")
	}
	if maps != nil {
		ins = ins.Where(maps)
	}
	return ins
}
