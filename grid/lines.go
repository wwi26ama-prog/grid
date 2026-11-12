package grid

// Diese Aufgaben liefern eine Linie aus dem Spielfeld: eine Zeile,
// eine Spalte oder eine Diagonale. Jede Linie ist eine neue Liste;
// wird sie geändert, bleibt das Spielfeld unverändert.

// Row erwartet ein Spielfeld und die Nummer `index` einer Zeile
// und liefert die Zellen dieser Zeile von links nach rechts.
func Row(cells [][]string, index int) []string {
	// Hinweis:
	// Kopieren Sie die Zellen der Zeile in eine neue Liste.
	// Würde die Zeile selbst zurückgegeben, könnte der Aufrufer
	// über das Ergebnis das Spielfeld ändern.

	line := []string{}
	for _, value := range cells[index] {
		line = append(line, value)
	}
	return line
}

// Col erwartet ein Spielfeld und die Nummer `index` einer Spalte
// und liefert die Zellen dieser Spalte von oben nach unten.
func Col(cells [][]string, index int) []string {
	// Hinweis:
	// Eine Spalte steht nicht am Stück im Spielfeld:
	// Aus jeder Zeile wird die Zelle an derselben Stelle geholt.

	line := []string{}
	for _, row := range cells {
		line = append(line, row[index])
	}
	return line
}

// Diagonal erwartet ein Spielfeld und liefert die Diagonale von links oben
// nach rechts unten: die Zellen, deren Zeile und Spalte gleich sind.
// Bei einem Spielfeld, das nicht quadratisch ist, endet sie am Rand.
func Diagonal(cells [][]string) []string {
	// Hinweis:
	// Die Diagonale endet, sobald die Zeilen oder die Spalten zu Ende sind.

	line := []string{}
	for i := 0; i < Rows(cells) && i < Cols(cells); i++ {
		line = append(line, cells[i][i])
	}
	return line
}

// AntiDiagonal erwartet ein Spielfeld und liefert die Diagonale von rechts oben
// nach links unten. Bei einem Spielfeld, das nicht quadratisch ist,
// endet sie am Rand.
func AntiDiagonal(cells [][]string) []string {
	// Hinweis:
	// Die Zeile läuft von oben nach unten, die Spalte gleichzeitig
	// von rechts (`Cols(cells) - 1`) nach links.

	line := []string{}
	cols := Cols(cells)
	for i := 0; i < Rows(cells) && i < cols; i++ {
		line = append(line, cells[i][cols-1-i])
	}
	return line
	// Anmerkung:
	// Ende der Schleife:
	// Die Bedingung `i < cols` beendet die Schleife, bevor die Spalte
	// `cols-1-i` negativ würde.
}
