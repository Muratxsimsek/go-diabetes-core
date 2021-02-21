package main

type DiabetesChart struct {
	MinSugarValue        int16                `json:"minSugarValue"`
	MaxSugarValue        int16                `json:"maxSugarValue"`
	TotalDiabetesChart   TotalDiabetesChart   `json:"total"`
	FastingDiabetesChart FastingDiabetesChart `json:"fasting"`
	EatingDiabetesChart  EatingDiabetesChart  `json:"eating"`
	OtherDiabetesChart   OtherDiabetesChart   `json:"other"`
}

type TotalDiabetesChart struct {
	SugarValues []int16 `json:"sugarValue"`
	Dates       []int64 `json:"dates"`
}

type FastingDiabetesChart struct {
	SugarValues []int16 `json:"sugarValue"`
}

type EatingDiabetesChart struct {
	SugarValues []int16 `json:"sugarValue"`
}

type OtherDiabetesChart struct {
	SugarValues []int16 `json:"sugarValue"`
}
