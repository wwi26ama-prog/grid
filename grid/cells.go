// Package grid enthält Funktionen für Spielfelder aus Zeilen und Spalten,
// wie sie z. B. Tic Tac Toe oder Vier Gewinnt brauchen.
package grid

// Diese Aufgaben sind die Grundlagen: ein Spielfeld anlegen, seine Größe
// abfragen, einzelne Zellen lesen und schreiben und das Spielfeld ausgeben.
// Ein Spielfeld ist eine Liste von Zeilen, jede Zeile eine Liste von Zellen,
// also `[][]string`. Zeilen und Spalten werden ab `0` gezählt.

import "strings"

// New erwartet die Anzahl der Zeilen `rows` und Spalten `cols` und einen Wert
// `value` und liefert ein neues Spielfeld, in dem jede Zelle `value` enthält.
// Mit `"."` als Wert entsteht z. B. ein leeres Spielfeld.
func New(rows, cols int, value string) [][]string {
	// Hinweis:
	// Bauen Sie in einer inneren Schleife eine Zeile auf und hängen Sie
	// in einer äußeren Schleife die Zeilen an das Spielfeld an.
	// Jede Zeile muss eine eigene Liste sein.

	cells := [][]string{}
	for r := 0; r < rows; r++ {
		row := []string{}
		for c := 0; c < cols; c++ {
			row = append(row, value)
		}
		cells = append(cells, row)
	}
	return cells
	// Anmerkung:
	// Eigene Zeilen:
	// Weil jede Zeile in der äußeren Schleife neu angelegt wird,
	// ändert sich beim Setzen einer Zelle nur diese eine.
}

// Rows erwartet ein Spielfeld und liefert die Anzahl seiner Zeilen.
func Rows(cells [][]string) int {
	// Hinweis:
	// Ein Spielfeld ist eine Liste von Zeilen.

	return len(cells)
}

// Cols erwartet ein Spielfeld und liefert die Anzahl seiner Spalten,
// also die Länge der ersten Zeile. Ein Spielfeld ohne Zeilen hat `0` Spalten.
func Cols(cells [][]string) int {
	// Hinweis:
	// Alle Zeilen sind gleich lang. Die erste Zeile gibt es aber nur,
	// wenn das Spielfeld überhaupt Zeilen hat.

	if len(cells) == 0 {
		return 0
	}
	return len(cells[0])
}

// Get erwartet ein Spielfeld und die Zeile `row` und Spalte `col` einer Zelle
// und liefert ihren Inhalt.
func Get(cells [][]string, row, col int) string {
	// Hinweis:
	// `cells[row]` ist eine Zeile, also wieder eine Liste.

	return cells[row][col]
}

// Set erwartet ein Spielfeld, die Zeile `row` und Spalte `col` einer Zelle
// und einen Wert `value` und schreibt den Wert in die Zelle.
// Die Änderung ist auch beim Aufrufer zu sehen.
func Set(cells [][]string, row, col int, value string) {
	// Hinweis:
	// Eine Liste als Parameter verweist auf dieselben Zeilen und Zellen
	// wie die Liste des Aufrufers. Die Funktion muss deshalb nichts zurückgeben.

	cells[row][col] = value
}

// Format erwartet ein Spielfeld und liefert es als Text zur Ausgabe:
// eine Textzeile pro Zeile des Spielfelds, die Zellen durch Leerzeichen
// getrennt, jede Textzeile mit einem Zeilenumbruch "\n" am Ende.
func Format(cells [][]string) string {
	// Hinweis:
	// `strings.Join(row, " ")` setzt die Zellen einer Zeile zu einem Text
	// zusammen, mit einem Leerzeichen zwischen je zwei Zellen.

	text := ""
	for _, row := range cells {
		text += strings.Join(row, " ") + "\n"
	}
	return text
}
