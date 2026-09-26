---
id: ai-industry-kb-2026-wave6/17-other-chinese-labs/part-7
title: "§17. Other Chinese Labs (part 7)"
domain: other-chinese-labs
role: deep-dive
task: actor-profile
actors: ["Falcon", "Huawei", "MiniMax", "Moonshot", "OpenAI", "OpenRouter", "Poolside", "Z.ai"]
dates: []
keywords: ["ascend", "benchmark", "benchmarks", "cost", "distribution", "gguf", "glm", "gpt-5.6", "ipo", "kimi", "latency", "license"]
source: docs/RAG/ai-industry-knowledge-base-2026-wave6.md
source_anchor: ""
source_lines: [8521, 8554]
section: "§17. Other Chinese Labs"
delta_of: ai-industry-kb-2026
sha256: 90b8609a7a3d708e65a4c65b1bd26086e0ecc7e74a59077841b010adef8d0132
---

# §17. Other Chinese Labs (part 7)

| Hy4 Preview backbone / experts | 78 layers (1 dense + 77 MoE); 256 routed + 1 shared, top-8 | S41 (single source) | [SECONDARY] |
| Hy4 Preview MTP layer | 10B total / 0.7B active | S41 (single source) | [SECONDARY] |
| Hy4 Preview API pricing | $0.834 / $2.501 / $0.042 per M (in/out/cache-hit) | S41, S42 | [SECONDARY] |
| Hy4 Preview measured throughput | ~40 tok/s, 3.49 s avg latency (OpenRouter) | S42 (single source) | [SECONDARY] |
| Hy4 Preview BF16 weight size | ~1.56 TB | S42 (single source) | [SECONDARY] |
| Hy4 Preview GGUF builds | STQ1_0 213.66 GiB; Q4_K_M 435.20 GiB; UD-IQ1_M 219.83 GiB | S42 (single source) | [SECONDARY] |
| Hy4 Preview blind eval (vendor) | 2.99/4.00 vs GLM-5.3 2.92, Kimi K3 2.94 (163 experts, 203 tasks) | S41 (single source) | [VENDOR] |
| openPangu 2.0 training | 34T tokens on Ascend NPUs | S45 (single source) | [SECONDARY] |
| openPangu Flash CUDA build (community) | ~56.9 GB resident memory | S78 (single source) | [COMMUNITY] |
| Falcon H1R 7B | DeepConf reasoning model | S47, S48 | [SECONDARY] |
| Falcon-H1 Arabic sizes | 3B / 7B / 34B; up to 256K context | S48 (single source) | [VENDOR] |
| Falcon-H1 Arabic OALL (vendor) | 61.87% / 71.47% / 75.36% (3B/7B/34B) | S48 (single source) | [VENDOR] |
| Poolside deal new terms | $12B pre-money (for $1B); offers to 109/~115 staff; co-founders stay | S50, S51 | [SECONDARY] |
| Poolside fee distribution | to existing investors by end of 2027 | S51 (single source) | [SECONDARY] |
| Poolside raised since 2023 | $1.6B | S51 (single source) | [SECONDARY] |
| Poolside Laguna XS 2.1 / M.1 | 33B/3B, 225B/23B (OpenMDW-1.1) | S51 (single source) | [SECONDARY] |
| Poolside Laguna reported benchmarks | M.1 72.5% SWE-bench Verified; XS 2.1 70.9%; S 2.1 70.2% Terminal-Bench 2.1 | S51 (single source) | [VENDOR] |
| ERNIE 5.0 architecture | ~2.4T native-multimodal MoE; <3% active/query | S55 (single source) | [SECONDARY] |
| ERNIE 5.1 claimed efficiency | 94% lower pre-training cost; total 1/3, active 1/2; −35% latency | S55, S56 | [VENDOR] |
| ERNIE 5.1 third-source pricing | $3/M in, $12/M out (Qianfan API) | S57 (single source) | [SECONDARY] |
| Ling-3.0-Flash license (community) | MIT attributed (was [UNVERIFIED] in base) | S59, S60 | [COMMUNITY] |
| Ling-3.0-Flash context ceiling | scalable to 1M (native 256K) | S59 (single source) | [COMMUNITY] |
| Ling-3.0-Flash AA Index / Omniscience | 38; hallucination 97% → 44% | S58 (single source) | [SECONDARY] |
| Ling-3.0-Flash API pricing | $0.075/M in, $0.22/M out | S59 (single source) | [COMMUNITY] |
| Ling FinFIRST benchmark | 123 tasks / 701 criteria / 12,300 rubric points | S61 (single source) | [VENDOR] |
| Ling-3.0-tiny | 7.9B / 1.3B, fully local | S61 (single source) | [VENDOR] |
| Step 5 Preview active density | ~4.5% (27B/600B) | S63 (single source) | [SECONDARY] |
| Step 5 Preview AA Index | 44; ~99.8 tok/s; $0.71/index task | S63, S64 | [SECONDARY] |
| Step 5 Preview vs GPT-5.6 Sol | 44 vs 47 (max) / 44 (matched speed); 7.4× cheaper output | S66 (single source) | [SECONDARY] |
| MiniMax IPO | HK$4.8–5.54B raised; closed HK$345 (+109.1–109.9%) | S69, S70, S71 | [SECONDARY] |
| MiniMax IPO oversubscription | 1,837× public, 37× international | S71 (single source) | [SECONDARY] |
| Zhipu IPO | HK$4.35B (~$552M); +3.3% open, +13% close day one | S69, S72, S74 | [SECONDARY] |
| MiniMax Stock Connect (Aug 2026) | +22.8% Aug 6; HK$10.6B mainland inflow; 9.7% stake | S73 (single source) | [SECONDARY] |

