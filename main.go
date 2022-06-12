package main

import (
	"fmt"
	"github.com/egasimov/num2words/helper"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(1)
	result := helper.ConvertIntPart("70")

	fmt.Print(result)
}
