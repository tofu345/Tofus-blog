package db

type Comment struct {
	BaseModel
	ID      int    `gorm:"primarykey" json:"id"`
	Author  string `json:"author"`
	Message string `json:"message"`
}
