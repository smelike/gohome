The basic types that serve as building blocks for data structures in a Go program; they are the atoms of our universe.

The molecules(分子) created by combining the basic types in various way.

Four such types - arrays, slices, maps and structs.

How structured data using these types can be encoded as and parsed from JSON data and used to generate HTML from templates.

Arrays and structs are aggregate types; their values are concatenations pf other values in memory.

Arrays are homogeneous - their elements all have the same type.
[homogeneous, 同质]

Whereas structs are heterogeneous.
[heterogeneous, 异质]

[fixed size]Both arrays and structs are fixed size. 
[dynamic data structures]In contrast, slices and maps are dynamic data structures that grow as values are added.


[size]int
[size]byte
[size]string
[...]int{1, 2, 3}

"..." appears in place of the length, the array length is determined by the number of initializers. For example:

```
q := [...]int{1, 2, 3}
fmt.Printf("%T \n", q) // %T -> type

```

The size of an array is part of its type, so [3]int and [4]int are different types. Thr size must be a constant expression, that is, an expression whose value can be computed asa the program is being compiled.


- %x to print all the elements in hexadecimal

- %t to show boolean

- %T to display the type of a value


When a function is called, a copy of each argument value is assigned to the corresponding parameter variable, so the function receives a copy, not the original. Passing large arrays in this way can be inefficient, and any changes that the function makes to array elements affect only the copy, not the original.

**Implicitly pass arrays by reference.** For example:

Of course, we can explicitly pass a pointer to an array so that any modifications the function makes to array elements will be visible to the caller. This function zeroes the contents of a [32]byte array:

```
func zero(ptr *[32]byte) {
    for i := range ptr {
        ptr[i] = 0
    }
}
```

The array literal [32]byte{} yields an array of 32 bytes. Each element of the array has the zero value for byte, which is zero. We can use that fact to write a different version of zero:

```
func zero(ptr *[32]byte) {
    *ptr = [32]byte{}
}
```

Using a pointer to an array is efficient and allows the called function to mutate the caller's variable, but arrays are still inherently inflexible because of their fixed size.

The zero function will not accept a pointer to a [16]byte variable, for example, nor is there any way to add ot remove array elements.

For these reasons, other than special cases like SHA256's fixed-size hash, arrays are seldom used as function parameters; instead, we use slices.


Exercise 4.1: Write a function that counts the number of bits that are different in twp SHA256 hashes.

Exercise 4.2: Write a program that prints the SHA256 hash of its standard input by default but supports a command-line flag to print the SHA384 or SHA512 hash instead.


4.2. Slices

Slices represent variable-length sequences whose elements all have the same type.

A slice type is written []T, where the elements have type T; it looks like an array type without a size.

[indices 是 index 的复数形式(plural noun)]


Slices are not comparable, we cannot use == to test whether two slices contain the same elements.

**bytes.Equal** function compare two slices of bytes ([]byte).

But for other type of slice, we must do the comparison ourselves.




Whether a slice is empty, use len(s) == 0, not s == nil.

```
var s []int

s = []int{} // s != nil
s = nil
s = []int(nil) // conversion expression 

fmt.Println(s == nil)
```

The built-in function make creates a slice of specified element type, length, and capacity. The capacity argument may be omitted, in which case the capacity equals the length.

```
make([]T, len)
make([]T, len, cap) // same as make([]T, cap)[:len]

```

Updating the slice variable is required not just when calling append, but for any function that may change the length or capacity of a slice or make it refer to a different underlying array.

To use slices correctly, it's important to bear in mind that although the elements of the underlying array are indirect, the slice's pointer, length, and capacity are not. 

To update them requires an assignment like the one above. In this respect, slices are not "pure" reference types but resemble an aggregate type such as this struct:

```
type IntSlice struct {
    ptr *int
    length, cap int
}
```


an array pointer ??


### 4.3 Map

`map[K]V`, where K and V are the types of its keys and values.

All of the keys in a given map are of the same type, and all of the values are of the same type, but the keys need not be of the same type as the values.

[the key type K's restrictions]: The key type K must be comparable using ==, so that the map can test whether a given key is equal to one already within it. Though floating-point numbers are comparable, it's a bad idea to compare floats for equality, especially bad if NaN if a posible value.

[the value type V]: There  are no restrictions on the value type V.


```
ages := map[string]int{
    "alice": 30,
    "charlie": 34,
}

// ages := make(map[string]int)
```

[*] A map element is not a variable, and we cannot take its address:

> _ = &ages["bob"] // compile error: cannnot take the address of map element

[*] Storing to a nil map causes a panic:
You must allocate the map before you can store into it.

```
var ages map[string]int // a nil map

ages["bob"] = 90 // panic: assginment to entry in nil map

ages == nil // true

##2
var ages = make(map[string]int) // it's a not nil map.
ages == nil // false

```

How to distinguish between a nonexistent element and an element that happens to have the value zero?

[an nonexistent element and the zero value]

```
age, ok := ages["bob"]

if !ok { /* not a key in this map; */}

// two statements combined, like this

if age, ok := ages["bob"]; !ok {
    // not a key in this map
}
```


The value type of a map itself be a composite type, such as a map or slice.

```
var graph = make(map[string]map[string]bool)

func addEdge(from, to string) {
    edges := graph[from]
    if edges == nil {
        edges = make(map[string]bool)
        graph[from] = edges
    }
    edges[to] = true
}

func hasEdge(from, to string) bool {
    return graph[from][to]
}
```

### 4.4 structs

A struct is an aggregate data type that groups together zero or more named values of arbitrary
types as a single entity.

[A struct -> an aggregate data type - groups together zero or more named values of arbitrary types as a single entity.]

Each value is called a field.

The classic example of a struct from data processing is the employee record, whose fields are a unique ID, the employee's name, address, date of birth, position, salary, manager, and the like.
[from data processing - a unique ID, name, address, date of birth, position, salary, manager, and the like]

All of these fields are collected into a single entity that can be copied as a unit, pass to functions and returned by them, stored in arrays, and so on.

These two statements declare a struct type called `Employee` and a variable called `dilbert` that is an instance of an Employee:

```
type Employee struct {
    ID int
    Name string
    Address string
    DoB time.Time
    Position string
    Salary int
    ManagerID int
}

var dilbert Employee

```

The zero value of a struct is composed of the zero valuees of each of its fields. It is usually desirable that the zero value be a natural or sensible default. For example, in bytes.Buffer, the initial value of the struct is a ready-to-use empty buffer, and the zero value of sync.Mutex is a ready-to-use unlocked mutex. Sometimes this sensible initial behavior happens for free, but sometimes the type designer has to work at it.

The struct type with no fields is called the empty struct, written struct{}. It has size zero and carries no information but may be useful nonetheless. Some Go programmers use it instead of bool as the value type of a map that represents a set, to emphasize that only the keys are significant, but the space saving is marginal and the syntax more cumbersome, so we generally avoid it.

`empty struct: struct{}`, no fileds.

```
// the key type is string, the value type is empty struct.
seen := make(map[string]struct{}) 

// ...
if _, ok := seen[s]; !ok {
    seen[s] = struct{}{}
    // ... first time seeing s...
}
```

### 4.4.1 Struct Literals

A value of a struct type can be written using a struct literal that specifies values for its fields.

```
type Point struct { X, Y int }
p := Point{1, 2}

```

There are two forms of struct literal. The first form, shown above, requires that a value be specified for every field, in the right order.

It burdens the writer (and reader) with remembering exactly what the fields are, and it makes the code fragile should the set of fields later grow or be reordered.

Accordingly, this form tends to be used only within the package that defines the struct type, or with smaller struct type for which there is an obvious field ordering convention, like `image.Point{x, y}` or `color.RGBA{red, green, blue, alpha}`.

[within the package that defines the struct type; 
smaller struct type for which there is an obvious field ordering convention]


The second form is used, in which a struct value is initialized by listing some or all of the field names and their corresponding values.
For example, `anim := gif.GIF{LoopCount: nframes}`.

If a field is omitted in this kind of literal, it is set to the zero value for its type. Because names are provided, the order of fields doesn't matter.


The two forms cannot be mixed in the same literal. Nor can you use the (order-based) first form of literal to sneak around the rule that `unexported identifiers` may not be referred to from another package.

```
package p
typpe T struct{a, b int} // a and b are not exported

package q
import "p"

var _ = p.T{a: 1, b:2}  // compile error: can't reference a, b
var _ = p.T{1, 2}   // compile error: can't reference a, b

```

Struct values can `be passed as arguments` to functions and returned from them. For instance, this function scales a `Point` by a specified factor:

```
func Scale(p Point, factor int) Point {
    return Point{p.X * factor, p.Y * factor}
}
fmt.Println(Scale(Point{1,  2}, 5))
```


For efficiently, larger struct type are usually passed to or returned from functions indirectly using a pointer,

Since in a `call-by-value` language like Go, the called function receives only a copy of an argument, not a reference to the original argument.


Shorthand notation to create and initialize a struct variable and obtain its address:

`pp := &Point{12, 4}` it is exactly equivalent to 

```
pp := new(Point)
*pp = Point{1, 2}
```

`&Point{1, 2}` can be used directly within an expression, such as a function call.


### 4.4.2 Comparing Structs

[Definition]If all the fields of a struct are comparable, the struct itself is comparable.


`==` operation compares the corresponding fields of the two structs in order.

Comparable struct types, like other comparable types, may be used as the key type of a map.

```
type address struct {
    hostname string
    port int
}

hits := make(map[address]int)
hits[address{"golabg,org", 443}]++
```


### 4.4.3 Struct Embeding and Anonymous Fields

Struct embedding mechanism lets us use one named struct type as an anonymous field of another struct type, providing a convenient syntactic shortcut so that a simple dot expression like `x.f` can stand for a chain of fields like `x.d.e.f`.

Consider a 2-D drawing program that provides a library of shapes, such as rectangles, ellipses, stars, and wheels. Here are two of the types it might define:

```
type Circle struct {
    X, Y, Radius int
}

type Wheel struct {
    X, Y, Radius, Spokes int
}
```

A circle has fields for the X and Y coordinates of its center, and a Radius. A wheel has all the features of a Circle, plus Spokes, the number of inscribed radial spokes. Let's create a wheel:

```
var w wheel
w.X = 8 // this is statement, not declaration
w.Y = 8 // the same as above
w.Radius = 5
w.Spokes = 20

```

As the set of shapes grows, we're bound to notice similarities and repetition among them, so it may be convenient to factor out their common parts:

```
type Point struct {
    X, Y int
}

type Circle struct {
    Center Point
    Radius int
}

type Wheel struct {
    Circle Circle
    Spokes int
}

```

[Anonymoua fields] Declare a field with a type but no name; such field are called anonymous fields. The type of the field must be a named type or a pointer to a named type.
[A named type or a pointer to a named type]

[intervening name, 中间名]refer to the names at the leaves of the implicit tree without giving the intervening names:

`w.X = 42 // instead of w.Circle.Point.X`

> fmt.Printf("%#v\n", w)

Notice how the `# adverb` causes Printf's %v to display values in a form similar to Go syntax. For struct values, this form includes the name of each field.

[two anonymous fiels of the same type-conflict]Because "anonymous" fields do have implicit names, you can't have `two anonymous fields of the same type` since their names would conflict. And because the name of the field is implicitly determined by its type, so too is the visibility of the field. In the examples above, the Point and Circle anonymous fields are exported. Had they been unexported (point and circle), we could still use the shorthand form

`w.X = 8 // equivalent to w.circle.point.X = 8`

but the explicit long form shown in the comment would be forbidden outside the declaring package because circle and point would be inaccessible.

What we've seen so far of struct embedding is just a sprinkling of syntactic sugar on the dot notation used to select struct fields. Later, we'll see that anonymous fields need not be struct types; any named type or pointer to a named type will do. But why would you want to embed a type that has no subfields?

[the outer struct type gains not just the fields of the embedded type but its methods too.]The answer has to do with methods. The shorthand notation used for selecting the fields of an embedded type works for selecting its methods as well. In effect, **the outer struct type gains not just the fields of the embedded type but its methods too**. This mechanism is the main way that complex object behaviors are composed from simpler ones. Compisition is central to object-oriented programming in Go.



### 4.5 JSON

JavaScript Object Notation (JSON) is a standard notation for sending and receiving structured information.

Go has excellent support for encoding and decoding these formats, provided by the standard library package encoding/json, encoding/xml, encoding/asn1, and so on, and these packages all have similar APIs.


JSON is an encoding of JavaScript values - strings, numbers, booleans, arrays, and objects - as Unicode text. It's an efficient yet readable representation for the basic data types of Chapter 3 and the composite types of this chapter - arrays, slices, structs, and maps.

[**]JSON's \Uhhhh numeric escapes denote UTF-16 codes, not runes.


- json.Marshal

- json.MarshalIndent

Only exported fields are marshaled.

The field tags, a field tag is a string of metadata associated at compile time with the field of a struct:

```
Year int `json:"released"`
Color bool `json:"color,omitempty"`

```
A field tag may be any literal string, but it is conventionally interpreted as a space-separated list of key:"value" pairs since they contain double quotation marks, field tags are usually written with raw string literals. 

The json key controls the behavior of the encoding/json package, and other enconding/... packages follow this convention.

The first part of the json field tag specifies an alternative JSON name for the Go field. Field tags are often used to **specify an idiomatic JSON name like total_count** for a Go field named TotalCount.

The tag for Color has an additional option, omitempty, which indicates that no JSON output should be produced if the field has the zero value for its type (false, here) or is otherwise empty. Sure enough, the JSON output for Casablanca, a black-and-white movie, has no color field.


### 4.6. Text and HTML Templates

The previous example does only the simplest possible formatting, for which `Printf` is entirely adequate. But sometimes formatting must be more elaborate, and it's desirable to separate the format from the code more completely. This can be done with the text/template and html/template packages, which provide a mechanism for **substituting** the values of variables into a text or HTML template.

A template is a string or file one or more portions enclosed in double braces, `{{...}}`, called actions. Most of the string is printed literally, but the actions trigger other behaviors.

Each action contains an expression in the template language, a simple but powerful notation for printing values, selecting struct fields, calling functions and methods, expressing control flow such as if-else statements and range loops, and instantitating other templates.

A simple template string is shown below:

```
> issuereport

const templ = `{{.TotalCount}} issues
{{range .Items}}------------------------------
Number: {{.Number}}
User:   {{.User.Login}}
Title:  {{.Title | printf "%.64s"}}
Age:    {{.CreateedAt | daysAgo}} days
{{end}}`

```