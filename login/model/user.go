package model

type User struct {
	Id       int    `gorm:"column:id" json:"id"`             //id
	Email    string `gorm:"column:email" json:"email"`       //邮箱
	Password string `gorm:"column:password" json:"password"` //密码
}
