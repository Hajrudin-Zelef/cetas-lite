---
id: ai-industry-kb-2026/16-hugging-face/part-2
title: "16. Hugging Face (part 2)"
domain: hugging-face
role: deep-dive
task: platform
actors: ["China", "DeepSeek", "Groq", "Hugging Face", "Moonshot", "Nvidia", "OpenAI", "SGLang", "United States", "Z.ai", "vLLM"]
dates: ["2025-04", "2025-07", "2025-08-05", "2025-12", "2026-01", "2026-02-20", "2026-03", "2026-03-21", "2026-04-23", "2026-04-28", "2026-07", "2026-07-27", "2026-08-05", "2026-08-26", "2026-08-27", "2026-09-01", "2026-09-02", "2026-09-03", "2026-09-14", "2026-09-15", "2026-09-16", "2026-09-22"]
keywords: ["acquisition", "chatgpt", "compute", "deepseek", "distribution", "glm", "humanoid", "incident", "inference", "kimi", "llama", "llama.cpp"]
source: docs/RAG/ai-industry-knowledge-base-2026.md
source_anchor: ""
source_lines: [7897, 7920]
section: "16. Hugging Face"
sha256: 16e732ff1c6d1234c946dc598dc47a4c18be3e6dc72ed0f5c01e6d7d216d56cc
---

# 16. Hugging Face (part 2)

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

