---
id: collect-261001-general-networking/general-networking/etape5-tracka-nvidia-4
title: "Step 5 — Track A: NVIDIA GPUs (B200/B300/Rubin), DGX Systems & Networking"
domain: general-networking
role: reference
task: reference
actors: ["AWS", "Cohere", "CoreWeave", "Google", "Groq", "Mistral", "Nvidia", "OpenAI", "Perplexity"]
dates: ["2026-09-22"]
keywords: ["gpu", "gpus", "nvidia", "rubin", "acquisition", "agent", "aws", "blackwell", "cohere", "compute", "cpo", "disaggregated"]
source: docs/RAG/collect-261001-general-networking/etape5_trackA_nvidia.md
source_anchor: ""
source_lines: [230, 269]
sha256: a800479dea07d7260ae57a4e0e19ae89dfd75077d742bb1a91dc4edf23b9e1a8
---

# Step 5 — Track A: NVIDIA GPUs (B200/B300/Rubin), DGX Systems & Networking

- **Vera Rubin full-stack platform:** full production confirmed; 7 chips, 5 rack-scale systems; partner availability H2 2026 [independent/secondary]
- **Feynman architecture:** first public preview — Rosa CPU, LP40 LPU, BlueField-5, CX10, Kyber co-packaged optics, NVLink 8, Spectrum-7 [secondary]
- **Groq acquisition:** NVIDIA confirmed acquisition of Groq's team and technology (LPUs integrated into Rubin/Feynman inference strategy) [secondary — https://acaderesearch.com/nvidia-gtc-2026-the-inference-inflection-vera-rubin-and-the-ai-factory-era/]
- **NemoClaw:** open-source enterprise AI agent framework (Wired leak March 10; briefings to Salesforce, Cisco, Google, Adobe, CrowdStrike) [secondary]
- **Dynamo:** inference operating system / cluster orchestration for disaggregated inference [secondary]
- **Nemotron 3:** Super/Ultra/Nano model families; Nemotron Coalition partners (Mistral, Perplexity, Cursor, Cohere) [secondary]
- **DLSS 5:** neural rendering beyond frame generation; "1,000,000× leap" in path-tracing performance claimed [vendor-reported/secondary]
- **Physical AI:** DRIVE Hyperion partnerships (Hyundai, BYD, Nissan) for Level 4 autonomy; T-Mobile AI-RAN partnership; Nokia deploying RTX PRO 4500 Blackwell Server Edition in AI-RAN base stations [secondary]
- **Spectrum-X CPO switch:** world's first mass-produced co-packaged-optics switch [secondary]
- **AWS partnership:** 1M+ NVIDIA GPUs across AWS regions (Blackwell + Rubin, RTX PRO Blackwell Server Edition, Groq 3 LPUs) [secondary]
- **Space-1 Vera Rubin:** orbital AI data centers concept [secondary — marketing announcement, treat as ambition]
- **Demand projection:** AI infrastructure demand doubled to **$1T through 2027** [vendor-reported]
- Context: keynote followed a ~14% stock pullback after a record $68.1B revenue quarter [secondary]

---

## 7. Verification log (open items)

1. **B200 usable memory** — 192 GB vs 180 GB "usable" conflict across secondary spec tables. [unverified]
2. **B300 FP8/FP4 sparse figures** — 9/18 vs 9/20 PFLOPS (B200) and ~28–30 PFLOPS sparse (B300) vary by source. [unverified]
3. **B300 FP8 compute** — 4.5 vs ~11–13.5 PFLOPS conflict. [unverified]
4. **336B-transistor Rubin figure** — from secondary sources; not verified against NVIDIA. [unverified]
5. **"H300" rebranding of Rubin R100** at AWS/Google Cloud — single secondary source. [unverified]
6. **Rubin Ultra configuration** — NVL576/576 GPUs vs 144 vertical-tray GPUs conflict. [unverified]
7. **OpenAI Q3 2026 Rubin deployment at scale** — Bloomberg via secondary relay. [unverified]
8. **DGX Station 2026 refresh** — no verified details located. [gap]
9. **DGX Cloud 2026 pricing/availability news** — no verified details located. [gap]
10. **RTX Spark launch timing** — "weeks before launch" per eTeknix Sep 21, 2026. [unverified]
11. **HBM4 sold out through 2026** — GTC 2026 secondary reporting. [unverified]
12. **CoreWeave 10× token output claim** — secondary relay of vendor/partner claim. [unverified]
13. **Space-1 orbital datacenters** — marketing concept, no deployment details. [unverified]
14. **Groq acquisition terms** — confirmed as announced but terms/price not verified in fetched sources. [unverified]

---

## 8. Collection metadata

- **Research date:** September 22, 2026
- **Sources prioritized:** NVIDIA official (docs.nvidia.com, investor.nvidia.com, datasheets), DCD, EE Times, VideoCardz, ServeTheHome-adjacent coverage, SemiAnalysis, Tom's Hardware, LMSYS, IntuitionLabs, GPUaaS, GPUSmith
- **Conventions:** all prices in USD unless noted; cloud rates are per-GPU-hour; "street" prices are retailer snapshots dated Sep 2026; sparse vs dense FLOPS always distinguished where sources allow
