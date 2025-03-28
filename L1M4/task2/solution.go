package main

import (
	"fmt"
	"container/list"
)

 
type State struct {
	ML, CL, MR, CR int     
	Boat           string  
	Parent         *State  
}

 
func (s State) isValid() bool {
	if (s.ML < s.CL && s.ML > 0) || (s.MR < s.CR && s.MR > 0) {
		return false
	}
	return s.ML >= 0 && s.CL >= 0 && s.MR >= 0 && s.CR >= 0
}

 
func (s State) isGoal() bool {
	return s.ML == 0 && s.CL == 0 && s.MR == 3 && s.CR == 3
}
 
var possibleMoves = [][]int{
	{1, 0}, {2, 0}, {0, 1}, {0, 2}, {1, 1},  
}


func bfs() {
	start := State{3, 3, 0, 0, "L", nil}
	visited := make(map[State]bool)
	queue := list.New()

	queue.PushBack(start)

	for queue.Len() > 0 {
		curr := queue.Remove(queue.Front()).(State)

		if curr.isGoal() {
			printSolution(&curr)
			return
		}

		for _, move := range possibleMoves {
			newState := getNextState(curr, move)

			if newState.isValid() && !visited[newState] {
				newState.Parent = &curr
				visited[newState] = true
				queue.PushBack(newState)
			}
		}
	}
	fmt.Println("No solution found.")
}

// getNextState generates the next possible state.
func getNextState(s State, move []int) State {
	m, c := move[0], move[1]
	if s.Boat == "L" {
		return State{s.ML - m, s.CL - c, s.MR + m, s.CR + c, "R", nil}
	}
	return State{s.ML + m, s.CL + c, s.MR - m, s.CR - c, "L", nil}
}

 
func printSolution(s *State) {
	var path []*State
	for s != nil {
		path = append([]*State{s}, path...)
		s = s.Parent
	}

	fmt.Println("Solution Path:")
	for _, state := range path {
		fmt.Printf("Left: (%dM, %dC) | Right: (%dM, %dC) | Boat: %s\n",
			state.ML, state.CL, state.MR, state.CR, state.Boat)
	}
}

func main() {
	bfs()
}
