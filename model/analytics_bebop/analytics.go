// Code generated by bebopc-go; DO NOT EDIT.

package analytics_bebop

import (
	"io"
	"time"
	"github.com/200sc/bebop"
	"github.com/200sc/bebop/iohelp"
)

var _ bebop.Record = &AnalyticsRecord{}

type AnalyticsRecord struct {
	Method *string
	Host *string
	Path *string
	RawPath *string
	ContentLength *int64
	Network *NetworkStats
	Timestamp *time.Time
	ExpireAt *time.Time
	Useragent *string
	Day *int32
	Year *int32
	Month *int32
	Hour *int32
	ResponseCode *int32
	ApiKey *string
	ApiVersion *string
	ApiName *string
	ApiID *string
	OrgID *string
	OauthID *string
	RequestTime *int64
	RawRequest *string
	RawResponse *string
	IpAddress *string
	Geo *GeoData
	Tags *[]string
	Latency *Latency
	TrackPath *bool
}

func (bbp AnalyticsRecord) MarshalBebopTo(buf []byte) int {
	at := 0
	iohelp.WriteUint32Bytes(buf[at:], uint32(bbp.Size()-4))
	at += 4
	if bbp.Method != nil {
		buf[at] = 1
		at++
		iohelp.WriteUint32Bytes(buf[at:], uint32(len(*bbp.Method)))
		copy(buf[at+4:at+4+len(*bbp.Method)], []byte(*bbp.Method))
		at += 4 + len(*bbp.Method)
	}
	if bbp.Host != nil {
		buf[at] = 2
		at++
		iohelp.WriteUint32Bytes(buf[at:], uint32(len(*bbp.Host)))
		copy(buf[at+4:at+4+len(*bbp.Host)], []byte(*bbp.Host))
		at += 4 + len(*bbp.Host)
	}
	if bbp.Path != nil {
		buf[at] = 3
		at++
		iohelp.WriteUint32Bytes(buf[at:], uint32(len(*bbp.Path)))
		copy(buf[at+4:at+4+len(*bbp.Path)], []byte(*bbp.Path))
		at += 4 + len(*bbp.Path)
	}
	if bbp.RawPath != nil {
		buf[at] = 4
		at++
		iohelp.WriteUint32Bytes(buf[at:], uint32(len(*bbp.RawPath)))
		copy(buf[at+4:at+4+len(*bbp.RawPath)], []byte(*bbp.RawPath))
		at += 4 + len(*bbp.RawPath)
	}
	if bbp.ContentLength != nil {
		buf[at] = 5
		at++
		iohelp.WriteInt64Bytes(buf[at:], *bbp.ContentLength)
		at += 8
	}
	if bbp.Network != nil {
		buf[at] = 6
		at++
		(*bbp.Network).MarshalBebopTo(buf[at:])
		at += (*bbp.Network).Size()
	}
	if bbp.Timestamp != nil {
		buf[at] = 7
		at++
		if *bbp.Timestamp != (time.Time{}) {
			iohelp.WriteInt64Bytes(buf[at:], ((*bbp.Timestamp).UnixNano()/100))
		} else {
			iohelp.WriteInt64Bytes(buf[at:], 0)
		}
		at += 8
	}
	if bbp.ExpireAt != nil {
		buf[at] = 8
		at++
		if *bbp.ExpireAt != (time.Time{}) {
			iohelp.WriteInt64Bytes(buf[at:], ((*bbp.ExpireAt).UnixNano()/100))
		} else {
			iohelp.WriteInt64Bytes(buf[at:], 0)
		}
		at += 8
	}
	if bbp.Useragent != nil {
		buf[at] = 9
		at++
		iohelp.WriteUint32Bytes(buf[at:], uint32(len(*bbp.Useragent)))
		copy(buf[at+4:at+4+len(*bbp.Useragent)], []byte(*bbp.Useragent))
		at += 4 + len(*bbp.Useragent)
	}
	if bbp.Day != nil {
		buf[at] = 10
		at++
		iohelp.WriteInt32Bytes(buf[at:], *bbp.Day)
		at += 4
	}
	if bbp.Year != nil {
		buf[at] = 11
		at++
		iohelp.WriteInt32Bytes(buf[at:], *bbp.Year)
		at += 4
	}
	if bbp.Month != nil {
		buf[at] = 12
		at++
		iohelp.WriteInt32Bytes(buf[at:], *bbp.Month)
		at += 4
	}
	if bbp.Hour != nil {
		buf[at] = 13
		at++
		iohelp.WriteInt32Bytes(buf[at:], *bbp.Hour)
		at += 4
	}
	if bbp.ResponseCode != nil {
		buf[at] = 14
		at++
		iohelp.WriteInt32Bytes(buf[at:], *bbp.ResponseCode)
		at += 4
	}
	if bbp.ApiKey != nil {
		buf[at] = 15
		at++
		iohelp.WriteUint32Bytes(buf[at:], uint32(len(*bbp.ApiKey)))
		copy(buf[at+4:at+4+len(*bbp.ApiKey)], []byte(*bbp.ApiKey))
		at += 4 + len(*bbp.ApiKey)
	}
	if bbp.ApiVersion != nil {
		buf[at] = 16
		at++
		iohelp.WriteUint32Bytes(buf[at:], uint32(len(*bbp.ApiVersion)))
		copy(buf[at+4:at+4+len(*bbp.ApiVersion)], []byte(*bbp.ApiVersion))
		at += 4 + len(*bbp.ApiVersion)
	}
	if bbp.ApiName != nil {
		buf[at] = 17
		at++
		iohelp.WriteUint32Bytes(buf[at:], uint32(len(*bbp.ApiName)))
		copy(buf[at+4:at+4+len(*bbp.ApiName)], []byte(*bbp.ApiName))
		at += 4 + len(*bbp.ApiName)
	}
	if bbp.ApiID != nil {
		buf[at] = 18
		at++
		iohelp.WriteUint32Bytes(buf[at:], uint32(len(*bbp.ApiID)))
		copy(buf[at+4:at+4+len(*bbp.ApiID)], []byte(*bbp.ApiID))
		at += 4 + len(*bbp.ApiID)
	}
	if bbp.OrgID != nil {
		buf[at] = 19
		at++
		iohelp.WriteUint32Bytes(buf[at:], uint32(len(*bbp.OrgID)))
		copy(buf[at+4:at+4+len(*bbp.OrgID)], []byte(*bbp.OrgID))
		at += 4 + len(*bbp.OrgID)
	}
	if bbp.OauthID != nil {
		buf[at] = 20
		at++
		iohelp.WriteUint32Bytes(buf[at:], uint32(len(*bbp.OauthID)))
		copy(buf[at+4:at+4+len(*bbp.OauthID)], []byte(*bbp.OauthID))
		at += 4 + len(*bbp.OauthID)
	}
	if bbp.RequestTime != nil {
		buf[at] = 21
		at++
		iohelp.WriteInt64Bytes(buf[at:], *bbp.RequestTime)
		at += 8
	}
	if bbp.RawRequest != nil {
		buf[at] = 22
		at++
		iohelp.WriteUint32Bytes(buf[at:], uint32(len(*bbp.RawRequest)))
		copy(buf[at+4:at+4+len(*bbp.RawRequest)], []byte(*bbp.RawRequest))
		at += 4 + len(*bbp.RawRequest)
	}
	if bbp.RawResponse != nil {
		buf[at] = 23
		at++
		iohelp.WriteUint32Bytes(buf[at:], uint32(len(*bbp.RawResponse)))
		copy(buf[at+4:at+4+len(*bbp.RawResponse)], []byte(*bbp.RawResponse))
		at += 4 + len(*bbp.RawResponse)
	}
	if bbp.IpAddress != nil {
		buf[at] = 24
		at++
		iohelp.WriteUint32Bytes(buf[at:], uint32(len(*bbp.IpAddress)))
		copy(buf[at+4:at+4+len(*bbp.IpAddress)], []byte(*bbp.IpAddress))
		at += 4 + len(*bbp.IpAddress)
	}
	if bbp.Geo != nil {
		buf[at] = 25
		at++
		(*bbp.Geo).MarshalBebopTo(buf[at:])
		at += (*bbp.Geo).Size()
	}
	if bbp.Tags != nil {
		buf[at] = 26
		at++
		iohelp.WriteUint32Bytes(buf[at:], uint32(len(*bbp.Tags)))
		at += 4
		for _, v2 := range *bbp.Tags {
			iohelp.WriteUint32Bytes(buf[at:], uint32(len(v2)))
			copy(buf[at+4:at+4+len(v2)], []byte(v2))
			at += 4 + len(v2)
		}
	}
	if bbp.Latency != nil {
		buf[at] = 27
		at++
		(*bbp.Latency).MarshalBebopTo(buf[at:])
		at += (*bbp.Latency).Size()
	}
	if bbp.TrackPath != nil {
		buf[at] = 28
		at++
		iohelp.WriteBoolBytes(buf[at:], *bbp.TrackPath)
		at += 1
	}
	return at
}

func (bbp *AnalyticsRecord) UnmarshalBebop(buf []byte) (err error) {
	at := 0
	_ = iohelp.ReadUint32Bytes(buf[at:])
	buf = buf[4:]
	for {
		switch buf[at] {
		case 1:
			at += 1
			bbp.Method = new(string)
			(*bbp.Method), err = iohelp.ReadStringBytes(buf[at:])
			if err != nil{
				return err
			}
			at += 4 + len((*bbp.Method))
		case 2:
			at += 1
			bbp.Host = new(string)
			(*bbp.Host), err = iohelp.ReadStringBytes(buf[at:])
			if err != nil{
				return err
			}
			at += 4 + len((*bbp.Host))
		case 3:
			at += 1
			bbp.Path = new(string)
			(*bbp.Path), err = iohelp.ReadStringBytes(buf[at:])
			if err != nil{
				return err
			}
			at += 4 + len((*bbp.Path))
		case 4:
			at += 1
			bbp.RawPath = new(string)
			(*bbp.RawPath), err = iohelp.ReadStringBytes(buf[at:])
			if err != nil{
				return err
			}
			at += 4 + len((*bbp.RawPath))
		case 5:
			at += 1
			bbp.ContentLength = new(int64)
			if len(buf[at:]) < 8 {
				 return io.ErrUnexpectedEOF
			}
			(*bbp.ContentLength) = iohelp.ReadInt64Bytes(buf[at:])
			at += 8
		case 6:
			at += 1
			bbp.Network = new(NetworkStats)
			(*bbp.Network), err = MakeNetworkStatsFromBytes(buf[at:])
			if err != nil{
				return err
			}
			at += ((*bbp.Network)).Size()
		case 7:
			at += 1
			bbp.Timestamp = new(time.Time)
			if len(buf[at:]) < 8 {
				 return io.ErrUnexpectedEOF
			}
			(*bbp.Timestamp) = iohelp.ReadDateBytes(buf[at:])
			at += 8
		case 8:
			at += 1
			bbp.ExpireAt = new(time.Time)
			if len(buf[at:]) < 8 {
				 return io.ErrUnexpectedEOF
			}
			(*bbp.ExpireAt) = iohelp.ReadDateBytes(buf[at:])
			at += 8
		case 9:
			at += 1
			bbp.Useragent = new(string)
			(*bbp.Useragent), err = iohelp.ReadStringBytes(buf[at:])
			if err != nil{
				return err
			}
			at += 4 + len((*bbp.Useragent))
		case 10:
			at += 1
			bbp.Day = new(int32)
			if len(buf[at:]) < 4 {
				 return io.ErrUnexpectedEOF
			}
			(*bbp.Day) = iohelp.ReadInt32Bytes(buf[at:])
			at += 4
		case 11:
			at += 1
			bbp.Year = new(int32)
			if len(buf[at:]) < 4 {
				 return io.ErrUnexpectedEOF
			}
			(*bbp.Year) = iohelp.ReadInt32Bytes(buf[at:])
			at += 4
		case 12:
			at += 1
			bbp.Month = new(int32)
			if len(buf[at:]) < 4 {
				 return io.ErrUnexpectedEOF
			}
			(*bbp.Month) = iohelp.ReadInt32Bytes(buf[at:])
			at += 4
		case 13:
			at += 1
			bbp.Hour = new(int32)
			if len(buf[at:]) < 4 {
				 return io.ErrUnexpectedEOF
			}
			(*bbp.Hour) = iohelp.ReadInt32Bytes(buf[at:])
			at += 4
		case 14:
			at += 1
			bbp.ResponseCode = new(int32)
			if len(buf[at:]) < 4 {
				 return io.ErrUnexpectedEOF
			}
			(*bbp.ResponseCode) = iohelp.ReadInt32Bytes(buf[at:])
			at += 4
		case 15:
			at += 1
			bbp.ApiKey = new(string)
			(*bbp.ApiKey), err = iohelp.ReadStringBytes(buf[at:])
			if err != nil{
				return err
			}
			at += 4 + len((*bbp.ApiKey))
		case 16:
			at += 1
			bbp.ApiVersion = new(string)
			(*bbp.ApiVersion), err = iohelp.ReadStringBytes(buf[at:])
			if err != nil{
				return err
			}
			at += 4 + len((*bbp.ApiVersion))
		case 17:
			at += 1
			bbp.ApiName = new(string)
			(*bbp.ApiName), err = iohelp.ReadStringBytes(buf[at:])
			if err != nil{
				return err
			}
			at += 4 + len((*bbp.ApiName))
		case 18:
			at += 1
			bbp.ApiID = new(string)
			(*bbp.ApiID), err = iohelp.ReadStringBytes(buf[at:])
			if err != nil{
				return err
			}
			at += 4 + len((*bbp.ApiID))
		case 19:
			at += 1
			bbp.OrgID = new(string)
			(*bbp.OrgID), err = iohelp.ReadStringBytes(buf[at:])
			if err != nil{
				return err
			}
			at += 4 + len((*bbp.OrgID))
		case 20:
			at += 1
			bbp.OauthID = new(string)
			(*bbp.OauthID), err = iohelp.ReadStringBytes(buf[at:])
			if err != nil{
				return err
			}
			at += 4 + len((*bbp.OauthID))
		case 21:
			at += 1
			bbp.RequestTime = new(int64)
			if len(buf[at:]) < 8 {
				 return io.ErrUnexpectedEOF
			}
			(*bbp.RequestTime) = iohelp.ReadInt64Bytes(buf[at:])
			at += 8
		case 22:
			at += 1
			bbp.RawRequest = new(string)
			(*bbp.RawRequest), err = iohelp.ReadStringBytes(buf[at:])
			if err != nil{
				return err
			}
			at += 4 + len((*bbp.RawRequest))
		case 23:
			at += 1
			bbp.RawResponse = new(string)
			(*bbp.RawResponse), err = iohelp.ReadStringBytes(buf[at:])
			if err != nil{
				return err
			}
			at += 4 + len((*bbp.RawResponse))
		case 24:
			at += 1
			bbp.IpAddress = new(string)
			(*bbp.IpAddress), err = iohelp.ReadStringBytes(buf[at:])
			if err != nil{
				return err
			}
			at += 4 + len((*bbp.IpAddress))
		case 25:
			at += 1
			bbp.Geo = new(GeoData)
			(*bbp.Geo), err = MakeGeoDataFromBytes(buf[at:])
			if err != nil{
				return err
			}
			at += ((*bbp.Geo)).Size()
		case 26:
			at += 1
			bbp.Tags = new([]string)
			if len(buf[at:]) < 4 {
				 return io.ErrUnexpectedEOF
			}
			(*bbp.Tags) = make([]string, iohelp.ReadUint32Bytes(buf[at:]))
			at += 4
			for i3 := range (*bbp.Tags) {
				((*bbp.Tags))[i3], err = iohelp.ReadStringBytes(buf[at:])
				if err != nil{
					return err
				}
				at += 4 + len(((*bbp.Tags))[i3])
			}
		case 27:
			at += 1
			bbp.Latency = new(Latency)
			(*bbp.Latency), err = MakeLatencyFromBytes(buf[at:])
			if err != nil{
				return err
			}
			at += ((*bbp.Latency)).Size()
		case 28:
			at += 1
			bbp.TrackPath = new(bool)
			if len(buf[at:]) < 1 {
				 return io.ErrUnexpectedEOF
			}
			(*bbp.TrackPath) = iohelp.ReadBoolBytes(buf[at:])
			at += 1
		default:
			return nil
		}
	}
}

func (bbp AnalyticsRecord) EncodeBebop(iow io.Writer) (err error) {
	w := iohelp.NewErrorWriter(iow)
	iohelp.WriteUint32(w, uint32(bbp.Size()-4))
	if bbp.Method != nil {
		w.Write([]byte{1})
		iohelp.WriteUint32(w, uint32(len(*bbp.Method)))
		w.Write([]byte(*bbp.Method))
	}
	if bbp.Host != nil {
		w.Write([]byte{2})
		iohelp.WriteUint32(w, uint32(len(*bbp.Host)))
		w.Write([]byte(*bbp.Host))
	}
	if bbp.Path != nil {
		w.Write([]byte{3})
		iohelp.WriteUint32(w, uint32(len(*bbp.Path)))
		w.Write([]byte(*bbp.Path))
	}
	if bbp.RawPath != nil {
		w.Write([]byte{4})
		iohelp.WriteUint32(w, uint32(len(*bbp.RawPath)))
		w.Write([]byte(*bbp.RawPath))
	}
	if bbp.ContentLength != nil {
		w.Write([]byte{5})
		iohelp.WriteInt64(w, *bbp.ContentLength)
	}
	if bbp.Network != nil {
		w.Write([]byte{6})
		err = (*bbp.Network).EncodeBebop(w)
		if err != nil{
			return err
		}
	}
	if bbp.Timestamp != nil {
		w.Write([]byte{7})
		if *bbp.Timestamp != (time.Time{}) {
			iohelp.WriteInt64(w, ((*bbp.Timestamp).UnixNano()/100))
		} else {
			iohelp.WriteInt64(w, 0)
		}
	}
	if bbp.ExpireAt != nil {
		w.Write([]byte{8})
		if *bbp.ExpireAt != (time.Time{}) {
			iohelp.WriteInt64(w, ((*bbp.ExpireAt).UnixNano()/100))
		} else {
			iohelp.WriteInt64(w, 0)
		}
	}
	if bbp.Useragent != nil {
		w.Write([]byte{9})
		iohelp.WriteUint32(w, uint32(len(*bbp.Useragent)))
		w.Write([]byte(*bbp.Useragent))
	}
	if bbp.Day != nil {
		w.Write([]byte{10})
		iohelp.WriteInt32(w, *bbp.Day)
	}
	if bbp.Year != nil {
		w.Write([]byte{11})
		iohelp.WriteInt32(w, *bbp.Year)
	}
	if bbp.Month != nil {
		w.Write([]byte{12})
		iohelp.WriteInt32(w, *bbp.Month)
	}
	if bbp.Hour != nil {
		w.Write([]byte{13})
		iohelp.WriteInt32(w, *bbp.Hour)
	}
	if bbp.ResponseCode != nil {
		w.Write([]byte{14})
		iohelp.WriteInt32(w, *bbp.ResponseCode)
	}
	if bbp.ApiKey != nil {
		w.Write([]byte{15})
		iohelp.WriteUint32(w, uint32(len(*bbp.ApiKey)))
		w.Write([]byte(*bbp.ApiKey))
	}
	if bbp.ApiVersion != nil {
		w.Write([]byte{16})
		iohelp.WriteUint32(w, uint32(len(*bbp.ApiVersion)))
		w.Write([]byte(*bbp.ApiVersion))
	}
	if bbp.ApiName != nil {
		w.Write([]byte{17})
		iohelp.WriteUint32(w, uint32(len(*bbp.ApiName)))
		w.Write([]byte(*bbp.ApiName))
	}
	if bbp.ApiID != nil {
		w.Write([]byte{18})
		iohelp.WriteUint32(w, uint32(len(*bbp.ApiID)))
		w.Write([]byte(*bbp.ApiID))
	}
	if bbp.OrgID != nil {
		w.Write([]byte{19})
		iohelp.WriteUint32(w, uint32(len(*bbp.OrgID)))
		w.Write([]byte(*bbp.OrgID))
	}
	if bbp.OauthID != nil {
		w.Write([]byte{20})
		iohelp.WriteUint32(w, uint32(len(*bbp.OauthID)))
		w.Write([]byte(*bbp.OauthID))
	}
	if bbp.RequestTime != nil {
		w.Write([]byte{21})
		iohelp.WriteInt64(w, *bbp.RequestTime)
	}
	if bbp.RawRequest != nil {
		w.Write([]byte{22})
		iohelp.WriteUint32(w, uint32(len(*bbp.RawRequest)))
		w.Write([]byte(*bbp.RawRequest))
	}
	if bbp.RawResponse != nil {
		w.Write([]byte{23})
		iohelp.WriteUint32(w, uint32(len(*bbp.RawResponse)))
		w.Write([]byte(*bbp.RawResponse))
	}
	if bbp.IpAddress != nil {
		w.Write([]byte{24})
		iohelp.WriteUint32(w, uint32(len(*bbp.IpAddress)))
		w.Write([]byte(*bbp.IpAddress))
	}
	if bbp.Geo != nil {
		w.Write([]byte{25})
		err = (*bbp.Geo).EncodeBebop(w)
		if err != nil{
			return err
		}
	}
	if bbp.Tags != nil {
		w.Write([]byte{26})
		iohelp.WriteUint32(w, uint32(len(*bbp.Tags)))
		for _, elem := range *bbp.Tags {
			iohelp.WriteUint32(w, uint32(len(elem)))
			w.Write([]byte(elem))
		}
	}
	if bbp.Latency != nil {
		w.Write([]byte{27})
		err = (*bbp.Latency).EncodeBebop(w)
		if err != nil{
			return err
		}
	}
	if bbp.TrackPath != nil {
		w.Write([]byte{28})
		iohelp.WriteBool(w, *bbp.TrackPath)
	}
	w.Write([]byte{0})
	return w.Err
}

func (bbp *AnalyticsRecord) DecodeBebop(ior io.Reader) (err error) {
	r := iohelp.NewErrorReader(ior)
	bodyLen := iohelp.ReadUint32(r)
	r.Reader = &io.LimitedReader{R:r.Reader, N:int64(bodyLen)}
	for {
		switch iohelp.ReadByte(r) {
		case 1:
			bbp.Method = new(string)
			*bbp.Method = iohelp.ReadString(r)
		case 2:
			bbp.Host = new(string)
			*bbp.Host = iohelp.ReadString(r)
		case 3:
			bbp.Path = new(string)
			*bbp.Path = iohelp.ReadString(r)
		case 4:
			bbp.RawPath = new(string)
			*bbp.RawPath = iohelp.ReadString(r)
		case 5:
			bbp.ContentLength = new(int64)
			*bbp.ContentLength = iohelp.ReadInt64(r)
		case 6:
			bbp.Network = new(NetworkStats)
			(*bbp.Network), err = MakeNetworkStats(r)
			if err != nil{
				return err
			}
		case 7:
			bbp.Timestamp = new(time.Time)
			*bbp.Timestamp = iohelp.ReadDate(r)
		case 8:
			bbp.ExpireAt = new(time.Time)
			*bbp.ExpireAt = iohelp.ReadDate(r)
		case 9:
			bbp.Useragent = new(string)
			*bbp.Useragent = iohelp.ReadString(r)
		case 10:
			bbp.Day = new(int32)
			*bbp.Day = iohelp.ReadInt32(r)
		case 11:
			bbp.Year = new(int32)
			*bbp.Year = iohelp.ReadInt32(r)
		case 12:
			bbp.Month = new(int32)
			*bbp.Month = iohelp.ReadInt32(r)
		case 13:
			bbp.Hour = new(int32)
			*bbp.Hour = iohelp.ReadInt32(r)
		case 14:
			bbp.ResponseCode = new(int32)
			*bbp.ResponseCode = iohelp.ReadInt32(r)
		case 15:
			bbp.ApiKey = new(string)
			*bbp.ApiKey = iohelp.ReadString(r)
		case 16:
			bbp.ApiVersion = new(string)
			*bbp.ApiVersion = iohelp.ReadString(r)
		case 17:
			bbp.ApiName = new(string)
			*bbp.ApiName = iohelp.ReadString(r)
		case 18:
			bbp.ApiID = new(string)
			*bbp.ApiID = iohelp.ReadString(r)
		case 19:
			bbp.OrgID = new(string)
			*bbp.OrgID = iohelp.ReadString(r)
		case 20:
			bbp.OauthID = new(string)
			*bbp.OauthID = iohelp.ReadString(r)
		case 21:
			bbp.RequestTime = new(int64)
			*bbp.RequestTime = iohelp.ReadInt64(r)
		case 22:
			bbp.RawRequest = new(string)
			*bbp.RawRequest = iohelp.ReadString(r)
		case 23:
			bbp.RawResponse = new(string)
			*bbp.RawResponse = iohelp.ReadString(r)
		case 24:
			bbp.IpAddress = new(string)
			*bbp.IpAddress = iohelp.ReadString(r)
		case 25:
			bbp.Geo = new(GeoData)
			(*bbp.Geo), err = MakeGeoData(r)
			if err != nil{
				return err
			}
		case 26:
			bbp.Tags = new([]string)
			*bbp.Tags = make([]string, iohelp.ReadUint32(r))
			for i := range *bbp.Tags {
				(*bbp.Tags)[i] = iohelp.ReadString(r)
			}
		case 27:
			bbp.Latency = new(Latency)
			(*bbp.Latency), err = MakeLatency(r)
			if err != nil{
				return err
			}
		case 28:
			bbp.TrackPath = new(bool)
			*bbp.TrackPath = iohelp.ReadBool(r)
		default:
			io.ReadAll(r)
			return r.Err
		}
	}
}

func (bbp AnalyticsRecord) Size() int {
	bodyLen := 5
	if bbp.Method != nil {
		bodyLen += 1
		bodyLen += 4 + len(*bbp.Method)
	}
	if bbp.Host != nil {
		bodyLen += 1
		bodyLen += 4 + len(*bbp.Host)
	}
	if bbp.Path != nil {
		bodyLen += 1
		bodyLen += 4 + len(*bbp.Path)
	}
	if bbp.RawPath != nil {
		bodyLen += 1
		bodyLen += 4 + len(*bbp.RawPath)
	}
	if bbp.ContentLength != nil {
		bodyLen += 1
		bodyLen += 8
	}
	if bbp.Network != nil {
		bodyLen += 1
		bodyLen += (*bbp.Network).Size()
	}
	if bbp.Timestamp != nil {
		bodyLen += 1
		bodyLen += 8
	}
	if bbp.ExpireAt != nil {
		bodyLen += 1
		bodyLen += 8
	}
	if bbp.Useragent != nil {
		bodyLen += 1
		bodyLen += 4 + len(*bbp.Useragent)
	}
	if bbp.Day != nil {
		bodyLen += 1
		bodyLen += 4
	}
	if bbp.Year != nil {
		bodyLen += 1
		bodyLen += 4
	}
	if bbp.Month != nil {
		bodyLen += 1
		bodyLen += 4
	}
	if bbp.Hour != nil {
		bodyLen += 1
		bodyLen += 4
	}
	if bbp.ResponseCode != nil {
		bodyLen += 1
		bodyLen += 4
	}
	if bbp.ApiKey != nil {
		bodyLen += 1
		bodyLen += 4 + len(*bbp.ApiKey)
	}
	if bbp.ApiVersion != nil {
		bodyLen += 1
		bodyLen += 4 + len(*bbp.ApiVersion)
	}
	if bbp.ApiName != nil {
		bodyLen += 1
		bodyLen += 4 + len(*bbp.ApiName)
	}
	if bbp.ApiID != nil {
		bodyLen += 1
		bodyLen += 4 + len(*bbp.ApiID)
	}
	if bbp.OrgID != nil {
		bodyLen += 1
		bodyLen += 4 + len(*bbp.OrgID)
	}
	if bbp.OauthID != nil {
		bodyLen += 1
		bodyLen += 4 + len(*bbp.OauthID)
	}
	if bbp.RequestTime != nil {
		bodyLen += 1
		bodyLen += 8
	}
	if bbp.RawRequest != nil {
		bodyLen += 1
		bodyLen += 4 + len(*bbp.RawRequest)
	}
	if bbp.RawResponse != nil {
		bodyLen += 1
		bodyLen += 4 + len(*bbp.RawResponse)
	}
	if bbp.IpAddress != nil {
		bodyLen += 1
		bodyLen += 4 + len(*bbp.IpAddress)
	}
	if bbp.Geo != nil {
		bodyLen += 1
		bodyLen += (*bbp.Geo).Size()
	}
	if bbp.Tags != nil {
		bodyLen += 1
		bodyLen += 4
		for _, elem := range *bbp.Tags {
			bodyLen += 4 + len(elem)
		}
	}
	if bbp.Latency != nil {
		bodyLen += 1
		bodyLen += (*bbp.Latency).Size()
	}
	if bbp.TrackPath != nil {
		bodyLen += 1
		bodyLen += 1
	}
	return bodyLen
}

func (bbp AnalyticsRecord) MarshalBebop() []byte {
	buf := make([]byte, bbp.Size())
	bbp.MarshalBebopTo(buf)
	return buf
}

func MakeAnalyticsRecord(r iohelp.ErrorReader) (AnalyticsRecord, error) {
	v := AnalyticsRecord{}
	err := v.DecodeBebop(r)
	return v, err
}

func MakeAnalyticsRecordFromBytes(buf []byte) (AnalyticsRecord, error) {
	v := AnalyticsRecord{}
	err := v.UnmarshalBebop(buf)
	return v, err
}

var _ bebop.Record = &NetworkStats{}

type NetworkStats struct {
	OpenConnections *int64
	ClosedConnections *int64
	BytesIn *int64
	BytesOut *int64
}

func (bbp NetworkStats) MarshalBebopTo(buf []byte) int {
	at := 0
	iohelp.WriteUint32Bytes(buf[at:], uint32(bbp.Size()-4))
	at += 4
	if bbp.OpenConnections != nil {
		buf[at] = 1
		at++
		iohelp.WriteInt64Bytes(buf[at:], *bbp.OpenConnections)
		at += 8
	}
	if bbp.ClosedConnections != nil {
		buf[at] = 2
		at++
		iohelp.WriteInt64Bytes(buf[at:], *bbp.ClosedConnections)
		at += 8
	}
	if bbp.BytesIn != nil {
		buf[at] = 3
		at++
		iohelp.WriteInt64Bytes(buf[at:], *bbp.BytesIn)
		at += 8
	}
	if bbp.BytesOut != nil {
		buf[at] = 4
		at++
		iohelp.WriteInt64Bytes(buf[at:], *bbp.BytesOut)
		at += 8
	}
	return at
}

func (bbp *NetworkStats) UnmarshalBebop(buf []byte) (err error) {
	at := 0
	_ = iohelp.ReadUint32Bytes(buf[at:])
	buf = buf[4:]
	for {
		switch buf[at] {
		case 1:
			at += 1
			bbp.OpenConnections = new(int64)
			if len(buf[at:]) < 8 {
				 return io.ErrUnexpectedEOF
			}
			(*bbp.OpenConnections) = iohelp.ReadInt64Bytes(buf[at:])
			at += 8
		case 2:
			at += 1
			bbp.ClosedConnections = new(int64)
			if len(buf[at:]) < 8 {
				 return io.ErrUnexpectedEOF
			}
			(*bbp.ClosedConnections) = iohelp.ReadInt64Bytes(buf[at:])
			at += 8
		case 3:
			at += 1
			bbp.BytesIn = new(int64)
			if len(buf[at:]) < 8 {
				 return io.ErrUnexpectedEOF
			}
			(*bbp.BytesIn) = iohelp.ReadInt64Bytes(buf[at:])
			at += 8
		case 4:
			at += 1
			bbp.BytesOut = new(int64)
			if len(buf[at:]) < 8 {
				 return io.ErrUnexpectedEOF
			}
			(*bbp.BytesOut) = iohelp.ReadInt64Bytes(buf[at:])
			at += 8
		default:
			return nil
		}
	}
}

func (bbp NetworkStats) EncodeBebop(iow io.Writer) (err error) {
	w := iohelp.NewErrorWriter(iow)
	iohelp.WriteUint32(w, uint32(bbp.Size()-4))
	if bbp.OpenConnections != nil {
		w.Write([]byte{1})
		iohelp.WriteInt64(w, *bbp.OpenConnections)
	}
	if bbp.ClosedConnections != nil {
		w.Write([]byte{2})
		iohelp.WriteInt64(w, *bbp.ClosedConnections)
	}
	if bbp.BytesIn != nil {
		w.Write([]byte{3})
		iohelp.WriteInt64(w, *bbp.BytesIn)
	}
	if bbp.BytesOut != nil {
		w.Write([]byte{4})
		iohelp.WriteInt64(w, *bbp.BytesOut)
	}
	w.Write([]byte{0})
	return w.Err
}

func (bbp *NetworkStats) DecodeBebop(ior io.Reader) (err error) {
	r := iohelp.NewErrorReader(ior)
	bodyLen := iohelp.ReadUint32(r)
	r.Reader = &io.LimitedReader{R:r.Reader, N:int64(bodyLen)}
	for {
		switch iohelp.ReadByte(r) {
		case 1:
			bbp.OpenConnections = new(int64)
			*bbp.OpenConnections = iohelp.ReadInt64(r)
		case 2:
			bbp.ClosedConnections = new(int64)
			*bbp.ClosedConnections = iohelp.ReadInt64(r)
		case 3:
			bbp.BytesIn = new(int64)
			*bbp.BytesIn = iohelp.ReadInt64(r)
		case 4:
			bbp.BytesOut = new(int64)
			*bbp.BytesOut = iohelp.ReadInt64(r)
		default:
			io.ReadAll(r)
			return r.Err
		}
	}
}

func (bbp NetworkStats) Size() int {
	bodyLen := 5
	if bbp.OpenConnections != nil {
		bodyLen += 1
		bodyLen += 8
	}
	if bbp.ClosedConnections != nil {
		bodyLen += 1
		bodyLen += 8
	}
	if bbp.BytesIn != nil {
		bodyLen += 1
		bodyLen += 8
	}
	if bbp.BytesOut != nil {
		bodyLen += 1
		bodyLen += 8
	}
	return bodyLen
}

func (bbp NetworkStats) MarshalBebop() []byte {
	buf := make([]byte, bbp.Size())
	bbp.MarshalBebopTo(buf)
	return buf
}

func MakeNetworkStats(r iohelp.ErrorReader) (NetworkStats, error) {
	v := NetworkStats{}
	err := v.DecodeBebop(r)
	return v, err
}

func MakeNetworkStatsFromBytes(buf []byte) (NetworkStats, error) {
	v := NetworkStats{}
	err := v.UnmarshalBebop(buf)
	return v, err
}

var _ bebop.Record = &GeoData{}

type GeoData struct {
	Country *Country
	City *City
	Location *Location
}

func (bbp GeoData) MarshalBebopTo(buf []byte) int {
	at := 0
	iohelp.WriteUint32Bytes(buf[at:], uint32(bbp.Size()-4))
	at += 4
	if bbp.Country != nil {
		buf[at] = 1
		at++
		(*bbp.Country).MarshalBebopTo(buf[at:])
		at += (*bbp.Country).Size()
	}
	if bbp.City != nil {
		buf[at] = 2
		at++
		(*bbp.City).MarshalBebopTo(buf[at:])
		at += (*bbp.City).Size()
	}
	if bbp.Location != nil {
		buf[at] = 3
		at++
		(*bbp.Location).MarshalBebopTo(buf[at:])
		at += (*bbp.Location).Size()
	}
	return at
}

func (bbp *GeoData) UnmarshalBebop(buf []byte) (err error) {
	at := 0
	_ = iohelp.ReadUint32Bytes(buf[at:])
	buf = buf[4:]
	for {
		switch buf[at] {
		case 1:
			at += 1
			bbp.Country = new(Country)
			(*bbp.Country), err = MakeCountryFromBytes(buf[at:])
			if err != nil{
				return err
			}
			at += ((*bbp.Country)).Size()
		case 2:
			at += 1
			bbp.City = new(City)
			(*bbp.City), err = MakeCityFromBytes(buf[at:])
			if err != nil{
				return err
			}
			at += ((*bbp.City)).Size()
		case 3:
			at += 1
			bbp.Location = new(Location)
			(*bbp.Location), err = MakeLocationFromBytes(buf[at:])
			if err != nil{
				return err
			}
			at += ((*bbp.Location)).Size()
		default:
			return nil
		}
	}
}

func (bbp GeoData) EncodeBebop(iow io.Writer) (err error) {
	w := iohelp.NewErrorWriter(iow)
	iohelp.WriteUint32(w, uint32(bbp.Size()-4))
	if bbp.Country != nil {
		w.Write([]byte{1})
		err = (*bbp.Country).EncodeBebop(w)
		if err != nil{
			return err
		}
	}
	if bbp.City != nil {
		w.Write([]byte{2})
		err = (*bbp.City).EncodeBebop(w)
		if err != nil{
			return err
		}
	}
	if bbp.Location != nil {
		w.Write([]byte{3})
		err = (*bbp.Location).EncodeBebop(w)
		if err != nil{
			return err
		}
	}
	w.Write([]byte{0})
	return w.Err
}

func (bbp *GeoData) DecodeBebop(ior io.Reader) (err error) {
	r := iohelp.NewErrorReader(ior)
	bodyLen := iohelp.ReadUint32(r)
	r.Reader = &io.LimitedReader{R:r.Reader, N:int64(bodyLen)}
	for {
		switch iohelp.ReadByte(r) {
		case 1:
			bbp.Country = new(Country)
			(*bbp.Country), err = MakeCountry(r)
			if err != nil{
				return err
			}
		case 2:
			bbp.City = new(City)
			(*bbp.City), err = MakeCity(r)
			if err != nil{
				return err
			}
		case 3:
			bbp.Location = new(Location)
			(*bbp.Location), err = MakeLocation(r)
			if err != nil{
				return err
			}
		default:
			io.ReadAll(r)
			return r.Err
		}
	}
}

func (bbp GeoData) Size() int {
	bodyLen := 5
	if bbp.Country != nil {
		bodyLen += 1
		bodyLen += (*bbp.Country).Size()
	}
	if bbp.City != nil {
		bodyLen += 1
		bodyLen += (*bbp.City).Size()
	}
	if bbp.Location != nil {
		bodyLen += 1
		bodyLen += (*bbp.Location).Size()
	}
	return bodyLen
}

func (bbp GeoData) MarshalBebop() []byte {
	buf := make([]byte, bbp.Size())
	bbp.MarshalBebopTo(buf)
	return buf
}

func MakeGeoData(r iohelp.ErrorReader) (GeoData, error) {
	v := GeoData{}
	err := v.DecodeBebop(r)
	return v, err
}

func MakeGeoDataFromBytes(buf []byte) (GeoData, error) {
	v := GeoData{}
	err := v.UnmarshalBebop(buf)
	return v, err
}

var _ bebop.Record = &Country{}

type Country struct {
	ISOCode *string
}

func (bbp Country) MarshalBebopTo(buf []byte) int {
	at := 0
	iohelp.WriteUint32Bytes(buf[at:], uint32(bbp.Size()-4))
	at += 4
	if bbp.ISOCode != nil {
		buf[at] = 1
		at++
		iohelp.WriteUint32Bytes(buf[at:], uint32(len(*bbp.ISOCode)))
		copy(buf[at+4:at+4+len(*bbp.ISOCode)], []byte(*bbp.ISOCode))
		at += 4 + len(*bbp.ISOCode)
	}
	return at
}

func (bbp *Country) UnmarshalBebop(buf []byte) (err error) {
	at := 0
	_ = iohelp.ReadUint32Bytes(buf[at:])
	buf = buf[4:]
	for {
		switch buf[at] {
		case 1:
			at += 1
			bbp.ISOCode = new(string)
			(*bbp.ISOCode), err = iohelp.ReadStringBytes(buf[at:])
			if err != nil{
				return err
			}
			at += 4 + len((*bbp.ISOCode))
		default:
			return nil
		}
	}
}

func (bbp Country) EncodeBebop(iow io.Writer) (err error) {
	w := iohelp.NewErrorWriter(iow)
	iohelp.WriteUint32(w, uint32(bbp.Size()-4))
	if bbp.ISOCode != nil {
		w.Write([]byte{1})
		iohelp.WriteUint32(w, uint32(len(*bbp.ISOCode)))
		w.Write([]byte(*bbp.ISOCode))
	}
	w.Write([]byte{0})
	return w.Err
}

func (bbp *Country) DecodeBebop(ior io.Reader) (err error) {
	r := iohelp.NewErrorReader(ior)
	bodyLen := iohelp.ReadUint32(r)
	r.Reader = &io.LimitedReader{R:r.Reader, N:int64(bodyLen)}
	for {
		switch iohelp.ReadByte(r) {
		case 1:
			bbp.ISOCode = new(string)
			*bbp.ISOCode = iohelp.ReadString(r)
		default:
			io.ReadAll(r)
			return r.Err
		}
	}
}

func (bbp Country) Size() int {
	bodyLen := 5
	if bbp.ISOCode != nil {
		bodyLen += 1
		bodyLen += 4 + len(*bbp.ISOCode)
	}
	return bodyLen
}

func (bbp Country) MarshalBebop() []byte {
	buf := make([]byte, bbp.Size())
	bbp.MarshalBebopTo(buf)
	return buf
}

func MakeCountry(r iohelp.ErrorReader) (Country, error) {
	v := Country{}
	err := v.DecodeBebop(r)
	return v, err
}

func MakeCountryFromBytes(buf []byte) (Country, error) {
	v := Country{}
	err := v.UnmarshalBebop(buf)
	return v, err
}

var _ bebop.Record = &City{}

type City struct {
	GeoNameID *uint32
}

func (bbp City) MarshalBebopTo(buf []byte) int {
	at := 0
	iohelp.WriteUint32Bytes(buf[at:], uint32(bbp.Size()-4))
	at += 4
	if bbp.GeoNameID != nil {
		buf[at] = 1
		at++
		iohelp.WriteUint32Bytes(buf[at:], *bbp.GeoNameID)
		at += 4
	}
	return at
}

func (bbp *City) UnmarshalBebop(buf []byte) (err error) {
	at := 0
	_ = iohelp.ReadUint32Bytes(buf[at:])
	buf = buf[4:]
	for {
		switch buf[at] {
		case 1:
			at += 1
			bbp.GeoNameID = new(uint32)
			if len(buf[at:]) < 4 {
				 return io.ErrUnexpectedEOF
			}
			(*bbp.GeoNameID) = iohelp.ReadUint32Bytes(buf[at:])
			at += 4
		default:
			return nil
		}
	}
}

func (bbp City) EncodeBebop(iow io.Writer) (err error) {
	w := iohelp.NewErrorWriter(iow)
	iohelp.WriteUint32(w, uint32(bbp.Size()-4))
	if bbp.GeoNameID != nil {
		w.Write([]byte{1})
		iohelp.WriteUint32(w, *bbp.GeoNameID)
	}
	w.Write([]byte{0})
	return w.Err
}

func (bbp *City) DecodeBebop(ior io.Reader) (err error) {
	r := iohelp.NewErrorReader(ior)
	bodyLen := iohelp.ReadUint32(r)
	r.Reader = &io.LimitedReader{R:r.Reader, N:int64(bodyLen)}
	for {
		switch iohelp.ReadByte(r) {
		case 1:
			bbp.GeoNameID = new(uint32)
			*bbp.GeoNameID = iohelp.ReadUint32(r)
		default:
			io.ReadAll(r)
			return r.Err
		}
	}
}

func (bbp City) Size() int {
	bodyLen := 5
	if bbp.GeoNameID != nil {
		bodyLen += 1
		bodyLen += 4
	}
	return bodyLen
}

func (bbp City) MarshalBebop() []byte {
	buf := make([]byte, bbp.Size())
	bbp.MarshalBebopTo(buf)
	return buf
}

func MakeCity(r iohelp.ErrorReader) (City, error) {
	v := City{}
	err := v.DecodeBebop(r)
	return v, err
}

func MakeCityFromBytes(buf []byte) (City, error) {
	v := City{}
	err := v.UnmarshalBebop(buf)
	return v, err
}

var _ bebop.Record = &Location{}

type Location struct {
	Latitude *float64
	Longitude *float64
	Timezone *string
}

func (bbp Location) MarshalBebopTo(buf []byte) int {
	at := 0
	iohelp.WriteUint32Bytes(buf[at:], uint32(bbp.Size()-4))
	at += 4
	if bbp.Latitude != nil {
		buf[at] = 1
		at++
		iohelp.WriteFloat64Bytes(buf[at:], *bbp.Latitude)
		at += 8
	}
	if bbp.Longitude != nil {
		buf[at] = 2
		at++
		iohelp.WriteFloat64Bytes(buf[at:], *bbp.Longitude)
		at += 8
	}
	if bbp.Timezone != nil {
		buf[at] = 3
		at++
		iohelp.WriteUint32Bytes(buf[at:], uint32(len(*bbp.Timezone)))
		copy(buf[at+4:at+4+len(*bbp.Timezone)], []byte(*bbp.Timezone))
		at += 4 + len(*bbp.Timezone)
	}
	return at
}

func (bbp *Location) UnmarshalBebop(buf []byte) (err error) {
	at := 0
	_ = iohelp.ReadUint32Bytes(buf[at:])
	buf = buf[4:]
	for {
		switch buf[at] {
		case 1:
			at += 1
			bbp.Latitude = new(float64)
			if len(buf[at:]) < 8 {
				 return io.ErrUnexpectedEOF
			}
			(*bbp.Latitude) = iohelp.ReadFloat64Bytes(buf[at:])
			at += 8
		case 2:
			at += 1
			bbp.Longitude = new(float64)
			if len(buf[at:]) < 8 {
				 return io.ErrUnexpectedEOF
			}
			(*bbp.Longitude) = iohelp.ReadFloat64Bytes(buf[at:])
			at += 8
		case 3:
			at += 1
			bbp.Timezone = new(string)
			(*bbp.Timezone), err = iohelp.ReadStringBytes(buf[at:])
			if err != nil{
				return err
			}
			at += 4 + len((*bbp.Timezone))
		default:
			return nil
		}
	}
}

func (bbp Location) EncodeBebop(iow io.Writer) (err error) {
	w := iohelp.NewErrorWriter(iow)
	iohelp.WriteUint32(w, uint32(bbp.Size()-4))
	if bbp.Latitude != nil {
		w.Write([]byte{1})
		iohelp.WriteFloat64(w, *bbp.Latitude)
	}
	if bbp.Longitude != nil {
		w.Write([]byte{2})
		iohelp.WriteFloat64(w, *bbp.Longitude)
	}
	if bbp.Timezone != nil {
		w.Write([]byte{3})
		iohelp.WriteUint32(w, uint32(len(*bbp.Timezone)))
		w.Write([]byte(*bbp.Timezone))
	}
	w.Write([]byte{0})
	return w.Err
}

func (bbp *Location) DecodeBebop(ior io.Reader) (err error) {
	r := iohelp.NewErrorReader(ior)
	bodyLen := iohelp.ReadUint32(r)
	r.Reader = &io.LimitedReader{R:r.Reader, N:int64(bodyLen)}
	for {
		switch iohelp.ReadByte(r) {
		case 1:
			bbp.Latitude = new(float64)
			*bbp.Latitude = iohelp.ReadFloat64(r)
		case 2:
			bbp.Longitude = new(float64)
			*bbp.Longitude = iohelp.ReadFloat64(r)
		case 3:
			bbp.Timezone = new(string)
			*bbp.Timezone = iohelp.ReadString(r)
		default:
			io.ReadAll(r)
			return r.Err
		}
	}
}

func (bbp Location) Size() int {
	bodyLen := 5
	if bbp.Latitude != nil {
		bodyLen += 1
		bodyLen += 8
	}
	if bbp.Longitude != nil {
		bodyLen += 1
		bodyLen += 8
	}
	if bbp.Timezone != nil {
		bodyLen += 1
		bodyLen += 4 + len(*bbp.Timezone)
	}
	return bodyLen
}

func (bbp Location) MarshalBebop() []byte {
	buf := make([]byte, bbp.Size())
	bbp.MarshalBebopTo(buf)
	return buf
}

func MakeLocation(r iohelp.ErrorReader) (Location, error) {
	v := Location{}
	err := v.DecodeBebop(r)
	return v, err
}

func MakeLocationFromBytes(buf []byte) (Location, error) {
	v := Location{}
	err := v.UnmarshalBebop(buf)
	return v, err
}

var _ bebop.Record = &Latency{}

type Latency struct {
	Total *int64
	Upstream *int64
}

func (bbp Latency) MarshalBebopTo(buf []byte) int {
	at := 0
	iohelp.WriteUint32Bytes(buf[at:], uint32(bbp.Size()-4))
	at += 4
	if bbp.Total != nil {
		buf[at] = 1
		at++
		iohelp.WriteInt64Bytes(buf[at:], *bbp.Total)
		at += 8
	}
	if bbp.Upstream != nil {
		buf[at] = 2
		at++
		iohelp.WriteInt64Bytes(buf[at:], *bbp.Upstream)
		at += 8
	}
	return at
}

func (bbp *Latency) UnmarshalBebop(buf []byte) (err error) {
	at := 0
	_ = iohelp.ReadUint32Bytes(buf[at:])
	buf = buf[4:]
	for {
		switch buf[at] {
		case 1:
			at += 1
			bbp.Total = new(int64)
			if len(buf[at:]) < 8 {
				 return io.ErrUnexpectedEOF
			}
			(*bbp.Total) = iohelp.ReadInt64Bytes(buf[at:])
			at += 8
		case 2:
			at += 1
			bbp.Upstream = new(int64)
			if len(buf[at:]) < 8 {
				 return io.ErrUnexpectedEOF
			}
			(*bbp.Upstream) = iohelp.ReadInt64Bytes(buf[at:])
			at += 8
		default:
			return nil
		}
	}
}

func (bbp Latency) EncodeBebop(iow io.Writer) (err error) {
	w := iohelp.NewErrorWriter(iow)
	iohelp.WriteUint32(w, uint32(bbp.Size()-4))
	if bbp.Total != nil {
		w.Write([]byte{1})
		iohelp.WriteInt64(w, *bbp.Total)
	}
	if bbp.Upstream != nil {
		w.Write([]byte{2})
		iohelp.WriteInt64(w, *bbp.Upstream)
	}
	w.Write([]byte{0})
	return w.Err
}

func (bbp *Latency) DecodeBebop(ior io.Reader) (err error) {
	r := iohelp.NewErrorReader(ior)
	bodyLen := iohelp.ReadUint32(r)
	r.Reader = &io.LimitedReader{R:r.Reader, N:int64(bodyLen)}
	for {
		switch iohelp.ReadByte(r) {
		case 1:
			bbp.Total = new(int64)
			*bbp.Total = iohelp.ReadInt64(r)
		case 2:
			bbp.Upstream = new(int64)
			*bbp.Upstream = iohelp.ReadInt64(r)
		default:
			io.ReadAll(r)
			return r.Err
		}
	}
}

func (bbp Latency) Size() int {
	bodyLen := 5
	if bbp.Total != nil {
		bodyLen += 1
		bodyLen += 8
	}
	if bbp.Upstream != nil {
		bodyLen += 1
		bodyLen += 8
	}
	return bodyLen
}

func (bbp Latency) MarshalBebop() []byte {
	buf := make([]byte, bbp.Size())
	bbp.MarshalBebopTo(buf)
	return buf
}

func MakeLatency(r iohelp.ErrorReader) (Latency, error) {
	v := Latency{}
	err := v.DecodeBebop(r)
	return v, err
}

func MakeLatencyFromBytes(buf []byte) (Latency, error) {
	v := Latency{}
	err := v.UnmarshalBebop(buf)
	return v, err
}

