package main

import (
	"fmt"
	"serialization-benchmark/model"
	"time"
)

func main() {
	test := model.Record{}
	test.APIID = "api_1"
	test.TimeStamp = time.Now()
	fmt.Println("test:", test.APIID, "-", "-", test.TimeStamp)

	proto := test.ToNormalProto()

	fmt.Println("newTest:", proto.APIID, "-", "-", proto.TimeStamp.AsTime())

}

/*
 - Bebop
 - gotiny (chino)

*/
