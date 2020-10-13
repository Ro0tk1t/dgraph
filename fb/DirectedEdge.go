// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fb

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type DirectedEdge struct {
	_tab flatbuffers.Table
}

func GetRootAsDirectedEdge(buf []byte, offset flatbuffers.UOffsetT) *DirectedEdge {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &DirectedEdge{}
	x.Init(buf, n+offset)
	return x
}

func (rcv *DirectedEdge) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *DirectedEdge) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *DirectedEdge) Entity() uint64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetUint64(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *DirectedEdge) MutateEntity(n uint64) bool {
	return rcv._tab.MutateUint64Slot(4, n)
}

func (rcv *DirectedEdge) Attr() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *DirectedEdge) Value(j int) byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.GetByte(a + flatbuffers.UOffsetT(j*1))
	}
	return 0
}

func (rcv *DirectedEdge) ValueLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *DirectedEdge) ValueBytes() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *DirectedEdge) MutateValue(j int, n byte) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.MutateByte(a+flatbuffers.UOffsetT(j*1), n)
	}
	return false
}

func (rcv *DirectedEdge) ValueType() PostingValueType {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		return PostingValueType(rcv._tab.GetInt32(o + rcv._tab.Pos))
	}
	return 0
}

func (rcv *DirectedEdge) MutateValueType(n PostingValueType) bool {
	return rcv._tab.MutateInt32Slot(10, int32(n))
}

func (rcv *DirectedEdge) ValueId() uint64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(12))
	if o != 0 {
		return rcv._tab.GetUint64(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *DirectedEdge) MutateValueId(n uint64) bool {
	return rcv._tab.MutateUint64Slot(12, n)
}

func (rcv *DirectedEdge) Label() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(14))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *DirectedEdge) Lang() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(16))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *DirectedEdge) Op() DirectedEdgeOp {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(18))
	if o != 0 {
		return DirectedEdgeOp(rcv._tab.GetInt32(o + rcv._tab.Pos))
	}
	return 0
}

func (rcv *DirectedEdge) MutateOp(n DirectedEdgeOp) bool {
	return rcv._tab.MutateInt32Slot(18, int32(n))
}

func (rcv *DirectedEdge) Facets(obj *Facet, j int) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(20))
	if o != 0 {
		x := rcv._tab.Vector(o)
		x += flatbuffers.UOffsetT(j) * 4
		x = rcv._tab.Indirect(x)
		obj.Init(rcv._tab.Bytes, x)
		return true
	}
	return false
}

func (rcv *DirectedEdge) FacetsLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(20))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *DirectedEdge) AllowedPreds(j int) []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(22))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.ByteVector(a + flatbuffers.UOffsetT(j*4))
	}
	return nil
}

func (rcv *DirectedEdge) AllowedPredsLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(22))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func DirectedEdgeStart(builder *flatbuffers.Builder) {
	builder.StartObject(10)
}
func DirectedEdgeAddEntity(builder *flatbuffers.Builder, entity uint64) {
	builder.PrependUint64Slot(0, entity, 0)
}
func DirectedEdgeAddAttr(builder *flatbuffers.Builder, attr flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(attr), 0)
}
func DirectedEdgeAddValue(builder *flatbuffers.Builder, value flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(2, flatbuffers.UOffsetT(value), 0)
}
func DirectedEdgeStartValueVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(1, numElems, 1)
}
func DirectedEdgeAddValueType(builder *flatbuffers.Builder, valueType PostingValueType) {
	builder.PrependInt32Slot(3, int32(valueType), 0)
}
func DirectedEdgeAddValueId(builder *flatbuffers.Builder, valueId uint64) {
	builder.PrependUint64Slot(4, valueId, 0)
}
func DirectedEdgeAddLabel(builder *flatbuffers.Builder, label flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(5, flatbuffers.UOffsetT(label), 0)
}
func DirectedEdgeAddLang(builder *flatbuffers.Builder, lang flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(6, flatbuffers.UOffsetT(lang), 0)
}
func DirectedEdgeAddOp(builder *flatbuffers.Builder, op DirectedEdgeOp) {
	builder.PrependInt32Slot(7, int32(op), 0)
}
func DirectedEdgeAddFacets(builder *flatbuffers.Builder, facets flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(8, flatbuffers.UOffsetT(facets), 0)
}
func DirectedEdgeStartFacetsVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func DirectedEdgeAddAllowedPreds(builder *flatbuffers.Builder, allowedPreds flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(9, flatbuffers.UOffsetT(allowedPreds), 0)
}
func DirectedEdgeStartAllowedPredsVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func DirectedEdgeEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
