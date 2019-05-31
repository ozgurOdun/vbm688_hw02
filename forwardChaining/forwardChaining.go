package forwardChaining

import (
	"fmt"
	"github.com/ozgurOdun/vbm688_hw02/utils"
)

func Entails(kb [][]string, goal string) bool {
	var agenda []string
	for i := 0; i < len(kb); i++ {
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
			return true
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
	}
	return false
}
