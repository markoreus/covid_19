package main

import (
	"strconv"
	"strings"
)

func (a Data) Len() int      { return len(a.Countries) }
func (a Data) Swap(i, j int) { a.Countries[i], a.Countries[j] = a.Countries[j], a.Countries[i] }
func (a Data) Less(i, j int) bool {
	return a.Countries[i].TotalConfirmed.Value > a.Countries[j].TotalConfirmed.Value
}

//Statistics generates statistics for every country
func Statistics(c *Country) Statistic {
	pop := int(((float64(c.TotalConfirmed.Value) / c.CasesPerMllP.Value) * 1000000))
	cp := (float64(c.TotalConfirmed.Value) / float64(pop)) * 100
	dp := (float64(c.TotalDeaths.Value) / float64(c.TotalConfirmed.Value)) * 100
	rp := (float64(c.TotalRecovered.Value) / float64(c.TotalConfirmed.Value)) * 100
	ap := (float64(c.ActiveCases.Value) / float64(c.TotalConfirmed.Value)) * 100
	crp := (float64(c.CriticalCases.Value) / float64(c.TotalConfirmed.Value)) * 100
	return Statistic{
		Population:       pop,
		ConfirmedPercent: cp,
		DeathPercent:     dp,
		RecoveredPercent: rp,
		ActivePercent:    ap,
		CriticalPercent:  crp,
	}
}

//UnmarshalJSON for SpecialInt
func (si *SpecialInt) UnmarshalJSON(input []byte) error {
	strInput := string(input)
	strInput = strings.ReplaceAll(strInput, ",", "")
	strInput = strings.ReplaceAll(strInput, "\"", "")
	newInt, err := strconv.Atoi(strInput)
	if err != nil {
		return err
	}
	si.Value = newInt
	return nil
}

//UnmarshalJSON for SpecialFloat
func (si *SpecialFloat) UnmarshalJSON(input []byte) error {
	strInput := string(input)
	strInput = strings.ReplaceAll(strInput, ",", "")
	strInput = strings.ReplaceAll(strInput, "\"", "")
	newFloat, err := strconv.ParseFloat(strInput, 64)
	if err != nil {
		return err
	}
	si.Value = newFloat
	return nil
}
