package main

import (
	"time"
)

// UserProfile model
type UserProfile struct {
	ID          uint      `gorm:"primary_key" json:"id"`
	Username    string    `gorm:"type:varchar(150);unique_index:uidx_name" json:"username,omitempty"`
	Password    string    `gorm:"type:varchar(128);default:'666'"`
	IsSuperUser bool      `gorm:"column:is_superuser;default:false"  json:"-"`
	IsStuff     bool      `gorm:"column:is_stuff;default:false"  json:"-"`
	IsActive    bool      `gorm:"column:is_active;default:false"  json:"-"`
	FirstName   string    `gorm:"column:first_name;type:varchar(30)" json:"first_name,omitempty"`
	LastName    string    `gorm:"column:last_name;type:varchar(150)" json:"last_name,omitempty"`
	Email       string    `gorm:"type:varchar(254)"`
	DateJoined  time.Time `gorm:"column:date_joined"  json:"-"`
	LastLogin   time.Time `gorm:"column:last_login"  json:"-"`
	Company     string    `gorm:"type:varchar(255)"`
	Telephone   string    `gorm:"type:varchar(11)"`
	ThirdApp    uint      `gorm:"not null;default:0" json:"-"`
	SecretKey   string    `gorm:"type:varchar(20)" json:"-"`
	LastActive  time.Time `gorm:"index"  json:"-"`
}

// TableName Set UserProfile's table name
func (UserProfile) TableName() string {
	return "users_userprofile"
}

// NewUserProfile create a user object by default
func NewUserProfile() *UserProfile {
	return &UserProfile{
		DateJoined: time.Now(),
	}
}

// IsActived return true if user is active,otherwise false
func (u UserProfile) IsActived() bool {

	return u.IsActive
}

// CheckPassword Check whether the user's password is the same as the given password
func (u UserProfile) CheckPassword(pw string) bool {

	return u.Password == pw
}

// UsernameColumnName Return the table field name corresponding to the username field
func (u UserProfile) UsernameColumnName() string {

	return "username"
}
