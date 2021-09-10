package jsonstructs

type Review struct {
	SubjectId               int `json:"subject_id"`
	IncorrectMeaningAnswers int `json:"incorrect_meaning_answers"`
	IncorrectReadingAnswers int `json:"incorrect_reading_answers"`
}

type CreateReviewJson struct {
	Review Review `json:"review"`
}
