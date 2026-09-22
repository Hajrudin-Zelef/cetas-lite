---
id: etape4-tracka-vllm-sglang/01-part-1-vllm/12-key-uncertainties-gaps
title: "12. Key uncertainties & gaps"
domain: part-1-vllm
role: deep-dive
task: reference
actors: ["AMD", "AWS", "DeepSeek", "Google", "Huawei", "Meta", "Mistral", "Moonshot", "OpenAI", "Oracle", "SGLang", "TensorRT-LLM", "vLLM"]
dates: ["2025-01-27", "2026-04-01", "2026-05", "2026-06-30", "2026-07-27", "2026-08-25", "2026-09-01", "2026-09-20", "2026-09-22"]
keywords: ["amd", "ascend", "aws", "benchmark", "benchmarks", "deepseek", "gpu", "inference", "kimi", "mistral", "quantization", "research"]
source: docs/RAG/etape4_trackA_vllm_sglang.md
source_anchor: ""
source_lines: [589, 638]
section: "PART 1 — vLLM"
sha256: 97d17b7e19fcb4153ebbf04809c18dbcd19c25c2a5e82c8483fbec10b4aeda49
---

# 12. Key uncertainties & gaps

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


