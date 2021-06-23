package parser

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "time"
)

type Data struct {
    Rows []Row
}

type Row struct {
    Key string
    ColumnFamilys []ColumnFamily
}

type ColumnFamily struct {
    Name string
    Columns []Column
}

type Column struct {
    Name string
    Cells []Cell
}

type Cell struct {
    Value string
    Timestamp time.Duration
}

func Parse(r io.Reader) (*Data, error) {
    scanner := bufio.NewScanner(os.Stdin)
    for scanner.Scan() {
        fmt.Println(scanner.Text())
    }
    if err := scanner.Err(); err != nil {
        return nil, fmt.Errorf("error occured to scan: %v", err)
    }
    return nil, nil
}
