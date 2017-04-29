package print_test

import (
    "fmt"
    p "github.com/kwmt/go-print"
)

func ExamplePrint() {
    fmt.Println(p.Print("そうだ京都へ行Go"))
    // Output:
    // そうだ京都へ行Go
}