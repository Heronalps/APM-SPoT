package main

import (
	"fmt"
	"math"
)
type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string{
	if (e < 0){
		return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
	} else {
		return fmt.Sprintf("%v", float64(e))
	}
}


func Sqrt(x float64) error {
	if (x < 0) {
		err := ErrNegativeSqrt(x)
		return err
	}
	delta := 1e-9
	current := x
	last := 0.0
	for math.Abs(current - last) > delta {
		last = current
		current = current - (current * current - x) / (2 * current)
	}
	return ErrNegativeSqrt(current)
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
