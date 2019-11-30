package calc

import "math"

type PonderalIndex interface {
	Calc(weight, height float64) (result float64, err error)
}

// Result is a ponderal index calculation result
type Result struct {
	Classification string
	Index          float64
	IndexType      string
	Status         string
}

const (
	ClassNormal         = "Normal"
	ClassSevereThinness = "Severe thinness"
	ClassUnderWeight    = "Underweight"
	ClassOverWight      = "Overweight"
	ClassObesity        = "Obesity"
	ClassSevereObesity  = "Severe obesity"
	StatusNormal        = "Normal"
	StatusWarn          = "Warning"
	StatusFatal         = "Fatal"
	TypeBMI             = "Body Mass Index"
	TypeKaup            = "Kaup Index"
	TypeRohrer          = "Rohrer Index"
)

func Run(w, h, age float64) (r Result, err error) {
	switch {
	case age < 6:
		r, err = CalcKaup(w, h, age)
	case age >= 6 && age < 15:
		r, err = CalcRohrer(w, h)
	default:
		r, err = CalcBMI(w, h)
	}
	return
}

// Round returns nearest float value
func Round(val float64, place int) float64 {
	shift := math.Pow(10, float64(place))
	return math.Floor(val*shift+.5) / shift
}
