package main

import (
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
	"strings"
	"github.com/xuri/excelize"
	"fmt"
	"os"
	"strconv"
)

func openfile() {
    xlsx, err := excelize.OpenFile("./data.xls")
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
    // Get value from cell by given sheet index and axis.
    cell, _ := xlsx.GetCellValue("Sheet1", "A2")
    fmt.Println(cell)
    // Get sheet index.
    index := xlsx.GetSheetIndex("Sheet1")
    // Get all the rows in a sheet.
    rows, _ := xlsx.GetRows("sheet" + strconv.Itoa(index))
    for _, row := range rows {
        for _, colCell := range row {
            fmt.Print(colCell, "\t")
        }
        fmt.Println()
    }
}

func main() {
	var inTE, outTE *walk.TextEdit

	openfile()

	MainWindow{
		Title:   "SCREAMO",
		MinSize: Size{600, 400},
		Layout:  VBox{},
		Children: []Widget{
			HSplitter{
				Children: []Widget{
					TextEdit{AssignTo: &inTE},
					TextEdit{AssignTo: &outTE, ReadOnly: true},
				},
			},
			PushButton{
				Text: "SCREAM",
				OnClicked: func() {
					outTE.SetText(strings.ToUpper(inTE.Text()))
				},
			},
		},
	}.Run()
}
