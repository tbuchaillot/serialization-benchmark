package model

import (
	"serialization-benchmark/model/normal_proto"

	"github.com/TykTechnologies/tyk-pump/analytics"
	"github.com/jinzhu/copier"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Record struct {
	analytics.AnalyticsRecord
}

type FutureRecord struct {
	analytics.AnalyticsRecord
	NewData string
}

func (r Record) ToNormalProto() normal_proto.AnalyticsRecord {
	new := normal_proto.AnalyticsRecord{}
	copier.Copy(&new, &r)

	new.TimeStamp = timestamppb.New(r.TimeStamp)

	return new
}
