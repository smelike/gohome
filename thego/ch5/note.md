## ch5 Functions

A function lets us `wrap up a sequence of statements` as a unit that can be called from elsewhere in a program, perhaps multiple times.
[wrap up a sequence of statements as a unit]

Functions make it possible to break a big job into smaller pieces that might well be written by different people separated by both time and space.

A function hides its implementation details from its users. For all of these reasons, functions are a critical part of any programming language.


We've seen many functions already. Now let's take time for a more thorough discussion. The running example of this chapter is a web crawler,
that is, the component of a web search engine responsible for fetching web pages, discovering the links within them, fetching the pages identified by those links, and so on. A web crawler gives us ample opportunity to explore recursion, anonymous functions, error handling, and aspects of functions that are unique to Go.

a web crawler: fetch web pages, discover the links within pages.

aspects of function: recursion, anonymous, error handling and aspects of functions

### 5.1. Function Declarations

A function declaration has a name, a list of parameters, an optional list of results, and a body:

```
func name(parameter-list) (result-list) {
    body
}
```

The parameter list specifies the names and types of the function's parameters, which are the local variable whose values or arguments are supplied by the caller. The result list specifies the types of the values that the funtion returns. If the function returns one unamed result or no results at all, parentheses are optional and usually omitted. Leaving off the result list entirely declares a function that does not return any value and is called only for its effects. In the hypot function,

```
func hypot(x, y float64) float64 {
    return math.Sqrt(x*x + y*y)
}
fmt.Println(hypot(3, 4))
```

**x and y are parameters in the declaration**, **3 and 4 are arguments of the call**, and the function returns a float64 value.


The type of a function is sometimes called its signature. Two functions have the same type or signature if they have the same sequence of parameter types and the same sequence of result types. The names of parameters and results don't affect the type, nore does whether or not they were declared using the factored form.

Every function call must provide an argument for each parameter, in the order in which the parameters were declared. Go has no concept of default parameter values, nor any way to specify arguments by name, so the names of parameters and results don't matter to the caller except as documentation.

Parameters are local variable within the body of the function, with their initial values set to the arguments supplied by the caller. Function parameters and named results are variables in the same lexical block as the function's outermost local variables.
[in the same lexical block as the function's outermost local variables]


Arguments are passed by value, so the function receives a copy of each argument; modifications to the copy do not affect the caller. However, if the argument contains some kind of reference, like a pointer, slice, map, function, or channel, then **the caller may be affected** by any modifications the function makes to variables indirectly referred to by the argument.

[a function declaration without body]You may occasionally encounter a function declaration without a body, indicating that the function is implemented in a language other than Go. Such a declaration defines the function signature.

[other than, 以外]

```
package math

func Sin(X float64) float64 // implemented in assembly language

```

### 5.2. Recursion

[recursive-data-struture]Functions may be recursive, that is, they may call themselves, either directly or indirectly. Recursion is a powerful technique for many problems, and of course it's essential for processing recursive data struture. 

In Section 4.4, we used recursion over a tree to implement a simple insertion sort. In this section, we'll use it again for processing HTML documents.

The example program below uses a non-standard package, golang.org/x/net/html, which provides an HTML parser. The golang.org/x/... respositories hold package designed and maintained by the Go team for applications such as networking, internationalized text processing, monbile platforms, image manipulation, cryptography, and developer tools. These packages are not in the standard library because they're still under development or because they're rarely needed by the majority by the majority of Go programmers.

The parts of the golang.org/x/net/html API that we'll need are show below. The function html.Parse reads a sequence of bytes, parse them, and returns the root of the HTML document, which is an html.Node.

HTML has several kinds of nodes - text, comments, and so on - but there are concerned only with element nodes of the form `<name key='value'>`.

```
// golang.org/x/net/html

type Node struct {
    Type NodeType
    Data string
    Attr []Attribute
    FirstChild, NextSibling *Node
}

type NodeType int32

const (
    ErrorNode NodeType = iota
    TextNode
    DocumentNode
    ElementNode
    CommentNode
    DoctypeNode
)

type Attribute struct {
    Key, Val string
}

func Parse(r io.Reader) (*Node, error)

```