package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	f, err := os.Create("test.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()

	r := strings.NewReader("wassup guys")

	io.Copy(f, r)

}
