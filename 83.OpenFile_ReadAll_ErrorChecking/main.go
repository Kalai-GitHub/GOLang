package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	i, err := os.Open("name.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer i.Close()
	bs, err := ioutil.ReadAll(i)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(bs))
}
