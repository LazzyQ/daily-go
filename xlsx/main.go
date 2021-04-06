package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/tealeg/xlsx/v3"
)

func main() {
	f, _ := xlsx.OpenFile(os.Args[1])
	sheet := f.Sheets[0]

	sheet.ForEachRow(func(row *xlsx.Row) error {
		var line = make([]string, 0, sheet.MaxCol)
		_ = row.ForEachCell(func(cell *xlsx.Cell) error {
			v, _ := cell.GeneralNumericWithoutScientific()
			line = append(line, strings.TrimSpace(v))
			return nil
		})
		fmt.Printf("line: %v, len: %d\n", line, len(line))
		return nil
		// var line = make([]string, 0, len(row.Cells))
		// for _, cell := range row.Cells {
		// 	v, _ := cell.GeneralNumericWithoutScientific()
		// 	line = append(line, strings.TrimSpace(v))
		// }
		// fmt.Printf("index: %d, %v len(line): %d\n", index, line, len(line))
	})

}
