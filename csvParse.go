package main

import (
	"encoding/csv"
	"fmt"
	"os"
	s "strings"
)

type Response struct {
	Category string
	ResType  string
	Code     string
	Url      string
	Name     string
}

type RCollection struct {
	Data   *[]Response
	Size   int
	UtSize int
}

func main() {
	data := new(RCollection)
	data.ReadFromCsv("data1.csv")
	data.GetUrls()
}

//func (c *RCollection) WriteToCSV(fileName string) {}

//make it pointer
func (c *RCollection) GetUrls() []string {

	urlData := make([]string, c.Size)
	for i, resp := range *c.Data {
		urlData[i] = processUrl(resp.Url)
	}
	return urlData
}
func (c *RCollection) ReadFromCsv(fileName string) {

	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	// automatically call Close() at the end of current method
	defer file.Close()

	reader := csv.NewReader(file)
	reader.Comma = ','
	lines, err := reader.ReadAll()
	// if error print and retrun
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	size := len(lines)
	uSize := size * 2
	collection := make([]Response, size)
	for i, line := range lines {
		resp := Response{line[0], line[1], line[2], line[3], line[4]}
		collection[i] = resp
	}

	c.Data = &collection
	c.Size = size
	c.UtSize = uSize
}

func processUrl(x string) (y string) {
	var p = s.Split(x, "//")
	return p[1]
}
