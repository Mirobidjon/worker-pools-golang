package excel

import (
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/manveru/faker"
)

var (
	fakeData     *faker.Faker
	ReadDuration time.Duration
)

type Data struct {
	ID       string `db:"id"`
	Name     string `db:"name"`
	Phone    string `db:"phone"`
	ParentID string `db:"parent_id"`
	Key      int64  `db:"key"`
}

func Importdb(db *sqlx.DB) {
	fakeData, _ = faker.New("en")
	for i := 0; i < 20000; i++ {
		id, _ := uuid.NewRandom()
		parent_id, _ := uuid.NewRandom()
		name := fakeData.Name()
		phone := fakeData.PhoneNumber()

		query := `INSERT INTO exceldb (
								id,
								name,
								phone,
								parent_id)
						VALUES ($1, $2, $3, $4) `

		_, err := db.Exec(query, id, name, phone, parent_id)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func GetDB(db *sqlx.DB, page, limit int64) []Data {
	var values []Data
	start := time.Now()
	query := `SELECT
					id,
					name,
					phone,
					parent_id
				FROM exceldb OFFSET %d LIMIT %d`

	err := db.Select(&values, fmt.Sprintf(query, page, limit))
	if err != nil {
		log.Fatal(err)
	}
	duration := time.Since(start)
	ReadDuration += duration
	// fmt.Println("O'qish uchun Ketgan vaqt:", duration.Milliseconds())
	return values
}
