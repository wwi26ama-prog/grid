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

// Empty ist der Inhalt eines freien Felds.
const Empty = "."

// NewBoard liefert ein neues, leeres Spielfeld für Tic Tac Toe,
// in dem jedes Feld `Empty` enthält.
//
// TODO
func NewBoard() [][]string {
	// Hinweis:
	// Die Bibliothek `grid` legt Spielfelder beliebiger Größe an.

	return nil
}

// Place erwartet ein Spielfeld, die Nummer `field` eines Felds von `1` bis `9`
// und das Symbol `player` und setzt das Symbol in dieses Feld.
// Sie liefert `true`, wenn das geklappt hat, und `false`, wenn es das Feld
// nicht gibt oder es schon besetzt ist. Dann bleibt das Spielfeld unverändert.
//
// TODO
func Place(cells [][]string, field int, player string) bool {
	// Hinweis:
	// Die Felder werden ab `1` gezählt, Zeilen und Spalten ab `0`.
	// Jede Zeile hat drei Felder: Mit `field - 1` liefert die Ganzzahldivision
	// durch `3` die Zeile und der Rest die Spalte.

	return false
}

// NextPlayer erwartet das Symbol, das gerade am Zug war, und liefert
// das Symbol, das als Nächstes am Zug ist: "O" nach "X" und "X" nach "O".
//
// TODO
func NextPlayer(player string) string {
	return ""
}

// Wins erwartet ein Spielfeld und ein Symbol und liefert `true`, wenn eine
// Zeile, eine Spalte oder eine Diagonale ganz mit diesem Symbol gefüllt ist.
//
// TODO
func Wins(cells [][]string, player string) bool {
	// Hinweis:
	// Prüfen Sie in einer Schleife alle drei Zeilen und Spalten, danach
	// die beiden Diagonalen. Die Bibliothek liefert jede dieser Linien
	// und prüft, ob sie nur ein Symbol enthält.

	return false
}

// Winner erwartet ein Spielfeld und liefert das Symbol, das gewonnen hat,
// oder "", wenn noch niemand gewonnen hat.
//
// TODO
func Winner(cells [][]string) string {
	return ""
}

// IsFull erwartet ein Spielfeld und liefert `true`, wenn kein Feld mehr frei ist.
//
// TODO
func IsFull(cells [][]string) bool {
	return false
}

// GameOver erwartet ein Spielfeld und liefert `true`, wenn das Spiel vorbei ist,
// weil jemand gewonnen hat oder kein Feld mehr frei ist.
//
// TODO
func GameOver(cells [][]string) bool {
	return false
}
