package gptr

func NewInt(v int) *int {
	return &v
}

func ToInt(p *int, d int) int {
	if p != nil {
		return *p
	}
	return d
}
