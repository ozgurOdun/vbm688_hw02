package main

import (
	"fmt"
	"github.com/ozgurOdun/vbm688_hw02/utils"
	"os"
	"strings"
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

	kbRaw := utils.InputParser("kb.txt")
	if kbRaw == nil {
		return
	}

	kb := make([][]string, len(kbRaw))
	for i := 0; i < len(kbRaw); i++ {
		fmt.Println(i, kbRaw[i])
		tmp := strings.Split(kbRaw[i], " ")
		kb[i] = make([]string, len(tmp))
		for j := 0; j < len(tmp); j++ {
			kb[i][j] = tmp[j]
		}

	}
	/*kb[0] = []string{"A"}
	kb[1] = []string{"B"}
	kb[2] = []string{"!A", "R"}
	kb[3] = []string{"!B", "!R", "S"}*/
	fmt.Println("kb")
	fmt.Println(kb[0])
	fmt.Println(kb[1])
	fmt.Println(kb[2])
	fmt.Println(kb[3])

	fmt.Println(strings.Compare(kb[3][2], "S"))
	var agenda []string
	for i := 0; i < len(kbRaw); i++ {
		if len(kb[i]) == 1 {
			fmt.Println(kb[i][0] + "is fact")
			agenda = append(agenda, kb[i][0])
		}
	}
	count := map[int]int{}
	var g int
	for range kb {
		count[g] = len(kb[g])
		g++
	}
	symbols := utils.SymbolTable(kb)
	infered := map[string]bool{}
	for v := range symbols {
		infered[symbols[v]] = false
	}

	for ok := true; ok; ok = len(agenda) != 0 {
		p := agenda[0]
		agenda = utils.RemoveElementFromSlice(agenda, 0)
		if utils.IsGoal(goal, p) {
			fmt.Println("successsssssssssssssssssssssss", goal, p)
			return
		}
		if infered[p] == false {
			infered[p] = true
			for a := range kb {
				if utils.SymbolIsInPremise(kb[a], p) {
					count[a]--
					loc := utils.ElementLocator(kb[a], p)
					tmp := make([]string, len(kb[a]))
					utils.CopySlice(tmp, kb[a])
					kb[a] = utils.RemoveElementFromSlice(kb[a], loc)
					fmt.Println("steps", p, tmp, kb[a])
					if count[a] == 1 {
						agenda = append(agenda, kb[a][0])
					}
				}

			}

		}
		fmt.Println("falseeeeeeeeeee")
		//return
	}

	fmt.Println("End...", time.Now())
	elapsed := time.Since(startTime)
	fmt.Println("Process took:", elapsed)
	utils.PrintMemUsage()
	return
}
