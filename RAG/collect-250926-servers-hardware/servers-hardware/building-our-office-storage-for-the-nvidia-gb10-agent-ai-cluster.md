---
id: collect-250926-servers-hardware/servers-hardware/building-our-office-storage-for-the-nvidia-gb10-agent-ai-cluster
title: "building-our-office-storage-for-the-nvidia-gb10-agent-ai-cluster"
domain: servers-hardware
role: reference
task: reference
actors: ["AMD", "Apple", "Nvidia"]
dates: []
keywords: ["agent", "nvidia", "amd", "gpu", "gpus", "memory", "nand"]
source: docs/RAG/clean4/building-our-office-storage-for-the-nvidia-gb10-agent-ai-cluster.md
source_anchor: ""
source_lines: [1, 23]
sha256: c7ac1e4bf4bb08ee4170dd42456f8c9985de2e81653dffb006964e42b1390b10
---

# building-our-office-storage-for-the-nvidia-gb10-agent-ai-cluster

In the STH studio, we have been working on building a neat cluster. While we have big GPUs that can run small models with great speed, we have instead been focusing on building a cluster of different types of edge AI systems, all with a common thread: they have high memory capacity and can run larger models. The cluster has had 4-5 NVIDIA GB10 machines, plus four AMD Ryzen AI Max+ 395 128GB systems, and an Apple Mac Studio M3 Ultra 512GB. We started a storage solution for the NVIDIA GB10 cluster, and then we hooked other machines up to it. As such, it was time for a little write-up.

As a quick note, we are using SSDs from Solidigm, a NAS from QNAP, NVIDIA GB10’s from NVIDIA and Dell, a MikroTik switch from MikroTik, and likely other hardware provided by other vendors. We have to say this is sponsored. This involved some ideas, plus quite a bit of trial and error to get to the point that we had a solution that worked how we wanted.

## Building Our Office Agent AI Cluster Storage for NVIDIA GB10

First off, why are we doing this project? One of the big reasons is that we have had 4-5 NVIDIA GB10 boxes with high-speed networking. An advantage of hosting models and also fine tuning model data over the network, instead of locally, is that the $1000 delta between 4TB and 1TB nodes starts to get less exciting as the number of machines scales. A 60GB model at these storage prices costs roughly $20 to store locally. When you have five machines that becomes $100 just to store a single model. There are also larger models, and you may have many of these downloaded onto a machine.

Also, the Apple Mac Studio with the M3 Ultra, which we use for high memory capacity in a single node, costs the same $1000 to move from 1TB to 4TB of storage.

That only has 10Gbase-T networking built in, but we have 25GbE Thunderbolt adapters which make storage reasonably fast.

We also have 10GbE or faster networking on five AMD Ryzen AI Max+ 395 servers that we have in the studio.

That does not include the GPU workstations and servers that we have for higher performance, but lower memory capacity.

A few weeks ago, we did a piece on Building New STH Studio Storage NAS. In that, we mentioned that we were building not just a storage repository for footage, but we were also using this for AI.

Key to this effort has been the Solidigm D5-P5336 SSDs. While these are far from the fastest SSDs on the market, they allowed us to scale the storage capacity to a decent figure while also not being wildly expensive for capacity.

Although these use QLC NAND, our workloads tend to be much heavier on the read side. Just to give you some sense of the impact versus the hard drives, loading a model like gpt-oss-120b from this Solidigm D5-P5336 versus an older hard drive-based QNAP TS-1655 12-Bay Plus 4-Bay NAS can often see 30-60% reductions in model load times. When you need to load an use a model, that can be a very notable difference.

Still, to get this working takes a bit of work, so let us get to that next.
