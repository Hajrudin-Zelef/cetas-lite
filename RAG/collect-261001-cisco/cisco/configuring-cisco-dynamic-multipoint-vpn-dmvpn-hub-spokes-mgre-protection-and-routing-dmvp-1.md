---
id: collect-261001-cisco/cisco/configuring-cisco-dynamic-multipoint-vpn-dmvpn-hub-spokes-mgre-protection-and-routing-dmvp-1
title: "Ent Peer NBMA Addr Peer Tunnel Add State UpDn Tm Attrb ----- ------------- --------------- ----- ------- ----- 1 2.2.2.10 172.16.0.2 UP 00:04:58 D 1 3.3.3.10 172.16.0.3 UP 00:04:12 D"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: ["ethernet"]
source: docs/RAG/collect-261001-cisco/configuring-cisco-dynamic-multipoint-vpn-dmvpn-hub-spokes-mgre-protection-and-routing-dmvpn-configur.md
source_anchor: ""
source_lines: [1, 218]
sha256: 5d88aa1d7f26d85a73277ef9fe00197c4c7974f6fc336780bb1c6076ac38abf3
---

# Ent Peer NBMA Addr Peer Tunnel Add State UpDn Tm Attrb ----- ------------- --------------- ----- ------- ----- 1 2.2.2.10 172.16.0.2 UP 00:04:58 D 1 3.3.3.10 172.16.0.3 UP 00:04:12 D

Our DMVPN Introduction article covered the DMVPN concept and deployment designs. We explained how DMVPN combines a number of technologies that give it its flexibility, low administrative overhead and ease of configuration. This article will cover the configuration of a Cisco DMVPN including Hub, Spokes, Routing and Protecting the mGRE Tunnel.

It is highly advisable for those who haven’t read our Introduction to DMVPN to do so as it contains basic concepts and theory that are important to the configuration process.

Configuring DMVPN is simple, if you’ve worked with GRE tunnels before. If the GRE Tunnel concept is new to you, we would recommend reading through our Point-to-Point GRE IPSec Tunnel Configuration article before proceeding with DMVPN configuration.

DMVPN as a design concept is essentially the configuration combination of protected GRE Tunnel and Next Hop Routing Protocol (NHRP).

Before diving into the configuration of our routers, we’ll briefly explain how the DMVPN is expected to work. This will help in understanding how DMVPN operates in a network:

Each spoke has a permanent IPSec tunnel to the hub but not to the other spokes within the network.

Each spoke registers as a client of the NHRP server. The Hub router undertakes the role of the NHRP server.

When a spoke needs to send a packet to a destination (private) subnet on another spoke, it queries the NHRP server for the real (outside) address of the destination (target) spoke.

After the originating spoke learns the peer address of the target spoke, it can initiate a dynamic IPSec tunnel to the target spoke.

The spoke-to-spoke tunnel is built over the multipoint GRE (mGRE) interface.

The spoke-to-spoke links are established on demand whenever there is traffic between the spokes. Thereafter, packets are able to bypass the hub and use the spoke-to-spoke tunnel.

All data traversing the GRE tunnel is encrypted using IPSecurity (optional)

Our DMVPN Network

The diagram below depicts our DMVPN example network. Our goal is to connect the two remote networks (Remote 1 & 2) with the company headquarters. The headquarters router R1 is the central Hub router that will hold the NHRP database containing all spoke routers, their public IP addresses and LAN networks.

Four Steps To Fully Configure Cisco DMVPN

To help simplify the configuration of DMVPN we’ve split the process into 4 easy-to-follow steps. Each step is required to be completed before moving to the next one. These steps are:

Configure the DMVPN Hub

Configure the DMVPN Spoke(s)

Protect the mGRE tunnels with IPSecurity (optional)

Configure Routing Between DMVPN mGRE Tunnels (static routing or routing protocol)

Configuring The DMVPN Hub – R1 Router

Configuring the Hub router (R1) is simple. After configuring the router’s LAN and WAN interfaces we create our mGRE tunnel interface. Let's start with the router’s Ethernet interfaces:

interface FastEthernet0/0

description LAN-Network

ip address 192.168.1.1 255.255.255.0

duplex auto

speed auto

!

interface FastEthernet0/1

description WAN-Network

ip address 1.1.1.10 255.255.255.0

duplex auto

speed auto

Next, we configure the Tunnel0 interface. Notice this is an almost typical tunnel interface configuration with some minor but important changes that have been highlighted:

interface Tunnel0

description mGRE - DMVPN Tunnel

ip address 172.16.0.1 255.255.255.0

no ip redirects

ip nhrp authentication firewall

ip nhrp map multicast dynamic

ip nhrp network-id 1

tunnel source 1.1.1.10

tunnel mode gre multipoint

Engineers familiar with GRE Tunnels will immediately notice the absence of the tunnel destination command. It has been replaced with the tunnel mode gre multipointcommand, which designates this tunnel as a multipoint GRE tunnel.

The ip nhrp map multicast dynamic command enables the forwarding of multicast traffic across the tunnel to dynamic spokes. This is usually required by routing protocols such as OSPF and EIGRP. In most cases, DMVPN is accompanied by a routing protocol to send and receive dynamic updates about the private networks.

The ip nhrp network-id 1 command is used to identify this DMVPN cloud. All routers participating in this DMVPN cloud must have the same network-id configured in order for tunnels to form between them.

The ip nhrp authentication command is used to allow the authenticated updates and queries to the NHRP Database, ensuring unwanted queries are not provided with any information about the DMVPN network.

Configuring The DMVPN Spokes – R2 & R3 Routers

Spoke router configuration is similar to that of the hub. First configure the LAN and WAN interfaces:

interface FastEthernet0/0

description LAN-Network

ip address 192.168.2.1 255.255.255.0

duplex auto

speed auto

!

interface FastEthernet0/1

description WAN-Network

ip address 2.2.2.10 255.255.255.0

duplex auto

speed auto

Next, it’s time to build that mGRE tunnel:

interface Tunnel0

description R2 mGRE - DMVPN Tunnel

ip address 172.16.0.2 255.255.255.0

no ip redirects

ip nhrp authentication firewall

ip nhrp map multicast dynamic

ip nhrp map 172.16.0.1 1.1.1.10

ip nhrp map multicast 1.1.1.10

ip nhrp network-id 1

ip nhrp nhs 172.16.0.1

tunnel source FastEthernet0/1

tunnel mode gre multipoint

After a couple of seconds, we receive confirmation that our tunnel interface is up:

*Sep 9 21:27:29.774: %LINEPROTO-5-UPDOWN: Line protocol on Interface Tunnel0, changed state to up

The ip nhrp nhs 172.16.0.1 command tells our spoke router who the Next Hop Server (NHS) is, while the ip nhrp map 172.16.0.1 1.1.1.10 command maps the NHS address (172.16.0.1) to the Hub’s (R1) public IP address (1.1.1.10).

The ip nhrp map multicast 1.1.1.10 ensures multicast traffic is sent only from spokes to the hub and not from spoke to spoke. All multicast traffic should be received by the hub, processed and then updates are sent out to the spokes.

Lastly, notice that tunnel source FastEthernet0/1 command. All spokes with dynamic WAN IP address must be configured to bind the physical WAN interface as the tunnel source. This way, when the spoke’s WAN IP changes, it will be able to update the NHS server with its new WAN IP address.

Note: In R2’s configuration, we’ve configured a static IP address on its WAN interface FastEthernet0/1, but for the sake of this example, let us assume it was dynamically provided by the ISP.

R3’s configuration follows, similar to that of the R2 spoke router:

interface FastEthernet0/0

description LAN-Network

ip address 192.168.3.1 255.255.255.0

duplex auto

speed auto

!

interface FastEthernet0/1

description WAN-Network

ip address 3.3.3.10 255.255.255.0

duplex auto

speed auto

Next, our tunnel configuration:

interface Tunnel0

description R3 mGRE - DMVPN Tunnel

ip address 172.16.0.3 255.255.255.0

no ip redirects

ip nhrp authentication firewall

ip nhrp map multicast dynamic

ip nhrp map 172.16.0.1 1.1.1.10

ip nhrp map multicast 1.1.1.10

ip nhrp network-id 1

ip nhrp nhs 172.16.0.1

tunnel source FastEthernet0/1

tunnel mode gre multipoint

Note: In R3’s configuration, we’ve configured a static IP address on its WAN interface FastEthernet0/1, but for the sake of this example, let us assume it was dynamically provided by the ISP.

This completes the DMVPN configuration on our central hub and two spoke routers. It is now time to verify the DMVPNs are working correctly.

Verifying DMVPN Functionality At The R1 HUB Router

After completing our routers configuration, it’s time to verify everything is working as planned.

First we turn to our main hub, R1, and check the DMVPN by using the show dmvpn command:

