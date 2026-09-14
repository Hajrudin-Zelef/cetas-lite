//go:build windows

package mcp

import (
	"fmt"
	"os/exec"
	"unsafe"

	"golang.org/x/sys/windows"
)

func isolatePreStart(cmd *exec.Cmd) {}

// isolatePostStart assigne le serveur MCP a un Job Object (au mieux) : a la
// fermeture, TerminateJobObject tue tout l'arbre, sans orphelins.
func isolatePostStart(cmd *exec.Cmd) uintptr {
	if cmd.Process == nil {
		return 0
	}
	job, err := windows.CreateJobObject(nil, nil)
	if err != nil {
		return 0
	}
	var info windows.JOBOBJECT_EXTENDED_LIMIT_INFORMATION
	info.BasicLimitInformation.LimitFlags = windows.JOB_OBJECT_LIMIT_KILL_ON_JOB_CLOSE
	if _, err := windows.SetInformationJobObject(job, windows.JobObjectExtendedLimitInformation,
		uintptr(unsafe.Pointer(&info)), uint32(unsafe.Sizeof(info))); err != nil {
		windows.CloseHandle(job)
		return 0
	}
	ph, err := openProcHandle(cmd)
	if err != nil {
		windows.CloseHandle(job)
		return 0
	}
	err = windows.AssignProcessToJobObject(job, ph)
	windows.CloseHandle(ph) // le job garde sa propre reference
	if err != nil {
		windows.CloseHandle(job)
		return 0
	}
	return uintptr(job)
}

// openProcHandle ouvre un handle sur le processus avec les droits requis
// pour l'assignation a un Job Object (os.Process n'expose pas son handle).
func openProcHandle(cmd *exec.Cmd) (windows.Handle, error) {
	if cmd.Process == nil {
		return 0, fmt.Errorf("processus absent")
	}
	return windows.OpenProcess(
		windows.PROCESS_SET_QUOTA|windows.PROCESS_TERMINATE,
		false, uint32(cmd.Process.Pid))
}

// killTree tue tout l'arbre du serveur MCP via le Job Object.
func killTree(cmd *exec.Cmd, job uintptr) {
	if job != 0 {
		h := windows.Handle(job)
		_ = windows.TerminateJobObject(h, 1)
		windows.CloseHandle(h)
	}
	if cmd.Process != nil {
		_ = cmd.Process.Kill()
	}
}
