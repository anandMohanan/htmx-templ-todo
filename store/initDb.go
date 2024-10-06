package store

import "database/sql"

func InitDb(db *sql.DB) {
	db.Exec(`CREATE TABLE IF NOT EXISTS todo (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    completed BOOLEAN NOT NULL
)`)
}
