package model

type ArticleInfo struct {
	Id           int    `json:"id"`
	Title        string `json:"title"`        //标题
	Context      string `json:"context"`      //内容
	Author       string `json:"author"`       //作者
	CreationTime int64  `json:"creationtime"` //创建时间
	ModifyTime   int64  `json:"modifytime"`   //修改时间
	Value        string `json:"value"`        //查询内容
}
