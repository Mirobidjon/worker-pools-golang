package main

import (
	"fmt"
	"log"
	"time"

	"github.com/Mirobidjon/excel_export/excel"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func main() {
	conn := `host=localhost port=5490 user=postgres password=qwerty dbname=excel sslmode=disable `

	db, err := sqlx.Open("postgres", conn)

	if err != nil {
		log.Fatal("error while connecting database")
		return
	}
	start := time.Now()
	// excel.Importdb(db)
	excel.ExportExcel(db)
	duration := time.Since(start)
	fmt.Println("Ish bajarildi! Ketgan vaqt:", duration.Milliseconds())
}
