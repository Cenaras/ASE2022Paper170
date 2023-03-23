// Code generated by generate-map.go for AnalysisIntraprocess. DO NOT EDIT.

package lattice

import "Goat/analysis/defs"

type AnalysisIntraprocessLattice struct {
	mapLatticeBase
}

func (m *AnalysisIntraprocessLattice) Eq(o Lattice) bool {
	switch o := o.(type) {
	case *AnalysisIntraprocessLattice:
		return true
	case *Lifted:
		return m.Eq(o.Lattice)
	case *Dropped:
		return m.Eq(o.Lattice)
	default:
		return false
	}
}

func (m *AnalysisIntraprocessLattice) Top() Element {
	panic(errUnsupportedOperation)
}

func (m *AnalysisIntraprocessLattice) AnalysisIntraprocess() *AnalysisIntraprocessLattice {
	return m
}

type AnalysisIntraprocess struct {
	element
	base baseMap
}

// Map methods
func (w AnalysisIntraprocess) Size() int {
	return w.base.Size()
}

func (w AnalysisIntraprocess) Height() int {
	return w.base.Height()
}

func (w AnalysisIntraprocess) Get(key defs.CtrLoc) (AnalysisState, bool) {
	v, found := w.base.Get(key)
	return v.(AnalysisState), found
}

func (w AnalysisIntraprocess) GetOrDefault(key defs.CtrLoc, dflt AnalysisState) AnalysisState {
	return w.base.GetOrDefault(key, dflt).(AnalysisState)
}

func (w AnalysisIntraprocess) GetUnsafe(key defs.CtrLoc) AnalysisState {
	return w.base.GetUnsafe(key).(AnalysisState)
}

func (w AnalysisIntraprocess) Update(key defs.CtrLoc, value AnalysisState) AnalysisIntraprocess {
	w.base = w.base.Update(key, value)
	return w
}

func (w AnalysisIntraprocess) WeakUpdate(key defs.CtrLoc, value AnalysisState) AnalysisIntraprocess {
	w.base = w.base.WeakUpdate(key, value)
	return w
}

func (w AnalysisIntraprocess) Remove(key defs.CtrLoc) AnalysisIntraprocess {
	w.base = w.base.Remove(key)
	return w
}

func (w AnalysisIntraprocess) ForEach(f func(defs.CtrLoc, AnalysisState)) {
	w.base.ForEach(func(key interface{}, value Element) {
		f(key.(defs.CtrLoc), value.(AnalysisState))
	})
}

func (w AnalysisIntraprocess) Find(f func(defs.CtrLoc, AnalysisState) bool) (zk defs.CtrLoc, zv AnalysisState, b bool) {
	k, e, found := w.base.Find(func(k interface{}, e Element) bool {
		return f(k.(defs.CtrLoc), e.(AnalysisState))
	})
	if found {
		return k.(defs.CtrLoc), e.(AnalysisState), true
	}
	return zk, zv, b
}

// Lattice element methods
func (w AnalysisIntraprocess) Leq(e Element) bool {
	checkLatticeMatch(w.lattice, e.Lattice(), "⊑")
	return w.leq(e)
}

func (w AnalysisIntraprocess) leq(e Element) bool {
	return w.base.leq(e.(AnalysisIntraprocess).base)
}

func (w AnalysisIntraprocess) Geq(e Element) bool {
	checkLatticeMatch(w.lattice, e.Lattice(), "⊒")
	return w.geq(e)
}

func (w AnalysisIntraprocess) geq(e Element) bool {
	return w.base.geq(e.(AnalysisIntraprocess).base)
}

func (w AnalysisIntraprocess) Eq(e Element) bool {
	checkLatticeMatch(w.lattice, e.Lattice(), "=")
	return w.eq(e)
}

func (w AnalysisIntraprocess) eq(e Element) bool {
	return w.base.eq(e.(AnalysisIntraprocess).base)
}

func (w AnalysisIntraprocess) Join(o Element) Element {
	checkLatticeMatch(w.lattice, o.Lattice(), "⊔")
	return w.join(o)
}

func (w AnalysisIntraprocess) join(o Element) Element {
	return w.MonoJoin(o.(AnalysisIntraprocess))
}

func (w AnalysisIntraprocess) MonoJoin(o AnalysisIntraprocess) AnalysisIntraprocess {
	w.base = w.base.MonoJoin(o.base)
	return w
}

func (w AnalysisIntraprocess) Meet(o Element) Element {
	panic(errUnsupportedOperation)
}

func (w AnalysisIntraprocess) meet(o Element) Element {
	panic(errUnsupportedOperation)
}

func (w AnalysisIntraprocess) String() string {
	return w.base.String()
}

// Type conversion
func (w AnalysisIntraprocess) AnalysisIntraprocess() AnalysisIntraprocess {
	return w
}
