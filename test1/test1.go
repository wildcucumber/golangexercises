package main

import (
	"fmt"
	"encoding/hex"
)

func main() {
	var (
		i, j, k int
	)
	fmt.Println(hex.Dump([]byte("Dumper returns a WriteCloser that writes a hex dump of all written data to w. The format of the dump matches the output of `hexdump -C` on the command line.")), i, j, k)
}
