package demo

import (
	"serialization-benchmark/model"
	"serialization-benchmark/model/analytics_bebop"
	"serialization-benchmark/model/analyticsflatbuffers"
	"serialization-benchmark/model/normal_proto"

	"github.com/jinzhu/copier"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func TransformToNormalProto(recs []model.Record) []normal_proto.AnalyticsRecord {
	transformedRecs := make([]normal_proto.AnalyticsRecord, len(recs))

	for i, rec := range recs {
		new := normal_proto.AnalyticsRecord{}
		copier.Copy(&new, &recs[i])

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

func TransofrmToBebop(recs []model.Record) []analytics_bebop.AnalyticsRecord {
	transformedRecs := make([]analytics_bebop.AnalyticsRecord, len(recs))

	for i := range recs {

		newRec := &analytics_bebop.AnalyticsRecord{}
		newRec.ApiID = &recs[i].APIID
		newRec.ApiKey = &recs[i].APIKey
		newRec.OrgID = &recs[i].OrgID
		newRec.ApiName = &recs[i].APIName
		newRec.ApiVersion = &recs[i].APIVersion
		newRec.ContentLength = &recs[i].ContentLength
		newRec.RawPath = &recs[i].RawPath
		newRec.RequestTime = &recs[i].RequestTime
		newRec.RawRequest = &recs[i].RawRequest
		newRec.RawResponse = &recs[i].RawResponse
		newRec.Useragent = &recs[i].UserAgent
		newRec.OauthID = &recs[i].OauthID
		newRec.Host = &recs[i].Host
		newRec.ExpireAt = &recs[i].ExpireAt
		newRec.Timestamp = &recs[i].TimeStamp
		newRec.IpAddress = &recs[i].IPAddress
		newRec.TrackPath = &recs[i].TrackPath
		newRec.Tags = &recs[i].Tags

		responseCode := int32(recs[i].ResponseCode)
		newRec.ResponseCode = &responseCode
		d := int32(recs[i].Day)
		h := int32(recs[i].Hour)
		month := int32(recs[i].Month)
		year := int32(recs[i].Year)
		newRec.Day = &d
		newRec.Hour = &h
		newRec.Month = &month
		newRec.Year = &year

		geonameID := uint32(recs[i].Geo.City.GeoNameID)
		newRec.Geo = &analytics_bebop.GeoData{}
		newRec.Geo.City = &analytics_bebop.City{
			GeoNameID: &geonameID,
		}
		newRec.Geo.Location = &analytics_bebop.Location{
			Latitude:  &recs[i].Geo.Location.Latitude,
			Longitude: &recs[i].Geo.Location.Longitude,
			Timezone:  &recs[i].Geo.Location.TimeZone,
		}
		newRec.Geo.Country = &analytics_bebop.Country{
			ISOCode: &recs[i].Geo.Country.ISOCode,
		}

		newRec.Network = &analytics_bebop.NetworkStats{}
		newRec.Network.BytesIn = &recs[i].Network.BytesIn
		newRec.Network.BytesOut = &recs[i].Network.BytesOut
		newRec.Network.ClosedConnections = &recs[i].Network.ClosedConnection
		newRec.Network.OpenConnections = &recs[i].Network.OpenConnections
		newRec.Latency = &analytics_bebop.Latency{}
		newRec.Latency.Total = &recs[i].Latency.Total
		newRec.Latency.Upstream = &recs[i].Latency.Upstream

		transformedRecs[i] = *newRec
	}
	return transformedRecs

}
