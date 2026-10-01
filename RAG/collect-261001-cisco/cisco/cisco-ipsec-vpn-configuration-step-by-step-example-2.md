---
id: collect-261001-cisco/cisco/cisco-ipsec-vpn-configuration-step-by-step-example-2
title: "cisco-ipsec-vpn-configuration-step-by-step-example"
domain: cisco
role: reference
task: reference
actors: ["Huawei"]
dates: []
keywords: []
source: docs/RAG/collect-261001-cisco/cisco-ipsec-vpn-configuration-step-by-step-example.md
source_anchor: ""
source_lines: [220, 251]
sha256: a4096fb66bd931a65c51b03b24e0899846bb27767d553b7a968d1d74085f5f27
---

# cisco-ipsec-vpn-configuration-step-by-step-example

**“show crypto session remote *ip-address* detail”** command shows the details of the crypto session.

```
RouterA# 
```
**show crypto session remote 20.20.20.20 detail**
Crypto session current status
Interface: GigabitEthernet1/0
Uptime: 00:11:11
Session status: UP-ACTIVE   >>>>> Status of the VPN
Peer: 20.20.20.20 port 500 fvrf: (none) ivrf: (none)
Phase1_id: 20.20.20.20
Desc: (none)
Session ID: 0
IKEv1 SA: local 10.10.10.10/500 remote 20.20.20.20/500 Active
Capabilities:(none) connid:1012 lifetime:14:10:23
IPSEC FLOW: permit ip 192.168.1.0/255.255.255.0 172.16.0.0/255.255.255.0
Active SAs: 2, origin: crypto map
Inbound:  #pkts dec'ed 25 drop 0 life (KB/Sec) 3458240/1743
Outbound: #pkts enc'ed 25 drop 0 life (KB/Sec) 3458240/1743
In this **IPSec for VPN Configuration** example, we have learned the details of how to configure **IPSEC VPN on Cisco routers**. You can try **Cisco IPSec Configuration** with different encryption, hashing and authentication methods.


Gokhan Kosem is a Network Engineer, Instructor and the Founder of IPCisco.com with 15+ years of experience in Cisco, Nokia, Huawei, Juniper, Linux, Service Provider Networks, Routing and Switching technologies.

He has worked on the backbone networks of major service providers and network vendors including Nortel, Alcatel-Lucent (Nokia) and has extensive hands-on experience with Cisco, Huawei, Juniper and Nokia networking technologies.

He has trained thousands of networking students worldwide through IPCisco.com, Udemy, books, labs, quizzes, and educational content across multiple social media platforms.

IPCisco.com | Best Route to Your Dreams

## Leave a Reply
