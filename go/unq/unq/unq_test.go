package unq_test

import (
	"fmt"
	"log"

	"github.com/lucasherediadv/lab/go/unq/unq"
)

func ExampleBytes() {
	bytes := fmt.Sprintf("%x", unq.Bytes(4))
	fmt.Println(len(bytes))

	// Output:
	// 8
}

func ExampleUUID() {
	uuid := unq.UUID()
	fmt.Println(len(uuid))

	// Output:
	// 36
}

func ExampleBase32() {
	uid32 := unq.Base32()
	fmt.Println(len(uid32))

	// Output:
	// 32
}

func ExampleHex() {
	hex1 := unq.Hex(1)
	hex3 := unq.Hex(3)
	hex18 := unq.Hex(18)
	fmt.Println(len(hex1))
	fmt.Println(len(hex3))
	fmt.Println(len(hex18))

	// Output:
	// 2
	// 6
	// 36
}

func ExampleSecond() {
	sec := unq.Second()
	fmt.Println(len(sec))
}

func ExampleIsosec() {
	sec := unq.Isosec()
	log.Print(sec)
}

func ExampleIsonan() {
	sec := unq.Isonan()
	log.Print(sec)
}

func ExampleIsodate() {
	d := unq.Isodate()
	log.Print(d)
}
