package main

import (
	"fmt"
	"os"
	"strconv"
	"github.com/lucasherediadv/lab/go/unq/unq"
)

const helpMessage = `A collection of universal unique identifiers.

USAGE
  unq <command>

COMMANDS
  uuid:    Generates a standard UUID v4.
  uid32:   Generates a 32-character Base32 ID.
  isosec:  Generates an ISO8601 timestamp (no punctuation).
  isosect: Generates an ISO8601 timestamp with a 'T'.
  isodate: Generates a readable ISO8601 timestamp.
  isonan:  Generates a nanosecond-granularity timestamp.
  epoch:   Generates the current Unix epoch time in seconds.
  randhex: Generates a random hexadecimal string.

EXAMPLES
  $ unq uuid
  $ unq epoch
  $ unq isosec

`

func main() {
	if len(os.Args) < 2 || os.Args[1] == "help" || os.Args[1] == "-h" {
		fmt.Print(helpMessage)
		return
	}

	command := os.Args[1]
	switch command {
	case "uuid":
		fmt.Println(unq.UUID())
	case "uid32":
		fmt.Println(unq.Base32())
	case "isosec":
		fmt.Println(unq.Isosec())
	case "isosect":
		fmt.Println(unq.IsosecT())
	case "isodate":
		fmt.Println(unq.Isodate())
	case "isonan":
		fmt.Println(unq.Isonan())
	case "epoch":
		fmt.Println(unq.Second())
	case "randhex":
		if len(os.Args) < 3 {
			fmt.Println("Usage: unq randhex <count>")
			os.Exit(1)
		}
		count, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("Error: count must be an integer.")
			os.Exit(1)
		}
		fmt.Println(unq.Hex(count))
	default:
		fmt.Printf("Unknown command: %s\n", command)
		fmt.Print(helpMessage)
		os.Exit(1)
	}
}
