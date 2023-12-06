# Pair Package

The pair package provides a concrete implementation of a Pair of values. A Pair is a data structure that holds two values of any type.

## Usage

To use the pair package, import it into your project:

```go
import (
    ...
    "github.com/arthurcgc/pair"
    ...
)
```

To create a new pair, use the `NewPair` function:

```go
p := pair.NewPair(1, 2.0)
```

To get the first and second values of the pair, use the `First` and `Second` methods, respectively:

```go
p := pair.NewPair(1, "hello")
fmt.Println(p.First) // 1
fmt.Println(p.Second) // hello
```

To get the string representation of the pair, use the `String` method:

```go
p := pair.NewPair(1, 2.0)
fmt.Println(p.String()) // (1, 2.0)
```

To check if two pairs are equal, use the `Equals` function:

```go
p1 := pair.NewPair(1, 2)
p2 := pair.NewPair(1, 2)
fmt.Println(Equals(p1, p2)) // true
```

```go
p1 := pair.NewPair(1, "hello")
p2 := pair.NewPair(1, "world")
fmt.Println(Equals(p1, p2)) // false
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

I intend to add support for more types in the future.

#### License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
