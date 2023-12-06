package pair

import (
	"fmt"
	"testing"
)

func TestPairString(t *testing.T) {
	tests := []struct {
		name string
		p    Pair
		want string
	}{
		{
			name: "Test intPair",
			p:    &intPair{p1: 10, p2: 20},
			want: "(10, 20)",
		},
		{
			name: "Test int32Pair",
			p:    &int32Pair{p1: 30, p2: 40},
			want: "(30, 40)",
		},
		{
			name: "Test int64Pair",
			p:    &int64Pair{p1: 50, p2: 60},
			want: "(50, 60)",
		},
		{
			name: "Test uintPair",
			p:    &uintPair{p1: 70, p2: 80},
			want: "(70, 80)",
		},
		{
			name: "Test uint32Pair",
			p:    &uint32Pair{p1: 90, p2: 100},
			want: "(90, 100)",
		},
		{
			name: "Test uint64Pair",
			p:    &uint64Pair{p1: 110, p2: 120},
			want: "(110, 120)",
		},
		{
			name: "Test float32Pair",
			p:    &float32Pair{p1: 1.23, p2: 4.56},
			want: fmt.Sprintf("(%f, %f)", float32(1.23), float32(4.56)),
		},
		{
			name: "Test float64Pair",
			p:    &float64Pair{p1: 7.89, p2: 10.11},
			want: fmt.Sprintf("(%f, %f)", float64(7.89), float64(10.11)),
		},
		{
			name: "Test stringPair",
			p:    &stringPair{p1: "hello", p2: "world"},
			want: "(hello, world)",
		},
		{
			name: "Test boolPair",
			p:    &boolPair{p1: true, p2: false},
			want: "(true, false)",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.p.String()
			if got != tt.want {
				t.Errorf("got %s, want %s", got, tt.want)
			}
		})
	}
}

func TestNew(t *testing.T) {
	tests := []struct {
		name string
		p1   interface{}
		p2   interface{}
		want Pair
	}{
		{
			name: "Test intPair",
			p1:   10,
			p2:   20,
			want: &intPair{p1: 10, p2: 20},
		},
		{
			name: "Test stringPair",
			p1:   "hello",
			p2:   "world",
			want: &stringPair{p1: "hello", p2: "world"},
		},
		{
			name: "Test int32Pair",
			p1:   int32(30),
			p2:   int32(40),
			want: &int32Pair{p1: 30, p2: 40},
		},
		{
			name: "Test int64Pair",
			p1:   int64(50),
			p2:   int64(60),
			want: &int64Pair{p1: 50, p2: 60},
		},
		{
			name: "Test uintPair",
			p1:   uint(70),
			p2:   uint(80),
			want: &uintPair{p1: 70, p2: 80},
		},
		{
			name: "Test uint32Pair",
			p1:   uint32(90),
			p2:   uint32(100),
			want: &uint32Pair{p1: 90, p2: 100},
		},
		{
			name: "Test uint64Pair",
			p1:   uint64(110),
			p2:   uint64(120),
			want: &uint64Pair{p1: 110, p2: 120},
		},
		{
			name: "Test float32Pair",
			p1:   float32(1.23),
			p2:   float32(4.56),
			want: &float32Pair{p1: 1.23, p2: 4.56},
		},
		{
			name: "Test float64Pair",
			p1:   float64(7.89),
			p2:   float64(10.11),
			want: &float64Pair{p1: 7.89, p2: 10.11},
		},
		{
			name: "Test boolPair",
			p1:   true,
			p2:   false,
			want: &boolPair{p1: true, p2: false},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := New(tt.p1, tt.p2)
			if fmt.Sprintf("%T", got) != fmt.Sprintf("%T", tt.want) {
				t.Errorf("got %T, want %T", got, tt.want)
			}
		})
	}
}

func TestNewWithInvalidTypes(t *testing.T) {
	tests := []struct {
		name string
		p1   interface{}
		p2   interface{}
		want Pair
	}{
		{
			name: "Test int32Pair and intPair",
			p1:   int32(10),
			p2:   20,
			want: nil,
		},
		{
			name: "Test intPair",
			p1:   10,
			p2:   "hello",
			want: nil,
		},
		{
			name: "Test stringPair",
			p1:   "hello",
			p2:   20,
			want: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				if r := recover(); r == nil {
					t.Errorf("should have panicked")
				}
			}()
			got := New(tt.p1, tt.p2)
			if fmt.Sprintf("%T", got) != fmt.Sprintf("%T", tt.want) {
				t.Errorf("got %T, want %T", got, tt.want)
			}
		})
	}
}
func TestFirst(t *testing.T) {
	tests := []struct {
		name     string
		p        Pair
		expected interface{}
	}{
		{
			name:     "Test intPair First",
			p:        &intPair{p1: 10, p2: 20},
			expected: 10,
		},
		{
			name:     "Test stringPair First",
			p:        &stringPair{p1: "hello", p2: "world"},
			expected: "hello",
		},
		{
			name:     "Test int32Pair First",
			p:        &int32Pair{p1: 30, p2: 40},
			expected: int32(30),
		},
		{
			name:     "Test int64Pair First",
			p:        &int64Pair{p1: 50, p2: 60},
			expected: int64(50),
		},
		{
			name:     "Test uintPair First",
			p:        &uintPair{p1: 70, p2: 80},
			expected: uint(70),
		},
		{
			name:     "Test uint32Pair First",
			p:        &uint32Pair{p1: 90, p2: 100},
			expected: uint32(90),
		},
		{
			name:     "Test uint64Pair First",
			p:        &uint64Pair{p1: 110, p2: 120},
			expected: uint64(110),
		},
		{
			name:     "Test float32Pair First",
			p:        &float32Pair{p1: 1.23, p2: 4.56},
			expected: float32(1.23),
		},
		{
			name:     "Test float64Pair First",
			p:        &float64Pair{p1: 7.89, p2: 10.11},
			expected: float64(7.89),
		},
		{
			name:     "Test boolPair First",
			p:        &boolPair{p1: true, p2: false},
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.p.First()
			if got != tt.expected {
				t.Errorf("got %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestSecond(t *testing.T) {
	tests := []struct {
		name     string
		p        Pair
		expected interface{}
	}{
		{
			name:     "Test intPair Second",
			p:        &intPair{p1: 10, p2: 20},
			expected: 20,
		},
		{
			name:     "Test stringPair Second",
			p:        &stringPair{p1: "hello", p2: "world"},
			expected: "world",
		},
		{
			name:     "Test int32Pair Second",
			p:        &int32Pair{p1: 30, p2: 40},
			expected: int32(40),
		},
		{
			name:     "Test int64Pair Second",
			p:        &int64Pair{p1: 50, p2: 60},
			expected: int64(60),
		},
		{
			name:     "Test uintPair Second",
			p:        &uintPair{p1: 70, p2: 80},
			expected: uint(80),
		},
		{
			name:     "Test uint32Pair Second",
			p:        &uint32Pair{p1: 90, p2: 100},
			expected: uint32(100),
		},
		{
			name:     "Test uint64Pair Second",
			p:        &uint64Pair{p1: 110, p2: 120},
			expected: uint64(120),
		},
		{
			name:     "Test float32Pair Second",
			p:        &float32Pair{p1: 1.23, p2: 4.56},
			expected: float32(4.56),
		},
		{
			name:     "Test float64Pair Second",
			p:        &float64Pair{p1: 7.89, p2: 10.11},
			expected: float64(10.11),
		},
		{
			name:     "Test boolPair Second",
			p:        &boolPair{p1: true, p2: false},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.p.Second()
			if got != tt.expected {
				t.Errorf("got %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestEqual(t *testing.T) {
	tests := []struct {
		name string
		p1   Pair
		p2   Pair
		want bool
	}{
		{
			name: "Test intPair",
			p1:   &intPair{p1: 10, p2: 20},
			p2:   &intPair{p1: 10, p2: 20},
			want: true,
		},
		{
			name: "Test intPair with different values",
			p1:   &intPair{p1: 10, p2: 20},
			p2:   &intPair{p1: 30, p2: 40},
			want: false,
		},
		{
			name: "Test stringPair",
			p1:   &stringPair{p1: "hello", p2: "world"},
			p2:   &stringPair{p1: "hello", p2: "world"},
			want: true,
		},
		{
			name: "Test stringPair with different values",
			p1:   &stringPair{p1: "hello", p2: "world"},
			p2:   &stringPair{p1: "foo", p2: "bar"},
			want: false,
		},
		{
			name: "Test int32Pair",
			p1:   &int32Pair{p1: 30, p2: 40},
			p2:   &int32Pair{p1: 30, p2: 40},
			want: true,
		},
		{
			name: "Test int32Pair with different values",
			p1:   &int32Pair{p1: 30, p2: 40},
			p2:   &int32Pair{p1: 50, p2: 60},
			want: false,
		},
		{
			name: "Test int64Pair",
			p1:   &int64Pair{p1: 50, p2: 60},
			p2:   &int64Pair{p1: 50, p2: 60},
			want: true,
		},
		{
			name: "Test int64Pair with different values",
			p1:   &int64Pair{p1: 50, p2: 60},
			p2:   &int64Pair{p1: 70, p2: 80},
			want: false,
		},
		{
			name: "Test uintPair",
			p1:   &uintPair{p1: 70, p2: 80},
			p2:   &uintPair{p1: 70, p2: 80},
			want: true,
		},
		{
			name: "Test uintPair with different values",
			p1:   &uintPair{p1: 70, p2: 80},
			p2:   &uintPair{p1: 90, p2: 100},
			want: false,
		},
		{
			name: "Test uint32Pair",
			p1:   &uint32Pair{p1: 90, p2: 100},
			p2:   &uint32Pair{p1: 90, p2: 100},
			want: true,
		},
		{
			name: "Test uint32Pair with different values",
			p1:   &uint32Pair{p1: 90, p2: 100},
			p2:   &uint32Pair{p1: 110, p2: 120},
			want: false,
		},
		{
			name: "Test uint64Pair",
			p1:   &uint64Pair{p1: 110, p2: 120},
			p2:   &uint64Pair{p1: 110, p2: 120},
			want: true,
		},
		{
			name: "Test uint64Pair with different values",
			p1:   &uint64Pair{p1: 110, p2: 120},
			p2:   &uint64Pair{p1: 130, p2: 140},
			want: false,
		},
		{
			name: "Test float32Pair",
			p1:   &float32Pair{p1: 1.23, p2: 4.56},
			p2:   &float32Pair{p1: 1.23, p2: 4.56},
			want: true,
		},
		{
			name: "Test float32Pair with different values",
			p1:   &float32Pair{p1: 1.23, p2: 4.56},
			p2:   &float32Pair{p1: 7.89, p2: 10.11},
			want: false,
		},
		{
			name: "Test float64Pair",
			p1:   &float64Pair{p1: 7.89, p2: 10.11},
			p2:   &float64Pair{p1: 7.89, p2: 10.11},
			want: true,
		},
		{
			name: "Test float64Pair with different values",
			p1:   &float64Pair{p1: 7.89, p2: 10.11},
			p2:   &float64Pair{p1: 12.34, p2: 56.78},
			want: false,
		},
		{
			name: "Test boolPair",
			p1:   &boolPair{p1: true, p2: false},
			p2:   &boolPair{p1: true, p2: false},
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Equal(tt.p1, tt.p2)
			if got != tt.want {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}
