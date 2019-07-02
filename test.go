package main

import (
	"encoding/csv"
	"gin-blog/pkg/export"
	"os"
)

func main() {
	testExcel()
}

func testExcel() {
	file, err := os.Create(export.GetExcelFullPath() + "text.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	//\xEF\xBB\xBF 是 UTF-8 BOM 的 16 进制格式
	file.WriteString("\xEF\xBB\xBF")
	w := csv.NewWriter(file)
	data := [][]string{
		{"1", "test1", "test1-1"},
		{"2", "test2", "test2-1"},
		{"3", "test3", "test3-1"},
	}
	w.WriteAll(data)
}
