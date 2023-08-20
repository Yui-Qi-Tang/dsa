package strategy

func NewAddStrategy(in InputTwoValues) Add {
	return Add{in}
}

// Add is the addtion strategy
type Add struct {
	InputTwoValues
}

func (a Add) Caculate() int {
	return a.Value1 + a.Value2
}
