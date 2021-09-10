package wkapi

import (
	"encoding/json"
	"github.com/ostrea/wanikanileechsuspender/pkg/jsonstructs"
	"log"
	"strconv"
	"strings"
)

func GetSpecifiedSubjects(subjectIds []int) []jsonstructs.SubjectObject {
	var allSpecifiedSubjects []jsonstructs.SubjectObject
	stringIds := make([]string, len(subjectIds))
	for i, subjectId := range subjectIds {
		stringIds[i] = strconv.Itoa(subjectId)
	}
	// TODO It fails if url is too long. Probably won't matter in reality.
	subjectStatsJson, nextUrl := getPage(baseApiUrl + "/subjects?ids=" + strings.Join(stringIds, ","))
	allSpecifiedSubjects = append(allSpecifiedSubjects, extractSubjectFromPage(subjectStatsJson)...)

	for nextUrl != nil {
		subjectStatsJson, nextUrl = getPage(*nextUrl)
		allSpecifiedSubjects = append(allSpecifiedSubjects, extractSubjectFromPage(subjectStatsJson)...)
	}

	return allSpecifiedSubjects
}

func extractSubjectFromPage(subjectStatsJson []byte) []jsonstructs.SubjectObject {
	var parsedJson jsonstructs.SubjectsResponse
	err := json.Unmarshal(subjectStatsJson, &parsedJson)
	if err != nil {
		log.Fatal(err)
	}

	return parsedJson.Data
}
