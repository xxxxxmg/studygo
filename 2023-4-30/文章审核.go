package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	str := audit("./article.txt", "./thesaurus.txt")
	fmt.Println("您上传的文章有违规内容:", str)
}

func audit(article, thesaurus string) string {
	art := readerArticle(article)
	thes := readerThesaurus(thesaurus)
	str := ""
	for _, v := range thes {
		result := strings.Contains(art, v)
		if result == true {
			str += v + " "
		}
	}
	return str
}

func readerArticle(fileName string) string {
	fp, err := os.OpenFile(fileName, os.O_RDWR, 0)
	if err != nil {
		fmt.Println("Open Err", err)
	}
	defer fp.Close()

	str, err := io.ReadAll(fp)
	if err != nil {
		fmt.Println("ReadAll Err", err)
	}
	return string(str)
}

func readerThesaurus(thesaurus string) []string {
	fp, err := os.OpenFile(thesaurus, os.O_RDWR, 0)
	if err != nil {
		fmt.Println("Open Err", err)
	}
	defer fp.Close()
	reader := bufio.NewReader(fp)
	str, err := reader.ReadString(',')
	if err != nil && err == io.EOF {
		//fmt.Println(err)
	}
	return strings.Split(str, "，")
}
