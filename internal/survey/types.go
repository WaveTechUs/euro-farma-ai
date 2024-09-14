package survey

type Survey struct {
	Id          int
	CreatedAt   []uint8
	Description string
	Topic       string
	Status      string
	Summary      string
	Conclusions string
	Method      string
	Keywords    string
}

type SurveyService interface {
	GetSurveys() ([]Survey, error)
}
