package model

type User struct {
	Id           int    `gorm:"column:id"`
	OpenId       string `gorm:"column:openid"`
	Username     string `gorm:"column:username"`
	IdCard       string `gorm:"column:id_card"`
	Mobile       string `gorm:"column:mobile"`
	Photo        string `gorm:"column:photo"`
	WorkPlace    string `gorm:"column:work_place"`
	ComputerInfo string `gorm:"column:computer_info"`
	ExtNum       string `gorm:"column:ext_num"`
	Engineer     int    `gorm:"column:engineer"`
	PositionId   int    `gorm:"column:position_id"`
	DepartmentId int    `gorm:"column:department_id"`
	MenuId       string `gorm:"column:menu_id"`
	CreateDate   string `gorm:"column:CreateDate"`
	UpdateDate   string `gorm:"column:UpdateDate"`
	Ip           string `gorm:"column:ip"`
	Email        string `gorm:"column:email"`
	State        int    `gorm:"column:state"`

	Password string `gorm:"column:password"`
}
