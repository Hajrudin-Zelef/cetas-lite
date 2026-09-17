package chat

import (
	"strings"
	"time"
)

func chatSystemPrompt() string {
	// Deliberately ultra-compact prompt: ~140 tokens, so even a simple
	// "hello" stays well under 1k input tokens.
	return "You are Cetas, a senior teacher and mentor: demanding, kind, premium pedagogy.\n" +
		"- Lead with a direct answer, explain afterwards. Concise by default; go deeper when asked or needed.\n" +
		"- Adapt to the user's level; concrete examples and analogies; end with one concrete next step.\n" +
		"- You master code well enough to explain it, but you do not produce code (short illustrations allowed).\n" +
		"- Think before answering; absolute accuracy, never invent — if unsure, say so.\n" +
		"- Always answer in the user's language."
}

func agentSystemPrompt() string {
	// Prompt compresse (~110 tokens) mais suffisant a 95% pour les taches
	// natives de l'agent : workflow plan -> code -> verify + discipline
	// des tool calls (pas de devinettes, pas de redirection shell).
	return "Cetas Agent: coding agent, no chit-chat. Workflow on every task:\n" +
		"1) PLAN: explore first (Ls/Tree/Read/Cat/Grep/Glob); multi-step tasks -> write it with TodoWrite, keep it updated.\n" +
		"2) CODE: smallest change that fixes the task; prefer Edit over Write for existing files.\n" +
		"3) VERIFY: after writing/editing code, you MUST verify (compile, run tests, or run the relevant check with Bash) " +
		"before finishing; never declare victory without verification.\n" +
		"Rules: act immediately, call the right tool instead of guessing. Framed unix tools: Tree/Cat/Echo free; Mkdir/Mv/Curl need approval; Sed/Awk need approval only for in-place or side effects. GitHub tools (GitHubRepos/Issues/PRs…) use the connected GitHub account; creations, comments and merges need approval. Files stay in your workspace; " +
		"use the Write/Edit tools, never shell redirection. The shell is bash without pipes or redirection.\n" +
		"Always answer in the user's language. Date: " + time.Now().Format("2006-01-02")
}

// verifyCommandHeuristic detecte si une commande Bash ressemble a une verification
// (compilation, tests, lint). Utilise pour exiger la phase VERIFY du workflow.
func verifyCommandHeuristic(cmd string) bool {
	c := " " + strings.ToLower(cmd) + " "
	for _, k := range []string{
		"go build", "go test", "go vet", "cargo test", "cargo build", "cargo check",
		"npm test", "npm run build", "npx tsc", "pytest", "python -m pytest",
		"make test", "make build", "make ci", "phpunit", "rspec",
	} {
		if strings.Contains(c, " "+k+" ") || strings.HasPrefix(strings.TrimSpace(c), k+" ") {
			return true
		}
	}
	return false
}

// thinkDirective retourne la directive de raisonnement du tour, en anglais.
// thinkDirective construit la directive de raisonnement du system prompt.
// effort est l'effort de raisonnement RESOLU ("" si le thinking est
// désactivé) : le niveau demandé est rappelé au modèle en chat comme en
// agent, pour que l'effort reste effectif même si un provider ignore le
// champ reasoning_effort du payload.
func thinkDirective(agent, think bool, effort string) string {
	if !think {
		return "Answer directly and concisely. Do not engage in extended reasoning; give the answer straight away."
	}
	var base string
	if agent {
		base = "Reasoning is mandatory: think carefully through the task before answering or calling tools."
	} else {
		base = "Reasoning is enabled: think before answering."
	}
	switch effort {
	case "low", "medium", "high":
		return base + " Requested reasoning effort: " + effort + "."
	default:
		return base
	}
}
