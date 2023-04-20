package alias

import (
	"go/types"

	"golang.org/x/tools/go/pointer"
	"golang.org/x/tools/go/ssa"
)

type alias struct {
	channels   []ssa.Value
	signatures []ssa.Value
	interfaces []ssa.Value
	maps       []ssa.Value
	slices     []ssa.Value
	pointers   []ssa.Value
}

func GetAlias(result *pointer.Result) int {
	al := &alias{}
	for v, _ := range result.Queries {
		al.checkType(v, v.Type())
	}
	channels := checkAlias(result.Queries, al.channels)
	signatures := checkAlias(result.Queries, al.signatures)
	interfaces := checkAlias(result.Queries, al.interfaces)
	maps := checkAlias(result.Queries, al.maps)
	slices := checkAlias(result.Queries, al.slices)
	pointers := checkAlias(result.Queries, al.pointers)
	return len(channels) + len(signatures) + len(interfaces) + len(maps) + len(slices) + len(pointers)
}

func (al *alias) checkType(v ssa.Value, t types.Type) {
	switch t := t.(type) {
	case *types.Named:
		al.checkType(v, t.Underlying())
	case *types.Chan:
		al.channels = append(al.channels, v)
	case *types.Signature:
		al.signatures = append(al.signatures, v)
	case *types.Interface:
		al.interfaces = append(al.interfaces, v)
	case *types.Map:
		al.maps = append(al.maps, v)
	// case *types.Array:
	// 	if include.All || include.Array {
	// 		return _DIRECT
	// 	}
	case *types.Slice:
		al.slices = append(al.slices, v)
	case *types.Pointer:
		// If pointers are not considered for the PTA, add indirect queues
		// for the types which are included
		al.pointers = append(al.pointers, v)
	}
}

type pair struct {
	one ssa.Value
	two ssa.Value
}

func checkAlias(queries map[ssa.Value]pointer.Pointer, vals []ssa.Value) []*pair {
	var result []*pair
	for i, x := range vals {
		for _, y := range vals[i+1:] {
			if queries[x].MayAlias(queries[y]) {
				result = append(result, &pair{x, y})
			}
		}
	}
	return result
}
