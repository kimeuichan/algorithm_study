package main

import (
	"bufio"
	"fmt"
	"os"
	"stack"
)


func main() {
	r := bufio.NewReader(os.Stdin)

	var pipeCount int
	var prevChar = ' '

	line, _, _ := r.ReadLine()
	pipeString := string(line)
	s := stack.Stack(100000)

	for _, char := range pipeString{
		if char == '('{
			s.Push(int(char))
		} else { // ')' 인 경우
			s.Pop()
			if prevChar == '('{ // 레이저
				pipeCount += s.Size()
			} else {  // 파이프 끝
				pipeCount += 1
			}
		}
		prevChar = char
	}

	fmt.Printf("%d", pipeCount)
}

