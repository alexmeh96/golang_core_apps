package closing_buffer

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"testing"
)

type BytesClosingBuffer struct {
	*bytes.Buffer
	io.Closer
}

func NewBytesClosingBuffer() *BytesClosingBuffer {
	return &BytesClosingBuffer{
		Buffer: new(bytes.Buffer),
	}
}

func (b *BytesClosingBuffer) Close() error {
	fmt.Println("closing!")
	return nil
}

func writeTo(wc io.WriteCloser, msg []byte) error {
	defer wc.Close()
	_, err := wc.Write(msg)
	return err
}

func TestBuffer(t *testing.T) {
	buf := NewBytesClosingBuffer()

	if err := writeTo(buf, []byte("Hello!")); err != nil {
		log.Fatal(err)
	}

	fmt.Println(buf.String())
}
