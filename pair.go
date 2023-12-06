// package pair is a package that defines a pair of values.
package pair

import "fmt"

// Pair is an interface that defines a pair of values.
type Pair interface {
	String() string
	First() interface{}
	Second() interface{}
}

// intPair is a struct that implements the Pair interface.
type intPair struct {
	p1 int
	p2 int
}

// String returns a string representation of the intPair.
func (p *intPair) String() string {
	return fmt.Sprintf("(%d, %d)", p.p1, p.p2)
}

// First returns the first value of the intPair.
func (p *intPair) First() interface{} {
	return p.p1
}

// Second returns the second value of the intPair.
func (p *intPair) Second() interface{} {
	return p.p2
}

// int32Pair is a struct that implements the Pair interface.
type int32Pair struct {
	p1 int32
	p2 int32
}

// String returns a string representation of the int32Pair.
func (p *int32Pair) String() string {
	return fmt.Sprintf("(%d, %d)", p.p1, p.p2)
}

// First returns the first value of the int32Pair.
func (p *int32Pair) First() interface{} {
	return p.p1
}

// Second returns the second value of the int32Pair.
func (p *int32Pair) Second() interface{} {
	return p.p2
}

// int64Pair is a struct that implements the Pair interface.
type int64Pair struct {
	p1 int64
	p2 int64
}

// String returns a string representation of the int64Pair.
func (p *int64Pair) String() string {
	return fmt.Sprintf("(%d, %d)", p.p1, p.p2)
}

// First returns the first value of the int64Pair.
func (p *int64Pair) First() interface{} {
	return p.p1
}

// Second returns the second value of the int64Pair.
func (p *int64Pair) Second() interface{} {
	return p.p2
}

// uintPair is a struct that implements the Pair interface.
type uintPair struct {
	p1 uint
	p2 uint
}

// String returns a string representation of the uintPair.
func (p *uintPair) String() string {
	return fmt.Sprintf("(%d, %d)", p.p1, p.p2)
}

// First returns the first value of the uintPair.
func (p *uintPair) First() interface{} {
	return p.p1
}

// Second returns the second value of the uintPair.
func (p *uintPair) Second() interface{} {
	return p.p2
}

// uint32Pair is a struct that implements the Pair interface.
type uint32Pair struct {
	p1 uint32
	p2 uint32
}

// String returns a string representation of the uint32Pair.
func (p *uint32Pair) String() string {
	return fmt.Sprintf("(%d, %d)", p.p1, p.p2)
}

// First returns the first value of the uint32Pair.
func (p *uint32Pair) First() interface{} {
	return p.p1
}

// Second returns the second value of the uint32Pair.
func (p *uint32Pair) Second() interface{} {
	return p.p2
}

// uint64Pair is a struct that implements the Pair interface.
type uint64Pair struct {
	p1 uint64
	p2 uint64
}

// String returns a string representation of the uint64Pair.
func (p *uint64Pair) String() string {
	return fmt.Sprintf("(%d, %d)", p.p1, p.p2)
}

// First returns the first value of the uint64Pair.
func (p *uint64Pair) First() interface{} {
	return p.p1
}

// Second returns the second value of the uint64Pair.
func (p *uint64Pair) Second() interface{} {
	return p.p2
}

// float32Pair is a struct that implements the Pair interface.
type float32Pair struct {
	p1 float32
	p2 float32
}

// String returns a string representation of the float32Pair.
func (p *float32Pair) String() string {
	return fmt.Sprintf("(%f, %f)", p.p1, p.p2)
}

// First returns the first value of the float32Pair.
func (p *float32Pair) First() interface{} {
	return p.p1
}

// Second returns the second value of the float32Pair.
func (p *float32Pair) Second() interface{} {
	return p.p2
}

// float64Pair is a struct that implements the Pair interface.
type float64Pair struct {
	p1 float64
	p2 float64
}

// String returns a string representation of the float64Pair.
func (p *float64Pair) String() string {
	return fmt.Sprintf("(%f, %f)", p.p1, p.p2)
}

// First returns the first value of the float64Pair.
func (p *float64Pair) First() interface{} {
	return p.p1
}

// Second returns the second value of the float64Pair.
func (p *float64Pair) Second() interface{} {
	return p.p2
}

// stringPair is a struct that implements the Pair interface.
type stringPair struct {
	p1 string
	p2 string
}

// String returns a string representation of the stringPair.
func (p *stringPair) String() string {
	return fmt.Sprintf("(%s, %s)", p.p1, p.p2)
}

// First returns the first value of the stringPair.
func (p *stringPair) First() interface{} {
	return p.p1
}

// Second returns the second value of the stringPair.
func (p *stringPair) Second() interface{} {
	return p.p2
}

// boolPair is a struct that implements the Pair interface.
type boolPair struct {
	p1 bool
	p2 bool
}

// String returns a string representation of the boolPair.
func (p *boolPair) String() string {
	return fmt.Sprintf("(%t, %t)", p.p1, p.p2)
}

// First returns the first value of the boolPair.
func (p *boolPair) First() interface{} {
	return p.p1
}

// Second returns the second value of the boolPair.
func (p *boolPair) Second() interface{} {
	return p.p2
}

// New returns a new Pair.
func New(p1, p2 interface{}) Pair {
	switch v1 := p1.(type) {
	case uint:
		v2 := p2.(uint)
		return &uintPair{p1: v1, p2: v2}
	case uint32:
		v2 := p2.(uint32)
		return &uint32Pair{p1: v1, p2: v2}
	case uint64:
		v2 := p2.(uint64)
		return &uint64Pair{p1: v1, p2: v2}
	case int:
		v2 := p2.(int)
		return &intPair{p1: v1, p2: v2}
	case int32:
		v2 := p2.(int32)
		return &int32Pair{p1: v1, p2: v2}
	case int64:
		v2 := p2.(int64)
		return &int64Pair{p1: v1, p2: v2}
	case float32:
		v2 := p2.(float32)
		return &float32Pair{p1: v1, p2: v2}
	case float64:
		v2 := p2.(float64)
		return &float64Pair{p1: v1, p2: v2}
	case string:
		v2 := p2.(string)
		return &stringPair{p1: v1, p2: v2}
	case bool:
		v2 := p2.(bool)
		return &boolPair{p1: v1, p2: v2}
	}
	return nil
}

// Equal returns true if the two Pairs are equal.
func Equal(p1, p2 Pair) bool {
	switch v1 := p1.(type) {
	case *uintPair:
		v2 := p2.(*uintPair)
		return v1.p1 == v2.p1 && v1.p2 == v2.p2
	case *uint32Pair:
		v2 := p2.(*uint32Pair)
		return v1.p1 == v2.p1 && v1.p2 == v2.p2
	case *uint64Pair:
		v2 := p2.(*uint64Pair)
		return v1.p1 == v2.p1 && v1.p2 == v2.p2
	case *intPair:
		v2 := p2.(*intPair)
		return v1.p1 == v2.p1 && v1.p2 == v2.p2
	case *int32Pair:
		v2 := p2.(*int32Pair)
		return v1.p1 == v2.p1 && v1.p2 == v2.p2
	case *int64Pair:
		v2 := p2.(*int64Pair)
		return v1.p1 == v2.p1 && v1.p2 == v2.p2
	case *float32Pair:
		v2 := p2.(*float32Pair)
		return v1.p1 == v2.p1 && v1.p2 == v2.p2
	case *float64Pair:
		v2 := p2.(*float64Pair)
		return v1.p1 == v2.p1 && v1.p2 == v2.p2
	case *stringPair:
		v2 := p2.(*stringPair)
		return v1.p1 == v2.p1 && v1.p2 == v2.p2
	case *boolPair:
		v2 := p2.(*boolPair)
		return v1.p1 == v2.p1 && v1.p2 == v2.p2
	}
	return false
}
