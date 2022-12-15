package backend

import (
	"time"
)

type BaseModel struct {
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Post struct {
	ID     int    `gorm:"primarykey" json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
	Author string `json:"author"`
	Views  int    `json:"views"`
	Likes  int    `json:"likes"`
	BaseModel
}

type User struct {
	BaseModel
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Username  string `json:"username"`
	Email     string `json:"email"`
}
