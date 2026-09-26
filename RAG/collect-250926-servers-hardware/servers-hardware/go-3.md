---
id: collect-250926-servers-hardware/servers-hardware/go-3
title: "Go"
domain: servers-hardware
role: reference
task: reference
actors: ["Alibaba", "DeepSeek", "LongCat", "Meta", "MiniMax", "Moonshot", "Z.ai", "xAI"]
dates: []
keywords: ["deepseek", "glm", "gpt-5.6", "gpt-6", "gpu", "grok", "grok 4", "kimi", "luna", "muse", "muse spark"]
source: docs/RAG/clean4/go.md
source_anchor: ""
source_lines: [230, 318]
sha256: 6a67aae4a8341a4cf76f29d22c19e580a10486369bbbbb2ee4085ca7d839aadc
---

# Go

Les limites dâutilisation peuvent changer au fur et Ã mesure que nous tirons des enseignements des premiÃ¨res utilisations et des retours.

Si vous avez Ã©galement des crÃ©dits sur votre solde Zen, vous pouvez activer lâoption **Use balance** dans la console. Lorsquâelle est activÃ©e, Go se rabattra sur votre solde Zen aprÃ¨s que vous ayez atteint vos limites dâutilisation au lieu de bloquer les requÃªtes.

Avec Go, vous payez 10 $/mois, et lâutilisation mensuelle incluse varie selon le modÃ¨le.

Pour la plupart des modÃ¨les, nous y parvenons grÃ¢ce Ã des remises sur volume et Ã une capacitÃ© GPU rÃ©servÃ©e. Nous vous faisons ensuite bÃ©nÃ©ficier de ces Ã©conomies sous la forme dâune utilisation mensuelle plus Ã©levÃ©e.

Pour certains modÃ¨les, nous nâavons pas encore eu lâoccasion de nÃ©gocier une remise ou de les hÃ©berger Ã moindre coÃ»t, soit parce que le modÃ¨le est nouveau, soit parce que son tarif public est dÃ©jÃ rÃ©duit.

Pour ces modÃ¨les, vous obtenez tout de mÃªme un peu plus que si vous payiez directement les fournisseurs de modÃ¨les ; câest pourquoi leur utilisation mensuelle incluse est plus faible.

Vous pouvez Ã©galement accÃ©der aux modÃ¨les Go via les points de terminaison dâAPI suivants.

| ModÃ¨le | ID de modÃ¨le | Point de terminaison | Package AI SDK | 
|---|---|---|---|
| Grok 4.7 | grok-4.7 | `https://opencode.ai/zen/go/v1/responses` | `@ai-sdk/openai` | 
| Grok 4.6 | grok-4.6 | `https://opencode.ai/zen/go/v1/responses` | `@ai-sdk/openai` | 
| GPT 6 Luna | gpt-6-luna | `https://opencode.ai/zen/go/v1/responses` | `@ai-sdk/openai` | 
| GPT 5.6 Luna | gpt-5.6-luna | `https://opencode.ai/zen/go/v1/responses` | `@ai-sdk/openai` | 
| GLM-5.3-Flash | glm-5.3-flash | `https://opencode.ai/zen/go/v1/chat/completions` | `@ai-sdk/openai-compatible` | 
| GLM-5.3 | glm-5.3 | `https://opencode.ai/zen/go/v1/chat/completions` | `@ai-sdk/openai-compatible` | 
| GLM-5.2 | glm-5.2 | `https://opencode.ai/zen/go/v1/chat/completions` | `@ai-sdk/openai-compatible` | 
| GLM-5.1 | glm-5.1 | `https://opencode.ai/zen/go/v1/chat/completions` | `@ai-sdk/openai-compatible` | 
| Kimi K3 | kimi-k3 | `https://opencode.ai/zen/go/v1/chat/completions` | `@ai-sdk/openai-compatible` | 
| Kimi K2.7 Code | kimi-k2.7-code | `https://opencode.ai/zen/go/v1/chat/completions` | `@ai-sdk/openai-compatible` | 
| Kimi K2.6 | kimi-k2.6 | `https://opencode.ai/zen/go/v1/chat/completions` | `@ai-sdk/openai-compatible` | 
| LongCat-2.0 | longcat-2.0 | `https://opencode.ai/zen/go/v1/chat/completions` | `@ai-sdk/openai-compatible` | 
| DeepSeek V4.1 Flash | deepseek-v4.1-flash | `https://opencode.ai/zen/go/v1/chat/completions` | `@ai-sdk/openai-compatible` | 
| DeepSeek V4 Pro | deepseek-v4-pro | `https://opencode.ai/zen/go/v1/chat/completions` | `@ai-sdk/openai-compatible` | 
| DeepSeek V4 Flash | deepseek-v4-flash | `https://opencode.ai/zen/go/v1/chat/completions` | `@ai-sdk/openai-compatible` | 
| DeepSeek V4 Flash Vision Exp | deepseek-v4-flash-vision-exp | `https://opencode.ai/zen/go/v1/chat/completions` | `@ai-sdk/openai-compatible` | 
| MiMo-V2.6-Flash | mimo-v2.6-flash | `https://opencode.ai/zen/go/v1/chat/completions` | `@ai-sdk/openai-compatible` | 
| MiMo-V2.6-Pro | mimo-v2.6-pro | `https://opencode.ai/zen/go/v1/chat/completions` | `@ai-sdk/openai-compatible` | 
| MiMo-V2.5 | mimo-v2.5 | `https://opencode.ai/zen/go/v1/chat/completions` | `@ai-sdk/openai-compatible` | 
| MiMo-V2.5-Pro | mimo-v2.5-pro | `https://opencode.ai/zen/go/v1/chat/completions` | `@ai-sdk/openai-compatible` | 
| MiniMax M3 | minimax-m3 | `https://opencode.ai/zen/go/v1/messages` | `@ai-sdk/anthropic` | 
| MiniMax M2.7 | minimax-m2.7 | `https://opencode.ai/zen/go/v1/messages` | `@ai-sdk/anthropic` | 
| MiniMax M2.5 | minimax-m2.5 | `https://opencode.ai/zen/go/v1/messages` | `@ai-sdk/anthropic` | 
| Muse Spark 1.3 Contributor | muse-spark-1.3-contributor | `https://opencode.ai/zen/go/v1/responses` | `@ai-sdk/openai` | 
| Muse Spark 1.2 Contributor | muse-spark-1.2-contributor | `https://opencode.ai/zen/go/v1/responses` | `@ai-sdk/openai` | 
| Qwen3.8 Max | qwen3.8-max | `https://opencode.ai/zen/go/v1/messages` | `@ai-sdk/anthropic` | 
| Qwen3.8 Flash | qwen3.8-flash | `https://opencode.ai/zen/go/v1/messages` | `@ai-sdk/anthropic` | 
| Qwen3.7 Max | qwen3.7-max | `https://opencode.ai/zen/go/v1/messages` | `@ai-sdk/anthropic` | 
| Qwen3.7 Plus | qwen3.7-plus | `https://opencode.ai/zen/go/v1/messages` | `@ai-sdk/anthropic` | 
| Qwen3.6 Plus | qwen3.6-plus | `https://opencode.ai/zen/go/v1/messages` | `@ai-sdk/anthropic` | 
| Hy4 preview | hy4-preview | `https://opencode.ai/zen/go/v1/chat/completions` | `@ai-sdk/openai-compatible` | 
| Hy3 | hy3 | `https://opencode.ai/zen/go/v1/chat/completions` | `@ai-sdk/openai-compatible` | 
| Space Bunny Free | space-bunny-free | `https://opencode.ai/zen/go/v1/chat/completions` | `@ai-sdk/openai-compatible` | 

LâID de modÃ¨le dans votre configuration OpenCode utilise le format `opencode-go/<model-id>`. Par exemple, pour Kimi K3, vous utiliseriez `opencode-go/kimi-k3` dans votre configuration.

Vous pouvez rÃ©cupÃ©rer la liste complÃ¨te des modÃ¨les disponibles et leurs mÃ©tadonnÃ©es Ã partir de :

| ModÃ¨le | EntraÃ®nement des modÃ¨les | Conservation des donnÃ©es | 
|---|---|---|
| Grok 4.7 | Non utilisÃ© | 30 jours | 
| Grok 4.6 | Non utilisÃ© | 30 jours | 
| GPT 6 Luna | Non utilisÃ© | 30 jours | 
| GPT 5.6 Luna | Non utilisÃ© | 30 jours | 
| GLM-5.3-Flash | Non utilisÃ© | 0 jour | 
| GLM-5.3 | Non utilisÃ© | 0 jour | 
| GLM-5.2 | Non utilisÃ© | 0 jour | 
| GLM-5.1 | Non utilisÃ© | 0 jour | 
| Kimi K3 | Non utilisÃ© | 0 jour | 
| Kimi K2.7 Code | Non utilisÃ© | 0 jour | 
| Kimi K2.6 | Non utilisÃ© | 0 jour | 
| LongCat-2.0 | Non utilisÃ© | 0 jour | 
| MiMo-V2.6-Pro | Non utilisÃ© | 0 jour | 
| MiMo-V2.6-Flash | Non utilisÃ© | 0 jour | 
| MiMo-V2.5-Pro | Non utilisÃ© | 0 jour | 
| MiMo-V2.5 | Non utilisÃ© | 0 jour | 
| Qwen3.8 Max | Non utilisÃ© | 0 jour | 
| Qwen3.8 Flash | Non utilisÃ© | 0 jour | 
| Qwen3.7 Max | Non utilisÃ© | 0 jour | 
| Qwen3.7 Plus | Non utilisÃ© | 0 jour | 
| Qwen3.6 Plus | Non utilisÃ© | 0 jour | 
| MiniMax M3 | Non utilisÃ© | 0 jour | 
| MiniMax M2.7 | Non utilisÃ© | 0 jour | 
| Muse Spark 1.3 Contributor | Oui | Pas de ZDR | 
| Muse Spark 1.2 Contributor | Oui | Pas de ZDR | 
| DeepSeek V4.1 Flash | Non utilisÃ© | 0 jour | 
| DeepSeek V4 Pro | Non utilisÃ© | 0 jour | 
| DeepSeek V4 Flash | Non utilisÃ© | 0 jour | 
| DeepSeek V4 Flash Vision Exp | Non utilisÃ© | 0 jour | 
| Hy4 preview | Non utilisÃ© | 0 jour | 
| Hy3 | Non utilisÃ© | 0 jour | 
| Space Bunny Free | Non utilisÃ© | 0 jour | 

