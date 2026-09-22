---
id: ai-industry-kb-2026/03-chinese-ai-labs-deep-file-glm-kimi-minimax/main-actors
title: "Main actors"
domain: chinese-ai-labs-deep-file-glm-kimi-minimax
role: deep-dive
task: actor-profile
actors: ["Alibaba", "Anthropic", "ByteDance", "China", "DeepSeek", "EU", "MiniMax", "Moonshot", "OpenAI", "United States", "Z.ai"]
dates: ["2026-01", "2026-06-23", "2026-07", "2026-07-01", "2026-08-10", "2026-08-14", "2026-08-17", "2026-08-20", "2026-08-25", "2026-09", "2026-09-22"]
keywords: ["agent", "agentic", "attention", "benchmark", "benchmarks", "claude", "compute", "cost", "cyber", "cybersecurity", "deepseek", "fable 5"]
source: docs/RAG/ai-industry-knowledge-base-2026.md
source_anchor: ""
source_lines: [1233, 1366]
section: "3. Chinese AI Labs — Deep File: GLM, Kimi, MiniMax"
sha256: 7c26d67fa0cec188459853a3bf217ef36c38af06c82ec211ce5f5db021abb22c
---

# Main actors

## Main actors

### The labs

- **Z.ai / Zhipu AI** (Beijing): the GLM line (GLM-5.2, GLM-5.3, GLM-5.3-Flash); IndexShare; the GLM Coding Plan subscription and ZCode agent; OpenVuln/VulnHunter. Listed publicly in Hong Kong before MiniMax (Reuters notes MiniMax was second of the "AI tigers" to list, following Z.AI).
- **Moonshot AI** (Beijing): the Kimi line (K2.5 → K2.6 → K3); Modified MIT open-weight licensing with the 100M-MAU / $20M-revenue display threshold; MoonViT vision encoder; Stable LatentMoE and Quantile Balancing as 2026 MoE design contributions.
- **MiniMax** (Shanghai, founded 2022): one of China's "AI tigers" — the well-funded startup group competing with domestic tech giants and US labs. Publicly listed in Hong Kong in **January 2026**. The "H" (video/Hailuo) vs "M" (text LLM) naming convention disambiguates H3 (video) from the M2/M3 text models. [UNVERIFIED] 2.7T-parameter LLM reportedly in development (Reuters, early July 2026).
- **ByteDance Seed** (Volcano Engine): Seed 2.1 family (base, Pro, Turbo) launched at the FORCE conference 2026-06-23/24; Seedance video models (Seedance 2.5 canonical in wave3/06 §1.3). Positioned by secondary review around "productivity delivery" — end-to-end software delivery from planning to verification — rather than raw benchmark leadership.

### Individuals

- **Yang Zhilin** — Moonshot AI founder (one secondary source notes he was a student of Tang Jie; biographical, low relevance).
- **Tang Jie** — the professor who later co-founded Z.AI — the biographical link between the two labs.

### Benchmark counterparties (for comparative framing, not lab actors)

- Claude Opus 4.6 (53.4% SWE-Bench Pro — the number K2.6 beat); Claude Opus 4.8 (the model K3 beats on GDPval-AA v2, benchmark-dependently); GPT-5.4 (57.7% SWE-Bench Pro); Anthropic's Fable 5 / Mythos 5 (the export-control suspension that GLM-5.2's launch followed); DeepSeek and Qwen3.8-Max (the open-weight frontier context for the 2.7T rumor tier: Kimi K3 2.8T, Qwen3.8-Max 2.4T).

### Business context (Reuters launch-day coverage, MiniMax H3)

- MiniMax framed H3 as stepping up competition in a market **"led by rivals ByteDance and Kuaishou"** — consistent with wave3/06's "China-led, closed-model race" thesis for text-to-video.
- The H3 launch commercial claim (2K video at <1/3 rival cost) [VENDOR] and the January 2026 HK listing together mark MiniMax as the capitalized challenger in the video-model race.

### Positioning notes

- **ByteDance's thesis** (secondary review): "productivity delivery" — end-to-end software delivery from planning to verification — rather than raw benchmark leadership. This is how the RAG should frame ByteDance's 2026 lab strategy versus the frontier-chasing posture of Moonshot and Z.ai.
- **Kuaishou** appears alongside ByteDance as the market co-leader in Reuters' framing of the video-model race — relevant context for H3's competitive positioning, not a model actor in this file.
- **Z.ai vs Moonshot founder link**: Moonshot's Yang Zhilin was reportedly a student of Tang Jie, who later co-founded Z.AI — the two labs' leadership shares an academic lineage (low relevance, recorded for completeness).
- **DeepSeek and Qwen** are the open-weight frontier reference points around these releases: DeepSeek Sparse Attention is GLM-5.2's attention substrate, and Qwen3.8-Max (2.4T) defines the tier the rumored MiniMax 2.7T LLM would enter.

## Timeline and context

The April-to-September arc for these three labs runs as a single competitive sequence. **April**: Moonshot opens the Kimi K2.6 preview (04-13) and reaches GA (04-20/21) with 1T/32B open weights, 384 experts, a 300-sub-agent swarm mode, and SWE-Bench Pro 58.6% — ahead of Claude Opus 4.6's 53.4%, a data point the brief undersold as "near." **June**: Z.ai ships GLM-5.2 (weights 06-16) the day after the US export-control order suspending Anthropic's Fable 5 / Mythos 5 — the open-weight counter-move, with 1M context, IndexShare, and the "best open-weight text model" crown valid only at that moment [DIRECTIONAL]; ByteDance launches the Seed 2.1 family at FORCE (06-23/24), the Turbo variant resting on a single trade report. **July**: Moonshot launches Kimi K3 (07-16; weights ~07-27) at 2.8T with the bandwidth-aware latent-MoE design, and MiniMax launches H3 = Hailuo 3.0 (07-31) — the first open model to top an AI video ranking — while Reuters reports the [UNVERIFIED] 2.7T MiniMax LLM in development. **August**: Z.ai announces GLM-5.3 (08-14) on the same 743B base, all gains from scaled post-training — the clearest 2026 instance of the post-training-scaling thesis (Terminal-Bench 3.0: 4.6 → 28.3 without new pretraining); the cyber capability outgrew Z.ai's training expectations, triggering the GLM line's first weight hold (~2 weeks; weights 08-28/29 under a bespoke license, a step down from MIT). **September 22**: the MiniMax 2.7T rumor and the Moonshot HK IPO rumor remain unconfirmed watch items.

Exact dates live in §2 (Key dated facts); the rumor-phase rule (pre-2026-08-14 GLM-5.3 mentions stay [UNVERIFIED]) applies to anything cited from the August speculation window.

### Month-by-month spine (April → September 2026)

- **April**: Kimi K2.6 — the agent-swarm open-weight release; 1T/32B, Modified MIT, SWE-Bench Pro ahead of Opus 4.6. Sets the 2026 template: trillion-scale open weights with agentic modes.
- **June**: GLM-5.2 (06-16) — the post-Anthropic-export-control counter-move; 753B/40B, 1M context, IndexShare; the "best open-weight text model" moment [DIRECTIONAL]. Seed 2.1 family at FORCE (06-23/24).
- **July**: Kimi K3 (07-16 launch; ~07-27 weights) — the bandwidth-aware 2.8T push (latent routing, Quantile Balancing, MXFP4). MiniMax H3 (07-31) — the open-video milestone; Reuters' [UNVERIFIED] 2.7T LLM report.
- **August**: GLM-5.3 — the post-training-scaling demonstration and the cyber-driven weight hold; rumor phase → 08-14 announcement → 08-18 pricing → 08-28/29 weights.
- **September**: verification cutoff 2026-09-22 — both watch items (MiniMax 2.7T, Moonshot HK IPO filing) still unconfirmed.

## Implications

### For the knowledge base design

- **Alias discipline is load-bearing.** Three names for one MiniMax model (H3 / Hailuo 3.0 / "Hailuo 03"), two numbers for one GLM-5.2 total ("744B" shorthand vs 753.3B full), three phases for one K2.6 release (preview / GA-day-1 / GA-day-2), and two dates for one Seed 2.1 Turbo launch (FORCE 06-24 vs tracker first-seen 08-10/12) — the KB must merge these into single entries with phase footnotes, or retrieval will triple-count. The consolidation notes (§8) carry the merge rules.
- **Counting conventions must be footnoted, not picked.** The GLM-5.2 case (backbone vs +MTP vs active vs FP8-build shorthand) is the template: pin the canonical total (753B/40B) and record the convention behind each competing figure.
- **Temporal bounds on superlatives.** "Best open-weight text model," "beats Opus 4.8," "first open model to top an AI video ranking" — each is [DIRECTIONAL] and valid only at its moment; the KB should date every leadership claim.

### Competitive implications

- **Open-weight frontier pressure from Chinese labs**: 753B/40B (June), 1T/32B (April), and 2.8T (July) open weights shipped within four months, all with permissive-to-pragmatic licensing (MIT; Modified MIT with a high display-gating threshold).
- **The post-training-scaling thesis, demonstrated**: GLM-5.3's 6× Terminal-Bench 3.0 jump (4.6 → 28.3) on the *same* 743B base, via RL environments, task diversity, longer trajectories, and more RL compute — no new pretraining. This is the strongest 2026 evidence point for post-training as the scaling lever.
- **Bandwidth-aware MoE as the 2026 design pattern**: K3's Stable LatentMoE (ℓ=3,584 vs 7,168, ~halving routed traffic), Quantile Balancing (single all-reduce), and MXFP4 (E8M0 per 32 weights) show how top-16-of-896 routing is made trainable — granularity is paid for with co-designed compression, not claimed as free.
- **Video**: H3's #1 on Artificial Analysis (Video Editing with Audio) marks the first open model topping an AI video ranking, inside wave3/06's China-led video-race thesis.

### Safety and policy implications

- **Cyber capability as a release-governance trigger**: GLM-5.3's vulnerability-discovery ability (2,436 findings, 53 CVEs via OpenVuln) grew faster than Z.ai expected *during training*, producing the GLM line's first weight hold (~2 weeks of safety hardening) and a bespoke non-MIT license for the 5.3 weights. This is the KB's canonical 2026 case of a capability overhang changing a release plan mid-flight.
- **License territorial restrictions as a new axis**: MiniMax's H3 Community License excludes the US, EU, UK, and South Korea from local deployment — a licensing-geography dimension the KB's license index should track separately from royalty/copy-left terms.
- **Rumor-phase hygiene**: the GLM-5.3 August window (rumor → announcement 08-14 → weights 08-28/29) is the template rule — never merge pre-announcement speculation dates with post-announcement confirmations.

### If the watch items resolve

- **MiniMax 2.7T LLM confirmed**: it would join the 3T-class open-weight tier (Kimi K3 2.8T, Qwen3.8-Max 2.4T); the KB would need a new model entry with the same alias/dating discipline applied to H3, plus a check on whether the "M" naming convention holds.
- **Moonshot HK IPO filing confirmed**: Moonshot joins Z.AI and MiniMax as listed "AI tigers" — update the business-context entries, not the model entries.

### The 2026 Chinese open-weight playbook (synthesis)

Four moves repeat across Z.ai, Moonshot, and MiniMax in this file's window, and the KB should treat them as the house strategy rather than isolated releases: (1) **ship the biggest open weights of the season** — 1T/32B (April), 753B/40B (June), 2.8T (July) — each timed against a Western frontier event (Opus 4.6 on SWE-Bench, the Anthropic export-control suspension); (2) **license pragmatically** — MIT or Modified MIT for reach, with the display-gating threshold (100M MAU / $20M revenue) set far above community-license norms, and bespoke or territorial terms reserved for the most capable artifact; (3) **gate access subscription-first, API-second, weights-last** — the GLM Coding Plan / ZCode pattern, with weights trailing by ~2 weeks (or indefinitely for closed modules like H3's 2K upscaler); (4) **scale post-training, not just pretraining** — GLM-5.3 is the purest 2026 instance: same 743B base, all gains from RL environments, task diversity, longer trajectories, more RL compute. ByteDance is the deliberate exception: closed weights, no parameter counts, "productivity delivery" positioning instead of benchmark leadership.

### Retrieval design rules derived from this file

- Architecture claims (expert counts, routing formulas, attention mechanisms) and systems claims (bandwidth, VRAM, speedups) must be stored as separate facts with separate evidence grades — the 2026 MoE literature shows they move in opposite directions for fine-grained designs.
- Every benchmark figure in this file needs three attached attributes: the harness version (TB 2.1 vs 3.0 scores are not comparable), the evidence label ([VENDOR] vs secondary), and the date of the claim.
- License entries need a fourth attribute beyond terms: **territorial scope** (the H3 exclusion) and **temporal scope** (GLM-5.3's bespoke license vs 5.2's MIT).

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

