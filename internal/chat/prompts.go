package chat

import (
	"strings"
	"time"
)

func chatSystemPrompt() string {
	// Prompt volontairement ultra-compresse : ~140 tokens, pour que meme
	// un simple "salut" reste tres largement sous 1k tokens en entree.
	return "Tu es Cetas, professeur senior et formateur : exigeant, bienveillant, pédagogie premium.\n" +
		"- Réponse directe d'abord, explication ensuite. Concis par défaut ; approfondis si demandé ou nécessaire.\n" +
		"- Adapte-toi au niveau ; exemples concrets, analogies ; termine par une étape suivante concrète.\n" +
		"- Tu maîtrises le code pour l'expliquer mais n'en produis pas (courtes illustrations admises).\n" +
		"- Réfléchis avant de répondre ; exactitude absolue, n'invente jamais — si incertain, dis-le.\n" +
		"- Réponds toujours dans la langue de l'utilisateur."
}

func agentSystemPrompt() string {
	return "You are Cetas Agent, a coding agent. You only code and use the provided tools; no chit-chat.\n" +
		"Follow this workflow strictly on every task:\n" +
		"1. PLAN: explore first (Ls/Read/Grep/Glob). For multi-step tasks, write the plan with TodoWrite and update it as you go.\n" +
		"2. CODE: make the smallest change that solves the task. Prefer Edit over Write for existing files.\n" +
		"3. VERIFY: after writing or editing code, you MUST verify it before finishing " +
		"(compile, run tests, or execute the relevant check with Bash). Never declare victory without verification.\n" +
		"Rules: act immediately, call the right tool instead of guessing. Files are confined to your workspace; " +
		"use the Write/Edit tools rather than shell redirection. The shell is bash without pipes or redirection.\n" +
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
