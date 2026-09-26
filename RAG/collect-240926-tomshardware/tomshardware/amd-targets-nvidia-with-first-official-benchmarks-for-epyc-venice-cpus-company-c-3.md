---
id: collect-240926-tomshardware/tomshardware/amd-targets-nvidia-with-first-official-benchmarks-for-epyc-venice-cpus-company-c-3
title: "amd-targets-nvidia-with-first-official-benchmarks-for-epyc-venice-cpus-company-c"
domain: tomshardware
role: reference
task: reference
actors: ["AMD", "Intel", "Meta", "Nvidia", "OpenAI"]
dates: []
keywords: ["amd", "nvidia", "dram", "gpu", "gpus", "intel", "lpddr5x", "memory", "nvlink"]
source: docs/RAG/clean_en/tomshardware/amd-targets-nvidia-with-first-official-benchmarks-for-epyc-venice-cpus-company-c.md
source_anchor: ""
source_lines: [110, 124]
sha256: 137bbfa51db47ef8d5a591944229558cdd48226f2da99face9f2981b67d52d84
---

# amd-targets-nvidia-with-first-official-benchmarks-for-epyc-venice-cpus-company-c

 Well, you're counting DRAM power for GPUs, but not CPUs. So, it's a mismatched comparison. I've heard that the cumulative power utilization of server memory can even surpass that of the CPU, although I doubt that will still hold true for CPUs burning 650W to 700W.Stomx said:If CEOs of AMD and Intel ever try to make their morning tea/coffee with 600W boiler instead of standard 1500W they'd probably finally understand that :)
 
 Obviously, Nvidia's Vera is an exception, due to its use of LPDDR5X and lower total capacity.
 
 
 Um, what? A PCIe 6.0 x16 slot is good for 128 GB/s per direction and each CPU has ~128 lanes (i.e. 1024 GB/s in each direction).Stomx said:With NVIDIA, you buy an 8-GPU motherboard and then populate it with 1, 2, 4, or all 8 GPUs (16 were possible before, too) — the 900 GB/s NVLink is already there. With CPUs all we ever had was a dual-socket design with a barely-there ~40 GB/s, and even that will most probably go extinct.
 
That said, you know about AMD's MIx00A series, right? They have infinity fabric links for running in a mesh, similar to NVLink. Some future version of Intel Xeons will have NVLink integrated directly.
- 
Reply
 AMD does that, too.LordVile said:Doesn’t really matter how much better they are than Nvidia’s when Nvidia gives their products out for free, sorry for pinky promised payments and stock.
 https://www.tomshardware.com/tech-industry/amd-to-supply-anthropic-with-2-gigawatts-of-instinct-mi450-gpus
 They have different sorts of deals with OpenAI and Meta, however. Also, I'm not aware of them doing anything like the VC fund that Nvidia raised to finance their customers.
 
So, you're right that the scale on which Nvidia is financing the purchase of its products is unmatched.
