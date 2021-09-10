package wkapi

import "os"

var bearer = "Bearer " + os.Getenv("API_TOKEN")
var baseApiUrl = "https://api.wanikani.com/v2"
