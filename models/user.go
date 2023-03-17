package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name      string `gorm:"column:name;type:varchar(100)" json:"name"`        //用户名
	Password  string `gorm:"column:password;type:varchar(32)" json:"password"` //密码
	Identity  string `gorm:"column:identity;type:varchar(36)" json:"identity"` //用户的唯一标识
	Phone     string `gorm:"column:phone;type:varchar(20)" json:"phone"`       //手机号
	Email     string `gorm:"column:email;type:varchar(100)" json:"email"`      //邮箱
	PassNum   int    `gorm:"column:pass_num;type:int(11)" json:"pass_num"`
	SubmitNum int    `gorm:"column:submit_num;type:int(11)" json:"submit_num"`
	IsAdmin   int    `gorm:"column:is_admin;type:tinyint(1);" json:"is_admin"`
}

func (table *User) TableName() string {
	return "user"
}
