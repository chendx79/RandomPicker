package main

import (
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
	"github.com/xuri/excelize"
	"fmt"
	"os"
	"log"
	"math/rand"
	"time"
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
			Label{
				MaxSize: Size{300, 20},
				Font:     Font{PointSize: 12},
				Alignment:AlignHCenterVCenter,
				Text:    `请点击随机选取按钮`,
				AssignTo:&mw.labelMsg,
			},
			PushButton{
				Text: "重置",
				Font:     Font{PointSize: 12},
				OnClicked: func() {
					mw.randModel.EmptyPersons()
					mw.fromModel.items = make([]string, len(mw.fromModel.itemsImported))
					copy(mw.fromModel.items, mw.fromModel.itemsImported)
					mw.lbFrom.SetModel(mw.fromModel.items)
					mw.labelMsg.SetText("列表已重置")
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
				Font:     Font{PointSize: 12},
				OnClicked: func() {
					if len(mw.fromModel.items) > 0{
						r := rand.New(rand.NewSource(time.Now().Unix()))
						randIndex := r.Intn(len(mw.fromModel.items))
						
						name := mw.fromModel.items[randIndex]
						mw.randModel.AddPerson(name)
						mw.fromModel.items = append(mw.fromModel.items[:randIndex], mw.fromModel.items[randIndex+1:]...)
						mw.lbFrom.SetModel(mw.fromModel.items)
						msg := fmt.Sprintf("顺序%d, %s", len(mw.randModel.items), name)
						mw.labelMsg.SetText(msg)
					}
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
	labelMsg  *walk.Label
}

type FromModel struct {
	walk.ListModelBase
	itemsImported []string
	items []string
}

func NewFromModel() *FromModel {
	var names []string
	names = readNamesFromExcelFile()

	m := &FromModel{items: make([]string, len(names)), itemsImported: make([]string, len(names))}

	for i, e := range names {
		m.items[i] = e
	}

	copy(m.itemsImported, m.items)

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
	walk.TableModelBase
	items []*Person
}
 
func NewRandModel() *RandModel {
	m := &RandModel{items: make([]*Person, 0)}
	return m
}

func (m *RandModel) EmptyPersons() {
	m.items = make([]*Person, 0)
	m.PublishRowsReset()
}

func (m *RandModel) AddPerson(name string) {
	m.items = append(m.items, &Person{Index:len(m.items) + 1, Name: name})
	m.PublishRowsReset()
}

func (m *RandModel) RowCount() int {
	return len(m.items)
}

// Called by the TableView when it needs the text to display for a given cell.
func (m *RandModel) Value(row, col int) interface{} {
	item := m.items[row]

	switch col {
	case 0:
		return item.Index

	case 1:
		return item.Name
	}

	panic("unexpected col")
}
