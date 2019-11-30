package calc

import (
	"fmt"
	"math"
)

type Kaup struct{}

func CalcKaup(w, h, age float64) (result Result, err error) {
	k := new(Kaup)
	i, err := k.Calc(w, h)
	if err != nil {
		return
	}
	c, err := k.Classification(i, age)
	if err != nil {
		return
	}
	s, err := k.StatusByClassification(c)
	if err != nil {
		return
	}

	return Result{
		Classification: c,
		Index:          i,
		IndexType:      TypeKaup,
		Status:         s,
	}, nil
}

func (k *Kaup) Calc(w, h float64) (r float64, err error) {
	if w <= 0 || h <= 0 {
		err = fmt.Errorf("request parameters are invalid")
	}
	r = w / math.Pow(h/100, 2)
	return
}

func (k *Kaup) Classification(pi float64, age float64) (r string, err error) {
	if pi <= 0 {
		err = fmt.Errorf("request ponderal index is incorrect")
		return
	}
	if age < 0 {
		err = fmt.Errorf("age is bigger than zero")
		return
	}
	switch {
	case age < 1:
		switch {
		case pi < 14.5:
			r = ClassSevereThinness
		case pi >= 14.5 && pi < 16:
			r = ClassUnderWeight
		case pi >= 16 && pi < 18:
			r = ClassNormal
		case pi >= 18 && pi < 20:
			r = ClassOverWight
		default:
			r = ClassSevereObesity
		}
	case age >= 1 && age < 1.5:
		switch {
		case pi < 14.5:
			r = ClassSevereThinness
		case pi >= 14.5 && pi < 15.5:
			r = ClassUnderWeight
		case pi >= 15.5 && pi < 17.5:
			r = ClassNormal
		case pi >= 17.5 && pi < 19.5:
			r = ClassOverWight
		default:
			r = ClassSevereObesity
		}
	case age >= 1.5 && age < 2:
		switch {
		case pi < 14:
			r = ClassSevereThinness
		case pi >= 14 && pi < 15:
			r = ClassUnderWeight
		case pi >= 15 && pi < 17:
			r = ClassNormal
		case pi >= 17 && pi < 19:
			r = ClassOverWight
		default:
			r = ClassSevereObesity
		}
	case age >= 2 && age < 3:
		switch {
		case pi < 13.5:
			r = ClassSevereThinness
		case pi >= 13.5 && pi < 15:
			r = ClassUnderWeight
		case pi >= 15 && pi < 17:
			r = ClassNormal
		case pi >= 17 && pi < 18.5:
			r = ClassOverWight
		default:
			r = ClassSevereObesity
		}
	case age >= 3 && age < 4:
		switch {
		case pi < 13.5:
			r = ClassSevereThinness
		case pi >= 13.5 && pi < 14.5:
			r = ClassUnderWeight
		case pi >= 14.5 && pi < 16.5:
			r = ClassNormal
		case pi >= 16.5 && pi < 18:
			r = ClassOverWight
		default:
			r = ClassSevereObesity
		}
	case age >= 4 && age < 5:
		switch {
		case pi < 13:
			r = ClassSevereThinness
		case pi >= 13 && pi < 14.5:
			r = ClassUnderWeight
		case pi >= 14.5 && pi < 16.5:
			r = ClassNormal
		case pi >= 16.5 && pi < 18:
			r = ClassOverWight
		default:
			r = ClassSevereObesity
		}
	default:
		switch {
		case pi < 13:
			r = ClassSevereThinness
		case pi >= 13 && pi < 14.5:
			r = ClassUnderWeight
		case pi >= 14.5 && pi < 16.5:
			r = ClassNormal
		case pi >= 16.5 && pi < 18.5:
			r = ClassOverWight
		default:
			r = ClassSevereObesity
		}
	}
	return
}

func (k *Kaup) Status(pi, age float64) (r string, err error) {
	c, err := k.Classification(pi, age)
	if err != nil {
		err = fmt.Errorf("failt to get classification : %v", err)
	}
	return k.StatusByClassification(c)
}

func (k *Kaup) StatusByClassification(c string) (r string, err error) {
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
