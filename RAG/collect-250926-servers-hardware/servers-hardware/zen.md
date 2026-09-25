---
id: collect-250926-servers-hardware/servers-hardware/zen
title: "Zen"
domain: servers-hardware
role: reference
task: reference
actors: ["Alibaba", "Anthropic", "DeepSeek", "Google", "Meta", "MiniMax", "Moonshot", "Nvidia", "OpenAI", "OpenRouter", "United States", "Z.ai", "xAI"]
dates: ["2026-02-06", "2026-02-16", "2026-03-06", "2026-03-09", "2026-03-15", "2026-05-14", "2026-06-15", "2026-07-23", "2026-08-05"]
keywords: ["agent", "agents", "astra", "benchmark", "claude", "deepseek", "fable 5", "gemini", "gemini 3.8", "glm", "gpt-5.6", "gpt-6"]
source: docs/RAG/clean4/zen.md
source_anchor: ""
source_lines: [1, 307]
sha256: 2146a9859dd93cb38c0a6225f9a17ce8494757c83beee8fc74d010cd63d7452d
---

# Zen

Liste organisÃ©e de modÃ¨les fournis par OpenCode.

OpenCode Zen est une liste de modÃ¨les testÃ©s et vÃ©rifiÃ©s fournie par lâÃ©quipe OpenCode.

Zen fonctionne comme nâimporte quel autre fournisseur dans OpenCode. Vous vous connectez Ã  OpenCode Zen et obtenez votre clÃ© API. Câest **entiÃ¨rement facultatif** et vous nâavez pas besoin de lâutiliser pour utiliser OpenCode.

Il existe un grand nombre de modÃ¨les, mais seuls quelques-uns fonctionnent bien comme agents de codage. De plus, la plupart des fournisseurs sont configurÃ©s trÃ¨s diffÃ©remment ; vous obtenez donc des performances et une qualitÃ© trÃ¨s variables.

Donc, si vous utilisez un modÃ¨le via quelque chose comme OpenRouter, vous ne pouvez jamais Ãªtre sÃ»r dâobtenir la meilleure version du modÃ¨le que vous voulez.

Pour corriger cela, nous avons fait plusieurs choses :

1. Nous avons testÃ© un groupe sÃ©lectionnÃ© de modÃ¨les et discutÃ© avec leurs Ã©quipes de la meilleure faÃ§on de les faire fonctionner.
2. Nous avons ensuite travaillÃ© avec quelques fournisseurs pour nous assurer quâils Ã©taient correctement servis.
3. Enfin, nous avons comparÃ© la combinaison modÃ¨le/fournisseur et Ã©tabli une liste que nous recommandons en toute confiance.

OpenCode Zen est une passerelle AI qui vous donne accÃ¨s Ã ces modÃ¨les.

OpenCode Zen fonctionne comme nâimporte quel autre fournisseur dans OpenCode.

1. Vous vous connectez Ã  **OpenCode Zen** , ajoutez vos informations de facturation et copiez votre clÃ© API.
2. Vous exÃ©cutez la commande `/connect` dans le TUI, sÃ©lectionnez OpenCode Zen et collez votre clÃ© API.
3. ExÃ©cutez `/models` dans le TUI pour voir la liste des modÃ¨les que nous recommandons.

La facturation se fait Ã la requÃªte et vous pouvez ajouter des crÃ©dits Ã votre compte.

Vous pouvez Ã©galement accÃ©der Ã nos modÃ¨les via les points de terminaison API suivants.

| ModÃ¨le | ID du modÃ¨le | Point de terminaison | Package AI SDK | 
|---|---|---|---|
| GPT 6 Astra | gpt-6-astra | `https://opencode.ai/zen/v1/responses` | `@ai-sdk/openai` | 
| GPT 6 Sol | gpt-6-sol | `https://opencode.ai/zen/v1/responses` | `@ai-sdk/openai` | 
| GPT 6 Luna | gpt-6-luna | `https://opencode.ai/zen/v1/responses` | `@ai-sdk/openai` | 
| GPT 5.6 Sol | gpt-5.6-sol | `https://opencode.ai/zen/v1/responses` | `@ai-sdk/openai` | 
| GPT 5.6 Terra | gpt-5.6-terra | `https://opencode.ai/zen/v1/responses` | `@ai-sdk/openai` | 
| GPT 5.6 Luna | gpt-5.6-luna | `https://opencode.ai/zen/v1/responses` | `@ai-sdk/openai` | 
| GPT 5.5 | gpt-5.5 | `https://opencode.ai/zen/v1/responses` | `@ai-sdk/openai` | 
| GPT 5.5 Pro | gpt-5.5-pro | `https://opencode.ai/zen/v1/responses` | `@ai-sdk/openai` | 
| GPT 5.4 | gpt-5.4 | `https://opencode.ai/zen/v1/responses` | `@ai-sdk/openai` | 
| GPT 5.4 Pro | gpt-5.4-pro | `https://opencode.ai/zen/v1/responses` | `@ai-sdk/openai` | 
| GPT 5.4 Mini | gpt-5.4-mini | `https://opencode.ai/zen/v1/responses` | `@ai-sdk/openai` | 
| GPT 5.4 Nano | gpt-5.4-nano | `https://opencode.ai/zen/v1/responses` | `@ai-sdk/openai` | 
| GPT 5.3 Codex | gpt-5.3-codex | `https://opencode.ai/zen/v1/responses` | `@ai-sdk/openai` | 
| GPT 5.3 Codex Spark | gpt-5.3-codex-spark | `https://opencode.ai/zen/v1/responses` | `@ai-sdk/openai` | 
| GPT 5.2 | gpt-5.2 | `https://opencode.ai/zen/v1/responses` | `@ai-sdk/openai` | 
| GPT 5.2 Codex | gpt-5.2-codex | `https://opencode.ai/zen/v1/responses` | `@ai-sdk/openai` | 
| GPT 5.1 | gpt-5.1 | `https://opencode.ai/zen/v1/responses` | `@ai-sdk/openai` | 
| GPT 5.1 Codex | gpt-5.1-codex | `https://opencode.ai/zen/v1/responses` | `@ai-sdk/openai` | 
| GPT 5.1 Codex Max | gpt-5.1-codex-max | `https://opencode.ai/zen/v1/responses` | `@ai-sdk/openai` | 
| GPT 5.1 Codex Mini | gpt-5.1-codex-mini | `https://opencode.ai/zen/v1/responses` | `@ai-sdk/openai` | 
| GPT 5 | gpt-5 | `https://opencode.ai/zen/v1/responses` | `@ai-sdk/openai` | 
| GPT 5 Codex | gpt-5-codex | `https://opencode.ai/zen/v1/responses` | `@ai-sdk/openai` | 
| GPT 5 Nano | gpt-5-nano | `https://opencode.ai/zen/v1/responses` | `@ai-sdk/openai` | 
| Claude Fable 5.1 | claude-fable-5-1 | `https://opencode.ai/zen/v1/messages` | `@ai-sdk/anthropic` | 
| Claude Fable 5 | claude-fable-5 | `https://opencode.ai/zen/v1/messages` | `@ai-sdk/anthropic` | 
| Claude Opus 5.5 | claude-opus-5-5 | `https://opencode.ai/zen/v1/messages` | `@ai-sdk/anthropic` | 
| Claude Opus 5 | claude-opus-5 | `https://opencode.ai/zen/v1/messages` | `@ai-sdk/anthropic` | 
| Claude Opus 4.8 | claude-opus-4-8 | `https://opencode.ai/zen/v1/messages` | `@ai-sdk/anthropic` | 
| Claude Opus 4.7 | claude-opus-4-7 | `https://opencode.ai/zen/v1/messages` | `@ai-sdk/anthropic` | 
| Claude Opus 4.6 | claude-opus-4-6 | `https://opencode.ai/zen/v1/messages` | `@ai-sdk/anthropic` | 
| Claude Opus 4.5 | claude-opus-4-5 | `https://opencode.ai/zen/v1/messages` | `@ai-sdk/anthropic` | 
| Claude Sonnet 5 | claude-sonnet-5 | `https://opencode.ai/zen/v1/messages` | `@ai-sdk/anthropic` | 
| Claude Sonnet 4.6 | claude-sonnet-4-6 | `https://opencode.ai/zen/v1/messages` | `@ai-sdk/anthropic` | 
| Claude Sonnet 4.5 | claude-sonnet-4-5 | `https://opencode.ai/zen/v1/messages` | `@ai-sdk/anthropic` | 
| Claude Haiku 4.5 | claude-haiku-4-5 | `https://opencode.ai/zen/v1/messages` | `@ai-sdk/anthropic` | 
| Gemini 3.8 Flash | gemini-3.8-flash | `https://opencode.ai/zen/v1/models/gemini-3.8-flash` | `@ai-sdk/google` | 
| Gemini 3.7 Flash | gemini-3.7-flash | `https://opencode.ai/zen/v1/models/gemini-3.7-flash` | `@ai-sdk/google` | 
| Gemini 3.6 Flash | gemini-3.6-flash | `https://opencode.ai/zen/v1/models/gemini-3.6-flash` | `@ai-sdk/google` | 
| Gemini 3.5 Flash | gemini-3.5-flash | `https://opencode.ai/zen/v1/models/gemini-3.5-flash` | `@ai-sdk/google` | 
| Gemini 3.5 Flash Lite | gemini-3.5-flash-lite | `https://opencode.ai/zen/v1/models/gemini-3.5-flash-lite` | `@ai-sdk/google` | 
| Gemini 3.1 Pro | gemini-3.1-pro | `https://opencode.ai/zen/v1/models/gemini-3.1-pro` | `@ai-sdk/google` | 
| Gemini 3 Flash | gemini-3-flash | `https://opencode.ai/zen/v1/models/gemini-3-flash` | `@ai-sdk/google` | 
| Grok 4.7 | grok-4.7 | `https://opencode.ai/zen/v1/responses` | `@ai-sdk/openai` | 
| Grok 4.6 | grok-4.6 | `https://opencode.ai/zen/v1/responses` | `@ai-sdk/openai` | 
| Grok 4.5 | grok-4.5 | `https://opencode.ai/zen/v1/responses` | `@ai-sdk/openai` | 
| Grok Build 0.1 | grok-build-0.1 | `https://opencode.ai/zen/v1/responses` | `@ai-sdk/openai` | 
| Muse Spark 1.3 | muse-spark-1.3 | `https://opencode.ai/zen/v1/responses` | `@ai-sdk/openai` | 
| Muse Spark 1.2 | muse-spark-1.2 | `https://opencode.ai/zen/v1/responses` | `@ai-sdk/openai` | 
| Qwen3.8 Flash | qwen3.8-flash | `https://opencode.ai/zen/v1/messages` | `@ai-sdk/anthropic` | 
| Qwen3.7 Max | qwen3.7-max | `https://opencode.ai/zen/v1/messages` | `@ai-sdk/anthropic` | 
| Qwen3.7 Plus | qwen3.7-plus | `https://opencode.ai/zen/v1/messages` | `@ai-sdk/anthropic` | 
| Qwen3.6 Plus | qwen3.6-plus | `https://opencode.ai/zen/v1/messages` | `@ai-sdk/anthropic` | 
| Qwen3.5 Plus | qwen3.5-plus | `https://opencode.ai/zen/v1/messages` | `@ai-sdk/anthropic` | 
| DeepSeek V4.1 Flash | deepseek-v4.1-flash | `https://opencode.ai/zen/v1/chat/completions` | `@ai-sdk/openai-compatible` | 
| DeepSeek V4 Pro | deepseek-v4-pro | `https://opencode.ai/zen/v1/chat/completions` | `@ai-sdk/openai-compatible` | 
| DeepSeek V4 Flash | deepseek-v4-flash | `https://opencode.ai/zen/v1/chat/completions` | `@ai-sdk/openai-compatible` | 
| DeepSeek V4 Flash Vision Exp | deepseek-v4-flash-vision-exp | `https://opencode.ai/zen/v1/chat/completions` | `@ai-sdk/openai-compatible` | 
| MiniMax M3 | minimax-m3 | `https://opencode.ai/zen/v1/chat/completions` | `@ai-sdk/openai-compatible` | 
| MiniMax M2.7 | minimax-m2.7 | `https://opencode.ai/zen/v1/chat/completions` | `@ai-sdk/openai-compatible` | 
| MiniMax M2.5 | minimax-m2.5 | `https://opencode.ai/zen/v1/chat/completions` | `@ai-sdk/openai-compatible` | 
| GLM 5.3 Flash | glm-5.3-flash | `https://opencode.ai/zen/v1/chat/completions` | `@ai-sdk/openai-compatible` | 
| GLM 5.3 | glm-5.3 | `https://opencode.ai/zen/v1/chat/completions` | `@ai-sdk/openai-compatible` | 
| GLM 5.2 | glm-5.2 | `https://opencode.ai/zen/v1/chat/completions` | `@ai-sdk/openai-compatible` | 
| GLM 5.1 | glm-5.1 | `https://opencode.ai/zen/v1/chat/completions` | `@ai-sdk/openai-compatible` | 
| GLM 5 | glm-5 | `https://opencode.ai/zen/v1/chat/completions` | `@ai-sdk/openai-compatible` | 
| Kimi K2.5 | kimi-k2.5 | `https://opencode.ai/zen/v1/chat/completions` | `@ai-sdk/openai-compatible` | 
| Kimi K2.6 | kimi-k2.6 | `https://opencode.ai/zen/v1/chat/completions` | `@ai-sdk/openai-compatible` | 
| Kimi K2.7 Code | kimi-k2.7-code | `https://opencode.ai/zen/v1/chat/completions` | `@ai-sdk/openai-compatible` | 
| Kimi K3 | kimi-k3 | `https://opencode.ai/zen/v1/chat/completions` | `@ai-sdk/openai-compatible` | 
| Jev 1.13 | jev-1.13 | `https://opencode.ai/zen/v1/systemone` | - | 
| Jev 1.13 Free | jev-1.13-free | `https://opencode.ai/zen/v1/systemone` | - | 
| Big Pickle | big-pickle | `https://opencode.ai/zen/v1/chat/completions` | `@ai-sdk/openai-compatible` | 
| Space Bunny Free | space-bunny-free | `https://opencode.ai/zen/v1/chat/completions` | `@ai-sdk/openai-compatible` | 
| MiMo-V2.6-Flash Free | mimo-v2.6-flash-free | `https://opencode.ai/zen/v1/chat/completions` | `@ai-sdk/openai-compatible` | 
| MiMo-V2.5 Free | mimo-v2.5-free | `https://opencode.ai/zen/v1/chat/completions` | `@ai-sdk/openai-compatible` | 
| Ling 3.0 Flash Fin Free | ling-3.0-flash-fin-free | `https://opencode.ai/zen/v1/chat/completions` | `@ai-sdk/openai-compatible` | 
| Nemotron 3 Ultra Free | nemotron-3-ultra-free | `https://opencode.ai/zen/v1/chat/completions` | `@ai-sdk/openai-compatible` | 
| Nemotron 3.5 Lightning Free | nemotron-3.5-lightning-free | `https://opencode.ai/zen/v1/chat/completions` | `@ai-sdk/openai-compatible` | 
| Muse Spark 1.3 Contributor Free | muse-spark-1.3-contributor-free | `https://opencode.ai/zen/v1/responses` | `@ai-sdk/openai` | 

Le model id dans votre configuration OpenCode utilise le format `opencode/<model-id>`. Par exemple, pour GPT 5.5, vous utiliseriez `opencode/gpt-5.5` dans votre configuration.

Vous pouvez rÃ©cupÃ©rer la liste complÃ¨te des modÃ¨les disponibles et leurs mÃ©tadonnÃ©es Ã partir de :

Jev est un modÃ¨le System One de TypeSafe AI conÃ§u pour prendre rapidement des dÃ©cisions structurÃ©es. Au lieu de gÃ©nÃ©rer du texte, il Ã©value un `state` Ã  partir de questions typÃ©es et renvoie des valeurs et des probabilitÃ©s directement exploitables par votre code. Il prend en charge les questions oui/non (`noul`), Ã  choix multiples (`choice`) et fondÃ©es sur une grille dâÃ©valuation (`score`).

Utilisez votre clÃ© API OpenCode Zen avec lâendpoint `https://opencode.ai/zen/v1/systemone`. Cet exemple vÃ©rifie si une demande dâassistance est urgente :

Vous pouvez poser plusieurs questions dans une mÃªme requÃªte. Jev les Ã©value en parallÃ¨le et renvoie chaque rÃ©ponse sous lâID de la question correspondante :

Utilisez `jev-1.13-free` au lieu de `jev-1.13` pour utiliser le modÃ¨le gratuit proposÃ© pour une durÃ©e limitÃ©e. Consultez la documentation TypeSafe AI pour en savoir plus sur les types de questions et les champs de rÃ©ponse.

Nous prenons en charge un modÃ¨le de paiement Ã  lâutilisation. Vous trouverez ci-dessous les prix **par 1M tokens**.

| ModÃ¨le | Input | Output | Cached Read | Cached Write | 
|---|---|---|---|---|
| Big Pickle | Free | Free | Free | - | 
| Space Bunny Free | Free | Free | Free | - | 
| MiMo-V2.6-Flash Free | Free | Free | Free | - | 
| MiMo-V2.5 Free | Free | Free | Free | - | 
| Ling 3.0 Flash Fin Free | Free | Free | Free | - | 
| Nemotron 3 Ultra Free | Free | Free | Free | - | 
| Nemotron 3.5 Lightning Free | Free | Free | Free | - | 
| Muse Spark 1.3 Contributor Free | Free | Free | Free | - | 
| Jev 1.13 Free | Gratuit | Gratuit | - | - | 
| Jev 1.13 | $0.042 | Gratuit | - | - | 
| MiniMax M3 | $0.30 | $1.20 | $0.06 | - | 
| MiniMax M2.7 | $0.30 | $1.20 | $0.06 | - | 
| MiniMax M2.5 | $0.30 | $1.20 | $0.06 | - | 
| GLM 5.3 Flash | $0.15 | $0.50 | $0.03 | - | 
| GLM 5.3 | $1.40 | $4.40 | $0.26 | - | 
| GLM 5.2 | $1.40 | $4.40 | $0.26 | - | 
| GLM 5.1 | $1.40 | $4.40 | $0.26 | - | 
| GLM 5 | $1.00 | $3.20 | $0.20 | - | 
| Kimi K2.7 Code | $0.95 | $4.00 | $0.19 | - | 
| Kimi K3 | $3.00 | $15.00 | $0.30 | - | 
| Kimi K2.6 | $0.95 | $4.00 | $0.16 | - | 
| Kimi K2.5 | $0.60 | $3.00 | $0.10 | - | 
| Qwen3.8 Flash | $0.15 | $0.47 | $0.016 | $0.20 | 
| Qwen3.7 Max | $2.50 | $7.50 | $0.50 | $3.125 | 
| Qwen3.7 Plus | $0.40 | $1.60 | $0.04 | $0.50 | 
| Qwen3.6 Plus | $0.50 | $3.00 | $0.05 | $0.625 | 
| Qwen3.5 Plus | $0.20 | $1.20 | $0.02 | $0.25 | 
| DeepSeek V4.1 Flash | $0.30 | $1.20 | $0.006 | - | 
| DeepSeek V4 Pro | $1.74 | $3.48 | $0.145 | - | 
| DeepSeek V4 Flash | $0.14 | $0.28 | $0.028 | - | 
| DeepSeek V4 Flash Vision Exp | $0.14 | $0.28 | $0.028 | - | 
| Claude Fable 5.1 | $10.00 | $50.00 | $0.25 | $12.50 | 
| Claude Fable 5 | $10.00 | $50.00 | $1.00 | $12.50 | 
| Claude Opus 5.5 | $4.00 | $20.00 | $0.20 | $5.00 | 
| Claude Opus 5 | $5.00 | $25.00 | $0.50 | $6.25 | 
| Claude Opus 4.8 | $5.00 | $25.00 | $0.50 | $6.25 | 
| Claude Opus 4.7 | $5.00 | $25.00 | $0.50 | $6.25 | 
| Claude Opus 4.6 | $5.00 | $25.00 | $0.50 | $6.25 | 
| Claude Opus 4.5 | $5.00 | $25.00 | $0.50 | $6.25 | 
| Claude Sonnet 5 | $2.00 | $10.00 | $0.20 | $2.50 | 
| Claude Sonnet 4.6 | $3.00 | $15.00 | $0.30 | $3.75 | 
| Claude Sonnet 4.5 (â¤ 200K tokens) | $3.00 | $15.00 | $0.30 | $3.75 | 
| Claude Sonnet 4.5 (> 200K tokens) | $6.00 | $22.50 | $0.60 | $7.50 | 
| Claude Haiku 4.5 | $1.00 | $5.00 | $0.10 | $1.25 | 
| Gemini 3.8 Flash | $1.50 | $7.50 | $0.15 | - | 
| Gemini 3.7 Flash | $1.50 | $7.50 | $0.15 | - | 
| Gemini 3.6 Flash | $1.50 | $7.50 | $0.15 | - | 
| Gemini 3.5 Flash | $1.50 | $9.00 | $0.15 | - | 
| Gemini 3.5 Flash Lite | $0.30 | $2.50 | $0.03 | - | 
| Gemini 3.1 Pro (â¤ 200K tokens) | $2.00 | $12.00 | $0.20 | - | 
| Gemini 3.1 Pro (> 200K tokens) | $4.00 | $18.00 | $0.40 | - | 
| Gemini 3 Flash | $0.50 | $3.00 | $0.05 | - | 
| Grok 4.7 (â¤ 200K tokens) | $2.00 | $6.00 | $0.50 | - | 
| Grok 4.7 (> 200K tokens) | $4.00 | $12.00 | $1.00 | - | 
| Grok 4.6 (â¤ 200K tokens) | $2.00 | $6.00 | $0.50 | - | 
| Grok 4.6 (> 200K tokens) | $4.00 | $12.00 | $1.00 | - | 
| Grok 4.5 (â¤ 200K tokens) | $2.00 | $6.00 | $0.30 | - | 
| Grok 4.5 (> 200K tokens) | $4.00 | $12.00 | $0.60 | - | 
| Grok Build 0.1 | $1.00 | $2.00 | $0.20 | - | 
| Muse Spark 1.3 | $1.25 | $4.25 | $0.15 | - | 
| Muse Spark 1.2 | $1.25 | $4.25 | $0.15 | - | 
| GPT 6 Astra (â¤ 272K tokens) | $10.00 | $50.00 | $1.00 | $12.50 | 
| GPT 6 Astra (> 272K tokens) | $20.00 | $75.00 | $2.00 | $25.00 | 
| GPT 6 Sol (â¤ 272K tokens) | $2.00 | $10.00 | $0.20 | $2.50 | 
| GPT 6 Sol (> 272K tokens) | $4.00 | $15.00 | $0.40 | $5.00 | 
| GPT 6 Luna (â¤ 272K tokens) | $0.10 | $0.50 | $0.01 | $0.125 | 
| GPT 6 Luna (> 272K tokens) | $0.20 | $0.75 | $0.02 | $0.25 | 
| GPT 5.6 Sol (â¤ 272K tokens) | $4.00 | $20.00 | $0.40 | $5.00 | 
| GPT 5.6 Sol (> 272K tokens) | $8.00 | $30.00 | $0.80 | $10.00 | 
| GPT 5.6 Terra (â¤ 272K tokens) | $2.00 | $12.00 | $0.20 | $2.50 | 
| GPT 5.6 Terra (> 272K tokens) | $4.00 | $18.00 | $0.40 | $5.00 | 
| GPT 5.6 Luna (â¤ 272K tokens) | $0.20 | $1.20 | $0.02 | $0.25 | 
| GPT 5.6 Luna (> 272K tokens) | $0.40 | $1.80 | $0.04 | $0.50 | 
| GPT 5.5 (â¤ 272K tokens) | $5.00 | $30.00 | $0.50 | - | 
| GPT 5.5 (> 272K tokens) | $10.00 | $45.00 | $1.00 | - | 
| GPT 5.5 Pro | $30.00 | $180.00 | $30.00 | - | 
| GPT 5.4 (â¤ 272K tokens) | $2.50 | $15.00 | $0.25 | - | 
| GPT 5.4 (> 272K tokens) | $5.00 | $22.50 | $0.50 | - | 
| GPT 5.4 Pro | $30.00 | $180.00 | $30.00 | - | 
| GPT 5.4 Mini | $0.75 | $4.50 | $0.075 | - | 
| GPT 5.4 Nano | $0.20 | $1.25 | $0.02 | - | 
| GPT 5.3 Codex Spark | $1.75 | $14.00 | $0.175 | - | 
| GPT 5.3 Codex | $1.75 | $14.00 | $0.175 | - | 
| GPT 5.2 | $1.75 | $14.00 | $0.175 | - | 
| GPT 5.2 Codex | $1.75 | $14.00 | $0.175 | - | 
| GPT 5.1 | $1.07 | $8.50 | $0.107 | - | 
| GPT 5.1 Codex | $1.07 | $8.50 | $0.107 | - | 
| GPT 5.1 Codex Max | $1.25 | $10.00 | $0.125 | - | 
| GPT 5.1 Codex Mini | $0.25 | $2.00 | $0.025 | - | 
| GPT 5 | $1.07 | $8.50 | $0.107 | - | 
| GPT 5 Codex | $1.07 | $8.50 | $0.107 | - | 
| GPT 5 Nano | $0.05 | $0.40 | $0.005 | - | 

**DeepSeek V4 Flash Vision Exp:** Les images sont converties en tokens selon leurs dimensions et facturÃ©es comme tokens dâentrÃ©e avec les tokens de texte. En savoir plus.

Vous remarquerez peut-Ãªtre des modÃ¨les Ã faible coÃ»t, tels que Haiku, Nano ou Flash, dans votre historique dâutilisation. OpenCode utilise ces modÃ¨les pour gÃ©nÃ©rer les titres des sessions.

Les modÃ¨les gratuits :

- MiMo-V2.6-Flash Free est disponible sur OpenCode pour une durÃ©e limitÃ©e. LâÃ©quipe utilise cette pÃ©riode pour recueillir des retours et amÃ©liorer le modÃ¨le.
- MiMo-V2.5 Free est disponible sur OpenCode pour une durÃ©e limitÃ©e. LâÃ©quipe utilise cette pÃ©riode pour recueillir des retours et amÃ©liorer le modÃ¨le.
- Ling 3.0 Flash Fin Free est disponible sur OpenCode pour une durÃ©e limitÃ©e. LâÃ©quipe utilise cette pÃ©riode pour recueillir des retours et amÃ©liorer le modÃ¨le.
- Nemotron 3 Ultra Free est disponible sur OpenCode pour une durÃ©e limitÃ©e. LâÃ©quipe utilise cette pÃ©riode pour recueillir des retours et amÃ©liorer le modÃ¨le.
- Nemotron 3.5 Lightning Free est disponible sur OpenCode pour une durÃ©e limitÃ©e. LâÃ©quipe utilise cette pÃ©riode pour recueillir des retours et amÃ©liorer le modÃ¨le.
- Big Pickle est un modÃ¨le stealth gratuit sur OpenCode pour une durÃ©e limitÃ©e. LâÃ©quipe utilise cette pÃ©riode pour recueillir des retours et amÃ©liorer le modÃ¨le.
- Space Bunny Free est un modÃ¨le stealth gratuit sur OpenCode pour une durÃ©e limitÃ©e. Son fournisseur applique une politique de conservation nulle et nâutilise pas vos donnÃ©es pour entraÃ®ner des modÃ¨les.
- Muse Spark 1.3 Contributor Free est disponible sur OpenCode pour une durÃ©e limitÃ©e. LâÃ©quipe utilise cette pÃ©riode pour recueillir des retours et amÃ©liorer le modÃ¨le.
- Jev 1.13 Free est disponible sur OpenCode pour une durÃ©e limitÃ©e.

Contactez-nous si vous avez des questions.

Si votre solde passe sous $5, Zen rechargera automatiquement $20.

Vous pouvez modifier le montant du rechargement automatique. Vous pouvez Ã©galement dÃ©sactiver complÃ¨tement le rechargement automatique.

Vous pouvez Ã©galement dÃ©finir une limite dâutilisation mensuelle pour lâensemble de lâespace de travail et pour chaque membre de votre Ã©quipe.

Par exemple, si vous dÃ©finissez une limite dâutilisation mensuelle Ã $20, Zen nâutilisera pas plus de $20 sur un mois. Mais si le rechargement automatique est activÃ©, Zen peut finir par vous facturer plus de $20 si votre solde passe sous $5.

| ModÃ¨le | Date de dÃ©prÃ©ciation | 
|---|---|
| GPT 5.2 Codex | July 23, 2026 | 
| GPT 5.1 Codex | July 23, 2026 | 
| GPT 5.1 Codex Max | July 23, 2026 | 
| GPT 5.1 Codex Mini | July 23, 2026 | 
| GPT 5 Codex | July 23, 2026 | 
| Claude Opus 4.1 | August 5, 2026 | 
| Claude Sonnet 4 | June 15, 2026 | 
| Claude Haiku 3.5 | February 16, 2026 | 
| Gemini 3 Pro | March 9, 2026 | 
| MiniMax M2.5 | August 5, 2026 | 
| MiniMax M2.1 | March 15, 2026 | 
| GLM 5 | May 14, 2026 | 
| GLM 4.7 | March 15, 2026 | 
| GLM 4.6 | March 15, 2026 | 
| Kimi K2.5 | August 5, 2026 | 
| Kimi K2 Thinking | March 6, 2026 | 
| Kimi K2 | March 6, 2026 | 
| Qwen3 Coder 480B | February 6, 2026 | 

Tous nos modÃ¨les sont hÃ©bergÃ©s aux US. Nos fournisseurs suivent une politique de rÃ©tention zÃ©ro et nâutilisent pas vos donnÃ©es pour lâentraÃ®nement des modÃ¨les, avec les exceptions suivantes :

- Big Pickle : Pendant sa pÃ©riode gratuite, les donnÃ©es collectÃ©es peuvent Ãªtre utilisÃ©es pour amÃ©liorer le modÃ¨le.
- MiMo-V2.6-Flash Free : Pendant sa pÃ©riode gratuite, les donnÃ©es collectÃ©es peuvent Ãªtre utilisÃ©es pour amÃ©liorer le modÃ¨le.
- MiMo-V2.5 Free : Pendant sa pÃ©riode gratuite, les donnÃ©es collectÃ©es peuvent Ãªtre utilisÃ©es pour amÃ©liorer le modÃ¨le.
- Ling 3.0 Flash Fin Free : Pendant sa pÃ©riode gratuite, les donnÃ©es collectÃ©es peuvent Ãªtre utilisÃ©es pour amÃ©liorer le modÃ¨le.
- Nemotron 3 Ultra Free (endpoints NVIDIA gratuits) : RÃ©servÃ© Ã un usage dâessai â nâenvoyez pas de donnÃ©es personnelles ou confidentielles. Votre utilisation est journalisÃ©e Ã des fins de sÃ©curitÃ© et pour amÃ©liorer les produits et services de NVIDIA. Les donnÃ©es de session journalisÃ©es Ã des fins dâamÃ©lioration ne sont pas liÃ©es Ã votre identitÃ© ni Ã un quelconque identifiant persistant. Pour plus dâinformations sur nos pratiques de traitement des donnÃ©es, consultez notre Politique de confidentialitÃ©. En interagissant avec cet endpoint, vous consentez Ã notre collecte, Ã notre enregistrement et Ã notre utilisation de ces informations ainsi quâaux NVIDIA API Trial Terms of Service.
- Nemotron 3.5 Lightning Free (endpoints NVIDIA gratuits) : RÃ©servÃ© Ã un usage dâessai â nâenvoyez pas de donnÃ©es personnelles ou confidentielles. Votre utilisation est journalisÃ©e Ã des fins de sÃ©curitÃ© et pour amÃ©liorer les produits et services de NVIDIA. Les donnÃ©es de session journalisÃ©es Ã des fins dâamÃ©lioration ne sont pas liÃ©es Ã votre identitÃ© ni Ã un quelconque identifiant persistant. Pour plus dâinformations sur nos pratiques de traitement des donnÃ©es, consultez notre Politique de confidentialitÃ©. En interagissant avec cet endpoint, vous consentez Ã notre collecte, Ã notre enregistrement et Ã notre utilisation de ces informations ainsi quâaux NVIDIA API Trial Terms of Service.
- OpenAI APIs : Les requÃªtes sont conservÃ©es pendant 30 jours conformÃ©ment Ã OpenAIâs Data Policies.
- Anthropic APIs : Les requÃªtes sont conservÃ©es pendant 30 jours conformÃ©ment Ã Anthropicâs Data Policies.
- Muse Spark 1.3 Contributor Free : Tarification des tokens fortement rÃ©duite en Ã©change de lâautorisation dâutiliser vos prompts et complÃ©tions pour entraÃ®ner les futurs modÃ¨les Meta. En savoir plus.

Zen fonctionne aussi trÃ¨s bien pour les Ã©quipes. Vous pouvez inviter des coÃ©quipiers, attribuer des rÃ´les, sÃ©lectionner les modÃ¨les que votre Ã©quipe utilise, et plus encore.

La gestion de votre espace de travail est actuellement gratuite pour les Ã©quipes dans le cadre de la version bÃªta. Nous partagerons bientÃ´t plus de dÃ©tails sur la tarification.

Vous pouvez inviter des coÃ©quipiers dans votre espace de travail et attribuer des rÃ´les :

- **Admin** : GÃ©rer les modÃ¨les, les membres, les clÃ©s API et la facturation
- **Member** : GÃ©rer uniquement ses propres clÃ©s API

Les administrateurs peuvent Ã©galement dÃ©finir des limites de dÃ©penses mensuelles pour chaque membre afin de garder les coÃ»ts sous contrÃ´le.

Les administrateurs peuvent activer ou dÃ©sactiver des modÃ¨les spÃ©cifiques pour lâespace de travail. Les requÃªtes effectuÃ©es vers un modÃ¨le dÃ©sactivÃ© renverront une erreur.

Cela est utile si vous souhaitez dÃ©sactiver lâutilisation dâun modÃ¨le qui collecte des donnÃ©es.

Vous pouvez utiliser vos propres clÃ©s API OpenAI ou Anthropic tout en accÃ©dant Ã dâautres modÃ¨les dans Zen.

Lorsque vous utilisez vos propres clÃ©s, les tokens sont facturÃ©s directement par le fournisseur, pas par Zen.

Par exemple, votre organisation a peut-Ãªtre dÃ©jÃ une clÃ© pour OpenAI ou Anthropic et vous souhaitez lâutiliser Ã la place de celle fournie par Zen.

Nous avons crÃ©Ã© OpenCode Zen pour :

1. **Benchmark** les meilleurs modÃ¨les/fournisseurs pour les agents de codage.
2. Avoir accÃ¨s aux options de **la plus haute qualitÃ©** sans dÃ©grader les performances ni basculer vers des fournisseurs moins chers.
3. RÃ©percuter toute **baisse de prix** en vendant au prix coÃ»tant ; la seule marge sert Ã  couvrir nos frais de traitement.
4. Nâavoir **aucun lock-in** en vous permettant de lâutiliser avec nâimporte quel autre agent de codage. Et toujours vous permettre dâutiliser nâimporte quel autre fournisseur avec OpenCode Ã©galement.
