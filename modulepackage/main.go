package main

import (
	"fmt"

	print "github.com/learning-go-book/package_example/formatter"
	"github.com/learning-go-book/package_example/math"
)

func main() {
	num := math.Double(5)
	output := print.Format(num)
	fmt.Println(output)
}
