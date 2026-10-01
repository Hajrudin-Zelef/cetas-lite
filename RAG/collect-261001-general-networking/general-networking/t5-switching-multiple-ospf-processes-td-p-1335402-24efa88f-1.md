---
id: collect-261001-general-networking/general-networking/t5-switching-multiple-ospf-processes-td-p-1335402-24efa88f-1
title: "t5-switching-multiple-ospf-processes-td-p-1335402-24efa88f"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["cost"]
source: docs/RAG/collect-261001-general-networking/t5-switching-multiple-ospf-processes-td-p-1335402-24efa88f.md
source_anchor: ""
source_lines: [1, 55]
sha256: 350b8a41b5913d4d1e4c0c46c4f834a9db75f714d511ba27bd23acf89732f532
---

# t5-switching-multiple-ospf-processes-td-p-1335402-24efa88f

- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
07-13-2009 07:50 PM - edited 03-06-2019 06:45 AM
Why would someone run multiple OSPF processes on the same router?
What are the implications?
Are there two SEPARATE databases that are established, one for each process?
Are these processes totally isolated?
What about the RIB? Are there effectively separate routing tables?
lets say a route is learned through both OSPF processes, what happens then....?
Another quick question:
router ospf 499
router-id 10.199.254.38
log-adjacency-changes
auto-cost reference-bandwidth 1000000
nsf
area 1 authentication message-digest
redistribute static metric-type 1 subnets route-map into-ospf
network 10.195.48.40 0.0.0.7 area 1
network 10.195.48.48 0.0.0.15 area 1
network 10.199.254.38 0.0.0.0 area 1
network 99.207.255.8 0.0.0.0 area 1
!
router ospf 599
log-adjacency-changes
auto-cost reference-bandwidth 1000000
nsf
area 0 authentication message-digest
area 2 authentication message-digest
redistribute static
passive-interface default
no passive-interface GigabitEthernet4/1
no passive-interface Vlan98
no passive-interface Vlan99
network 10.195.48.8 0.0.0.3 area 2
network 10.195.48.32 0.0.0.7 area 2
network 10.195.48.248 0.0.0.7 area 0
network 10.195.49.0 0.0.0.127 area 2
network 10.195.50.0 0.0.0.127 area 2
network 10.195.64.0 0.0.15.255 area 2
network 10.195.80.0 0.0.0.255 area 2
<b>interface GigabitEthernet4/1
ip address 10.195.48.53 255.255.255.240</b>
Notice the IP address of this interface. OSPF 499 is running OSPF on this interface (see the network statement), yet the interface is un-passived out in OSPF599....why is that? Typos?
Thanks
Solved! Go to Solution.
- Labels:
- 
						
							
		
