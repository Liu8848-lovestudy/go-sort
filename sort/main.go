package main

import "fmt"

func main() {
	arr := [...]int{10, 9, 11, 23, 13, 67, 99, 59}
	for i := 0; i < len(arr)-1; i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] > arr[j] {
				temp := arr[i]
				arr[i] = arr[j]
				arr[j] = temp
			}
		}
	}
	fmt.Println(arr)
}
