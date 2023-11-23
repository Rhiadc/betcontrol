package infra

import (
	"fmt"

	"github.com/xuri/excelize/v2"
)

type Sheet struct {
	Name         string
	DocumentName string
}

type ColumnName struct {
	Column map[string]string
}

func NewSheet(name string, documentName string) *Sheet {
	return &Sheet{Name: name, DocumentName: documentName}
}

func (s *Sheet) CreateSpreadSheet(columnName ColumnName) {
	f := excelize.NewFile()
	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	index, err := f.NewSheet(s.Name)
	if err != nil {
		fmt.Println(err)
		return
	}

	s.nameColumns(s.Name, columnName, f)

	f.SetActiveSheet(index)
	if err := f.SaveAs(s.DocumentName); err != nil {
		fmt.Println(err)
	}
}

func (s *Sheet) nameColumns(sheet string, columnName ColumnName, f *excelize.File) {
	for k, v := range columnName.Column {
		f.SetCellValue(sheet, k, v)
	}
}

func (s *Sheet) addData(f *excelize.File) {

}
