package main

import (
	"fmt"
	"stack"
)

type Pair struct {
	x int
	y int
}

var dx = [4]int{1, 0, -1, 0}
var dy = [4]int{0, 1, 0, -1}

var board = [8][8]int{
	{0, 0, 0},
	{0, 1, 0},
	{0, 0, 0},
	{0, 0, 1},
	{1, 0, 1},
}
var vis = [8][8]int{}

func main() {
	n, m := 6, 3

	s := stack.Stack(100)

	startIndex := Pair{0, 0}
	vis[startIndex.x][startIndex.y] = 0

	s.Push(startIndex)

	for ;s.Empty() != 1; {
		curPair := s.Pop().(Pair)

		for dir:=0; dir<4; dir++ {
			nx := curPair.x + dx[dir]
			ny := curPair.y + dy[dir]

			if nx < 0 || nx >= n || ny < 0 || ny >= m {
				continue
			}

			if vis[nx][ny] != 0 || board[nx][ny] != 0{
				continue
			}

			s.Push(Pair{nx, ny})
			vis[nx][ny] = 1
		}
	}

	for _, val := range vis{
		fmt.Println(val)
	}
}
