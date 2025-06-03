package main

import (
	"fmt"
	"shieldBuilder/shieldBuilder"
)

func main() {
	builder := shieldbuilder.NewShieldBuilder()
	shield := builder.RaiseBack().RaiseFront().Build()
	fmt.Printf("shield: %v\n", shield)
}
