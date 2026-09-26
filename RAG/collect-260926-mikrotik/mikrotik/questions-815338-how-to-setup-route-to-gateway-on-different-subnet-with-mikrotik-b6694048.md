---
id: collect-260926-mikrotik/mikrotik/questions-815338-how-to-setup-route-to-gateway-on-different-subnet-with-mikrotik-b6694048
title: "questions-815338-how-to-setup-route-to-gateway-on-different-subnet-with-mikrotik-b6694048"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/lot-mikrotik/forum/routing-bgp-ospf/questions-815338-how-to-setup-route-to-gateway-on-different-subnet-with-mikrotik-b6694048.md
source_anchor: ""
source_lines: [1, 32]
sha256: a7417a78c7c494094c8fe8d876835d0eaa6a691b1d9ddeb866177401a7ca0750
---

# questions-815338-how-to-setup-route-to-gateway-on-different-subnet-with-mikrotik-b6694048

We have the following setup which I need to get working:
- Location A:
  - Subnet 1:
    - Network: 192.168.1.0
    - Mask: 255.255.255.0
    - Default Gateway: 192.168.1.1
  - Subnet 2:
    - Network: 192.168.2.0
    - Mask: 255.255.255.0
    - Default Gateway: 192.168.2.1
- Subnet 1:
- Location B:
  - Subnet 3:
    - Network: 192.168.3.0
    - Mask: 255.255.255.0
    - Default Gateway: 192.168.3.1
- Subnet 3:
The default gateways of location A and B are connected via a VPN and ONLY route subnets 1 and 3. I cannot change the config of these gateways. What I need to achieve is, that traffic going to a public subnet, e.g. 193.197.0.0 is routed through the default gateway of subnet 2.
What I did so far:
- I added a MikroTik router at location A and B which is configured as follows:
  - Location A:
    - IP: 192.168.1.254, 192.168.2.254
    - Route: 193.197.0.0 -> 192.168.2.254
  - Location B:
    - IP: 192.168.3.254
- Location A:
Now, for location A and subnet 1 this works fine. However, I fail to setup the MikroTik at location 3 properly. I tried:
- Route: 193.197.0.0/12 -> 192.168.1.254 (that does not work)
- Setup IPIP-tunnel between the two MikroTik and route 193.197.0.0/12 -> (IPIP-tunnel)
Background: clients in subnet 1 and 3 need to access a certain website. However, the website can only be contacted through the gateway in subnet 2 (some high security stuff).
I'm pretty sure that this scenario can be solved with these routers but I need your help! Any ideas?
location 3, but you don't give any details about the network or addresses assigned to the router at location 3, unless you are using confusing names. Are the VPNs terminating on these routers? If so, they have probably additional IP addresses and subnets you haven't told us about. Those are the addresses you probably should be using for your routes. Anyway, you need to fill out the details, and proofread to make sure that someone else can follow your description./ip address printon each router, or just look at the IP -> Addresses screen and tell us all the IPs.
