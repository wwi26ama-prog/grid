package extensions

// Diese Aufgaben ergänzen Funktionen für Spielfelder. Verwenden Sie die
// Funktionen der Bibliothek `grid`, statt sie nachzubauen.

import "grid/grid"

// CountValue erwartet ein Spielfeld und einen Wert und liefert,
// wie viele Zellen diesen Wert enthalten.
func CountValue(cells [][]string, value string) int {
	// Hinweis:
	// Wie bei `grid.Contains` brauchen Sie zwei verschachtelte Schleifen,
	// zählen aber alle Treffer, statt beim ersten aufzuhören.

	count := 0
	for _, row := range cells {
		for _, cell := range row {
			if cell == value {
				count++
			}
		}
	}
	return count
}

// Transpose erwartet ein Spielfeld und liefert ein neues Spielfeld,
// in dem die Spalten zu Zeilen geworden sind: Die erste Spalte von `cells`
// ist die erste Zeile des Ergebnisses und so weiter. Aus zwei Zeilen
// mit drei Spalten werden so drei Zeilen mit zwei Spalten.
func Transpose(cells [][]string) [][]string {
	// Hinweis:
	// Die Bibliothek liefert jede Spalte als Liste.
	// Hängen Sie die Spalten nacheinander als Zeilen an.

	result := [][]string{}
	for c := 0; c < grid.Cols(cells); c++ {
		result = append(result, grid.Col(cells, c))
	}
	return result
}

// ContainsRun erwartet eine Linie, einen Wert und eine Länge `length` ab `1`
// und liefert `true`, wenn die Linie mindestens `length` Zellen mit diesem
// Wert direkt hintereinander enthält. Bei Vier Gewinnt gewinnt z. B., wer
// vier eigene Steine hintereinander in einer Linie hat.
func ContainsRun(line []string, value string, length int) bool {
	// Hinweis:
	// Zählen Sie mit, wie viele Zellen mit dem Wert gerade hintereinander
	// stehen. Eine andere Zelle setzt den Zähler wieder auf `0`.

	run := 0
	for _, cell := range line {
		if cell == value {
			run++
			if run >= length {
				return true
			}
		} else {
			run = 0
		}
	}
	return false
}
