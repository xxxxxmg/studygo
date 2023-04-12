package main

import (
	"fmt"
	"user/db"
	"user/model"
)

func init() {
	//调用db.InitMysql方法，并把数据传给host
	db.InitMysql("root:root@tcp(192.168.100.44:3306)/tongzi?charset=utf8&parseTime=True&loc=Local")
}

func main() {

	var user []model.User

	//新增数据
	// db.Mysql().Create(&model.User{
	// 	Email:    "zhousan@163.com",
	// 	Password: "888888",
	// })
	//修改数据
	//db.Mysql().Model(user).Debug().Where("email = ?", "lisi@163.com").Update("email", "wangwu@163.com")
	//查询所有
	//db.Mysql().Find(&user)
	//查询单条数据
	db.Mysql().Debug().Where("email=?", "zhousan@163.com").Find(&user)
	//修改数据
	//db.Mysql().Model(user).Debug().Where("email = ?", "lisi@163.com").Update("email", "wangwu@163.com")
	//删除指定数据
	//db.Mysql().Debug().Where("email", "zhaoliu@163.com").Delete(&user)

	fmt.Println(user)
}
