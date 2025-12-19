package gosql

import (
	_ "github.com/mattn/go-sqlite3"
)

type Essay struct {
	ID              int
	Topic           string
	Content         string
	DifficultyLevel string
	CreatedAt       string
}

func (data *Data) AddEssay(e Essay) int {
	res, err := data.DB.Exec("INSERT INTO essays (topic, content, difficulty_level) VALUES (?, ?, ?)",
		e.Topic, e.Content, e.DifficultyLevel)
	if err != nil {
		return 0
	}
	id, err := res.LastInsertId()
	if err == nil {
		return int(id)
	}
	return 0
}
func (data *Data) GetAllEssays() ([]Essay, error) {
	rows, err := data.DB.Query("SELECT id, topic, content, difficulty_level, created_at FROM essays")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var essays []Essay
	for rows.Next() {
		var e Essay
		err := rows.Scan(&e.ID, &e.Topic, &e.Content, &e.DifficultyLevel, &e.CreatedAt)
		if err != nil {
			return nil, err
		}
		essays = append(essays, e)
	}
	return essays, nil
}

func (data *Data) GetEssayByID(id int) (*Essay, error) {
	row := data.DB.QueryRow("SELECT id, topic, content, difficulty_level, created_at FROM essays WHERE id = ?", id)
	var e Essay
	err := row.Scan(&e.ID, &e.Topic, &e.Content, &e.DifficultyLevel, &e.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &e, nil
}
func (data *Data) UpdateEssay(e Essay) error {
	_, err := data.DB.Exec("UPDATE essays SET topic = ?, content = ?, difficulty_level = ? WHERE id = ?",
		e.Topic, e.Content, e.DifficultyLevel, e.ID)
	return err
}
func (data *Data) DeleteEssay(id int) error {
	_, err := data.DB.Exec("DELETE FROM essays WHERE id = ?", id)
	return err
}
