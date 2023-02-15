package main

import (
	"fmt"
	"k8s-dm/dm3"
	"strconv"
	"time"
)

func main() {
	fmt.Println("0. <Main function>")
	dm3.Dm3()
	fmt.Println("time is: ", time.UnixDate)
	a := "123"
	i, _ := strconv.ParseInt(a, 10, 32)
	b := int32(i)
	fmt.Println("conv string to int32, ", b)
}
