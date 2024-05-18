package main

import "fmt"

func main() {
  sum := 1
  // Go has only one looping construct, the "for" loop
  // Unlike other languages there are no parentheses surrounding the
  // three components of the for statement

  // for i := 0; i < 10; i++ {
    // sum += i
  // }

  // The init and post statements are optional
  // At that point you cand drop the semicolons
  // C's "while" is spelled "for" in Go
  for sum < 1000 {
    sum += sum
  }
  
  // If you omit the loop condition it loops forever
  // for {
  // }

  fmt.Println(sum)
}

