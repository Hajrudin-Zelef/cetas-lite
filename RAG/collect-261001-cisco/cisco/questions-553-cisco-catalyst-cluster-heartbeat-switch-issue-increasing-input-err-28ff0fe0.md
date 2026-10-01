---
id: collect-261001-cisco/cisco/questions-553-cisco-catalyst-cluster-heartbeat-switch-issue-increasing-input-err-28ff0fe0
title: "questions-553-cisco-catalyst-cluster-heartbeat-switch-issue-increasing-input-err-28ff0fe0"
domain: cisco
role: reference
task: reference
actors: ["Broadcom", "Intel"]
dates: []
keywords: ["full-duplex", "intel"]
source: docs/RAG/collect-261001-cisco/questions-553-cisco-catalyst-cluster-heartbeat-switch-issue-increasing-input-err-28ff0fe0.md
source_anchor: ""
source_lines: [1, 24]
sha256: a2b5d922f8baa0b620b095338de5b2db3bcdeb612a0037ca5bc5df34877c2548
---

# questions-553-cisco-catalyst-cluster-heartbeat-switch-issue-increasing-input-err-28ff0fe0

PROBLEM: The servers in two clusters keep losing heartbeat connectivity with each other thus causing database outages. Outages are brief but disruptive.
SETUP:
- There are two clusters of three servers each.
- Each server has one NIC connected to a single Layer 2 switch (Catalyst 2950) with the switch ports hard-coded at 100Mb/full-duplex.
- The DBAs confirm that each heartbeat NIC is hard-coded to 100Mb/full-duplex.
- There are two clusters configured in VLAN 100 and in the same subnet (10.40.60.0/24).
- The management IP address is on a separate subnet (10.40.1.0/24) and it's switch port is in VLAN 1.
SYMPTOMS:
- I see an ever-increasing error count on the switch ports. For the three servers in one cluster the input errors (all CRC) are about 3% of total input packets. There are no output errors. The other cluster is at about 6% of total input packets.
- Transmit and receive load on the switch ports is light, under 20/255 on txload and rxload.
- The switch log shows the switch ports bouncing:
May 16 11:15:31 PDT: %LINEPROTO-5-UPDOWN: Line protocol on Interface FastEthernet0/13, changed state to down 
 May 16 11:15:32 PDT: %LINK-3-UPDOWN: Interface FastEthernet0/13, changed state to down
 May 16 11:15:34 PDT: %LINK-3-UPDOWN: Interface FastEthernet0/13, changed state to up
May 16 11:15:35 PDT: %LINEPROTO-5-UPDOWN: Line protocol on Interface FastEthernet0/13, changed state to up
- I replaced the old Cat5 cabling between the server heartbeat NIC and the switch with new Cat6 -- no effect.
- I created a new VLAN 200 in a new subnet (10.40.61.0/24) and had the DBAs re-IP their heartbeat NICs on one cluster -- no effect.
- We tried every combination of speed and duplex on the switch port and the NIC -- no effect, went back to 100Mb/full-duplex on both.
- The DBAs upgraded the Broadcom drivers on both clusters to the latest -- drop in error percentage on the 6% cluster down to 4%, the other cluster is still at 3%.
MY PROPOSED NEXT STEPS:
- There are Intel NICs on the servers. Try moving the cluster heartbeat to an Intel NIC. Maybe it's a Broadcom issue?
- Change out the switch to a gig capable switch. There is a Catalayst 3560x available, but taking it will delay a project. Maybe gig on the switch port and NIC will play nicer?
THOUGHTS?
Is there something I can configure on the existing 2950 switch to mitigate the errors? What additional troubleshooting steps should I take?
