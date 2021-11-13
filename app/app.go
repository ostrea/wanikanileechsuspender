package app

import (
	"github.com/ostrea/wanikanileechsuspender/pkg/database"
	"github.com/ostrea/wanikanileechsuspender/pkg/wkapi"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"math"
	"os"
)

func Run() {
	log.Println("Leech suspension started.")
	var leechIds []int

	allSubjectStats := wkapi.GetAllSubjectStats()
	burnedSubjectIds := wkapi.GetBurnedSubjectIds()
	for _, subject := range allSubjectStats {
		if subject.SubjectType == "kanji" {
			continue
		}

		// Some kanji moved beyond current level, and have 0 in streak.
		// They were unlocked ages ago, so they have an entry in stats.
		if subject.MeaningCurrentStreak == 0 {
			continue
		}

		if _, ok := burnedSubjectIds[subject.SubjectId]; ok {
			continue
		}

		leechScoreMeaning := float64(subject.MeaningIncorrect) / (1.5 * float64(subject.MeaningCurrentStreak))
		var leechScoreReading float64
		if subject.SubjectType != "radical" {
			leechScoreReading = float64(subject.ReadingIncorrect) / (1.5 * float64(subject.ReadingCurrentStreak))
		}

		leechScore := math.Max(leechScoreMeaning, leechScoreReading)
		if leechScore >= 2 {
			leechIds = append(leechIds, subject.SubjectId)
		}
	}

	leechSubjects := wkapi.GetSpecifiedSubjects(leechIds)
	db, err := gorm.Open(postgres.Open(os.Getenv("DATABASE_URL")), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	err = db.AutoMigrate(&database.Leech{})
	if err != nil {
		log.Fatal(err)
	}

	database.SaveLeeches(db, leechSubjects)

	currentlyAvailableForReviewIds := wkapi.GetCurrentlyAvailableForReviewIds()
	var leechIdsInReview []int
	db.Model(&database.Leech{}).Where("subject_id IN ?", currentlyAvailableForReviewIds).Pluck("subject_id", &leechIdsInReview)
	wkapi.CreateReviews(leechIdsInReview)

	log.Println("Leech suspension ended.")
}
