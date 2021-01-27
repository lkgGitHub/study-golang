package cnpc

import (
	"fmt"

	"github.com/tealeg/xlsx"
)

var (
	inFile = "/Users/mac/cnpc/数字图书馆/20210119-索引.xlsx"
)

func ReadExcel() {
	// 打开文件
	xlFile, err := xlsx.OpenFile(inFile)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	// 遍历sheet页读取
	for _, sheet := range xlFile.Sheets {
		fmt.Println("sheet name: ", sheet.Name)
		//遍历行读取
		for _, row := range sheet.Rows {
			// 遍历每行的列读取
			for _, cell := range row.Cells {
				text := cell.String()
				fmt.Printf("%20s", text)
			}
			fmt.Print("\n")
		}
	}
	fmt.Println("\n\nimport success")
}
