---
id: collect-261001-meraki/meraki/r-meraki-comments-9iguzw-mx-firewalling-5dd83a2b
title: "r-meraki-comments-9iguzw-mx-firewalling-5dd83a2b"
domain: meraki
role: reference
task: reference
actors: ["Google"]
dates: []
keywords: []
source: docs/RAG/collect-261001-meraki/r-meraki-comments-9iguzw-mx-firewalling-5dd83a2b.md
source_anchor: ""
source_lines: [1, 28]
sha256: bc5764a00caab1e5b49502713612cf49ea4ceca6ebffec58adff762146ef8e33
---

# r-meraki-comments-9iguzw-mx-firewalling-5dd83a2b

MX Firewalling
Can someone help me understand how the Meraki firewall works please. I can see that there is no option for inbound firewall rules, but I came across this post.
https://community.meraki.com/t5/Security-SD-WAN/Inbound-firewall-rules/td-p/22679
Meraki has a unique way of doing firewall rules compared to a traditional firewall. Here is an example. If you were trying to prevent a network server at 8.8.8.8 from being able to ping anything in your environment. On a traditional firewall you could prevent incoming icmp from 8.8.8.8. On the MX you'd instead create an outgoing rule to prevent ICMP to 8.8.8.8. It accomplishes the same thing of ultimately blocking the incoming traffic but it does it via blocking the response. It took me a while to wrap my head around this difference since I was used to traditional Cisco ACLs and Sonicwalls
This makes it sound like the firewall doesn’t block the inbound traffic, but just blocks the return traffic? Is this correct?
What happens if you want your server to ping 8.8.8.8 but don’t want 8.8.8.8 to ping your server?
Is this really the case?
Section des commentaires
Yes and no.
8.8.8.8 cant just ping your server unless you have a NAT rule, and then allow that traffic in that NAT rule.
The difference between an ASA and a MX is that the ACL is attached to the NAT rules on your 1:1 NATs or on your 1:Many PAT. There aren't interface ACL.
What are you trying to accomplish, since I'm guessing stopping Google DNS from pinging you isn't it? 😊
Hadn't thought about the NAT part of it regarding the internet. We are going to have our MPLS on an SVI hosted on the MX. If I wanted to allow my traffic behind the MX to contact a server on the MPLS, but not the other way round, how would I achieve that?
Got it. So that comes back to the No in my "yes and no". ☹️.
So there isn't really a way to do one way permissions in the MX that I am aware of. I can test something out in my lab and let you know though. Just to be clear, both networks are using an SVI on the MX as their gateway?
All firewalling is flow-based. All inbound flows are denied by default (with the exception of ICMP ping, which is part of a special set of rules under "security appliance services"). All outbound flows are allowed by default.
The only way to allow inbound flows is to enable a port-forward or 1:1 NAT rule, in which case you can configure a whitelist of hosts that are allowed initiate flows to that port.
You cannot configure the firewall such that outbound packets for a given flow are allowed, but inbound packets are not.
Inbound flows between SVI's are denied by default? I'm not talking about the internet/WAN interface here. I'm talking about two VLAN's on the MX talking to each other. Like this
https://documentation.meraki.com/MX/Firewall_and_Traffic_Shaping/Creating_a_DMZ_with_the_MX_Security_Appliance
Ah, I get it now.
All flows initiated from behind the MX can be thought of as "outbound" flows, even if they are to another network segment behind the MX. Thus the "outbound" firewall can be used to block access to different local network segments.
eg you have VLAN A is 10.0.0.0/24, VLAN B is 10.0.1.0/24. To deny access between them you need to add the following rules:
deny, any any any src 10.0.0.0/24 dst 10.0.1.0/24
deny, any any any src 10.0.1.0/24 dst 10.0.0.0/24
Only using one of those rules allows uni-directional access. (ie one subnet can initiate flows to the other, but the reverse is not true)
This can get annoying when using a lot of VLANs, but if all your VLANs are in typical NAT address spaces you can use larger CIDRs in the rules (eg in the above example, you could use a single deny any any any src 10.0.0.0/8 dst 10.0.0.0/8 to accomplish the same thing)
You've got it. That's exactly how it works.
