---
id: collect-261001-general-networking/general-networking/questions-443580-udp-packet-size-max-ping-failing-to-3g-recalcitrant-software-ve-492b602d
title: "questions-443580-udp-packet-size-max-ping-failing-to-3g-recalcitrant-software-ve-492b602d"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-general-networking/questions-443580-udp-packet-size-max-ping-failing-to-3g-recalcitrant-software-ve-492b602d.md
source_anchor: ""
source_lines: [1, 15]
sha256: b1a7cd1b63cd5e1e235eefb842ec2286cbe6ed90127860952ee380461cf284dc
---

# questions-443580-udp-packet-size-max-ping-failing-to-3g-recalcitrant-software-ve-492b602d

We have a network of 330 clients (Service Vans) connected via 3G to our central office. Currently these clients are connected via a small ruggedized 3G router that establishes the 3G connection and then an IPSec VPN tunnel (OpenSwan stack)
A Windows 7 PC (A mobile data terminal) is plugged into this router and it runs a small piece of commercial software that uses UDP to communicate details of jobs that the Service Van has to attend.
During User Acceptance Testing of this product we did note some communication failures on WIRED clients when the NIC was set to 10MB/s or when an additional switch was placed in the communication pathway.. We never got to the bottom of this issue.. fortunately the software did actually work over our 3G connections. The software, despite some minor issues, has been working quite well since implementation.
Recently we kicked off a project to retire the 3G routers (which have a reliability issue) with an Internal 3G modem that is mounted in the Mobile Data Terminal that the Service Technician uses. Our test units are workig well, we can establish a 3G connection and then a VPN tunnel using the internal modem and a VPN client..
However the commercial software are using is now failing.. all other network traffic is unaffected, for example we can successfully use Remote Desktop over the link, copy files, ping back and forth etc..
The Vendor is not keen to investigate this issue in any depth, but I feel its almost certainly related to how they are using UDP.. we have seen some interesting PING results when using large packet sizes..
When pinging our 3G routers we can successfully ping with a buffer size of about 51Kb before failures start to occur..
When pinging our new 3G modems inside our Mobile Data Terminals we start to see failures with a buffer size of about 5Kb, at around 12Kb we have at least a 50% failure rate. At 51Kb we have total failure.
It seems clear that the 3G Router is coping much better with the large packet size.. I just dont know why.
I further feel that if we can resolve this issue the Vendors software will start to work again. I would much rather the Vendor have a more tolerant protocol but I fear that the software "is what it is".. and that it wont change.
The backend server for this application resides on a 32Bit Windows Server 2003 machine. The VPN Gateway is a Fortigate 310b.
I am seriously out of my depth and would appreciate any thoughts that anyone might have on how to approach the problem..
edit: more information available..
I wrote a small utility to send/recieve UDP datagrams.. the existing 3G router will happily accept a datagram up to about 50KB in size.. the internal modem maxes out at about a 11KB datagram...
When pinging our new 3G modems inside our Mobile Data Terminals we start to see failures with a buffer size of about 5Kb, at around 12Kb we have at least a 50% failure rate.So you are pinging from your office or similar yes? What do you mean buffer size of 5kb and 12kb, are you using standard ICMP ping increasing the payload size incrementally? Have you discoverd what the MTU of the new 3G modems are and old/existing 3G routers?
