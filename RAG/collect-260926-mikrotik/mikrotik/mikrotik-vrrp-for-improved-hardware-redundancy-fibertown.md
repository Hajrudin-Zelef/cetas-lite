---
id: collect-260926-mikrotik/mikrotik/mikrotik-vrrp-for-improved-hardware-redundancy-fibertown
title: "mikrotik-vrrp-for-improved-hardware-redundancy-fibertown"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: ["preemption"]
source: docs/RAG/lot-mikrotik/forum/dhcp-dns/mikrotik-vrrp-for-improved-hardware-redundancy-fibertown.md
source_anchor: ""
source_lines: [1, 31]
sha256: bb936d6cd26c81ea7809a04bc80c5a961419c52322eb444028f846d04e57b70a
---

# mikrotik-vrrp-for-improved-hardware-redundancy-fibertown

I hear plenty of data center buzz words…and what good colocation sales guy doesn’t work one into every other sentence? One of those phrases you hear is “concurrently maintainable.” This means you can sustain loss in your infrastructure and still be up. We are talking about hardware redundancy.

If you aren’t using BGP (a protocol backing the core routing decisions on the Internet), you would do well to use VRRP (a redundancy protocol designed to increase the availability of the default gateway servicing hosts on the same subnet).

You have two routers connected to the same layer 2 segment. You have a subnet configured that is /29 or larger. You configure a physical IP on the interfaces, then you create a VRRP interface on each router associated with those connected interfaces. You then assign the same VRRP IP address on both routers to the VRRP interface.

The VRRP router that has the higher priority(default is 100) is the master. The master responds to ARP requests for the VRRP IP. If the master router fails, then the backup router takes over and owns the VRRP IP.

Your default gateway points towards the VRRP IP so that if the master fails and the backup takes over, your default route is still valid! There is also a concept of premption. By default, preemption will migrate the VRRP IP over to the router with the highest priority.


So here’s our demo configuration:


So what happens when one of our providers fail?


*Provider fails on one link. The backup guy takes over the VRRP IP. Our default route points to 10.0.0.1 so we still route out!*


We drop half of our network gear, but have no fear. The ISP was pointing towards 10.0.0.6 to route to me, so all is good in the hood.

**Router 10.0.0.4**

Create the VRRP interface *assign it higher priority – default is 100*:

This is great for the WAN side, but is quite often used for the LAN also. You can also run two separate VRRP groups on a single interface which will allow you to load balance with redundancy.

Go forth and route my peoples!

(Learn more about Mikrotik at my blog)
