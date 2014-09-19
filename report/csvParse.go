package report

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	s "strings"
)

type Site struct {
	Category string
	ResType  string
	Code     string
	Url      string
	Name     string
}

type RCollection struct {
	Data   *[]Site
	Size   int
	UtSize int
}

func (c *RCollection) ToCSV(fileName string) (ok bool) {

	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close()
	writer := csv.NewWriter(file)
	for _, wd := range *(c.Data) {
		if err = writer.Write([]string{wd.Category, wd.ResType, wd.Code, wd.Url, wd.Name}); err != nil {
			log.Println("Error:", err.Error)
			return
		}
		log.Println("wrote")
	}
	ok = true
	writer.Flush()
	return
}
func (c *RCollection) PrintData() {

	for _, rp := range *c.Data {

		fmt.Println(rp.Category)
		fmt.Println(rp.Url)

	}
}
func (c *RCollection) FromCsv(fileName string) (ok bool) {

	file, err := os.Open(fileName)
	if err != nil {
		log.Println("Error:", err.Error)
		return
	}
	// automatically call Close() at the end of current method
	defer file.Close()

	reader := csv.NewReader(file)
	reader.Comma = ','
	lines, err := reader.ReadAll()
	// if error print and retrun
	if err != nil {
		log.Println("Error:", err.Error)
		return
	}

	size := len(lines)
	uSize := size * 2
	collection := make([]Site, size)
	for i, line := range lines {
		if v, ok := c.ProcessUrl(line[3]); ok {
			resp := Site{line[0], line[1], line[2], v, line[4]}
			collection[i] = resp
		}
	}

	c.Data = &collection
	c.Size = size
	c.UtSize = uSize

	ok = true
	return
}

func (c *RCollection) ProcessUrl(x string) (y string, ok bool) {
	if len(x) > 0 {
		p := s.Split(x, "//")

		if len(p) == 2 {
			y = p[1]
		} else {
			y = x
		}
		ok = true
	}
	return
}
