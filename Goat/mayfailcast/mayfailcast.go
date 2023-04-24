package mayfailcast

import (
	"go/types"

	"golang.org/x/tools/go/pointer"
	"golang.org/x/tools/go/ssa"
	"golang.org/x/tools/go/ssa/ssautil"
)

func MayFailCast(prog *ssa.Program, result *pointer.Result) ([]*ssa.TypeAssert, int, int) {
	var mayFails []*ssa.TypeAssert
	totalCasts := 0
	totalOkCasts := 0
	for fun := range ssautil.AllFunctions(prog) {
		for _, block := range fun.Blocks {
			for _, insn := range block.Instrs {
				switch v := insn.(type) {
				case *ssa.TypeAssert:
					res := result.Queries[v.X]
					if !res.IsValidPointer() {
						break
					}
					if v.CommaOk {
						totalOkCasts++
						break
					}
					totalCasts++
					ts := res.DynamicTypes().Keys()
					for _, k := range ts {
						if !types.AssignableTo(k, v.AssertedType) {
							mayFails = append(mayFails, v)
							break
						}
					}
				}
			}
		}
	}
	return mayFails, totalCasts, totalOkCasts
}
