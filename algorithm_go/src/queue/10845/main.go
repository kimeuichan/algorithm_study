package main

import (
	"bufio"
	"fmt"
	"os"
	"queue"
	"strconv"
	"strings"
)

func main() {
	que := queue.Queue()

	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)

	var n int

	fmt.Fscanln(r, &n)

	for i :=0; i<n; i++{
		line, _, _ := r.ReadLine()
		cmd := strings.Split(string(line), " ")

		if cmd[0] == "push" {
			number, _ := strconv.Atoi(cmd[1])
			que.Push(number)
		} else if cmd[0] == "pop" {
			fmt.Fprintln(w, que.Pop())
		} else if cmd[0] == "size" {
			fmt.Fprintln(w, que.Size())
		} else if cmd[0] == "empty" {
			fmt.Fprintln(w, que.Empty())
		} else if cmd[0] == "front" {
			fmt.Fprintln(w, que.Front())
		} else if cmd[0] == "back" {
			fmt.Fprintln(w, que.Back())
		}
	}

	w.Flush()

}
