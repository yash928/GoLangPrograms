package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{}

func (logWriter) Write(b []byte) (int, error) {
	cont := string(b)
	fmt.Println(cont)
	fmt.Println("Number of bytes read=", len(cont))
	return len(cont), nil
}

func main() {
	res, err := http.Get("https://google.com")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	lw := logWriter{}
	io.Copy(lw, res.Body)
}
