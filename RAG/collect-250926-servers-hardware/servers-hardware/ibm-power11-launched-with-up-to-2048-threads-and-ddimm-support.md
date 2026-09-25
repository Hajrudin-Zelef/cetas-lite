---
id: collect-250926-servers-hardware/servers-hardware/ibm-power11-launched-with-up-to-2048-threads-and-ddimm-support
title: "ibm-power11-launched-with-up-to-2048-threads-and-ddimm-support"
domain: servers-hardware
role: reference
task: reference
actors: []
dates: []
keywords: ["accelerator", "cost", "custom silicon", "memory", "rack-scale"]
source: docs/RAG/clean4/ibm-power11-launched-with-up-to-2048-threads-and-ddimm-support.md
source_anchor: ""
source_lines: [1, 49]
sha256: bb05e6e71ca3f216aa5bc4536b1a542ea3fa82ecfa97a6c4e18f593d6609a5f0
---

# ibm-power11-launched-with-up-to-2048-threads-and-ddimm-support

IBM has a new generation of POWER with the new Power11 line of servers. Instead of doing multiple launches in different segments, the company is showing off a top-to-bottom stack ranging from lower-cost single CPU servers up to 16-socket 2048-thread full rack servers.

## IBM Power11 Launched with Up To 2048 Threads and DDIMM Support

Last week, I sat on a pre-brief call for the Power11 launch, and there was not a lot of focus on the hardware. This, however, is STH. So, of course, we will augment. IBM was focused on providing a newer, more modern platform versus the IBM POWER10 we covered in 2020, which was for 2021-2022 servers. It is fun as we even had a early “blue door studio” version of that announcement. That is not just across the hardware layers, but also all of the enterprise software and services that go around these hardware platforms.

The goal of Power11 is to provide a reliable and secure platform for enterprises that is faster than the previous generation.

One of the big moves in this generation is to also focus on AI. A big one here might be the IBM watsonx Code Assistant for i. One challenge in an x86 and Arm dominated world (or maybe it is CUDA?) is getting folks to program for another architecture. AI coding assistants have the capability to open up architectures beyond traditional ones, especially as those coding assistants are rapidly leveling up.

Part of this is alos IBM’s custom accelerators. We went into the IBM Spyre AI accelerator at Hot Chips last year and got to see these during the IBM z17 launch. IBM gets to leverage its investment in these cards beyond IBM Z mainframes with its POWER servers now.

IBM also has a hybrid cloud offering with IBM PowerVS where an organization can have on-prem and cloud servers. For x86 and Arm, there are plenty of clouds that have modern servers to run on. For IBM POWER, hybrid cloud would be less straightforward if IBM did not have PowerVS.

In the Q&A slide, there was a bit about the hardware. I made my own slide based on that sidebar. IBM is launching a family from an entry-class 2U system to a rack-scale system in this generation.

Starting with the IBM Power E1180 we have four sockets per system, and four connected together for up to 16 socket systems. Four of these system nodes with four CPUs of 16 cores each, and 8 threads per core give us 256 cores and 2048 threads in the scale-up system.

Somethign fun about these IBM Power11 systems is the RAM. Instead of standard DDR5 modules, these use DDIMMs. We covered OpenCAPI memory as a precursor to CXL in standard servers but this is one way IBM is able to get the density they need in these systems.

The next system is the IBM Power E1150. This is the four socket server with up to 64 OMI memory channels for 64 of those DDIMMs.

One the smaller side, there is the 4U IBM Power S1124. This is a dual socket server that is designed for operating reliability at a smaller scale.

Just to give you some sense of reliability built into these systems, IBM has a feature called Spare Cores. If there is an issue with a core, then fully functional idle cores that are present but not active can become active to swap into the system.

There is also a smaller edge server, the IBM Power S1122 that is only 2U but reduces some features like having half of the memory chapacity of the S1124.

Many years ago when there was a direction with IBM Power9 of selling scale-out systems, we looked at hosting STH on IBM Power9. Here is the CPU socket of one of the IBM LC921 servers that we bought to host STH.

Maybe one day we need to open one of those systems up to show you inside. One of the fun bits is that the LC921 and LC922 were actually made by Supermicro in the Power9 generation.

## Final Words

IBM told us in a call before the announcement that although they were still getting new customers, the main focus was selling into the existing installed base. There are still a lot of shops out there that run IBM POWER. Providing those customers with not just the hardware to get more performance and features in newer systems, but also the software and services aroudn it will continue to make this a neat business going forward.

Of course, if you want to see some other cool IBM engineering including the IBM Fishkill facility used to develop many of its high-end custom silicon efforts, feel free to check out pieces we did earlier this year like the The IBM z17 Mainframe Brings AI with Telum II and Spyre and The New IBM z17 Telum II Processor Module Cut Open Down to Silicon.

It is great hardware, but the software licensing and software maintenance costs are way over the top for just about anything you want to run on it. Yes, there is an installed base of POWER globally, and there is an appetite for running applications at the tier levels these systems support.

But there is no ecosystem to foster knowledge or understanding outside of the blue circle surrounding this platform. This is what keeps it so niche and the resources to maintain them so expensive.

I can’t quite support that. I am an admin of exactly these servers (from Power7-Power11) . You can use the server with e.g. AlmaLinux or other open platform software without high software software lizensen. But they wanted Enterprise Support.

The IBM Power Platform, for example, has established the Openpower foundation.

Crank that POWER to 11 baby!! All jokes aside, this is actually really neat. And great timing to boot as I added a POWER9 Sforza processor to my CPU collection today, completing my set of POWER9s (I’ve had a Monza for a while now and I got a Lagrange last year). Next up, POWER10, one each of the SCM and DCM variants. C’mon eBay, make it happen!

Comments are closed.
