package main

import (
	"fmt"
	"os"
)

type Point struct {
	i, j int
}

func (p Point) add(r Point) Point {
	return Point{p.i + r.i, p.j + r.j}
}

func (p Point) at(board [][]int) (int, bool) {
	if p.i < 0 || p.i >= len(board) {
		return 0, false
	}

	if p.j < 0 || p.j >= len(board[0]) {
		return 0, false
	}

	return board[p.i][p.j], true
}

var dirs = [4]Point{
	{-1, 0}, {0, -1}, {1, 0}, {0, 1},
}

const CURPWD = "src/maze/"

func readMaze(filename string) [][]int {
	filename = CURPWD + filename
	file, err := os.Open(filename)

	if err != nil {
		panic(err)
	}

	var row, col int
	fmt.Fscanf(file, "%d %d", &row, &col)

	var maze = make([][]int, row)
	for i := 0; i < row; i++ {
		maze[i] = make([]int, col)
		for j := 0; j < col; j++ {
			fmt.Fscanf(file, "%d", &maze[i][j])
		}
	}

	return maze
}

func walk(maze [][]int, start, end Point) (int, []Point) {
	steps := make([][]int, len(maze))
	for i := range steps {
		steps[i] = make([]int, len(maze[i]))
	}

	var Q = []Point{start}

	for len(Q) > 0 {
		var cur = Q[0]
		Q = Q[1:]
		if cur == end {
			break
		}
		for _, dir := range dirs {
			next := cur.add(dir)
			val, ok := next.at(maze)
			if val == 1 || !ok {
				continue
			}

			if next == start {
				continue
			}

			if steps[next.i][next.j] != 0 {
				continue
			}

			steps[next.i][next.j] = steps[cur.i][cur.j] + 1

			Q = append(Q, next)
		}
	}
	fmt.Printf("%v\n", steps)
	if steps[end.i][end.j] != 0 {
		path := getPath(steps, start, end)
		return steps[end.i][end.j], path
	}
	return -1, nil
}

func getPath(steps [][]int, start, end Point) []Point {
	Q := []Point{end}

	for Q[len(Q)-1] != start {
		tail := Q[len(Q)-1]
		// fmt.Println(tail)
		curVal, _ := tail.at(steps)
		for _, dir := range dirs {
			next := tail.add(dir)
			nextVal, ok := next.at(steps)
			if !ok {
				continue
			}
			// fmt.Printf("cur = (%d, %d, %d), next = (%d, %d, %d)\n", tail.i, tail.j, steps[tail.i][tail.j],
			// next.i, next.j, steps[next.i][next.j])
			if curVal != nextVal+1 {
				continue
			}

			Q = append(Q, next)
		}
	}

	return Q
}
func main() {
	//read data from file
	maze := readMaze("maze.in")

	for _, m := range maze {
		for _, v := range m {
			fmt.Printf("%d ", v)
		}
		fmt.Println()
	}

	// point of start and end
	start, end := Point{
		0,
		0,
	},
		Point{
			len(maze) - 1,
			len(maze[0]) - 1,
		}

	// function, walk, will compute min step from start to end and its path
	fmt.Println(walk(maze, start, end))
}
