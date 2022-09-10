package repair

// View
type TaskItem struct {
	Id         int    `gorm:"column:id"`
	Text       string `gorm:"column:text"`
	Urgent     bool   `gorm:"column:urgent"`
	QrCode     string `gorm:"column:qr_code"`
	Pic1       string `gorm:"column:pic1"`
	Pic2       string `gorm:"column:pic2"`
	Pic3       string `gorm:"column:pic3"`
	Pic4       string `gorm:"column:pic4"`
	Pic5       string `gorm:"column:pic5"`
	Voice      string `gorm:"column:voice"`
	UserId     int    `gorm:"column:user_id"`
	UserName   string `gorm:"column:user_name"`
	CreateDate string `gorm:"column:CreateDate"`
	State      int    `gorm:"column:state"`

	TypeId       int    `gorm:"column:type_id"`
	TypeName     string `gorm:"column:type_name"`
	LogId        int    `gorm:"column:log_id"`
	Out          int    `gorm:"column:out"`
	EngineerId   int    `gorm:"column:engineer_id"`
	EngineerName string `gorm:"column:engineer_name"`

	SolutionId int `gorm:"column:solution_id"`
}
