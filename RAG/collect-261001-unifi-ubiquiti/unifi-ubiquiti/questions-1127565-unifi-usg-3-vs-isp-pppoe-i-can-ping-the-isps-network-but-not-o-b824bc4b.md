---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/questions-1127565-unifi-usg-3-vs-isp-pppoe-i-can-ping-the-isps-network-but-not-o-b824bc4b
title: "questions-1127565-unifi-usg-3-vs-isp-pppoe-i-can-ping-the-isps-network-but-not-o-b824bc4b"
domain: unifi-ubiquiti
role: reference
task: reference
actors: []
dates: []
keywords: ["ethernet"]
source: docs/RAG/collect-261001-unifi-ubiquiti/questions-1127565-unifi-usg-3-vs-isp-pppoe-i-can-ping-the-isps-network-but-not-o-b824bc4b.md
source_anchor: ""
source_lines: [1, 19]
sha256: 7d137e519637ad48db9690ea65ed5e8d8febad5496d421d3c273513f76496810
---

# questions-1127565-unifi-usg-3-vs-isp-pppoe-i-can-ping-the-isps-network-but-not-o-b824bc4b

Our network is based on Ubiquiti Unifi devices and has some printers and assorted clients, and a Synology NAS running DSM7.
A few days ago the network suddenly developed problems, although I have made zero changes to any of the network devices these past two weeks. It's beginning to become a critical issue, but I'm stumped.
It's not a large network:
- Fiber connection enters the premises
- Fiber/ethernet media converter (zero configuration options, and owned by ISP)
- Ubiqiti USG ("Unifi Security Gateway")
- Ubiqiti Unifi Switch
- Ubiqiti Unifi access points
- misc. clients, printers, ...
- Synology NAS (the Unifi controller is running in a Docker container, along with other server tasks)
Of note is that the USG has a PPPoE configuration (username, password, VLAN) that must match what is provided by the ISP. This has been working fine "since forever". However, a few days ago the internet connection just ... dropped. From talking with the ISP tech support, their connection up to the USG checks out, and from what I can tell via Unifi our network also checks out. And yet, no connection.
Or is there?
It turns out that I can actually ping the ISP's default gateway (94.144.1.4), and I can also ping other addresses on the ISP's network (variations on our own public IP; variations on their gateway IP) but I cannot ping anything outside of the ISP's network (eg. 1.1.1.1, 8.8.8.8) -- heck, not even their own preferred DNS servers (eg. 212.242.40.3)!
Because they say that there is a connection into our premises, the problem must be on our end. I have been unable to obtain any explanation as to why I am able to access only their network but nothing outside.
To me this suggests that the ISP must have made some change that has affected our connection.
What gives? How can I troubleshoot this?
It may be related to my other issue, that the NAS (also serving as a Unifi controller) has been giving me grief with logging into it. I do not suspect a malicious attack, for whatever that's worth.
Edit 2: Now that I've received their official equipment, I see the exact same behaviour. At least they're now willing to sic their network team on the issue, so there is hope it will be resolved within the week.
Edit 1: I can't check the routing table as the interface of the USG in stand-alone mode is very simplistic:
