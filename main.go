package main

import (
	"encoding/json"
	"fmt"
	"serialization-benchmark/model"
	"serialization-benchmark/model/analytics_bebop"
	"time"

	"github.com/niubaoshu/gotiny"
)

func main() {
	test := model.Record{}
	test.APIID = "api_1"
	test.TimeStamp = time.Now()

	bop := analytics_bebop.AnalyticsRecord{}

	jBytes, _ := json.Marshal(test)
	fmt.Println(string(jBytes))
	bop.UnmarshalBebop(jBytes)

	fmt.Println(bop.ApiID, bop.Timestamp)
}

func tiny() {
	//Playing with gotiny
	test := model.Record{}
	test.APIID = "api_1"
	test.TimeStamp = time.Now()

	lot := []model.Record{}
	lot = append(lot, test)
	test.APIID = "api_2"
	lot = append(lot, test)

	fmt.Println("test:", lot)

	byt := gotiny.Marshal(&lot)

	newLots := []model.Record{}
	gotiny.Unmarshal(byt, &newLots)

	fmt.Println("newTest:", newLots)

}
