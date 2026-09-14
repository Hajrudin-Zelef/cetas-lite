package chat

import "cetas-lite/internal/provider"

type LogEvent struct {
	Seq   int            `json:"seq"`
	TS    int64          `json:"ts"`
	Delta map[string]any `json:"delta"`
}

type snapshot struct {
	ID       string             `json:"id"`
	Messages []provider.Message `json:"messages"`
	Log      []LogEvent         `json:"log"`
	Seq      int                `json:"seq"`
	Turn     *snapshotTurn      `json:"turn,omitempty"`
}

type snapshotTurn struct {
	Family      string   `json:"family"`
	Mode        string   `json:"mode"`
	Text        string   `json:"text"`
	Web         bool     `json:"web,omitempty"`
	MCP         bool     `json:"mcp,omitempty"`
	Think       bool     `json:"think,omitempty"`
	Effort      string   `json:"effort,omitempty"`
	Approve     bool     `json:"approve,omitempty"`
	Plan        bool     `json:"plan,omitempty"`
	Worktree    bool     `json:"worktree,omitempty"`
	Repo        string   `json:"repo,omitempty"`
	ProjectID   string   `json:"project_id,omitempty"`
	Attachments []string `json:"attachments,omitempty"`
}
