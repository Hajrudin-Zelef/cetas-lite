---
id: collect-261001-ia-llm/ia-llm/etape4-tracka-vllm-sglang-14
title: "Step 4 — Inference Serving Frameworks: vLLM + SGLang (February 1 → September 22, 2026)"
domain: ia-llm
role: reference
task: reference
actors: ["AMD", "DeepSeek", "Google", "Huawei", "Intel", "Nvidia", "SGLang", "vLLM", "xAI"]
dates: ["2025-08", "2026-01-21", "2026-05", "2026-09-04", "2026-09-22"]
keywords: ["inference", "sglang", "vllm", "agentic", "agents", "amd", "ascend", "cost", "deepseek", "disclosure", "governance", "gpu"]
source: docs/RAG/collect-261001-ia-llm/etape4_trackA_vllm_sglang.md
source_anchor: ""
source_lines: [905, 941]
sha256: 9fefca5892a4b8f85cf65b9145b31ebc6e0991a1d4a67908248c1fbff8adfa89
---

# Step 4 — Inference Serving Frameworks: vLLM + SGLang (February 1 → September 22, 2026)

- **RadixArk — the commercial company behind SGLang** [secondary, TechCrunch 2026-01-21](https://techcrunch.com/2026/01/21/sources-project-sglang-spins-out-as-radixark-with-400m-valuation-as-inference-market-explodes/):
  - Spun out of the Berkeley SGLang team; announced August 2025.
  - Valued at **~$400M** in a round led by **Accel** (TechCrunch, citing two people familiar; round size unconfirmed by TechCrunch).
  - A later summary reports a **$100M seed at $400M post-money (May 2026), led by Accel and Spark Capital, with NVIDIA's NVentures and AMD participating** [secondary](https://www.beri.net/article/vllm-vs-tensorrt-llm-vs-sglang-inference-runtime-2026) — the $100M figure is single-sourced; treat as [unverified].
  - Founders: **Ying Sheng** (CEO; ex-xAI, ex-Databricks research scientist) and **Banghua Zhu** [secondary]. Earlier angel backing incl. Intel CEO Lip-Bu Tan [secondary].
  - Product direction: keep developing SGLang as open-source engine; **Miles** RL framework; **paid hosting services** launched [secondary](https://techfundingnews.com/radixark-sglang-spinoff-400m-valuation-ai-inference/).
  - **TPU partnership:** "RadixArk and Google bring full SGLang features to TPUs" (2026-07) [official README news].
- **Enterprise contact:** `sglang@lmsys.org` for "adopting or deploying SGLang at scale, including technical consulting, sponsorship opportunities, or partnership inquiries" [official README].
- **Known production users (project-stated):** xAI ("default inference engine for xAI" per a 2025 community deck [secondary]), Cursor, LinkedIn, plus the cloud/GPU providers in §1 [official README].
- **Vendor-built platforms on SGLang:** **Atlas Cloud "Atlas Inference"** — press release claims up to 2.1× throughput vs competitors on 12 nodes, sub-5 s TTFT, 100 ms ITL at 10,000+ concurrent sessions, via PD disaggregation + DeepEP + two-batch overlap; SGLang core developer Yineng Zhang quoted [vendor-reported](https://www.newscom.com/accesswire?p_p_id=listenersearchresults_WAR_searchportlet&p_p_lifecycle=0&p_p_state=normal&p_p_mode=view&p_p_col_pos=1&p_p_col_count=2&_listenersearchresults_WAR_searchportlet_pageNumber=1&_listenersearchresults_WAR_searchportlet_struts.portlet.action=%2Fview%2FshowDetail&_listenersearchresults_WAR_searchportlet_tagId=awire192676).
- **Security posture (enterprise-relevant caveat):** August-2026 CERT/CC-coordinated disclosure of **six SGLang vulnerabilities** (pickle-based IPC in expert-backup path — mitigation `SGLANG_USE_PICKLE_IPC=false`; opt-in dumper `DUMPER_SERVER_PORT`; unauthenticated `/server_info`; NCCL weight-broadcast exfiltration; `update_weights_from_disk`/`update_weights_from_huggingface` fetching; mitigations: API key, network isolation, HF repo allowlisting) [secondary](https://github.com/wrg-11/wrg-sigma-rules/blob/HEAD/docs/detection-notes/sglang-2026-08-disclosure-series-detection-2026-09-04.md). A Sept-2026 comparison notes SGLang's repo had **no SECURITY.md** and no documented triage/disclosure process, vs vLLM's PyTorch-Foundation governance [secondary](https://www.beri.net/article/vllm-vs-tensorrt-llm-vs-sglang-inference-runtime-2026). Status of upstream fixes as of 2026-09-22: [unverified] — confirm against current release notes before production sign-off.

---

## 10. Positioning vs vLLM

| Dimension | SGLang | vLLM |
|---|---|---|
| Core scheduling idea | **RadixAttention**: radix-tree prefix KV reuse across requests (automatic) | **PagedAttention**: block-level KV management; prefix caching opt-in (`--enable-prefix-caching`) |
| Sweet spot | Prefix-heavy: agents, multi-turn chat, RAG, few-shot evals (up to 6.4× claimed) | General production, broadest model coverage, most mature ecosystem |
| Raw throughput (H100, Llama 3.1 8B, high concurrency) | ~16,200 tok/s [secondary] | ~12,500 tok/s [secondary] |
| Single-stream unique prompts | Roughly parity to −12% [secondary] | Slightly ahead in one old test [secondary] |
| DeepSeek models | Day-0 support (V3/R1/V4), optimized MLA/DSA/MTP, claimed 3.1× on V3 [official/secondary] | Supported; SGLang is the "officially recommended engine" per one comparison [secondary] |
| Speculative decoding | EAGLE/DFlash2/DSpark/KDA/NEXTN/MTP + beam search (v0.5.19) | EAGLE3 etc. |
| PD disaggregation | First-class (Mooncake/NIXL, DCP) | Available; llm-d covers both |
| Hardware breadth | NVIDIA + AMD (strong) + Intel XPU/CPU + TPU + Ascend + MUSA; HPU experimental | NVIDIA + AMD + TPU + XPU + HPU (Habana fork, mature-ish) |
| Structured output | XGrammar, compressed-FSM JSON (3× claim), frontend DSL | Outlines/XGrammar guidance |
| Governance | LMSYS non-profit + RadixArk commercial spin-out ($400M) | PyTorch Foundation (neutral IP), formal security process; Red Hat sells supported vLLM/llm-d |
| GitHub traction | 36.3K stars / 9.1K forks (2026-09-22) | (not pulled in this pass) |
| Default port | 30000 | 8000 |

**Practitioner summary:** choose SGLang when prefix reuse dominates (agentic/RAG/multi-turn), when serving DeepSeek-family models (day-0 kernels), or when AMD MI300X/MI355X is the fleet (feature parity + strong independent cost numbers in 2026). Choose vLLM for the broadest ecosystem, formal enterprise support channels (e.g., Red Hat), and Gaudi/HPU fleets. The raw-throughput gap is real at high concurrency but collapses on unique-prompt workloads.

---

## 11. Source list

