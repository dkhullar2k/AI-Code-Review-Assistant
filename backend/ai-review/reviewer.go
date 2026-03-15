package aireview

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type GroqRequest struct {
	Model    string    `json:"model"`
	Messages []Message `json:"messages"`
}

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type AIReview struct {
	Score          int    `json:"score"`
	ReviewMarkdown string `json:"review_markdown"`
}

func ReviewCode(diff string) AIReview {

	apiKey := os.Getenv("GROQ_API_KEY")

	prompt := fmt.Sprintf(`
	You are a senior software engineer performing a code review.

	Analyze the code diff and produce:

	1. Bugs
	2. Performance issues
	3. Security concerns
	4. Code quality suggestions

	Return your response ONLY in JSON with the following structure:

	{
	"score": number,
	"review_markdown": "well formatted markdown review"
	}

	The review_markdown should contain clear sections like:

	### Possible Bugs
	### Performance Issues
	### Security Concerns
	### Code Quality Suggestions
	### Score

	Respond ONLY with raw JSON.
	Do NOT wrap the JSON in markdown or code blocks.
	Code diff:
	%s
	`, diff)

	reqBody := GroqRequest{
		Model: "llama-3.3-70b-versatile",
		Messages: []Message{
			{
				Role:    "user",
				Content: prompt,
			},
		},
	}

	bodyBytes, _ := json.Marshal(reqBody)

	req, _ := http.NewRequest(
		"POST",
		"https://api.groq.com/openai/v1/chat/completions",
		bytes.NewBuffer(bodyBytes),
	)

	req.Header.Set("Authorization", "Bearer "+apiKey)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}

	resp, err := client.Do(req)

	if err != nil {
		fmt.Println("AI request failed:", err)
		return AIReview{}
	}

	defer resp.Body.Close()

	var result map[string]interface{}

	json.NewDecoder(resp.Body).Decode(&result)

	choices := result["choices"].([]interface{})

	firstChoice := choices[0].(map[string]interface{})

	message := firstChoice["message"].(map[string]interface{})

	content := message["content"].(string)

	var review AIReview

	resultError := json.Unmarshal([]byte(content), &review)

	if resultError != nil {
		fmt.Println("Failed to parse AI JSON:", resultError)
	}

	return review
}
