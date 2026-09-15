//go:build !windows

package mcp

import (
	"os/exec"
	"syscall"
)

// isolatePreStart isole le serveur MCP dans son propre groupe de processus
// (avant demarrage) : a la fermeture, tout l'arbre est tue, sans orphelins.
func isolatePreStart(cmd *exec.Cmd) {
	cmd.SysProcAttr = &syscall.SysProcAttr{Setpgid: true}
}

func isolatePostStart(cmd *exec.Cmd) uintptr { return 0 }

// killTree tue tout le groupe de processus du serveur MCP.
func killTree(cmd *exec.Cmd, _ uintptr) {
	if cmd.Process != nil {
		_ = syscall.Kill(-cmd.Process.Pid, syscall.SIGKILL)
	}
}
