---
id: collect-261001-general-networking/general-networking/manual-how-tos-nat-reflection-html-7747e718-1
title: "manual-how-tos-nat-reflection-html-7747e718"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["attention"]
source: docs/RAG/collect-261001-general-networking/manual-how-tos-nat-reflection-html-7747e718.md
source_anchor: ""
source_lines: [1, 57]
sha256: 3baeeb5f70b9f3dfc6ce48716679d60d909b16c11eaa5a42399218d94c6601e1
---

# manual-how-tos-nat-reflection-html-7747e718

Reflection and Hairpin NAT
Networks used in this How-To section
| Interface | IPv4 Subnet | Hosts | Gateway | 
|---|---|---|---|
| WAN | 203.0.113.0/24 | 203.0.113.1 - OPNsense | 203.0.113.254 - OPNsense | 
| DMZ | 172.16.1.0/24 | 172.16.1.1 - Webserver | 172.16.1.254 - OPNsense | 
| LAN | 192.168.1.0/24 | 192.168.1.1 - Client | 192.168.1.254 - OPNsense | 
NAT - Quick Overview
Because there are not enough available IPv4 addresses, a workaround called NAT (Network Address Translation) was implemented into the IPv4 Standard. It basically enables a router like the OPNsense to translate IPv4 addresses to other IPv4 addresses. Most of the time it is used to translate the limited external IPv4 address space to the shared internal IPv4 address space (RFC 1918, 192.168.0.0/16 - 172.16.0.0/12 - 10.0.0.0/8) and vice versa.
Note
- SNAT - Source Network Address Translation
  - Changes the source IP of a packet
  - Firewall –> NAT –> Source NAT (Outbound) using the option Translate Source IP in a rule
- DNAT - Destination Network Address Translation
  - Changes the destination IP of a packet
  - Firewall –> NAT –> Destination NAT (Port Forward) using the option Redirect target IP in a rule
- PAT - Port Address Translation
  - Changes the destination port of a packet
  - Firewall –> NAT –> Destination NAT (Port Forward) using the option Redirect target port in a rule
If you create a DNAT rule, you enable all clients in the WAN access to an internal IPv4 address. The OPNsense acts like a translator, translating IPv4 addresses between client and server. The OPNsense writes all translations into a file called the NAT table. It knows exactly how traffic should flow back and forth with the translations in place.
Warning
NAT is not a security feature. It only acts as a translator. If you want security, you need firewall rules in addition.
Introduction to Reflection and Hairpin NAT
For example, you have a Webserver example.com with the internal IP 172.16.1.1 in your DMZ. It has a public DNS Record of example.com in A 203.0.113.1.
Your internal client 192.168.1.1 can’t reach the Webserver if it resolves the DNS A-Record 203.0.113.1. When the OPNsense receives the packet from the client 192.168.1.1 with the destination IP 203.0.113.1, it chooses itself as the target, and not 172.16.1.1. That’s because the external IPv4 address 203.0.113.1 is mapped to the WAN interface of the OPNsense.
That’s where Reflection NAT comes into play. It creates NAT rules which help your internal client 192.168.1.1 to communicate with your webserver 203.0.113.1, by using the OPNsense as the “translator” to the actual destination 172.16.1.1.
Attention
You should choose your preferred Reflection NAT method from the three possible choices presented here. They’re exclusive to each other, picking one method and sticking to it will prevent mistakes.
- Method 1 - Creating manual Port-Forward NAT (DNAT), manual Source NAT (SNAT), and automatic firewall rules
- Method 2 - Creating automatic Port-Forward NAT (DNAT), manual Source NAT (SNAT), and manual firewall rules
- Method 3 - Creating automatic Port-Forward NAT (DNAT), automatic Source NAT (SNAT), and manual firewall rules
Note
- Reflection NAT: The client and the server are in different subnets (layer 2 broadcast domains) and the OPNsense routes traffic between them. They can’t communicate directly by resolving ARP requests. You only need DNAT.
- Hairpin NAT: The client and the server are in the same subnet (layer 2 broadcast domain). They can communicate directly with each other by resolving ARP requests. You need SNAT and DNAT.
Note
When using IPsec, by default NAT only matches on policy based VPN. NAT on VTI (Virtual Tunnel Interfaces) won’t match unless some tunables are set. These tunables change the behavior of firewall filter and NAT on if_enc and if_ipsec interfaces. You can read more about the tunables in IPsec VTI - Route based setup
Best Practice
The best way to do Reflection NAT in the OPNsense is not to use the legacy Reflection options in (Advanced) Settings. Creating the NAT rules manually with Method 1 prevents unwanted traffic and makes auditing easy. There will be no hidden rules. All rules will be perfectly visible in the GUI and .xml config exports.
Start of the How-To Section:
The goal is to access the Webserver 172.16.1.1 on port 443 with it’s external IP 203.0.113.1 from a client in WAN, LAN and DMZ.
Method 1 - Creating manual Port-Forward NAT (DNAT), manual Source NAT (SNAT), and automatic firewall rules
- Go to
- Disable Reflection for Destination NAT (Port Forwards), Reflection for 1:1 and Automatic Source NAT (Outbound) for Reflection
- Go to
- Select + to create a new Destination NAT (Port Forward) rule. Interface: Select WAN ,DMZ andLAN - Select all interfaces in which clients are that should access the webserver. This will create a linked Firewall rule in  which allows the traffic.Protocol: Select TCPSource: Select AnySource port range: Select AnyDestination: Input 203.0.113.1 - It’s the external IPv4 address of the webserver.Destination port range: Input 443 - Or select the aliasHTTPSRedirect target IP: Input 172.16.1.1 - It’s the Webserver’s internal IPv4 address in the DMZ.Redirect target port: Input 443 - Or select the aliasHTTPSDescription: Input Reflection NAT Rule Webserver 443 - Add a description because the linked Filter rule association will use that as its name and the  rule will have it in the description.NAT reflection: Use system default Filter rule association: Add associated filter rule
Tip
Reading the DNAT rule like a sentence makes it clearer:
If a packet is received by the OPNsense on any of the interfaces WAN, DMZ and LAN with protocol TCP from the source IP ANY and the source port range ANY to destination
IP 203.0.113.1 and destination port 443 –> rewrite the destination IP to 172.16.1.1 and the destination port to 443.
Note
Due to “Add associated filter rule”, the added linked firewall rule in   will allow traffic to the destination IP 172.16.1.1 because NAT rules match before Firewall rules. That means the firewall receives the packet and the NAT rule converts the destination from 203.0.113.1 to 172.16.1.1 first, before passing the packet to the firewall filter. You could also set “Filter rule association: Pass”, but then the resulting firewall rule would be invisible.
Note
In some setups (e.g. an external IP address is bound on an additional VPN interface) you need to set “Filter rule association: None” and create your own Firewall rules. One of those firewall rules should match only on the VPN interface, and in “advanced features” of that rule “reply-to” should be your VPN interface. The other firewall rule (without “reply-to”) should match the remaining interfaces.
Attention
Now you have Reflection NAT. The traffic from the internal LAN client 192.168.1.1 and any WAN client reaches the Webserver.
But there is a caveat - any DMZ client and the Webserver itself are still unable reach the external IP 203.0.113.1. For that you need Hairpin NAT, which involves an additional SNAT rule.
- Go to
