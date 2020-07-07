package main

import (
	"bufio"
	"fmt"
	"os"
	"queue"
)

type Pair struct {
	x int
	y int
}

var board [105]string
var vis [105][105]int

var dx = [4]int{1, 0, -1, 0}
var dy = [4]int{0, 1, 0, -1}


func main() {
	/*
	given
4 6
101111
101010
101011
111011
	 */
	/*
	then
15
	 */

	r := bufio.NewReader(os.Stdin)

	var n, m int

	fmt.Fscanln(r, &n, &m)

	for i := 0; i<n; i++{
		fmt.Fscanln(r, &board[i])
	}

	q := queue.Queue()
	vis[0][0] = 1

	q.Push(Pair{0, 0})
	for ;q.Empty() != 1; {
		curCursor := q.Pop().(Pair)
		currentPoint := vis[curCursor.x][curCursor.y]

		for dir:=0; dir<4; dir++{
			nx := curCursor.x + dx[dir]
			ny := curCursor.y + dy[dir]

			if nx < 0 || nx >= n || ny < 0 || ny >= m{
				continue
			}

			if board[nx][ny] != '1' || vis[nx][ny] != 0{
				continue
			}
			vis[nx][ny] = currentPoint + 1
			q.Push(Pair{nx, ny})
		}
	}

	fmt.Println(vis[n-1][m-1])
}
