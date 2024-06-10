package main

import (
	"fmt"
)

func main() {
	// создать пустой срез
	aSlice := []float64{}
	// и длина, и емкость равны 0, поскольку aSlice пуст
	fmt.Println(aSlice, len(aSlice), cap(aSlice))

	aSlice = append(aSlice, 1234.23)
	aSlice = append(aSlice, -439)
	fmt.Println(aSlice, "with length", len(aSlice))

	t := make([]int, 4)
	t[0] = -1
	t[1] = -2
	t[2] = -3
	t[3] = -4
	t = append(t, -5)
	fmt.Println(t)

	// двумерный срез
	// у вас может быть столько измерений, сколько необходимо
	twoD := [][]int{{1, 2, 3}, {4, 5, 6}}
	// просмотр всех элементов 2D-среза
	// с помощью двойного цикла for
	for i, v := range twoD {
		for j, k := range v {
			fmt.Print(i, " ", j, " ", k, "| ")
		}
		fmt.Println()
	}

	make2D := make([][]int, 3)
	fmt.Println(make2D)
	make2D[0] = []int{1, 2, 3, 4}
	make2D[1] = []int{-1, -2, -3, -4}
	make2D[2] = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(make2D)
	fmt.Println(make2D[0][1])
}