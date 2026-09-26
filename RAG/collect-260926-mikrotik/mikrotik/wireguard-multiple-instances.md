---
id: collect-260926-mikrotik/mikrotik/wireguard-multiple-instances
title: "wireguard-multiple-instances"
domain: mikrotik
role: reference
task: reference
actors: ["Apple"]
dates: []
keywords: []
source: docs/RAG/lot-mikrotik/forum/wireguard/wireguard-multiple-instances.md
source_anchor: ""
source_lines: [1, 65]
sha256: b268c9e5f6aa00f663594a471663b84ec17fb1ca7cf2131397d694ad0bb897c7
---

# wireguard-multiple-instances

Thanks anav,

The reason the first site is disabled is so that my remote site works.

I do want to keep them separate ( one for iPhones/Laptops/Tablets ) road warrior scenario and second for the S2S .

the problem is that for some reason the S2S (Kubara) somehow wants to connect to the WG-HQ site even thought the ports are different for both. Only if I disable the WG-HQ site (interface) then it skips over to the WG-S2S. Should it do it ?

Just to reiterate the S2S configuration is working and routing if the WG-HQ is disabled. What I do not get is why the remote site is trying to connect on port 13231 even thought it is configured for 13232..

To answer other questions I don’t need to have the phones to connect to the remote site.

I do have the routes set as you have stated. The configuration is huge so I thought I would only post the one related to WG.

so here it is again updated with routes



Main Site

/interface wireguard

add disabled=yes listen-port=13231 mtu=1420 name=WG-HQ

add listen-port=13232 mtu=1420 name=WG-S2S

/interface wireguard peers

add allowed-address=10.10.100.2/32,192.168.220.0/24 interface=WG-S2S name=WG-Kubara public-key=“-+PK MK interface±”

add allowed-address=10.10.0.3/32 comment=iPhone16Pro interface=WG-HQ name=iphone16pro public-key=“-+phone PK±”

add address=10.10.100.1/24 comment=“WG 2412” interface=WG-S2S network=10.10.100.0



/ip route

add comment=WireGuard-Kubara disabled=no distance=1 dst-address=192.168.220.0/24 gateway=WG-S2S routing-table=main scope=30 suppress-hw-offload=no  target-scope=10

this route is added once the interface is started

DAc  10.10.100.0/24    WG-S2S                0

Remote Site

/interface wireguard

add listen-port=13232 mtu=1420 name=WG-Kubara

/interface wireguard peers

add allowed-address=10.10.100.1/32,192.168.26.0 endpoint-address=myhomeendpoint endpoint-port=13232 interface=WG-Kubara name=“WG-S2S Kubara” persistent-keepalive=10m public-key=“-+PK MK interface±”

add address=10.10.100.2/24 interface=WG-Kubara network=10.10.100.0

/ip route

add comment=“WireGuard - 6.x Backups” disabled=no distance=1 dst-address=192.168.26.0/24 gateway=WG-Kubara pref-src=“” routing-table=main scope=30 suppress-hw-offload=no target-scope=10

this route is added once the interface is started

DAc  10.10.100.0/24    WG-Kubara                0

Is there any other configuration related to WG I need to share  to help me solve this issue?
