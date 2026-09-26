---
id: ai-industry-kb-2026/08-quantization-model-formats/part-20
title: "8. Quantization & Model Formats (part 20)"
domain: quantization-model-formats
role: deep-dive
task: quantization
actors: ["AMD", "Alibaba", "Intel", "Nvidia", "SGLang", "vLLM"]
dates: ["2026-07", "2026-08-27"]
keywords: ["quantization", "accelerator", "amd", "attention", "awq", "benchmark", "bitnet", "blackwell", "consumer", "cost", "datacenter", "decode"]
source: docs/RAG/ai-industry-knowledge-base-2026.md
source_anchor: ""
source_lines: [4511, 4525]
section: "8. Quantization & Model Formats"
sha256: af7976424478c2db8edaf42146f9532acb435e645ee030d07252d5eaba3d8b12
---

# 8. Quantization & Model Formats (part 20)

- **Dynamic v3.0 independence**: partially addressed by the 2026-08-27 community test (KLD corroborated, 9.5% decode-speed regression measured, one UD variant shipped a llama-server bug) — Divergence-300 specifically and benchmark-score replication still open.
- **FP4 KV on consumer Blackwell**: stock vLLM path datacenter-only; community FA2+XQA route (1.6× fp8 KV pool, July 2026) needs wider validation before production use.
- **BitNet at scale**: the 100B-on-a-single-CPU claim is theoretical; whether natively-trained ternary models scale past 3B without quality collapse is unanswered (Bonsai's converted 1-bit at 27B is the parallel bet).
- **MXFP4 beyond NVIDIA**: Dell's 6.1× datapoint on AMD MI355X suggests the OCP standard could become the cross-vendor 4-bit default; AMD next-gen and Intel accelerator coverage is the thing to track.
- **Sub-2-bit training economics**: Bonsai proved conversion-to-1-bit at 27B with >90% retention; native 1-bit training (BitNet-style) at 27B+ has not been demonstrated publicly.
- **NVFP4 W4A16 vs W4A4 confusion at download time**: checkpoint naming and model-card flags remain the failure point — always check the card before downloading.
- **AWQ's tooling tail**: the format remains the dominant compat INT4 path, but new-model coverage (Qwen3.5+), llm-compressor pinning, and calibration-OOM workarounds show where the "standard" frays.

- **NVFP4 W4A16 vs W4A4 confusion at download time**: checkpoint naming and model-card flags remain the failure point — always check the card before downloading; the industry needs machine-readable dtype flags.
- **AutoRound vs llm-compressor adoption split**: AutoRound is the strongest measured 4-bit PTQ, but Marlin rejects its symmetric `uint4`/`zero_point=false` form on Ampere — adoption hinges on compressed-tensors export paths.
- **INT4 KV in vLLM/SGLang mainline**: OpenVINO 2026.2 shipped it on Intel GPU; whether vLLM/SGLang mainline adopt INT4 KV (vs the FP8 default) decides long-context economics on Hopper hardware.
- **MXFP4 kernel coverage on AMD next-gen and Intel accelerators** beyond the single MI355X datapoint — the open-standard bet's progress metric.
- **Consumer-Blackwell W4A4 maturity**: GB10 production-ready and RTX 5090 software-dequant routes exist, but the stock vLLM NVFP4 KV/attention path remains datacenter-only — watch whether SM120 gets first-class FP4 KV in 2027.
- **Quantization labor vs format lock-in**: QAT's <1%-loss promise commits a model to one quantization target; the retraining cost means PTQ-plus-good-kernels (FP8, MXFP4 via llm-compressor) keeps winning on economics unless the deployment is pinned to one format for years.

