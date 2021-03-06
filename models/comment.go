package models

import (
	//"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

type Comment struct {
	Post_id    int       `json:"post_id"`
	User_id    int       `json:"user_id"`
	Id         int       `json:"id" gorm:"primary_key;auto_increment;not_null"`
	Comment    string    `json:"comment"`
	Created_at time.Time `json:"created_at"`
}
