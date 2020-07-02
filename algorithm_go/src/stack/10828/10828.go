package main

import (
	"bufio"
	"fmt"
	"os"
	"stack"
	"strconv"
	"strings"
)

func main(){
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)

	var n int
	var cmd string

	s := stack.Stack(100)

	fmt.Fscanln(r, &n)

	for i := 0; i< n; i++ {

		line, _, _ := r.ReadLine()

		parts := strings.Split(string(line), " ")
		cmd = parts[0]

		if cmd == "push"{
			number, _ := strconv.Atoi(parts[1])
			s.Push(number)
		} else if cmd == "pop" {
			fmt.Fprintln(w, s.Pop())
		} else if cmd == "size" {
			fmt.Fprintln(w, s.Size())
		} else if cmd == "top" {
			fmt.Fprintln(w, s.Top())
		} else if cmd == "empty" {
			fmt.Fprintln(w, s.Empty());
		}
	}
	w.Flush()


}
