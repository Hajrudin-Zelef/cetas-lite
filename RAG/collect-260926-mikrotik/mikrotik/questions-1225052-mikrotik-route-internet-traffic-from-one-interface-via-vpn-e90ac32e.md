---
id: collect-260926-mikrotik/mikrotik/questions-1225052-mikrotik-route-internet-traffic-from-one-interface-via-vpn-e90ac32e
title: "jul/02/2017 19:49:03 by RouterOS 6.39.2"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: ["ethernet"]
source: docs/RAG/lot-mikrotik/forum/vpn/questions-1225052-mikrotik-route-internet-traffic-from-one-interface-via-vpn-e90ac32e.md
source_anchor: ""
source_lines: [1, 52]
sha256: ab893e55416c826aa957ca2fa0ef23bc50d8ea574d500cac4029da958ce15071
---

# jul/02/2017 19:49:03 by RouterOS 6.39.2

below is my current config of my mikrotik hAP lite. As you see the config is pretty basic. I have another router in my network (192.168.178.1) that does DHCP and provides internet connectivity. The port ether1 is the uplink to this router.
All devices connected to ether2, ether3 and the wlan should be in the network as if they were just connected via a normal switch.
This works, but the config might not be optimal for that. If I can do better, please advise.
The device on ether4 should also be available in the network, and should also be able to access all other devices on the network as normal, however when the device connected at ether4 wants to access the internet (= i.e. make DNS requests to or sends traffic via the default gateway 192.168.178.1), this traffic should be re-routed and be sent via the configured VPN instead. The default gateway of the VPN is dynamically assigned (its the l2tpclient interface named my-vpn) and currently has the ip 10.9.9.1.
As you can probably tell I tried that with the firewall mangle rule, but that did not work.
What do I need to remove / change / add to make this work?
Cheers,
Sebastian
# jul/02/2017 19:49:03 by RouterOS 6.39.2
/interface bridge
add admin-mac=AA:BB:CC:AA:BB:CC auto-mac=no comment=defconf fast-forward=no \
    name=bridge
/interface wireless
set [ find default-name=wlan1 ] band=2ghz-b/g/n country=germany disabled=no \
    frequency=auto mode=ap-bridge ssid=test wireless-protocol=802.11
/interface ethernet
set [ find default-name=ether2 ] master-port=ether1
set [ find default-name=ether3 ] master-port=ether1
/interface l2tp-client
add connect-to=some.vpn.com disabled=no name=my-vpn password=\
    test user=test
/interface wireless security-profiles
set [ find default=yes ] authentication-types=wpa2-psk eap-methods="" \
    group-ciphers=tkip,aes-ccm mode=dynamic-keys unicast-ciphers=tkip,aes-ccm \
    wpa-pre-shared-key=test wpa2-pre-shared-key=test
/interface bridge port
add bridge=bridge interface=ether1
add bridge=bridge interface=wlan1
add bridge=bridge interface=ether4
/ip dhcp-client
add comment=defconf dhcp-options=hostname,clientid disabled=no interface=\
    bridge
/ip dns
set allow-remote-requests=yes
/ip dns static
add address=192.168.88.1 name=router
/ip firewall filter
add action=fasttrack-connection chain=forward comment="defconf: fasttrack" \
    connection-state=established,related
add action=accept chain=forward comment="defconf: accept established,related" \
    connection-state=established,related
add action=drop chain=forward comment="defconf: drop invalid" \
    connection-state=invalid
# in/out-interface matcher not possible when interface (ether1) is slave - use master instead (bridge)
add action=drop chain=forward comment=\
    "defconf:  drop all from WAN not DSTNATed" connection-nat-state=!dstnat \
    connection-state=new in-interface=ether1
/ip firewall mangle
add action=route chain=prerouting dst-address=192.168.178.1 log=yes \
    log-prefix=test passthrough=yes route-dst=10.9.9.1
/system clock
set time-zone-name=Europe/Berlin
