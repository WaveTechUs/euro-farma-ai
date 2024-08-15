package gemini

type GeminiService interface {
    Gemini()
}

type GeminiRequest struct {
    Text string
}
