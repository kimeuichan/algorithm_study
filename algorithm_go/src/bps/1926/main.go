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

var board[505][505]int
var vis[505][505]bool

var dx = [4]int{1, 0, -1, 0}
var dy = [4]int{0, 1, 0, -1}

func main() {
	w := bufio.NewWriter(os.Stdout)
	r := bufio.NewReader(os.Stdin)

	var n, m, currentPictureWidth, maxPictureWidth, pictureCount int

	q := queue.Queue()

	fmt.Fscan(r, &n, &m)

	for i := 0; i<n; i++{
		for j:= 0; j<m; j++{
			fmt.Fscan(r, &board[i][j])
		}
	}


	for i := 0; i<n; i++{
		for j:= 0; j<m; j++{
			if vis[i][j] || board[i][j] != 1{
				continue
			}

			pictureCount++
			q.Push(Pair{i, j})

			vis[i][j] = true

			currentPictureWidth = 0
			for ; q.Empty() != 1; {
				curPair := q.Pop().(Pair)
				currentPictureWidth++

				for dir := 0; dir< 4; dir++{
					nx := curPair.x + dx[dir]
					ny := curPair.y + dy[dir]

					if nx < 0 || nx >= n || ny < 0 || ny >= m{ // i가 x이기 때문에 m, j가 y이기 때문에 n => x, y가 반대라고 생각하면 편하다.
						continue
					}

					if vis[nx][ny] || board[nx][ny] != 1{
						continue
					}

					vis[nx][ny] = true
					q.Push(Pair{nx, ny})
				}
			}

			if maxPictureWidth < currentPictureWidth{
				maxPictureWidth = currentPictureWidth
			}
		}
	}

	fmt.Fprintln(w, pictureCount)
	fmt.Fprintln(w, maxPictureWidth)
	w.Flush()
}
