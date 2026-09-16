// Package keysetup — gestion des clés API des providers pour CETAS Lite,
// adaptée du setup.py du CETAS original.
//
// Principes :
//   - les clés sont saisies en aveugle (jamais d'écho terminal) ;
//   - chaque clé est VALIDÉE par un appel réel au provider avant stockage ;
//   - le coffre (internal/securevault, format V4) est la seule source de vérité ;
//   - le fichier .env exporté ne contient JAMAIS de clé en clair : chaque entrée
//     est scellée (AES-256-GCM) avec une proxy_key conservée dans le coffre.
//     Format byte-identique au setup.py d'origine :
//     <nom>_key=<iv hex>:<ciphertext hex>   (AAD = <nom>)
package keysetup

// TestMode — protocole utilisé pour valider une clé.
type TestMode int

const (
	// ModeChat — POST {base}/chat/completions façon OpenAI.
	ModeChat TestMode = iota
	// ModeResponses — POST {base}/responses (Groq).
	ModeResponses
	// ModeAnthropic — POST /v1/messages (API Anthropic native).
	ModeAnthropic
)

// Provider décrit un fournisseur de modèles configurable.
type Provider struct {
	// ID canonique : nom dans le coffre (api_keys.<ID>) et base du nom .env (<ID>_key).
	ID string
	// Label affiché dans le menu.
	Label string
	// Hint affiché à la saisie (préfixe typique de la clé).
	Hint string
	// TestURL vide => provider désactivé (ex. FreeLLMAPI sans CETAS_FREELLM_URL).
	TestURL string
	// TestModel utilisé pour l'appel de validation.
	TestModel string
	// Mode de validation.
	Mode TestMode
	// Headers supplémentaires pour la validation (ex. OpenRouter).
	Headers map[string]string
	// AppID : ID du provider dans CETAS Lite (sync). Vide = non géré par l'app.
	AppID string
}

// Providers — catalogue repris du setup.py d'origine.
var Providers = []Provider{
	{ID: "nvidia", Label: "NVIDIA NIM", Hint: "nvapi-...",
		TestURL:   "https://integrate.api.nvidia.com/v1/chat/completions",
		TestModel: "nvidia/nemotron-3.5-lightning-30b-a3b", Mode: ModeChat},
	{ID: "groq", Label: "Groq", Hint: "gsk_...",
		TestURL:   "https://api.groq.com/openai/v1/responses",
		TestModel: "openai/gpt-oss-20b", Mode: ModeResponses},
	{ID: "openrouter", Label: "OpenRouter", Hint: "sk-or-...",
		TestURL:   "https://openrouter.ai/api/v1/chat/completions",
		TestModel: "openai/gpt-4o-mini", Mode: ModeChat,
		Headers: map[string]string{
			"HTTP-Referer": "https://cetas.local",
			"X-Title":      "cetas-lite",
		},
		AppID: "openrouter"},
	{ID: "deepseek", Label: "DeepSeek", Hint: "sk-...",
		TestURL:   "https://api.deepseek.com/v1/chat/completions",
		TestModel: "deepseek-chat", Mode: ModeChat,
		AppID: "deepseek"},
	{ID: "freellmapi", Label: "FreeLLMAPI", Hint: "fllm-...",
		TestURL:   freellmURL(),
		TestModel: "llama-3.3-70b-versatile", Mode: ModeChat},
	{ID: "anthropic", Label: "Anthropic", Hint: "sk-ant-...",
		TestURL:   "https://api.anthropic.com/v1/messages",
		TestModel: "claude-3-5-sonnet-20241022", Mode: ModeAnthropic},
	{ID: "openai", Label: "OpenAI", Hint: "sk-...",
		TestURL:   "https://api.openai.com/v1/chat/completions",
		TestModel: "gpt-4o-mini", Mode: ModeChat},
	{ID: "grok", Label: "Grok", Hint: "xai-...",
		TestURL:   "https://api.x.ai/v1/chat/completions",
		TestModel: "grok-beta", Mode: ModeChat},
	{ID: "perplexity", Label: "Perplexity", Hint: "pplx-...",
		TestURL:   "https://api.perplexity.ai/chat/completions",
		TestModel: "llama-3.1-sonar-small-128k-online", Mode: ModeChat},
	{ID: "google", Label: "Google", Hint: "AIza...",
		TestURL:   "https://generativelanguage.googleapis.com/v1beta/openai/chat/completions",
		TestModel: "gemini-3.5-flash-lite", Mode: ModeChat},
	{ID: "mistral", Label: "Mistral", Hint: "sk-...",
		TestURL:   "https://api.mistral.ai/v1/chat/completions",
		TestModel: "open-mistral-nemo", Mode: ModeChat},
	{ID: "qwen", Label: "Qwen", Hint: "sk-...",
		TestURL:   "https://dashscope.aliyuncs.com/compatible-mode/v1/chat/completions",
		TestModel: "qwen-max", Mode: ModeChat},
	{ID: "kimi", Label: "Kimi", Hint: "sk-...",
		TestURL:   "https://api.moonshot.cn/v1/chat/completions",
		TestModel: "moonshot-v1-8k", Mode: ModeChat},
	{ID: "glm", Label: "GLM", Hint: "sk-...",
		TestURL:   "https://open.bigmodel.cn/api/paas/v4/chat/completions",
		TestModel: "glm-4-plus", Mode: ModeChat},
	{ID: "opencode", Label: "OpenCode Zen", Hint: "sk-...",
		TestURL:   "https://opencode.ai/zen/v1/chat/completions",
		TestModel: "glm-5.3", Mode: ModeChat,
		Headers: map[string]string{openCodeSessionHeader: openCodeSessionValue},
		AppID:   "opencode"},
	{ID: "opencode-go", Label: "OpenCode Go", Hint: "sk-...",
		TestURL:   "https://opencode.ai/zen/go/v1/chat/completions",
		TestModel: "hy3", Mode: ModeChat,
		Headers: map[string]string{openCodeSessionHeader: openCodeSessionValue},
		AppID:   "opencode-go"},
}

// openCodeSessionHeader — depuis 2026-09-05, la gateway OpenCode (Zen et Go)
// exige un x-opencode-session (ID stable par conversation, pour le routage
// et la réutilisation du cache de prompt) ; sans lui : HTTP 400
// « Request is missing x-opencode-session ». Pour la validation (test en
// un coup), un UUID fixe suffit.
const openCodeSessionHeader = "x-opencode-session"
const openCodeSessionValue = "6ba7b811-9dad-11d1-80b4-00c04fd430c8"

// freellmURL — FreeLLMAPI n'est testable que si CETAS_FREELLM_URL est défini.
func freellmURL() string {
	return envOr("CETAS_FREELLM_URL", "")
}

// ByID retrouve un provider par son ID canonique.
func ByID(id string) (Provider, bool) {
	for _, p := range Providers {
		if p.ID == id {
			return p, true
		}
	}
	return Provider{}, false
}

// EnvName — nom d'entrée .env pour un ID : <id>_key.
// Byte-identique au setup.py (les IDs sont déjà normalisés ; les cas
// spéciaux OpenCode Zen -> "opencode" et OpenCode Go -> "opencode-go"
// sont portés par les IDs eux-mêmes).
func EnvName(id string) string { return id + "_key" }
