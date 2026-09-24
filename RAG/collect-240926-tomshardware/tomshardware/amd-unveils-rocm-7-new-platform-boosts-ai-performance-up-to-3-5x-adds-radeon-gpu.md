---
id: collect-240926-tomshardware/tomshardware/amd-unveils-rocm-7-new-platform-boosts-ai-performance-up-to-3-5x-adds-radeon-gpu
title: "amd-unveils-rocm-7-new-platform-boosts-ai-performance-up-to-3-5x-adds-radeon-gpu"
domain: tomshardware
role: reference
task: reference
actors: ["AMD", "Alibaba", "Google", "Intel", "Nvidia", "SGLang", "vLLM"]
dates: []
keywords: ["amd", "gpu", "attention", "compute", "consumer", "deepseek", "fp4", "gpus", "inference", "intel", "llama", "memory"]
source: docs/RAG/clean_en/tomshardware/amd-unveils-rocm-7-new-platform-boosts-ai-performance-up-to-3-5x-adds-radeon-gpu.md
source_anchor: ""
source_lines: [1, 71]
sha256: 0b7155fdec3e54bc1261942feb6d55297a0b2f1059949f0209b38b001796e89f
---

# amd-unveils-rocm-7-new-platform-boosts-ai-performance-up-to-3-5x-adds-radeon-gpu

<!-- source: https://www.tomshardware.com/pc-components/gpus/amd-unveils-rocm-7-new-platform-boosts-ai-performance-up-to-3-5x-adds-radeon-gpu-support -->

AMD this week introduced the 7th version of its ROCm (Radeon open compute) open-source software stack for accelerated computing that substantially improves the performance of AI inference on existing hardware compared to ROCm 6, as well as adds support for distributed workloads, and expands to Windows and Radeon GPUs. In addition, ROCm 7 adds support for FP4 and FP6 low-precision formats for the latest Instinct MI350X/MI355X processors.

The biggest change brought by ROCm 7 for client PCs is the extension of ROCm to Windows and Radeon GPUs, which allows the use of discrete and integrated GPUs for AI workloads, but only on Ryzen-based PCs. Starting in the second half of 2025, developers will be able to build and run AI programs on Ryzen desktops and laptops with Radeon GPUs, which could be a big deal for those who want to run higher-end AI LLMs locally.

One of the reasons for AMD's weak position in the AI hardware market is imperfect software. But it looks like the situation is improving as AMD's Instinct MI300X with ROCm 7 delivers over 3.5 times the inference performance and 3 times the training throughput compared to ROCm 6, according to AMD. The company conducted tests using an 8-way Instinct MI300X machine running Llama 3.1-70B, Qwen 72B, and Deepseek-R1 models with batch sizes ranging from 1 to 256, and the only difference was the usage of ROCm 7 over ROCm 6. AMD says that such improvements are enabled by enhancements in GPU utilization as well as data movement, though it does not provide any more details.

The new release also introduces support for distributed inference through integration with open frameworks such as vLLM, SGLang, and llm-d. AMD worked with these partners to build shared components and primitives, allowing the software to scale efficiently across multiple GPUs.

Furthermore, ROCm 7 adds support for lower-precision data types like FP4 and FP6, which will bring tangible improvements for the company's latest CDNA 4-based Instinct MI350X/MI355X processors as well as upcoming CDNA 5-based MI400X and next-generation Instinct MI500X-series products that will succeed the Instinct MI300-series in 2026 and 2027, respectively.

In addition, along with ROCm 7, AMD introduced its ROCm Enterprise AI MLOps solution tailored for enterprise use. The platform offers tools for refining models using domain-specific datasets and supports integration into both structured and unstructured workflows. AMD said it works with ecosystem partners to build reference implementations for applications such as chatbots and document summarization in a bid to make AMD hardware suitable for rapid deployment in production environments.

Last but not least, AMD also launched its Developer Cloud that provides ready-to-use access to MI300X hardware with configurations ranging from a single-GPU MI300X with 192 GB of memory to eight-way MI300X setups with 1536 GB of memory. For starters, AMD provides 25 free usage hours, and additional credits are available through developer programs. Early support for Instinct MI350X-based systems is also planned.

Get Tom's Hardware's best news and in-depth reviews, straight to your inbox.

*Follow* *Tom's Hardware on Google News* *to get our up-to-date news, analysis, and reviews in your feeds. Make sure to click the Follow button.*

- 
...yet there's no mention of which Radeon GPUs are supported. There's conflicting information on the website. Here this "hardware spec" list seems to imply that the stack supports most modern Radeons down to the 6600 (plus Vega, the only arch ever really supported by ROCm anyway):Reply
 https://rocm.docs.amd.com/en/latest/reference/gpu-arch-specs.html
 But this page only mentions Radeon Pro W 7000 as the cards one should buy to run ROCm 7:
 https://www.amd.com/en/developer/resources/rocm-hub/dev-ai.html#tabs-d13ff37f6d-item-94bf1a0f36-tab
Let's see if AMD has chosen to be competitive with Nvidia from 2020 or still prefers to be stuck in 2017.
- 
Reply
 The place to look will be in the release notes. I don't yet see them for release 7, but here's the place in the 6.4.1 release notes, where hardware compatibility is specified:virgult said:...yet there's no mention of which Radeon GPUs are supported.
https://rocm.docs.amd.com/projects/install-on-linux/en/latest/reference/system-requirements.html
- 
"The biggest change brought by ROCm 7 for client PCs is the extension of ROCm to Windows and Radeon GPUs, which allows the use of discrete and integrated GPUs for AI workloads,Reply***but only on Ryzen-based PCs.*** "
 
 I wonder where that comes from and if it will in fact be true. There is no technical reason why you shouldn't be able to run ROCm 7 workloads on an RX 9000 dGPU using an Intel CPU underneath.
 
 It might not be able to use whatever Intel xPU facilities reside within Intel's APUs or SoCs, but for discrete GPUs that would be emulating some of Intel's worst behavior.
 
 But what litlte information is out there only gives "positive" discrimination in that it promises support for the latest "AI" APUs and RX9000 series GPUs.
 
 I doesn't mention RDNA 3, so unless they are supported already consumer type APUs and GPUs of these older generations seem likely to miss out.
 
In any case this is only an announcement so far, nothing I'd base a purchasing decision on before it's been validated by the likes of Phoronix.
- 
I think "but only Ryzen-based PCs" means the integrated GPUs on Ryzen processors are the only ones that work with ROCm 7. The need to say such an obvious thing makes more sense when one considers that HIP is cross-platform with backends to target both Nvidia and AMD accelerators.Reply
 
Note also AMD had a line of APUs with integrated Radeon graphics before Ryzen that could to some extent be used for GPGPU computing. These will not work with ROCm.
- 
Reply
 Perhaps the "only on Ryzen-based PCs" was meant as a qualification only applying to the integrated GPUs?abufrejoval said:"The biggest change brought by ROCm 7 for client PCs is the extension of ROCm to Windows and Radeon GPUs, which allows the use of discrete and integrated GPUs for AI workloads,***but only on Ryzen-based PCs.*** "
 
I wonder where that comes from and if it will in fact be true.
- 
Reply
 HIP is designed as a cross-platform solution, but I have no idea about the current state of its support for running atop Intel iGPUs.ejolson said:The need to say such an obvious thing makes more sense when one considers that HIP is cross-platform with backends to target both Nvidia and AMD accelerators.
 
 I consider HIP a vendor-specific API, even though the theoretical (and maybe rather real) possibility exists to use it atop other hardware. We cannot depend on them to maintain other backends with the same attention to quality, reliability, accuracy, and functional completeness as they do for their own hardware.
 
If I truly wanted a hardware-agnostic solution, I would opt for something that's not backed by a specific vendor, because they're always going to be biased towards their own hardware.
- 
Reply
 Another possibility is that ROCm 7 just doesn't support Intel CPU's. This doesn't mean it wouldn't work and most likely wouldn't be an issue at all. If this is the case and not the Radeon iGPU theory, there needs to be pushback on this as vendor lock-in is something we need to get away from in the world of AI compute (yes, I'm attacking CUDA) and compute in general.abufrejoval said:"The biggest change brought by ROCm 7 for client PCs is the extension of ROCm to Windows and Radeon GPUs, which allows the use of discrete and integrated GPUs for AI workloads,***but only on Ryzen-based PCs.*** "
 
 I wonder where that comes from and if it will in fact be true. There is no technical reason why you shouldn't be able to run ROCm 7 workloads on an RX 9000 dGPU using an Intel CPU underneath.
 ...
 
ROCm 7 is apparently still in preview, and I'm not seeing a lot of new details and clarifications from AMD since this news article was posted.
- 
Reply
 Let's just hope for the best...DS426 said:Another possibility is that ROCm 7 just doesn't support Intel CPU's. This doesn't mean it wouldn't work and most likely wouldn't be an issue at all. If this is the case and not the Radeon iGPU theory, there needs to be pushback on this as vendor lock-in is something we need to get away from in the world of AI compute (yes, I'm attacking CUDA) and compute in general.
 
ROCm 7 is apparently still in preview, and I'm not seeing a lot of new details and clarifications from AMD since this news article was posted
