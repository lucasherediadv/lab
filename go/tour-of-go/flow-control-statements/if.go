package main

import (
    "fmt"
    "math"
)

func pow(x, n, lim float64) float64 {
    // The expression need not be surrounded by parentheses
    // The "if" statement can start with a short statement to execute
    // before the condition
    if v := math.Pow(x, n); v < lim {
        return v
    // Variables declares inside an "if" are alse available inside any of the
    // "else" blocks
    } else {
        fmt.Printf("%g >= %g\n", v, lim)
    }
    return lim
}

func main() {
    fmt.Println(
        pow(3, 2, 10),
        pow(3, 3, 20),
    )
}

