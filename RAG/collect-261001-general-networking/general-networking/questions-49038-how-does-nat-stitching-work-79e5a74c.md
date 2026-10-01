---
id: collect-261001-general-networking/general-networking/questions-49038-how-does-nat-stitching-work-79e5a74c
title: "questions-49038-how-does-nat-stitching-work-79e5a74c"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["cybersecurity"]
source: docs/RAG/collect-261001-general-networking/questions-49038-how-does-nat-stitching-work-79e5a74c.md
source_anchor: ""
source_lines: [1, 14]
sha256: f2d7816aeef1dd1839be1b188f0773af659dc0dd8d19d67d771fb17fea19d85e
---

# questions-49038-how-does-nat-stitching-work-79e5a74c

You must remember that a flow using NAT will look like two different flows: a flow pre-NAT, and a flow post-NAT. This is because NAT is changing one or more of the addresses in the packets. This can present a distorted view of your flows.
As Cisco explains it, NAT stitching will stitch the (apparently) separate flows to give you the single flow view:
  Exporting NetFlow from the NAT devices will stitch both pre & post NAT
  flows together.
The Cisco Press book NetFlow for Cybersecurity goes into more detail:
  Lancope’s StealthWatch solution supports a feature called network
  address translation (NAT) stitching. NAT stitching uses data from
  network devices to combine NAT information from inside a firewall (or
  a NAT device) with information from outside the firewall (or a NAT
  device) to identify which IP addresses and users are part of a
  specific flow. A great feature of the StealthWatch solution is its
  ability to perform “NetFlow deduplication.” This feature allows you to
  deploy several NetFlow collectors within your organization without
  worrying about double or triple counting the traffic.
