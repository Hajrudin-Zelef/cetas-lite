package metrics

// macOS : les implementations /proc (Linux) et API noyau (Windows) ne
// s'appliquent pas. Sample() renvoie des champs nil (indisponibles), ce que
// l'interface sait deja afficher.

func cpuPercent() *float64 { return nil }

func ramStat() *MemStat { return nil }

func diskStat() *MemStat { return nil }

func netStat() *NetStat { return nil }
