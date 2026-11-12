// Package tictactoe enthält die Regeln von Tic Tac Toe.
// Es verwendet die Bibliothek `grid` für das Spielfeld.
package tictactoe

// Diese Aufgaben setzen die Regeln von Tic Tac Toe um. Das Spielfeld hat drei
// Zeilen und drei Spalten; zwei Personen setzen abwechselnd ihr Symbol
// "X" oder "O" in ein freies Feld. Wer zuerst eine ganze Zeile, Spalte
// oder Diagonale mit dem eigenen Symbol füllt, gewinnt.
// Die Felder sind für die Eingabe von `1` bis `9` nummeriert:
//
//	1 2 3
//	4 5 6
//	7 8 9
//
// Das Spielfeld wird nur über die Funktionen der Bibliothek `grid`
// gelesen und geändert.

import "grid/grid"

// Empty ist der Inhalt eines freien Felds.
const Empty = "."

// NewBoard liefert ein neues, leeres Spielfeld für Tic Tac Toe,
// in dem jedes Feld `Empty` enthält.
func NewBoard() [][]string {
	// Hinweis:
	// Die Bibliothek `grid` legt Spielfelder beliebiger Größe an.

	return grid.New(3, 3, Empty)
}

// Place erwartet ein Spielfeld, die Nummer `field` eines Felds von `1` bis `9`
// und das Symbol `player` und setzt das Symbol in dieses Feld.
// Sie liefert `true`, wenn das geklappt hat, und `false`, wenn es das Feld
// nicht gibt oder es schon besetzt ist. Dann bleibt das Spielfeld unverändert.
func Place(cells [][]string, field int, player string) bool {
	// Hinweis:
	// Die Felder werden ab `1` gezählt, Zeilen und Spalten ab `0`.
	// Jede Zeile hat drei Felder: Mit `field - 1` liefert die Ganzzahldivision
	// durch `3` die Zeile und der Rest die Spalte.

	if field < 1 || field > 9 {
		return false
	}
	row := (field - 1) / 3
	col := (field - 1) % 3
	if grid.Get(cells, row, col) != Empty {
		return false
	}
	grid.Set(cells, row, col, player)
	return true
	// Anmerkung:
	// Beispiel:
	// Feld `6` ist Zeile `5 / 3`, also `1`, und Spalte `5 % 3`, also `2`.
}

// NextPlayer erwartet das Symbol, das gerade am Zug war, und liefert
// das Symbol, das als Nächstes am Zug ist: "O" nach "X" und "X" nach "O".
func NextPlayer(player string) string {
	if player == "X" {
		return "O"
	}
	return "X"
}

// Wins erwartet ein Spielfeld und ein Symbol und liefert `true`, wenn eine
// Zeile, eine Spalte oder eine Diagonale ganz mit diesem Symbol gefüllt ist.
func Wins(cells [][]string, player string) bool {
	// Hinweis:
	// Prüfen Sie in einer Schleife alle drei Zeilen und Spalten, danach
	// die beiden Diagonalen. Die Bibliothek liefert jede dieser Linien
	// und prüft, ob sie nur ein Symbol enthält.

	for i := 0; i < 3; i++ {
		if grid.AllEqual(grid.Row(cells, i), player) || grid.AllEqual(grid.Col(cells, i), player) {
			return true
		}
	}
	return grid.AllEqual(grid.Diagonal(cells), player) ||
		grid.AllEqual(grid.AntiDiagonal(cells), player)
}

// Winner erwartet ein Spielfeld und liefert das Symbol, das gewonnen hat,
// oder "", wenn noch niemand gewonnen hat.
func Winner(cells [][]string) string {
	if Wins(cells, "X") {
		return "X"
	}
	if Wins(cells, "O") {
		return "O"
	}
	return ""
}

// IsFull erwartet ein Spielfeld und liefert `true`, wenn kein Feld mehr frei ist.
func IsFull(cells [][]string) bool {
	return !grid.Contains(cells, Empty)
}

// GameOver erwartet ein Spielfeld und liefert `true`, wenn das Spiel vorbei ist,
// weil jemand gewonnen hat oder kein Feld mehr frei ist.
func GameOver(cells [][]string) bool {
	return Winner(cells) != "" || IsFull(cells)
}
