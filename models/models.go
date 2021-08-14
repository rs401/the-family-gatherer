package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	DisplayName string `json:"display_name" gorm:"unique;not null"`
	Email       string `json:"email" gorm:"unique;not null"`
	Password    []byte `json:"-"`
	Posts       []Post
	Threads     []Thread
	Forums      []Forum
}

type Forum struct {
	gorm.Model
	Name    string `json:"name" gorm:"not null" form:"name"`
	UserID  uint   `json:"userid" gorm:"not null"`
	User    User
	Threads []Thread
}

type Thread struct {
	gorm.Model
	Title   string `json:"title" gorm:"not null"`
	Body    string `json:"body" gorm:"not null"`
	UserID  uint   `json:"userid" gorm:"not null"`
	User    User
	ForumID uint `json:"forum_id" gorm:"not null"`
	Forum   Forum
	Posts   []Post `json:""`
}

type Post struct {
	gorm.Model
	Body     string `json:"body" gorm:"not null"`
	UserID   uint   `json:"userid" gorm:"not null"`
	User     User
	ThreadID uint `json:"thread_id" gorm:"not null"`
	Thread   Thread
	// Thread Thread //`json:"thread"`
}

type Session struct {
	gorm.Model
	UserID uint `json:"userid" gorm:"not null"`
	Email  string
}
