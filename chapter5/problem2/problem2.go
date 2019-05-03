package main

import (
	"bytes"
	"fmt"
)

func btos(number float32) string {
	var buffer bytes.Buffer
	buffer.WriteString("0.")
	for i := 0; i < 32; i++ {
		number *= 2
		if number > 1 {
			buffer.WriteString("1")
			number--
		} else {
			buffer.WriteString("0")
		}

	}
	return buffer.String()
}

func main() {
	fmt.Println(btos(0.72))

}
