// A defer statement defers the execution of a function untile the surrounding
// function return

// The deferred call's arguments are evaluated inmediately, but the function
// call is not executed until the surrounding function returns.

package main

import "fmt"

func main() {
    fmt.Println("counting")

    // Deferred function calls are pushed onto a stack
    // When a function returns, its deferred calls are
    // executed in last-in-first-out order
    for i := 0; i < 10; i++ {
        defer fmt.Println(i)
    }

    fmt.Println("done")
}

