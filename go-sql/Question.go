package gosql

import (
	_ "github.com/mattn/go-sqlite3"
)

type Question struct {
	ID              int
	QuestionType    string
	QuestionText    string
	OptionA         string
	OptionB         string
	OptionC         string
	OptionD         string
	CorrectOption   string
	DifficultyLevel string
	CreatedAt       string
}

func (data *Data) AddQuestion(q Question) int {
	res, err := data.DB.Exec("INSERT INTO questions (question_type, question_text, option_a, option_b, option_c, option_d, correct_option, difficulty_level) VALUES (?, ?, ?, ?, ?, ?, ?, ?)",
		q.QuestionType, q.QuestionText, q.OptionA, q.OptionB, q.OptionC, q.OptionD, q.CorrectOption, q.DifficultyLevel)
	if err != nil {
		return 0
	}
	id, err := res.LastInsertId()
	if err == nil {
		return int(id)
	}
	return 0
}
func (data *Data) GetAllQuestions() ([]Question, error) {
	rows, err := data.DB.Query("SELECT id, question_type, question_text, option_a, option_b, option_c, option_d, correct_option, difficulty_level, created_at FROM questions")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var questions []Question
	for rows.Next() {
		var q Question
		err := rows.Scan(&q.ID, &q.QuestionType, &q.QuestionText, &q.OptionA, &q.OptionB, &q.OptionC, &q.OptionD, &q.CorrectOption, &q.DifficultyLevel, &q.CreatedAt)
		if err != nil {
			return nil, err
		}
		questions = append(questions, q)
	}
	return questions, nil
}

func (data *Data) GetQuestionByID(id int) (*Question, error) {
	row := data.DB.QueryRow("SELECT id, question_type, question_text, option_a, option_b, option_c, option_d, correct_option, difficulty_level, created_at FROM questions WHERE id = ?", id)
	var q Question
	err := row.Scan(&q.ID, &q.QuestionType, &q.QuestionText, &q.OptionA, &q.OptionB, &q.OptionC, &q.OptionD, &q.CorrectOption, &q.DifficultyLevel, &q.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &q, nil
}
func (data *Data) UpdateQuestion(q Question) error {
	_, err := data.DB.Exec("UPDATE questions SET question_type = ?, question_text = ?, option_a = ?, option_b = ?, option_c = ?, option_d = ?, correct_option = ?, difficulty_level = ? WHERE id = ?",
		q.QuestionType, q.QuestionText, q.OptionA, q.OptionB, q.OptionC, q.OptionD, q.CorrectOption, q.DifficultyLevel, q.ID)
	return err
}
func (data *Data) DeleteQuestion(id int) error {
	_, err := data.DB.Exec("DELETE FROM questions WHERE id = ?", id)
	return err
}
