package main

import "fmt"

func InsertSlice(argSlice []int, err error) []int {
	var value int
	if err != nil {
		return argSlice
	}
	number, err := fmt.Scanf("%d", &value)
	if number == 1 {
		argSlice = append(argSlice, value)
	}
	return InsertSlice(argSlice, err)
}

func InsertionSort(argSlice []int) []int {
	var x = len(argSlice)
	
	for n := 1; n < x; n++ {
		v := n
		for v > 0 {
			if argSlice[v-1] > argSlice[v] {
				argSlice[v-1], argSlice[v] = argSlice[v], argSlice[v-1]
			}
			v = v - 1
		}
	}
	return argSlice
}

func main() {

	fmt.Println("Enter the values of slice")
	mainSlice := InsertSlice([]int{}, nil)
	fmt.Println("Current slice is : ", mainSlice)
	x := InsertionSort(mainSlice)
	fmt.Println("Sorted slice is : ", x)

	fmt.Println("Enter the values of slice")
	second := InsertSlice(mainSlice, nil)
	fmt.Println("Current slice is : ", second)
	y := InsertionSort(second)
	fmt.Println("Sorted slice is : ", y)
}
