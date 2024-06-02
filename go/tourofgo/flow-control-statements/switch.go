package main

import (
    "fmt"
    "runtime"
)

func main() {
    fmt.Println("Go runs on")

    // It runs the first case whose value is equal to the condition expression
    // Go only runs the selected case, not all the cases that follow
    // In effect the "break" statement is provided automatically in Go
    // Go's switch cases need not be constants
    // The values involved need not be integers
    switch os := runtime.GOOS; os {
    case "darwin":
        fmt.Println("OS X.")
    case "linux":
        fmt.Println("Linux.")
    default:
        // freebsd, openbsd,
        // plan9, windows...
        fmt.Printf("%s. \n", os)
    }
}

// Switch cases evaluate cases from top to bottom, stopping when a case succeeds
// For example

//switch i {
//case 0:
//case f():
//}

// Does not call "f" if "i==0"

