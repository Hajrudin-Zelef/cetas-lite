---
id: ai-industry-kb-2026/16-hugging-face/figures-and-metrics
title: "Figures and metrics"
domain: hugging-face
role: deep-dive
task: platform
actors: ["AWS", "Cerebras", "Cohere", "Fireworks AI", "Google", "Groq", "Hugging Face", "Nscale", "Nvidia", "OpenAI"]
dates: ["2019-03-11", "2020-04-27", "2025-07", "2025-12", "2026-04-23", "2026-04-27", "2026-04-28", "2026-05", "2026-06", "2026-07", "2026-08", "2026-09", "2026-09-01", "2026-09-03", "2026-09-22"]
keywords: ["acquisition", "apache", "attribution", "cohere", "embeddings", "gpu", "humanoid", "inference", "nvidia", "pricing", "research", "revenue"]
source: docs/RAG/ai-industry-knowledge-base-2026.md
source_anchor: ""
source_lines: [7921, 8011]
section: "16. Hugging Face"
sha256: 9a4ca5e307c38c1efe195395b2296ae2aa39555134f14e0e6200e03290cc7705
---

# Figures and metrics

## Figures and metrics

### Deal terms (per NVIDIA SEC Form 8-K, September 2026)

| Component | Value |
|---|---|
| Base consideration to stockholders | ~$11.9 billion |
| Equity-based employee retention program (up to) | ~$1.0 billion |
| **Total transaction value** | **$12.9303 billion** |
| Employees | ~750 |
| Largest prior NVIDIA acquisition (closed) | Mellanox — agreed 2019-03-11 at ~$6.9B enterprise value, closed 2020-04-27 at ~$7B transaction value |
| Deal size vs Mellanox | ~1.9× (largest announced acquisition if completed) |
| HF annualized revenue (Sacra, August 2026) | ~$150M [DIRECTIONAL, unaudited estimate] |
| Implied revenue multiple | ~86× [DIRECTIONAL] |
| Last venture round | 2023-08: $235M led by Salesforce Ventures at $4.5B valuation (Google, Amazon, IBM, NVIDIA participated) |
| Total raised to date | $395M+ (Crunchbase) |
| Deal price vs last round valuation | ~2.9× ($4.5B → $12.93B) |
| Deal price vs rejected 2025 $7B mark | ~1.8× |
| Status as of 2026-09-22 | Definitive agreement signed; **not closed**; regulatory approval pending |
| Expected close | **H1 2027** (no closing date disclosed) |

### Platform scale — source-by-source evolution table

| Date | Models | Datasets | Spaces/Apps | Source |
|---|---|---|---|---|
| May 2026 | 1.8M+ | 450K | 720K Spaces | ayinedjimi-consultants entity profile |
| May 2026 | 2.4M+ | 730K+ | ~1M Spaces | metacto guide (updated May 2026) |
| June 2026 | 2M+ | 500K+ | 1M+ Spaces | Medium (johirbuet) |
| ~July 2026 | ~3M public models | ~1M datasets | new repo every ~7s | fintech-radar, citing Delangue [VENDOR, self-reported/unaudited] |
| Sept 2026 (Wave 1 synthesis) | 2.4M+ public models | 730k+ | ~1M Spaces | HF central-hub section; DOE filing: 11M users, 50,000+ organizations |
| **2026-09-03** | **3M** | **500K ("half a million")** | **1M applications** | TechCrunch; Intelligent CIO; Reuters; TBPN |
| 2026-09-03 (deal coverage) | 3M models | — | 1M apps; **18M developers; 200K companies** | TechCrunch; Intelligent CIO; podcast summaries |

### Datasets-count methodology note (why the figures disagree)

The "927K+ datasets" figure is **[UNVERIFIED — no supporting source found]** and is **not used** in
this consolidation. The three live figures reflect different counting methods: **500K** (September
3, 2026 deal press — likely public, curated/counted-by-press repositories); **730K+** (May 2026
practitioner guides — likely total visible repositories); **~1M** (Delangue, July 2026 —
self-reported, unaudited [VENDOR], likely including private/gated repos and de-duplicated-at-source
counts). Consolidation rule: cite **500K with the Sept 3 deal-coverage attribution**, note the
730K–1M vendor/guide range, and never the unsupported 927K+.

### Inference economics layer (2026 list pricing)

| Product | Detail |
|---|---|
| Inference Providers router | router.huggingface.co — OpenAI-compatible `POST /v1/chat/completions`, plus embeddings, image, video, audio routes |
| Backend partners (15+, 2026) | Cerebras, Cohere, Fal AI, Featherless AI, Fireworks, Groq, HF Inference, Hyperbolic, Novita, Nscale, Replicate, SambaNova, Together, WaveSpeedAI |
| Provider selection | Automatic `:fastest` / `:cheapest` suffixes; failover between providers; **no markup** on provider costs |
| Auth model | One HF token, or your own provider key swapped transparently |
| Free credits | **$0.10 free tier / $2.00 PRO**, monthly; pay-as-you-go past that |
| Legacy serverless Inference API | Renamed **hf-inference**; since July 2025 mostly CPU inference (embeddings, reranking, classification, small/legacy LLMs); serverless routes limited to **<10GB models** on some paths |
| Inference Endpoints (dedicated GPU, $/hr) | T4 $0.50; L4 $0.80; A10G $1.00; L40S $1.80; A100 80GB $2.50; H200 $5.00; H100 (GCP) $10.00; TPU v5e from $1.20 — multi-GPU scales linearly (4× A100 = $10/hr; 8× H200 = $40/hr) |
| Platform tiers | Free / **Pro $9/mo** / Enterprise Hub (SSO, audit, compliance) |
| Transformers library | **113M+ monthly downloads** (late 2025) |
| Datasets library | 730k+ datasets, 500k+ covering **8,000+ languages** (Wave 1) |

### Robotics line (specs)

| Milestone | Detail |
|---|---|
| LeRobot launch | 2024 — open models, datasets, tools for robotics; 12K–24K GitHub stars across 2026 (growth over time, not contradiction) |
| Pollen Robotics acquisition | 2026-04 — Bordeaux, France; original Reachy; Reachy 2 ~$70K; 100+ units in 20+ countries; Orbita 7-DOF joints; Apache-2.0 software |
| Reachy Mini debut | **2025-07** — 3D-printed prototype; Lite $299 / Full-Wireless $449; ~11 in (28 cm) tall; 15 pre-installed demos (facial recognition, hand tracking, voice-activated AI conversations); integrates with HF Hub models/datasets |
| Reachy Mini sales (trade press) | $500K in first 24h (Digital Trends) / $1M in first 5 days (one source) [COMMUNITY — not HF financials] |
| Mass production | Late 2025 — 3,000+ pre-orders in a week; Seeed Studio partnership |
| CES 2026 | 2026-01 — 3,000 units shipped; **Jensen Huang highlighted Reachy Mini in his CES keynote** ("AI robots for the masses") |
| 2026-09-01 unveiling | **HopeJR** — full-size humanoid, **66 actuated degrees of freedom**, walking + arm manipulation, **~$3,000**, open-source; **refreshed Reachy Mini** — $250–$300, shipping by year-end, open-source; enabled by the Pollen acquisition |

### CVE-2026-25874 — fact sheet

| Field | Value |
|---|---|
| CVE ID | CVE-2026-25874 |
| Severity | CVSS **9.8** (NVD 3.1) / **9.3** (CVSS 4.0) — the two numbers are the v3.1 vs v4.0 scoring difference, not a contradiction |
| Component | LeRobot **async inference PolicyServer** — `pickle.loads()` on gRPC input with `add_insecure_port()` (no authentication, no TLS) |
| Attack | Unauthenticated, network-reachable attacker sends crafted pickle payload → **arbitrary OS command execution** on the GPU-backed inference host |
| Affected versions | **v0.4.3 verified vulnerable; NVD scopes impact through v0.5.1** |
| NVD publication | **2026-04-23** |
| Independent discovery 1 | Researcher 'chenpinji', **December 2025** — HF responded the implementation was "experimental" and needed "near-complete refactoring" (Lyrie) |
| Independent discovery 2 + PoC | Valentin Lobstein ("chocapikk"), PoC published **2026-04-28** |
| Press coverage | 2026-04-27/28 (The Hacker News, Resecurity, Cloud Security Alliance); Lyrie Research two-part analysis 2026-04-28/29 |
| Patch status | **Still unpatched as of September 2026**; fix planned in **0.6.0** (pickle → safetensors/JSON) |
| Mitigations documented | Replace pickle with JSON/protobuf/safetensors; `add_secure_port()` + TLS; gRPC token-auth interceptor |
| The safetensors irony | Hugging Face created **SafeTensors** specifically to prevent pickle-deserialization attacks — yet LeRobot developers used `pickle.loads()` on network input with a **`# nosec` comment to silence security scanners** (Lyrie Research) |
| Physical-safety dimension | RCE reaches the robot's joint-control path — blast radius extends from data theft to physical sabotage of hardware |
| Blast radius amplification | LeRobot's popularity (12K–24K GitHub stars, 2026) + vulnerable default config shipped in releases = any exposed deployment is effectively a backdoor (Lyrie Research) |

### Model-count velocity and curation signal

