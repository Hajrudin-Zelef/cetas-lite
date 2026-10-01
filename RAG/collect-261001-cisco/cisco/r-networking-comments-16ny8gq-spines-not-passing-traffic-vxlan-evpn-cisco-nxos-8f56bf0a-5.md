---
id: collect-261001-cisco/cisco/r-networking-comments-16ny8gq-spines-not-passing-traffic-vxlan-evpn-cisco-nxos-8f56bf0a-5
title: "r-networking-comments-16ny8gq-spines-not-passing-traffic-vxlan-evpn-cisco-nxos-8f56bf0a"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: ["asic", "ethernet"]
source: docs/RAG/collect-261001-cisco/r-networking-comments-16ny8gq-spines-not-passing-traffic-vxlan-evpn-cisco-nxos-8f56bf0a.md
source_anchor: ""
source_lines: [241, 255]
sha256: 3f8a58a89558ef4a82391da015640fcbdf54bdbc46e1ad1be2c415aabb6b8f18
---

# r-networking-comments-16ny8gq-spines-not-passing-traffic-vxlan-evpn-cisco-nxos-8f56bf0a

I believe you see evidence of this in your SPAN-to-CPU. The detailed Ethanalyzer output you shared indicates that there is no 802.1Q header between the outer Ethernet header and the outer IPv4 header. This is because the ASIC physically cannot insert the 802.1Q header after having already encapsulated your native IPv4 packet with a VXLAN header. The spine is most likely receiving this packet and dropping it on ingress.
(Note that I believe this is the case - I'm 97.3% sure that a packet punted to the CPU via SPAN-to-CPU will retain the 802.1Q header [such that it's visible in detailed Ethanalyzer output], but I would need to test that just to be absolutely positive.)
Another key piece of evidence that suggests the spine is dropping this packet on ingress (meaning, it's not forwarding it out to the leaf) is in your non-detailed Ethanalyzer output:
Recall you have the following SPAN-to-CPU configuration:
With this configuration, if you have a packet ingress Ethernet1/2 and successfully egress Ethernet1/4, I would expect to see the same packet twice in the Ethanalyzer output - once as it ingresses Ethernet1/2, and once as it egresses Ethernet1/4. That's not the case here, because we see each ICMP Echo sequence number once, and the timestamps are very evenly spaced out with one second in between. I'd expect to see sub-millisecond differences in timestamps if the packet was being duplicated on ingress and egress, so it does feel like the spine is dropping this packet on ingress.
A "fun fact" - this same double-encap hardware limitation is also the root cause for a number of limitations with other features (e.g. why you cannot ERSPAN towards an IP where the egress interface is a tunnel, why you cannot configure Q-in-Q such that a frame ingresses with no 802.1Q headers and egresses with two 802.1Q headers, etc.)
As a side note, this likely worked just fine with Nexus 9000v because in 9000v, all of the forwarding is done in software. There's no ASIC/hardware for forwarding to be delegated to on a 9000v, and therefore, you cannot encounter hardware-specific limitations like this one when using 9000v (although interestingly, the lack of hardware/TCAM is the root cause of why some features aren't supported on 9000v).
Lastly, just to clear up any confusion, let's address the documentation you originally quoted:
This was taken from the Cisco Nexus 9000 Series NX-OS VXLAN Configuration Guide, Release 9.3(x) document. However, the bullet point immediately following this statement is as follows:
So, the intent of these two bullet points is to state that before 9.3(5), you could not have non-VXLAN traffic egress a subinterface if the parent interface was a VXLAN uplink. However, this capability was introduced starting with 9.3(5), but you still cannot forward VXLAN traffic over a subinterface due to the aforementioned hardware limitation.
I hope this helps clear things up - let me know if you have any questions!
This was, indeed, the issue. Thank you so much for this detailed response! So great!
VXLAN/EVPN is the biggest scam network vendors managed to sell to their customers.
What.. Why?
lol wut
