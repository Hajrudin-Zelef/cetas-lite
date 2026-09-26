---
id: collect-260926-mikrotik/mikrotik/questions-867418-limit-bandwith-of-a-wireless-ap-connected-to-mikrotik-routerboa-1773057e
title: "questions-867418-limit-bandwith-of-a-wireless-ap-connected-to-mikrotik-routerboa-1773057e"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/lot-mikrotik/forum/routing-bgp-ospf/questions-867418-limit-bandwith-of-a-wireless-ap-connected-to-mikrotik-routerboa-1773057e.md
source_anchor: ""
source_lines: [1, 8]
sha256: 2c1127499b27cf82ba06427f71d53bf77787fa318bc083b943bab13794526690
---

# questions-867418-limit-bandwith-of-a-wireless-ap-connected-to-mikrotik-routerboa-1773057e

We don't know the exact configuration (for instance I don't know whether A or B is handling DHCP), but I'm going to assume all 3 are effectively just on the same network and A manages everything.
As mentioned by pilsetnieks, if you could limit the guests to a specific IP range, it could be limited that way, although that may not be trivial or even possible depending on the exact hardware and setup.
My suggestion would be to do the following (Which is effectively the same as running a new cable from A - C, without having to run a cable).
- Add a small vlan switch next to B. Seeing as you're already using Mikrotik, something like a RB260GS would be fine.
- Plug A and C into the switch, and add a new cable between B and the switch
This just requires enough space next to B to place a small switch, a power outlet and a patch cable to connect it to B.
You can now configure the ports for B and C on the switch to be on independent vlans, and configure the port going to A to carry both. There'll also be some configuration on A to handle the vlans, but once configured this with effectively allow you to treat A-B and A-C as separate subnets with independent IP ranges. (I haven't done anything like this on Mikrotik for a while so I probably wouldn't be able to advise working config unfortunately. I think there's a few sample configs on the Mikrotik wiki)
This has the added benefit of splitting the guest traffic from everything else (which you really should do if it is actually guest traffic). You can use the firewall on A to allow or block traffic between the two networks if needed.
