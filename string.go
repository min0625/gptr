package gptr

func NewString(v string) *string {
	return &v
}

func ToString(p *string, d string) string {
	if p != nil {
		return *p
	}
	return d
}
