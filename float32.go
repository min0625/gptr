package gptr

func NewFloat32(v float32) *float32 {
	return &v
}

func ToFloat32(p *float32, d float32) float32 {
	if p != nil {
		return *p
	}
	return d
}
