package main

import (
	"login/db"
	"login/model"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func init() {
	//调用db.InitMysql方法，并把数据传给host
	db.InitMysql("root:root@tcp(127.0.0.1:3306)/test?charset=utf8&parseTime=True&loc=Local")
}

//登录
func Login(c *gin.Context) {
	c.HTML(200, "login.html", nil)
}

//判断邮箱密码是否正确
func DoLogin(c *gin.Context) {
	//获取页面输入的邮箱
	email := c.PostForm("email")
	//获取页面输入的密码
	password := c.PostForm("password")
	//查询数据库中这个邮箱是否存在
	var user model.User
	db.Mysql().Where("email = ?", email).Find(&user)
	//判断数据是否正确
	if user.Email == email && user.Password == password {
		//数据写入到指定的HTML网页中，并响应200
		c.HTML(200, "welcomelogin.html", gin.H{
			"Email": email,
		})
	} else {
		c.String(200, "账号或者密码错误")
	}
}

//注册
func Regist(c *gin.Context) {
	//数据写入到指定的HTML网页中，并响应200
	c.HTML(200, "regist.html", nil)
}

//判断邮箱是否存在
func DoRegist(c *gin.Context) {
	//获取邮箱
	email := c.PostForm("email")
	//获取密码
	password := c.PostForm("password")
	//获取数据库数据
	var user model.User
	db.Mysql().Debug().Table("users").Where("email = ?", email).Find(&user)
	//判断邮箱是否已存在
	if email == user.Email {
		c.String(200, "邮箱已存在")
	} else {
		//不存在，写入数据库
		db.Mysql().Create(&model.User{
			Email:    email,
			Password: password,
		})
		//数据写入到指定的HTML网页中，并响应200
		c.HTML(200, "welcomeregist.html", gin.H{
			"Email":    email,
			"Password": password,
		})
	}
}

//查询所有用户
func GetAllUser(c *gin.Context) {
	var user []model.User
	db.Mysql().Debug().Table("users").Find(&user)
	c.String(200, "用户ID \t 用户Email\n")
	for _, v := range user {
		userID := strconv.Itoa(v.Id)
		c.String(200, userID+"\t\t"+v.Email+"\n")
	}
}

//添加好友
func AddFriend(c *gin.Context) {
	userid := c.PostForm("userid")
	friendid := c.PostForm("friendid")
	useid, _ := strconv.ParseInt(userid, 10, 64)
	fid, _ := strconv.ParseInt(friendid, 10, 64)
	//获取用户数据
	var use model.Userfriend
	db.Mysql().Debug().Where("user_id = ?", useid).Where("friend_id = ?", fid).First(&use)

	if use.FriendId != int(fid) {
		var user model.User
		db.Mysql().Debug().Table("users").Where("id = ?", fid).First(&user)
		if user.Id == int(fid) {
			//相互添加好友，增加两条数据
			db.Mysql().Debug().Create(&model.Userfriend{
				UserId:       int(useid),
				FriendId:     int(fid),
				CreationTime: int(time.Now().Unix()),
			})
			db.Mysql().Debug().Create(&model.Userfriend{
				UserId:       int(fid),
				FriendId:     int(useid),
				CreationTime: int(time.Now().Unix()),
			})
			c.String(200, "添加好友成功")
		} else {
			c.String(200, "好友不存在")
		}
	} else {
		c.String(200, "你们已经是好友了")
	}
}

//查询好友列表
func GetFriends(c *gin.Context) {
	userid := c.PostForm("userid")
	useid, _ := strconv.ParseInt(userid, 10, 64)
	//根据用户id查出所有好友对象
	var friends []model.Userfriend
	db.Mysql().Debug().Where("user_id = ?", int(useid)).Find(&friends)
	//将所有好友ID添加到切片中
	ids := make([]int, 0, len(friends))
	for _, v := range friends {
		ids = append(ids, v.FriendId)
	}
	//根据好友ID找出好友信息
	var usernames []model.User
	db.Mysql().Debug().Where("id in (?)", ids).Find(&usernames)

	//遍历出好友邮箱
	c.String(200, "好友Email\n")
	for _, v := range usernames {
		c.String(200, v.Email+"\n")
	}
}

//删除好友
func DelFriend(c *gin.Context) {
	userid := c.PostForm("userid")
	friendid := c.PostForm("friendid")
	useid, _ := strconv.ParseInt(userid, 10, 64)
	fid, _ := strconv.ParseInt(friendid, 10, 64)

	var userf model.Userfriend

	//互删好友
	db.Mysql().Debug().Where("user_id = ?", useid).Where("friend_id = ?", fid).Delete(&userf)
	db.Mysql().Debug().Where("user_id = ?", fid).Where("friend_id = ?", useid).Delete(&userf)
	c.String(200, "删除成功")
}

//根据ID或者用户名查找用户
func GetUser(c *gin.Context) {
	keyword := c.PostForm("keyword")
	var user []model.User
	db.Mysql().Debug().Table("users").Where("id = ? or email like ?", keyword, "%"+keyword+"%").Find(&user)
	c.String(200, "用户ID \t 用户Email\n")
	for _, v := range user {
		userid := strconv.Itoa(v.Id)
		c.String(200, userid+"\t\t"+v.Email+"\n")
	}
}

//添加黑名单
func AddBlacklist(c *gin.Context) {
	userid := c.PostForm("userid")
	friendid := c.PostForm("friendid")

	//只能在好友列表里面将好友添加进黑名单，这里就不需要再判断好友是否存在
	var friend model.Userfriend
	db.Mysql().Debug().Where("user_id = ? and friend_id = ?", userid, friendid).Find(&friend)
	if friend.BlacklistState == 0 {
		db.Mysql().Debug().Table("userfriends").Where("user_id = ? and friend_id = ?", userid, friendid).Update("blacklist_state", 1)
		c.String(200, "添加黑名单成功")
	} else {
		c.String(200, "对方已经在你的黑名单中")
	}

}

//查看黑名单
func GetBlacklist(c *gin.Context) {
	userid := c.PostForm("userid")

	var friend []model.Userfriend
	db.Mysql().Debug().Where("user_id = ? and blacklist_state = '1'", userid).Find(&friend)
	//查找出用户所有的黑名单好友ID
	blacklistid := make([]int, 0, len(friend))
	for _, v := range friend {
		blacklistid = append(blacklistid, v.FriendId)
	}
	//查找出来的黑名单ID不等于0，再去查找ID对应的邮箱号
	if len(blacklistid) != 0 {
		var users []model.User
		db.Mysql().Debug().Where("id in (?)", blacklistid).Find(&users)
		c.String(200, "黑名单Email\n")
		for _, v := range users {
			c.String(200, v.Email+"\n")
		}
	} else {
		c.String(200, "你没有好友在黑名单中")
	}

}

//移除黑名单
func DelBlacklist(c *gin.Context) {
	userid := c.PostForm("userid")
	friendid := c.PostForm("friendid")

	//只能在好友列表里面将好友添加进黑名单，这里就不需要再判断好友是否存在
	var friend model.Userfriend
	db.Mysql().Debug().Where("user_id = ? and friend_id = ?", userid, friendid).Find(&friend)
	if friend.BlacklistState == 1 {
		db.Mysql().Debug().Table("userfriends").Where("user_id = ? and friend_id = ?", userid, friendid).Update("blacklist_state", 0)
		c.String(200, "移除黑名单成功")
	} else {
		c.String(200, "对方不在你的黑名单中")
	}
}

func main() {
	e := gin.Default()
	e.LoadHTMLGlob("cshtml/*") //获取当前路径下的cshtml文件夹下的所有html文件
	e.GET("/login", Login)
	e.POST("/login", DoLogin) //POST请求交给DoLogin处理
	e.GET("/regist", Regist)
	e.POST("/regist", DoRegist)           //POST请求交给DoRegist处理
	e.POST("/addfriend", AddFriend)       //添加好友
	e.POST("/getfriends", GetFriends)     //获取好友列表
	e.POST("/delfriend", DelFriend)       //删除好友
	e.POST("/getalluser", GetAllUser)     //查询所有用户
	e.POST("/getuser", GetUser)           //查找好友
	e.POST("/addblacklist", AddBlacklist) //添加黑名单
	e.POST("/getblacklist", GetBlacklist) //查看黑名单
	e.POST("/delblacklist", DelBlacklist) //移除黑名单

	e.Run()
}
