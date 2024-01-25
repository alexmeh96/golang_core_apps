package buffer

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"testing"
)

func writeTo(w io.Writer, msg []byte) error {
	_, err := w.Write(msg)
	return err
}

func TestBuffer(t *testing.T) {
	buf := new(bytes.Buffer)
	if err := writeTo(buf, []byte("Hello!")); err != nil {
		log.Fatal(err)
	}

	fmt.Println(buf.String())
}
