package lib

import (
	"fmt"
	"github.com/tealeg/xlsx"
)

//CreateExel is
type CreateExel struct {
}

var (
	file     *xlsx.File
	sheet    *xlsx.Sheet
	row      *xlsx.Row
	cell     *xlsx.Cell
	err      error
	control  bool
	fileName string
)

//Open is
func (c CreateExel) Open() {
	file = xlsx.NewFile()
	control = true
}

//Create is
func (c CreateExel) Create(setFileName string, sheetName string) {
	if setFileName != "" && sheetName != "" {
		sheet, err = file.AddSheet(sheetName)

		if err != nil {
			if !control {
				fmt.Println("Please open")
			}
			fmt.Printf(err.Error())
		}
		fileName = setFileName
	}
}

//CreateNewRow is
func (c CreateExel) CreateNewRow() {
	row = sheet.AddRow()
	cell = row.AddCell()

	//cell = sheet.Cell(5, 3)
	if err != nil {
		fmt.Println("CreateNewRow")
		fmt.Printf(err.Error())
	}
}

//CreateRow is
func (c CreateExel) CreateRow() {
	row = sheet.AddRow()
	//cell = sheet.Cell(0, 0)
	cell = sheet.AddRow().Sheet.Cell(3, 5)
	if err != nil {
		fmt.Println("CreateNewRow")
		fmt.Printf(err.Error())
	}
}

func (c CreateExel) Write(setValue []string) {
	for i := 0; i < len(setValue); i++ {
		cell = row.AddCell()
		cell.Value = setValue[i]
		err = file.Save(fileName + ".xlsx")
		if err != nil {
			fmt.Println("Write")
			fmt.Printf(err.Error())
		}
	}
}

func (c CreateExel) WriteByOne(setValue string) {

	cell = row.AddCell()
	cell.Value = setValue
	err = file.Save(fileName + ".xlsx")
	if err != nil {
		fmt.Println("Write")
		fmt.Printf(err.Error())
	}
}
