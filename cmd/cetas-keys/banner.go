package main

import "fmt"

// printSetupBanner — bannière reproduite à l'identique depuis setup.py
// (couleurs ANSI et espacements d'origine, inchangés).
func printSetupBanner() {
	fmt.Println("\x1b[38;5;34m  ╔════════════════════════════════════════════════════════════════╗\x1b[0m")
	fmt.Println("\x1b[38;5;34m  ║░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░║\x1b[0m")
	fmt.Println("\x1b[38;5;34m  ║                                                                ║\x1b[0m")
	fmt.Println("\x1b[38;5;34m  ║\x1b[1m   █████╗ ███████╗████████╗ █████╗ ███████╗\x1b[0m     \x1b[38;5;34m║\x1b[0m")
	fmt.Println("\x1b[38;5;34m  ║\x1b[1m  ██╔══██╗██╔════╝╚══██╔══╝██╔══██╗██╔════╝\x1b[0m     \x1b[38;5;34m║\x1b[0m")
	fmt.Println("\x1b[38;5;34m  ║\x1b[1m  ███████║█████╗     ██║   ███████║███████╗\x1b[0m     \x1b[38;5;34m║\x1b[0m")
	fmt.Println("\x1b[38;5;34m  ║\x1b[38;5;28m  ██╔══██║██╔══╝     ██║   ██╔══██║╚════██║\x1b[0m     \x1b[38;5;34m║\x1b[0m")
	fmt.Println("\x1b[38;5;34m  ║\x1b[38;5;28m  ██║  ██║███████╗   ██║   ██║  ██║███████║\x1b[0m     \x1b[38;5;34m║\x1b[0m")
	fmt.Println("\x1b[38;5;34m  ║\x1b[38;5;28m  ╚═╝  ╚═╝╚══════╝   ╚═╝   ╚═╝  ╚═╝╚══════╝\x1b[0m     \x1b[38;5;34m║\x1b[0m")
	fmt.Println("\x1b[38;5;34m  ║\x1b[0m                                        \x1b[2m\x1b[97mby Marexsoft Corporation\x1b[0m  \x1b[38;5;34m║\x1b[0m")
	fmt.Println("\x1b[38;5;34m  ╠────────────────────────────────────────────────────────────────╣\x1b[0m")
	fmt.Println("\x1b[38;5;34m  ║\x1b[0m   \x1b[38;5;34m▶\x1b[0m \x1b[1m\x1b[97mApp\x1b[0m  \x1b[38;5;34m│\x1b[0m  \x1b[97mCETAS\x1b[0m                              \x1b[2m\x1b[97mv1.0\x1b[0m           \x1b[38;5;34m║\x1b[0m")
	fmt.Println("\x1b[38;5;34m  ║\x1b[0m   \x1b[2m\x1b[3m\x1b[97mTous droits réservés © Marexsoft Corporation\x1b[0m                    \x1b[38;5;34m║\x1b[0m")
	fmt.Println("\x1b[38;5;34m  ╚════════════════════════════════════════════════════════════════╝\x1b[0m")
}
