---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/blog-it-2022-02-02-wireguad-and-split-vpn-on-unifi-dream-machine-pro-se-24e3a435-4
title: "This script downloads the latest split-vpn and installs it"
domain: unifi-ubiquiti
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-unifi-ubiquiti/blog-it-2022-02-02-wireguad-and-split-vpn-on-unifi-dream-machine-pro-se-24e3a435.md
source_anchor: ""
source_lines: [255, 285]
sha256: dccbd91c23652cd2ede62e5e4f80a2060d7874810b322bbdc561add0a5413a68
---

# This script downloads the latest split-vpn and installs it

If a data packet is compared to a transport truck, with the header being the truck itself and the payload being the trailer and cargo, then MSS is like a scale that measures only the trailer. If the trailer weighs too much, then the truck is not allowed to continue to its destination.
More specifically, MSS is the largest TCP (Transport Control Protocol) segment size that a network-connected device can receive. MSS defines “segment” as only the length of the payload, not any attached headers. MSS is measured in bytes.
MSS is determined by another metric that has to do with packet size: MTU, or the maximum transmission unit, which does include the TCP and IP (Internet Protocol) headers. To continue the analogy, MTU measures the total weight of the truck and its trailer and cargo, instead of just the trailer and cargo.
Essentially, the MSS is equal to MTU minus the size of a TCP header and an IP header:
 MSS = MTU - (TCP header + IP header) 
So for Wireguard, the MSS is 1420 bytes - 20 (IP header) bytes - 20 (TCP header) byte = 1380 byte
One of the key differences between MTU and MSS is that if a packet exceeds a device’s MTU, it is broken up into smaller pieces, or “fragmented.” In contrast, if a packet exceeds the MSS, it is dropped and not delivered.
- What is MSS clamping?
Occasionally, a router along a network path has an MTU value set lower than the typical 1,500 bytes. This can result in packet loss and can be difficult to discover.
To ensure packets still reach their destination in this situation, one option is to reduce the size of incoming packet payloads. This can be achieved by configuring the server to apply an MSS clamp: during the TCP handshake, the server can signal the MSS for packets it is willing to receive, “clamping” the maximum payload size from the other server. For example, if servers A and B are establishing a TCP connection and server B communicates an MSS of 1436 bytes, server A will send packets with a maximum payload size of 1436 bytes for the duration of the connection.
Another application of MSS clamping is in the case of GRE tunneling, where a 24 bytes header is added to the original packet in order to send it to a new destination. If the original packet was larger than 1476 bytes, this could make the new packet exceed the typical 1500 bytes MTU; an MSS clamp can be applied to require incoming packets to be less than 1,500 bytes even after the GRE header is applied.
Note that usually, it is not needed to set MSS clamping manually, but some VPN connections stall if the MSS clamping is not set correctly. Typical values range from 1240 to 1460 bytes, but it could be lower.
 I fixe the MSS clamping to 1380 in vpn.conf with MSS_CLAMPING_IPV4 directive
Tips: Wireguad® allowed IPs calculator Permalink
Follow instructions here to calculate allowed IP.
Tips: Test DNS leak Permalink
What is a DNS leak and why should I care?
When using an anonymity or privacy service, it is extremely important that all traffic originating from your computer is routed through the anonymity network. If any traffic leaks outside of the secure connection to the network, any adversary monitoring your traffic will be able to log your activity.
DNS or the domain name system is used to translate domain names such as www.eff.org into numerical IP addresses (e.g. 123.123.123.123) which are required to route packets of data on the Internet. Whenever your computer needs to contact a server on the Internet, such as when you enter a URL into your browser, your computer contacts a DNS server and requests the IP address. Most Internet service providers assign their customers a DNS server which they control and use for logging and recording your Internet activities.
Under certain conditions, even when connected to the anonymity network, the operating system will continue to use its default DNS servers instead of the anonymous DNS servers assigned to your computer by the anonymity network. DNS leaks are a major privacy threat since the anonymity network may be providing a false sense of security while private data is leaking.
If you are concerned about DNS leaks, you should also understand transparent DNS proxy technology to ensure that the solution you choose will stop
You can test DNS leak on dnsleaktest.com or bash.ws
dnsleaktest.com only shows the v4 servers queried. bash.ws lists both v4 and v6.
Tips: How do you check your clients are on the VPN? Permalink
On your client, check if you are seeing the VPN IPs when you visit https://browserleaks.com or https://ifconfig.co.
You can also test from command line, by running the following commands from your clients. Make sure you are not seeing your real IP anywhere, either IPv4 or IPv6.
curl -4 ifconfig.co   #IPv4
curl -6 ifconfig.co   #IPv6
If you are seeing your real IPv6 address above, make sure that you are forcing your client through IPv6 as well as IPv4, by forcing through interface, MAC address, or the IPv6 directly. If IPv6 is not supported by your VPN provider, the IPv6 check will time out and not return anything. You should never see your real IPv6 address.
Tips: VPN provider doesn’t support IPv6 Permalink
If your VPN provider doesn’t support IPv6, it is recommended to disable IPv6 for that VLAN in the UDM settings, or on the client, so that you don’t encounter any delays. If you don’t disable IPv6, clients on that network will try to communicate over IPv6 first and fail, then fallback to IPv4. This creates a delay that can be avoided if IPv6 is turned off completely for that network or client.
