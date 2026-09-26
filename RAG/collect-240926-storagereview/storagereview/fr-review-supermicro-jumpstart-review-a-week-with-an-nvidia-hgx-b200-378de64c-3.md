---
id: collect-240926-storagereview/storagereview/fr-review-supermicro-jumpstart-review-a-week-with-an-nvidia-hgx-b200-378de64c-3
title: "fr-review-supermicro-jumpstart-review-a-week-with-an-nvidia-hgx-b200-378de64c"
domain: storagereview
role: reference
task: reference
actors: ["Meta", "Nvidia", "vLLM"]
dates: []
keywords: ["nvidia", "benchmark", "benchmarks", "compute", "fp4", "fp8", "gpu", "inference", "latency", "llama", "nvfp4", "parameters"]
source: docs/RAG/clean_en/storagereview/fr-review-supermicro-jumpstart-review-a-week-with-an-nvidia-hgx-b200-378de64c.md
source_anchor: ""
source_lines: [23, 57]
sha256: 48f2f615437da7b0273309da82f7f3cd94347781f9bb448778f635e15c8b995c
---

# fr-review-supermicro-jumpstart-review-a-week-with-an-nvidia-hgx-b200-378de64c

The real acceleration occurred from eight threads onward, with most workloads leveling off around 30 GiB/s. Large blocks performed best, with consistently excellent results for 5 MB and 10 MB sizes regardless of thread count.
Throughput ultimately peaked at around 43 GiB/s with a 10 MB block size at 256 threads, representing the highest sustained sequential read rate we observed in this test.
GDSIO Sequential Read Latency
In terms of latency, the workload started with excellent responsiveness: single-thread reads ranged from 0.06 to 0.1 ms for small blocks. As thread count increased, latency gradually decreased, remaining below 1 ms up to 8 threads for most workloads.
Beyond 16 threads, larger block sizes began to produce latencies of several milliseconds as the storage path became increasingly saturated. Peak latency was observed at the end of the test, with a 10 MB block size and 256 threads, spiking at just over 1.2 seconds (about 1,200 ms), corresponding to an extreme load designed to saturate the system.
GDSIO Sequential Write Throughput
For sequential writes, performance was much more stable than for reads. The workload leveled off quickly, with most thread and block size combinations settling around 6.3 to 6.5 GiB/s. This indicates that write throughput hits a consistent ceiling very early, likely tied to the storage medium and buffering rather than GPU or PCIe limitations.
Increasing thread count had no significant impact, with throughput remaining virtually unchanged between 2 and 128 threads. The only notable result occurred at the end of the test: with a 10 MB block size and 256 threads, throughput peaked at 18.2 GiB/s, demonstrating a slight advantage when the system fully exploits deep queues and write aggregation.
GDSIO Sequential Write Latency
Write latency was initially relatively low, ranging from 0.15 to 0.5 ms for single-threaded workloads with small blocks. As thread count increased, latency grew much faster than for reads, reaching 1 to 4 ms at four threads and 4 to 9 ms at eight threads.
Once we reached 32 threads with larger blocks, latency rose sharply, with 5 MB and 10 MB blocks entering the 170 to 350 ms range. The most extreme case was 10 MB blocks with 256 threads, which peaked at just under 3 seconds (about 2,900 ms), clearly illustrating how quickly the write path saturates under heavy parallel load.
GDSIO Random Read Throughput
For random reads, the workload ramped up quickly. Single-thread performance ranged from about 11 to 31 GiB/s, depending on block size, with larger blocks immediately benefiting from higher bandwidth. With two and then four threads, throughput reached 20 to 36 GiB/s, demonstrating excellent scalability from the earliest stages.
From 8 threads onward, the system leveled off around 30 GiB/s, a result very similar to that observed in the sequential read test. The best result was achieved with a 10 MB block size at 256 threads, peaking at about 42.7 GiB/s.
GDSIO Random Read Latency
Random read latency was initially very low, on the order of 0.15 to 0.4 ms for single-threaded workloads with small blocks. As thread count increased, latency rose progressively, remaining below 1 ms up to 4 threads and reaching about 1 to 3 ms at 8 threads.
Once we moved to 32 threads and larger blocks, latency increased more sharply, with 5 MB and 10 MB transfers reaching values between 30 and 55 ms. The most extreme case was observed with 256 threads and a 10 MB block size, where latency peaked at just over 1.1 seconds (about 1,180 ms).
GDSIO Random Write Throughput
Random write performance was very consistent, with most block sizes and thread counts settling around 5.8 to 6.1 GiB/s. The workload reached this level almost instantly and improved only marginally with additional threads, indicating that the write path quickly reaches its limit.
The only notable exception appeared at the end of the test, where a 10 MB block size at 256 threads briefly reached 12.5 GiB/s, likely benefiting from deep queuing and write aggregation under heavy parallel load.
GDSIO Random Write Latency
Random write latency was initially relatively low, with single-core performance ranging from about 0.6 to 1.2 ms for small blocks and 5 to 12 ms for the largest transfers. As concurrency increased, latency rose rapidly, reaching 4 to 9 ms at eight threads and 18 to 38 ms at 32 threads.
Beyond that threshold, the write path became saturated. With larger blocks, latency reached 150 to 380 ms at 64 threads, and continued increases in block size drove latency up dramatically. The worst case was observed with 10 MB blocks and 256 threads, with peak latency of about 4.4 seconds.
vLLM Online Serving – LLM Inference Performance
vLLM is the most popular high-throughput inference and serving engine for LLMs. The vLLM online serving benchmark is a performance evaluation tool that measures the engine's real serving capabilities under concurrent requests. It simulates production workloads by sending requests to a running vLLM server with configurable parameters, such as request rate, input/output length, and number of concurrent clients. The benchmark measures key metrics including throughput (tokens per second), time to first token, and time per output token (TPOT), allowing users to understand vLLM performance under different load conditions.
We tested inference performance across a comprehensive suite of models covering various architectures, parameter scales, and quantization strategies to evaluate throughput under different concurrency profiles.
Dense Model Performance
Dense models follow the classic LLM architecture, where all parameters and activations are used during inference, resulting in more compute-intensive processing than their sparse counterparts. To comprehensively evaluate performance across different models based on size and quantization strategies, we compared several dense model configurations from the Llama 3.1 8B family.
Our test suite included evaluations of Meta Llama 3.1 8B in three precision formats: the standard configuration, plus FP8 and FP4 quantized versions using NVIDIA's NVFP4 format. It's important to note that vLLM currently uses the Marlin kernel for NVFP4 quantized models, and the optimal performance gains from this quantization format are not yet fully visible in these benchmarks. Future vLLM optimizations targeting native NVFP4 tensor core operations could deliver additional performance improvements. This model selection strategy enables direct performance comparison while isolating the impact of progressive quantization on inference throughput.
Llama 3.1 8B Performance
The Llama 3.1 8B model, in standard precision, exhibits the following scaling characteristics based on concurrency level. In single-user mode (BS=1), it reaches 279.27 tok/s per user, for a total throughput of 1,727.62 tok/s and a TPOT of 3.37 ms. As batch size increases, per-user throughput decreases while total throughput increases. At BS=8, the model reaches 82.85 tok/s per user, for a total throughput of 3,386.48 tok/s and a TPOT of 3.28 ms. Performance continues to improve up to BS=32 (56.46 tok/s per user, 8,274.66 tok/s total) and BS=64 (52.70 tok/s per user, 13,707.66 tok/s total).
The model reaches its maximum total throughput at BS=256, with 32,797.67 tok/s, or 30.64 tok/s per user and a TPOT of 16.13 ms. This represents a 19x increase in total throughput compared to single-user performance. TPOT values remain between 3 and 4 ms up to BS=64, then increase to 16.13 ms at BS=256.
Llama 3.1 8B FP8 Performance
The FP8 quantized variant exhibits different characteristics. At BS=1, it reaches 149.46 tok/s per user, with a total throughput of 9,565.20 tok/s and a TPOT of 3.44 ms. Pareto analysis reveals three optimal points (BS=1, BS=128, BS=256).
At BS=128, the model delivers 44.12 tok/s per user with a total of 19,198.40 tok/s and a response time of 11.04 ms.
