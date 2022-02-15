package services

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

var ApiKey = "uzVaR3JVeMZD9d5Qmxj0OFNqVrT6o0ii"

func EnrichmentService(postal_code string) (string, error) {
	response, err := http.Get(fmt.Sprintf("https://api.tomtom.com/search/2/geocode/%s.json?storeResult=false&countrySet=BR&view=Unified&key=%s", postal_code, ApiKey))
	if err != nil {
		return "", err
	}
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}
	return string(responseData), nil
}
