// automatically generated by the FlatBuffers compiler, do not modify

package schema

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type PatchRequest struct {
	_tab flatbuffers.Table
}

func GetRootAsPatchRequest(buf []byte, offset flatbuffers.UOffsetT) *PatchRequest {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &PatchRequest{}
	x.Init(buf, n+offset)
	return x
}

func (rcv *PatchRequest) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *PatchRequest) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *PatchRequest) Operations(obj *Operation, j int) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		x := rcv._tab.Vector(o)
		x += flatbuffers.UOffsetT(j) * 4
		x = rcv._tab.Indirect(x)
		obj.Init(rcv._tab.Bytes, x)
		return true
	}
	return false
}

func (rcv *PatchRequest) OperationsLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *PatchRequest) Score() float64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.GetFloat64(o + rcv._tab.Pos)
	}
	return 0.0
}

func (rcv *PatchRequest) MutateScore(n float64) bool {
	return rcv._tab.MutateFloat64Slot(6, n)
}

func (rcv *PatchRequest) MaxSize() uint64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.GetUint64(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *PatchRequest) MutateMaxSize(n uint64) bool {
	return rcv._tab.MutateUint64Slot(8, n)
}

func (rcv *PatchRequest) Expiry() uint64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		return rcv._tab.GetUint64(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *PatchRequest) MutateExpiry(n uint64) bool {
	return rcv._tab.MutateUint64Slot(10, n)
}

func PatchRequestStart(builder *flatbuffers.Builder) {
	builder.StartObject(4)
}
func PatchRequestAddOperations(builder *flatbuffers.Builder, operations flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(operations), 0)
}
func PatchRequestStartOperationsVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func PatchRequestAddScore(builder *flatbuffers.Builder, score float64) {
	builder.PrependFloat64Slot(1, score, 0.0)
}
func PatchRequestAddMaxSize(builder *flatbuffers.Builder, maxSize uint64) {
	builder.PrependUint64Slot(2, maxSize, 0)
}
func PatchRequestAddExpiry(builder *flatbuffers.Builder, expiry uint64) {
	builder.PrependUint64Slot(3, expiry, 0)
}
func PatchRequestEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
