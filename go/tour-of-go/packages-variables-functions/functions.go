package main

import "fmt"

// Type comes after varialbe name
// When two or more consecutive named function parameters share a type
// you can omit the type from all but the last
func add(x, y int) int {
    return x + y
}

func main() {
    fmt.Println(add(42, 13))
}

