An interface type is defined as a set of method signatures.

A method is a function with a special **receiver** argument.

```
// interface
type Abser intreface {
 Abs() float64
}
// method
type Vertex struct {
 X, Y float
}

func (v Vertex) Abs() float64 {
  return math.Sqrt(v.X * v.X + v.Y * v.Y)
}

```