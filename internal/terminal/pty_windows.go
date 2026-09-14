//go:build windows

package terminal

import (
	"errors"
	"fmt"
	"io"
	"os"
	"sync"
	"unsafe"

	"golang.org/x/sys/windows"
)

// conptyProc encapsule un processus Windows rattache a un pseudo-terminal
// (ConPTY). L'entree/sortie passe par des pipes anonymes, ce qui donne une
// vraie emulation de terminal (programmes interactifs, redimensionnement,
// signaux) contrairement a de simples pipes sur cmd.exe.
type conptyProc struct {
	hPC     windows.Handle
	hProc   windows.Handle
	hThread windows.Handle
	stdin   *os.File
	stdout  *os.File
	closeMu sync.Mutex
	closed  bool
}

func newConpty(shell, cwd string, cols, rows int) (*conptyProc, error) {
	// Pipes : nous ecrivons dans inW (la console lit inR),
	// nous lisons depuis outR (la console ecrit outW).
	var inR, inW, outR, outW windows.Handle
	if err := windows.CreatePipe(&inR, &inW, nil, 0); err != nil {
		return nil, fmt.Errorf("pipe stdin: %w", err)
	}
	if err := windows.CreatePipe(&outR, &outW, nil, 0); err != nil {
		windows.CloseHandle(inR)
		windows.CloseHandle(inW)
		return nil, fmt.Errorf("pipe stdout: %w", err)
	}
	cleanup := func() {
		windows.CloseHandle(inR)
		windows.CloseHandle(inW)
		windows.CloseHandle(outR)
		windows.CloseHandle(outW)
	}

	var hPC windows.Handle
	if err := windows.CreatePseudoConsole(
		windows.Coord{X: int16(cols), Y: int16(rows)},
		inR, outW, 0, &hPC,
	); err != nil {
		cleanup()
		return nil, fmt.Errorf("CreatePseudoConsole: %w", err)
	}
	cleanupPC := func() {
		windows.ClosePseudoConsole(hPC)
		cleanup()
	}

	// Liste d'attributs avec le pseudo-terminal.
	al, err := windows.NewProcThreadAttributeList(1)
	if err != nil {
		cleanupPC()
		return nil, fmt.Errorf("attributs processus: %w", err)
	}
	defer al.Delete()
	// Note : la doc Microsoft ("Creating a Pseudoconsole session") passe le
	// HPCON lui-meme (cast en pointeur), pas un pointeur vers le handle.
	// `go vet` (analyseur unsafeptr) signale ce cast sous Windows : c'est un
	// faux positif, le pattern est impose par l'API.
	if err := al.Update(
		windows.PROC_THREAD_ATTRIBUTE_PSEUDOCONSOLE,
		unsafe.Pointer(hPC),
		unsafe.Sizeof(hPC),
	); err != nil {
		cleanupPC()
		return nil, fmt.Errorf("attribut pseudoconsole: %w", err)
	}

	siEx := windows.StartupInfoEx{}
	siEx.StartupInfo.Cb = uint32(unsafe.Sizeof(siEx))
	siEx.ProcThreadAttributeList = al.List()

	cmdLine, err := windows.UTF16PtrFromString(shell)
	if err != nil {
		cleanupPC()
		return nil, err
	}
	var dir *uint16
	if cwd != "" {
		dir, err = windows.UTF16PtrFromString(cwd)
		if err != nil {
			cleanupPC()
			return nil, err
		}
	}
	var pi windows.ProcessInformation
	// StartupInfoEx commence par un StartupInfo : le cast est sur.
	err = windows.CreateProcess(
		nil, cmdLine,
		nil, nil, false,
		windows.EXTENDED_STARTUPINFO_PRESENT,
		nil, dir,
		(*windows.StartupInfo)(unsafe.Pointer(&siEx)),
		&pi,
	)
	if err != nil {
		cleanupPC()
		return nil, fmt.Errorf("CreateProcess: %w", err)
	}
	// La console a duplique les extremites qui lui etaient confiees ;
	// nous pouvons fermer nos copies.
	windows.CloseHandle(inR)
	windows.CloseHandle(outW)

	return &conptyProc{
		hPC:     hPC,
		hProc:   pi.Process,
		hThread: pi.Thread,
		stdin:   os.NewFile(uintptr(inW), "conpty-stdin"),
		stdout:  os.NewFile(uintptr(outR), "conpty-stdout"),
	}, nil
}

// close libere le pseudo-terminal et les handles (idempotent). Fermer le
// pseudo-terminal provoque EOF sur stdout, ce qui debloque la boucle de
// lecture du gestionnaire.
func (p *conptyProc) close() {
	p.closeMu.Lock()
	defer p.closeMu.Unlock()
	if p.closed {
		return
	}
	p.closed = true
	windows.ClosePseudoConsole(p.hPC)
	windows.CloseHandle(p.hThread)
	windows.CloseHandle(p.hProc)
	if p.stdin != nil {
		p.stdin.Close()
	}
	if p.stdout != nil {
		p.stdout.Close()
	}
}

// startShell lance un shell interactif dans un vrai pseudo-terminal Windows
// (ConPTY). stdout transporte la sortie brute du terminal (sequences ANSI
// incluses), prete pour xterm.js.
func startShell(cwd string) (stdin io.WriteCloser, stdout io.Reader, resize func(cols, rows int) error, kill func() error, wait func() error, err error) {
	shell := os.Getenv("COMSPEC")
	if shell == "" {
		shell = "cmd.exe"
	}
	p, err := newConpty(shell, cwd, 120, 30)
	if err != nil {
		return nil, nil, nil, nil, nil, err
	}
	resize = func(cols, rows int) error {
		if cols < 1 || rows < 1 {
			return errors.New("dimensions invalides")
		}
		return windows.ResizePseudoConsole(p.hPC, windows.Coord{X: int16(cols), Y: int16(rows)})
	}
	kill = func() error {
		_ = windows.TerminateProcess(p.hProc, 1)
		p.close()
		return nil
	}
	wait = func() error {
		// Attend la fin du processus, puis ferme le pseudo-terminal pour
		// que la boucle de lecture recoive EOF.
		_, _ = windows.WaitForSingleObject(p.hProc, windows.INFINITE)
		var code uint32
		_ = windows.GetExitCodeProcess(p.hProc, &code)
		p.close()
		if code != 0 {
			return fmt.Errorf("processus termine (code %d)", code)
		}
		return nil
	}
	return p.stdin, p.stdout, resize, kill, wait, nil
}
