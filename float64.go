package gptr

func NewFloat64(v float64) *float64 {
	return &v
}

func ToFloat64(p *float64, d float64) float64 {
	if p != nil {
		return *p
	}
	return d
}
