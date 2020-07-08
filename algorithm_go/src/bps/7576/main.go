package main

import (
	"bufio"
	"fmt"
	"os"
	"queue"
)


var board = [1005][1005]int{}
var vis = [1005][1005]int{}

type Pair struct {
	x int
	y int
}

var dx = [4]int{1, 0, -1, 0}
var dy = [4]int{0, 1, 0, -1}

func main() {
	var n, m, maxDate int

	r := bufio.NewReader(os.Stdin)

	fmt.Scanln(&m, &n)

	q := queue.Queue()

	for i:=0; i<n; i++{
		for j:=0; j<m; j++{
			fmt.Fscan(r, &board[i][j])

			if board[i][j] == 1 {
				q.Push(Pair{i, j})
			}

			if board[i][j] == 0 {
				vis[i][j] = -1
			}
		}
	}

	for ;q.Empty() != 1;{
		curPair := q.Pop().(Pair)
		curDay := vis[curPair.x][curPair.y]

		for dir:=0; dir<4; dir++{
			nx := curPair.x + dx[dir]
			ny := curPair.y + dy[dir]

			if nx < 0 || nx >= n || ny < 0 || ny >= m {
				continue
			}
			if vis[nx][ny] >= 0 {
				continue
			}

			q.Push(Pair{nx, ny})
			vis[nx][ny] = curDay + 1
		}
	}

	for i:=0; i<n; i++ {
		for j := 0; j < m; j++ {
			if vis[i][j] == -1 {
				fmt.Println(-1)
				return
			}

			if maxDate < vis[i][j]{
				maxDate = vis[i][j]
			}
		}
	}

	fmt.Println(maxDate)


}
