package main

import (
	"fmt"
	"os"

	"github.com/smiyaguchi/bt/internal/parser"
)

func main() {
	fi, err := os.Stdin.Stat()
	if err != nil {
		panic(err)
	}
	if fi.Mode()&os.ModeNamedPipe != 0 {
		t, err := parser.Parse(os.Stdin)
		if err != nil {
			panic(err)
		}
		for _, r := range t.Rows {
			for k, v := range r.Value {
				fmt.Printf("key: %s, value: %s\n", k, v)
			}
		}
	}
}
