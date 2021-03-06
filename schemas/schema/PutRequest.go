// automatically generated by the FlatBuffers compiler, do not modify

package schema

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type PutRequest struct {
	_tab flatbuffers.Table
}

func GetRootAsPutRequest(buf []byte, offset flatbuffers.UOffsetT) *PutRequest {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &PutRequest{}
	x.Init(buf, n+offset)
	return x
}

func (rcv *PutRequest) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *PutRequest) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *PutRequest) Score() float64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetFloat64(o + rcv._tab.Pos)
	}
	return 0.0
}

func (rcv *PutRequest) MutateScore(n float64) bool {
	return rcv._tab.MutateFloat64Slot(4, n)
}

func (rcv *PutRequest) MaxSize() uint64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.GetUint64(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *PutRequest) MutateMaxSize(n uint64) bool {
	return rcv._tab.MutateUint64Slot(6, n)
}

func (rcv *PutRequest) Expiry() uint64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.GetUint64(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *PutRequest) MutateExpiry(n uint64) bool {
	return rcv._tab.MutateUint64Slot(8, n)
}

func (rcv *PutRequest) Records(obj *PutRecord, j int) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		x := rcv._tab.Vector(o)
		x += flatbuffers.UOffsetT(j) * 4
		x = rcv._tab.Indirect(x)
		obj.Init(rcv._tab.Bytes, x)
		return true
	}
	return false
}

func (rcv *PutRequest) RecordsLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func PutRequestStart(builder *flatbuffers.Builder) {
	builder.StartObject(4)
}
func PutRequestAddScore(builder *flatbuffers.Builder, score float64) {
	builder.PrependFloat64Slot(0, score, 0.0)
}
func PutRequestAddMaxSize(builder *flatbuffers.Builder, maxSize uint64) {
	builder.PrependUint64Slot(1, maxSize, 0)
}
func PutRequestAddExpiry(builder *flatbuffers.Builder, expiry uint64) {
	builder.PrependUint64Slot(2, expiry, 0)
}
func PutRequestAddRecords(builder *flatbuffers.Builder, records flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(3, flatbuffers.UOffsetT(records), 0)
}
func PutRequestStartRecordsVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func PutRequestEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
