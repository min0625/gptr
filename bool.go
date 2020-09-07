package gptr

func NewBool(v bool) *bool {
	return &v
}

func ToBool(p *bool, d bool) bool {
	if p != nil {
		return *p
	}
	return d
}
