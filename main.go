package main

import (
	libs "gitlab.com/nabeelsaabna/go-private-nabeel"
	_ "github.com/google/go-cmp/cmp"
)

func Sum(a, b int) int {
	return lib.Multiply(a, b)
}

func PrintSum(a, b int) {
	libs.PrintSum(a, b)
}
