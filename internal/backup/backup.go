package backup

import (
	"archive/tar"
	"compress/gzip"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"time"
)

const (
	DBName  = "cetas-lite.db"
	MCPName = "mcp.json"
)

func Create(dbPath, mcpPath, out string) error {
	if _, err := os.Stat(dbPath); err != nil {
		return fmt.Errorf("base introuvable: %w", err)
	}
	tmp := out + ".tmp"
	f, err := os.Create(tmp)
	if err != nil {
		return err
	}
	gz := gzip.NewWriter(f)
	tw := tar.NewWriter(gz)

	fail := func(err error) error {
		_ = tw.Close()
		_ = gz.Close()
		_ = f.Close()
		_ = os.Remove(tmp)
		return err
	}

	for _, entry := range []struct{ path, name string }{
		{dbPath, DBName},
		{mcpPath, MCPName},
	} {
		st, err := os.Stat(entry.path)
		if err != nil {
			continue
		}
		if err := tw.WriteHeader(&tar.Header{
			Name: entry.name, Mode: 0o600, Size: st.Size(), ModTime: time.Now(),
		}); err != nil {
			return fail(err)
		}
		in, err := os.Open(entry.path)
		if err != nil {
			return fail(err)
		}
		_, err = io.Copy(tw, in)
		_ = in.Close()
		if err != nil {
			return fail(err)
		}
	}

	if err := tw.Close(); err != nil {
		return fail(err)
	}
	if err := gz.Close(); err != nil {
		return fail(err)
	}
	if err := f.Close(); err != nil {
		_ = os.Remove(tmp)
		return err
	}
	return os.Rename(tmp, out)
}

func Restore(bundle, home string) error {
	f, err := os.Open(bundle)
	if err != nil {
		return err
	}
	defer f.Close()
	gz, err := gzip.NewReader(f)
	if err != nil {
		return fmt.Errorf("bundle illisible: %w", err)
	}
	defer gz.Close()
	tr := tar.NewReader(gz)

	seenDB := false
	for {
		hdr, err := tr.Next()
		if errors.Is(err, io.EOF) {
			break
		}
		if err != nil {
			return err
		}
		name := filepath.Base(hdr.Name)
		if name != DBName && name != MCPName {
			continue
		}
		dst := filepath.Join(home, name)
		if name == DBName {
			if _, err := os.Stat(dst); err == nil {
				_ = copyFile(dst, dst+".bak-"+time.Now().Format("20060102_150405"))
			}
			seenDB = true
		}
		tmp := dst + ".restore-tmp"
		if err := writeFile(tmp, tr); err != nil {
			return err
		}
		if err := os.Rename(tmp, dst); err != nil {
			return err
		}
	}
	if !seenDB {
		return errors.New("bundle sans " + DBName)
	}
	return nil
}

func writeFile(path string, r io.Reader) error {
	out, err := os.OpenFile(path, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0o600)
	if err != nil {
		return err
	}
	if _, err := io.Copy(out, io.LimitReader(r, 1<<31)); err != nil {
		_ = out.Close()
		return err
	}
	return out.Close()
}

func copyFile(src, dst string) error {
	in, err := os.Open(src)
	if err != nil {
		return err
	}
	defer in.Close()
	return writeFile(dst, in)
}
