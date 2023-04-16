package stu

import "fmt"

type Stu struct {
	Name  string
	Age   int
	Score int
}

func (s *Stu) GetStu() string {
	return fmt.Sprintf("姓名：%s,年龄：%d,得分：%d", s.Name, s.Age, s.Score)
}
