---
id: collect-261001-cisco/cisco/questions-24609-cisco-small-business-sg300-10p-and-ubiquiti-wireless-aps-is-is-p-7537c29a
title: "questions-24609-cisco-small-business-sg300-10p-and-ubiquiti-wireless-aps-is-is-p-7537c29a"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: ["research"]
source: docs/RAG/collect-261001-cisco/questions-24609-cisco-small-business-sg300-10p-and-ubiquiti-wireless-aps-is-is-p-7537c29a.md
source_anchor: ""
source_lines: [1, 28]
sha256: e197307f22611c022ec4fda40095edf3f10e9ba42bab9689b2cd71845bd84fca
---

# questions-24609-cisco-small-business-sg300-10p-and-ubiquiti-wireless-aps-is-is-p-7537c29a

Network Engineering is part of Stack Overflow’s open communities: specialist spaces where curiosity is welcome, knowledge is shared freely, and the best answers rise to the top.
This question shows research effort; it is useful and clear
1
This question does not show any research effort; it is unclear or not useful
Save this question.
Show activity on this post.
Problem: We have our Ubiquiti wireless APs hooked up to an SG300-10P. The user's MAC can roam from AP to AP without asking DHCP for an address. We want to have IP Source Guard enabled as a best practice, to prevent someone wreaking havoc on our wireless network with a static IP address.
Unfortunately we had to disable IP Source Guard because it appears to lock the MAC to a port as well as an IP. Lets say a user connects to an AP on Port 1 and pulls an address via DHCP. They then begin to walk across the office and migrate to another AP on port 2. At that point, IP Source Guard drops all of their traffic until they pull a new DHCP address.
That is what source guard is supposed to do. It protects the network by making sure that each address is only coming from one port on the switch. This prevents someone from spoofing to receive traffic destined for a different port. IP source guard doesn't do this by MAC address, but by IP address, so it isn't locking your MAC to a single port, it is locking the IP address to a single port to prevent spoofing. It uses DHCP snooping to lock the IP address to the port. To use the same device from a different port, you would need to pull an address from DHCP on the new port.
You seem to want it to work only for static addresses, but that is not how IP source guard works or what it is designed to do. It require DHCP snooping to allow an IP address on a port, and that is how it blocks statically assigned IP addresses, but there is no feature that only blocks statically assigned addresses; it is sort of a side effect which you can get around by putting a static binding on a port to allow a statically addressed device to work on a port.
IP source guard provides source IP address filtering on a Layer 2 port
to prevent a malicious host from impersonating a legitimate host by
assuming the legitimate host's IP address. The feature uses dynamic
DHCP snooping and static IP source binding to match IP addresses to
hosts on untrusted Layer 2 access ports.
Initially, all IP traffic on the protected port is blocked except for
DHCP packets. After a client receives an IP address from the DHCP
server, or after static IP source binding is configured by the
administrator, all traffic with that IP source address is permitted
from that client. Traffic from other hosts is denied. This filtering
limits a host's ability to attack the network by claiming a neighbor
host's IP address. IP source guard is a port-based feature that
automatically creates an implicit port access control list (PACL).
Ran into a similar issue with Cisco WS-3850 switches connected to both Aruba and Aerohive systems. Wireless stations roamed between WAPs ( and because the director of IT didn't believe in site surveys, even stationary clients fast roamed ) and the switches threw errors and warnings about "MAC-ADDRESS-FLAP xxxx.xxxx.xxxx".
Wasn't really able to resolve this, but potentially one solution - at least for Enterprise WAPs that have a controller appliance or controller server - would be to switch the WAPs from "bridge mode" in which traffic is directly bridged to your wired network and into "tunnel mode" in which traffic is routed either IP in IP or in GRE tunnels back to your controller appliance/server.
If you have a campus type network setup or you're in a single building this is not so bad, but the amount of overhead you incur depends greatly on the location of your controller relative to your WAPs. If your clients access internal resources that are not in a centralized location then the overhead is higher.
I think Cisco APs will work with a controller feature built into the WS-3850 switches that will alleviate this problem.
I "resolved" it by disabling ip device tracking on the switches, but this was not an ideal solution.
