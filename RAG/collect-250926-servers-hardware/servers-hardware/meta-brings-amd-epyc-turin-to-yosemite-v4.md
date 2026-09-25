---
id: collect-250926-servers-hardware/servers-hardware/meta-brings-amd-epyc-turin-to-yosemite-v4
title: "meta-brings-amd-epyc-turin-to-yosemite-v4"
domain: servers-hardware
role: reference
task: reference
actors: ["AMD", "Intel", "Meta"]
dates: []
keywords: ["amd", "compute", "intel"]
source: docs/RAG/clean4/meta-brings-amd-epyc-turin-to-yosemite-v4.md
source_anchor: ""
source_lines: [1, 15]
sha256: 17b952e2485e43e14e5d948e74f1905349015b89cf1089dbbf9360148c1f36b0
---

# meta-brings-amd-epyc-turin-to-yosemite-v4

Going through photos from OCP Summit 2024, we found this tasty one. Yosemite v4 was first detailed at OCP Summit 2023. This year, Meta showed off the new compute chassis, saying it would start deploying the AMD EPYC Turin systems along with CXL in early 2025.

## Meta Brings AMD EPYC Turin to Yosemite v4

The Meta Yosemite V4 is the company’s newest scale-out chassis architecture. In 7OU of rack space, it incorporates 8, 6, or 4 compute nodes. These compute nodes can vary in factors like the amount of local storage and accelerators they offer. Still, at the OCP Summit, we saw an 8-node configuration.

The bottom of the chassis has the management module and the networking tray. One of the big reasons we have so many multi-host adapters on the market today is that the teams at Meta drove that with their Yosemite program. These adapters help lower power requirements and costs versus traditional blade servers and even multi-node setups.

On the rear, we get fans and an OCP Rack V3 power tap. I believe the breakout for power is an assembly called Medusa.

## Final Words

Yosemite has been a major platform at Meta / Facebook for many years. It was big news when we covered the Meta AMD EPYC North Dome CPU and Platform Details three years ago. Meta had been a longtime Intel shop, making that announcement a big deal at the time. Now, we have at least some decent clue that AMD will be making its way into Yosemite v4 and with CXL support. Or at least, in the Meta booth, the placard on Yosemite v4 said that it would have a CXL-expansion-enabled AMD Turin version.

If you want to learn more about the new AMD EPYC 9005 Turin generation, you can see more here:
