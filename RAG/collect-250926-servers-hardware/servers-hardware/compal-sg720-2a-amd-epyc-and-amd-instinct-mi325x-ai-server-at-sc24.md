---
id: collect-250926-servers-hardware/servers-hardware/compal-sg720-2a-amd-epyc-and-amd-instinct-mi325x-ai-server-at-sc24
title: "compal-sg720-2a-amd-epyc-and-amd-instinct-mi325x-ai-server-at-sc24"
domain: servers-hardware
role: reference
task: reference
actors: ["AMD", "Nvidia"]
dates: []
keywords: ["amd", "gpu", "gpus", "nvidia", "nvlink"]
source: docs/RAG/clean4/compal-sg720-2a-amd-epyc-and-amd-instinct-mi325x-ai-server-at-sc24.md
source_anchor: ""
source_lines: [1, 23]
sha256: f134c812eaf2362e7a44a463be68fbcfd6beed01a9ce5d02fd73a667c3f9bd97
---

# compal-sg720-2a-amd-epyc-and-amd-instinct-mi325x-ai-server-at-sc24

We are gearing up for a fun week but wanted to quickly cover a server we did not show previously from SC24. This is the Compal SG720-2A, an AMD EPYC and AMD Instinct MI325X AI server we saw on the show floor.

## Compal SG720-2A AMD EPYC and AMD Instinct MI325X AI Server

Here is the quick spec sheet, which includes SSD storage, and many expansion slots. The prominent feature is that this AMD EPYC 9005 Turin system supports up to 500W TDP CPUs.

The front of this server is a fan aficionado’s dream. The top portion is three rows of fans, followed by the U.2 drive bays, and then more fans.

Here is what the other side of the fan wall looks like.

The bottom section has the U.2 drive bays, but then a neat feature. This is actually a front I/O server with the management port and front KVM hookups. On either side, we have OCP NIC 3.0 slots.

Turning to the rear of the server, we have a big GPU tray on top. Usually top-tier designs have GPUs on top since that minimizes the PCIe trace lengths. Likewise, a tray design is important to ensure that systems are easy to service while in racks.

On the bottom rear, we have six power supply slots which is a 5+1 design. The PCIe slots are also on serviceable trays.

As an AMD Instinct MI325X (and MI300X) server, this offers something different than the NVIDIA HGX platforms. AMD’s platform does not have the NVLink switch architecture, and instead uses the OCP UBB.

Still, the two CPU and eight GPU option might be attractive to many users.

## Final Words

We are going to hear plenty around NVIDIA servers since we are in the middle of doing in-depth reviews of three NVIDIA HGX H200 platforms. We wanted to at least take a moment to show the AMD competition given that.
