package main

import "fmt"

func main() {
	arr := [...]int{10, 9, 11, 23, 13, 67, 99, 59}
	for i := 0; i < len(arr)-1; i++ { //循环趟数
		for j := 0; j < len(arr)-1-i; j++ { //相邻元素两两相比较，冒泡出大数并放到后面
			if arr[j] > arr[j+1] {
				temp := arr[j]
				arr[j] = arr[j+1]
				arr[j+1] = temp
			}
		}
	}
	fmt.Println(arr)
}
