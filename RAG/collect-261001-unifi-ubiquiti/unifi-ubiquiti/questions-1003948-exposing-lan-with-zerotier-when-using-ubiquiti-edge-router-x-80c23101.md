---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/questions-1003948-exposing-lan-with-zerotier-when-using-ubiquiti-edge-router-x-80c23101
title: "questions-1003948-exposing-lan-with-zerotier-when-using-ubiquiti-edge-router-x-80c23101"
domain: unifi-ubiquiti
role: reference
task: reference
actors: ["Google"]
dates: []
keywords: []
source: docs/RAG/collect-261001-unifi-ubiquiti/questions-1003948-exposing-lan-with-zerotier-when-using-ubiquiti-edge-router-x-80c23101.md
source_anchor: ""
source_lines: [1, 39]
sha256: 6d8c3fd033b4ddf93d5cc0a1526d92280bd4a2e8220ac26ff465dfb71acd8d22
---

# questions-1003948-exposing-lan-with-zerotier-when-using-ubiquiti-edge-router-x-80c23101

Here's the solution:
- Update the Ubiquiti EdgeRouter-X (ER-X) with Firmware 2.0.8 (you can Google this very easily)
- Install ZeroTier from the command line (Skipping the device mapping)
ubnt@edgerouter# sudo -i
root@edgerouter:~# curl -s https://install.zerotier.com | sudo bash
At this point the ER-X is unable to "interact" with ZeroTier so we need to patch.
- Patch with the zerotier-edgeos (by Dennis Kruyt) 
Download this file from https://github.com/dkruyt/resources/raw/master/zerotier-edgeos.tgz
root@edgerouter:/# tar -C / -xvzf /home/admin/zerotier-edgeos.tgz
At this point after reboot the ZeroTier interface will now show in the router dashboard. 
- Setup DNAT and SNAT rule in the ER-X router, taking into account the Upstream IP address and the ZeroTier IP address and the switch0 address, additionally, to keep the setup simple a Network Group was created, the Upstream and ZeroTier group respectively (nothing special here just set the Group following the standard Ubiquiti guide
Then
ubnt@edgerouter# configure
edit service nat rule 1
set description "ZeroTier DNAT" 
set destination group network-group ZeroTier
set inbound-interface ztzlgezugz
set inside-address address 192.168.0.157
set log disable
set protocol all
set type destination
top
edit service nat rule 5000
set description "ZeroTier SNAT"
set log disable
set outbound-interface ztzlgezugz
set outside-address address 10.147.17.15
set protocol all
set source group network-group Upstream
set type source
top
commit
save exit
Where the IP 192.168.0.157 is the upstream IP i.e. eth0 and  the 10.147.17.15 is the ZT IP
- Setup Managed Route in ZeroTier:
Note: 192.168.1.0/24 here is the switch0 address where the LAN devices are connected, eth1, eth2, eth3 and eth4 in my case since I have ER-X router
Done.
From here, you will be able to access the LAN devices connected to switch0 over the internet. 
References:
