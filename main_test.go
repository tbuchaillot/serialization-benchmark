package main

import (
	"encoding/json"
	"math/rand"
	"serialization-benchmark/demo"
	"serialization-benchmark/model"
	"serialization-benchmark/model/analytics_bebop"
	"serialization-benchmark/model/analyticsflatbuffers"
	"serialization-benchmark/model/normal_proto"
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

func Benchmark_GoTiny_Marshal(b *testing.B) {
	b.Helper()
	data := data
	b.ReportAllocs()
	b.ResetTimer()
	var serialSize int

	for n := 0; n < b.N; n++ {
		o := data[rand.Intn(len(data))]
		bytes := gotiny.Marshal(&o)

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
