// automatically generated by the FlatBuffers compiler, do not modify

package schema

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type PostRecord struct {
	_tab flatbuffers.Table
}

func GetRootAsPostRecord(buf []byte, offset flatbuffers.UOffsetT) *PostRecord {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &PostRecord{}
	x.Init(buf, n+offset)
	return x
}

func (rcv *PostRecord) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *PostRecord) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *PostRecord) Typ() int8 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetInt8(o + rcv._tab.Pos)
	}
	return 2
}

func (rcv *PostRecord) MutateTyp(n int8) bool {
	return rcv._tab.MutateInt8Slot(4, n)
}

func (rcv *PostRecord) Id(obj *Id) *Id {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		x := rcv._tab.Indirect(o + rcv._tab.Pos)
		if obj == nil {
			obj = new(Id)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

func (rcv *PostRecord) OwnerId(obj *Id) *Id {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		x := rcv._tab.Indirect(o + rcv._tab.Pos)
		if obj == nil {
			obj = new(Id)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

func (rcv *PostRecord) Updated() uint64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		return rcv._tab.GetUint64(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *PostRecord) MutateUpdated(n uint64) bool {
	return rcv._tab.MutateUint64Slot(10, n)
}

func (rcv *PostRecord) Expiry() uint64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(12))
	if o != 0 {
		return rcv._tab.GetUint64(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *PostRecord) MutateExpiry(n uint64) bool {
	return rcv._tab.MutateUint64Slot(12, n)
}

func (rcv *PostRecord) Reserved() uint64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(14))
	if o != 0 {
		return rcv._tab.GetUint64(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *PostRecord) MutateReserved(n uint64) bool {
	return rcv._tab.MutateUint64Slot(14, n)
}

func (rcv *PostRecord) Cost(obj *Cost) *Cost {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(16))
	if o != 0 {
		x := rcv._tab.Indirect(o + rcv._tab.Pos)
		if obj == nil {
			obj = new(Cost)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

func (rcv *PostRecord) TransactionId(obj *Id) *Id {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(18))
	if o != 0 {
		x := rcv._tab.Indirect(o + rcv._tab.Pos)
		if obj == nil {
			obj = new(Id)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

func PostRecordStart(builder *flatbuffers.Builder) {
	builder.StartObject(8)
}
func PostRecordAddTyp(builder *flatbuffers.Builder, typ int8) {
	builder.PrependInt8Slot(0, typ, 2)
}
func PostRecordAddId(builder *flatbuffers.Builder, id flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(id), 0)
}
func PostRecordAddOwnerId(builder *flatbuffers.Builder, ownerId flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(2, flatbuffers.UOffsetT(ownerId), 0)
}
func PostRecordAddUpdated(builder *flatbuffers.Builder, updated uint64) {
	builder.PrependUint64Slot(3, updated, 0)
}
func PostRecordAddExpiry(builder *flatbuffers.Builder, expiry uint64) {
	builder.PrependUint64Slot(4, expiry, 0)
}
func PostRecordAddReserved(builder *flatbuffers.Builder, reserved uint64) {
	builder.PrependUint64Slot(5, reserved, 0)
}
func PostRecordAddCost(builder *flatbuffers.Builder, cost flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(6, flatbuffers.UOffsetT(cost), 0)
}
func PostRecordAddTransactionId(builder *flatbuffers.Builder, transactionId flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(7, flatbuffers.UOffsetT(transactionId), 0)
}
func PostRecordEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
