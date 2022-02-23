// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package analyticsflatbuffers

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type NetworkStats struct {
	_tab flatbuffers.Table
}

func GetRootAsNetworkStats(buf []byte, offset flatbuffers.UOffsetT) *NetworkStats {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &NetworkStats{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsNetworkStats(buf []byte, offset flatbuffers.UOffsetT) *NetworkStats {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &NetworkStats{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *NetworkStats) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *NetworkStats) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *NetworkStats) OpenConnections() int64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetInt64(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *NetworkStats) MutateOpenConnections(n int64) bool {
	return rcv._tab.MutateInt64Slot(4, n)
}

func (rcv *NetworkStats) ClosedConnections() int64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.GetInt64(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *NetworkStats) MutateClosedConnections(n int64) bool {
	return rcv._tab.MutateInt64Slot(6, n)
}

func (rcv *NetworkStats) BytesIn() int64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.GetInt64(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *NetworkStats) MutateBytesIn(n int64) bool {
	return rcv._tab.MutateInt64Slot(8, n)
}

func (rcv *NetworkStats) BytesOut() int64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		return rcv._tab.GetInt64(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *NetworkStats) MutateBytesOut(n int64) bool {
	return rcv._tab.MutateInt64Slot(10, n)
}

func NetworkStatsStart(builder *flatbuffers.Builder) {
	builder.StartObject(4)
}
func NetworkStatsAddOpenConnections(builder *flatbuffers.Builder, openConnections int64) {
	builder.PrependInt64Slot(0, openConnections, 0)
}
func NetworkStatsAddClosedConnections(builder *flatbuffers.Builder, closedConnections int64) {
	builder.PrependInt64Slot(1, closedConnections, 0)
}
func NetworkStatsAddBytesIn(builder *flatbuffers.Builder, bytesIn int64) {
	builder.PrependInt64Slot(2, bytesIn, 0)
}
func NetworkStatsAddBytesOut(builder *flatbuffers.Builder, bytesOut int64) {
	builder.PrependInt64Slot(3, bytesOut, 0)
}
func NetworkStatsEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
