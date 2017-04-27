package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (reader *rot13Reader) Read(b []byte) (int, error) {
	newReader := (*reader).r
	n, err := newReader.Read(b)
	for index := range b{
		if (b[index] >= 65 && b[index] <= 77) || (b[index] >= 97 && b[index] <= 109) {
			b[index] += 13
		} else if (b[index] >= 78 && b[index] <= 90) || (b[index] >= 110 && b[index] <= 122) {
			b[index] -= 13
		}
	}
	return n, err
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
