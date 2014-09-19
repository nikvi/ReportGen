package report

import (
	"encoding/json"
	"fmt"
	gD "github.com/vly/go-gadata"
)

type Response struct {
	Columns []struct {
		Name  string `json:"name"`
		CType string `json:"columnType"`
		DType string `json:"dataType"`
	} `json:"columnHeaders"`
	Total map[string]string `json:"totalsForAllResults"`
	Rows  [][]string        `json:"rows"`
}

type GATest struct {
	Resp string
}

func (resp *Response) ToString() (rpData []string, ok bool) {

	header := ""
	for _, cls := range resp.Columns {
		header += cls.Name + ", "
	}
	data := make([]string, len(resp.Rows))
	for i, cols := range resp.Rows {
		for _, s := range cols {
			data[i] += s
		}
	}
	rpData := make([]string, len(resp.Rows)+1)
	// need to combine both

	ok = true
	return
}

func (g *GATest) Query() bool {
	// initialise GAData
	gaTest := new(gD.GAData)

	// initialise instance incl. authentication
	gaTest.Init()

	// build a basic GA query, replace your ga ID
	testRequest := gD.GaRequest{"ga:43047246", // GA id
		"2014-01-01", // start date
		"2014-01-02", // end date
		"ga:visits",  // metrics
		"ga:day",     // dimensions
		"",           // filters
		"",           // segments
		"",           // sort
		100}          // results no.

	// launch data
	result := gaTest.GetData(1, &testRequest)
	fmt.Printf("results: %s\n", result)
	return true
}

func ProcessResponse(input string) (data Response, ok bool) {
	tmp := new(Response)
	if err := json.Unmarshal([]byte(input), tmp); err == nil {
		//fmt.Printf("%s, %v", input, tmp)
		//fmt.Printf("%v", tmp)
		if data, ok := tmp.ToString; ok {
			for _, r := range data {
				fmt.Printf(r)
			}
		}
	} else {
		fmt.Printf(err.Error())
	}
	return
}
