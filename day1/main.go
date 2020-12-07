package main

import "fmt"

func main() {
	list := []int{1721, 979, 366, 299, 675, 1456}

	var result int

	for i := 0; i < len(list); i++ {
		for j := 0; j < len(list)-1; j++ {
			temp := list[i] + list[j+1]

			if temp == 2020 {
				result = list[i] * list[j+1]

				break
			}
		}
	}

	fmt.Println("hey", result)
}
