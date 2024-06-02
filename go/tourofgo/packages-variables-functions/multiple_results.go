package main

import "fmt"

// A function can return any number of results
// The swap function returns two strings
func swap(x, y string) (string, string) {
    return y, x
}

func main() {
    // ":=" is a  short variable declaration
    // Declare a new variable and assign a value to it in a single step
    // Inside a function the short assignment statement can be used in place
    // of a var declaration with implicit type
    // Outside a function, every statement begins with a keyword and
    // so the ":=" construct is not available
    a, b := swap("hello", "world")
    fmt.Println(a, b)
}

