package jsonstructs

import "time"

type BaseCollection struct {
	Pages Pages
}

type Pages struct {
	NextUrl *string `json:"next_url"`
}

type ReviewStatisticResponse struct {
	Data []ReviewStatistic
}

type ReviewStatistic struct {
	Data SubjectStats
}

type SubjectStats struct {
	SubjectId            int    `json:"subject_id"`
	SubjectType          string `json:"subject_type"`
	MeaningIncorrect     int    `json:"meaning_incorrect"`
	MeaningCurrentStreak int    `json:"meaning_current_streak"`
	ReadingIncorrect     int    `json:"reading_incorrect"`
	ReadingCurrentStreak int    `json:"reading_current_streak"`
}

type AssignmentResponse struct {
	Data []AssignmentObject
}

type AssignmentObject struct {
	Data Assignment
}

type Assignment struct {
	SubjectId     int        `json:"subject_id"`
	ResurrectedAt *time.Time `json:"resurrected_at"`
}

type SubjectsResponse struct {
	Data []SubjectObject
}

type SubjectObject struct {
	Id   int
	Type string `json:"object"`
	Data Subject
}

type Subject struct {
	Level       int
	DocumentUrl string `json:"document_url"`
	Characters  string
	Meanings    []Meaning
	Readings    []Reading
}

type Meaning struct {
	Meaning        string
	AcceptedAnswer bool `json:"accepted_answer"`
}

type Reading struct {
	Reading        string
	AcceptedAnswer bool `json:"accepted_answer"`
}
