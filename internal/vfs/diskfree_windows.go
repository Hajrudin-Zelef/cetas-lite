//go:build windows

package vfs

// diskFreeBytes : espace libre indisponible sans appel systeme supplementaire
// (syscall.Statfs n'existe pas sous Windows). ok=false -> le garde-fou F6.5
// laisse le cache fonctionner ; la taille du miroir reste bornee par
// CacheL2MaxBytes.
func diskFreeBytes(string) (int64, bool) { return 0, false }
