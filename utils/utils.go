package utils

import (
	"fmt"
	"os"
	"runtime"
	"strings"
)

func checkErr(e error) {
	if e != nil {
		panic(e)
	}
}

func InputParser(fileName string) []string {
	board := make([][]int, 3)
	for i := 0; i < 3; i++ {
		board[i] = make([]int, 3)
	}

	file, err := os.Open(fileName)
	checkErr(err)

	rawInput := make([]byte, 100)
	n, err := file.Read(rawInput)
	if n > 0 {
		fmt.Println("Raw input is: ", string(rawInput))
	}
	var inpStr []string
	inputLines := strings.Split(string(rawInput), "\n")
	fmt.Println(len(inputLines))
	for i := 0; i < len(inputLines); i++ {
		fmt.Printf("Input line %d is %s\n", i, inputLines[i])
		inpStr = strings.Split(inputLines[i], " ")
		fmt.Println("len of inpStr is:", len(inpStr))

	}

	return inpStr
}

/////// MEMORY METHODS
func PrintMemUsage() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	// For info on each, see: https://golang.org/pkg/runtime/#MemStats
	fmt.Printf("Alloc = %v MiB", bToMb(m.Alloc))
	fmt.Printf("\tTotalAlloc = %v MiB", bToMb(m.TotalAlloc))
	fmt.Printf("\tSys = %v MiB", bToMb(m.Sys))
	fmt.Printf("\tNumGC = %v\n", m.NumGC)
}

func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}
