// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package analyticsflatbuffers

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type Country struct {
	_tab flatbuffers.Table
}

func GetRootAsCountry(buf []byte, offset flatbuffers.UOffsetT) *Country {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &Country{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsCountry(buf []byte, offset flatbuffers.UOffsetT) *Country {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &Country{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *Country) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *Country) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *Country) IsoCode() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func CountryStart(builder *flatbuffers.Builder) {
	builder.StartObject(1)
}
func CountryAddIsoCode(builder *flatbuffers.Builder, isoCode flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(isoCode), 0)
}
func CountryEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}