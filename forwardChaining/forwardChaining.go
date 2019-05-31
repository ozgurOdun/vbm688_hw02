package forwardChaining

import (
	"github.com/ozgurOdun/vbm688_hw02/utils"
)

func Entails(kb [][]string, goal string) (bool, [][]string) {
	proof := make([][]string, 100)
	var stepCounter int
	var agenda []string
	for i := 0; i < len(kb); i++ {
		if len(kb[i]) == 1 {
			//fmt.Println(kb[i][0] + "is fact")
			if len(kb[i][0]) > 0 {
				agenda = append(agenda, kb[i][0])
			}
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
		//fmt.Println("agenda", agenda)
		agenda = utils.RemoveElementFromSlice(agenda, 0)
		if utils.IsGoal(goal, p) {
			return true, proof
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

					ok, step := utils.ProofBuilder(p, tmp, kb[a])
					if ok {
						proof[stepCounter] = make([]string, 3)
						//utils.CopySlice(proof[stepCounter], step)
						proof[stepCounter] = step
						//fmt.Printf("step:%s,%s,%s\n", step[0], step[1], step[2])
						stepCounter++
					}

					if count[a] == 1 {
						if len(kb[a][0]) > 0 {
							agenda = append(agenda, kb[a][0])
							//fmt.Println("agenda1", agenda)
						}
					}
				}
			}
		}
	}
	return false, nil
}
