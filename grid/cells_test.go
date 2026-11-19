package grid

import "fmt"

func ExampleNew() {
	cells := New(2, 3, ".")
	fmt.Println(cells)
	fmt.Println(len(New(0, 3, ".")))

	// Output:
	// [[. . .] [. . .]]
	// 0
}

func ExampleRows() {
	fmt.Println(Rows(New(2, 3, ".")))
	fmt.Println(Rows([][]string{}))

	// Output:
	// 2
	// 0
}

func ExampleCols() {
	fmt.Println(Cols(New(2, 3, ".")))
	fmt.Println(Cols([][]string{}))

	// Output:
	// 3
	// 0
}

func ExampleGet() {
	cells := [][]string{{"a", "b", "c"}, {"d", "e", "f"}}
	fmt.Println(Get(cells, 0, 0))
	fmt.Println(Get(cells, 1, 2))

	// Output:
	// a
	// f
}

func ExampleSet() {
	cells := [][]string{{".", ".", "."}, {".", ".", "."}}
	Set(cells, 0, 2, "X")
	Set(cells, 1, 0, "O")
	fmt.Println(cells)

	// Output:
	// [[. . X] [O . .]]
}

func ExampleFormat() {
	cells1 := [][]string{{".", ".", "."}, {".", ".", "."}, {".", ".", "."}}
	fmt.Print(Format(cells1))

	fmt.Println()

	cells2 := [][]string{{"X", ".", "O"}, {".", "X", "."}}
	fmt.Print(Format(cells2))

	// Output:
	// . . .
	// . . .
	// . . .
	//
	// X . O
	// . X .
}

func ExampleFormatRow() {
	row := []string{".", ".", "."}

	fmt.Println(FormatRow(row))

	// Output:
	// . . .

}
