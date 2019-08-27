package main

import (
	// "github.com/lxn/walk"
	// . "github.com/lxn/walk/declarative"
	// "strings"
	"github.com/xuri/excelize"
	"fmt"
	"os"
)

func openfile() {
    xlsx, err := excelize.OpenFile("./data.xlsx")
    if err != nil {
        fmt.Println("Excel file openning error:{0}", err)
        os.Exit(1)
	}
	
	// Get all the rows in a sheet.
	rows, _ := xlsx.GetRows("Sheet1")
    for i, row := range rows {
		if i == 0{
			continue
		}
		if len(row) == 0{
			break
		}
		fmt.Println(row[0])
    }
}

func main() {
	// var inTE, outTE *walk.TextEdit

	openfile()

	// MainWindow{
	// 	Title:   "SCREAMO",
	// 	MinSize: Size{600, 400},
	// 	Layout:  VBox{},
	// 	Children: []Widget{
	// 		HSplitter{
	// 			Children: []Widget{
	// 				TextEdit{AssignTo: &inTE},
	// 				TextEdit{AssignTo: &outTE, ReadOnly: true},
	// 			},
	// 		},
	// 		PushButton{
	// 			Text: "SCREAM",
	// 			OnClicked: func() {
	// 				outTE.SetText(strings.ToUpper(inTE.Text()))
	// 			},
	// 		},
	// 	},
	// }.Run()
}
