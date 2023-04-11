package main

import (
	"fmt"

	"github.com/xuri/excelize/v2"
)

func main() {
	f, err := excelize.OpenFile("D:\\cs\\Book1.xlsx")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()
	//获取表格中指定单元格的值
	// cell, err := f.GetCellValue("Sheet1", "B2")
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// fmt.Println(cell)
	//获取sheet1表格上的所有单元格
	rows, err := f.GetRows("Sheet1")
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, row := range rows {
		for _, colCell := range row {
			fmt.Println(colCell, "\t")
		}
		fmt.Println()
	}

}

// package main

// import (
// 	"fmt"

// 	"github.com/xuri/excelize/v2"
// )

// func main() {
// 	f := excelize.NewFile()
// 	// 创建一个工作表
// 	index, err := f.NewSheet("Sheet2")
// 	if err != nil {
// 		return
// 	}
// 	// 设置单元格的值
// 	f.SetCellValue("Sheet2", "A2", "Hello world.")
// 	f.SetCellValue("Sheet1", "B2", 100)
// 	// 设置工作簿的默认工作表
// 	f.SetActiveSheet(index)
// 	// 根据指定路径保存文件
// 	if err := f.SaveAs("D:\\cs\\Book1.xlsx"); err != nil {
// 		fmt.Println(err)
// 	}
// }
