---
id: collect-250926-servers-hardware/servers-hardware/zen-3
title: "Zen"
domain: servers-hardware
role: reference
task: reference
actors: ["Alibaba", "Anthropic", "DeepSeek", "Google", "Meta", "MiniMax", "Moonshot", "OpenAI", "Z.ai", "xAI"]
dates: []
keywords: ["astra", "claude", "deepseek", "fable 5", "gemini", "gemini 3.8", "glm", "grok", "grok 4", "kimi", "luna", "muse"]
source: docs/RAG/clean4/zen.md
source_anchor: ""
source_lines: [110, 227]
sha256: 55382896a72d7448e0427ec513bf9c99bf2d6584ac60d8da838597bfe3cc193e
---

# Zen

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

