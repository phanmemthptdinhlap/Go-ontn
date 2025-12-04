package gosql

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

	CREATE TABLE IF NOT EXISTS questions (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		question_text TEXT NOT NULL,
		option_a TEXT NOT NULL,
		option_b TEXT NOT NULL,
		option_c TEXT NOT NULL,
		option_d TEXT NOT NULL,
		correct_option TEXT NOT NULL,
		difficulty_level TEXT NOT NULL,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP
	);
	CREATE TABLE IF NOT EXISTS true_false_questions (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		context TEXT NOT NULL,
		question_text TEXT NOT NULL,
		statement_1 TEXT NOT NULL,
		statement_2 TEXT NOT NULL,
		statement_3 TEXT NOT NULL,
		statement_4 TEXT NOT NULL,
		correct_statements TEXT NOT NULL,
		difficulty_level TEXT NOT NULL,
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
func InsertConfig(db *sql.DB, api, vaiTro, huongDan, nhiemVu, noiDung string) error {
	_, err := db.Exec(`INSERT INTO config (api, vai_tro, huong_dan, nhiem_vu, noi_dung) VALUES (?, ?, ?, ?, ?)`, api, vaiTro, huongDan, nhiemVu, noiDung)
	return err
}

func DeleteConfig(db *sql.DB, id int) error {
	_, err := db.Exec(`DELETE FROM config WHERE id = ?`, id)
	return err
}

func GetConfigs(db *sql.DB) ([]map[string]interface{}, error) {
	rows, err := db.Query(`SELECT * FROM config`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var configs []map[string]interface{}
	for rows.Next() {
		var id int
		var api, vaiTro, huongDan, nhiemVu, noiDung string
		if err := rows.Scan(&id, &api, &vaiTro, &huongDan, &nhiemVu, &noiDung); err != nil {
			return nil, err
		}
		configs = append(configs, map[string]interface{}{
			"id":        id,
			"api":       api,
			"vai_tro":   vaiTro,
			"huong_dan": huongDan,
			"nhiem_vu":  nhiemVu,
			"noi_dung":  noiDung,
		})
	}
	return configs, nil
}
