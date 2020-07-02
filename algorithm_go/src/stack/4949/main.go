package main

import (
	"bufio"
	"fmt"
	"os"
	"stack"
)

func main() {
	w := bufio.NewWriter(os.Stdout)
	r := bufio.NewReader(os.Stdin)

	for true {
		line, _, _ := r.ReadLine()

		realLine := string(line)

		if realLine == "."{
			break
		}

		tagMap := map[int32]int32{
			']': '[',
			')': '(',
		}

		s := stack.Stack(10000)

		for _, char := range realLine {
			if char == '(' || char == '[' {
				s.Push(int(char))
			} else if char == ')' ||  char == ']'{
				closingTag := tagMap[char]

				if s.Top() == int(closingTag) {
					s.Pop()
				} else {
					s.Push(int(char))
					break
				}
			}
		}

		if s.Empty() == 1 {
			fmt.Fprintln(w, "yes")
		} else {
			fmt.Fprintln(w, "no")
		}
	}

	w.Flush()
}
