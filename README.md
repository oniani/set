# set

A generic set implementation for the Go programming language.

The implementation makes use of the built-in `map` data structure.

**TODO**:

- Add tests
- Add docs

## API

```go
import "github.com/oniani/set"

func main() {
    // Two ways to create and populate a set
    s := Set.New("Apple", "Banana", "Cherry")
    t := Set.New[string]()
    t.Add("Raspberry")
    t.Add("Strawberry")
    t.Add("Apple")

    // Check the results
    fmt.Println(s, s.Len())
    fmt.Println(t, t.Len())

    // Union, intersection, and difference
    m := s.Union(t)
    n := s.Intersection(t)
    o := s.Difference(t)

    // Check the Results
    fmt.Println(m, m.Len())
    fmt.Println(n, n.Len())
    fmt.Println(o, o.Len())

    sHasApple, sHasCar = s.Contains("Apple"), s.Contains("Car")
    sHasAll1, sHasAll2 = s.All("Car", "Laptop", "Popcorn"), s.All("Car", "Laptop", "Apple")
    sHasAny1, sHasAny2 = m.Any("Car", "Laptop", "Popcorn"), m.Any("Car", "Laptop", "Apple")

    // Check the results
    fmt.Println(sHasApple, sHasCar)
    fmt.Println(sHasAll1, sHasAll2)
    fmt.Println(s_has_any1, sHasAny2)

    // Clone, remove, and clear
    p := s.Clone()
    p.Remove("Apple")
    s.Clear()

    // Check the results
    fmt.Println(s, s.Len())
    fmt.Println(p, p.Len())

    // Map, filter, and slice conversion
    // TODO
}
```

## License

[MIT License][license]

[license]: LICENSE
