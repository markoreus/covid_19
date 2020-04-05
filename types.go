package main

//Country struct
type Country struct {
	Name           string       `json:"country"`
	TotalConfirmed SpecialInt   `json:"total_cases"`
	NewConfirmed   SpecialInt   `json:"new_cases"`
	TotalDeaths    SpecialInt   `json:"total_deaths"`
	NewDeaths      SpecialInt   `json:"new_deaths"`
	TotalRecovered SpecialInt   `json:"total_recovered"`
	ActiveCases    SpecialInt   `json:"active_cases"`
	CriticalCases  SpecialInt   `json:"serious_critical"`
	CasesPerMllP   SpecialFloat `json:"cases_per_mill_pop"`
}

//Countries contains the data of all the countries
type Countries struct {
	Data []Country `json:"Countries"`
}

//Data that contains the countries info and date
type Data struct {
	Date      string    `json:"last_update"`
	Countries []Country `json:"rows"`
}

//JSON root
type JSON struct {
	Data Data `json:"data"`
}

//Statistic values of each country
type Statistic struct {
	Population       int
	ConfirmedPercent float64
	DeathPercent     float64
	RecoveredPercent float64
	ActivePercent    float64
	CriticalPercent  float64
}

//SpecialInt is a wrapper to unmarshall ints
type SpecialInt struct {
	Value int
}

//SpecialFloat is a wrapper to unmarshall floats
type SpecialFloat struct {
	Value float64
}
