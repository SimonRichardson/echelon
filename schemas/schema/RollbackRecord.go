// automatically generated by the FlatBuffers compiler, do not modify

package schema

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type RollbackRecord struct {
	_tab flatbuffers.Table
}

func GetRootAsRollbackRecord(buf []byte, offset flatbuffers.UOffsetT) *RollbackRecord {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &RollbackRecord{}
	x.Init(buf, n+offset)
	return x
}

func (rcv *RollbackRecord) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *RollbackRecord) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *RollbackRecord) Typ() int8 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetInt8(o + rcv._tab.Pos)
	}
	return 5
}

func (rcv *RollbackRecord) MutateTyp(n int8) bool {
	return rcv._tab.MutateInt8Slot(4, n)
}

func (rcv *RollbackRecord) Id(obj *Id) *Id {
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

func (rcv *RollbackRecord) OwnerId(obj *Id) *Id {
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

func (rcv *RollbackRecord) Updated() uint64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		return rcv._tab.GetUint64(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *RollbackRecord) MutateUpdated(n uint64) bool {
	return rcv._tab.MutateUint64Slot(10, n)
}

func (rcv *RollbackRecord) TransactionId(obj *Id) *Id {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(12))
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

func RollbackRecordStart(builder *flatbuffers.Builder) {
	builder.StartObject(5)
}
func RollbackRecordAddTyp(builder *flatbuffers.Builder, typ int8) {
	builder.PrependInt8Slot(0, typ, 5)
}
func RollbackRecordAddId(builder *flatbuffers.Builder, id flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(id), 0)
}
func RollbackRecordAddOwnerId(builder *flatbuffers.Builder, ownerId flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(2, flatbuffers.UOffsetT(ownerId), 0)
}
func RollbackRecordAddUpdated(builder *flatbuffers.Builder, updated uint64) {
	builder.PrependUint64Slot(3, updated, 0)
}
func RollbackRecordAddTransactionId(builder *flatbuffers.Builder, transactionId flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(4, flatbuffers.UOffsetT(transactionId), 0)
}
func RollbackRecordEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
