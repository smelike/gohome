Exported names

In Go, a name is exported if it begins with a capital letter. For example, Pizza is an exported name, as is Pi, which is exported from the math package.

Funtions

```
func add(x int, y int) int {
    return x + y
}
```

When two or more consecutive named function parameters share a type, you can omit the type from all but the last. In this example, we shortened `x int, y int` to `x, y int`.


Functions: Multiple results

A function can return any number of results. The `swap` function returns two strings.

```
func swap(x, y string) (string, string) {
    return y, x
}
```

Functions: Named return values

Go's return values may be named. If so, they are treated as variables defined at the top of the function. These names should be used to document the meaning of the return values.

A return statement without arguments returns the named return values. This is known as a "naked" return.

Naked return statements should be used only in short functions, as with the example shown here. They can harm readability in longer functions.

```
func slit(sum int) (x, y int) {
    x = sum * 4 / 9
    y = sum - x
    return
}
```

Variables

The `var` statement declares a list of variables; as in function argument list, the type is last. A `var` statement can be at package or function level. We see both in this example.

```
package main

import "fmt"

var c, python, java bool

func main() {
    var i int
    fmt.Println(i, c, python, java)
}

```

Variables: variables with initializers

A var declaration can include initializers, one per variable. If an initializer is present, the type can omitted; the variable will take the type of the initializer.
```
package main

import "fmt"

var i, j int = 1, 2 // var i, j = 1, 2
func main() {
    var c, python, java = true, false, "no!"
    fmt.Println(c, python, java)
}
```

Variables: short variable declarations

Inside a function, the `:=` short assignment statement can be used in place of a `var` declaration with implicit type.

Outside a function, every statement begins  with a keyword (`var`, `func`, and so on) and so the `:=` construct is not available.

```
func main() {
    var i, j int = 1, 2
    k := 3 // inside a function level
    c, python, java := true, false, "no!" // inside a function level
    fmt.Println(i, j, k, c, python, java)
}
```

Basic types
bool
string
int int8 int16 int32 int64
uint uint8 uint16 uint32 uint64 uintptr

byte    // alias for uint8
rune    // alias for int32
        // represents a Unicode code point
float32 float64
complex64   complex128

---
The example shows variables of several types, and also that variable declarations may be "factored" into blocks, as with import statements.

```
package main

// import block
import (
    "fmt"
    "math/cmplx"
)

var (
    ToBe bool = false
    MaxInt uint64 = 1<<64 - 1
    z complex128    = cmplx.Sqrt(-5 + 12i)
)

func main() {
    fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
    fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
    fmt.Printf("Type: %T Value: %v\n", z, z)
}
```

Zero Values: Variables declared without an explicit initial value are given their zero value. [explicit vs. implicit]

The zero value is:
0 for numeric type,
false for the boolean type, and ""(the empty string) for strings.


Type conversions: The expression `T(v)` converts the value `v` to the type `T`.

Some numeric conversions:
```
var i int = 42
var f float64 = float64(i)
var u uint = uint(f)
```
Or put more simply:
```
i := 42
f := float64(i)
u := uint(f)
```

```
// type-conversions.go
import (
    "fmt"
    "math"
)
func main() {
    var x, y int = 3, 4
    var f float64 = math.Sqrt(float64(x*x + y *y))
    var z uint = uint(f)
    fmt.Println(x, y, z)
}
```

Type inference
When declaring a variable without specifying an explicit type (either by using the `:=` syntax or `var=` experssion syntax), the variable's type is inferred from the value on the right hand side.

When the right hand side of the declaration is typed, the new variable is of that same type: 
```
var i int
j := i // j is an int
```
[an untyped numeric constant]But when the right hand side contains an untyped numeric constant, the new variable may be an `int`, `float64`, or `complex128` depending on the precision of the constant:

```
i := 42 // int
f := 3.142 // float64
g := 0.867 + 0.5i // complex128
```

Constants
Constants are declared like variables, but with the `const` keyword.
Constants can be character, string, boolean, or numeric values.
Constants cannot be declared using the `:=` syntax.

```
const Pi = 3.14
const World = "世界"
```

Numeric Constants
Numeric constants are high-precision values.
An untyped constant takes the type needed by its context.

```
const (
    Big = 1 << 100
    Small = Big >> 99
)
```

---
## Flowcontrol

If with a short statement
Like `for`, the `if` statement can start with a short statement to execute before the condition. Variables declared by the statement are only in scopee until the end of the if. [only in scope until...]

Exercise: Loops and Functions

Let's implement a square root function: given a number x, we want to find the number z for which z^2 is most nearly x.

寻找 z，使得 z^2 最接近于 x，即是求 x 平方根的最近似值。

```
func Sqrt(x float64) float64 {
    // implement
}
```

Swith with no condition

Switch without a condition is the same as `switch` true. This construct can be a clean way to write long if-then-else chains.

```
func main() {
    t := time.Now()
    switch {
        case t.Hour() < 12:
            fmt.Println("Good morning!")
        case t.Hour() < 17:
            fmt.Println("Good afternoon.")
        default:
            fmt.Println("Good evening.")
    }
}

``

Defer

A defer statement defers the execution of a function until the surrounding returns. The deferred call's arguments are evaluated immediately, but the function call is not executed until the surrouding function returns.

![defer.go 结果图](./snapshots/defer.go.png)

Stacking defers
Deferred function calls are pushed onto a stack. When a function returns, its deferred calls are executed in last-in-first-out order.

```
func main() {
    fmt.Println("counting")
    for i :=0; i < 10; i++ {
        defer fmt.Println(i)
    }
    fmt.Println("Done")
}

/*
    Output:
    counting
    Done
    9
    ...
    0
*/
```
---

Moretypes

Moretypes | Pointers

Go has pointers. A pointer holds the memory address of a value.
The type `*T` is a pointer to a `T` value. Its zero value is `nil`.
```
var p *int
```
The `&` operator generates a pointer to its operand.
```
i := 42
p = &i // point to i
```
The `*` operator denotes the pointer's underlying value.
```
fmt.Println(*p) // read i through the pointer p
*p = 21 // 
```
This is known as "dereferencing(*)" or "indrecting(&)".
Structs: A `struct` is a collection of fields.

```
package main

import "fmt"

type Vertex struct {
    X int
    Y int
}

func main() {
    fmt.Println(Vertex{1, 2})
}
```
Struct Fields
Struct fields are accessed using a dot.

```
type Vertex struct {
    X int
    Y int
}
func main() {
    v := Vertex{1, 2}
    v.X = 4
    fmt.Println(v.X)
}
```

Pointers to structs
Struct fields can be accessed through a struct pointer.
To access the field X of a struct when have the struct pointer p we could write `(*p).X`. However, that notation is cumbersome, so the language permits us instead to write just `p.X`, without the explicit dereference.[(*p).X]

```
func main() {
    v := Vertex{1, 2}
    p := &v
    p.X = 1e9
    fmt.Println(v)
}
```
Struct Literals

A struct literal denotes a newly allocated struct value by listing the values of its fields.
You can list just a subset of fields by using the `Name:` syntax. (And the order of named fields is irrelevant.)
The special prefix `&` returns a pointer to the struct value.

```
type Vertex struct {
    X, Y int
}
var (
    v1 = Vertex{1, 2}   // has type Vertex
    v2 = Vertex{X: 1}   // Y:0 is implicit
    v3 = Vertex{}   // X: 0 and Y:0
    p = &Vertex{1, 2}   // has type *Vertex
)

```

Arrays | [n]T
The type `[n]T` is an array of `n` values of type `T`.

The expression `var a [10]int` declares a variable `a` as an array of ten integers.
An arrray's length is part of its type, so arrays cannot be resized. This seems limiting, but don't worry; Go provides a convenient way of working with arrays.

```
func main() {
    var a [2]string
    a[0] = "Hello"
    a[1] = "World"
    fmt.Println(a[0], a[1])
    fmt.Println(a)

    primes := [6]int{2, 3, 5, 7, 11, 13}
    fmt.Println(primes)
}
```

Slices
An array has a  fixed size. A slice, on the other hand, is a dynamically-sized, flexible view into the elements of an array. In practice, slice are much more common than arrays.

The type `[]T` is a slice with elements of type `T`.
A slice is formed by specifying two indices, a low and a high bound, separated by a colon: `a[low : high]`

This selects a half-open range which includes the first element, but excludes the last one. The following expression creates a slice which includes elements `1` through `3` of `a`: `a[1:4]`

An empty slice declaration: `var names []string`.

Slices are like references to arrays

A slice does not store any data, it just describes a section of an underlying array. Changing the elements of a slice modifies the corresponding elements of its underlying array. Other slices that share the same underlying array will see those changes.


```
func main() {
    names := []string{
        "John",
        "Paul",
        "George",
        "Ringo",
    }
    fmt.Println(names) // [John XXXX George Ringo]

    b := names[1:3]
    b[0] = "XXXX"
    fmt.Println(names) // [John XXXX George Ringo]
}
```

Slice literals

A slice literal is like an array literal without the length.
This is an array literal: `[3]bool{true, true, false}`.
And this creates the same array as above, then builds a slice that references it:
```
[]bool{true, true, false}
```

```
// fiel-value pairs
s := []struct{
    i int
    b bool
}{
    {2, true},
    {3, true},
    {5, true},
    {7, true},
    {11, false},
    {13, false},
}
```

Slice defaults
When slicing, you may omit the high or low bounds to use their defaults instead. The default is zero for the low bound and the length of the slice for the high bound.

For the array `var [10]int`, these slice expressions are equivalent:
```
a[0:10]
a[:10]
a[0:]
a[:]
```

For example,

```
func main() {
	s := []int{2, 3, 5, 7, 11, 13}

	s = s[1:4] // [3,5,7]
	fmt.Println(s)

	s = s[:2] // [3,5] // exclude the index 2
	fmt.Println(s) 

	s = s[1:] // [5] // the lenght of slice is 2.
	fmt.Println(s)
}
```

Slice length and capacity

A slice has both a length and a capacity.
The length of a slice is the number of elements it contains.
The capacity of a slice is the number of elements in the underlying array, counting from the first element in the slice.
The length and capacity of a slice `s` can be obtained using the expressions `len(s)` and `cap(s)`.

![slice-len.go](./snapshots/slice-len.go.png)
(定义、标准、因果关系、价值观)