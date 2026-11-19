package main

// Mit diesem Programm spielen zwei Personen Tic Tac Toe gegeneinander,
// abwechselnd an derselben Tastatur. Die Regeln stehen im Paket `tictactoe`,
// das Spielfeld in `grid`.
// Der Ablauf des Spiels in `main` ist vorgegeben; die Aufgabe ist das Ergebnis.

import (
	"fmt"

	"grid/grid"
	"grid/tictactoe"
)

func main() {
	fmt.Println("Die Felder sind so nummeriert:")
	fmt.Println("1 2 3")
	fmt.Println("4 5 6")
	fmt.Println("7 8 9")

	cells := tictactoe.NewBoard()
	player := "X"
	for !tictactoe.GameOver(cells) {
		fmt.Println()
		fmt.Print(grid.Format(cells))
		fmt.Printf("%s ist am Zug, Feld: ", player)
		var field int
		_, err := fmt.Scan(&field)
		if err != nil {
			fmt.Println("Keine gültige Eingabe, das Spiel wird beendet.")
			return
		}
		if !tictactoe.Place(cells, field, player) {
			fmt.Println("Dieses Feld gibt es nicht, oder es ist schon besetzt.")
			continue
		}
		player = tictactoe.NextPlayer(player)
	}

	fmt.Println()
	fmt.Print(grid.Format(cells))
	printResult(cells)
	// Der Ablauf eines Spiels:
	// - Die Schleife läuft, bis das Spiel vorbei ist.
	// - Lässt sich das Symbol nicht setzen, weil es das Feld nicht gibt oder es
	//   besetzt ist, beginnt mit `continue` der nächste Durchlauf,
	//   ohne dass der Zug wechselt.
	// - Liefert `fmt.Scan` einen Fehler, z. B. weil die Eingabe keine Zahl ist
	//   oder zu Ende ist (Strg+D), endet das Programm. Am Ende der Eingabe
	//   liefert `fmt.Scan` bei jedem Aufruf sofort wieder einen Fehler;
	//   ohne `return` liefe die Schleife dann endlos.
	// - `cells` wird nur über die Funktionen der Pakete geändert und gelesen.
}

// printResult erwartet ein Spielfeld am Ende eines Spiels und gibt aus,
// wer gewonnen hat, z. B. "X hat gewonnen!", oder "Unentschieden.".
func printResult(cells [][]string) {
	if tictactoe.Wins(cells, "X") {
		fmt.Println("X hat gewonnen!")
		return
	}
	if tictactoe.Wins(cells, "O") {
		fmt.Println("O hat gewonnen!")
		return
	}
	fmt.Println("Unentschieden.")

}
