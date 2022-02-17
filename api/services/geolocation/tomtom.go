package tomtom

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type ApiResponse struct {
	Results []struct {
		Address struct {
			StreetName         string
			ExtendedPostalCode string
		}
		Position struct {
			Lat float64
			Lon float64
		}
	}
}

func FetchAddress(params string) (ApiResponse, error) {
	var ApiKey = os.Getenv("TOMTOM_API_KEY")

	response, err := http.Get(
		fmt.Sprintf(
			"https://api.tomtom.com/search/2/geocode/%s.json?storeResult=false&countrySet=BR&view=Unified&key=%s",
			params,
			ApiKey,
		),
	)
	if err != nil {
		return ApiResponse{}, err
	}
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return ApiResponse{}, err
	}

	var apiResponse ApiResponse
	if err := json.Unmarshal(responseData, &apiResponse); err != nil {
		fmt.Println("Can not unmarshal JSON")
	}

	return apiResponse, nil
}
