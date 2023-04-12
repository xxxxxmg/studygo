package model

type User struct {
	Id       int    `gorm:"column:id" json:"id"`             //id
	Email    string `grom:"column:email" json:"email"`       //邮箱
	Password string `grom:"column:password" json:"password"` //密码
}
