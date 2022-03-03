//go:generate msgp
package model

import (
	"time"
)

type Record struct {
	Method        string       `json:"method" msg:"method" gorm:"column:method"`
	Host          string       `json:"host" msg:"host" gorm:"column:host"`
	Path          string       `json:"path" msg:"path" gorm:"column:path"`
	RawPath       string       `json:"raw_path" msg:"raw_path" gorm:"column:rawpath"`
	ContentLength int64        `json:"content_length" msg:"content_length" gorm:"column:contentlength"`
	UserAgent     string       `json:"user_agent" msg:"user_agent" gorm:"column:useragent"`
	Day           int          `json:"day" msg:"day" sql:"-"`
	Month         int          `json:"month"  msg:"month" sql:"-"`
	Year          int          `json:"year" msg:"year" sql:"-"`
	Hour          int          `json:"hour" msg:"hour" sql:"-"`
	ResponseCode  int          `json:"response_code" msg:"response_code" gorm:"column:responsecode;index"`
	APIKey        string       `json:"api_key" msg:"api_key" gorm:"column:apikey;index"`
	TimeStamp     time.Time    `json:"timestamp" msg:"timestamp" gorm:"column:timestamp;index"`
	APIVersion    string       `json:"api_version" msg:"api_version" gorm:"column:apiversion"`
	APIName       string       `json:"api_name" msg:"api_name" sql:"-"`
	APIID         string       `json:"api_id" msg:"api_id" gorm:"column:apiid;index"`
	OrgID         string       `json:"org_id" msg:"org_id" gorm:"column:orgid;index"`
	OauthID       string       `json:"oauth_id" msg:"oauth_id" gorm:"column:oauthid;index"`
	RequestTime   int64        `json:"request_time" msg:"request_time" gorm:"column:requesttime"`
	RawRequest    string       `json:"raw_request" msg:"raw_request" gorm:"column:rawrequest"`
	RawResponse   string       `json:"raw_response" msg:"raw_response" gorm:"column:rawresponse"`
	IPAddress     string       `json:"ip_address" msg:"ip_address" gorm:"column:ipaddress"`
	Geo           GeoData      `json:"geo" msg:"geo" gorm:"embedded"`
	Network       NetworkStats `json:"network" msg:"network" `
	Latency       Latency      `json:"latency" msg:"latency" `
	Tags          []string     `json:"tags" msg:"tags" `
	Alias         string       `json:"alias" msg:"alias" `
	TrackPath     bool         `json:"track_path" msg:"track_path" gorm:"column:trackpath"`
	ExpireAt      time.Time    `bson:"expireAt" msg:"expireAt" json:"expireAt"`
}
type Country struct {
	ISOCode string `maxminddb:"iso_code" msg:"iso_code" json:"iso_code"`
}

type GeoData struct {
	Country Country `maxminddb:"country" msg:"country" json:"country"`

	City struct {
		GeoNameID uint              `maxminddb:"geoname_id" msg:"geoname_id" json:"geoname_id"`
		Names     map[string]string `maxminddb:"names" msg:"names" json:"names"`
	} `maxminddb:"city" msg:"city" json:"city"`

	Location struct {
		Latitude  float64 `maxminddb:"latitude" msg:"latitude" json:"latitude"`
		Longitude float64 `maxminddb:"longitude" msg:"longitude" json:"longitude"`
		TimeZone  string  `maxminddb:"time_zone" msg:"time_zone" json:"time_zone"`
	} `maxminddb:"location" msg:"location" json:"location"`
}
type NetworkStats struct {
	OpenConnections  int64 `msg:"open_connections" json:"open_connections"`
	ClosedConnection int64 `msg:"closed_connections" json:"closed_connections"`
	BytesIn          int64 `msg:"bytes_in" json:"bytes_in"`
	BytesOut         int64 `msg:"bytes_out" json:"bytes_out"`
}

type Latency struct {
	Total    int64 `msg:"total" json:"total"`
	Upstream int64 `msg:"upstream" json:"upstream"`
}
type FutureRecord struct {
	Record
	NewData string
}
