package pair

import (
	"testing"
)

func TestPairIntBool(t *testing.T) {
	tests := []struct {
		name     string
		p        *Pair[int, bool]
		expected string
	}{
		{
			name:     "Test with integer and true",
			p:        New(10, true),
			expected: "(10, true)",
		},
		{
			name:     "Test with integer and false",
			p:        New(10, false),
			expected: "(10, false)",
		},
		{
			name:     "Test with zero integer and true",
			p:        New(0, true),
			expected: "(0, true)",
		},
		{
			name:     "Test with zero integer and false",
			p:        New(0, false),
			expected: "(0, false)",
		},
		{
			name:     "Test with negative integer and true",
			p:        New(-10, true),
			expected: "(-10, true)",
		},
		{
			name:     "Test with negative integer and false",
			p:        New(-10, false),
			expected: "(-10, false)",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.p.String(); got != tt.expected {
				t.Errorf("String returned %s, expected %s", got, tt.expected)
			}
		})
	}
}

func TestPairIntString(t *testing.T) {
	tests := []struct {
		name     string
		p        *Pair[int, string]
		expected string
	}{
		{
			name:     "Test with integer and string",
			p:        New(10, "hello"),
			expected: "(10, hello)",
		},
		{
			name:     "Test with negative integer and string",
			p:        New(-10, "hello"),
			expected: "(-10, hello)",
		},
		{
			name:     "Test with zero integer and string",
			p:        New(0, "hello"),
			expected: "(0, hello)",
		},
		{
			name:     "Test with integer and empty string",
			p:        New(10, ""),
			expected: "(10, )",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.p.String(); got != tt.expected {
				t.Errorf("String() returned %s, expected %s", got, tt.expected)
			}
		})
	}
}

func TestPairIntInt(t *testing.T) {
	tests := []struct {
		name     string
		p        *Pair[int, int]
		expected string
	}{
		{
			name:     "Test with integer and integer",
			p:        New(10, 20),
			expected: "(10, 20)",
		},
		{
			name:     "Test with negative integer and integer",
			p:        New(-10, -20),
			expected: "(-10, -20)",
		},
		{
			name:     "Test with zero integer and integer",
			p:        New(0, 0),
			expected: "(0, 0)",
		},
		{
			name:     "Test with integer and zero integer",
			p:        New(10, 0),
			expected: "(10, 0)",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.p.String(); got != tt.expected {
				t.Errorf("String() returned %s, expected %s", got, tt.expected)
			}
		})
	}
}

func TestPairIntFloat(t *testing.T) {
	tests := []struct {
		name     string
		p        *Pair[int, float64]
		expected string
	}{
		{
			name:     "Test with integer and float",
			p:        New(10, 10.5),
			expected: "(10, 10.5)",
		},
		{
			name:     "Test with negative integer and float",
			p:        New(-10, -10.5),
			expected: "(-10, -10.5)",
		},
		{
			name:     "Test with zero integer and float",
			p:        New(0, 0.0),
			expected: "(0, 0)",
		},
		{
			name:     "Test with integer and zero float",
			p:        New(10, 0.0),
			expected: "(10, 0)",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.p.String(); got != tt.expected {
				t.Errorf("String() returned %s, expected %s", got, tt.expected)
			}
		})
	}
}

func TestPairFloatBool(t *testing.T) {
	tests := []struct {
		name     string
		p        *Pair[float64, bool]
		expected string
	}{
		{
			name:     "Test with float and true",
			p:        New(10.5, true),
			expected: "(10.5, true)",
		},
		{
			name:     "Test with float and false",
			p:        New(10.5, false),
			expected: "(10.5, false)",
		},
		{
			name:     "Test with zero float and true",
			p:        New(0.0, true),
			expected: "(0, true)",
		},
		{
			name:     "Test with zero float and false",
			p:        New(0.0, false),
			expected: "(0, false)",
		},
		{
			name:     "Test with negative float and true",
			p:        New(-10.5, true),
			expected: "(-10.5, true)",
		},
		{
			name:     "Test with negative float and false",
			p:        New(-10.5, false),
			expected: "(-10.5, false)",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.p.String(); tt.p.String() != tt.expected {
				t.Errorf("String() returned %s, expected %s", got, tt.expected)
			}
		})
	}
}

func TestPairFloatString(t *testing.T) {
	tests := []struct {
		name     string
		p        *Pair[float64, string]
		expected string
	}{
		{
			name:     "Test with float and string",
			p:        New(10.5, "hello"),
			expected: "(10.5, hello)",
		},
		{
			name:     "Test with negative float and string",
			p:        New(-10.5, "hello"),
			expected: "(-10.5, hello)",
		},
		{
			name:     "Test with zero float and string",
			p:        New(0.0, "hello"),
			expected: "(0, hello)",
		},
		{
			name:     "Test with float and empty string",
			p:        New(10.5, ""),
			expected: "(10.5, )",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.p.String(); got != tt.expected {
				t.Errorf("String() returned %s, expected %s", got, tt.expected)
			}
		})
	}
}

func TestPairFloatInt(t *testing.T) {
	tests := []struct {
		name     string
		p        *Pair[float64, int]
		expected string
	}{
		{
			name:     "Test with float and integer",
			p:        New(10.5, 10),
			expected: "(10.5, 10)",
		},
		{
			name:     "Test with negative float and integer",
			p:        New(-10.5, -10),
			expected: "(-10.5, -10)",
		},
		{
			name:     "Test with zero float and integer",
			p:        New(0.0, 0),
			expected: "(0, 0)",
		},
		{
			name:     "Test with float and zero integer",
			p:        New(10.5, 0),
			expected: "(10.5, 0)",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.p.String(); got != tt.expected {
				t.Errorf("String() returned %s, expected %s", got, tt.expected)
			}
		})
	}
}

func TestPairFloatFloat(t *testing.T) {
	tests := []struct {
		name     string
		p        *Pair[float64, float64]
		expected string
	}{
		{
			name:     "Test with float and float",
			p:        New(10.5, 10.5),
			expected: "(10.5, 10.5)",
		},
		{
			name:     "Test with negative float and float",
			p:        New(-10.5, -10.5),
			expected: "(-10.5, -10.5)",
		},
		{
			name:     "Test with zero float and float",
			p:        New(0.0, 0.0),
			expected: "(0, 0)",
		},
		{
			name:     "Test with float and zero float",
			p:        New(10.5, 0.0),
			expected: "(10.5, 0)",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.p.String(); got != tt.expected {
				t.Errorf("String() returned %s, expected %s", got, tt.expected)
			}
		})
	}
}

func TestPairStringBool(t *testing.T) {
	tests := []struct {
		name     string
		p        *Pair[string, bool]
		expected string
	}{
		{
			name:     "Test with string and true",
			p:        New("hello", true),
			expected: "(hello, true)",
		},
		{
			name:     "Test with string and false",
			p:        New("hello", false),
			expected: "(hello, false)",
		},
		{
			name:     "Test with empty string and true",
			p:        New("", true),
			expected: "(, true)",
		},
		{
			name:     "Test with empty string and false",
			p:        New("", false),
			expected: "(, false)",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.p.String(); tt.p.String() != tt.expected {
				t.Errorf("String() returned %s, expected %s", got, tt.expected)
			}
		})
	}
}

func TestPairStringString(t *testing.T) {
	tests := []struct {
		name     string
		p        *Pair[string, string]
		expected string
	}{
		{
			name:     "Test with string and string",
			p:        New("hello", "world"),
			expected: "(hello, world)",
		},
		{
			name:     "Test with empty string and string",
			p:        New("", "world"),
			expected: "(, world)",
		},
		{
			name:     "Test with string and empty string",
			p:        New("hello", ""),
			expected: "(hello, )",
		},
		{
			name:     "Test with empty string and empty string",
			p:        New("", ""),
			expected: "(, )",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.p.String(); got != tt.expected {
				t.Errorf("String() returned %s, expected %s", got, tt.expected)
			}
		})
	}
}

func TestPairStringInt(t *testing.T) {
	tests := []struct {
		name     string
		p        *Pair[string, int]
		expected string
	}{
		{
			name:     "Test with string and integer",
			p:        New("hello", 10),
			expected: "(hello, 10)",
		},
		{
			name:     "Test with empty string and integer",
			p:        New("", 10),
			expected: "(, 10)",
		},
		{
			name:     "Test with string and zero integer",
			p:        New("hello", 0),
			expected: "(hello, 0)",
		},
		{
			name:     "Test with empty string and zero integer",
			p:        New("", 0),
			expected: "(, 0)",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.p.String(); got != tt.expected {
				t.Errorf("String() returned %s, expected %s", got, tt.expected)
			}
		})
	}
}

func TestPairStringFloat(t *testing.T) {
	tests := []struct {
		name     string
		p        *Pair[string, float64]
		expected string
	}{
		{
			name:     "Test with string and float",
			p:        New("hello", 10.5),
			expected: "(hello, 10.5)",
		},
		{
			name:     "Test with empty string and float",
			p:        New("", 10.5),
			expected: "(, 10.5)",
		},
		{
			name:     "Test with string and zero float",
			p:        New("hello", 0.0),
			expected: "(hello, 0)",
		},
		{
			name:     "Test with empty string and zero float",
			p:        New("", 0.0),
			expected: "(, 0)",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.p.String(); got != tt.expected {
				t.Errorf("String() returned %s, expected %s", got, tt.expected)
			}
		})
	}
}

func TestPairBoolString(t *testing.T) {
	tests := []struct {
		name     string
		p        *Pair[bool, string]
		expected string
	}{
		{
			name:     "Test with true and string",
			p:        New(true, "hello"),
			expected: "(true, hello)",
		},
		{
			name:     "Test with false and string",
			p:        New(false, "hello"),
			expected: "(false, hello)",
		},
		{
			name:     "Test with true and empty string",
			p:        New(true, ""),
			expected: "(true, )",
		},
		{
			name:     "Test with false and empty string",
			p:        New(false, ""),
			expected: "(false, )",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.p.String(); got != tt.expected {
				t.Errorf("String() returned %s, expected %s", got, tt.expected)
			}
		})
	}
}

func TestPairBoolInt(t *testing.T) {
	tests := []struct {
		name     string
		p        *Pair[bool, int]
		expected string
	}{
		{
			name:     "Test with true and integer",
			p:        New(true, 10),
			expected: "(true, 10)",
		},
		{
			name:     "Test with false and integer",
			p:        New(false, 10),
			expected: "(false, 10)",
		},
		{
			name:     "Test with true and zero integer",
			p:        New(true, 0),
			expected: "(true, 0)",
		},
		{
			name:     "Test with false and zero integer",
			p:        New(false, 0),
			expected: "(false, 0)",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.p.String(); got != tt.expected {
				t.Errorf("String() returned %s, expected %s", got, tt.expected)
			}
		})
	}
}

func TestPairBoolFloat(t *testing.T) {
	tests := []struct {
		name     string
		p        *Pair[bool, float64]
		expected string
	}{
		{
			name:     "Test with true and float",
			p:        New(true, 10.5),
			expected: "(true, 10.5)",
		},
		{
			name:     "Test with false and float",
			p:        New(false, 10.5),
			expected: "(false, 10.5)",
		},
		{
			name:     "Test with true and zero float",
			p:        New(true, 0.0),
			expected: "(true, 0)",
		},
		{
			name:     "Test with false and zero float",
			p:        New(false, 0.0),
			expected: "(false, 0)",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.p.String(); got != tt.expected {
				t.Errorf("String() returned %s, expected %s", got, tt.expected)
			}
		})
	}
}

func TestPairEqualIntString(t *testing.T) {
	tests := []struct {
		name     string
		p1       *Pair[int, string]
		p2       *Pair[int, string]
		expected bool
	}{
		{
			name:     "Test with equal pairs",
			p1:       New(10, "hello"),
			p2:       New(10, "hello"),
			expected: true,
		},
		{
			name:     "Test with different pairs",
			p1:       New(10, "hello"),
			p2:       New(20, "world"),
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Equal(tt.p1, tt.p2); got != tt.expected {
				t.Errorf("Equal() returned %v, expected %v", got, tt.expected)
			}
		})
	}
}

func TestPairEqualIntInt(t *testing.T) {
	tests := []struct {
		name     string
		p1       *Pair[int, int]
		p2       *Pair[int, int]
		expected bool
	}{
		{
			name:     "Test with equal pairs",
			p1:       New(10, 20),
			p2:       New(10, 20),
			expected: true,
		},
		{
			name:     "Test with different pairs",
			p1:       New(10, 20),
			p2:       New(20, 10),
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Equal(tt.p1, tt.p2); got != tt.expected {
				t.Errorf("Equal() returned %v, expected %v", got, tt.expected)
			}
		})
	}
}

func TestPairEqualIntFloat(t *testing.T) {
	tests := []struct {
		name     string
		p1       *Pair[int, float64]
		p2       *Pair[int, float64]
		expected bool
	}{
		{
			name:     "Test with equal pairs",
			p1:       New(10, 10.5),
			p2:       New(10, 10.5),
			expected: true,
		},
		{
			name:     "Test with different pairs",
			p1:       New(10, 10.5),
			p2:       New(20, 20.5),
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Equal(tt.p1, tt.p2); got != tt.expected {
				t.Errorf("Equal() returned %v, expected %v", got, tt.expected)
			}
		})
	}
}

func TestPairEqualFloatString(t *testing.T) {
	tests := []struct {
		name     string
		p1       *Pair[float64, string]
		p2       *Pair[float64, string]
		expected bool
	}{
		{
			name:     "Test with equal pairs",
			p1:       New(10.5, "hello"),
			p2:       New(10.5, "hello"),
			expected: true,
		},
		{
			name:     "Test with different pairs",
			p1:       New(10.5, "hello"),
			p2:       New(20.5, "world"),
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Equal(tt.p1, tt.p2); got != tt.expected {
				t.Errorf("Equal() returned %v, expected %v", got, tt.expected)
			}
		})
	}
}

func TestPairEqualFloatInt(t *testing.T) {
	tests := []struct {
		name     string
		p1       *Pair[float64, int]
		p2       *Pair[float64, int]
		expected bool
	}{
		{
			name:     "Test with equal pairs",
			p1:       New(10.5, 10),
			p2:       New(10.5, 10),
			expected: true,
		},
		{
			name:     "Test with different pairs",
			p1:       New(10.5, 10),
			p2:       New(20.5, 20),
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Equal(tt.p1, tt.p2); got != tt.expected {
				t.Errorf("Equal() returned %v, expected %v", got, tt.expected)
			}
		})
	}
}

func TestPairEqualFloatFloat(t *testing.T) {
	tests := []struct {
		name     string
		p1       *Pair[float64, float64]
		p2       *Pair[float64, float64]
		expected bool
	}{
		{
			name:     "Test with equal pairs",
			p1:       New(10.5, 10.5),
			p2:       New(10.5, 10.5),
			expected: true,
		},
		{
			name:     "Test with different pairs",
			p1:       New(10.5, 10.5),
			p2:       New(20.5, 20.5),
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Equal(tt.p1, tt.p2); got != tt.expected {
				t.Errorf("Equal() returned %v, expected %v", got, tt.expected)
			}
		})
	}
}

func TestPairEqualStringString(t *testing.T) {
	tests := []struct {
		name     string
		p1       *Pair[string, string]
		p2       *Pair[string, string]
		expected bool
	}{
		{
			name:     "Test with equal pairs",
			p1:       New("hello", "world"),
			p2:       New("hello", "world"),
			expected: true,
		},
		{
			name:     "Test with different pairs",
			p1:       New("hello", "world"),
			p2:       New("world", "hello"),
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Equal(tt.p1, tt.p2); got != tt.expected {
				t.Errorf("Equal() returned %v, expected %v", got, tt.expected)
			}
		})
	}
}

func TestPairEqualStringInt(t *testing.T) {
	tests := []struct {
		name     string
		p1       *Pair[string, int]
		p2       *Pair[string, int]
		expected bool
	}{
		{
			name:     "Test with equal pairs",
			p1:       New("hello", 10),
			p2:       New("hello", 10),
			expected: true,
		},
		{
			name:     "Test with different pairs",
			p1:       New("hello", 10),
			p2:       New("world", 20),
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Equal(tt.p1, tt.p2); got != tt.expected {
				t.Errorf("Equal() returned %v, expected %v", got, tt.expected)
			}
		})
	}
}

func TestPairEqualStringFloat(t *testing.T) {
	tests := []struct {
		name     string
		p1       *Pair[string, float64]
		p2       *Pair[string, float64]
		expected bool
	}{
		{
			name:     "Test with equal pairs",
			p1:       New("hello", 10.5),
			p2:       New("hello", 10.5),
			expected: true,
		},
		{
			name:     "Test with different pairs",
			p1:       New("hello", 10.5),
			p2:       New("world", 20.5),
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Equal(tt.p1, tt.p2); got != tt.expected {
				t.Errorf("Equal() returned %v, expected %v", got, tt.expected)
			}
		})
	}
}

func TestPairEqualBoolString(t *testing.T) {
	tests := []struct {
		name     string
		p1       *Pair[bool, string]
		p2       *Pair[bool, string]
		expected bool
	}{
		{
			name:     "Test with equal pairs",
			p1:       New(true, "hello"),
			p2:       New(true, "hello"),
			expected: true,
		},
		{
			name:     "Test with different pairs",
			p1:       New(true, "hello"),
			p2:       New(false, "world"),
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Equal(tt.p1, tt.p2); got != tt.expected {
				t.Errorf("Equal() returned %v, expected %v", got, tt.expected)
			}
		})
	}
}

func TestPairEqualBoolInt(t *testing.T) {
	tests := []struct {
		name     string
		p1       *Pair[bool, int]
		p2       *Pair[bool, int]
		expected bool
	}{
		{
			name:     "Test with equal pairs",
			p1:       New(true, 10),
			p2:       New(true, 10),
			expected: true,
		},
		{
			name:     "Test with different pairs",
			p1:       New(true, 10),
			p2:       New(false, 20),
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Equal(tt.p1, tt.p2); got != tt.expected {
				t.Errorf("Equal() returned %v, expected %v", got, tt.expected)
			}
		})
	}
}

func TestPairEqualBoolFloat(t *testing.T) {
	tests := []struct {
		name     string
		p1       *Pair[bool, float64]
		p2       *Pair[bool, float64]
		expected bool
	}{
		{
			name:     "Test with equal pairs",
			p1:       New(true, 10.5),
			p2:       New(true, 10.5),
			expected: true,
		},
		{
			name:     "Test with different pairs",
			p1:       New(true, 10.5),
			p2:       New(false, 20.5),
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Equal(tt.p1, tt.p2); got != tt.expected {
				t.Errorf("Equal() returned %v, expected %v", got, tt.expected)
			}
		})
	}
}

func TestPairEqualBoolBool(t *testing.T) {
	tests := []struct {
		name     string
		p1       *Pair[bool, bool]
		p2       *Pair[bool, bool]
		expected bool
	}{
		{
			name:     "Test with equal pairs",
			p1:       New(true, true),
			p2:       New(true, true),
			expected: true,
		},
		{
			name:     "Test with different pairs",
			p1:       New(true, true),
			p2:       New(false, false),
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Equal(tt.p1, tt.p2); got != tt.expected {
				t.Errorf("Equal() returned %v, expected %v", got, tt.expected)
			}
		})
	}
}
