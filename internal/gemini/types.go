package gemini

import (
	S "farmaIA/internal/survey"
)

type GeminiService interface {
	Gemini()
	GetSurveys() ([]S.Survey, error)
}

type History struct {
	OldText string
	Role    string
}

type GeminiRequest struct {
	Text    string
	History []History
}

type GeminiResponse struct {
	Parts []string
	Role  string
}
