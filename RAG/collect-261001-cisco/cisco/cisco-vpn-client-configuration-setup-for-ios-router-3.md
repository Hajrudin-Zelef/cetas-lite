---
id: collect-261001-cisco/cisco/cisco-vpn-client-configuration-setup-for-ios-router-3
title: "cisco-vpn-client-configuration-setup-for-ios-router"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-cisco/cisco-vpn-client-configuration-setup-for-ios-router.md
source_anchor: ""
source_lines: [109, 139]
sha256: a580b6a9d8c688f8f511a509ee461057277195c92e832b24c957b96f03bcf57c
---

# cisco-vpn-client-configuration-setup-for-ios-router

To help cut down the configuration to just a couple of lines, this is the alternative code that would be used and have the same effect:

Do not NAT any traffic from our LANs toward VPN clients, but NAT everything else destined to the Internet:

R1(config)# access-list 100 remark [Deny NAT for VPN Clients]=- R1(config)# access-list 100 deny ip 10.0.0.0 0.0.0.255 192.168.0.0 0.0.0.255 R1(config)# access-list 100 deny ip 10.10.10.0 0.0.0.255 192.168.0.0 0.0.0.255 R1(config)# access-list 100 deny ip 192.168.0.0 0.0.0.255 192.168.0.0 0.0.0.255 R1(config)# access-list 100 remark R1(config)# access-list 100 remark -=[Internet NAT Service]=- R1(config)# access-list 100 permit ip 10.0.0.0 0.0.0.255 any R1(config)# access-list 100 permit ip 10.10.10.0 0.0.0.255 any R1(config)# access-list 100 permit ip 192.168.0.0 0.0.0.255 any

The access-list 120 instructs the router to tunnel all traffic from the three networks to our VPN clients who's IP address will be in the 192.168.0.0/24 range!

So, if the VPN client received from the VPN Pool, IP address 192.168.0.23 or 192.168.0.49, it really wouldn't matter as the '192.168.0.0 0.0.0.255' statement at the end of each access-list 120 covers both 192.168.0.23 & 192.168.0.49. Even replacing the '192.168.0.0 0.0.0.255' with the 'any' statement would have the same effect.

For 'access-list 100' that controls the NAT service, we cannot use the 'any' statement at the end of the DENY portion of the ACLs, because it would exclude NAT for all networks (public and private) therefore completely disabling NAT and as a result, Internet access.

As a last note, if it was required the VPN clients to be provided with an IP address range different from that of the internal network (e.g 192.168.50.0/24), then the following minor changes to the configuration would have to be made:

R1(config)# crypto isakmp client configuration group CCLIENT-VPN R1(config-isakmp-group)# key firewall.cx R1(config-isakmp-group)# dns 10.0.0.10 R1(config-isakmp-group)# pool VPN-Pool R1(config-isakmp-group)# acl 120 R1(config-isakmp-group)# max-users 5 R1(config-isakmp-group)# exit R1(config)# R1(config)# ip local pool VPN-Pool 192.168.50.10 192.168.50.25 R1(config)# R1(config)# interface Virtual-Template2 type tunnel R1(config-if)# ip address 192.168.50.1 255.255.255.0 R1(config-if)# tunnel mode ipsec ipv4 R1(config-if)# tunnel protection ipsec profile VPN-Profile-1

Assuming 3 internal networks Mark VPN Traffic to be tunnelled:

Do not NAT any traffic from our LANs toward VPN clients, but NAT everything else destined to the Internet:

R1(config)# access-list 100 remark [Deny NAT for VPN Clients]=- R1(config)# access-list 100 deny ip 10.0.0.0 0.0.0.255 192.168.50.0 0.0.0.255 R1(config)# access-list 100 deny ip 10.10.10.0 0.0.0.255 192.168.50.0 0.0.0.255 R1(config)# access-list 100 deny ip 192.168.0.0 0.0.0.255 192.168.50.0 0.0.0.255 R1(config)# access-list 100 remark R1(config)# access-list 100 remark -=[Internet NAT Service]=- R1(config)# access-list 100 permit ip 10.0.0.0 0.0.0.255 any R1(config)# access-list 100 permit ip 10.10.10.0 0.0.0.255 any R1(config)# access-list 100 permit ip 192.168.0.0 0.0.0.255 any

Summary

This article explained the fundamentals of Cisco's VPN client and features it offers to allow the remote and secure connection of users to their corporate networks from anywhere in the world.

We examined the necessary steps and commands required on a Cisco router to setup and configure it to accept Cisco VPN client connections. Detailed explanation was provided for every configuration step, along with the necessary diagrams and screenshots.

Split tunneling was explained and covered, showing how to configure the Cisco VPN clients access only to the required internal networks while maintaining access to the Internet.

Lastly, a few tips were presented to help make the Cisco VPN configuration a lot easier for large and more complex networks.
