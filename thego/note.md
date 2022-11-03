## ch #1
---

Like the for and if statements, a switch may inclue an optional simple statement - a short variable declaration, an increment or assignment statement, or a function call - that can be used to set a value before it is tested.

The break and continue statements modify the flow of control. A break causes control to resume at the next statement after the innermost for,switch, or select statement.

A continue causes the innermost for loop to start its next iteration. Statements may be labeled so that break and continue can refer to them, for instance to break out of several nested loops at once or to start the next iteration of the outermost loop. There is even a goto statement, though it's intended for machine-generated code, not regular use by programmers.

**Named types:** A type declaration makes it possible to give a name to an existing type. Since struct type are often long, they are always named. A familiar example is the definition of a Point type for a 2-D graphics system:

type Point struct {
    X, Y int
}
var p Point

**Pointers:** values that contain the address of a variable.  The & operator yields the address of a variable, and the * operator retrieves the variable that the pointer refers to, but there is no pointer arithmetic.


**Methods and interface:**  A method is a function associated with a named type; Go is unusual in that methods may be attached to almost any named type. __Interfaces are abstract types__ that let us treat different concrete types in the same way based on what methods they have, not how they are represented or implemented.

**Packages:** The standard library packages at https://golang.org/pkg and the packages contributed by the community at https://godoc.org. The go doc tool makes these documents easily accessible from the command line:

>$ go doc http.ListenAndServe

---
## ch2 Program Structure
---


**25 keywords:**
    break       default     func    interface   select
    case        defer       go      map         struct
    chan        else        goto    package     switch
    const       fallthrough if      range       type
    continue    for         import  return      var

**predeclared names:**
    Constants: true false iota nil
    Types:
        int int8 int16 int32 int64
        uint uint8 uint16 uint32 uint64
        float32 float64 complex128 complex64
        bool byte rune string error

    Functions: make len cap new append copy close delete
                complex real imag panic recover

**Declaration Visibility:**
    - If an entity is declared within a function, it is local to that function.

    - If declared outside of a function, however, it is visible in all files of the package to which it belongs.

    - The case of the first letter of a name determines its visibility across package boundaries. If the name begins with an upper-case letter, it is exported, which means that it is visible and accessible outside of its own package and may be referred to by other parts of the program, as with Printf in the fmt package. 

    - Package names themselves are always in lower case.

    - Local variables with small scopes; you are much more likely to see variables named i than theLoopIndex.

    - The larger the scope of a name, the longer and more meaningful it should be.

### 2.2 Declarations

Four major kinds of declarations: var, const, type, and func.

Each file begins with a **package** declaration that says what package the file is part of.
The package declaration is followed by any import declarations, and then a sequence of package-level declarations of types, variables, constants, and functions, in any order.


Each declaration has the general form 
> var name type = expression
Either the type or the = expression part may be omitted, but not both.

If the type is omitted, it is determined by the initializer expression.

If the expression is omitted, the initial value is the zero value for the type, which is 0 for numbers, false for booleans, "" for strings, and nil for interfaces and reference types (slice,pointer,map,channel,function).

The zero value of an aggregate type like an array or a struct has the zero value of all of its elements or fields.

declare and optionally initialize a set of variables in a single declaration, with a matching list of expressions:
> var i, j, k int // int, int, int

Omitting the type allows declaration of multiple variables of different types:
> var b, f, s = true, 2.3, "four" // bool, float64, string

__Package-level variables__ are initialized before main begins.
And __Local variables__ are initialized as their declarations are encountered **during function execution**.


A set of variables can also be initialized by calling a function that returns multiple values:
> var f, err = os.Open(name) // os.Open returns a file and an error

### 2.3 Short Variable Declarations

Within a function, a short variable declaration may be used to declare and initialize local variables.

The form __name := expression__
the type of name is determined by the type of expression.

**A var declaration** tends to be reserved for local variables that need an explicit type that differs from that of the initializer expression, or for when the variable will be assigned a value later and its initial value is unimportant.

```
i := 100 // an int
var boiling float64 = 100 // a float64

var names []string
var err error
var p Point

```

**A tuple assignment**
> i, j = j, i // swap values of i and j

A short variable declaration must declare at least one new variable, so this code will not compile:

```
f, err := os.Open(infile)

// ...

f, err := os.Open(outfile) // compile error: no new variables
```

In the code below, the first statement declares both in and err. The second declares out **but only assigns a value to the existing err variable.**

```
in, err := os.Open(infile)
// ...
out, err := os.Create(outfile)
```

### 2.3.2 Pointers

A variable is a piece of storage containing a value.

A pointer value is the address of a variable.
A pointer is thus the location at which a value is stored.
Not every value has an address, but every variable does.
With a pointer, we can read or update the value of a variable indirectly, without using or even knowing the name of the variable, if indeed it has a name.

```
x := 1
p := &x // p, of type *int, points to x
fmt.Println(*p) // "1"
fmt.Println(p) // pointer

*p = 2  // equivalent to x = 2
fmt.Println(x) // "2"
```
If a variable is declared __var x int__, the expression **&x("address of x")** yields a pointer to an integer variable(x), that is, a value of type *int, which pronounced "pointer to int."
If this value is called p, we say "p points to x", or equivalently "p contains the address of x". The variable to which **[p points]** is written *p.

[
    **A pointer to an integer variable, that is a value of type *int.**
    A pointer to an string variable, that is a value of typr *string.
]
The expression *p yields the value of that variable(x), an int, but since __*p denotes(代表) a variable__, it may also appear on the left-hand side of an assignment, in which case the assignment updates the variable.


Pointers are comparable; two pointers are equal if and only if they point to the same variable or both are nil.

```
var x, y int

fmt.Println(&x == &x, &x == &y, &x == nil) // true false false

```

It is perfectly safe for a function to return the address of a local variable. For intance, in the code below, the local variable v created by this particular call to f will remain in existence even after the call has returned, and the pointer p will still refer to it:


```
var p = f()

func f() *int{
    v := 2
    return &v
}
```

Each call of f returns a distinct value:
> fmt.Println(f() == f()) // false

Because a pointer contains the address of a variable, passing a pointer argument to a function makes it possible for the function to update the variable that was indirectly passed.

For example, this function increments the variable that its argument points to and returns the new value of the variable so it may be used in an expression:

```
func incr(p *int) int {
    *p++
    return *p
}

v := 1
incr(&v) // side effect: v is now 2
fmt.Println(incr(&v)) // 3
```


### 2.3.3 The new Function

Another way to create a variable is to use the built-in function **new**.

The expression new(T) creates an unnamed variable of type T, initializes it to the zero value of T, and returns its address, which is a value of type *T. (T maybe int, string, slice, and so on)

```
p := new(int) // p, of type *int, points to an unamed int variable
fmt.Println(*p) // 0
*p = 2 // sets the unamed int to 2
fmt.Println(*p) // 2

```

The new is only a syntactic convenience, not a fundamental notion:
the two newInt functions below have identical behaviors:
```
func newInt() *int {
    return new(int)
}

func newInt() *int {
    var dummy int
    return &dummy
}
```

Each call to new returns a distincet variable with a unique address:

```
p := new(int)
q := new(int)
fmt.Println(p == q)  // false
```

There is one exceotion to this rule: two variables whose type carries no information and is therefore of size zero, such as struct{} or [0]int, depending on the implementation, have the same address.

Since new is a predeclared function, not a keyword, it's possible to redefine the name for something else within a function, for example:

> func delte(old, new int) int { return new - old }


### 2.3.4 Lifetime of Variabales

The lifetime of a variable is the interval of time during which it exists as the program executes.
The lifetime of a package-level variable is the entire execution of the program.
By contrast, local variables have dynamic lifetime: a new instance is created each time the declaration statement is executed, and the variable lives on until it becomes unreachable, at which point its storage may be recycled.

Function parameters and results are local variables too; they are created each time their enclosing function is called.

For example, in this excerpt from the Lissajous program of Section 1.4,

```
for t := 0.0; t < cycles*2*math.Pi; t+=res {
    x := math.Sin(t)
    y := math.Sin(t*freq + phase)
    img.SetColorIndex(size + int(x*size+0.5), size+int(y*size+0.5), blackIndex)
}
```
the variable t is created each time the **for loog** begins, and new variables x and y are created on each iteration of the loop.


A compiler may choose to allocate local variables __on the heap or on the stack__ but, perhaps surprisingly, this choice is not determined by whether var or new was used to declare the variable.

```
var global *int

func f() {
    var x int
    x = 1
    global = &x
}

func g() {
    y := new(int)
    *y = 1
}

```

Here, **x must be heap-allocated** because it is still reachable from the variable global after f has return, despite being declared as a local variable; we say x escapes from f.

Conversely, when g returns, the variable *y becomes unreachable and can be recycled. **Since *y does not escape from g, it's safe for the compiler to allocate *y on the stack, even though it was allocated with new.**

In any case, the notion of escaping is not something that you need to worry about in order to write correct code, though it's good to keep in mind during performance optimization, since each variable that escape requires an extra memory allocation.

Garbage collection is a tremendous help in writing correct programs, but it does not relieve you of the burden of thinking about memory. You don't need to explicitly allocate and free memory, but to write efficient programs you still need to be aware of the lifetime of variables.

For example, keeping unnecessary pointers to short-lived objects within long-lived object, especialy global variables, will prevent the garbage collector from reclaiming the short-lived objects.


### 2.4 Assginments

```
x = 1 //named variable
*p = true // indirect variable
person.name = "bod" // struct field
count[x] = count[x] * scale // array or slice or map element

```
Each of the arithmetic and bitwise binary operators has a corresponding assignment operator allowing, for example, the last statement to be written as `count[x] *= scale` which saves us from having to repeat (and re-evaluate) the expression for the variable.

Numeric variables can also be incremented and decremented by ++ and -- statements:

```
v := 1 
v++ // same as v = v + 1 becomes 2
v-- // same as v = v - 1 becomes 1 again

```

#### 2.4.1 Tuple Assignment

tuple asignment, allows several variables to be assigned at once.

All of the right-hand side expressions are evaluated before any of the variables are updated, making this form most useful when some of the variables appear on both sides of the assignment, as happens, for example, when swapping the values of two variables:

```
x, y = y, x
a[i], a[j] = a[j], a[i]

```

A function with multiple results, produce several values.

```
f, err = os.Open("foo.txt") // function call returns two values

// a map lookup
v, ok = m[key]

// type assertion (section 7.10)
v, ok = x.(T)

// channel receive (section 8.4.2)
v, ok = <-ch
```

As with variable declaration, we can assign unwanted values to the blank identifier:

```
_, err = io.Copy(dst, src) // discard byte count
_, ok = x.(T) check type but discard result

```

### 2.4.2 Assignability

Assignment statements are an explicit form of assignment.

An assignment occurs implicity: a function call implicitly assigns the argument values to the corresponding parameter variables; a return statement implicitly assigns the return operands to the corresponding result variables; and a literal expression for a composite type (section 4.2) such as this slice:

```
medals := []string{"gold", "silver", "bronze"}
```

### 2.5 Type Declaration

The type of a variable or expression defines the characteristics of the values it may take on, such as their size (number of bits or number of elements, perhaps), how they are represented internally, the intrinsic operations that can be performed on them, and the methods associated with them.

**type name underlying-type**

type conversion like Celsius(t) or Fahrenheit(t).

For every type T, there is a corresponding conversion operation **T(x)** that converts the value x to type T.

(both have the same underlying type, or if both are unamed pointer types that point to variables of the same underlying type.)

Arithmetic operator and comparison operators 

```
type Celsius float64
type Fahrenheit float64

var c Celsius
var f Fahrenheit

fmt.Println(c == Celsius(f)) // true
```

type conversion `Celsius(f)` does not change the value of its argument, just its type. The test is true because c and f are both zero.

Named types also make it possible to define new behaviors for values of the type. [**new behaviors**]

The declaration below, in which Celsius parameter c appears before the function name, associate with the Celsius type a method named String that return c's numeric value followed by  °C.

> func (c Celsius) String() string {return fmt.Sprintf("%g\ °C", c)}


### 2.6 Packages and Files

the range loop uses only the index, it could be written as

`for i := range pc{`

equivalent to 

`for i, _ := range pc{`



localhost variable and short hand variable statment (:=)

```
var cwd string

func init() {
    cwd, err := os.Getwd() // compile error: unused: cwd
    if err != nil {
        log.Fatal("os.Getwd failed: %v", err)
    }
}
```

the := statement declares both of them(cwd, err) as local variables.


```
var cwd string

func init() {
    var err error
    cwd, err = os.Getwd()

    if err != nil {
        log.Fatal("os.Getwd failed: %v", err)
    }
}
```

How packages, files, declarations, and statements express the structure of programs.

## ch3 Basic Data Types

(the structure of data.)

Go's type fall into categories: basic types, aggregate types, reference types, and interface types.


Basic types: numbers, strings, and booleans

Aggregate types: arrays, structs, form more complicated data types by combining values of several simple ones.

Reference types: pointers, slices, maps, functions, and channels, but what they have in common is that they all refer to program variables or state indirectly, so that the effect of ann operation applied to one reference is observed by all copies of that reference.

### 3.1 Integer


**rune** is an synonym for int32 and conventionally indicates that a value is a Unicode code point.

**byte** is an synonym for uint8.

Go's binary operators for arithmetic, logic, and comparison are listed here in order of decreasing precedence:

```
 * / % <<  >> & &^
 + - | ^
 == != < <= >  >=
 &&
 ||
```


The remainder operator & applies only to integer.

In Go, the sign of the remainder is always the same as the sign of the dividend, so -5%3 and -5%-3 are both -2. (被除数（dividend）)

Thre behavior of  / depends on whether its operands are integers, __so 5.0/4.0 is 1.25, but 5/4 is 1__ because integer division truncates the result toward zero.


```
var u uint8 = 255

fmt.Println(u + 1, u*u) // 0 1

// uint8: 0 ~ 255
// 为什么 u*u 结果为 1 呢？

```

Two integeers of the same type may be compared using the binary comparison operators below; the type of a comparison expression is a boolean.

 == equal to
 != not equal to
 < less than
 <= less than or equal to
 > greater than
 >= greater than or equal to


 [Unary operations] As unary operations have only one operand they are evaluated before other operations containing them.


 Bitwise binary operators

 &  bitwise AND
 |  bitwise OR
 ^  bitwise XOR
 __&^ bitwise clear (AND NOT) ??__
 << left shift
 >> right shift

 The operator ^ is bitwise exclusive OR (XOR) when used as a binary operator, but when used as a unary operator it is bitwise negation or complement; that is, it returns a value with each bit in its operand inverted.

 The &^ operator is bit clear (AND NOT): in the expression `z=x &^ y`, each bit of z is 0 if the corresponding bit of y is 1; otherwise it equals the corresponding bit of x.

 &^ 将运算符[左边数据]相异的位保留，相同位清零。（左操作数？）
功能与 a&(^b) 相同。

 ```
fmt.Println(0&^0) // 0
fmt.Println(0&^1) // 0
fmt.Println(1&^0) // 1
fmt.Println(1&^1) // 0
 ``` 


In general, an explicit conversion is required to convert a value from one type to another, and binary operators for arithmetic and logic (except shift) must have operands of the same type.



### 3.2 Floating-Point Numbers

A float32 provide approximately six decimal digits of precision, whereas a float64 provides about 15 digits.


Any comparison with NaN yields false.

> nan := math.NaN()

(0/0 or Sqrt(-1))

If a  function that returns a floating-point result might fail it's better to report the failure separately, likt this:

```
func compute() (value float64, ok bool) {
    // ...

    if failed {
        return 0, false
    }
    return result, true
}
```


### 3.3 Complex Numbers

Two sizes of complex numbers, complex64 and complex128.

Built-in real and imag function extract those components:

```
var x complex128 = complex(1,2)
var y complex128 = complex(3,4)
fmt.Println(x*y)
fmt.Println(real(x*y))
fmt.Println(imag(x*y))
```


### 3.4 Booleans

Boolean values can be combined with the && (AND) and || (OR) operators, which short-circuit behavior: if the answer is already determined by the value of the left operand, the right operand is not evaluated, making it safe to write expressions like this


### 3.5 Strings

A string is an immutable sequence of bytes.

Text strings are conventionally interpreted as UTF-8-encoded sequences of Unicode code points (runes).

The i-th byte of a string is not necessarily the i-th character of a string, because the UTF-8 encoding of a non-ASCII code point requires two or more bytes.

