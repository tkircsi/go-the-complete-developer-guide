package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{}

func main() {
	resp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	// fmt.Println(resp)
	// bs := make([]byte, 15000)
	// n, err := resp.Body.Read(bs)
	// resp.Body.Close()
	// fmt.Println("Nr of bytes read:", n)
	// fmt.Println("Body:", string(bs))

	// file, _ := os.Create("google")
	// io.Copy(file, resp.Body)
	// io.Copy(os.Stdout, resp.Body)
	lw := logWriter{}
	io.Copy(lw, resp.Body)
}

func (logWriter) Write(bs []byte) (int, error) {
	fmt.Print(string(bs))
	return len(bs), nil
}
