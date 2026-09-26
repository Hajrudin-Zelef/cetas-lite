---
id: ai-industry-kb-2026-wave6/08-mimo-and-xiaomi/v2-6-pro-v2-6-flash-live-rl-training-released-2026-09-21-22-
title: "V2.6-Pro / V2.6-Flash Live RL training (released 2026-09-21/22; details beyond base section)"
domain: mimo-and-xiaomi
role: deep-dive
task: training
actors: ["Alibaba", "Anthropic", "Hugging Face", "Moonshot", "OpenAI", "Xiaomi", "Z.ai", "xAI"]
dates: ["2026-03-19", "2026-09-21"]
keywords: ["training", "agent", "benchmark", "cost", "distillation", "fable 5", "glm", "gpt-5.6", "grok", "grok 4", "kimi", "multimodal"]
source: docs/RAG/ai-industry-knowledge-base-2026-wave6.md
source_anchor: ""
source_lines: [3673, 3707]
section: "§8. MiMo and Xiaomi"
delta_of: ai-industry-kb-2026
sha256: 630f68aee7174a52feb333b040d41c300ba6f410ec9d575c2aaff07ec6a60aa9
---

# V2.6-Pro / V2.6-Flash Live RL training (released 2026-09-21/22; details beyond base section)

### V2.6-Pro / V2.6-Flash Live RL training (released 2026-09-21/22; details beyond base section)
- V2.6-Pro and V2.6-Flash were trained with fully asynchronous GRPO (group relative policy optimization) reinforcement learning. [VENDOR, S21][SECONDARY, S22]
- The RL run used 1,568 prompts per training step. [VENDOR, S24][SECONDARY, S25]
- Each prompt generated 16 rollouts per step. [VENDOR, S24][SECONDARY, S25]
- Each RL step processed approximately 3.5B to 3.7B tokens. [VENDOR, S24 — single source]
- RL training ran at up to 1M-token context. [VENDOR, S24 — single source]
- Xiaomi reported three named RL innovations: Groupwise Reward Synthesis (GRS). [VENDOR, S21 — single source]
- Xiaomi reported Groupwise Advantage Redistribution (GAR). [VENDOR, S21 — single source]
- Xiaomi reported MOPD2, an on-policy distillation technique. [VENDOR, S24 — single source]
- On held-out DeepSWE v1.1 evaluation, Flash improved from 48.8 to 65.68 over the RL run. [VENDOR, S21][SECONDARY, S25]
- On held-out DeepSWE v1.1 evaluation, Pro improved from 58.4 to 72.57 over the RL run. [VENDOR, S21][SECONDARY, S25]
- The vendor launch benchmark table reports V2.6-Pro at 71.9 on DeepSWE v1.1 (vs 67.9 Flash, 19.0 V2.5-Pro), distinct from the 72.57 post-training figure — secondary coverage reproduces both figures and both settings. [VENDOR, S25][SECONDARY, S51][SECONDARY, S53][SECONDARY, S54]
- The vendor launch table reports V2.6-Pro at 53.1 on AutomationBench v1.0.6 — beating closed rivals' 50.3/45.8/46.2 per secondary reproduction. [VENDOR, S25][SECONDARY, S51][SECONDARY, S53]
- The vendor launch table reports V2.6-Pro at 76.9 on Toolathlon-Verified. [VENDOR, S25][SECONDARY, S51][SECONDARY, S53]
- The vendor launch table reports V2.6-Pro at 89.9 on Terminal-Bench 2.1 — best in the table per secondary reproduction. [VENDOR, S25][SECONDARY, S52][SECONDARY, S53]
- The vendor launch table reports V2.6-Pro at 62.0 on JobBench. [VENDOR, S25][SECONDARY, S52]
- The vendor launch table reports V2.6-Pro at 94.0 on CyberGym (Flash 95.1 vs GLM 5.3's 84.5). [VENDOR, S25][SECONDARY, S51][SECONDARY, S53]
- The vendor launch table reports V2.6-Pro at 72.3 on MiMo Visual Coding (ahead of Opus 5's 70.0 and Fable 5's 69.1, behind GPT-5.6 Sol's 73.4). [VENDOR, S25][SECONDARY, S52][SECONDARY, S53]
- Hugging Face checkpoints for the V2.6 pair appeared at approximately 2026-09-21 15:39 UTC. [COMMUNITY, S24 — single source]
- The Flash and Pro checkpoint uploads were approximately 18 seconds apart. [COMMUNITY, S24 — single source]
- Reported RL training cost for V2.6-Flash is approximately $0.85M; the run spanned 30 RL steps over ~750,000 trajectories in under six days with 1,568 samples per update. [SECONDARY, S24][SECONDARY, S54][SECONDARY, S56]
- Reported RL training cost for V2.6-Pro is approximately $2.62M, roughly 3.1× the Flash cost; Xiaomi also released 7,000+ task environments and RL resources alongside the models. [SECONDARY, S24][SECONDARY, S51][SECONDARY, S56]
- Xiaomi released the V2.6 technical report, the RL training environments, and training code alongside the weights. [SECONDARY, S22][SECONDARY, S21]
- The RL training process was publicly observable ("watch a 1T AI model train live") during the six-day publicly tracked run. [SECONDARY, S24][SECONDARY, S56]
- V2.6-Pro debuted as the top open-weights model at AA Intelligence Index v4.3 (46.32, tied with Grok 4.7), ahead of Kimi K3 and Qwen3.8 Max; the model is MIT-licensed with 1M context and native multimodal inputs. [SECONDARY, S21][SECONDARY, S55][SECONDARY, S53]

### Xiaomi AI investment and ecosystem (new to this section)
- CEO Lei Jun announced on 2026-03-19 that Xiaomi will invest at least CNY60B (approximately $8.7B) in AI over the next three years. [SECONDARY, S9][SECONDARY, S10][SECONDARY, S11]
- The "at least" phrasing indicates a floor commitment, not a capped budget. [SECONDARY, S9][SECONDARY, S10]
- Xiaomi's 2026 AI R&D budget alone exceeded CNY16B (approximately $2.3B). [SECONDARY, S11 — single source]
- Xiaomi's broader five-year core-technology budget is CNY200B, spanning chips, AI, and operating systems. [SECONDARY, S11 — single source]
- Xiaomi shares (1810.HK) rose approximately 5% following the investment announcement. [SECONDARY, S12 — single source]
- MiClaw, a smartphone AI agent, entered closed beta testing. [SECONDARY, S8 — single source]
- MiClaw is integrated into Xiaomi's operating system and the Human × Car × Home ecosystem strategy. [SECONDARY, S8 — single source]

