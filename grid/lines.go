package grid

// Diese Aufgaben liefern eine Linie aus dem Spielfeld: eine Zeile,
// eine Spalte oder eine Diagonale. Jede Linie ist eine neue Liste;
// wird sie geändert, bleibt das Spielfeld unverändert.

// Row erwartet ein Spielfeld und die Nummer `index` einer Zeile
// und liefert die Zellen dieser Zeile von links nach rechts.
func Row(cells [][]string, index int) []string {
	result := []string{}

	for _, value := range cells[index] {
		result = append(result, value)
	}
	return result
}

// RowV2 erwartet ein Spielfeld und die Nummer `index` einer Zeile
// und liefert die Zellen dieser Zeile von links nach rechts.
func RowV2(cells [][]string, index int) []string {
	result := make([]string, len(cells[index]))

	copy(result, cells[index])

	return result
}

// Col erwartet ein Spielfeld und die Nummer `index` einer Spalte
// und liefert die Zellen dieser Spalte von oben nach unten.
func Col(cells [][]string, index int) []string {
	result := []string{}

	for row := range cells {
		result = append(result, cells[row][index])
	}

	return result
}

// Diagonal erwartet ein Spielfeld und liefert die Diagonale von links oben
// nach rechts unten: die Zellen, deren Zeile und Spalte gleich sind.
// Bei einem Spielfeld, das nicht quadratisch ist, endet sie am Rand.
func Diagonal(cells [][]string) []string {
	result := []string{}

	for i := range Min(Rows(cells), Cols(cells)) {
		k := i
		result = append(result, cells[i][k])
	}

	return result

	// Alternative mit der generalisierten Diagonalfunktion unten:
	// return GeneralDiag(cells, func(i int) int { return i })
}

// AntiDiagonal erwartet ein Spielfeld und liefert die Diagonale von rechts oben
// nach links unten. Bei einem Spielfeld, das nicht quadratisch ist,
// endet sie am Rand.
func AntiDiagonal(cells [][]string) []string {
	result := []string{}

	for i := range Min(Rows(cells), Cols(cells)) {
		k := Cols(cells) - 1 - i
		result = append(result, cells[i][k])
	}

	return result

	// Alternative mit der generalisierten Diagonalfunktion unten:
	// return GeneralDiag(cells, func(i int) int { return Cols(cells) - 1 - i })
}

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func GeneralDiag(cells [][]string, col func(int) int) []string {
	result := []string{}

	for i := range Min(Rows(cells), Cols(cells)) {
		k := col(i)
		result = append(result, cells[i][k])
	}

	return result
}
