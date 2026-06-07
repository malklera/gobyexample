package main

import (
	"io"
	"os"
	"testing"
)

func check2(b *testing.B, e error) {
	switch e {
	case nil:
		break
	case io.EOF:
		break
	default:
		b.Fatal()
	}
}

func genContent() string {
	lines := ""
	text := "2026/05/31 19:07:57 Success.\n"
	target := "2026/05/31 17:59:32 Error.\n"
	for n := range 5000 {
		if n == 4995 {
			lines += target
		}
		lines += text
	}
	return lines
}

func BenchmarkUsingScanner(b *testing.B) {
	path := "example.log"
	file, err := os.Create(path)
	check2(b, err)
	defer os.Remove(path)
	_, err = file.WriteString(genContent())
	check2(b, err)
	file.Close()
	fileScanner, err := os.Open(path)
	check2(b, err)
	b.ResetTimer()
	for b.Loop() {
		fileScanner.Seek(0, io.SeekStart)
		_ = usingScanner(fileScanner)
	}
}

func BenchmarkUsingSeek(b *testing.B) {
	path := "example.log"
	file, err := os.Create(path)
	check2(b, err)
	defer os.Remove(path)
	_, err = file.WriteString(genContent())
	check2(b, err)
	file.Close()
	fileSeek, err := os.Open(path)
	check2(b, err)
	b.ResetTimer()
	for b.Loop() {
		fileSeek.Seek(0, io.SeekStart)
		_ = usingSeek(fileSeek)
	}
}
