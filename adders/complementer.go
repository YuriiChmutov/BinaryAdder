package adders

type Complementer struct {
	Inversion bool
}

func NewComplementer(inversion bool) *Complementer {
	return &Complementer{inversion}
}
