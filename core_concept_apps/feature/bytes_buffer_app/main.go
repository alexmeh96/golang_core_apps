package main

import (
	"bytes"
	"fmt"
)

func main() {
	buf := new(bytes.Buffer)
	buf.Write([]byte("foo"))
	buf.WriteString("bar")
	fmt.Println(buf.Len())
	fmt.Println(buf.String())
}
