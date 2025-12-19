package prompt

import (
	"fmt"
	"go-ontn/sqldb"
)
type Exam struct {
    Role       string `json:"Role"`
    Task       Task   `json:"Task"`
    Language   string `json:"Language"`
    Constraints string `json:"Constraints"`
    Output     Output `json:"Output"`
}

type Task struct {
    Type       string    `json:"Type"`
    Subject    string    `json:"Subject"`
    Grade      string    `json:"Grade"`
    ExamPeriod string    `json:"Exam Period"`
    Questions  Questions `json:"Questions"`
}

type Questions struct {
    MultipleChoice []MultipleChoiceQuestion `json:"Multiple Choice"`
    TrueFalse      []TrueFalseQuestion      `json:"True/False"`
}

type MultipleChoiceQuestion struct {
    Group             string `json:"Group"`
    NumberOfQuestions int    `json:"Number of Questions"`
    Difficulty        string `json:"Difficulty"`
    Content           string `json:"content"`
}

type TrueFalseQuestion struct {
    Question   string `json:"Question"`
    Difficulty string `json:"Difficulty"`
    Content    string `json:"content"`
}

type Output struct {
    Type      string `json:"type"`
    Structure string `json:"structure"`
}