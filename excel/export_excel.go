package excel

import (
	"fmt"

	"github.com/360EntSecGroup-Skylar/excelize/v2"
	"github.com/jmoiron/sqlx"
)

func ExportExcel(db *sqlx.DB) {
	f := excelize.NewFile()
	var i, key int64
	page := 25
	limit := 2000
	style, _ := f.NewStyle(`{"alignment":{"horizontal":"center", "vertical": "center", "wrap_text": true}}`)
	f.SetCellStyle("Sheet1", "A1", fmt.Sprintf("G%d", page*limit), style)
	f.SetColWidth("Sheet1", "A", "D", 30)
	f.SetRowHeight("Sheet1", 1, 35)

	jobsCount := page * limit
	jobs := make(chan Data, jobsCount)
	results := make(chan bool, jobsCount)
	for k := 1; k <= 2; k++ {
		go Writers(jobs, results, f)
	}

	key = 2
	for i = 1; i <= int64(page); i++ {
		data := GetDB(db, i, int64(limit))
		for _, v := range data {
			v.Key = key - 1
			jobs <- v
			key++
		}
	}
	close(jobs)

	for k := 0; k < jobsCount; k++ {
		result := <-results
		if !result {
			break
		}
	}

	err := f.SaveAs("01.create.xlsx")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("page : %d  limit : %d \n\n", page, limit)
}

func Writers(jobs <-chan Data, results chan<- bool, f *excelize.File) {
	for v := range jobs {
		f.SetRowHeight("Sheet1", int(v.Key), 30)
		f.SetCellValue("Sheet1", fmt.Sprintf("%c%d", 'A', v.Key), v.ID)
		f.SetCellValue("Sheet1", fmt.Sprintf("%c%d", 'B', v.Key), v.Name)
		f.SetCellValue("Sheet1", fmt.Sprintf("%c%d", 'C', v.Key), v.Phone)
		f.SetCellValue("Sheet1", fmt.Sprintf("%c%d", 'D', v.Key), v.ParentID)
		f.SetCellValue("Sheet1", fmt.Sprintf("%c%d", 'E', v.Key), v.Key)
		results <- true
	}
}
