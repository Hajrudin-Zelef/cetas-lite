---
id: collect-240926-tomshardware/tomshardware/amd-s-256-core-epyc-9996-venice-claims-up-to-a-3-4x-jump-over-intel-xeon-competi-2
title: "amd-s-256-core-epyc-9996-venice-claims-up-to-a-3-4x-jump-over-intel-xeon-competi"
domain: tomshardware
role: reference
task: reference
actors: ["AMD", "AWS", "China", "Google", "Intel", "Meta", "Moonshot", "Nvidia", "TSMC"]
dates: []
keywords: ["amd", "intel", "2nm", "3nm", "agent", "agentic", "agents", "agi", "aws", "benchmark", "benchmarks", "disaggregated"]
source: docs/RAG/clean_en/tomshardware/amd-s-256-core-epyc-9996-venice-claims-up-to-a-3-4x-jump-over-intel-xeon-competi.md
source_anchor: ""
source_lines: [89, 148]
sha256: 1ca8b01395c4c001e72102494cb4a4f420cd034baa287b47e04a3867ac496cb0
---

# amd-s-256-core-epyc-9996-venice-claims-up-to-a-3-4x-jump-over-intel-xeon-competi

In vectorized workloads, AMD claims a 1.6x gen-on-gen improvement and 2.3x improvement compared to the Xeon 6980P. For this test, AMD used Meta’s open-source FAISS (Facebook AI Similarity Search) library to search for similar vectors in the siftm1 dataset. The metric here is QPS, or queries processed per second, looking at overall query throughput.

| **Chip** | **QPS** | 
| Intel Xeon 6980P | 316,069 | 
| AWS Graviton5 | 119,179 | 
| AMD Epyc 9755 | 369,252 | 
| AMD Epyc 9965 | 472,079 | 
| AMD Epyc 9996 | 751,453 | 

For its “enterprise tools” benchmarks, AMD ran several tests, including TPC-H, TPC-C, and Redis, and it reports the results as “geomean throughput.” We have actual numbers here, but they’re a geomean representing several different tests rather than a single benchmark. Broadly, however, AMD claims a 1.6x gen-on-gen improvement in these workloads, and a 2.6x improvement compared to Intel.

| **Chip** | **Geomean throughput** | 
| Intel Xeon 6980P | 2,284,701 | 
| AWS Graviton5 | 2,982,203 | 
| AMD Epyc 9755 | 2,546,290 | 
| AMD Epyc 9965 | 3,867,149 | 
| AMD Epyc 9996 | 6,054,748 | 

A lot of agentic workloads are applicable outside of agents, but AMD also tested a few agents directly. It replayed five different agent personas across the chips and, once again, gathered a throughput geomean. We don’t have the metrics here, nor for the previous benchmark, so it’s possible there’s an angle of performance that we’re not seeing with the data provided by AMD.

Regardless, the company claims a 1.5x gen-on-gen improvement in this test, and a 2.5x improvement compared to the 6980P.

| **Chip** | **Geomean throughput** | 
| Intel Xeon 6980P | 1.779 | 
| AWS Graviton5 | 2.505 | 
| AMD Epyc 9755 | 2.317 | 
| AMD Epyc 9965 | 2.97 | 
| AMD Epyc 9996 | 4.451 | 

AMD ran these tests earlier in the month. But just a few days ago, Nvidia published its __first SPEC CPU 2026 results for Vera__. AMD ran some tests of its own using the same compiler for a comparison between Vera and Venice. AMD’s Kuppuswamy says, “everything is apples-to-apples comparison, same compiler.” That’s GNU 15.2, if you’re curious.

In throughput, AMD claims a 2.2x improvement in the SPECrate integer suite, compared to Vera using the dense Venice design with 256 Zen 6c cores. More importantly, AMD claims a 1.2x improvement in per-core performance when comparing Vera to a 96-core “High Frequency” Venice chip. AMD says it used Nvidia’s results as the basis for comparison. With both Venice designs, AMD used a 600W TDP.

In SPEC CPU 2017 (again using SPECrate with integer workloads), AMD has data comparing Venice to Intel’s 6980P and __Arm’s new AGI__, showing 2x throughput compared to Intel, and 1.3x per-core performance. Note the core counts here for AMD. SPECrate is a throughput test, and AMD stepping down to a 128-core model suggests that performance will likely drop off as the core count increases.

Although AMD wants to focus Venice performance on agentic workloads, it shared a range of what are now being called “legacy” workloads across the cloud and HPC. Some of the results are repeated from the earlier slides, such as Redis and NGINX, but there are some additional data points, including NAMD and SQL. The performance improvements here are large, though not surprising. You can see that across these tests, even Turin beats the competition from Intel and AWS.

*Follow* *Tom's Hardware on Google News**, or* *add us as a preferred source**, to get our latest news, analysis, & reviews in your feeds.*

- 
Reply
 I don't know how you get to that takeaway. If you normalise the core count differences against the power usage and then the additional performance by AMD considering the extra power used it is basically a wash in terms of efficiency, even though AMD is using a disaggregated design against a monolithic design (which Nvidia made a whole song and dance about when comparing against a 2 year old Zen 5 part). It seems like AMD have some answers to Vera ready now and further answers with Verano.Gururu said:All need to just bow to nVidia already. It's going to be all over within two years with total domination of entire processor market.
 
It'll be interesting to see what Intel can do because they're at risk of Nvidia squeezing them out too and it just being a race between Nvidia and AMD.
- 
Reply
 The article says N2 (2nm):jackt said:idk... in H2 2027 ? 3nm ? it risks to born old.
 Zen 6 is built on TSMC’s N2 (this has been previously confirmed). AMD confirmed that there are 32 cores on a CCD, along with two IODs. Keep in mind that the 32-core CCD is using Zen 6c, not full Zen 6.
- 
Reply
 They have way too much money and the GPU market is already nearly a monopoly. It looks like they already got a CPU light years ahead of anything China can make and just two steps from AMD and Intel, who can't touch nVidia for pure resources at the moment. The AI surge is depressing and nVidia is at the forefront, but if they can shift their incalculable resources towards CPUs while AI is hot, AMD and Intel are dead.richardnpaul said:I don't know how you get to that takeaway. If you normalise the core count differences against the power usage and then the additional performance by AMD considering the extra power used it is basically a wash in terms of efficiency, even though AMD is using a disaggregated design against a monolithic design (which Nvidia made a whole song and dance about when comparing against a 2 year old Zen 5 part). It seems like AMD have some answers to Vera ready now and further answers with Verano.
 
It'll be interesting to see what Intel can do because they're at risk of Nvidia squeezing them out too and it just being a race between Nvidia and AMD.
- 
Somebody who has subscription to Phoronics can you tell them ( I tried several times to subscribe but this site was created so long ago, that it conflicts and refuses to take my money) when doing in the future comparisons to include also performance with the local AI models like GLM5.2 and Kimi k3.Reply
 
 Interesting how much larger memory bandwidth and new binary AVX512bmm will translate into inference token per second on decent 4-5 bit models. Turin has 5 t/s on llama.cpp and nearly lossless Q4 GLM5.2 but ideally needs at least 3x more ( of course it's faster on lower quantization ).
 
 Other than that not much new is in "consumers" Venice vs Turin. Max core count of 9756 is the same 128, the TDP is the same 500W. Clock is even lower. Yes memory bandwidth is larger but will 2-3x bandwidth translate in at least 20% more performance? In AI hopefully yes. Not very likely everywhere else, the socket-to-socket bandwidth is still barely doubled to PCIe6. Based on first impressions and memory market conditions, I'd even call it, sorry for my French, a fart in public place. NVIDIA does not afraid to reach 1400W in TDP of their GPUs, while all processor manufacturers behave like NVIDIA inflicted them an inferiority complex.
 
