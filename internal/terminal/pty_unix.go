//go:build !windows

package terminal

import (
	"io"
	"os"
	"os/exec"

	"github.com/creack/pty"
)

// startShell lance un shell interactif dans un vrai PTY.
func startShell(cwd string) (stdin io.WriteCloser, stdout io.Reader, resize func(cols, rows int) error, kill func() error, wait func() error, err error) {
	shell := os.Getenv("SHELL")
	if shell == "" {
		shell = "/bin/sh"
	}
	cmd := exec.Command(shell)
	cmd.Dir = cwd
	cmd.Env = append(os.Environ(), "TERM=xterm-256color", "COLORTERM=truecolor")
	ptmx, err := pty.Start(cmd)
	if err != nil {
		return nil, nil, nil, nil, nil, err
	}
	resize = func(cols, rows int) error {
		return pty.Setsize(ptmx, &pty.Winsize{Cols: uint16(cols), Rows: uint16(rows)})
	}
	kill = func() error {
		if cmd.Process != nil {
			_ = cmd.Process.Kill()
		}
		return ptmx.Close()
	}
	wait = cmd.Wait
	return ptmx, ptmx, resize, kill, wait, nil
}
