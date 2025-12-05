package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type InputPrompt struct {
	Role                string `json:"role"`
	Task                string `json:"task"`
	Output_Requirements string `json:"output_requirements"`
	Content_Input       string `json:"content_input"`
}

type QuestionOption struct {
	Text      string `json:"text"`
	IsCorrect bool   `json:"is_correct"`
}

type OutputQuestion struct {
	Topic        string           `json:"topic"`
	GradeLevel   int              `json:"grade_level"`
	QuestionText string           `json:"question_text"`
	Options      []QuestionOption `json:"options"`
	Rationale    string           `json:"rationale"`
}

func fetchQuestion(prompt InputPrompt, apiURL string, apikey string) (OutputQuestion, error) {
	requestBody, err := json.Marshal(prompt)
	if err != nil {
		return OutputQuestion{}, err
	}
	client := &http.Client{Timeout: 30 * time.Second}

	req, err := http.NewRequest("POST", apiURL, bytes.NewBuffer(requestBody))
	if err != nil {
		return OutputQuestion{}, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+apikey)

	resp, err := client.Do(req)
	if err != nil {
		return OutputQuestion{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return OutputQuestion{}, fmt.Errorf("API request failed with status: %s", resp.Status)
	}

	var output OutputQuestion
	err = json.NewDecoder(resp.Body).Decode(&output)
	if err != nil {
		return OutputQuestion{}, err
	}

	return output, nil

}

func main() {
	// Dữ liệu mẫu (sử dụng cấu trúc Prompt JSON bạn đã cung cấp)
	promptData := InputPrompt{
		Role:                "Senior High School Informatics Pedagogy Expert...",
		Task:                "Generate one (1) High-Quality Multiple Choice Question...",
		Output_Requirements: "The JSON must include 'topic', 'grade_level'...",
		Content_Input:       "Nguyên tắc cơ bản của thuật toán sắp xếp nổi bọt (Bubble Sort) là lặp lại việc so sánh cặp phần tử liền kề và hoán đổi chúng nếu chúng không theo đúng thứ tự...",
	}

	// Cấu hình API (Thay thế bằng URL và Key thực tế của nhà cung cấp AI, ví dụ: Gemini API)
	// Để đơn giản, ta dùng giá trị giả định.
	apiURL := "https://generativelanguage.googleapis.com/v1beta/models/gemini-2.5-flash"
	apiKey := "AIzaSyBZTUj8uu6phUdqPxrO93UsZLwW-YZHyW4" // Lấy từ biến môi trường để bảo mật

	if apiKey == "" {
		fmt.Println("Vui lòng đặt biến môi trường AI_API_KEY để chạy mã này.")
		// Vì không thể gọi API thực, ta sẽ kết thúc ở đây.
		return
	}

	fmt.Printf("Đang gửi yêu cầu tạo câu hỏi cho nội dung: %s\n", promptData.Content_Input)

	// Gọi hàm
	mcq, err := fetchQuestion(promptData, apiURL, apiKey)
	if err != nil {
		fmt.Printf("Lỗi gọi API: %v\n", err)
		return
	}

	// In ra kết quả JSON đã nhận
	outputJSON, _ := json.MarshalIndent(mcq, "", "  ")
	fmt.Println("\n--- Câu hỏi Trắc nghiệm đã nhận từ AI ---")
	fmt.Println(string(outputJSON))
}
