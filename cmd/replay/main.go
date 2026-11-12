package main

// Dieses Programm spielt eine Partie Tic Tac Toe aus einer Liste von Zügen nach,
// z. B. um eine Partie zu analysieren oder zu testen. Es verwendet die Pakete
// `grid` und `tictactoe`; das Programm `play` ist ähnlich aufgebaut.
// Die Züge in `main` sind vorgegeben; die Aufgabe ist `replay`.

func main() {
	replay([]int{5, 1, 9, 3, 2, 8, 7, 4, 6})
}

// replay erwartet eine Liste von Feldnummern und spielt sie als Züge nach,
// abwechselnd für "X" und "O", beginnend mit "X". Ist ein Zug nicht möglich,
// gibt sie z. B. "Ungültiger Zug: 5" aus, und mit dem nächsten Zug der Liste
// ist dasselbe Symbol noch einmal an der Reihe. Ist das Spiel vorbei, werden
// weitere Züge nicht mehr gespielt. Am Ende gibt sie das Spielfeld aus und
// darunter das Ergebnis: "X hat gewonnen!", "O hat gewonnen!",
// "Unentschieden." oder "Das Spiel ist noch nicht vorbei.".
//
// TODO
func replay(moves []int) {
	// Hinweis:
	// Statt der Eingabe in `play` liefert hier die Liste die Züge.
	// Mit `break` endet eine Schleife vorzeitig.
}
