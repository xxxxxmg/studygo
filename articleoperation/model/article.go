package model

type ArticleInfo struct {
	Id      int    `json:"id"`
	Title   string `json:"title"`   //标题
	Context string `json:"context"` //内容
	Time    int64  `json:"time"`    //时间
	Author  string `json:"author"`  //作者
}
