package main

type DiabetesChart struct {
	MinSugarValue        int16                  `json:"minSugarValue"`
	MaxSugarValue        int16                  `json:"maxSugarValue"`
	TotalDiabetesChart   []TotalDiabetesChart   `json:"total"`
	FastingDiabetesChart []FastingDiabetesChart `json:"fasting"`
	EatingDiabetesChart  []EatingDiabetesChart  `json:"eating"`
	OtherDiabetesChart   []OtherDiabetesChart   `json:"other"`
}

type TotalDiabetesChart struct {
	Dates       int64 `json:"x"`
	SugarValues int16 `json:"y"`
}

type FastingDiabetesChart struct {
	Dates       int64 `json:"x"`
	SugarValues int16 `json:"y"`
}

type EatingDiabetesChart struct {
	Dates       int64 `json:"x"`
	SugarValues int16 `json:"y"`
}

type OtherDiabetesChart struct {
	Dates       int64 `json:"x"`
	SugarValues int16 `json:"y"`
}
