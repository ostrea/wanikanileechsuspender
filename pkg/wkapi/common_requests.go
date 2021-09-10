package wkapi

import (
	"encoding/json"
	"github.com/ostrea/wanikanileechsuspender/pkg/jsonstructs"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

func getPage(url string) ([]byte, *string) {
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}
	request.Header.Set("Authorization", bearer)

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		log.Fatal(err)
	}
	defer func(body io.ReadCloser) {
		err := body.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(response.Body)

	data, _ := ioutil.ReadAll(response.Body)

	if response.StatusCode != http.StatusOK {
		log.Fatalf("Error non 200 response from wk api. Status code: %v, response data: %v", response.StatusCode, string(data))
	}

	var parsedJson jsonstructs.BaseCollection
	err = json.Unmarshal(data, &parsedJson)
	if err != nil {
		log.Fatal(err)
	}

	return data, parsedJson.Pages.NextUrl
}
