// Every Go program is made up of packages
// Programs start runing in package main
package main

// Can also write:
// import "fmt"
// import "math"
// But is is good style to use the factored import statement "import ()"
import (
    // The package name is the same as the last element of the import path
    "fmt"
    "math/rand"
)

func main() {
    // A name is exported if it begins with a capital letter
    // "rand.intn" will not be exported
    fmt.Println("My favorite number is", rand.Intn(10))
}

