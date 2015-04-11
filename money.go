package money

import (
	"fmt"
	"math/big"

	"github.com/juju/errors"
)

// Money represents a monetary value
type Money struct {
	rat *big.Rat
}

// New creates a new instance with a zero value.
func New() *Money {
	return &Money{
		rat: big.NewRat(0, 1),
	}
}

// Parse a string to create a new money value.
func Parse(s string) (*Money, error) {
	rat := new(big.Rat)
	if _, err := fmt.Sscan(s, rat); err != nil {
		return nil, errors.Trace(err)
	}

	return &Money{rat}, nil
}

// Format the money value with a specific deciaml precision.
func (money *Money) Format(prec int) string {
	return money.rat.FloatString(prec)
}

// Mul multiplies the money value n times and returns the result.
func (money *Money) Mul(n int64) *Money {
	b := big.NewRat(n, 1)
	result := New()
	result.rat.Mul(money.rat, b)
	return result
}

// Add two money values together and returns the result.
func (money *Money) Add(other *Money) *Money {
	result := New()
	result.rat.Add(money.rat, other.rat)
	return result
}

// Sub subtracts two money values and returns the result.
func (money *Money) Sub(other *Money) *Money {
	result := New()
	result.rat.Sub(money.rat, other.rat)
	return result
}

// Div divides two money values and returns the result.
func (money *Money) Div(other *Money) *Money {
	result := New()
	result.rat.Quo(money.rat, other.rat)
	return result
}

// LessThan returns true if a money value is less than the other.
func (money *Money) LessThan(other *Money) bool {
	return money.rat.Cmp(other.rat) == -1
}
