package model

type Article struct {
	Id       int    `gorm:"column:id;primaryKey;autoIncrement" db:"id" json:"id"`
	Title    string `gorm:"column:title;comment:标题" db:"title" json:"title"`               //标题
	Context  string `gorm:"column:context;comment:内容" db:"context" json:"context"`         //内容
	Author   string `gorm:"column:author;comment:作者" db:"author" json:"author"`            //作者
	UpdateAt int64  `gorm:"column:update_at;comment:创建时间" db:"update_at" json:"update_at"` //修改时间
	CreateAt int64  `gorm:"column:create_at;comment:创建时间" db:"create_at" json:"create_at"` //创建时间
}

// TableName ...
func (*Article) TableName() string {
	return "article"
}

type AddArticle struct {
	Id       int    `gorm:"column:id;primaryKey;autoIncrement" db:"id" json:"id"`
	Title    string `gorm:"column:title;comment:标题" db:"title" json:"title"`               //标题
	Context  string `gorm:"column:context;comment:内容" db:"context" json:"context"`         //内容
	Author   string `gorm:"column:author;comment:作者" db:"author" json:"author"`            //作者
	CreateAt int64  `gorm:"column:create_at;comment:创建时间" db:"create_at" json:"create_at"` //创建时间
}

type DelArticle struct {
	Id int `gorm:"column:id;primaryKey;autoIncrement" db:"id" json:"id"`
}

type UpArticle struct {
	Id       int    `gorm:"column:id;primaryKey;autoIncrement" db:"id" json:"id"`
	Title    string `gorm:"column:title;comment:标题" db:"title" json:"title"`               //标题
	Context  string `gorm:"column:context;comment:内容" db:"context" json:"context"`         //内容
	Author   string `gorm:"column:author;comment:作者" db:"author" json:"author"`            //作者
	UpdateAt int64  `gorm:"column:update_at;comment:创建时间" db:"update_at" json:"update_at"` //修改时间
}

type SelArticle struct {
	Id       int    `gorm:"column:id;primaryKey;autoIncrement" db:"id" json:"id"`
	Title    string `gorm:"column:title;comment:标题" db:"title" json:"title"`               //标题
	Context  string `gorm:"column:context;comment:内容" db:"context" json:"context"`         //内容
	Author   string `gorm:"column:author;comment:作者" db:"author" json:"author"`            //作者
	UpdateAt int64  `gorm:"column:update_at;comment:创建时间" db:"update_at" json:"update_at"` //修改时间
	CreateAt int64  `gorm:"column:create_at;comment:创建时间" db:"create_at" json:"create_at"` //创建时间
	Value    string `gorm:"column:value;comment:查询内容" db:"value" json:"value"`             //查询内容
}
