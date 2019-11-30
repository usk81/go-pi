package calc

import (
	"fmt"
	"math"
)

type Rohrer struct{}

func CalcRohrer(w, h float64) (result Result, err error) {
	ro := new(Rohrer)
	i, err := ro.Calc(w, h)
	if err != nil {
		return
	}
	c, err := ro.Classification(i)
	if err != nil {
		return
	}
	s, err := ro.StatusByClassification(c)
	if err != nil {
		return
	}

	return Result{
		Classification: c,
		Index:          i,
		IndexType:      TypeRohrer,
		Status:         s,
	}, nil
}

func (ro *Rohrer) Calc(w, h float64) (r float64, err error) {
	if w <= 0 || h <= 0 {
		err = fmt.Errorf("request parameters are invalid")
	}
	r = (w / math.Pow(h/100, 3)) * 10
	return
}

func (ro *Rohrer) Classification(pi float64) (r string, err error) {
	if pi <= 0 {
		err = fmt.Errorf("request ponderal index is incorrect")
		return
	}
	switch {
	case pi < 100:
		r = ClassSevereThinness
	case pi >= 100 && pi < 115:
		r = ClassUnderWeight
	case pi >= 115 && pi < 145:
		r = ClassNormal
	case pi >= 145 && pi < 160:
		r = ClassOverWight
	default:
		r = ClassSevereObesity
	}
	return
}

func (ro *Rohrer) Status(pi float64) (r string, err error) {
	c, err := ro.Classification(pi)
	if err != nil {
		err = fmt.Errorf("failt to get classification : %v", err)
	}
	return ro.StatusByClassification(c)
}

func (ro *Rohrer) StatusByClassification(c string) (r string, err error) {
	switch c {
	case ClassNormal:
		r = StatusNormal
	case ClassUnderWeight, ClassOverWight:
		r = StatusWarn
	case ClassSevereThinness, ClassObesity, ClassSevereObesity:
		r = StatusFatal
	default:
		err = fmt.Errorf("%s is unknown classification", c)
		return
	}
	return
}
