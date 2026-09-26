---
id: ai-industry-kb-2026/11-multimodality-open-weight-omni-models/1-11-any-to-any-scoreboard-update-september-2026
title: "1.11 Any-to-any scoreboard update (September 2026)"
domain: multimodality-open-weight-omni-models
role: deep-dive
task: multimodal
actors: ["Alibaba", "ByteDance", "DeepSeek", "EU", "Google", "MiniMax", "Moonshot", "OpenAI", "United States", "Z.ai"]
dates: ["2026-09", "2026-09-18", "2026-09-22"]
keywords: ["apache", "benchmark", "deepseek", "gemini", "glm", "kimi", "leaderboard", "license", "multimodal", "omni", "open weights", "open-weight"]
source: docs/RAG/ai-industry-knowledge-base-2026.md
source_anchor: ""
source_lines: [5873, 5892]
section: "11. Multimodality — Open-Weight Omni Models"
sha256: df5ae750e239d32acf1b77dcd8021c41965d2e0424a9daf84585ebd9b889b294
---

# 1.11 Any-to-any scoreboard update (September 2026)

- **MiniMax M3 license is disputed**: best-of-ai and Morph report **Apache 2.0**; the agentone research file reports **MiniMax Community License** (commercially restricted). The HF repo (`MiniMaxAI/MiniMax-M3`) is the tiebreaker — verify the current license file before commercial use.
- **MiniMax H3 license**: territorially restricted — US, EU, UK, South Korea excluded from local deployment; not Apache-grade despite "open weights".
- **Qwen3.8-Omni-Flash**: no open weights were announced at launch (September 18, 2026); do not classify the 1M-context omni flagship as open-weight. The Qwen omni line otherwise remains open (Qwen3-Omni Apache 2.0, Qwen3.5-Omni Light dense open).
- **Kimi K3**: open weights under a bespoke Kimi K3 License with a MaaS resale gate — downloadable but not fully open-source in the OSI sense.
- **Vendor vs independent**: MiMo-V2.6, MiniMax M3, Qwen3.5-Omni, GLM-5.3-Flash, DeepSeek V4.1-Flash, HiDream-O1-Video-1.0 headline numbers are vendor-run. Independent confirmations exist for GLM-5.3-Flash (Ed Yau: 100% pass / 9.3/10 over 28 tasks, $0.28 vs $1.43 for GPT-5.5), MiniMax M3 (Morph fact-check Sept 7, 2026), Gemma 4 31B (LMArena #3). MiMo-V2.6's benchmark independence is pending as of September 22, 2026.
- **"Native" vs "modular"**: "native omni" should not be read as "single encoder" — MiMo-V2.6 (ViT + audio tokenizer + patch encoder), MoonViT-family models, Qwen3-Omni (Thinker/Talker), and MiniMax H3 (Qwen3-VL-32B as text encoder) are all natively multimodal *and* multi-encoder. Only Gemma 4 12B Unified claims an encoder-free single-transformer design.

### 1.11 Any-to-any scoreboard update (September 2026)

Wave 1 §1-era snapshot named Gemini Omni Flash as the only closed any-to-any model. September 2026 additions change the board:

| Model | Lab | Access | Inputs | Outputs | Status / caveat |
|---|---|---|---|---|---|
| Gemini Omni Flash | Google | Closed API | text, image, audio, video | video + synchronized audio (10s clips) | Still the benchmark-setter; $0.10/sec 720p (see §1) |
| HiDream-O1-Video-1.0 | HiDream.ai | API (closed weights) | text, image, video | 1080p video (5–20s) + synchronized audio | Company-reported only (Sept 17) [VENDOR] |
| Seedance 2.5 | ByteDance | Closed | text (reported) | 30s clips with built-in audio | Single secondary mention [UNVERIFIED]; full coverage in §12 |
| MiniMax H3 | MiniMax | Open weights (restrictive license) | text, image, video, audio | video + native stereo audio (4–15s, 768p local) | First open model to top an AI video ranking; US/EU/UK/KR excluded from local deployment |

The open side now has a genuine any-to-any video entry on at least one public leaderboard, but under a license that disqualifies it as Apache-grade open weights. The thesis "closed labs lead productized any-to-any output" survives, narrowed: it is now a *productization-and-licensing* lead, not a capability lead.

