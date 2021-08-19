package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	// Id       uint   `json:"id"`
	Name     string   `json:"Name"`
	Email    string   `json:"Email" gorm:"unique"`
	Password []byte   `json:"-"`
	Posts    []Post   `json:"-"`
	Threads  []Thread `json:"-"`
	Forums   []Forum  `json:"-"`
}

type Forum struct {
	gorm.Model
	Name    string   `json:"Name" gorm:"not null" form:"name"`
	UserID  uint     `json:"UserId" gorm:"not null"`
	User    User     `json:"User"`
	Threads []Thread `json:"Threads"`
}

type Thread struct {
	gorm.Model
	Title   string `json:"Title" gorm:"not null"`
	Body    string `json:"Body" gorm:"not null"`
	UserID  uint   `json:"UserId" gorm:"not null"`
	User    User   `json:"User"`
	ForumID uint   `json:"ForumId" gorm:"not null"`
	Forum   Forum  `json:"-"`
	Posts   []Post `json:"Posts"`
}

type Post struct {
	gorm.Model
	Body     string `json:"Body" gorm:"not null"`
	UserID   uint   `json:"UserId" gorm:"not null"`
	User     User   `json:"User"`
	ThreadID uint   `json:"ThreadId" gorm:"not null"`
	Thread   Thread `json:"-"`
}
