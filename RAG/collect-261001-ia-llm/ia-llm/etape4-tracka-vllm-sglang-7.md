---
id: collect-261001-ia-llm/ia-llm/etape4-tracka-vllm-sglang-7
title: "Step 4 — Inference Serving Frameworks: vLLM + SGLang (February 1 → September 22, 2026)"
domain: ia-llm
role: reference
task: reference
actors: ["AMD", "AWS", "DeepSeek", "Google", "Huawei", "Hugging Face", "Intel", "Meta", "Mistral", "Moonshot", "Nvidia", "OpenAI", "Oracle", "SGLang", "TensorRT-LLM", "Z.ai", "vLLM"]
dates: ["2025-01-27", "2025-12", "2026-04-01", "2026-05", "2026-06-30", "2026-07-27", "2026-08-25", "2026-09", "2026-09-01", "2026-09-20", "2026-09-22"]
keywords: ["inference", "sglang", "vllm", "agentic", "amd", "ascend", "aws", "benchmark", "benchmarks", "cost", "deepseek", "glm"]
source: docs/RAG/collect-261001-ia-llm/etape4_trackA_vllm_sglang.md
source_anchor: ""
source_lines: [564, 652]
sha256: 0e791bf3813116f1574cbba968fcaa05ec353e1820bfc9e675c92207a4e4a85e
---

# Step 4 — Inference Serving Frameworks: vLLM + SGLang (February 1 → September 22, 2026)

| Dimension | vLLM | SGLang | Notes |
|---|---|---|---|
| Core innovation | PagedAttention (block paging) | RadixAttention (radix tree) | Architectural differentiator |
| GitHub (2026-09-22) | 92,444★ / 22,525 forks | 36,323★ / 9,061 forks | vLLM ~2.5× stars [official] |
| Prefix-heavy throughput | ~12,500 tok/s | ~16,200 tok/s (H100, Llama-3.1-8B) | ~29% SGLang edge **only with shared prefixes**; unique prompts within a few % (RunPod, 2026) [secondary] |
| Prefix caching | Block-level hash (APC, always-on in V1) | Token-level radix tree (better for branching/agentic) | SGLang wins tree-structured reuse (MCTS, tool rollbacks) [secondary] |
| Structured output | xgrammar/guidance; noticeable overhead at high batch (per TECHSY) | Minimal overhead (overlapped mask gen) | SGLang edge on constrained decoding [secondary] |
| Hardware breadth | NVIDIA, AMD, Intel (XPU/Gaudi), TPU, CPU, Ascend | NVIDIA, AMD (narrower) | vLLM's main differentiator: breadth [secondary] |
| Model adoption speed | Fastest day-0-ish support (Kimi K3, DeepSeek V4.x, GLM-5.x) | Fast, smaller kernel team | [secondary] |
| Ease of setup | `pip install vllm` | More friction (mem-fraction, backend flags) | Independent benchmark author noted "real setup cost" for SGLang [independent] |
| Quantization | 29 `--quantization` values; Marlin W4A16; NVFP4/MXFP4; online quant | Comparable breadth; quark_int4fp8_moe, petit_nvfp | Rough parity mid-2026 [secondary] |
| Disaggregation | NIXL + Mooncake; DCP/PCP | Mooncake/NIXL backends | Both support [secondary] |
| Long-context TTFT | Weaker (L40S: 63.8 s vs TensorRT-LLM 8.6 s @ 8k ctx) | Similar to vLLM (68.6 s) | TensorRT-LLM leads this niche [independent] |

- Positioning summary (multiple 2026 sources converge): **vLLM = breadth** (models × hardware × deployment,
  largest ecosystem, fastest new-model support); **SGLang = depth on prefix-heavy/agentic workloads**
  (RadixAttention, structured generation); **TensorRT-LLM = NVIDIA peak throughput** at the cost of a
  compile step; LMDeploy leads quantized-model serving in some comparisons. [secondary][independent]
- ⚠️ TECHSY (Sep 2026) claims "Hugging Face put TGI into maintenance mode in December 2025 and now points
  teams toward vLLM or SGLang" — could not independently verify; treat as [unverified].
- ⚠️ Star-count comparisons in third-party articles are stale (e.g. "17k+ vs 15k+", "72.4k as of Mar 2026");
  current counts (Sep 2026) are in §9.1. [official]

---

## 12. Key uncertainties & gaps

1. ⚠️ **v0.30.0 released 2026-09-22 (today)** — release notes read the same day; adoption/bug reports not
   yet available. Flag early-adopter risk, esp. breaking changes (scale-out endpoints opt-in, `g_idx`
   removal, gRPC entrypoint deprecation). [official]
2. ⚠️ **MRV1 removal targeted at v0.32** — exact v0.32 date unknown; feature gaps (sequence parallelism,
   elastic EP, custom logits processors) still fall back to MRV1. [official]
3. ⚠️ **Production-user claims** (Meta/LinkedIn/Mistral/HF/Uber/JPMorgan) are third-party assertions; no
   company-issued confirmation found in this research. [unverified]
4. ⚠️ **AWS Trainium** support status not confirmed from an authoritative 2026 source. [unverified]
5. ⚠️ **TGI maintenance-mode claim** (TECHSY) unverified. [unverified]
6. ⚠️ **Sponsor list** copied from forked READMEs; verify against the canonical README before citation. [official]
7. ⚠️ PR-reported performance percentages in release notes (e.g. "−2.94% E2E TPOT", "+15% throughput")
   are micro-benchmark claims from the PR author, not audited results. [official — vendor-reported]
8. ⚠️ BitsAndBytes migration to out-of-tree plugin (v0.28) — plugin's maintenance/availability not verified
   in this research. [unverified]

---

## 13. Sources

- GitHub releases (official): https://github.com/vllm-project/vllm/releases (API + v0.26.0/v0.29.0/v0.30.0 notes)
- vLLM blog — Kimi K3 (2026-07-27): https://github.com/vllm-project/vllm-project.github.io/blob/HEAD/_posts/2026-07-27-k3.md
- vLLM V1 alpha post (2025-01-27): https://github.com/vllm-project/vllm-project.github.io/blob/HEAD/_posts/2025-01-27-v1-alpha-release.md
- vLLM Production Stack docs: https://github.com/vllm-project/production-stack
- vLLM Gaudi plugin release notes: https://github.com/vllm-project/vllm-gaudi/blob/HEAD/docs/release_notes_v0.26.0.md
- vLLM Ascend release notes (v0.22.1rc1, 2026-06-30; v0.18.0rc1, 2026-04-01)
- vLLM recipes: https://github.com/vllm-project/recipes ; TPU inference: https://github.com/vllm-project/tpu-inference
- SemiAnalysis InferenceX blogs (2026-08-25): DeepSeek-V4-Pro-AgentX MI355X vs B200; MI355X Kimi-K2.5 AITER 7× speedup; Kimi-K3 AgentX MI355X vs GB300 NVL72
- dstack DeepSeek-R1 benchmark (H200/MI300X): https://github.com/dstackai/dstack docs/blog
- inference-bench (L4/A100 vLLM vs SGLang, May 2026): https://github.com/ree2raz/inference-bench
- ai-lab-benchmarks engines-2026-08: https://github.com/marianvid/ai-lab-benchmarks
- arXiv 2603.10031 (Feb 2026): Architecture-Aware LLM Inference on AMD MI325X
- RunPod comparison (vLLM vs TensorRT-LLM vs SGLang, 2026): https://www.runpod.io/articles/comparison/vllm-vs-tensorrt-llm
- TECHSY vLLM vs SGLang (Sep 2026): https://techsy.io/en/blog/vllm-vs-sglang
- explore.n1n.ai — SGLang vs vLLM architecture (2026-09-20)
- Red Hat Developer — LLM Compressor / nm-vllm: https://developers.redhat.com/articles/2024/08/14/llm-compressor-here-faster-inference-vllm
- ISTA — Neural Magic acquired by Red Hat: https://ist.ac.at/en/news/neural-magic-ai-startup-with-ista-mit-roots-acquired-by-red-hat/
- Red Hat AI Inference Server: https://www.press.in.th/red-hat-unlocks-generative-ai/ ; digitalisationworld (2026)
- TechTimes — VMware AI Factory (2026-09-01): https://www.techtimes.com/articles/326160/20260901/enterprise-gpu-chaos-ends-vmware-ai-factory-ships-full-stack-bare-metal-model.htm
- Google Cloud TPU docs: https://docs.cloud.google.com/tpu/docs/tpu-inference ; vLLM TPU announcement (CloudSteak)
- premai.io — LLMs on Kubernetes guide (2026): https://www.premai.io/blog/deploying-llms-on-kubernetes-vllm-ray-serve-gpu-scheduling-guide-2026/
- Oracle — Deploy OpenAI vLLM Production Stack on OKE: http://docs.oracle.com/en/learn/deploy-vllm-production-stack-oke/
- Quantization catalog (third-party, source-verified): https://github.com/kriptoburak/air-gapped-skills vllm-quantization SKILL.md ; vllm-quant-deep-dive.md (allanschramm/local-model-autotuning)
- oss-atlas vLLM page (Jul 2026): https://github.com/zhenninglang/oss-atlas vllm.md
- vllm-daily digests (2026-02/2026-04): https://github.com/vllm-project/vllm-daily
- V1 architecture skill note: https://github.com/tylertitsworth/skills vllm/SKILL.md
- candle-vllm (ericlbuehler): https://github.com/ericlbuehler/candle-vllm


# PART 2 — SGLang

# SGLang — Inference-Serving Framework (LMSYS) — Research Report

> **Research date:** 2026-09-22 (UTC)
> **Project:** RAG data collection, Step 4 (AI infra — inference/training)
> **Scope:** SGLang news and state as of September 2026
> **Provenance legend:** `[official]` = sgl-project GitHub / docs.sglang.io / lmsys.org blog; `[vendor-reported]` = company marketing claims; `[independent]` = third-party measurements (SemiAnalysis InferenceX, dstack, community); `[secondary]` = press/blogs summarizing others; `[unverified]` = could not be confirmed.
> Every fact below carries its date. Uncertainties are marked explicitly.

---

## 1. Project identity & ecosystem snapshot

