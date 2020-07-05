package main

import (
	"fmt"
	"stack"
)

// 참고
// https://suriisurii.tistory.com/53


func main() {
	// given
	// (()[[]])([])
	// then
	// 28

	var line string
	fmt.Scanln(&line)

	var s, m, answer int
	mul := 1

	st := stack.Stack(1000)
	for i, char := range line {
		if char == '['{
			m++
			st.Push(int(char))
			mul *= 3
			if i+1 < len(line) && line[i+1] == ']'{
				answer += mul
			}
		} else if char == '('{
			s++
			st.Push(int(char))
			mul *= 2
			if i+1 < len(line) && line[i+1] == ')'{
				answer += mul
			}
		} else if char == ']'{
			m--
			if st.Top() != '['{
				break
			}
			st.Pop()
			mul /= 3
		} else if char == ')'{
			s--

			if st.Top() != '('{
				break
			}
			st.Pop()
			mul /= 2
		}
	}

	if st.Empty() != 1 || s != 0 || m != 0 {
		fmt.Println(0)
	} else {
		fmt.Println(answer)
	}
}
