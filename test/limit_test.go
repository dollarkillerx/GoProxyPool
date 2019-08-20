package test

import (
	"fmt"
	"testing"
)

func TestLimit(t *testing.T) {
	ints := make(chan int, 10)

	ints <- 10
	ints <- 10

	i := len(ints)
	fmt.Println(i)
}
