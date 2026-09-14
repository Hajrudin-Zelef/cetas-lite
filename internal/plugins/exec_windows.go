//go:build windows

package plugins

import (
	"bytes"
	"context"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"sync"
	"time"
	"unsafe"

	"golang.org/x/sys/windows"
)

// quoteWindowsArg cite un argument selon les regles MSVCRT (guillemets et
// antislashs echappes) pour reconstruire une ligne de commande exacte.
func quoteWindowsArg(s string) string {
	if s == "" {
		return `""`
	}
	if !strings.ContainsAny(s, " \t\n\v\"") {
		return s
	}
	var b strings.Builder
	b.WriteByte('"')
	slashes := 0
	flushSlashes := func() {
		for i := 0; i < slashes; i++ {
			b.WriteByte('\\')
		}
		slashes = 0
	}
	for _, r := range s {
		switch r {
		case '\\':
			slashes++
		case '"':
			for i := 0; i < slashes; i++ {
				b.WriteString(`\\`)
			}
			slashes = 0
			b.WriteString(`\"`)
		default:
			flushSlashes()
			b.WriteRune(r)
		}
	}
	for i := 0; i < slashes; i++ {
		b.WriteString(`\\`)
	}
	b.WriteByte('"')
	return b.String()
}

// createKillJob cree un Job Object qui tue tous ses processus a la
// fermeture (filet de securite) et sur TerminateJobObject.
func createKillJob() (windows.Handle, error) {
	job, err := windows.CreateJobObject(nil, nil)
	if err != nil {
		return 0, err
	}
	var info windows.JOBOBJECT_EXTENDED_LIMIT_INFORMATION
	info.BasicLimitInformation.LimitFlags = windows.JOB_OBJECT_LIMIT_KILL_ON_JOB_CLOSE
	if _, err := windows.SetInformationJobObject(job, windows.JobObjectExtendedLimitInformation,
		uintptr(unsafe.Pointer(&info)), uint32(unsafe.Sizeof(info))); err != nil {
		windows.CloseHandle(job)
		return 0, err
	}
	return job, nil
}

// inheritPipe cree un tube anonyme dont les deux extremites sont heritables.
func inheritPipe() (r, w windows.Handle, err error) {
	sa := &windows.SecurityAttributes{InheritHandle: 1}
	if err := windows.CreatePipe(&r, &w, sa, 0); err != nil {
		return 0, 0, err
	}
	return r, w, nil
}

// noInherit retire le flag d'heritage (extremites gardees par le parent).
func noInherit(h windows.Handle) {
	windows.SetHandleInformation(h, windows.HANDLE_FLAG_INHERIT, 0)
}

// envBlock construit le bloc d'environnement UTF-16 pour CreateProcess.
func envBlock(extra map[string]string) *uint16 {
	merged := map[string]string{}
	for _, kv := range os.Environ() {
		if i := strings.IndexByte(kv, '='); i > 0 {
			merged[kv[:i]] = kv[i+1:]
		}
	}
	for k, v := range extra {
		merged[k] = v
	}
	var block []uint16
	for k, v := range merged {
		kv, _ := windows.UTF16FromString(k + "=" + v)
		block = append(block, kv...)
	}
	block = append(block, 0)
	return &block[0]
}

// runExecRaw execute la commande brute (args JSON sur stdin, stdout brut en
// retour). Le plugin demarre SUSPENDU, est assigne a un Job Object, puis
// reprend : en cas de timeout (ou d'annulation), TerminateJobObject tue tout
// l'arbre (petits-enfants inclus) — aucun orphelin ne survit sous Windows.
func runExecRaw(ctx context.Context, t *loadedTool, argsJSON string) ([]byte, error) {
	ex := t.def.Exec
	workdir, err := confinedDir(t.dir, ex.Dir)
	if err != nil {
		return nil, err
	}
	timeout := timeoutOf(ex.TimeoutSec)
	cctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	bin, err := exec.LookPath(ex.Command[0])
	if err != nil {
		return nil, fmt.Errorf("exec: demarrage: %w", err)
	}
	parts := make([]string, 0, len(ex.Command))
	parts = append(parts, quoteWindowsArg(bin))
	for _, a := range ex.Command[1:] {
		parts = append(parts, quoteWindowsArg(a))
	}
	cmdLine, err := windows.UTF16PtrFromString(strings.Join(parts, " "))
	if err != nil {
		return nil, fmt.Errorf("exec: demarrage: %w", err)
	}
	dirPtr, err := windows.UTF16PtrFromString(workdir)
	if err != nil {
		return nil, fmt.Errorf("exec: demarrage: %w", err)
	}

	job, err := createKillJob()
	if err != nil {
		return nil, fmt.Errorf("exec: job: %w", err)
	}
	defer windows.CloseHandle(job)

	stdinR, stdinW, err := inheritPipe()
	if err != nil {
		return nil, fmt.Errorf("exec: tube: %w", err)
	}
	stdoutR, stdoutW, err := inheritPipe()
	if err != nil {
		windows.CloseHandle(stdinR)
		windows.CloseHandle(stdinW)
		return nil, fmt.Errorf("exec: tube: %w", err)
	}
	stderrR, stderrW, err := inheritPipe()
	if err != nil {
		windows.CloseHandle(stdinR)
		windows.CloseHandle(stdinW)
		windows.CloseHandle(stdoutR)
		windows.CloseHandle(stdoutW)
		return nil, fmt.Errorf("exec: tube: %w", err)
	}
	// Extremites parent : non heritables.
	noInherit(stdinW)
	noInherit(stdoutR)
	noInherit(stderrR)

	si := &windows.StartupInfo{}
	si.Cb = uint32(unsafe.Sizeof(*si))
	si.Flags = windows.STARTF_USESTDHANDLES
	si.StdInput = stdinR
	si.StdOutput = stdoutW
	si.StdErr = stderrW
	var pi windows.ProcessInformation
	// Demarrage suspendu : aucun petit-enfant ne peut echapper au job entre
	// la creation du processus et son assignation.
	err = windows.CreateProcess(nil, cmdLine, nil, nil, true,
		windows.CREATE_SUSPENDED, envBlock(map[string]string{
			"CETAS_PLUGIN_NAME": t.def.Name,
			"CETAS_PLUGIN_DIR":  t.dir,
		}), dirPtr, si, &pi)
	// L'enfant possede ses copies des extremites heritees : on ferme les
	// notres dans tous les cas.
	windows.CloseHandle(stdinR)
	windows.CloseHandle(stdoutW)
	windows.CloseHandle(stderrW)
	if err != nil {
		windows.CloseHandle(stdinW)
		windows.CloseHandle(stdoutR)
		windows.CloseHandle(stderrR)
		return nil, fmt.Errorf("exec: demarrage: %w", err)
	}
	defer windows.CloseHandle(pi.Process)
	defer windows.CloseHandle(pi.Thread)

	if err := windows.AssignProcessToJobObject(job, pi.Process); err != nil {
		windows.TerminateProcess(pi.Process, 1)
		windows.CloseHandle(stdinW)
		windows.CloseHandle(stdoutR)
		windows.CloseHandle(stderrR)
		return nil, fmt.Errorf("exec: job: %w", err)
	}
	if _, err := windows.ResumeThread(pi.Thread); err != nil {
		windows.TerminateJobObject(job, 1)
		windows.CloseHandle(stdinW)
		windows.CloseHandle(stdoutR)
		windows.CloseHandle(stderrR)
		return nil, fmt.Errorf("exec: reprise: %w", err)
	}

	var stdout, stderr bytes.Buffer
	lwOut := &limitedWriter{w: &stdout, n: maxStdout}
	lwErr := &limitedWriter{w: &stderr, n: maxStderr}
	var wg sync.WaitGroup
	wg.Add(3)
	// stdin : ecriture du (petit) payload JSON puis fermeture.
	go func() {
		defer wg.Done()
		defer windows.CloseHandle(stdinW)
		var done uint32
		_ = windows.WriteFile(stdinW, []byte(argsJSON), &done, nil)
	}()
	readTo := func(h windows.Handle, lw *limitedWriter) {
		defer wg.Done()
		var buf [8192]byte
		for {
			var n uint32
			if err := windows.ReadFile(h, buf[:], &n, nil); err != nil {
				return
			}
			if n > 0 {
				lw.Write(buf[:n])
			}
		}
	}
	go readTo(stdoutR, lwOut)
	go readTo(stderrR, lwErr)

	type waitRes struct {
		code uint32
		err  error
	}
	wch := make(chan waitRes, 1)
	go func() {
		_, werr := windows.WaitForSingleObject(pi.Process, windows.INFINITE)
		var code uint32
		_ = windows.GetExitCodeProcess(pi.Process, &code)
		wch <- waitRes{code: code, err: werr}
	}()

	var rerr error
	select {
	case r := <-wch:
		if r.err != nil {
			rerr = fmt.Errorf("exec: attente: %w", r.err)
		} else if r.code != 0 {
			msg := strings.TrimSpace(stderr.String())
			if msg == "" {
				msg = fmt.Sprintf("code de sortie %d", r.code)
			}
			rerr = fmt.Errorf("exec: %s", truncateRunes(msg, 500))
		}
	case <-cctx.Done():
		// Timeout ou annulation : tout l'arbre du job est tue d'un coup.
		_ = windows.TerminateJobObject(job, 1)
		<-wch // le job est mort : les pipes se ferment, les lecteurs reviennent
		if ctx.Err() != nil {
			rerr = ctx.Err()
		} else {
			rerr = fmt.Errorf("exec: timeout depasse (%ds)", int(timeout.Seconds()))
		}
	}

	// Les lecteurs terminent quand les pipes se ferment. Un petit-enfant
	// survivant (sortie normale) peut retenir un pipe : grace de 2 s, puis
	// on ferme les extremites de lecture pour ne jamais rester bloque.
	closeReads := func() {
		windows.CloseHandle(stdoutR)
		windows.CloseHandle(stderrR)
	}
	readersDone := make(chan struct{})
	go func() { wg.Wait(); close(readersDone) }()
	select {
	case <-readersDone:
		closeReads()
	case <-time.After(2 * time.Second):
		closeReads() // debloque les ReadFile en attente
		<-readersDone
	}

	if rerr != nil {
		return nil, rerr
	}
	return stdout.Bytes(), nil
}
