// automatically generated by the FlatBuffers compiler, do not modify

package schema

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type OKKeyFieldScoreTxnValue struct {
	_tab flatbuffers.Table
}

func GetRootAsOKKeyFieldScoreTxnValue(buf []byte, offset flatbuffers.UOffsetT) *OKKeyFieldScoreTxnValue {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &OKKeyFieldScoreTxnValue{}
	x.Init(buf, n+offset)
	return x
}

func (rcv *OKKeyFieldScoreTxnValue) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *OKKeyFieldScoreTxnValue) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *OKKeyFieldScoreTxnValue) Duration() int64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetInt64(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *OKKeyFieldScoreTxnValue) MutateDuration(n int64) bool {
	return rcv._tab.MutateInt64Slot(4, n)
}

func (rcv *OKKeyFieldScoreTxnValue) Records(obj *KeyFieldScoreTxnValue) *KeyFieldScoreTxnValue {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		x := rcv._tab.Indirect(o + rcv._tab.Pos)
		if obj == nil {
			obj = new(KeyFieldScoreTxnValue)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

func OKKeyFieldScoreTxnValueStart(builder *flatbuffers.Builder) {
	builder.StartObject(2)
}
func OKKeyFieldScoreTxnValueAddDuration(builder *flatbuffers.Builder, duration int64) {
	builder.PrependInt64Slot(0, duration, 0)
}
func OKKeyFieldScoreTxnValueAddRecords(builder *flatbuffers.Builder, records flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(records), 0)
}
func OKKeyFieldScoreTxnValueEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}