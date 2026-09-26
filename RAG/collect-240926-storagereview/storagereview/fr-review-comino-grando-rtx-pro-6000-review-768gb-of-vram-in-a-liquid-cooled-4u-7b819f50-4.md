---
id: collect-240926-storagereview/storagereview/fr-review-comino-grando-rtx-pro-6000-review-768gb-of-vram-in-a-liquid-cooled-4u-7b819f50-4
title: "fr-review-comino-grando-rtx-pro-6000-review-768gb-of-vram-in-a-liquid-cooled-4u--7b819f50"
domain: storagereview
role: reference
task: reference
actors: ["AMD", "Anthropic", "MiniMax", "Nvidia", "OpenRouter", "vLLM"]
dates: []
keywords: ["amd", "benchmarks", "claude", "cost", "gpu", "gpus", "inference", "latency", "memory", "moe", "nvidia", "opus 4"]
source: docs/RAG/clean_en/storagereview/fr-review-comino-grando-rtx-pro-6000-review-768gb-of-vram-in-a-liquid-cooled-4u--7b819f50.md
source_anchor: ""
source_lines: [76, 100]
sha256: 081b955baf4af2ba28c9d4c0ca908726e3777be89019f6b357ce9aedbafbcb8d
---

# fr-review-comino-grando-rtx-pro-6000-review-768gb-of-vram-in-a-liquid-cooled-4u--7b819f50

While our tests focused on AI inference workloads, the Grando proves equally useful for creative professionals and engineers who need significant local GPU computing power. Its cumulative 768 GB of VRAM, spread across eight RTX PRO 6000 cards, offers capabilities that classic workstation configurations cannot match.
VFX artists and motion graphics professionals can now render complex scenes with massive texture sets entirely in VRAM, eliminating the disk-swapping bottlenecks that affect productions using 8K sequences or high-resolution environments. CAD engineers performing computational fluid dynamics or structural simulations can process assemblies of unprecedented complexity without having to partition their models into multiple runs. Video editors working with multi-stream 8K RAW timelines, colorists applying machine-learning-based noise reduction at full resolution, and 3D artists performing final ray-traced rendering locally rather than waiting for cloud server availability all benefit from this density of memory and GPU computing power.
The Grando does not require a full eight-GPU configuration. Comino offers the platform in four, six, and eight GPU configurations, all available immediately. Small studios, independent creators, and engineering teams can thus tailor their investment to their current needs while benefiting from a simple upgrade path to accompany their business growth.
Platform Trade-offs: Density vs. Expandability
The Grando's compact design offers exceptional GPU density and thermal management in a standard 4U form factor, but this density implies architectural trade-offs that should be understood before deployment.
The chassis is compatible with EATX and EEB form factor motherboards, but not with the extended server motherboards found on classic dual-processor platforms. This limits the total number of PCIe lanes available for peripherals other than GPUs. In our eight-GPU configuration, the 128 PCIe Gen 5 lanes of the AMD EPYC processor are almost entirely used by the GPUs, leaving little bandwidth for additional NVMe storage or high-speed network connectivity beyond the integrated 10 GbE ports.
This contrasts with the eight-GPU platforms we tested from Dell, HPE, and Supermicro. These systems use larger chassis, dual-socket configurations, and PCIe switch topologies to support significantly greater peripheral connectivity. They typically accommodate four to eight additional network cards or data processing units (DPUs), in addition to the full set of GPUs, as well as eight or more hot-swappable NVMe bays, making them perfectly suited for distributed inference workloads requiring high-bandwidth interconnects.
However, this increased capability comes at a significant cost. Power consumption exceeds 8 kW. Thermal loads require dedicated data center cooling infrastructure. Noise levels prevent any deployment outside specially designed machine rooms. Additionally, lead times often extend from six to eighteen months due to persistent supply constraints on enterprise GPU platforms.
The Grando occupies a different position. For organizations that prioritize rapid deployment, easy-to-manage operating environments, and inference or creative workloads rather than large-scale distributed training, the trade-offs are often advantageous. Teams that need their hardware immediately, in an operational environment, may find the Grando's approach to density more practical than waiting for a platform they cannot deploy effectively once available.
Comino Grando Performance Test Results
System Configuration
- Chassis: Comino Grando
- Motherboard: ASRock GENOAD8X-2T/BCM support
- CPU: AMD EPYC 9474F 48C
- Memory: 512 GB DDR5 memory
- GPU: 8 x NVIDIA RTX PRO 6000
- Storage: M.2 SSD
Claude Code Service – MiniMax M2.5
Beyond traditional raw LLM inference benchmarks, we wanted to evaluate this hardware's performance in an automated coding workflow, specifically by managing multiple simultaneous Claude Code sessions using a locally hosted model. This use case has a direct impact on development team productivity: how many engineers can simultaneously use an AI coding assistant hosted on a single node before the experience degrades?
To test this, we designed a test environment that generates a dataset of medium-difficulty programming problems (such as implementing an LRU cache, creating a command-line task management application, writing a Markdown converter, and building a REST API) and runs each Claude Code session in a separate Docker container, interacting with the local vLLM server. A transparent proxy, placed between the sessions and the inference endpoint, captures per-request metrics for each Claude Code instance. The model used was MiniMax M2.5, deployed via vLLM on the system's eight NVIDIA RTX PRO 6000 GPUs. Although it is not the highest-ranked coding model in public leaderboards, M2.5 is a capable model that many users, including our developer colleagues, use locally.
To establish a baseline, we use the average output throughput of Anthropic's Claude Opus 4.6, measured via OpenRouter.ai, one of the most widely used routing services for production API access. This baseline throughput corresponds to approximately 37 tokens per second per API request.
We measured two key indicators: the average number of output tokens per second per Claude Code session (what each developer experiences) and the total number of output tokens per second across all sessions (the total work produced by the server).
Based on the results, a single simultaneous Claude Code session delivers a throughput of 67.3 tok/s per user and an aggregate throughput of 64.7 tok/s. With two simultaneous sessions, per-instance throughput drops slightly to 57.4 tok/s, while aggregate throughput reaches 95.1 tok/s, with vLLM's batching beginning to amortize overhead. Four simultaneous sessions maintain a throughput of 49.2 tok/s per user, providing a still highly responsive experience for interactive coding workflows, while aggregate throughput reaches 177.2 tok/s. Eight sessions represent the optimal configuration for aggregate throughput, peaking at 206.7 tok/s, while per-instance throughput stabilizes at 38.7 tok/s, a comfortable level for real-time code generation and iteration.
With 16 simultaneous sessions, the system exhibits the classic batching trade-off: per-instance throughput drops to 31.1 tok/s and aggregate throughput to 105.8 tok/s. This suggests that at this level of concurrency, the 230B MiniMax M2.5 model reaches the limits of what eight GPUs can support without introducing significant latency for each user. The overall decline observed between 8 and 16 sessions reflects the memory bandwidth requirements of a large MoE architecture under significant simultaneous decoding load, rather than a scheduling inefficiency.
For organizations evaluating self-hosted AI infrastructure for their development tools, the Grando is a compelling solution. Equipped with a cutting-edge 230B model, it can comfortably handle up to eight simultaneous Claude Code sessions with throughput providing genuine interactivity, and per-user speeds exceeding 38 requests/second at peak aggregate production. Teams of four to eight engineers can thus work at near-optimal throughput without perceptible loss of responsiveness.
