package model

type Userfriend struct {
	Id             int `gorm:"column:id" json:"id"`                           //id
	UserId         int `gorm:"column:user_id" json:"user_id"`                 //用户ID
	FriendId       int `gorm:"column:friend_id" json:"friend_id"`             //好友ID
	CreationTime   int `gorm:"column:creation_time" json:"creation_time"`     //创建时间
	BlacklistState int `gorm:"column:blacklist_state" json:"blacklist_state"` //黑名单状态
}
