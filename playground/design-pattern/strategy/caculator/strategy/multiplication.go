package strategy

func NewMultiplication(in InputTwoValues) Multiplication {
	return Multiplication{in}
}

// Multiplication is the multiplication strategy
type Multiplication struct {
	InputTwoValues
}

func (m Multiplication) Caculate() int {
	return m.Value1 * m.Value2
}
