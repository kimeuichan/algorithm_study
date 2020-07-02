package main

import (
	"bufio"
	"deque"
	"fmt"
	"os"
	"strconv"
	"strings"
)


func main(){
	dq := deque.Deque(10000*2)
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)

	var n int

	fmt.Fscanln(r, &n)

	for i := 0; i<n; i++{
		line, _, _ := r.ReadLine()

		cmd := strings.Split(string(line), " ")

		if cmd[0] == "push_front"{
			number, _ := strconv.Atoi(cmd[1])
			dq.PushFront(number)
		} else if cmd[0] == "push_back"{
			number, _ := strconv.Atoi(cmd[1])
			dq.PushBack(number)
		} else if cmd[0] == "pop_front"{
			fmt.Fprintln(w, dq.PopFront())
		} else if cmd[0] == "pop_back"{
			fmt.Fprintln(w, dq.PopBack())
		} else if cmd[0] == "size"{
			fmt.Fprintln(w, dq.Size())
		} else if cmd[0] == "empty"{
			fmt.Fprintln(w, dq.Empty())
		} else if cmd[0] == "front"{
			fmt.Fprintln(w, dq.Front())
		} else if cmd[0] == "back"{
			fmt.Fprintln(w, dq.Back())
		}
	}

	w.Flush()
}
