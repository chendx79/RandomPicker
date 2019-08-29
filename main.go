package main

import (
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
	"github.com/xuri/excelize"
	"fmt"
	"os"
	"log"
)

func readNamesFromExcelFile() []string{
	var names []string

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
		names = append(names, row[0])
	}
	return names
}

func main() {
	mw := &MyMainWindow{model: NewEnvModel()}

	if _, err := (MainWindow{
		AssignTo: &mw.MainWindow,
		Title:    "随机选择工具",
		MinSize:  Size{240, 320},
		Size:     Size{300, 400},
		Layout:   VBox{MarginsZero: true},
		Children: []Widget{
			HSplitter{
				Children: []Widget{
					ListBox{
						AssignTo: &mw.lb,
						Model:    mw.model,
					},
					TextEdit{
						AssignTo: &mw.te,
						ReadOnly: true,
					},
				},
			},
		},
	}.Run()); err != nil {
		log.Fatal(err)
	}
}

type MyMainWindow struct {
	*walk.MainWindow
	model *EnvModel
	lb    *walk.ListBox
	te    *walk.TextEdit
}

type EnvModel struct {
	walk.ListModelBase
	items []string
}

func NewEnvModel() *EnvModel {
	var names []string
	names = readNamesFromExcelFile()

	m := &EnvModel{items: make([]string, len(names))}

	for i, e := range names {
		m.items[i] = e
	}

	return m
}

func (m *EnvModel) ItemCount() int {
	return len(m.items)
}

func (m *EnvModel) Value(index int) interface{} {
	return m.items[index]
}
