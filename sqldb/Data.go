package sqldb

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func InitDB(dbPath string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, err
	}

	createTableSQL := `
	CREATE TABLE IF NOT EXISTS config (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		api TEXT NOT NULL,
		vai_tro TEXT NOT NULL,
		huong_dan TEXT NOT NULL,
		nhiem_vu TEXT NOT NULL,
		noi_dung TEXT NOT NULL,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP
	);
	CREATE TABLE IF NOT EXISTS essays (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		topic TEXT NOT NULL,
		content TEXT NOT NULL,
		difficulty_level TEXT NOT NULL,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP
	);

	CREATE TABLE IF NOT EXISTS questions (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		question_type TEXT NOT NULL,
		question_text TEXT NOT NULL,
		option_a TEXT NOT NULL,
		option_b TEXT NOT NULL,
		option_c TEXT NOT NULL,
		option_d TEXT NOT NULL,
		correct_option TEXT NOT NULL,
		difficulty_level TEXT NOT NULL,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP
	);
	CREATE TABLE IF NOT EXISTS images (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		image_name TEXT NOT NULL,
		image_data BLOB NOT NULL,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP
	);
	`

	_, err = db.Exec(createTableSQL)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return db, nil
}

type Data struct {
	DB *sql.DB
}

func NewData(dbPath string) (*Data, error) {
	db, err := InitDB(dbPath)
	if err != nil {
		return nil, err
	}
	return &Data{DB: db}, nil
}
func (d *Data) Close() {
	if d.DB != nil {
		d.DB.Close()
	}
}
