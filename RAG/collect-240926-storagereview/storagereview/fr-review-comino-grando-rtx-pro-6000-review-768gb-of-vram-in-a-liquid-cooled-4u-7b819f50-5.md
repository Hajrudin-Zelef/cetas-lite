---
id: collect-240926-storagereview/storagereview/fr-review-comino-grando-rtx-pro-6000-review-768gb-of-vram-in-a-liquid-cooled-4u-7b819f50-5
title: "fr-review-comino-grando-rtx-pro-6000-review-768gb-of-vram-in-a-liquid-cooled-4u--7b819f50"
domain: storagereview
role: reference
task: reference
actors: ["Alibaba", "MiniMax", "Mistral", "Nvidia", "vLLM"]
dates: []
keywords: ["benchmark", "blackwell", "cost", "decode", "fp4", "fp8", "gpu", "inference", "inference engine", "latency", "liquid cooling", "llama"]
source: docs/RAG/clean_en/storagereview/fr-review-comino-grando-rtx-pro-6000-review-768gb-of-vram-in-a-liquid-cooled-4u--7b819f50.md
source_anchor: ""
source_lines: [101, 142]
sha256: f85c5dc5a7033d771ad78f2906a233e06f302cdf51fbf4083f1daf6ab3f7d4d2
---

# fr-review-comino-grando-rtx-pro-6000-review-768gb-of-vram-in-a-liquid-cooled-4u--7b819f50

The liquid cooling architecture makes this level of computing power usable in environments where traditional GPU servers cannot operate. The system is quiet enough to be installed in a startup office, a small machine room, or a dedicated corner of an open workspace. Air-cooled systems with similar GPU density typically reach 90 dB or more, a noise level high enough to require a dedicated space in a data center or, at minimum, a closed server room with effective acoustic treatment. The Grando adapts to the team using it. With full data localization, no per-token API costs, and total control over model choice, it offers a self-hosted solution that scales with a growing development team without requiring data center infrastructure or systematic cost increases.
vLLM Online Service – LLM Inference Performance
vLLM is one of the most popular high-throughput inference and serving engines for LLMs. vLLM's online serving benchmark evaluates this inference engine's real-world performance under simultaneous requests. It simulates production workloads by sending requests to a running vLLM server, with configurable parameters such as request rate, input and output lengths, and number of concurrent clients. This benchmark measures key indicators, including throughput (tokens per second), time to first token, and time per output token (TPOT), allowing users to understand vLLM's performance under different load conditions.
We tested inference performance across a comprehensive suite of models covering various architectures, parameter scales, and quantization strategies to evaluate throughput under different concurrency profiles.
Summary of Results
| Model | Precision | Equal (256/256) | Heavy Prefill (8k/1k) | Decode-Heavy (1k/8k) | 
|---|---|---|---|---|
| Comino Grando with 8× RTX PRO 6000 Blackwell — vLLM Inference Results (tok/s, peak at BS=256) |  |  |  |  | 
| GPT-OSS 20B | ep_dp1 | 17,280 | 32,061 | 11,187 | 
| GPT-OSS 120B | ep_dp1 | 11,726 | 21,636 | 7,570 | 
| Llama 3.1 8B Instruct | FP8 | 12,109 | 20,137 | 7,353 | 
| Llama 3.1 8B Instruct | FP4 | 11,954 | 20,206 | 7,239 | 
| Llama 3.1 8B Instruct | BF16 | 11,752 | 17,346 | 6,155 | 
| Qwen3 Coder 30B A3B | FP8 | 10,985 | 16,659 | 4,907 | 
| Qwen3 Coder 30B A3B | BF16 | 10,588 | 16,680 | 4,829 | 
| Mistral Small 3.1 24B | BF16 | 8,925 | 11,846 | 4,975 | 
| MiniMax M2.5 (230B) | ep_dp1 | 5,753 | 7,357 * | 2,555 | 
| All values are expressed in tok/s, maximum throughput at BS=256. *MiniMax M2.5 heavy prefill peaked at BS=128 (7,357 tok/s); BS=256 was 7,141 tok/s. |  |  |  |  | 
GPT-OSS 120B and 20B
The GPT-OSS model family was tested in 120B and 20B configurations on the Comino Grando.
GPT-OSS 120B
Under an equal workload (256/256), the 120B model reaches 268.85 tok/s at BS=1, 6,666.23 tok/s at BS=64, and a peak of 11,726.04 tok/s at BS=256. With heavy prefill (8k/1k), throughput starts at 1,375.69 tok/s, climbs to 16,374.19 tok/s at BS=64 and 17,944.55 tok/s at BS=128, reaching a peak of 21,636.41 tok/s at BS=256. Decode-heavy (1k/8k) goes from 196.28 tok/s at BS=1 to 7,569.97 tok/s at BS=256, with latency well controlled at lower concurrency levels.
GPT-OSS 20B
The 20B model achieves a throughput of 334.80 tok/s at BS=1 under an equal workload, 10,303.56 tok/s at BS=64, and a peak of 17,280.12 tok/s at BS=256. In prefill-heavy mode, throughput starts at 2,007.90 tok/s, climbs to 24,990.46 tok/s at BS=64 and 26,866.25 tok/s at BS=128, reaching a peak of 32,060.72 tok/s at BS=256, the highest absolute prefill throughput recorded for both model sizes. Heavy decode throughput goes from 286.08 tok/s at BS=1 to 11,187.36 tok/s at BS=256, delivering approximately 1.5 times the decode throughput of the 120B at maximum concurrency while maintaining lower latency.
Qwen3 Coder 30B A3B Instruct and FP8 Instruct
The Qwen3-Coder-30B-A3B-Instruct model was tested with BF16 and FP8 precision.
Qwen3-Coder-30B-A3B-Instruct (BF16)
Under an equal workload (256/256), the BF16 model reaches 1,902.32 tok/s at BS=8, 6,683.58 tok/s at BS=64, and a peak of 10,587.56 tok/s at BS=256. With heavy prefill (8k/1k), it starts at 1,256.03 tok/s at BS=1, climbs to 14,400.57 tok/s at BS=64 and 15,308.35 tok/s at BS=128, peaking at 16,679.52 tok/s at BS=256. Decode-heavy (1k/8k) goes from 169.19 tok/s at BS=1 to 4,828.82 tok/s at BS=256, with latency well controlled at lower concurrency levels.
Qwen3-Coder-30B-A3B-Instruct (FP8)
The FP8 model offers throughput comparable to BF16 in most scenarios, with an equivalent workload reaching 6,478.54 tok/s at BS=64 and peaking at 10,984.61 tok/s at BS=256, a slight improvement over BF16 during high concurrency. With heavy prefill, throughput starts at 987.48 tok/s at BS=1, climbs to 14,036.46 tok/s at BS=64 and 15,156.69 tok/s at BS=128, peaking at 16,658.98 tok/s at BS=256. Decode-heavy goes from 130.70 tok/s at BS=1 to 4,906.51 tok/s at BS=256, slightly surpassing BF16 at maximum concurrency while both configurations remain closely matched across the rest of the concurrency range.
Mistral Small 3.1 24B Instruct 2503
Under an equal workload (256/256), the model reaches 1,598.79 tok/s at BS=8, 4,713.84 tok/s at BS=64, and climbs sharply to 8,925.12 tok/s at BS=256. With heavy prefill (8k/1k), throughput starts at 897.84 tok/s at BS=1, climbs to 9,632.58 tok/s at BS=64 and 11,488.13 tok/s at BS=128, reaching a peak of 11,846.15 tok/s at BS=256. Decode-heavy throughput (1k/8k) goes from 124.98 tok/s at BS=1 to 2,653.82 tok/s at BS=64, then accelerates significantly at higher concurrency levels, reaching 4,262.53 tok/s at BS=128 and peaking at 4,975.06 tok/s at BS=256, reflecting the model's ability to maintain high decode throughput as concurrency increases.
Llama 3.1 8B Instruct
The Llama-3.1-8B-Instruct model was tested in three precision configurations on the Comino, providing a clear view of how quantization affects throughput for this model size.
Llama 3.1 8B Instruct BF16
Under an equal workload (256/256), the BF16 model reaches 2,776.42 tok/s at BS=8, 7,369.01 tok/s at BS=64, and a peak of 11,751.56 tok/s at BS=256. With heavy prefill (8k/1k), it starts at 1,645.29 tok/s at BS=1, climbs to 14,990.47 tok/s at BS=64 and 17,140.71 tok/s at BS=128, peaking at 17,345.80 tok/s at BS=256. Decode-heavy (1k/8k) goes from 234.78 tok/s at BS=1 to 6,154.73 tok/s at BS=256.
Llama 3.1 8B Instruct FP8
FP8 quantization offers a significant gain across all scenarios. At equal workload, throughput reaches 7,530.39 tok/s at BS=64 and peaks at 12,108.98 tok/s at BS=256. With heavy prefill, throughput climbs to 16,546.53 tok/s at BS=64 and 19,306.49 tok/s at BS=128, peaking at 20,137.35 tok/s at BS=256, a gain of approximately 16% over BF16 at full concurrency. With decode-heavy, throughput peaks at 7,353.40 tok/s at BS=256, approximately 19% more than BF16.
Llama 3.1 8B Instruct FP4
FP4 offers very competitive throughput compared to FP8 at high concurrency levels, although it is slightly behind for small batches. At equal workload, maximum throughput reaches 11,954.40 tok/s at BS=256, and with heavy prefill, it peaks at 20,205.57 tok/s at BS=256, narrowly surpassing FP8 at full concurrency. With decode-heavy, maximum throughput reaches 7,239.29 tok/s at BS=256, remaining only a few percent behind FP8, making FP4 an attractive option when memory efficiency is paramount without significant throughput sacrifice.
MiniMax M2.5
The MiniMax-M2.5 230B, tested on the Comino Grando, was the largest and most demanding model we used.
