---
id: collect-261001-cisco/cisco/questions-1790-cisco-nexus-7000-sup-2-and-2e-why-no-cmp-f04e33c7
title: "questions-1790-cisco-nexus-7000-sup-2-and-2e-why-no-cmp-f04e33c7"
domain: cisco
role: reference
task: reference
actors: ["Intel"]
dates: []
keywords: ["intel", "reasoning", "research"]
source: docs/RAG/collect-261001-cisco/questions-1790-cisco-nexus-7000-sup-2-and-2e-why-no-cmp-f04e33c7.md
source_anchor: ""
source_lines: [1, 22]
sha256: 0a9f348ee3fc92c491540aa62c1bb71f90b3da2159a35d903f0fd15e7f985ecc
---

# questions-1790-cisco-nexus-7000-sup-2-and-2e-why-no-cmp-f04e33c7

Network Engineering is part of Stack Overflow’s open communities: specialist spaces where curiosity is welcome, knowledge is shared freely, and the best answers rise to the top.
This question shows research effort; it is useful and clear
9
This question does not show any research effort; it is unclear or not useful
Save this question.
Show activity on this post.
I am looking at an upgrade from Nexus 7000 Supervisor 1's to Supervisor 2E's.
One of the differences that I am seeing is a lack of the CMP (Connectivity Managment Processor) and the CMP port. Even on the config guide for CMP, here, it says:
Note: The CMP is available only on the Supervisor 1 modules, not on
the Supervisor 2 nor Supervisor 2E modules.
Currently we're utilizing those ports for OOB management, but will lose that functionality in the Sup-2E's. It isn't the end of the world, but I will need to document/work-around this in the implementation plan.
I'm assuming this was an intentional choice on Cisco's part, so does any one know why, or have any documentation with the reasoning behind this choice?
Friend of mine asked this question in Vegas in Cisco Live and Cisco cited lack of customer demand and thermal as main reasons.
CMP is absolutely the right solution and we need it in every vendor switch and router, only proper OOB. Server guys have had this for over decade.
So please start adding this as scoring item in your RFQ, so vendors know there is demand.
Because of the thermal comment I emailed Freescale (most routers and switches use their SOC in control-plane) and asked if they are planning to support anything like Intel vPro, which helps implement OOB functionality...
From our marketing team updates I heard something similar will be
implemented in the next generation of our multicore processors, but
this is under discussion yet, so I do not have any details. Ok, I will
pass your suggestion to our core design team, hope they will take it
into account.
There is no guarantee whatsoever that Cisco would implement OOB on Nexus 7k in future supervisors, but I thought it is interesting that Freescale is considering adding a feature like this to their silicon.
