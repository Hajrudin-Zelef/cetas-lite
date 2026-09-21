//go:build !windows

package vfs

import "syscall"

// diskFreeBytes retourne l'espace disque libre (octets) du volume contenant
// path. ok=false si l'information est indisponible. Unix uniquement :
// syscall.Statfs n'existe pas sous Windows (voir diskfree_windows.go).
func diskFreeBytes(path string) (int64, bool) {
	var st syscall.Statfs_t
	if err := syscall.Statfs(path, &st); err != nil {
		return 0, false
	}
	return int64(st.Bavail) * int64(st.Bsize), true
}
