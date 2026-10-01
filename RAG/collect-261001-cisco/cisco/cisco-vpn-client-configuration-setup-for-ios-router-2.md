---
id: collect-261001-cisco/cisco/cisco-vpn-client-configuration-setup-for-ios-router-2
title: "cisco-vpn-client-configuration-setup-for-ios-router"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-cisco/cisco-vpn-client-configuration-setup-for-ios-router.md
source_anchor: ""
source_lines: [71, 108]
sha256: 6b65b0490edc24ec4782c540ebf82902a8a0fe20748bdb9caa2a9b51593fea4f
---

# cisco-vpn-client-configuration-setup-for-ios-router

R1(config)# crypto isakmp profile vpn-ike-profile-1 R1(conf-isa-prof)# match identity group CCLIENT-VPN R1(conf-isa-prof)# client authentication list vpn_xauth_ml_1 R1(conf-isa-prof)# isakmp authorization list vpn_group_ml_1 R1(conf-isa-prof)# client configuration address respond R1(conf-isa-prof)# virtual-template 2

Last step is the creation of our access lists that will control the VPN traffic to be tunnelled, effectively controlling what our VPN users are able to access remotely.

Once that's done, we need to add a 'no NAT' statement so that traffic exiting the router and heading toward the VPN user is preserved with its private IP address, otherwise packets sent through the tunnel by the router, will be NAT'ed and therefore rejected by the remote VPN Client.

When NAT is enabled through a VPN tunnel, the remote user sees the tunnelled traffic coming from the router's public IP address, when in fact it should be from the router's private IP address.

We assume the following standard NAT configuration to provide Internet access to the company's LAN network:

R1#show running-config <output omitted> ip nat inside source list 100 interface Dialer1 overload access-list 100 remark -=[Internet NAT Service]=- access-list 100 permit ip 192.168.0.0 0.0.0.255 any access-list 100 remark

Based on the above, we proceed with our configuration. First, we need to restrict access to our remote VPN users, so that they only access our SQL server with IP address 192.168.0.6 (access-list 120), then we deny NAT (access-list 100) to our remote VPN Pool IP range:

R1(config)# no access-list 100 R1(config)# access-list 100 remark [Deny NAT for VPN Clients]=- R1(config)# access-list 100 deny ip 192.168.0.0 0.0.0.255 host 192.168.0.20 R1(config)# access-list 100 deny ip 192.168.0.0 0.0.0.255 host 192.168.0.21 R1(config)# access-list 100 deny ip 192.168.0.0 0.0.0.255 host 192.168.0.22 R1(config)# access-list 100 deny ip 192.168.0.0 0.0.0.255 host 192.168.0.23 R1(config)# access-list 100 deny ip 192.168.0.0 0.0.0.255 host 192.168.0.24 R1(config)# access-list 100 deny ip 192.168.0.0 0.0.0.255 host 192.168.0.25 R1(config)# access-list 100 remark R1(config)# access-list 100 remark -=[Internet NAT Service]=- R1(config)# access-list 100 permit ip 192.168.0.0 0.0.0.255 any

Note that for access-list 100, we could either 'deny ip host 192.168.0.6' to our remote clients, or as shown, deny the 192.168.0.0/24 network. What's the difference? Practically none. Denying your whole network the NAT service toward your remote clients, will make it easier for any future additions.

If for example there was a need to deny NAT for another 5 servers so they can reach remote VPN clients, then the access-list 100 would need to be edited to include these new hosts, where as now it's already taken care of. Remember, with access-list 100 we are simply controlling the NAT function , not the access the remote clients have (done with access-list 120 in our example.

At this point, the Cisco VPN configuration is complete and fully functional.

VPN - Split Tunneling

We mentioned in the beginning of this article that we would cover split tunneling and full tunneling methods for our VPN clients. You'll be pleased to know that this functionality is solely determined by the group's access-lists, which our case is access-list 120.

If we wanted to tunnel all traffic from the VPN client to our network, we would use the following access-list 120 configuration:

R1(config)# access-list 120 remark ==[Cisco VPN Users]== R1(config)# access-list 120 permit ip any host 192.168.0.20 R1(config)# access-list 120 permit ip any host 192.168.0.21 R1(config)# access-list 120 permit ip any host 192.168.0.22 R1(config)# access-list 120 permit ip any host 192.168.0.23 R1(config)# access-list 120 permit ip any host 192.168.0.24 R1(config)# access-list 120 permit ip any host 192.168.0.25

In another example, if we wanted to provide our VPN clients access to networks 10.0.0.0/24, 10.10.10.0/24 & 192.168.0.0/24, here's what the access-list 120 would look like (this scenario requires modification of NAT access-list 100 as well):

R1(config)# access-list 120 remark ==[Cisco VPN Users]== R1(config)# access-list 120 permit ip 10.0.0.0 0.0.0.255 host 192.168.0.20 R1(config)# access-list 120 permit ip 10.0.0.0 0.0.0.255 host 192.168.0.21 R1(config)# access-list 120 permit ip 10.0.0.0 0.0.0.255 host 192.168.0.22 R1(config)# access-list 120 permit ip 10.0.0.0 0.0.0.255 host 192.168.0.23 R1(config)#access-list 120 permit ip 10.0.0.0 0.0.0.255 host 192.168.0.24 R1(config)# access-list 120 permit ip 10.0.0.0 0.0.0.255 host 192.168.0.25 R1(config)# R1(config)# access-list 120 permit ip 10.10.10.0 0.0.0.255 host 192.168.0.20 R1(config)# access-list 120 permit ip 10.10.10.0 0.0.0.255 host 192.168.0.21 R1(config)# access-list 120 permit ip 10.10.10.0 0.0.0.255 host 192.168.0.22 R1(config)# access-list 120 permit ip 10.10.10.0 0.0.0.255 host 192.168.0.23 R1(config)# access-list 120 permit ip 10.10.10.0 0.0.0.255 host 192.168.0.24 R1(config)#access-list 120 permit ip 10.10.10.0 0.0.0.255 host 192.168.0.25 R1(config)# R1(config)# R1(config)# access-list 120 permit ip 192.168.0.0 0.0.0.255 host 192.168.0.20 R1(config)# access-list 120 permit ip 192.168.0.0 0.0.0.255 host 192.168.0.21 R1(config)# access-list 120 permit ip 192.168.0.0 0.0.0.255 host 192.168.0.22 R1(config)#access-list 120 permit ip 192.168.0.0 0.0.0.255 host 192.168.0.23 R1(config)# access-list 120 permit ip 192.168.0.0 0.0.0.255 host 192.168.0.24 R1(config)# access-list 120 permit ip 192.168.0.0 0.0.0.255 host 192.168.0.25 R1(config)# R1(config)# R1(config)# no access-list 100 R1(config)# access-list 100 remark [Deny NAT for VPN Clients]=- R1(config)# access-list 100 deny ip 10.0.0.0 0.0.0.255 host 192.168.0.20 R1(config)#access-list 100 deny ip 10.0.0.0 0.0.0.255 host 192.168.0.21 R1(config)# access-list 100 deny ip 10.0.0.0 0.0.0.255 host 192.168.0.22 R1(config)# access-list 100 deny ip 10.0.0.0 0.0.0.255 host 192.168.0.23 R1(config)# access-list 100 deny ip 10.0.0.0 0.0.0.255 host 192.168.0.24 R1(config)# access-list 100 deny ip 10.0.0.0 0.0.0.255 host 192.168.0.25 R1(config)# R1(config)# R1(config)#access-list 100 deny ip 10.10.10.0 0.0.0.255 host 192.168.0.20 R1(config)#access-list 100 deny ip 10.10.10.0 0.0.0.255 host 192.168.0.21 R1(config)# access-list 100 deny ip 10.10.10.0 0.0.0.255 host 192.168.0.22 R1(config)#access-list 100 deny ip 10.10.10.0 0.0.0.255 host 192.168.0.23 R1(config)# access-list 100 deny ip 10.10.10.0 0.0.0.255 host 192.168.0.24 R1(config)# access-list 100 deny ip 10.10.10.0 0.0.0.255 host 192.168.0.25 R1(config)# R1(config)# R1(config)# access-list 100 deny ip 192.168.0.0 0.0.0.255 host 192.168.0.20 R1(config)# access-list 100 deny ip 192.168.0.0 0.0.0.255 host 192.168.0.21 R1(config)#access-list 100 deny ip 192.168.0.0 0.0.0.255 host 192.168.0.22 R1(config)# access-list 100 deny ip 192.168.0.0 0.0.0.255 host 192.168.0.23 R1(config)# access-list 100 deny ip 192.168.0.0 0.0.0.255 host 192.168.0.24 R1(config)#access-list 100 deny ip 192.168.0.0 0.0.0.255 host 192.168.0.25 R1(config)# access-list 100 remark R1(config)#access-list 100 remark -=[Internet NAT Service]=- R1(config)# access-list 100 permit ip 10.0.0.0 0.0.0.255 any R1(config)# access-list 100 permit ip 10.10.10.0 0.0.0.255 any R1(config)# access-list 100 permit ip 192.168.0.0 0.0.0.255 any

When the VPN client connects, should we go to the connection's statistics, we would see the 3 networks under the secure routes, indicating all traffic toward these networks is tunnelled through the VPN:

It is evident from our last example with the tunneling of our 3 networks, that should our VPN IP address pool be larger, for example 50 IP addresses, then we would have to enter 50 IPs x 3 Networks = 150 lines of code just for the access-list 120, plus another 150 lines for access-list 100 (no NAT)! That is quite a task indeed!

