package utils

import (
	"fmt"
	"strconv"

	"github.com/lzhy87/domain_spider/model"

	"github.com/tealeg/xlsx"
)

// //ReadXls 读取指定的xls文件内容
// func ReadXls(filePath string) ([]model.Ete, error) {
// 	// defer model.Wg.Done()
// 	var orders []model.Ete
// 	xlFile, err := xlsx.OpenFile(filePath)
// 	if err != nil {

// 		fmt.Println("打开excel文件失败", err)
// 		return nil, err
// 	}

// 	for _, sheet := range xlFile.Sheets {
// 		order := model.Ete{}
// 		for _, row := range sheet.Rows {
// 			var str []string
// 			for _, cell := range row.Cells {

// 				text := cell.String()
// 				str = append(str, strings.Trim(text, " "))

// 			}
// 			if str != nil {
// 				order.OrderID = str[0]
// 				order.Name = str[1]
// 				order.Tel = str[2]
// 				order.ID = str[3]
// 				order.Addr = str[4]
// 				// order.Balance = "nil"
// 				orders = append(orders, order)
// 			} else {
// 				//fmt.Println("请检查excel文件里面的内容是否正确！")
// 				return nil, fmt.Errorf("404")
// 			}

// 		}

// 	}
// 	return orders, nil
// }

//WriteXls 写入数据
func WriteXls(orders []*model.Domain, savePath string) {
	// defer model.Wg.Done()
	var file *xlsx.File
	var sheet *xlsx.Sheet
	var row *xlsx.Row

	var cell *xlsx.Cell
	var err error

	file = xlsx.NewFile()
	sheet, err = file.AddSheet("Sheet1")
	if err != nil {
		fmt.Printf(err.Error())
	}

	sheet.SetColWidth(10, 10, 60)
	row = sheet.AddRow()
	row.SetHeightCM(0.5)
	cell = row.AddCell()
	cell.Value = "序号"
	cell = row.AddCell()
	cell.Value = "域名"
	cell = row.AddCell()
	cell.Value = "后缀"
	cell = row.AddCell()
	cell.Value = "网址"
	cell = row.AddCell()
	cell.Value = "已注册"

	for _, i := range orders {
		if i.OrderID == "序号" {
			continue
		}
		var row1 *xlsx.Row
		row1 = sheet.AddRow()
		row1.SetHeightCM(0.5)

		cell = row1.AddCell()
		cell.Value = i.OrderID
		cell = row1.AddCell()
		cell.Value = i.Name
		cell = row1.AddCell()
		cell.Value = i.Suffix
		cell = row1.AddCell()
		cell.Value = i.UserAddr
		cell = row1.AddCell()
		cell.Value = strconv.FormatBool(i.IsRegister)

	}
	err = file.Save(savePath)
	if err != nil {
		fmt.Printf(err.Error())
	}
	fmt.Println("数据写入excel表完成")

}
