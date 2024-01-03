package main

import (
	"fmt"
	"log"
	"mysql-to-excel/initialize"
	"mysql-to-excel/model"
	"mysql-to-excel/utils"
	"strconv"
	"strings"

	"github.com/xuri/excelize/v2"
)

func main() {
	// 初始化GORM
	db := initialize.InitGorm()

	// 获取所有表名
	var tableNames []string
	db.Raw("SHOW TABLES").Scan(&tableNames)
	fmt.Println(tableNames)

	// 新建一个Excel文件
	file := excelize.NewFile()

	for _, tableName := range tableNames {
		var tableColumns []model.TableInfo
		db.Raw("SHOW FULL COLUMNS FROM " + tableName).Scan(&tableColumns)
		// fmt.Println(tableColumns)

		// 新建工作表，命名为表名
		file.NewSheet(tableName)
		// 设置表头
		headers := []string{"列名", "字段类型", "长度", "是否可为空", "默认值", "备注"}
		for colIndex, header := range headers {
			cellName := utils.ToAlphaString(colIndex+1) + "1"
			file.SetCellValue(tableName, cellName, header)
			// fmt.Println(header, " --写入单元格--> ", ToAlphaString(colIndex+1)+"1")
		}

		// 设置数据
		for rowIndex, column := range tableColumns {
			// column.Field, column.Type, Len, column.Null, column.Default
			file.SetCellValue(tableName, "A"+fmt.Sprint(rowIndex+2), column.Field)

			// 分割varchar(256)，取出数字
			if strings.Contains(column.Type, "varchar") || strings.Contains(column.Type, "char") {
				lenStr := strings.Split(column.Type, "(")[1]
				len, _ := strconv.Atoi(strings.Split(lenStr, ")")[0])

				file.SetCellValue(tableName, "B"+fmt.Sprint(rowIndex+2), strings.Split(column.Type, "(")[0])
				file.SetCellValue(tableName, "C"+fmt.Sprint(rowIndex+2), len)
			} else {
				file.SetCellValue(tableName, "B"+fmt.Sprint(rowIndex+2), column.Type)
				file.SetCellValue(tableName, "C"+fmt.Sprint(rowIndex+2), "(NULL)")
			}

			file.SetCellValue(tableName, "D"+fmt.Sprint(rowIndex+2), column.Null)

			if column.Default == "" {
				file.SetCellValue(tableName, "E"+fmt.Sprint(rowIndex+2), "(NULL)")
			} else {
				file.SetCellValue(tableName, "E"+fmt.Sprint(rowIndex+2), column.Default)
			}

			file.SetCellValue(tableName, "F"+fmt.Sprint(rowIndex+2), column.Comment)
		}
	}

	// Save the Excel file
	fileName := "数据库表格.xlsx"
	if err := file.SaveAs(fileName); err != nil {
		log.Fatalf("error saving Excel file: %s", err)
	}

	fmt.Printf("Excel file '%s' saved successfully.\n", fileName)
}
