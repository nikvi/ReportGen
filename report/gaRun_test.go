package report

import (
	// "encoding/json"
	// "log"
	// gD "github.com/vly/go-gadata"
	"testing"
)

const SampleResponse = `{"kind":"analytics#gaData","id":"https://www.googleapis.com/analytics/v3/data/ga?ids=ga:43047246&dimensions=ga:day&metrics=ga:visits&start-date=2014-01-01&end-date=2014-01-02&max-results=100","query":{"start-date":"2014-01-01","end-date":"2014-01-02","ids":"ga:43047246","dimensions":"ga:day","metrics":["ga:visits"],"start-index":1,"max-results":100},"itemsPerPage":100,"totalResults":2,"selfLink":"https://www.googleapis.com/analytics/v3/data/ga?ids=ga:43047246&dimensions=ga:day&metrics=ga:visits&start-date=2014-01-01&end-date=2014-01-02&max-results=100","profileInfo":{"profileId":"43047246","accountId":"21795369","webPropertyId":"UA-21795369-1","internalWebPropertyId":"43063907","profileName":"General - rollup: main","tableId":"ga:43047246"},"containsSampledData":false,"columnHeaders":[{"name":"ga:day","columnType":"DIMENSION","dataType":"STRING"},{"name":"ga:visits","columnType":"METRIC","dataType":"INTEGER"}],"totalsForAllResults":{"ga:visits":"98447"},"rows":[["01","33800"],["02","64647"]]}`

func TestProcessResponse(t *testing.T) {
	// gaTest := new(gD.GAData)
	if _, ok := ProcessResponse(SampleResponse); !ok {
		t.Errorf("Didn't process sample response")
	}
}

func TestToString(t *testing.T) {
	// gaTest := new(gD.GAData)
	// data := new(Responsse)
	// if dt, ok := data.ToString(); ok {
	// 	for _, r := range data {
	// 		t.Errorf(r)
	// 	}
	// 	t.Errorf("did not process")
	// }
}
