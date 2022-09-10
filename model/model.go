// 数据库表对应的结构体

package model

type Task struct {
	Id         int    `gorm:"column:id"`
	TypeId     int    `gorm:"column:type_id"`
	UserId     int    `gorm:"column:user_id"`
	Text       string `gorm:"column:text"`
	Urgent     int    `gorm:"column:urgent"`
	QrCode     string `gorm:"column:qr_code"`
	Pic1       string `gorm:"column:pic1"`
	Pic2       string `gorm:"column:pic2"`
	Pic3       string `gorm:"column:pic3"`
	Pic4       string `gorm:"column:pic4"`
	Pic5       string `gorm:"column:pic5"`
	Voice      string `gorm:"column:voice"`
	CreateDate string `gorm:"column:CreateDate"`
}

type TaskLog struct {
	Id          int     `gorm:"column:id"`
	TaskId      int     `gorm:"column:task_id"`
	EngineerId  int     `gorm:"column:engineer_id"`
	State       int     `gorm:"column:state"`
	Mode        int     `gorm:"column:mode"`
	CreateDate  *string `gorm:"column:CreateDate"`
	StartDate   *string `gorm:"column:StartDate"`
	FinishDate  *string `gorm:"column:FinishDate"`
	Out         int     `gorm:"column:out"`
	ProcessTime int     `gorm:"column:process_time"`
}

type TaskSolution struct {
	Id          int    `gorm:"column:id"`
	TaskId      int    `gorm:"column:task_id"`
	LogId       int    `gorm:"column:log_id"`
	Text        string `gorm:"column:text"`
	Pic1        string `gorm:"column:pic1"`
	Pic2        string `gorm:"column:pic2"`
	Pic3        string `gorm:"column:pic3"`
	Pic4        string `gorm:"column:pic4"`
	Pic5        string `gorm:"column:pic5"`
	CreateDate  string `gorm:"column:CreateDate"`
	RelatedInfo string `gorm:"column:related_info"`
}

type TaskType struct {
	Id         int    `gorm:"column:id"`
	Name       string `gorm:"column:name"`
	FatherId   int    `gorm:"column:father_id"`
	OrderNo    int    `gorm:"column:OrderNo"`
	Createdate string `gorm:"column:createdate"`
	UpdateDate string `gorm:"column:Updatedate"`
	UserId     int    `gorm:"column:user_id"`
}

type TaskMenu struct { // 即将废弃
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
	LendId        int     `gorm:"column:id"`
	ResourceId    int     `gorm:"column:resource_id"`
	ResourceState int     `gorm:"column:state"`
	LendUserId    *int    `gorm:"column:lend_user_id"`
	ReceiveUserId *int    `gorm:"column:receive_user_id"`
	EngineerId    *int    `gorm:"column:engineer_id"`
	LendDate      *string `gorm:"column:LendDate"`
	ReceiveDate   *string `gorm:"column:ReceiveDate"`
	ReturnDate    *string `gorm:"column:ReturnDate"`
}

type Resource struct {
	Id         int    `gorm:"column:id"`
	CodeId     int    `gorm:"column:code_id"`
	Usecase    string `gorm:"column:usecase"`
	Picture    string `gorm:"column:pic"`
	UserId     int    `gorm:"column:user_id"`
	Value1     string `gorm:"column:value1"`
	Value2     string `gorm:"column:value2"`
	Value3     string `gorm:"column:value3"`
	Value4     string `gorm:"column:value4"`
	Value5     string `gorm:"column:value5"`
	Value6     string `gorm:"column:value6"`
	Value7     string `gorm:"column:value7"`
	Value8     string `gorm:"column:value8"`
	Value9     string `gorm:"column:value9"`
	Value10    string `gorm:"column:value10"`
	Value11    string `gorm:"column:value11"`
	CreateDate string `gorm:"column:CreateDate"`
	UpdateDate string `gorm:"column:UpdateDate"`
	CWId       int    `gorm:"column:cw_id"`
	CWState    int    `gorm:"column:cw_state"`
	State      int    `gorm:"column:state"`
}

type Company struct {
	Id         int    `gorm:"column:id"`
	Group      string `gorm:"column:group"`
	Tyue       int    `gorm:"column:type"`
	Name       string `gorm:"column:name"`
	FullName   string `gorm:"column:fullname"`
	ManagerId  string `gorm:"column:manager_id"`
	CreateDate string `gorm:"column:CreateDate"`
	UpdateDate string `gorm:"column:UpdateDate"`
	OrderNo    int    `gorm:"column:OrderNo"`
	State      int    `gorm:"column:state"`
}

type Department struct {
	Id         int    `gorm:"column:id"`
	Name       string `gorm:"column:name"`
	ManagerId  string `gorm:"column:manager_id"`
	FatherId   int    `gorm:"column:father_id"`
	CompanyId  int    `gorm:"column:company_id"`
	CreateDate string `gorm:"column:CreateDate"`
	UpdateDate string `gorm:"column:UpdateDate"`
	OrderNo    int    `gorm:"column:OrderNo"`
	State      int    `gorm:"column:state"`
}

type Position struct {
	Id         int    `gorm:"column:id"`
	Name       string `gorm:"column:name"`
	CreateDate string `gorm:"column:CreateDate"`
	UpdateDate string `gorm:"column:UpdateDate"`
}

type AuthenticationLog struct {
	OpenId   string `gorm:"column:openid"`
	Username string `gorm:"column:username"`
	DepId    int    `gorm:"column:dep_id"`
	Mobile   string `gorm:"column:mobile"`
	IdCard   string `gorm:"column:id_card"`
}

type Group struct {
	GroupId      int    `gorm:"column:group_id"`
	GroupName    string `gorm:"column:group_name"`
	GroupNote    string `gorm:"column:group_note"`
	GroupVerify  int    `gorm:"column:group_verify"`
	CreateUserId string `gorm:"column:create_userid"`
	CreateDate   string `gorm:"column:CreateDate"`
	State        int    `gorm:"column:state"`
}

type GroupUser struct {
	Id         int    `gorm:"column:id"`
	GroupId    int    `gorm:"column:group_id"`
	UserId     string `gorm:"column:user_id"`
	Admin      int    `gorm:"column:admin"`
	Userstate  int    `gorm:"column:userstate"`
	CreateDate string `gorm:"column:createdate"`
	DeleteDate string `gorm:"column:deletedate"`
	State      int    `gorm:"column:state"`
}

type GroupMessage struct {
	Id          int    `gorm:"column:id"`
	UserId      string `gorm:"column:user_id"`
	Title       string `gorm:"column:title"`
	Abstract    string `gorm:"column:abstract"`
	Content     string `gorm:"column:content"`
	SendGroup   int    `gorm:"column:send_group"`
	Form        string `gorm:"column:form"`
	Choice      string `gorm:"column:choice"`
	Visible     int    `gorm:"column:visible"`
	RemindDelay int    `gorm:"column:remind_delay"`
	AutoClose   int    `gorm:"column:autoclose"`
	Repeate     int    `gorm:"column:repeate"`
	SendDate    string `gorm:"column:SendDate"`
	CreateDate  string `gorm:"column:CreateDate"`
	CloseDate   string `gorm:"column:CloseDate"`
	Draft       int    `gorm:"column:draft"`
}

type MessageResponse struct {
	Id         int    `gorm:"column:id"`
	MessageId  int    `gorm:"column:message_id"`
	UserId     string `gorm:"column:user_id"`
	Choice     string `gorm:"column:choice"`
	Form       string `gorm:"column:form"`
	SubmitTime string `gorm:"column:SubmitTime"`
}
