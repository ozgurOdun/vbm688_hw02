package main

import (
	"fmt"
	"github.com/ozgurOdun/vbm688_hw02/forwardChaining"
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
	kb := utils.InputParser("kb.txt")
	if kb == nil {
		return
	}
	fmt.Println("Starting...", time.Now())
	startTime := time.Now()
	entails, proof := forwardChaining.Entails(kb, goal)
	//fmt.Println("proof,", proof)
	if entails {
		utils.ProofPrinter(proof)
	} else {
		fmt.Println("FALSE")
	}
	fmt.Println("End...", time.Now())
	elapsed := time.Since(startTime)
	fmt.Println("Process took:", elapsed)
	utils.PrintMemUsage()
	return
}
