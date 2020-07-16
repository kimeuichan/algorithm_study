package main

import (
	"fmt"
	"queue"
)

var vis = [100005]int{}


/*
given
5 17

then
4
 */
func main() {
	var n, m int

	fmt.Scanln(&n, &m)

	q := queue.Queue()
	q.Push(n)

	for i, _ := range vis{
		vis[i] = -1
	}

	vis[n] = 0

	for ; vis[m] == -1; {
		cur := q.Pop().(int)
		for _, v := range []int{cur-1, cur + 1, cur * 2}{

			if v < 0 || v > 100000 {
				continue
			}

			if vis[v] != -1 {
				continue
			}
			vis[v] = vis[cur] + 1
			q.Push(v)
		}
	}

	fmt.Println(vis[m])

}
