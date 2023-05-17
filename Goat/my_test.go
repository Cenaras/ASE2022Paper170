package main

import (
	"Goat/analysis/cfg"
	u "Goat/analysis/upfront"
	"Goat/pkgutil"
	"context"
	"fmt"
	"log"
	"testing"
	"time"

	"golang.org/x/tools/go/pointer"
	"golang.org/x/tools/go/ssa/ssautil"
)

func TestProfiling(t *testing.T) {
	gopath := "/home/cenaras/uni/masters/ForkASE/ASE2022Paper170/Goat/external/gfuzz"
	modulepath := "external/gfuzz/tidb"
	path := "github.com/pingcap/tidb/cmd/benchdb"

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


	preanalysisPipeline := func(includes u.IncludeType) (*pointer.Result, *cfg.Cfg) {
		fmt.Println()

		//strategy := opts.AnalysisStrategy()
		log.Printf("Performing points-to analysis with\n...")

		ctx, _ := context.WithTimeout(context.Background(), 5*time.Minute)
		c := make(chan *pointer.Result)

		start := time.Now()
		go func() {
			ptaResult := u.Andersen(prog, mains, includes, "")
			c <- ptaResult
		}()
		var ptaResult *pointer.Result
		select {
		case ptaResult = <-c:
		case <-ctx.Done():
			panic("Analysis timed out")
		}

		elapsed := time.Since(start)
		log.Printf("Points-to analysis took: %f seconds", elapsed.Seconds())
		log.Println("Points-to analysis done")
		fmt.Println()

		//log.Println("Extending CFG...")
		progCfg := cfg.GetCFG(prog, mains, ptaResult)
		//log.Println("CFG extensions done")
		//fmt.Println()

		opts.OnVerbose(func() {
			for val, ptr := range ptaResult.Queries {
				fmt.Printf("Points to information for \"%s\" at %d (%s):\n",
					val, val.Pos(), prog.Fset.Position(val.Pos()))
				for _, label := range ptr.PointsTo().Labels() {
					fmt.Printf("%s : %d (%s), ", label, (*label).Pos(), prog.Fset.Position((*label).Pos()))
				}
				fmt.Print("\n\n")
			}
		})

		return ptaResult, progCfg
	}

	preanalysisPipeline(u.IncludeType{All: true})
}