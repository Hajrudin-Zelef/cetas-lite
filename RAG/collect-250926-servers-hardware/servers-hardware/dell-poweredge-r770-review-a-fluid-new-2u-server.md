---
id: collect-250926-servers-hardware/servers-hardware/dell-poweredge-r770-review-a-fluid-new-2u-server
title: "dell-poweredge-r770-review-a-fluid-new-2u-server"
domain: servers-hardware
role: reference
task: reference
actors: ["Intel", "Nvidia"]
dates: []
keywords: ["gpus", "intel", "nvidia"]
source: docs/RAG/clean4/dell-poweredge-r770-review-a-fluid-new-2u-server.md
source_anchor: ""
source_lines: [1, 49]
sha256: 3b698e8514116dc6fb627d5da29740d4d07953d677a36efcfff1cfb15e35bb6d
---

# dell-poweredge-r770-review-a-fluid-new-2u-server

The Dell PowerEdge R770 is a 2U server that does just about everything. In a simplistic view, it is another 2U dual Intel Xeon 6 server. When we started taking the system apart, it became apparent that Dell designed an enormous amount of flexibility into this system while offering a big upgrade over the Dell PowerEdge R760. I am going to call this “fluid” because when you see how Dell’s engineers designed the system it is meant to do everything.

## Dell PowerEdge R770 External Hardware Overview

Starting off, the PowerEdge R770 is a 2U server with a fairly standard 801.51mm or 31.56″ depth. Even the depth is configurable depending on if it is a front I/O or rear I/O configuration and some other options, it can get up to just over 32″ deep.

The front of the chassis can take up to 40 E3.S SSDs, but there are also options for up to 24x 2.5″ SSDs. Our system is a 16x E3.S SSD configuration. We do not have the option, but in the rear we can get up to four E3.S SSDs as well.

In the center, instead of drive bays we have vents for airflow.

We have eight E3.S SSD slots on the right and another eight on the right.

The E3.S SSDs are the eventual replacements for 2.5″ U.2/ U.3 SSDs as we covered in 2021 with E1 and E3 EDSFF to Take Over from M.2 and 2.5 in SSDs. Admittedly, the transition is a bit slower than we would have expected.

Dell has plenty of options for SSDs. We even used some Kioxia CM7 drives that we had in the lab.

Then we get to the rear and chaos takes hold. The rear of these systems is wildly expandable so take what we have as showing a single option. We will let you look those options up, and just go through what we have here because there is plenty just with this one configuration.

On the left side, we have one power supply and a double-width PCIe riser.

In this riser, we have a NVIDIA H100 NVL. This system is designed for up to two double-width GPUs or seven 75W single-width GPUs. Some of Dell’s competitors have designs for more GPUs in 2U, but Dell also has historically had the xa series like the Dell EMC PowerEdge R750xa for four double-width accelerators in 2U.

The rear I/O is in an OCP inspired design that offers the VGA, dual USB Type-A and iDRAC port for out-of-band management.

In the center, we have both a top OCP slot as well as another full-height riser slot.

Below those center risers, we have two more full-height risers. Again, there are a ton of different options for all of these risers.

On the bottom we have the Dell BOSS which is Dell’s dual SSD boot solution. We have seen the Dell BOSS for several generations at this point. For some reason, despite the fact that I went to the same school in NJ as “The Boss” (different years), Dell’s boss always makes me think Kelis – Bossy instead. Word association aside, the advantage of this is that it allows for two boot drives without using a front hot swap bay. SSDs generally have an AFR of under 0.25%, but still this makes it easier to service without opening the chassis.

An uncommon feature on 2U servers that we have seen from Dell for some time is the handle.

On the right side, we have two more full-height slots and another power supply.

In our system, this riser is another dual width single slot riser.

Next, let us get into the server to see how Dell is making such a fluid system.

If only there was some fluid in there for cooling.

Indeed, I had assumed this was a review of one of the liquid-cooled models.

With all the dust the word “fluid” doesn’t come to mind.

Our hopes for a review of a liquid cooled server evaporated as we drank in the first paragraph. Like the bursting of a dam our hopes washed away, until we were somewhat buoyed by the review’s inclusion of E3.S SSDs instead of 2.5? SSDs (but the option to use them, for those that have them, was refreshing).

So offering help, try: “adaptive”, “versatile” or “highly configurable”.

Comments are closed.
