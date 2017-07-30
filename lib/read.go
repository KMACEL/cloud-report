package lib

import (
	"fmt"
	"github.com/tealeg/xlsx"
)

//ReadExel is
type ReadExel struct {
}

var (
	excelFileName string
	xlsFile       *xlsx.File
	errs          error
)

//SetReadFilePath is
func (r ReadExel) SetReadFilePath(path string) {
	xlsFile, errs = xlsx.OpenFile(path)
	if err != nil {
		fmt.Println("Hata")
	}
}

//ReadAll is all read
func (r ReadExel) ReadAll() {

	for _, sheet := range xlsFile.Sheets {
		for _, row := range sheet.Rows {
			for _, cell := range row.Cells {
				text := cell.String()
				fmt.Printf("%s ", text)
			}
			fmt.Printf("\n\n\n")
		}
	}
}

//ReadColumnAll is
func (r ReadExel) ReadColumnAll(readColumnNumber int) []string {
	returnValue := make([]string, 5)
	for _, sheet := range xlsFile.Sheets {
		for _, row := range sheet.Rows {
			if row.Cells[1].String() == "" {
				fmt.Println("NOT User")
				fmt.Printf("%s\n", row.Cells[readColumnNumber])
				fmt.Println("--------------------------")
				returnValue = append(returnValue, row.Cells[0].String())
				fmt.Printf("\n\n\n")
			}
		}
	}
	return returnValue
}

//ReadColumnByOne is
func (r ReadExel) ReadColumnByOne(readColumnNumber1 int, deviceID func(string), columnNumber1 func(string)) {
	for _, sheet := range xlsFile.Sheets {
		for _, row := range sheet.Rows {
			deviceID(row.Cells[0].String())
			columnNumber1(row.Cells[readColumnNumber1].String())
		}
	}
}

//ReadColumnByTwo is
func (r ReadExel) ReadColumnByTwo(readColumnNumber1 int, readColumnNumber2 int, deviceID func(string), columnNumber1 func(string), columnNumber2 func(string)) {
	for _, sheet := range xlsFile.Sheets {
		for _, row := range sheet.Rows {
			deviceID(row.Cells[0].String())
			columnNumber1(row.Cells[readColumnNumber1].String())
			columnNumber2(row.Cells[readColumnNumber2].String())
		}
	}
}
