package main

import (
	"fmt"

	"dsa-problems.yqtang.org/playground/design-pattern/strategy/caculator"
	"dsa-problems.yqtang.org/playground/design-pattern/strategy/caculator/strategy"
)

func main() {
	c := caculator.New() // create Caculator context

	input := strategy.InputTwoValues{Value1: 100, Value2: -100}
	addStrategy := strategy.NewAddStrategy(input) // create addition strategy
	c.Set(addStrategy)                            // using the strategy
	fmt.Println(c.Caculate())                     // 0

	// using another strategy
	anotherStrategy := strategy.NewMultiplication(input)
	c.Set(anotherStrategy)
	fmt.Println(c.Caculate()) // -10000
}
