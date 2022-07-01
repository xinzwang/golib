package model

type User struct {
	OpenId     string `gorm:"column:openid"`
	IdCard     string `gorm:"column:id_card"`
	UserId     string `gorm:"column:user_id"`
	Username   string `gorm:"column:username"`
	Password   string `gorm:"column:password"`
	PositionId int    `gorm:"column:position_id"`
	Mobile     string `gorm:"column:mobile"`
	DepId      string `gorm:"column:dep_id"`
	CreateDate string `gorm:"column:create_date"`
	UpdateDate string `gorm:"column:update_date"`
	UserSate   int    `gorm:"column:usersate"`
	Engineer   int    `gorm:"column:engineer"`
	menuId     string `gorm:"column:menu_id"`
	Photo      string `gorm:"column:photo"`
}

type TaskList struct {
	Id           int    `gorm:"column:id"`
	TaskId       string `gorm:"column:task_id"`
	TypeId       string `gorm:"column:type_id"`
	UserId       string `gorm:"column:user_id"`
	TaskProblem  string `gorm:"column:task_problem"`
	TaskUserPic1 string `gorm:"column:task_userpic1"`
	TaskUserPic2 string `gorm:"column:task_userpic2"`
	TaskUserPic3 string `gorm:"column:task_userpic3"`
	TaskUserPic4 string `gorm:"column:task_userpic4"`
	TaskUserPic5 string `gorm:"column:task_userpic5"`
	TaskVoice    string `gorm:"column:task_voice"`
	CreateDate   string `gorm:"column:CreateDate"`
	EngineerId   string `gorm:"column:engineer_id"`
	Urgent       bool   `gorm:"column:urgent"`
	QrCode       string `gorm:"column:qr_code"`
}

type TaskLog struct {
	Id          int    `gorm:"column:id"`
	ActivityId  string `gorm:"column:activity_id"`
	TaskId      string `gorm:"column:task_id"`
	TaskState   int    `gorm:"column:task_state"`
	EngineerId  string `gorm:"column:engineer_id"`
	TaskOut     int    `gorm:"column:task_out"`
	TaskLogDate string `gorm:"column:task_logdate"`
	ProcessTime int    `gorm:"column:process_time"`
}

type TaskWay struct {
	Id          int    `gorm:"column:id"`
	SolutionId  string `gorm:"column:solution_id"`
	TaskId      string `gorm:"column:task_id"`
	WayContent  string `gorm:"column:way_content"`
	WayPic1     string `gorm:"column:way_pic1"`
	WayPic2     string `gorm:"column:way_pic2"`
	WayPic3     string `gorm:"column:way_pic3"`
	WayPic4     string `gorm:"column:way_pic4"`
	WayPic5     string `gorm:"column:way_pic5"`
	RelatedInfo string `gorm:"column:related_info"`
	CreateDate  string `gorm:"column:CreateDate"`
}

type TaskMenu struct {
	Id         int    `gorm:"column:id"`
	TypeId     string `gorm:"column:type_id"`
	TaskName   string `gorm:"column:task_name"`
	TaskRid    string `gorm:"column:task_Rid"`
	OrderNo    int    `gorm:"column:OrderNo"`
	Createdate string `gorm:"column:createdate"`
	NewDate    string `gorm:"column:newdate"`
	UserId     string `gorm:"column:user_id"`
}

type LendLog struct {
	LendId        string `gorm:"column:lend_id"`
	ResourceId    string `gorm:"column:resource_id"`
	ResourceState int    `gorm:"column:resource_state"`
	LendUserId    string `gorm:"column:lend_user_id"`
	LendDate      string `gorm:"column:lend_date"`
	EngineerId    string `gorm:"column:engineer_id"`
	ReceiveUserId string `gorm:"column:receive_user_id"`
	ReceiveDate   string `gorm:"column:receive_date"`
	ReturnDate    string `gorm:"column:return_date"`
}

type Resource struct {
	ResourceId      string `gorm:"column:resource_id"`
	CodeId          string `gorm:"column:code_id"`
	ResourceValue1  string `gorm:"column:resource_value_1"`
	ResourceValue2  string `gorm:"column:resource_value_2"`
	ResourceValue3  string `gorm:"column:resource_value_3"`
	ResourceValue4  string `gorm:"column:resource_value_4"`
	ResourceValue5  string `gorm:"column:resource_value_5"`
	ResourceValue6  string `gorm:"column:resource_value_6"`
	ResourceValue7  string `gorm:"column:resource_value_7"`
	ResourceValue8  string `gorm:"column:resource_value_8"`
	ResourceValue9  string `gorm:"column:resource_value_9"`
	ResourceValue10 string `gorm:"column:resource_value_10"`
	ResourceValue11 string `gorm:"column:resource_value_11"`
	UserId          string `gorm:"column:user_id"`
	Usecase         string `gorm:"column:usecase"`
	CreateDate      string `gorm:"column:createdate"`
	NewDate         string `gorm:"column:newdate"`
	ResourceState   int    `gorm:"column:resource_state"`
	ResourceCWId    string `gorm:"column:resource_cw_id"`
	ResourceCWState int    `gorm:"column:resource_cw_state"`
}

type Department struct {
	DepId        string `gorm:"column:dep_id"`
	DepGroup     string `gorm:"column:dep_group"`
	DepTyue      string `gorm:"column:dep_type"`
	DepName      string `gorm:"column:dep_name"`
	DepName1     string `gorm:"column:dep_name_1"`
	DepName2     string `gorm:"column:dep_name_2"`
	DepName3     string `gorm:"column:dep_name_3"`
	DepName4     string `gorm:"column:dep_name_4"`
	CreateDate   string `gorm:"column:CreateDate"`
	State        string `gorm:"column:dep_state"`
	OrderNo      string `gorm:"column:OrderNo"`
	NewDate      string `gorm:"column:newdate"`
	DepFullName  string `gorm:"column:dep_fullname"`
	DepManagerId string `gorm:"column:dep_manager_id"`
}

type AuthenticationLog struct {
	OpenId   string `gorm:"column:openid"`
	Username string `gorm:"column:username"`
	DepId    string `gorm:"column:dep_id"`
	Mobile   string `gorm:"column:mobile"`
	IdCard   string `gorm:"column:id_card"`
}
