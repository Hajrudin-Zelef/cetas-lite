---
id: ai-industry-kb-2026-wave6/02-deepseek/v4-1-flash-ga-2026-09-10-ground-up-redesign-secondary
title: "V4.1-Flash — GA 2026-09-10, ground-up redesign [SECONDARY]"
domain: deepseek
role: deep-dive
task: actor-profile
actors: ["Baseten", "DeepSeek", "Meta", "Nvidia", "OpenAI", "Z.ai"]
dates: ["2023-11", "2023-11-29", "2024-05-06", "2024-09-05", "2024-12-26", "2025-01-20", "2025-08-21", "2025-12-01", "2026-04-24", "2026-09-08", "2026-09-10", "2026-09-11"]
keywords: ["agent", "agentic", "attention", "benchmark", "compute", "context window", "deepseek", "distillation", "fp4", "fp8", "glm", "gpt-5.6"]
source: docs/RAG/ai-industry-knowledge-base-2026-wave6.md
source_anchor: ""
source_lines: [680, 721]
section: "§2. DeepSeek"
delta_of: ai-industry-kb-2026
sha256: e18c5dbda7ff65a27afe04f6a5ce246edc0e57c5de1d53e672532d202609474d
---

# V4.1-Flash — GA 2026-09-10, ground-up redesign [SECONDARY]

### V4.1-Flash — GA 2026-09-10, ground-up redesign [SECONDARY]
- **2026-09-08** — two-day beta preview under the temporary model ID `deepseek-v4.1-flash-expires-on-0910`, capped at **20 concurrent requests**; the beta window closed and **GA landed 2026-09-10**. [SECONDARY]
- Scale ledger: **552B backbone + 196B Engram** (conditional memory module accessed sparsely via N-gram-hashing token lookup) — **CONTRADICTION**: secondary write-ups compute the total as **~748B** (552+196) or **~763B**. Both are secondary; the 552B/196B split is consistent across sources, only the summed headline differs. [SECONDARY]
- **Causal Encoder-Decoder (CED)**: instead of recomputing KV cache per decoder layer, the decoder's global KV is **projected from the encoder's final hidden states** — roughly **halving prefill computation**, which matters most at long context and high cache-miss rates. [SECONDARY]
- **CSA2 three modes**: Full (compute KV and select attention indices from scratch), Reindex (reuse shared KV, pick fresh indices), Reuse (borrow both KV and indices from a prior layer). Cross-layer sharing is the main reason the cache shrinks. One independent reconstruction: main KV entry = 512 channels × 4 bits (256 B) + 32 one-byte scales = 288 B; indexer-K = 128 × 4 bits (64 B) + 4 scales = 68 B; encoder 3 Full layers at 2:1 compression, decoder 1 Full layer at 1:1 → (3×288 + 3×68)/2 + 288 + 68 = **890 bytes/token**. [SECONDARY]
- **SWA Bounded Replay**: instead of persisting full sliding-window attention state to SSD, reconstruct by replaying only the most recent window — persistent cache footprint ≈ **1/8 of V4 Flash**. [SECONDARY]
- **DSpark speculative decoding**: a lightweight **3-block drafter** generating **5 draft tokens per forward pass**, trained separately after pre-training and kept aligned during post-training; it **replaces the MTP module** from DeepSeek V3. [SECONDARY]
- **FP4 KV caching in MXFP4 (E2M1)** format; the 890 B/token figure is **~1/4 of V4 Flash** and DeepSeek's headline **437× vs DeepSeek V1** — but V1 shipped November 2023 with a 4K context window, pre-MLA. The honest comparison from the same sources: **4× vs V4 Flash** (the model it actually replaces); full-cache comparison is V1 1.59 GB vs V4.1-Flash 0.93 GB = **256× the context in 41% less memory**. The corpus should cite the 4×/256× figures, not the 437×. [SECONDARY]
- Serving economics (community analysis): **510.3 GB across 48 files**, ~313 GB resident once the 196.9B-parameter engram tables move to host memory; on **8× NVIDIA H200 SXM (141 GB each)**, ~815 GB free → **~915 concurrent users each holding a full million tokens**, vs 40 for DeepSeek V3 on the same node. [COMMUNITY]
- Training: **from scratch on 45T tokens**, context extended to 1M at the **34T-token mark**. Post-training: SFT, RL, on-policy distillation; the agentic gains are credited to **large-scale automated synthesis of agent tasks and environments** — synthetic agent-task generation at scale. **Reasoning effort continuously controllable from 1 to 100** per request. [SECONDARY]
- Observed serving (third-party): **215 output tokens/s**, p50 TTFT 842 ms, p95 2.40 s, 1.2% error rate. **MIT license** — downloadable, self-hostable, fine-tunable, commercial use without royalties or usage restrictions. [SECONDARY]
- Community extreme: a single **DGX Spark (GB10)** recipe serves the full FP4 checkpoint — 73.8 GB hot-expert set (25.6% of 15,360 routed FP4 experts) resident, remainder streamed off NVMe, DSpark worth ~1.5× — full checkpoint quality, no requantization. [COMMUNITY]

### V4.1-Flash vendor benchmark table — full three-generation comparison [SECONDARY]
- KDnuggets published DeepSeek's reported three-generation table (all vendor claims):
  - **DeepSWE v1.1**: V4-Flash 54.4 → V4-Pro 62.7 → **V4.1-Flash 74.2**
  - **Terminal-Bench 2.1**: 82.7 → 87.9 → **90.6**
  - **CyberGym**: 76.7 → 83.3 → **88.1**
  - **AutomationBench**: 37.7 → 43.2 → **54.8**
  - **Agent's Last Exam**: 25.2 → 25.7 → **31.8**
- Note: these index numbers must not be compared to AA Index figures for these models — different instruments. [SECONDARY]

### V4.1-Flash — launch week ecosystem (2026-09-10/11) [SECONDARY]
- **Baseten** added V4.1-Flash to its Model APIs on **2026-09-11** — the day after launch — describing it as DeepSeek's **third open-weight flash release of 2026** and "the only model of its scale" using CED; support for Baseten's Loops training product was announced as coming soon. [SECONDARY]
- **First non-experimental native image input**: V4.1-Flash is DeepSeek's first non-experimental model with native image understanding — previously limited to the experimental **V4-Flash-Vision-Exp**. Card figures: **Chartography 78.9 vs 64.3**, **ZeroBench 49 vs 35** (new vs experimental). [SECONDARY]
- Launch-day cross-lab comparisons (vendor card): **GPQA Diamond 90.9**, **Codeforces 3471**, **HLE with tools 63.9**; **CyberGym 88.1 vs 84.5 for GPT-5.6 Sol and GLM 5.3** (per one secondary write-up). [SECONDARY]
- The size dispute, reconciled: one community analysis (r/LocalLLaMA, cited by OrcaRouter) reads the checkpoint as **748B total, ~510 GB in FP8** — the reconciliation offered is that DeepSeek's "552B" counts only the compute backbone while the download also carries a large sparsely-accessed memory table looked up rather than computed through. Treat 552B as vendor-reported, 748B as an unrefuted community reading. [COMMUNITY]

### DeepSeek model lineage — consolidated table [SECONDARY]
- One community inference guide consolidates the full lineage (all MIT open-weight MoE unless noted):
  - **DeepSeek-LLM 7B/67B** (2023-11-29): 7B/67B dense, 4K, MIT.
  - **DeepSeek-Coder 1.3B–33B** (2023-11): 16K, MIT/DeepSeek License.
  - **DeepSeek-V2/Lite** (2024-05-06): 236B (16B Lite), 21B (2.4B) active, 128K (32K Lite).
  - **DeepSeek-V2.5** (2024-09-05): 236B, 21B active, 128K.
  - **DeepSeek-R1/R1-Zero** (2025-01-20): 671B, 37B active, 128K, MIT.
  - **DeepSeek-V3** (2024-12-26): 671B, 37B, 64K (128K via V3.1).
  - **DeepSeek-V3.1/Terminus** (2025-08-21): 671B, 37B, 128K, hybrid reasoning.
  - **DeepSeek-V3.2/Speciale** (2025-12-01): 671B, 37B, 128K, MoE + DSA.
  - **DeepSeek-V4-Flash/Base** (2026-04-24): 284B, 13B, 1M.
  - **DeepSeek-V4-Pro/Base** (2026-04-24): 1.6T, 49B, 1M.
- Cadence note: from ~1 major family/year (V2→V3) to multiple substantial releases/year by 2025–2026, "with a clear shift toward long-context efficiency and agentic/reasoning-first capabilities." [COMMUNITY]

