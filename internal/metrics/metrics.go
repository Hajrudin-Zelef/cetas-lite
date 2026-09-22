// Package metrics expose les metriques systeme (CPU, RAM, disque, reseau)
// sans dependance externe. Disponibles sur Linux et Windows ; indisponibles
// sur macOS (champs nil).
package metrics

// MemStat decrit une ressource memoire/disque en octets.
type MemStat struct {
	Used  uint64 `json:"used"`
	Total uint64 `json:"total"`
}

// NetStat cumule les octets reseau depuis le demarrage (hors loopback).
type NetStat struct {
	RxBytes uint64 `json:"rx_bytes"`
	TxBytes uint64 `json:"tx_bytes"`
}

// Metrics regroupe les metriques ; un champ nil signifie indisponible.
type Metrics struct {
	CPU  *float64 `json:"cpu"`
	RAM  *MemStat `json:"ram"`
	Disk *MemStat `json:"disk"`
	Net  *NetStat `json:"net"`
}

// Sample lit les metriques systeme actuelles.
func Sample() Metrics {
	return Metrics{
		CPU:  cpuPercent(),
		RAM:  ramStat(),
		Disk: diskStat(),
		Net:  netStat(),
	}
}

func fptr(f float64) *float64 { return &f }
