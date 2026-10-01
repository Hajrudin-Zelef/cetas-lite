---
id: collect-261001-cisco/cisco/configuring-cisco-site-to-site-ipsec-vpn-with-dynamic-ip-endpoint-cisco-routers-2
title: "configuring-cisco-site-to-site-ipsec-vpn-with-dynamic-ip-endpoint-cisco-routers"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-cisco/configuring-cisco-site-to-site-ipsec-vpn-with-dynamic-ip-endpoint-cisco-routers.md
source_anchor: ""
source_lines: [157, 255]
sha256: 2f2ef5c0260f3fd4900290d8f76c414ab8475f44b121a706251224c0272da79e
---

# configuring-cisco-site-to-site-ipsec-vpn-with-dynamic-ip-endpoint-cisco-routers

For the headquarter router, deny NAT for packets destined to the remote VPN networks, but allow NAT for all other networks (Internet):

ip nat inside source list 100 interface fastethernet0/1 overload ! access-list 100 remark -=[Define NAT Service]=- access-list 100 deny ip 10.10.10.0 0.0.0.255 20.20.20.0 0.0.0.255 access-list 100 deny ip 10.10.10.0 0.0.0.255 30.30.30.0 0.0.0.255 access-list 100 permit ip 10.10.10.0 0.0.0.255 any access-list 100 remark

For Remote Site 1 Router, deny NAT for packets destined to the headquarter network:

ip nat inside source list 100 interface fastethernet0/1 overload

!

access-list 100 remark -=[Define NAT Service]=-

access-list 100 deny ip 20.20.20.0 0.0.0.255 10.10.10.0 0.0.0.255

access-list 100 permit ip 20.20.20.0 0.0.0.255 any

access-list 100 remark

For Remote Site 2 Router, deny NAT for packets destined to the headquarter network:

ip nat inside source list 100 interface fastethernet0/1 overload

!

access-list 100 remark -=[Define NAT Service]=-

access-list 100 deny ip 30.30.30.0 0.0.0.255 10.10.10.0 0.0.0.255

access-list 100 permit ip 30.30.30.0 0.0.0.255 any

access-list 100 remark

Bringing Up & Verifying The VPN Tunnel

At this point, we’ve completed our configuration and the VPN Tunnel is ready to be brought up. To initiate the VPN Tunnel, we need to force one packet to traverse the VPN and this can be achieved by pinging from one router to another. There is however one caveat that was mentioned in the beginning of this article:

Site to Site VPN networks with Dynamic remote Public IP addresses can only be brought up by the remote sites.

The reason for this is simple and logical. Only the remote site routers are aware of the headquarter’s public IP address (74.200.90.5) because it is static, and therefore only the remote router can initiate the VPN tunnel.

From Remote Site 1, let’s ping the headquarter router:

R2# ping 10.10.10.1 source fastethernet0/1 Type escape sequence to abort. Sending 5, 100-byte ICMP Echos to 10.10.10.1, timeout is 2 seconds: Packet sent with a source address of 73.54.120.100 .!!!! Success rate is 80 percent (4/5), round-trip min/avg/max = 42/46/5

The first ping received a timeout, but the rest received a reply, as expected. The time required to bring up the VPN Tunnel is sometimes slightly more than 2 seconds, causing the first ping to timeout.

To verify the VPN Tunnel, use the show crypto session command:

R2# show crypto session

Crypto session current status

Interface: FastEthernet0/1

Session status: UP-ACTIVE

Peer: 74.200.90.5 port 500

IKE SA: local 73.54.120.100/500 remote 74.200.90.5 /500 Active

IPSEC FLOW: permit ip 20.20.20.0/255.255.255.0 10.10.10.0/255.255.255.0

Active SAs: 2, origin: crypto map

From Remote Site 2, let’s ping the headquarter router:

R3# ping 10.10.10.1 source fastethernet0/1

Type escape sequence to abort.

Sending 5, 100-byte ICMP Echos to 10.10.10.1, timeout is 2 seconds:

Packet sent with a source address of 85.100.120.5

.!!!!

Success rate is 80 percent (4/5), round-trip min/avg/max = 47/50/53 ms

Again, the first ping received a timeout, but the rest received a reply, as expected. The time required to bring up the VPN Tunnel is sometimes slightly more than 2 seconds, causing the first ping to timeout.

To verify the VPN Tunnel, use the show crypto session command:

R3# show crypto session

Crypto session current status

Interface: FastEthernet0/1

Session status: UP-ACTIVE

Peer: 74.200.90.5 port 500

IKE SA: local 85.100.120.5/500 remote 74.200.90.5 /500 Active

IPSEC FLOW: permit ip 30.30.30.0/255.255.255.0 10.10.10.0/255.255.255.0

Active SAs: 2, origin: crypto map

Issuing the show crypto session command at the headquarter router will reveal all remote routers public IP addresses. This is usually a good shortcut when trying to figure out the public IP address of your remote routers.
