package wkapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/ostrea/wanikanileechsuspender/pkg/jsonstructs"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

func CreateReviews(ids []int) {
	for _, id := range ids[:1] {
		newReview := jsonstructs.Review{
			SubjectId:               id,
			IncorrectMeaningAnswers: 0,
			IncorrectReadingAnswers: 0,
		}
		err := sendRequest(newReview)

		if err != nil {
			log.Printf("Creation of a review for %v failed with the following error: %v", id, err)
		}

	}
}

func sendRequest(review jsonstructs.Review) error {
	jsonBytes, err := json.Marshal(jsonstructs.CreateReviewJson{Review: review})
	if err != nil {
		return err
	}

	request, err := http.NewRequest("POST", baseApiUrl+"/reviews", bytes.NewBuffer(jsonBytes))
	if err != nil {
		return err
	}
	request.Header.Set("Authorization", bearer)
	request.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return err
	}
	defer func(body io.ReadCloser) {
		err := body.Close()
		if err != nil {
			log.Printf("Error closing response body: %v", err)
		}
	}(response.Body)

	if response.StatusCode != http.StatusCreated {
		data, _ := ioutil.ReadAll(response.Body)
		return fmt.Errorf("error non 200 response from wk api. Status code: %v, response data: %v", response.StatusCode, string(data))
	}

	return nil
}
