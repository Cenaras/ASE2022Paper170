package main

import (
	u "Goat/analysis/upfront"
	"Goat/mayfailcast"
	"Goat/pkgutil"
	"fmt"
	"log"
	"os"
	"testing"
	"time"

	"golang.org/x/tools/go/ssa/ssautil"
)

func TestProfiling(t *testing.T) {
	gopath := "/home/cenaras/uni/masters/ForkASE/ASE2022Paper170/Goat/external/gfuzz"
	modulepath := "external/gfuzz/tidb"
	path := "github.com/pingcap/tidb/bindinfo"

	pkgs, _ := pkgutil.LoadPackages(pkgutil.LoadConfig{
		GoPath:       gopath,
		ModulePath:   modulepath,
		IncludeTests: true,
	}, path)

	prog, _ := ssautil.AllPackages(pkgs, 0)
	prog.Build()

	mains := ssautil.MainPackages(prog.AllPackages())

	if len(mains) == 0 {
		log.Println("No main packages detected")
		t.Fail()
	}

	allPackages := pkgutil.AllPackages(prog)
	pkgutil.GetLocalPackages(mains, allPackages)

	strategies := []string{"insens", "default", "1Obj", "2Obj+H", "1Call", "1Call+H", "U1Obj", "U2Obj+H",
		"SB1Obj", "SA1Obj", "2SObj+H", "2Call+H"}

	f, _ := os.Create("results.txt")
	defer f.Close()
	for _, strategy := range strategies {
		fmt.Println("Running strategy ", strategy)
		start := time.Now()
		ptaResult := u.Andersen(prog, mains, u.IncludeType{All: true}, strategy)
		elapsed := time.Since(start)
		f.WriteString(fmt.Sprintf("Running strategy %s\n", strategy))
		f.WriteString(fmt.Sprintf("Analysis took %fs\n", elapsed.Seconds()))
		f.WriteString(fmt.Sprintf("Number of nodes: %d\n", ptaResult.NoNodes))
		mayFails, totalCasts, totalOkCasts := mayfailcast.MayFailCast(prog, ptaResult)
		f.WriteString(fmt.Sprintf("Number of non-ok casts: %d\n", totalCasts))
		f.WriteString(fmt.Sprintf("Number of ok casts: %d\n", totalOkCasts))
		f.WriteString(fmt.Sprintf("Number of May Fail Casts: %d\n\n", len(mayFails)))
	}
}
