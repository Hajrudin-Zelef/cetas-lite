---
id: collect-261001-general-networking/general-networking/t5-network-security-inter-vlan-routing-configuration-in-firepower-1010-m-p-43149-4c299e4d
title: "t5-network-security-inter-vlan-routing-configuration-in-firepower-1010-m-p-43149-4c299e4d"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-general-networking/t5-network-security-inter-vlan-routing-configuration-in-firepower-1010-m-p-43149-4c299e4d.md
source_anchor: ""
source_lines: [1, 120]
sha256: c6590e194efe92c578cb4e8590782bd3e7767182813bb3ac43a290cc8e6abda1
---

# t5-network-security-inter-vlan-routing-configuration-in-firepower-1010-m-p-43149-4c299e4d

- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
03-26-2021 02:58 AM - edited 03-26-2021 01:57 PM
hello i have 3 V LAN, hosts can ping their default gateway  can connect to FM C but can not ping each other, on A C P i Allow everything, ft d can ping all hosts. what can be  a problem?
Solved! Go to Solution.
Accepted Solutions
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
03-27-2021 10:25 AM
The result of that packet-tracer output was an "allow", which means the traffic should be permitted by the FTD.
I suggest you disable the local firewall on the hosts for testing and try again.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
03-26-2021 02:36 PM - edited 03-26-2021 02:38 PM
Do you have sub-interfaces on the FTD?
Is the FTD the default gateway for each VLAN?
Do you have NAT exemption rules in place to ensure the inter-vlan traffic is not unintentially natted?
Please can you run packet-tracer from the CLI and provide the output for review. Example: packet-tracer input <interface> <protocol> <src ip> <src port> <dst ip> <dst port>
Provide some output of your FTD and switch configuration.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
03-27-2021 09:58 AM
sorry i do not know how to run paket-tracer from CLI, i have sub-interface, host can ping ftd and all vlan can connect to fmc. default gateway is correct on host it is ftd sub-interface ip.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
03-27-2021 10:11 AM
@shotalezhava Run the following from the CLI of the FTD and provide the output:-
packet-tracer input managment icmp 192.168.77.11 8 0 192.168.10.10
Can the FTD ping a host in each of the vlans?
Does the host you are trying to ping (192.168.10.10) have a local firewall turned on that could be preventing a ping response?
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
03-27-2021 10:13 AM
local firewall allow icmp Ipv4
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
03-27-2021 10:17 AM
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
03-27-2021 10:25 AM
The result of that packet-tracer output was an "allow", which means the traffic should be permitted by the FTD.
I suggest you disable the local firewall on the hosts for testing and try again.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
03-27-2021 10:33 AM
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
03-27-2021 10:44 AM
Right fine, but you are troubleshooting a problem with the FTD....you can easily eliminate that a potential issue by temporarily disabling the local firewall on the windows host and then test.
Aside from that, is the FTD the default gateway definately of the windows host?
Run the command "system support firewall-engine-debug" from the FTD CLI and filter on the ip address(es) you are testing with and then run some tests to generate some traffic.
You could also run packet capture on the windows host.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
03-29-2021 04:04 AM
it was windows firewall problem! thanks
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
05-12-2021 02:29 PM
hello i create HA and i have error Frequent drain of Archives & Cores & File Logs
