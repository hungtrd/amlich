package amlich

import (
	"encoding/json"
	"io"
	"os"
	"testing"
)

type TestCase struct {
	Solar []int `json:"solar"`
	Lunar []int `json:"lunar"`
}

func TestSolar2Lunar(t *testing.T) {
	testCases := readTestData()
	for _, v := range testCases {
		d, m, y, leap := Solar2Lunar(v.Solar[0], v.Solar[1], v.Solar[2], v.Solar[3])

		if d != v.Lunar[0] || m != v.Lunar[1] || y != v.Lunar[2] || v.Lunar[3] != leap {
			t.Errorf("Failed test case: %v\nExpect: %v\nGot: %v %v %v %v", v.Solar, v.Lunar, d, m, y, leap)
		}
	}
}

func TestLunar2Solar(t *testing.T) {
	testCases := readTestData()
	for _, v := range testCases {
		d, m, y := Lunar2Solar(v.Lunar[0], v.Lunar[1], v.Lunar[2], v.Lunar[3], 7)

		if d != v.Solar[0] || m != v.Solar[1] || y != v.Solar[2] {
			t.Errorf("Failed test case: %v\nExpect: %v\nGot: %v %v %v", v.Lunar, v.Solar, d, m, y)
		}
	}
}

func readTestData() []TestCase {
	var testCases []TestCase
	jsonFile, err := os.Open("./testdata/testcase.json")
	if err != nil {
		panic(err)
	}
	defer jsonFile.Close()
	byteValue, _ := io.ReadAll(jsonFile)

	err = json.Unmarshal(byteValue, &testCases)
	if err != nil {
		panic(err)
	}
	return testCases
}
