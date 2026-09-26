---
id: ai-industry-kb-2026/16-hugging-face/main-actors
title: "Main actors"
domain: hugging-face
role: deep-dive
task: platform
actors: ["AWS", "EU", "Google", "Groq", "Hugging Face", "Nvidia", "United States"]
dates: ["2023-08", "2025-04", "2025-07", "2025-12", "2026-01", "2026-02-20", "2026-03-21", "2026-04-27", "2026-04-28", "2026-06", "2026-08", "2026-08-26", "2026-09", "2026-09-01", "2026-09-02", "2026-09-03"]
keywords: ["acquisition", "advisory", "agent", "apache", "arr", "compute", "distribution", "energy", "funding", "gguf", "inference", "llama"]
source: docs/RAG/ai-industry-knowledge-base-2026.md
source_anchor: ""
source_lines: [8039, 8113]
section: "16. Hugging Face"
sha256: 9238904f407194c5df71a0e4c22a75840d0678f08bc50bc3c68cb19a2de8ec39
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

