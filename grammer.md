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

寻找 z，使得 z^2 最接近于 x，即是求 x的平方根的最近似值。

 x - z^2 = 0

导数方程是 x = 2z


```
func Sqrt(x float64) float64 {
    // implement
}
```
Switch

```
func main() {
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS;os{
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.\n", os)
	}
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

An empty(nil) slice declaration: `var names []string`.

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

Nil slices
The zero value of a slice is `nil`.
A nil slice has a length and capacity of 0 and has no underlying array.

```
func main() {
    var s []int
    fmt.Println(s, len(s), cap(s))
    if s == nil {
        fmt.Println("Nil!")
    }
}
```

Create a slice with make
Slices can be created with the built-in `make` function; this is how you create dynamically-size arrays.
The make function allocates a zeroed array and returns a slice that refers to that array: `a := make([]int, 5)`, len(a) = 5.

To specify a capacity, pass a third argument to `make`:

```
b := make([]int, 0, 5)  // len(b) = 0, cap(b) = 5
b = b[:cap(b)]  // len(b) = 5, cap(b) = 5
b = b[1:]   // len(b) = 4, cap(b) = 4
```

Slice of slices
Slices can contain any type, including other slices.

```
func main() {
    board : = [][]string{
        []string{"_", "_", "_"},
        []string{"_", "_", "_"},
        []string{"_", "_", "_"}
    }
}
```

Appending to a slice

It is common to append new elements to a slice, and so Go provides a built-in `append` function.

```
func append(s []T, vs ...T) []T
```
The first parameter `s` of append is a slice of type, and the rest are `T` values to append to the slice.

The resulting value of `append` is a slice containing all the elements of the original slice plus the provided values.

If the backing array of `s` is too small to fit all the given values, a bigger array will be allocated. The returned slice will point to the newly allocated array.

注意：slice 与内存释放。从一个大文件中筛选内容时，使用 slice 保存了选中结果，但存放结果的 slice 所依赖的数组还是该大文件。当 slice 没被 GC 释放回收，则该大文件同样存在于内存。要想释放大文件的内存，应在返回选中结果的slice 前，先用 make 新建一个slice(c := make([]byte, len(b)))，再复制(copy(c, b))。

Range
 
 The `range` form of the `for` loop iterates over a slice or a map.
When ranging over slice, two values are returned for each iteration. The first is the index, and the second is a copy of the element at that index.

```
var pow =[]int{1, 2, 4, 8, 16, 32, 64, 128}

func main() {

    for i, v := range pow {
        fmt.Println("2**%d = %d\n", i, v)
    }
}

```
Range continued

You can skip the index or value by assign to `_`.
```
for i, _ := range pow
for _, value := range pow
```
If you only want the index, you can omit the second variable.
```
for i := range pow
```

Exercise: Slices

Implement `Pic`. It should return a slice of length `dy`, each element of which is a slice of `dx` 8-bit unsigned integers. When you run the program, it will display your picture, interpreting the integers as grays (well, bluescale) values.
The choice of image is up to you. Interesting function include `(x+y)/2`, `x*y`, and `x^y`.

```
package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {

}

func main() {
    pic.Show(pic)
}
```

Maps 

A map maps keys to values.
The zero value of a map is `nil`. A `nil` map has no keys, nor can keys be added.
The `make` function returns a map of the given type, initialized and ready for use.

```
type Vertex struct {
    Lat, Long float64
}

var m map[string]Vertex // zero value is nil

func main() {
    m = make(map[string]Vertex) // initialize map
    // m := make(map[string]Veretex) // if not declare var m map[string]Vertex
    m["Bell Labs"] = Vertex{
        40.089, -74.399967,
    }
    fmt.Println(m["Bell Labs"])
}
```

Map literals

Map literals are like struct literals, but the keys are required.

```
type Vertex struct {
    Lat, Long float64
}

var m = map[string]Vertex{
    "Bell Labs": Vertex {
        40.68, -74.3999,
    },
    "Google": Vertex{
        37.42, -122.08,
    }
}
```

Map literals continued

If the top-level type is just a type name, you can omit it fromthe elements of the literal.

```
type Vertex struct {
    Lat, Long float64
}

var m = map[string]Vertex{
    "Bell Labs": {40.68, -74.399},
    "Google": {37.422, -122.084},
}
```
Mutating Maps

Insert or update an element in map `m`:   m[key] = elem

Retrieve an element: elem = m[key]

Delete an element: delete(m, key)

Test that a key is present with a two-value assignment: elem, ok = m[key]

If key is in m, ok is true. If not, ok is false.

If key is not in the map, then elem is the zero value for the map's element type.

Note: if elem or ok have not yet been declared you could use a short declaration form: `elem, ok := m[key]`

Exercise: Maps

Implement `WordCount`. It should return a map of the counts of each "word" in the string `s`. The `wc.Test` function runs a test suite against the provided function and prints success or failure. You might find strings.Fields helpful.

```
func WordCount(s string) map[string]int {
    count := make(map[string]int)

    for _, f := range strings.Fields(s) {
        if _, ok = count[f]; !ok {
            count[f] = 1
            continue
        }
        count[f] += 1
    }
}

```
Function values

Functions are values too. They can be passed around just like other values.
Function values may be used as function arguments and return values.
[function arguments] [return values]


```
// function value 内的函数参数是没有 name的
func compute(fn func(float64, float64) float64) float64 {
    return fn(3, 4)
}

func main() {
    hypot := func(x, y float64) float64 {
        return math.Sqrt(x*x + y*y)
    }
    fmt.Println(hypot(5, 12))
    fmt.Println(compute(hypot))
    fmt.Println(compute(math.Pow))
}
```

Function closures

Go functions may be closure. A closure is a function value that references variables from outside its body. The function may access and assign to the referenced variables; in this sense the function is "bound" to the variables.

For example, the `adder` function returns a closure. Each closure is bound to its own `sum` variable.

```
func adder() func(int) int {
    sum := 0
    return func(x int) int {
        sum += x
        return sum
    }
}

func main() {
    pos, neg := adder(), adder()
    for i := 0; i < 10; i++ {
        fmt.Println(
            pos(i),
            neg(-2*i), // a newline must end with ,
        )
    }
}
```

Exercise: Fibonacci closure

Let's have some fun with functions.
Implement a `fibonacci` function that returns a function (a closure) that returns successive `[fibonacci numbers](https://en.wikipedia.org/wiki/Fibonacci_number)`(0 1 1 2 3 5, ...)

```
func fibonacci () func() int {
    x := 0
    y := 1
    return func() int {
        y += x
        x, y = y,x
        return y
    }
}
```

## Methods

Methods

Go does not have classes. However, you can define methods on types.
A method is a function with a special `receiver` argument.

The receiver appears in its own argument list between the `func` keyword and the method name.

In this example, the `Abs` method has a receiver of type `Vertex` named `v`.

```
type Vertex struct {
    X, Y float64
}

func (v Vertex) Abs() float64 {
    return math.Sqrt(v.X * v.X + v.Y * v.Y)
}

func main() {
    v := Vertex{3, 4}
    fmt.Println(v.Abs())
}
```

Methods are functions

Remember: a method is just a function with a receiver argument.

Here's Abs written  as a regular function with no change in funtionality.

```
type Vertex struct {
    X, Y float64
}

func Abs(v Vertex) float64 {
    return math.Sqrt(v.X * v.X + v.Y *v.Y)
}

func main() {
    v := Vertex{3, 4}
    fmt.Println(Abs(v))
}
```

Method continued

You can declare a method on non-struct types, too.
In this example we see a numeric type `MyFloat` with an `Abs` method.

You can only declare a method with a receiver whose type is defined in the same package as the method. You cannot declare a method with a receiver whose type is defined in another package (which includes the built-in types such as `int`).


```
type MyFloat float64

func (f MyFloat) Abs() float64 {
    if f < 0 {
        return float64(-f)
    }
    return float64(f)
}
func main() {
    f := MyFloat64(-math.Sqrt2) // convert -math.Sqrt2 to MyFloat64, expression T()
    fmt.Println(f.Abs())
}
```

Pointer receivers
[value receiver] [pointer receiver]
You can declare methods with pointer receivers.
This means the receiver type has the literal syntax `*T` for some type `T`. (Also, T cannot itself be a pointer such as `*int`.)

For example, the `Scale` method here is defined on `*Vertex`.

Methods with pointer receivers can modify the value to which the receiver points(as Scale does here). Since methods often need to modify their receiver, pointer receivers are common than value receivers.

With a value receiver, the Scale method operates on a copy of the original Vertex value.(This is the same behavior as for any other function argument.) The Scale must have a pointer receiver to change the Vertex value declared in the main function.

```
// method-pointers.go

type Vertex struct {
    X， Y float64
}

// receiver - value receiver or pointer receiver
func (v *Vertex) Abs() float64 {
    return math.Sqrt(v.X * v.X + v.Y * v.Y)
}
func (v *Vertex) Scale(f float64) {
    v.X = v.X * f
    v.Y = v.Y * f
}
```

Pointers and functions

Here we see the `Abs` and `Scale` methods rewritten as functions. Again, try removing the `*` from line 16. Can you see why the behavior changes? Why else did you need to change for the example to compile? (If you're not sure, continue to the next page.)

```
type Vertex struct {
    X, Y float64
}

func Abs(v Vertex) float64 {
    return math.Sqrt(v.X * v.X + v.Y * v.Y)
}

func Scale(v *Vertex, f float64) {
    v.X = v.X * f
    v.Y = v.Y * f
}

func main() {
    v := Vertex{3, 4}
    Scale(&v, 10)
    fmt.Println(Abs(v))
}
```

Methods and pointer indirection

Comparing the previous two program, you might notice that functions with a pointer argument must take a pointer:

```
var v Vertex
ScaleFunc(v, 5) // compile error!
ScaleFunc(&v, 5) // ok
```
While methods with pointer receivers take either a value or a pointer as the receiver when they are called:

```
var v Vertex
v.Scale(5) // ok
p := &v
p.Scale(10) // ok
```
For teh statement v.Scale(5), even though v is a value and not a pointer, the method with the pointer receiver is called automatically. That is, as a convenience, Go interprets the statement v.Scale(5) as (&v).Scale(5) since the Scale method has a pointer receiver.

```
type Vertex struct {
    X, Y float64
}

func (v *Vertex) Scale(f float64) {
    v.X = v.X * f
    v.Y = v.Y * f
}

func ScaleFunc(v *Vertex, f float64) {
    v.X = v.X * f
    v.Y = v.Y * f
}

func main() {
    v := Vertex{3, 4}
    v.Scale(2)
    ScaleFunc(&v, 10) // &v pointer
    p := &Vertex{4, 3} // pointer
    p.Scale(3)
    ScaleFunc(p, 8) // p is a pointer value

    fmt.Println(v, p) // v is a value, p is a poniter value.
}
```
Methods and pointer indirection (2)

The equivalent happens in the reverse direction.

Functions that take a value argument must take a value of that specific type:
```
var v Vertex
fmt.Println(AbsFunc(v)) // ok
fmt.Println(AbsFunc(&v)) // compile error!
```
Whie methods with value receivers take either a value or a pointer as the receiver when they are called:

```
var v Vertex
fmt.Println(v.Abs()) // ok

p := &v
fmt.Println(p.Abs()) // ok

```
In this case, the method call p.Abs() is interpreted as `(*p).Abs()`.

```
type Vertex struct {
    X, Y float64
}

func (v Vertex) Abs() float64 {
    return math.Sqrt(v.X * v.X + v.Y * v.Y)
}

func AbsFunc(v Vertex) float64{
    return math.Sqrt(v.X * v.X + v.Y * v.Y)
}

func main() {
    v := Vertex{3, 4}
    fmt.Println(v.Abs())
    fmt.Println(AnsFunc(v))

    p := &Vertex{4, 3}
    fmt.Println(p.Abs())
    fmt.Println(AbsFunc(*p))
}
```

Choosing a value or a pointer receiver

There are two reasons to use a pointer receiver.
The first is so that the method can modiy the value that its receiver points to.
The second is to avoid copying the value on each method call. This can be more efficient if the receiver is a large struct, for example.

In this example, both Scale and Abs are with receiver type `*Vertex`, even though the `Abs` method needn't modify its receiver.

In general, all methods on a given type should have either value or pointer receivers, but not a mmixture of both. (We'll see why over the next few pages.)

```
type Vertex struct {
    X, Y float64
}

func (v *Vertex) Abs() float64 {
    return math.Sqrt(v.X * v.X + v.Y * v.Y)
}

func (v *Vertex) Scale(f float64) {
    v.X = v.X * f
    v.Y = v.Y * f
}

func main() {
    v := &Vertex{3, 4}
    fmt.Printf("Before scaling: %+v, Abs: %v\n", v, v.Abs())
    v.Scale(5)
    fmt.Printf("After Scaling: %+v, Abs: %v\n", v, v.Abs())
}
```

Interfaces

[A set of method signatures]An __interface type__ is defined as a set of method signatures.

A value of interface type can hold any value that implements those methods.

Note: There is an error in the example code on line 22. `Vertex`(the value type) doesn't implement `Abser`
because the `Abs` method is defined only on `*Vertex`(the pointer type).


```
package main

import (
    "fmt"
    "math"
)

// a set of method signatures
type Abser interface {
    Abs() float64
}

func main () {
    var a Abser
    f := MyFloat(-math.Sqrt2)
    v := Vertex{3 ,4}
    
    a = f   // a MyFloat implements Abser
    a = &v  // a *Vertex implements Abser
    
    // In the following line, v is a Vertex (not *Vertex) and does NOT implement Abser.
    // a = v
    fmt.Println(a.Abs())
}

// type MyFloat
type MyFloat float64

func (f MyFloat) Abs() float64 {
    if f < 0 {
        return float64(-f)
    }
    return float64(f)
}

// type Vertex
type Vertex struct {
    X, Y float64
}

func (v *Vertex) Abs() float64 {
    return math.Sqrt(v.X * v.X + v.Y * v.Y)
}
```
Interface are implemented implicitly

A type implements an interface by implementing its methods. There is no explicit declaration of intent, no "implements" keyword.

Implicit interface decouple the definition of an interface from its implementation, which could then appear in any package without prearrangement.


```
package main

type I interface {
    M()
}

type T struct {
    S string
}


// This method means type T implements the interface I, but we don't need to explicitly(显式) declare that it does so.
func (t T) M() {
    fmt.Println(t.S)
}

func main() {
    var i I = T{"Hello, I am type T struct"}
    i.M() // access function M
}
```

Interface values

Under the hood, interface values can be thought of as a tuple of a value and a concrete type: `(value, type)`
An interface value holds a value of a specific underlying concrete type.
Calling a method on an interface value executes the method of the same name on its underlying type.
[a method on an interface value] [the method of the same name on its underlying type]

```
type I interface {
    M()
}

type T struct {
    S string
}

// pointer receiver
func (t *T)M() {
    fmt.Println(t.S)
}

type MyFloat float64

func (f MyFloat)M() {
    fmt.Println(f)
}

func main() {
    var i I
    i = &T{"Hello"}
    describe(i)
    i.M() // underlying type is T

    i = MyFloat(-math.Sqrt2)
    describe(i)
    i.M() // underlying type is F
}
```
Interface values with nil underlying values

If the concrete value inside the interface is nil, the method will be called with a nil receiver.
In some languages this would trigger a null pointer exception, but in Go it is common to write methods that gracefully handle being called with a nil receiver (as with the method M in this example).

Note that an interface value that holds a nil concrete value is itself non-nil.

```
type I interface {
    M()
}

type T struct {
    S string
}

func (t *T) M() {
    if t == nil {
        fmt.Println("<nil>")
        return
    }
    fmt.Println(t.S)
}

func main() {
    var i I
    var t *T
    i = t // t is nil
    describe(i)
    i.M() // called method M()

    i = &T{"hello"}
    describe(i)
    i.M()
}

func describe(i I) {
    fmt.Printf("(%v, %T)\n", i, i)
}
```

Nil interface values

A nil interface value holds neither value nor concrete type.

Calling a method on a nil interface is a run-time error because there is no type inside the interface tuple to indicate which `concrete` method to call.

```
type I interface {
    M()
}

func main() {
    var i I // nil
    describe(i)
    i.M() // panic: runtime error
}

func describe(i I) {
    fmt.Printf("(%v, %T)\n", i, i)
}
```

The empty interface

The interface type that specifies zero methods is known as the __empty interface__: `interface{}`

An empty interface may hold values of any type. (Every type implements at least zero methods.)

Empty interfaces are used by code that handles values of unknown type. For example, `fmt.Print` takes any numbre of arguments of type `interface{}`.

```
func main() {
    var i interface{}
    describe(i)
    
    // empty interface are used by code that handles values of unknown type.
    i = 42 // int
    describe(i)

    i = "hello" // string
    describe(i)
}

func describe(i interface{}) {
    fmt.Printf("(%v, %T)\n", i, i)
}
```

Type assertions
(类型判断)

A __type assertion__ provide access to an interface value's underlying concrete value.

```
t := i.(T)
```
This statement asserts that the interface value `i` holds the concrete type `T` and assigns the underlying `T` value to the variable `t`.

If `i` does not hold a `T`, the statement will trigger a panic.

To test whether an interface value holds a specific type, a type assertion can return two values: the underlying value and a boolean value that reports whether the assertion succeeded.

```
t, ok := i.(T)
```
If `i` holds a `T`, then `t` will be the underlying value and `ok` will be true.
If not, `ok` will be false and `t` will be the zero value of type `T`, and no panic occurs.
[the zero value of type T]

Note the similarity between this syntax and that of reading from a map. 

```
// map
var m map[string]int

v, ok = map[k]
```

```
// type assertion

func main() {
    var i interface{} = "hello"

    s := i.(String)
    fmt.Println(s)

    s, ok := i.(String)
    fmt.Println(s, ok)

    f, ok := i.(float64)
    fmt.Println(f, ok) // the zero value of type T(float64)

    f = i.(float64) // panic
    fmt.Print(f)
}
```

Type switches

A __type switch__ is a construct that permits several type assertions in series.

A type switch is like a regular switch statement, but the cases in a type switch specify types (not values), and those values are compared against the type of the value held by the given interface value.

```
switch v := i.(type) {
    case T:
        // here v has type T
    case S:
        // here v has type S
    default:
        // no match; here v has the same type as i
}
```

The declaration in a type switch has the same syntax as a type assertion `i.(T)`, but the specific type `T` is replaced with the keyword `type`.

This switch statement tests whether the interface value `i` holds a value of type `T` or `S`. In each of the `T` and `S` cases, the variable `v` will be of type `T` or `S` respectively and hold the value held by `i`. In the default case (where there is no match), the variable `v` is of the same interface type and value as `i`.

```
func do(i interface{}) {
    switch v: = i.(type) {
        case int:
        fmt.Printf("Twice %v is %v \n", v, v*2)
        case string:
        fmt.Printf("%q is %v bytes long\n", v, v)
        default:
        fmt.Printf("I don't know about type %T\n", v)
    }
}


```

Stringers

One of the most ubiquitous interface is `Stringer` defined by the `fmt` package.

```
type Stringer interface {
    String() string
}
```

A `Stringer` is a type that can describe itself as a string. The `fmt` package (and many others) look for this interface to print values.


```
type Person struct{
    Name string
    Age int
}

func (p Person) String() string{
    return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

func main() {
    a := Person{"Arthur Dent", 42}
    z := Person{"Zaphod Beeblebrox", 9000}
    fmt.Println(a,z)
}
```

Exercise: Stringers

Make the `IPAddr` type implement `fmt.Stringer` to print the address as a dotted quad.
For instance, `IPAddr{1, 2, 3, 4}` should print as "1.2.3.4".

```
package main

import (
    "fmt"
    "strings"
)

type IPAddr [4]byte // array of bytes

func (ip IPAddr) String() string {
    var o []string

    for _, p := range ip {
        o = append(o, fmt.Sprint(int(p)))
    }
    return strings.Join(o, ".")
}

func main() {
    hosts := map[string]IPAddr{
        "loopback": {127, 0, 0, 1},
        "googleDNS": {8, 8, 8, 8}
    }

    for name, ip := range hosts {
        fmt.Printf("%v: %v\n", name, ip)
    }
}
```

Errors

Go programs express error state with error values.

The error type is a built-in interface similar to `fmt.Stringer`:

```
type error interface {
    Error() string
}
```
(As with fmt.Stringer, the fmt package looks for the error interface when printing values.)

Functions often return an error value, and calling code should handle errors by testing whether the error equals nil.

```
i, err := strconv.Atoi("42")

if err != nil {
    fmt.Printf("could't convert number: %v\n", err)
    return
}
fmt.Println("Converted integer:", i)
```
A nil error denotes success; a non-nil error denotes failure.

// errors.go
```
package main

import(
    "fmt"
    "time"
)

type MyError struct{
    When time.Time
    What string
}

func(e *MyError) Error() string {
    return fmt.Sprintf("at (%v), %s", e.When, e.What)
}

func run() error {
    return &MyError{
        time.Now(),
        "it didn't work",
    }
}

func main() {
    if err := run(); err != nil {
        fmt.Println(err)
    }
}

```

Exercise: Errors
Copy your `Sqrt` function from the earlier exercise and modify it to return an `error` value.

`Sqrt` should return a non-nil error value when given a negative number, as it doesn't support complex numbers.

Create a new type
```
type ErrNegativeSqrt float64
``
and make it an `error` by giving it a 
```
func (e ErrNegativeSqrt) Error() string
```
method such that `ErrNegativeSqrt(-2).Error()` returns "cannot Sqrt negative number: -2".

Note: A call to `fmt.Sprint(e)` inside the `Error` method will send the program into an infinite loop. You can avoid this by converting `e` first: `fmt.Sprint(float64(e))`. Why?
Change your `Sqrt` function to return an `ErrNegativeSqrt` value when given a negative number.

---

## methods

Readers

The `io` package specifies the `io.Reader` interface, which presents the read end of a  stream of data.

The Go standard library contains many implementations of this interface, including files, network connnections, compressors, ciphers, and others.

The `io.Reader` interface has a `Read` method:

```
func (T) Read(b []byte) (n int, err error)
```
`Read` populate the given byte slice with data and returns the number of bytes populated and an error value. It returns an `io.EOF` error when the stream ends.

The example code creates a `string.Reader` and consumes its output 8 bytes at a time.

```
package main

import (
    "fmt"
    "strings"
    "io"
)

func main() {
    f := strings.NewReader("Hello, Reader!")

    b := make([]byte, 8)

    for {
        // populate the byte slice(b) with data
        // return the number of bytes populated
        n, err := r.Read(b)  
        fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
        fmt.Printf("b[:n]=%q\n", b[:n])
    }
}
```

---
Built-in

func cap ¶
func cap(v Type) int
The cap built-in function returns the capacity of v, according to its type:
Array: the number of elements in v (same as len(v)).
Pointer to array: the number of elements in *v (same as len(v)).
Slice: the maximum length the slice can reach when resliced;
if v is nil, cap(v) is zero.
Channel: the channel buffer capacity, in units of elements;
if v is nil, cap(v) is zero.

(定义、标准、因果关系、价值观)


-----
黄执中说话课程
-----
我有一个需求，请问如何让大家来满足我？很多人都是有类同的需求，底层是同样的需求模式。
如何让大家知道我喜欢他们？

能看到别人有需求的人，才能产生影响力？

跟别人相处，不需去展露你的博学，而是要展露你的好奇，你的感知（认同情绪等）。

聊天不是交换信息，而是交换情绪。

没碰到情绪，就是没触碰到你。***
年轻人不是不给情绪，而是只给了自己人，如：KTV，遇到喜欢的朋友。

更多人的问题是自己不听自己的话。
