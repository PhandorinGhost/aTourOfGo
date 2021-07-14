package main

import (
	"io"
	"log"
	"os"
	"strings"
	"unicode"
)

func rot13(bt byte) byte {
	if !unicode.IsLetter(rune(bt)) {
		return bt
	}
	var rez byte
	if (bt >= 'a' && bt <= 'm') || (bt >= 'A' && bt <= 'M') {
		rez = bt + 13
	} else if (bt >= 'n' && bt <= 'z') || (bt >= 'N' && bt <= 'Z') {
		rez = bt - 13
	}

	return rez
}

type rot13Reader struct {
	r io.Reader
}

func (rot rot13Reader) Read(b []byte) (int, error) {
	number, err := rot.r.Read(b)
	if err != nil {
		return 0, err
	}

	for i := 0; i < number; i++ {
		b[i] = rot13(b[i])
	}

	return number, err
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	_, err := io.Copy(os.Stdout, &r)
	if err != nil {
		log.Println(err)
	}
}
