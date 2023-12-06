package pair

import "fmt"

func ExampleNew() {
	pair := New(1, 2)
	fmt.Println(pair)
	// Output: (1, 2)
}

func ExampleEqual() {
	pair1 := New(1, 2)
	pair2 := New(1, 2)
	fmt.Println(Equal(pair1, pair2))
	// Output: true
}
