package slice

import "fmt"

func Example_slicing() {
	nums := []int{1, 2, 3, 4, 5}
	fmt.Println(nums)
	fmt.Println(nums[1:3])
	fmt.Println(nums[2:])
	fmt.Println(nums[:3])

	// Output:
	//[1 2 3 4 5]
	//[2 3]
	//[3 4 5]
	//[1 2 3]
}

func Example_append() {
	f1 := []string{"사과", "바나나", "딸기"}
	f2 := []string{"포도", "딸기"}
	f3 := append(f1, f2...) //이어붙이기
	f4 := append(f1[:2], f2...)

	fmt.Println(f1)
	fmt.Println(f2)
	fmt.Println(f3)
	fmt.Println(f4)

	//Output:
	//[사과 바나나 딸기]
	//[포도 딸기]
	//[사과 바나나 딸기 포도 딸기]
	//[사과 바나나 포도 딸기]
}
