package fast

type GroupItem struct {
	Id         int    `gorm:"column:id"`
	Name       string `gorm:"column:name"`
	Note       string `gorm:"column:note"`
	Verify     int    `gorm:"column:verify"`
	UserId     int    `gorm:"column:user_id"`
	CreateDate string `gorm:"column:CreateDate"`
	State      int    `gorm:"column:state"`
	Num        int    `gorm:"column:num"`
}

type GroupMember struct {
	UserId   int    `gorm:"column:user_id"`
	Username string `gorm:"column:username"`
	Role     int    `gorm:"column:role"`
	Photo    string `gorm:"column:photo"`
}
