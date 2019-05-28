package main

import (
	"fmt"
	"github.com/ozgurOdun/vbm688_hw02/utils"
	"os"
	"time"
)

func main() {
	utils.PrintMemUsage()
	args := os.Args
	if len(args) != 2 {
		fmt.Println("Usage: hw02 {goal}")
		return
	}

	var goal string
	goal = args[1]
	fmt.Println("Goal is ", goal)
	fmt.Println("Starting...", time.Now())
	startTime := time.Now()

	board := utils.InputParser("kb.txt")
	if board == nil {
		return
	}

	fmt.Println("End...", time.Now())
	elapsed := time.Since(startTime)
	fmt.Println("Process took:", elapsed)
	utils.PrintMemUsage()
	return
}
