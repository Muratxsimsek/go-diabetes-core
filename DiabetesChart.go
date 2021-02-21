package main

type DiabetesChart struct {
	MinSugarValue int16   `json:"minSugarValue"`
	MaxSugarValue int16   `json:"maxSugarValue"`
	SugarValues   []int16 `json:"sugarValue"`
	Dates         []int64 `json:"dates"`
}
