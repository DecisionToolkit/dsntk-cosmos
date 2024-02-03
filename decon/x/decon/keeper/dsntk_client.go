package keeper

import (
	"bytes"
	"encoding/json"
	"net/http"
)

const Uri = "http://0.0.0.0:22022/evaluate/io/dsntk/DecisionContract/"
const SlaUri = Uri + "SLA"
const FineUri = Uri + "Fine"
const ContentType = "application/json"
const Multiplier = 100000000.0

type SlaParams struct {
	YearsAsCustomer uint64 `json:"YearsAsCustomer"`
	NumberOfUnits   uint64 `json:"NumberOfUnits"`
}

type SlaResult struct {
	Data uint64 `json:"data"`
}

type FineParams struct {
	YearsAsCustomer uint64  `json:"YearsAsCustomer"`
	NumberOfUnits   uint64  `json:"NumberOfUnits"`
	DefectiveUnits  float64 `json:"DefectiveUnits"`
}

type FineResult struct {
	Data float64 `json:"data"`
}

func querySla(yearsAsCustomer uint64, numberOfUnits uint64) uint64 {
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

func queryFine(yearsAsCustomer uint64, numberOfUnits uint64, defectiveUnits uint64) uint64 {
	fineParams := FineParams{
		YearsAsCustomer: yearsAsCustomer,
		NumberOfUnits:   numberOfUnits,
		DefectiveUnits:  float64(defectiveUnits) / Multiplier,
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
	return uint64(fineResult.Data * Multiplier)
}
