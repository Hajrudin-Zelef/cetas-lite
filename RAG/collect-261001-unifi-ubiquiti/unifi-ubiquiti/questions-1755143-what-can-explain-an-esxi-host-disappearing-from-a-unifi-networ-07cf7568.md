---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/questions-1755143-what-can-explain-an-esxi-host-disappearing-from-a-unifi-networ-07cf7568
title: "questions-1755143-what-can-explain-an-esxi-host-disappearing-from-a-unifi-networ-07cf7568"
domain: unifi-ubiquiti
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-unifi-ubiquiti/questions-1755143-what-can-explain-an-esxi-host-disappearing-from-a-unifi-networ-07cf7568.md
source_anchor: ""
source_lines: [1, 5]
sha256: a5ac5c7af084e89686f620037e026c453c344cdc37a1b17562ee26e11314cf0e
---

# questions-1755143-what-can-explain-an-esxi-host-disappearing-from-a-unifi-networ-07cf7568

My ESXi host is running headless and contains 2 VMs. Both VMs are set to autostart after a reboot and come up without any problem. There is only a single network interface both for the management and the VMs hence I'd expect both the VMs and the management interface to be reachable.
The ESXi host interface was detected on the network (in Unifi client log) for ~10 minutes after a reboot, but is not happening reliably after every reboot. The host is connected to a Unifi USW-16-PoE on firmware 6.3.13
I paired manually the host mac address with static IP using arp in a client, but the host remained unreachable.
Another ESXi host in my network (with a different hardware) works almost fine (i.e. it disappears from the list of active hosts in the Unifi list, but is shown in an arp -a and remains reachable). The working host connected through a Unifi Flex switch (on firmware 1.8.6), not directly to the main switch.
I looked at a similar thread (Out of 2 VMs and their 1 host, only 2 at a time are reachable from the LAN) but there isn't a configuration that limits the clients per port on Unifi's controller that I can find.
