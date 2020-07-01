package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main(){
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)

	stack := []int{}

	var n int
	var cmd string

	fmt.Fscanln(r, &n)

	for i := 0; i< n; i++ {

		line, _, _ := r.ReadLine()

		parts := strings.Split(string(line), " ")
		cmd = parts[0]

		if cmd == "push"{
			number, _ := strconv.Atoi(parts[1])
			stack = append(stack, number)
		} else if cmd == "pop" {
			if len(stack) == 0 {
				fmt.Fprintln(w, -1)
			} else {
				temp := stack[len(stack)-1]
				stack = stack[:len(stack) - 1]
				fmt.Fprintln(w, temp)
			}

		} else if cmd == "size" {
			fmt.Fprintln(w, len(stack))
		} else if cmd == "top" {
			if len(stack) == 0 {
				fmt.Println(-1)
			} else {
				fmt.Fprintln(w, stack[len(stack) - 1])
			}
		} else if cmd == "empty" {
			if len(stack) == 0{
				fmt.Fprintln(w, 1);
			} else {
				fmt.Fprintln(w, 0);
			}
		}
	}
	w.Flush()


}
