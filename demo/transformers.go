package demo

import (
	"serialization-benchmark/model"
	"serialization-benchmark/model/analyticsflatbuffers"
	"serialization-benchmark/model/normal_proto"

	"github.com/jinzhu/copier"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func TransformToNormalProto(recs []model.Record) []normal_proto.AnalyticsRecord {
	transformedRecs := make([]normal_proto.AnalyticsRecord, len(recs))

	for i, rec := range recs {
		new := normal_proto.AnalyticsRecord{}
		copier.Copy(&new, &rec)

		new.TimeStamp = timestamppb.New(rec.TimeStamp)
		transformedRecs[i] = new
	}
	return transformedRecs
}

func TransformToFlatBuffer(recs []model.Record) []analyticsflatbuffers.AnalyticsRecord {
	transformedRecs := make([]analyticsflatbuffers.AnalyticsRecord, len(recs))

	for i, rec := range recs {
		new := analyticsflatbuffers.AnalyticsRecord{}
		//new
		copier.Copy(&new, &rec)

		//new.TimeStamp = timestamppb.New(rec.TimeStamp)
		transformedRecs[i] = new
	}
	return transformedRecs
}
