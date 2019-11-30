package calc

import (
	"fmt"
	"math"
)

type BMI struct{}

func CalcBMI(w, h float64) (result Result, err error) {
	bi := new(BMI)
	i, err := bi.Calc(w, h)
	if err != nil {
		return
	}
	c, err := bi.Classification(i)
	if err != nil {
		return
	}
	s, err := bi.StatusByClassification(c)
	if err != nil {
		return
	}

	return Result{
		Classification: c,
		Index:          i,
		IndexType:      TypeBMI,
		Status:         s,
	}, nil
}

func (bi *BMI) Calc(w, h float64) (r float64, err error) {
	if w <= 0 || h <= 0 {
		err = fmt.Errorf("request parameters are invalid")
	}
	r = w / math.Pow(h/100, 2)
	return
}

func (bi *BMI) Classification(pi float64) (r string, err error) {
	if pi <= 0 {
		err = fmt.Errorf("request ponderal index is incorrect")
		return
	}
	switch {
	case pi < 16:
		r = ClassSevereThinness
	case pi >= 16 && pi < 18.5:
		r = ClassUnderWeight
	case pi >= 18.5 && pi < 25:
		r = ClassNormal
	case pi >= 25 && pi < 30:
		r = ClassOverWight
	case pi >= 30 && pi < 35:
		r = ClassObesity
	default:
		r = ClassSevereObesity
	}
	return
}

func (bi *BMI) Status(pi float64) (r string, err error) {
	if pi <= 0 {
		err = fmt.Errorf("invalid request")
		return
	}
	c, err := bi.Classification(pi)
	if err != nil {
		err = fmt.Errorf("failt to get classification : %v", err)
	}
	return bi.StatusByClassification(c)
}

func (bi *BMI) StatusByClassification(c string) (r string, err error) {
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
