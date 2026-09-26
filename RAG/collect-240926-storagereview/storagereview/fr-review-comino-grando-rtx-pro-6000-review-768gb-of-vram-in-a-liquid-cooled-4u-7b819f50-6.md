---
id: collect-240926-storagereview/storagereview/fr-review-comino-grando-rtx-pro-6000-review-768gb-of-vram-in-a-liquid-cooled-4u-7b819f50-6
title: "fr-review-comino-grando-rtx-pro-6000-review-768gb-of-vram-in-a-liquid-cooled-4u--7b819f50"
domain: storagereview
role: reference
task: reference
actors: ["Anthropic", "MiniMax", "Nvidia"]
dates: []
keywords: ["agent", "claude", "compute", "cost", "decode", "gpu", "gpus", "inference", "liquid cooling", "llama", "memory", "moe"]
source: docs/RAG/clean_en/storagereview/fr-review-comino-grando-rtx-pro-6000-review-768gb-of-vram-in-a-liquid-cooled-4u--7b819f50.md
source_anchor: ""
source_lines: [143, 152]
sha256: b3ca38a80de5c17ad7d01268d5c700b26de8468cce3a083a2e9907a8dae46b60
---

# fr-review-comino-grando-rtx-pro-6000-review-768gb-of-vram-in-a-liquid-cooled-4u--7b819f50

Under an equal workload (256/256), the model starts at 16.35 tok/s at BS=1, reaches 2,751.25 tok/s at BS=64, and scales sharply with higher concurrency, peaking at 5,753.24 tok/s at BS=256. With heavy prefill (8k/1k), throughput starts at 606.97 tok/s at BS=1, increases steadily to 5,351.02 tok/s at BS=32 and 6,557.92 tok/s at BS=64, peaking at 7,357.26 tok/s at BS=128 before declining slightly to 7,140.74 tok/s at BS=256, suggesting the model approaches saturation in prefill throughput beyond BS=128. Decode-heavy (1k/8k) increases steadily from 82.21 tok/s at BS=1 to 1,485.28 tok/s at BS=64, peaking at 2,554.87 tok/s at BS=256, reflecting the expected memory bandwidth requirements of a 230B MoE architecture under sustained decode workloads.
Conclusion
The Comino Grando is above all a system designed to fully harness the potential of eight NVIDIA RTX PRO 6000 GPUs. Every major design choice, from the inverted motherboard layout to the cooling system and integrated monitoring platform, aims to ensure continuous operation of these GPUs at their maximum 600 W TDP, without thermal or power constraints.
What makes the Grando so compelling is not any single feature, but the coherence of the entire system. Liquid cooling is not a mere add-on; it is an integral part of the architecture. The power supply is redundant, hot-swappable, and sized to support the 4,800 W load of eight 600 W cards, with comfortable headroom. The monitoring system does not merely display temperatures; it automatically protects the hardware in the event of a problem. Here, nothing appears to have been added hastily.
The observed performance confirms this coherence. Across a varied range of models, from Llama 3.1 8B to MiniMax M2.5 230B, the Grando delivered thoroughly satisfactory throughput for a self-hosted platform. Claude Code concurrency tests highlighted its practical value: eight engineers can simultaneously run multi-agent coding sessions on a locally hosted 230B model, at interactive speeds, with per-user throughput exceeding 38 tok/s at peak aggregate production. Teams of four to eight people can work at near-optimal throughput without perceptible degradation.
The value of this configuration extends beyond AI inference. With 96 GB of VRAM per GPU and a high-performance multi-GPU architecture, the platform is perfectly suited to demanding creative and engineering workloads, such as VFX rendering, large-scale simulation, and complex CAD pipelines. The system is scalable to four- and two-GPU configurations, making this level of performance accessible to smaller studios and teams that nonetheless need workstation-class compute density.
Where the Grando most distinguishes itself from the enterprise eight-GPU platforms we tested is in its ease of deployment. Those systems offer greater PCIe lane capacity, more NIC slots, and more extensive storage connectivity, but they require dedicated data center infrastructure, consume well over 8 kW, and their lead times can exceed a year. The Grando prioritizes a system quiet enough to be installed in the same room, generates less heat, and is available immediately, at the cost of some peripheral expansion capability. For organizations that prioritize rapid deployment and easy-to-manage operating environments over maximum network connectivity, this trade-off is advantageous.
Product Page – Comino Grando
Comino Configurator – Page
Ranking: The Comino Grando RTX PRO 6000 takes first place in the ranking of the best desktop computers for local AI.
