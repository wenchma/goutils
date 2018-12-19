package algorithm

import (
	"fmt"
)

/* 有一个M*N的矩阵, 从左上角出发, 只能往右或下走，直到最右下角,
 求最佳路径，使得路径上的数字加起来和最大.
 动态规划的关键点在于定义一个状态：1)问题的答案是某种状态，或可由状态推导
 2)当前状态可由之前状态推导而来

 Calc the best path(from first to end) for M*N Matrix
 */

const RIGHT, DOWN = 1, 2

func calcMax(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func calcBestPath (matrix [][]int) (bestV int, bestP []int) {
	m := len(matrix)
	n := len(matrix[0])
	var best = make([][]int, m)
	for i :=0; i < m; i++ {
		best[i] = make([]int, n)
	}
	var bestPath = make([]int, m+n-2)

	// Calc the best value
	for i:= 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 && j == 0 {
				best[i][j] = matrix[i][j]
			} else if i == 0 {
				best[i][j] = best[i][j-1] + matrix[i][j]
			} else if j ==0 {
				best[i][j] = best[i-1][j] + matrix[i][j]
			} else {
				best[i][j] = calcMax(best[i][j-1], best[i-1][j]) + matrix[i][j]
			}
		}
	}

	//Calc the best path
	x, y := m-1, n-1
	for x >= 0 && y >= 0 && x + y > 0 {
		if x == 0 {
			bestPath[x+y-1] = RIGHT
			y--
		} else if y == 0 {
			bestPath[x+y-1] = DOWN
			x--
		} else {
			if best[x-1][y] > best[x][y-1] {
				bestPath[x+y-1] = DOWN
				x--
			} else {
				bestPath[x+y-1] = RIGHT
				y--
			}
		}
	}
	return best[m-1][n-1], bestPath
}

func printPath(path []int) {
	for i :=0; i < len(path); i++ {
		if path[i] == RIGHT {
			fmt.Print("RIGHT ")
		} else {
			fmt.Print("DOWN ")
		}
	}
	fmt.Print("\n")
}

func main() {
	var matrix1 = [][]int{
		{300, 500, 560, 400, 160},
		{1000, 100, 200, 340, 690},
		{600, 500, 500, 460, 320},
		{300, 400, 250, 210, 760},
	}
	var matrix2 = [][]int{
		{300, 500, 2560, 400},
		{1000, 100, 200, 340},
		{600, 500, 500, 460},
		{300, 400, 250, 210},
		{860, 690, 320, 760},
	}

	v1, p1 := calcBestPath(matrix1)
	fmt.Println(v1)
	printPath(p1)

	v2, p2 := calcBestPath(matrix2)
	fmt.Println(v2)
	printPath(p2)
}
