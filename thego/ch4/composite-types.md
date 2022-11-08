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