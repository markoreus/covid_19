package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

func writeCsv(data Data) error {
	recordFile, err := os.Create("./coronavirus.csv")
	fmt.Println("Writing data to: " + recordFile.Name())
	if err != nil {
		return err
	}
	writer := csv.NewWriter(recordFile)
	headers := []string{
		"Country",
		"Total Confirmed",
		"New Confirmed",
		"Total Deaths",
		"New Deaths",
		"Total Recovered",
		"Active Cases",
		"Critical Cases",
		"Cases Per Million Pop",
		"Date",
		"Total Population",
		"Confirmed Percent",
		"Death Percent",
		"Recovered Percent",
		"Active Percent",
		"Critical Percent",
	}
	writer.Write(headers)
	date := data.Date
	for _, value := range data.Countries {
		stats := Statistics(&value)
		err := writer.Write([]string{
			value.Name,
			strconv.Itoa(value.TotalConfirmed.Value),
			strconv.Itoa(value.NewConfirmed.Value),
			strconv.Itoa(value.TotalDeaths.Value),
			strconv.Itoa(value.NewDeaths.Value),
			strconv.Itoa(value.TotalRecovered.Value),
			strconv.Itoa(value.ActiveCases.Value),
			strconv.Itoa(value.CriticalCases.Value),
			strconv.FormatFloat(value.CasesPerMllP.Value, 'f', 10, 64),
			date,
			strconv.Itoa(stats.Population),
			strconv.FormatFloat(stats.ConfirmedPercent, 'f', 10, 64),
			strconv.FormatFloat(stats.DeathPercent, 'f', 10, 64),
			strconv.FormatFloat(stats.RecoveredPercent, 'f', 10, 64),
			strconv.FormatFloat(stats.ActivePercent, 'f', 10, 64),
			strconv.FormatFloat(stats.CriticalPercent, 'f', 10, 64),
		})
		if err != nil {
			return err
		}
	}
	writer.Flush()
	recordFile.Close()
	fmt.Println("Done")
	return nil
}
