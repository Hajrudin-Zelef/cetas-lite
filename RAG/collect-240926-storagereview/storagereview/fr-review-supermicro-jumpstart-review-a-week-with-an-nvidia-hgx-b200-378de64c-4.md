---
id: collect-240926-storagereview/storagereview/fr-review-supermicro-jumpstart-review-a-week-with-an-nvidia-hgx-b200-378de64c-4
title: "fr-review-supermicro-jumpstart-review-a-week-with-an-nvidia-hgx-b200-378de64c"
domain: storagereview
role: reference
task: reference
actors: ["Alibaba", "DeepSeek", "Nvidia", "OpenAI"]
dates: []
keywords: ["nvidia", "benchmark", "deepseek", "fp4", "fp8", "latency", "llama", "mixture of experts", "moe", "nvfp4", "parameters", "quantization"]
source: docs/RAG/clean_en/storagereview/fr-review-supermicro-jumpstart-review-a-week-with-an-nvidia-hgx-b200-378de64c.md
source_anchor: ""
source_lines: [58, 85]
sha256: ccf87d8ecaf9c5712df9d65eabc6a38c2f9e4b4bdfb74c8e2583d779185a959a
---

# fr-review-supermicro-jumpstart-review-a-week-with-an-nvidia-hgx-b200-378de64c

Maximum total throughput is reached at BS=256 with a total of 29,219.67 tok/s, or 30.13 tok/s per user, and a TPOT of 13.26 ms. The FP8 variant achieves lower maximum total throughput than standard precision (29.2 K vs. 32.8 K tok/s).
Llama 3.1 8B FP4 Performance
The FP4 quantized configuration shows the following results: in single-user mode (BS=1), it reaches 279.73 tok/s per user, a total throughput of 830.46 tok/s and a TPOT of 3.43 ms.
The FP4 model has seven Pareto frontier points. At BS=2, it delivers 159.95 tok/s per user, or a total throughput of 928.60 tok/s, and a TPOT of 3.43 ms. Performance continues to improve up to BS=4 (76.36 tok/s per user, or a total throughput of 1,631.69 tok/s) and BS=32 (76.09 tok/s per user, or a total throughput of 9,014.29 tok/s). The model reaches its maximum total throughput at BS=256, with 29,340.89 tok/s, 30.13 tok/s per user and a TPOT of 16.18 ms.
Sparse Model Performance
Sparse models, notably Mixture of Experts (MoE) architectures, represent an emerging approach to efficiently scaling language models. These architectures maintain a high total parameter count while activating only a subset of parameters per token, potentially improving performance per active parameter.
We evaluated two MoE architectures: DeepSeek-R1, a reasoning-focused model, and Qwen3 Coder 30B-A3B, a sparse architecture specialized in code generation. DeepSeek-R1 is the most widely used reasoning-focused model, delivering significantly higher performance than traditional language models. The Qwen3 Coder model retains all of its 30 billion parameters while activating only 3 billion per generated token. We compared Qwen3 Coder's performance with its default and FP8 floating-point quantization variants to analyze variations by quantization strategy.
DeepSeek-R1 Performance
The DeepSeek-R1 model exhibits interesting scaling behavior based on batch size. In single-user mode (BS=1), it reaches 30.24 tok/s per user, a total throughput of 88.13 tok/s and a TPOT of 29.85 ms. At BS=4, performance reaches 29.77 tok/s per user and a total throughput of 266.40 tok/s, for a TPOT of 32.04 ms, the maximum total throughput achieved across all configurations.
DeepSeek-R1 performance plateaus beyond BS=4. At BS=8, per-user throughput drops sharply to 14.98 tok/s, and this decline continues for larger batch sizes, reaching only 0.46 tok/s per user at BS=256. Total throughput remains relatively stable between 200 and 260 tok/s from BS=4 to BS=256, as the model cannot handle an increased number of concurrent requests on a single node, resulting in increased latency without significant throughput gains. It's worth noting that the B200 DGX is one of the few single-server solutions capable of running this complex model.
Qwen3 Coder 30B-A3B Performance
The Qwen3 Coder in standard precision shows the following performance based on concurrency level. In single-user mode (BS=1), the model reaches 178.30 tok/s per user, a total throughput of 527.25 tok/s and a TPOT of 5.46 ms. At BS=2, performance reaches 174.56 tok/s per user, a total throughput of 718.70 tok/s and a TPOT of 5.60 ms. As batch size increases, per-user throughput decreases while total throughput continues to grow: at BS=16, we get 127.76 tok/s per user, a total throughput of 4,204.40 tok/s and a TPOT of 6.93 ms.
The model reaches its maximum total throughput at BS=256, totaling 22,305.88 tok/s, with 46.16 tok/s per user and a TPOT of 17.64 ms. The Pareto frontier includes eight distinct points, with a duplicated entry at BS=32 (72.97 tok/s and 93.50 tok/s per user, likely corresponding to different configurations). TPOT values remain between 5 and 9 ms up to BS=64.
Qwen3 Coder 30B-A3B FP8 Performance
The FP8 quantized variant exhibits the following performance characteristics: in single-user mode (BS=1), it delivers 107.46 tok/s per user, a total throughput of 317.75 tok/s and a TPOT of 9.16 ms. In BS=2 mode, the model reaches 99.55 tok/s per user, a total throughput of 409.87 tok/s and a TPOT of 9.86 ms.
Moving to larger batch sizes: at BS=8, throughput is 54.60 tok/s per user, or a total of 1,383.13 tok/s and a TPOT of 10.24 ms. At BS=32, this throughput reaches 48.78 tok/s per user, or a total of 3,874.21 tok/s and a TPOT of 10.67 ms. Maximum total throughput is achieved at BS=256, with 19,114.86 tok/s, 36.38 tok/s per user and a TPOT of 20.00 ms. This represents approximately 86% of the maximum throughput in standard precision.
Microscaling Data Type Performance
Microscaling represents an advanced quantization approach that applies precise scaling factors to small blocks of weights rather than uniform quantization across large groups of parameters. NVIDIA's NVFP4 format implements this technique through a block floating-point representation where each microscaling block of 8 to 32 values shares a common exponent as a scaling factor. This granular approach preserves numerical accuracy while achieving a 4-bit representation, thereby preserving the dynamic range essential for transformer architectures. This format integrates with NVIDIA's Tensor Core architecture, enabling efficient mixed-precision computation with on-the-fly decompression during matrix operations.
We evaluated OpenAI's GPT OSS models at two parameter scales using NVFP4 quantization: the 20B variant and the larger 120B variant. These benchmark results demonstrate microscaling quantization performance across different model sizes.
GPT-OSS-20B Performance
The 20-billion-parameter model delivers the following performance across different batch sizes. In single-user mode (BS=1), it reaches 299.28 tok/s per user, or a total throughput of 943.43 tok/s and a TPOT of 3.23 ms. In BS=2 mode, the model maintains 299.19 tok/s per user, a total throughput of 1,356.87 tok/s and a TPOT of 3.19 ms.
Moving to larger batch sizes: BS=8 reaches 259.02 tok/s per user, for a total of 5,149.59 tok/s and a TPOT of 3.42 ms, while BS=16 delivers 200.69 tok/s per user, for a total of 7,765.73 tok/s and a TPOT of 3.77 ms. The model continues to improve up to BS=32 (168.34 tok/s per user, for a total of 12,411.72 tok/s) and BS=64 (123.96 tok/s per user, for a total of 16,931.47 tok/s).
Total throughput reaches 38,258.50 tok/s at BS=256, or 65.08 tok/s per user and a TPOT of 9.39 ms. This represents a 40.5x increase in total throughput compared to single-user performance. TPOT values remain between 3 and 5 ms up to BS=32. The Pareto frontier includes eight distinct points.
GPT-OSS-120B Performance
The 120-billion-parameter model maintains the following performance despite the increased parameter count: in single-user mode (BS=1), it reaches 248.62 tok/s per user, a total throughput of 783.73 tok/s and a TPOT of 3.89 ms; in BS=2 mode, it reaches 240.99 tok/s per user, a total throughput of 1,092.91 tok/s and a TPOT of 3.99 ms.
Performance continues to improve as batch size increases: BS=4 reaches 190.63 tok/s per user, for a total of 2,096.73 tok/s and a TPOT of 4.22 ms, while BS=8 delivers 172.66 tok/s per user, for a total of 3,692.10 tok/s and a TPOT of 4.54 ms. Throughput growth continues with BS=16 (138.28 tok/s per user, for a total of 5,751.41 tok/s), BS=32 (111.63 tok/s per user, for a total of 8,646.05 tok/s) and BS=64 (88.64 tok/s per user, for a total of 13,027.97 tok/s).
The model reaches its maximum total throughput at BS=256, with 29,976.99 tok/s, or 48.64 tok/s per user and a TPOT of 12.53 ms. This represents a 38.2x increase in total throughput compared to single-user performance. Maximum total throughput represents approximately 78% of the 20B model's maximum throughput. The Pareto frontier includes nine distinct points, the highest number among all tested models.
Unexpected Quantized Model Performance
