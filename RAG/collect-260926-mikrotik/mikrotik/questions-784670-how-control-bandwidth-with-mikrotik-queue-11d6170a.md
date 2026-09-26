---
id: collect-260926-mikrotik/mikrotik/questions-784670-how-control-bandwidth-with-mikrotik-queue-11d6170a
title: "questions-784670-how-control-bandwidth-with-mikrotik-queue-11d6170a"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: ["ethernet"]
source: docs/RAG/lot-mikrotik/forum/qos/questions-784670-how-control-bandwidth-with-mikrotik-queue-11d6170a.md
source_anchor: ""
source_lines: [1, 100]
sha256: 915fb837f8b7a1e4cf4634a95a4b39e3250f7ff9ca62269cd0aff5708b8ba8b5
---

# questions-784670-how-control-bandwidth-with-mikrotik-queue-11d6170a

I have a mikrotik router and i want to limit the speed of wlan users. I create a queue for target = wlan and set the limit downloa and upload = 64k, but when i test it, this seems not working correctly! Test download speed > 500k. The router config is default. How can i do that
/interface bridge
add admin-mac=E4:8D:8C:46:C7:0F auto-mac=no comment=defconf name=bridge
/interface wireless
set [ find default-name=wlan1 ] band=2ghz-b/g/n channel-width=20/40mhz-Ce \
    disabled=no distance=indoors frequency=auto mode=ap-bridge ssid=PNG1 \
    wireless-protocol=802.11
/interface ethernet
set [ find default-name=ether2 ] name=ether2-master
set [ find default-name=ether3 ] master-port=ether2-master
set [ find default-name=ether4 ] master-port=ether2-master
set [ find default-name=ether5 ] master-port=ether2-master
/ip neighbor discovery
set ether1 discover=no
set bridge comment=defconf
/interface wireless security-profiles
set [ find default=yes ] authentication-types=wpa-psk,wpa2-psk mode=\
    dynamic-keys wpa-pre-shared-key=noor3664 wpa2-pre-shared-key=noor3664
/ip hotspot profile
add dns-name=so.zzz hotspot-address=10.5.50.1 login-by=http-pap name=hsprof1 \
    use-radius=yes
/ip pool
add name=dhcp ranges=192.168.88.10-192.168.88.254
add name=hs-pool-2 ranges=10.5.50.2-10.5.50.254
/ip dhcp-server
add address-pool=dhcp authoritative=after-10sec-delay disabled=no interface=\
    bridge name=defconf
add address-pool=hs-pool-2 interface=ether2-master lease-time=1h name=dhcp1
/ip hotspot
add address-pool=hs-pool-2 addresses-per-mac=1 interface=ether2-master name=\
    hotspot1 profile=hsprof1
/ip hotspot user profile
add address-pool=hs-pool-2 advertise=yes advertise-interval="" advertise-url="" \
    name=uprof1 open-status-page=http-login transparent-proxy=yes
/queue type
set 0 pfifo-limit=200
/queue interface
set ether2-master queue=default-small
/queue simple
add limit-at=30k/30k max-limit=30k/30k name=queue1 queue=default/default \
    target=bridge total-queue=default
/interface bridge port
add bridge=bridge comment=defconf interface=ether2-master
add auto-isolate=yes bridge=bridge interface=wlan1
/ip address
add address=192.168.88.1/24 comment=defconf interface=bridge network=\
    192.168.88.0
add address=10.5.50.1/24 comment="hotspot network" interface=ether2-master \
    network=10.5.50.0
/ip dhcp-client
add comment=defconf dhcp-options=hostname,clientid disabled=no interface=ether1
/ip dhcp-server network
add address=10.5.50.0/24 comment="hotspot network" gateway=10.5.50.1
add address=192.168.88.0/24 comment=defconf gateway=192.168.88.1
/ip dns
set allow-remote-requests=yes servers=4.2.2.4
/ip dns static
add address=10.5.50.1 name=router
/ip firewall filter
add chain=input comment="defconf: accept ICMP" protocol=icmp
add chain=input comment="defconf: accept established,related" connection-state=\
    established,related
add action=drop chain=input comment="defconf: drop all from WAN" in-interface=\
    ether1
add action=fasttrack-connection chain=forward comment="defconf: fasttrack" \
    connection-state=established,related
add chain=forward comment="defconf: accept established,related" \
    connection-state=established,related
add action=drop chain=forward comment="defconf: drop invalid" connection-state=\
    invalid
add action=drop chain=forward comment=\
    "defconf:  drop all from WAN not DSTNATed" connection-nat-state=!dstnat \
    connection-state=new in-interface=ether1
/ip firewall nat
add action=masquerade chain=srcnat comment="defconf: masquerade" out-interface=\
    ether1
add action=masquerade chain=srcnat comment="masquerade hotspot net
    src-address=10.5.50.0/24
add action=masquerade chain=srcnat comment="masquerade hotspot net
    src-address=10.5.50.0/24
add action=masquerade chain=srcnat comment="masquerade hotspot net
    src-address=10.5.50.0/24
/ip hotspot user
add name=admin
/radius
add address=192.168.88.1 service=hotspot
add address=192.168.88.1 service=hotspot
/system clock
set time-zone-name=Asia/Tehran
/system leds
set 0 interface=wlan1
/system routerboard settings
set protected-routerboot=disabled
/tool mac-server
set [ find default=yes ] disabled=yes
add interface=bridge
/tool mac-server mac-winbox
set [ find default=yes ] disabled=yes
add interface=bridge
[admin@MikroTik] >
