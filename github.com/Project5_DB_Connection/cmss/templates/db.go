package cmss

import (
	"database/sql"

	_ "github.com/lib/pq" //	//Use the postgres SQL driver
)

type PgStore struct {
	DB *sql.DB
}

func newDB() *PgStore {
	db, err := sql.Open("postgres", "user=goprojects dbname=goprojects sslmode=disable")
	if err != nil {
		panic(err)
	}

	return &PgStore{
		DB: db,
	}
}

func CreatePage(p *Page) (int, error) {
	var id int
	err := store.DB.QueryRow("Insert into pages(title, content) values($1, $2) returning id", p.Title, p.Content).Scan(&id)

	return id, err
}
