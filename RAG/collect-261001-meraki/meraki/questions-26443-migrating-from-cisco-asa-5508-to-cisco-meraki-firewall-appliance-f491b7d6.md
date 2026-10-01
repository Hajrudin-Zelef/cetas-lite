---
id: collect-261001-meraki/meraki/questions-26443-migrating-from-cisco-asa-5508-to-cisco-meraki-firewall-appliance-f491b7d6
title: "questions-26443-migrating-from-cisco-asa-5508-to-cisco-meraki-firewall-appliance-f491b7d6"
domain: meraki
role: reference
task: reference
actors: []
dates: []
keywords: ["research"]
source: docs/RAG/collect-261001-meraki/questions-26443-migrating-from-cisco-asa-5508-to-cisco-meraki-firewall-appliance-f491b7d6.md
source_anchor: ""
source_lines: [1, 13]
sha256: 91c92a2d757c558bac92656ca2bc4a3bbd0df0b3dc22cf2d805f9738e81b8c6a
---

# questions-26443-migrating-from-cisco-asa-5508-to-cisco-meraki-firewall-appliance-f491b7d6

Network Engineering is part of Stack Overflow’s open communities: specialist spaces where curiosity is welcome, knowledge is shared freely, and the best answers rise to the top.
This question shows research effort; it is useful and clear
3
This question does not show any research effort; it is unclear or not useful
Save this question.
Show activity on this post.
At a small remote office my boss decided to buy a Meraki Firewall device with Wireless to replace an old Cisco ASA. The setup looks pretty straight forward. We have 5 public IP addresses for an e-mail server, 2 web servers, our ASA outside interface and the last one is used for PAT on the ASA.
In the Meraki dashboard it was easy to set up 1:1 NAT for our servers and I have the outside interface public IP configured but where do I conigure the PAT internet address that all of our internal clients use to get on the internet?
I have ran into several situations similar to this and engaged Meraki support to confirm.
1:Many NAT only allows you to specify a WAN IP address to have ports forwarded to multiple internal hosts.
1:1 NAT only links 1 WAN IP to 1 LAN IP
Port Forwarding is only tied to the external IP of the Meraki device itself
That said those are your only options within Meraki, you cannot PAT an entire subnet to a different WAN IP Address. If this is required you will need to either reconfigure your internal network to make the Meraki work or get another ASA.
