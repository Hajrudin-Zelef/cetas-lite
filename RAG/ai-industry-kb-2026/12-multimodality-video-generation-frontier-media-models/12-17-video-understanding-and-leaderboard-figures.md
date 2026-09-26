---
id: ai-industry-kb-2026/12-multimodality-video-generation-frontier-media-models/12-17-video-understanding-and-leaderboard-figures
title: "12.17 Video-understanding and leaderboard figures"
domain: multimodality-video-generation-frontier-media-models
role: deep-dive
task: multimodal
actors: ["Google", "MiniMax"]
dates: ["2025-12", "2026-04"]
keywords: ["leaderboard", "agentic", "arr", "compute", "consumer", "cost", "gemini", "nvfp4", "research", "revenue", "text-to-image", "text-to-video"]
source: docs/RAG/ai-industry-knowledge-base-2026.md
source_anchor: ""
source_lines: [6328, 6346]
section: "12. Multimodality — Video Generation & Frontier Media Models"
sha256: fd6af9006c5f30d9eb2b62974b7d088ef785df063e06f2f49d4e34eda9ae80e1
---

# 12.17 Video-understanding and leaderboard figures

- Sora: reported ~$5.4B annual compute spend vs ~$2.1M lifetime revenue [UNVERIFIED — CIOL, no named methodology]; 1M+ downloads in first five days (ngram).
- Kling: reported **$240M ARR in December 2025** [UNVERIFIED — single secondary deep-research compilation].
- Google AI Ultra plan: **$249.99/mo** (Project Genie consumer gate).
- Adobe Firefly: **30+ video models** by April 2026 (Veo 3.1, Runway Gen-4.5, Kling 3.0 among them).
- gpt-image-1.5: $0.011–$0.167 per 1024x1024 by quality; Mini variant $0.005–$0.036; ~4s per 1024x1024 (~4x faster than prior); 5 input images preserved with high fidelity; 1B+ images generated via Gemini (Nano Banana Pro counterpoint).
- Nano Banana 2 Lite: $0.034 per 1K images; ~4s text-to-image at ~1K; $0.25/$1.50 per M tokens (gigazine).
- fal H3 Max: 5s 768p render in under 3s; launch discount expiry Sept 15, 2026 quadrupled 5s-clip price $0.10→$0.40; community NVFP4 quants ~175s per 10s clip on a single RTX 5090.

### 12.17 Video-understanding and leaderboard figures

- Google Agentic Video Understanding (Sept 2026): **−88% tokens, +~7% accuracy, −66% cost**; Gemini 3.7 Flash on the accuracy-vs-cost Pareto frontier on 1H-VideoQA.
- StreamArena (2026-08): 243 full-length videos (avg 88.8 min), 3,646 manually validated open-ended tasks, 774 proactive-monitoring tasks with ground-truth trigger times.
- LVBench: 103 videos, 117 hours total, avg 4,101s, 1,549 human QA pairs, 6 capability axes.
- CFD on LVBench: **52.9** overall vs MemVid 44.4 / VCA 41.3 / VideoTree 28.8 / VideoAgent 29.3 / AdaReTaKe-72B 53.3.
- Leaderboard claims (secondary/unverified unless noted): Seedance 2.0 claimed #1 Artificial Analysis text-to-video at launch [UNVERIFIED — secondary blog quoting the leaderboard]; MiniMax H3 #1 AA Video Editing (With Audio), #2 Text-to-Video, #3 Image-to-Video; H3 ~1476 Arena.ai image-to-video Elo; HiDream-O1-Video-1.0 company-reported #4 AA Image-to-Video (With Audio), #8 Arena.ai [VENDOR — no independent reproduction].
- Genie 3: 720p/24fps explorable worlds lasting several minutes; worlds shareable via public link; real-time interactive, not exportable files.

### 12.18 Clip-spec comparison (flagship video outputs, Sept 2026)

