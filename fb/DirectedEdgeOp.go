// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fb

import "strconv"

type DirectedEdgeOp int32

const (
	DirectedEdgeOpSET DirectedEdgeOp = 0
	DirectedEdgeOpDEL DirectedEdgeOp = 1
)

var EnumNamesDirectedEdgeOp = map[DirectedEdgeOp]string{
	DirectedEdgeOpSET: "SET",
	DirectedEdgeOpDEL: "DEL",
}

var EnumValuesDirectedEdgeOp = map[string]DirectedEdgeOp{
	"SET": DirectedEdgeOpSET,
	"DEL": DirectedEdgeOpDEL,
}

func (v DirectedEdgeOp) String() string {
	if s, ok := EnumNamesDirectedEdgeOp[v]; ok {
		return s
	}
	return "DirectedEdgeOp(" + strconv.FormatInt(int64(v), 10) + ")"
}
