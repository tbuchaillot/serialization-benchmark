package main

import (
	"fmt"
	"serialization-benchmark/model"
	"time"

	"github.com/niubaoshu/gotiny"
)

func main() {
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
