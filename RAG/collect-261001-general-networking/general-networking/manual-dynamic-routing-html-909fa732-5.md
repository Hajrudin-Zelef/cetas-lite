---
id: collect-261001-general-networking/general-networking/manual-dynamic-routing-html-909fa732-5
title: "manual-dynamic-routing-html-909fa732"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-general-networking/manual-dynamic-routing-html-909fa732.md
source_anchor: ""
source_lines: [299, 311]
sha256: 27df09740b8b7b9296f5be629ad202220ce8a0f8520bc94612705268b172bb76
---

# manual-dynamic-routing-html-909fa732

Bidirectional Forwarding Detection (BFD) is a lightweight protocol used to detect faults between routers or switches by sending periodic Hello packets (asynchronous mode). BFD quickly identifies failing links, making it a useful companion to routing protocols like OSPF and BGP for faster convergence.
STATIC (Static Routes Daemon)
| Options | Description | 
|---|---|
| Enable | This will activate the staticd service | 
| Options | Description | 
|---|---|
| Enabled | Enable/Disable | 
| Network | Defines the target for the static route, in CIDR notation. | 
| Gateway (optional) | Optional gateway IP address for this route. | 
| Interface | Select the interface where this setting applies. | 
| BFD | Mark the route as dependent on the BFD neighbor session with the next hop. | 
STATIC is a daemon that handles the installation and deletion of static routes. These routes can be used supplemental to dynamic routes. It is beneficial for fine grained control over routes in more complex network environments, if redistributing directly attached routes is not an option.
