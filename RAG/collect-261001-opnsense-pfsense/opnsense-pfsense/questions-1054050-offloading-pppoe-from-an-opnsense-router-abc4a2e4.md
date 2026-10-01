---
id: collect-261001-opnsense-pfsense/opnsense-pfsense/questions-1054050-offloading-pppoe-from-an-opnsense-router-abc4a2e4
title: "Offloading PPPoE from an OPNsense router"
domain: opnsense-pfsense
role: reference
task: reference
actors: ["Broadcom", "Intel"]
dates: []
keywords: ["ethernet", "intel", "open source", "pricing"]
source: docs/RAG/collect-261001-opnsense-pfsense/questions-1054050-offloading-pppoe-from-an-opnsense-router-abc4a2e4.md
source_anchor: ""
source_lines: [1, 52]
sha256: b2b20db74c0cb5ff1979d862c804ebe22e0828c88191670be11606e691905462
---

# Offloading PPPoE from an OPNsense router

*Score : 3 | Source : https://serverfault.com/questions/1054050/offloading-pppoe-from-an-opnsense-router*

I'm running opnSense, a FreeBSD-based firewall and router similar to pfSense, in a virtual machine under VMware ESXi 7 on a Dell PowerEdge R230, as a router for my home network. No other VMs are running or even set up on the host, just this one.
My ISP uses symmetrical gigabit fiber. You connect to it via PPPoE with VLAN tagging. The IP is dynamically assigned, although it's unclear if it's DHCP or not. Either way, I get assigned a single IP not of my choosing. I don't have any static IPs nor a static block, something for which my ISP charges a non-trivial amount. Reconnecting often does give me a new IP, so I can't assume it would stay the same all the time.
I set up a gateway for the ISP in opnSense. opnSense does the PPPoE connection and for simplicity ESXi itself does the VLAN tagging so it's transparent to the underlying switch. The topology ends up looking like this:
       Internet
           ↓
      Fiber line
           ↓
       Fiber ONT
           ↓
    Ethernet cable
           ↓
     Physical NIC
           ↓
Virtual switch/port group
           ↓
      Virtual NIC
           ↓
      opnSense VM
The NICs in ESXi look like this. The first two are the onboard LAN. The last four are the PCI-E card: Dell Intel I350-T4 Quad Port 1GbE PCI-E card. The ISP's connection is using vmnic2, the first port on the PCI-E card.
But, the problem is I'm not getting the full speed. This is because FreeBSD has a bug with some NICs where PPPoE does not use all the CPU cores. The VM has what I'd consider overkill assigned to it: 4 vCPUs from a Xeon E3-1240 v5 (3.5 GHz) and 8 GB RAM, but apparently the single core performance isn't fast enough to do gigabit due to PPPoE. When I do a speed test, the CPU usage within the VM predictably goes to 25% (one of four cores) and I only get about 600 Mbps peak download, and for some reason upload is much worse than that too. When I use other hardware, like a Ubiquiti UniFi USG-3P or the ISP's provided modem/router, I get the nearly the full speed both directions. I've tried making the virtual NIC both vmxnet3 and e1000e with no difference in performance.
So with this setup, how can I get full speed with a PPPoE connection? I see several options but I don't know which is the best or the most likely chance of success:
- Use some other network driver in opnSense, although I have no idea how to do that or if an applicable one even exists. Is this even an option?
- Offload the PPPoE unwrapping in a Linux VM which doesn't have the PPPoE bug, and then have opnSense use that VM as the gateway and not the ISP directly. With this option there are two approaches I can think of:
  - Use a Linux VM to act as a very simple router with NAT and treat opnSense as a DMZ, so all the port forwarding, etc. applies only to opnSense and I don't have to do it two places. This does directly expose the Linux VM to the internet without the opnSense firewall to protect it. Is this safe?
  - Use a Linux VM to only unwrap the PPPoE aspect of the connection and send the unwrapped Ethernet frames straight to another interface; i.e., opnSense still gets the assigned the external IP and there's no NAT. I saw a tutorial on that but I don't think it works in my case because I don't have a static IP or static IP block. This sounds like the cheapest option with no new hardware or service charges if it works, but is it a good idea and is it safe?
- Use a hardware device of some kind to do the PPPoE-to-Ethernet conversion. My ISP-provided modem/router, which I'm not using right now, is a Zyxel C1100Z. It has an option for "transparent bridging", but best I can tell, the PPPoE handling would still happen on opnSense in this scenario which doesn't help. I'm not using the UniFi USG-3P because it is very buggy with software updates and also doesn't do the WAN failover reliably at all.
- Replace the PCI-E NIC with something else, or passthrough the physical PCI-E card to the VM instead of virtualizing it. From what I've seen so far, other vendors' cards like Realtek are even more unreliable in opnSense. Is there a NIC that's known to work for this scenario? Would the onboard LAN work any better?
- Install opnSense directly on the host or run it directly off a USB stick to eliminate the VMware side. I'm really hoping I can do a VM-based setup for snapshots and whatnot for safety. Would this be any improvement at all given the underlying hardware won't change?
- Replace opnSense with some other virtual appliance not based on FreeBSD, although I haven't yet found something as comparatively easy to use or feature-rich, or they're expensive or require subscription pricing. I prefer UIs over text configuration and want to use open source self-hosted software. Does such a product exist?

---

### Reponse — score 2

Have you tried this.
There is another way. By default, FreeBSD kernel performs all processing of received PPPoE frame within driver interrupt context: decapsulation, optional decompression/decryption, network address translation, routing lookups, packet filtering and so on. This can result in overloaded single CPU core in default configuration when sysctl net.isr.dispatch=direct. Since FreeBSD 8 we have netisr(8) network dispatch service allowing any NIC driver just enqueue received ethernet frame and cease its following processing freeing this CPU core. Other kernel threads using other CPU cores will then dequeue received frames to complete decapsilation etc. loading all CPU cores evenly.
So, one just should make sure it has "net.isr.maxthreads" and "net.isr.numthreads" greater than 1 and switch net.isr.dispatch to "deferred" value that permits NIC drivers to use netisr(9) queues to distribute load between CPU core
From
https://bugs.freebsd.org/bugzilla/show_bug.cgi?id=203856#c11

---

### Reponse — score 1

Here's a few suggestions:
Use a bridge interface from ESXi for your WAN interface as opposed to PCI passthrough
Use a Broadcom NIC, or an older Intel NIC using the em(4) driver. It may sound weird, but Broadcom NICs may be better for PPPoE.
em(4)
