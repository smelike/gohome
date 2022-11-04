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



