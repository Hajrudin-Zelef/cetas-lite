---
id: collect-240926-storagereview/storagereview/fr-review-scaling-ai-checkpoints-the-impact-of-high-capacity-ssds-on-model-train-5dfc8c98-3
title: "fr-review-scaling-ai-checkpoints-the-impact-of-high-capacity-ssds-on-model-train-5dfc8c98"
domain: storagereview
role: reference
task: reference
actors: ["Nvidia"]
dates: []
keywords: ["benchmark", "cost", "gpu", "gpus", "latency", "memory", "nvidia", "parameters", "training"]
source: docs/RAG/clean_en/storagereview/fr-review-scaling-ai-checkpoints-the-impact-of-high-capacity-ssds-on-model-train-5dfc8c98.md
source_anchor: ""
source_lines: [27, 49]
sha256: 02fb640fce26185e93e43656064329587adb617d090d70528e0a848330eee033
---

# fr-review-scaling-ai-checkpoints-the-impact-of-high-capacity-ssds-on-model-train-5dfc8c98

For enterprises looking to minimize checkpoint intervals, the TLC-based Gen5 PS1010 offers the advantage of providing the fastest processing time. If the goal is to retain many checkpoints cost-effectively, the QLC-based Gen4 P5336 can do so. We measured an average checkpoint time difference of less than 17% between the two drives in passes two and three.
GPUDirect Storage bandwidth
Although DLIO shows flash performance in an AI workflow, the workload is entirely write-based until a checkpoint is restored. To paint a more complete picture of the Solidigm D7-PS1010 and D5-P5336 in AI workloads, we included read bandwidth measurements using GDSIO.
How GPUDirect Storage works
Traditionally, when a GPU processes data stored on an NVMe drive, the data must first pass through the CPU and system memory before reaching the GPU. This process introduces bottlenecks, because the CPU becomes an intermediary, adding latency and consuming valuable system resources. GPUDirect Storage eliminates this inefficiency by allowing the GPU to access data directly from the storage device via the PCIe bus. This direct path reduces the overhead associated with moving data, enabling faster and more efficient data transfers.
AI workloads, especially those involving deep learning, are extremely data-intensive. Training large neural networks requires processing terabytes of data, and any delay in data transfer can lead to GPU underutilization and longer training times. GPUDirect Storage addresses this challenge by ensuring that data is delivered to the GPU as quickly as possible, minimizing idle time and maximizing computational efficiency.
As with the DLIO test, the goal is to better understand and characterize the differences between high-speed Gen5 SSDs and high-capacity QLC drives. Not all AI workloads are the same, and each drive offers distinct advantages, depending on the need.
Test configuration matrix
We systematically tested each combination of the following parameters with an NVIDIA L4 in our test platform:
- Block sizes: 1 M, 128 K, 64 K, 16 K, 8 K
- Number of threads: 128, 64, 32, 16, 8, 4, 1
- Number of jobs: 16
- Batch sizes: 16
Our first overview was the QLC-based D5-P5336, which reached 4.2 GiB/s with a transfer size of 1 M at an I/O depth of 128. The effect of block sizes produced a substantial increase in bandwidth, from 8 to 32 M. The benefit of increased I/O depth began to diminish at 32, where workloads began to stabilize.
Next, we look at the Gen5 PS-1010, which can scale up to 6.2 GiB/s with a block size of 1 M and an I/O depth of 128. Overall, it outperformed the Gen4-based P5336, with particular workloads demonstrating substantial improvement. One notable area of improvement came at the 128 K block size, where with an I/O depth of 64 and 128, the PS1010 delivered twice the read bandwidth of the P5336.
It is important to note that both SSDs were tested with the NVIDIA L4 GPU. Although the Gen4 D5-P5336 is a high-end model, higher-performance NVIDIA GPUs, such as the H100, have demonstrated better efficiency with the D7-PS1010. Drive speed is the primary selection criterion for some customers, while others prioritize overall density. Solidigm offers solutions for both, with its QLC and TLC SSDs.
Conclusion
As the scale and complexity of AI training continue to grow, the underlying storage infrastructure must not only keep pace, but also set the tone. Our tests with two very different SSDs illustrate the importance of aligning storage solutions with specific training priorities, such as minimizing checkpoint latency or maximizing checkpoint density for cost-effective scalability.
In our evaluation, we tested the Solidigm D5-P5336 (61.44 TB) and the D7-PS1010 (7.68 TB) under realistic AI training conditions using the DLIO benchmark and a full hybrid-parallel LLM checkpointing workflow. We captured measurements reflecting checkpoint write performance across multiple runs as the drives filled, highlighting the differences in completion time between the QLC-based Gen4 D5-P5336 and the TLC-based Gen5 D7-PS1010.
While the D7-PS1010 offered the fastest possible checkpoint writes, the D5-P5336 demonstrated compelling advantages in cost-effectiveness and capacity with only a modest performance decline. We also examined GPUDirect Storage read bandwidth with GDSIO via an NVIDIA L4 GPU. We found that the Solidigm D5-P5336 delivered up to 4.2 GiB/s of read bandwidth with a transfer size of 1 MB, while the D7-PS1010 delivered a substantial increase up to 6.2 GiB/s. You will get even higher performance by leveraging an even larger GPU, such as the NVIDIA L40s or H100/H200.
Looking ahead, the unprecedented capacity of the 122 TB Solidigm D5-P5336 SSD is expected to reshape AI training and deployment. As model sizes and checkpoint requirements continue to increase, these massive drives pave the way for new levels of efficiency and flexibility, enabling training strategies that were previously out of reach. Solidigm's leadership in high-capacity SSD solutions enables organizations to store more data and checkpoints on fewer drives and helps future-proof their infrastructure for the next wave of AI complexity.
Solidigm D5-P5336 122 TB SSD
This report is sponsored by Solidigm. All views and opinions expressed in this report are based on our impartial assessment of the product(s) under study.
