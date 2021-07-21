package main

import (
	"fmt"
	"strconv"
	"testing"
	"unsafe"
)

func bInt(n int8) string {
	fmt.Println(*(*uint8)(unsafe.Pointer(&n)))
	return strconv.FormatUint(uint64(*(*uint8)(unsafe.Pointer(&n))), 2)
}

func TestParseInt(t *testing.T) {
	fmt.Println(bInt(-1))

	i, err := strconv.ParseInt("1111", 2, 16)
	fmt.Println(i, err)
	i, err = strconv.ParseInt("-255", 10, 16)
	fmt.Println(i, err)
	i, err = strconv.ParseInt("4FDC", 16, 16)
	fmt.Println(i, err)
}
