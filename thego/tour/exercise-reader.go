package main

import "golang.org/x/tour/reader"

type MyReader struct{}

// ToDo: Add a Read([]byte) (int, error) method to MyReader
// byte    // alias for uint8
// rune    // alias for int32

// the ASCII character 'A'
func (r MyReader) Read(b []byte) (int, error) {
	for x := range b {
		b[x] = 'A'
		// b[x] = "A"
		//  cannot use "A" (untyped string constant) as byte value in assignment
	}
	return len(b), nil
}

func main() {
	reader.Validate(MyReader{})
}
