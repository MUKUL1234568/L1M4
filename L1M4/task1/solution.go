package main

import (
	"fmt"
	"math"
)


func minSteps(m, n, d int) int {

	if d > int(math.Max(float64(m), float64(n))) {
		return -1
	}

	 
	queue := [][]int{{0, 0, 0}} // (jug1, jug2, steps)
	visited := make([][]bool, m+1)
	for i := range visited {
		visited[i] = make([]bool, n+1)
	}
	visited[0][0] = true

	 
	for len(queue) > 0 {
	 
		curr := queue[0]
		queue = queue[1:]

		jug1, jug2, steps := curr[0], curr[1], curr[2]
 
		if jug1 == d  {
			return steps
		}

		 

		// 1: Fill jug1
		if !visited[m][jug2] {
			visited[m][jug2] = true
			queue = append(queue, []int{m, jug2, steps + 1})
		}

		// 2: Fill jug2
		if !visited[jug1][n] {
			visited[jug1][n] = true
			queue = append(queue, []int{jug1, n, steps + 1})
		}

		// 3: Empty jug1
		if !visited[0][jug2] {
			visited[0][jug2] = true
			queue = append(queue, []int{0, jug2, steps + 1})
		}

		// 4: Empty jug2
		if !visited[jug1][0] {
			visited[jug1][0] = true
			queue = append(queue, []int{jug1, 0, steps + 1})
		}

		// 5: Pour jug1 into jug2
		pour1to2 := min(jug1, n-jug2)
		if !visited[jug1-pour1to2][jug2+pour1to2] {
			visited[jug1-pour1to2][jug2+pour1to2] = true
			queue = append(queue, []int{jug1 - pour1to2, jug2 + pour1to2, steps + 1})
		}

		// 6: Pour jug2 into jug1
		pour2to1 := min(jug2, m-jug1)
		if !visited[jug1+pour2to1][jug2-pour2to1] {
			visited[jug1+pour2to1][jug2-pour2to1] = true
			queue = append(queue, []int{jug1 + pour2to1, jug2 - pour2to1, steps + 1})
		}
	}

	
	return -1
}

 
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
 
	m, n, d := 4, 3, 2
	fmt.Println(minSteps(m, n, d))
}
