package main

import (
	"fmt"
	"queue"
)

type Pair struct{
	x int
	y int
}

var board[501][501]int
var visit[501][501]bool

var dy = [4]int{1, 0, -1, 0}
var dx = [4]int{0, 1, 0, -1}

const widthMax = 7
const heightMax = 5

func main() {
	q := queue.Queue()

	visit[0][0] = true
	q.Push(Pair{0, 0})

	for ;q.Empty() != 1; {
		curPair := q.Pop().(Pair)

		fmt.Printf("<%d, %d>\n", curPair.x, curPair.y)

		for dir := 0; dir< 4; dir++ {
			nx := curPair.x + dx[dir]
			ny := curPair.y + dy[dir]

			if nx < 0 || nx >= widthMax || ny < 0 || ny >= heightMax {
				continue
			}
			if visit[nx][ny] || board[nx][ny] != 1{
				continue
			}

			visit[nx][ny] = true
			q.Push(Pair{nx, ny})
		}
	}
}
