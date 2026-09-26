---
id: ai-industry-kb-2026/03-chinese-ai-labs-deep-file-glm-kimi-minimax/sources-and-urls
title: "Sources and URLs"
domain: chinese-ai-labs-deep-file-glm-kimi-minimax
role: deep-dive
task: actor-profile
actors: ["China", "MiniMax", "Moonshot", "Z.ai"]
dates: ["2026-07", "2026-07-01", "2026-08-10", "2026-08-14", "2026-08-17", "2026-08-20", "2026-08-25"]
keywords: ["agent", "benchmarks", "claude", "cyber", "cybersecurity", "fp8", "glm", "kimi", "moe", "mxfp4", "omni", "open-weight"]
source: docs/RAG/ai-industry-knowledge-base-2026.md
source_anchor: ""
source_lines: [1313, 1366]
section: "3. Chinese AI Labs — Deep File: GLM, Kimi, MiniMax"
sha256: 196d298907c15022da6179c9089ac3dcf18430a1ad35ef7abd333a87f5427bff
---

# Sources and URLs

## Sources and URLs

### GLM-5.2 — parameter reconciliation and architecture [COMMUNITY / secondary research notes]

- https://github.com/t0msilver/working-set/blob/HEAD/research/model_glm52.md — [COMMUNITY] config-derived research note; source of the 743.4B / 753.3B / 39.3B reconciliation and the FP8-build origin of "744B"
- https://github.com/premdesai/pilot-shell/blob/HEAD/docs/docusaurus/blog/2026-07-01-glm-5-2.md — [COMMUNITY] GLM-5.2 practitioner write-up, July 2026
- https://github.com/zhongkaifu/tensorsharp/blob/HEAD/docs/models/glm.md — [COMMUNITY] tensorsharp architecture page with per-layer shapes (IndexShare mechanics)
- https://github.com/mattrobenolt/pi-configs/blob/HEAD/skills/model-guide/research/glm-5p2.md — [COMMUNITY] GLM-5.2 model guide research note

### GLM-5.3 — launch coverage and explainers

- https://github.com/bofai/docs/blob/HEAD/docs/llmservice/models/glm-5-3.md — [COMMUNITY] GLM-5.3 model documentation note
- https://www.unite.ai/z-ai-launches-glm-5-3-with-frontier-coding-and-a-cyber-capability-that-outgrew-its-training/ — [VENDOR-adjacent press] GLM-5.3 launch, frontier coding and cyber capability
- https://betanews.com/article/zai-glm-5-3-cybersecurity-delay/ — [press] GLM-5.3 cybersecurity weight-hold coverage; explicitly flags figures as company-reported
- https://emergent.sh/learn/what-is-glm-5-3 — [secondary] GLM-5.3 explainer
- https://memeburn.com/glm-5-3-is-here-benchmarks-pricing-coding-and-whats-new/ — [secondary] GLM-5.3 benchmarks, pricing, coding summary
- https://github.com/zerx-lab/zerx-lab-website/blob/HEAD/src/content/posts/daily-tech-news-2026-08-14/en.md — [COMMUNITY] daily tech news, 2026-08-14 (announcement-day record)
- https://github.com/diclogic/ai-daily-digest/blob/HEAD/digests/2026-08-20.md — [COMMUNITY] AI daily digest, 2026-08-20 ("Ox Alpha" / Flash rumor phase)

### Kimi K2.6 — launch coverage

- https://tpsreport.news/news/moonshot-ai-kimi-k2-6-1t-parameter-moe-agent-swarm — [secondary] Kimi K2.6 1T-parameter MoE agent-swarm coverage with architecture details
- https://datanorth.ai/news/moonshot-ai-releases-kimi-k2-6 — [secondary] Kimi K2.6 release (dates GA April 21)
- https://awesomeagents.ai/news/kimi-k2-6-agent-swarm-open-weight/ — [secondary] Kimi K2.6 agent-swarm open-weight coverage (dates GA April 20)
- https://kimi-k2.org/kimi-k26 — [secondary] Kimi K2.6 reference page (dates GA April 21)
- https://www.verdent.ai/guides/what-is-kimi-k2-6 — [secondary] Kimi K2.6 guide (dates GA April 20)
- https://apidog.com/blog/what-is-kimi-k2-6/ — [secondary] Kimi K2.6 explainer
- https://aitoolsrecap.com/Reviews/kimi-k2-6-review-2026 — [secondary] Kimi K2.6 review 2026
- https://www.frontiernews.ai/news/article/how-moonshot-ais-kimi-k2-became-a-top-open-weight-2961dd1f — [secondary] Moonshot Kimi K2 lineage background

### Kimi K3 — primary technical report and architecture inspections

- https://raw.githubusercontent.com/MoonshotAI/Kimi-K3/master/k3_tech_report.pdf — [PRIMARY] Kimi K3 technical report (Moonshot AI); source of Stable LatentMoE, Quantile Balancing, SiTU-GLU, MXFP4, ~2.5× scaling efficiency over K2
- https://github.com/sqliteai/warp/blob/HEAD/docs/K3.md — [COMMUNITY] K3 architecture inspection note
- https://github.com/tellebma/the-claude-codex/blob/HEAD/content/en/kimi-k3-tech-report.mdx — [COMMUNITY] K3 tech-report reading notes
- https://github.com/hyan-yao/homepage/blob/HEAD/content_zh/blog/kimi-k3.md — [COMMUNITY] K3 technical analysis (Chinese)
- https://arxiviq.substack.com/p/kimi-k3-open-frontier-intelligence — [secondary] Kimi K3 technical analysis

### MiniMax H3 / Hailuo 3.0 — launch and alias evidence

- https://srnnews.com/chinas-minimax-releases-h3-video-model/ — [press] Reuters launch-day coverage of MiniMax H3 (release July 30/31; Reuters also carried the 2.7T LLM report, early July 2026)
- https://www.marktechpost.com/2026/08/01/minimax-releases-minimax-h3-an-omni-modal-video-model-that-generates-15-second-2k-clips-with-native-stereo-audio/ — [secondary] MiniMax H3 launch coverage; uses "MiniMax H3" for the model live in the Hailuo AI app (alias evidence)
- http://news.minnesotaheadlines.com/story/628001/hailuo-ai-launches-h3-with-native-2k-video-and-onepass-synchronised-audio.html — [secondary] Hailuo AI H3 FAQ ("MiniMax H3, also known as Hailuo 3.0" — alias evidence)
- https://note.com/natty_quoll9640/n/n39ab69bca9fd — [COMMUNITY] Japanese name table: "MiniMax H3 ＝ Hailuo 3.0 ＝ Hailuo 03" (alias evidence)

### Seed 2.1 Turbo — dating and catalog evidence

- https://tpsreport.news/news/bytedance-seed-2-1-turbo-262k-context — [secondary] Seed 2.1 Turbo 262K-context coverage
- https://picksbymodel.com/models/bytedance-seed-seed-2-1-turbo/ — [catalog] Seed 2.1 Turbo model listing (specs, pricing)
- https://github.com/brainmox/agent-365/blob/HEAD/journal/2026/08/2026-08-25-seed-2-1-turbo-launched-twice.md — [COMMUNITY] engineering journal documenting the "launched twice, 47 days apart" dating discrepancy
- https://github.com/brainmox/agent-365/blob/HEAD/journal/2026/08/2026-08-17-seed-2-1-pro-and-the-all-superlatives-launch.md — [COMMUNITY] engineering journal on the chart-only-superlatives FORCE launch
- https://github.com/theopenco/llmgateway/commit/61d07d0f475a2131014021a9fa61e98d3d9fd119 — [catalog artifact] LLM Gateway PR #3580 commit adding Seed 2.1 Turbo (2026-08-10 first-seen evidence)
- https://writingmate.ai/models/bytedance-seed/seed-2-1-turbo — [catalog] Seed 2.1 Turbo model listing

