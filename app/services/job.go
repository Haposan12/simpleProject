package services

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"project/freelance/danspro/simpleProject/app/models"
)

func GetJobList(reqParam models.JobListRequestParam) (interface{}, error) {
	values := url.Values{
		"description": []string{reqParam.Description},
		"location":    []string{reqParam.Location},
		"page":        []string{reqParam.Page},
	}

	if reqParam.FullTime != "" {
		values.Set("full_time", reqParam.FullTime)
	}

	return hitDansMultiProAPI("http://dev3.dansmultipro.co.id/api/recruitment/positions.json?", values)
}

func GetJobDetail(positionsId string) (models.Position, error) {
	var position models.Position
	values := url.Values{}

	url := fmt.Sprintf("http://dev3.dansmultipro.co.id/api/recruitment/positions/%s", positionsId)

	responseDansPro, err := hitDansMultiProAPI(url, values)
	if err != nil {
		return position, err
	}

	dataByte, err := json.Marshal(responseDansPro)
	if err != nil {
		return position, err
	}

	err = json.Unmarshal(dataByte, &position)

	return position, err
}

func hitDansMultiProAPI(urlString string, valuesParam url.Values) (interface{}, error) {
	var positions interface{}

	resp, err := http.Get(urlString + valuesParam.Encode())
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		err = errors.New(http.StatusText(resp.StatusCode))
		return nil, err
	}

	if err = json.NewDecoder(resp.Body).Decode(&positions); err != nil {
		return nil, err
	}

	return positions, nil
}
