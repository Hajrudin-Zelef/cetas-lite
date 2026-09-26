---
id: ai-industry-kb-2026-wave6/17-other-chinese-labs/part-3
title: "§17. Other Chinese Labs (part 3)"
domain: other-chinese-labs
role: deep-dive
task: actor-profile
actors: ["Anthropic", "ByteDance", "China", "Hugging Face", "Moonshot", "OpenAI", "StepFun"]
dates: ["2025-06", "2025-09-16", "2025-11-26", "2026-02", "2026-06", "2026-07", "2026-09"]
keywords: ["agent", "agents", "apache", "benchmarks", "claude", "cost", "fp8", "gguf", "ipo", "kimi", "mcp", "moe"]
source: docs/RAG/ai-industry-knowledge-base-2026-wave6.md
source_anchor: ""
source_lines: [8359, 8385]
section: "§17. Other Chinese Labs"
delta_of: ai-industry-kb-2026
sha256: 510fe6278f7f6f07f9f611d030b91ccf88a30de5d9113e130fc70dde14af7763
---

# §17. Other Chinese Labs (part 3)

34. Volcano Engine president Tan Dai demonstrated Doubao 2.1 Pro executing a chip-design RTL engineering process — simulation, testing, inspection — over nine iterations in about 18 hours [SECONDARY]. (S23 — single secondary coverage)
35. Average daily token call volume for the Doubao large model surpassed 180 trillion as of June 2026, over tenfold growth in a year; IDC reports Volcano Engine holds 49.5% of China's public cloud MaaS market [SECONDARY]. (S23 — single secondary coverage)
36. In September 2026 Volcano Engine updated Doubao 2.1 Pro to a 0915 revision: multimodal understanding improved for video reasoning and 3D object recognition, image/video reasoning token use fell more than 30% versus the prior generation, and the model reportedly dispatched multiple sub-agents in parallel — repairing 83% of 1,000 historical issues on the ~387K-line Luanti game repository in ~36 hours [VENDOR]. (S24, S84)
37. ByteDance launched Doubao-Seed-Code as an AI coding agent at RMB 9.9 for the first month (~63% below market average), integrated with the Trae app, processing up to 256,000 words per query and scoring 78.8% on SWE-Bench Verified [SECONDARY]. (S25, S85)
38. The same event cycle saw Seedance 2.5 video generation (early July release), Seedream 5.0 Pro image generation, and Seed-Audio 1.0 opening for invite-only testing [SECONDARY]. (S21, S23)
39. Doubao's realtime voice model is described by ByteDance as a truly end-to-end speech system primarily for Chinese contexts, serving hundreds of millions of users; external testing found users more satisfied than with GPT-4o on voice naturalness — vendor framing [VENDOR]. (S26 — single vendor source coverage)
40. StepFun released Step-3.7-Flash on May 28–29, 2026 (dates conflict across sources — preserved as contradiction): a 198B-parameter sparse MoE vision-language model activating ~11B parameters per token, composed of a 196B-parameter language backbone plus a 1.8B-parameter vision encoder [SECONDARY]. (S27, S28)
41. Step 3.7 Flash supports 256K-token context, three selectable reasoning levels (low/medium/high), and throughput up to 400 tokens/second [SECONDARY]. (S27, S28)
42. Native multimodal support (images, GUIs, documents) is new in 3.7 — Step 3.5 Flash was text-only; the DataLLM Lab review independently notes "native vision is a key differentiator" [SECONDARY]. (S28, S31)
43. StepFun-reported benchmarks: SimpleVQA (Search) 79.2 (first place), V* (Python) 95.3, ClawEval-1.1 67.1 (vs 59.8 second place), SWE-Bench PRO 56.3 (second place), Toolathlon 49.5, HLE w/Tool 48.1, Terminal-Bench 2.1 59.5, GDPVal-AA 45.8 — vendor figures, not independent [VENDOR]. (S27, S30)
44. Advisor Mode scores 76.3–76.5% on SWE-Bench Verified at $0.19/task versus Claude Opus 4.6 at $1.76 [SECONDARY]. (S28, S27)
45. Cross-harness coding variance narrowed from 43–73% (Step 3.5 Flash) to 64.5–71.5% (Step 3.7 Flash), with the vendor's per-harness table reproduced by secondary coverage [VENDOR]. (S28, S81)
46. An independent review by DataLLM Lab tested Step 3.7 Flash on nine coding tasks: 8/9 pass (a notch behind 9/9 models), real per-task cost ~$2.66/1,000 versus Opus 4.6's $4.05 (~1.5x cheaper), concluding "sticker price ≠ spend" — per-token 1/9-cost claims do not translate to per-task savings [SECONDARY]. (S31 — single secondary coverage)
47. Step 3.7 Flash pricing: $0.20/M input (cache miss), $0.04/M input (cache hit), $1.15/M output — corroborated by StepFun's own FP8 model card on Hugging Face [VENDOR]. (S28, S76)
48. Step 3.7 Flash is Apache 2.0 on Hugging Face (confirmed by the DataLLM Lab review: "Apache 2.0 on Hugging Face (with GGUF)") and listed as compatible with Claude Code, KiloCode, Hermes Agent and OpenClaw harnesses per announcement coverage [SECONDARY]. (S28, S31)
49. Hunyuan3D 3.0 was announced September 16, 2025 at the Tencent Global Digital Ecosystem Conference in Shenzhen — a 2025 release, not a 2026 launch; it must not be presented as 2026 [SECONDARY]. (S32, S33)
50. Hunyuan3D 3.0 uses a 3D-DiT "hierarchical sculpting" model achieving 1536³ geometric resolution and 3.6 billion voxels, claimed 3x accuracy versus prior [SECONDARY]. (S32, S33)
51. Hunyuan3D 3.0 is closed: hosted API only, not open weights — community research explicitly flags blogs claiming "2.5 is open source" as wrong [COMMUNITY]. (S34, S35)
52. Hunyuan3D 3.1 followed around February 2026 as a closed hosted refinement of 3.0 (Replicate listing `tencent/hunyuan-3d-3.1`), adding multi-view input types (top/bottom/left_front/right_front); Sketch mode remains 3.0-only [COMMUNITY]. (S35 — single community source coverage)
53. The open-weights Hunyuan3D line stalls at 2.1 (June 2025, PBR pipeline); open 2025–2026 artefacts include Hunyuan3D-Omni (Sept 2025, open weights, "ControlNet of 3D" on 2.1) and HY3D-Bench (Feb 4, 2026 — open dataset of 252K+ watertight meshes / ~11 TB, 240K+ part-decomposed objects / ~5.0 TB, 125K+ synthetic objects across 1,252 categories / ~6.5 TB) [COMMUNITY]. (S35, S36)
54. Tencent Cloud's international station and Hunyuan3D API for overseas users launched November 26, 2025; the Hunyuan 3D Engine global launch offers 20 free generations/day for individuals and 200 free credits for Tencent Cloud API users [COMMUNITY]. (S35, S33)
55. No `Tencent-Hunyuan/Hunyuan3D-3.0` repository exists on the Tencent-Hunyuan GitHub org — the 3.x specs come from Tencent's own X/press posts and secondary coverage, with no arXiv technical report found [UNVERIFIED]. (S35)
56. HunyuanImage 3.0 and Hy3 are separate products from Hunyuan3D and must not be conflated; likewise Hunyuan3D 3.0 ≠ Hunyuan3D-World 1.0-Lite (Aug 2025, FP8, <17 GB VRAM) [COMMUNITY]. (S35, S36)
57. ERNIE-4.5-VL-28B-A3B-Thinking is a 2025-era model and must not be framed as a 2026 launch (standing correction carried from a prior wave — no new source asserted) [DIRECTIONAL].
58. Moonshot's Kimi K3 launched in July 2026 and the company is reportedly seeking a Hong Kong IPO under confidential terms, per NDTV reporting — single-source, unconfirmed by Moonshot [UNVERIFIED]. (S38)
59. Midscene v1.10 adopted Doubao-Seed-2.1-turbo as its recommended model and formally retired MCP server packages (final MCP version pinned to 1.9.8), showing Doubao-Seed ecosystem penetration in agent tooling [COMMUNITY]. (S39 — single community source coverage)

