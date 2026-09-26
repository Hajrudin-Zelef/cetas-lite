---
id: etape4-tracka-vllm-sglang/02-part-2-sglang/9-enterprise-support-commercial-backers
title: "9. Enterprise support & commercial backers"
domain: part-2-sglang
role: deep-dive
task: reference
actors: ["AMD", "Google", "Intel", "Nvidia", "SGLang", "vLLM", "xAI"]
dates: ["2025-08", "2026-01-21", "2026-05", "2026-09-04", "2026-09-22"]
keywords: ["amd", "copilot", "diffusion", "disclosure", "governance", "gpu", "inference", "inference engine", "intel", "mcp", "nvidia", "omni"]
source: docs/RAG/etape4_trackA_vllm_sglang.md
source_anchor: ""
source_lines: [894, 918]
section: "PART 2 — SGLang"
sha256: a45bc0cf17964259be0608b4a650152b9e5f14a69f64a1bac57ffaedc1ee99fd
---

# 9. Enterprise support & commercial backers

- **SGLang Router:** cache-aware load balancer/router (Rust `sgl-router` + newer Rust server integration); the original v0.4 (Dec 2024) shipped a "cache-aware load balancer" [official]. Note: a third-party compat doc flags the *classic* `sgl-router` as missing `max_model_len` in `/v1/models` and no `/server_info` route — the **SGLang Model Gateway** (`sgl-model-gateway`) proxies to workers and is the supported path [secondary](https://github.com/fuzzifikation/vllm-copilot/blob/HEAD/docs/sglang-compat-plan.md).
- **SGLang Model Gateway (SMG):** high-performance model-routing gateway — load balancing, PD disaggregation, multi-model routing, gRPC, MCP, Kubernetes service discovery, history-store/privacy controls [secondary docs ref](https://github.com/hygon-ai/sglang-das/blob/HEAD/docs/docs/advanced_features/sgl_model_gateway.mdx); gateway-v0.3.1 (Jan 2026) with 10–12× routing speedup [official].
- **Related repos (sgl-project org):** `sglang-omni` (audio/TTS/ASR, v0.1.6, Sept 2026) [official](https://github.com/sgl-project/sglang-omni); `sgl-learning-materials` (slides/adoption); diffusion serving inside the main repo (SGLang-Diffusion). Downstream/community: `bytedance-iaas/sglang` fork, `hygon-ai/sglang-das`, LocalAI SGLang backend, LMCache integration work [secondary].
- **Dependencies (v0.5.20):** FlashInfer 0.6.18, sglang-kernel 0.4.7 (split out as standalone `sgl-kernel` package since v0.4.1), sgl-deep-gemm 0.2.0, sgl-deep-ep 0.1.2, mooncake 0.3.13, TileLang 0.1.12, compressed-tensors 0.18.0, transformers 5.3.0, torch 2.11–2.13, CUDA 13.x [official].
- **Contributors:** v0.5.20: 713 PRs from 237 contributors; v0.5.19: 786 PRs / 214; v0.5.18: 710 / 212 [official]. Release captains seen in API: Qiaolin-Yu, Fridge003, Kangyan-Zhou [official].
- **Sponsors/grants:** a16z Open Source AI Grant, third batch (announced 2025-06, per README news) [official]; PyTorch Ecosystem member (Mar 2025) [official]. Commercial company: **RadixArk** (see §9).

---

## 9. Enterprise support & commercial backers

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

