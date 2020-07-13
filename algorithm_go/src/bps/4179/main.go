package main

import (
	"bufio"
	"fmt"
	"os"
	"queue"
	"strings"
)

type Pair struct{
	x int
	y int
}

var dx = [4]int{1, 0, -1, 0}
var dy = [4]int{0, 1, 0, -1}

/*
given
4 4
####
#JF#
#..#
#..#

then
3
*/


func main() {
	var board = [1005]string{}
	var vis = [1005][1005]int{}
	var n, m int
	var playerPair Pair

	r := bufio.NewReader(os.Stdin)

	fmt.Fscanln(r, &n, &m)

	q := queue.Queue()
	for i := 0; i<n; i++{
		fmt.Fscanln(r, &board[i])
		fireIndex := strings.Index(board[i], "F")
		playerIndex := strings.Index(board[i], "J")

		if fireIndex != -1 {
			q.Push(Pair{i, fireIndex})
			vis[i][fireIndex] = 1
		}

		if playerIndex != -1 {
			playerPair = Pair{i, playerIndex}
		}
	}

	for ;q.Empty() != 1;{
		curPair := q.Pop().(Pair)
		curFireMove := vis[curPair.x][curPair.y]

		for dir:=0; dir<4; dir++{
			nx := curPair.x + dx[dir]
			ny := curPair.y + dy[dir]

			if nx < 0 || nx >= n || ny < 0 || ny >= m {
				continue
			}

			if vis[nx][ny] > 0 || board[nx][ny] == '#'{
				continue
			}

			q.Push(Pair{nx, ny})
			vis[nx][ny] = curFireMove + 1
		}
	}

	q.Push(playerPair)
	vis[playerPair.x][playerPair.y] = 1

	for ; q.Empty() != 1; {
		curPair := q.Pop().(Pair)
		currentMove := vis[curPair.x][curPair.y]

		for dir:=0; dir<4; dir++{
			nx := curPair.x + dx[dir]
			ny := curPair.y + dy[dir]

			if nx < 0 || nx >= n || ny < 0 || ny >=m {
				fmt.Println(currentMove)
				return
			}

			if board[nx][ny] == '#'{
				continue
			}

			if vis[nx][ny] <= currentMove {
				continue
			}

			vis[nx][ny] = currentMove + 1
			q.Push(Pair{nx, ny})
		}
	}
	fmt.Println("IMPOSSIBLE")
}
