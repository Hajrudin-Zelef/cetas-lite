---
id: collect-250926-servers-hardware/servers-hardware/they-let-me-bring-a-camera-into-a-top-classified-us-supercomputer-el-capitan
title: "they-let-me-bring-a-camera-into-a-top-classified-us-supercomputer-el-capitan"
domain: servers-hardware
role: reference
task: reference
actors: ["AMD", "California", "Nvidia", "United States"]
dates: []
keywords: ["compute", "amd", "energy", "gpu", "hbm", "liquid cooling", "memory", "nvidia", "open source", "power delivery"]
source: docs/RAG/clean4/they-let-me-bring-a-camera-into-a-top-classified-us-supercomputer-el-capitan.md
source_anchor: ""
source_lines: [1, 51]
sha256: 36dc07e4fb783191e3ba94c2ff24cefc6e0ef5bd08b640bc11ade7e5d75db326
---

# they-let-me-bring-a-camera-into-a-top-classified-us-supercomputer-el-capitan

Today was a big day. The US Department of Energy’s new top supercomputer and the #1 machine on the Top500, El Capitan, had its dedication ceremony at LLNL in California. After the dedication ceremony, they let me loose inside El Capitan with my phone. Normally, this is an area where electronics are not allowed, and phones do not get to enter. Today, I got in.

## The El Capitan Dedication Ceremony

After submitting my information to get validated for an on-site visit, we were then given a badge and escorted to an auditorium. There, the LLNL, NNSA, DoE, and other folks gave talks about El Capitan.

Apparently Antonio Neri, HPE’s CEO lived in Livermore for a period of time. I asked Antonio about how this work on El Capitan (and other HPC clusters) translates to AI sales. He said that all of the underlying technologies including the GPU compute, networking, liquid cooling, power delivery, and so forth, being deployed at scale, directly translate to AI clusters.

AMD’s CEO, Lisa Su brought a de-lidded AMD Instinct MI300A and pulled it out at the podium. Lisa also had an insightful answer to the question of how this translates to AI. Her response was that this was yet another proof point of deployment and operation at scale with well over 40,000 accelerators in El Capitan. That means the AMD and HPE teams needed to design for reliability to operate the system on simulations that can take months.

For folks working on the project, it has been a long time. I remember being invited to HPE’s headquarters just before the pandemic for a small room where the HPE-Cray and AMD Win was announced. Anything pre-pandemic feels like ancient history at this point.

In today’s presentations, we also got a quick look at some of the simulations run like this one showing a shocked Tin surface across 139 billion zones on only 2048 of the nodes in El Capitan. For those who do not know, El Capitan’s mission is classified, but generally it is to support US weapons programs. That makes seeing inside the system unusual.

Of course, this is STH, so I was eagerly awiting a promised chance to see the system. To my surprise, at the electronics drop area, they allowed me to bring my phone and to take photos. At first I thought I would just get to see the nodes, which were impressive. CoolIT is providing the liquid cooling blocks. I was able to show and hold a (very heavy) Frontier node in Calgary during our How Liquid Cooling is Prototyped and Tested in the CoolIT Liquid Lab Tour a few years ago.

The annotations, by the way are:

1. Node
2. SIVOC (Power Regulator)
3. Slingshot NIC Mezzanine Card
4. Cold Plate
5. AMD Instinct MI300A APU.

There was also a tray with four populated APU sockets and four unpopulated socket only nodes. Unlike most systems today, each socket and package is an integrated set of chiplets spanning CPU cores, GPU cores, and high-bandwidth memory (HBM) so we have a uniform set of sockets and without DIMM slots on the sides of each socket.

Luckily, that was not all on display as El Capitan was open just as it was about to start its classified mission. Next, let us get to the running system.

I was a system administrator for a supercomputer center at Lockheed in the late 80s. We had a Cray X-MP, a Y-MP, a Connection Machine and several DEC VAXs computers as front end compute nodes. It was an exciting time with so much raw computational power that today can be had in a smart phone.

Your visit to Aurora brought back lots of great memories.

But you had to leave your packet capture gear at home, right?

Some seriously nice kit, thanks for sharing.

If you ever find yourself in Lugano, Switzerland then you could attempt a visit to CSCS to see their HPC/Cray machines

You are missing a ‘Us’ in the title of yhis article :)

And I can’t spell ‘this’ :(

The Rabbit software is entirely open source

https://nearnodeflash.github.io/latest/

Any insights into why AMD won this vs NVIDIA back in 2020?

Probably one good reason is that AMD committed to the HPC market and FP64 while NVIDIA has been focused on mixed precision

Comments are closed.
