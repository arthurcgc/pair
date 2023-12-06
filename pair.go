// pair is a package that defines a generic Pair type.
package pair

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

// pairType is a constraint that permits any type that satisfies the Ordered interface and bool.
type pairType interface {
	constraints.Ordered | bool
}

// Pair is a generic pair of values.
type Pair[T1 pairType, T2 pairType | bool] struct {
	First  T1
	Second T2
}

// String returns a string representation of the Pair.
func (p *Pair[T1, T2]) String() string {
	return fmt.Sprintf("(%v, %v)", p.First, p.Second)
}

// New returns a new Pair.
func New[T1, T2 pairType](p1 T1, p2 T2) *Pair[T1, T2] {
	return &Pair[T1, T2]{First: p1, Second: p2}
}

// Equal returns true if the two Pairs are equal.
func Equal[T1, T2 pairType](pair1, pair2 *Pair[T1, T2]) bool {
	return pair1.First == pair2.First && pair1.Second == pair2.Second
}
