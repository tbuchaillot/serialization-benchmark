package main

import (
	"encoding/json"
	"math/rand"
	"reflect"
	"serialization-benchmark/demo"
	"serialization-benchmark/model"
	"serialization-benchmark/model/analytics_bebop"
	"serialization-benchmark/model/analyticsflatbuffers"
	"serialization-benchmark/model/normal_proto"
	"serialization-benchmark/model/protobuf-import/github.com/bmizerany/assert"
	"testing"

	"github.com/niubaoshu/gotiny"

	"github.com/golang/protobuf/proto"
	msgpack "gopkg.in/vmihailenco/msgpack.v2"
)

var data []model.Record = demo.GenerateDemoData()
var normalProtoData []normal_proto.AnalyticsRecord = demo.TransformToNormalProto(demo.GenerateDemoData())
var normalFlatData []analyticsflatbuffers.AnalyticsRecord = demo.TransformToFlatBuffer(demo.GenerateDemoData())
var normalBebopData []analytics_bebop.AnalyticsRecord = demo.TransofrmToBebop(demo.GenerateDemoData())

func Benchmark_JSON_Marshal(b *testing.B) {
	b.Helper()
	data := data
	b.ReportAllocs()
	b.ResetTimer()
	var serialSize int

	for n := 0; n < b.N; n++ {
		o := data[rand.Intn(len(data))]
		bytes, _ := json.Marshal(&o)

		serialSize += len(bytes)
	}
	b.ReportMetric(float64(serialSize)/float64(b.N), "B/serial")
}

func Benchmark_Msgpack_Marshal(b *testing.B) {
	b.Helper()
	data := data
	b.ReportAllocs()
	b.ResetTimer()
	var serialSize int

	for n := 0; n < b.N; n++ {
		o := data[rand.Intn(len(data))]
		bytes, _ := msgpack.Marshal(&o)

		serialSize += len(bytes)
	}
	b.ReportMetric(float64(serialSize)/float64(b.N), "B/serial")
}

func Benchmark_NormalProto_Marshal(b *testing.B) {
	b.Helper()
	data := normalProtoData
	b.ReportAllocs()
	b.ResetTimer()
	var serialSize int

	for n := 0; n < b.N; n++ {
		o := normalProtoData[rand.Intn(len(data))]
		bytes, _ := proto.Marshal(&o)

		serialSize += len(bytes)
	}
	b.ReportMetric(float64(serialSize)/float64(b.N), "B/serial")

}

func Benchmark_Bebop_Marshal(b *testing.B) {
	b.Helper()
	b.ReportAllocs()
	b.ResetTimer()
	var serialSize int

	for n := 0; n < b.N; n++ {
		o := normalBebopData[rand.Intn(len(normalBebopData))]
		bytes := o.MarshalBebop()
		serialSize += len(bytes)
	}
	b.ReportMetric(float64(serialSize)/float64(b.N), "B/serial")

}

func Benchmark_MsgpGen_Marshal(b *testing.B) {
	b.Helper()
	data := data
	b.ReportAllocs()
	b.ResetTimer()
	var serialSize int
	for n := 0; n < b.N; n++ {
		o := data[rand.Intn(len(data))]
		bytes, _ := o.MarshalMsg(nil)
		serialSize += len(bytes)
	}
	b.ReportMetric(float64(serialSize)/float64(b.N), "B/serial")
}

func Benchmark_GoTiny_Marshal(b *testing.B) {
	b.Helper()
	data := data
	b.ReportAllocs()
	b.ResetTimer()
	var serialSize int
	ot := reflect.TypeOf(data[rand.Intn(len(data))])
	enc := gotiny.NewEncoderWithType(ot)
	for n := 0; n < b.N; n++ {
		o := data[rand.Intn(len(data))]
		bytes := enc.Encode(&o)

		serialSize += len(bytes)
	}
	b.ReportMetric(float64(serialSize)/float64(b.N), "B/serial")
}

/*
func Benchmark_Flatbuffer_Marshal(b *testing.B) {
	b.Helper()
	data := normalFlatData
	b.ReportAllocs()
	b.ResetTimer()
	var serialSize int

	for n := 0; n < b.N; n++ {
		o := data[rand.Intn(len(data))]


		serialSize += len(bytes)
	}
	b.ReportMetric(float64(serialSize)/float64(b.N), "B/serial")
}
*/

func Test_GoTiny_Future(t *testing.T) {
	//Playing with gotiny
	old := data[rand.Intn(len(data))]

	ot := reflect.TypeOf(old)
	enc := gotiny.NewEncoderWithType(ot)
	dec := gotiny.NewDecoderWithType(ot)

	byt := enc.Encode(old)

	new := model.FutureRecord{}
	dec.Decode(byt, &new)

	assert.Equal(t, old.APIID, new.APIID)
	assert.Equal(t, old.TimeStamp, new.TimeStamp)
	assert.Equal(t, old.Latency, new.Latency)
}

func Test_MsgpGen_Future(t *testing.T) {
	//Playing with gotiny
	old := data[rand.Intn(len(data))]

	bytes, _ := old.MarshalMsg(nil)

	new := model.FutureRecord{}
	new.UnmarshalMsg(bytes)

	assert.Equal(t, old.APIID, new.APIID)
	assert.Equal(t, old.TimeStamp, new.TimeStamp)
	assert.Equal(t, old.Latency, new.Latency)
}
