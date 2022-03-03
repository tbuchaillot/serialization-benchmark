package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"reflect"
	"serialization-benchmark/demo"
	"serialization-benchmark/model"
	"serialization-benchmark/model/analytics_bebop"

	"github.com/niubaoshu/gotiny"
)

func main() {
	var bebp []analytics_bebop.AnalyticsRecord = demo.TransofrmToBebop(demo.GenerateDemoData())
	var randomData []model.Record = demo.GenerateDemoData()
	fmt.Println("-===== gotiny")
	tiny(randomData)
	fmt.Println("===== bbop")
	bbop(bebp)

	fmt.Println("===== gMsgp")
	generatedMsgp(randomData)
}

func generatedMsgp(randomData []model.Record) {
	//Playing with gotiny
	old := randomData[rand.Intn(len(randomData))]

	bytes, _ := old.MarshalMsg(nil)

	new := model.FutureRecord{}
	new.UnmarshalMsg(bytes)

	fmt.Println(new.APIID, new.TimeStamp)
}

func tiny(randomData []model.Record) {
	//Playing with gotiny
	ot := reflect.TypeOf(randomData[rand.Intn(len(randomData))])
	enc := gotiny.NewEncoderWithType(ot)
	dec := gotiny.NewDecoderWithType(ot)

	byt := enc.Encode(randomData[rand.Intn(len(randomData))])

	newLots := model.FutureRecord{}
	dec.Decode(byt, &newLots)

	fmt.Println(newLots.APIID, newLots.TimeStamp)

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
