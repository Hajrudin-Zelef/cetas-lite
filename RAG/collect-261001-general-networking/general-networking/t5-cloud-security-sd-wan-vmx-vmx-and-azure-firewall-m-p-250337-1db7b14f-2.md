---
id: collect-261001-general-networking/general-networking/t5-cloud-security-sd-wan-vmx-vmx-and-azure-firewall-m-p-250337-1db7b14f-2
title: "t5-cloud-security-sd-wan-vmx-vmx-and-azure-firewall-m-p-250337-1db7b14f"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-general-networking/t5-cloud-security-sd-wan-vmx-vmx-and-azure-firewall-m-p-250337-1db7b14f.md
source_anchor: ""
source_lines: [162, 182]
sha256: d73b779ec0daba24ccbd7880f145968c626a3d2b1b8987ebfcd94c54d48c984d
---

# t5-cloud-security-sd-wan-vmx-vmx-and-azure-firewall-m-p-250337-1db7b14f

Does your NVAs support BGP peering? If so you could set up eBGP peerings between the NVAs and the vMX pair in azure. Once the packets reach the NVA it would know how to get to the azure spokes.
If you can't change your VPN Gateway to fith the RS requirements then it's not much else you can do.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
10-08-2024 10:14 AM
Thanks for this - Is it worth as this is in lab to look at utilising vWAN instead get that set up and then add in a firewall and adjust the routing.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
10-08-2024 10:57 AM
Depends on your architecture. But VWAN can get expensive fast. If you just plan on using meraki as a vpnc which acts as the only entrypoint to your environment i would say that VWAN is overkill.
Also, for existing azure environments the route server setup is easier to "shim" in without uprooting the entire architecture. Just my two cents
