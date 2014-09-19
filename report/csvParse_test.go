package report

import (
	// "log"
	"testing"
)

func TestProcessURL(t *testing.T) {
	samples := map[string]string{
		"http://testing.com": "testing.com",
		"https://abc.com.au": "abc.com.au",
		"test.com/":          "test.com/",
	}
	data := new(RCollection)
	for req, res := range samples {
		if val, ok := data.ProcessUrl(req); ok {
			if val != res {
				t.Errorf("Unexpected result, expected %s, got %s", res, val)
			}

		} else {
			t.Errorf("Failed to process URL: %s, returned %s", req)
		}

	}
}

func TestFromCSV(t *testing.T) {
	data := new(RCollection)
	if ok := data.FromCsv("data3.csv"); !ok {
		t.Errorf("Failed to open CSV file")
	}
	return
}

func TestToCSV(t *testing.T) {

	//create mock data to write into csv file
	ste := Site{"New starter", "E", "HR 4",
		"hr.unimelb.edu.au/__data/assets/word_doc/0003/143949/HR4_updated.doc", "Honorary Appointment"}
	data := RCollection{&([]Site{ste}), 1, 1}
	if ok := data.ToCSV("data3.csv"); !ok {
		t.Errorf("Failed to write to CSV file")
	}
	return

}
