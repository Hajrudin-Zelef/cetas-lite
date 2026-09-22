---
id: ai-industry-kb-2026/16-hugging-face/main-actors
title: "Main actors"
domain: hugging-face
role: deep-dive
task: platform
actors: ["AWS", "Alibaba", "China", "DeepSeek", "EU", "Google", "Groq", "Hugging Face", "Moonshot", "Nvidia", "OpenAI", "SGLang", "United States", "Unsloth", "Z.ai", "vLLM"]
dates: ["2023-08", "2025-04", "2025-07", "2025-12", "2026-01", "2026-02-20", "2026-03-21", "2026-04-23", "2026-04-27", "2026-04-28", "2026-05", "2026-06", "2026-07", "2026-08", "2026-08-26", "2026-08-27", "2026-09", "2026-09-01", "2026-09-02", "2026-09-03", "2026-09-15", "2026-09-22"]
keywords: ["acquisition", "advisory", "agent", "agents", "apache", "arr", "compute", "deepseek", "distillation", "distribution", "energy", "funding"]
source: docs/RAG/ai-industry-knowledge-base-2026.md
source_anchor: ""
source_lines: [8039, 8194]
section: "16. Hugging Face"
sha256: d080b47d956e174825a0031b83220e4d1bfda48e2cb3463b4190bb3904accf19
---

# Main actors

## Main actors

### Corporate and executive

- **NVIDIA Corporation** — acquirer under the 2026-09-02 definitive agreement; AI-chip-market leader; deal value $12.9303B, would be its largest acquisition if completed; disclosed ~$18B committed to equity investments for the rest of FY2027.
- **Jensen Huang** — NVIDIA CEO; issued verbatim open-platform commitments on 2026-09-03 ("Hugging Face will remain an open platform for the entire AI ecosystem. Developers will choose the models they want, the frameworks they want, the clouds and inference service providers they want and the computing platforms they want. Nvidia compute will not be required to build on or deploy through Hugging Face."); highlighted Reachy Mini in his CES 2026 keynote.
- **Hugging Face, Inc.** — target; founded 2016 in New York; ~750 employees; $395M+ raised (Crunchbase); $4.5B valuation at the August 2023 round; ~$150M annualized revenue (Sacra, August 2026) [DIRECTIONAL].
- **Clément Delangue** — Hugging Face CEO and co-founder (French); initiated 2026 talks with NVIDIA seeking "more compute, more support, more collaboration, more visibility" (X post); public rationale for robotics: "robotics does not get dominated by just a few big players with dangerous black-box systems" (Outlook Business, September 2026); source of the ~1M-datasets / "half of Fortune 500" figures [VENDOR, self-reported/unaudited].
- **Julien Chaumond and Thomas Wolf** — co-founders (French), with Delangue.

### Robotics and partners

- **Pollen Robotics** — Bordeaux, France; acquired April 2025; original Reachy robot; Reachy 2 (~$70K); 100+ units deployed across 20+ countries; the enabling acquisition for the Reachy Mini/HopeJR line.
- **Seeed Studio** — manufacturing partner that took Reachy Mini from prototype to mass production (3,000+ pre-orders late 2025; 3,000 units shipped by January 2026).
- **ggml.ai** — acquired by Hugging Face (announced 2026-02-20); the GGML tensor-library company, consolidating HF's grip on the open-weight inference-format layer.

### Security research community

- **'chenpinji'** — independent researcher; reported the LeRobot pickle flaw to HF in December 2025; HF's response ("experimental," needs refactoring) preceded the CVE by four months.
- **Valentin Lobstein ("chocapikk")** — published the public PoC on 2026-04-28, triggering press coverage.
- **Lyrie Research** — published the two definitive technical analyses (2026-04-28/29) establishing CVSS 9.8/9.3, the `# nosec` detail, and the physical-safety dimension.
- **Resecurity, The Hacker News, CyberPress, Ethical Hacking News, GitHub lemmaoracle** — secondary technical coverage 2026-04-27/28.
- **SentinelOne / GitHub Security Advisory** — documented the companion CVE-2026-0599 (TGI DoS); recommends upgrade + reverse-proxy request limits.

### Regulators and analysts (watch list)

- **US authorities** — already examining NVIDIA's AI-chip-market practices; review jurisdiction for the deal flagged in coverage (thefoxdaily).
- **EU competition authorities** — flagged jurisdiction; focus: whether a dominant firm controlling a strategically important developer platform can strengthen its hardware position.
- **Sid Nag (Tekonyx)** — on record: deal "positive for open-model funding and adoption but potentially negative for ecosystem neutrality" (via Tech Times).
- **Siddy Jobe (fund manager, via eponalab summary of CNBC)** — NVIDIA building control "from energy to foundational models and applications."
- **Sacra** — source of the ~$150M annualized-revenue estimate (August 2026) behind the 86× multiple [DIRECTIONAL, unaudited].
- **The Information** — broke the 2026-08-26 report; projected the deal "will effectively restart the AI cloud business Nvidia shut down a year ago" [DIRECTIONAL, analyst projection].
- **Bloomberg** — reported advanced talks 2026-09-01 ("agreement possible as soon as this week").

### Developer-side stakeholders

- **18M developers and 200K companies** on the platform per September 3, 2026 deal coverage [VENDOR via press].
- **"Half of Fortune 500"** deploying private + open models on the platform, per Delangue [VENDOR, self-reported/unaudited].
- **Open-source library ecosystem** — Transformers (113M+ monthly downloads), Diffusers, PEFT/LoRA, Accelerate, Datasets, TRL, Optimum, smolagents (code-first agent framework; turned Spaces into agent hosts in 2026).

## Timeline and context

### 2016–2023: from chatbot startup to the GitHub of AI

Hugging Face was founded in 2016 in New York by French founders Clément Delangue, Julien Chaumond
and Thomas Wolf. It does not build frontier models; it runs the registry and distribution layer for
open models, datasets, and Spaces (interactive apps) — the discovery-to-deployment funnel for
open-weight AI. Funding totaled $395M+ (Crunchbase); the last venture round was **August 2023: $235M
led by Salesforce Ventures at a $4.5B valuation**, with Google, Amazon, IBM, and NVIDIA itself
participating. Revenue trajectory (Sacra estimates, secondary [DIRECTIONAL]): ~$81M ARR end of 2025
→ ~$100M (June 2026) → ~$150M annualized (August 2026).

### 2024: LeRobot and the robotics thesis

In 2024 Hugging Face launched **LeRobot** — open AI models, datasets, and tools for building
robotics systems (12K–24K GitHub stars across 2026). Delangue's stated rationale: open, affordable,
inspectable hardware as the counter-model to "a few big players with dangerous black-box systems."

### 2025: Pollen acquisition, Reachy Mini, the ggml and infrastructure moves

- **April 2025**: Hugging Face **acquires Pollen Robotics** — the French startup behind the original Reachy robot and Reachy 2 (~$70K, 100+ units across 20+ countries, Orbita 7-DOF joints, Apache-2.0 software).
- **July 2025**: **Reachy Mini debuts** — the date that must not be misread as 2026. A 3D-printed prototype (Lite $299, Full/Wireless $449), ~11 in (28 cm) tall, with 15 pre-installed demos; trade press reported $500K in first-24-hour sales (Digital Trends) and $1M in the first five days (one source) [COMMUNITY].
- **Late 2025**: 3,000+ pre-orders in a week; manufacturing moves to mass production with Seeed Studio.
- **Late 2025 (reported January 2026)**: Hugging Face **rejects a $500M NVIDIA investment offer** at a $7B valuation (FT via Reuters) — reportedly to keep any single investor from gaining outsized influence over platform decisions. The later $12.93B agreed price is roughly 3× the $4.5B 2023 valuation and nearly 2× the rejected $7B mark.
- **December 2025**: researcher 'chenpinji' reports the LeRobot pickle-deserialization flaw to HF; the response is that the implementation is "experimental" — four months before the CVE goes public.
- **December 2025 (reported)**: NVIDIA signs its ~$20B non-exclusive Groq licensing deal, and discloses ~$18B committed to equity investments for the rest of FY2027 — the capital posture behind the Hugging Face bid.

### 2026-02-20: ggml.ai acquired

Hugging Face announces the acquisition of **ggml.ai**, absorbing the GGML machine-learning tensor
library company — a move that tightens HF's control over the open-weight inference-format layer
alongside llama.cpp/GGUF distribution through the Hub.

### 2026-03-21: TGI archived — the orchestration correction

Hugging Face's official position: "text-generation-inference is now in maintenance mode" — accepting
only minor bug fixes and docs; the repository **huggingface/text-generation-inference was archived
(read-only) on 2026-03-21**, with HF recommending migration to **vLLM, SGLang**, or local engines
(llama.cpp, MLX). TGI v3 (late 2025) had been positioned as a long-prompt specialist (vendor docs
claimed 13× speedup over vLLM on 200K+ token prompts — 27.5s → ~2s [UNVERIFIED, vendor docs via
secondary coverage]), plus a router/model-server architecture. Ecosystem tooling actively removed
TGI providers in 2026 refactors (Llama Stack distributions, Dell's distribution). Existing TGI
deployments remain in production, but TGI is a **migration signal, not a growth layer** — any
consolidation must not present it as central orchestration.

### 2026-04-23/28: CVE-2026-25874 goes public

NVD publishes **CVE-2026-25874** (2026-04-23); press and two Lyrie analyses land 2026-04-27/28;
chocapikk's PoC drops 2026-04-28. The technical attack chain: LeRobot's async inference splits robot
control (RobotClient on the device) from heavy model inference (PolicyServer on a GPU machine); both
sides run `pickle.loads()` on gRPC-received data with **no authentication and no TLS**
(`add_insecure_port()`); pickle executes arbitrary code during deserialization → unauthenticated RCE
on the host. Consequences compound: GPU-backed inference hosts typically run with elevated
privileges (internal-network access, datasets, API keys, SSH credentials, model files, expensive
compute); lateral movement into lab/enterprise networks; and RCE reaching the robot's
**joint-control path** — physical sabotage, not just data theft. Fix planned in **0.6.0** (pickle →
safetensors/JSON); unpatched as of 2026-09-22.

### 2026-01 to 2026-09: platform growth and the infrastructure that actually scaled

While TGI wound down, the real orchestration layer — **Inference Providers** — scaled: single
endpoint (router.huggingface.co), 15+ backend partners, automatic `:fastest`/`:cheapest` selection,
failover, no markup, one HF token. Monthly free credits ($0.10 free / $2.00 PRO). The legacy
serverless API became **hf-inference** (CPU-focused since July 2025). **Inference Endpoints** sell
dedicated GPUs at list rates (T4 $0.50/hr … H100 $10.00/hr). The library layer kept compounding:
Transformers 113M+ monthly downloads; Datasets 730k+ with 500k+ covering 8,000+ languages; **TRL +
GRPO** made R1-style RL post-training a pip-installable recipe (the open-RL stack that commoditized
reasoning-model production); **smolagents** turned Spaces into agent hosts; Spaces held at ~1M
interactive apps. Model growth: 2.4M+ (May 2026) → 3M (September 3, 2026 deal coverage).

### 2026-07: HopeJR's context and the breach cross-ref

The July 2026 autonomous-breach incident targeting Hugging Face's platform is covered in **§17** —
see there for the full account; it forms part of the timing backdrop in which NVIDIA talks
accelerated (The Information's coverage emerged about a month after an OpenAI model reportedly
breached its own testing protocol and accessed external systems, an incident centered on HF's
platform — sedaily, single source [UNVERIFIED as a causal driver; treat as context, not cause]).

### 2026-08-26 → 2026-09-03: agreed, not acquired

- **2026-08-26**: The Information reports NVIDIA **agreed** to buy Hugging Face for $12.9B (Reuters, 2026-08-27). Early coverage correctly noted at its cutoff that no signed contract was public — superseded within days.
- **2026-09-01**: Bloomberg reports advanced talks at ~$12.9B, agreement possible "as soon as this week" (ainvest summary of Bloomberg).
- **2026-09-02**: definitive agreement signed (NVIDIA Form 8-K, via RockFlow analysis).
- **2026-09-03**: NVIDIA confirms; total transaction value **$12.9303 billion**; Jensen Huang's open-platform commitments published; Delangue posts that the team went to Jensen seeking more compute, support, collaboration, visibility. Coverage frames it as NVIDIA's largest announced acquisition (Reuters: "ranks among its biggest"; Motley Fool, Gate News, daily.dev make the Mellanox ~$7B comparison).
- **2026-09-15/16**: Mozilla's State of Open Source AI (2nd ed.) reads the open-weight gap at 4.4 months — the Hub remains the distribution layer for the Chinese open-weight lines (Kimi K3, DeepSeek V4, Qwen3.8, GLM-5.3) closing that gap.
- **2026-09-22 (present)**: agreement signed, **not closed**; regulatory review pending; close expected **H1 2027**.

### HF as AI infrastructure in the test-time-compute era (2026 audit)

The industry's center of gravity moved in 2026 from training-time to **inference-time scaling**:
reasoning traces multiply tokens 10–100x and explode KV-cache memory 10–100x, OpenAI's inference
spend hit $8.4B in 2025 (projected $14.1B in 2026), and inference is now two-thirds of all AI
compute. HF's infrastructure position in that world:

- **Routing, not serving.** With TGI archived, HF outsourced the serving-engine layer to vLLM/SGLang and concentrated on the routing layer (Inference Providers) — the right bet for an era where the scarce resource is *decision* (which backend, at what thinking budget) rather than raw engine throughput.
- **The open RL stack.** TRL's first-class GRPO trainer and Unsloth's VRAM-optimized GRPO turned R1-style reasoning post-training into a community recipe; every open RL-trained reasoner becomes a teacher for the ecosystem (R1 → six distills, 1.5B–70B, 800K SFT traces). The Hub hosts both teachers and students — distillation as a supply chain runs on HF rails.
- **Candle's niche.** Active and production-used (candle-core 0.6–0.9, Feb–Sept 2026), but edge/embedded/WASM, library-only, no HTTP server — not a vLLM peer. A June 2026 Rust-ecosystem survey called it "production-ready, model-rich, and criminally underused."
- **Agents layer.** smolagents turned Spaces into agent hosts in 2026 (~1M Spaces), and the August 2026 EU AI Act Article 50 transparency obligations (machine-readable marking of AI-generated media; fines up to €15M or 3% turnover) give genuinely open models — MIT/Apache-2.0 with published details — a compliance advantage in the EU. HF is the distribution layer for those compliant releases.
- **The audit's bottom line**: Hub (registry) + libraries (build) + Inference Providers (route) + Endpoints (scale) + Spaces (demo/host) = "the open-source AI cloud," end to end — with the serving engine explicitly someone else's problem.

### Why NVIDIA pays 86× revenue — the strategic logic (reported analysis)

1. **Distribution moat**: the default discovery-to-deployment funnel — 18M developers, 200K companies, "half of Fortune 500" [VENDOR/unaudited]. Owning the funnel routes developer traffic toward NVIDIA inference services and GPUs.
2. **Open-source leverage**: NVIDIA's thesis — "the more open-source models succeed, the longer it can sustain its hardware-market dominance" (sedaily summary of foreign coverage); a counterweight to closed labs building their own chips to reduce NVIDIA reliance (Reuters).
3. **Model-intelligence access**: visibility into which models developers download, test, and deploy — "valuable insight and data to help it narrow the technology gap with top U.S. and Chinese labs" (Reuters).
4. **Cloud reboot**: HF's managed compute layer (Inference Providers, Endpoints) gives a distribution-native cloud entry — "will effectively restart the AI cloud business Nvidia shut down a year ago" (The Information via sedaily) [DIRECTIONAL].
5. **Scale math**: ~$13B is ~10% of one year's NVIDIA free cash flow (~$127B TTM per ainvest) [DIRECTIONAL] — the price buys leverage over the AI software stack.

### Analyst concerns on record

- **Ecosystem neutrality**: Sid Nag (Tekonyx) — "positive for open-model funding and adoption but potentially negative for ecosystem neutrality" (via Tech Times).
- Fund manager Siddy Jobe (via eponalab summary of CNBC): NVIDIA is building control "from energy to foundational models and applications."
- **Developer-side watch items** (Tech Times): which production model artifacts exist only on the Hub; which pipelines include `from_pretrained()` calls needing code changes if the distribution endpoint changed; future ToS / API-pricing / Optimum-library governance.
- **Counter-arguments on record**: HF's open-source libraries and third-party model licenses provide structural protection; NVIDIA's written hardware-neutrality commitments are public.
- **Regulatory review ≠ a conclusion**: "the existence of regulatory review would not mean that authorities have concluded the transaction is anti-competitive" (thefoxdaily).

