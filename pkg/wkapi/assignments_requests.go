package wkapi

import (
	"encoding/json"
	"github.com/ostrea/wanikanileechsuspender/pkg/jsonstructs"
	"log"
)

func GetBurnedSubjectIds() map[int]struct{} {
	var allBurnedAssignments []jsonstructs.Assignment
	burnedJson, nextUrl := getPage(baseApiUrl + "/assignments?burned=true")
	allBurnedAssignments = append(allBurnedAssignments, extractAssignmentFromPage(burnedJson)...)

	for nextUrl != nil {
		burnedJson, nextUrl = getPage(*nextUrl)
		allBurnedAssignments = append(allBurnedAssignments, extractAssignmentFromPage(burnedJson)...)
	}

	allBurnedAssignmentIds := make(map[int]struct{})
	for _, assignment := range allBurnedAssignments {
		if assignment.ResurrectedAt != nil {
			continue
		}

		allBurnedAssignmentIds[assignment.SubjectId] = struct{}{}
	}
	return allBurnedAssignmentIds
}

func extractAssignmentFromPage(subjectStatsJson []byte) []jsonstructs.Assignment {
	var parsedJson jsonstructs.AssignmentResponse
	err := json.Unmarshal(subjectStatsJson, &parsedJson)
	if err != nil {
		log.Fatal(err)
	}

	assignments := make([]jsonstructs.Assignment, len(parsedJson.Data))
	for i, assignmentObject := range parsedJson.Data {
		assignments[i] = assignmentObject.Data
	}
	return assignments
}
