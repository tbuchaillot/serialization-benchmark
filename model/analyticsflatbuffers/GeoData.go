// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package analyticsflatbuffers

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type GeoData struct {
	_tab flatbuffers.Table
}

func GetRootAsGeoData(buf []byte, offset flatbuffers.UOffsetT) *GeoData {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &GeoData{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsGeoData(buf []byte, offset flatbuffers.UOffsetT) *GeoData {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &GeoData{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *GeoData) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *GeoData) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *GeoData) Country(obj *Country) *Country {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		x := rcv._tab.Indirect(o + rcv._tab.Pos)
		if obj == nil {
			obj = new(Country)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

func (rcv *GeoData) City(obj *City) *City {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		x := rcv._tab.Indirect(o + rcv._tab.Pos)
		if obj == nil {
			obj = new(City)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

func (rcv *GeoData) Location(obj *Location) *Location {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		x := rcv._tab.Indirect(o + rcv._tab.Pos)
		if obj == nil {
			obj = new(Location)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

func GeoDataStart(builder *flatbuffers.Builder) {
	builder.StartObject(3)
}
func GeoDataAddCountry(builder *flatbuffers.Builder, country flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(country), 0)
}
func GeoDataAddCity(builder *flatbuffers.Builder, city flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(city), 0)
}
func GeoDataAddLocation(builder *flatbuffers.Builder, location flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(2, flatbuffers.UOffsetT(location), 0)
}
func GeoDataEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
