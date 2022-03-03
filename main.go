package main

import (
	"encoding/json"
	"fmt"
	"serialization-benchmark/demo"
	"serialization-benchmark/model"
	"serialization-benchmark/model/analytics_bebop"
	"time"

	"github.com/niubaoshu/gotiny"
)

func main() {
	var bebp []analytics_bebop.AnalyticsRecord = demo.TransofrmToBebop(demo.GenerateDemoData())
	var randomData []model.Record = demo.GenerateDemoData()
	fmt.Println("-===== gotiny")
	tiny()
	fmt.Println("===== bbop")
	bbop(bebp)

	fmt.Println("===== gMsgp")
	generatedMsgp(randomData)
}

func generatedMsgp(randomData []model.Record) {
	o := randomData[0]
	bytes, _ := o.MarshalMsg(nil)

	new := model.Record{}
	leftover, err := new.UnmarshalMsg(bytes)
	fmt.Println(leftover)
	fmt.Println(err)
	fmt.Println(new.APIID, new.TimeStamp)
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
func bbop(bebp []analytics_bebop.AnalyticsRecord) {

	bop := analytics_bebop.AnalyticsRecord{}

	if len(bebp) == 0 {
		panic("0")
	}
	jBytes, _ := json.Marshal(bebp[0])
	bop.UnmarshalBebop(jBytes)

	fmt.Println(bop.ApiID, bop.Timestamp)
	fmt.Println()
}
