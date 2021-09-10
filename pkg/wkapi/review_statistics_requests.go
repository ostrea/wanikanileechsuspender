package wkapi

import (
	"encoding/json"
	"github.com/ostrea/wanikanileechsuspender/pkg/jsonstructs"
	"log"
)

func GetAllSubjectStats() []jsonstructs.SubjectStats {
	var allSubjectStats []jsonstructs.SubjectStats
	subjectStatsJson, nextUrl := getPage(baseApiUrl + "/review_statistics")
	allSubjectStats = append(allSubjectStats, extractSubjectStatsFromPage(subjectStatsJson)...)

	for nextUrl != nil {
		subjectStatsJson, nextUrl = getPage(*nextUrl)
		allSubjectStats = append(allSubjectStats, extractSubjectStatsFromPage(subjectStatsJson)...)
	}

	return allSubjectStats
}

func extractSubjectStatsFromPage(subjectStatsJson []byte) []jsonstructs.SubjectStats {
	var parsedJson jsonstructs.ReviewStatisticResponse
	err := json.Unmarshal(subjectStatsJson, &parsedJson)
	if err != nil {
		log.Fatal(err)
	}

	subjectStats := make([]jsonstructs.SubjectStats, len(parsedJson.Data))
	for i, reviewStatistic := range parsedJson.Data {
		subjectStats[i] = reviewStatistic.Data
	}
	return subjectStats
}
