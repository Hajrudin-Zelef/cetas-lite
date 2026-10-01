---
id: collect-261001-opnsense-pfsense/opnsense-pfsense/r-opnsense-comments-14joc3g-opnsense-cant-get-vlans-to-work-afb9ae3c
title: "r-opnsense-comments-14joc3g-opnsense-cant-get-vlans-to-work-afb9ae3c"
domain: opnsense-pfsense
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-opnsense-pfsense/r-opnsense-comments-14joc3g-opnsense-cant-get-vlans-to-work-afb9ae3c.md
source_anchor: ""
source_lines: [1, 20]
sha256: 1027c4fd1303f752be9199493250b61ff9b5f8cd721c96c8d2ad562208f5c739
---

# r-opnsense-comments-14joc3g-opnsense-cant-get-vlans-to-work-afb9ae3c

Hi guys, gals,

New to opnsense, and relatively new to this kind of networking.

Have the basics up and running but struggling with vlans.

    Have VLAN set up (say `192.168.30.0/24`), with the LAN interface as parent, assigned, and DHCP enabled. Pinging to the gateway (`192.168.30.1`) of that subnet works (from LAN).
  

However, when I try to connect to the VLAN itself, via a smart switch, I never get an IP address. It just times out.

    When I manually assign an ip address in the proper ip range (say `192.168.30.10` with subnet mask `255.255.255.0` and gateway `192.168.30.1`), I can't even ping the gateway.
  

I assume this has nothing to do with firewall rules, it is in the same subnet after all, no firewalls in place there I assume.

    The switch (netgear gs108Ev3) on which I am directly connected has the VLAN created, assigned (as `untagged`) to the connected port, and the port PVID set to 30 as well.
  

What am I missing? Thanks for your input!
