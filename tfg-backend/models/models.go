package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	// Id       uint   `json:"id"`
	Name     string   `json:"name"`
	Email    string   `json:"email" gorm:"unique"`
	Password []byte   `json:"-"`
	Posts    []Post   `json:"-"`
	Threads  []Thread `json:"-"`
	Forums   []Forum  `json:"-"`
}

type Forum struct {
	gorm.Model
	Name    string   `json:"name" gorm:"not null" form:"name"`
	UserID  uint     `json:"userid" gorm:"not null"`
	User    User     `json:"-"`
	Threads []Thread `json:"-"`
}

type Thread struct {
	gorm.Model
	Title   string `json:"title" gorm:"not null"`
	Body    string `json:"body" gorm:"not null"`
	UserID  uint   `json:"userid" gorm:"not null"`
	User    User   `json:"-"`
	ForumID uint   `json:"forum_id" gorm:"not null"`
	Forum   Forum  `json:"-"`
	Posts   []Post `json:"-"`
}

type Post struct {
	gorm.Model
	Body     string `json:"body" gorm:"not null"`
	UserID   uint   `json:"userid" gorm:"not null"`
	User     User   `json:"-"`
	ThreadID uint   `json:"thread_id" gorm:"not null"`
	Thread   Thread `json:"-"`
	// Thread Thread //`json:"thread"`
}
