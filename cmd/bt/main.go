package main

import (
	"os"

    "github.com/smiyaguchi/bt/internal/parser"
)

func main() {
    fi, err := os.Stdin.Stat()
    if err != nil {
        panic(err)
    }
    if fi.Mode() & os.ModeNamedPipe != 0 {
        _, err = parser.Parse(os.Stdin)
        if err != nil {
            panic(err)
        }
    }
}
