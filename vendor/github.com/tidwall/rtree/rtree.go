// generated; DO NOT EDIT!

package rtree

import "math"

type Iterator func(item Item) bool
type Item interface {
	Rect(ctx interface{}) (min []float64, max []float64)
}

type RTree struct {
	ctx  interface{}
	tr1  *d1RTree
	tr2  *d2RTree
	tr3  *d3RTree
	tr4  *d4RTree
	tr5  *d5RTree
	tr6  *d6RTree
	tr7  *d7RTree
	tr8  *d8RTree
	tr9  *d9RTree
	tr10 *d10RTree
	tr11 *d11RTree
	tr12 *d12RTree
	tr13 *d13RTree
	tr14 *d14RTree
	tr15 *d15RTree
	tr16 *d16RTree
	tr17 *d17RTree
	tr18 *d18RTree
	tr19 *d19RTree
	tr20 *d20RTree
}

func New(ctx interface{}) *RTree {
	return &RTree{
		ctx:  ctx,
		tr1:  d1New(),
		tr2:  d2New(),
		tr3:  d3New(),
		tr4:  d4New(),
		tr5:  d5New(),
		tr6:  d6New(),
		tr7:  d7New(),
		tr8:  d8New(),
		tr9:  d9New(),
		tr10: d10New(),
		tr11: d11New(),
		tr12: d12New(),
		tr13: d13New(),
		tr14: d14New(),
		tr15: d15New(),
		tr16: d16New(),
		tr17: d17New(),
		tr18: d18New(),
		tr19: d19New(),
		tr20: d20New(),
	}
}

func (tr *RTree) Insert(item Item) {
	if item == nil {
		panic("nil item being added to RTree")
	}
	min, max := item.Rect(tr.ctx)
	if len(min) != len(max) {
		return // just return
		panic("invalid item rectangle")
	}
	switch len(min) {
	default:
		return // just return
		panic("invalid dimension")
	case 1:
		var amin, amax [1]float64
		for i := 0; i < len(min); i++ {
			amin[i], amax[i] = min[i], max[i]
		}
		tr.tr1.Insert(amin, amax, item)
	case 2:
		var amin, amax [2]float64
		for i := 0; i < len(min); i++ {
			amin[i], amax[i] = min[i], max[i]
		}
		tr.tr2.Insert(amin, amax, item)
	case 3:
		var amin, amax [3]float64
		for i := 0; i < len(min); i++ {
			amin[i], amax[i] = min[i], max[i]
		}
		tr.tr3.Insert(amin, amax, item)
	case 4:
		var amin, amax [4]float64
		for i := 0; i < len(min); i++ {
			amin[i], amax[i] = min[i], max[i]
		}
		tr.tr4.Insert(amin, amax, item)
	case 5:
		var amin, amax [5]float64
		for i := 0; i < len(min); i++ {
			amin[i], amax[i] = min[i], max[i]
		}
		tr.tr5.Insert(amin, amax, item)
	case 6:
		var amin, amax [6]float64
		for i := 0; i < len(min); i++ {
			amin[i], amax[i] = min[i], max[i]
		}
		tr.tr6.Insert(amin, amax, item)
	case 7:
		var amin, amax [7]float64
		for i := 0; i < len(min); i++ {
			amin[i], amax[i] = min[i], max[i]
		}
		tr.tr7.Insert(amin, amax, item)
	case 8:
		var amin, amax [8]float64
		for i := 0; i < len(min); i++ {
			amin[i], amax[i] = min[i], max[i]
		}
		tr.tr8.Insert(amin, amax, item)
	case 9:
		var amin, amax [9]float64
		for i := 0; i < len(min); i++ {
			amin[i], amax[i] = min[i], max[i]
		}
		tr.tr9.Insert(amin, amax, item)
	case 10:
		var amin, amax [10]float64
		for i := 0; i < len(min); i++ {
			amin[i], amax[i] = min[i], max[i]
		}
		tr.tr10.Insert(amin, amax, item)
	case 11:
		var amin, amax [11]float64
		for i := 0; i < len(min); i++ {
			amin[i], amax[i] = min[i], max[i]
		}
		tr.tr11.Insert(amin, amax, item)
	case 12:
		var amin, amax [12]float64
		for i := 0; i < len(min); i++ {
			amin[i], amax[i] = min[i], max[i]
		}
		tr.tr12.Insert(amin, amax, item)
	case 13:
		var amin, amax [13]float64
		for i := 0; i < len(min); i++ {
			amin[i], amax[i] = min[i], max[i]
		}
		tr.tr13.Insert(amin, amax, item)
	case 14:
		var amin, amax [14]float64
		for i := 0; i < len(min); i++ {
			amin[i], amax[i] = min[i], max[i]
		}
		tr.tr14.Insert(amin, amax, item)
	case 15:
		var amin, amax [15]float64
		for i := 0; i < len(min); i++ {
			amin[i], amax[i] = min[i], max[i]
		}
		tr.tr15.Insert(amin, amax, item)
	case 16:
		var amin, amax [16]float64
		for i := 0; i < len(min); i++ {
			amin[i], amax[i] = min[i], max[i]
		}
		tr.tr16.Insert(amin, amax, item)
	case 17:
		var amin, amax [17]float64
		for i := 0; i < len(min); i++ {
			amin[i], amax[i] = min[i], max[i]
		}
		tr.tr17.Insert(amin, amax, item)
	case 18:
		var amin, amax [18]float64
		for i := 0; i < len(min); i++ {
			amin[i], amax[i] = min[i], max[i]
		}
		tr.tr18.Insert(amin, amax, item)
	case 19:
		var amin, amax [19]float64
		for i := 0; i < len(min); i++ {
			amin[i], amax[i] = min[i], max[i]
		}
		tr.tr19.Insert(amin, amax, item)
	case 20:
		var amin, amax [20]float64
		for i := 0; i < len(min); i++ {
			amin[i], amax[i] = min[i], max[i]
		}
		tr.tr20.Insert(amin, amax, item)
	}
}

func (tr *RTree) Remove(item Item) {
	if item == nil {
		panic("nil item being added to RTree")
	}
	min, max := item.Rect(tr.ctx)
	if len(min) != len(max) {
		return // just return
		panic("invalid item rectangle")
	}
	switch len(min) {
	default:
		return // just return
		panic("invalid dimension")
	case 1:
		var amin, amax [1]float64
		for i := 0; i < len(min); i++ {
			amin[i], amax[i] = min[i], max[i]
		}
		tr.tr1.Remove(amin, amax, item)
	case 2:
		var amin, amax [2]float64
		for i := 0; i < len(min); i++ {
			amin[i], amax[i] = min[i], max[i]
		}
		tr.tr2.Remove(amin, amax, item)
	case 3:
		var amin, amax [3]float64
		for i := 0; i < len(min); i++ {
			amin[i], amax[i] = min[i], max[i]
		}
		tr.tr3.Remove(amin, amax, item)
	case 4:
		var amin, amax [4]float64
		for i := 0; i < len(min); i++ {
			amin[i], amax[i] = min[i], max[i]
		}
		tr.tr4.Remove(amin, amax, item)
	case 5:
		var amin, amax [5]float64
		for i := 0; i < len(min); i++ {
			amin[i], amax[i] = min[i], max[i]
		}
		tr.tr5.Remove(amin, amax, item)
	case 6:
		var amin, amax [6]float64
		for i := 0; i < len(min); i++ {
			amin[i], amax[i] = min[i], max[i]
		}
		tr.tr6.Remove(amin, amax, item)
	case 7:
		var amin, amax [7]float64
		for i := 0; i < len(min); i++ {
			amin[i], amax[i] = min[i], max[i]
		}
		tr.tr7.Remove(amin, amax, item)
	case 8:
		var amin, amax [8]float64
		for i := 0; i < len(min); i++ {
			amin[i], amax[i] = min[i], max[i]
		}
		tr.tr8.Remove(amin, amax, item)
	case 9:
		var amin, amax [9]float64
		for i := 0; i < len(min); i++ {
			amin[i], amax[i] = min[i], max[i]
		}
		tr.tr9.Remove(amin, amax, item)
	case 10:
		var amin, amax [10]float64
		for i := 0; i < len(min); i++ {
			amin[i], amax[i] = min[i], max[i]
		}
		tr.tr10.Remove(amin, amax, item)
	case 11:
		var amin, amax [11]float64
		for i := 0; i < len(min); i++ {
			amin[i], amax[i] = min[i], max[i]
		}
		tr.tr11.Remove(amin, amax, item)
	case 12:
		var amin, amax [12]float64
		for i := 0; i < len(min); i++ {
			amin[i], amax[i] = min[i], max[i]
		}
		tr.tr12.Remove(amin, amax, item)
	case 13:
		var amin, amax [13]float64
		for i := 0; i < len(min); i++ {
			amin[i], amax[i] = min[i], max[i]
		}
		tr.tr13.Remove(amin, amax, item)
	case 14:
		var amin, amax [14]float64
		for i := 0; i < len(min); i++ {
			amin[i], amax[i] = min[i], max[i]
		}
		tr.tr14.Remove(amin, amax, item)
	case 15:
		var amin, amax [15]float64
		for i := 0; i < len(min); i++ {
			amin[i], amax[i] = min[i], max[i]
		}
		tr.tr15.Remove(amin, amax, item)
	case 16:
		var amin, amax [16]float64
		for i := 0; i < len(min); i++ {
			amin[i], amax[i] = min[i], max[i]
		}
		tr.tr16.Remove(amin, amax, item)
	case 17:
		var amin, amax [17]float64
		for i := 0; i < len(min); i++ {
			amin[i], amax[i] = min[i], max[i]
		}
		tr.tr17.Remove(amin, amax, item)
	case 18:
		var amin, amax [18]float64
		for i := 0; i < len(min); i++ {
			amin[i], amax[i] = min[i], max[i]
		}
		tr.tr18.Remove(amin, amax, item)
	case 19:
		var amin, amax [19]float64
		for i := 0; i < len(min); i++ {
			amin[i], amax[i] = min[i], max[i]
		}
		tr.tr19.Remove(amin, amax, item)
	case 20:
		var amin, amax [20]float64
		for i := 0; i < len(min); i++ {
			amin[i], amax[i] = min[i], max[i]
		}
		tr.tr20.Remove(amin, amax, item)
	}
}
func (tr *RTree) Reset() {
	tr.tr1 = d1New()
	tr.tr2 = d2New()
	tr.tr3 = d3New()
	tr.tr4 = d4New()
	tr.tr5 = d5New()
	tr.tr6 = d6New()
	tr.tr7 = d7New()
	tr.tr8 = d8New()
	tr.tr9 = d9New()
	tr.tr10 = d10New()
	tr.tr11 = d11New()
	tr.tr12 = d12New()
	tr.tr13 = d13New()
	tr.tr14 = d14New()
	tr.tr15 = d15New()
	tr.tr16 = d16New()
	tr.tr17 = d17New()
	tr.tr18 = d18New()
	tr.tr19 = d19New()
	tr.tr20 = d20New()
}
func (tr *RTree) Count() int {
	count := 0
	count += tr.tr1.Count()
	count += tr.tr2.Count()
	count += tr.tr3.Count()
	count += tr.tr4.Count()
	count += tr.tr5.Count()
	count += tr.tr6.Count()
	count += tr.tr7.Count()
	count += tr.tr8.Count()
	count += tr.tr9.Count()
	count += tr.tr10.Count()
	count += tr.tr11.Count()
	count += tr.tr12.Count()
	count += tr.tr13.Count()
	count += tr.tr14.Count()
	count += tr.tr15.Count()
	count += tr.tr16.Count()
	count += tr.tr17.Count()
	count += tr.tr18.Count()
	count += tr.tr19.Count()
	count += tr.tr20.Count()
	return count
}
func (tr *RTree) Search(bounds Item, iter Iterator) {
	if bounds == nil {
		panic("nil bounds being used for search")
	}
	min, max := bounds.Rect(tr.ctx)
	if len(min) != len(max) {
		return // just return
		panic("invalid item rectangle")
	}
	switch len(min) {
	default:
		return // just return
		panic("invalid dimension")
	case 1:
	case 2:
	case 3:
	case 4:
	case 5:
	case 6:
	case 7:
	case 8:
	case 9:
	case 10:
	case 11:
	case 12:
	case 13:
	case 14:
	case 15:
	case 16:
	case 17:
	case 18:
	case 19:
	case 20:
	}
	if !tr.search1(min, max, iter) {
		return
	}
	if !tr.search2(min, max, iter) {
		return
	}
	if !tr.search3(min, max, iter) {
		return
	}
	if !tr.search4(min, max, iter) {
		return
	}
	if !tr.search5(min, max, iter) {
		return
	}
	if !tr.search6(min, max, iter) {
		return
	}
	if !tr.search7(min, max, iter) {
		return
	}
	if !tr.search8(min, max, iter) {
		return
	}
	if !tr.search9(min, max, iter) {
		return
	}
	if !tr.search10(min, max, iter) {
		return
	}
	if !tr.search11(min, max, iter) {
		return
	}
	if !tr.search12(min, max, iter) {
		return
	}
	if !tr.search13(min, max, iter) {
		return
	}
	if !tr.search14(min, max, iter) {
		return
	}
	if !tr.search15(min, max, iter) {
		return
	}
	if !tr.search16(min, max, iter) {
		return
	}
	if !tr.search17(min, max, iter) {
		return
	}
	if !tr.search18(min, max, iter) {
		return
	}
	if !tr.search19(min, max, iter) {
		return
	}
	if !tr.search20(min, max, iter) {
		return
	}
}

func (tr *RTree) search1(min, max []float64, iter Iterator) bool {
	var amin, amax [1]float64
	for i := 0; i < 1; i++ {
		if i < len(min) {
			amin[i] = min[i]
			amax[i] = max[i]
		} else {
			amin[i] = math.Inf(-1)
			amax[i] = math.Inf(+1)
		}
	}
	ended := false
	tr.tr1.Search(amin, amax, func(dataID interface{}) bool {
		if !iter(dataID.(Item)) {
			ended = true
			return false
		}
		return true
	})
	return !ended
}

func (tr *RTree) search2(min, max []float64, iter Iterator) bool {
	var amin, amax [2]float64
	for i := 0; i < 2; i++ {
		if i < len(min) {
			amin[i] = min[i]
			amax[i] = max[i]
		} else {
			amin[i] = math.Inf(-1)
			amax[i] = math.Inf(+1)
		}
	}
	ended := false
	tr.tr2.Search(amin, amax, func(dataID interface{}) bool {
		if !iter(dataID.(Item)) {
			ended = true
			return false
		}
		return true
	})
	return !ended
}

func (tr *RTree) search3(min, max []float64, iter Iterator) bool {
	var amin, amax [3]float64
	for i := 0; i < 3; i++ {
		if i < len(min) {
			amin[i] = min[i]
			amax[i] = max[i]
		} else {
			amin[i] = math.Inf(-1)
			amax[i] = math.Inf(+1)
		}
	}
	ended := false
	tr.tr3.Search(amin, amax, func(dataID interface{}) bool {
		if !iter(dataID.(Item)) {
			ended = true
			return false
		}
		return true
	})
	return !ended
}

func (tr *RTree) search4(min, max []float64, iter Iterator) bool {
	var amin, amax [4]float64
	for i := 0; i < 4; i++ {
		if i < len(min) {
			amin[i] = min[i]
			amax[i] = max[i]
		} else {
			amin[i] = math.Inf(-1)
			amax[i] = math.Inf(+1)
		}
	}
	ended := false
	tr.tr4.Search(amin, amax, func(dataID interface{}) bool {
		if !iter(dataID.(Item)) {
			ended = true
			return false
		}
		return true
	})
	return !ended
}

func (tr *RTree) search5(min, max []float64, iter Iterator) bool {
	var amin, amax [5]float64
	for i := 0; i < 5; i++ {
		if i < len(min) {
			amin[i] = min[i]
			amax[i] = max[i]
		} else {
			amin[i] = math.Inf(-1)
			amax[i] = math.Inf(+1)
		}
	}
	ended := false
	tr.tr5.Search(amin, amax, func(dataID interface{}) bool {
		if !iter(dataID.(Item)) {
			ended = true
			return false
		}
		return true
	})
	return !ended
}

func (tr *RTree) search6(min, max []float64, iter Iterator) bool {
	var amin, amax [6]float64
	for i := 0; i < 6; i++ {
		if i < len(min) {
			amin[i] = min[i]
			amax[i] = max[i]
		} else {
			amin[i] = math.Inf(-1)
			amax[i] = math.Inf(+1)
		}
	}
	ended := false
	tr.tr6.Search(amin, amax, func(dataID interface{}) bool {
		if !iter(dataID.(Item)) {
			ended = true
			return false
		}
		return true
	})
	return !ended
}

func (tr *RTree) search7(min, max []float64, iter Iterator) bool {
	var amin, amax [7]float64
	for i := 0; i < 7; i++ {
		if i < len(min) {
			amin[i] = min[i]
			amax[i] = max[i]
		} else {
			amin[i] = math.Inf(-1)
			amax[i] = math.Inf(+1)
		}
	}
	ended := false
	tr.tr7.Search(amin, amax, func(dataID interface{}) bool {
		if !iter(dataID.(Item)) {
			ended = true
			return false
		}
		return true
	})
	return !ended
}

func (tr *RTree) search8(min, max []float64, iter Iterator) bool {
	var amin, amax [8]float64
	for i := 0; i < 8; i++ {
		if i < len(min) {
			amin[i] = min[i]
			amax[i] = max[i]
		} else {
			amin[i] = math.Inf(-1)
			amax[i] = math.Inf(+1)
		}
	}
	ended := false
	tr.tr8.Search(amin, amax, func(dataID interface{}) bool {
		if !iter(dataID.(Item)) {
			ended = true
			return false
		}
		return true
	})
	return !ended
}

func (tr *RTree) search9(min, max []float64, iter Iterator) bool {
	var amin, amax [9]float64
	for i := 0; i < 9; i++ {
		if i < len(min) {
			amin[i] = min[i]
			amax[i] = max[i]
		} else {
			amin[i] = math.Inf(-1)
			amax[i] = math.Inf(+1)
		}
	}
	ended := false
	tr.tr9.Search(amin, amax, func(dataID interface{}) bool {
		if !iter(dataID.(Item)) {
			ended = true
			return false
		}
		return true
	})
	return !ended
}

func (tr *RTree) search10(min, max []float64, iter Iterator) bool {
	var amin, amax [10]float64
	for i := 0; i < 10; i++ {
		if i < len(min) {
			amin[i] = min[i]
			amax[i] = max[i]
		} else {
			amin[i] = math.Inf(-1)
			amax[i] = math.Inf(+1)
		}
	}
	ended := false
	tr.tr10.Search(amin, amax, func(dataID interface{}) bool {
		if !iter(dataID.(Item)) {
			ended = true
			return false
		}
		return true
	})
	return !ended
}

func (tr *RTree) search11(min, max []float64, iter Iterator) bool {
	var amin, amax [11]float64
	for i := 0; i < 11; i++ {
		if i < len(min) {
			amin[i] = min[i]
			amax[i] = max[i]
		} else {
			amin[i] = math.Inf(-1)
			amax[i] = math.Inf(+1)
		}
	}
	ended := false
	tr.tr11.Search(amin, amax, func(dataID interface{}) bool {
		if !iter(dataID.(Item)) {
			ended = true
			return false
		}
		return true
	})
	return !ended
}

func (tr *RTree) search12(min, max []float64, iter Iterator) bool {
	var amin, amax [12]float64
	for i := 0; i < 12; i++ {
		if i < len(min) {
			amin[i] = min[i]
			amax[i] = max[i]
		} else {
			amin[i] = math.Inf(-1)
			amax[i] = math.Inf(+1)
		}
	}
	ended := false
	tr.tr12.Search(amin, amax, func(dataID interface{}) bool {
		if !iter(dataID.(Item)) {
			ended = true
			return false
		}
		return true
	})
	return !ended
}

func (tr *RTree) search13(min, max []float64, iter Iterator) bool {
	var amin, amax [13]float64
	for i := 0; i < 13; i++ {
		if i < len(min) {
			amin[i] = min[i]
			amax[i] = max[i]
		} else {
			amin[i] = math.Inf(-1)
			amax[i] = math.Inf(+1)
		}
	}
	ended := false
	tr.tr13.Search(amin, amax, func(dataID interface{}) bool {
		if !iter(dataID.(Item)) {
			ended = true
			return false
		}
		return true
	})
	return !ended
}

func (tr *RTree) search14(min, max []float64, iter Iterator) bool {
	var amin, amax [14]float64
	for i := 0; i < 14; i++ {
		if i < len(min) {
			amin[i] = min[i]
			amax[i] = max[i]
		} else {
			amin[i] = math.Inf(-1)
			amax[i] = math.Inf(+1)
		}
	}
	ended := false
	tr.tr14.Search(amin, amax, func(dataID interface{}) bool {
		if !iter(dataID.(Item)) {
			ended = true
			return false
		}
		return true
	})
	return !ended
}

func (tr *RTree) search15(min, max []float64, iter Iterator) bool {
	var amin, amax [15]float64
	for i := 0; i < 15; i++ {
		if i < len(min) {
			amin[i] = min[i]
			amax[i] = max[i]
		} else {
			amin[i] = math.Inf(-1)
			amax[i] = math.Inf(+1)
		}
	}
	ended := false
	tr.tr15.Search(amin, amax, func(dataID interface{}) bool {
		if !iter(dataID.(Item)) {
			ended = true
			return false
		}
		return true
	})
	return !ended
}

func (tr *RTree) search16(min, max []float64, iter Iterator) bool {
	var amin, amax [16]float64
	for i := 0; i < 16; i++ {
		if i < len(min) {
			amin[i] = min[i]
			amax[i] = max[i]
		} else {
			amin[i] = math.Inf(-1)
			amax[i] = math.Inf(+1)
		}
	}
	ended := false
	tr.tr16.Search(amin, amax, func(dataID interface{}) bool {
		if !iter(dataID.(Item)) {
			ended = true
			return false
		}
		return true
	})
	return !ended
}

func (tr *RTree) search17(min, max []float64, iter Iterator) bool {
	var amin, amax [17]float64
	for i := 0; i < 17; i++ {
		if i < len(min) {
			amin[i] = min[i]
			amax[i] = max[i]
		} else {
			amin[i] = math.Inf(-1)
			amax[i] = math.Inf(+1)
		}
	}
	ended := false
	tr.tr17.Search(amin, amax, func(dataID interface{}) bool {
		if !iter(dataID.(Item)) {
			ended = true
			return false
		}
		return true
	})
	return !ended
}

func (tr *RTree) search18(min, max []float64, iter Iterator) bool {
	var amin, amax [18]float64
	for i := 0; i < 18; i++ {
		if i < len(min) {
			amin[i] = min[i]
			amax[i] = max[i]
		} else {
			amin[i] = math.Inf(-1)
			amax[i] = math.Inf(+1)
		}
	}
	ended := false
	tr.tr18.Search(amin, amax, func(dataID interface{}) bool {
		if !iter(dataID.(Item)) {
			ended = true
			return false
		}
		return true
	})
	return !ended
}

func (tr *RTree) search19(min, max []float64, iter Iterator) bool {
	var amin, amax [19]float64
	for i := 0; i < 19; i++ {
		if i < len(min) {
			amin[i] = min[i]
			amax[i] = max[i]
		} else {
			amin[i] = math.Inf(-1)
			amax[i] = math.Inf(+1)
		}
	}
	ended := false
	tr.tr19.Search(amin, amax, func(dataID interface{}) bool {
		if !iter(dataID.(Item)) {
			ended = true
			return false
		}
		return true
	})
	return !ended
}

func (tr *RTree) search20(min, max []float64, iter Iterator) bool {
	var amin, amax [20]float64
	for i := 0; i < 20; i++ {
		if i < len(min) {
			amin[i] = min[i]
			amax[i] = max[i]
		} else {
			amin[i] = math.Inf(-1)
			amax[i] = math.Inf(+1)
		}
	}
	ended := false
	tr.tr20.Search(amin, amax, func(dataID interface{}) bool {
		if !iter(dataID.(Item)) {
			ended = true
			return false
		}
		return true
	})
	return !ended
}

func d1fmin(a, b float64) float64 {
	if a < b {
		return a
	}
	return b
}
func d1fmax(a, b float64) float64 {
	if a > b {
		return a
	}
	return b
}

const (
	d1numDims            = 1
	d1maxNodes           = 8
	d1minNodes           = d1maxNodes / 2
	d1useSphericalVolume = true // Better split classification, may be slower on some systems
)

var d1unitSphereVolume = []float64{
	0.000000, 2.000000, 3.141593, // Dimension  0,1,2
	4.188790, 4.934802, 5.263789, // Dimension  3,4,5
	5.167713, 4.724766, 4.058712, // Dimension  6,7,8
	3.298509, 2.550164, 1.884104, // Dimension  9,10,11
	1.335263, 0.910629, 0.599265, // Dimension  12,13,14
	0.381443, 0.235331, 0.140981, // Dimension  15,16,17
	0.082146, 0.046622, 0.025807, // Dimension  18,19,20
}[d1numDims]

type d1RTree struct {
	root *d1nodeT ///< Root of tree
}

/// Minimal bounding rectangle (n-dimensional)
type d1rectT struct {
	min [d1numDims]float64 ///< Min dimensions of bounding box
	max [d1numDims]float64 ///< Max dimensions of bounding box
}

/// May be data or may be another subtree
/// The parents level determines this.
/// If the parents level is 0, then this is data
type d1branchT struct {
	rect  d1rectT     ///< Bounds
	child *d1nodeT    ///< Child node
	data  interface{} ///< Data Id or Ptr
}

/// d1nodeT for each branch level
type d1nodeT struct {
	count  int                   ///< Count
	level  int                   ///< Leaf is zero, others positive
	branch [d1maxNodes]d1branchT ///< Branch
}

func (node *d1nodeT) isInternalNode() bool {
	return (node.level > 0) // Not a leaf, but a internal node
}
func (node *d1nodeT) isLeaf() bool {
	return (node.level == 0) // A leaf, contains data
}

/// A link list of nodes for reinsertion after a delete operation
type d1listNodeT struct {
	next *d1listNodeT ///< Next in list
	node *d1nodeT     ///< Node
}

const d1notTaken = -1 // indicates that position

/// Variables for finding a split partition
type d1partitionVarsT struct {
	partition [d1maxNodes + 1]int
	total     int
	minFill   int
	count     [2]int
	cover     [2]d1rectT
	area      [2]float64

	branchBuf      [d1maxNodes + 1]d1branchT
	branchCount    int
	coverSplit     d1rectT
	coverSplitArea float64
}

func d1New() *d1RTree {
	// We only support machine word size simple data type eg. integer index or object pointer.
	// Since we are storing as union with non data branch
	return &d1RTree{
		root: &d1nodeT{},
	}
}

/// Insert entry
/// \param a_min Min of bounding rect
/// \param a_max Max of bounding rect
/// \param a_dataId Positive Id of data.  Maybe zero, but negative numbers not allowed.
func (tr *d1RTree) Insert(min, max [d1numDims]float64, dataId interface{}) {
	var branch d1branchT
	branch.data = dataId
	for axis := 0; axis < d1numDims; axis++ {
		branch.rect.min[axis] = min[axis]
		branch.rect.max[axis] = max[axis]
	}
	d1insertRect(&branch, &tr.root, 0)
}

/// Remove entry
/// \param a_min Min of bounding rect
/// \param a_max Max of bounding rect
/// \param a_dataId Positive Id of data.  Maybe zero, but negative numbers not allowed.
func (tr *d1RTree) Remove(min, max [d1numDims]float64, dataId interface{}) {
	var rect d1rectT
	for axis := 0; axis < d1numDims; axis++ {
		rect.min[axis] = min[axis]
		rect.max[axis] = max[axis]
	}
	d1removeRect(&rect, dataId, &tr.root)
}

/// Find all within d1search rectangle
/// \param a_min Min of d1search bounding rect
/// \param a_max Max of d1search bounding rect
/// \param a_searchResult d1search result array.  Caller should set grow size. Function will reset, not append to array.
/// \param a_resultCallback Callback function to return result.  Callback should return 'true' to continue searching
/// \param a_context User context to pass as parameter to a_resultCallback
/// \return Returns the number of entries found
func (tr *d1RTree) Search(min, max [d1numDims]float64, resultCallback func(data interface{}) bool) int {
	var rect d1rectT
	for axis := 0; axis < d1numDims; axis++ {
		rect.min[axis] = min[axis]
		rect.max[axis] = max[axis]
	}
	foundCount, _ := d1search(tr.root, rect, 0, resultCallback)
	return foundCount
}

/// Count the data elements in this container.  This is slow as no internal counter is maintained.
func (tr *d1RTree) Count() int {
	var count int
	d1countRec(tr.root, &count)
	return count
}

/// Remove all entries from tree
func (tr *d1RTree) RemoveAll() {
	// Delete all existing nodes
	tr.root = &d1nodeT{}
}

func d1countRec(node *d1nodeT, count *int) {
	if node.isInternalNode() { // not a leaf node
		for index := 0; index < node.count; index++ {
			d1countRec(node.branch[index].child, count)
		}
	} else { // A leaf node
		*count += node.count
	}
}

// Inserts a new data rectangle into the index structure.
// Recursively descends tree, propagates splits back up.
// Returns 0 if node was not split.  Old node updated.
// If node was split, returns 1 and sets the pointer pointed to by
// new_node to point to the new node.  Old node updated to become one of two.
// The level argument specifies the number of steps up from the leaf
// level to insert; e.g. a data rectangle goes in at level = 0.
func d1insertRectRec(branch *d1branchT, node *d1nodeT, newNode **d1nodeT, level int) bool {
	// recurse until we reach the correct level for the new record. data records
	// will always be called with a_level == 0 (leaf)
	if node.level > level {
		// Still above level for insertion, go down tree recursively
		var otherNode *d1nodeT
		//var newBranch d1branchT

		// find the optimal branch for this record
		index := d1pickBranch(&branch.rect, node)

		// recursively insert this record into the picked branch
		childWasSplit := d1insertRectRec(branch, node.branch[index].child, &otherNode, level)

		if !childWasSplit {
			// Child was not split. Merge the bounding box of the new record with the
			// existing bounding box
			node.branch[index].rect = d1combineRect(&branch.rect, &(node.branch[index].rect))
			return false
		} else {
			// Child was split. The old branches are now re-partitioned to two nodes
			// so we have to re-calculate the bounding boxes of each node
			node.branch[index].rect = d1nodeCover(node.branch[index].child)
			var newBranch d1branchT
			newBranch.child = otherNode
			newBranch.rect = d1nodeCover(otherNode)

			// The old node is already a child of a_node. Now add the newly-created
			// node to a_node as well. a_node might be split because of that.
			return d1addBranch(&newBranch, node, newNode)
		}
	} else if node.level == level {
		// We have reached level for insertion. Add rect, split if necessary
		return d1addBranch(branch, node, newNode)
	} else {
		// Should never occur
		return false
	}
}

// Insert a data rectangle into an index structure.
// d1insertRect provides for splitting the root;
// returns 1 if root was split, 0 if it was not.
// The level argument specifies the number of steps up from the leaf
// level to insert; e.g. a data rectangle goes in at level = 0.
// InsertRect2 does the recursion.
//
func d1insertRect(branch *d1branchT, root **d1nodeT, level int) bool {
	var newNode *d1nodeT

	if d1insertRectRec(branch, *root, &newNode, level) { // Root split

		// Grow tree taller and new root
		newRoot := &d1nodeT{}
		newRoot.level = (*root).level + 1

		var newBranch d1branchT

		// add old root node as a child of the new root
		newBranch.rect = d1nodeCover(*root)
		newBranch.child = *root
		d1addBranch(&newBranch, newRoot, nil)

		// add the split node as a child of the new root
		newBranch.rect = d1nodeCover(newNode)
		newBranch.child = newNode
		d1addBranch(&newBranch, newRoot, nil)

		// set the new root as the root node
		*root = newRoot

		return true
	}
	return false
}

// Find the smallest rectangle that includes all rectangles in branches of a node.
func d1nodeCover(node *d1nodeT) d1rectT {
	rect := node.branch[0].rect
	for index := 1; index < node.count; index++ {
		rect = d1combineRect(&rect, &(node.branch[index].rect))
	}
	return rect
}

// Add a branch to a node.  Split the node if necessary.
// Returns 0 if node not split.  Old node updated.
// Returns 1 if node split, sets *new_node to address of new node.
// Old node updated, becomes one of two.
func d1addBranch(branch *d1branchT, node *d1nodeT, newNode **d1nodeT) bool {
	if node.count < d1maxNodes { // Split won't be necessary
		node.branch[node.count] = *branch
		node.count++
		return false
	} else {
		d1splitNode(node, branch, newNode)
		return true
	}
}

// Disconnect a dependent node.
// Caller must return (or stop using iteration index) after this as count has changed
func d1disconnectBranch(node *d1nodeT, index int) {
	// Remove element by swapping with the last element to prevent gaps in array
	node.branch[index] = node.branch[node.count-1]
	node.branch[node.count-1].data = nil
	node.branch[node.count-1].child = nil
	node.count--
}

// Pick a branch.  Pick the one that will need the smallest increase
// in area to accomodate the new rectangle.  This will result in the
// least total area for the covering rectangles in the current node.
// In case of a tie, pick the one which was smaller before, to get
// the best resolution when searching.
func d1pickBranch(rect *d1rectT, node *d1nodeT) int {
	var firstTime bool = true
	var increase float64
	var bestIncr float64 = -1
	var area float64
	var bestArea float64
	var best int
	var tempRect d1rectT

	for index := 0; index < node.count; index++ {
		curRect := &node.branch[index].rect
		area = d1calcRectVolume(curRect)
		tempRect = d1combineRect(rect, curRect)
		increase = d1calcRectVolume(&tempRect) - area
		if (increase < bestIncr) || firstTime {
			best = index
			bestArea = area
			bestIncr = increase
			firstTime = false
		} else if (increase == bestIncr) && (area < bestArea) {
			best = index
			bestArea = area
			bestIncr = increase
		}
	}
	return best
}

// Combine two rectangles into larger one containing both
func d1combineRect(rectA, rectB *d1rectT) d1rectT {
	var newRect d1rectT

	for index := 0; index < d1numDims; index++ {
		newRect.min[index] = d1fmin(rectA.min[index], rectB.min[index])
		newRect.max[index] = d1fmax(rectA.max[index], rectB.max[index])
	}

	return newRect
}

// Split a node.
// Divides the nodes branches and the extra one between two nodes.
// Old node is one of the new ones, and one really new one is created.
// Tries more than one method for choosing a partition, uses best result.
func d1splitNode(node *d1nodeT, branch *d1branchT, newNode **d1nodeT) {
	// Could just use local here, but member or external is faster since it is reused
	var localVars d1partitionVarsT
	parVars := &localVars

	// Load all the branches into a buffer, initialize old node
	d1getBranches(node, branch, parVars)

	// Find partition
	d1choosePartition(parVars, d1minNodes)

	// Create a new node to hold (about) half of the branches
	*newNode = &d1nodeT{}
	(*newNode).level = node.level

	// Put branches from buffer into 2 nodes according to the chosen partition
	node.count = 0
	d1loadNodes(node, *newNode, parVars)
}

// Calculate the n-dimensional volume of a rectangle
func d1rectVolume(rect *d1rectT) float64 {
	var volume float64 = 1
	for index := 0; index < d1numDims; index++ {
		volume *= rect.max[index] - rect.min[index]
	}
	return volume
}

// The exact volume of the bounding sphere for the given d1rectT
func d1rectSphericalVolume(rect *d1rectT) float64 {
	var sumOfSquares float64 = 0
	var radius float64

	for index := 0; index < d1numDims; index++ {
		halfExtent := (rect.max[index] - rect.min[index]) * 0.5
		sumOfSquares += halfExtent * halfExtent
	}

	radius = math.Sqrt(sumOfSquares)

	// Pow maybe slow, so test for common dims just use x*x, x*x*x.
	if d1numDims == 5 {
		return (radius * radius * radius * radius * radius * d1unitSphereVolume)
	} else if d1numDims == 4 {
		return (radius * radius * radius * radius * d1unitSphereVolume)
	} else if d1numDims == 3 {
		return (radius * radius * radius * d1unitSphereVolume)
	} else if d1numDims == 2 {
		return (radius * radius * d1unitSphereVolume)
	} else {
		return (math.Pow(radius, d1numDims) * d1unitSphereVolume)
	}
}

// Use one of the methods to calculate retangle volume
func d1calcRectVolume(rect *d1rectT) float64 {
	if d1useSphericalVolume {
		return d1rectSphericalVolume(rect) // Slower but helps certain merge cases
	} else { // RTREE_USE_SPHERICAL_VOLUME
		return d1rectVolume(rect) // Faster but can cause poor merges
	} // RTREE_USE_SPHERICAL_VOLUME
}

// Load branch buffer with branches from full node plus the extra branch.
func d1getBranches(node *d1nodeT, branch *d1branchT, parVars *d1partitionVarsT) {
	// Load the branch buffer
	for index := 0; index < d1maxNodes; index++ {
		parVars.branchBuf[index] = node.branch[index]
	}
	parVars.branchBuf[d1maxNodes] = *branch
	parVars.branchCount = d1maxNodes + 1

	// Calculate rect containing all in the set
	parVars.coverSplit = parVars.branchBuf[0].rect
	for index := 1; index < d1maxNodes+1; index++ {
		parVars.coverSplit = d1combineRect(&parVars.coverSplit, &parVars.branchBuf[index].rect)
	}
	parVars.coverSplitArea = d1calcRectVolume(&parVars.coverSplit)
}

// Method #0 for choosing a partition:
// As the seeds for the two groups, pick the two rects that would waste the
// most area if covered by a single rectangle, i.e. evidently the worst pair
// to have in the same group.
// Of the remaining, one at a time is chosen to be put in one of the two groups.
// The one chosen is the one with the greatest difference in area expansion
// depending on which group - the rect most strongly attracted to one group
// and repelled from the other.
// If one group gets too full (more would force other group to violate min
// fill requirement) then other group gets the rest.
// These last are the ones that can go in either group most easily.
func d1choosePartition(parVars *d1partitionVarsT, minFill int) {
	var biggestDiff float64
	var group, chosen, betterGroup int

	d1initParVars(parVars, parVars.branchCount, minFill)
	d1pickSeeds(parVars)

	for ((parVars.count[0] + parVars.count[1]) < parVars.total) &&
		(parVars.count[0] < (parVars.total - parVars.minFill)) &&
		(parVars.count[1] < (parVars.total - parVars.minFill)) {
		biggestDiff = -1
		for index := 0; index < parVars.total; index++ {
			if d1notTaken == parVars.partition[index] {
				curRect := &parVars.branchBuf[index].rect
				rect0 := d1combineRect(curRect, &parVars.cover[0])
				rect1 := d1combineRect(curRect, &parVars.cover[1])
				growth0 := d1calcRectVolume(&rect0) - parVars.area[0]
				growth1 := d1calcRectVolume(&rect1) - parVars.area[1]
				diff := growth1 - growth0
				if diff >= 0 {
					group = 0
				} else {
					group = 1
					diff = -diff
				}

				if diff > biggestDiff {
					biggestDiff = diff
					chosen = index
					betterGroup = group
				} else if (diff == biggestDiff) && (parVars.count[group] < parVars.count[betterGroup]) {
					chosen = index
					betterGroup = group
				}
			}
		}
		d1classify(chosen, betterGroup, parVars)
	}

	// If one group too full, put remaining rects in the other
	if (parVars.count[0] + parVars.count[1]) < parVars.total {
		if parVars.count[0] >= parVars.total-parVars.minFill {
			group = 1
		} else {
			group = 0
		}
		for index := 0; index < parVars.total; index++ {
			if d1notTaken == parVars.partition[index] {
				d1classify(index, group, parVars)
			}
		}
	}
}

// Copy branches from the buffer into two nodes according to the partition.
func d1loadNodes(nodeA, nodeB *d1nodeT, parVars *d1partitionVarsT) {
	for index := 0; index < parVars.total; index++ {
		targetNodeIndex := parVars.partition[index]
		targetNodes := []*d1nodeT{nodeA, nodeB}

		// It is assured that d1addBranch here will not cause a node split.
		d1addBranch(&parVars.branchBuf[index], targetNodes[targetNodeIndex], nil)
	}
}

// Initialize a d1partitionVarsT structure.
func d1initParVars(parVars *d1partitionVarsT, maxRects, minFill int) {
	parVars.count[0] = 0
	parVars.count[1] = 0
	parVars.area[0] = 0
	parVars.area[1] = 0
	parVars.total = maxRects
	parVars.minFill = minFill
	for index := 0; index < maxRects; index++ {
		parVars.partition[index] = d1notTaken
	}
}

func d1pickSeeds(parVars *d1partitionVarsT) {
	var seed0, seed1 int
	var worst, waste float64
	var area [d1maxNodes + 1]float64

	for index := 0; index < parVars.total; index++ {
		area[index] = d1calcRectVolume(&parVars.branchBuf[index].rect)
	}

	worst = -parVars.coverSplitArea - 1
	for indexA := 0; indexA < parVars.total-1; indexA++ {
		for indexB := indexA + 1; indexB < parVars.total; indexB++ {
			oneRect := d1combineRect(&parVars.branchBuf[indexA].rect, &parVars.branchBuf[indexB].rect)
			waste = d1calcRectVolume(&oneRect) - area[indexA] - area[indexB]
			if waste > worst {
				worst = waste
				seed0 = indexA
				seed1 = indexB
			}
		}
	}

	d1classify(seed0, 0, parVars)
	d1classify(seed1, 1, parVars)
}

// Put a branch in one of the groups.
func d1classify(index, group int, parVars *d1partitionVarsT) {
	parVars.partition[index] = group

	// Calculate combined rect
	if parVars.count[group] == 0 {
		parVars.cover[group] = parVars.branchBuf[index].rect
	} else {
		parVars.cover[group] = d1combineRect(&parVars.branchBuf[index].rect, &parVars.cover[group])
	}

	// Calculate volume of combined rect
	parVars.area[group] = d1calcRectVolume(&parVars.cover[group])

	parVars.count[group]++
}

// Delete a data rectangle from an index structure.
// Pass in a pointer to a d1rectT, the tid of the record, ptr to ptr to root node.
// Returns 1 if record not found, 0 if success.
// d1removeRect provides for eliminating the root.
func d1removeRect(rect *d1rectT, id interface{}, root **d1nodeT) bool {
	var reInsertList *d1listNodeT

	if !d1removeRectRec(rect, id, *root, &reInsertList) {
		// Found and deleted a data item
		// Reinsert any branches from eliminated nodes
		for reInsertList != nil {
			tempNode := reInsertList.node

			for index := 0; index < tempNode.count; index++ {
				// TODO go over this code. should I use (tempNode->m_level - 1)?
				d1insertRect(&tempNode.branch[index], root, tempNode.level)
			}
			reInsertList = reInsertList.next
		}

		// Check for redundant root (not leaf, 1 child) and eliminate TODO replace
		// if with while? In case there is a whole branch of redundant roots...
		if (*root).count == 1 && (*root).isInternalNode() {
			tempNode := (*root).branch[0].child
			*root = tempNode
		}
		return false
	} else {
		return true
	}
}

// Delete a rectangle from non-root part of an index structure.
// Called by d1removeRect.  Descends tree recursively,
// merges branches on the way back up.
// Returns 1 if record not found, 0 if success.
func d1removeRectRec(rect *d1rectT, id interface{}, node *d1nodeT, listNode **d1listNodeT) bool {
	if node.isInternalNode() { // not a leaf node
		for index := 0; index < node.count; index++ {
			if d1overlap(*rect, node.branch[index].rect) {
				if !d1removeRectRec(rect, id, node.branch[index].child, listNode) {
					if node.branch[index].child.count >= d1minNodes {
						// child removed, just resize parent rect
						node.branch[index].rect = d1nodeCover(node.branch[index].child)
					} else {
						// child removed, not enough entries in node, eliminate node
						d1reInsert(node.branch[index].child, listNode)
						d1disconnectBranch(node, index) // Must return after this call as count has changed
					}
					return false
				}
			}
		}
		return true
	} else { // A leaf node
		for index := 0; index < node.count; index++ {
			if node.branch[index].data == id {
				d1disconnectBranch(node, index) // Must return after this call as count has changed
				return false
			}
		}
		return true
	}
}

// Decide whether two rectangles d1overlap.
func d1overlap(rectA, rectB d1rectT) bool {
	for index := 0; index < d1numDims; index++ {
		if rectA.min[index] > rectB.max[index] ||
			rectB.min[index] > rectA.max[index] {
			return false
		}
	}
	return true
}

// Add a node to the reinsertion list.  All its branches will later
// be reinserted into the index structure.
func d1reInsert(node *d1nodeT, listNode **d1listNodeT) {
	newListNode := &d1listNodeT{}
	newListNode.node = node
	newListNode.next = *listNode
	*listNode = newListNode
}

// d1search in an index tree or subtree for all data retangles that d1overlap the argument rectangle.
func d1search(node *d1nodeT, rect d1rectT, foundCount int, resultCallback func(data interface{}) bool) (int, bool) {
	if node.isInternalNode() {
		// This is an internal node in the tree
		for index := 0; index < node.count; index++ {
			if d1overlap(rect, node.branch[index].rect) {
				var ok bool
				foundCount, ok = d1search(node.branch[index].child, rect, foundCount, resultCallback)
				if !ok {
					// The callback indicated to stop searching
					return foundCount, false
				}
			}
		}
	} else {
		// This is a leaf node
		for index := 0; index < node.count; index++ {
			if d1overlap(rect, node.branch[index].rect) {
				id := node.branch[index].data
				foundCount++
				if !resultCallback(id) {
					return foundCount, false // Don't continue searching
				}

			}
		}
	}
	return foundCount, true // Continue searching
}

func d2fmin(a, b float64) float64 {
	if a < b {
		return a
	}
	return b
}
func d2fmax(a, b float64) float64 {
	if a > b {
		return a
	}
	return b
}

const (
	d2numDims            = 2
	d2maxNodes           = 8
	d2minNodes           = d2maxNodes / 2
	d2useSphericalVolume = true // Better split classification, may be slower on some systems
)

var d2unitSphereVolume = []float64{
	0.000000, 2.000000, 3.141593, // Dimension  0,1,2
	4.188790, 4.934802, 5.263789, // Dimension  3,4,5
	5.167713, 4.724766, 4.058712, // Dimension  6,7,8
	3.298509, 2.550164, 1.884104, // Dimension  9,10,11
	1.335263, 0.910629, 0.599265, // Dimension  12,13,14
	0.381443, 0.235331, 0.140981, // Dimension  15,16,17
	0.082146, 0.046622, 0.025807, // Dimension  18,19,20
}[d2numDims]

type d2RTree struct {
	root *d2nodeT ///< Root of tree
}

/// Minimal bounding rectangle (n-dimensional)
type d2rectT struct {
	min [d2numDims]float64 ///< Min dimensions of bounding box
	max [d2numDims]float64 ///< Max dimensions of bounding box
}

/// May be data or may be another subtree
/// The parents level determines this.
/// If the parents level is 0, then this is data
type d2branchT struct {
	rect  d2rectT     ///< Bounds
	child *d2nodeT    ///< Child node
	data  interface{} ///< Data Id or Ptr
}

/// d2nodeT for each branch level
type d2nodeT struct {
	count  int                   ///< Count
	level  int                   ///< Leaf is zero, others positive
	branch [d2maxNodes]d2branchT ///< Branch
}

func (node *d2nodeT) isInternalNode() bool {
	return (node.level > 0) // Not a leaf, but a internal node
}
func (node *d2nodeT) isLeaf() bool {
	return (node.level == 0) // A leaf, contains data
}

/// A link list of nodes for reinsertion after a delete operation
type d2listNodeT struct {
	next *d2listNodeT ///< Next in list
	node *d2nodeT     ///< Node
}

const d2notTaken = -1 // indicates that position

/// Variables for finding a split partition
type d2partitionVarsT struct {
	partition [d2maxNodes + 1]int
	total     int
	minFill   int
	count     [2]int
	cover     [2]d2rectT
	area      [2]float64

	branchBuf      [d2maxNodes + 1]d2branchT
	branchCount    int
	coverSplit     d2rectT
	coverSplitArea float64
}

func d2New() *d2RTree {
	// We only support machine word size simple data type eg. integer index or object pointer.
	// Since we are storing as union with non data branch
	return &d2RTree{
		root: &d2nodeT{},
	}
}

/// Insert entry
/// \param a_min Min of bounding rect
/// \param a_max Max of bounding rect
/// \param a_dataId Positive Id of data.  Maybe zero, but negative numbers not allowed.
func (tr *d2RTree) Insert(min, max [d2numDims]float64, dataId interface{}) {
	var branch d2branchT
	branch.data = dataId
	for axis := 0; axis < d2numDims; axis++ {
		branch.rect.min[axis] = min[axis]
		branch.rect.max[axis] = max[axis]
	}
	d2insertRect(&branch, &tr.root, 0)
}

/// Remove entry
/// \param a_min Min of bounding rect
/// \param a_max Max of bounding rect
/// \param a_dataId Positive Id of data.  Maybe zero, but negative numbers not allowed.
func (tr *d2RTree) Remove(min, max [d2numDims]float64, dataId interface{}) {
	var rect d2rectT
	for axis := 0; axis < d2numDims; axis++ {
		rect.min[axis] = min[axis]
		rect.max[axis] = max[axis]
	}
	d2removeRect(&rect, dataId, &tr.root)
}

/// Find all within d2search rectangle
/// \param a_min Min of d2search bounding rect
/// \param a_max Max of d2search bounding rect
/// \param a_searchResult d2search result array.  Caller should set grow size. Function will reset, not append to array.
/// \param a_resultCallback Callback function to return result.  Callback should return 'true' to continue searching
/// \param a_context User context to pass as parameter to a_resultCallback
/// \return Returns the number of entries found
func (tr *d2RTree) Search(min, max [d2numDims]float64, resultCallback func(data interface{}) bool) int {
	var rect d2rectT
	for axis := 0; axis < d2numDims; axis++ {
		rect.min[axis] = min[axis]
		rect.max[axis] = max[axis]
	}
	foundCount, _ := d2search(tr.root, rect, 0, resultCallback)
	return foundCount
}

/// Count the data elements in this container.  This is slow as no internal counter is maintained.
func (tr *d2RTree) Count() int {
	var count int
	d2countRec(tr.root, &count)
	return count
}

/// Remove all entries from tree
func (tr *d2RTree) RemoveAll() {
	// Delete all existing nodes
	tr.root = &d2nodeT{}
}

func d2countRec(node *d2nodeT, count *int) {
	if node.isInternalNode() { // not a leaf node
		for index := 0; index < node.count; index++ {
			d2countRec(node.branch[index].child, count)
		}
	} else { // A leaf node
		*count += node.count
	}
}

// Inserts a new data rectangle into the index structure.
// Recursively descends tree, propagates splits back up.
// Returns 0 if node was not split.  Old node updated.
// If node was split, returns 1 and sets the pointer pointed to by
// new_node to point to the new node.  Old node updated to become one of two.
// The level argument specifies the number of steps up from the leaf
// level to insert; e.g. a data rectangle goes in at level = 0.
func d2insertRectRec(branch *d2branchT, node *d2nodeT, newNode **d2nodeT, level int) bool {
	// recurse until we reach the correct level for the new record. data records
	// will always be called with a_level == 0 (leaf)
	if node.level > level {
		// Still above level for insertion, go down tree recursively
		var otherNode *d2nodeT
		//var newBranch d2branchT

		// find the optimal branch for this record
		index := d2pickBranch(&branch.rect, node)

		// recursively insert this record into the picked branch
		childWasSplit := d2insertRectRec(branch, node.branch[index].child, &otherNode, level)

		if !childWasSplit {
			// Child was not split. Merge the bounding box of the new record with the
			// existing bounding box
			node.branch[index].rect = d2combineRect(&branch.rect, &(node.branch[index].rect))
			return false
		} else {
			// Child was split. The old branches are now re-partitioned to two nodes
			// so we have to re-calculate the bounding boxes of each node
			node.branch[index].rect = d2nodeCover(node.branch[index].child)
			var newBranch d2branchT
			newBranch.child = otherNode
			newBranch.rect = d2nodeCover(otherNode)

			// The old node is already a child of a_node. Now add the newly-created
			// node to a_node as well. a_node might be split because of that.
			return d2addBranch(&newBranch, node, newNode)
		}
	} else if node.level == level {
		// We have reached level for insertion. Add rect, split if necessary
		return d2addBranch(branch, node, newNode)
	} else {
		// Should never occur
		return false
	}
}

// Insert a data rectangle into an index structure.
// d2insertRect provides for splitting the root;
// returns 1 if root was split, 0 if it was not.
// The level argument specifies the number of steps up from the leaf
// level to insert; e.g. a data rectangle goes in at level = 0.
// InsertRect2 does the recursion.
//
func d2insertRect(branch *d2branchT, root **d2nodeT, level int) bool {
	var newNode *d2nodeT

	if d2insertRectRec(branch, *root, &newNode, level) { // Root split

		// Grow tree taller and new root
		newRoot := &d2nodeT{}
		newRoot.level = (*root).level + 1

		var newBranch d2branchT

		// add old root node as a child of the new root
		newBranch.rect = d2nodeCover(*root)
		newBranch.child = *root
		d2addBranch(&newBranch, newRoot, nil)

		// add the split node as a child of the new root
		newBranch.rect = d2nodeCover(newNode)
		newBranch.child = newNode
		d2addBranch(&newBranch, newRoot, nil)

		// set the new root as the root node
		*root = newRoot

		return true
	}
	return false
}

// Find the smallest rectangle that includes all rectangles in branches of a node.
func d2nodeCover(node *d2nodeT) d2rectT {
	rect := node.branch[0].rect
	for index := 1; index < node.count; index++ {
		rect = d2combineRect(&rect, &(node.branch[index].rect))
	}
	return rect
}

// Add a branch to a node.  Split the node if necessary.
// Returns 0 if node not split.  Old node updated.
// Returns 1 if node split, sets *new_node to address of new node.
// Old node updated, becomes one of two.
func d2addBranch(branch *d2branchT, node *d2nodeT, newNode **d2nodeT) bool {
	if node.count < d2maxNodes { // Split won't be necessary
		node.branch[node.count] = *branch
		node.count++
		return false
	} else {
		d2splitNode(node, branch, newNode)
		return true
	}
}

// Disconnect a dependent node.
// Caller must return (or stop using iteration index) after this as count has changed
func d2disconnectBranch(node *d2nodeT, index int) {
	// Remove element by swapping with the last element to prevent gaps in array
	node.branch[index] = node.branch[node.count-1]
	node.branch[node.count-1].data = nil
	node.branch[node.count-1].child = nil
	node.count--
}

// Pick a branch.  Pick the one that will need the smallest increase
// in area to accomodate the new rectangle.  This will result in the
// least total area for the covering rectangles in the current node.
// In case of a tie, pick the one which was smaller before, to get
// the best resolution when searching.
func d2pickBranch(rect *d2rectT, node *d2nodeT) int {
	var firstTime bool = true
	var increase float64
	var bestIncr float64 = -1
	var area float64
	var bestArea float64
	var best int
	var tempRect d2rectT

	for index := 0; index < node.count; index++ {
		curRect := &node.branch[index].rect
		area = d2calcRectVolume(curRect)
		tempRect = d2combineRect(rect, curRect)
		increase = d2calcRectVolume(&tempRect) - area
		if (increase < bestIncr) || firstTime {
			best = index
			bestArea = area
			bestIncr = increase
			firstTime = false
		} else if (increase == bestIncr) && (area < bestArea) {
			best = index
			bestArea = area
			bestIncr = increase
		}
	}
	return best
}

// Combine two rectangles into larger one containing both
func d2combineRect(rectA, rectB *d2rectT) d2rectT {
	var newRect d2rectT

	for index := 0; index < d2numDims; index++ {
		newRect.min[index] = d2fmin(rectA.min[index], rectB.min[index])
		newRect.max[index] = d2fmax(rectA.max[index], rectB.max[index])
	}

	return newRect
}

// Split a node.
// Divides the nodes branches and the extra one between two nodes.
// Old node is one of the new ones, and one really new one is created.
// Tries more than one method for choosing a partition, uses best result.
func d2splitNode(node *d2nodeT, branch *d2branchT, newNode **d2nodeT) {
	// Could just use local here, but member or external is faster since it is reused
	var localVars d2partitionVarsT
	parVars := &localVars

	// Load all the branches into a buffer, initialize old node
	d2getBranches(node, branch, parVars)

	// Find partition
	d2choosePartition(parVars, d2minNodes)

	// Create a new node to hold (about) half of the branches
	*newNode = &d2nodeT{}
	(*newNode).level = node.level

	// Put branches from buffer into 2 nodes according to the chosen partition
	node.count = 0
	d2loadNodes(node, *newNode, parVars)
}

// Calculate the n-dimensional volume of a rectangle
func d2rectVolume(rect *d2rectT) float64 {
	var volume float64 = 1
	for index := 0; index < d2numDims; index++ {
		volume *= rect.max[index] - rect.min[index]
	}
	return volume
}

// The exact volume of the bounding sphere for the given d2rectT
func d2rectSphericalVolume(rect *d2rectT) float64 {
	var sumOfSquares float64 = 0
	var radius float64

	for index := 0; index < d2numDims; index++ {
		halfExtent := (rect.max[index] - rect.min[index]) * 0.5
		sumOfSquares += halfExtent * halfExtent
	}

	radius = math.Sqrt(sumOfSquares)

	// Pow maybe slow, so test for common dims just use x*x, x*x*x.
	if d2numDims == 5 {
		return (radius * radius * radius * radius * radius * d2unitSphereVolume)
	} else if d2numDims == 4 {
		return (radius * radius * radius * radius * d2unitSphereVolume)
	} else if d2numDims == 3 {
		return (radius * radius * radius * d2unitSphereVolume)
	} else if d2numDims == 2 {
		return (radius * radius * d2unitSphereVolume)
	} else {
		return (math.Pow(radius, d2numDims) * d2unitSphereVolume)
	}
}

// Use one of the methods to calculate retangle volume
func d2calcRectVolume(rect *d2rectT) float64 {
	if d2useSphericalVolume {
		return d2rectSphericalVolume(rect) // Slower but helps certain merge cases
	} else { // RTREE_USE_SPHERICAL_VOLUME
		return d2rectVolume(rect) // Faster but can cause poor merges
	} // RTREE_USE_SPHERICAL_VOLUME
}

// Load branch buffer with branches from full node plus the extra branch.
func d2getBranches(node *d2nodeT, branch *d2branchT, parVars *d2partitionVarsT) {
	// Load the branch buffer
	for index := 0; index < d2maxNodes; index++ {
		parVars.branchBuf[index] = node.branch[index]
	}
	parVars.branchBuf[d2maxNodes] = *branch
	parVars.branchCount = d2maxNodes + 1

	// Calculate rect containing all in the set
	parVars.coverSplit = parVars.branchBuf[0].rect
	for index := 1; index < d2maxNodes+1; index++ {
		parVars.coverSplit = d2combineRect(&parVars.coverSplit, &parVars.branchBuf[index].rect)
	}
	parVars.coverSplitArea = d2calcRectVolume(&parVars.coverSplit)
}

// Method #0 for choosing a partition:
// As the seeds for the two groups, pick the two rects that would waste the
// most area if covered by a single rectangle, i.e. evidently the worst pair
// to have in the same group.
// Of the remaining, one at a time is chosen to be put in one of the two groups.
// The one chosen is the one with the greatest difference in area expansion
// depending on which group - the rect most strongly attracted to one group
// and repelled from the other.
// If one group gets too full (more would force other group to violate min
// fill requirement) then other group gets the rest.
// These last are the ones that can go in either group most easily.
func d2choosePartition(parVars *d2partitionVarsT, minFill int) {
	var biggestDiff float64
	var group, chosen, betterGroup int

	d2initParVars(parVars, parVars.branchCount, minFill)
	d2pickSeeds(parVars)

	for ((parVars.count[0] + parVars.count[1]) < parVars.total) &&
		(parVars.count[0] < (parVars.total - parVars.minFill)) &&
		(parVars.count[1] < (parVars.total - parVars.minFill)) {
		biggestDiff = -1
		for index := 0; index < parVars.total; index++ {
			if d2notTaken == parVars.partition[index] {
				curRect := &parVars.branchBuf[index].rect
				rect0 := d2combineRect(curRect, &parVars.cover[0])
				rect1 := d2combineRect(curRect, &parVars.cover[1])
				growth0 := d2calcRectVolume(&rect0) - parVars.area[0]
				growth1 := d2calcRectVolume(&rect1) - parVars.area[1]
				diff := growth1 - growth0
				if diff >= 0 {
					group = 0
				} else {
					group = 1
					diff = -diff
				}

				if diff > biggestDiff {
					biggestDiff = diff
					chosen = index
					betterGroup = group
				} else if (diff == biggestDiff) && (parVars.count[group] < parVars.count[betterGroup]) {
					chosen = index
					betterGroup = group
				}
			}
		}
		d2classify(chosen, betterGroup, parVars)
	}

	// If one group too full, put remaining rects in the other
	if (parVars.count[0] + parVars.count[1]) < parVars.total {
		if parVars.count[0] >= parVars.total-parVars.minFill {
			group = 1
		} else {
			group = 0
		}
		for index := 0; index < parVars.total; index++ {
			if d2notTaken == parVars.partition[index] {
				d2classify(index, group, parVars)
			}
		}
	}
}

// Copy branches from the buffer into two nodes according to the partition.
func d2loadNodes(nodeA, nodeB *d2nodeT, parVars *d2partitionVarsT) {
	for index := 0; index < parVars.total; index++ {
		targetNodeIndex := parVars.partition[index]
		targetNodes := []*d2nodeT{nodeA, nodeB}

		// It is assured that d2addBranch here will not cause a node split.
		d2addBranch(&parVars.branchBuf[index], targetNodes[targetNodeIndex], nil)
	}
}

// Initialize a d2partitionVarsT structure.
func d2initParVars(parVars *d2partitionVarsT, maxRects, minFill int) {
	parVars.count[0] = 0
	parVars.count[1] = 0
	parVars.area[0] = 0
	parVars.area[1] = 0
	parVars.total = maxRects
	parVars.minFill = minFill
	for index := 0; index < maxRects; index++ {
		parVars.partition[index] = d2notTaken
	}
}

func d2pickSeeds(parVars *d2partitionVarsT) {
	var seed0, seed1 int
	var worst, waste float64
	var area [d2maxNodes + 1]float64

	for index := 0; index < parVars.total; index++ {
		area[index] = d2calcRectVolume(&parVars.branchBuf[index].rect)
	}

	worst = -parVars.coverSplitArea - 1
	for indexA := 0; indexA < parVars.total-1; indexA++ {
		for indexB := indexA + 1; indexB < parVars.total; indexB++ {
			oneRect := d2combineRect(&parVars.branchBuf[indexA].rect, &parVars.branchBuf[indexB].rect)
			waste = d2calcRectVolume(&oneRect) - area[indexA] - area[indexB]
			if waste > worst {
				worst = waste
				seed0 = indexA
				seed1 = indexB
			}
		}
	}

	d2classify(seed0, 0, parVars)
	d2classify(seed1, 1, parVars)
}

// Put a branch in one of the groups.
func d2classify(index, group int, parVars *d2partitionVarsT) {
	parVars.partition[index] = group

	// Calculate combined rect
	if parVars.count[group] == 0 {
		parVars.cover[group] = parVars.branchBuf[index].rect
	} else {
		parVars.cover[group] = d2combineRect(&parVars.branchBuf[index].rect, &parVars.cover[group])
	}

	// Calculate volume of combined rect
	parVars.area[group] = d2calcRectVolume(&parVars.cover[group])

	parVars.count[group]++
}

// Delete a data rectangle from an index structure.
// Pass in a pointer to a d2rectT, the tid of the record, ptr to ptr to root node.
// Returns 1 if record not found, 0 if success.
// d2removeRect provides for eliminating the root.
func d2removeRect(rect *d2rectT, id interface{}, root **d2nodeT) bool {
	var reInsertList *d2listNodeT

	if !d2removeRectRec(rect, id, *root, &reInsertList) {
		// Found and deleted a data item
		// Reinsert any branches from eliminated nodes
		for reInsertList != nil {
			tempNode := reInsertList.node

			for index := 0; index < tempNode.count; index++ {
				// TODO go over this code. should I use (tempNode->m_level - 1)?
				d2insertRect(&tempNode.branch[index], root, tempNode.level)
			}
			reInsertList = reInsertList.next
		}

		// Check for redundant root (not leaf, 1 child) and eliminate TODO replace
		// if with while? In case there is a whole branch of redundant roots...
		if (*root).count == 1 && (*root).isInternalNode() {
			tempNode := (*root).branch[0].child
			*root = tempNode
		}
		return false
	} else {
		return true
	}
}

// Delete a rectangle from non-root part of an index structure.
// Called by d2removeRect.  Descends tree recursively,
// merges branches on the way back up.
// Returns 1 if record not found, 0 if success.
func d2removeRectRec(rect *d2rectT, id interface{}, node *d2nodeT, listNode **d2listNodeT) bool {
	if node.isInternalNode() { // not a leaf node
		for index := 0; index < node.count; index++ {
			if d2overlap(*rect, node.branch[index].rect) {
				if !d2removeRectRec(rect, id, node.branch[index].child, listNode) {
					if node.branch[index].child.count >= d2minNodes {
						// child removed, just resize parent rect
						node.branch[index].rect = d2nodeCover(node.branch[index].child)
					} else {
						// child removed, not enough entries in node, eliminate node
						d2reInsert(node.branch[index].child, listNode)
						d2disconnectBranch(node, index) // Must return after this call as count has changed
					}
					return false
				}
			}
		}
		return true
	} else { // A leaf node
		for index := 0; index < node.count; index++ {
			if node.branch[index].data == id {
				d2disconnectBranch(node, index) // Must return after this call as count has changed
				return false
			}
		}
		return true
	}
}

// Decide whether two rectangles d2overlap.
func d2overlap(rectA, rectB d2rectT) bool {
	for index := 0; index < d2numDims; index++ {
		if rectA.min[index] > rectB.max[index] ||
			rectB.min[index] > rectA.max[index] {
			return false
		}
	}
	return true
}

// Add a node to the reinsertion list.  All its branches will later
// be reinserted into the index structure.
func d2reInsert(node *d2nodeT, listNode **d2listNodeT) {
	newListNode := &d2listNodeT{}
	newListNode.node = node
	newListNode.next = *listNode
	*listNode = newListNode
}

// d2search in an index tree or subtree for all data retangles that d2overlap the argument rectangle.
func d2search(node *d2nodeT, rect d2rectT, foundCount int, resultCallback func(data interface{}) bool) (int, bool) {
	if node.isInternalNode() {
		// This is an internal node in the tree
		for index := 0; index < node.count; index++ {
			if d2overlap(rect, node.branch[index].rect) {
				var ok bool
				foundCount, ok = d2search(node.branch[index].child, rect, foundCount, resultCallback)
				if !ok {
					// The callback indicated to stop searching
					return foundCount, false
				}
			}
		}
	} else {
		// This is a leaf node
		for index := 0; index < node.count; index++ {
			if d2overlap(rect, node.branch[index].rect) {
				id := node.branch[index].data
				foundCount++
				if !resultCallback(id) {
					return foundCount, false // Don't continue searching
				}

			}
		}
	}
	return foundCount, true // Continue searching
}

func d3fmin(a, b float64) float64 {
	if a < b {
		return a
	}
	return b
}
func d3fmax(a, b float64) float64 {
	if a > b {
		return a
	}
	return b
}

const (
	d3numDims            = 3
	d3maxNodes           = 8
	d3minNodes           = d3maxNodes / 2
	d3useSphericalVolume = true // Better split classification, may be slower on some systems
)

var d3unitSphereVolume = []float64{
	0.000000, 2.000000, 3.141593, // Dimension  0,1,2
	4.188790, 4.934802, 5.263789, // Dimension  3,4,5
	5.167713, 4.724766, 4.058712, // Dimension  6,7,8
	3.298509, 2.550164, 1.884104, // Dimension  9,10,11
	1.335263, 0.910629, 0.599265, // Dimension  12,13,14
	0.381443, 0.235331, 0.140981, // Dimension  15,16,17
	0.082146, 0.046622, 0.025807, // Dimension  18,19,20
}[d3numDims]

type d3RTree struct {
	root *d3nodeT ///< Root of tree
}

/// Minimal bounding rectangle (n-dimensional)
type d3rectT struct {
	min [d3numDims]float64 ///< Min dimensions of bounding box
	max [d3numDims]float64 ///< Max dimensions of bounding box
}

/// May be data or may be another subtree
/// The parents level determines this.
/// If the parents level is 0, then this is data
type d3branchT struct {
	rect  d3rectT     ///< Bounds
	child *d3nodeT    ///< Child node
	data  interface{} ///< Data Id or Ptr
}

/// d3nodeT for each branch level
type d3nodeT struct {
	count  int                   ///< Count
	level  int                   ///< Leaf is zero, others positive
	branch [d3maxNodes]d3branchT ///< Branch
}

func (node *d3nodeT) isInternalNode() bool {
	return (node.level > 0) // Not a leaf, but a internal node
}
func (node *d3nodeT) isLeaf() bool {
	return (node.level == 0) // A leaf, contains data
}

/// A link list of nodes for reinsertion after a delete operation
type d3listNodeT struct {
	next *d3listNodeT ///< Next in list
	node *d3nodeT     ///< Node
}

const d3notTaken = -1 // indicates that position

/// Variables for finding a split partition
type d3partitionVarsT struct {
	partition [d3maxNodes + 1]int
	total     int
	minFill   int
	count     [2]int
	cover     [2]d3rectT
	area      [2]float64

	branchBuf      [d3maxNodes + 1]d3branchT
	branchCount    int
	coverSplit     d3rectT
	coverSplitArea float64
}

func d3New() *d3RTree {
	// We only support machine word size simple data type eg. integer index or object pointer.
	// Since we are storing as union with non data branch
	return &d3RTree{
		root: &d3nodeT{},
	}
}

/// Insert entry
/// \param a_min Min of bounding rect
/// \param a_max Max of bounding rect
/// \param a_dataId Positive Id of data.  Maybe zero, but negative numbers not allowed.
func (tr *d3RTree) Insert(min, max [d3numDims]float64, dataId interface{}) {
	var branch d3branchT
	branch.data = dataId
	for axis := 0; axis < d3numDims; axis++ {
		branch.rect.min[axis] = min[axis]
		branch.rect.max[axis] = max[axis]
	}
	d3insertRect(&branch, &tr.root, 0)
}

/// Remove entry
/// \param a_min Min of bounding rect
/// \param a_max Max of bounding rect
/// \param a_dataId Positive Id of data.  Maybe zero, but negative numbers not allowed.
func (tr *d3RTree) Remove(min, max [d3numDims]float64, dataId interface{}) {
	var rect d3rectT
	for axis := 0; axis < d3numDims; axis++ {
		rect.min[axis] = min[axis]
		rect.max[axis] = max[axis]
	}
	d3removeRect(&rect, dataId, &tr.root)
}

/// Find all within d3search rectangle
/// \param a_min Min of d3search bounding rect
/// \param a_max Max of d3search bounding rect
/// \param a_searchResult d3search result array.  Caller should set grow size. Function will reset, not append to array.
/// \param a_resultCallback Callback function to return result.  Callback should return 'true' to continue searching
/// \param a_context User context to pass as parameter to a_resultCallback
/// \return Returns the number of entries found
func (tr *d3RTree) Search(min, max [d3numDims]float64, resultCallback func(data interface{}) bool) int {
	var rect d3rectT
	for axis := 0; axis < d3numDims; axis++ {
		rect.min[axis] = min[axis]
		rect.max[axis] = max[axis]
	}
	foundCount, _ := d3search(tr.root, rect, 0, resultCallback)
	return foundCount
}

/// Count the data elements in this container.  This is slow as no internal counter is maintained.
func (tr *d3RTree) Count() int {
	var count int
	d3countRec(tr.root, &count)
	return count
}

/// Remove all entries from tree
func (tr *d3RTree) RemoveAll() {
	// Delete all existing nodes
	tr.root = &d3nodeT{}
}

func d3countRec(node *d3nodeT, count *int) {
	if node.isInternalNode() { // not a leaf node
		for index := 0; index < node.count; index++ {
			d3countRec(node.branch[index].child, count)
		}
	} else { // A leaf node
		*count += node.count
	}
}

// Inserts a new data rectangle into the index structure.
// Recursively descends tree, propagates splits back up.
// Returns 0 if node was not split.  Old node updated.
// If node was split, returns 1 and sets the pointer pointed to by
// new_node to point to the new node.  Old node updated to become one of two.
// The level argument specifies the number of steps up from the leaf
// level to insert; e.g. a data rectangle goes in at level = 0.
func d3insertRectRec(branch *d3branchT, node *d3nodeT, newNode **d3nodeT, level int) bool {
	// recurse until we reach the correct level for the new record. data records
	// will always be called with a_level == 0 (leaf)
	if node.level > level {
		// Still above level for insertion, go down tree recursively
		var otherNode *d3nodeT
		//var newBranch d3branchT

		// find the optimal branch for this record
		index := d3pickBranch(&branch.rect, node)

		// recursively insert this record into the picked branch
		childWasSplit := d3insertRectRec(branch, node.branch[index].child, &otherNode, level)

		if !childWasSplit {
			// Child was not split. Merge the bounding box of the new record with the
			// existing bounding box
			node.branch[index].rect = d3combineRect(&branch.rect, &(node.branch[index].rect))
			return false
		} else {
			// Child was split. The old branches are now re-partitioned to two nodes
			// so we have to re-calculate the bounding boxes of each node
			node.branch[index].rect = d3nodeCover(node.branch[index].child)
			var newBranch d3branchT
			newBranch.child = otherNode
			newBranch.rect = d3nodeCover(otherNode)

			// The old node is already a child of a_node. Now add the newly-created
			// node to a_node as well. a_node might be split because of that.
			return d3addBranch(&newBranch, node, newNode)
		}
	} else if node.level == level {
		// We have reached level for insertion. Add rect, split if necessary
		return d3addBranch(branch, node, newNode)
	} else {
		// Should never occur
		return false
	}
}

// Insert a data rectangle into an index structure.
// d3insertRect provides for splitting the root;
// returns 1 if root was split, 0 if it was not.
// The level argument specifies the number of steps up from the leaf
// level to insert; e.g. a data rectangle goes in at level = 0.
// InsertRect2 does the recursion.
//
func d3insertRect(branch *d3branchT, root **d3nodeT, level int) bool {
	var newNode *d3nodeT

	if d3insertRectRec(branch, *root, &newNode, level) { // Root split

		// Grow tree taller and new root
		newRoot := &d3nodeT{}
		newRoot.level = (*root).level + 1

		var newBranch d3branchT

		// add old root node as a child of the new root
		newBranch.rect = d3nodeCover(*root)
		newBranch.child = *root
		d3addBranch(&newBranch, newRoot, nil)

		// add the split node as a child of the new root
		newBranch.rect = d3nodeCover(newNode)
		newBranch.child = newNode
		d3addBranch(&newBranch, newRoot, nil)

		// set the new root as the root node
		*root = newRoot

		return true
	}
	return false
}

// Find the smallest rectangle that includes all rectangles in branches of a node.
func d3nodeCover(node *d3nodeT) d3rectT {
	rect := node.branch[0].rect
	for index := 1; index < node.count; index++ {
		rect = d3combineRect(&rect, &(node.branch[index].rect))
	}
	return rect
}

// Add a branch to a node.  Split the node if necessary.
// Returns 0 if node not split.  Old node updated.
// Returns 1 if node split, sets *new_node to address of new node.
// Old node updated, becomes one of two.
func d3addBranch(branch *d3branchT, node *d3nodeT, newNode **d3nodeT) bool {
	if node.count < d3maxNodes { // Split won't be necessary
		node.branch[node.count] = *branch
		node.count++
		return false
	} else {
		d3splitNode(node, branch, newNode)
		return true
	}
}

// Disconnect a dependent node.
// Caller must return (or stop using iteration index) after this as count has changed
func d3disconnectBranch(node *d3nodeT, index int) {
	// Remove element by swapping with the last element to prevent gaps in array
	node.branch[index] = node.branch[node.count-1]
	node.branch[node.count-1].data = nil
	node.branch[node.count-1].child = nil
	node.count--
}

// Pick a branch.  Pick the one that will need the smallest increase
// in area to accomodate the new rectangle.  This will result in the
// least total area for the covering rectangles in the current node.
// In case of a tie, pick the one which was smaller before, to get
// the best resolution when searching.
func d3pickBranch(rect *d3rectT, node *d3nodeT) int {
	var firstTime bool = true
	var increase float64
	var bestIncr float64 = -1
	var area float64
	var bestArea float64
	var best int
	var tempRect d3rectT

	for index := 0; index < node.count; index++ {
		curRect := &node.branch[index].rect
		area = d3calcRectVolume(curRect)
		tempRect = d3combineRect(rect, curRect)
		increase = d3calcRectVolume(&tempRect) - area
		if (increase < bestIncr) || firstTime {
			best = index
			bestArea = area
			bestIncr = increase
			firstTime = false
		} else if (increase == bestIncr) && (area < bestArea) {
			best = index
			bestArea = area
			bestIncr = increase
		}
	}
	return best
}

// Combine two rectangles into larger one containing both
func d3combineRect(rectA, rectB *d3rectT) d3rectT {
	var newRect d3rectT

	for index := 0; index < d3numDims; index++ {
		newRect.min[index] = d3fmin(rectA.min[index], rectB.min[index])
		newRect.max[index] = d3fmax(rectA.max[index], rectB.max[index])
	}

	return newRect
}

// Split a node.
// Divides the nodes branches and the extra one between two nodes.
// Old node is one of the new ones, and one really new one is created.
// Tries more than one method for choosing a partition, uses best result.
func d3splitNode(node *d3nodeT, branch *d3branchT, newNode **d3nodeT) {
	// Could just use local here, but member or external is faster since it is reused
	var localVars d3partitionVarsT
	parVars := &localVars

	// Load all the branches into a buffer, initialize old node
	d3getBranches(node, branch, parVars)

	// Find partition
	d3choosePartition(parVars, d3minNodes)

	// Create a new node to hold (about) half of the branches
	*newNode = &d3nodeT{}
	(*newNode).level = node.level

	// Put branches from buffer into 2 nodes according to the chosen partition
	node.count = 0
	d3loadNodes(node, *newNode, parVars)
}

// Calculate the n-dimensional volume of a rectangle
func d3rectVolume(rect *d3rectT) float64 {
	var volume float64 = 1
	for index := 0; index < d3numDims; index++ {
		volume *= rect.max[index] - rect.min[index]
	}
	return volume
}

// The exact volume of the bounding sphere for the given d3rectT
func d3rectSphericalVolume(rect *d3rectT) float64 {
	var sumOfSquares float64 = 0
	var radius float64

	for index := 0; index < d3numDims; index++ {
		halfExtent := (rect.max[index] - rect.min[index]) * 0.5
		sumOfSquares += halfExtent * halfExtent
	}

	radius = math.Sqrt(sumOfSquares)

	// Pow maybe slow, so test for common dims just use x*x, x*x*x.
	if d3numDims == 5 {
		return (radius * radius * radius * radius * radius * d3unitSphereVolume)
	} else if d3numDims == 4 {
		return (radius * radius * radius * radius * d3unitSphereVolume)
	} else if d3numDims == 3 {
		return (radius * radius * radius * d3unitSphereVolume)
	} else if d3numDims == 2 {
		return (radius * radius * d3unitSphereVolume)
	} else {
		return (math.Pow(radius, d3numDims) * d3unitSphereVolume)
	}
}

// Use one of the methods to calculate retangle volume
func d3calcRectVolume(rect *d3rectT) float64 {
	if d3useSphericalVolume {
		return d3rectSphericalVolume(rect) // Slower but helps certain merge cases
	} else { // RTREE_USE_SPHERICAL_VOLUME
		return d3rectVolume(rect) // Faster but can cause poor merges
	} // RTREE_USE_SPHERICAL_VOLUME
}

// Load branch buffer with branches from full node plus the extra branch.
func d3getBranches(node *d3nodeT, branch *d3branchT, parVars *d3partitionVarsT) {
	// Load the branch buffer
	for index := 0; index < d3maxNodes; index++ {
		parVars.branchBuf[index] = node.branch[index]
	}
	parVars.branchBuf[d3maxNodes] = *branch
	parVars.branchCount = d3maxNodes + 1

	// Calculate rect containing all in the set
	parVars.coverSplit = parVars.branchBuf[0].rect
	for index := 1; index < d3maxNodes+1; index++ {
		parVars.coverSplit = d3combineRect(&parVars.coverSplit, &parVars.branchBuf[index].rect)
	}
	parVars.coverSplitArea = d3calcRectVolume(&parVars.coverSplit)
}

// Method #0 for choosing a partition:
// As the seeds for the two groups, pick the two rects that would waste the
// most area if covered by a single rectangle, i.e. evidently the worst pair
// to have in the same group.
// Of the remaining, one at a time is chosen to be put in one of the two groups.
// The one chosen is the one with the greatest difference in area expansion
// depending on which group - the rect most strongly attracted to one group
// and repelled from the other.
// If one group gets too full (more would force other group to violate min
// fill requirement) then other group gets the rest.
// These last are the ones that can go in either group most easily.
func d3choosePartition(parVars *d3partitionVarsT, minFill int) {
	var biggestDiff float64
	var group, chosen, betterGroup int

	d3initParVars(parVars, parVars.branchCount, minFill)
	d3pickSeeds(parVars)

	for ((parVars.count[0] + parVars.count[1]) < parVars.total) &&
		(parVars.count[0] < (parVars.total - parVars.minFill)) &&
		(parVars.count[1] < (parVars.total - parVars.minFill)) {
		biggestDiff = -1
		for index := 0; index < parVars.total; index++ {
			if d3notTaken == parVars.partition[index] {
				curRect := &parVars.branchBuf[index].rect
				rect0 := d3combineRect(curRect, &parVars.cover[0])
				rect1 := d3combineRect(curRect, &parVars.cover[1])
				growth0 := d3calcRectVolume(&rect0) - parVars.area[0]
				growth1 := d3calcRectVolume(&rect1) - parVars.area[1]
				diff := growth1 - growth0
				if diff >= 0 {
					group = 0
				} else {
					group = 1
					diff = -diff
				}

				if diff > biggestDiff {
					biggestDiff = diff
					chosen = index
					betterGroup = group
				} else if (diff == biggestDiff) && (parVars.count[group] < parVars.count[betterGroup]) {
					chosen = index
					betterGroup = group
				}
			}
		}
		d3classify(chosen, betterGroup, parVars)
	}

	// If one group too full, put remaining rects in the other
	if (parVars.count[0] + parVars.count[1]) < parVars.total {
		if parVars.count[0] >= parVars.total-parVars.minFill {
			group = 1
		} else {
			group = 0
		}
		for index := 0; index < parVars.total; index++ {
			if d3notTaken == parVars.partition[index] {
				d3classify(index, group, parVars)
			}
		}
	}
}

// Copy branches from the buffer into two nodes according to the partition.
func d3loadNodes(nodeA, nodeB *d3nodeT, parVars *d3partitionVarsT) {
	for index := 0; index < parVars.total; index++ {
		targetNodeIndex := parVars.partition[index]
		targetNodes := []*d3nodeT{nodeA, nodeB}

		// It is assured that d3addBranch here will not cause a node split.
		d3addBranch(&parVars.branchBuf[index], targetNodes[targetNodeIndex], nil)
	}
}

// Initialize a d3partitionVarsT structure.
func d3initParVars(parVars *d3partitionVarsT, maxRects, minFill int) {
	parVars.count[0] = 0
	parVars.count[1] = 0
	parVars.area[0] = 0
	parVars.area[1] = 0
	parVars.total = maxRects
	parVars.minFill = minFill
	for index := 0; index < maxRects; index++ {
		parVars.partition[index] = d3notTaken
	}
}

func d3pickSeeds(parVars *d3partitionVarsT) {
	var seed0, seed1 int
	var worst, waste float64
	var area [d3maxNodes + 1]float64

	for index := 0; index < parVars.total; index++ {
		area[index] = d3calcRectVolume(&parVars.branchBuf[index].rect)
	}

	worst = -parVars.coverSplitArea - 1
	for indexA := 0; indexA < parVars.total-1; indexA++ {
		for indexB := indexA + 1; indexB < parVars.total; indexB++ {
			oneRect := d3combineRect(&parVars.branchBuf[indexA].rect, &parVars.branchBuf[indexB].rect)
			waste = d3calcRectVolume(&oneRect) - area[indexA] - area[indexB]
			if waste > worst {
				worst = waste
				seed0 = indexA
				seed1 = indexB
			}
		}
	}

	d3classify(seed0, 0, parVars)
	d3classify(seed1, 1, parVars)
}

// Put a branch in one of the groups.
func d3classify(index, group int, parVars *d3partitionVarsT) {
	parVars.partition[index] = group

	// Calculate combined rect
	if parVars.count[group] == 0 {
		parVars.cover[group] = parVars.branchBuf[index].rect
	} else {
		parVars.cover[group] = d3combineRect(&parVars.branchBuf[index].rect, &parVars.cover[group])
	}

	// Calculate volume of combined rect
	parVars.area[group] = d3calcRectVolume(&parVars.cover[group])

	parVars.count[group]++
}

// Delete a data rectangle from an index structure.
// Pass in a pointer to a d3rectT, the tid of the record, ptr to ptr to root node.
// Returns 1 if record not found, 0 if success.
// d3removeRect provides for eliminating the root.
func d3removeRect(rect *d3rectT, id interface{}, root **d3nodeT) bool {
	var reInsertList *d3listNodeT

	if !d3removeRectRec(rect, id, *root, &reInsertList) {
		// Found and deleted a data item
		// Reinsert any branches from eliminated nodes
		for reInsertList != nil {
			tempNode := reInsertList.node

			for index := 0; index < tempNode.count; index++ {
				// TODO go over this code. should I use (tempNode->m_level - 1)?
				d3insertRect(&tempNode.branch[index], root, tempNode.level)
			}
			reInsertList = reInsertList.next
		}

		// Check for redundant root (not leaf, 1 child) and eliminate TODO replace
		// if with while? In case there is a whole branch of redundant roots...
		if (*root).count == 1 && (*root).isInternalNode() {
			tempNode := (*root).branch[0].child
			*root = tempNode
		}
		return false
	} else {
		return true
	}
}

// Delete a rectangle from non-root part of an index structure.
// Called by d3removeRect.  Descends tree recursively,
// merges branches on the way back up.
// Returns 1 if record not found, 0 if success.
func d3removeRectRec(rect *d3rectT, id interface{}, node *d3nodeT, listNode **d3listNodeT) bool {
	if node.isInternalNode() { // not a leaf node
		for index := 0; index < node.count; index++ {
			if d3overlap(*rect, node.branch[index].rect) {
				if !d3removeRectRec(rect, id, node.branch[index].child, listNode) {
					if node.branch[index].child.count >= d3minNodes {
						// child removed, just resize parent rect
						node.branch[index].rect = d3nodeCover(node.branch[index].child)
					} else {
						// child removed, not enough entries in node, eliminate node
						d3reInsert(node.branch[index].child, listNode)
						d3disconnectBranch(node, index) // Must return after this call as count has changed
					}
					return false
				}
			}
		}
		return true
	} else { // A leaf node
		for index := 0; index < node.count; index++ {
			if node.branch[index].data == id {
				d3disconnectBranch(node, index) // Must return after this call as count has changed
				return false
			}
		}
		return true
	}
}

// Decide whether two rectangles d3overlap.
func d3overlap(rectA, rectB d3rectT) bool {
	for index := 0; index < d3numDims; index++ {
		if rectA.min[index] > rectB.max[index] ||
			rectB.min[index] > rectA.max[index] {
			return false
		}
	}
	return true
}

// Add a node to the reinsertion list.  All its branches will later
// be reinserted into the index structure.
func d3reInsert(node *d3nodeT, listNode **d3listNodeT) {
	newListNode := &d3listNodeT{}
	newListNode.node = node
	newListNode.next = *listNode
	*listNode = newListNode
}

// d3search in an index tree or subtree for all data retangles that d3overlap the argument rectangle.
func d3search(node *d3nodeT, rect d3rectT, foundCount int, resultCallback func(data interface{}) bool) (int, bool) {
	if node.isInternalNode() {
		// This is an internal node in the tree
		for index := 0; index < node.count; index++ {
			if d3overlap(rect, node.branch[index].rect) {
				var ok bool
				foundCount, ok = d3search(node.branch[index].child, rect, foundCount, resultCallback)
				if !ok {
					// The callback indicated to stop searching
					return foundCount, false
				}
			}
		}
	} else {
		// This is a leaf node
		for index := 0; index < node.count; index++ {
			if d3overlap(rect, node.branch[index].rect) {
				id := node.branch[index].data
				foundCount++
				if !resultCallback(id) {
					return foundCount, false // Don't continue searching
				}

			}
		}
	}
	return foundCount, true // Continue searching
}

func d4fmin(a, b float64) float64 {
	if a < b {
		return a
	}
	return b
}
func d4fmax(a, b float64) float64 {
	if a > b {
		return a
	}
	return b
}

const (
	d4numDims            = 4
	d4maxNodes           = 8
	d4minNodes           = d4maxNodes / 2
	d4useSphericalVolume = true // Better split classification, may be slower on some systems
)

var d4unitSphereVolume = []float64{
	0.000000, 2.000000, 3.141593, // Dimension  0,1,2
	4.188790, 4.934802, 5.263789, // Dimension  3,4,5
	5.167713, 4.724766, 4.058712, // Dimension  6,7,8
	3.298509, 2.550164, 1.884104, // Dimension  9,10,11
	1.335263, 0.910629, 0.599265, // Dimension  12,13,14
	0.381443, 0.235331, 0.140981, // Dimension  15,16,17
	0.082146, 0.046622, 0.025807, // Dimension  18,19,20
}[d4numDims]

type d4RTree struct {
	root *d4nodeT ///< Root of tree
}

/// Minimal bounding rectangle (n-dimensional)
type d4rectT struct {
	min [d4numDims]float64 ///< Min dimensions of bounding box
	max [d4numDims]float64 ///< Max dimensions of bounding box
}

/// May be data or may be another subtree
/// The parents level determines this.
/// If the parents level is 0, then this is data
type d4branchT struct {
	rect  d4rectT     ///< Bounds
	child *d4nodeT    ///< Child node
	data  interface{} ///< Data Id or Ptr
}

/// d4nodeT for each branch level
type d4nodeT struct {
	count  int                   ///< Count
	level  int                   ///< Leaf is zero, others positive
	branch [d4maxNodes]d4branchT ///< Branch
}

func (node *d4nodeT) isInternalNode() bool {
	return (node.level > 0) // Not a leaf, but a internal node
}
func (node *d4nodeT) isLeaf() bool {
	return (node.level == 0) // A leaf, contains data
}

/// A link list of nodes for reinsertion after a delete operation
type d4listNodeT struct {
	next *d4listNodeT ///< Next in list
	node *d4nodeT     ///< Node
}

const d4notTaken = -1 // indicates that position

/// Variables for finding a split partition
type d4partitionVarsT struct {
	partition [d4maxNodes + 1]int
	total     int
	minFill   int
	count     [2]int
	cover     [2]d4rectT
	area      [2]float64

	branchBuf      [d4maxNodes + 1]d4branchT
	branchCount    int
	coverSplit     d4rectT
	coverSplitArea float64
}

func d4New() *d4RTree {
	// We only support machine word size simple data type eg. integer index or object pointer.
	// Since we are storing as union with non data branch
	return &d4RTree{
		root: &d4nodeT{},
	}
}

/// Insert entry
/// \param a_min Min of bounding rect
/// \param a_max Max of bounding rect
/// \param a_dataId Positive Id of data.  Maybe zero, but negative numbers not allowed.
func (tr *d4RTree) Insert(min, max [d4numDims]float64, dataId interface{}) {
	var branch d4branchT
	branch.data = dataId
	for axis := 0; axis < d4numDims; axis++ {
		branch.rect.min[axis] = min[axis]
		branch.rect.max[axis] = max[axis]
	}
	d4insertRect(&branch, &tr.root, 0)
}

/// Remove entry
/// \param a_min Min of bounding rect
/// \param a_max Max of bounding rect
/// \param a_dataId Positive Id of data.  Maybe zero, but negative numbers not allowed.
func (tr *d4RTree) Remove(min, max [d4numDims]float64, dataId interface{}) {
	var rect d4rectT
	for axis := 0; axis < d4numDims; axis++ {
		rect.min[axis] = min[axis]
		rect.max[axis] = max[axis]
	}
	d4removeRect(&rect, dataId, &tr.root)
}

/// Find all within d4search rectangle
/// \param a_min Min of d4search bounding rect
/// \param a_max Max of d4search bounding rect
/// \param a_searchResult d4search result array.  Caller should set grow size. Function will reset, not append to array.
/// \param a_resultCallback Callback function to return result.  Callback should return 'true' to continue searching
/// \param a_context User context to pass as parameter to a_resultCallback
/// \return Returns the number of entries found
func (tr *d4RTree) Search(min, max [d4numDims]float64, resultCallback func(data interface{}) bool) int {
	var rect d4rectT
	for axis := 0; axis < d4numDims; axis++ {
		rect.min[axis] = min[axis]
		rect.max[axis] = max[axis]
	}
	foundCount, _ := d4search(tr.root, rect, 0, resultCallback)
	return foundCount
}

/// Count the data elements in this container.  This is slow as no internal counter is maintained.
func (tr *d4RTree) Count() int {
	var count int
	d4countRec(tr.root, &count)
	return count
}

/// Remove all entries from tree
func (tr *d4RTree) RemoveAll() {
	// Delete all existing nodes
	tr.root = &d4nodeT{}
}

func d4countRec(node *d4nodeT, count *int) {
	if node.isInternalNode() { // not a leaf node
		for index := 0; index < node.count; index++ {
			d4countRec(node.branch[index].child, count)
		}
	} else { // A leaf node
		*count += node.count
	}
}

// Inserts a new data rectangle into the index structure.
// Recursively descends tree, propagates splits back up.
// Returns 0 if node was not split.  Old node updated.
// If node was split, returns 1 and sets the pointer pointed to by
// new_node to point to the new node.  Old node updated to become one of two.
// The level argument specifies the number of steps up from the leaf
// level to insert; e.g. a data rectangle goes in at level = 0.
func d4insertRectRec(branch *d4branchT, node *d4nodeT, newNode **d4nodeT, level int) bool {
	// recurse until we reach the correct level for the new record. data records
	// will always be called with a_level == 0 (leaf)
	if node.level > level {
		// Still above level for insertion, go down tree recursively
		var otherNode *d4nodeT
		//var newBranch d4branchT

		// find the optimal branch for this record
		index := d4pickBranch(&branch.rect, node)

		// recursively insert this record into the picked branch
		childWasSplit := d4insertRectRec(branch, node.branch[index].child, &otherNode, level)

		if !childWasSplit {
			// Child was not split. Merge the bounding box of the new record with the
			// existing bounding box
			node.branch[index].rect = d4combineRect(&branch.rect, &(node.branch[index].rect))
			return false
		} else {
			// Child was split. The old branches are now re-partitioned to two nodes
			// so we have to re-calculate the bounding boxes of each node
			node.branch[index].rect = d4nodeCover(node.branch[index].child)
			var newBranch d4branchT
			newBranch.child = otherNode
			newBranch.rect = d4nodeCover(otherNode)

			// The old node is already a child of a_node. Now add the newly-created
			// node to a_node as well. a_node might be split because of that.
			return d4addBranch(&newBranch, node, newNode)
		}
	} else if node.level == level {
		// We have reached level for insertion. Add rect, split if necessary
		return d4addBranch(branch, node, newNode)
	} else {
		// Should never occur
		return false
	}
}

// Insert a data rectangle into an index structure.
// d4insertRect provides for splitting the root;
// returns 1 if root was split, 0 if it was not.
// The level argument specifies the number of steps up from the leaf
// level to insert; e.g. a data rectangle goes in at level = 0.
// InsertRect2 does the recursion.
//
func d4insertRect(branch *d4branchT, root **d4nodeT, level int) bool {
	var newNode *d4nodeT

	if d4insertRectRec(branch, *root, &newNode, level) { // Root split

		// Grow tree taller and new root
		newRoot := &d4nodeT{}
		newRoot.level = (*root).level + 1

		var newBranch d4branchT

		// add old root node as a child of the new root
		newBranch.rect = d4nodeCover(*root)
		newBranch.child = *root
		d4addBranch(&newBranch, newRoot, nil)

		// add the split node as a child of the new root
		newBranch.rect = d4nodeCover(newNode)
		newBranch.child = newNode
		d4addBranch(&newBranch, newRoot, nil)

		// set the new root as the root node
		*root = newRoot

		return true
	}
	return false
}

// Find the smallest rectangle that includes all rectangles in branches of a node.
func d4nodeCover(node *d4nodeT) d4rectT {
	rect := node.branch[0].rect
	for index := 1; index < node.count; index++ {
		rect = d4combineRect(&rect, &(node.branch[index].rect))
	}
	return rect
}

// Add a branch to a node.  Split the node if necessary.
// Returns 0 if node not split.  Old node updated.
// Returns 1 if node split, sets *new_node to address of new node.
// Old node updated, becomes one of two.
func d4addBranch(branch *d4branchT, node *d4nodeT, newNode **d4nodeT) bool {
	if node.count < d4maxNodes { // Split won't be necessary
		node.branch[node.count] = *branch
		node.count++
		return false
	} else {
		d4splitNode(node, branch, newNode)
		return true
	}
}

// Disconnect a dependent node.
// Caller must return (or stop using iteration index) after this as count has changed
func d4disconnectBranch(node *d4nodeT, index int) {
	// Remove element by swapping with the last element to prevent gaps in array
	node.branch[index] = node.branch[node.count-1]
	node.branch[node.count-1].data = nil
	node.branch[node.count-1].child = nil
	node.count--
}

// Pick a branch.  Pick the one that will need the smallest increase
// in area to accomodate the new rectangle.  This will result in the
// least total area for the covering rectangles in the current node.
// In case of a tie, pick the one which was smaller before, to get
// the best resolution when searching.
func d4pickBranch(rect *d4rectT, node *d4nodeT) int {
	var firstTime bool = true
	var increase float64
	var bestIncr float64 = -1
	var area float64
	var bestArea float64
	var best int
	var tempRect d4rectT

	for index := 0; index < node.count; index++ {
		curRect := &node.branch[index].rect
		area = d4calcRectVolume(curRect)
		tempRect = d4combineRect(rect, curRect)
		increase = d4calcRectVolume(&tempRect) - area
		if (increase < bestIncr) || firstTime {
			best = index
			bestArea = area
			bestIncr = increase
			firstTime = false
		} else if (increase == bestIncr) && (area < bestArea) {
			best = index
			bestArea = area
			bestIncr = increase
		}
	}
	return best
}

// Combine two rectangles into larger one containing both
func d4combineRect(rectA, rectB *d4rectT) d4rectT {
	var newRect d4rectT

	for index := 0; index < d4numDims; index++ {
		newRect.min[index] = d4fmin(rectA.min[index], rectB.min[index])
		newRect.max[index] = d4fmax(rectA.max[index], rectB.max[index])
	}

	return newRect
}

// Split a node.
// Divides the nodes branches and the extra one between two nodes.
// Old node is one of the new ones, and one really new one is created.
// Tries more than one method for choosing a partition, uses best result.
func d4splitNode(node *d4nodeT, branch *d4branchT, newNode **d4nodeT) {
	// Could just use local here, but member or external is faster since it is reused
	var localVars d4partitionVarsT
	parVars := &localVars

	// Load all the branches into a buffer, initialize old node
	d4getBranches(node, branch, parVars)

	// Find partition
	d4choosePartition(parVars, d4minNodes)

	// Create a new node to hold (about) half of the branches
	*newNode = &d4nodeT{}
	(*newNode).level = node.level

	// Put branches from buffer into 2 nodes according to the chosen partition
	node.count = 0
	d4loadNodes(node, *newNode, parVars)
}

// Calculate the n-dimensional volume of a rectangle
func d4rectVolume(rect *d4rectT) float64 {
	var volume float64 = 1
	for index := 0; index < d4numDims; index++ {
		volume *= rect.max[index] - rect.min[index]
	}
	return volume
}

// The exact volume of the bounding sphere for the given d4rectT
func d4rectSphericalVolume(rect *d4rectT) float64 {
	var sumOfSquares float64 = 0
	var radius float64

	for index := 0; index < d4numDims; index++ {
		halfExtent := (rect.max[index] - rect.min[index]) * 0.5
		sumOfSquares += halfExtent * halfExtent
	}

	radius = math.Sqrt(sumOfSquares)

	// Pow maybe slow, so test for common dims just use x*x, x*x*x.
	if d4numDims == 5 {
		return (radius * radius * radius * radius * radius * d4unitSphereVolume)
	} else if d4numDims == 4 {
		return (radius * radius * radius * radius * d4unitSphereVolume)
	} else if d4numDims == 3 {
		return (radius * radius * radius * d4unitSphereVolume)
	} else if d4numDims == 2 {
		return (radius * radius * d4unitSphereVolume)
	} else {
		return (math.Pow(radius, d4numDims) * d4unitSphereVolume)
	}
}

// Use one of the methods to calculate retangle volume
func d4calcRectVolume(rect *d4rectT) float64 {
	if d4useSphericalVolume {
		return d4rectSphericalVolume(rect) // Slower but helps certain merge cases
	} else { // RTREE_USE_SPHERICAL_VOLUME
		return d4rectVolume(rect) // Faster but can cause poor merges
	} // RTREE_USE_SPHERICAL_VOLUME
}

// Load branch buffer with branches from full node plus the extra branch.
func d4getBranches(node *d4nodeT, branch *d4branchT, parVars *d4partitionVarsT) {
	// Load the branch buffer
	for index := 0; index < d4maxNodes; index++ {
		parVars.branchBuf[index] = node.branch[index]
	}
	parVars.branchBuf[d4maxNodes] = *branch
	parVars.branchCount = d4maxNodes + 1

	// Calculate rect containing all in the set
	parVars.coverSplit = parVars.branchBuf[0].rect
	for index := 1; index < d4maxNodes+1; index++ {
		parVars.coverSplit = d4combineRect(&parVars.coverSplit, &parVars.branchBuf[index].rect)
	}
	parVars.coverSplitArea = d4calcRectVolume(&parVars.coverSplit)
}

// Method #0 for choosing a partition:
// As the seeds for the two groups, pick the two rects that would waste the
// most area if covered by a single rectangle, i.e. evidently the worst pair
// to have in the same group.
// Of the remaining, one at a time is chosen to be put in one of the two groups.
// The one chosen is the one with the greatest difference in area expansion
// depending on which group - the rect most strongly attracted to one group
// and repelled from the other.
// If one group gets too full (more would force other group to violate min
// fill requirement) then other group gets the rest.
// These last are the ones that can go in either group most easily.
func d4choosePartition(parVars *d4partitionVarsT, minFill int) {
	var biggestDiff float64
	var group, chosen, betterGroup int

	d4initParVars(parVars, parVars.branchCount, minFill)
	d4pickSeeds(parVars)

	for ((parVars.count[0] + parVars.count[1]) < parVars.total) &&
		(parVars.count[0] < (parVars.total - parVars.minFill)) &&
		(parVars.count[1] < (parVars.total - parVars.minFill)) {
		biggestDiff = -1
		for index := 0; index < parVars.total; index++ {
			if d4notTaken == parVars.partition[index] {
				curRect := &parVars.branchBuf[index].rect
				rect0 := d4combineRect(curRect, &parVars.cover[0])
				rect1 := d4combineRect(curRect, &parVars.cover[1])
				growth0 := d4calcRectVolume(&rect0) - parVars.area[0]
				growth1 := d4calcRectVolume(&rect1) - parVars.area[1]
				diff := growth1 - growth0
				if diff >= 0 {
					group = 0
				} else {
					group = 1
					diff = -diff
				}

				if diff > biggestDiff {
					biggestDiff = diff
					chosen = index
					betterGroup = group
				} else if (diff == biggestDiff) && (parVars.count[group] < parVars.count[betterGroup]) {
					chosen = index
					betterGroup = group
				}
			}
		}
		d4classify(chosen, betterGroup, parVars)
	}

	// If one group too full, put remaining rects in the other
	if (parVars.count[0] + parVars.count[1]) < parVars.total {
		if parVars.count[0] >= parVars.total-parVars.minFill {
			group = 1
		} else {
			group = 0
		}
		for index := 0; index < parVars.total; index++ {
			if d4notTaken == parVars.partition[index] {
				d4classify(index, group, parVars)
			}
		}
	}
}

// Copy branches from the buffer into two nodes according to the partition.
func d4loadNodes(nodeA, nodeB *d4nodeT, parVars *d4partitionVarsT) {
	for index := 0; index < parVars.total; index++ {
		targetNodeIndex := parVars.partition[index]
		targetNodes := []*d4nodeT{nodeA, nodeB}

		// It is assured that d4addBranch here will not cause a node split.
		d4addBranch(&parVars.branchBuf[index], targetNodes[targetNodeIndex], nil)
	}
}

// Initialize a d4partitionVarsT structure.
func d4initParVars(parVars *d4partitionVarsT, maxRects, minFill int) {
	parVars.count[0] = 0
	parVars.count[1] = 0
	parVars.area[0] = 0
	parVars.area[1] = 0
	parVars.total = maxRects
	parVars.minFill = minFill
	for index := 0; index < maxRects; index++ {
		parVars.partition[index] = d4notTaken
	}
}

func d4pickSeeds(parVars *d4partitionVarsT) {
	var seed0, seed1 int
	var worst, waste float64
	var area [d4maxNodes + 1]float64

	for index := 0; index < parVars.total; index++ {
		area[index] = d4calcRectVolume(&parVars.branchBuf[index].rect)
	}

	worst = -parVars.coverSplitArea - 1
	for indexA := 0; indexA < parVars.total-1; indexA++ {
		for indexB := indexA + 1; indexB < parVars.total; indexB++ {
			oneRect := d4combineRect(&parVars.branchBuf[indexA].rect, &parVars.branchBuf[indexB].rect)
			waste = d4calcRectVolume(&oneRect) - area[indexA] - area[indexB]
			if waste > worst {
				worst = waste
				seed0 = indexA
				seed1 = indexB
			}
		}
	}

	d4classify(seed0, 0, parVars)
	d4classify(seed1, 1, parVars)
}

// Put a branch in one of the groups.
func d4classify(index, group int, parVars *d4partitionVarsT) {
	parVars.partition[index] = group

	// Calculate combined rect
	if parVars.count[group] == 0 {
		parVars.cover[group] = parVars.branchBuf[index].rect
	} else {
		parVars.cover[group] = d4combineRect(&parVars.branchBuf[index].rect, &parVars.cover[group])
	}

	// Calculate volume of combined rect
	parVars.area[group] = d4calcRectVolume(&parVars.cover[group])

	parVars.count[group]++
}

// Delete a data rectangle from an index structure.
// Pass in a pointer to a d4rectT, the tid of the record, ptr to ptr to root node.
// Returns 1 if record not found, 0 if success.
// d4removeRect provides for eliminating the root.
func d4removeRect(rect *d4rectT, id interface{}, root **d4nodeT) bool {
	var reInsertList *d4listNodeT

	if !d4removeRectRec(rect, id, *root, &reInsertList) {
		// Found and deleted a data item
		// Reinsert any branches from eliminated nodes
		for reInsertList != nil {
			tempNode := reInsertList.node

			for index := 0; index < tempNode.count; index++ {
				// TODO go over this code. should I use (tempNode->m_level - 1)?
				d4insertRect(&tempNode.branch[index], root, tempNode.level)
			}
			reInsertList = reInsertList.next
		}

		// Check for redundant root (not leaf, 1 child) and eliminate TODO replace
		// if with while? In case there is a whole branch of redundant roots...
		if (*root).count == 1 && (*root).isInternalNode() {
			tempNode := (*root).branch[0].child
			*root = tempNode
		}
		return false
	} else {
		return true
	}
}

// Delete a rectangle from non-root part of an index structure.
// Called by d4removeRect.  Descends tree recursively,
// merges branches on the way back up.
// Returns 1 if record not found, 0 if success.
func d4removeRectRec(rect *d4rectT, id interface{}, node *d4nodeT, listNode **d4listNodeT) bool {
	if node.isInternalNode() { // not a leaf node
		for index := 0; index < node.count; index++ {
			if d4overlap(*rect, node.branch[index].rect) {
				if !d4removeRectRec(rect, id, node.branch[index].child, listNode) {
					if node.branch[index].child.count >= d4minNodes {
						// child removed, just resize parent rect
						node.branch[index].rect = d4nodeCover(node.branch[index].child)
					} else {
						// child removed, not enough entries in node, eliminate node
						d4reInsert(node.branch[index].child, listNode)
						d4disconnectBranch(node, index) // Must return after this call as count has changed
					}
					return false
				}
			}
		}
		return true
	} else { // A leaf node
		for index := 0; index < node.count; index++ {
			if node.branch[index].data == id {
				d4disconnectBranch(node, index) // Must return after this call as count has changed
				return false
			}
		}
		return true
	}
}

// Decide whether two rectangles d4overlap.
func d4overlap(rectA, rectB d4rectT) bool {
	for index := 0; index < d4numDims; index++ {
		if rectA.min[index] > rectB.max[index] ||
			rectB.min[index] > rectA.max[index] {
			return false
		}
	}
	return true
}

// Add a node to the reinsertion list.  All its branches will later
// be reinserted into the index structure.
func d4reInsert(node *d4nodeT, listNode **d4listNodeT) {
	newListNode := &d4listNodeT{}
	newListNode.node = node
	newListNode.next = *listNode
	*listNode = newListNode
}

// d4search in an index tree or subtree for all data retangles that d4overlap the argument rectangle.
func d4search(node *d4nodeT, rect d4rectT, foundCount int, resultCallback func(data interface{}) bool) (int, bool) {
	if node.isInternalNode() {
		// This is an internal node in the tree
		for index := 0; index < node.count; index++ {
			if d4overlap(rect, node.branch[index].rect) {
				var ok bool
				foundCount, ok = d4search(node.branch[index].child, rect, foundCount, resultCallback)
				if !ok {
					// The callback indicated to stop searching
					return foundCount, false
				}
			}
		}
	} else {
		// This is a leaf node
		for index := 0; index < node.count; index++ {
			if d4overlap(rect, node.branch[index].rect) {
				id := node.branch[index].data
				foundCount++
				if !resultCallback(id) {
					return foundCount, false // Don't continue searching
				}

			}
		}
	}
	return foundCount, true // Continue searching
}

func d5fmin(a, b float64) float64 {
	if a < b {
		return a
	}
	return b
}
func d5fmax(a, b float64) float64 {
	if a > b {
		return a
	}
	return b
}

const (
	d5numDims            = 5
	d5maxNodes           = 8
	d5minNodes           = d5maxNodes / 2
	d5useSphericalVolume = true // Better split classification, may be slower on some systems
)

var d5unitSphereVolume = []float64{
	0.000000, 2.000000, 3.141593, // Dimension  0,1,2
	4.188790, 4.934802, 5.263789, // Dimension  3,4,5
	5.167713, 4.724766, 4.058712, // Dimension  6,7,8
	3.298509, 2.550164, 1.884104, // Dimension  9,10,11
	1.335263, 0.910629, 0.599265, // Dimension  12,13,14
	0.381443, 0.235331, 0.140981, // Dimension  15,16,17
	0.082146, 0.046622, 0.025807, // Dimension  18,19,20
}[d5numDims]

type d5RTree struct {
	root *d5nodeT ///< Root of tree
}

/// Minimal bounding rectangle (n-dimensional)
type d5rectT struct {
	min [d5numDims]float64 ///< Min dimensions of bounding box
	max [d5numDims]float64 ///< Max dimensions of bounding box
}

/// May be data or may be another subtree
/// The parents level determines this.
/// If the parents level is 0, then this is data
type d5branchT struct {
	rect  d5rectT     ///< Bounds
	child *d5nodeT    ///< Child node
	data  interface{} ///< Data Id or Ptr
}

/// d5nodeT for each branch level
type d5nodeT struct {
	count  int                   ///< Count
	level  int                   ///< Leaf is zero, others positive
	branch [d5maxNodes]d5branchT ///< Branch
}

func (node *d5nodeT) isInternalNode() bool {
	return (node.level > 0) // Not a leaf, but a internal node
}
func (node *d5nodeT) isLeaf() bool {
	return (node.level == 0) // A leaf, contains data
}

/// A link list of nodes for reinsertion after a delete operation
type d5listNodeT struct {
	next *d5listNodeT ///< Next in list
	node *d5nodeT     ///< Node
}

const d5notTaken = -1 // indicates that position

/// Variables for finding a split partition
type d5partitionVarsT struct {
	partition [d5maxNodes + 1]int
	total     int
	minFill   int
	count     [2]int
	cover     [2]d5rectT
	area      [2]float64

	branchBuf      [d5maxNodes + 1]d5branchT
	branchCount    int
	coverSplit     d5rectT
	coverSplitArea float64
}

func d5New() *d5RTree {
	// We only support machine word size simple data type eg. integer index or object pointer.
	// Since we are storing as union with non data branch
	return &d5RTree{
		root: &d5nodeT{},
	}
}

/// Insert entry
/// \param a_min Min of bounding rect
/// \param a_max Max of bounding rect
/// \param a_dataId Positive Id of data.  Maybe zero, but negative numbers not allowed.
func (tr *d5RTree) Insert(min, max [d5numDims]float64, dataId interface{}) {
	var branch d5branchT
	branch.data = dataId
	for axis := 0; axis < d5numDims; axis++ {
		branch.rect.min[axis] = min[axis]
		branch.rect.max[axis] = max[axis]
	}
	d5insertRect(&branch, &tr.root, 0)
}

/// Remove entry
/// \param a_min Min of bounding rect
/// \param a_max Max of bounding rect
/// \param a_dataId Positive Id of data.  Maybe zero, but negative numbers not allowed.
func (tr *d5RTree) Remove(min, max [d5numDims]float64, dataId interface{}) {
	var rect d5rectT
	for axis := 0; axis < d5numDims; axis++ {
		rect.min[axis] = min[axis]
		rect.max[axis] = max[axis]
	}
	d5removeRect(&rect, dataId, &tr.root)
}

/// Find all within d5search rectangle
/// \param a_min Min of d5search bounding rect
/// \param a_max Max of d5search bounding rect
/// \param a_searchResult d5search result array.  Caller should set grow size. Function will reset, not append to array.
/// \param a_resultCallback Callback function to return result.  Callback should return 'true' to continue searching
/// \param a_context User context to pass as parameter to a_resultCallback
/// \return Returns the number of entries found
func (tr *d5RTree) Search(min, max [d5numDims]float64, resultCallback func(data interface{}) bool) int {
	var rect d5rectT
	for axis := 0; axis < d5numDims; axis++ {
		rect.min[axis] = min[axis]
		rect.max[axis] = max[axis]
	}
	foundCount, _ := d5search(tr.root, rect, 0, resultCallback)
	return foundCount
}

/// Count the data elements in this container.  This is slow as no internal counter is maintained.
func (tr *d5RTree) Count() int {
	var count int
	d5countRec(tr.root, &count)
	return count
}

/// Remove all entries from tree
func (tr *d5RTree) RemoveAll() {
	// Delete all existing nodes
	tr.root = &d5nodeT{}
}

func d5countRec(node *d5nodeT, count *int) {
	if node.isInternalNode() { // not a leaf node
		for index := 0; index < node.count; index++ {
			d5countRec(node.branch[index].child, count)
		}
	} else { // A leaf node
		*count += node.count
	}
}

// Inserts a new data rectangle into the index structure.
// Recursively descends tree, propagates splits back up.
// Returns 0 if node was not split.  Old node updated.
// If node was split, returns 1 and sets the pointer pointed to by
// new_node to point to the new node.  Old node updated to become one of two.
// The level argument specifies the number of steps up from the leaf
// level to insert; e.g. a data rectangle goes in at level = 0.
func d5insertRectRec(branch *d5branchT, node *d5nodeT, newNode **d5nodeT, level int) bool {
	// recurse until we reach the correct level for the new record. data records
	// will always be called with a_level == 0 (leaf)
	if node.level > level {
		// Still above level for insertion, go down tree recursively
		var otherNode *d5nodeT
		//var newBranch d5branchT

		// find the optimal branch for this record
		index := d5pickBranch(&branch.rect, node)

		// recursively insert this record into the picked branch
		childWasSplit := d5insertRectRec(branch, node.branch[index].child, &otherNode, level)

		if !childWasSplit {
			// Child was not split. Merge the bounding box of the new record with the
			// existing bounding box
			node.branch[index].rect = d5combineRect(&branch.rect, &(node.branch[index].rect))
			return false
		} else {
			// Child was split. The old branches are now re-partitioned to two nodes
			// so we have to re-calculate the bounding boxes of each node
			node.branch[index].rect = d5nodeCover(node.branch[index].child)
			var newBranch d5branchT
			newBranch.child = otherNode
			newBranch.rect = d5nodeCover(otherNode)

			// The old node is already a child of a_node. Now add the newly-created
			// node to a_node as well. a_node might be split because of that.
			return d5addBranch(&newBranch, node, newNode)
		}
	} else if node.level == level {
		// We have reached level for insertion. Add rect, split if necessary
		return d5addBranch(branch, node, newNode)
	} else {
		// Should never occur
		return false
	}
}

// Insert a data rectangle into an index structure.
// d5insertRect provides for splitting the root;
// returns 1 if root was split, 0 if it was not.
// The level argument specifies the number of steps up from the leaf
// level to insert; e.g. a data rectangle goes in at level = 0.
// InsertRect2 does the recursion.
//
func d5insertRect(branch *d5branchT, root **d5nodeT, level int) bool {
	var newNode *d5nodeT

	if d5insertRectRec(branch, *root, &newNode, level) { // Root split

		// Grow tree taller and new root
		newRoot := &d5nodeT{}
		newRoot.level = (*root).level + 1

		var newBranch d5branchT

		// add old root node as a child of the new root
		newBranch.rect = d5nodeCover(*root)
		newBranch.child = *root
		d5addBranch(&newBranch, newRoot, nil)

		// add the split node as a child of the new root
		newBranch.rect = d5nodeCover(newNode)
		newBranch.child = newNode
		d5addBranch(&newBranch, newRoot, nil)

		// set the new root as the root node
		*root = newRoot

		return true
	}
	return false
}

// Find the smallest rectangle that includes all rectangles in branches of a node.
func d5nodeCover(node *d5nodeT) d5rectT {
	rect := node.branch[0].rect
	for index := 1; index < node.count; index++ {
		rect = d5combineRect(&rect, &(node.branch[index].rect))
	}
	return rect
}

// Add a branch to a node.  Split the node if necessary.
// Returns 0 if node not split.  Old node updated.
// Returns 1 if node split, sets *new_node to address of new node.
// Old node updated, becomes one of two.
func d5addBranch(branch *d5branchT, node *d5nodeT, newNode **d5nodeT) bool {
	if node.count < d5maxNodes { // Split won't be necessary
		node.branch[node.count] = *branch
		node.count++
		return false
	} else {
		d5splitNode(node, branch, newNode)
		return true
	}
}

// Disconnect a dependent node.
// Caller must return (or stop using iteration index) after this as count has changed
func d5disconnectBranch(node *d5nodeT, index int) {
	// Remove element by swapping with the last element to prevent gaps in array
	node.branch[index] = node.branch[node.count-1]
	node.branch[node.count-1].data = nil
	node.branch[node.count-1].child = nil
	node.count--
}

// Pick a branch.  Pick the one that will need the smallest increase
// in area to accomodate the new rectangle.  This will result in the
// least total area for the covering rectangles in the current node.
// In case of a tie, pick the one which was smaller before, to get
// the best resolution when searching.
func d5pickBranch(rect *d5rectT, node *d5nodeT) int {
	var firstTime bool = true
	var increase float64
	var bestIncr float64 = -1
	var area float64
	var bestArea float64
	var best int
	var tempRect d5rectT

	for index := 0; index < node.count; index++ {
		curRect := &node.branch[index].rect
		area = d5calcRectVolume(curRect)
		tempRect = d5combineRect(rect, curRect)
		increase = d5calcRectVolume(&tempRect) - area
		if (increase < bestIncr) || firstTime {
			best = index
			bestArea = area
			bestIncr = increase
			firstTime = false
		} else if (increase == bestIncr) && (area < bestArea) {
			best = index
			bestArea = area
			bestIncr = increase
		}
	}
	return best
}

// Combine two rectangles into larger one containing both
func d5combineRect(rectA, rectB *d5rectT) d5rectT {
	var newRect d5rectT

	for index := 0; index < d5numDims; index++ {
		newRect.min[index] = d5fmin(rectA.min[index], rectB.min[index])
		newRect.max[index] = d5fmax(rectA.max[index], rectB.max[index])
	}

	return newRect
}

// Split a node.
// Divides the nodes branches and the extra one between two nodes.
// Old node is one of the new ones, and one really new one is created.
// Tries more than one method for choosing a partition, uses best result.
func d5splitNode(node *d5nodeT, branch *d5branchT, newNode **d5nodeT) {
	// Could just use local here, but member or external is faster since it is reused
	var localVars d5partitionVarsT
	parVars := &localVars

	// Load all the branches into a buffer, initialize old node
	d5getBranches(node, branch, parVars)

	// Find partition
	d5choosePartition(parVars, d5minNodes)

	// Create a new node to hold (about) half of the branches
	*newNode = &d5nodeT{}
	(*newNode).level = node.level

	// Put branches from buffer into 2 nodes according to the chosen partition
	node.count = 0
	d5loadNodes(node, *newNode, parVars)
}

// Calculate the n-dimensional volume of a rectangle
func d5rectVolume(rect *d5rectT) float64 {
	var volume float64 = 1
	for index := 0; index < d5numDims; index++ {
		volume *= rect.max[index] - rect.min[index]
	}
	return volume
}

// The exact volume of the bounding sphere for the given d5rectT
func d5rectSphericalVolume(rect *d5rectT) float64 {
	var sumOfSquares float64 = 0
	var radius float64

	for index := 0; index < d5numDims; index++ {
		halfExtent := (rect.max[index] - rect.min[index]) * 0.5
		sumOfSquares += halfExtent * halfExtent
	}

	radius = math.Sqrt(sumOfSquares)

	// Pow maybe slow, so test for common dims just use x*x, x*x*x.
	if d5numDims == 5 {
		return (radius * radius * radius * radius * radius * d5unitSphereVolume)
	} else if d5numDims == 4 {
		return (radius * radius * radius * radius * d5unitSphereVolume)
	} else if d5numDims == 3 {
		return (radius * radius * radius * d5unitSphereVolume)
	} else if d5numDims == 2 {
		return (radius * radius * d5unitSphereVolume)
	} else {
		return (math.Pow(radius, d5numDims) * d5unitSphereVolume)
	}
}

// Use one of the methods to calculate retangle volume
func d5calcRectVolume(rect *d5rectT) float64 {
	if d5useSphericalVolume {
		return d5rectSphericalVolume(rect) // Slower but helps certain merge cases
	} else { // RTREE_USE_SPHERICAL_VOLUME
		return d5rectVolume(rect) // Faster but can cause poor merges
	} // RTREE_USE_SPHERICAL_VOLUME
}

// Load branch buffer with branches from full node plus the extra branch.
func d5getBranches(node *d5nodeT, branch *d5branchT, parVars *d5partitionVarsT) {
	// Load the branch buffer
	for index := 0; index < d5maxNodes; index++ {
		parVars.branchBuf[index] = node.branch[index]
	}
	parVars.branchBuf[d5maxNodes] = *branch
	parVars.branchCount = d5maxNodes + 1

	// Calculate rect containing all in the set
	parVars.coverSplit = parVars.branchBuf[0].rect
	for index := 1; index < d5maxNodes+1; index++ {
		parVars.coverSplit = d5combineRect(&parVars.coverSplit, &parVars.branchBuf[index].rect)
	}
	parVars.coverSplitArea = d5calcRectVolume(&parVars.coverSplit)
}

// Method #0 for choosing a partition:
// As the seeds for the two groups, pick the two rects that would waste the
// most area if covered by a single rectangle, i.e. evidently the worst pair
// to have in the same group.
// Of the remaining, one at a time is chosen to be put in one of the two groups.
// The one chosen is the one with the greatest difference in area expansion
// depending on which group - the rect most strongly attracted to one group
// and repelled from the other.
// If one group gets too full (more would force other group to violate min
// fill requirement) then other group gets the rest.
// These last are the ones that can go in either group most easily.
func d5choosePartition(parVars *d5partitionVarsT, minFill int) {
	var biggestDiff float64
	var group, chosen, betterGroup int

	d5initParVars(parVars, parVars.branchCount, minFill)
	d5pickSeeds(parVars)

	for ((parVars.count[0] + parVars.count[1]) < parVars.total) &&
		(parVars.count[0] < (parVars.total - parVars.minFill)) &&
		(parVars.count[1] < (parVars.total - parVars.minFill)) {
		biggestDiff = -1
		for index := 0; index < parVars.total; index++ {
			if d5notTaken == parVars.partition[index] {
				curRect := &parVars.branchBuf[index].rect
				rect0 := d5combineRect(curRect, &parVars.cover[0])
				rect1 := d5combineRect(curRect, &parVars.cover[1])
				growth0 := d5calcRectVolume(&rect0) - parVars.area[0]
				growth1 := d5calcRectVolume(&rect1) - parVars.area[1]
				diff := growth1 - growth0
				if diff >= 0 {
					group = 0
				} else {
					group = 1
					diff = -diff
				}

				if diff > biggestDiff {
					biggestDiff = diff
					chosen = index
					betterGroup = group
				} else if (diff == biggestDiff) && (parVars.count[group] < parVars.count[betterGroup]) {
					chosen = index
					betterGroup = group
				}
			}
		}
		d5classify(chosen, betterGroup, parVars)
	}

	// If one group too full, put remaining rects in the other
	if (parVars.count[0] + parVars.count[1]) < parVars.total {
		if parVars.count[0] >= parVars.total-parVars.minFill {
			group = 1
		} else {
			group = 0
		}
		for index := 0; index < parVars.total; index++ {
			if d5notTaken == parVars.partition[index] {
				d5classify(index, group, parVars)
			}
		}
	}
}

// Copy branches from the buffer into two nodes according to the partition.
func d5loadNodes(nodeA, nodeB *d5nodeT, parVars *d5partitionVarsT) {
	for index := 0; index < parVars.total; index++ {
		targetNodeIndex := parVars.partition[index]
		targetNodes := []*d5nodeT{nodeA, nodeB}

		// It is assured that d5addBranch here will not cause a node split.
		d5addBranch(&parVars.branchBuf[index], targetNodes[targetNodeIndex], nil)
	}
}

// Initialize a d5partitionVarsT structure.
func d5initParVars(parVars *d5partitionVarsT, maxRects, minFill int) {
	parVars.count[0] = 0
	parVars.count[1] = 0
	parVars.area[0] = 0
	parVars.area[1] = 0
	parVars.total = maxRects
	parVars.minFill = minFill
	for index := 0; index < maxRects; index++ {
		parVars.partition[index] = d5notTaken
	}
}

func d5pickSeeds(parVars *d5partitionVarsT) {
	var seed0, seed1 int
	var worst, waste float64
	var area [d5maxNodes + 1]float64

	for index := 0; index < parVars.total; index++ {
		area[index] = d5calcRectVolume(&parVars.branchBuf[index].rect)
	}

	worst = -parVars.coverSplitArea - 1
	for indexA := 0; indexA < parVars.total-1; indexA++ {
		for indexB := indexA + 1; indexB < parVars.total; indexB++ {
			oneRect := d5combineRect(&parVars.branchBuf[indexA].rect, &parVars.branchBuf[indexB].rect)
			waste = d5calcRectVolume(&oneRect) - area[indexA] - area[indexB]
			if waste > worst {
				worst = waste
				seed0 = indexA
				seed1 = indexB
			}
		}
	}

	d5classify(seed0, 0, parVars)
	d5classify(seed1, 1, parVars)
}

// Put a branch in one of the groups.
func d5classify(index, group int, parVars *d5partitionVarsT) {
	parVars.partition[index] = group

	// Calculate combined rect
	if parVars.count[group] == 0 {
		parVars.cover[group] = parVars.branchBuf[index].rect
	} else {
		parVars.cover[group] = d5combineRect(&parVars.branchBuf[index].rect, &parVars.cover[group])
	}

	// Calculate volume of combined rect
	parVars.area[group] = d5calcRectVolume(&parVars.cover[group])

	parVars.count[group]++
}

// Delete a data rectangle from an index structure.
// Pass in a pointer to a d5rectT, the tid of the record, ptr to ptr to root node.
// Returns 1 if record not found, 0 if success.
// d5removeRect provides for eliminating the root.
func d5removeRect(rect *d5rectT, id interface{}, root **d5nodeT) bool {
	var reInsertList *d5listNodeT

	if !d5removeRectRec(rect, id, *root, &reInsertList) {
		// Found and deleted a data item
		// Reinsert any branches from eliminated nodes
		for reInsertList != nil {
			tempNode := reInsertList.node

			for index := 0; index < tempNode.count; index++ {
				// TODO go over this code. should I use (tempNode->m_level - 1)?
				d5insertRect(&tempNode.branch[index], root, tempNode.level)
			}
			reInsertList = reInsertList.next
		}

		// Check for redundant root (not leaf, 1 child) and eliminate TODO replace
		// if with while? In case there is a whole branch of redundant roots...
		if (*root).count == 1 && (*root).isInternalNode() {
			tempNode := (*root).branch[0].child
			*root = tempNode
		}
		return false
	} else {
		return true
	}
}

// Delete a rectangle from non-root part of an index structure.
// Called by d5removeRect.  Descends tree recursively,
// merges branches on the way back up.
// Returns 1 if record not found, 0 if success.
func d5removeRectRec(rect *d5rectT, id interface{}, node *d5nodeT, listNode **d5listNodeT) bool {
	if node.isInternalNode() { // not a leaf node
		for index := 0; index < node.count; index++ {
			if d5overlap(*rect, node.branch[index].rect) {
				if !d5removeRectRec(rect, id, node.branch[index].child, listNode) {
					if node.branch[index].child.count >= d5minNodes {
						// child removed, just resize parent rect
						node.branch[index].rect = d5nodeCover(node.branch[index].child)
					} else {
						// child removed, not enough entries in node, eliminate node
						d5reInsert(node.branch[index].child, listNode)
						d5disconnectBranch(node, index) // Must return after this call as count has changed
					}
					return false
				}
			}
		}
		return true
	} else { // A leaf node
		for index := 0; index < node.count; index++ {
			if node.branch[index].data == id {
				d5disconnectBranch(node, index) // Must return after this call as count has changed
				return false
			}
		}
		return true
	}
}

// Decide whether two rectangles d5overlap.
func d5overlap(rectA, rectB d5rectT) bool {
	for index := 0; index < d5numDims; index++ {
		if rectA.min[index] > rectB.max[index] ||
			rectB.min[index] > rectA.max[index] {
			return false
		}
	}
	return true
}

// Add a node to the reinsertion list.  All its branches will later
// be reinserted into the index structure.
func d5reInsert(node *d5nodeT, listNode **d5listNodeT) {
	newListNode := &d5listNodeT{}
	newListNode.node = node
	newListNode.next = *listNode
	*listNode = newListNode
}

// d5search in an index tree or subtree for all data retangles that d5overlap the argument rectangle.
func d5search(node *d5nodeT, rect d5rectT, foundCount int, resultCallback func(data interface{}) bool) (int, bool) {
	if node.isInternalNode() {
		// This is an internal node in the tree
		for index := 0; index < node.count; index++ {
			if d5overlap(rect, node.branch[index].rect) {
				var ok bool
				foundCount, ok = d5search(node.branch[index].child, rect, foundCount, resultCallback)
				if !ok {
					// The callback indicated to stop searching
					return foundCount, false
				}
			}
		}
	} else {
		// This is a leaf node
		for index := 0; index < node.count; index++ {
			if d5overlap(rect, node.branch[index].rect) {
				id := node.branch[index].data
				foundCount++
				if !resultCallback(id) {
					return foundCount, false // Don't continue searching
				}

			}
		}
	}
	return foundCount, true // Continue searching
}

func d6fmin(a, b float64) float64 {
	if a < b {
		return a
	}
	return b
}
func d6fmax(a, b float64) float64 {
	if a > b {
		return a
	}
	return b
}

const (
	d6numDims            = 6
	d6maxNodes           = 8
	d6minNodes           = d6maxNodes / 2
	d6useSphericalVolume = true // Better split classification, may be slower on some systems
)

var d6unitSphereVolume = []float64{
	0.000000, 2.000000, 3.141593, // Dimension  0,1,2
	4.188790, 4.934802, 5.263789, // Dimension  3,4,5
	5.167713, 4.724766, 4.058712, // Dimension  6,7,8
	3.298509, 2.550164, 1.884104, // Dimension  9,10,11
	1.335263, 0.910629, 0.599265, // Dimension  12,13,14
	0.381443, 0.235331, 0.140981, // Dimension  15,16,17
	0.082146, 0.046622, 0.025807, // Dimension  18,19,20
}[d6numDims]

type d6RTree struct {
	root *d6nodeT ///< Root of tree
}

/// Minimal bounding rectangle (n-dimensional)
type d6rectT struct {
	min [d6numDims]float64 ///< Min dimensions of bounding box
	max [d6numDims]float64 ///< Max dimensions of bounding box
}

/// May be data or may be another subtree
/// The parents level determines this.
/// If the parents level is 0, then this is data
type d6branchT struct {
	rect  d6rectT     ///< Bounds
	child *d6nodeT    ///< Child node
	data  interface{} ///< Data Id or Ptr
}

/// d6nodeT for each branch level
type d6nodeT struct {
	count  int                   ///< Count
	level  int                   ///< Leaf is zero, others positive
	branch [d6maxNodes]d6branchT ///< Branch
}

func (node *d6nodeT) isInternalNode() bool {
	return (node.level > 0) // Not a leaf, but a internal node
}
func (node *d6nodeT) isLeaf() bool {
	return (node.level == 0) // A leaf, contains data
}

/// A link list of nodes for reinsertion after a delete operation
type d6listNodeT struct {
	next *d6listNodeT ///< Next in list
	node *d6nodeT     ///< Node
}

const d6notTaken = -1 // indicates that position

/// Variables for finding a split partition
type d6partitionVarsT struct {
	partition [d6maxNodes + 1]int
	total     int
	minFill   int
	count     [2]int
	cover     [2]d6rectT
	area      [2]float64

	branchBuf      [d6maxNodes + 1]d6branchT
	branchCount    int
	coverSplit     d6rectT
	coverSplitArea float64
}

func d6New() *d6RTree {
	// We only support machine word size simple data type eg. integer index or object pointer.
	// Since we are storing as union with non data branch
	return &d6RTree{
		root: &d6nodeT{},
	}
}

/// Insert entry
/// \param a_min Min of bounding rect
/// \param a_max Max of bounding rect
/// \param a_dataId Positive Id of data.  Maybe zero, but negative numbers not allowed.
func (tr *d6RTree) Insert(min, max [d6numDims]float64, dataId interface{}) {
	var branch d6branchT
	branch.data = dataId
	for axis := 0; axis < d6numDims; axis++ {
		branch.rect.min[axis] = min[axis]
		branch.rect.max[axis] = max[axis]
	}
	d6insertRect(&branch, &tr.root, 0)
}

/// Remove entry
/// \param a_min Min of bounding rect
/// \param a_max Max of bounding rect
/// \param a_dataId Positive Id of data.  Maybe zero, but negative numbers not allowed.
func (tr *d6RTree) Remove(min, max [d6numDims]float64, dataId interface{}) {
	var rect d6rectT
	for axis := 0; axis < d6numDims; axis++ {
		rect.min[axis] = min[axis]
		rect.max[axis] = max[axis]
	}
	d6removeRect(&rect, dataId, &tr.root)
}

/// Find all within d6search rectangle
/// \param a_min Min of d6search bounding rect
/// \param a_max Max of d6search bounding rect
/// \param a_searchResult d6search result array.  Caller should set grow size. Function will reset, not append to array.
/// \param a_resultCallback Callback function to return result.  Callback should return 'true' to continue searching
/// \param a_context User context to pass as parameter to a_resultCallback
/// \return Returns the number of entries found
func (tr *d6RTree) Search(min, max [d6numDims]float64, resultCallback func(data interface{}) bool) int {
	var rect d6rectT
	for axis := 0; axis < d6numDims; axis++ {
		rect.min[axis] = min[axis]
		rect.max[axis] = max[axis]
	}
	foundCount, _ := d6search(tr.root, rect, 0, resultCallback)
	return foundCount
}

/// Count the data elements in this container.  This is slow as no internal counter is maintained.
func (tr *d6RTree) Count() int {
	var count int
	d6countRec(tr.root, &count)
	return count
}

/// Remove all entries from tree
func (tr *d6RTree) RemoveAll() {
	// Delete all existing nodes
	tr.root = &d6nodeT{}
}

func d6countRec(node *d6nodeT, count *int) {
	if node.isInternalNode() { // not a leaf node
		for index := 0; index < node.count; index++ {
			d6countRec(node.branch[index].child, count)
		}
	} else { // A leaf node
		*count += node.count
	}
}

// Inserts a new data rectangle into the index structure.
// Recursively descends tree, propagates splits back up.
// Returns 0 if node was not split.  Old node updated.
// If node was split, returns 1 and sets the pointer pointed to by
// new_node to point to the new node.  Old node updated to become one of two.
// The level argument specifies the number of steps up from the leaf
// level to insert; e.g. a data rectangle goes in at level = 0.
func d6insertRectRec(branch *d6branchT, node *d6nodeT, newNode **d6nodeT, level int) bool {
	// recurse until we reach the correct level for the new record. data records
	// will always be called with a_level == 0 (leaf)
	if node.level > level {
		// Still above level for insertion, go down tree recursively
		var otherNode *d6nodeT
		//var newBranch d6branchT

		// find the optimal branch for this record
		index := d6pickBranch(&branch.rect, node)

		// recursively insert this record into the picked branch
		childWasSplit := d6insertRectRec(branch, node.branch[index].child, &otherNode, level)

		if !childWasSplit {
			// Child was not split. Merge the bounding box of the new record with the
			// existing bounding box
			node.branch[index].rect = d6combineRect(&branch.rect, &(node.branch[index].rect))
			return false
		} else {
			// Child was split. The old branches are now re-partitioned to two nodes
			// so we have to re-calculate the bounding boxes of each node
			node.branch[index].rect = d6nodeCover(node.branch[index].child)
			var newBranch d6branchT
			newBranch.child = otherNode
			newBranch.rect = d6nodeCover(otherNode)

			// The old node is already a child of a_node. Now add the newly-created
			// node to a_node as well. a_node might be split because of that.
			return d6addBranch(&newBranch, node, newNode)
		}
	} else if node.level == level {
		// We have reached level for insertion. Add rect, split if necessary
		return d6addBranch(branch, node, newNode)
	} else {
		// Should never occur
		return false
	}
}

// Insert a data rectangle into an index structure.
// d6insertRect provides for splitting the root;
// returns 1 if root was split, 0 if it was not.
// The level argument specifies the number of steps up from the leaf
// level to insert; e.g. a data rectangle goes in at level = 0.
// InsertRect2 does the recursion.
//
func d6insertRect(branch *d6branchT, root **d6nodeT, level int) bool {
	var newNode *d6nodeT

	if d6insertRectRec(branch, *root, &newNode, level) { // Root split

		// Grow tree taller and new root
		newRoot := &d6nodeT{}
		newRoot.level = (*root).level + 1

		var newBranch d6branchT

		// add old root node as a child of the new root
		newBranch.rect = d6nodeCover(*root)
		newBranch.child = *root
		d6addBranch(&newBranch, newRoot, nil)

		// add the split node as a child of the new root
		newBranch.rect = d6nodeCover(newNode)
		newBranch.child = newNode
		d6addBranch(&newBranch, newRoot, nil)

		// set the new root as the root node
		*root = newRoot

		return true
	}
	return false
}

// Find the smallest rectangle that includes all rectangles in branches of a node.
func d6nodeCover(node *d6nodeT) d6rectT {
	rect := node.branch[0].rect
	for index := 1; index < node.count; index++ {
		rect = d6combineRect(&rect, &(node.branch[index].rect))
	}
	return rect
}

// Add a branch to a node.  Split the node if necessary.
// Returns 0 if node not split.  Old node updated.
// Returns 1 if node split, sets *new_node to address of new node.
// Old node updated, becomes one of two.
func d6addBranch(branch *d6branchT, node *d6nodeT, newNode **d6nodeT) bool {
	if node.count < d6maxNodes { // Split won't be necessary
		node.branch[node.count] = *branch
		node.count++
		return false
	} else {
		d6splitNode(node, branch, newNode)
		return true
	}
}

// Disconnect a dependent node.
// Caller must return (or stop using iteration index) after this as count has changed
func d6disconnectBranch(node *d6nodeT, index int) {
	// Remove element by swapping with the last element to prevent gaps in array
	node.branch[index] = node.branch[node.count-1]
	node.branch[node.count-1].data = nil
	node.branch[node.count-1].child = nil
	node.count--
}

// Pick a branch.  Pick the one that will need the smallest increase
// in area to accomodate the new rectangle.  This will result in the
// least total area for the covering rectangles in the current node.
// In case of a tie, pick the one which was smaller before, to get
// the best resolution when searching.
func d6pickBranch(rect *d6rectT, node *d6nodeT) int {
	var firstTime bool = true
	var increase float64
	var bestIncr float64 = -1
	var area float64
	var bestArea float64
	var best int
	var tempRect d6rectT

	for index := 0; index < node.count; index++ {
		curRect := &node.branch[index].rect
		area = d6calcRectVolume(curRect)
		tempRect = d6combineRect(rect, curRect)
		increase = d6calcRectVolume(&tempRect) - area
		if (increase < bestIncr) || firstTime {
			best = index
			bestArea = area
			bestIncr = increase
			firstTime = false
		} else if (increase == bestIncr) && (area < bestArea) {
			best = index
			bestArea = area
			bestIncr = increase
		}
	}
	return best
}

// Combine two rectangles into larger one containing both
func d6combineRect(rectA, rectB *d6rectT) d6rectT {
	var newRect d6rectT

	for index := 0; index < d6numDims; index++ {
		newRect.min[index] = d6fmin(rectA.min[index], rectB.min[index])
		newRect.max[index] = d6fmax(rectA.max[index], rectB.max[index])
	}

	return newRect
}

// Split a node.
// Divides the nodes branches and the extra one between two nodes.
// Old node is one of the new ones, and one really new one is created.
// Tries more than one method for choosing a partition, uses best result.
func d6splitNode(node *d6nodeT, branch *d6branchT, newNode **d6nodeT) {
	// Could just use local here, but member or external is faster since it is reused
	var localVars d6partitionVarsT
	parVars := &localVars

	// Load all the branches into a buffer, initialize old node
	d6getBranches(node, branch, parVars)

	// Find partition
	d6choosePartition(parVars, d6minNodes)

	// Create a new node to hold (about) half of the branches
	*newNode = &d6nodeT{}
	(*newNode).level = node.level

	// Put branches from buffer into 2 nodes according to the chosen partition
	node.count = 0
	d6loadNodes(node, *newNode, parVars)
}

// Calculate the n-dimensional volume of a rectangle
func d6rectVolume(rect *d6rectT) float64 {
	var volume float64 = 1
	for index := 0; index < d6numDims; index++ {
		volume *= rect.max[index] - rect.min[index]
	}
	return volume
}

// The exact volume of the bounding sphere for the given d6rectT
func d6rectSphericalVolume(rect *d6rectT) float64 {
	var sumOfSquares float64 = 0
	var radius float64

	for index := 0; index < d6numDims; index++ {
		halfExtent := (rect.max[index] - rect.min[index]) * 0.5
		sumOfSquares += halfExtent * halfExtent
	}

	radius = math.Sqrt(sumOfSquares)

	// Pow maybe slow, so test for common dims just use x*x, x*x*x.
	if d6numDims == 5 {
		return (radius * radius * radius * radius * radius * d6unitSphereVolume)
	} else if d6numDims == 4 {
		return (radius * radius * radius * radius * d6unitSphereVolume)
	} else if d6numDims == 3 {
		return (radius * radius * radius * d6unitSphereVolume)
	} else if d6numDims == 2 {
		return (radius * radius * d6unitSphereVolume)
	} else {
		return (math.Pow(radius, d6numDims) * d6unitSphereVolume)
	}
}

// Use one of the methods to calculate retangle volume
func d6calcRectVolume(rect *d6rectT) float64 {
	if d6useSphericalVolume {
		return d6rectSphericalVolume(rect) // Slower but helps certain merge cases
	} else { // RTREE_USE_SPHERICAL_VOLUME
		return d6rectVolume(rect) // Faster but can cause poor merges
	} // RTREE_USE_SPHERICAL_VOLUME
}

// Load branch buffer with branches from full node plus the extra branch.
func d6getBranches(node *d6nodeT, branch *d6branchT, parVars *d6partitionVarsT) {
	// Load the branch buffer
	for index := 0; index < d6maxNodes; index++ {
		parVars.branchBuf[index] = node.branch[index]
	}
	parVars.branchBuf[d6maxNodes] = *branch
	parVars.branchCount = d6maxNodes + 1

	// Calculate rect containing all in the set
	parVars.coverSplit = parVars.branchBuf[0].rect
	for index := 1; index < d6maxNodes+1; index++ {
		parVars.coverSplit = d6combineRect(&parVars.coverSplit, &parVars.branchBuf[index].rect)
	}
	parVars.coverSplitArea = d6calcRectVolume(&parVars.coverSplit)
}

// Method #0 for choosing a partition:
// As the seeds for the two groups, pick the two rects that would waste the
// most area if covered by a single rectangle, i.e. evidently the worst pair
// to have in the same group.
// Of the remaining, one at a time is chosen to be put in one of the two groups.
// The one chosen is the one with the greatest difference in area expansion
// depending on which group - the rect most strongly attracted to one group
// and repelled from the other.
// If one group gets too full (more would force other group to violate min
// fill requirement) then other group gets the rest.
// These last are the ones that can go in either group most easily.
func d6choosePartition(parVars *d6partitionVarsT, minFill int) {
	var biggestDiff float64
	var group, chosen, betterGroup int

	d6initParVars(parVars, parVars.branchCount, minFill)
	d6pickSeeds(parVars)

	for ((parVars.count[0] + parVars.count[1]) < parVars.total) &&
		(parVars.count[0] < (parVars.total - parVars.minFill)) &&
		(parVars.count[1] < (parVars.total - parVars.minFill)) {
		biggestDiff = -1
		for index := 0; index < parVars.total; index++ {
			if d6notTaken == parVars.partition[index] {
				curRect := &parVars.branchBuf[index].rect
				rect0 := d6combineRect(curRect, &parVars.cover[0])
				rect1 := d6combineRect(curRect, &parVars.cover[1])
				growth0 := d6calcRectVolume(&rect0) - parVars.area[0]
				growth1 := d6calcRectVolume(&rect1) - parVars.area[1]
				diff := growth1 - growth0
				if diff >= 0 {
					group = 0
				} else {
					group = 1
					diff = -diff
				}

				if diff > biggestDiff {
					biggestDiff = diff
					chosen = index
					betterGroup = group
				} else if (diff == biggestDiff) && (parVars.count[group] < parVars.count[betterGroup]) {
					chosen = index
					betterGroup = group
				}
			}
		}
		d6classify(chosen, betterGroup, parVars)
	}

	// If one group too full, put remaining rects in the other
	if (parVars.count[0] + parVars.count[1]) < parVars.total {
		if parVars.count[0] >= parVars.total-parVars.minFill {
			group = 1
		} else {
			group = 0
		}
		for index := 0; index < parVars.total; index++ {
			if d6notTaken == parVars.partition[index] {
				d6classify(index, group, parVars)
			}
		}
	}
}

// Copy branches from the buffer into two nodes according to the partition.
func d6loadNodes(nodeA, nodeB *d6nodeT, parVars *d6partitionVarsT) {
	for index := 0; index < parVars.total; index++ {
		targetNodeIndex := parVars.partition[index]
		targetNodes := []*d6nodeT{nodeA, nodeB}

		// It is assured that d6addBranch here will not cause a node split.
		d6addBranch(&parVars.branchBuf[index], targetNodes[targetNodeIndex], nil)
	}
}

// Initialize a d6partitionVarsT structure.
func d6initParVars(parVars *d6partitionVarsT, maxRects, minFill int) {
	parVars.count[0] = 0
	parVars.count[1] = 0
	parVars.area[0] = 0
	parVars.area[1] = 0
	parVars.total = maxRects
	parVars.minFill = minFill
	for index := 0; index < maxRects; index++ {
		parVars.partition[index] = d6notTaken
	}
}

func d6pickSeeds(parVars *d6partitionVarsT) {
	var seed0, seed1 int
	var worst, waste float64
	var area [d6maxNodes + 1]float64

	for index := 0; index < parVars.total; index++ {
		area[index] = d6calcRectVolume(&parVars.branchBuf[index].rect)
	}

	worst = -parVars.coverSplitArea - 1
	for indexA := 0; indexA < parVars.total-1; indexA++ {
		for indexB := indexA + 1; indexB < parVars.total; indexB++ {
			oneRect := d6combineRect(&parVars.branchBuf[indexA].rect, &parVars.branchBuf[indexB].rect)
			waste = d6calcRectVolume(&oneRect) - area[indexA] - area[indexB]
			if waste > worst {
				worst = waste
				seed0 = indexA
				seed1 = indexB
			}
		}
	}

	d6classify(seed0, 0, parVars)
	d6classify(seed1, 1, parVars)
}

// Put a branch in one of the groups.
func d6classify(index, group int, parVars *d6partitionVarsT) {
	parVars.partition[index] = group

	// Calculate combined rect
	if parVars.count[group] == 0 {
		parVars.cover[group] = parVars.branchBuf[index].rect
	} else {
		parVars.cover[group] = d6combineRect(&parVars.branchBuf[index].rect, &parVars.cover[group])
	}

	// Calculate volume of combined rect
	parVars.area[group] = d6calcRectVolume(&parVars.cover[group])

	parVars.count[group]++
}

// Delete a data rectangle from an index structure.
// Pass in a pointer to a d6rectT, the tid of the record, ptr to ptr to root node.
// Returns 1 if record not found, 0 if success.
// d6removeRect provides for eliminating the root.
func d6removeRect(rect *d6rectT, id interface{}, root **d6nodeT) bool {
	var reInsertList *d6listNodeT

	if !d6removeRectRec(rect, id, *root, &reInsertList) {
		// Found and deleted a data item
		// Reinsert any branches from eliminated nodes
		for reInsertList != nil {
			tempNode := reInsertList.node

			for index := 0; index < tempNode.count; index++ {
				// TODO go over this code. should I use (tempNode->m_level - 1)?
				d6insertRect(&tempNode.branch[index], root, tempNode.level)
			}
			reInsertList = reInsertList.next
		}

		// Check for redundant root (not leaf, 1 child) and eliminate TODO replace
		// if with while? In case there is a whole branch of redundant roots...
		if (*root).count == 1 && (*root).isInternalNode() {
			tempNode := (*root).branch[0].child
			*root = tempNode
		}
		return false
	} else {
		return true
	}
}

// Delete a rectangle from non-root part of an index structure.
// Called by d6removeRect.  Descends tree recursively,
// merges branches on the way back up.
// Returns 1 if record not found, 0 if success.
func d6removeRectRec(rect *d6rectT, id interface{}, node *d6nodeT, listNode **d6listNodeT) bool {
	if node.isInternalNode() { // not a leaf node
		for index := 0; index < node.count; index++ {
			if d6overlap(*rect, node.branch[index].rect) {
				if !d6removeRectRec(rect, id, node.branch[index].child, listNode) {
					if node.branch[index].child.count >= d6minNodes {
						// child removed, just resize parent rect
						node.branch[index].rect = d6nodeCover(node.branch[index].child)
					} else {
						// child removed, not enough entries in node, eliminate node
						d6reInsert(node.branch[index].child, listNode)
						d6disconnectBranch(node, index) // Must return after this call as count has changed
					}
					return false
				}
			}
		}
		return true
	} else { // A leaf node
		for index := 0; index < node.count; index++ {
			if node.branch[index].data == id {
				d6disconnectBranch(node, index) // Must return after this call as count has changed
				return false
			}
		}
		return true
	}
}

// Decide whether two rectangles d6overlap.
func d6overlap(rectA, rectB d6rectT) bool {
	for index := 0; index < d6numDims; index++ {
		if rectA.min[index] > rectB.max[index] ||
			rectB.min[index] > rectA.max[index] {
			return false
		}
	}
	return true
}

// Add a node to the reinsertion list.  All its branches will later
// be reinserted into the index structure.
func d6reInsert(node *d6nodeT, listNode **d6listNodeT) {
	newListNode := &d6listNodeT{}
	newListNode.node = node
	newListNode.next = *listNode
	*listNode = newListNode
}

// d6search in an index tree or subtree for all data retangles that d6overlap the argument rectangle.
func d6search(node *d6nodeT, rect d6rectT, foundCount int, resultCallback func(data interface{}) bool) (int, bool) {
	if node.isInternalNode() {
		// This is an internal node in the tree
		for index := 0; index < node.count; index++ {
			if d6overlap(rect, node.branch[index].rect) {
				var ok bool
				foundCount, ok = d6search(node.branch[index].child, rect, foundCount, resultCallback)
				if !ok {
					// The callback indicated to stop searching
					return foundCount, false
				}
			}
		}
	} else {
		// This is a leaf node
		for index := 0; index < node.count; index++ {
			if d6overlap(rect, node.branch[index].rect) {
				id := node.branch[index].data
				foundCount++
				if !resultCallback(id) {
					return foundCount, false // Don't continue searching
				}

			}
		}
	}
	return foundCount, true // Continue searching
}

func d7fmin(a, b float64) float64 {
	if a < b {
		return a
	}
	return b
}
func d7fmax(a, b float64) float64 {
	if a > b {
		return a
	}
	return b
}

const (
	d7numDims            = 7
	d7maxNodes           = 8
	d7minNodes           = d7maxNodes / 2
	d7useSphericalVolume = true // Better split classification, may be slower on some systems
)

var d7unitSphereVolume = []float64{
	0.000000, 2.000000, 3.141593, // Dimension  0,1,2
	4.188790, 4.934802, 5.263789, // Dimension  3,4,5
	5.167713, 4.724766, 4.058712, // Dimension  6,7,8
	3.298509, 2.550164, 1.884104, // Dimension  9,10,11
	1.335263, 0.910629, 0.599265, // Dimension  12,13,14
	0.381443, 0.235331, 0.140981, // Dimension  15,16,17
	0.082146, 0.046622, 0.025807, // Dimension  18,19,20
}[d7numDims]

type d7RTree struct {
	root *d7nodeT ///< Root of tree
}

/// Minimal bounding rectangle (n-dimensional)
type d7rectT struct {
	min [d7numDims]float64 ///< Min dimensions of bounding box
	max [d7numDims]float64 ///< Max dimensions of bounding box
}

/// May be data or may be another subtree
/// The parents level determines this.
/// If the parents level is 0, then this is data
type d7branchT struct {
	rect  d7rectT     ///< Bounds
	child *d7nodeT    ///< Child node
	data  interface{} ///< Data Id or Ptr
}

/// d7nodeT for each branch level
type d7nodeT struct {
	count  int                   ///< Count
	level  int                   ///< Leaf is zero, others positive
	branch [d7maxNodes]d7branchT ///< Branch
}

func (node *d7nodeT) isInternalNode() bool {
	return (node.level > 0) // Not a leaf, but a internal node
}
func (node *d7nodeT) isLeaf() bool {
	return (node.level == 0) // A leaf, contains data
}

/// A link list of nodes for reinsertion after a delete operation
type d7listNodeT struct {
	next *d7listNodeT ///< Next in list
	node *d7nodeT     ///< Node
}

const d7notTaken = -1 // indicates that position

/// Variables for finding a split partition
type d7partitionVarsT struct {
	partition [d7maxNodes + 1]int
	total     int
	minFill   int
	count     [2]int
	cover     [2]d7rectT
	area      [2]float64

	branchBuf      [d7maxNodes + 1]d7branchT
	branchCount    int
	coverSplit     d7rectT
	coverSplitArea float64
}

func d7New() *d7RTree {
	// We only support machine word size simple data type eg. integer index or object pointer.
	// Since we are storing as union with non data branch
	return &d7RTree{
		root: &d7nodeT{},
	}
}

/// Insert entry
/// \param a_min Min of bounding rect
/// \param a_max Max of bounding rect
/// \param a_dataId Positive Id of data.  Maybe zero, but negative numbers not allowed.
func (tr *d7RTree) Insert(min, max [d7numDims]float64, dataId interface{}) {
	var branch d7branchT
	branch.data = dataId
	for axis := 0; axis < d7numDims; axis++ {
		branch.rect.min[axis] = min[axis]
		branch.rect.max[axis] = max[axis]
	}
	d7insertRect(&branch, &tr.root, 0)
}

/// Remove entry
/// \param a_min Min of bounding rect
/// \param a_max Max of bounding rect
/// \param a_dataId Positive Id of data.  Maybe zero, but negative numbers not allowed.
func (tr *d7RTree) Remove(min, max [d7numDims]float64, dataId interface{}) {
	var rect d7rectT
	for axis := 0; axis < d7numDims; axis++ {
		rect.min[axis] = min[axis]
		rect.max[axis] = max[axis]
	}
	d7removeRect(&rect, dataId, &tr.root)
}

/// Find all within d7search rectangle
/// \param a_min Min of d7search bounding rect
/// \param a_max Max of d7search bounding rect
/// \param a_searchResult d7search result array.  Caller should set grow size. Function will reset, not append to array.
/// \param a_resultCallback Callback function to return result.  Callback should return 'true' to continue searching
/// \param a_context User context to pass as parameter to a_resultCallback
/// \return Returns the number of entries found
func (tr *d7RTree) Search(min, max [d7numDims]float64, resultCallback func(data interface{}) bool) int {
	var rect d7rectT
	for axis := 0; axis < d7numDims; axis++ {
		rect.min[axis] = min[axis]
		rect.max[axis] = max[axis]
	}
	foundCount, _ := d7search(tr.root, rect, 0, resultCallback)
	return foundCount
}

/// Count the data elements in this container.  This is slow as no internal counter is maintained.
func (tr *d7RTree) Count() int {
	var count int
	d7countRec(tr.root, &count)
	return count
}

/// Remove all entries from tree
func (tr *d7RTree) RemoveAll() {
	// Delete all existing nodes
	tr.root = &d7nodeT{}
}

func d7countRec(node *d7nodeT, count *int) {
	if node.isInternalNode() { // not a leaf node
		for index := 0; index < node.count; index++ {
			d7countRec(node.branch[index].child, count)
		}
	} else { // A leaf node
		*count += node.count
	}
}

// Inserts a new data rectangle into the index structure.
// Recursively descends tree, propagates splits back up.
// Returns 0 if node was not split.  Old node updated.
// If node was split, returns 1 and sets the pointer pointed to by
// new_node to point to the new node.  Old node updated to become one of two.
// The level argument specifies the number of steps up from the leaf
// level to insert; e.g. a data rectangle goes in at level = 0.
func d7insertRectRec(branch *d7branchT, node *d7nodeT, newNode **d7nodeT, level int) bool {
	// recurse until we reach the correct level for the new record. data records
	// will always be called with a_level == 0 (leaf)
	if node.level > level {
		// Still above level for insertion, go down tree recursively
		var otherNode *d7nodeT
		//var newBranch d7branchT

		// find the optimal branch for this record
		index := d7pickBranch(&branch.rect, node)

		// recursively insert this record into the picked branch
		childWasSplit := d7insertRectRec(branch, node.branch[index].child, &otherNode, level)

		if !childWasSplit {
			// Child was not split. Merge the bounding box of the new record with the
			// existing bounding box
			node.branch[index].rect = d7combineRect(&branch.rect, &(node.branch[index].rect))
			return false
		} else {
			// Child was split. The old branches are now re-partitioned to two nodes
			// so we have to re-calculate the bounding boxes of each node
			node.branch[index].rect = d7nodeCover(node.branch[index].child)
			var newBranch d7branchT
			newBranch.child = otherNode
			newBranch.rect = d7nodeCover(otherNode)

			// The old node is already a child of a_node. Now add the newly-created
			// node to a_node as well. a_node might be split because of that.
			return d7addBranch(&newBranch, node, newNode)
		}
	} else if node.level == level {
		// We have reached level for insertion. Add rect, split if necessary
		return d7addBranch(branch, node, newNode)
	} else {
		// Should never occur
		return false
	}
}

// Insert a data rectangle into an index structure.
// d7insertRect provides for splitting the root;
// returns 1 if root was split, 0 if it was not.
// The level argument specifies the number of steps up from the leaf
// level to insert; e.g. a data rectangle goes in at level = 0.
// InsertRect2 does the recursion.
//
func d7insertRect(branch *d7branchT, root **d7nodeT, level int) bool {
	var newNode *d7nodeT

	if d7insertRectRec(branch, *root, &newNode, level) { // Root split

		// Grow tree taller and new root
		newRoot := &d7nodeT{}
		newRoot.level = (*root).level + 1

		var newBranch d7branchT

		// add old root node as a child of the new root
		newBranch.rect = d7nodeCover(*root)
		newBranch.child = *root
		d7addBranch(&newBranch, newRoot, nil)

		// add the split node as a child of the new root
		newBranch.rect = d7nodeCover(newNode)
		newBranch.child = newNode
		d7addBranch(&newBranch, newRoot, nil)

		// set the new root as the root node
		*root = newRoot

		return true
	}
	return false
}

// Find the smallest rectangle that includes all rectangles in branches of a node.
func d7nodeCover(node *d7nodeT) d7rectT {
	rect := node.branch[0].rect
	for index := 1; index < node.count; index++ {
		rect = d7combineRect(&rect, &(node.branch[index].rect))
	}
	return rect
}

// Add a branch to a node.  Split the node if necessary.
// Returns 0 if node not split.  Old node updated.
// Returns 1 if node split, sets *new_node to address of new node.
// Old node updated, becomes one of two.
func d7addBranch(branch *d7branchT, node *d7nodeT, newNode **d7nodeT) bool {
	if node.count < d7maxNodes { // Split won't be necessary
		node.branch[node.count] = *branch
		node.count++
		return false
	} else {
		d7splitNode(node, branch, newNode)
		return true
	}
}

// Disconnect a dependent node.
// Caller must return (or stop using iteration index) after this as count has changed
func d7disconnectBranch(node *d7nodeT, index int) {
	// Remove element by swapping with the last element to prevent gaps in array
	node.branch[index] = node.branch[node.count-1]
	node.branch[node.count-1].data = nil
	node.branch[node.count-1].child = nil
	node.count--
}

// Pick a branch.  Pick the one that will need the smallest increase
// in area to accomodate the new rectangle.  This will result in the
// least total area for the covering rectangles in the current node.
// In case of a tie, pick the one which was smaller before, to get
// the best resolution when searching.
func d7pickBranch(rect *d7rectT, node *d7nodeT) int {
	var firstTime bool = true
	var increase float64
	var bestIncr float64 = -1
	var area float64
	var bestArea float64
	var best int
	var tempRect d7rectT

	for index := 0; index < node.count; index++ {
		curRect := &node.branch[index].rect
		area = d7calcRectVolume(curRect)
		tempRect = d7combineRect(rect, curRect)
		increase = d7calcRectVolume(&tempRect) - area
		if (increase < bestIncr) || firstTime {
			best = index
			bestArea = area
			bestIncr = increase
			firstTime = false
		} else if (increase == bestIncr) && (area < bestArea) {
			best = index
			bestArea = area
			bestIncr = increase
		}
	}
	return best
}

// Combine two rectangles into larger one containing both
func d7combineRect(rectA, rectB *d7rectT) d7rectT {
	var newRect d7rectT

	for index := 0; index < d7numDims; index++ {
		newRect.min[index] = d7fmin(rectA.min[index], rectB.min[index])
		newRect.max[index] = d7fmax(rectA.max[index], rectB.max[index])
	}

	return newRect
}

// Split a node.
// Divides the nodes branches and the extra one between two nodes.
// Old node is one of the new ones, and one really new one is created.
// Tries more than one method for choosing a partition, uses best result.
func d7splitNode(node *d7nodeT, branch *d7branchT, newNode **d7nodeT) {
	// Could just use local here, but member or external is faster since it is reused
	var localVars d7partitionVarsT
	parVars := &localVars

	// Load all the branches into a buffer, initialize old node
	d7getBranches(node, branch, parVars)

	// Find partition
	d7choosePartition(parVars, d7minNodes)

	// Create a new node to hold (about) half of the branches
	*newNode = &d7nodeT{}
	(*newNode).level = node.level

	// Put branches from buffer into 2 nodes according to the chosen partition
	node.count = 0
	d7loadNodes(node, *newNode, parVars)
}

// Calculate the n-dimensional volume of a rectangle
func d7rectVolume(rect *d7rectT) float64 {
	var volume float64 = 1
	for index := 0; index < d7numDims; index++ {
		volume *= rect.max[index] - rect.min[index]
	}
	return volume
}

// The exact volume of the bounding sphere for the given d7rectT
func d7rectSphericalVolume(rect *d7rectT) float64 {
	var sumOfSquares float64 = 0
	var radius float64

	for index := 0; index < d7numDims; index++ {
		halfExtent := (rect.max[index] - rect.min[index]) * 0.5
		sumOfSquares += halfExtent * halfExtent
	}

	radius = math.Sqrt(sumOfSquares)

	// Pow maybe slow, so test for common dims just use x*x, x*x*x.
	if d7numDims == 5 {
		return (radius * radius * radius * radius * radius * d7unitSphereVolume)
	} else if d7numDims == 4 {
		return (radius * radius * radius * radius * d7unitSphereVolume)
	} else if d7numDims == 3 {
		return (radius * radius * radius * d7unitSphereVolume)
	} else if d7numDims == 2 {
		return (radius * radius * d7unitSphereVolume)
	} else {
		return (math.Pow(radius, d7numDims) * d7unitSphereVolume)
	}
}

// Use one of the methods to calculate retangle volume
func d7calcRectVolume(rect *d7rectT) float64 {
	if d7useSphericalVolume {
		return d7rectSphericalVolume(rect) // Slower but helps certain merge cases
	} else { // RTREE_USE_SPHERICAL_VOLUME
		return d7rectVolume(rect) // Faster but can cause poor merges
	} // RTREE_USE_SPHERICAL_VOLUME
}

// Load branch buffer with branches from full node plus the extra branch.
func d7getBranches(node *d7nodeT, branch *d7branchT, parVars *d7partitionVarsT) {
	// Load the branch buffer
	for index := 0; index < d7maxNodes; index++ {
		parVars.branchBuf[index] = node.branch[index]
	}
	parVars.branchBuf[d7maxNodes] = *branch
	parVars.branchCount = d7maxNodes + 1

	// Calculate rect containing all in the set
	parVars.coverSplit = parVars.branchBuf[0].rect
	for index := 1; index < d7maxNodes+1; index++ {
		parVars.coverSplit = d7combineRect(&parVars.coverSplit, &parVars.branchBuf[index].rect)
	}
	parVars.coverSplitArea = d7calcRectVolume(&parVars.coverSplit)
}

// Method #0 for choosing a partition:
// As the seeds for the two groups, pick the two rects that would waste the
// most area if covered by a single rectangle, i.e. evidently the worst pair
// to have in the same group.
// Of the remaining, one at a time is chosen to be put in one of the two groups.
// The one chosen is the one with the greatest difference in area expansion
// depending on which group - the rect most strongly attracted to one group
// and repelled from the other.
// If one group gets too full (more would force other group to violate min
// fill requirement) then other group gets the rest.
// These last are the ones that can go in either group most easily.
func d7choosePartition(parVars *d7partitionVarsT, minFill int) {
	var biggestDiff float64
	var group, chosen, betterGroup int

	d7initParVars(parVars, parVars.branchCount, minFill)
	d7pickSeeds(parVars)

	for ((parVars.count[0] + parVars.count[1]) < parVars.total) &&
		(parVars.count[0] < (parVars.total - parVars.minFill)) &&
		(parVars.count[1] < (parVars.total - parVars.minFill)) {
		biggestDiff = -1
		for index := 0; index < parVars.total; index++ {
			if d7notTaken == parVars.partition[index] {
				curRect := &parVars.branchBuf[index].rect
				rect0 := d7combineRect(curRect, &parVars.cover[0])
				rect1 := d7combineRect(curRect, &parVars.cover[1])
				growth0 := d7calcRectVolume(&rect0) - parVars.area[0]
				growth1 := d7calcRectVolume(&rect1) - parVars.area[1]
				diff := growth1 - growth0
				if diff >= 0 {
					group = 0
				} else {
					group = 1
					diff = -diff
				}

				if diff > biggestDiff {
					biggestDiff = diff
					chosen = index
					betterGroup = group
				} else if (diff == biggestDiff) && (parVars.count[group] < parVars.count[betterGroup]) {
					chosen = index
					betterGroup = group
				}
			}
		}
		d7classify(chosen, betterGroup, parVars)
	}

	// If one group too full, put remaining rects in the other
	if (parVars.count[0] + parVars.count[1]) < parVars.total {
		if parVars.count[0] >= parVars.total-parVars.minFill {
			group = 1
		} else {
			group = 0
		}
		for index := 0; index < parVars.total; index++ {
			if d7notTaken == parVars.partition[index] {
				d7classify(index, group, parVars)
			}
		}
	}
}

// Copy branches from the buffer into two nodes according to the partition.
func d7loadNodes(nodeA, nodeB *d7nodeT, parVars *d7partitionVarsT) {
	for index := 0; index < parVars.total; index++ {
		targetNodeIndex := parVars.partition[index]
		targetNodes := []*d7nodeT{nodeA, nodeB}

		// It is assured that d7addBranch here will not cause a node split.
		d7addBranch(&parVars.branchBuf[index], targetNodes[targetNodeIndex], nil)
	}
}

// Initialize a d7partitionVarsT structure.
func d7initParVars(parVars *d7partitionVarsT, maxRects, minFill int) {
	parVars.count[0] = 0
	parVars.count[1] = 0
	parVars.area[0] = 0
	parVars.area[1] = 0
	parVars.total = maxRects
	parVars.minFill = minFill
	for index := 0; index < maxRects; index++ {
		parVars.partition[index] = d7notTaken
	}
}

func d7pickSeeds(parVars *d7partitionVarsT) {
	var seed0, seed1 int
	var worst, waste float64
	var area [d7maxNodes + 1]float64

	for index := 0; index < parVars.total; index++ {
		area[index] = d7calcRectVolume(&parVars.branchBuf[index].rect)
	}

	worst = -parVars.coverSplitArea - 1
	for indexA := 0; indexA < parVars.total-1; indexA++ {
		for indexB := indexA + 1; indexB < parVars.total; indexB++ {
			oneRect := d7combineRect(&parVars.branchBuf[indexA].rect, &parVars.branchBuf[indexB].rect)
			waste = d7calcRectVolume(&oneRect) - area[indexA] - area[indexB]
			if waste > worst {
				worst = waste
				seed0 = indexA
				seed1 = indexB
			}
		}
	}

	d7classify(seed0, 0, parVars)
	d7classify(seed1, 1, parVars)
}

// Put a branch in one of the groups.
func d7classify(index, group int, parVars *d7partitionVarsT) {
	parVars.partition[index] = group

	// Calculate combined rect
	if parVars.count[group] == 0 {
		parVars.cover[group] = parVars.branchBuf[index].rect
	} else {
		parVars.cover[group] = d7combineRect(&parVars.branchBuf[index].rect, &parVars.cover[group])
	}

	// Calculate volume of combined rect
	parVars.area[group] = d7calcRectVolume(&parVars.cover[group])

	parVars.count[group]++
}

// Delete a data rectangle from an index structure.
// Pass in a pointer to a d7rectT, the tid of the record, ptr to ptr to root node.
// Returns 1 if record not found, 0 if success.
// d7removeRect provides for eliminating the root.
func d7removeRect(rect *d7rectT, id interface{}, root **d7nodeT) bool {
	var reInsertList *d7listNodeT

	if !d7removeRectRec(rect, id, *root, &reInsertList) {
		// Found and deleted a data item
		// Reinsert any branches from eliminated nodes
		for reInsertList != nil {
			tempNode := reInsertList.node

			for index := 0; index < tempNode.count; index++ {
				// TODO go over this code. should I use (tempNode->m_level - 1)?
				d7insertRect(&tempNode.branch[index], root, tempNode.level)
			}
			reInsertList = reInsertList.next
		}

		// Check for redundant root (not leaf, 1 child) and eliminate TODO replace
		// if with while? In case there is a whole branch of redundant roots...
		if (*root).count == 1 && (*root).isInternalNode() {
			tempNode := (*root).branch[0].child
			*root = tempNode
		}
		return false
	} else {
		return true
	}
}

// Delete a rectangle from non-root part of an index structure.
// Called by d7removeRect.  Descends tree recursively,
// merges branches on the way back up.
// Returns 1 if record not found, 0 if success.
func d7removeRectRec(rect *d7rectT, id interface{}, node *d7nodeT, listNode **d7listNodeT) bool {
	if node.isInternalNode() { // not a leaf node
		for index := 0; index < node.count; index++ {
			if d7overlap(*rect, node.branch[index].rect) {
				if !d7removeRectRec(rect, id, node.branch[index].child, listNode) {
					if node.branch[index].child.count >= d7minNodes {
						// child removed, just resize parent rect
						node.branch[index].rect = d7nodeCover(node.branch[index].child)
					} else {
						// child removed, not enough entries in node, eliminate node
						d7reInsert(node.branch[index].child, listNode)
						d7disconnectBranch(node, index) // Must return after this call as count has changed
					}
					return false
				}
			}
		}
		return true
	} else { // A leaf node
		for index := 0; index < node.count; index++ {
			if node.branch[index].data == id {
				d7disconnectBranch(node, index) // Must return after this call as count has changed
				return false
			}
		}
		return true
	}
}

// Decide whether two rectangles d7overlap.
func d7overlap(rectA, rectB d7rectT) bool {
	for index := 0; index < d7numDims; index++ {
		if rectA.min[index] > rectB.max[index] ||
			rectB.min[index] > rectA.max[index] {
			return false
		}
	}
	return true
}

// Add a node to the reinsertion list.  All its branches will later
// be reinserted into the index structure.
func d7reInsert(node *d7nodeT, listNode **d7listNodeT) {
	newListNode := &d7listNodeT{}
	newListNode.node = node
	newListNode.next = *listNode
	*listNode = newListNode
}

// d7search in an index tree or subtree for all data retangles that d7overlap the argument rectangle.
func d7search(node *d7nodeT, rect d7rectT, foundCount int, resultCallback func(data interface{}) bool) (int, bool) {
	if node.isInternalNode() {
		// This is an internal node in the tree
		for index := 0; index < node.count; index++ {
			if d7overlap(rect, node.branch[index].rect) {
				var ok bool
				foundCount, ok = d7search(node.branch[index].child, rect, foundCount, resultCallback)
				if !ok {
					// The callback indicated to stop searching
					return foundCount, false
				}
			}
		}
	} else {
		// This is a leaf node
		for index := 0; index < node.count; index++ {
			if d7overlap(rect, node.branch[index].rect) {
				id := node.branch[index].data
				foundCount++
				if !resultCallback(id) {
					return foundCount, false // Don't continue searching
				}

			}
		}
	}
	return foundCount, true // Continue searching
}

func d8fmin(a, b float64) float64 {
	if a < b {
		return a
	}
	return b
}
func d8fmax(a, b float64) float64 {
	if a > b {
		return a
	}
	return b
}

const (
	d8numDims            = 8
	d8maxNodes           = 8
	d8minNodes           = d8maxNodes / 2
	d8useSphericalVolume = true // Better split classification, may be slower on some systems
)

var d8unitSphereVolume = []float64{
	0.000000, 2.000000, 3.141593, // Dimension  0,1,2
	4.188790, 4.934802, 5.263789, // Dimension  3,4,5
	5.167713, 4.724766, 4.058712, // Dimension  6,7,8
	3.298509, 2.550164, 1.884104, // Dimension  9,10,11
	1.335263, 0.910629, 0.599265, // Dimension  12,13,14
	0.381443, 0.235331, 0.140981, // Dimension  15,16,17
	0.082146, 0.046622, 0.025807, // Dimension  18,19,20
}[d8numDims]

type d8RTree struct {
	root *d8nodeT ///< Root of tree
}

/// Minimal bounding rectangle (n-dimensional)
type d8rectT struct {
	min [d8numDims]float64 ///< Min dimensions of bounding box
	max [d8numDims]float64 ///< Max dimensions of bounding box
}

/// May be data or may be another subtree
/// The parents level determines this.
/// If the parents level is 0, then this is data
type d8branchT struct {
	rect  d8rectT     ///< Bounds
	child *d8nodeT    ///< Child node
	data  interface{} ///< Data Id or Ptr
}

/// d8nodeT for each branch level
type d8nodeT struct {
	count  int                   ///< Count
	level  int                   ///< Leaf is zero, others positive
	branch [d8maxNodes]d8branchT ///< Branch
}

func (node *d8nodeT) isInternalNode() bool {
	return (node.level > 0) // Not a leaf, but a internal node
}
func (node *d8nodeT) isLeaf() bool {
	return (node.level == 0) // A leaf, contains data
}

/// A link list of nodes for reinsertion after a delete operation
type d8listNodeT struct {
	next *d8listNodeT ///< Next in list
	node *d8nodeT     ///< Node
}

const d8notTaken = -1 // indicates that position

/// Variables for finding a split partition
type d8partitionVarsT struct {
	partition [d8maxNodes + 1]int
	total     int
	minFill   int
	count     [2]int
	cover     [2]d8rectT
	area      [2]float64

	branchBuf      [d8maxNodes + 1]d8branchT
	branchCount    int
	coverSplit     d8rectT
	coverSplitArea float64
}

func d8New() *d8RTree {
	// We only support machine word size simple data type eg. integer index or object pointer.
	// Since we are storing as union with non data branch
	return &d8RTree{
		root: &d8nodeT{},
	}
}

/// Insert entry
/// \param a_min Min of bounding rect
/// \param a_max Max of bounding rect
/// \param a_dataId Positive Id of data.  Maybe zero, but negative numbers not allowed.
func (tr *d8RTree) Insert(min, max [d8numDims]float64, dataId interface{}) {
	var branch d8branchT
	branch.data = dataId
	for axis := 0; axis < d8numDims; axis++ {
		branch.rect.min[axis] = min[axis]
		branch.rect.max[axis] = max[axis]
	}
	d8insertRect(&branch, &tr.root, 0)
}

/// Remove entry
/// \param a_min Min of bounding rect
/// \param a_max Max of bounding rect
/// \param a_dataId Positive Id of data.  Maybe zero, but negative numbers not allowed.
func (tr *d8RTree) Remove(min, max [d8numDims]float64, dataId interface{}) {
	var rect d8rectT
	for axis := 0; axis < d8numDims; axis++ {
		rect.min[axis] = min[axis]
		rect.max[axis] = max[axis]
	}
	d8removeRect(&rect, dataId, &tr.root)
}

/// Find all within d8search rectangle
/// \param a_min Min of d8search bounding rect
/// \param a_max Max of d8search bounding rect
/// \param a_searchResult d8search result array.  Caller should set grow size. Function will reset, not append to array.
/// \param a_resultCallback Callback function to return result.  Callback should return 'true' to continue searching
/// \param a_context User context to pass as parameter to a_resultCallback
/// \return Returns the number of entries found
func (tr *d8RTree) Search(min, max [d8numDims]float64, resultCallback func(data interface{}) bool) int {
	var rect d8rectT
	for axis := 0; axis < d8numDims; axis++ {
		rect.min[axis] = min[axis]
		rect.max[axis] = max[axis]
	}
	foundCount, _ := d8search(tr.root, rect, 0, resultCallback)
	return foundCount
}

/// Count the data elements in this container.  This is slow as no internal counter is maintained.
func (tr *d8RTree) Count() int {
	var count int
	d8countRec(tr.root, &count)
	return count
}

/// Remove all entries from tree
func (tr *d8RTree) RemoveAll() {
	// Delete all existing nodes
	tr.root = &d8nodeT{}
}

func d8countRec(node *d8nodeT, count *int) {
	if node.isInternalNode() { // not a leaf node
		for index := 0; index < node.count; index++ {
			d8countRec(node.branch[index].child, count)
		}
	} else { // A leaf node
		*count += node.count
	}
}

// Inserts a new data rectangle into the index structure.
// Recursively descends tree, propagates splits back up.
// Returns 0 if node was not split.  Old node updated.
// If node was split, returns 1 and sets the pointer pointed to by
// new_node to point to the new node.  Old node updated to become one of two.
// The level argument specifies the number of steps up from the leaf
// level to insert; e.g. a data rectangle goes in at level = 0.
func d8insertRectRec(branch *d8branchT, node *d8nodeT, newNode **d8nodeT, level int) bool {
	// recurse until we reach the correct level for the new record. data records
	// will always be called with a_level == 0 (leaf)
	if node.level > level {
		// Still above level for insertion, go down tree recursively
		var otherNode *d8nodeT
		//var newBranch d8branchT

		// find the optimal branch for this record
		index := d8pickBranch(&branch.rect, node)

		// recursively insert this record into the picked branch
		childWasSplit := d8insertRectRec(branch, node.branch[index].child, &otherNode, level)

		if !childWasSplit {
			// Child was not split. Merge the bounding box of the new record with the
			// existing bounding box
			node.branch[index].rect = d8combineRect(&branch.rect, &(node.branch[index].rect))
			return false
		} else {
			// Child was split. The old branches are now re-partitioned to two nodes
			// so we have to re-calculate the bounding boxes of each node
			node.branch[index].rect = d8nodeCover(node.branch[index].child)
			var newBranch d8branchT
			newBranch.child = otherNode
			newBranch.rect = d8nodeCover(otherNode)

			// The old node is already a child of a_node. Now add the newly-created
			// node to a_node as well. a_node might be split because of that.
			return d8addBranch(&newBranch, node, newNode)
		}
	} else if node.level == level {
		// We have reached level for insertion. Add rect, split if necessary
		return d8addBranch(branch, node, newNode)
	} else {
		// Should never occur
		return false
	}
}

// Insert a data rectangle into an index structure.
// d8insertRect provides for splitting the root;
// returns 1 if root was split, 0 if it was not.
// The level argument specifies the number of steps up from the leaf
// level to insert; e.g. a data rectangle goes in at level = 0.
// InsertRect2 does the recursion.
//
func d8insertRect(branch *d8branchT, root **d8nodeT, level int) bool {
	var newNode *d8nodeT

	if d8insertRectRec(branch, *root, &newNode, level) { // Root split

		// Grow tree taller and new root
		newRoot := &d8nodeT{}
		newRoot.level = (*root).level + 1

		var newBranch d8branchT

		// add old root node as a child of the new root
		newBranch.rect = d8nodeCover(*root)
		newBranch.child = *root
		d8addBranch(&newBranch, newRoot, nil)

		// add the split node as a child of the new root
		newBranch.rect = d8nodeCover(newNode)
		newBranch.child = newNode
		d8addBranch(&newBranch, newRoot, nil)

		// set the new root as the root node
		*root = newRoot

		return true
	}
	return false
}

// Find the smallest rectangle that includes all rectangles in branches of a node.
func d8nodeCover(node *d8nodeT) d8rectT {
	rect := node.branch[0].rect
	for index := 1; index < node.count; index++ {
		rect = d8combineRect(&rect, &(node.branch[index].rect))
	}
	return rect
}

// Add a branch to a node.  Split the node if necessary.
// Returns 0 if node not split.  Old node updated.
// Returns 1 if node split, sets *new_node to address of new node.
// Old node updated, becomes one of two.
func d8addBranch(branch *d8branchT, node *d8nodeT, newNode **d8nodeT) bool {
	if node.count < d8maxNodes { // Split won't be necessary
		node.branch[node.count] = *branch
		node.count++
		return false
	} else {
		d8splitNode(node, branch, newNode)
		return true
	}
}

// Disconnect a dependent node.
// Caller must return (or stop using iteration index) after this as count has changed
func d8disconnectBranch(node *d8nodeT, index int) {
	// Remove element by swapping with the last element to prevent gaps in array
	node.branch[index] = node.branch[node.count-1]
	node.branch[node.count-1].data = nil
	node.branch[node.count-1].child = nil
	node.count--
}

// Pick a branch.  Pick the one that will need the smallest increase
// in area to accomodate the new rectangle.  This will result in the
// least total area for the covering rectangles in the current node.
// In case of a tie, pick the one which was smaller before, to get
// the best resolution when searching.
func d8pickBranch(rect *d8rectT, node *d8nodeT) int {
	var firstTime bool = true
	var increase float64
	var bestIncr float64 = -1
	var area float64
	var bestArea float64
	var best int
	var tempRect d8rectT

	for index := 0; index < node.count; index++ {
		curRect := &node.branch[index].rect
		area = d8calcRectVolume(curRect)
		tempRect = d8combineRect(rect, curRect)
		increase = d8calcRectVolume(&tempRect) - area
		if (increase < bestIncr) || firstTime {
			best = index
			bestArea = area
			bestIncr = increase
			firstTime = false
		} else if (increase == bestIncr) && (area < bestArea) {
			best = index
			bestArea = area
			bestIncr = increase
		}
	}
	return best
}

// Combine two rectangles into larger one containing both
func d8combineRect(rectA, rectB *d8rectT) d8rectT {
	var newRect d8rectT

	for index := 0; index < d8numDims; index++ {
		newRect.min[index] = d8fmin(rectA.min[index], rectB.min[index])
		newRect.max[index] = d8fmax(rectA.max[index], rectB.max[index])
	}

	return newRect
}

// Split a node.
// Divides the nodes branches and the extra one between two nodes.
// Old node is one of the new ones, and one really new one is created.
// Tries more than one method for choosing a partition, uses best result.
func d8splitNode(node *d8nodeT, branch *d8branchT, newNode **d8nodeT) {
	// Could just use local here, but member or external is faster since it is reused
	var localVars d8partitionVarsT
	parVars := &localVars

	// Load all the branches into a buffer, initialize old node
	d8getBranches(node, branch, parVars)

	// Find partition
	d8choosePartition(parVars, d8minNodes)

	// Create a new node to hold (about) half of the branches
	*newNode = &d8nodeT{}
	(*newNode).level = node.level

	// Put branches from buffer into 2 nodes according to the chosen partition
	node.count = 0
	d8loadNodes(node, *newNode, parVars)
}

// Calculate the n-dimensional volume of a rectangle
func d8rectVolume(rect *d8rectT) float64 {
	var volume float64 = 1
	for index := 0; index < d8numDims; index++ {
		volume *= rect.max[index] - rect.min[index]
	}
	return volume
}

// The exact volume of the bounding sphere for the given d8rectT
func d8rectSphericalVolume(rect *d8rectT) float64 {
	var sumOfSquares float64 = 0
	var radius float64

	for index := 0; index < d8numDims; index++ {
		halfExtent := (rect.max[index] - rect.min[index]) * 0.5
		sumOfSquares += halfExtent * halfExtent
	}

	radius = math.Sqrt(sumOfSquares)

	// Pow maybe slow, so test for common dims just use x*x, x*x*x.
	if d8numDims == 5 {
		return (radius * radius * radius * radius * radius * d8unitSphereVolume)
	} else if d8numDims == 4 {
		return (radius * radius * radius * radius * d8unitSphereVolume)
	} else if d8numDims == 3 {
		return (radius * radius * radius * d8unitSphereVolume)
	} else if d8numDims == 2 {
		return (radius * radius * d8unitSphereVolume)
	} else {
		return (math.Pow(radius, d8numDims) * d8unitSphereVolume)
	}
}

// Use one of the methods to calculate retangle volume
func d8calcRectVolume(rect *d8rectT) float64 {
	if d8useSphericalVolume {
		return d8rectSphericalVolume(rect) // Slower but helps certain merge cases
	} else { // RTREE_USE_SPHERICAL_VOLUME
		return d8rectVolume(rect) // Faster but can cause poor merges
	} // RTREE_USE_SPHERICAL_VOLUME
}

// Load branch buffer with branches from full node plus the extra branch.
func d8getBranches(node *d8nodeT, branch *d8branchT, parVars *d8partitionVarsT) {
	// Load the branch buffer
	for index := 0; index < d8maxNodes; index++ {
		parVars.branchBuf[index] = node.branch[index]
	}
	parVars.branchBuf[d8maxNodes] = *branch
	parVars.branchCount = d8maxNodes + 1

	// Calculate rect containing all in the set
	parVars.coverSplit = parVars.branchBuf[0].rect
	for index := 1; index < d8maxNodes+1; index++ {
		parVars.coverSplit = d8combineRect(&parVars.coverSplit, &parVars.branchBuf[index].rect)
	}
	parVars.coverSplitArea = d8calcRectVolume(&parVars.coverSplit)
}

// Method #0 for choosing a partition:
// As the seeds for the two groups, pick the two rects that would waste the
// most area if covered by a single rectangle, i.e. evidently the worst pair
// to have in the same group.
// Of the remaining, one at a time is chosen to be put in one of the two groups.
// The one chosen is the one with the greatest difference in area expansion
// depending on which group - the rect most strongly attracted to one group
// and repelled from the other.
// If one group gets too full (more would force other group to violate min
// fill requirement) then other group gets the rest.
// These last are the ones that can go in either group most easily.
func d8choosePartition(parVars *d8partitionVarsT, minFill int) {
	var biggestDiff float64
	var group, chosen, betterGroup int

	d8initParVars(parVars, parVars.branchCount, minFill)
	d8pickSeeds(parVars)

	for ((parVars.count[0] + parVars.count[1]) < parVars.total) &&
		(parVars.count[0] < (parVars.total - parVars.minFill)) &&
		(parVars.count[1] < (parVars.total - parVars.minFill)) {
		biggestDiff = -1
		for index := 0; index < parVars.total; index++ {
			if d8notTaken == parVars.partition[index] {
				curRect := &parVars.branchBuf[index].rect
				rect0 := d8combineRect(curRect, &parVars.cover[0])
				rect1 := d8combineRect(curRect, &parVars.cover[1])
				growth0 := d8calcRectVolume(&rect0) - parVars.area[0]
				growth1 := d8calcRectVolume(&rect1) - parVars.area[1]
				diff := growth1 - growth0
				if diff >= 0 {
					group = 0
				} else {
					group = 1
					diff = -diff
				}

				if diff > biggestDiff {
					biggestDiff = diff
					chosen = index
					betterGroup = group
				} else if (diff == biggestDiff) && (parVars.count[group] < parVars.count[betterGroup]) {
					chosen = index
					betterGroup = group
				}
			}
		}
		d8classify(chosen, betterGroup, parVars)
	}

	// If one group too full, put remaining rects in the other
	if (parVars.count[0] + parVars.count[1]) < parVars.total {
		if parVars.count[0] >= parVars.total-parVars.minFill {
			group = 1
		} else {
			group = 0
		}
		for index := 0; index < parVars.total; index++ {
			if d8notTaken == parVars.partition[index] {
				d8classify(index, group, parVars)
			}
		}
	}
}

// Copy branches from the buffer into two nodes according to the partition.
func d8loadNodes(nodeA, nodeB *d8nodeT, parVars *d8partitionVarsT) {
	for index := 0; index < parVars.total; index++ {
		targetNodeIndex := parVars.partition[index]
		targetNodes := []*d8nodeT{nodeA, nodeB}

		// It is assured that d8addBranch here will not cause a node split.
		d8addBranch(&parVars.branchBuf[index], targetNodes[targetNodeIndex], nil)
	}
}

// Initialize a d8partitionVarsT structure.
func d8initParVars(parVars *d8partitionVarsT, maxRects, minFill int) {
	parVars.count[0] = 0
	parVars.count[1] = 0
	parVars.area[0] = 0
	parVars.area[1] = 0
	parVars.total = maxRects
	parVars.minFill = minFill
	for index := 0; index < maxRects; index++ {
		parVars.partition[index] = d8notTaken
	}
}

func d8pickSeeds(parVars *d8partitionVarsT) {
	var seed0, seed1 int
	var worst, waste float64
	var area [d8maxNodes + 1]float64

	for index := 0; index < parVars.total; index++ {
		area[index] = d8calcRectVolume(&parVars.branchBuf[index].rect)
	}

	worst = -parVars.coverSplitArea - 1
	for indexA := 0; indexA < parVars.total-1; indexA++ {
		for indexB := indexA + 1; indexB < parVars.total; indexB++ {
			oneRect := d8combineRect(&parVars.branchBuf[indexA].rect, &parVars.branchBuf[indexB].rect)
			waste = d8calcRectVolume(&oneRect) - area[indexA] - area[indexB]
			if waste > worst {
				worst = waste
				seed0 = indexA
				seed1 = indexB
			}
		}
	}

	d8classify(seed0, 0, parVars)
	d8classify(seed1, 1, parVars)
}

// Put a branch in one of the groups.
func d8classify(index, group int, parVars *d8partitionVarsT) {
	parVars.partition[index] = group

	// Calculate combined rect
	if parVars.count[group] == 0 {
		parVars.cover[group] = parVars.branchBuf[index].rect
	} else {
		parVars.cover[group] = d8combineRect(&parVars.branchBuf[index].rect, &parVars.cover[group])
	}

	// Calculate volume of combined rect
	parVars.area[group] = d8calcRectVolume(&parVars.cover[group])

	parVars.count[group]++
}

// Delete a data rectangle from an index structure.
// Pass in a pointer to a d8rectT, the tid of the record, ptr to ptr to root node.
// Returns 1 if record not found, 0 if success.
// d8removeRect provides for eliminating the root.
func d8removeRect(rect *d8rectT, id interface{}, root **d8nodeT) bool {
	var reInsertList *d8listNodeT

	if !d8removeRectRec(rect, id, *root, &reInsertList) {
		// Found and deleted a data item
		// Reinsert any branches from eliminated nodes
		for reInsertList != nil {
			tempNode := reInsertList.node

			for index := 0; index < tempNode.count; index++ {
				// TODO go over this code. should I use (tempNode->m_level - 1)?
				d8insertRect(&tempNode.branch[index], root, tempNode.level)
			}
			reInsertList = reInsertList.next
		}

		// Check for redundant root (not leaf, 1 child) and eliminate TODO replace
		// if with while? In case there is a whole branch of redundant roots...
		if (*root).count == 1 && (*root).isInternalNode() {
			tempNode := (*root).branch[0].child
			*root = tempNode
		}
		return false
	} else {
		return true
	}
}

// Delete a rectangle from non-root part of an index structure.
// Called by d8removeRect.  Descends tree recursively,
// merges branches on the way back up.
// Returns 1 if record not found, 0 if success.
func d8removeRectRec(rect *d8rectT, id interface{}, node *d8nodeT, listNode **d8listNodeT) bool {
	if node.isInternalNode() { // not a leaf node
		for index := 0; index < node.count; index++ {
			if d8overlap(*rect, node.branch[index].rect) {
				if !d8removeRectRec(rect, id, node.branch[index].child, listNode) {
					if node.branch[index].child.count >= d8minNodes {
						// child removed, just resize parent rect
						node.branch[index].rect = d8nodeCover(node.branch[index].child)
					} else {
						// child removed, not enough entries in node, eliminate node
						d8reInsert(node.branch[index].child, listNode)
						d8disconnectBranch(node, index) // Must return after this call as count has changed
					}
					return false
				}
			}
		}
		return true
	} else { // A leaf node
		for index := 0; index < node.count; index++ {
			if node.branch[index].data == id {
				d8disconnectBranch(node, index) // Must return after this call as count has changed
				return false
			}
		}
		return true
	}
}

// Decide whether two rectangles d8overlap.
func d8overlap(rectA, rectB d8rectT) bool {
	for index := 0; index < d8numDims; index++ {
		if rectA.min[index] > rectB.max[index] ||
			rectB.min[index] > rectA.max[index] {
			return false
		}
	}
	return true
}

// Add a node to the reinsertion list.  All its branches will later
// be reinserted into the index structure.
func d8reInsert(node *d8nodeT, listNode **d8listNodeT) {
	newListNode := &d8listNodeT{}
	newListNode.node = node
	newListNode.next = *listNode
	*listNode = newListNode
}

// d8search in an index tree or subtree for all data retangles that d8overlap the argument rectangle.
func d8search(node *d8nodeT, rect d8rectT, foundCount int, resultCallback func(data interface{}) bool) (int, bool) {
	if node.isInternalNode() {
		// This is an internal node in the tree
		for index := 0; index < node.count; index++ {
			if d8overlap(rect, node.branch[index].rect) {
				var ok bool
				foundCount, ok = d8search(node.branch[index].child, rect, foundCount, resultCallback)
				if !ok {
					// The callback indicated to stop searching
					return foundCount, false
				}
			}
		}
	} else {
		// This is a leaf node
		for index := 0; index < node.count; index++ {
			if d8overlap(rect, node.branch[index].rect) {
				id := node.branch[index].data
				foundCount++
				if !resultCallback(id) {
					return foundCount, false // Don't continue searching
				}

			}
		}
	}
	return foundCount, true // Continue searching
}

func d9fmin(a, b float64) float64 {
	if a < b {
		return a
	}
	return b
}
func d9fmax(a, b float64) float64 {
	if a > b {
		return a
	}
	return b
}

const (
	d9numDims            = 9
	d9maxNodes           = 8
	d9minNodes           = d9maxNodes / 2
	d9useSphericalVolume = true // Better split classification, may be slower on some systems
)

var d9unitSphereVolume = []float64{
	0.000000, 2.000000, 3.141593, // Dimension  0,1,2
	4.188790, 4.934802, 5.263789, // Dimension  3,4,5
	5.167713, 4.724766, 4.058712, // Dimension  6,7,8
	3.298509, 2.550164, 1.884104, // Dimension  9,10,11
	1.335263, 0.910629, 0.599265, // Dimension  12,13,14
	0.381443, 0.235331, 0.140981, // Dimension  15,16,17
	0.082146, 0.046622, 0.025807, // Dimension  18,19,20
}[d9numDims]

type d9RTree struct {
	root *d9nodeT ///< Root of tree
}

/// Minimal bounding rectangle (n-dimensional)
type d9rectT struct {
	min [d9numDims]float64 ///< Min dimensions of bounding box
	max [d9numDims]float64 ///< Max dimensions of bounding box
}

/// May be data or may be another subtree
/// The parents level determines this.
/// If the parents level is 0, then this is data
type d9branchT struct {
	rect  d9rectT     ///< Bounds
	child *d9nodeT    ///< Child node
	data  interface{} ///< Data Id or Ptr
}

/// d9nodeT for each branch level
type d9nodeT struct {
	count  int                   ///< Count
	level  int                   ///< Leaf is zero, others positive
	branch [d9maxNodes]d9branchT ///< Branch
}

func (node *d9nodeT) isInternalNode() bool {
	return (node.level > 0) // Not a leaf, but a internal node
}
func (node *d9nodeT) isLeaf() bool {
	return (node.level == 0) // A leaf, contains data
}

/// A link list of nodes for reinsertion after a delete operation
type d9listNodeT struct {
	next *d9listNodeT ///< Next in list
	node *d9nodeT     ///< Node
}

const d9notTaken = -1 // indicates that position

/// Variables for finding a split partition
type d9partitionVarsT struct {
	partition [d9maxNodes + 1]int
	total     int
	minFill   int
	count     [2]int
	cover     [2]d9rectT
	area      [2]float64

	branchBuf      [d9maxNodes + 1]d9branchT
	branchCount    int
	coverSplit     d9rectT
	coverSplitArea float64
}

func d9New() *d9RTree {
	// We only support machine word size simple data type eg. integer index or object pointer.
	// Since we are storing as union with non data branch
	return &d9RTree{
		root: &d9nodeT{},
	}
}

/// Insert entry
/// \param a_min Min of bounding rect
/// \param a_max Max of bounding rect
/// \param a_dataId Positive Id of data.  Maybe zero, but negative numbers not allowed.
func (tr *d9RTree) Insert(min, max [d9numDims]float64, dataId interface{}) {
	var branch d9branchT
	branch.data = dataId
	for axis := 0; axis < d9numDims; axis++ {
		branch.rect.min[axis] = min[axis]
		branch.rect.max[axis] = max[axis]
	}
	d9insertRect(&branch, &tr.root, 0)
}

/// Remove entry
/// \param a_min Min of bounding rect
/// \param a_max Max of bounding rect
/// \param a_dataId Positive Id of data.  Maybe zero, but negative numbers not allowed.
func (tr *d9RTree) Remove(min, max [d9numDims]float64, dataId interface{}) {
	var rect d9rectT
	for axis := 0; axis < d9numDims; axis++ {
		rect.min[axis] = min[axis]
		rect.max[axis] = max[axis]
	}
	d9removeRect(&rect, dataId, &tr.root)
}

/// Find all within d9search rectangle
/// \param a_min Min of d9search bounding rect
/// \param a_max Max of d9search bounding rect
/// \param a_searchResult d9search result array.  Caller should set grow size. Function will reset, not append to array.
/// \param a_resultCallback Callback function to return result.  Callback should return 'true' to continue searching
/// \param a_context User context to pass as parameter to a_resultCallback
/// \return Returns the number of entries found
func (tr *d9RTree) Search(min, max [d9numDims]float64, resultCallback func(data interface{}) bool) int {
	var rect d9rectT
	for axis := 0; axis < d9numDims; axis++ {
		rect.min[axis] = min[axis]
		rect.max[axis] = max[axis]
	}
	foundCount, _ := d9search(tr.root, rect, 0, resultCallback)
	return foundCount
}

/// Count the data elements in this container.  This is slow as no internal counter is maintained.
func (tr *d9RTree) Count() int {
	var count int
	d9countRec(tr.root, &count)
	return count
}

/// Remove all entries from tree
func (tr *d9RTree) RemoveAll() {
	// Delete all existing nodes
	tr.root = &d9nodeT{}
}

func d9countRec(node *d9nodeT, count *int) {
	if node.isInternalNode() { // not a leaf node
		for index := 0; index < node.count; index++ {
			d9countRec(node.branch[index].child, count)
		}
	} else { // A leaf node
		*count += node.count
	}
}

// Inserts a new data rectangle into the index structure.
// Recursively descends tree, propagates splits back up.
// Returns 0 if node was not split.  Old node updated.
// If node was split, returns 1 and sets the pointer pointed to by
// new_node to point to the new node.  Old node updated to become one of two.
// The level argument specifies the number of steps up from the leaf
// level to insert; e.g. a data rectangle goes in at level = 0.
func d9insertRectRec(branch *d9branchT, node *d9nodeT, newNode **d9nodeT, level int) bool {
	// recurse until we reach the correct level for the new record. data records
	// will always be called with a_level == 0 (leaf)
	if node.level > level {
		// Still above level for insertion, go down tree recursively
		var otherNode *d9nodeT
		//var newBranch d9branchT

		// find the optimal branch for this record
		index := d9pickBranch(&branch.rect, node)

		// recursively insert this record into the picked branch
		childWasSplit := d9insertRectRec(branch, node.branch[index].child, &otherNode, level)

		if !childWasSplit {
			// Child was not split. Merge the bounding box of the new record with the
			// existing bounding box
			node.branch[index].rect = d9combineRect(&branch.rect, &(node.branch[index].rect))
			return false
		} else {
			// Child was split. The old branches are now re-partitioned to two nodes
			// so we have to re-calculate the bounding boxes of each node
			node.branch[index].rect = d9nodeCover(node.branch[index].child)
			var newBranch d9branchT
			newBranch.child = otherNode
			newBranch.rect = d9nodeCover(otherNode)

			// The old node is already a child of a_node. Now add the newly-created
			// node to a_node as well. a_node might be split because of that.
			return d9addBranch(&newBranch, node, newNode)
		}
	} else if node.level == level {
		// We have reached level for insertion. Add rect, split if necessary
		return d9addBranch(branch, node, newNode)
	} else {
		// Should never occur
		return false
	}
}

// Insert a data rectangle into an index structure.
// d9insertRect provides for splitting the root;
// returns 1 if root was split, 0 if it was not.
// The level argument specifies the number of steps up from the leaf
// level to insert; e.g. a data rectangle goes in at level = 0.
// InsertRect2 does the recursion.
//
func d9insertRect(branch *d9branchT, root **d9nodeT, level int) bool {
	var newNode *d9nodeT

	if d9insertRectRec(branch, *root, &newNode, level) { // Root split

		// Grow tree taller and new root
		newRoot := &d9nodeT{}
		newRoot.level = (*root).level + 1

		var newBranch d9branchT

		// add old root node as a child of the new root
		newBranch.rect = d9nodeCover(*root)
		newBranch.child = *root
		d9addBranch(&newBranch, newRoot, nil)

		// add the split node as a child of the new root
		newBranch.rect = d9nodeCover(newNode)
		newBranch.child = newNode
		d9addBranch(&newBranch, newRoot, nil)

		// set the new root as the root node
		*root = newRoot

		return true
	}
	return false
}

// Find the smallest rectangle that includes all rectangles in branches of a node.
func d9nodeCover(node *d9nodeT) d9rectT {
	rect := node.branch[0].rect
	for index := 1; index < node.count; index++ {
		rect = d9combineRect(&rect, &(node.branch[index].rect))
	}
	return rect
}

// Add a branch to a node.  Split the node if necessary.
// Returns 0 if node not split.  Old node updated.
// Returns 1 if node split, sets *new_node to address of new node.
// Old node updated, becomes one of two.
func d9addBranch(branch *d9branchT, node *d9nodeT, newNode **d9nodeT) bool {
	if node.count < d9maxNodes { // Split won't be necessary
		node.branch[node.count] = *branch
		node.count++
		return false
	} else {
		d9splitNode(node, branch, newNode)
		return true
	}
}

// Disconnect a dependent node.
// Caller must return (or stop using iteration index) after this as count has changed
func d9disconnectBranch(node *d9nodeT, index int) {
	// Remove element by swapping with the last element to prevent gaps in array
	node.branch[index] = node.branch[node.count-1]
	node.branch[node.count-1].data = nil
	node.branch[node.count-1].child = nil
	node.count--
}

// Pick a branch.  Pick the one that will need the smallest increase
// in area to accomodate the new rectangle.  This will result in the
// least total area for the covering rectangles in the current node.
// In case of a tie, pick the one which was smaller before, to get
// the best resolution when searching.
func d9pickBranch(rect *d9rectT, node *d9nodeT) int {
	var firstTime bool = true
	var increase float64
	var bestIncr float64 = -1
	var area float64
	var bestArea float64
	var best int
	var tempRect d9rectT

	for index := 0; index < node.count; index++ {
		curRect := &node.branch[index].rect
		area = d9calcRectVolume(curRect)
		tempRect = d9combineRect(rect, curRect)
		increase = d9calcRectVolume(&tempRect) - area
		if (increase < bestIncr) || firstTime {
			best = index
			bestArea = area
			bestIncr = increase
			firstTime = false
		} else if (increase == bestIncr) && (area < bestArea) {
			best = index
			bestArea = area
			bestIncr = increase
		}
	}
	return best
}

// Combine two rectangles into larger one containing both
func d9combineRect(rectA, rectB *d9rectT) d9rectT {
	var newRect d9rectT

	for index := 0; index < d9numDims; index++ {
		newRect.min[index] = d9fmin(rectA.min[index], rectB.min[index])
		newRect.max[index] = d9fmax(rectA.max[index], rectB.max[index])
	}

	return newRect
}

// Split a node.
// Divides the nodes branches and the extra one between two nodes.
// Old node is one of the new ones, and one really new one is created.
// Tries more than one method for choosing a partition, uses best result.
func d9splitNode(node *d9nodeT, branch *d9branchT, newNode **d9nodeT) {
	// Could just use local here, but member or external is faster since it is reused
	var localVars d9partitionVarsT
	parVars := &localVars

	// Load all the branches into a buffer, initialize old node
	d9getBranches(node, branch, parVars)

	// Find partition
	d9choosePartition(parVars, d9minNodes)

	// Create a new node to hold (about) half of the branches
	*newNode = &d9nodeT{}
	(*newNode).level = node.level

	// Put branches from buffer into 2 nodes according to the chosen partition
	node.count = 0
	d9loadNodes(node, *newNode, parVars)
}

// Calculate the n-dimensional volume of a rectangle
func d9rectVolume(rect *d9rectT) float64 {
	var volume float64 = 1
	for index := 0; index < d9numDims; index++ {
		volume *= rect.max[index] - rect.min[index]
	}
	return volume
}

// The exact volume of the bounding sphere for the given d9rectT
func d9rectSphericalVolume(rect *d9rectT) float64 {
	var sumOfSquares float64 = 0
	var radius float64

	for index := 0; index < d9numDims; index++ {
		halfExtent := (rect.max[index] - rect.min[index]) * 0.5
		sumOfSquares += halfExtent * halfExtent
	}

	radius = math.Sqrt(sumOfSquares)

	// Pow maybe slow, so test for common dims just use x*x, x*x*x.
	if d9numDims == 5 {
		return (radius * radius * radius * radius * radius * d9unitSphereVolume)
	} else if d9numDims == 4 {
		return (radius * radius * radius * radius * d9unitSphereVolume)
	} else if d9numDims == 3 {
		return (radius * radius * radius * d9unitSphereVolume)
	} else if d9numDims == 2 {
		return (radius * radius * d9unitSphereVolume)
	} else {
		return (math.Pow(radius, d9numDims) * d9unitSphereVolume)
	}
}

// Use one of the methods to calculate retangle volume
func d9calcRectVolume(rect *d9rectT) float64 {
	if d9useSphericalVolume {
		return d9rectSphericalVolume(rect) // Slower but helps certain merge cases
	} else { // RTREE_USE_SPHERICAL_VOLUME
		return d9rectVolume(rect) // Faster but can cause poor merges
	} // RTREE_USE_SPHERICAL_VOLUME
}

// Load branch buffer with branches from full node plus the extra branch.
func d9getBranches(node *d9nodeT, branch *d9branchT, parVars *d9partitionVarsT) {
	// Load the branch buffer
	for index := 0; index < d9maxNodes; index++ {
		parVars.branchBuf[index] = node.branch[index]
	}
	parVars.branchBuf[d9maxNodes] = *branch
	parVars.branchCount = d9maxNodes + 1

	// Calculate rect containing all in the set
	parVars.coverSplit = parVars.branchBuf[0].rect
	for index := 1; index < d9maxNodes+1; index++ {
		parVars.coverSplit = d9combineRect(&parVars.coverSplit, &parVars.branchBuf[index].rect)
	}
	parVars.coverSplitArea = d9calcRectVolume(&parVars.coverSplit)
}

// Method #0 for choosing a partition:
// As the seeds for the two groups, pick the two rects that would waste the
// most area if covered by a single rectangle, i.e. evidently the worst pair
// to have in the same group.
// Of the remaining, one at a time is chosen to be put in one of the two groups.
// The one chosen is the one with the greatest difference in area expansion
// depending on which group - the rect most strongly attracted to one group
// and repelled from the other.
// If one group gets too full (more would force other group to violate min
// fill requirement) then other group gets the rest.
// These last are the ones that can go in either group most easily.
func d9choosePartition(parVars *d9partitionVarsT, minFill int) {
	var biggestDiff float64
	var group, chosen, betterGroup int

	d9initParVars(parVars, parVars.branchCount, minFill)
	d9pickSeeds(parVars)

	for ((parVars.count[0] + parVars.count[1]) < parVars.total) &&
		(parVars.count[0] < (parVars.total - parVars.minFill)) &&
		(parVars.count[1] < (parVars.total - parVars.minFill)) {
		biggestDiff = -1
		for index := 0; index < parVars.total; index++ {
			if d9notTaken == parVars.partition[index] {
				curRect := &parVars.branchBuf[index].rect
				rect0 := d9combineRect(curRect, &parVars.cover[0])
				rect1 := d9combineRect(curRect, &parVars.cover[1])
				growth0 := d9calcRectVolume(&rect0) - parVars.area[0]
				growth1 := d9calcRectVolume(&rect1) - parVars.area[1]
				diff := growth1 - growth0
				if diff >= 0 {
					group = 0
				} else {
					group = 1
					diff = -diff
				}

				if diff > biggestDiff {
					biggestDiff = diff
					chosen = index
					betterGroup = group
				} else if (diff == biggestDiff) && (parVars.count[group] < parVars.count[betterGroup]) {
					chosen = index
					betterGroup = group
				}
			}
		}
		d9classify(chosen, betterGroup, parVars)
	}

	// If one group too full, put remaining rects in the other
	if (parVars.count[0] + parVars.count[1]) < parVars.total {
		if parVars.count[0] >= parVars.total-parVars.minFill {
			group = 1
		} else {
			group = 0
		}
		for index := 0; index < parVars.total; index++ {
			if d9notTaken == parVars.partition[index] {
				d9classify(index, group, parVars)
			}
		}
	}
}

// Copy branches from the buffer into two nodes according to the partition.
func d9loadNodes(nodeA, nodeB *d9nodeT, parVars *d9partitionVarsT) {
	for index := 0; index < parVars.total; index++ {
		targetNodeIndex := parVars.partition[index]
		targetNodes := []*d9nodeT{nodeA, nodeB}

		// It is assured that d9addBranch here will not cause a node split.
		d9addBranch(&parVars.branchBuf[index], targetNodes[targetNodeIndex], nil)
	}
}

// Initialize a d9partitionVarsT structure.
func d9initParVars(parVars *d9partitionVarsT, maxRects, minFill int) {
	parVars.count[0] = 0
	parVars.count[1] = 0
	parVars.area[0] = 0
	parVars.area[1] = 0
	parVars.total = maxRects
	parVars.minFill = minFill
	for index := 0; index < maxRects; index++ {
		parVars.partition[index] = d9notTaken
	}
}

func d9pickSeeds(parVars *d9partitionVarsT) {
	var seed0, seed1 int
	var worst, waste float64
	var area [d9maxNodes + 1]float64

	for index := 0; index < parVars.total; index++ {
		area[index] = d9calcRectVolume(&parVars.branchBuf[index].rect)
	}

	worst = -parVars.coverSplitArea - 1
	for indexA := 0; indexA < parVars.total-1; indexA++ {
		for indexB := indexA + 1; indexB < parVars.total; indexB++ {
			oneRect := d9combineRect(&parVars.branchBuf[indexA].rect, &parVars.branchBuf[indexB].rect)
			waste = d9calcRectVolume(&oneRect) - area[indexA] - area[indexB]
			if waste > worst {
				worst = waste
				seed0 = indexA
				seed1 = indexB
			}
		}
	}

	d9classify(seed0, 0, parVars)
	d9classify(seed1, 1, parVars)
}

// Put a branch in one of the groups.
func d9classify(index, group int, parVars *d9partitionVarsT) {
	parVars.partition[index] = group

	// Calculate combined rect
	if parVars.count[group] == 0 {
		parVars.cover[group] = parVars.branchBuf[index].rect
	} else {
		parVars.cover[group] = d9combineRect(&parVars.branchBuf[index].rect, &parVars.cover[group])
	}

	// Calculate volume of combined rect
	parVars.area[group] = d9calcRectVolume(&parVars.cover[group])

	parVars.count[group]++
}

// Delete a data rectangle from an index structure.
// Pass in a pointer to a d9rectT, the tid of the record, ptr to ptr to root node.
// Returns 1 if record not found, 0 if success.
// d9removeRect provides for eliminating the root.
func d9removeRect(rect *d9rectT, id interface{}, root **d9nodeT) bool {
	var reInsertList *d9listNodeT

	if !d9removeRectRec(rect, id, *root, &reInsertList) {
		// Found and deleted a data item
		// Reinsert any branches from eliminated nodes
		for reInsertList != nil {
			tempNode := reInsertList.node

			for index := 0; index < tempNode.count; index++ {
				// TODO go over this code. should I use (tempNode->m_level - 1)?
				d9insertRect(&tempNode.branch[index], root, tempNode.level)
			}
			reInsertList = reInsertList.next
		}

		// Check for redundant root (not leaf, 1 child) and eliminate TODO replace
		// if with while? In case there is a whole branch of redundant roots...
		if (*root).count == 1 && (*root).isInternalNode() {
			tempNode := (*root).branch[0].child
			*root = tempNode
		}
		return false
	} else {
		return true
	}
}

// Delete a rectangle from non-root part of an index structure.
// Called by d9removeRect.  Descends tree recursively,
// merges branches on the way back up.
// Returns 1 if record not found, 0 if success.
func d9removeRectRec(rect *d9rectT, id interface{}, node *d9nodeT, listNode **d9listNodeT) bool {
	if node.isInternalNode() { // not a leaf node
		for index := 0; index < node.count; index++ {
			if d9overlap(*rect, node.branch[index].rect) {
				if !d9removeRectRec(rect, id, node.branch[index].child, listNode) {
					if node.branch[index].child.count >= d9minNodes {
						// child removed, just resize parent rect
						node.branch[index].rect = d9nodeCover(node.branch[index].child)
					} else {
						// child removed, not enough entries in node, eliminate node
						d9reInsert(node.branch[index].child, listNode)
						d9disconnectBranch(node, index) // Must return after this call as count has changed
					}
					return false
				}
			}
		}
		return true
	} else { // A leaf node
		for index := 0; index < node.count; index++ {
			if node.branch[index].data == id {
				d9disconnectBranch(node, index) // Must return after this call as count has changed
				return false
			}
		}
		return true
	}
}

// Decide whether two rectangles d9overlap.
func d9overlap(rectA, rectB d9rectT) bool {
	for index := 0; index < d9numDims; index++ {
		if rectA.min[index] > rectB.max[index] ||
			rectB.min[index] > rectA.max[index] {
			return false
		}
	}
	return true
}

// Add a node to the reinsertion list.  All its branches will later
// be reinserted into the index structure.
func d9reInsert(node *d9nodeT, listNode **d9listNodeT) {
	newListNode := &d9listNodeT{}
	newListNode.node = node
	newListNode.next = *listNode
	*listNode = newListNode
}

// d9search in an index tree or subtree for all data retangles that d9overlap the argument rectangle.
func d9search(node *d9nodeT, rect d9rectT, foundCount int, resultCallback func(data interface{}) bool) (int, bool) {
	if node.isInternalNode() {
		// This is an internal node in the tree
		for index := 0; index < node.count; index++ {
			if d9overlap(rect, node.branch[index].rect) {
				var ok bool
				foundCount, ok = d9search(node.branch[index].child, rect, foundCount, resultCallback)
				if !ok {
					// The callback indicated to stop searching
					return foundCount, false
				}
			}
		}
	} else {
		// This is a leaf node
		for index := 0; index < node.count; index++ {
			if d9overlap(rect, node.branch[index].rect) {
				id := node.branch[index].data
				foundCount++
				if !resultCallback(id) {
					return foundCount, false // Don't continue searching
				}

			}
		}
	}
	return foundCount, true // Continue searching
}

func d10fmin(a, b float64) float64 {
	if a < b {
		return a
	}
	return b
}
func d10fmax(a, b float64) float64 {
	if a > b {
		return a
	}
	return b
}

const (
	d10numDims            = 10
	d10maxNodes           = 8
	d10minNodes           = d10maxNodes / 2
	d10useSphericalVolume = true // Better split classification, may be slower on some systems
)

var d10unitSphereVolume = []float64{
	0.000000, 2.000000, 3.141593, // Dimension  0,1,2
	4.188790, 4.934802, 5.263789, // Dimension  3,4,5
	5.167713, 4.724766, 4.058712, // Dimension  6,7,8
	3.298509, 2.550164, 1.884104, // Dimension  9,10,11
	1.335263, 0.910629, 0.599265, // Dimension  12,13,14
	0.381443, 0.235331, 0.140981, // Dimension  15,16,17
	0.082146, 0.046622, 0.025807, // Dimension  18,19,20
}[d10numDims]

type d10RTree struct {
	root *d10nodeT ///< Root of tree
}

/// Minimal bounding rectangle (n-dimensional)
type d10rectT struct {
	min [d10numDims]float64 ///< Min dimensions of bounding box
	max [d10numDims]float64 ///< Max dimensions of bounding box
}

/// May be data or may be another subtree
/// The parents level determines this.
/// If the parents level is 0, then this is data
type d10branchT struct {
	rect  d10rectT    ///< Bounds
	child *d10nodeT   ///< Child node
	data  interface{} ///< Data Id or Ptr
}

/// d10nodeT for each branch level
type d10nodeT struct {
	count  int                     ///< Count
	level  int                     ///< Leaf is zero, others positive
	branch [d10maxNodes]d10branchT ///< Branch
}

func (node *d10nodeT) isInternalNode() bool {
	return (node.level > 0) // Not a leaf, but a internal node
}
func (node *d10nodeT) isLeaf() bool {
	return (node.level == 0) // A leaf, contains data
}

/// A link list of nodes for reinsertion after a delete operation
type d10listNodeT struct {
	next *d10listNodeT ///< Next in list
	node *d10nodeT     ///< Node
}

const d10notTaken = -1 // indicates that position

/// Variables for finding a split partition
type d10partitionVarsT struct {
	partition [d10maxNodes + 1]int
	total     int
	minFill   int
	count     [2]int
	cover     [2]d10rectT
	area      [2]float64

	branchBuf      [d10maxNodes + 1]d10branchT
	branchCount    int
	coverSplit     d10rectT
	coverSplitArea float64
}

func d10New() *d10RTree {
	// We only support machine word size simple data type eg. integer index or object pointer.
	// Since we are storing as union with non data branch
	return &d10RTree{
		root: &d10nodeT{},
	}
}

/// Insert entry
/// \param a_min Min of bounding rect
/// \param a_max Max of bounding rect
/// \param a_dataId Positive Id of data.  Maybe zero, but negative numbers not allowed.
func (tr *d10RTree) Insert(min, max [d10numDims]float64, dataId interface{}) {
	var branch d10branchT
	branch.data = dataId
	for axis := 0; axis < d10numDims; axis++ {
		branch.rect.min[axis] = min[axis]
		branch.rect.max[axis] = max[axis]
	}
	d10insertRect(&branch, &tr.root, 0)
}

/// Remove entry
/// \param a_min Min of bounding rect
/// \param a_max Max of bounding rect
/// \param a_dataId Positive Id of data.  Maybe zero, but negative numbers not allowed.
func (tr *d10RTree) Remove(min, max [d10numDims]float64, dataId interface{}) {
	var rect d10rectT
	for axis := 0; axis < d10numDims; axis++ {
		rect.min[axis] = min[axis]
		rect.max[axis] = max[axis]
	}
	d10removeRect(&rect, dataId, &tr.root)
}

/// Find all within d10search rectangle
/// \param a_min Min of d10search bounding rect
/// \param a_max Max of d10search bounding rect
/// \param a_searchResult d10search result array.  Caller should set grow size. Function will reset, not append to array.
/// \param a_resultCallback Callback function to return result.  Callback should return 'true' to continue searching
/// \param a_context User context to pass as parameter to a_resultCallback
/// \return Returns the number of entries found
func (tr *d10RTree) Search(min, max [d10numDims]float64, resultCallback func(data interface{}) bool) int {
	var rect d10rectT
	for axis := 0; axis < d10numDims; axis++ {
		rect.min[axis] = min[axis]
		rect.max[axis] = max[axis]
	}
	foundCount, _ := d10search(tr.root, rect, 0, resultCallback)
	return foundCount
}

/// Count the data elements in this container.  This is slow as no internal counter is maintained.
func (tr *d10RTree) Count() int {
	var count int
	d10countRec(tr.root, &count)
	return count
}

/// Remove all entries from tree
func (tr *d10RTree) RemoveAll() {
	// Delete all existing nodes
	tr.root = &d10nodeT{}
}

func d10countRec(node *d10nodeT, count *int) {
	if node.isInternalNode() { // not a leaf node
		for index := 0; index < node.count; index++ {
			d10countRec(node.branch[index].child, count)
		}
	} else { // A leaf node
		*count += node.count
	}
}

// Inserts a new data rectangle into the index structure.
// Recursively descends tree, propagates splits back up.
// Returns 0 if node was not split.  Old node updated.
// If node was split, returns 1 and sets the pointer pointed to by
// new_node to point to the new node.  Old node updated to become one of two.
// The level argument specifies the number of steps up from the leaf
// level to insert; e.g. a data rectangle goes in at level = 0.
func d10insertRectRec(branch *d10branchT, node *d10nodeT, newNode **d10nodeT, level int) bool {
	// recurse until we reach the correct level for the new record. data records
	// will always be called with a_level == 0 (leaf)
	if node.level > level {
		// Still above level for insertion, go down tree recursively
		var otherNode *d10nodeT
		//var newBranch d10branchT

		// find the optimal branch for this record
		index := d10pickBranch(&branch.rect, node)

		// recursively insert this record into the picked branch
		childWasSplit := d10insertRectRec(branch, node.branch[index].child, &otherNode, level)

		if !childWasSplit {
			// Child was not split. Merge the bounding box of the new record with the
			// existing bounding box
			node.branch[index].rect = d10combineRect(&branch.rect, &(node.branch[index].rect))
			return false
		} else {
			// Child was split. The old branches are now re-partitioned to two nodes
			// so we have to re-calculate the bounding boxes of each node
			node.branch[index].rect = d10nodeCover(node.branch[index].child)
			var newBranch d10branchT
			newBranch.child = otherNode
			newBranch.rect = d10nodeCover(otherNode)

			// The old node is already a child of a_node. Now add the newly-created
			// node to a_node as well. a_node might be split because of that.
			return d10addBranch(&newBranch, node, newNode)
		}
	} else if node.level == level {
		// We have reached level for insertion. Add rect, split if necessary
		return d10addBranch(branch, node, newNode)
	} else {
		// Should never occur
		return false
	}
}

// Insert a data rectangle into an index structure.
// d10insertRect provides for splitting the root;
// returns 1 if root was split, 0 if it was not.
// The level argument specifies the number of steps up from the leaf
// level to insert; e.g. a data rectangle goes in at level = 0.
// InsertRect2 does the recursion.
//
func d10insertRect(branch *d10branchT, root **d10nodeT, level int) bool {
	var newNode *d10nodeT

	if d10insertRectRec(branch, *root, &newNode, level) { // Root split

		// Grow tree taller and new root
		newRoot := &d10nodeT{}
		newRoot.level = (*root).level + 1

		var newBranch d10branchT

		// add old root node as a child of the new root
		newBranch.rect = d10nodeCover(*root)
		newBranch.child = *root
		d10addBranch(&newBranch, newRoot, nil)

		// add the split node as a child of the new root
		newBranch.rect = d10nodeCover(newNode)
		newBranch.child = newNode
		d10addBranch(&newBranch, newRoot, nil)

		// set the new root as the root node
		*root = newRoot

		return true
	}
	return false
}

// Find the smallest rectangle that includes all rectangles in branches of a node.
func d10nodeCover(node *d10nodeT) d10rectT {
	rect := node.branch[0].rect
	for index := 1; index < node.count; index++ {
		rect = d10combineRect(&rect, &(node.branch[index].rect))
	}
	return rect
}

// Add a branch to a node.  Split the node if necessary.
// Returns 0 if node not split.  Old node updated.
// Returns 1 if node split, sets *new_node to address of new node.
// Old node updated, becomes one of two.
func d10addBranch(branch *d10branchT, node *d10nodeT, newNode **d10nodeT) bool {
	if node.count < d10maxNodes { // Split won't be necessary
		node.branch[node.count] = *branch
		node.count++
		return false
	} else {
		d10splitNode(node, branch, newNode)
		return true
	}
}

// Disconnect a dependent node.
// Caller must return (or stop using iteration index) after this as count has changed
func d10disconnectBranch(node *d10nodeT, index int) {
	// Remove element by swapping with the last element to prevent gaps in array
	node.branch[index] = node.branch[node.count-1]
	node.branch[node.count-1].data = nil
	node.branch[node.count-1].child = nil
	node.count--
}

// Pick a branch.  Pick the one that will need the smallest increase
// in area to accomodate the new rectangle.  This will result in the
// least total area for the covering rectangles in the current node.
// In case of a tie, pick the one which was smaller before, to get
// the best resolution when searching.
func d10pickBranch(rect *d10rectT, node *d10nodeT) int {
	var firstTime bool = true
	var increase float64
	var bestIncr float64 = -1
	var area float64
	var bestArea float64
	var best int
	var tempRect d10rectT

	for index := 0; index < node.count; index++ {
		curRect := &node.branch[index].rect
		area = d10calcRectVolume(curRect)
		tempRect = d10combineRect(rect, curRect)
		increase = d10calcRectVolume(&tempRect) - area
		if (increase < bestIncr) || firstTime {
			best = index
			bestArea = area
			bestIncr = increase
			firstTime = false
		} else if (increase == bestIncr) && (area < bestArea) {
			best = index
			bestArea = area
			bestIncr = increase
		}
	}
	return best
}

// Combine two rectangles into larger one containing both
func d10combineRect(rectA, rectB *d10rectT) d10rectT {
	var newRect d10rectT

	for index := 0; index < d10numDims; index++ {
		newRect.min[index] = d10fmin(rectA.min[index], rectB.min[index])
		newRect.max[index] = d10fmax(rectA.max[index], rectB.max[index])
	}

	return newRect
}

// Split a node.
// Divides the nodes branches and the extra one between two nodes.
// Old node is one of the new ones, and one really new one is created.
// Tries more than one method for choosing a partition, uses best result.
func d10splitNode(node *d10nodeT, branch *d10branchT, newNode **d10nodeT) {
	// Could just use local here, but member or external is faster since it is reused
	var localVars d10partitionVarsT
	parVars := &localVars

	// Load all the branches into a buffer, initialize old node
	d10getBranches(node, branch, parVars)

	// Find partition
	d10choosePartition(parVars, d10minNodes)

	// Create a new node to hold (about) half of the branches
	*newNode = &d10nodeT{}
	(*newNode).level = node.level

	// Put branches from buffer into 2 nodes according to the chosen partition
	node.count = 0
	d10loadNodes(node, *newNode, parVars)
}

// Calculate the n-dimensional volume of a rectangle
func d10rectVolume(rect *d10rectT) float64 {
	var volume float64 = 1
	for index := 0; index < d10numDims; index++ {
		volume *= rect.max[index] - rect.min[index]
	}
	return volume
}

// The exact volume of the bounding sphere for the given d10rectT
func d10rectSphericalVolume(rect *d10rectT) float64 {
	var sumOfSquares float64 = 0
	var radius float64

	for index := 0; index < d10numDims; index++ {
		halfExtent := (rect.max[index] - rect.min[index]) * 0.5
		sumOfSquares += halfExtent * halfExtent
	}

	radius = math.Sqrt(sumOfSquares)

	// Pow maybe slow, so test for common dims just use x*x, x*x*x.
	if d10numDims == 5 {
		return (radius * radius * radius * radius * radius * d10unitSphereVolume)
	} else if d10numDims == 4 {
		return (radius * radius * radius * radius * d10unitSphereVolume)
	} else if d10numDims == 3 {
		return (radius * radius * radius * d10unitSphereVolume)
	} else if d10numDims == 2 {
		return (radius * radius * d10unitSphereVolume)
	} else {
		return (math.Pow(radius, d10numDims) * d10unitSphereVolume)
	}
}

// Use one of the methods to calculate retangle volume
func d10calcRectVolume(rect *d10rectT) float64 {
	if d10useSphericalVolume {
		return d10rectSphericalVolume(rect) // Slower but helps certain merge cases
	} else { // RTREE_USE_SPHERICAL_VOLUME
		return d10rectVolume(rect) // Faster but can cause poor merges
	} // RTREE_USE_SPHERICAL_VOLUME
}

// Load branch buffer with branches from full node plus the extra branch.
func d10getBranches(node *d10nodeT, branch *d10branchT, parVars *d10partitionVarsT) {
	// Load the branch buffer
	for index := 0; index < d10maxNodes; index++ {
		parVars.branchBuf[index] = node.branch[index]
	}
	parVars.branchBuf[d10maxNodes] = *branch
	parVars.branchCount = d10maxNodes + 1

	// Calculate rect containing all in the set
	parVars.coverSplit = parVars.branchBuf[0].rect
	for index := 1; index < d10maxNodes+1; index++ {
		parVars.coverSplit = d10combineRect(&parVars.coverSplit, &parVars.branchBuf[index].rect)
	}
	parVars.coverSplitArea = d10calcRectVolume(&parVars.coverSplit)
}

// Method #0 for choosing a partition:
// As the seeds for the two groups, pick the two rects that would waste the
// most area if covered by a single rectangle, i.e. evidently the worst pair
// to have in the same group.
// Of the remaining, one at a time is chosen to be put in one of the two groups.
// The one chosen is the one with the greatest difference in area expansion
// depending on which group - the rect most strongly attracted to one group
// and repelled from the other.
// If one group gets too full (more would force other group to violate min
// fill requirement) then other group gets the rest.
// These last are the ones that can go in either group most easily.
func d10choosePartition(parVars *d10partitionVarsT, minFill int) {
	var biggestDiff float64
	var group, chosen, betterGroup int

	d10initParVars(parVars, parVars.branchCount, minFill)
	d10pickSeeds(parVars)

	for ((parVars.count[0] + parVars.count[1]) < parVars.total) &&
		(parVars.count[0] < (parVars.total - parVars.minFill)) &&
		(parVars.count[1] < (parVars.total - parVars.minFill)) {
		biggestDiff = -1
		for index := 0; index < parVars.total; index++ {
			if d10notTaken == parVars.partition[index] {
				curRect := &parVars.branchBuf[index].rect
				rect0 := d10combineRect(curRect, &parVars.cover[0])
				rect1 := d10combineRect(curRect, &parVars.cover[1])
				growth0 := d10calcRectVolume(&rect0) - parVars.area[0]
				growth1 := d10calcRectVolume(&rect1) - parVars.area[1]
				diff := growth1 - growth0
				if diff >= 0 {
					group = 0
				} else {
					group = 1
					diff = -diff
				}

				if diff > biggestDiff {
					biggestDiff = diff
					chosen = index
					betterGroup = group
				} else if (diff == biggestDiff) && (parVars.count[group] < parVars.count[betterGroup]) {
					chosen = index
					betterGroup = group
				}
			}
		}
		d10classify(chosen, betterGroup, parVars)
	}

	// If one group too full, put remaining rects in the other
	if (parVars.count[0] + parVars.count[1]) < parVars.total {
		if parVars.count[0] >= parVars.total-parVars.minFill {
			group = 1
		} else {
			group = 0
		}
		for index := 0; index < parVars.total; index++ {
			if d10notTaken == parVars.partition[index] {
				d10classify(index, group, parVars)
			}
		}
	}
}

// Copy branches from the buffer into two nodes according to the partition.
func d10loadNodes(nodeA, nodeB *d10nodeT, parVars *d10partitionVarsT) {
	for index := 0; index < parVars.total; index++ {
		targetNodeIndex := parVars.partition[index]
		targetNodes := []*d10nodeT{nodeA, nodeB}

		// It is assured that d10addBranch here will not cause a node split.
		d10addBranch(&parVars.branchBuf[index], targetNodes[targetNodeIndex], nil)
	}
}

// Initialize a d10partitionVarsT structure.
func d10initParVars(parVars *d10partitionVarsT, maxRects, minFill int) {
	parVars.count[0] = 0
	parVars.count[1] = 0
	parVars.area[0] = 0
	parVars.area[1] = 0
	parVars.total = maxRects
	parVars.minFill = minFill
	for index := 0; index < maxRects; index++ {
		parVars.partition[index] = d10notTaken
	}
}

func d10pickSeeds(parVars *d10partitionVarsT) {
	var seed0, seed1 int
	var worst, waste float64
	var area [d10maxNodes + 1]float64

	for index := 0; index < parVars.total; index++ {
		area[index] = d10calcRectVolume(&parVars.branchBuf[index].rect)
	}

	worst = -parVars.coverSplitArea - 1
	for indexA := 0; indexA < parVars.total-1; indexA++ {
		for indexB := indexA + 1; indexB < parVars.total; indexB++ {
			oneRect := d10combineRect(&parVars.branchBuf[indexA].rect, &parVars.branchBuf[indexB].rect)
			waste = d10calcRectVolume(&oneRect) - area[indexA] - area[indexB]
			if waste > worst {
				worst = waste
				seed0 = indexA
				seed1 = indexB
			}
		}
	}

	d10classify(seed0, 0, parVars)
	d10classify(seed1, 1, parVars)
}

// Put a branch in one of the groups.
func d10classify(index, group int, parVars *d10partitionVarsT) {
	parVars.partition[index] = group

	// Calculate combined rect
	if parVars.count[group] == 0 {
		parVars.cover[group] = parVars.branchBuf[index].rect
	} else {
		parVars.cover[group] = d10combineRect(&parVars.branchBuf[index].rect, &parVars.cover[group])
	}

	// Calculate volume of combined rect
	parVars.area[group] = d10calcRectVolume(&parVars.cover[group])

	parVars.count[group]++
}

// Delete a data rectangle from an index structure.
// Pass in a pointer to a d10rectT, the tid of the record, ptr to ptr to root node.
// Returns 1 if record not found, 0 if success.
// d10removeRect provides for eliminating the root.
func d10removeRect(rect *d10rectT, id interface{}, root **d10nodeT) bool {
	var reInsertList *d10listNodeT

	if !d10removeRectRec(rect, id, *root, &reInsertList) {
		// Found and deleted a data item
		// Reinsert any branches from eliminated nodes
		for reInsertList != nil {
			tempNode := reInsertList.node

			for index := 0; index < tempNode.count; index++ {
				// TODO go over this code. should I use (tempNode->m_level - 1)?
				d10insertRect(&tempNode.branch[index], root, tempNode.level)
			}
			reInsertList = reInsertList.next
		}

		// Check for redundant root (not leaf, 1 child) and eliminate TODO replace
		// if with while? In case there is a whole branch of redundant roots...
		if (*root).count == 1 && (*root).isInternalNode() {
			tempNode := (*root).branch[0].child
			*root = tempNode
		}
		return false
	} else {
		return true
	}
}

// Delete a rectangle from non-root part of an index structure.
// Called by d10removeRect.  Descends tree recursively,
// merges branches on the way back up.
// Returns 1 if record not found, 0 if success.
func d10removeRectRec(rect *d10rectT, id interface{}, node *d10nodeT, listNode **d10listNodeT) bool {
	if node.isInternalNode() { // not a leaf node
		for index := 0; index < node.count; index++ {
			if d10overlap(*rect, node.branch[index].rect) {
				if !d10removeRectRec(rect, id, node.branch[index].child, listNode) {
					if node.branch[index].child.count >= d10minNodes {
						// child removed, just resize parent rect
						node.branch[index].rect = d10nodeCover(node.branch[index].child)
					} else {
						// child removed, not enough entries in node, eliminate node
						d10reInsert(node.branch[index].child, listNode)
						d10disconnectBranch(node, index) // Must return after this call as count has changed
					}
					return false
				}
			}
		}
		return true
	} else { // A leaf node
		for index := 0; index < node.count; index++ {
			if node.branch[index].data == id {
				d10disconnectBranch(node, index) // Must return after this call as count has changed
				return false
			}
		}
		return true
	}
}

// Decide whether two rectangles d10overlap.
func d10overlap(rectA, rectB d10rectT) bool {
	for index := 0; index < d10numDims; index++ {
		if rectA.min[index] > rectB.max[index] ||
			rectB.min[index] > rectA.max[index] {
			return false
		}
	}
	return true
}

// Add a node to the reinsertion list.  All its branches will later
// be reinserted into the index structure.
func d10reInsert(node *d10nodeT, listNode **d10listNodeT) {
	newListNode := &d10listNodeT{}
	newListNode.node = node
	newListNode.next = *listNode
	*listNode = newListNode
}

// d10search in an index tree or subtree for all data retangles that d10overlap the argument rectangle.
func d10search(node *d10nodeT, rect d10rectT, foundCount int, resultCallback func(data interface{}) bool) (int, bool) {
	if node.isInternalNode() {
		// This is an internal node in the tree
		for index := 0; index < node.count; index++ {
			if d10overlap(rect, node.branch[index].rect) {
				var ok bool
				foundCount, ok = d10search(node.branch[index].child, rect, foundCount, resultCallback)
				if !ok {
					// The callback indicated to stop searching
					return foundCount, false
				}
			}
		}
	} else {
		// This is a leaf node
		for index := 0; index < node.count; index++ {
			if d10overlap(rect, node.branch[index].rect) {
				id := node.branch[index].data
				foundCount++
				if !resultCallback(id) {
					return foundCount, false // Don't continue searching
				}

			}
		}
	}
	return foundCount, true // Continue searching
}

func d11fmin(a, b float64) float64 {
	if a < b {
		return a
	}
	return b
}
func d11fmax(a, b float64) float64 {
	if a > b {
		return a
	}
	return b
}

const (
	d11numDims            = 11
	d11maxNodes           = 8
	d11minNodes           = d11maxNodes / 2
	d11useSphericalVolume = true // Better split classification, may be slower on some systems
)

var d11unitSphereVolume = []float64{
	0.000000, 2.000000, 3.141593, // Dimension  0,1,2
	4.188790, 4.934802, 5.263789, // Dimension  3,4,5
	5.167713, 4.724766, 4.058712, // Dimension  6,7,8
	3.298509, 2.550164, 1.884104, // Dimension  9,10,11
	1.335263, 0.910629, 0.599265, // Dimension  12,13,14
	0.381443, 0.235331, 0.140981, // Dimension  15,16,17
	0.082146, 0.046622, 0.025807, // Dimension  18,19,20
}[d11numDims]

type d11RTree struct {
	root *d11nodeT ///< Root of tree
}

/// Minimal bounding rectangle (n-dimensional)
type d11rectT struct {
	min [d11numDims]float64 ///< Min dimensions of bounding box
	max [d11numDims]float64 ///< Max dimensions of bounding box
}

/// May be data or may be another subtree
/// The parents level determines this.
/// If the parents level is 0, then this is data
type d11branchT struct {
	rect  d11rectT    ///< Bounds
	child *d11nodeT   ///< Child node
	data  interface{} ///< Data Id or Ptr
}

/// d11nodeT for each branch level
type d11nodeT struct {
	count  int                     ///< Count
	level  int                     ///< Leaf is zero, others positive
	branch [d11maxNodes]d11branchT ///< Branch
}

func (node *d11nodeT) isInternalNode() bool {
	return (node.level > 0) // Not a leaf, but a internal node
}
func (node *d11nodeT) isLeaf() bool {
	return (node.level == 0) // A leaf, contains data
}

/// A link list of nodes for reinsertion after a delete operation
type d11listNodeT struct {
	next *d11listNodeT ///< Next in list
	node *d11nodeT     ///< Node
}

const d11notTaken = -1 // indicates that position

/// Variables for finding a split partition
type d11partitionVarsT struct {
	partition [d11maxNodes + 1]int
	total     int
	minFill   int
	count     [2]int
	cover     [2]d11rectT
	area      [2]float64

	branchBuf      [d11maxNodes + 1]d11branchT
	branchCount    int
	coverSplit     d11rectT
	coverSplitArea float64
}

func d11New() *d11RTree {
	// We only support machine word size simple data type eg. integer index or object pointer.
	// Since we are storing as union with non data branch
	return &d11RTree{
		root: &d11nodeT{},
	}
}

/// Insert entry
/// \param a_min Min of bounding rect
/// \param a_max Max of bounding rect
/// \param a_dataId Positive Id of data.  Maybe zero, but negative numbers not allowed.
func (tr *d11RTree) Insert(min, max [d11numDims]float64, dataId interface{}) {
	var branch d11branchT
	branch.data = dataId
	for axis := 0; axis < d11numDims; axis++ {
		branch.rect.min[axis] = min[axis]
		branch.rect.max[axis] = max[axis]
	}
	d11insertRect(&branch, &tr.root, 0)
}

/// Remove entry
/// \param a_min Min of bounding rect
/// \param a_max Max of bounding rect
/// \param a_dataId Positive Id of data.  Maybe zero, but negative numbers not allowed.
func (tr *d11RTree) Remove(min, max [d11numDims]float64, dataId interface{}) {
	var rect d11rectT
	for axis := 0; axis < d11numDims; axis++ {
		rect.min[axis] = min[axis]
		rect.max[axis] = max[axis]
	}
	d11removeRect(&rect, dataId, &tr.root)
}

/// Find all within d11search rectangle
/// \param a_min Min of d11search bounding rect
/// \param a_max Max of d11search bounding rect
/// \param a_searchResult d11search result array.  Caller should set grow size. Function will reset, not append to array.
/// \param a_resultCallback Callback function to return result.  Callback should return 'true' to continue searching
/// \param a_context User context to pass as parameter to a_resultCallback
/// \return Returns the number of entries found
func (tr *d11RTree) Search(min, max [d11numDims]float64, resultCallback func(data interface{}) bool) int {
	var rect d11rectT
	for axis := 0; axis < d11numDims; axis++ {
		rect.min[axis] = min[axis]
		rect.max[axis] = max[axis]
	}
	foundCount, _ := d11search(tr.root, rect, 0, resultCallback)
	return foundCount
}

/// Count the data elements in this container.  This is slow as no internal counter is maintained.
func (tr *d11RTree) Count() int {
	var count int
	d11countRec(tr.root, &count)
	return count
}

/// Remove all entries from tree
func (tr *d11RTree) RemoveAll() {
	// Delete all existing nodes
	tr.root = &d11nodeT{}
}

func d11countRec(node *d11nodeT, count *int) {
	if node.isInternalNode() { // not a leaf node
		for index := 0; index < node.count; index++ {
			d11countRec(node.branch[index].child, count)
		}
	} else { // A leaf node
		*count += node.count
	}
}

// Inserts a new data rectangle into the index structure.
// Recursively descends tree, propagates splits back up.
// Returns 0 if node was not split.  Old node updated.
// If node was split, returns 1 and sets the pointer pointed to by
// new_node to point to the new node.  Old node updated to become one of two.
// The level argument specifies the number of steps up from the leaf
// level to insert; e.g. a data rectangle goes in at level = 0.
func d11insertRectRec(branch *d11branchT, node *d11nodeT, newNode **d11nodeT, level int) bool {
	// recurse until we reach the correct level for the new record. data records
	// will always be called with a_level == 0 (leaf)
	if node.level > level {
		// Still above level for insertion, go down tree recursively
		var otherNode *d11nodeT
		//var newBranch d11branchT

		// find the optimal branch for this record
		index := d11pickBranch(&branch.rect, node)

		// recursively insert this record into the picked branch
		childWasSplit := d11insertRectRec(branch, node.branch[index].child, &otherNode, level)

		if !childWasSplit {
			// Child was not split. Merge the bounding box of the new record with the
			// existing bounding box
			node.branch[index].rect = d11combineRect(&branch.rect, &(node.branch[index].rect))
			return false
		} else {
			// Child was split. The old branches are now re-partitioned to two nodes
			// so we have to re-calculate the bounding boxes of each node
			node.branch[index].rect = d11nodeCover(node.branch[index].child)
			var newBranch d11branchT
			newBranch.child = otherNode
			newBranch.rect = d11nodeCover(otherNode)

			// The old node is already a child of a_node. Now add the newly-created
			// node to a_node as well. a_node might be split because of that.
			return d11addBranch(&newBranch, node, newNode)
		}
	} else if node.level == level {
		// We have reached level for insertion. Add rect, split if necessary
		return d11addBranch(branch, node, newNode)
	} else {
		// Should never occur
		return false
	}
}

// Insert a data rectangle into an index structure.
// d11insertRect provides for splitting the root;
// returns 1 if root was split, 0 if it was not.
// The level argument specifies the number of steps up from the leaf
// level to insert; e.g. a data rectangle goes in at level = 0.
// InsertRect2 does the recursion.
//
func d11insertRect(branch *d11branchT, root **d11nodeT, level int) bool {
	var newNode *d11nodeT

	if d11insertRectRec(branch, *root, &newNode, level) { // Root split

		// Grow tree taller and new root
		newRoot := &d11nodeT{}
		newRoot.level = (*root).level + 1

		var newBranch d11branchT

		// add old root node as a child of the new root
		newBranch.rect = d11nodeCover(*root)
		newBranch.child = *root
		d11addBranch(&newBranch, newRoot, nil)

		// add the split node as a child of the new root
		newBranch.rect = d11nodeCover(newNode)
		newBranch.child = newNode
		d11addBranch(&newBranch, newRoot, nil)

		// set the new root as the root node
		*root = newRoot

		return true
	}
	return false
}

// Find the smallest rectangle that includes all rectangles in branches of a node.
func d11nodeCover(node *d11nodeT) d11rectT {
	rect := node.branch[0].rect
	for index := 1; index < node.count; index++ {
		rect = d11combineRect(&rect, &(node.branch[index].rect))
	}
	return rect
}

// Add a branch to a node.  Split the node if necessary.
// Returns 0 if node not split.  Old node updated.
// Returns 1 if node split, sets *new_node to address of new node.
// Old node updated, becomes one of two.
func d11addBranch(branch *d11branchT, node *d11nodeT, newNode **d11nodeT) bool {
	if node.count < d11maxNodes { // Split won't be necessary
		node.branch[node.count] = *branch
		node.count++
		return false
	} else {
		d11splitNode(node, branch, newNode)
		return true
	}
}

// Disconnect a dependent node.
// Caller must return (or stop using iteration index) after this as count has changed
func d11disconnectBranch(node *d11nodeT, index int) {
	// Remove element by swapping with the last element to prevent gaps in array
	node.branch[index] = node.branch[node.count-1]
	node.branch[node.count-1].data = nil
	node.branch[node.count-1].child = nil
	node.count--
}

// Pick a branch.  Pick the one that will need the smallest increase
// in area to accomodate the new rectangle.  This will result in the
// least total area for the covering rectangles in the current node.
// In case of a tie, pick the one which was smaller before, to get
// the best resolution when searching.
func d11pickBranch(rect *d11rectT, node *d11nodeT) int {
	var firstTime bool = true
	var increase float64
	var bestIncr float64 = -1
	var area float64
	var bestArea float64
	var best int
	var tempRect d11rectT

	for index := 0; index < node.count; index++ {
		curRect := &node.branch[index].rect
		area = d11calcRectVolume(curRect)
		tempRect = d11combineRect(rect, curRect)
		increase = d11calcRectVolume(&tempRect) - area
		if (increase < bestIncr) || firstTime {
			best = index
			bestArea = area
			bestIncr = increase
			firstTime = false
		} else if (increase == bestIncr) && (area < bestArea) {
			best = index
			bestArea = area
			bestIncr = increase
		}
	}
	return best
}

// Combine two rectangles into larger one containing both
func d11combineRect(rectA, rectB *d11rectT) d11rectT {
	var newRect d11rectT

	for index := 0; index < d11numDims; index++ {
		newRect.min[index] = d11fmin(rectA.min[index], rectB.min[index])
		newRect.max[index] = d11fmax(rectA.max[index], rectB.max[index])
	}

	return newRect
}

// Split a node.
// Divides the nodes branches and the extra one between two nodes.
// Old node is one of the new ones, and one really new one is created.
// Tries more than one method for choosing a partition, uses best result.
func d11splitNode(node *d11nodeT, branch *d11branchT, newNode **d11nodeT) {
	// Could just use local here, but member or external is faster since it is reused
	var localVars d11partitionVarsT
	parVars := &localVars

	// Load all the branches into a buffer, initialize old node
	d11getBranches(node, branch, parVars)

	// Find partition
	d11choosePartition(parVars, d11minNodes)

	// Create a new node to hold (about) half of the branches
	*newNode = &d11nodeT{}
	(*newNode).level = node.level

	// Put branches from buffer into 2 nodes according to the chosen partition
	node.count = 0
	d11loadNodes(node, *newNode, parVars)
}

// Calculate the n-dimensional volume of a rectangle
func d11rectVolume(rect *d11rectT) float64 {
	var volume float64 = 1
	for index := 0; index < d11numDims; index++ {
		volume *= rect.max[index] - rect.min[index]
	}
	return volume
}

// The exact volume of the bounding sphere for the given d11rectT
func d11rectSphericalVolume(rect *d11rectT) float64 {
	var sumOfSquares float64 = 0
	var radius float64

	for index := 0; index < d11numDims; index++ {
		halfExtent := (rect.max[index] - rect.min[index]) * 0.5
		sumOfSquares += halfExtent * halfExtent
	}

	radius = math.Sqrt(sumOfSquares)

	// Pow maybe slow, so test for common dims just use x*x, x*x*x.
	if d11numDims == 5 {
		return (radius * radius * radius * radius * radius * d11unitSphereVolume)
	} else if d11numDims == 4 {
		return (radius * radius * radius * radius * d11unitSphereVolume)
	} else if d11numDims == 3 {
		return (radius * radius * radius * d11unitSphereVolume)
	} else if d11numDims == 2 {
		return (radius * radius * d11unitSphereVolume)
	} else {
		return (math.Pow(radius, d11numDims) * d11unitSphereVolume)
	}
}

// Use one of the methods to calculate retangle volume
func d11calcRectVolume(rect *d11rectT) float64 {
	if d11useSphericalVolume {
		return d11rectSphericalVolume(rect) // Slower but helps certain merge cases
	} else { // RTREE_USE_SPHERICAL_VOLUME
		return d11rectVolume(rect) // Faster but can cause poor merges
	} // RTREE_USE_SPHERICAL_VOLUME
}

// Load branch buffer with branches from full node plus the extra branch.
func d11getBranches(node *d11nodeT, branch *d11branchT, parVars *d11partitionVarsT) {
	// Load the branch buffer
	for index := 0; index < d11maxNodes; index++ {
		parVars.branchBuf[index] = node.branch[index]
	}
	parVars.branchBuf[d11maxNodes] = *branch
	parVars.branchCount = d11maxNodes + 1

	// Calculate rect containing all in the set
	parVars.coverSplit = parVars.branchBuf[0].rect
	for index := 1; index < d11maxNodes+1; index++ {
		parVars.coverSplit = d11combineRect(&parVars.coverSplit, &parVars.branchBuf[index].rect)
	}
	parVars.coverSplitArea = d11calcRectVolume(&parVars.coverSplit)
}

// Method #0 for choosing a partition:
// As the seeds for the two groups, pick the two rects that would waste the
// most area if covered by a single rectangle, i.e. evidently the worst pair
// to have in the same group.
// Of the remaining, one at a time is chosen to be put in one of the two groups.
// The one chosen is the one with the greatest difference in area expansion
// depending on which group - the rect most strongly attracted to one group
// and repelled from the other.
// If one group gets too full (more would force other group to violate min
// fill requirement) then other group gets the rest.
// These last are the ones that can go in either group most easily.
func d11choosePartition(parVars *d11partitionVarsT, minFill int) {
	var biggestDiff float64
	var group, chosen, betterGroup int

	d11initParVars(parVars, parVars.branchCount, minFill)
	d11pickSeeds(parVars)

	for ((parVars.count[0] + parVars.count[1]) < parVars.total) &&
		(parVars.count[0] < (parVars.total - parVars.minFill)) &&
		(parVars.count[1] < (parVars.total - parVars.minFill)) {
		biggestDiff = -1
		for index := 0; index < parVars.total; index++ {
			if d11notTaken == parVars.partition[index] {
				curRect := &parVars.branchBuf[index].rect
				rect0 := d11combineRect(curRect, &parVars.cover[0])
				rect1 := d11combineRect(curRect, &parVars.cover[1])
				growth0 := d11calcRectVolume(&rect0) - parVars.area[0]
				growth1 := d11calcRectVolume(&rect1) - parVars.area[1]
				diff := growth1 - growth0
				if diff >= 0 {
					group = 0
				} else {
					group = 1
					diff = -diff
				}

				if diff > biggestDiff {
					biggestDiff = diff
					chosen = index
					betterGroup = group
				} else if (diff == biggestDiff) && (parVars.count[group] < parVars.count[betterGroup]) {
					chosen = index
					betterGroup = group
				}
			}
		}
		d11classify(chosen, betterGroup, parVars)
	}

	// If one group too full, put remaining rects in the other
	if (parVars.count[0] + parVars.count[1]) < parVars.total {
		if parVars.count[0] >= parVars.total-parVars.minFill {
			group = 1
		} else {
			group = 0
		}
		for index := 0; index < parVars.total; index++ {
			if d11notTaken == parVars.partition[index] {
				d11classify(index, group, parVars)
			}
		}
	}
}

// Copy branches from the buffer into two nodes according to the partition.
func d11loadNodes(nodeA, nodeB *d11nodeT, parVars *d11partitionVarsT) {
	for index := 0; index < parVars.total; index++ {
		targetNodeIndex := parVars.partition[index]
		targetNodes := []*d11nodeT{nodeA, nodeB}

		// It is assured that d11addBranch here will not cause a node split.
		d11addBranch(&parVars.branchBuf[index], targetNodes[targetNodeIndex], nil)
	}
}

// Initialize a d11partitionVarsT structure.
func d11initParVars(parVars *d11partitionVarsT, maxRects, minFill int) {
	parVars.count[0] = 0
	parVars.count[1] = 0
	parVars.area[0] = 0
	parVars.area[1] = 0
	parVars.total = maxRects
	parVars.minFill = minFill
	for index := 0; index < maxRects; index++ {
		parVars.partition[index] = d11notTaken
	}
}

func d11pickSeeds(parVars *d11partitionVarsT) {
	var seed0, seed1 int
	var worst, waste float64
	var area [d11maxNodes + 1]float64

	for index := 0; index < parVars.total; index++ {
		area[index] = d11calcRectVolume(&parVars.branchBuf[index].rect)
	}

	worst = -parVars.coverSplitArea - 1
	for indexA := 0; indexA < parVars.total-1; indexA++ {
		for indexB := indexA + 1; indexB < parVars.total; indexB++ {
			oneRect := d11combineRect(&parVars.branchBuf[indexA].rect, &parVars.branchBuf[indexB].rect)
			waste = d11calcRectVolume(&oneRect) - area[indexA] - area[indexB]
			if waste > worst {
				worst = waste
				seed0 = indexA
				seed1 = indexB
			}
		}
	}

	d11classify(seed0, 0, parVars)
	d11classify(seed1, 1, parVars)
}

// Put a branch in one of the groups.
func d11classify(index, group int, parVars *d11partitionVarsT) {
	parVars.partition[index] = group

	// Calculate combined rect
	if parVars.count[group] == 0 {
		parVars.cover[group] = parVars.branchBuf[index].rect
	} else {
		parVars.cover[group] = d11combineRect(&parVars.branchBuf[index].rect, &parVars.cover[group])
	}

	// Calculate volume of combined rect
	parVars.area[group] = d11calcRectVolume(&parVars.cover[group])

	parVars.count[group]++
}

// Delete a data rectangle from an index structure.
// Pass in a pointer to a d11rectT, the tid of the record, ptr to ptr to root node.
// Returns 1 if record not found, 0 if success.
// d11removeRect provides for eliminating the root.
func d11removeRect(rect *d11rectT, id interface{}, root **d11nodeT) bool {
	var reInsertList *d11listNodeT

	if !d11removeRectRec(rect, id, *root, &reInsertList) {
		// Found and deleted a data item
		// Reinsert any branches from eliminated nodes
		for reInsertList != nil {
			tempNode := reInsertList.node

			for index := 0; index < tempNode.count; index++ {
				// TODO go over this code. should I use (tempNode->m_level - 1)?
				d11insertRect(&tempNode.branch[index], root, tempNode.level)
			}
			reInsertList = reInsertList.next
		}

		// Check for redundant root (not leaf, 1 child) and eliminate TODO replace
		// if with while? In case there is a whole branch of redundant roots...
		if (*root).count == 1 && (*root).isInternalNode() {
			tempNode := (*root).branch[0].child
			*root = tempNode
		}
		return false
	} else {
		return true
	}
}

// Delete a rectangle from non-root part of an index structure.
// Called by d11removeRect.  Descends tree recursively,
// merges branches on the way back up.
// Returns 1 if record not found, 0 if success.
func d11removeRectRec(rect *d11rectT, id interface{}, node *d11nodeT, listNode **d11listNodeT) bool {
	if node.isInternalNode() { // not a leaf node
		for index := 0; index < node.count; index++ {
			if d11overlap(*rect, node.branch[index].rect) {
				if !d11removeRectRec(rect, id, node.branch[index].child, listNode) {
					if node.branch[index].child.count >= d11minNodes {
						// child removed, just resize parent rect
						node.branch[index].rect = d11nodeCover(node.branch[index].child)
					} else {
						// child removed, not enough entries in node, eliminate node
						d11reInsert(node.branch[index].child, listNode)
						d11disconnectBranch(node, index) // Must return after this call as count has changed
					}
					return false
				}
			}
		}
		return true
	} else { // A leaf node
		for index := 0; index < node.count; index++ {
			if node.branch[index].data == id {
				d11disconnectBranch(node, index) // Must return after this call as count has changed
				return false
			}
		}
		return true
	}
}

// Decide whether two rectangles d11overlap.
func d11overlap(rectA, rectB d11rectT) bool {
	for index := 0; index < d11numDims; index++ {
		if rectA.min[index] > rectB.max[index] ||
			rectB.min[index] > rectA.max[index] {
			return false
		}
	}
	return true
}

// Add a node to the reinsertion list.  All its branches will later
// be reinserted into the index structure.
func d11reInsert(node *d11nodeT, listNode **d11listNodeT) {
	newListNode := &d11listNodeT{}
	newListNode.node = node
	newListNode.next = *listNode
	*listNode = newListNode
}

// d11search in an index tree or subtree for all data retangles that d11overlap the argument rectangle.
func d11search(node *d11nodeT, rect d11rectT, foundCount int, resultCallback func(data interface{}) bool) (int, bool) {
	if node.isInternalNode() {
		// This is an internal node in the tree
		for index := 0; index < node.count; index++ {
			if d11overlap(rect, node.branch[index].rect) {
				var ok bool
				foundCount, ok = d11search(node.branch[index].child, rect, foundCount, resultCallback)
				if !ok {
					// The callback indicated to stop searching
					return foundCount, false
				}
			}
		}
	} else {
		// This is a leaf node
		for index := 0; index < node.count; index++ {
			if d11overlap(rect, node.branch[index].rect) {
				id := node.branch[index].data
				foundCount++
				if !resultCallback(id) {
					return foundCount, false // Don't continue searching
				}

			}
		}
	}
	return foundCount, true // Continue searching
}

func d12fmin(a, b float64) float64 {
	if a < b {
		return a
	}
	return b
}
func d12fmax(a, b float64) float64 {
	if a > b {
		return a
	}
	return b
}

const (
	d12numDims            = 12
	d12maxNodes           = 8
	d12minNodes           = d12maxNodes / 2
	d12useSphericalVolume = true // Better split classification, may be slower on some systems
)

var d12unitSphereVolume = []float64{
	0.000000, 2.000000, 3.141593, // Dimension  0,1,2
	4.188790, 4.934802, 5.263789, // Dimension  3,4,5
	5.167713, 4.724766, 4.058712, // Dimension  6,7,8
	3.298509, 2.550164, 1.884104, // Dimension  9,10,11
	1.335263, 0.910629, 0.599265, // Dimension  12,13,14
	0.381443, 0.235331, 0.140981, // Dimension  15,16,17
	0.082146, 0.046622, 0.025807, // Dimension  18,19,20
}[d12numDims]

type d12RTree struct {
	root *d12nodeT ///< Root of tree
}

/// Minimal bounding rectangle (n-dimensional)
type d12rectT struct {
	min [d12numDims]float64 ///< Min dimensions of bounding box
	max [d12numDims]float64 ///< Max dimensions of bounding box
}

/// May be data or may be another subtree
/// The parents level determines this.
/// If the parents level is 0, then this is data
type d12branchT struct {
	rect  d12rectT    ///< Bounds
	child *d12nodeT   ///< Child node
	data  interface{} ///< Data Id or Ptr
}

/// d12nodeT for each branch level
type d12nodeT struct {
	count  int                     ///< Count
	level  int                     ///< Leaf is zero, others positive
	branch [d12maxNodes]d12branchT ///< Branch
}

func (node *d12nodeT) isInternalNode() bool {
	return (node.level > 0) // Not a leaf, but a internal node
}
func (node *d12nodeT) isLeaf() bool {
	return (node.level == 0) // A leaf, contains data
}

/// A link list of nodes for reinsertion after a delete operation
type d12listNodeT struct {
	next *d12listNodeT ///< Next in list
	node *d12nodeT     ///< Node
}

const d12notTaken = -1 // indicates that position

/// Variables for finding a split partition
type d12partitionVarsT struct {
	partition [d12maxNodes + 1]int
	total     int
	minFill   int
	count     [2]int
	cover     [2]d12rectT
	area      [2]float64

	branchBuf      [d12maxNodes + 1]d12branchT
	branchCount    int
	coverSplit     d12rectT
	coverSplitArea float64
}

func d12New() *d12RTree {
	// We only support machine word size simple data type eg. integer index or object pointer.
	// Since we are storing as union with non data branch
	return &d12RTree{
		root: &d12nodeT{},
	}
}

/// Insert entry
/// \param a_min Min of bounding rect
/// \param a_max Max of bounding rect
/// \param a_dataId Positive Id of data.  Maybe zero, but negative numbers not allowed.
func (tr *d12RTree) Insert(min, max [d12numDims]float64, dataId interface{}) {
	var branch d12branchT
	branch.data = dataId
	for axis := 0; axis < d12numDims; axis++ {
		branch.rect.min[axis] = min[axis]
		branch.rect.max[axis] = max[axis]
	}
	d12insertRect(&branch, &tr.root, 0)
}

/// Remove entry
/// \param a_min Min of bounding rect
/// \param a_max Max of bounding rect
/// \param a_dataId Positive Id of data.  Maybe zero, but negative numbers not allowed.
func (tr *d12RTree) Remove(min, max [d12numDims]float64, dataId interface{}) {
	var rect d12rectT
	for axis := 0; axis < d12numDims; axis++ {
		rect.min[axis] = min[axis]
		rect.max[axis] = max[axis]
	}
	d12removeRect(&rect, dataId, &tr.root)
}

/// Find all within d12search rectangle
/// \param a_min Min of d12search bounding rect
/// \param a_max Max of d12search bounding rect
/// \param a_searchResult d12search result array.  Caller should set grow size. Function will reset, not append to array.
/// \param a_resultCallback Callback function to return result.  Callback should return 'true' to continue searching
/// \param a_context User context to pass as parameter to a_resultCallback
/// \return Returns the number of entries found
func (tr *d12RTree) Search(min, max [d12numDims]float64, resultCallback func(data interface{}) bool) int {
	var rect d12rectT
	for axis := 0; axis < d12numDims; axis++ {
		rect.min[axis] = min[axis]
		rect.max[axis] = max[axis]
	}
	foundCount, _ := d12search(tr.root, rect, 0, resultCallback)
	return foundCount
}

/// Count the data elements in this container.  This is slow as no internal counter is maintained.
func (tr *d12RTree) Count() int {
	var count int
	d12countRec(tr.root, &count)
	return count
}

/// Remove all entries from tree
func (tr *d12RTree) RemoveAll() {
	// Delete all existing nodes
	tr.root = &d12nodeT{}
}

func d12countRec(node *d12nodeT, count *int) {
	if node.isInternalNode() { // not a leaf node
		for index := 0; index < node.count; index++ {
			d12countRec(node.branch[index].child, count)
		}
	} else { // A leaf node
		*count += node.count
	}
}

// Inserts a new data rectangle into the index structure.
// Recursively descends tree, propagates splits back up.
// Returns 0 if node was not split.  Old node updated.
// If node was split, returns 1 and sets the pointer pointed to by
// new_node to point to the new node.  Old node updated to become one of two.
// The level argument specifies the number of steps up from the leaf
// level to insert; e.g. a data rectangle goes in at level = 0.
func d12insertRectRec(branch *d12branchT, node *d12nodeT, newNode **d12nodeT, level int) bool {
	// recurse until we reach the correct level for the new record. data records
	// will always be called with a_level == 0 (leaf)
	if node.level > level {
		// Still above level for insertion, go down tree recursively
		var otherNode *d12nodeT
		//var newBranch d12branchT

		// find the optimal branch for this record
		index := d12pickBranch(&branch.rect, node)

		// recursively insert this record into the picked branch
		childWasSplit := d12insertRectRec(branch, node.branch[index].child, &otherNode, level)

		if !childWasSplit {
			// Child was not split. Merge the bounding box of the new record with the
			// existing bounding box
			node.branch[index].rect = d12combineRect(&branch.rect, &(node.branch[index].rect))
			return false
		} else {
			// Child was split. The old branches are now re-partitioned to two nodes
			// so we have to re-calculate the bounding boxes of each node
			node.branch[index].rect = d12nodeCover(node.branch[index].child)
			var newBranch d12branchT
			newBranch.child = otherNode
			newBranch.rect = d12nodeCover(otherNode)

			// The old node is already a child of a_node. Now add the newly-created
			// node to a_node as well. a_node might be split because of that.
			return d12addBranch(&newBranch, node, newNode)
		}
	} else if node.level == level {
		// We have reached level for insertion. Add rect, split if necessary
		return d12addBranch(branch, node, newNode)
	} else {
		// Should never occur
		return false
	}
}

// Insert a data rectangle into an index structure.
// d12insertRect provides for splitting the root;
// returns 1 if root was split, 0 if it was not.
// The level argument specifies the number of steps up from the leaf
// level to insert; e.g. a data rectangle goes in at level = 0.
// InsertRect2 does the recursion.
//
func d12insertRect(branch *d12branchT, root **d12nodeT, level int) bool {
	var newNode *d12nodeT

	if d12insertRectRec(branch, *root, &newNode, level) { // Root split

		// Grow tree taller and new root
		newRoot := &d12nodeT{}
		newRoot.level = (*root).level + 1

		var newBranch d12branchT

		// add old root node as a child of the new root
		newBranch.rect = d12nodeCover(*root)
		newBranch.child = *root
		d12addBranch(&newBranch, newRoot, nil)

		// add the split node as a child of the new root
		newBranch.rect = d12nodeCover(newNode)
		newBranch.child = newNode
		d12addBranch(&newBranch, newRoot, nil)

		// set the new root as the root node
		*root = newRoot

		return true
	}
	return false
}

// Find the smallest rectangle that includes all rectangles in branches of a node.
func d12nodeCover(node *d12nodeT) d12rectT {
	rect := node.branch[0].rect
	for index := 1; index < node.count; index++ {
		rect = d12combineRect(&rect, &(node.branch[index].rect))
	}
	return rect
}

// Add a branch to a node.  Split the node if necessary.
// Returns 0 if node not split.  Old node updated.
// Returns 1 if node split, sets *new_node to address of new node.
// Old node updated, becomes one of two.
func d12addBranch(branch *d12branchT, node *d12nodeT, newNode **d12nodeT) bool {
	if node.count < d12maxNodes { // Split won't be necessary
		node.branch[node.count] = *branch
		node.count++
		return false
	} else {
		d12splitNode(node, branch, newNode)
		return true
	}
}

// Disconnect a dependent node.
// Caller must return (or stop using iteration index) after this as count has changed
func d12disconnectBranch(node *d12nodeT, index int) {
	// Remove element by swapping with the last element to prevent gaps in array
	node.branch[index] = node.branch[node.count-1]
	node.branch[node.count-1].data = nil
	node.branch[node.count-1].child = nil
	node.count--
}

// Pick a branch.  Pick the one that will need the smallest increase
// in area to accomodate the new rectangle.  This will result in the
// least total area for the covering rectangles in the current node.
// In case of a tie, pick the one which was smaller before, to get
// the best resolution when searching.
func d12pickBranch(rect *d12rectT, node *d12nodeT) int {
	var firstTime bool = true
	var increase float64
	var bestIncr float64 = -1
	var area float64
	var bestArea float64
	var best int
	var tempRect d12rectT

	for index := 0; index < node.count; index++ {
		curRect := &node.branch[index].rect
		area = d12calcRectVolume(curRect)
		tempRect = d12combineRect(rect, curRect)
		increase = d12calcRectVolume(&tempRect) - area
		if (increase < bestIncr) || firstTime {
			best = index
			bestArea = area
			bestIncr = increase
			firstTime = false
		} else if (increase == bestIncr) && (area < bestArea) {
			best = index
			bestArea = area
			bestIncr = increase
		}
	}
	return best
}

// Combine two rectangles into larger one containing both
func d12combineRect(rectA, rectB *d12rectT) d12rectT {
	var newRect d12rectT

	for index := 0; index < d12numDims; index++ {
		newRect.min[index] = d12fmin(rectA.min[index], rectB.min[index])
		newRect.max[index] = d12fmax(rectA.max[index], rectB.max[index])
	}

	return newRect
}

// Split a node.
// Divides the nodes branches and the extra one between two nodes.
// Old node is one of the new ones, and one really new one is created.
// Tries more than one method for choosing a partition, uses best result.
func d12splitNode(node *d12nodeT, branch *d12branchT, newNode **d12nodeT) {
	// Could just use local here, but member or external is faster since it is reused
	var localVars d12partitionVarsT
	parVars := &localVars

	// Load all the branches into a buffer, initialize old node
	d12getBranches(node, branch, parVars)

	// Find partition
	d12choosePartition(parVars, d12minNodes)

	// Create a new node to hold (about) half of the branches
	*newNode = &d12nodeT{}
	(*newNode).level = node.level

	// Put branches from buffer into 2 nodes according to the chosen partition
	node.count = 0
	d12loadNodes(node, *newNode, parVars)
}

// Calculate the n-dimensional volume of a rectangle
func d12rectVolume(rect *d12rectT) float64 {
	var volume float64 = 1
	for index := 0; index < d12numDims; index++ {
		volume *= rect.max[index] - rect.min[index]
	}
	return volume
}

// The exact volume of the bounding sphere for the given d12rectT
func d12rectSphericalVolume(rect *d12rectT) float64 {
	var sumOfSquares float64 = 0
	var radius float64

	for index := 0; index < d12numDims; index++ {
		halfExtent := (rect.max[index] - rect.min[index]) * 0.5
		sumOfSquares += halfExtent * halfExtent
	}

	radius = math.Sqrt(sumOfSquares)

	// Pow maybe slow, so test for common dims just use x*x, x*x*x.
	if d12numDims == 5 {
		return (radius * radius * radius * radius * radius * d12unitSphereVolume)
	} else if d12numDims == 4 {
		return (radius * radius * radius * radius * d12unitSphereVolume)
	} else if d12numDims == 3 {
		return (radius * radius * radius * d12unitSphereVolume)
	} else if d12numDims == 2 {
		return (radius * radius * d12unitSphereVolume)
	} else {
		return (math.Pow(radius, d12numDims) * d12unitSphereVolume)
	}
}

// Use one of the methods to calculate retangle volume
func d12calcRectVolume(rect *d12rectT) float64 {
	if d12useSphericalVolume {
		return d12rectSphericalVolume(rect) // Slower but helps certain merge cases
	} else { // RTREE_USE_SPHERICAL_VOLUME
		return d12rectVolume(rect) // Faster but can cause poor merges
	} // RTREE_USE_SPHERICAL_VOLUME
}

// Load branch buffer with branches from full node plus the extra branch.
func d12getBranches(node *d12nodeT, branch *d12branchT, parVars *d12partitionVarsT) {
	// Load the branch buffer
	for index := 0; index < d12maxNodes; index++ {
		parVars.branchBuf[index] = node.branch[index]
	}
	parVars.branchBuf[d12maxNodes] = *branch
	parVars.branchCount = d12maxNodes + 1

	// Calculate rect containing all in the set
	parVars.coverSplit = parVars.branchBuf[0].rect
	for index := 1; index < d12maxNodes+1; index++ {
		parVars.coverSplit = d12combineRect(&parVars.coverSplit, &parVars.branchBuf[index].rect)
	}
	parVars.coverSplitArea = d12calcRectVolume(&parVars.coverSplit)
}

// Method #0 for choosing a partition:
// As the seeds for the two groups, pick the two rects that would waste the
// most area if covered by a single rectangle, i.e. evidently the worst pair
// to have in the same group.
// Of the remaining, one at a time is chosen to be put in one of the two groups.
// The one chosen is the one with the greatest difference in area expansion
// depending on which group - the rect most strongly attracted to one group
// and repelled from the other.
// If one group gets too full (more would force other group to violate min
// fill requirement) then other group gets the rest.
// These last are the ones that can go in either group most easily.
func d12choosePartition(parVars *d12partitionVarsT, minFill int) {
	var biggestDiff float64
	var group, chosen, betterGroup int

	d12initParVars(parVars, parVars.branchCount, minFill)
	d12pickSeeds(parVars)

	for ((parVars.count[0] + parVars.count[1]) < parVars.total) &&
		(parVars.count[0] < (parVars.total - parVars.minFill)) &&
		(parVars.count[1] < (parVars.total - parVars.minFill)) {
		biggestDiff = -1
		for index := 0; index < parVars.total; index++ {
			if d12notTaken == parVars.partition[index] {
				curRect := &parVars.branchBuf[index].rect
				rect0 := d12combineRect(curRect, &parVars.cover[0])
				rect1 := d12combineRect(curRect, &parVars.cover[1])
				growth0 := d12calcRectVolume(&rect0) - parVars.area[0]
				growth1 := d12calcRectVolume(&rect1) - parVars.area[1]
				diff := growth1 - growth0
				if diff >= 0 {
					group = 0
				} else {
					group = 1
					diff = -diff
				}

				if diff > biggestDiff {
					biggestDiff = diff
					chosen = index
					betterGroup = group
				} else if (diff == biggestDiff) && (parVars.count[group] < parVars.count[betterGroup]) {
					chosen = index
					betterGroup = group
				}
			}
		}
		d12classify(chosen, betterGroup, parVars)
	}

	// If one group too full, put remaining rects in the other
	if (parVars.count[0] + parVars.count[1]) < parVars.total {
		if parVars.count[0] >= parVars.total-parVars.minFill {
			group = 1
		} else {
			group = 0
		}
		for index := 0; index < parVars.total; index++ {
			if d12notTaken == parVars.partition[index] {
				d12classify(index, group, parVars)
			}
		}
	}
}

// Copy branches from the buffer into two nodes according to the partition.
func d12loadNodes(nodeA, nodeB *d12nodeT, parVars *d12partitionVarsT) {
	for index := 0; index < parVars.total; index++ {
		targetNodeIndex := parVars.partition[index]
		targetNodes := []*d12nodeT{nodeA, nodeB}

		// It is assured that d12addBranch here will not cause a node split.
		d12addBranch(&parVars.branchBuf[index], targetNodes[targetNodeIndex], nil)
	}
}

// Initialize a d12partitionVarsT structure.
func d12initParVars(parVars *d12partitionVarsT, maxRects, minFill int) {
	parVars.count[0] = 0
	parVars.count[1] = 0
	parVars.area[0] = 0
	parVars.area[1] = 0
	parVars.total = maxRects
	parVars.minFill = minFill
	for index := 0; index < maxRects; index++ {
		parVars.partition[index] = d12notTaken
	}
}

func d12pickSeeds(parVars *d12partitionVarsT) {
	var seed0, seed1 int
	var worst, waste float64
	var area [d12maxNodes + 1]float64

	for index := 0; index < parVars.total; index++ {
		area[index] = d12calcRectVolume(&parVars.branchBuf[index].rect)
	}

	worst = -parVars.coverSplitArea - 1
	for indexA := 0; indexA < parVars.total-1; indexA++ {
		for indexB := indexA + 1; indexB < parVars.total; indexB++ {
			oneRect := d12combineRect(&parVars.branchBuf[indexA].rect, &parVars.branchBuf[indexB].rect)
			waste = d12calcRectVolume(&oneRect) - area[indexA] - area[indexB]
			if waste > worst {
				worst = waste
				seed0 = indexA
				seed1 = indexB
			}
		}
	}

	d12classify(seed0, 0, parVars)
	d12classify(seed1, 1, parVars)
}

// Put a branch in one of the groups.
func d12classify(index, group int, parVars *d12partitionVarsT) {
	parVars.partition[index] = group

	// Calculate combined rect
	if parVars.count[group] == 0 {
		parVars.cover[group] = parVars.branchBuf[index].rect
	} else {
		parVars.cover[group] = d12combineRect(&parVars.branchBuf[index].rect, &parVars.cover[group])
	}

	// Calculate volume of combined rect
	parVars.area[group] = d12calcRectVolume(&parVars.cover[group])

	parVars.count[group]++
}

// Delete a data rectangle from an index structure.
// Pass in a pointer to a d12rectT, the tid of the record, ptr to ptr to root node.
// Returns 1 if record not found, 0 if success.
// d12removeRect provides for eliminating the root.
func d12removeRect(rect *d12rectT, id interface{}, root **d12nodeT) bool {
	var reInsertList *d12listNodeT

	if !d12removeRectRec(rect, id, *root, &reInsertList) {
		// Found and deleted a data item
		// Reinsert any branches from eliminated nodes
		for reInsertList != nil {
			tempNode := reInsertList.node

			for index := 0; index < tempNode.count; index++ {
				// TODO go over this code. should I use (tempNode->m_level - 1)?
				d12insertRect(&tempNode.branch[index], root, tempNode.level)
			}
			reInsertList = reInsertList.next
		}

		// Check for redundant root (not leaf, 1 child) and eliminate TODO replace
		// if with while? In case there is a whole branch of redundant roots...
		if (*root).count == 1 && (*root).isInternalNode() {
			tempNode := (*root).branch[0].child
			*root = tempNode
		}
		return false
	} else {
		return true
	}
}

// Delete a rectangle from non-root part of an index structure.
// Called by d12removeRect.  Descends tree recursively,
// merges branches on the way back up.
// Returns 1 if record not found, 0 if success.
func d12removeRectRec(rect *d12rectT, id interface{}, node *d12nodeT, listNode **d12listNodeT) bool {
	if node.isInternalNode() { // not a leaf node
		for index := 0; index < node.count; index++ {
			if d12overlap(*rect, node.branch[index].rect) {
				if !d12removeRectRec(rect, id, node.branch[index].child, listNode) {
					if node.branch[index].child.count >= d12minNodes {
						// child removed, just resize parent rect
						node.branch[index].rect = d12nodeCover(node.branch[index].child)
					} else {
						// child removed, not enough entries in node, eliminate node
						d12reInsert(node.branch[index].child, listNode)
						d12disconnectBranch(node, index) // Must return after this call as count has changed
					}
					return false
				}
			}
		}
		return true
	} else { // A leaf node
		for index := 0; index < node.count; index++ {
			if node.branch[index].data == id {
				d12disconnectBranch(node, index) // Must return after this call as count has changed
				return false
			}
		}
		return true
	}
}

// Decide whether two rectangles d12overlap.
func d12overlap(rectA, rectB d12rectT) bool {
	for index := 0; index < d12numDims; index++ {
		if rectA.min[index] > rectB.max[index] ||
			rectB.min[index] > rectA.max[index] {
			return false
		}
	}
	return true
}

// Add a node to the reinsertion list.  All its branches will later
// be reinserted into the index structure.
func d12reInsert(node *d12nodeT, listNode **d12listNodeT) {
	newListNode := &d12listNodeT{}
	newListNode.node = node
	newListNode.next = *listNode
	*listNode = newListNode
}

// d12search in an index tree or subtree for all data retangles that d12overlap the argument rectangle.
func d12search(node *d12nodeT, rect d12rectT, foundCount int, resultCallback func(data interface{}) bool) (int, bool) {
	if node.isInternalNode() {
		// This is an internal node in the tree
		for index := 0; index < node.count; index++ {
			if d12overlap(rect, node.branch[index].rect) {
				var ok bool
				foundCount, ok = d12search(node.branch[index].child, rect, foundCount, resultCallback)
				if !ok {
					// The callback indicated to stop searching
					return foundCount, false
				}
			}
		}
	} else {
		// This is a leaf node
		for index := 0; index < node.count; index++ {
			if d12overlap(rect, node.branch[index].rect) {
				id := node.branch[index].data
				foundCount++
				if !resultCallback(id) {
					return foundCount, false // Don't continue searching
				}

			}
		}
	}
	return foundCount, true // Continue searching
}

func d13fmin(a, b float64) float64 {
	if a < b {
		return a
	}
	return b
}
func d13fmax(a, b float64) float64 {
	if a > b {
		return a
	}
	return b
}

const (
	d13numDims            = 13
	d13maxNodes           = 8
	d13minNodes           = d13maxNodes / 2
	d13useSphericalVolume = true // Better split classification, may be slower on some systems
)

var d13unitSphereVolume = []float64{
	0.000000, 2.000000, 3.141593, // Dimension  0,1,2
	4.188790, 4.934802, 5.263789, // Dimension  3,4,5
	5.167713, 4.724766, 4.058712, // Dimension  6,7,8
	3.298509, 2.550164, 1.884104, // Dimension  9,10,11
	1.335263, 0.910629, 0.599265, // Dimension  12,13,14
	0.381443, 0.235331, 0.140981, // Dimension  15,16,17
	0.082146, 0.046622, 0.025807, // Dimension  18,19,20
}[d13numDims]

type d13RTree struct {
	root *d13nodeT ///< Root of tree
}

/// Minimal bounding rectangle (n-dimensional)
type d13rectT struct {
	min [d13numDims]float64 ///< Min dimensions of bounding box
	max [d13numDims]float64 ///< Max dimensions of bounding box
}

/// May be data or may be another subtree
/// The parents level determines this.
/// If the parents level is 0, then this is data
type d13branchT struct {
	rect  d13rectT    ///< Bounds
	child *d13nodeT   ///< Child node
	data  interface{} ///< Data Id or Ptr
}

/// d13nodeT for each branch level
type d13nodeT struct {
	count  int                     ///< Count
	level  int                     ///< Leaf is zero, others positive
	branch [d13maxNodes]d13branchT ///< Branch
}

func (node *d13nodeT) isInternalNode() bool {
	return (node.level > 0) // Not a leaf, but a internal node
}
func (node *d13nodeT) isLeaf() bool {
	return (node.level == 0) // A leaf, contains data
}

/// A link list of nodes for reinsertion after a delete operation
type d13listNodeT struct {
	next *d13listNodeT ///< Next in list
	node *d13nodeT     ///< Node
}

const d13notTaken = -1 // indicates that position

/// Variables for finding a split partition
type d13partitionVarsT struct {
	partition [d13maxNodes + 1]int
	total     int
	minFill   int
	count     [2]int
	cover     [2]d13rectT
	area      [2]float64

	branchBuf      [d13maxNodes + 1]d13branchT
	branchCount    int
	coverSplit     d13rectT
	coverSplitArea float64
}

func d13New() *d13RTree {
	// We only support machine word size simple data type eg. integer index or object pointer.
	// Since we are storing as union with non data branch
	return &d13RTree{
		root: &d13nodeT{},
	}
}

/// Insert entry
/// \param a_min Min of bounding rect
/// \param a_max Max of bounding rect
/// \param a_dataId Positive Id of data.  Maybe zero, but negative numbers not allowed.
func (tr *d13RTree) Insert(min, max [d13numDims]float64, dataId interface{}) {
	var branch d13branchT
	branch.data = dataId
	for axis := 0; axis < d13numDims; axis++ {
		branch.rect.min[axis] = min[axis]
		branch.rect.max[axis] = max[axis]
	}
	d13insertRect(&branch, &tr.root, 0)
}

/// Remove entry
/// \param a_min Min of bounding rect
/// \param a_max Max of bounding rect
/// \param a_dataId Positive Id of data.  Maybe zero, but negative numbers not allowed.
func (tr *d13RTree) Remove(min, max [d13numDims]float64, dataId interface{}) {
	var rect d13rectT
	for axis := 0; axis < d13numDims; axis++ {
		rect.min[axis] = min[axis]
		rect.max[axis] = max[axis]
	}
	d13removeRect(&rect, dataId, &tr.root)
}

/// Find all within d13search rectangle
/// \param a_min Min of d13search bounding rect
/// \param a_max Max of d13search bounding rect
/// \param a_searchResult d13search result array.  Caller should set grow size. Function will reset, not append to array.
/// \param a_resultCallback Callback function to return result.  Callback should return 'true' to continue searching
/// \param a_context User context to pass as parameter to a_resultCallback
/// \return Returns the number of entries found
func (tr *d13RTree) Search(min, max [d13numDims]float64, resultCallback func(data interface{}) bool) int {
	var rect d13rectT
	for axis := 0; axis < d13numDims; axis++ {
		rect.min[axis] = min[axis]
		rect.max[axis] = max[axis]
	}
	foundCount, _ := d13search(tr.root, rect, 0, resultCallback)
	return foundCount
}

/// Count the data elements in this container.  This is slow as no internal counter is maintained.
func (tr *d13RTree) Count() int {
	var count int
	d13countRec(tr.root, &count)
	return count
}

/// Remove all entries from tree
func (tr *d13RTree) RemoveAll() {
	// Delete all existing nodes
	tr.root = &d13nodeT{}
}

func d13countRec(node *d13nodeT, count *int) {
	if node.isInternalNode() { // not a leaf node
		for index := 0; index < node.count; index++ {
			d13countRec(node.branch[index].child, count)
		}
	} else { // A leaf node
		*count += node.count
	}
}

// Inserts a new data rectangle into the index structure.
// Recursively descends tree, propagates splits back up.
// Returns 0 if node was not split.  Old node updated.
// If node was split, returns 1 and sets the pointer pointed to by
// new_node to point to the new node.  Old node updated to become one of two.
// The level argument specifies the number of steps up from the leaf
// level to insert; e.g. a data rectangle goes in at level = 0.
func d13insertRectRec(branch *d13branchT, node *d13nodeT, newNode **d13nodeT, level int) bool {
	// recurse until we reach the correct level for the new record. data records
	// will always be called with a_level == 0 (leaf)
	if node.level > level {
		// Still above level for insertion, go down tree recursively
		var otherNode *d13nodeT
		//var newBranch d13branchT

		// find the optimal branch for this record
		index := d13pickBranch(&branch.rect, node)

		// recursively insert this record into the picked branch
		childWasSplit := d13insertRectRec(branch, node.branch[index].child, &otherNode, level)

		if !childWasSplit {
			// Child was not split. Merge the bounding box of the new record with the
			// existing bounding box
			node.branch[index].rect = d13combineRect(&branch.rect, &(node.branch[index].rect))
			return false
		} else {
			// Child was split. The old branches are now re-partitioned to two nodes
			// so we have to re-calculate the bounding boxes of each node
			node.branch[index].rect = d13nodeCover(node.branch[index].child)
			var newBranch d13branchT
			newBranch.child = otherNode
			newBranch.rect = d13nodeCover(otherNode)

			// The old node is already a child of a_node. Now add the newly-created
			// node to a_node as well. a_node might be split because of that.
			return d13addBranch(&newBranch, node, newNode)
		}
	} else if node.level == level {
		// We have reached level for insertion. Add rect, split if necessary
		return d13addBranch(branch, node, newNode)
	} else {
		// Should never occur
		return false
	}
}

// Insert a data rectangle into an index structure.
// d13insertRect provides for splitting the root;
// returns 1 if root was split, 0 if it was not.
// The level argument specifies the number of steps up from the leaf
// level to insert; e.g. a data rectangle goes in at level = 0.
// InsertRect2 does the recursion.
//
func d13insertRect(branch *d13branchT, root **d13nodeT, level int) bool {
	var newNode *d13nodeT

	if d13insertRectRec(branch, *root, &newNode, level) { // Root split

		// Grow tree taller and new root
		newRoot := &d13nodeT{}
		newRoot.level = (*root).level + 1

		var newBranch d13branchT

		// add old root node as a child of the new root
		newBranch.rect = d13nodeCover(*root)
		newBranch.child = *root
		d13addBranch(&newBranch, newRoot, nil)

		// add the split node as a child of the new root
		newBranch.rect = d13nodeCover(newNode)
		newBranch.child = newNode
		d13addBranch(&newBranch, newRoot, nil)

		// set the new root as the root node
		*root = newRoot

		return true
	}
	return false
}

// Find the smallest rectangle that includes all rectangles in branches of a node.
func d13nodeCover(node *d13nodeT) d13rectT {
	rect := node.branch[0].rect
	for index := 1; index < node.count; index++ {
		rect = d13combineRect(&rect, &(node.branch[index].rect))
	}
	return rect
}

// Add a branch to a node.  Split the node if necessary.
// Returns 0 if node not split.  Old node updated.
// Returns 1 if node split, sets *new_node to address of new node.
// Old node updated, becomes one of two.
func d13addBranch(branch *d13branchT, node *d13nodeT, newNode **d13nodeT) bool {
	if node.count < d13maxNodes { // Split won't be necessary
		node.branch[node.count] = *branch
		node.count++
		return false
	} else {
		d13splitNode(node, branch, newNode)
		return true
	}
}

// Disconnect a dependent node.
// Caller must return (or stop using iteration index) after this as count has changed
func d13disconnectBranch(node *d13nodeT, index int) {
	// Remove element by swapping with the last element to prevent gaps in array
	node.branch[index] = node.branch[node.count-1]
	node.branch[node.count-1].data = nil
	node.branch[node.count-1].child = nil
	node.count--
}

// Pick a branch.  Pick the one that will need the smallest increase
// in area to accomodate the new rectangle.  This will result in the
// least total area for the covering rectangles in the current node.
// In case of a tie, pick the one which was smaller before, to get
// the best resolution when searching.
func d13pickBranch(rect *d13rectT, node *d13nodeT) int {
	var firstTime bool = true
	var increase float64
	var bestIncr float64 = -1
	var area float64
	var bestArea float64
	var best int
	var tempRect d13rectT

	for index := 0; index < node.count; index++ {
		curRect := &node.branch[index].rect
		area = d13calcRectVolume(curRect)
		tempRect = d13combineRect(rect, curRect)
		increase = d13calcRectVolume(&tempRect) - area
		if (increase < bestIncr) || firstTime {
			best = index
			bestArea = area
			bestIncr = increase
			firstTime = false
		} else if (increase == bestIncr) && (area < bestArea) {
			best = index
			bestArea = area
			bestIncr = increase
		}
	}
	return best
}

// Combine two rectangles into larger one containing both
func d13combineRect(rectA, rectB *d13rectT) d13rectT {
	var newRect d13rectT

	for index := 0; index < d13numDims; index++ {
		newRect.min[index] = d13fmin(rectA.min[index], rectB.min[index])
		newRect.max[index] = d13fmax(rectA.max[index], rectB.max[index])
	}

	return newRect
}

// Split a node.
// Divides the nodes branches and the extra one between two nodes.
// Old node is one of the new ones, and one really new one is created.
// Tries more than one method for choosing a partition, uses best result.
func d13splitNode(node *d13nodeT, branch *d13branchT, newNode **d13nodeT) {
	// Could just use local here, but member or external is faster since it is reused
	var localVars d13partitionVarsT
	parVars := &localVars

	// Load all the branches into a buffer, initialize old node
	d13getBranches(node, branch, parVars)

	// Find partition
	d13choosePartition(parVars, d13minNodes)

	// Create a new node to hold (about) half of the branches
	*newNode = &d13nodeT{}
	(*newNode).level = node.level

	// Put branches from buffer into 2 nodes according to the chosen partition
	node.count = 0
	d13loadNodes(node, *newNode, parVars)
}

// Calculate the n-dimensional volume of a rectangle
func d13rectVolume(rect *d13rectT) float64 {
	var volume float64 = 1
	for index := 0; index < d13numDims; index++ {
		volume *= rect.max[index] - rect.min[index]
	}
	return volume
}

// The exact volume of the bounding sphere for the given d13rectT
func d13rectSphericalVolume(rect *d13rectT) float64 {
	var sumOfSquares float64 = 0
	var radius float64

	for index := 0; index < d13numDims; index++ {
		halfExtent := (rect.max[index] - rect.min[index]) * 0.5
		sumOfSquares += halfExtent * halfExtent
	}

	radius = math.Sqrt(sumOfSquares)

	// Pow maybe slow, so test for common dims just use x*x, x*x*x.
	if d13numDims == 5 {
		return (radius * radius * radius * radius * radius * d13unitSphereVolume)
	} else if d13numDims == 4 {
		return (radius * radius * radius * radius * d13unitSphereVolume)
	} else if d13numDims == 3 {
		return (radius * radius * radius * d13unitSphereVolume)
	} else if d13numDims == 2 {
		return (radius * radius * d13unitSphereVolume)
	} else {
		return (math.Pow(radius, d13numDims) * d13unitSphereVolume)
	}
}

// Use one of the methods to calculate retangle volume
func d13calcRectVolume(rect *d13rectT) float64 {
	if d13useSphericalVolume {
		return d13rectSphericalVolume(rect) // Slower but helps certain merge cases
	} else { // RTREE_USE_SPHERICAL_VOLUME
		return d13rectVolume(rect) // Faster but can cause poor merges
	} // RTREE_USE_SPHERICAL_VOLUME
}

// Load branch buffer with branches from full node plus the extra branch.
func d13getBranches(node *d13nodeT, branch *d13branchT, parVars *d13partitionVarsT) {
	// Load the branch buffer
	for index := 0; index < d13maxNodes; index++ {
		parVars.branchBuf[index] = node.branch[index]
	}
	parVars.branchBuf[d13maxNodes] = *branch
	parVars.branchCount = d13maxNodes + 1

	// Calculate rect containing all in the set
	parVars.coverSplit = parVars.branchBuf[0].rect
	for index := 1; index < d13maxNodes+1; index++ {
		parVars.coverSplit = d13combineRect(&parVars.coverSplit, &parVars.branchBuf[index].rect)
	}
	parVars.coverSplitArea = d13calcRectVolume(&parVars.coverSplit)
}

// Method #0 for choosing a partition:
// As the seeds for the two groups, pick the two rects that would waste the
// most area if covered by a single rectangle, i.e. evidently the worst pair
// to have in the same group.
// Of the remaining, one at a time is chosen to be put in one of the two groups.
// The one chosen is the one with the greatest difference in area expansion
// depending on which group - the rect most strongly attracted to one group
// and repelled from the other.
// If one group gets too full (more would force other group to violate min
// fill requirement) then other group gets the rest.
// These last are the ones that can go in either group most easily.
func d13choosePartition(parVars *d13partitionVarsT, minFill int) {
	var biggestDiff float64
	var group, chosen, betterGroup int

	d13initParVars(parVars, parVars.branchCount, minFill)
	d13pickSeeds(parVars)

	for ((parVars.count[0] + parVars.count[1]) < parVars.total) &&
		(parVars.count[0] < (parVars.total - parVars.minFill)) &&
		(parVars.count[1] < (parVars.total - parVars.minFill)) {
		biggestDiff = -1
		for index := 0; index < parVars.total; index++ {
			if d13notTaken == parVars.partition[index] {
				curRect := &parVars.branchBuf[index].rect
				rect0 := d13combineRect(curRect, &parVars.cover[0])
				rect1 := d13combineRect(curRect, &parVars.cover[1])
				growth0 := d13calcRectVolume(&rect0) - parVars.area[0]
				growth1 := d13calcRectVolume(&rect1) - parVars.area[1]
				diff := growth1 - growth0
				if diff >= 0 {
					group = 0
				} else {
					group = 1
					diff = -diff
				}

				if diff > biggestDiff {
					biggestDiff = diff
					chosen = index
					betterGroup = group
				} else if (diff == biggestDiff) && (parVars.count[group] < parVars.count[betterGroup]) {
					chosen = index
					betterGroup = group
				}
			}
		}
		d13classify(chosen, betterGroup, parVars)
	}

	// If one group too full, put remaining rects in the other
	if (parVars.count[0] + parVars.count[1]) < parVars.total {
		if parVars.count[0] >= parVars.total-parVars.minFill {
			group = 1
		} else {
			group = 0
		}
		for index := 0; index < parVars.total; index++ {
			if d13notTaken == parVars.partition[index] {
				d13classify(index, group, parVars)
			}
		}
	}
}

// Copy branches from the buffer into two nodes according to the partition.
func d13loadNodes(nodeA, nodeB *d13nodeT, parVars *d13partitionVarsT) {
	for index := 0; index < parVars.total; index++ {
		targetNodeIndex := parVars.partition[index]
		targetNodes := []*d13nodeT{nodeA, nodeB}

		// It is assured that d13addBranch here will not cause a node split.
		d13addBranch(&parVars.branchBuf[index], targetNodes[targetNodeIndex], nil)
	}
}

// Initialize a d13partitionVarsT structure.
func d13initParVars(parVars *d13partitionVarsT, maxRects, minFill int) {
	parVars.count[0] = 0
	parVars.count[1] = 0
	parVars.area[0] = 0
	parVars.area[1] = 0
	parVars.total = maxRects
	parVars.minFill = minFill
	for index := 0; index < maxRects; index++ {
		parVars.partition[index] = d13notTaken
	}
}

func d13pickSeeds(parVars *d13partitionVarsT) {
	var seed0, seed1 int
	var worst, waste float64
	var area [d13maxNodes + 1]float64

	for index := 0; index < parVars.total; index++ {
		area[index] = d13calcRectVolume(&parVars.branchBuf[index].rect)
	}

	worst = -parVars.coverSplitArea - 1
	for indexA := 0; indexA < parVars.total-1; indexA++ {
		for indexB := indexA + 1; indexB < parVars.total; indexB++ {
			oneRect := d13combineRect(&parVars.branchBuf[indexA].rect, &parVars.branchBuf[indexB].rect)
			waste = d13calcRectVolume(&oneRect) - area[indexA] - area[indexB]
			if waste > worst {
				worst = waste
				seed0 = indexA
				seed1 = indexB
			}
		}
	}

	d13classify(seed0, 0, parVars)
	d13classify(seed1, 1, parVars)
}

// Put a branch in one of the groups.
func d13classify(index, group int, parVars *d13partitionVarsT) {
	parVars.partition[index] = group

	// Calculate combined rect
	if parVars.count[group] == 0 {
		parVars.cover[group] = parVars.branchBuf[index].rect
	} else {
		parVars.cover[group] = d13combineRect(&parVars.branchBuf[index].rect, &parVars.cover[group])
	}

	// Calculate volume of combined rect
	parVars.area[group] = d13calcRectVolume(&parVars.cover[group])

	parVars.count[group]++
}

// Delete a data rectangle from an index structure.
// Pass in a pointer to a d13rectT, the tid of the record, ptr to ptr to root node.
// Returns 1 if record not found, 0 if success.
// d13removeRect provides for eliminating the root.
func d13removeRect(rect *d13rectT, id interface{}, root **d13nodeT) bool {
	var reInsertList *d13listNodeT

	if !d13removeRectRec(rect, id, *root, &reInsertList) {
		// Found and deleted a data item
		// Reinsert any branches from eliminated nodes
		for reInsertList != nil {
			tempNode := reInsertList.node

			for index := 0; index < tempNode.count; index++ {
				// TODO go over this code. should I use (tempNode->m_level - 1)?
				d13insertRect(&tempNode.branch[index], root, tempNode.level)
			}
			reInsertList = reInsertList.next
		}

		// Check for redundant root (not leaf, 1 child) and eliminate TODO replace
		// if with while? In case there is a whole branch of redundant roots...
		if (*root).count == 1 && (*root).isInternalNode() {
			tempNode := (*root).branch[0].child
			*root = tempNode
		}
		return false
	} else {
		return true
	}
}

// Delete a rectangle from non-root part of an index structure.
// Called by d13removeRect.  Descends tree recursively,
// merges branches on the way back up.
// Returns 1 if record not found, 0 if success.
func d13removeRectRec(rect *d13rectT, id interface{}, node *d13nodeT, listNode **d13listNodeT) bool {
	if node.isInternalNode() { // not a leaf node
		for index := 0; index < node.count; index++ {
			if d13overlap(*rect, node.branch[index].rect) {
				if !d13removeRectRec(rect, id, node.branch[index].child, listNode) {
					if node.branch[index].child.count >= d13minNodes {
						// child removed, just resize parent rect
						node.branch[index].rect = d13nodeCover(node.branch[index].child)
					} else {
						// child removed, not enough entries in node, eliminate node
						d13reInsert(node.branch[index].child, listNode)
						d13disconnectBranch(node, index) // Must return after this call as count has changed
					}
					return false
				}
			}
		}
		return true
	} else { // A leaf node
		for index := 0; index < node.count; index++ {
			if node.branch[index].data == id {
				d13disconnectBranch(node, index) // Must return after this call as count has changed
				return false
			}
		}
		return true
	}
}

// Decide whether two rectangles d13overlap.
func d13overlap(rectA, rectB d13rectT) bool {
	for index := 0; index < d13numDims; index++ {
		if rectA.min[index] > rectB.max[index] ||
			rectB.min[index] > rectA.max[index] {
			return false
		}
	}
	return true
}

// Add a node to the reinsertion list.  All its branches will later
// be reinserted into the index structure.
func d13reInsert(node *d13nodeT, listNode **d13listNodeT) {
	newListNode := &d13listNodeT{}
	newListNode.node = node
	newListNode.next = *listNode
	*listNode = newListNode
}

// d13search in an index tree or subtree for all data retangles that d13overlap the argument rectangle.
func d13search(node *d13nodeT, rect d13rectT, foundCount int, resultCallback func(data interface{}) bool) (int, bool) {
	if node.isInternalNode() {
		// This is an internal node in the tree
		for index := 0; index < node.count; index++ {
			if d13overlap(rect, node.branch[index].rect) {
				var ok bool
				foundCount, ok = d13search(node.branch[index].child, rect, foundCount, resultCallback)
				if !ok {
					// The callback indicated to stop searching
					return foundCount, false
				}
			}
		}
	} else {
		// This is a leaf node
		for index := 0; index < node.count; index++ {
			if d13overlap(rect, node.branch[index].rect) {
				id := node.branch[index].data
				foundCount++
				if !resultCallback(id) {
					return foundCount, false // Don't continue searching
				}

			}
		}
	}
	return foundCount, true // Continue searching
}

func d14fmin(a, b float64) float64 {
	if a < b {
		return a
	}
	return b
}
func d14fmax(a, b float64) float64 {
	if a > b {
		return a
	}
	return b
}

const (
	d14numDims            = 14
	d14maxNodes           = 8
	d14minNodes           = d14maxNodes / 2
	d14useSphericalVolume = true // Better split classification, may be slower on some systems
)

var d14unitSphereVolume = []float64{
	0.000000, 2.000000, 3.141593, // Dimension  0,1,2
	4.188790, 4.934802, 5.263789, // Dimension  3,4,5
	5.167713, 4.724766, 4.058712, // Dimension  6,7,8
	3.298509, 2.550164, 1.884104, // Dimension  9,10,11
	1.335263, 0.910629, 0.599265, // Dimension  12,13,14
	0.381443, 0.235331, 0.140981, // Dimension  15,16,17
	0.082146, 0.046622, 0.025807, // Dimension  18,19,20
}[d14numDims]

type d14RTree struct {
	root *d14nodeT ///< Root of tree
}

/// Minimal bounding rectangle (n-dimensional)
type d14rectT struct {
	min [d14numDims]float64 ///< Min dimensions of bounding box
	max [d14numDims]float64 ///< Max dimensions of bounding box
}

/// May be data or may be another subtree
/// The parents level determines this.
/// If the parents level is 0, then this is data
type d14branchT struct {
	rect  d14rectT    ///< Bounds
	child *d14nodeT   ///< Child node
	data  interface{} ///< Data Id or Ptr
}

/// d14nodeT for each branch level
type d14nodeT struct {
	count  int                     ///< Count
	level  int                     ///< Leaf is zero, others positive
	branch [d14maxNodes]d14branchT ///< Branch
}

func (node *d14nodeT) isInternalNode() bool {
	return (node.level > 0) // Not a leaf, but a internal node
}
func (node *d14nodeT) isLeaf() bool {
	return (node.level == 0) // A leaf, contains data
}

/// A link list of nodes for reinsertion after a delete operation
type d14listNodeT struct {
	next *d14listNodeT ///< Next in list
	node *d14nodeT     ///< Node
}

const d14notTaken = -1 // indicates that position

/// Variables for finding a split partition
type d14partitionVarsT struct {
	partition [d14maxNodes + 1]int
	total     int
	minFill   int
	count     [2]int
	cover     [2]d14rectT
	area      [2]float64

	branchBuf      [d14maxNodes + 1]d14branchT
	branchCount    int
	coverSplit     d14rectT
	coverSplitArea float64
}

func d14New() *d14RTree {
	// We only support machine word size simple data type eg. integer index or object pointer.
	// Since we are storing as union with non data branch
	return &d14RTree{
		root: &d14nodeT{},
	}
}

/// Insert entry
/// \param a_min Min of bounding rect
/// \param a_max Max of bounding rect
/// \param a_dataId Positive Id of data.  Maybe zero, but negative numbers not allowed.
func (tr *d14RTree) Insert(min, max [d14numDims]float64, dataId interface{}) {
	var branch d14branchT
	branch.data = dataId
	for axis := 0; axis < d14numDims; axis++ {
		branch.rect.min[axis] = min[axis]
		branch.rect.max[axis] = max[axis]
	}
	d14insertRect(&branch, &tr.root, 0)
}

/// Remove entry
/// \param a_min Min of bounding rect
/// \param a_max Max of bounding rect
/// \param a_dataId Positive Id of data.  Maybe zero, but negative numbers not allowed.
func (tr *d14RTree) Remove(min, max [d14numDims]float64, dataId interface{}) {
	var rect d14rectT
	for axis := 0; axis < d14numDims; axis++ {
		rect.min[axis] = min[axis]
		rect.max[axis] = max[axis]
	}
	d14removeRect(&rect, dataId, &tr.root)
}

/// Find all within d14search rectangle
/// \param a_min Min of d14search bounding rect
/// \param a_max Max of d14search bounding rect
/// \param a_searchResult d14search result array.  Caller should set grow size. Function will reset, not append to array.
/// \param a_resultCallback Callback function to return result.  Callback should return 'true' to continue searching
/// \param a_context User context to pass as parameter to a_resultCallback
/// \return Returns the number of entries found
func (tr *d14RTree) Search(min, max [d14numDims]float64, resultCallback func(data interface{}) bool) int {
	var rect d14rectT
	for axis := 0; axis < d14numDims; axis++ {
		rect.min[axis] = min[axis]
		rect.max[axis] = max[axis]
	}
	foundCount, _ := d14search(tr.root, rect, 0, resultCallback)
	return foundCount
}

/// Count the data elements in this container.  This is slow as no internal counter is maintained.
func (tr *d14RTree) Count() int {
	var count int
	d14countRec(tr.root, &count)
	return count
}

/// Remove all entries from tree
func (tr *d14RTree) RemoveAll() {
	// Delete all existing nodes
	tr.root = &d14nodeT{}
}

func d14countRec(node *d14nodeT, count *int) {
	if node.isInternalNode() { // not a leaf node
		for index := 0; index < node.count; index++ {
			d14countRec(node.branch[index].child, count)
		}
	} else { // A leaf node
		*count += node.count
	}
}

// Inserts a new data rectangle into the index structure.
// Recursively descends tree, propagates splits back up.
// Returns 0 if node was not split.  Old node updated.
// If node was split, returns 1 and sets the pointer pointed to by
// new_node to point to the new node.  Old node updated to become one of two.
// The level argument specifies the number of steps up from the leaf
// level to insert; e.g. a data rectangle goes in at level = 0.
func d14insertRectRec(branch *d14branchT, node *d14nodeT, newNode **d14nodeT, level int) bool {
	// recurse until we reach the correct level for the new record. data records
	// will always be called with a_level == 0 (leaf)
	if node.level > level {
		// Still above level for insertion, go down tree recursively
		var otherNode *d14nodeT
		//var newBranch d14branchT

		// find the optimal branch for this record
		index := d14pickBranch(&branch.rect, node)

		// recursively insert this record into the picked branch
		childWasSplit := d14insertRectRec(branch, node.branch[index].child, &otherNode, level)

		if !childWasSplit {
			// Child was not split. Merge the bounding box of the new record with the
			// existing bounding box
			node.branch[index].rect = d14combineRect(&branch.rect, &(node.branch[index].rect))
			return false
		} else {
			// Child was split. The old branches are now re-partitioned to two nodes
			// so we have to re-calculate the bounding boxes of each node
			node.branch[index].rect = d14nodeCover(node.branch[index].child)
			var newBranch d14branchT
			newBranch.child = otherNode
			newBranch.rect = d14nodeCover(otherNode)

			// The old node is already a child of a_node. Now add the newly-created
			// node to a_node as well. a_node might be split because of that.
			return d14addBranch(&newBranch, node, newNode)
		}
	} else if node.level == level {
		// We have reached level for insertion. Add rect, split if necessary
		return d14addBranch(branch, node, newNode)
	} else {
		// Should never occur
		return false
	}
}

// Insert a data rectangle into an index structure.
// d14insertRect provides for splitting the root;
// returns 1 if root was split, 0 if it was not.
// The level argument specifies the number of steps up from the leaf
// level to insert; e.g. a data rectangle goes in at level = 0.
// InsertRect2 does the recursion.
//
func d14insertRect(branch *d14branchT, root **d14nodeT, level int) bool {
	var newNode *d14nodeT

	if d14insertRectRec(branch, *root, &newNode, level) { // Root split

		// Grow tree taller and new root
		newRoot := &d14nodeT{}
		newRoot.level = (*root).level + 1

		var newBranch d14branchT

		// add old root node as a child of the new root
		newBranch.rect = d14nodeCover(*root)
		newBranch.child = *root
		d14addBranch(&newBranch, newRoot, nil)

		// add the split node as a child of the new root
		newBranch.rect = d14nodeCover(newNode)
		newBranch.child = newNode
		d14addBranch(&newBranch, newRoot, nil)

		// set the new root as the root node
		*root = newRoot

		return true
	}
	return false
}

// Find the smallest rectangle that includes all rectangles in branches of a node.
func d14nodeCover(node *d14nodeT) d14rectT {
	rect := node.branch[0].rect
	for index := 1; index < node.count; index++ {
		rect = d14combineRect(&rect, &(node.branch[index].rect))
	}
	return rect
}

// Add a branch to a node.  Split the node if necessary.
// Returns 0 if node not split.  Old node updated.
// Returns 1 if node split, sets *new_node to address of new node.
// Old node updated, becomes one of two.
func d14addBranch(branch *d14branchT, node *d14nodeT, newNode **d14nodeT) bool {
	if node.count < d14maxNodes { // Split won't be necessary
		node.branch[node.count] = *branch
		node.count++
		return false
	} else {
		d14splitNode(node, branch, newNode)
		return true
	}
}

// Disconnect a dependent node.
// Caller must return (or stop using iteration index) after this as count has changed
func d14disconnectBranch(node *d14nodeT, index int) {
	// Remove element by swapping with the last element to prevent gaps in array
	node.branch[index] = node.branch[node.count-1]
	node.branch[node.count-1].data = nil
	node.branch[node.count-1].child = nil
	node.count--
}

// Pick a branch.  Pick the one that will need the smallest increase
// in area to accomodate the new rectangle.  This will result in the
// least total area for the covering rectangles in the current node.
// In case of a tie, pick the one which was smaller before, to get
// the best resolution when searching.
func d14pickBranch(rect *d14rectT, node *d14nodeT) int {
	var firstTime bool = true
	var increase float64
	var bestIncr float64 = -1
	var area float64
	var bestArea float64
	var best int
	var tempRect d14rectT

	for index := 0; index < node.count; index++ {
		curRect := &node.branch[index].rect
		area = d14calcRectVolume(curRect)
		tempRect = d14combineRect(rect, curRect)
		increase = d14calcRectVolume(&tempRect) - area
		if (increase < bestIncr) || firstTime {
			best = index
			bestArea = area
			bestIncr = increase
			firstTime = false
		} else if (increase == bestIncr) && (area < bestArea) {
			best = index
			bestArea = area
			bestIncr = increase
		}
	}
	return best
}

// Combine two rectangles into larger one containing both
func d14combineRect(rectA, rectB *d14rectT) d14rectT {
	var newRect d14rectT

	for index := 0; index < d14numDims; index++ {
		newRect.min[index] = d14fmin(rectA.min[index], rectB.min[index])
		newRect.max[index] = d14fmax(rectA.max[index], rectB.max[index])
	}

	return newRect
}

// Split a node.
// Divides the nodes branches and the extra one between two nodes.
// Old node is one of the new ones, and one really new one is created.
// Tries more than one method for choosing a partition, uses best result.
func d14splitNode(node *d14nodeT, branch *d14branchT, newNode **d14nodeT) {
	// Could just use local here, but member or external is faster since it is reused
	var localVars d14partitionVarsT
	parVars := &localVars

	// Load all the branches into a buffer, initialize old node
	d14getBranches(node, branch, parVars)

	// Find partition
	d14choosePartition(parVars, d14minNodes)

	// Create a new node to hold (about) half of the branches
	*newNode = &d14nodeT{}
	(*newNode).level = node.level

	// Put branches from buffer into 2 nodes according to the chosen partition
	node.count = 0
	d14loadNodes(node, *newNode, parVars)
}

// Calculate the n-dimensional volume of a rectangle
func d14rectVolume(rect *d14rectT) float64 {
	var volume float64 = 1
	for index := 0; index < d14numDims; index++ {
		volume *= rect.max[index] - rect.min[index]
	}
	return volume
}

// The exact volume of the bounding sphere for the given d14rectT
func d14rectSphericalVolume(rect *d14rectT) float64 {
	var sumOfSquares float64 = 0
	var radius float64

	for index := 0; index < d14numDims; index++ {
		halfExtent := (rect.max[index] - rect.min[index]) * 0.5
		sumOfSquares += halfExtent * halfExtent
	}

	radius = math.Sqrt(sumOfSquares)

	// Pow maybe slow, so test for common dims just use x*x, x*x*x.
	if d14numDims == 5 {
		return (radius * radius * radius * radius * radius * d14unitSphereVolume)
	} else if d14numDims == 4 {
		return (radius * radius * radius * radius * d14unitSphereVolume)
	} else if d14numDims == 3 {
		return (radius * radius * radius * d14unitSphereVolume)
	} else if d14numDims == 2 {
		return (radius * radius * d14unitSphereVolume)
	} else {
		return (math.Pow(radius, d14numDims) * d14unitSphereVolume)
	}
}

// Use one of the methods to calculate retangle volume
func d14calcRectVolume(rect *d14rectT) float64 {
	if d14useSphericalVolume {
		return d14rectSphericalVolume(rect) // Slower but helps certain merge cases
	} else { // RTREE_USE_SPHERICAL_VOLUME
		return d14rectVolume(rect) // Faster but can cause poor merges
	} // RTREE_USE_SPHERICAL_VOLUME
}

// Load branch buffer with branches from full node plus the extra branch.
func d14getBranches(node *d14nodeT, branch *d14branchT, parVars *d14partitionVarsT) {
	// Load the branch buffer
	for index := 0; index < d14maxNodes; index++ {
		parVars.branchBuf[index] = node.branch[index]
	}
	parVars.branchBuf[d14maxNodes] = *branch
	parVars.branchCount = d14maxNodes + 1

	// Calculate rect containing all in the set
	parVars.coverSplit = parVars.branchBuf[0].rect
	for index := 1; index < d14maxNodes+1; index++ {
		parVars.coverSplit = d14combineRect(&parVars.coverSplit, &parVars.branchBuf[index].rect)
	}
	parVars.coverSplitArea = d14calcRectVolume(&parVars.coverSplit)
}

// Method #0 for choosing a partition:
// As the seeds for the two groups, pick the two rects that would waste the
// most area if covered by a single rectangle, i.e. evidently the worst pair
// to have in the same group.
// Of the remaining, one at a time is chosen to be put in one of the two groups.
// The one chosen is the one with the greatest difference in area expansion
// depending on which group - the rect most strongly attracted to one group
// and repelled from the other.
// If one group gets too full (more would force other group to violate min
// fill requirement) then other group gets the rest.
// These last are the ones that can go in either group most easily.
func d14choosePartition(parVars *d14partitionVarsT, minFill int) {
	var biggestDiff float64
	var group, chosen, betterGroup int

	d14initParVars(parVars, parVars.branchCount, minFill)
	d14pickSeeds(parVars)

	for ((parVars.count[0] + parVars.count[1]) < parVars.total) &&
		(parVars.count[0] < (parVars.total - parVars.minFill)) &&
		(parVars.count[1] < (parVars.total - parVars.minFill)) {
		biggestDiff = -1
		for index := 0; index < parVars.total; index++ {
			if d14notTaken == parVars.partition[index] {
				curRect := &parVars.branchBuf[index].rect
				rect0 := d14combineRect(curRect, &parVars.cover[0])
				rect1 := d14combineRect(curRect, &parVars.cover[1])
				growth0 := d14calcRectVolume(&rect0) - parVars.area[0]
				growth1 := d14calcRectVolume(&rect1) - parVars.area[1]
				diff := growth1 - growth0
				if diff >= 0 {
					group = 0
				} else {
					group = 1
					diff = -diff
				}

				if diff > biggestDiff {
					biggestDiff = diff
					chosen = index
					betterGroup = group
				} else if (diff == biggestDiff) && (parVars.count[group] < parVars.count[betterGroup]) {
					chosen = index
					betterGroup = group
				}
			}
		}
		d14classify(chosen, betterGroup, parVars)
	}

	// If one group too full, put remaining rects in the other
	if (parVars.count[0] + parVars.count[1]) < parVars.total {
		if parVars.count[0] >= parVars.total-parVars.minFill {
			group = 1
		} else {
			group = 0
		}
		for index := 0; index < parVars.total; index++ {
			if d14notTaken == parVars.partition[index] {
				d14classify(index, group, parVars)
			}
		}
	}
}

// Copy branches from the buffer into two nodes according to the partition.
func d14loadNodes(nodeA, nodeB *d14nodeT, parVars *d14partitionVarsT) {
	for index := 0; index < parVars.total; index++ {
		targetNodeIndex := parVars.partition[index]
		targetNodes := []*d14nodeT{nodeA, nodeB}

		// It is assured that d14addBranch here will not cause a node split.
		d14addBranch(&parVars.branchBuf[index], targetNodes[targetNodeIndex], nil)
	}
}

// Initialize a d14partitionVarsT structure.
func d14initParVars(parVars *d14partitionVarsT, maxRects, minFill int) {
	parVars.count[0] = 0
	parVars.count[1] = 0
	parVars.area[0] = 0
	parVars.area[1] = 0
	parVars.total = maxRects
	parVars.minFill = minFill
	for index := 0; index < maxRects; index++ {
		parVars.partition[index] = d14notTaken
	}
}

func d14pickSeeds(parVars *d14partitionVarsT) {
	var seed0, seed1 int
	var worst, waste float64
	var area [d14maxNodes + 1]float64

	for index := 0; index < parVars.total; index++ {
		area[index] = d14calcRectVolume(&parVars.branchBuf[index].rect)
	}

	worst = -parVars.coverSplitArea - 1
	for indexA := 0; indexA < parVars.total-1; indexA++ {
		for indexB := indexA + 1; indexB < parVars.total; indexB++ {
			oneRect := d14combineRect(&parVars.branchBuf[indexA].rect, &parVars.branchBuf[indexB].rect)
			waste = d14calcRectVolume(&oneRect) - area[indexA] - area[indexB]
			if waste > worst {
				worst = waste
				seed0 = indexA
				seed1 = indexB
			}
		}
	}

	d14classify(seed0, 0, parVars)
	d14classify(seed1, 1, parVars)
}

// Put a branch in one of the groups.
func d14classify(index, group int, parVars *d14partitionVarsT) {
	parVars.partition[index] = group

	// Calculate combined rect
	if parVars.count[group] == 0 {
		parVars.cover[group] = parVars.branchBuf[index].rect
	} else {
		parVars.cover[group] = d14combineRect(&parVars.branchBuf[index].rect, &parVars.cover[group])
	}

	// Calculate volume of combined rect
	parVars.area[group] = d14calcRectVolume(&parVars.cover[group])

	parVars.count[group]++
}

// Delete a data rectangle from an index structure.
// Pass in a pointer to a d14rectT, the tid of the record, ptr to ptr to root node.
// Returns 1 if record not found, 0 if success.
// d14removeRect provides for eliminating the root.
func d14removeRect(rect *d14rectT, id interface{}, root **d14nodeT) bool {
	var reInsertList *d14listNodeT

	if !d14removeRectRec(rect, id, *root, &reInsertList) {
		// Found and deleted a data item
		// Reinsert any branches from eliminated nodes
		for reInsertList != nil {
			tempNode := reInsertList.node

			for index := 0; index < tempNode.count; index++ {
				// TODO go over this code. should I use (tempNode->m_level - 1)?
				d14insertRect(&tempNode.branch[index], root, tempNode.level)
			}
			reInsertList = reInsertList.next
		}

		// Check for redundant root (not leaf, 1 child) and eliminate TODO replace
		// if with while? In case there is a whole branch of redundant roots...
		if (*root).count == 1 && (*root).isInternalNode() {
			tempNode := (*root).branch[0].child
			*root = tempNode
		}
		return false
	} else {
		return true
	}
}

// Delete a rectangle from non-root part of an index structure.
// Called by d14removeRect.  Descends tree recursively,
// merges branches on the way back up.
// Returns 1 if record not found, 0 if success.
func d14removeRectRec(rect *d14rectT, id interface{}, node *d14nodeT, listNode **d14listNodeT) bool {
	if node.isInternalNode() { // not a leaf node
		for index := 0; index < node.count; index++ {
			if d14overlap(*rect, node.branch[index].rect) {
				if !d14removeRectRec(rect, id, node.branch[index].child, listNode) {
					if node.branch[index].child.count >= d14minNodes {
						// child removed, just resize parent rect
						node.branch[index].rect = d14nodeCover(node.branch[index].child)
					} else {
						// child removed, not enough entries in node, eliminate node
						d14reInsert(node.branch[index].child, listNode)
						d14disconnectBranch(node, index) // Must return after this call as count has changed
					}
					return false
				}
			}
		}
		return true
	} else { // A leaf node
		for index := 0; index < node.count; index++ {
			if node.branch[index].data == id {
				d14disconnectBranch(node, index) // Must return after this call as count has changed
				return false
			}
		}
		return true
	}
}

// Decide whether two rectangles d14overlap.
func d14overlap(rectA, rectB d14rectT) bool {
	for index := 0; index < d14numDims; index++ {
		if rectA.min[index] > rectB.max[index] ||
			rectB.min[index] > rectA.max[index] {
			return false
		}
	}
	return true
}

// Add a node to the reinsertion list.  All its branches will later
// be reinserted into the index structure.
func d14reInsert(node *d14nodeT, listNode **d14listNodeT) {
	newListNode := &d14listNodeT{}
	newListNode.node = node
	newListNode.next = *listNode
	*listNode = newListNode
}

// d14search in an index tree or subtree for all data retangles that d14overlap the argument rectangle.
func d14search(node *d14nodeT, rect d14rectT, foundCount int, resultCallback func(data interface{}) bool) (int, bool) {
	if node.isInternalNode() {
		// This is an internal node in the tree
		for index := 0; index < node.count; index++ {
			if d14overlap(rect, node.branch[index].rect) {
				var ok bool
				foundCount, ok = d14search(node.branch[index].child, rect, foundCount, resultCallback)
				if !ok {
					// The callback indicated to stop searching
					return foundCount, false
				}
			}
		}
	} else {
		// This is a leaf node
		for index := 0; index < node.count; index++ {
			if d14overlap(rect, node.branch[index].rect) {
				id := node.branch[index].data
				foundCount++
				if !resultCallback(id) {
					return foundCount, false // Don't continue searching
				}

			}
		}
	}
	return foundCount, true // Continue searching
}

func d15fmin(a, b float64) float64 {
	if a < b {
		return a
	}
	return b
}
func d15fmax(a, b float64) float64 {
	if a > b {
		return a
	}
	return b
}

const (
	d15numDims            = 15
	d15maxNodes           = 8
	d15minNodes           = d15maxNodes / 2
	d15useSphericalVolume = true // Better split classification, may be slower on some systems
)

var d15unitSphereVolume = []float64{
	0.000000, 2.000000, 3.141593, // Dimension  0,1,2
	4.188790, 4.934802, 5.263789, // Dimension  3,4,5
	5.167713, 4.724766, 4.058712, // Dimension  6,7,8
	3.298509, 2.550164, 1.884104, // Dimension  9,10,11
	1.335263, 0.910629, 0.599265, // Dimension  12,13,14
	0.381443, 0.235331, 0.140981, // Dimension  15,16,17
	0.082146, 0.046622, 0.025807, // Dimension  18,19,20
}[d15numDims]

type d15RTree struct {
	root *d15nodeT ///< Root of tree
}

/// Minimal bounding rectangle (n-dimensional)
type d15rectT struct {
	min [d15numDims]float64 ///< Min dimensions of bounding box
	max [d15numDims]float64 ///< Max dimensions of bounding box
}

/// May be data or may be another subtree
/// The parents level determines this.
/// If the parents level is 0, then this is data
type d15branchT struct {
	rect  d15rectT    ///< Bounds
	child *d15nodeT   ///< Child node
	data  interface{} ///< Data Id or Ptr
}

/// d15nodeT for each branch level
type d15nodeT struct {
	count  int                     ///< Count
	level  int                     ///< Leaf is zero, others positive
	branch [d15maxNodes]d15branchT ///< Branch
}

func (node *d15nodeT) isInternalNode() bool {
	return (node.level > 0) // Not a leaf, but a internal node
}
func (node *d15nodeT) isLeaf() bool {
	return (node.level == 0) // A leaf, contains data
}

/// A link list of nodes for reinsertion after a delete operation
type d15listNodeT struct {
	next *d15listNodeT ///< Next in list
	node *d15nodeT     ///< Node
}

const d15notTaken = -1 // indicates that position

/// Variables for finding a split partition
type d15partitionVarsT struct {
	partition [d15maxNodes + 1]int
	total     int
	minFill   int
	count     [2]int
	cover     [2]d15rectT
	area      [2]float64

	branchBuf      [d15maxNodes + 1]d15branchT
	branchCount    int
	coverSplit     d15rectT
	coverSplitArea float64
}

func d15New() *d15RTree {
	// We only support machine word size simple data type eg. integer index or object pointer.
	// Since we are storing as union with non data branch
	return &d15RTree{
		root: &d15nodeT{},
	}
}

/// Insert entry
/// \param a_min Min of bounding rect
/// \param a_max Max of bounding rect
/// \param a_dataId Positive Id of data.  Maybe zero, but negative numbers not allowed.
func (tr *d15RTree) Insert(min, max [d15numDims]float64, dataId interface{}) {
	var branch d15branchT
	branch.data = dataId
	for axis := 0; axis < d15numDims; axis++ {
		branch.rect.min[axis] = min[axis]
		branch.rect.max[axis] = max[axis]
	}
	d15insertRect(&branch, &tr.root, 0)
}

/// Remove entry
/// \param a_min Min of bounding rect
/// \param a_max Max of bounding rect
/// \param a_dataId Positive Id of data.  Maybe zero, but negative numbers not allowed.
func (tr *d15RTree) Remove(min, max [d15numDims]float64, dataId interface{}) {
	var rect d15rectT
	for axis := 0; axis < d15numDims; axis++ {
		rect.min[axis] = min[axis]
		rect.max[axis] = max[axis]
	}
	d15removeRect(&rect, dataId, &tr.root)
}

/// Find all within d15search rectangle
/// \param a_min Min of d15search bounding rect
/// \param a_max Max of d15search bounding rect
/// \param a_searchResult d15search result array.  Caller should set grow size. Function will reset, not append to array.
/// \param a_resultCallback Callback function to return result.  Callback should return 'true' to continue searching
/// \param a_context User context to pass as parameter to a_resultCallback
/// \return Returns the number of entries found
func (tr *d15RTree) Search(min, max [d15numDims]float64, resultCallback func(data interface{}) bool) int {
	var rect d15rectT
	for axis := 0; axis < d15numDims; axis++ {
		rect.min[axis] = min[axis]
		rect.max[axis] = max[axis]
	}
	foundCount, _ := d15search(tr.root, rect, 0, resultCallback)
	return foundCount
}

/// Count the data elements in this container.  This is slow as no internal counter is maintained.
func (tr *d15RTree) Count() int {
	var count int
	d15countRec(tr.root, &count)
	return count
}

/// Remove all entries from tree
func (tr *d15RTree) RemoveAll() {
	// Delete all existing nodes
	tr.root = &d15nodeT{}
}

func d15countRec(node *d15nodeT, count *int) {
	if node.isInternalNode() { // not a leaf node
		for index := 0; index < node.count; index++ {
			d15countRec(node.branch[index].child, count)
		}
	} else { // A leaf node
		*count += node.count
	}
}

// Inserts a new data rectangle into the index structure.
// Recursively descends tree, propagates splits back up.
// Returns 0 if node was not split.  Old node updated.
// If node was split, returns 1 and sets the pointer pointed to by
// new_node to point to the new node.  Old node updated to become one of two.
// The level argument specifies the number of steps up from the leaf
// level to insert; e.g. a data rectangle goes in at level = 0.
func d15insertRectRec(branch *d15branchT, node *d15nodeT, newNode **d15nodeT, level int) bool {
	// recurse until we reach the correct level for the new record. data records
	// will always be called with a_level == 0 (leaf)
	if node.level > level {
		// Still above level for insertion, go down tree recursively
		var otherNode *d15nodeT
		//var newBranch d15branchT

		// find the optimal branch for this record
		index := d15pickBranch(&branch.rect, node)

		// recursively insert this record into the picked branch
		childWasSplit := d15insertRectRec(branch, node.branch[index].child, &otherNode, level)

		if !childWasSplit {
			// Child was not split. Merge the bounding box of the new record with the
			// existing bounding box
			node.branch[index].rect = d15combineRect(&branch.rect, &(node.branch[index].rect))
			return false
		} else {
			// Child was split. The old branches are now re-partitioned to two nodes
			// so we have to re-calculate the bounding boxes of each node
			node.branch[index].rect = d15nodeCover(node.branch[index].child)
			var newBranch d15branchT
			newBranch.child = otherNode
			newBranch.rect = d15nodeCover(otherNode)

			// The old node is already a child of a_node. Now add the newly-created
			// node to a_node as well. a_node might be split because of that.
			return d15addBranch(&newBranch, node, newNode)
		}
	} else if node.level == level {
		// We have reached level for insertion. Add rect, split if necessary
		return d15addBranch(branch, node, newNode)
	} else {
		// Should never occur
		return false
	}
}

// Insert a data rectangle into an index structure.
// d15insertRect provides for splitting the root;
// returns 1 if root was split, 0 if it was not.
// The level argument specifies the number of steps up from the leaf
// level to insert; e.g. a data rectangle goes in at level = 0.
// InsertRect2 does the recursion.
//
func d15insertRect(branch *d15branchT, root **d15nodeT, level int) bool {
	var newNode *d15nodeT

	if d15insertRectRec(branch, *root, &newNode, level) { // Root split

		// Grow tree taller and new root
		newRoot := &d15nodeT{}
		newRoot.level = (*root).level + 1

		var newBranch d15branchT

		// add old root node as a child of the new root
		newBranch.rect = d15nodeCover(*root)
		newBranch.child = *root
		d15addBranch(&newBranch, newRoot, nil)

		// add the split node as a child of the new root
		newBranch.rect = d15nodeCover(newNode)
		newBranch.child = newNode
		d15addBranch(&newBranch, newRoot, nil)

		// set the new root as the root node
		*root = newRoot

		return true
	}
	return false
}

// Find the smallest rectangle that includes all rectangles in branches of a node.
func d15nodeCover(node *d15nodeT) d15rectT {
	rect := node.branch[0].rect
	for index := 1; index < node.count; index++ {
		rect = d15combineRect(&rect, &(node.branch[index].rect))
	}
	return rect
}

// Add a branch to a node.  Split the node if necessary.
// Returns 0 if node not split.  Old node updated.
// Returns 1 if node split, sets *new_node to address of new node.
// Old node updated, becomes one of two.
func d15addBranch(branch *d15branchT, node *d15nodeT, newNode **d15nodeT) bool {
	if node.count < d15maxNodes { // Split won't be necessary
		node.branch[node.count] = *branch
		node.count++
		return false
	} else {
		d15splitNode(node, branch, newNode)
		return true
	}
}

// Disconnect a dependent node.
// Caller must return (or stop using iteration index) after this as count has changed
func d15disconnectBranch(node *d15nodeT, index int) {
	// Remove element by swapping with the last element to prevent gaps in array
	node.branch[index] = node.branch[node.count-1]
	node.branch[node.count-1].data = nil
	node.branch[node.count-1].child = nil
	node.count--
}

// Pick a branch.  Pick the one that will need the smallest increase
// in area to accomodate the new rectangle.  This will result in the
// least total area for the covering rectangles in the current node.
// In case of a tie, pick the one which was smaller before, to get
// the best resolution when searching.
func d15pickBranch(rect *d15rectT, node *d15nodeT) int {
	var firstTime bool = true
	var increase float64
	var bestIncr float64 = -1
	var area float64
	var bestArea float64
	var best int
	var tempRect d15rectT

	for index := 0; index < node.count; index++ {
		curRect := &node.branch[index].rect
		area = d15calcRectVolume(curRect)
		tempRect = d15combineRect(rect, curRect)
		increase = d15calcRectVolume(&tempRect) - area
		if (increase < bestIncr) || firstTime {
			best = index
			bestArea = area
			bestIncr = increase
			firstTime = false
		} else if (increase == bestIncr) && (area < bestArea) {
			best = index
			bestArea = area
			bestIncr = increase
		}
	}
	return best
}

// Combine two rectangles into larger one containing both
func d15combineRect(rectA, rectB *d15rectT) d15rectT {
	var newRect d15rectT

	for index := 0; index < d15numDims; index++ {
		newRect.min[index] = d15fmin(rectA.min[index], rectB.min[index])
		newRect.max[index] = d15fmax(rectA.max[index], rectB.max[index])
	}

	return newRect
}

// Split a node.
// Divides the nodes branches and the extra one between two nodes.
// Old node is one of the new ones, and one really new one is created.
// Tries more than one method for choosing a partition, uses best result.
func d15splitNode(node *d15nodeT, branch *d15branchT, newNode **d15nodeT) {
	// Could just use local here, but member or external is faster since it is reused
	var localVars d15partitionVarsT
	parVars := &localVars

	// Load all the branches into a buffer, initialize old node
	d15getBranches(node, branch, parVars)

	// Find partition
	d15choosePartition(parVars, d15minNodes)

	// Create a new node to hold (about) half of the branches
	*newNode = &d15nodeT{}
	(*newNode).level = node.level

	// Put branches from buffer into 2 nodes according to the chosen partition
	node.count = 0
	d15loadNodes(node, *newNode, parVars)
}

// Calculate the n-dimensional volume of a rectangle
func d15rectVolume(rect *d15rectT) float64 {
	var volume float64 = 1
	for index := 0; index < d15numDims; index++ {
		volume *= rect.max[index] - rect.min[index]
	}
	return volume
}

// The exact volume of the bounding sphere for the given d15rectT
func d15rectSphericalVolume(rect *d15rectT) float64 {
	var sumOfSquares float64 = 0
	var radius float64

	for index := 0; index < d15numDims; index++ {
		halfExtent := (rect.max[index] - rect.min[index]) * 0.5
		sumOfSquares += halfExtent * halfExtent
	}

	radius = math.Sqrt(sumOfSquares)

	// Pow maybe slow, so test for common dims just use x*x, x*x*x.
	if d15numDims == 5 {
		return (radius * radius * radius * radius * radius * d15unitSphereVolume)
	} else if d15numDims == 4 {
		return (radius * radius * radius * radius * d15unitSphereVolume)
	} else if d15numDims == 3 {
		return (radius * radius * radius * d15unitSphereVolume)
	} else if d15numDims == 2 {
		return (radius * radius * d15unitSphereVolume)
	} else {
		return (math.Pow(radius, d15numDims) * d15unitSphereVolume)
	}
}

// Use one of the methods to calculate retangle volume
func d15calcRectVolume(rect *d15rectT) float64 {
	if d15useSphericalVolume {
		return d15rectSphericalVolume(rect) // Slower but helps certain merge cases
	} else { // RTREE_USE_SPHERICAL_VOLUME
		return d15rectVolume(rect) // Faster but can cause poor merges
	} // RTREE_USE_SPHERICAL_VOLUME
}

// Load branch buffer with branches from full node plus the extra branch.
func d15getBranches(node *d15nodeT, branch *d15branchT, parVars *d15partitionVarsT) {
	// Load the branch buffer
	for index := 0; index < d15maxNodes; index++ {
		parVars.branchBuf[index] = node.branch[index]
	}
	parVars.branchBuf[d15maxNodes] = *branch
	parVars.branchCount = d15maxNodes + 1

	// Calculate rect containing all in the set
	parVars.coverSplit = parVars.branchBuf[0].rect
	for index := 1; index < d15maxNodes+1; index++ {
		parVars.coverSplit = d15combineRect(&parVars.coverSplit, &parVars.branchBuf[index].rect)
	}
	parVars.coverSplitArea = d15calcRectVolume(&parVars.coverSplit)
}

// Method #0 for choosing a partition:
// As the seeds for the two groups, pick the two rects that would waste the
// most area if covered by a single rectangle, i.e. evidently the worst pair
// to have in the same group.
// Of the remaining, one at a time is chosen to be put in one of the two groups.
// The one chosen is the one with the greatest difference in area expansion
// depending on which group - the rect most strongly attracted to one group
// and repelled from the other.
// If one group gets too full (more would force other group to violate min
// fill requirement) then other group gets the rest.
// These last are the ones that can go in either group most easily.
func d15choosePartition(parVars *d15partitionVarsT, minFill int) {
	var biggestDiff float64
	var group, chosen, betterGroup int

	d15initParVars(parVars, parVars.branchCount, minFill)
	d15pickSeeds(parVars)

	for ((parVars.count[0] + parVars.count[1]) < parVars.total) &&
		(parVars.count[0] < (parVars.total - parVars.minFill)) &&
		(parVars.count[1] < (parVars.total - parVars.minFill)) {
		biggestDiff = -1
		for index := 0; index < parVars.total; index++ {
			if d15notTaken == parVars.partition[index] {
				curRect := &parVars.branchBuf[index].rect
				rect0 := d15combineRect(curRect, &parVars.cover[0])
				rect1 := d15combineRect(curRect, &parVars.cover[1])
				growth0 := d15calcRectVolume(&rect0) - parVars.area[0]
				growth1 := d15calcRectVolume(&rect1) - parVars.area[1]
				diff := growth1 - growth0
				if diff >= 0 {
					group = 0
				} else {
					group = 1
					diff = -diff
				}

				if diff > biggestDiff {
					biggestDiff = diff
					chosen = index
					betterGroup = group
				} else if (diff == biggestDiff) && (parVars.count[group] < parVars.count[betterGroup]) {
					chosen = index
					betterGroup = group
				}
			}
		}
		d15classify(chosen, betterGroup, parVars)
	}

	// If one group too full, put remaining rects in the other
	if (parVars.count[0] + parVars.count[1]) < parVars.total {
		if parVars.count[0] >= parVars.total-parVars.minFill {
			group = 1
		} else {
			group = 0
		}
		for index := 0; index < parVars.total; index++ {
			if d15notTaken == parVars.partition[index] {
				d15classify(index, group, parVars)
			}
		}
	}
}

// Copy branches from the buffer into two nodes according to the partition.
func d15loadNodes(nodeA, nodeB *d15nodeT, parVars *d15partitionVarsT) {
	for index := 0; index < parVars.total; index++ {
		targetNodeIndex := parVars.partition[index]
		targetNodes := []*d15nodeT{nodeA, nodeB}

		// It is assured that d15addBranch here will not cause a node split.
		d15addBranch(&parVars.branchBuf[index], targetNodes[targetNodeIndex], nil)
	}
}

// Initialize a d15partitionVarsT structure.
func d15initParVars(parVars *d15partitionVarsT, maxRects, minFill int) {
	parVars.count[0] = 0
	parVars.count[1] = 0
	parVars.area[0] = 0
	parVars.area[1] = 0
	parVars.total = maxRects
	parVars.minFill = minFill
	for index := 0; index < maxRects; index++ {
		parVars.partition[index] = d15notTaken
	}
}

func d15pickSeeds(parVars *d15partitionVarsT) {
	var seed0, seed1 int
	var worst, waste float64
	var area [d15maxNodes + 1]float64

	for index := 0; index < parVars.total; index++ {
		area[index] = d15calcRectVolume(&parVars.branchBuf[index].rect)
	}

	worst = -parVars.coverSplitArea - 1
	for indexA := 0; indexA < parVars.total-1; indexA++ {
		for indexB := indexA + 1; indexB < parVars.total; indexB++ {
			oneRect := d15combineRect(&parVars.branchBuf[indexA].rect, &parVars.branchBuf[indexB].rect)
			waste = d15calcRectVolume(&oneRect) - area[indexA] - area[indexB]
			if waste > worst {
				worst = waste
				seed0 = indexA
				seed1 = indexB
			}
		}
	}

	d15classify(seed0, 0, parVars)
	d15classify(seed1, 1, parVars)
}

// Put a branch in one of the groups.
func d15classify(index, group int, parVars *d15partitionVarsT) {
	parVars.partition[index] = group

	// Calculate combined rect
	if parVars.count[group] == 0 {
		parVars.cover[group] = parVars.branchBuf[index].rect
	} else {
		parVars.cover[group] = d15combineRect(&parVars.branchBuf[index].rect, &parVars.cover[group])
	}

	// Calculate volume of combined rect
	parVars.area[group] = d15calcRectVolume(&parVars.cover[group])

	parVars.count[group]++
}

// Delete a data rectangle from an index structure.
// Pass in a pointer to a d15rectT, the tid of the record, ptr to ptr to root node.
// Returns 1 if record not found, 0 if success.
// d15removeRect provides for eliminating the root.
func d15removeRect(rect *d15rectT, id interface{}, root **d15nodeT) bool {
	var reInsertList *d15listNodeT

	if !d15removeRectRec(rect, id, *root, &reInsertList) {
		// Found and deleted a data item
		// Reinsert any branches from eliminated nodes
		for reInsertList != nil {
			tempNode := reInsertList.node

			for index := 0; index < tempNode.count; index++ {
				// TODO go over this code. should I use (tempNode->m_level - 1)?
				d15insertRect(&tempNode.branch[index], root, tempNode.level)
			}
			reInsertList = reInsertList.next
		}

		// Check for redundant root (not leaf, 1 child) and eliminate TODO replace
		// if with while? In case there is a whole branch of redundant roots...
		if (*root).count == 1 && (*root).isInternalNode() {
			tempNode := (*root).branch[0].child
			*root = tempNode
		}
		return false
	} else {
		return true
	}
}

// Delete a rectangle from non-root part of an index structure.
// Called by d15removeRect.  Descends tree recursively,
// merges branches on the way back up.
// Returns 1 if record not found, 0 if success.
func d15removeRectRec(rect *d15rectT, id interface{}, node *d15nodeT, listNode **d15listNodeT) bool {
	if node.isInternalNode() { // not a leaf node
		for index := 0; index < node.count; index++ {
			if d15overlap(*rect, node.branch[index].rect) {
				if !d15removeRectRec(rect, id, node.branch[index].child, listNode) {
					if node.branch[index].child.count >= d15minNodes {
						// child removed, just resize parent rect
						node.branch[index].rect = d15nodeCover(node.branch[index].child)
					} else {
						// child removed, not enough entries in node, eliminate node
						d15reInsert(node.branch[index].child, listNode)
						d15disconnectBranch(node, index) // Must return after this call as count has changed
					}
					return false
				}
			}
		}
		return true
	} else { // A leaf node
		for index := 0; index < node.count; index++ {
			if node.branch[index].data == id {
				d15disconnectBranch(node, index) // Must return after this call as count has changed
				return false
			}
		}
		return true
	}
}

// Decide whether two rectangles d15overlap.
func d15overlap(rectA, rectB d15rectT) bool {
	for index := 0; index < d15numDims; index++ {
		if rectA.min[index] > rectB.max[index] ||
			rectB.min[index] > rectA.max[index] {
			return false
		}
	}
	return true
}

// Add a node to the reinsertion list.  All its branches will later
// be reinserted into the index structure.
func d15reInsert(node *d15nodeT, listNode **d15listNodeT) {
	newListNode := &d15listNodeT{}
	newListNode.node = node
	newListNode.next = *listNode
	*listNode = newListNode
}

// d15search in an index tree or subtree for all data retangles that d15overlap the argument rectangle.
func d15search(node *d15nodeT, rect d15rectT, foundCount int, resultCallback func(data interface{}) bool) (int, bool) {
	if node.isInternalNode() {
		// This is an internal node in the tree
		for index := 0; index < node.count; index++ {
			if d15overlap(rect, node.branch[index].rect) {
				var ok bool
				foundCount, ok = d15search(node.branch[index].child, rect, foundCount, resultCallback)
				if !ok {
					// The callback indicated to stop searching
					return foundCount, false
				}
			}
		}
	} else {
		// This is a leaf node
		for index := 0; index < node.count; index++ {
			if d15overlap(rect, node.branch[index].rect) {
				id := node.branch[index].data
				foundCount++
				if !resultCallback(id) {
					return foundCount, false // Don't continue searching
				}

			}
		}
	}
	return foundCount, true // Continue searching
}

func d16fmin(a, b float64) float64 {
	if a < b {
		return a
	}
	return b
}
func d16fmax(a, b float64) float64 {
	if a > b {
		return a
	}
	return b
}

const (
	d16numDims            = 16
	d16maxNodes           = 8
	d16minNodes           = d16maxNodes / 2
	d16useSphericalVolume = true // Better split classification, may be slower on some systems
)

var d16unitSphereVolume = []float64{
	0.000000, 2.000000, 3.141593, // Dimension  0,1,2
	4.188790, 4.934802, 5.263789, // Dimension  3,4,5
	5.167713, 4.724766, 4.058712, // Dimension  6,7,8
	3.298509, 2.550164, 1.884104, // Dimension  9,10,11
	1.335263, 0.910629, 0.599265, // Dimension  12,13,14
	0.381443, 0.235331, 0.140981, // Dimension  15,16,17
	0.082146, 0.046622, 0.025807, // Dimension  18,19,20
}[d16numDims]

type d16RTree struct {
	root *d16nodeT ///< Root of tree
}

/// Minimal bounding rectangle (n-dimensional)
type d16rectT struct {
	min [d16numDims]float64 ///< Min dimensions of bounding box
	max [d16numDims]float64 ///< Max dimensions of bounding box
}

/// May be data or may be another subtree
/// The parents level determines this.
/// If the parents level is 0, then this is data
type d16branchT struct {
	rect  d16rectT    ///< Bounds
	child *d16nodeT   ///< Child node
	data  interface{} ///< Data Id or Ptr
}

/// d16nodeT for each branch level
type d16nodeT struct {
	count  int                     ///< Count
	level  int                     ///< Leaf is zero, others positive
	branch [d16maxNodes]d16branchT ///< Branch
}

func (node *d16nodeT) isInternalNode() bool {
	return (node.level > 0) // Not a leaf, but a internal node
}
func (node *d16nodeT) isLeaf() bool {
	return (node.level == 0) // A leaf, contains data
}

/// A link list of nodes for reinsertion after a delete operation
type d16listNodeT struct {
	next *d16listNodeT ///< Next in list
	node *d16nodeT     ///< Node
}

const d16notTaken = -1 // indicates that position

/// Variables for finding a split partition
type d16partitionVarsT struct {
	partition [d16maxNodes + 1]int
	total     int
	minFill   int
	count     [2]int
	cover     [2]d16rectT
	area      [2]float64

	branchBuf      [d16maxNodes + 1]d16branchT
	branchCount    int
	coverSplit     d16rectT
	coverSplitArea float64
}

func d16New() *d16RTree {
	// We only support machine word size simple data type eg. integer index or object pointer.
	// Since we are storing as union with non data branch
	return &d16RTree{
		root: &d16nodeT{},
	}
}

/// Insert entry
/// \param a_min Min of bounding rect
/// \param a_max Max of bounding rect
/// \param a_dataId Positive Id of data.  Maybe zero, but negative numbers not allowed.
func (tr *d16RTree) Insert(min, max [d16numDims]float64, dataId interface{}) {
	var branch d16branchT
	branch.data = dataId
	for axis := 0; axis < d16numDims; axis++ {
		branch.rect.min[axis] = min[axis]
		branch.rect.max[axis] = max[axis]
	}
	d16insertRect(&branch, &tr.root, 0)
}

/// Remove entry
/// \param a_min Min of bounding rect
/// \param a_max Max of bounding rect
/// \param a_dataId Positive Id of data.  Maybe zero, but negative numbers not allowed.
func (tr *d16RTree) Remove(min, max [d16numDims]float64, dataId interface{}) {
	var rect d16rectT
	for axis := 0; axis < d16numDims; axis++ {
		rect.min[axis] = min[axis]
		rect.max[axis] = max[axis]
	}
	d16removeRect(&rect, dataId, &tr.root)
}

/// Find all within d16search rectangle
/// \param a_min Min of d16search bounding rect
/// \param a_max Max of d16search bounding rect
/// \param a_searchResult d16search result array.  Caller should set grow size. Function will reset, not append to array.
/// \param a_resultCallback Callback function to return result.  Callback should return 'true' to continue searching
/// \param a_context User context to pass as parameter to a_resultCallback
/// \return Returns the number of entries found
func (tr *d16RTree) Search(min, max [d16numDims]float64, resultCallback func(data interface{}) bool) int {
	var rect d16rectT
	for axis := 0; axis < d16numDims; axis++ {
		rect.min[axis] = min[axis]
		rect.max[axis] = max[axis]
	}
	foundCount, _ := d16search(tr.root, rect, 0, resultCallback)
	return foundCount
}

/// Count the data elements in this container.  This is slow as no internal counter is maintained.
func (tr *d16RTree) Count() int {
	var count int
	d16countRec(tr.root, &count)
	return count
}

/// Remove all entries from tree
func (tr *d16RTree) RemoveAll() {
	// Delete all existing nodes
	tr.root = &d16nodeT{}
}

func d16countRec(node *d16nodeT, count *int) {
	if node.isInternalNode() { // not a leaf node
		for index := 0; index < node.count; index++ {
			d16countRec(node.branch[index].child, count)
		}
	} else { // A leaf node
		*count += node.count
	}
}

// Inserts a new data rectangle into the index structure.
// Recursively descends tree, propagates splits back up.
// Returns 0 if node was not split.  Old node updated.
// If node was split, returns 1 and sets the pointer pointed to by
// new_node to point to the new node.  Old node updated to become one of two.
// The level argument specifies the number of steps up from the leaf
// level to insert; e.g. a data rectangle goes in at level = 0.
func d16insertRectRec(branch *d16branchT, node *d16nodeT, newNode **d16nodeT, level int) bool {
	// recurse until we reach the correct level for the new record. data records
	// will always be called with a_level == 0 (leaf)
	if node.level > level {
		// Still above level for insertion, go down tree recursively
		var otherNode *d16nodeT
		//var newBranch d16branchT

		// find the optimal branch for this record
		index := d16pickBranch(&branch.rect, node)

		// recursively insert this record into the picked branch
		childWasSplit := d16insertRectRec(branch, node.branch[index].child, &otherNode, level)

		if !childWasSplit {
			// Child was not split. Merge the bounding box of the new record with the
			// existing bounding box
			node.branch[index].rect = d16combineRect(&branch.rect, &(node.branch[index].rect))
			return false
		} else {
			// Child was split. The old branches are now re-partitioned to two nodes
			// so we have to re-calculate the bounding boxes of each node
			node.branch[index].rect = d16nodeCover(node.branch[index].child)
			var newBranch d16branchT
			newBranch.child = otherNode
			newBranch.rect = d16nodeCover(otherNode)

			// The old node is already a child of a_node. Now add the newly-created
			// node to a_node as well. a_node might be split because of that.
			return d16addBranch(&newBranch, node, newNode)
		}
	} else if node.level == level {
		// We have reached level for insertion. Add rect, split if necessary
		return d16addBranch(branch, node, newNode)
	} else {
		// Should never occur
		return false
	}
}

// Insert a data rectangle into an index structure.
// d16insertRect provides for splitting the root;
// returns 1 if root was split, 0 if it was not.
// The level argument specifies the number of steps up from the leaf
// level to insert; e.g. a data rectangle goes in at level = 0.
// InsertRect2 does the recursion.
//
func d16insertRect(branch *d16branchT, root **d16nodeT, level int) bool {
	var newNode *d16nodeT

	if d16insertRectRec(branch, *root, &newNode, level) { // Root split

		// Grow tree taller and new root
		newRoot := &d16nodeT{}
		newRoot.level = (*root).level + 1

		var newBranch d16branchT

		// add old root node as a child of the new root
		newBranch.rect = d16nodeCover(*root)
		newBranch.child = *root
		d16addBranch(&newBranch, newRoot, nil)

		// add the split node as a child of the new root
		newBranch.rect = d16nodeCover(newNode)
		newBranch.child = newNode
		d16addBranch(&newBranch, newRoot, nil)

		// set the new root as the root node
		*root = newRoot

		return true
	}
	return false
}

// Find the smallest rectangle that includes all rectangles in branches of a node.
func d16nodeCover(node *d16nodeT) d16rectT {
	rect := node.branch[0].rect
	for index := 1; index < node.count; index++ {
		rect = d16combineRect(&rect, &(node.branch[index].rect))
	}
	return rect
}

// Add a branch to a node.  Split the node if necessary.
// Returns 0 if node not split.  Old node updated.
// Returns 1 if node split, sets *new_node to address of new node.
// Old node updated, becomes one of two.
func d16addBranch(branch *d16branchT, node *d16nodeT, newNode **d16nodeT) bool {
	if node.count < d16maxNodes { // Split won't be necessary
		node.branch[node.count] = *branch
		node.count++
		return false
	} else {
		d16splitNode(node, branch, newNode)
		return true
	}
}

// Disconnect a dependent node.
// Caller must return (or stop using iteration index) after this as count has changed
func d16disconnectBranch(node *d16nodeT, index int) {
	// Remove element by swapping with the last element to prevent gaps in array
	node.branch[index] = node.branch[node.count-1]
	node.branch[node.count-1].data = nil
	node.branch[node.count-1].child = nil
	node.count--
}

// Pick a branch.  Pick the one that will need the smallest increase
// in area to accomodate the new rectangle.  This will result in the
// least total area for the covering rectangles in the current node.
// In case of a tie, pick the one which was smaller before, to get
// the best resolution when searching.
func d16pickBranch(rect *d16rectT, node *d16nodeT) int {
	var firstTime bool = true
	var increase float64
	var bestIncr float64 = -1
	var area float64
	var bestArea float64
	var best int
	var tempRect d16rectT

	for index := 0; index < node.count; index++ {
		curRect := &node.branch[index].rect
		area = d16calcRectVolume(curRect)
		tempRect = d16combineRect(rect, curRect)
		increase = d16calcRectVolume(&tempRect) - area
		if (increase < bestIncr) || firstTime {
			best = index
			bestArea = area
			bestIncr = increase
			firstTime = false
		} else if (increase == bestIncr) && (area < bestArea) {
			best = index
			bestArea = area
			bestIncr = increase
		}
	}
	return best
}

// Combine two rectangles into larger one containing both
func d16combineRect(rectA, rectB *d16rectT) d16rectT {
	var newRect d16rectT

	for index := 0; index < d16numDims; index++ {
		newRect.min[index] = d16fmin(rectA.min[index], rectB.min[index])
		newRect.max[index] = d16fmax(rectA.max[index], rectB.max[index])
	}

	return newRect
}

// Split a node.
// Divides the nodes branches and the extra one between two nodes.
// Old node is one of the new ones, and one really new one is created.
// Tries more than one method for choosing a partition, uses best result.
func d16splitNode(node *d16nodeT, branch *d16branchT, newNode **d16nodeT) {
	// Could just use local here, but member or external is faster since it is reused
	var localVars d16partitionVarsT
	parVars := &localVars

	// Load all the branches into a buffer, initialize old node
	d16getBranches(node, branch, parVars)

	// Find partition
	d16choosePartition(parVars, d16minNodes)

	// Create a new node to hold (about) half of the branches
	*newNode = &d16nodeT{}
	(*newNode).level = node.level

	// Put branches from buffer into 2 nodes according to the chosen partition
	node.count = 0
	d16loadNodes(node, *newNode, parVars)
}

// Calculate the n-dimensional volume of a rectangle
func d16rectVolume(rect *d16rectT) float64 {
	var volume float64 = 1
	for index := 0; index < d16numDims; index++ {
		volume *= rect.max[index] - rect.min[index]
	}
	return volume
}

// The exact volume of the bounding sphere for the given d16rectT
func d16rectSphericalVolume(rect *d16rectT) float64 {
	var sumOfSquares float64 = 0
	var radius float64

	for index := 0; index < d16numDims; index++ {
		halfExtent := (rect.max[index] - rect.min[index]) * 0.5
		sumOfSquares += halfExtent * halfExtent
	}

	radius = math.Sqrt(sumOfSquares)

	// Pow maybe slow, so test for common dims just use x*x, x*x*x.
	if d16numDims == 5 {
		return (radius * radius * radius * radius * radius * d16unitSphereVolume)
	} else if d16numDims == 4 {
		return (radius * radius * radius * radius * d16unitSphereVolume)
	} else if d16numDims == 3 {
		return (radius * radius * radius * d16unitSphereVolume)
	} else if d16numDims == 2 {
		return (radius * radius * d16unitSphereVolume)
	} else {
		return (math.Pow(radius, d16numDims) * d16unitSphereVolume)
	}
}

// Use one of the methods to calculate retangle volume
func d16calcRectVolume(rect *d16rectT) float64 {
	if d16useSphericalVolume {
		return d16rectSphericalVolume(rect) // Slower but helps certain merge cases
	} else { // RTREE_USE_SPHERICAL_VOLUME
		return d16rectVolume(rect) // Faster but can cause poor merges
	} // RTREE_USE_SPHERICAL_VOLUME
}

// Load branch buffer with branches from full node plus the extra branch.
func d16getBranches(node *d16nodeT, branch *d16branchT, parVars *d16partitionVarsT) {
	// Load the branch buffer
	for index := 0; index < d16maxNodes; index++ {
		parVars.branchBuf[index] = node.branch[index]
	}
	parVars.branchBuf[d16maxNodes] = *branch
	parVars.branchCount = d16maxNodes + 1

	// Calculate rect containing all in the set
	parVars.coverSplit = parVars.branchBuf[0].rect
	for index := 1; index < d16maxNodes+1; index++ {
		parVars.coverSplit = d16combineRect(&parVars.coverSplit, &parVars.branchBuf[index].rect)
	}
	parVars.coverSplitArea = d16calcRectVolume(&parVars.coverSplit)
}

// Method #0 for choosing a partition:
// As the seeds for the two groups, pick the two rects that would waste the
// most area if covered by a single rectangle, i.e. evidently the worst pair
// to have in the same group.
// Of the remaining, one at a time is chosen to be put in one of the two groups.
// The one chosen is the one with the greatest difference in area expansion
// depending on which group - the rect most strongly attracted to one group
// and repelled from the other.
// If one group gets too full (more would force other group to violate min
// fill requirement) then other group gets the rest.
// These last are the ones that can go in either group most easily.
func d16choosePartition(parVars *d16partitionVarsT, minFill int) {
	var biggestDiff float64
	var group, chosen, betterGroup int

	d16initParVars(parVars, parVars.branchCount, minFill)
	d16pickSeeds(parVars)

	for ((parVars.count[0] + parVars.count[1]) < parVars.total) &&
		(parVars.count[0] < (parVars.total - parVars.minFill)) &&
		(parVars.count[1] < (parVars.total - parVars.minFill)) {
		biggestDiff = -1
		for index := 0; index < parVars.total; index++ {
			if d16notTaken == parVars.partition[index] {
				curRect := &parVars.branchBuf[index].rect
				rect0 := d16combineRect(curRect, &parVars.cover[0])
				rect1 := d16combineRect(curRect, &parVars.cover[1])
				growth0 := d16calcRectVolume(&rect0) - parVars.area[0]
				growth1 := d16calcRectVolume(&rect1) - parVars.area[1]
				diff := growth1 - growth0
				if diff >= 0 {
					group = 0
				} else {
					group = 1
					diff = -diff
				}

				if diff > biggestDiff {
					biggestDiff = diff
					chosen = index
					betterGroup = group
				} else if (diff == biggestDiff) && (parVars.count[group] < parVars.count[betterGroup]) {
					chosen = index
					betterGroup = group
				}
			}
		}
		d16classify(chosen, betterGroup, parVars)
	}

	// If one group too full, put remaining rects in the other
	if (parVars.count[0] + parVars.count[1]) < parVars.total {
		if parVars.count[0] >= parVars.total-parVars.minFill {
			group = 1
		} else {
			group = 0
		}
		for index := 0; index < parVars.total; index++ {
			if d16notTaken == parVars.partition[index] {
				d16classify(index, group, parVars)
			}
		}
	}
}

// Copy branches from the buffer into two nodes according to the partition.
func d16loadNodes(nodeA, nodeB *d16nodeT, parVars *d16partitionVarsT) {
	for index := 0; index < parVars.total; index++ {
		targetNodeIndex := parVars.partition[index]
		targetNodes := []*d16nodeT{nodeA, nodeB}

		// It is assured that d16addBranch here will not cause a node split.
		d16addBranch(&parVars.branchBuf[index], targetNodes[targetNodeIndex], nil)
	}
}

// Initialize a d16partitionVarsT structure.
func d16initParVars(parVars *d16partitionVarsT, maxRects, minFill int) {
	parVars.count[0] = 0
	parVars.count[1] = 0
	parVars.area[0] = 0
	parVars.area[1] = 0
	parVars.total = maxRects
	parVars.minFill = minFill
	for index := 0; index < maxRects; index++ {
		parVars.partition[index] = d16notTaken
	}
}

func d16pickSeeds(parVars *d16partitionVarsT) {
	var seed0, seed1 int
	var worst, waste float64
	var area [d16maxNodes + 1]float64

	for index := 0; index < parVars.total; index++ {
		area[index] = d16calcRectVolume(&parVars.branchBuf[index].rect)
	}

	worst = -parVars.coverSplitArea - 1
	for indexA := 0; indexA < parVars.total-1; indexA++ {
		for indexB := indexA + 1; indexB < parVars.total; indexB++ {
			oneRect := d16combineRect(&parVars.branchBuf[indexA].rect, &parVars.branchBuf[indexB].rect)
			waste = d16calcRectVolume(&oneRect) - area[indexA] - area[indexB]
			if waste > worst {
				worst = waste
				seed0 = indexA
				seed1 = indexB
			}
		}
	}

	d16classify(seed0, 0, parVars)
	d16classify(seed1, 1, parVars)
}

// Put a branch in one of the groups.
func d16classify(index, group int, parVars *d16partitionVarsT) {
	parVars.partition[index] = group

	// Calculate combined rect
	if parVars.count[group] == 0 {
		parVars.cover[group] = parVars.branchBuf[index].rect
	} else {
		parVars.cover[group] = d16combineRect(&parVars.branchBuf[index].rect, &parVars.cover[group])
	}

	// Calculate volume of combined rect
	parVars.area[group] = d16calcRectVolume(&parVars.cover[group])

	parVars.count[group]++
}

// Delete a data rectangle from an index structure.
// Pass in a pointer to a d16rectT, the tid of the record, ptr to ptr to root node.
// Returns 1 if record not found, 0 if success.
// d16removeRect provides for eliminating the root.
func d16removeRect(rect *d16rectT, id interface{}, root **d16nodeT) bool {
	var reInsertList *d16listNodeT

	if !d16removeRectRec(rect, id, *root, &reInsertList) {
		// Found and deleted a data item
		// Reinsert any branches from eliminated nodes
		for reInsertList != nil {
			tempNode := reInsertList.node

			for index := 0; index < tempNode.count; index++ {
				// TODO go over this code. should I use (tempNode->m_level - 1)?
				d16insertRect(&tempNode.branch[index], root, tempNode.level)
			}
			reInsertList = reInsertList.next
		}

		// Check for redundant root (not leaf, 1 child) and eliminate TODO replace
		// if with while? In case there is a whole branch of redundant roots...
		if (*root).count == 1 && (*root).isInternalNode() {
			tempNode := (*root).branch[0].child
			*root = tempNode
		}
		return false
	} else {
		return true
	}
}

// Delete a rectangle from non-root part of an index structure.
// Called by d16removeRect.  Descends tree recursively,
// merges branches on the way back up.
// Returns 1 if record not found, 0 if success.
func d16removeRectRec(rect *d16rectT, id interface{}, node *d16nodeT, listNode **d16listNodeT) bool {
	if node.isInternalNode() { // not a leaf node
		for index := 0; index < node.count; index++ {
			if d16overlap(*rect, node.branch[index].rect) {
				if !d16removeRectRec(rect, id, node.branch[index].child, listNode) {
					if node.branch[index].child.count >= d16minNodes {
						// child removed, just resize parent rect
						node.branch[index].rect = d16nodeCover(node.branch[index].child)
					} else {
						// child removed, not enough entries in node, eliminate node
						d16reInsert(node.branch[index].child, listNode)
						d16disconnectBranch(node, index) // Must return after this call as count has changed
					}
					return false
				}
			}
		}
		return true
	} else { // A leaf node
		for index := 0; index < node.count; index++ {
			if node.branch[index].data == id {
				d16disconnectBranch(node, index) // Must return after this call as count has changed
				return false
			}
		}
		return true
	}
}

// Decide whether two rectangles d16overlap.
func d16overlap(rectA, rectB d16rectT) bool {
	for index := 0; index < d16numDims; index++ {
		if rectA.min[index] > rectB.max[index] ||
			rectB.min[index] > rectA.max[index] {
			return false
		}
	}
	return true
}

// Add a node to the reinsertion list.  All its branches will later
// be reinserted into the index structure.
func d16reInsert(node *d16nodeT, listNode **d16listNodeT) {
	newListNode := &d16listNodeT{}
	newListNode.node = node
	newListNode.next = *listNode
	*listNode = newListNode
}

// d16search in an index tree or subtree for all data retangles that d16overlap the argument rectangle.
func d16search(node *d16nodeT, rect d16rectT, foundCount int, resultCallback func(data interface{}) bool) (int, bool) {
	if node.isInternalNode() {
		// This is an internal node in the tree
		for index := 0; index < node.count; index++ {
			if d16overlap(rect, node.branch[index].rect) {
				var ok bool
				foundCount, ok = d16search(node.branch[index].child, rect, foundCount, resultCallback)
				if !ok {
					// The callback indicated to stop searching
					return foundCount, false
				}
			}
		}
	} else {
		// This is a leaf node
		for index := 0; index < node.count; index++ {
			if d16overlap(rect, node.branch[index].rect) {
				id := node.branch[index].data
				foundCount++
				if !resultCallback(id) {
					return foundCount, false // Don't continue searching
				}

			}
		}
	}
	return foundCount, true // Continue searching
}

func d17fmin(a, b float64) float64 {
	if a < b {
		return a
	}
	return b
}
func d17fmax(a, b float64) float64 {
	if a > b {
		return a
	}
	return b
}

const (
	d17numDims            = 17
	d17maxNodes           = 8
	d17minNodes           = d17maxNodes / 2
	d17useSphericalVolume = true // Better split classification, may be slower on some systems
)

var d17unitSphereVolume = []float64{
	0.000000, 2.000000, 3.141593, // Dimension  0,1,2
	4.188790, 4.934802, 5.263789, // Dimension  3,4,5
	5.167713, 4.724766, 4.058712, // Dimension  6,7,8
	3.298509, 2.550164, 1.884104, // Dimension  9,10,11
	1.335263, 0.910629, 0.599265, // Dimension  12,13,14
	0.381443, 0.235331, 0.140981, // Dimension  15,16,17
	0.082146, 0.046622, 0.025807, // Dimension  18,19,20
}[d17numDims]

type d17RTree struct {
	root *d17nodeT ///< Root of tree
}

/// Minimal bounding rectangle (n-dimensional)
type d17rectT struct {
	min [d17numDims]float64 ///< Min dimensions of bounding box
	max [d17numDims]float64 ///< Max dimensions of bounding box
}

/// May be data or may be another subtree
/// The parents level determines this.
/// If the parents level is 0, then this is data
type d17branchT struct {
	rect  d17rectT    ///< Bounds
	child *d17nodeT   ///< Child node
	data  interface{} ///< Data Id or Ptr
}

/// d17nodeT for each branch level
type d17nodeT struct {
	count  int                     ///< Count
	level  int                     ///< Leaf is zero, others positive
	branch [d17maxNodes]d17branchT ///< Branch
}

func (node *d17nodeT) isInternalNode() bool {
	return (node.level > 0) // Not a leaf, but a internal node
}
func (node *d17nodeT) isLeaf() bool {
	return (node.level == 0) // A leaf, contains data
}

/// A link list of nodes for reinsertion after a delete operation
type d17listNodeT struct {
	next *d17listNodeT ///< Next in list
	node *d17nodeT     ///< Node
}

const d17notTaken = -1 // indicates that position

/// Variables for finding a split partition
type d17partitionVarsT struct {
	partition [d17maxNodes + 1]int
	total     int
	minFill   int
	count     [2]int
	cover     [2]d17rectT
	area      [2]float64

	branchBuf      [d17maxNodes + 1]d17branchT
	branchCount    int
	coverSplit     d17rectT
	coverSplitArea float64
}

func d17New() *d17RTree {
	// We only support machine word size simple data type eg. integer index or object pointer.
	// Since we are storing as union with non data branch
	return &d17RTree{
		root: &d17nodeT{},
	}
}

/// Insert entry
/// \param a_min Min of bounding rect
/// \param a_max Max of bounding rect
/// \param a_dataId Positive Id of data.  Maybe zero, but negative numbers not allowed.
func (tr *d17RTree) Insert(min, max [d17numDims]float64, dataId interface{}) {
	var branch d17branchT
	branch.data = dataId
	for axis := 0; axis < d17numDims; axis++ {
		branch.rect.min[axis] = min[axis]
		branch.rect.max[axis] = max[axis]
	}
	d17insertRect(&branch, &tr.root, 0)
}

/// Remove entry
/// \param a_min Min of bounding rect
/// \param a_max Max of bounding rect
/// \param a_dataId Positive Id of data.  Maybe zero, but negative numbers not allowed.
func (tr *d17RTree) Remove(min, max [d17numDims]float64, dataId interface{}) {
	var rect d17rectT
	for axis := 0; axis < d17numDims; axis++ {
		rect.min[axis] = min[axis]
		rect.max[axis] = max[axis]
	}
	d17removeRect(&rect, dataId, &tr.root)
}

/// Find all within d17search rectangle
/// \param a_min Min of d17search bounding rect
/// \param a_max Max of d17search bounding rect
/// \param a_searchResult d17search result array.  Caller should set grow size. Function will reset, not append to array.
/// \param a_resultCallback Callback function to return result.  Callback should return 'true' to continue searching
/// \param a_context User context to pass as parameter to a_resultCallback
/// \return Returns the number of entries found
func (tr *d17RTree) Search(min, max [d17numDims]float64, resultCallback func(data interface{}) bool) int {
	var rect d17rectT
	for axis := 0; axis < d17numDims; axis++ {
		rect.min[axis] = min[axis]
		rect.max[axis] = max[axis]
	}
	foundCount, _ := d17search(tr.root, rect, 0, resultCallback)
	return foundCount
}

/// Count the data elements in this container.  This is slow as no internal counter is maintained.
func (tr *d17RTree) Count() int {
	var count int
	d17countRec(tr.root, &count)
	return count
}

/// Remove all entries from tree
func (tr *d17RTree) RemoveAll() {
	// Delete all existing nodes
	tr.root = &d17nodeT{}
}

func d17countRec(node *d17nodeT, count *int) {
	if node.isInternalNode() { // not a leaf node
		for index := 0; index < node.count; index++ {
			d17countRec(node.branch[index].child, count)
		}
	} else { // A leaf node
		*count += node.count
	}
}

// Inserts a new data rectangle into the index structure.
// Recursively descends tree, propagates splits back up.
// Returns 0 if node was not split.  Old node updated.
// If node was split, returns 1 and sets the pointer pointed to by
// new_node to point to the new node.  Old node updated to become one of two.
// The level argument specifies the number of steps up from the leaf
// level to insert; e.g. a data rectangle goes in at level = 0.
func d17insertRectRec(branch *d17branchT, node *d17nodeT, newNode **d17nodeT, level int) bool {
	// recurse until we reach the correct level for the new record. data records
	// will always be called with a_level == 0 (leaf)
	if node.level > level {
		// Still above level for insertion, go down tree recursively
		var otherNode *d17nodeT
		//var newBranch d17branchT

		// find the optimal branch for this record
		index := d17pickBranch(&branch.rect, node)

		// recursively insert this record into the picked branch
		childWasSplit := d17insertRectRec(branch, node.branch[index].child, &otherNode, level)

		if !childWasSplit {
			// Child was not split. Merge the bounding box of the new record with the
			// existing bounding box
			node.branch[index].rect = d17combineRect(&branch.rect, &(node.branch[index].rect))
			return false
		} else {
			// Child was split. The old branches are now re-partitioned to two nodes
			// so we have to re-calculate the bounding boxes of each node
			node.branch[index].rect = d17nodeCover(node.branch[index].child)
			var newBranch d17branchT
			newBranch.child = otherNode
			newBranch.rect = d17nodeCover(otherNode)

			// The old node is already a child of a_node. Now add the newly-created
			// node to a_node as well. a_node might be split because of that.
			return d17addBranch(&newBranch, node, newNode)
		}
	} else if node.level == level {
		// We have reached level for insertion. Add rect, split if necessary
		return d17addBranch(branch, node, newNode)
	} else {
		// Should never occur
		return false
	}
}

// Insert a data rectangle into an index structure.
// d17insertRect provides for splitting the root;
// returns 1 if root was split, 0 if it was not.
// The level argument specifies the number of steps up from the leaf
// level to insert; e.g. a data rectangle goes in at level = 0.
// InsertRect2 does the recursion.
//
func d17insertRect(branch *d17branchT, root **d17nodeT, level int) bool {
	var newNode *d17nodeT

	if d17insertRectRec(branch, *root, &newNode, level) { // Root split

		// Grow tree taller and new root
		newRoot := &d17nodeT{}
		newRoot.level = (*root).level + 1

		var newBranch d17branchT

		// add old root node as a child of the new root
		newBranch.rect = d17nodeCover(*root)
		newBranch.child = *root
		d17addBranch(&newBranch, newRoot, nil)

		// add the split node as a child of the new root
		newBranch.rect = d17nodeCover(newNode)
		newBranch.child = newNode
		d17addBranch(&newBranch, newRoot, nil)

		// set the new root as the root node
		*root = newRoot

		return true
	}
	return false
}

// Find the smallest rectangle that includes all rectangles in branches of a node.
func d17nodeCover(node *d17nodeT) d17rectT {
	rect := node.branch[0].rect
	for index := 1; index < node.count; index++ {
		rect = d17combineRect(&rect, &(node.branch[index].rect))
	}
	return rect
}

// Add a branch to a node.  Split the node if necessary.
// Returns 0 if node not split.  Old node updated.
// Returns 1 if node split, sets *new_node to address of new node.
// Old node updated, becomes one of two.
func d17addBranch(branch *d17branchT, node *d17nodeT, newNode **d17nodeT) bool {
	if node.count < d17maxNodes { // Split won't be necessary
		node.branch[node.count] = *branch
		node.count++
		return false
	} else {
		d17splitNode(node, branch, newNode)
		return true
	}
}

// Disconnect a dependent node.
// Caller must return (or stop using iteration index) after this as count has changed
func d17disconnectBranch(node *d17nodeT, index int) {
	// Remove element by swapping with the last element to prevent gaps in array
	node.branch[index] = node.branch[node.count-1]
	node.branch[node.count-1].data = nil
	node.branch[node.count-1].child = nil
	node.count--
}

// Pick a branch.  Pick the one that will need the smallest increase
// in area to accomodate the new rectangle.  This will result in the
// least total area for the covering rectangles in the current node.
// In case of a tie, pick the one which was smaller before, to get
// the best resolution when searching.
func d17pickBranch(rect *d17rectT, node *d17nodeT) int {
	var firstTime bool = true
	var increase float64
	var bestIncr float64 = -1
	var area float64
	var bestArea float64
	var best int
	var tempRect d17rectT

	for index := 0; index < node.count; index++ {
		curRect := &node.branch[index].rect
		area = d17calcRectVolume(curRect)
		tempRect = d17combineRect(rect, curRect)
		increase = d17calcRectVolume(&tempRect) - area
		if (increase < bestIncr) || firstTime {
			best = index
			bestArea = area
			bestIncr = increase
			firstTime = false
		} else if (increase == bestIncr) && (area < bestArea) {
			best = index
			bestArea = area
			bestIncr = increase
		}
	}
	return best
}

// Combine two rectangles into larger one containing both
func d17combineRect(rectA, rectB *d17rectT) d17rectT {
	var newRect d17rectT

	for index := 0; index < d17numDims; index++ {
		newRect.min[index] = d17fmin(rectA.min[index], rectB.min[index])
		newRect.max[index] = d17fmax(rectA.max[index], rectB.max[index])
	}

	return newRect
}

// Split a node.
// Divides the nodes branches and the extra one between two nodes.
// Old node is one of the new ones, and one really new one is created.
// Tries more than one method for choosing a partition, uses best result.
func d17splitNode(node *d17nodeT, branch *d17branchT, newNode **d17nodeT) {
	// Could just use local here, but member or external is faster since it is reused
	var localVars d17partitionVarsT
	parVars := &localVars

	// Load all the branches into a buffer, initialize old node
	d17getBranches(node, branch, parVars)

	// Find partition
	d17choosePartition(parVars, d17minNodes)

	// Create a new node to hold (about) half of the branches
	*newNode = &d17nodeT{}
	(*newNode).level = node.level

	// Put branches from buffer into 2 nodes according to the chosen partition
	node.count = 0
	d17loadNodes(node, *newNode, parVars)
}

// Calculate the n-dimensional volume of a rectangle
func d17rectVolume(rect *d17rectT) float64 {
	var volume float64 = 1
	for index := 0; index < d17numDims; index++ {
		volume *= rect.max[index] - rect.min[index]
	}
	return volume
}

// The exact volume of the bounding sphere for the given d17rectT
func d17rectSphericalVolume(rect *d17rectT) float64 {
	var sumOfSquares float64 = 0
	var radius float64

	for index := 0; index < d17numDims; index++ {
		halfExtent := (rect.max[index] - rect.min[index]) * 0.5
		sumOfSquares += halfExtent * halfExtent
	}

	radius = math.Sqrt(sumOfSquares)

	// Pow maybe slow, so test for common dims just use x*x, x*x*x.
	if d17numDims == 5 {
		return (radius * radius * radius * radius * radius * d17unitSphereVolume)
	} else if d17numDims == 4 {
		return (radius * radius * radius * radius * d17unitSphereVolume)
	} else if d17numDims == 3 {
		return (radius * radius * radius * d17unitSphereVolume)
	} else if d17numDims == 2 {
		return (radius * radius * d17unitSphereVolume)
	} else {
		return (math.Pow(radius, d17numDims) * d17unitSphereVolume)
	}
}

// Use one of the methods to calculate retangle volume
func d17calcRectVolume(rect *d17rectT) float64 {
	if d17useSphericalVolume {
		return d17rectSphericalVolume(rect) // Slower but helps certain merge cases
	} else { // RTREE_USE_SPHERICAL_VOLUME
		return d17rectVolume(rect) // Faster but can cause poor merges
	} // RTREE_USE_SPHERICAL_VOLUME
}

// Load branch buffer with branches from full node plus the extra branch.
func d17getBranches(node *d17nodeT, branch *d17branchT, parVars *d17partitionVarsT) {
	// Load the branch buffer
	for index := 0; index < d17maxNodes; index++ {
		parVars.branchBuf[index] = node.branch[index]
	}
	parVars.branchBuf[d17maxNodes] = *branch
	parVars.branchCount = d17maxNodes + 1

	// Calculate rect containing all in the set
	parVars.coverSplit = parVars.branchBuf[0].rect
	for index := 1; index < d17maxNodes+1; index++ {
		parVars.coverSplit = d17combineRect(&parVars.coverSplit, &parVars.branchBuf[index].rect)
	}
	parVars.coverSplitArea = d17calcRectVolume(&parVars.coverSplit)
}

// Method #0 for choosing a partition:
// As the seeds for the two groups, pick the two rects that would waste the
// most area if covered by a single rectangle, i.e. evidently the worst pair
// to have in the same group.
// Of the remaining, one at a time is chosen to be put in one of the two groups.
// The one chosen is the one with the greatest difference in area expansion
// depending on which group - the rect most strongly attracted to one group
// and repelled from the other.
// If one group gets too full (more would force other group to violate min
// fill requirement) then other group gets the rest.
// These last are the ones that can go in either group most easily.
func d17choosePartition(parVars *d17partitionVarsT, minFill int) {
	var biggestDiff float64
	var group, chosen, betterGroup int

	d17initParVars(parVars, parVars.branchCount, minFill)
	d17pickSeeds(parVars)

	for ((parVars.count[0] + parVars.count[1]) < parVars.total) &&
		(parVars.count[0] < (parVars.total - parVars.minFill)) &&
		(parVars.count[1] < (parVars.total - parVars.minFill)) {
		biggestDiff = -1
		for index := 0; index < parVars.total; index++ {
			if d17notTaken == parVars.partition[index] {
				curRect := &parVars.branchBuf[index].rect
				rect0 := d17combineRect(curRect, &parVars.cover[0])
				rect1 := d17combineRect(curRect, &parVars.cover[1])
				growth0 := d17calcRectVolume(&rect0) - parVars.area[0]
				growth1 := d17calcRectVolume(&rect1) - parVars.area[1]
				diff := growth1 - growth0
				if diff >= 0 {
					group = 0
				} else {
					group = 1
					diff = -diff
				}

				if diff > biggestDiff {
					biggestDiff = diff
					chosen = index
					betterGroup = group
				} else if (diff == biggestDiff) && (parVars.count[group] < parVars.count[betterGroup]) {
					chosen = index
					betterGroup = group
				}
			}
		}
		d17classify(chosen, betterGroup, parVars)
	}

	// If one group too full, put remaining rects in the other
	if (parVars.count[0] + parVars.count[1]) < parVars.total {
		if parVars.count[0] >= parVars.total-parVars.minFill {
			group = 1
		} else {
			group = 0
		}
		for index := 0; index < parVars.total; index++ {
			if d17notTaken == parVars.partition[index] {
				d17classify(index, group, parVars)
			}
		}
	}
}

// Copy branches from the buffer into two nodes according to the partition.
func d17loadNodes(nodeA, nodeB *d17nodeT, parVars *d17partitionVarsT) {
	for index := 0; index < parVars.total; index++ {
		targetNodeIndex := parVars.partition[index]
		targetNodes := []*d17nodeT{nodeA, nodeB}

		// It is assured that d17addBranch here will not cause a node split.
		d17addBranch(&parVars.branchBuf[index], targetNodes[targetNodeIndex], nil)
	}
}

// Initialize a d17partitionVarsT structure.
func d17initParVars(parVars *d17partitionVarsT, maxRects, minFill int) {
	parVars.count[0] = 0
	parVars.count[1] = 0
	parVars.area[0] = 0
	parVars.area[1] = 0
	parVars.total = maxRects
	parVars.minFill = minFill
	for index := 0; index < maxRects; index++ {
		parVars.partition[index] = d17notTaken
	}
}

func d17pickSeeds(parVars *d17partitionVarsT) {
	var seed0, seed1 int
	var worst, waste float64
	var area [d17maxNodes + 1]float64

	for index := 0; index < parVars.total; index++ {
		area[index] = d17calcRectVolume(&parVars.branchBuf[index].rect)
	}

	worst = -parVars.coverSplitArea - 1
	for indexA := 0; indexA < parVars.total-1; indexA++ {
		for indexB := indexA + 1; indexB < parVars.total; indexB++ {
			oneRect := d17combineRect(&parVars.branchBuf[indexA].rect, &parVars.branchBuf[indexB].rect)
			waste = d17calcRectVolume(&oneRect) - area[indexA] - area[indexB]
			if waste > worst {
				worst = waste
				seed0 = indexA
				seed1 = indexB
			}
		}
	}

	d17classify(seed0, 0, parVars)
	d17classify(seed1, 1, parVars)
}

// Put a branch in one of the groups.
func d17classify(index, group int, parVars *d17partitionVarsT) {
	parVars.partition[index] = group

	// Calculate combined rect
	if parVars.count[group] == 0 {
		parVars.cover[group] = parVars.branchBuf[index].rect
	} else {
		parVars.cover[group] = d17combineRect(&parVars.branchBuf[index].rect, &parVars.cover[group])
	}

	// Calculate volume of combined rect
	parVars.area[group] = d17calcRectVolume(&parVars.cover[group])

	parVars.count[group]++
}

// Delete a data rectangle from an index structure.
// Pass in a pointer to a d17rectT, the tid of the record, ptr to ptr to root node.
// Returns 1 if record not found, 0 if success.
// d17removeRect provides for eliminating the root.
func d17removeRect(rect *d17rectT, id interface{}, root **d17nodeT) bool {
	var reInsertList *d17listNodeT

	if !d17removeRectRec(rect, id, *root, &reInsertList) {
		// Found and deleted a data item
		// Reinsert any branches from eliminated nodes
		for reInsertList != nil {
			tempNode := reInsertList.node

			for index := 0; index < tempNode.count; index++ {
				// TODO go over this code. should I use (tempNode->m_level - 1)?
				d17insertRect(&tempNode.branch[index], root, tempNode.level)
			}
			reInsertList = reInsertList.next
		}

		// Check for redundant root (not leaf, 1 child) and eliminate TODO replace
		// if with while? In case there is a whole branch of redundant roots...
		if (*root).count == 1 && (*root).isInternalNode() {
			tempNode := (*root).branch[0].child
			*root = tempNode
		}
		return false
	} else {
		return true
	}
}

// Delete a rectangle from non-root part of an index structure.
// Called by d17removeRect.  Descends tree recursively,
// merges branches on the way back up.
// Returns 1 if record not found, 0 if success.
func d17removeRectRec(rect *d17rectT, id interface{}, node *d17nodeT, listNode **d17listNodeT) bool {
	if node.isInternalNode() { // not a leaf node
		for index := 0; index < node.count; index++ {
			if d17overlap(*rect, node.branch[index].rect) {
				if !d17removeRectRec(rect, id, node.branch[index].child, listNode) {
					if node.branch[index].child.count >= d17minNodes {
						// child removed, just resize parent rect
						node.branch[index].rect = d17nodeCover(node.branch[index].child)
					} else {
						// child removed, not enough entries in node, eliminate node
						d17reInsert(node.branch[index].child, listNode)
						d17disconnectBranch(node, index) // Must return after this call as count has changed
					}
					return false
				}
			}
		}
		return true
	} else { // A leaf node
		for index := 0; index < node.count; index++ {
			if node.branch[index].data == id {
				d17disconnectBranch(node, index) // Must return after this call as count has changed
				return false
			}
		}
		return true
	}
}

// Decide whether two rectangles d17overlap.
func d17overlap(rectA, rectB d17rectT) bool {
	for index := 0; index < d17numDims; index++ {
		if rectA.min[index] > rectB.max[index] ||
			rectB.min[index] > rectA.max[index] {
			return false
		}
	}
	return true
}

// Add a node to the reinsertion list.  All its branches will later
// be reinserted into the index structure.
func d17reInsert(node *d17nodeT, listNode **d17listNodeT) {
	newListNode := &d17listNodeT{}
	newListNode.node = node
	newListNode.next = *listNode
	*listNode = newListNode
}

// d17search in an index tree or subtree for all data retangles that d17overlap the argument rectangle.
func d17search(node *d17nodeT, rect d17rectT, foundCount int, resultCallback func(data interface{}) bool) (int, bool) {
	if node.isInternalNode() {
		// This is an internal node in the tree
		for index := 0; index < node.count; index++ {
			if d17overlap(rect, node.branch[index].rect) {
				var ok bool
				foundCount, ok = d17search(node.branch[index].child, rect, foundCount, resultCallback)
				if !ok {
					// The callback indicated to stop searching
					return foundCount, false
				}
			}
		}
	} else {
		// This is a leaf node
		for index := 0; index < node.count; index++ {
			if d17overlap(rect, node.branch[index].rect) {
				id := node.branch[index].data
				foundCount++
				if !resultCallback(id) {
					return foundCount, false // Don't continue searching
				}

			}
		}
	}
	return foundCount, true // Continue searching
}

func d18fmin(a, b float64) float64 {
	if a < b {
		return a
	}
	return b
}
func d18fmax(a, b float64) float64 {
	if a > b {
		return a
	}
	return b
}

const (
	d18numDims            = 18
	d18maxNodes           = 8
	d18minNodes           = d18maxNodes / 2
	d18useSphericalVolume = true // Better split classification, may be slower on some systems
)

var d18unitSphereVolume = []float64{
	0.000000, 2.000000, 3.141593, // Dimension  0,1,2
	4.188790, 4.934802, 5.263789, // Dimension  3,4,5
	5.167713, 4.724766, 4.058712, // Dimension  6,7,8
	3.298509, 2.550164, 1.884104, // Dimension  9,10,11
	1.335263, 0.910629, 0.599265, // Dimension  12,13,14
	0.381443, 0.235331, 0.140981, // Dimension  15,16,17
	0.082146, 0.046622, 0.025807, // Dimension  18,19,20
}[d18numDims]

type d18RTree struct {
	root *d18nodeT ///< Root of tree
}

/// Minimal bounding rectangle (n-dimensional)
type d18rectT struct {
	min [d18numDims]float64 ///< Min dimensions of bounding box
	max [d18numDims]float64 ///< Max dimensions of bounding box
}

/// May be data or may be another subtree
/// The parents level determines this.
/// If the parents level is 0, then this is data
type d18branchT struct {
	rect  d18rectT    ///< Bounds
	child *d18nodeT   ///< Child node
	data  interface{} ///< Data Id or Ptr
}

/// d18nodeT for each branch level
type d18nodeT struct {
	count  int                     ///< Count
	level  int                     ///< Leaf is zero, others positive
	branch [d18maxNodes]d18branchT ///< Branch
}

func (node *d18nodeT) isInternalNode() bool {
	return (node.level > 0) // Not a leaf, but a internal node
}
func (node *d18nodeT) isLeaf() bool {
	return (node.level == 0) // A leaf, contains data
}

/// A link list of nodes for reinsertion after a delete operation
type d18listNodeT struct {
	next *d18listNodeT ///< Next in list
	node *d18nodeT     ///< Node
}

const d18notTaken = -1 // indicates that position

/// Variables for finding a split partition
type d18partitionVarsT struct {
	partition [d18maxNodes + 1]int
	total     int
	minFill   int
	count     [2]int
	cover     [2]d18rectT
	area      [2]float64

	branchBuf      [d18maxNodes + 1]d18branchT
	branchCount    int
	coverSplit     d18rectT
	coverSplitArea float64
}

func d18New() *d18RTree {
	// We only support machine word size simple data type eg. integer index or object pointer.
	// Since we are storing as union with non data branch
	return &d18RTree{
		root: &d18nodeT{},
	}
}

/// Insert entry
/// \param a_min Min of bounding rect
/// \param a_max Max of bounding rect
/// \param a_dataId Positive Id of data.  Maybe zero, but negative numbers not allowed.
func (tr *d18RTree) Insert(min, max [d18numDims]float64, dataId interface{}) {
	var branch d18branchT
	branch.data = dataId
	for axis := 0; axis < d18numDims; axis++ {
		branch.rect.min[axis] = min[axis]
		branch.rect.max[axis] = max[axis]
	}
	d18insertRect(&branch, &tr.root, 0)
}

/// Remove entry
/// \param a_min Min of bounding rect
/// \param a_max Max of bounding rect
/// \param a_dataId Positive Id of data.  Maybe zero, but negative numbers not allowed.
func (tr *d18RTree) Remove(min, max [d18numDims]float64, dataId interface{}) {
	var rect d18rectT
	for axis := 0; axis < d18numDims; axis++ {
		rect.min[axis] = min[axis]
		rect.max[axis] = max[axis]
	}
	d18removeRect(&rect, dataId, &tr.root)
}

/// Find all within d18search rectangle
/// \param a_min Min of d18search bounding rect
/// \param a_max Max of d18search bounding rect
/// \param a_searchResult d18search result array.  Caller should set grow size. Function will reset, not append to array.
/// \param a_resultCallback Callback function to return result.  Callback should return 'true' to continue searching
/// \param a_context User context to pass as parameter to a_resultCallback
/// \return Returns the number of entries found
func (tr *d18RTree) Search(min, max [d18numDims]float64, resultCallback func(data interface{}) bool) int {
	var rect d18rectT
	for axis := 0; axis < d18numDims; axis++ {
		rect.min[axis] = min[axis]
		rect.max[axis] = max[axis]
	}
	foundCount, _ := d18search(tr.root, rect, 0, resultCallback)
	return foundCount
}

/// Count the data elements in this container.  This is slow as no internal counter is maintained.
func (tr *d18RTree) Count() int {
	var count int
	d18countRec(tr.root, &count)
	return count
}

/// Remove all entries from tree
func (tr *d18RTree) RemoveAll() {
	// Delete all existing nodes
	tr.root = &d18nodeT{}
}

func d18countRec(node *d18nodeT, count *int) {
	if node.isInternalNode() { // not a leaf node
		for index := 0; index < node.count; index++ {
			d18countRec(node.branch[index].child, count)
		}
	} else { // A leaf node
		*count += node.count
	}
}

// Inserts a new data rectangle into the index structure.
// Recursively descends tree, propagates splits back up.
// Returns 0 if node was not split.  Old node updated.
// If node was split, returns 1 and sets the pointer pointed to by
// new_node to point to the new node.  Old node updated to become one of two.
// The level argument specifies the number of steps up from the leaf
// level to insert; e.g. a data rectangle goes in at level = 0.
func d18insertRectRec(branch *d18branchT, node *d18nodeT, newNode **d18nodeT, level int) bool {
	// recurse until we reach the correct level for the new record. data records
	// will always be called with a_level == 0 (leaf)
	if node.level > level {
		// Still above level for insertion, go down tree recursively
		var otherNode *d18nodeT
		//var newBranch d18branchT

		// find the optimal branch for this record
		index := d18pickBranch(&branch.rect, node)

		// recursively insert this record into the picked branch
		childWasSplit := d18insertRectRec(branch, node.branch[index].child, &otherNode, level)

		if !childWasSplit {
			// Child was not split. Merge the bounding box of the new record with the
			// existing bounding box
			node.branch[index].rect = d18combineRect(&branch.rect, &(node.branch[index].rect))
			return false
		} else {
			// Child was split. The old branches are now re-partitioned to two nodes
			// so we have to re-calculate the bounding boxes of each node
			node.branch[index].rect = d18nodeCover(node.branch[index].child)
			var newBranch d18branchT
			newBranch.child = otherNode
			newBranch.rect = d18nodeCover(otherNode)

			// The old node is already a child of a_node. Now add the newly-created
			// node to a_node as well. a_node might be split because of that.
			return d18addBranch(&newBranch, node, newNode)
		}
	} else if node.level == level {
		// We have reached level for insertion. Add rect, split if necessary
		return d18addBranch(branch, node, newNode)
	} else {
		// Should never occur
		return false
	}
}

// Insert a data rectangle into an index structure.
// d18insertRect provides for splitting the root;
// returns 1 if root was split, 0 if it was not.
// The level argument specifies the number of steps up from the leaf
// level to insert; e.g. a data rectangle goes in at level = 0.
// InsertRect2 does the recursion.
//
func d18insertRect(branch *d18branchT, root **d18nodeT, level int) bool {
	var newNode *d18nodeT

	if d18insertRectRec(branch, *root, &newNode, level) { // Root split

		// Grow tree taller and new root
		newRoot := &d18nodeT{}
		newRoot.level = (*root).level + 1

		var newBranch d18branchT

		// add old root node as a child of the new root
		newBranch.rect = d18nodeCover(*root)
		newBranch.child = *root
		d18addBranch(&newBranch, newRoot, nil)

		// add the split node as a child of the new root
		newBranch.rect = d18nodeCover(newNode)
		newBranch.child = newNode
		d18addBranch(&newBranch, newRoot, nil)

		// set the new root as the root node
		*root = newRoot

		return true
	}
	return false
}

// Find the smallest rectangle that includes all rectangles in branches of a node.
func d18nodeCover(node *d18nodeT) d18rectT {
	rect := node.branch[0].rect
	for index := 1; index < node.count; index++ {
		rect = d18combineRect(&rect, &(node.branch[index].rect))
	}
	return rect
}

// Add a branch to a node.  Split the node if necessary.
// Returns 0 if node not split.  Old node updated.
// Returns 1 if node split, sets *new_node to address of new node.
// Old node updated, becomes one of two.
func d18addBranch(branch *d18branchT, node *d18nodeT, newNode **d18nodeT) bool {
	if node.count < d18maxNodes { // Split won't be necessary
		node.branch[node.count] = *branch
		node.count++
		return false
	} else {
		d18splitNode(node, branch, newNode)
		return true
	}
}

// Disconnect a dependent node.
// Caller must return (or stop using iteration index) after this as count has changed
func d18disconnectBranch(node *d18nodeT, index int) {
	// Remove element by swapping with the last element to prevent gaps in array
	node.branch[index] = node.branch[node.count-1]
	node.branch[node.count-1].data = nil
	node.branch[node.count-1].child = nil
	node.count--
}

// Pick a branch.  Pick the one that will need the smallest increase
// in area to accomodate the new rectangle.  This will result in the
// least total area for the covering rectangles in the current node.
// In case of a tie, pick the one which was smaller before, to get
// the best resolution when searching.
func d18pickBranch(rect *d18rectT, node *d18nodeT) int {
	var firstTime bool = true
	var increase float64
	var bestIncr float64 = -1
	var area float64
	var bestArea float64
	var best int
	var tempRect d18rectT

	for index := 0; index < node.count; index++ {
		curRect := &node.branch[index].rect
		area = d18calcRectVolume(curRect)
		tempRect = d18combineRect(rect, curRect)
		increase = d18calcRectVolume(&tempRect) - area
		if (increase < bestIncr) || firstTime {
			best = index
			bestArea = area
			bestIncr = increase
			firstTime = false
		} else if (increase == bestIncr) && (area < bestArea) {
			best = index
			bestArea = area
			bestIncr = increase
		}
	}
	return best
}

// Combine two rectangles into larger one containing both
func d18combineRect(rectA, rectB *d18rectT) d18rectT {
	var newRect d18rectT

	for index := 0; index < d18numDims; index++ {
		newRect.min[index] = d18fmin(rectA.min[index], rectB.min[index])
		newRect.max[index] = d18fmax(rectA.max[index], rectB.max[index])
	}

	return newRect
}

// Split a node.
// Divides the nodes branches and the extra one between two nodes.
// Old node is one of the new ones, and one really new one is created.
// Tries more than one method for choosing a partition, uses best result.
func d18splitNode(node *d18nodeT, branch *d18branchT, newNode **d18nodeT) {
	// Could just use local here, but member or external is faster since it is reused
	var localVars d18partitionVarsT
	parVars := &localVars

	// Load all the branches into a buffer, initialize old node
	d18getBranches(node, branch, parVars)

	// Find partition
	d18choosePartition(parVars, d18minNodes)

	// Create a new node to hold (about) half of the branches
	*newNode = &d18nodeT{}
	(*newNode).level = node.level

	// Put branches from buffer into 2 nodes according to the chosen partition
	node.count = 0
	d18loadNodes(node, *newNode, parVars)
}

// Calculate the n-dimensional volume of a rectangle
func d18rectVolume(rect *d18rectT) float64 {
	var volume float64 = 1
	for index := 0; index < d18numDims; index++ {
		volume *= rect.max[index] - rect.min[index]
	}
	return volume
}

// The exact volume of the bounding sphere for the given d18rectT
func d18rectSphericalVolume(rect *d18rectT) float64 {
	var sumOfSquares float64 = 0
	var radius float64

	for index := 0; index < d18numDims; index++ {
		halfExtent := (rect.max[index] - rect.min[index]) * 0.5
		sumOfSquares += halfExtent * halfExtent
	}

	radius = math.Sqrt(sumOfSquares)

	// Pow maybe slow, so test for common dims just use x*x, x*x*x.
	if d18numDims == 5 {
		return (radius * radius * radius * radius * radius * d18unitSphereVolume)
	} else if d18numDims == 4 {
		return (radius * radius * radius * radius * d18unitSphereVolume)
	} else if d18numDims == 3 {
		return (radius * radius * radius * d18unitSphereVolume)
	} else if d18numDims == 2 {
		return (radius * radius * d18unitSphereVolume)
	} else {
		return (math.Pow(radius, d18numDims) * d18unitSphereVolume)
	}
}

// Use one of the methods to calculate retangle volume
func d18calcRectVolume(rect *d18rectT) float64 {
	if d18useSphericalVolume {
		return d18rectSphericalVolume(rect) // Slower but helps certain merge cases
	} else { // RTREE_USE_SPHERICAL_VOLUME
		return d18rectVolume(rect) // Faster but can cause poor merges
	} // RTREE_USE_SPHERICAL_VOLUME
}

// Load branch buffer with branches from full node plus the extra branch.
func d18getBranches(node *d18nodeT, branch *d18branchT, parVars *d18partitionVarsT) {
	// Load the branch buffer
	for index := 0; index < d18maxNodes; index++ {
		parVars.branchBuf[index] = node.branch[index]
	}
	parVars.branchBuf[d18maxNodes] = *branch
	parVars.branchCount = d18maxNodes + 1

	// Calculate rect containing all in the set
	parVars.coverSplit = parVars.branchBuf[0].rect
	for index := 1; index < d18maxNodes+1; index++ {
		parVars.coverSplit = d18combineRect(&parVars.coverSplit, &parVars.branchBuf[index].rect)
	}
	parVars.coverSplitArea = d18calcRectVolume(&parVars.coverSplit)
}

// Method #0 for choosing a partition:
// As the seeds for the two groups, pick the two rects that would waste the
// most area if covered by a single rectangle, i.e. evidently the worst pair
// to have in the same group.
// Of the remaining, one at a time is chosen to be put in one of the two groups.
// The one chosen is the one with the greatest difference in area expansion
// depending on which group - the rect most strongly attracted to one group
// and repelled from the other.
// If one group gets too full (more would force other group to violate min
// fill requirement) then other group gets the rest.
// These last are the ones that can go in either group most easily.
func d18choosePartition(parVars *d18partitionVarsT, minFill int) {
	var biggestDiff float64
	var group, chosen, betterGroup int

	d18initParVars(parVars, parVars.branchCount, minFill)
	d18pickSeeds(parVars)

	for ((parVars.count[0] + parVars.count[1]) < parVars.total) &&
		(parVars.count[0] < (parVars.total - parVars.minFill)) &&
		(parVars.count[1] < (parVars.total - parVars.minFill)) {
		biggestDiff = -1
		for index := 0; index < parVars.total; index++ {
			if d18notTaken == parVars.partition[index] {
				curRect := &parVars.branchBuf[index].rect
				rect0 := d18combineRect(curRect, &parVars.cover[0])
				rect1 := d18combineRect(curRect, &parVars.cover[1])
				growth0 := d18calcRectVolume(&rect0) - parVars.area[0]
				growth1 := d18calcRectVolume(&rect1) - parVars.area[1]
				diff := growth1 - growth0
				if diff >= 0 {
					group = 0
				} else {
					group = 1
					diff = -diff
				}

				if diff > biggestDiff {
					biggestDiff = diff
					chosen = index
					betterGroup = group
				} else if (diff == biggestDiff) && (parVars.count[group] < parVars.count[betterGroup]) {
					chosen = index
					betterGroup = group
				}
			}
		}
		d18classify(chosen, betterGroup, parVars)
	}

	// If one group too full, put remaining rects in the other
	if (parVars.count[0] + parVars.count[1]) < parVars.total {
		if parVars.count[0] >= parVars.total-parVars.minFill {
			group = 1
		} else {
			group = 0
		}
		for index := 0; index < parVars.total; index++ {
			if d18notTaken == parVars.partition[index] {
				d18classify(index, group, parVars)
			}
		}
	}
}

// Copy branches from the buffer into two nodes according to the partition.
func d18loadNodes(nodeA, nodeB *d18nodeT, parVars *d18partitionVarsT) {
	for index := 0; index < parVars.total; index++ {
		targetNodeIndex := parVars.partition[index]
		targetNodes := []*d18nodeT{nodeA, nodeB}

		// It is assured that d18addBranch here will not cause a node split.
		d18addBranch(&parVars.branchBuf[index], targetNodes[targetNodeIndex], nil)
	}
}

// Initialize a d18partitionVarsT structure.
func d18initParVars(parVars *d18partitionVarsT, maxRects, minFill int) {
	parVars.count[0] = 0
	parVars.count[1] = 0
	parVars.area[0] = 0
	parVars.area[1] = 0
	parVars.total = maxRects
	parVars.minFill = minFill
	for index := 0; index < maxRects; index++ {
		parVars.partition[index] = d18notTaken
	}
}

func d18pickSeeds(parVars *d18partitionVarsT) {
	var seed0, seed1 int
	var worst, waste float64
	var area [d18maxNodes + 1]float64

	for index := 0; index < parVars.total; index++ {
		area[index] = d18calcRectVolume(&parVars.branchBuf[index].rect)
	}

	worst = -parVars.coverSplitArea - 1
	for indexA := 0; indexA < parVars.total-1; indexA++ {
		for indexB := indexA + 1; indexB < parVars.total; indexB++ {
			oneRect := d18combineRect(&parVars.branchBuf[indexA].rect, &parVars.branchBuf[indexB].rect)
			waste = d18calcRectVolume(&oneRect) - area[indexA] - area[indexB]
			if waste > worst {
				worst = waste
				seed0 = indexA
				seed1 = indexB
			}
		}
	}

	d18classify(seed0, 0, parVars)
	d18classify(seed1, 1, parVars)
}

// Put a branch in one of the groups.
func d18classify(index, group int, parVars *d18partitionVarsT) {
	parVars.partition[index] = group

	// Calculate combined rect
	if parVars.count[group] == 0 {
		parVars.cover[group] = parVars.branchBuf[index].rect
	} else {
		parVars.cover[group] = d18combineRect(&parVars.branchBuf[index].rect, &parVars.cover[group])
	}

	// Calculate volume of combined rect
	parVars.area[group] = d18calcRectVolume(&parVars.cover[group])

	parVars.count[group]++
}

// Delete a data rectangle from an index structure.
// Pass in a pointer to a d18rectT, the tid of the record, ptr to ptr to root node.
// Returns 1 if record not found, 0 if success.
// d18removeRect provides for eliminating the root.
func d18removeRect(rect *d18rectT, id interface{}, root **d18nodeT) bool {
	var reInsertList *d18listNodeT

	if !d18removeRectRec(rect, id, *root, &reInsertList) {
		// Found and deleted a data item
		// Reinsert any branches from eliminated nodes
		for reInsertList != nil {
			tempNode := reInsertList.node

			for index := 0; index < tempNode.count; index++ {
				// TODO go over this code. should I use (tempNode->m_level - 1)?
				d18insertRect(&tempNode.branch[index], root, tempNode.level)
			}
			reInsertList = reInsertList.next
		}

		// Check for redundant root (not leaf, 1 child) and eliminate TODO replace
		// if with while? In case there is a whole branch of redundant roots...
		if (*root).count == 1 && (*root).isInternalNode() {
			tempNode := (*root).branch[0].child
			*root = tempNode
		}
		return false
	} else {
		return true
	}
}

// Delete a rectangle from non-root part of an index structure.
// Called by d18removeRect.  Descends tree recursively,
// merges branches on the way back up.
// Returns 1 if record not found, 0 if success.
func d18removeRectRec(rect *d18rectT, id interface{}, node *d18nodeT, listNode **d18listNodeT) bool {
	if node.isInternalNode() { // not a leaf node
		for index := 0; index < node.count; index++ {
			if d18overlap(*rect, node.branch[index].rect) {
				if !d18removeRectRec(rect, id, node.branch[index].child, listNode) {
					if node.branch[index].child.count >= d18minNodes {
						// child removed, just resize parent rect
						node.branch[index].rect = d18nodeCover(node.branch[index].child)
					} else {
						// child removed, not enough entries in node, eliminate node
						d18reInsert(node.branch[index].child, listNode)
						d18disconnectBranch(node, index) // Must return after this call as count has changed
					}
					return false
				}
			}
		}
		return true
	} else { // A leaf node
		for index := 0; index < node.count; index++ {
			if node.branch[index].data == id {
				d18disconnectBranch(node, index) // Must return after this call as count has changed
				return false
			}
		}
		return true
	}
}

// Decide whether two rectangles d18overlap.
func d18overlap(rectA, rectB d18rectT) bool {
	for index := 0; index < d18numDims; index++ {
		if rectA.min[index] > rectB.max[index] ||
			rectB.min[index] > rectA.max[index] {
			return false
		}
	}
	return true
}

// Add a node to the reinsertion list.  All its branches will later
// be reinserted into the index structure.
func d18reInsert(node *d18nodeT, listNode **d18listNodeT) {
	newListNode := &d18listNodeT{}
	newListNode.node = node
	newListNode.next = *listNode
	*listNode = newListNode
}

// d18search in an index tree or subtree for all data retangles that d18overlap the argument rectangle.
func d18search(node *d18nodeT, rect d18rectT, foundCount int, resultCallback func(data interface{}) bool) (int, bool) {
	if node.isInternalNode() {
		// This is an internal node in the tree
		for index := 0; index < node.count; index++ {
			if d18overlap(rect, node.branch[index].rect) {
				var ok bool
				foundCount, ok = d18search(node.branch[index].child, rect, foundCount, resultCallback)
				if !ok {
					// The callback indicated to stop searching
					return foundCount, false
				}
			}
		}
	} else {
		// This is a leaf node
		for index := 0; index < node.count; index++ {
			if d18overlap(rect, node.branch[index].rect) {
				id := node.branch[index].data
				foundCount++
				if !resultCallback(id) {
					return foundCount, false // Don't continue searching
				}

			}
		}
	}
	return foundCount, true // Continue searching
}

func d19fmin(a, b float64) float64 {
	if a < b {
		return a
	}
	return b
}
func d19fmax(a, b float64) float64 {
	if a > b {
		return a
	}
	return b
}

const (
	d19numDims            = 19
	d19maxNodes           = 8
	d19minNodes           = d19maxNodes / 2
	d19useSphericalVolume = true // Better split classification, may be slower on some systems
)

var d19unitSphereVolume = []float64{
	0.000000, 2.000000, 3.141593, // Dimension  0,1,2
	4.188790, 4.934802, 5.263789, // Dimension  3,4,5
	5.167713, 4.724766, 4.058712, // Dimension  6,7,8
	3.298509, 2.550164, 1.884104, // Dimension  9,10,11
	1.335263, 0.910629, 0.599265, // Dimension  12,13,14
	0.381443, 0.235331, 0.140981, // Dimension  15,16,17
	0.082146, 0.046622, 0.025807, // Dimension  18,19,20
}[d19numDims]

type d19RTree struct {
	root *d19nodeT ///< Root of tree
}

/// Minimal bounding rectangle (n-dimensional)
type d19rectT struct {
	min [d19numDims]float64 ///< Min dimensions of bounding box
	max [d19numDims]float64 ///< Max dimensions of bounding box
}

/// May be data or may be another subtree
/// The parents level determines this.
/// If the parents level is 0, then this is data
type d19branchT struct {
	rect  d19rectT    ///< Bounds
	child *d19nodeT   ///< Child node
	data  interface{} ///< Data Id or Ptr
}

/// d19nodeT for each branch level
type d19nodeT struct {
	count  int                     ///< Count
	level  int                     ///< Leaf is zero, others positive
	branch [d19maxNodes]d19branchT ///< Branch
}

func (node *d19nodeT) isInternalNode() bool {
	return (node.level > 0) // Not a leaf, but a internal node
}
func (node *d19nodeT) isLeaf() bool {
	return (node.level == 0) // A leaf, contains data
}

/// A link list of nodes for reinsertion after a delete operation
type d19listNodeT struct {
	next *d19listNodeT ///< Next in list
	node *d19nodeT     ///< Node
}

const d19notTaken = -1 // indicates that position

/// Variables for finding a split partition
type d19partitionVarsT struct {
	partition [d19maxNodes + 1]int
	total     int
	minFill   int
	count     [2]int
	cover     [2]d19rectT
	area      [2]float64

	branchBuf      [d19maxNodes + 1]d19branchT
	branchCount    int
	coverSplit     d19rectT
	coverSplitArea float64
}

func d19New() *d19RTree {
	// We only support machine word size simple data type eg. integer index or object pointer.
	// Since we are storing as union with non data branch
	return &d19RTree{
		root: &d19nodeT{},
	}
}

/// Insert entry
/// \param a_min Min of bounding rect
/// \param a_max Max of bounding rect
/// \param a_dataId Positive Id of data.  Maybe zero, but negative numbers not allowed.
func (tr *d19RTree) Insert(min, max [d19numDims]float64, dataId interface{}) {
	var branch d19branchT
	branch.data = dataId
	for axis := 0; axis < d19numDims; axis++ {
		branch.rect.min[axis] = min[axis]
		branch.rect.max[axis] = max[axis]
	}
	d19insertRect(&branch, &tr.root, 0)
}

/// Remove entry
/// \param a_min Min of bounding rect
/// \param a_max Max of bounding rect
/// \param a_dataId Positive Id of data.  Maybe zero, but negative numbers not allowed.
func (tr *d19RTree) Remove(min, max [d19numDims]float64, dataId interface{}) {
	var rect d19rectT
	for axis := 0; axis < d19numDims; axis++ {
		rect.min[axis] = min[axis]
		rect.max[axis] = max[axis]
	}
	d19removeRect(&rect, dataId, &tr.root)
}

/// Find all within d19search rectangle
/// \param a_min Min of d19search bounding rect
/// \param a_max Max of d19search bounding rect
/// \param a_searchResult d19search result array.  Caller should set grow size. Function will reset, not append to array.
/// \param a_resultCallback Callback function to return result.  Callback should return 'true' to continue searching
/// \param a_context User context to pass as parameter to a_resultCallback
/// \return Returns the number of entries found
func (tr *d19RTree) Search(min, max [d19numDims]float64, resultCallback func(data interface{}) bool) int {
	var rect d19rectT
	for axis := 0; axis < d19numDims; axis++ {
		rect.min[axis] = min[axis]
		rect.max[axis] = max[axis]
	}
	foundCount, _ := d19search(tr.root, rect, 0, resultCallback)
	return foundCount
}

/// Count the data elements in this container.  This is slow as no internal counter is maintained.
func (tr *d19RTree) Count() int {
	var count int
	d19countRec(tr.root, &count)
	return count
}

/// Remove all entries from tree
func (tr *d19RTree) RemoveAll() {
	// Delete all existing nodes
	tr.root = &d19nodeT{}
}

func d19countRec(node *d19nodeT, count *int) {
	if node.isInternalNode() { // not a leaf node
		for index := 0; index < node.count; index++ {
			d19countRec(node.branch[index].child, count)
		}
	} else { // A leaf node
		*count += node.count
	}
}

// Inserts a new data rectangle into the index structure.
// Recursively descends tree, propagates splits back up.
// Returns 0 if node was not split.  Old node updated.
// If node was split, returns 1 and sets the pointer pointed to by
// new_node to point to the new node.  Old node updated to become one of two.
// The level argument specifies the number of steps up from the leaf
// level to insert; e.g. a data rectangle goes in at level = 0.
func d19insertRectRec(branch *d19branchT, node *d19nodeT, newNode **d19nodeT, level int) bool {
	// recurse until we reach the correct level for the new record. data records
	// will always be called with a_level == 0 (leaf)
	if node.level > level {
		// Still above level for insertion, go down tree recursively
		var otherNode *d19nodeT
		//var newBranch d19branchT

		// find the optimal branch for this record
		index := d19pickBranch(&branch.rect, node)

		// recursively insert this record into the picked branch
		childWasSplit := d19insertRectRec(branch, node.branch[index].child, &otherNode, level)

		if !childWasSplit {
			// Child was not split. Merge the bounding box of the new record with the
			// existing bounding box
			node.branch[index].rect = d19combineRect(&branch.rect, &(node.branch[index].rect))
			return false
		} else {
			// Child was split. The old branches are now re-partitioned to two nodes
			// so we have to re-calculate the bounding boxes of each node
			node.branch[index].rect = d19nodeCover(node.branch[index].child)
			var newBranch d19branchT
			newBranch.child = otherNode
			newBranch.rect = d19nodeCover(otherNode)

			// The old node is already a child of a_node. Now add the newly-created
			// node to a_node as well. a_node might be split because of that.
			return d19addBranch(&newBranch, node, newNode)
		}
	} else if node.level == level {
		// We have reached level for insertion. Add rect, split if necessary
		return d19addBranch(branch, node, newNode)
	} else {
		// Should never occur
		return false
	}
}

// Insert a data rectangle into an index structure.
// d19insertRect provides for splitting the root;
// returns 1 if root was split, 0 if it was not.
// The level argument specifies the number of steps up from the leaf
// level to insert; e.g. a data rectangle goes in at level = 0.
// InsertRect2 does the recursion.
//
func d19insertRect(branch *d19branchT, root **d19nodeT, level int) bool {
	var newNode *d19nodeT

	if d19insertRectRec(branch, *root, &newNode, level) { // Root split

		// Grow tree taller and new root
		newRoot := &d19nodeT{}
		newRoot.level = (*root).level + 1

		var newBranch d19branchT

		// add old root node as a child of the new root
		newBranch.rect = d19nodeCover(*root)
		newBranch.child = *root
		d19addBranch(&newBranch, newRoot, nil)

		// add the split node as a child of the new root
		newBranch.rect = d19nodeCover(newNode)
		newBranch.child = newNode
		d19addBranch(&newBranch, newRoot, nil)

		// set the new root as the root node
		*root = newRoot

		return true
	}
	return false
}

// Find the smallest rectangle that includes all rectangles in branches of a node.
func d19nodeCover(node *d19nodeT) d19rectT {
	rect := node.branch[0].rect
	for index := 1; index < node.count; index++ {
		rect = d19combineRect(&rect, &(node.branch[index].rect))
	}
	return rect
}

// Add a branch to a node.  Split the node if necessary.
// Returns 0 if node not split.  Old node updated.
// Returns 1 if node split, sets *new_node to address of new node.
// Old node updated, becomes one of two.
func d19addBranch(branch *d19branchT, node *d19nodeT, newNode **d19nodeT) bool {
	if node.count < d19maxNodes { // Split won't be necessary
		node.branch[node.count] = *branch
		node.count++
		return false
	} else {
		d19splitNode(node, branch, newNode)
		return true
	}
}

// Disconnect a dependent node.
// Caller must return (or stop using iteration index) after this as count has changed
func d19disconnectBranch(node *d19nodeT, index int) {
	// Remove element by swapping with the last element to prevent gaps in array
	node.branch[index] = node.branch[node.count-1]
	node.branch[node.count-1].data = nil
	node.branch[node.count-1].child = nil
	node.count--
}

// Pick a branch.  Pick the one that will need the smallest increase
// in area to accomodate the new rectangle.  This will result in the
// least total area for the covering rectangles in the current node.
// In case of a tie, pick the one which was smaller before, to get
// the best resolution when searching.
func d19pickBranch(rect *d19rectT, node *d19nodeT) int {
	var firstTime bool = true
	var increase float64
	var bestIncr float64 = -1
	var area float64
	var bestArea float64
	var best int
	var tempRect d19rectT

	for index := 0; index < node.count; index++ {
		curRect := &node.branch[index].rect
		area = d19calcRectVolume(curRect)
		tempRect = d19combineRect(rect, curRect)
		increase = d19calcRectVolume(&tempRect) - area
		if (increase < bestIncr) || firstTime {
			best = index
			bestArea = area
			bestIncr = increase
			firstTime = false
		} else if (increase == bestIncr) && (area < bestArea) {
			best = index
			bestArea = area
			bestIncr = increase
		}
	}
	return best
}

// Combine two rectangles into larger one containing both
func d19combineRect(rectA, rectB *d19rectT) d19rectT {
	var newRect d19rectT

	for index := 0; index < d19numDims; index++ {
		newRect.min[index] = d19fmin(rectA.min[index], rectB.min[index])
		newRect.max[index] = d19fmax(rectA.max[index], rectB.max[index])
	}

	return newRect
}

// Split a node.
// Divides the nodes branches and the extra one between two nodes.
// Old node is one of the new ones, and one really new one is created.
// Tries more than one method for choosing a partition, uses best result.
func d19splitNode(node *d19nodeT, branch *d19branchT, newNode **d19nodeT) {
	// Could just use local here, but member or external is faster since it is reused
	var localVars d19partitionVarsT
	parVars := &localVars

	// Load all the branches into a buffer, initialize old node
	d19getBranches(node, branch, parVars)

	// Find partition
	d19choosePartition(parVars, d19minNodes)

	// Create a new node to hold (about) half of the branches
	*newNode = &d19nodeT{}
	(*newNode).level = node.level

	// Put branches from buffer into 2 nodes according to the chosen partition
	node.count = 0
	d19loadNodes(node, *newNode, parVars)
}

// Calculate the n-dimensional volume of a rectangle
func d19rectVolume(rect *d19rectT) float64 {
	var volume float64 = 1
	for index := 0; index < d19numDims; index++ {
		volume *= rect.max[index] - rect.min[index]
	}
	return volume
}

// The exact volume of the bounding sphere for the given d19rectT
func d19rectSphericalVolume(rect *d19rectT) float64 {
	var sumOfSquares float64 = 0
	var radius float64

	for index := 0; index < d19numDims; index++ {
		halfExtent := (rect.max[index] - rect.min[index]) * 0.5
		sumOfSquares += halfExtent * halfExtent
	}

	radius = math.Sqrt(sumOfSquares)

	// Pow maybe slow, so test for common dims just use x*x, x*x*x.
	if d19numDims == 5 {
		return (radius * radius * radius * radius * radius * d19unitSphereVolume)
	} else if d19numDims == 4 {
		return (radius * radius * radius * radius * d19unitSphereVolume)
	} else if d19numDims == 3 {
		return (radius * radius * radius * d19unitSphereVolume)
	} else if d19numDims == 2 {
		return (radius * radius * d19unitSphereVolume)
	} else {
		return (math.Pow(radius, d19numDims) * d19unitSphereVolume)
	}
}

// Use one of the methods to calculate retangle volume
func d19calcRectVolume(rect *d19rectT) float64 {
	if d19useSphericalVolume {
		return d19rectSphericalVolume(rect) // Slower but helps certain merge cases
	} else { // RTREE_USE_SPHERICAL_VOLUME
		return d19rectVolume(rect) // Faster but can cause poor merges
	} // RTREE_USE_SPHERICAL_VOLUME
}

// Load branch buffer with branches from full node plus the extra branch.
func d19getBranches(node *d19nodeT, branch *d19branchT, parVars *d19partitionVarsT) {
	// Load the branch buffer
	for index := 0; index < d19maxNodes; index++ {
		parVars.branchBuf[index] = node.branch[index]
	}
	parVars.branchBuf[d19maxNodes] = *branch
	parVars.branchCount = d19maxNodes + 1

	// Calculate rect containing all in the set
	parVars.coverSplit = parVars.branchBuf[0].rect
	for index := 1; index < d19maxNodes+1; index++ {
		parVars.coverSplit = d19combineRect(&parVars.coverSplit, &parVars.branchBuf[index].rect)
	}
	parVars.coverSplitArea = d19calcRectVolume(&parVars.coverSplit)
}

// Method #0 for choosing a partition:
// As the seeds for the two groups, pick the two rects that would waste the
// most area if covered by a single rectangle, i.e. evidently the worst pair
// to have in the same group.
// Of the remaining, one at a time is chosen to be put in one of the two groups.
// The one chosen is the one with the greatest difference in area expansion
// depending on which group - the rect most strongly attracted to one group
// and repelled from the other.
// If one group gets too full (more would force other group to violate min
// fill requirement) then other group gets the rest.
// These last are the ones that can go in either group most easily.
func d19choosePartition(parVars *d19partitionVarsT, minFill int) {
	var biggestDiff float64
	var group, chosen, betterGroup int

	d19initParVars(parVars, parVars.branchCount, minFill)
	d19pickSeeds(parVars)

	for ((parVars.count[0] + parVars.count[1]) < parVars.total) &&
		(parVars.count[0] < (parVars.total - parVars.minFill)) &&
		(parVars.count[1] < (parVars.total - parVars.minFill)) {
		biggestDiff = -1
		for index := 0; index < parVars.total; index++ {
			if d19notTaken == parVars.partition[index] {
				curRect := &parVars.branchBuf[index].rect
				rect0 := d19combineRect(curRect, &parVars.cover[0])
				rect1 := d19combineRect(curRect, &parVars.cover[1])
				growth0 := d19calcRectVolume(&rect0) - parVars.area[0]
				growth1 := d19calcRectVolume(&rect1) - parVars.area[1]
				diff := growth1 - growth0
				if diff >= 0 {
					group = 0
				} else {
					group = 1
					diff = -diff
				}

				if diff > biggestDiff {
					biggestDiff = diff
					chosen = index
					betterGroup = group
				} else if (diff == biggestDiff) && (parVars.count[group] < parVars.count[betterGroup]) {
					chosen = index
					betterGroup = group
				}
			}
		}
		d19classify(chosen, betterGroup, parVars)
	}

	// If one group too full, put remaining rects in the other
	if (parVars.count[0] + parVars.count[1]) < parVars.total {
		if parVars.count[0] >= parVars.total-parVars.minFill {
			group = 1
		} else {
			group = 0
		}
		for index := 0; index < parVars.total; index++ {
			if d19notTaken == parVars.partition[index] {
				d19classify(index, group, parVars)
			}
		}
	}
}

// Copy branches from the buffer into two nodes according to the partition.
func d19loadNodes(nodeA, nodeB *d19nodeT, parVars *d19partitionVarsT) {
	for index := 0; index < parVars.total; index++ {
		targetNodeIndex := parVars.partition[index]
		targetNodes := []*d19nodeT{nodeA, nodeB}

		// It is assured that d19addBranch here will not cause a node split.
		d19addBranch(&parVars.branchBuf[index], targetNodes[targetNodeIndex], nil)
	}
}

// Initialize a d19partitionVarsT structure.
func d19initParVars(parVars *d19partitionVarsT, maxRects, minFill int) {
	parVars.count[0] = 0
	parVars.count[1] = 0
	parVars.area[0] = 0
	parVars.area[1] = 0
	parVars.total = maxRects
	parVars.minFill = minFill
	for index := 0; index < maxRects; index++ {
		parVars.partition[index] = d19notTaken
	}
}

func d19pickSeeds(parVars *d19partitionVarsT) {
	var seed0, seed1 int
	var worst, waste float64
	var area [d19maxNodes + 1]float64

	for index := 0; index < parVars.total; index++ {
		area[index] = d19calcRectVolume(&parVars.branchBuf[index].rect)
	}

	worst = -parVars.coverSplitArea - 1
	for indexA := 0; indexA < parVars.total-1; indexA++ {
		for indexB := indexA + 1; indexB < parVars.total; indexB++ {
			oneRect := d19combineRect(&parVars.branchBuf[indexA].rect, &parVars.branchBuf[indexB].rect)
			waste = d19calcRectVolume(&oneRect) - area[indexA] - area[indexB]
			if waste > worst {
				worst = waste
				seed0 = indexA
				seed1 = indexB
			}
		}
	}

	d19classify(seed0, 0, parVars)
	d19classify(seed1, 1, parVars)
}

// Put a branch in one of the groups.
func d19classify(index, group int, parVars *d19partitionVarsT) {
	parVars.partition[index] = group

	// Calculate combined rect
	if parVars.count[group] == 0 {
		parVars.cover[group] = parVars.branchBuf[index].rect
	} else {
		parVars.cover[group] = d19combineRect(&parVars.branchBuf[index].rect, &parVars.cover[group])
	}

	// Calculate volume of combined rect
	parVars.area[group] = d19calcRectVolume(&parVars.cover[group])

	parVars.count[group]++
}

// Delete a data rectangle from an index structure.
// Pass in a pointer to a d19rectT, the tid of the record, ptr to ptr to root node.
// Returns 1 if record not found, 0 if success.
// d19removeRect provides for eliminating the root.
func d19removeRect(rect *d19rectT, id interface{}, root **d19nodeT) bool {
	var reInsertList *d19listNodeT

	if !d19removeRectRec(rect, id, *root, &reInsertList) {
		// Found and deleted a data item
		// Reinsert any branches from eliminated nodes
		for reInsertList != nil {
			tempNode := reInsertList.node

			for index := 0; index < tempNode.count; index++ {
				// TODO go over this code. should I use (tempNode->m_level - 1)?
				d19insertRect(&tempNode.branch[index], root, tempNode.level)
			}
			reInsertList = reInsertList.next
		}

		// Check for redundant root (not leaf, 1 child) and eliminate TODO replace
		// if with while? In case there is a whole branch of redundant roots...
		if (*root).count == 1 && (*root).isInternalNode() {
			tempNode := (*root).branch[0].child
			*root = tempNode
		}
		return false
	} else {
		return true
	}
}

// Delete a rectangle from non-root part of an index structure.
// Called by d19removeRect.  Descends tree recursively,
// merges branches on the way back up.
// Returns 1 if record not found, 0 if success.
func d19removeRectRec(rect *d19rectT, id interface{}, node *d19nodeT, listNode **d19listNodeT) bool {
	if node.isInternalNode() { // not a leaf node
		for index := 0; index < node.count; index++ {
			if d19overlap(*rect, node.branch[index].rect) {
				if !d19removeRectRec(rect, id, node.branch[index].child, listNode) {
					if node.branch[index].child.count >= d19minNodes {
						// child removed, just resize parent rect
						node.branch[index].rect = d19nodeCover(node.branch[index].child)
					} else {
						// child removed, not enough entries in node, eliminate node
						d19reInsert(node.branch[index].child, listNode)
						d19disconnectBranch(node, index) // Must return after this call as count has changed
					}
					return false
				}
			}
		}
		return true
	} else { // A leaf node
		for index := 0; index < node.count; index++ {
			if node.branch[index].data == id {
				d19disconnectBranch(node, index) // Must return after this call as count has changed
				return false
			}
		}
		return true
	}
}

// Decide whether two rectangles d19overlap.
func d19overlap(rectA, rectB d19rectT) bool {
	for index := 0; index < d19numDims; index++ {
		if rectA.min[index] > rectB.max[index] ||
			rectB.min[index] > rectA.max[index] {
			return false
		}
	}
	return true
}

// Add a node to the reinsertion list.  All its branches will later
// be reinserted into the index structure.
func d19reInsert(node *d19nodeT, listNode **d19listNodeT) {
	newListNode := &d19listNodeT{}
	newListNode.node = node
	newListNode.next = *listNode
	*listNode = newListNode
}

// d19search in an index tree or subtree for all data retangles that d19overlap the argument rectangle.
func d19search(node *d19nodeT, rect d19rectT, foundCount int, resultCallback func(data interface{}) bool) (int, bool) {
	if node.isInternalNode() {
		// This is an internal node in the tree
		for index := 0; index < node.count; index++ {
			if d19overlap(rect, node.branch[index].rect) {
				var ok bool
				foundCount, ok = d19search(node.branch[index].child, rect, foundCount, resultCallback)
				if !ok {
					// The callback indicated to stop searching
					return foundCount, false
				}
			}
		}
	} else {
		// This is a leaf node
		for index := 0; index < node.count; index++ {
			if d19overlap(rect, node.branch[index].rect) {
				id := node.branch[index].data
				foundCount++
				if !resultCallback(id) {
					return foundCount, false // Don't continue searching
				}

			}
		}
	}
	return foundCount, true // Continue searching
}

func d20fmin(a, b float64) float64 {
	if a < b {
		return a
	}
	return b
}
func d20fmax(a, b float64) float64 {
	if a > b {
		return a
	}
	return b
}

const (
	d20numDims            = 20
	d20maxNodes           = 8
	d20minNodes           = d20maxNodes / 2
	d20useSphericalVolume = true // Better split classification, may be slower on some systems
)

var d20unitSphereVolume = []float64{
	0.000000, 2.000000, 3.141593, // Dimension  0,1,2
	4.188790, 4.934802, 5.263789, // Dimension  3,4,5
	5.167713, 4.724766, 4.058712, // Dimension  6,7,8
	3.298509, 2.550164, 1.884104, // Dimension  9,10,11
	1.335263, 0.910629, 0.599265, // Dimension  12,13,14
	0.381443, 0.235331, 0.140981, // Dimension  15,16,17
	0.082146, 0.046622, 0.025807, // Dimension  18,19,20
}[d20numDims]

type d20RTree struct {
	root *d20nodeT ///< Root of tree
}

/// Minimal bounding rectangle (n-dimensional)
type d20rectT struct {
	min [d20numDims]float64 ///< Min dimensions of bounding box
	max [d20numDims]float64 ///< Max dimensions of bounding box
}

/// May be data or may be another subtree
/// The parents level determines this.
/// If the parents level is 0, then this is data
type d20branchT struct {
	rect  d20rectT    ///< Bounds
	child *d20nodeT   ///< Child node
	data  interface{} ///< Data Id or Ptr
}

/// d20nodeT for each branch level
type d20nodeT struct {
	count  int                     ///< Count
	level  int                     ///< Leaf is zero, others positive
	branch [d20maxNodes]d20branchT ///< Branch
}

func (node *d20nodeT) isInternalNode() bool {
	return (node.level > 0) // Not a leaf, but a internal node
}
func (node *d20nodeT) isLeaf() bool {
	return (node.level == 0) // A leaf, contains data
}

/// A link list of nodes for reinsertion after a delete operation
type d20listNodeT struct {
	next *d20listNodeT ///< Next in list
	node *d20nodeT     ///< Node
}

const d20notTaken = -1 // indicates that position

/// Variables for finding a split partition
type d20partitionVarsT struct {
	partition [d20maxNodes + 1]int
	total     int
	minFill   int
	count     [2]int
	cover     [2]d20rectT
	area      [2]float64

	branchBuf      [d20maxNodes + 1]d20branchT
	branchCount    int
	coverSplit     d20rectT
	coverSplitArea float64
}

func d20New() *d20RTree {
	// We only support machine word size simple data type eg. integer index or object pointer.
	// Since we are storing as union with non data branch
	return &d20RTree{
		root: &d20nodeT{},
	}
}

/// Insert entry
/// \param a_min Min of bounding rect
/// \param a_max Max of bounding rect
/// \param a_dataId Positive Id of data.  Maybe zero, but negative numbers not allowed.
func (tr *d20RTree) Insert(min, max [d20numDims]float64, dataId interface{}) {
	var branch d20branchT
	branch.data = dataId
	for axis := 0; axis < d20numDims; axis++ {
		branch.rect.min[axis] = min[axis]
		branch.rect.max[axis] = max[axis]
	}
	d20insertRect(&branch, &tr.root, 0)
}

/// Remove entry
/// \param a_min Min of bounding rect
/// \param a_max Max of bounding rect
/// \param a_dataId Positive Id of data.  Maybe zero, but negative numbers not allowed.
func (tr *d20RTree) Remove(min, max [d20numDims]float64, dataId interface{}) {
	var rect d20rectT
	for axis := 0; axis < d20numDims; axis++ {
		rect.min[axis] = min[axis]
		rect.max[axis] = max[axis]
	}
	d20removeRect(&rect, dataId, &tr.root)
}

/// Find all within d20search rectangle
/// \param a_min Min of d20search bounding rect
/// \param a_max Max of d20search bounding rect
/// \param a_searchResult d20search result array.  Caller should set grow size. Function will reset, not append to array.
/// \param a_resultCallback Callback function to return result.  Callback should return 'true' to continue searching
/// \param a_context User context to pass as parameter to a_resultCallback
/// \return Returns the number of entries found
func (tr *d20RTree) Search(min, max [d20numDims]float64, resultCallback func(data interface{}) bool) int {
	var rect d20rectT
	for axis := 0; axis < d20numDims; axis++ {
		rect.min[axis] = min[axis]
		rect.max[axis] = max[axis]
	}
	foundCount, _ := d20search(tr.root, rect, 0, resultCallback)
	return foundCount
}

/// Count the data elements in this container.  This is slow as no internal counter is maintained.
func (tr *d20RTree) Count() int {
	var count int
	d20countRec(tr.root, &count)
	return count
}

/// Remove all entries from tree
func (tr *d20RTree) RemoveAll() {
	// Delete all existing nodes
	tr.root = &d20nodeT{}
}

func d20countRec(node *d20nodeT, count *int) {
	if node.isInternalNode() { // not a leaf node
		for index := 0; index < node.count; index++ {
			d20countRec(node.branch[index].child, count)
		}
	} else { // A leaf node
		*count += node.count
	}
}

// Inserts a new data rectangle into the index structure.
// Recursively descends tree, propagates splits back up.
// Returns 0 if node was not split.  Old node updated.
// If node was split, returns 1 and sets the pointer pointed to by
// new_node to point to the new node.  Old node updated to become one of two.
// The level argument specifies the number of steps up from the leaf
// level to insert; e.g. a data rectangle goes in at level = 0.
func d20insertRectRec(branch *d20branchT, node *d20nodeT, newNode **d20nodeT, level int) bool {
	// recurse until we reach the correct level for the new record. data records
	// will always be called with a_level == 0 (leaf)
	if node.level > level {
		// Still above level for insertion, go down tree recursively
		var otherNode *d20nodeT
		//var newBranch d20branchT

		// find the optimal branch for this record
		index := d20pickBranch(&branch.rect, node)

		// recursively insert this record into the picked branch
		childWasSplit := d20insertRectRec(branch, node.branch[index].child, &otherNode, level)

		if !childWasSplit {
			// Child was not split. Merge the bounding box of the new record with the
			// existing bounding box
			node.branch[index].rect = d20combineRect(&branch.rect, &(node.branch[index].rect))
			return false
		} else {
			// Child was split. The old branches are now re-partitioned to two nodes
			// so we have to re-calculate the bounding boxes of each node
			node.branch[index].rect = d20nodeCover(node.branch[index].child)
			var newBranch d20branchT
			newBranch.child = otherNode
			newBranch.rect = d20nodeCover(otherNode)

			// The old node is already a child of a_node. Now add the newly-created
			// node to a_node as well. a_node might be split because of that.
			return d20addBranch(&newBranch, node, newNode)
		}
	} else if node.level == level {
		// We have reached level for insertion. Add rect, split if necessary
		return d20addBranch(branch, node, newNode)
	} else {
		// Should never occur
		return false
	}
}

// Insert a data rectangle into an index structure.
// d20insertRect provides for splitting the root;
// returns 1 if root was split, 0 if it was not.
// The level argument specifies the number of steps up from the leaf
// level to insert; e.g. a data rectangle goes in at level = 0.
// InsertRect2 does the recursion.
//
func d20insertRect(branch *d20branchT, root **d20nodeT, level int) bool {
	var newNode *d20nodeT

	if d20insertRectRec(branch, *root, &newNode, level) { // Root split

		// Grow tree taller and new root
		newRoot := &d20nodeT{}
		newRoot.level = (*root).level + 1

		var newBranch d20branchT

		// add old root node as a child of the new root
		newBranch.rect = d20nodeCover(*root)
		newBranch.child = *root
		d20addBranch(&newBranch, newRoot, nil)

		// add the split node as a child of the new root
		newBranch.rect = d20nodeCover(newNode)
		newBranch.child = newNode
		d20addBranch(&newBranch, newRoot, nil)

		// set the new root as the root node
		*root = newRoot

		return true
	}
	return false
}

// Find the smallest rectangle that includes all rectangles in branches of a node.
func d20nodeCover(node *d20nodeT) d20rectT {
	rect := node.branch[0].rect
	for index := 1; index < node.count; index++ {
		rect = d20combineRect(&rect, &(node.branch[index].rect))
	}
	return rect
}

// Add a branch to a node.  Split the node if necessary.
// Returns 0 if node not split.  Old node updated.
// Returns 1 if node split, sets *new_node to address of new node.
// Old node updated, becomes one of two.
func d20addBranch(branch *d20branchT, node *d20nodeT, newNode **d20nodeT) bool {
	if node.count < d20maxNodes { // Split won't be necessary
		node.branch[node.count] = *branch
		node.count++
		return false
	} else {
		d20splitNode(node, branch, newNode)
		return true
	}
}

// Disconnect a dependent node.
// Caller must return (or stop using iteration index) after this as count has changed
func d20disconnectBranch(node *d20nodeT, index int) {
	// Remove element by swapping with the last element to prevent gaps in array
	node.branch[index] = node.branch[node.count-1]
	node.branch[node.count-1].data = nil
	node.branch[node.count-1].child = nil
	node.count--
}

// Pick a branch.  Pick the one that will need the smallest increase
// in area to accomodate the new rectangle.  This will result in the
// least total area for the covering rectangles in the current node.
// In case of a tie, pick the one which was smaller before, to get
// the best resolution when searching.
func d20pickBranch(rect *d20rectT, node *d20nodeT) int {
	var firstTime bool = true
	var increase float64
	var bestIncr float64 = -1
	var area float64
	var bestArea float64
	var best int
	var tempRect d20rectT

	for index := 0; index < node.count; index++ {
		curRect := &node.branch[index].rect
		area = d20calcRectVolume(curRect)
		tempRect = d20combineRect(rect, curRect)
		increase = d20calcRectVolume(&tempRect) - area
		if (increase < bestIncr) || firstTime {
			best = index
			bestArea = area
			bestIncr = increase
			firstTime = false
		} else if (increase == bestIncr) && (area < bestArea) {
			best = index
			bestArea = area
			bestIncr = increase
		}
	}
	return best
}

// Combine two rectangles into larger one containing both
func d20combineRect(rectA, rectB *d20rectT) d20rectT {
	var newRect d20rectT

	for index := 0; index < d20numDims; index++ {
		newRect.min[index] = d20fmin(rectA.min[index], rectB.min[index])
		newRect.max[index] = d20fmax(rectA.max[index], rectB.max[index])
	}

	return newRect
}

// Split a node.
// Divides the nodes branches and the extra one between two nodes.
// Old node is one of the new ones, and one really new one is created.
// Tries more than one method for choosing a partition, uses best result.
func d20splitNode(node *d20nodeT, branch *d20branchT, newNode **d20nodeT) {
	// Could just use local here, but member or external is faster since it is reused
	var localVars d20partitionVarsT
	parVars := &localVars

	// Load all the branches into a buffer, initialize old node
	d20getBranches(node, branch, parVars)

	// Find partition
	d20choosePartition(parVars, d20minNodes)

	// Create a new node to hold (about) half of the branches
	*newNode = &d20nodeT{}
	(*newNode).level = node.level

	// Put branches from buffer into 2 nodes according to the chosen partition
	node.count = 0
	d20loadNodes(node, *newNode, parVars)
}

// Calculate the n-dimensional volume of a rectangle
func d20rectVolume(rect *d20rectT) float64 {
	var volume float64 = 1
	for index := 0; index < d20numDims; index++ {
		volume *= rect.max[index] - rect.min[index]
	}
	return volume
}

// The exact volume of the bounding sphere for the given d20rectT
func d20rectSphericalVolume(rect *d20rectT) float64 {
	var sumOfSquares float64 = 0
	var radius float64

	for index := 0; index < d20numDims; index++ {
		halfExtent := (rect.max[index] - rect.min[index]) * 0.5
		sumOfSquares += halfExtent * halfExtent
	}

	radius = math.Sqrt(sumOfSquares)

	// Pow maybe slow, so test for common dims just use x*x, x*x*x.
	if d20numDims == 5 {
		return (radius * radius * radius * radius * radius * d20unitSphereVolume)
	} else if d20numDims == 4 {
		return (radius * radius * radius * radius * d20unitSphereVolume)
	} else if d20numDims == 3 {
		return (radius * radius * radius * d20unitSphereVolume)
	} else if d20numDims == 2 {
		return (radius * radius * d20unitSphereVolume)
	} else {
		return (math.Pow(radius, d20numDims) * d20unitSphereVolume)
	}
}

// Use one of the methods to calculate retangle volume
func d20calcRectVolume(rect *d20rectT) float64 {
	if d20useSphericalVolume {
		return d20rectSphericalVolume(rect) // Slower but helps certain merge cases
	} else { // RTREE_USE_SPHERICAL_VOLUME
		return d20rectVolume(rect) // Faster but can cause poor merges
	} // RTREE_USE_SPHERICAL_VOLUME
}

// Load branch buffer with branches from full node plus the extra branch.
func d20getBranches(node *d20nodeT, branch *d20branchT, parVars *d20partitionVarsT) {
	// Load the branch buffer
	for index := 0; index < d20maxNodes; index++ {
		parVars.branchBuf[index] = node.branch[index]
	}
	parVars.branchBuf[d20maxNodes] = *branch
	parVars.branchCount = d20maxNodes + 1

	// Calculate rect containing all in the set
	parVars.coverSplit = parVars.branchBuf[0].rect
	for index := 1; index < d20maxNodes+1; index++ {
		parVars.coverSplit = d20combineRect(&parVars.coverSplit, &parVars.branchBuf[index].rect)
	}
	parVars.coverSplitArea = d20calcRectVolume(&parVars.coverSplit)
}

// Method #0 for choosing a partition:
// As the seeds for the two groups, pick the two rects that would waste the
// most area if covered by a single rectangle, i.e. evidently the worst pair
// to have in the same group.
// Of the remaining, one at a time is chosen to be put in one of the two groups.
// The one chosen is the one with the greatest difference in area expansion
// depending on which group - the rect most strongly attracted to one group
// and repelled from the other.
// If one group gets too full (more would force other group to violate min
// fill requirement) then other group gets the rest.
// These last are the ones that can go in either group most easily.
func d20choosePartition(parVars *d20partitionVarsT, minFill int) {
	var biggestDiff float64
	var group, chosen, betterGroup int

	d20initParVars(parVars, parVars.branchCount, minFill)
	d20pickSeeds(parVars)

	for ((parVars.count[0] + parVars.count[1]) < parVars.total) &&
		(parVars.count[0] < (parVars.total - parVars.minFill)) &&
		(parVars.count[1] < (parVars.total - parVars.minFill)) {
		biggestDiff = -1
		for index := 0; index < parVars.total; index++ {
			if d20notTaken == parVars.partition[index] {
				curRect := &parVars.branchBuf[index].rect
				rect0 := d20combineRect(curRect, &parVars.cover[0])
				rect1 := d20combineRect(curRect, &parVars.cover[1])
				growth0 := d20calcRectVolume(&rect0) - parVars.area[0]
				growth1 := d20calcRectVolume(&rect1) - parVars.area[1]
				diff := growth1 - growth0
				if diff >= 0 {
					group = 0
				} else {
					group = 1
					diff = -diff
				}

				if diff > biggestDiff {
					biggestDiff = diff
					chosen = index
					betterGroup = group
				} else if (diff == biggestDiff) && (parVars.count[group] < parVars.count[betterGroup]) {
					chosen = index
					betterGroup = group
				}
			}
		}
		d20classify(chosen, betterGroup, parVars)
	}

	// If one group too full, put remaining rects in the other
	if (parVars.count[0] + parVars.count[1]) < parVars.total {
		if parVars.count[0] >= parVars.total-parVars.minFill {
			group = 1
		} else {
			group = 0
		}
		for index := 0; index < parVars.total; index++ {
			if d20notTaken == parVars.partition[index] {
				d20classify(index, group, parVars)
			}
		}
	}
}

// Copy branches from the buffer into two nodes according to the partition.
func d20loadNodes(nodeA, nodeB *d20nodeT, parVars *d20partitionVarsT) {
	for index := 0; index < parVars.total; index++ {
		targetNodeIndex := parVars.partition[index]
		targetNodes := []*d20nodeT{nodeA, nodeB}

		// It is assured that d20addBranch here will not cause a node split.
		d20addBranch(&parVars.branchBuf[index], targetNodes[targetNodeIndex], nil)
	}
}

// Initialize a d20partitionVarsT structure.
func d20initParVars(parVars *d20partitionVarsT, maxRects, minFill int) {
	parVars.count[0] = 0
	parVars.count[1] = 0
	parVars.area[0] = 0
	parVars.area[1] = 0
	parVars.total = maxRects
	parVars.minFill = minFill
	for index := 0; index < maxRects; index++ {
		parVars.partition[index] = d20notTaken
	}
}

func d20pickSeeds(parVars *d20partitionVarsT) {
	var seed0, seed1 int
	var worst, waste float64
	var area [d20maxNodes + 1]float64

	for index := 0; index < parVars.total; index++ {
		area[index] = d20calcRectVolume(&parVars.branchBuf[index].rect)
	}

	worst = -parVars.coverSplitArea - 1
	for indexA := 0; indexA < parVars.total-1; indexA++ {
		for indexB := indexA + 1; indexB < parVars.total; indexB++ {
			oneRect := d20combineRect(&parVars.branchBuf[indexA].rect, &parVars.branchBuf[indexB].rect)
			waste = d20calcRectVolume(&oneRect) - area[indexA] - area[indexB]
			if waste > worst {
				worst = waste
				seed0 = indexA
				seed1 = indexB
			}
		}
	}

	d20classify(seed0, 0, parVars)
	d20classify(seed1, 1, parVars)
}

// Put a branch in one of the groups.
func d20classify(index, group int, parVars *d20partitionVarsT) {
	parVars.partition[index] = group

	// Calculate combined rect
	if parVars.count[group] == 0 {
		parVars.cover[group] = parVars.branchBuf[index].rect
	} else {
		parVars.cover[group] = d20combineRect(&parVars.branchBuf[index].rect, &parVars.cover[group])
	}

	// Calculate volume of combined rect
	parVars.area[group] = d20calcRectVolume(&parVars.cover[group])

	parVars.count[group]++
}

// Delete a data rectangle from an index structure.
// Pass in a pointer to a d20rectT, the tid of the record, ptr to ptr to root node.
// Returns 1 if record not found, 0 if success.
// d20removeRect provides for eliminating the root.
func d20removeRect(rect *d20rectT, id interface{}, root **d20nodeT) bool {
	var reInsertList *d20listNodeT

	if !d20removeRectRec(rect, id, *root, &reInsertList) {
		// Found and deleted a data item
		// Reinsert any branches from eliminated nodes
		for reInsertList != nil {
			tempNode := reInsertList.node

			for index := 0; index < tempNode.count; index++ {
				// TODO go over this code. should I use (tempNode->m_level - 1)?
				d20insertRect(&tempNode.branch[index], root, tempNode.level)
			}
			reInsertList = reInsertList.next
		}

		// Check for redundant root (not leaf, 1 child) and eliminate TODO replace
		// if with while? In case there is a whole branch of redundant roots...
		if (*root).count == 1 && (*root).isInternalNode() {
			tempNode := (*root).branch[0].child
			*root = tempNode
		}
		return false
	} else {
		return true
	}
}

// Delete a rectangle from non-root part of an index structure.
// Called by d20removeRect.  Descends tree recursively,
// merges branches on the way back up.
// Returns 1 if record not found, 0 if success.
func d20removeRectRec(rect *d20rectT, id interface{}, node *d20nodeT, listNode **d20listNodeT) bool {
	if node.isInternalNode() { // not a leaf node
		for index := 0; index < node.count; index++ {
			if d20overlap(*rect, node.branch[index].rect) {
				if !d20removeRectRec(rect, id, node.branch[index].child, listNode) {
					if node.branch[index].child.count >= d20minNodes {
						// child removed, just resize parent rect
						node.branch[index].rect = d20nodeCover(node.branch[index].child)
					} else {
						// child removed, not enough entries in node, eliminate node
						d20reInsert(node.branch[index].child, listNode)
						d20disconnectBranch(node, index) // Must return after this call as count has changed
					}
					return false
				}
			}
		}
		return true
	} else { // A leaf node
		for index := 0; index < node.count; index++ {
			if node.branch[index].data == id {
				d20disconnectBranch(node, index) // Must return after this call as count has changed
				return false
			}
		}
		return true
	}
}

// Decide whether two rectangles d20overlap.
func d20overlap(rectA, rectB d20rectT) bool {
	for index := 0; index < d20numDims; index++ {
		if rectA.min[index] > rectB.max[index] ||
			rectB.min[index] > rectA.max[index] {
			return false
		}
	}
	return true
}

// Add a node to the reinsertion list.  All its branches will later
// be reinserted into the index structure.
func d20reInsert(node *d20nodeT, listNode **d20listNodeT) {
	newListNode := &d20listNodeT{}
	newListNode.node = node
	newListNode.next = *listNode
	*listNode = newListNode
}

// d20search in an index tree or subtree for all data retangles that d20overlap the argument rectangle.
func d20search(node *d20nodeT, rect d20rectT, foundCount int, resultCallback func(data interface{}) bool) (int, bool) {
	if node.isInternalNode() {
		// This is an internal node in the tree
		for index := 0; index < node.count; index++ {
			if d20overlap(rect, node.branch[index].rect) {
				var ok bool
				foundCount, ok = d20search(node.branch[index].child, rect, foundCount, resultCallback)
				if !ok {
					// The callback indicated to stop searching
					return foundCount, false
				}
			}
		}
	} else {
		// This is a leaf node
		for index := 0; index < node.count; index++ {
			if d20overlap(rect, node.branch[index].rect) {
				id := node.branch[index].data
				foundCount++
				if !resultCallback(id) {
					return foundCount, false // Don't continue searching
				}

			}
		}
	}
	return foundCount, true // Continue searching
}
