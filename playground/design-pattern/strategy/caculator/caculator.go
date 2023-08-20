package caculator

// Strategy represents the strategy interface
type Strategy interface {
	// Strategy represents the method for caculating
	Caculate() int
}

func New() *Caculator {
	return &Caculator{}
}

// Caculator is the context that use a strategy
type Caculator struct {
	stratey Strategy
}

// Set sets the strategy
func (c *Caculator) Set(s Strategy) {
	c.stratey = s
}

func (c *Caculator) Caculate() int {
	return c.stratey.Caculate()
}
