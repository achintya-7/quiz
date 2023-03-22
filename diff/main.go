package main

import (
	"fmt"

	"github.com/sergi/go-diff/diffmatchpatch"
)

func main() {

	value1 := "{\"ping\":\"pong\"}"
	value2 := "{\"pingy\":\"pongy\"}"

	dmp := diffmatchpatch.New()
	diffs := dmp.DiffMain(value1, value2, false)
	fmt.Println(dmp.DiffPrettyText(diffs))
}
