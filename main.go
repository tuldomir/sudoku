package main

import (
	"fmt"
)

func main() {
	fmt.Println("sudoku solver started")
	arr := [][]int{
		// {0, 5, 0, 1, 0, 0, 4, 7, 9}, // 2-5-6-
		// {7, 0, 3, 5, 9, 0, 1, 6, 2},
		// {0, 9, 1, 7, 6, 0, 5, 3, 8},
		// {9, 0, 0, 8, 0, 1, 3, 4, 0},
		// {0, 1, 0, 3, 0, 9, 7, 2, 0},
		// {0, 3, 0, 2, 0, 6, 8, 9, 1},
		// {3, 0, 0, 9, 1, 5, 6, 8, 7},
		// {5, 0, 0, 4, 2, 7, 9, 1, 3},
		// {1, 7, 9, 6, 0, 0, 2, 5, 4},
		{0, 0, 2, 6, 0, 4, 7, 0, 0}, // 2-5-6-
		{4, 6, 0, 0, 7, 0, 0, 9, 3},
		{5, 0, 0, 9, 0, 3, 0, 0, 6},
		{2, 0, 6, 0, 0, 0, 4, 0, 9},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{8, 0, 3, 0, 0, 0, 5, 0, 7},
		{6, 0, 0, 8, 0, 1, 0, 0, 5},
		{3, 2, 0, 0, 4, 0, 0, 6, 1},
		{0, 0, 9, 2, 0, 6, 3, 0, 0},
	}

	res, ok := solver(arr)
	if !ok {
		fmt.Println("cant resolve sudoku")
		return
	}
	fmt.Println("input")
	print(arr)
	fmt.Println("result")
	print(res)
}

type pos struct {
	x, y, sec, v int
}

func solver(arr [][]int) ([][]int, bool) {

	newArr := make([][]int, len(arr))
	for i, v := range arr {
		newArr[i] = make([]int, len(v))
		for j, k := range v {
			newArr[i][j] = k
		}
	}

	vert := make([]map[int]bool, 9)
	horiz := make([]map[int]bool, 9)
	// 3*row + col
	sect := make([]map[int]bool, 9)
	empty := make([]*pos, 0)

	// make maps
	for i := 0; i < 9; i++ {
		vert[i] = make(map[int]bool)
		horiz[i] = make(map[int]bool)
		sect[i] = make(map[int]bool)
	}

	// fill digits  in maps for vert, horiz, sect
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr); j++ {
			colPos := i / 3
			rowPos := j / 3
			// sector position
			sectNum := 3*colPos + rowPos

			if arr[i][j] == 0 {
				empty = append(empty, &pos{
					x: i, y: j, sec: sectNum, v: 0,
				})
				continue
			}

			h := arr[i][j]
			v := arr[j][i]

			vert[i][v] = true
			horiz[i][h] = true
			sect[sectNum][h] = true
		}
	}
	cur := 0

	//  run
	for cur < len(empty) {
		if cur < 0 {
			return [][]int{}, false
		}
		p := empty[cur]

		var found bool
		// loop for inserting number in sudoku
		for k := p.v + 1; k < 10; k++ {

			if vert[p.y][k] || horiz[p.x][k] || sect[p.sec][k] {
				continue
			}

			newArr[p.x][p.y] = k
			vert[p.y][k] = true
			horiz[p.x][k] = true
			sect[p.sec][k] = true
			p.v = k
			found = true
			break
		}

		if found {
			cur++
			// fmt.Println(newArr)
			// time.Sleep(1 * time.Second)
			continue
		}

		cur--
		prev := empty[cur]
		vert[prev.y][prev.v] = false
		horiz[prev.x][prev.v] = false
		sect[prev.sec][prev.v] = false
		newArr[prev.x][prev.y] = 0
		p.v = 0

	}

	return newArr, true
}

func print(arr [][]int) {
	for _, v := range arr {
		fmt.Println(v)
	}
}
