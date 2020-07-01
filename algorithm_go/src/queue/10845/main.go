package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	//var que queue.Queue

	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)


	var n int

	fmt.Fscanln(r, &n)
	fmt.Fprintln(w, -1)

}
