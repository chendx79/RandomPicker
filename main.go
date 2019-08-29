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
	mw := &MyMainWindow{fromModel: NewFromModel(), randModel: NewRandModel()}

	if _, err := (MainWindow{
		AssignTo: &mw.MainWindow,
		Title:    "随机选择工具",
		MinSize:  Size{240, 320},
		Size:     Size{300, 400},
		Layout:   VBox{MarginsZero: true},
		Children: []Widget{
			PushButton{
				Text: "重置",
				OnClicked: func() {

				},
			},
			HSplitter{
				Children: []Widget{
					ListBox{
						AssignTo: &mw.lbFrom,
						Model:    mw.fromModel,
					},
					TableView{
						AssignTo:      &mw.lbRand,
						Columns: []TableViewColumn{
							TableViewColumn{
								DataMember: "顺序",
								Width:      50,
							},
							TableViewColumn{
								DataMember: "姓名",
								Width:      64,
							},
						},
						Model: mw.randModel,
					},
				},
			},
			PushButton{
				Text: "随机选取",
				OnClicked: func() {
					mw.randModel.RandPerson()
				},
			},
		},
	}.Run()); err != nil {
		log.Fatal(err)
	}
}

type MyMainWindow struct {
	*walk.MainWindow
	fromModel *FromModel
	randModel *RandModel
	lbFrom    *walk.ListBox
	lbRand    *walk.TableView
}

type FromModel struct {
	walk.ListModelBase
	items []string
}

func NewFromModel() *FromModel {
	var names []string
	names = readNamesFromExcelFile()

	m := &FromModel{items: make([]string, len(names))}

	for i, e := range names {
		m.items[i] = e
	}

	return m
}

func (m *FromModel) ItemCount() int {
	return len(m.items)
}

func (m *FromModel) Value(index int) interface{} {
	return m.items[index]
}

type Person struct {
	Index int
	Name  string
}

type RandModel struct {
	walk.SortedReflectTableModelBase
	items []Person
}

func NewRandModel() *RandModel {
	m := &RandModel{items: make([]Person, 0)}

	return m
}

func (m *RandModel) Items() interface{} {
	return m.items
}

func (m *RandModel) RandPerson() {
	// newItems := append(m.items, "1")
	// m = &RandModel{items: newItems}
}