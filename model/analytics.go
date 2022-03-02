package model

import (
	"time"
)

type Record struct {
	Method        string       `json:"method" gorm:"column:method"`
	Host          string       `json:"host" gorm:"column:host"`
	Path          string       `json:"path" gorm:"column:path"`
	RawPath       string       `json:"raw_path" gorm:"column:rawpath"`
	ContentLength int64        `json:"content_length" gorm:"column:contentlength"`
	UserAgent     string       `json:"user_agent" gorm:"column:useragent"`
	Day           int          `json:"day" sql:"-"`
	Month         time.Month   `json:"month" sql:"-"`
	Year          int          `json:"year" sql:"-"`
	Hour          int          `json:"hour" sql:"-"`
	ResponseCode  int          `json:"response_code" gorm:"column:responsecode;index"`
	APIKey        string       `json:"api_key" gorm:"column:apikey;index"`
	TimeStamp     time.Time    `json:"timestamp" gorm:"column:timestamp;index"`
	APIVersion    string       `json:"api_version" gorm:"column:apiversion"`
	APIName       string       `json:"api_name" sql:"-"`
	APIID         string       `json:"api_id" gorm:"column:apiid;index"`
	OrgID         string       `json:"org_id" gorm:"column:orgid;index"`
	OauthID       string       `json:"oauth_id" gorm:"column:oauthid;index"`
	RequestTime   int64        `json:"request_time" gorm:"column:requesttime"`
	RawRequest    string       `json:"raw_request" gorm:"column:rawrequest"`
	RawResponse   string       `json:"raw_response" gorm:"column:rawresponse"`
	IPAddress     string       `json:"ip_address" gorm:"column:ipaddress"`
	Geo           GeoData      `json:"geo" gorm:"embedded"`
	Network       NetworkStats `json:"network"`
	Latency       Latency      `json:"latency"`
	Tags          []string     `json:"tags"`
	Alias         string       `json:"alias"`
	TrackPath     bool         `json:"track_path" gorm:"column:trackpath"`
	ExpireAt      time.Time    `bson:"expireAt" json:"expireAt"`
}
type Country struct {
	ISOCode string `maxminddb:"iso_code" json:"iso_code"`
}

type GeoData struct {
	Country Country `maxminddb:"country" json:"country"`

	City struct {
		GeoNameID uint              `maxminddb:"geoname_id" json:"geoname_id"`
		Names     map[string]string `maxminddb:"names" json:"names"`
	} `maxminddb:"city" json:"city"`

	Location struct {
		Latitude  float64 `maxminddb:"latitude" json:"latitude"`
		Longitude float64 `maxminddb:"longitude" json:"longitude"`
		TimeZone  string  `maxminddb:"time_zone" json:"time_zone"`
	} `maxminddb:"location" json:"location"`
}
type NetworkStats struct {
	OpenConnections  int64 `json:"open_connections"`
	ClosedConnection int64 `json:"closed_connections"`
	BytesIn          int64 `json:"bytes_in"`
	BytesOut         int64 `json:"bytes_out"`
}

type Latency struct {
	Total    int64 `json:"total"`
	Upstream int64 `json:"upstream"`
}
type FutureRecord struct {
	Record
	NewData string
}
