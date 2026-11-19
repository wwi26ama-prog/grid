// Package grid enthält Funktionen für Spielfelder aus Zeilen und Spalten,
// wie sie z. B. Tic Tac Toe oder Vier Gewinnt brauchen.
package grid

// Diese Aufgaben sind die Grundlagen: ein Spielfeld anlegen, seine Größe
// abfragen, einzelne Zellen lesen und schreiben und das Spielfeld ausgeben.
// Ein Spielfeld ist eine Liste von Zeilen, jede Zeile eine Liste von Zellen,
// also `[][]string`. Zeilen und Spalten werden ab `0` gezählt.

// New erwartet die Anzahl der Zeilen `rows` und Spalten `cols` und einen Wert
// `value` und liefert ein neues Spielfeld, in dem jede Zelle `value` enthält.
// Mit `"."` als Wert entsteht z. B. ein leeres Spielfeld.
func New(rows, cols int, value string) [][]string {
	// Ein leeres Grid erzeugen.
	grid := [][]string{}

	// Die Zeilen erzeugen.
	for range rows {
		row := []string{}
		for range cols {
			row = append(row, value)
		}
		grid = append(grid, row)
	}

	return grid
}

// Rows erwartet ein Spielfeld und liefert die Anzahl seiner Zeilen.
func Rows(cells [][]string) int {
	return len(cells)
}

// Cols erwartet ein Spielfeld und liefert die Anzahl seiner Spalten,
// also die Länge der ersten Zeile. Ein Spielfeld ohne Zeilen hat `0` Spalten.
func Cols(cells [][]string) int {
	if Rows(cells) == 0 {
		return 0
	}
	return len(cells[0])
}

// Get erwartet ein Spielfeld und die Zeile `row` und Spalte `col` einer Zelle
// und liefert ihren Inhalt.
func Get(cells [][]string, row, col int) string {
	return cells[row][col]
}

// Set erwartet ein Spielfeld, die Zeile `row` und Spalte `col` einer Zelle
// und einen Wert `value` und schreibt den Wert in die Zelle.
// Die Änderung ist auch beim Aufrufer zu sehen.
func Set(cells [][]string, row, col int, value string) {
	cells[row][col] = value
}

// Format erwartet ein Spielfeld und liefert es als Text zur Ausgabe:
// eine Textzeile pro Zeile des Spielfelds, die Zellen durch Leerzeichen
// getrennt, jede Textzeile mit einem Zeilenumbruch "\n" am Ende.
func Format(cells [][]string) string {
	result := ""

	// Alle Zeilen ausgeben.
	for _, row := range cells {
		result += FormatRow(row)
		result += "\n"
	}

	return result
}

// FormatRow erwartet eine Zeile eines Grids
// und liefert einen String, der die Elemente dieser Zeile
// getrennt durch Leerzeichen enthält.
func FormatRow(row []string) string {
	result := ""

	for _, value := range row {
		result += value
		result += " "
	}

	result = result[:len(result)-1]

	return result
}
