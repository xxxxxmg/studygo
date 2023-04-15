package main

import (
	"fmt"
	"math"
	"user/db"
	"user/model"
)

func init() {
	//调用db.InitMysql方法，并把数据传给host
	db.InitMysql("root:root@tcp(192.168.100.44:3306)/tongzi?charset=utf8&parseTime=True&loc=Local")
}

func main() {

	var userfriend []model.Userfriend

	//新增数据
	// db.Mysql().Create(&model.User{
	// 	Email:    "zhousan@163.com",
	// 	Password: "888888",
	// })
	//修改数据
	//db.Mysql().Model(user).Debug().Where("email = ?", "lisi@163.com").Update("email", "wangwu@163.com")
	//查询所有
	db.Mysql().Find(&userfriend)
	//查询单条数据
	//db.Mysql().Debug().Where("email=?", "zhousan@163.com").First(&user)
	//修改数据
	//db.Mysql().Model(user).Debug().Where("email = ?", "lisi@163.com").Update("email", "wangwu@163.com")
	//删除指定数据
	//db.Mysql().Debug().Where("email", "zhaoliu@163.com").Delete(&user)

	if len(userfriend) >= 10 {
		var tempid int = 0
		var tempCreationTime int = math.MaxInt
		for _, v := range userfriend {
			fmt.Println(v.CreationTime)
			if tempCreationTime > v.CreationTime {
				tempCreationTime = v.CreationTime
				tempid = v.Id
			}
		}
		fmt.Println(tempid)
		fmt.Println(tempCreationTime)
		fmt.Println("这条数据将被删除", tempid)
		db.Mysql().Where("id = ?", tempid).Delete(&userfriend)
	}
}
