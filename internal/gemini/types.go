package gemini

type GeminiService interface {
	Gemini()
}

type History struct {
    OldText string
    Role string
}

type GeminiRequest struct {
	Text string
    History []History
}

type GeminiResponse struct {
	Parts []string
	Role  string
}
