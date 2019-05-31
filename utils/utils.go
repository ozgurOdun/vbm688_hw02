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
	file, err := os.Open(fileName)
	checkErr(err)

	rawInput := make([]byte, 100)
	n, err := file.Read(rawInput)
	if n > 0 {
		fmt.Println("Raw input is: ", string(rawInput))
	}
	inputLines := strings.Split(string(rawInput), "\r\n")
	fmt.Println(len(inputLines))
	return inputLines
}

func IsGoal(goal, query string) bool {
	if strings.Compare(goal, query) == 0 {
		return true
	}
	return false
}

func IsTrue(query string) bool {
	if strings.HasPrefix(query, "!") {
		return true
	}
	return false
}
func stripNeg(query string) string {
	if strings.HasPrefix(query, "!") {
		return strings.Replace(query, "!", "", -1)
	}
	return query
}
func stripSpace(query string) string {
	return strings.Replace(query, " ", "", -1)

}

func removeDuplicatesUnordered(elements []string) []string {
	encountered := map[string]bool{}

	// Create a map of all unique elements.
	for v := range elements {
		encountered[elements[v]] = true
	}
	// Place all keys from the map into a slice.
	result := []string{}
	for key := range encountered {
		result = append(result, key)
	}
	return result
}

func SymbolTable(kb [][]string) []string {
	var symbolsDuplicate []string
	for i := 0; i < len(kb); i++ {
		for j := 0; j < len(kb[i]); j++ {
			symbolsDuplicate = append(symbolsDuplicate, stripNeg(kb[i][j]))
		}
	}
	symbols := removeDuplicatesUnordered(symbolsDuplicate)
	return symbols
}

func SymbolIsInPremise(premise []string, p string) bool {
	for v := range premise {
		if strings.Contains(premise[v], p) {
			return true
		}

	}
	return false
}
func Pop(slice []string, s int) []string {
	return append(slice[:s], slice[s+1:]...)
}
func RemoveElementFromSlice(slice []string, index int) []string {
	slice[index] = slice[len(slice)-1]
	slice[len(slice)-1] = ""
	slice = slice[:len(slice)-1]
	return slice
}

func ElementLocator(slice []string, find string) int {
	for v := range slice {
		if strings.Contains(slice[v], find) {
			return v
		}
	}
	return -1
}

func CopySlice(dst []string, src []string) {
	for v := range src {
		dst[v] = src[v]
	}
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
