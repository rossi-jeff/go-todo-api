type ToDo struct {
	BaseModel
	Task      string `gorm:"size:255"`
	Completed bool   `gorm:"default:false"`
	UserId    *uint
	User      *User
}