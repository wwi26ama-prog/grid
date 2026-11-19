package grid

import "fmt"

func ExampleRow() {
	cells := [][]string{{"a", "b", "c"}, {"d", "e", "f"}, {"g", "h", "i"}}
	fmt.Println(Row(cells, 0))
	fmt.Println(Row(cells, 2))

	line := Row(cells, 1)
	for i := range line {
		line[i] = "z"
	}
	fmt.Println(cells[1])

	// Output:
	// [a b c]
	// [g h i]
	// [d e f]
}

func ExampleCol() {
	cells := [][]string{{"a", "b", "c"}, {"d", "e", "f"}, {"g", "h", "i"}}
	fmt.Println(Col(cells, 0))
	fmt.Println(Col(cells, 2))

	// Output:
	// [a d g]
	// [c f i]
}

func ExampleDiagonal() {
	fmt.Println(Diagonal([][]string{{"a", "b", "c"}, {"d", "e", "f"}, {"g", "h", "i"}}))
	fmt.Println(Diagonal([][]string{{"a", "b", "c"}, {"d", "e", "f"}}))
	fmt.Println(Diagonal([][]string{{"a", "b"}, {"c", "d"}, {"e", "f"}}))

	// Output:
	// [a e i]
	// [a e]
	// [a d]
}

func ExampleAntiDiagonal() {
	fmt.Println(AntiDiagonal([][]string{{"a", "b", "c"}, {"d", "e", "f"}, {"g", "h", "i"}}))
	fmt.Println(AntiDiagonal([][]string{{"a", "b", "c"}, {"d", "e", "f"}}))

	// Output:
	// [c e g]
	// [c e]
}
