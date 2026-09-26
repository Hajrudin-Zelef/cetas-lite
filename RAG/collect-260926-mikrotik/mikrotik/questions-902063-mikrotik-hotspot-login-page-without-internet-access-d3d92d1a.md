---
id: collect-260926-mikrotik/mikrotik/questions-902063-mikrotik-hotspot-login-page-without-internet-access-d3d92d1a
title: "questions-902063-mikrotik-hotspot-login-page-without-internet-access-d3d92d1a"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: ["ethernet"]
source: docs/RAG/lot-mikrotik/forum/dhcp-dns/questions-902063-mikrotik-hotspot-login-page-without-internet-access-d3d92d1a.md
source_anchor: ""
source_lines: [1, 99]
sha256: 41aa00075bd1580aef3906e207ad186e4f9ea45f427a2856ae51c304b0c75dfb
---

# questions-902063-mikrotik-hotspot-login-page-without-internet-access-d3d92d1a

Good day. I created hotspot on Mikrotik device (RouterOS v. 6.27) When I plug in internet cable, everything works - any url opened by client is redirected to the login page. When logged in then no redirection.
But I need the same functionality without internet - connect to wifi, get html page no matter on URL entered, nothing more.
The problem is that redirection does not work when no internet is plugged in. The hotspot page can be open, and ask for the login, but it has no effect.
/interface bridge
add admin-mac=4C:5E:0C:0B:D5:17 auto-mac=no name=bridge-local
/interface wireless
set [ find default-name=wlan1 ] band=2ghz-b/g/n channel-width=\
    20/40mhz-Ce disabled=no distance=indoors frequency=auto \
    l2mtu=1600 mode=ap-bridge rx-chains=0,1 ssid="VVR HotSpot" \
    tx-chains=0,1 wireless-protocol=802.11
/interface ethernet
set [ find default-name=ether1 ] name=ether1-gateway
set [ find default-name=ether2 ] name=ether2-master-local
set [ find default-name=ether3 ] master-port=ether2-master-local \
    name=ether3-slave-local
set [ find default-name=ether4 ] master-port=ether2-master-local \
    name=ether4-slave-local
/ip neighbor discovery
set ether1-gateway discover=no
/ip hotspot profile
add dns-name=hotspot.vvr hotspot-address=10.10.10.1 login-by=\
    cookie,http-pap name=hsprof2
/ip pool
add name=default-dhcp ranges=192.168.88.10-192.168.88.254
add name=hs-pool-5 ranges=10.10.10.2-10.10.10.254
/ip dhcp-server
add address-pool=default-dhcp disabled=no interface=bridge-local \
    name=default
add address-pool=hs-pool-5 disabled=no interface=wlan1 \
    lease-time=1h name=dhcp1
/ip hotspot
add address-pool=hs-pool-5 disabled=no interface=wlan1 name=\
    hotspot2 profile=hsprof2
/interface bridge port
add bridge=bridge-local interface=ether2-master-local
add bridge=bridge-local disabled=yes interface=wlan1
/ip address
add address=192.168.88.1/24 comment="default configuration" \
    interface=bridge-local network=192.168.88.0
add address=10.10.10.1/24 comment="hotspot network" interface=\
    wlan1 network=10.10.10.0
/ip dhcp-client
add comment="default configuration" dhcp-options=\
    hostname,clientid disabled=no interface=ether1-gateway
/ip dhcp-server network
add address=10.10.10.0/24 comment="hotspot network" dns-server=\
    10.10.10.1 gateway=10.10.10.1
add address=192.168.88.0/24 comment="default configuration" \
    gateway=192.168.88.1
/ip dns
set allow-remote-requests=yes cache-max-ttl=0s \
    query-server-timeout=0ms query-total-timeout=0ms
/ip dns static
add address=192.168.88.1 name=router
/ip firewall filter
add action=passthrough chain=unused-hs-chain comment=\
    "place hotspot rules here" disabled=yes
add chain=input comment="default configuration" protocol=icmp
add chain=input comment="default configuration" connection-state=\
    established,related
add action=drop chain=input comment="default configuration" \
    in-interface=ether1-gateway
add chain=forward comment="default configuration" \
    connection-state=established,related
add action=drop chain=forward comment="default configuration" \
    connection-state=invalid
add action=drop chain=forward comment="default configuration" \
    connection-nat-state=!dstnat connection-state=new \
    in-interface=ether1-gateway
/ip firewall nat
add action=passthrough chain=unused-hs-chain comment=\
    "place hotspot rules here" disabled=yes
add action=masquerade chain=srcnat comment=\
    "default configuration" out-interface=ether1-gateway
add action=masquerade chain=srcnat comment=\
    "masquerade hotspot network" src-address=10.10.10.0/24
add action=masquerade chain=srcnat comment=\
    "masquerade hotspot network" src-address=10.10.10.0/24
/ip hotspot user
add name=admin password=vvradmin
/system clock
set time-zone-name=Europe/Kiev
/system routerboard settings
set cpu-frequency=650MHz protected-routerboot=disabled
/tool mac-server
set [ find default=yes ] disabled=yes
add interface=ether2-master-local
add interface=ether3-slave-local
add interface=ether4-slave-local
add interface=wlan1
add interface=bridge-local
/tool mac-server mac-winbox
set [ find default=yes ] disabled=yes
add interface=ether2-master-local
add interface=ether3-slave-local
add interface=ether4-slave-local
add interface=wlan1
add interface=bridge-local
Help me please to find a problem or configure it from starting point.
