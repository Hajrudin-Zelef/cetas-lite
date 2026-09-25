---
id: collect-250926-servers-hardware/servers-hardware/nvidias-groq-3-lpu-accelerators-for-heterogeneous-ai-compute-at-hot-chips-2026
title: "nvidias-groq-3-lpu-accelerators-for-heterogeneous-ai-compute-at-hot-chips-2026"
domain: servers-hardware
role: reference
task: reference
actors: ["Groq", "Nvidia"]
dates: []
keywords: ["compute", "nvidia", "agentic", "attention", "benchmarks", "decode", "disaggregated", "fp8", "gpu", "gpus", "inference", "kv cache"]
source: docs/RAG/clean4/nvidias-groq-3-lpu-accelerators-for-heterogeneous-ai-compute-at-hot-chips-2026.md
source_anchor: ""
source_lines: [1, 81]
sha256: b2771d5ea661961aa2f89ce6fb7350670bf14be71c91b1594de37fdae70df4bf
---

# nvidias-groq-3-lpu-accelerators-for-heterogeneous-ai-compute-at-hot-chips-2026

The second AI presentation of the afternoon comes from NVIDIA, who besides doing talks on the Vera CPU and Rubin GPU, are also presenting a talk on the use of language processing units (LPUs) in their Vera Rubin racks.

This talk is arguably a bit of an unusual one, because NVIDIA does not currently produce their own LPUs (though they are under development). The LPUs being used in Vera Rubin generation racks – and specifically the dedicated LPX racks – come from Groq, whom NVIDIA is buying the chips from as part of a broader acquihire of the company. None the less, with the bulk of Groq’s talent now at NVIDIA, it falls to NVIDIA to promote the chips.

Groq’s LPUs are designed to fill a weak spot in NVIDIA’s Vera Rubin stack. GPUs are great for pre-fill, and depending on how the design of the chip is optimized, so-so at decode. Whereas LPUs – essentially purpose-built chips with large amounts of on-die SRAM to keep model data as close to the compute hardware as possible. Fittingly, the biggest rationale for the use of LPUs in a Vera Rubin server cluster is to boost performance at low latencies, taking advantage of that SRAM to get results back ASAP so that AI models can move on to the next token.

This article is being written live from the presentation, so please excuse any typos.


## NVIDIA’s Groq 3 LPU Accelerators for Heterogeneous AI Compute at Hot Chips 2026

Yesterday NVIDIA released the first third-party benchmarks of an LPX rack, so the timing of that and this talk is not coincidental. Nor is the fact that NVIDIA has been promoting the use of LPUs in both their Vera and Rubin talks. Today’s talk is reiterating parts of that for the Hot Chips crowd, as well as laying out the case for using LPUs with Vera Rubin and how they fit in to the broader ecosystem.

Recapping a common refrain through all of NVIDIA’s talks at Hot Chips, the company considers agenetic AI to be the most complex computing workload in history. And to that end, it has required new approaches to hardware development – as well as a whole lot more hardware.

LPUs are the building block of the LPX rack, which is one of several racks that make up the larger Vera Rubin hardware ecosystem.

And once again showing off NVIDIA’s performance curves/frontiers for the Vera Rubin ecosystem, plotting tokens/sec/watt versus tokens/sec/user. The Rubin GPU hardware has its own performance curve, but overall performance can drop pretty hard if you try to push higher user interactivity (i.e. lowering token latency). This is where the LPX rack comes in, adding a second layer of hardware to significantly boost the per-user token rate and reduces the latency accordingly. Or as the Groq team puts it: this is ludicrous mode.

Agentic AI systems need a lot of processing time due to the large amounts of context tokens in play. Decode is the largest portion of the compute workload in term of the amount of time taken. And that is where the LPU comes in to handle this portion of the inference task more efficiently and quickly.

Presenting at Hot Chips is apparently a bit of a “pinch me” moment for the Groq team members who moved over to NVIDIA.

Diving into the LPX rack, a single rack can decode 11,000 tokens per second on a 31B model (Gemma 4).

NVIDIA this week published third-party benchmarks, which were conducted by Artificial Analysis. The LPX rack offed a 4x higher token output rate than the next-closest public competitor, which NVIDIA doesn’t name but looks to be a Cerebas CS-3.

Diving into the LPU architecture, one of the key elements of the chip design is close connection between memory and compute. Groq has put quite a bit of SRAM very close to their compute elements.

This is a fully deterministic chip. There is no branching or other sources of non-determinism.

Because the chip is deterministic, the instruction scheduling is handled in software as well. Which means there’s no need for significant scheduling hardware on the chip.

And here is a step-by-step example of how a thread is scheduled and executed over an LPU, with an emphasis on how threads can be handled concurrently.








Ultimately an LPU isn’t going to be working alone. It’s going to be thousands of chips working together, effectively acting as a single massive LPU core.

Looking at the LPU networking, each chip acts as both a processor and a router. Again owing to the determinism, the compiler schedules all resources, both processing and networking.

Interestingly, there is no hardware flow control or virtual channels here. Groq keeps things very bare and basic.

Getting back to the physical world, there are 256 LPUs in a single LPX rack. This combines to make for a 128GB of SRAM for memory, offering a total of 40PB/second of aggregate SRAM bandwidth. Or in terms of compute, there are 315 PFLOPS of FP8 compute performance.

One focus has been on ensuring the LPX rack and trays are optimized for mass production and are reliable once assembled. The LPU team is using NVIDIA’s MGX rack architecture here, which means using a standardized and well-tested piece of hardware across NVIDIA’s ecosystem.

The emphasis on determinism also means that power consumption is very deterministic. That means that the LPX can employ look-ahead techniques to smooth power consumption, and even order current from power regulators ahead of time so that it can be available when it’s needed in a moment. The net effect is that it cuts down on both droop (60%) and overshoot, flattening out the overall curve.

(And the audience just applauded at this)

The software backing the LPX also takes into consideration thermal considerations to better spread out the workload over the chip to avoid thermal hot spots and thermal throttling. This being another benefit of deterministic execution.

And that’s the LPU hardware in a nutshell.

Looking at the bigger picture, LPX is but one part of the larger Vera Rubin hardware ecosystem. NVIDIA will be using the LPXes in conjunction with their GPU hardware to maximize overall performance. Which means co-designing the GPU and LPU.

Getting the LPUs and GPUs talking to each other and efficiently interacting is trickier than it may first appear. The LPUs are in a synchronous domain, but talking to the GPUs or going to an external KV cache is an async operation. So NVIDIA employs a FPGA to function as an async bridge between the two worlds.

Larger mixed clusters also have to account for the fact that not everything within LPX’s domain is perfectly deterministic. So there is an need to handle static scheduling of dynamic workloads.

LPU and GPU racks keep their own KV caches. Only draft tokens are exchanged between the two racks.

NVIDIA offloads the attention portion of the decode process back to the GPUs, making decode a disaggregated process.

And, of course, prefill and decode are disaggreated as well, with prefill taking place in the GPUs while most of the decode process takes place in the LPUs.

NVIDIA uses micro-batches of workloads that are executed concurrently on the GPUs and LPUs. This allows NVIDIA to overlap these workloads and their communication to help keep up the utilization of the hardware.

The integration of LPUs into the NVIDIA hardware ecosystem means that the CUDA software platform needs to be similarly updated. NVIDIA has adding LPU support to CUDA ecosystem, making them a fully-supported target for CUDA.

Looking at the culmination of the LPU hardware, the rest of the Vera Rubin hardware, and NVIDIA’s software changes significantly alters the performance curve/pareto frontier. Vera Rubin can push a lot of tokens overall at low interactivity, but tapers off quickly at higher interactivity rates. Combining this with the LPUs extends these curves further, and the more that is offloaded to the LPUs the higher the interactivity rates gets, up to a 5x improvement at the highest token/user/second rate. Though the trade-off is that total throughput efficiency is dropping the more the LPUs are used; GPUs are still the king of efficiency when total throughput is all that matters and high latencies are acceptable.

And that is the Groq 3 LPU, and NVIDIA’s Vera Rubin LPX rack. The disaggregation unlocks new performance possibilities for the NVIDIA ecosystem, and gives the platform the tools needed to offer much faster low-latency inference than what GPUs can provide on their own.

We have a full recap on the Substack, including where all of these pieces ranked in terms of popularity:
