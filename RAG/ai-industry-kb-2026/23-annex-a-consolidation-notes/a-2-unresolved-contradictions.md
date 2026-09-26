---
id: ai-industry-kb-2026/23-annex-a-consolidation-notes/a-2-unresolved-contradictions
title: "A.2 Unresolved contradictions"
domain: appendix
role: appendix
task: reference
actors: ["Alibaba", "Anthropic", "ByteDance", "Cerebras", "DeepSeek", "EU", "MiniMax", "Mistral", "Moonshot", "OpenAI", "United States", "Unsloth", "Z.ai"]
dates: ["2025-06-18", "2025-08", "2025-09-30", "2025-11-25", "2026-03-25", "2026-04", "2026-05-24", "2026-06", "2026-06-08", "2026-06-10", "2026-06-24", "2026-07", "2026-07-23", "2026-07-28", "2026-09-30"]
keywords: ["benchmark", "benchmarks", "chatgpt", "compute", "deepseek", "diffusion", "glm", "gpu", "ipo", "kill switch", "kimi", "license"]
source: docs/RAG/ai-industry-knowledge-base-2026.md
source_anchor: ""
source_lines: [11036, 11064]
section: "Annex A — Consolidation notes"
sha256: 7e3b2a80fc15a38c6de98a9633166b4c209f6da2ea825fa85b377aba6bc94694
---

# A.2 Unresolved contradictions

28. **MiniMax H3 vs Hailuo 3.0 vs "Hailuo 03"** — §3. One model, three names — verified across independent sources (Reuters-day coverage, FAQ, Japanese name table). Merged into a single index entry; never listed separately. Distinct from M3 (428B/23B text omni MoE, June 2026) and from H3's territorial-license restriction (excludes US/EU/UK/South Korea local deployment).
29. **Seed 2.1 Turbo (June 24 vs August 10–12)** — §3, §4. **2026-06-24** (FORCE) is the release, with a single-source caveat on Turbo's presence (DataNorth trade report; the official ByteDance launch blog does not name the Turbo); August 10–12 are Western-aggregator first-seen dates. Parameters, cutoff, license, and independent benchmarks stay [UNVERIFIED].
30. **MCP adoption figures ("78% of enterprise AI teams")** — §13. The April 2026 "78%" figure is RETRACTED (no traceable source; retracted by at least one publication); replaced by **Stacklok 41% / 30% / 29% (n=100, 2026-01)**. The ~500M/month aggregator claim stays excluded; the 30+ CVEs / 82% path-traversal claims stay flagged as single-sourced.
31. **MCP "110M monthly downloads (July 2026)"** — §13. Corrected to **June 2026** (Dev Summit India retrospective); 97M = 2026-03-25, 9,652 registry records = 2026-05-24, 12,000+ across 33 registries = April 2026 — all VERIFIED; the 17,468 census / 90,000 TechRT / 3–4× private / 21,500 Glama counts are different registers or definitions, never merged into one number.
32. **MCP spec RC caveat vs Final** — §13. The RC-window caveat (researched 2026-06-10) is OBSOLETE: 2026-07-28 is **Final** — replaced by the adoption-lag caveat (most deployed clients speak 2025-06-18/2025-11-25). The spec's own versioning semantics apply: "modern" (2026-07-28+), "legacy" (2025-11-25 and earlier), "dual-era" (both).
33. **"TCP/IP moment" vs "USB-C of AI"** — §13. "TCP/IP moment" DROPPED as quotation (no source quote found; [EDITORIAL] synthesis at most); "USB-C of AI" KEPT as supported synthesis (Anthropic's own positioning, broad independent usage).
34. **Mistral "Medium 3.1" → Medium 3.5** — §2. "Medium 3.1" is an August 2025 refresh; the April 2026 release is **Mistral Medium 3.5** (announced Apr 28, released Apr 29–30, 2026).
35. **CED = "Compressed Expert Dispatch" vs "Causal Encoder-Decoder"** — §5. Only sanctioned expansion: **Causal Encoder-Decoder** (20-layer causal encoder + 20-layer decoder; decoder KV projected from encoder final hidden states). The assembled draft must be grepped; every occurrence of "Compressed Expert Dispatch" is deleted.
36. **"2026 = the year MoE becomes the default architecture" (Wave 1 editorial)** — §9. Restated: MoE became the default for very large open-weight models; dense counterexamples (Gemma 4 31B, Qwen3.8-27B, gpt-oss-20b, the ~150M controlled study) and undisclosed closed architectures keep the absolute form false. Wave 1's stronger "without hurting inter-GPU bandwidth" for fine-grained experts is marked [UNVERIFIED / contradicted as stated] per MoE Parallel Folding (arXiv 2504.14960v2); Kimi K3's Stable LatentMoE + Quantile Balancing + MXFP4 stands as the documented bandwidth-preserving co-design, not proof that granularity is free.
37. **Kimi K3 active params (104B vs ~50B)** — §9. Adjudicated as **104B** (technical-report sources, majority); the ~50B single-Medium-analysis dispute is recorded as a do-not-cite-without-checking note, not silently dropped. (The per-expert geometry inconsistency is separately open — see A.2.4.)
38. **Diffusion speed "2x" vs "1.2–1.7x"** — §9, §10. The later **v0.1.808-beta** figure (1.2–1.7x) is used; the downgrade is a vendor-issued correction, not an independent finding. Unsloth's "12x" vs "3–5x" vs component numbers are all recorded with dates, not collapsed to one.
39. **GLM-5.3-Flash context (1M vs 300K)** — §9. The **1M (config)** figure is used; the basis of the 300K LumaDock-eval figure is unclear. Always-on-reasoning (a breaking change) and the ~2-week safety hold with 08-28/29 weight shipment are recorded.
40. **Warner "intrusion" framing (June 11) vs authorized red-team drill** — §17. The "intrusion" reading is PARTIALLY CONTRADICTED: authorized drill, relayed secondhand; walked back by The Economist's Shashank Joshi on June 21. The quote itself is VERIFIED as a quote. The 17:21 ET BIS-directive timestamp is [UNVERIFIED] as primary-sourced (attributed to 9to5Mac).
41. **OpenAI confidential IPO filing date (June 8 vs May 22)** — §21. **June 8, 2026** [VENDOR] is used everywhere; the May 22 digest claim appears in one digest and one aggregator only and stays [UNVERIFIED].
42. **Sora 2 launch "2026-09-30"** — §12, §19. Corrected to **2025-09-30** (VentureBeat); a secondary page propagated the wrong year. "Sora shut down" refers to the standalone product and public API only (model survives in ChatGPT paid tiers; team continues as world-simulation research unit per CIOL, secondary).
43. **Cerebras valuation conflict ($56.4B at pricing vs ~$95B first-day)** — §15. Resolved as different measurement points (pricing vs first-day close), kept with caveats.
44. **SMCI share-drop figures (~8% vs ~28%)** — §15. Different measurement windows; kept [SECONDARY] with windows stated.
45. **Mooncake figures (paper-era +115%/+107%/+75% vs 2026 measurements)** — §7. The paper-era figures predate the 2026 waves; the verified 2026 measurement is the Tsinghua ToS 2025 evaluation of Mooncake Store (TCP 17–22%, RDMA 26–33% vs Redis). Never present paper-era figures as 2026 production benchmarks.
46. **DeepSeek-V4 launch day (April 23 vs April 24)** — §6, §7. Secondary-source disagreement; adjudicated as immaterial and not worth resolving in the final document.
47. **SAG vs Counterpoint robotics application splits** — §22. Preserved as incompatible (70%+ industrial vs >60% entertainment/research), not blended. Tracker denominators (19,100 → >40,000) reported as a range. (Carried as preserved-not-blended; see also A.2.7 for what is unresolved.)
48. **"Beats Opus 4.8" (Kimi K3)** — §3. Adjudicated as **benchmark-dependent** (GDPval-AA v2 1687 vs 1600; AA Intelligence ~57); never a blanket verdict. Brief "Kimi K2.6 close to Opus 4.6" superseded by **ahead on SWE-Bench Pro (58.6% vs 53.4%)**.
49. **"H.R. 11" bill number** — §13, §17, §18. Stays [UNVERIFIED]: the released Kill Switch Act text carries a blank placeholder; never cited as an assigned bill number. (Kill Switch Act itself introduced Thursday, **July 23, 2026** [Lieu + Moran; ≥$500M revenue / ≥$100M compute; VERIFIED].)

---

## A.2 Unresolved contradictions

**Count: 14** unresolved contradictions (listed below). Sources conflict and no adjudication was possible; each entry carries both claims, provenance labels, and what would settle it. Minor ±1-day date ambiguities (Mistral Medium 3.5 Apr 29 vs 30; V4-Flash-Vision-Exp Aug 21 vs Aug 31; Opus 4.7 Apr 16 vs Apr 17) are recorded with ±1-day flags in A.3 rather than counted here.

