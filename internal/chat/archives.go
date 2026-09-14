package chat

import (
	"encoding/json"
	"sort"
	"strings"
	"time"
)

type ArchiveInfo struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	Updated  int64  `json:"updated"`
	Messages int    `json:"messages"`
}

func (e *Engine) ListArchives(user string) []ArchiveInfo {
	if e.st == nil {
		return nil
	}
	ids, _ := e.st.ListArchives(user)
	out := make([]ArchiveInfo, 0, len(ids))
	for _, id := range ids {
		raw, ok := e.st.GetArchive(user, id)
		if !ok {
			continue
		}
		var s snapshot
		if json.Unmarshal(raw, &s) != nil {
			continue
		}
		out = append(out, archiveInfo(id, s))
	}
	sort.Slice(out, func(i, j int) bool {
		if out[i].Updated != out[j].Updated {
			return out[i].Updated > out[j].Updated
		}
		return out[i].ID > out[j].ID
	})
	return out
}

func (e *Engine) DeleteArchive(user, id string) bool {
	if e.st == nil {
		return false
	}
	if _, ok := e.st.GetArchive(user, id); !ok {
		return false
	}
	_ = e.st.DeleteArchive(user, id)
	return true
}

func archiveInfo(id string, s snapshot) ArchiveInfo {
	title := ""
	for _, m := range s.Messages {
		if m.Role != "user" {
			continue
		}
		if txt, ok := m.Content.(string); ok && strings.TrimSpace(txt) != "" {
			title = condense(txt, 80)
			break
		}
	}
	if title == "" {
		title = "(sans titre)"
	}
	var updated int64
	if len(s.Log) > 0 {
		updated = s.Log[len(s.Log)-1].TS
	}
	if updated == 0 {
		updated = archiveIDTime(id)
	}
	n := 0
	for _, m := range s.Messages {
		if m.Role == "user" || m.Role == "assistant" {
			n++
		}
	}
	return ArchiveInfo{ID: id, Title: title, Updated: updated, Messages: n}
}

func condense(s string, max int) string {
	s = strings.Join(strings.Fields(s), " ")
	r := []rune(s)
	if len(r) > max {
		return strings.TrimSpace(string(r[:max])) + "…"
	}
	return s
}

func archiveIDTime(id string) int64 {
	if len(id) < 15 {
		return 0
	}
	t, err := time.ParseInLocation("20060102_150405", id[:15], time.Local)
	if err != nil {
		return 0
	}
	return t.UnixMilli()
}
