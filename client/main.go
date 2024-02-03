package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

const Uri = "http://0.0.0.0:22022/evaluate/io/dsntk/DecisionContract/"
const SlaUri = Uri + "SLA"
const FineUri = Uri + "Fine"
const ContentType = "application/json"

type SlaParams struct {
	YearsAsCustomer int64 `json:"YearsAsCustomer"`
	NumberOfUnits   int64 `json:"NumberOfUnits"`
}

type SlaResult struct {
	Data int64 `json:"data"`
}

type FineParams struct {
	YearsAsCustomer int64   `json:"YearsAsCustomer"`
	NumberOfUnits   int64   `json:"NumberOfUnits"`
	DefectiveUnits  float64 `json:"DefectiveUnits"`
}

type FineResult struct {
	Data float64 `json:"data"`
}

func querySla(yearsAsCustomer int64, numberOfUnits int64) int64 {
	slaParams := SlaParams{
		YearsAsCustomer: yearsAsCustomer,
		NumberOfUnits:   numberOfUnits,
	}

	var body bytes.Buffer
	err := json.NewEncoder(&body).Encode(&slaParams)
	if err != nil {
		panic(err)
	}

	response, err := http.Post(SlaUri, ContentType, &body)
	if err != nil {
		panic(err)
	}

	slaResult := SlaResult{}
	err = json.NewDecoder(response.Body).Decode(&slaResult)
	if err != nil {
		panic(err)
	}
	return slaResult.Data
}

func queryFine(yearsAsCustomer int64, numberOfUnits int64, defectiveUnits float64) float64 {
	fineParams := FineParams{
		YearsAsCustomer: yearsAsCustomer,
		NumberOfUnits:   numberOfUnits,
		DefectiveUnits:  defectiveUnits,
	}

	var body bytes.Buffer
	err := json.NewEncoder(&body).Encode(&fineParams)
	if err != nil {
		panic(err)
	}

	response, err := http.Post(FineUri, ContentType, &body)
	if err != nil {
		panic(err)
	}

	fineResult := FineResult{}
	err = json.NewDecoder(response.Body).Decode(&fineResult)
	if err != nil {
		panic(err)
	}
	return fineResult.Data
}

func main() {
	fmt.Printf("SLA = %d\n", querySla(1, 1000))
	fmt.Printf("Fine = %.0f%%\n", queryFine(1, 1000, 0.034)*100)
}
