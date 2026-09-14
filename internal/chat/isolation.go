package chat

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"time"
)

const (
	IsolationNone  = "none"
	IsolationBwrap = "bwrap"
)

var bwrapROPaths = []string{"/usr", "/bin", "/sbin", "/lib", "/lib32", "/lib64", "/libx32", "/etc", "/opt"}

func ResolveIsolation(mode, probeRoot string) (string, error) {
	switch mode {
	case "", IsolationNone:
		return IsolationNone, nil
	case IsolationBwrap:
		if !bwrapAvailable() {
			return "", fmt.Errorf("CETAS_LITE_SANDBOX=bwrap mais bwrap introuvable")
		}
		if !probeBwrap(probeRoot) {
			return "", fmt.Errorf("CETAS_LITE_SANDBOX=bwrap mais la sonde a echoue (namespaces indisponibles ?)")
		}
		return IsolationBwrap, nil
	case "auto":
		if bwrapAvailable() && probeBwrap(probeRoot) {
			return IsolationBwrap, nil
		}
		return IsolationNone, nil
	}
	return "", fmt.Errorf("mode d'isolation invalide: %q", mode)
}

func bwrapAvailable() bool {
	_, err := exec.LookPath("bwrap")
	return err == nil
}

func bwrapArgv(root string, inner []string) []string {
	argv := []string{
		"bwrap",
		"--unshare-all",
		"--share-net",
		"--die-with-parent",
		"--new-session",
		"--proc", "/proc",
		"--dev", "/dev",
		"--tmpfs", "/tmp",
	}
	for _, p := range bwrapROPaths {
		if _, err := os.Stat(p); err == nil {
			argv = append(argv, "--ro-bind", p, p)
		}
	}
	argv = append(argv,
		"--bind", root, root,
		"--chdir", root,
		"--setenv", "HOME", root,
		"--",
	)
	return append(argv, inner...)
}

func probeBwrap(root string) bool {
	if root == "" {
		return false
	}
	argv := bwrapArgv(root, []string{"true"})
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	return exec.CommandContext(ctx, argv[0], argv[1:]...).Run() == nil
}

func wrapCommand(isolation, root string, tokens []string) []string {
	if isolation != IsolationBwrap || len(tokens) == 0 {
		return tokens
	}
	return bwrapArgv(root, tokens)
}
