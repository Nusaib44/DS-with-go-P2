package main

import "fmt"

func Input(arrSlice []int, err error) []int {
	var value int

	if err != nil {
		return arrSlice
	}
	number, err := fmt.Scanf("%d", &value)

	if number == 1 {
		arrSlice = append(arrSlice, value)
	}
	return Input(arrSlice, err)
}

func BubbleSort(bubbleSlice []int) []int {
	for i := 0; i < len(bubbleSlice)-1; i++ {
		for j := 0; j < len(bubbleSlice)-i-1; j++ {
			if bubbleSlice[j] > bubbleSlice[j+1] {
				bubbleSlice[j], bubbleSlice[j+1] = bubbleSlice[j+1], bubbleSlice[j]
			}
		}
	}
	return bubbleSlice
}

func main() {
	fmt.Println("Enter the values of slice")
	mainSlice := Input([]int{}, nil)
	fmt.Println("Values on slice is", mainSlice)
	x := BubbleSort(mainSlice)
	fmt.Println("Sorted list", x)

	fmt.Println("Enter the values of slice")
	second := Input(mainSlice, nil)
	fmt.Println("Current slice is : ", second)
	y := BubbleSort(second)
	fmt.Println("Sorted slice is : ", y)
}
