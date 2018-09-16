package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type SubObj struct {
	a float64
	b float64
}

type Test struct {
	Foo       float64
	Bar       float64
	Greeting  string
	SubObject SubObj
}

func main() {
	var (
		rr      Test
		textObj = []byte("{\"foo\":11,\"bar\":144,\"greeting\":\"Hello, World!\",\"subObject\":{\"a\":1,\"b\":2}}")
	)
	err := json.Unmarshal(textObj, &rr)
	if err != nil {
		fmt.Printf("error occured %s", err)
		os.Exit(1)
	}

	fmt.Printf("data:\n%#v\n", rr)
}
