// automatically generated by the FlatBuffers compiler, do not modify

package schema

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type Dates struct {
	_tab flatbuffers.Table
}

func GetRootAsDates(buf []byte, offset flatbuffers.UOffsetT) *Dates {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &Dates{}
	x.Init(buf, n+offset)
	return x
}

func (rcv *Dates) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *Dates) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *Dates) Start() uint64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetUint64(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Dates) MutateStart(n uint64) bool {
	return rcv._tab.MutateUint64Slot(4, n)
}

func (rcv *Dates) End() uint64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.GetUint64(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Dates) MutateEnd(n uint64) bool {
	return rcv._tab.MutateUint64Slot(6, n)
}

func DatesStart(builder *flatbuffers.Builder) {
	builder.StartObject(2)
}
func DatesAddStart(builder *flatbuffers.Builder, start uint64) {
	builder.PrependUint64Slot(0, start, 0)
}
func DatesAddEnd(builder *flatbuffers.Builder, end uint64) {
	builder.PrependUint64Slot(1, end, 0)
}
func DatesEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
