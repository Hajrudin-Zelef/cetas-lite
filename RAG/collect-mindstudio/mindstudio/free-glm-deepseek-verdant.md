---
id: collect-mindstudio/mindstudio/free-glm-deepseek-verdant
title: "How to Get GLM 5.3 Flash and DeepSeek V4 Flash Free in Verdant"
domain: mindstudio
role: reference
task: article
actors: ["Anthropic", "DeepSeek", "Google", "Moonshot", "OpenAI", "Z.ai"]
dates: ["2026-09-23"]
keywords: ["deepseek", "glm", "agent", "agentic", "agents", "chatgpt", "claude", "cost", "fable 5", "gemini", "inference", "kimi"]
source: docs/RAG/Collect RAG/02_mindstudio/free-glm-deepseek-verdant.md
source_anchor: ""
source_lines: [1, 66]
sha256: da8dd79e3f23f6fdb97bbffe0197bc497719797bbb722d1b27d5fcb7342ece62
---

# How to Get GLM 5.3 Flash and DeepSeek V4 Flash Free in Verdant

## Metadata

- **Source** : https://www.mindstudio.ai/blog/free-glm-deepseek-verdant
- **Site** : MindStudio
- **Type** : Article
- **Language** : en
- **Verification status** : ✅ reachable
- **Collection date** : 2026-09-23

## Full summary

This article covers Verdant, a desktop multi-agent coding app, currently offering two newly released models completely free: GLM 5.3 Flash from ZAI (Zhipu AI) and DeepSeek V4 Flash. Both carry usage limits measured on a 5-hour and weekly basis rather than a hard cap — wide enough that most individual users won't hit them during normal daily coding work.

GLM 5.3 Flash details: a mixture-of-experts model with 320 billion total parameters but only 18 billion active per token, which is why it runs fast inside an agent loop. MIT licensed, and the first natively multimodal model in the GLM 5 series. ZAI trained it to look at rendered interfaces, gameplay, and 3D output and judge its own work through visual feedback — Verdant calls it a "visual work partner." In practice, when an agent builds a front end it can inspect the actual rendered result and fix visual issues rather than assuming the code compiled correctly. This self-checking loop is the more useful trait than raw speed for UI-heavy/front-end work.

DeepSeek V4 Flash (the 0731 version): the second free model, positioned for deep reasoning and agentic tasks with cost-efficient inference — heavier logic and multi-step planning rather than visual checks. The two models split work naturally: GLM 5.3 Flash covers front-end/visual/multimodal tasks; DeepSeek V4 Flash covers reasoning-heavy and agentic tasks. Both were released within the last month and are free during this promotional window — between them, a user can plausibly handle most day-to-day coding tasks without paying a single credit.

Verdant itself: a desktop app built around multi-agent coding — you give it a task and it plans, writes code, and verifies results on its own, with the option to run multiple agents in parallel. Includes skills, plugins, a plan mode, and the ability to deploy cloud apps directly. The model picker shows GLM 5.3 Flash with a free tag and listed at 0.03x normal credit usage (currently costs nothing within stated time/weekly limits). DeepSeek V4 Flash sits alongside under the same free arrangement.

Pricing beyond free models: free trial gives 7 days and 100 credits with access to flagship models (Claude Opus 5, Sonnet 5, GPT 5.6, Gemini 3.1 Pro, GLM 5.2, Kimi K3); trial also supports bring-your-own API key and includes one cloud app with 10GB project storage and a 500MB database. Paid tiers: Light $5/month (Eco mode; GPT 5.6, Luna, GLM 5.2, Kimi K3, K2.7 Code, DeepSeek V4 Pro; two cloud apps, 20GB storage, optional add-on credits), Starter $19/month (bonus: 480 credits/month vs base 320; Claude Fable 5, Opus 5, Sonnet 5, GPT 5.6, Gemini 3.1 Pro, Kimi K3, GLM 5.2; Eco mode doesn't consume credits), Pro $59/month (1,500 credits after bonus), Max $179/month (4,500 credits after bonus). Comparison: Claude Pro and ChatGPT Plus cost $20/month with tighter limits and a single model family; Verdant's $5 Light plan + free flash models compares favorably on raw value.

Limitations: free access is explicitly limited-time — pricing and bonus credits could change; flash models are fast/efficient, not flagship-tier reasoning engines — for genuinely hard problems, larger models (Fable 5, Opus 5, full-size GLM 5.3) still outperform. Sensible approach: free models as default for everyday coding/visual checks; reserve paid credits/flagship models for deeper-reasoning tasks.

## Key points

- Verdant (multi-agent coding app) currently offers GLM 5.3 Flash (ZAI) and DeepSeek V4 Flash (0731) completely free with 5-hour + weekly usage limits instead of hard quotas.
- GLM 5.3 Flash: 320B total / 18B active MoE, MIT licensed, first natively multimodal GLM 5 model — visually inspects its own rendered output (UI, gameplay, 3D).
- DeepSeek V4 Flash: built for deep reasoning and agentic tasks with cost-efficient inference.
- Complementary split: GLM 5.3 Flash for front-end/visual/multimodal; DeepSeek V4 Flash for reasoning-heavy/agentic work.
- Free trial: 7 days + 100 credits with flagship access (Claude Opus 5, GPT 5.6, Gemini 3.1 Pro, etc.).
- Paid tiers: Light $5/mo, Starter $19/mo (480 credits w/ bonus), Pro $59/mo (1,500), Max $179/mo (4,500).
- Both free offers are limited-time; flash models aren't flagship-tier for hard reasoning.

## Technical data / figures

| Model | Architecture | License | Positioning |
|---|---|---|---|
| GLM 5.3 Flash | MoE: 320B total / 18B active per token | MIT | Natively multimodal; visual self-checking; 0.03x credit usage (free now) |
| DeepSeek V4 Flash (0731) | — | — | Deep reasoning + agentic, cost-efficient inference |

| Verdant plan | Price | Highlights |
|---|---|---|
| Free trial | $0 (7 days) | 100 credits; Opus 5, GPT 5.6, Gemini 3.1 Pro, GLM 5.2, Kimi K3; 1 cloud app, 10GB storage, 500MB DB |
| Light | $5/mo | Eco mode; GPT 5.6, Luna, GLM 5.2, Kimi K3, K2.7 Code, DeepSeek V4 Pro; 2 cloud apps, 20GB |
| Starter | $19/mo | 480 credits (bonus); Claude Fable 5, Opus 5, Sonnet 5, GPT 5.6, Gemini 3.1 Pro, Kimi K3 |
| Pro | $59/mo | 1,500 credits (after bonus) |
| Max | $179/mo | 4,500 credits (after bonus) |

| Feature | Detail |
|---|---|
| GLM 5.3 Flash credit cost | 0.03x normal; currently free |
| Limits type | 5-hour + weekly (not hard cap) |
| Comparison | Verdant $5 Light vs Claude Pro / ChatGPT Plus $20/mo |
| Flagship models still outperform flash | Fable 5, Opus 5, full-size GLM 5.3 |

## Why this source matters for the RAG

Provides current, time-sensitive information about free access to two recent frontier-flash models (GLM 5.3 Flash, DeepSeek V4 Flash) and Verdant's multi-agent pricing — useful for cost modeling and model availability in agentic/RAG workflows. Includes concrete architecture specs (320B/18B MoE, multimodal visual self-checking) and comparative pricing data for model/tool selection.

## Related context from the article

- GLM 5.3 Flash is MIT-licensed — relevant to self-hosting/local inference options.
- Both free models were released within the last month (as of Aug 2026); free window is promotional.
- Verdant free mode / pricing are covered in separate related MindStudio articles.
