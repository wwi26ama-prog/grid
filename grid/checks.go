package grid

// Diese Aufgaben prüfen eine Linie oder das ganze Spielfeld auf einen Wert,
// z. B. ob eine Zeile nur ein bestimmtes Symbol enthält
// oder ob noch ein Feld frei ist.

// AllEqual erwartet eine Linie, z. B. eine Zeile oder Spalte, und einen Wert
// und liefert `true`, wenn jede Zelle der Linie diesen Wert enthält.
// Für eine leere Linie liefert sie `true`, weil keine Zelle widerspricht.
func AllEqual(line []string, value string) bool {
	// Hinweis:
	// Ein Gegenbeispiel genügt: Sobald eine Zelle einen anderen Wert hat,
	// steht das Ergebnis fest.

	for _, cell := range line {
		if cell != value {
			return false
		}
	}
	return true
}

// Contains erwartet ein Spielfeld und einen Wert und liefert `true`,
// wenn mindestens eine Zelle diesen Wert enthält.
func Contains(cells [][]string, value string) bool {
	// Hinweis:
	// Eine äußere Schleife geht die Zeilen durch, eine innere die Zellen einer Zeile.

	for _, row := range cells {
		for _, cell := range row {
			if cell == value {
				return true
			}
		}
	}
	return false
	// Anmerkung:
	// Vorzeitiges Ende:
	// Das `return` beendet beide Schleifen und die Funktion,
	// sobald der Wert gefunden ist.
}
