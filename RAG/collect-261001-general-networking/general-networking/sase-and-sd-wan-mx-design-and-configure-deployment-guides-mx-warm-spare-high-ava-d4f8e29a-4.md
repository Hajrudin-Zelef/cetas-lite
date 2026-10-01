---
id: collect-261001-general-networking/general-networking/sase-and-sd-wan-mx-design-and-configure-deployment-guides-mx-warm-spare-high-ava-d4f8e29a-4
title: "sase-and-sd-wan-mx-design-and-configure-deployment-guides-mx-warm-spare-high-ava-d4f8e29a"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-general-networking/sase-and-sd-wan-mx-design-and-configure-deployment-guides-mx-warm-spare-high-ava-d4f8e29a.md
source_anchor: ""
source_lines: [134, 159]
sha256: 7ab5cd5d7dfb9259acfbaaee2ee827d4a0bfba3e474b8d6a129ac99d21cdab51
---

# sase-and-sd-wan-mx-design-and-configure-deployment-guides-mx-warm-spare-high-ava-d4f8e29a

  - If both appliances are consistently reporting in the "active" state, check their LAN connection and make sure they can communicate with each other.
  - If the spare MX is intermittently reporting as active while the primary remains online and active, check that both MXs can communicate with each other on all VLANs. Additionally, ensure there are no bad cables connecting the two devices or any other physical issue that could result in unreliable communication.
  - Take a packet capture on the LAN side of each MX, to get a clear picture of where the VRRP heartbeats are being lost (for packet capturing guidance, refer to Packet Capture Overview article).
- If the HA pair is configured to use a virtual IP on the uplink, make sure that each pair of WAN connections (WAN 1 on each MX, for example) share the same broadcast domain so they can both be seen by the upstream device.
Firmware Upgrade Behavior
When MX appliances are configured to operate in High Availability (HA) (either in NAT/routed mode or when operating as one-armed VPN concentrators), the dashboard will automatically take steps to ensure a zero-downtime MX upgrade. This is achieved through the following automated process:
- 
    The primary MX downloads firmware.
- 
    The primary MX stops advertising VRRP.
- 
    The secondary MX becomes active.
- 
    The primary MX reboots.
- 
    The primary MX comes online again.
- 
    The primary MX starts advertising VRRP again.
- 
    The primary MX becomes active again.
- 
    The secondary MX downloads firmware* (approximately 15 minutes after the original upgrade is scheduled; the 15-minute delay for the secondary MX is a safety buffer to ensure the primary is stable).
- 
    The secondary MX stops advertising VRRP.
- 
    The secondary MX reboots and comes back online. Note: The secondary MX will attempt to upgrade regardless of whether the primary MX upgrade was successful or not.
