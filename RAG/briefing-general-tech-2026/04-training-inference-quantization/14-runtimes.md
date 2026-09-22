---
id: briefing-general-tech-2026/04-training-inference-quantization/14-runtimes
title: "Runtimes: vLLM, TensorRT-LLM, SGLang, llama.cpp"
domain: training-inference-quantization
role: deep-dive
task: architecture
actors: ["Intel", "Nvidia", "PrismML"]
dates: []
keywords: ["llama", "llama.cpp", "sglang", "tensorrt", "vllm", "awq", "benchmark", "blackwell", "fp4", "fp8", "gptq", "gpu"]
source: docs/RAG/briefing-general-tech-2026-en.md
source_anchor: "#g05-14"
source_lines: [5401, 5453]
sha256: 52e8945799a9665a0841f83b41a1e10309907c91d87f69653e1ed7e67f1c3c50
---

# Runtimes: vLLM, TensorRT-LLM, SGLang, llama.cpp

<a id="g05-14"></a>
### 5.14 Runtimes: vLLM, TensorRT-LLM, SGLang, llama.cpp

Formats and methods are only as real as the runtimes that serve them. A
quantization format with no serving support is a file-compression demo; a
method with no kernel is a paper. The 2026 serving landscape has four names
that matter, and their quantization support maps exactly onto the production
tiers this chapter established — which is why the runtime map is the reality
check on everything above.

**vLLM** is the open-source serving standard and the broadest quantization surface in the industry. Its confirmed format support — FP8, NVFP4, MXFP4, AWQ and GPTQ, all natively supported per the project's documentation — reads as the index of the production-format list itself: if a format works in vLLM, it is deployable on commodity infrastructure without vendor lock-in. The Marlin/Machete kernels live here too, providing the hand-tuned 4-bit execution paths on Ampere and Hopper that turn AWQ's and GPTQ's compressed weights into actual throughput rather than merely smaller memory footprints. vLLM's role in the 2026 story is standardization-through-openness: it is the reason the "safe tier / aggressive tier" framing of Sections 5.8–5.9 is not analyst taxonomy but deployment reality — operators configure FP8 everywhere, 4-bit where the evals pass, and the runtime doesn't care which tier you chose.

**TensorRT-LLM** is Nvidia's performance ceiling: the same confirmed format coverage (FP8, NVFP4, MXFP4, AWQ, GPTQ) with the deepest hardware-specific optimization, including the Blackwell-only NVFP4 paths that the open runtimes expose but Nvidia tunes hardest. The vLLM/TensorRT-LLM pair structures the high-end serving decision in 2026: vLLM for openness, portability and the broadest hardware base; TensorRT-LLM for maximum tokens-per-second on Nvidia silicon, where the last few percent of kernel efficiency translate directly into cost-per-token margin. Both serve the same tiers; they differ in how far down the optimization curve each deployment needs to go.

**SGLang** is the third production runtime and the home of the KV-cache frontier. Its TurboQuant integration via `--kv-cache-dtype turboquant` (Section 5.13) makes it the serving stack where the weight-quantization story and the KV-cache story meet in a single deployment: quantized weights *and* a 3-bit KV cache behind two flags. On Blackwell fleets the 2026 routing is clear — vLLM or SGLang for the newest formats, with SGLang differentiated where long-context KV-cache economics dominate the workload. SGLang's significance in this chapter is as the proof that the second quantization frontier (KV cache) cleared the lab-to-runtime transition within the same year its paper was presented at ICLR.

**llama.cpp** occupies the distinct edge, and the distinction matters. It is optimized for classic GGUF quantized models and remains the standard for CPU inference, consumer hardware, laptops, and the phone-scale deployments that Bonsai (Section 5.12) targets — the GGUF and MLX builds of the open-weights ecosystem run here. But it does *not* target NVFP4 or the Blackwell-specific format paths: Blackwell-class datacenter serving routes to vLLM/SGLang, not to llama.cpp. The division of labor is clean and stable: llama.cpp owns the edge and the laptop, the GPU serving runtimes own the datacenter, and neither is trying to be the other.

Two smaller names complete the picture and confirm the method verdict of
Section 5.10: **LMDeploy** and **Hugging Face TGI** (Text Generation
Inference) both support AWQ — further evidence that AWQ, not GPTQ, is the most
broadly served 4-bit method in the 2026 landscape, the method whose support
you assume rather than check.

The 2026 format map, in one line — the line an operator pins above the serving
fleet's configuration management: **FP8 and INT4 dominate production;
NVFP4/MXFP4 are rising fast on the Blackwell fleet; the KV-cache frontier
(TurboQuant) lives in SGLang; the edge lives in llama.cpp/GGUF.**

| Runtime | Quantization surface (confirmed) | 2026 role |
|---|---|---|
| vLLM | FP8, NVFP4, MXFP4, AWQ, GPTQ natively; Marlin/Machete kernels (Ampere/Hopper) | Open-source serving standard; broadest format coverage; openness as standardization |
| TensorRT-LLM | FP8, NVFP4, MXFP4, AWQ, GPTQ; deepest optimization incl. Blackwell-only NVFP4 paths | Nvidia performance ceiling; last-percent kernel efficiency → cost-per-token margin |
| SGLang | TurboQuant KV-cache flag (`--kv-cache-dtype turboquant`) + modern format support | Home of the KV-cache frontier; Blackwell-fleet router alongside vLLM |
| llama.cpp | Classic GGUF optimization; does NOT target NVFP4/Blackwell paths | CPU/consumer/edge inference standard; Bonsai-class phone-scale deployments |
| LMDeploy, TGI | AWQ support | Confirm AWQ as the most broadly served 4-bit method |

For the operator reader, the chapter compresses into a decision tree — the
tiered policy that Section 5.15's second consequence describes. Start every
new deployment at FP8: it's the safe tier, the evals are routine, the tooling
is assumed. Where the quality gates pass with headroom — and the
DeepSeek-R1-0528 ≤1% datapoint suggests they often will — promote weights to
INT4 with grouping via AWQ, served through vLLM or TensorRT-LLM with Marlin-
class kernels on Ampere/Hopper. Where the fleet is Blackwell and the evals
specifically favor it, trial NVFP4 or MXFP4 — but benchmark against GPTQ-INT4
on your own models first, per the Intel KLD lesson, because the FP4-vs-INT4
hierarchy is empirical, not settled. Where context lengths are long and the KV
cache dominates memory, add TurboQuant behind SGLang's flag. Where the
deployment target is a phone or a laptop, reach for llama.cpp and the GGUF
ecosystem — and keep an eye on Bonsai-class artifacts, evaluated, not just
downloaded. The tree's shape is the chapter's shape: tiers, gates, and the bit
ladder descending.

