package extensions

// Diese Aufgaben verwenden die Bibliothek `grid` und das Paket `tictactoe`
// für weitere Spiele: einen Überblick über die freien Felder,
// z. B. für einen Computergegner, und den ersten Schritt zu Vier Gewinnt.

// FreeFields erwartet ein Spielfeld für Tic Tac Toe und liefert die Nummern
// der freien Felder von `1` bis `9` in aufsteigender Reihenfolge.
// Für ein leeres Spielfeld sind das alle Nummern von `1` bis `9`.
//
// TODO
func FreeFields(cells [][]string) []int {
	// Hinweis:
	// Gehen Sie alle Zeilen und Spalten durch. Die Nummer eines Felds ergibt
	// sich aus Zeile und Spalte umgekehrt wie in `tictactoe.Place`:
	// Zeile mal `3`, plus Spalte, plus `1`.

	return nil
}

// DropPiece erwartet ein Spielfeld für Vier Gewinnt, die Nummer `col` einer
// Spalte, gezählt ab `0`, und ein Symbol `piece`. Wie bei Tic Tac Toe enthalten freie Felder
// `tictactoe.Empty`. Bei Vier Gewinnt fällt ein Stein in einer Spalte nach
// unten, bis er auf den Boden oder einen anderen Stein trifft. Die Funktion
// setzt das Symbol deshalb in das unterste freie Feld der Spalte und liefert
// `true`. Ist die Spalte schon voll, liefert sie `false`.
//
// Weiterführende Aufgabe.
//
// TODO
func DropPiece(cells [][]string, col int, piece string) bool {
	// Hinweis:
	// Die unterste Zeile hat die Nummer `grid.Rows(cells) - 1`.
	// Gehen Sie die Zeilen von unten nach oben durch.

	return false
}
