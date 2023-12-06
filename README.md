# Pair Package

The Pair package provides a simple way to handle pairs of different types in Go. It defines a `Pair` interface and several structs that implement this interface for different types.

## Usage

To create a new pair, use the `New` function. It takes two parameters of any type and returns a `Pair`:

```go
p := pair.New(1, 2) // returns an intPair
```

To access the values of the pair, use the `First` and `Second` methods:

```go
p := pair.New(1, 2)
fmt.Println(p.First()) // prints 1
fmt.Println(p.Second()) // prints 2
```

To check if two pairs are equal, use the `Equals` method:

```go
p1 := pair.New(1, 2)
p2 := pair.New(1, 2)
fmt.Println(p1.Equals(p2)) // prints true
```

Each pair type has a String method that returns a string representation of the pair:

```go
p := pair.New(1, 2)
fmt.Println(p.String()) // prints "(1, 2)"
```

### Supported Types

The Pair package supports the following types:

- `int`
- `int32`
- `int64`
- `uint`
- `uint32`
- `uint64`
- `float32`
- `float64`
- `string`
- `bool`

Each type has a corresponding pair struct (e.g., intPair, stringPair) that implements the Pair interface. Each pair struct has First and Second methods that return the first and second values of the pair, respectively.