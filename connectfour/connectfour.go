package connectfour

import "grid/grid"

const Empty = "."

func NewBoard() [][]string {
	return grid.New(6, 7, Empty)
}

// Place setzt einen Stein in der angegebenen Spalte
// des Spielfelds ein. Der Stein wandert so weit wie möglich nach unten.
// Liefert false zurück, falls der Zug nicht erlaubt ist.
//
// TODO
func Place(cells [][]string, column int, player string) bool {
	return false
}

// Aufgaben:
// Weitere Funktionen aus TicTacToe übernehmen und anpassen:
// - Wins
// - Winner
