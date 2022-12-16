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

(定义、标准、因果关系、价值观)