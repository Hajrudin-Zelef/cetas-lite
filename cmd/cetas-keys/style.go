package main

import "fmt"

// Style visuel repris de setup.py : cadres cyan, textes blancs,
// succès verts, erreurs rouges, infos secondaires grises.
// Tout est indenté de 2 espaces, comme l'original.

const (
	sReset  = "\x1b[0m"
	sBold   = "\x1b[1m"
	sGreen  = "\x1b[92m"
	sCyan   = "\x1b[96m"
	sYellow = "\x1b[93m"
	sRed    = "\x1b[91m"
	sWhite  = "\x1b[97m"
	sGray   = "\x1b[90m"
)

func paint(col, s string) string { return col + s + sReset }

func white(s string) string  { return paint(sWhite, s) }
func green(s string) string  { return paint(sGreen, s) }
func red(s string) string    { return paint(sRed, s) }
func gray(s string) string   { return paint(sGray, s) }
func cyanS(s string) string  { return paint(sCyan, s) }
func yellow(s string) string { return paint(sYellow, s) }
func bold(s string) string   { return paint(sBold, s) }

// frame — cadre de section façon setup.py : "  ┌─ TITRE ─────┐" en cyan.
func frame(title string) {
	inner := "─ " + title + " "
	for len([]rune(inner)) < 46 {
		inner += "─"
	}
	fmt.Println("  " + paint(sCyan, "┌"+inner+"┐"))
}

// subtitle — sous-titre façon setup.py : "  --- texte ---" en cyan.
func subtitle(text string) {
	fmt.Println("  " + paint(sCyan, "--- "+text+" ---"))
}

// Messages standard (vert = succès, rouge = erreur, gris = info secondaire).
func okLine(s string)   { fmt.Println(green("  " + s)) }
func errLine(s string)  { fmt.Println(red("  " + s)) }
func infoLine(s string) { fmt.Println(gray("  " + s)) }
