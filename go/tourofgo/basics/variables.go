package main

import "fmt"

// The var statement declares a list of variables
// As in function argument lists, the type is last
// var c, python, java bool

// A var declaration can include initializers, one per variable
var i, j int = 1, 2

func main() {
    // A var statement can be at package or function level
    // var i int

    // If an initializer is present, the type can be omitted
    // The variable will take the type of the initializer
    var c, python, java = true, false, "no!"

    fmt.Println(i, j, c, python, java)
}

