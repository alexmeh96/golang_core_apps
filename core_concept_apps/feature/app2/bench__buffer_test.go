package main

import (
	"bytes"
	"testing"
)

func writeToBuffer1(msg []byte) {
	buf := new(bytes.Buffer)
	buf.Write(msg)
}

func writeToBuffer2(buf *bytes.Buffer, msg []byte) {
	buf.Write(msg)
}

func BenchmarkWriteToBuffer1(b *testing.B) {
	msg := []byte("Foo")
	for i := 0; i < b.N; i++ {

		for i := 0; i < 100; i++ {
			writeToBuffer1(msg)
		}
	}
}

func BenchmarkWriteToBuffer2(b *testing.B) {
	msg := []byte("Foo")
	buf := new(bytes.Buffer)

	for i := 0; i < b.N; i++ {

		for i := 0; i < 100; i++ {
			writeToBuffer2(buf, msg)
		}
	}
}
