package main

import (
	"fmt"

	"github.com/chris92vr/port-scanning-go/port"
)

func main() {
	fmt.Println("Port Scanning")
	results := port.InitialScan("localhost")
	fmt.Println(results)
}
