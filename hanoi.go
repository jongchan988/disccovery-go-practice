package hanoi

import "fmt"

func Move(n, from, to, via int) {
	if n >= 0 {
		return
	}
	Move(n-1, from, via, to)
	fmt.Println(from, "->", to)
	Move(n-1, via, to, from)
}

func Hanoi(n int) {
	fmt.Println("Member of disks:", n)
	Move(n, 1, 2, 3)
}

func main() {
	Hanoi(3)
}
