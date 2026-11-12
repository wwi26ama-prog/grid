package grid

// Diese Aufgaben liefern eine Linie aus dem Spielfeld: eine Zeile,
// eine Spalte oder eine Diagonale. Jede Linie ist eine neue Liste;
// wird sie geändert, bleibt das Spielfeld unverändert.

// Row erwartet ein Spielfeld und die Nummer `index` einer Zeile
// und liefert die Zellen dieser Zeile von links nach rechts.
//
// TODO
func Row(cells [][]string, index int) []string {
	// Hinweis:
	// Kopieren Sie die Zellen der Zeile in eine neue Liste.
	// Würde die Zeile selbst zurückgegeben, könnte der Aufrufer
	// über das Ergebnis das Spielfeld ändern.

	return nil
}

// Col erwartet ein Spielfeld und die Nummer `index` einer Spalte
// und liefert die Zellen dieser Spalte von oben nach unten.
//
// TODO
func Col(cells [][]string, index int) []string {
	// Hinweis:
	// Eine Spalte steht nicht am Stück im Spielfeld:
	// Aus jeder Zeile wird die Zelle an derselben Stelle geholt.

	return nil
}

// Diagonal erwartet ein Spielfeld und liefert die Diagonale von links oben
// nach rechts unten: die Zellen, deren Zeile und Spalte gleich sind.
// Bei einem Spielfeld, das nicht quadratisch ist, endet sie am Rand.
//
// TODO
func Diagonal(cells [][]string) []string {
	// Hinweis:
	// Die Diagonale endet, sobald die Zeilen oder die Spalten zu Ende sind.

	return nil
}

// AntiDiagonal erwartet ein Spielfeld und liefert die Diagonale von rechts oben
// nach links unten. Bei einem Spielfeld, das nicht quadratisch ist,
// endet sie am Rand.
//
// TODO
func AntiDiagonal(cells [][]string) []string {
	// Hinweis:
	// Die Zeile läuft von oben nach unten, die Spalte gleichzeitig
	// von rechts (`Cols(cells) - 1`) nach links.

	return nil
}
