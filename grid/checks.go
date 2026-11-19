package grid

import "slices"

// Diese Aufgaben prüfen eine Linie oder das ganze Spielfeld auf einen Wert,
// z. B. ob eine Zeile nur ein bestimmtes Symbol enthält
// oder ob noch ein Feld frei ist.

// AllEqual erwartet eine Linie, z. B. eine Zeile oder Spalte, und einen Wert
// und liefert `true`, wenn jede Zelle der Linie diesen Wert enthält.
// Für eine leere Linie liefert sie `true`, weil keine Zelle widerspricht.
func AllEqual(line []string, value string) bool {
	for _, element := range line {
		if value != element {
			return false
		}
	}
	return true
}

// Contains erwartet ein Spielfeld und einen Wert und liefert `true`,
// wenn mindestens eine Zelle diesen Wert enthält.
func Contains(cells [][]string, value string) bool {
	for _, row := range cells {

		if slices.Contains(row, value) {
			return true
		}

		// Handgeschriebene Alternative zu slices.Contains:
		// for _, cell := range row {
		// 	if cell == value {
		// 		return true
		// 	}
		// }
	}
	return false
}
