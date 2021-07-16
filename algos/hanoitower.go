package algos

import "fmt"

func TOHUtil(n int, from string, to string, temp string) {
	/* If only 1 disk, make the move and return */ if n == 1 {
		fmt.Println("Move disk ", n, " from peg ", from, " to peg ", to)
		return
	}
	/* Move top n-1 disks from A to B, using C as auxiliary */
	TOHUtil(n-1, from, temp, to)
	/* Move remaining disks from A to C */
	fmt.Println("Move disk ", n, " from peg ", from, " to peg ", to)
	/* Move n-1 disks from B to C using A as auxiliary */
	TOHUtil(n-1, temp, to, from)
}

func TowersOfHanoi(n int) {
	TOHUtil(n, "A", "C", "B")
}
