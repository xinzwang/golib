package fast

type Group struct {
	Id         int    `gorm:"column:id"`
	Name       string `gorm:"column:name"`
	Note       string `gorm:"column:note"`
	Verify     int    `gorm:"column:verify"`
	UserId     int    `gorm:"column:user_id"`
	CreateDate string `gorm:"column:CreateDate"`
	State      int    `gorm:"column:state"`
}

type GroupUser struct {
	Id         int     `gorm:"column:id"`
	GroupId    int     `gorm:"column:group_id"`
	UserId     int     `gorm:"column:user_id"`
	Role       int     `gorm:"column:role"`
	CreateDate *string `gorm:"column:CreateDate"`
	QuitDate   *string `gorm:"column:QuitDate"`
	State      int     `gorm:"column:state"`
}
