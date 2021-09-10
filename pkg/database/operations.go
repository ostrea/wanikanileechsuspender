package database

import (
	"github.com/ostrea/wanikanileechsuspender/pkg/jsonstructs"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func SaveLeeches(db *gorm.DB, leechSubjects []jsonstructs.SubjectObject) {
	dbObjects := make([]Leech, len(leechSubjects))
	for i, leechSubject := range leechSubjects {
		var acceptedMeaning string
		for _, meaning := range leechSubject.Data.Meanings {
			if meaning.AcceptedAnswer {
				acceptedMeaning = meaning.Meaning
				break
			}
		}
		var acceptedReading *string
		for _, reading := range leechSubject.Data.Readings {
			if reading.AcceptedAnswer {
				acceptedReading = &reading.Reading
				break
			}
		}
		dbObject := Leech{
			SubjectId: leechSubject.Id,
			Type:      leechSubject.Type,
			Level:     leechSubject.Data.Level,
			Url:       leechSubject.Data.DocumentUrl,
			Value:     leechSubject.Data.Characters,
			Meaning:   acceptedMeaning,
			Reading:   acceptedReading,
		}
		dbObjects[i] = dbObject
	}

	db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "subject_id"}},
		UpdateAll: true,
	}).Create(&dbObjects)
}
