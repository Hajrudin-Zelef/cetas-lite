---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/questions-24780-vyatta-edgeos-remote-access-vpn-without-nat-f905ced4
title: "questions-24780-vyatta-edgeos-remote-access-vpn-without-nat-f905ced4"
domain: unifi-ubiquiti
role: reference
task: reference
actors: []
dates: []
keywords: ["research"]
source: docs/RAG/collect-261001-unifi-ubiquiti/questions-24780-vyatta-edgeos-remote-access-vpn-without-nat-f905ced4.md
source_anchor: ""
source_lines: [1, 9]
sha256: 405d4946998870d8ff320eea0ff13481dbd77a82af153228182349466c60651e
---

# questions-24780-vyatta-edgeos-remote-access-vpn-without-nat-f905ced4

Network Engineering is part of Stack Overflow’s open communities: specialist spaces where curiosity is welcome, knowledge is shared freely, and the best answers rise to the top.
This question shows research effort; it is useful and clear
2
This question does not show any research effort; it is unclear or not useful
Save this question.
Show activity on this post.
I currently have a ubiquiti edgerouter setup with an ipsec site-to-site vpn and l2tp/ipsec remote access vpn. Both work as intended.
Is it possible to configure the remote access (l2tp/ipsec) vpn to work when the client is not behind NAT? I'm having trouble finding information about this since NAT is so prevalent and everyone assumes that it will be in use.
You might be able to do this using OpenVPN instead, although you'd end up needing to install client software. Assuming that the need is for something more secure than PPTP on the client end...
