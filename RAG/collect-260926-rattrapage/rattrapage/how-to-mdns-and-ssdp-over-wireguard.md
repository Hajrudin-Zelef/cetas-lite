---
id: collect-260926-rattrapage/rattrapage/how-to-mdns-and-ssdp-over-wireguard
title: "how-to-mdns-and-ssdp-over-wireguard"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["ethernet"]
source: docs/RAG/lot-rattrapage/servers-reviews/how-to-mdns-and-ssdp-over-wireguard.md
source_anchor: ""
source_lines: [1, 145]
sha256: 61c69837a4817c7347d950beef8e461e53ced7616788cfd3e0f9f48622ebfdb5
---

# how-to-mdns-and-ssdp-over-wireguard

This is a guide for getting mDNS (Bonjour) and SSDP (for DLNA) working across a Wireguard interface linking two Mikrotik routers running ROS7.7 or greater without needing fluff like IGMP Proxy/PIM/Avahi/Containers.

The attachment below shows the implementation.  Bold above the flag symbol indicates actual interfaces.  The heavy vertical lines are a shared network layer.

The routers’ bridges are not using VLAN-filtering as it’s usually not necessary in this case for home routers.



**Wireguard**

Side A:

```
/interface wireguard
add listen-port=13231 mtu=1412 name=Wireguard
/interface ireguard peers
add allowed-address=172.16.200.0/24 endpoint-address=site-b.com \
    endpoint-port=13231 interface=Wireguard public-key=\
    "<side a's public key>"
/ip route
add disabled=no distance=1 dst-address=172.16.200.0/24 gateway=Wireguard \
    pref-src="" routing-table=main scope=30 suppress-hw-offload=no \
    target-scope=10
```

Side B:

```
/interface wireguard
add listen-port=13231 mtu=1412 name=Wireguard
/interface ireguard peers
add allowed-address=172.16.100.0/24 endpoint-address=site-a.com \
    endpoint-port=13231 interface=Wireguard public-key=\
    "<side b's public key>"
/ip route
add disabled=no distance=1 dst-address=172.16.100.0/24 gateway=Wireguard \
    pref-src="" routing-table=main scope=30 suppress-hw-offload=no \
    target-scope=10
```

This is a typical Wireguard config, don’t forget to allow your firewall to accept UDP on port 13231 on the input chain for Wireguard traffic.

In this case I have set the MTU to 1412 down from the default 1420 as one side of the link uses PPPoE.  You will need to adjust this to suit your connection.

The routes are needed of course so each router can find the subnet on the opposing side.



**EoIP**

Side A:

```
/interface eoip
add !keepalive local-address=172.16.100.254 mtu=\
    1500 name=EoIP remote-address=172.16.200.254 tunnel-id=1
```

Side B:

```
/interface eoip
add !keepalive local-address=172.16.200.254 mtu=\
    1500 name=EoIP remote-address=172.16.100.254 tunnel-id=1
```

Here we set up the EoIP interface.  No IPSEC is needed as it runs through the Wireguard link.

Don’t forget to add the EoIP port to the main bridge at each end.

Side A and B:

```
/interface bridge port
add bridge=BridgeMain interface=EoIP
```

**Bridge Filtering**

At this stage both bridges are linked in the broadcast domain which will be a disaster if left unfiltered.  Any broadcasts including DHCP requests and replies will flow both ways.

We just want to let mDNS and SSDP broadcasts through and absolutely nothing else and this can be done by using the Bridge Filter - I think a powerful feature of ROS usually forgotten lying in a dusty corner.

Side A and B:

```
/interface bridge filter
add action=accept chain=forward comment="Allow mDNS" dst-address=224.0.0.251/32 \
    dst-mac-address=01:00:5E:00:00:FB/FF:FF:FF:FF:FF:FF dst-port=5353 \
    ip-protocol=udp mac-protocol=ip out-interface=EoIP src-port=5353
add action=accept chain=forward comment="Allow SSDP" dst-address=239.255.255.250/32 \
    dst-mac-address=01:00:5E:7F:FF:FA/FF:FF:FF:FF:FF:FF dst-port=1900 \
    ip-protocol=udp log-prefix=SSDP mac-protocol=ip out-interface=EoIP
add action=drop chain=forward out-interface=EoIP
add action=drop chain=output out-interface=EoIP
```

This filtering will preserve the ethernet frames’ source MAC addresses which start with 01: and are needed to ensure proper flooding on the networks at the other side.

mDNS traffic both ways seems to all be done with broadcasts.  The contents of the mDNS packets will contain IP addresses of the services and once a client learns of the service will communicate over the normal Wireguard route on layer 3.

SSDP (DLNA) discovery traffic is a broadcast from the client to find out what servers are available.  The server replies with a unicast message on layer 3 to that client by sending a UDP packet back to the source IP and UDP port the client send the broadcast from.

In my case the DLNA server is MythTV but due to a security issue from 2014 it now only replies to client broadcasts from the subnets it’s a member of.  Other DLNA servers might have the same behaviour.  I had to make some DSTNAT and SRCNAT rules to fool it.

On the router as the same side as MythTV I had these NAT rules.

The router at that side has a gateway address of 172.16.100.254 so for my own clarity I added an address of 172.16.100.253.  The TV on the other side is 172.16.200.237, MythTV is 172.16.100.200.

The src-nat rule makes the TV’s IP address appear to come from the same subnet as the MythTV when the discovery broadcast comes through.  It still preserves the MAC frame source address of 01:etc. so it can be flooded to the subnet.

The dst-nat rule takes the unicast reply from MythTV which thinks it’s replying to 172.16.100.253 (the routers other address) and rewrites it to the TV’s address which then goes over Wireguard.  Any further communication between MythTV and the TV Client is done over the normally routed unicast and doesn’t need NATting.




```
/ip firewall nat
add action=src-nat chain=srcnat comment="SSDP broadcast" dst-address=239.255.255.250 \
    src-address=172.16.200.237 to-addresses=172.16.100.253
add action=dst-nat chain=dstnat comment="SSDP unicast reply" dst-address=172.16.100.253 \
    src-address=172.16.100.200 to-addresses=172.16.200.237
```

Using this example as a framework you could possibly use the following substitutions but it’s out of the scope of this document.

EoIP: VLANX, VPLS (?), OpenVPN TAP

Wireguard: L2TP, PPP, IPSec, OpenVPN TUN, GRE



*References:*

Wireguard: https://help.mikrotik.com/docs/display/ROS/WireGuard

EoIP: https://help.mikrotik.com/docs/display/ROS/EoIP

Bridge Firewall: https://help.mikrotik.com/docs/display/ROS/Bridging+and+Switching#BridgingandSwitching-BridgeFirewall

mDNS: https://en.wikipedia.org/wiki/Multicast_DNS

SSDP: https://en.wikipedia.org/wiki/Simple_Service_Discovery_Protocol

Forum discussion: http://forum.mikrotik.com/t/mdns-repeater-feature/148334/179
