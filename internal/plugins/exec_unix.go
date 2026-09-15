//go:build !windows

package plugins

import (
	"bytes"
	"context"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"syscall"
)

// setProcessGroup isole le plugin dans son propre groupe de processus.
func setProcessGroup(cmd *exec.Cmd) {
	cmd.SysProcAttr = &syscall.SysProcAttr{Setpgid: true}
}

// killProcessTree tue tout le groupe de processus du plugin.
func killProcessTree(cmd *exec.Cmd) error {
	if cmd.Process == nil {
		return nil
	}
	return syscall.Kill(-cmd.Process.Pid, syscall.SIGKILL)
}

// runExecRaw execute la commande brute (args JSON sur stdin, stdout brut en
// retour). Le plugin tourne dans son propre groupe de processus : a la fin
// du timeout (ou a l'annulation), tout l'arbre est tue — aucun orphelin ne
// survit, et cmd.Wait ne reste jamais bloque sur un pipe herite.
func runExecRaw(ctx context.Context, t *loadedTool, argsJSON string) ([]byte, error) {
	ex := t.def.Exec
	workdir, err := confinedDir(t.dir, ex.Dir)
	if err != nil {
		return nil, err
	}
	timeout := timeoutOf(ex.TimeoutSec)
	cctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	cmd := exec.Command(ex.Command[0], ex.Command[1:]...)
	setProcessGroup(cmd)
	cmd.Dir = workdir
	cmd.Stdin = strings.NewReader(argsJSON)
	cmd.Env = append(os.Environ(),
		"CETAS_PLUGIN_NAME="+t.def.Name,
		"CETAS_PLUGIN_DIR="+t.dir,
	)
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &limitedWriter{w: &stdout, n: maxStdout}
	cmd.Stderr = &limitedWriter{w: &stderr, n: maxStderr}

	if err := cmd.Start(); err != nil {
		return nil, fmt.Errorf("exec: demarrage: %w", err)
	}
	done := make(chan error, 1)
	go func() { done <- cmd.Wait() }()
	select {
	case werr := <-done:
		if werr != nil {
			msg := strings.TrimSpace(stderr.String())
			if msg == "" {
				msg = werr.Error()
			}
			return nil, fmt.Errorf("exec: %s", truncateRunes(msg, 500))
		}
	case <-cctx.Done():
		_ = killProcessTree(cmd)
		<-done // le groupe est mort : les pipes se ferment, Wait revient
		if ctx.Err() != nil {
			return nil, ctx.Err()
		}
		return nil, fmt.Errorf("exec: timeout depasse (%ds)", int(timeout.Seconds()))
	}
	return stdout.Bytes(), nil
}
