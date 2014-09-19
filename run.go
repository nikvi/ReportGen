package main

import "./report"

func main() {
	data := new(report.RCollection)
	data.FromCsv("data1.csv")
	data.ToCSV("data4.csv")
}
