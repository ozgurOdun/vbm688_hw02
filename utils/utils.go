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

func InputParser(fileName string) [][]string {
	file, err := os.Open(fileName)
	checkErr(err)
	rawInput := make([]byte, 100)
	//var rawInput []byte
	n, err := file.Read(rawInput)
	if n > 0 {
		//fmt.Println("Raw input is: ", string(rawInput), n)
	}
	inputLines := strings.Split(string(rawInput[:n]), "\r\n")
	//fmt.Println(len(inputLines))

	kb := make([][]string, len(inputLines))
	for i := 0; i < len(inputLines)-1; i++ {
		tmp := strings.Split(inputLines[i], " ")
		kb[i] = make([]string, len(tmp))
		for j := 0; j < len(tmp); j++ {
			kb[i][j] = tmp[j]
		}

	}
	//fmt.Println(kb)
	return kb
}

func ProofBuilder(clause1 string, clause2, resolvent []string) (bool, []string) {
	var clause2Str, resolventStr string
	for v := range clause2 {
		if v == 0 {
			clause2Str = clause2[v]
		} else {
			clause2Str = clause2Str + " " + clause2[v]
		}
	}
	for v := range resolvent {
		if v == 0 {
			resolventStr = resolvent[v]
		} else {
			resolventStr = resolventStr + " " + resolvent[v]
		}
	}
	if strings.EqualFold(resolventStr, "") {
		return false, nil
	}

	step := []string{clause1, clause2Str, resolventStr}
	return true, step
}

func ProofPrinter(proof [][]string) {
	fmt.Printf("\n\n|Clause1   |Clause2   |Resolvent |\n")
	for v := range proof {
		if len(proof[v]) == 3 {
			fmt.Printf("|%-10s|%-10s|%-10s|\n", proof[v][0], proof[v][1], proof[v][2])
		}
	}
	fmt.Println("\n")
}

func IsGoal(goal, query string) bool {
	if strings.EqualFold(goal, query) {
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

func RemoveElementFromSlice(slice []string, index int) []string {
	//slice[index] = slice[len(slice)-1]
	//slice[len(slice)-1] = ""
	//slice = slice[:len(slice)-1]
	slice = slice[1:]
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
