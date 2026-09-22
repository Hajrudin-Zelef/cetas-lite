---
id: ai-industry-kb-2026/16-hugging-face/overview
title: "16. Hugging Face"
domain: hugging-face
role: deep-dive
task: platform
actors: ["AMD", "AWS", "Alibaba", "China", "DeepSeek", "EU", "Google", "Groq", "Hugging Face", "Intel", "Moonshot", "Nvidia", "OpenAI", "SGLang", "United States", "Unsloth", "Z.ai", "vLLM"]
dates: ["2025-04", "2025-07", "2025-08-05", "2025-12", "2026-01", "2026-02-20", "2026-03", "2026-03-21", "2026-04-23", "2026-04-24", "2026-04-28", "2026-05", "2026-07", "2026-07-27", "2026-08", "2026-08-05", "2026-08-12", "2026-08-26", "2026-08-27", "2026-09", "2026-09-01", "2026-09-02", "2026-09-03", "2026-09-14", "2026-09-15", "2026-09-16", "2026-09-22"]
keywords: ["acquisition", "amd", "antitrust", "aws", "chatgpt", "compute", "deepseek", "distribution", "gguf", "glm", "humanoid", "incident"]
source: docs/RAG/ai-industry-knowledge-base-2026.md
source_anchor: ""
source_lines: [7830, 7920]
section: "16. Hugging Face"
sha256: 60df84e958d93088d8fab505b87ec9cf7341abd7ef85ade15d6eaa91a3702067
---

# 16. Hugging Face
Keywords: Hugging Face, NVIDIA agreed to acquire, $12.9303 billion, Jensen Huang, Form 8-K, definitive agreement, regulatory approval, H1 2027 close, Clément Delangue, model hub scale, 3 million models, 500K datasets, Reachy Mini, Pollen Robotics, HopeJR, LeRobot, CVE-2026-25874, pickle deserialization RCE, CVSS 9.8, safetensors irony, Inference Providers, TGI maintenance mode, Candle,
smolagents, ggml.ai acquisition, EU AI Act open-source exemption

## Summary

Hugging Face — "the GitHub of AI," founded 2016 in New York by Clément Delangue, Julien Chaumond and
Thomas Wolf — became the most strategically valuable infrastructure asset in the 2026 AI market
without building a single frontier model. Between February and September 2026, the company (a)
announced the acquisition of **ggml.ai** (2026-02-20), (b) absorbed a critical unpatched security
flaw (**CVE-2026-25874**, unauthenticated remote code execution via pickle deserialization in the
LeRobot async PolicyServer, CVSS 9.8, published 2026-04-23), (c) unveiled a second robotics wave on
**2026-09-01** (the full-size humanoid **HopeJR** plus a refreshed Reachy Mini), and (d) entered a
**definitive agreement, signed 2026-09-02 and confirmed 2026-09-03**, under which **NVIDIA agreed to
acquire it for a total transaction value of $12.9303 billion** — ~$11.9B payable to stockholders
plus up to $1.0B in an equity-based employee retention program. The transaction is **not closed**:
it is subject to customary closing conditions and regulatory review, with closing expected in **H1
2027**. If completed, it would be NVIDIA's largest acquisition ever — nearly double the $6.9B
Mellanox deal — at roughly **86× annualized revenue** (~$150M annualized revenue, Sacra estimate,
August 2026) [DIRECTIONAL].

Wording discipline matters throughout this section: NVIDIA **agreed to acquire** Hugging Face; it
has not acquired it. Early coverage cautioned that, at its cutoff, no signed contract was public
(The IT Guys, 2026-08-27) — that caution was superseded by NVIDIA's SEC Form 8-K documenting a
signed definitive agreement, and by NVIDIA's public confirmation on 2026-09-03. The consolidation
resolves the record toward **agreed-to-acquire** on the strength of the primary sources, never
toward completed purchase.

Platform scale claims require the same discipline. The Hub reached **3 million models** by September
2026 (deal coverage, 2026-09-03; 2.4M+ in May 2026 guides), ~1M applications/Spaces, 18M developers
and 200K companies [VENDOR/self-reported via deal coverage]. The **"927K+ datasets" figure is
unverified — no source supports it and this consolidation does not use it.** Contemporaneous sources
state ~500K datasets (TechCrunch; Intelligent CIO; TBPN, all September 3, 2026), while HF's own
figures range from 730K+ (May 2026 guides) to ~1M (Delangue, July 2026, self-reported/unaudited)
[VENDOR]; the variance reflects methodology differences (public vs total repositories, raw vs
de-duplicated counts) and is documented in the figures section below.

On infrastructure, 2026 delivered a hard correction: Hugging Face's own serving stack **TGI
(text-generation-inference) entered maintenance mode and its repository was archived on
2026-03-21**, with HF itself recommending migration to vLLM or SGLang. What grew into HF's
orchestration layer instead is **Inference Providers** — a unified router at router.huggingface.co
across 15+ backend partners with automatic `:fastest`/`:cheapest` selection, zero markup, and
one-token access. Candle (HF's pure-Rust ML framework) is active and production-used in 2026 but
remains a niche (edge/embedded/WASM), not a serving-engine peer. The full TRL/GRPO open RL stack
(the commoditized reasoning-training pipeline) is the quieter infrastructure story of the year:
R1-style RL post-training became a pip-installable recipe, collapsing the barrier to producing
reasoning models.

Regulatory review of the NVIDIA agreement is a **live watch item as of 2026-09-22**: US and EU
scrutiny is flagged in coverage, with the novel antitrust question being whether NVIDIA can use
ownership of the dominant AI developer platform to strengthen its hardware position — across access,
interoperability, performance optimization, commercial terms, cloud relationships, and support for
competing accelerators (AMD Instinct, Intel Gaudi, Google TPUs, AWS Trainium).

The Hub's 2026 release record is the concrete evidence behind the valuation. Every major open-weight
release landed on Hugging Face within days of announcement: DeepSeek's V4 family (April 24, 2026,
MIT), Kimi K3 (weights July 27, 2026, Modified MIT — "appeared on Hugging Face at 2:00 AM Beijing
time"), Qwen3.8-Max weights (August 12, 2026), GLM-5.3-Flash (August 26, 2026, MIT), and OpenAI's
gpt-oss (4.3M+ downloads). The Unsloth quantization catalog (GGUF/NVFP4 for Qwen3.x, GLM-5.3, Kimi,
DeepSeek-V4, Gemma 4, Nemotron) is distributed primarily through the Hub, as is the GGUF standard's
home alongside llama.cpp and Ollama. Epoch's curated snapshot (1,340 open-weight models) and ATOM
(~1,500 mainline models) track the *selectable* universe against the Hub's raw 2.4M–3M count — a
reminder that raw repository count includes derivatives, conversions, and fine-tunes, and that model
selection is now the hard problem.

## Key dated facts

- **2026-02-20** — Hugging Face announces the acquisition of **ggml.ai** (the company behind the GGML machine-learning tensor library / llama.cpp ecosystem adjacency), consolidating control over the open-weight inference-format layer.
- **2026-03-21** — Hugging Face archives the **text-generation-inference (TGI)** GitHub repository (read-only); TGI enters maintenance mode (minor bug fixes and docs only); HF officially recommends migration to vLLM, SGLang, or local engines.
- **2026-04-23** — NVD publishes **CVE-2026-25874**: unauthenticated RCE in Hugging Face LeRobot's async inference PolicyServer via insecure pickle deserialization over exposed gRPC endpoints. CVSS 9.8 (v3.1) / 9.3 (v4.0). Affected versions: **v0.4.3 through v0.5.1**; **no fixed release as of 2026-09-22**; fix planned in 0.6.0 (pickle replaced with safetensors/JSON).
- **2026-04-28** — Researcher Valentin Lobstein ("chocapikk") publishes a proof of concept for CVE-2026-25874; independent discovery followed a December 2025 report by 'chenpinji', which Hugging Face had dismissed as "experimental."
- **2026-08-26** — The Information first reports that **NVIDIA agreed to acquire Hugging Face for ~$12.9B** (via Reuters, 2026-08-27).
- **2026-09-01** — Bloomberg reports advanced talks at ~$12.9B, agreement possible "as soon as this week"; Hugging Face unveils **HopeJR** (full-size humanoid, 66 actuated degrees of freedom, ~$3,000) and a refreshed **Reachy Mini** ($250–$300, shipping by year-end), both open-source.
- **2026-09-02** — **Definitive agreement signed**, per NVIDIA's SEC Form 8-K (RockFlow analysis).
- **2026-09-03** — NVIDIA confirms: total transaction value **$12.9303 billion** — ~$11.9B base consideration to Hugging Face stockholders plus up to ~$1.0B equity-based employee retention (company: ~750 employees). Coverage: TechCrunch, Reuters, Intelligent CIO, Motley Fool. Jensen Huang publishes open-platform commitments: "Hugging Face will remain an open platform for the entire AI ecosystem... Nvidia compute will not be required to build on or deploy through Hugging Face."
- **2026-09-15** — Mozilla's State of Open Source AI (2nd ed., previewed September 16, 2026) puts Chinese open-weight models 4.4 months behind the US frontier reading off the Artificial Analysis Intelligence Index — with the Hub as the distribution layer for those models.
- **2026-09-22** — As of this date: agreement signed but **not closed**; regulatory approval pending; closing expected **H1 2027**; CVE-2026-25874 **still unpatched**; no closing date disclosed.
- **July 2025** (correction, not 2026) — Original **Reachy Mini** debut (3D-printed prototype; Lite $299 / Full-Wireless $449); 3,000 units shipped by CES 2026 (January 2026), where Jensen Huang highlighted it in his keynote.
- **2026-05 (metacto guide, updated)** — Hub counts cited as 2.4M+ models, 730K+ datasets, ~1M Spaces — the pre-deal baseline.
- **2026-07-27** — Kimi K3 open weights (2.8T MoE, Modified MIT) land on the Hub — "appeared on Hugging Face at 2:00 AM Beijing time"; the largest open-weight release ever at that point.
- **2026-08-05** — OpenAI's gpt-oss turns one year old (released 2025-08-05); **4.3M+ downloads on Hugging Face** — among the most-downloaded open-weight models ever, and a marker of the Hub's release-day gravitational pull.
- **2026-08-26** — GLM-5.3-Flash weights (320B MoE, MIT) published on Hugging Face — the second 300B+ MIT multimodal checkpoint in one week (with DeepSeek V4-Flash-Vision-Exp).
- **2026-09-14** — ChatGPT Plus/Pro lose automatic Instant→Thinking switching (thinking becomes a manual dial) — part of the test-time-compute era context in which HF's inference routing layer matters (see §"Timeline and context").
- **April 2025** — Hugging Face acquired **Pollen Robotics** (Bordeaux, France; original Reachy robot; Reachy 2 at ~$70K; 100+ units across 20+ countries), the enabling deal behind the robotics line.
- **2024** — Hugging Face launched **LeRobot** (open AI models, datasets, and tools for robotics; 12K–24K GitHub stars across 2026 depending on source/date).
- **March 2026** — Hugging Face files a DOE Genesis Mission AI Workforce RFI citing **11M users** and **50,000+ organizations** [VENDOR].
- **Late 2025 (reported January 2026)** — FT: Hugging Face **rejected a $500M NVIDIA investment offer** in late 2025 that would have valued it at $7B, reportedly to avoid any single investor's outsized influence over platform decisions (reported via Reuters, 2026-08-27).
- **December 2025 (reported)** — NVIDIA's ~$20B non-exclusive Groq licensing deal; NVIDIA disclosed $18B committed to equity investments for the rest of FY2027 — the capital context in which the $12.93B agreement lands.
- The July 2026 autonomous-breach incident targeting Hugging Face's platform is documented in **§17** (not repeated here).
- Robotics deployment detail beyond the chronology above lives in **§19** (not repeated here).

