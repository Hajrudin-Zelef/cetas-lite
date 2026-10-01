---
id: collect-261001-general-networking/general-networking/questions-1128003-tcp-communication-to-internet-broken-no-syn-ack-received-c3a0e908
title: "TCP communication to internet broken - no SYN-ACK received"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-general-networking/questions-1128003-tcp-communication-to-internet-broken-no-syn-ack-received-c3a0e908.md
source_anchor: ""
source_lines: [1, 33]
sha256: b325713dd30057ca2800dca020004f8fd1822f68280e4755000245e77e2a7f7e
---

# TCP communication to internet broken - no SYN-ACK received

*Score : 0 | Source : https://serverfault.com/questions/1128003/tcp-communication-to-internet-broken-no-syn-ack-received*

Given following network setup:
Debian bullseye host (Proxmox VE) -> OPNsense firewall doing NAT -> DSL modem (connected via PPPoE) -> Internet
The host isn't able to establish a tcp connection to any host in the internet. I tried to reach different hosts, also a server where I am able to capture traffic.
I captured some traffic using port mirroring of the DSL modem switch port ensuring I really see traffic passing the line to the internet.
Traffic from another debian machine in the network using the same topology looks fine.
On both machines I executed curl -4 http://ftp.de.debian.org/debian and curl -4 http://www.google.de and compared the TCP SYN packets.
They look the same to me, differing only in
These differences are expected.
Following a screenshot capturing traffic at the DSL modem port. The capture contains traffic from the working host up until package 308. Traffic from packet 309 onwards is from the broken host.
I filtered on packets having the SYN flag set, as these packets are relevant only.
As mentioned above, I also took a traffic capture from the server side. It looks exactly the same as visible in the screens (packet 309, 310, 320, ...).
Additional info:
What could cause the internet hosts not to answer the SYN packets with SYN-ACK?

---

### Reponse (acceptee) — score 1

Short answer: the checksums get calculated incorrectly. Enabling checksum verification within wireshark makes this obvious:
Long answer:
The problem is caused by the fact, that OPNsense runs virtualized on top of Proxmox VE. Therefore I disabled some offloading features of the NIC within /etc/network/interfaces:
[...]
pre-up ethtool -K enp0s20f1 gro off lro off gso off
post-up ethtool -K vmbr1 gro off lro off gso off
This snippet disables generic receive offload and large receive offload. This is to prevent the NIC from reassembling multiple TCP packets into larger ones that might be larger than MTU. As modifying packets is a bad idea when simple routing should be performed, this causes connectivity problems in many cases. Disabling generic segmentation offset prevents the NIC from splitting large packets into multiple.
The snippet above has to contain tx off
pre-up ethtool -K enp0s20f1 gro off lro off gso off tx off
post-up ethtool -K vmbr1 gro off lro off gso off tx off
to enforce software checksum calculation.
