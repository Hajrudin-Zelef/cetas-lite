---
id: collect-260926-mikrotik/mikrotik/problem-with-port-forwarding-on-l009uigs-double-nat-dynamic-wanip
title: "2024-02-20 16:30:40 by RouterOS 7.12.1"
domain: mikrotik
role: reference
task: reference
actors: []
dates: ["2024-02-20"]
keywords: []
source: docs/RAG/lot-mikrotik/forum/firewall-nat/problem-with-port-forwarding-on-l009uigs-double-nat-dynamic-wanip.md
source_anchor: ""
source_lines: [1, 193]
sha256: 062edf66cb3fd6cd3095806e20f64ea3b8ccf11b3a8f3fa8ed0583c1d4b30cec
---

# 2024-02-20 16:30:40 by RouterOS 7.12.1

Hello

Yet another forwarding problem 

I have followed: https://help.mikrotik.com/docs/display/RKB/Port+forwarding

also tried using the: Quick Set > Port mapping

but they both have failed.

I have 2 Unifi APs connected to the router. The whole network is on 192.168.88.x

Let’s say I want to redirect the 8000 port to 192.168.88.253

I have ISP router plugged in ETH1 as WAN and for now I have double NAT (working on that) with DMZ.

The other subnet is 192.168.100.1 (and Mikrotik has .100 static IP)

I think I either get the In. interface wrong or it is some default firewall rule that blocks my forwarding.

It is my first time with Mikrotik so any help would be great!

Thanks!


Strange sidenote.

I am running python webserver to test it:

```
python3 -m http.server
```

And at first it works with both: http://192.168.88.253:8000/ or http://localhost:8000/

But when i try to connect later with http://192.168.88.1:8000/ or http://192.168.100.100:8000/ or http://my-public-ip:8000/

The localhost and local ip stops working as well. It all hangs on my mac & I need to restart the python webserver.

Even wget/curl hangs. Strange?




```
# 2024-02-20 16:30:40 by RouterOS 7.12.1
# software id = ABCD-8UDB
#
# model = L009UiGS
# serial number = ABCD
/interface bridge
add admin-mac=78:9A:18:62:43:F8 auto-mac=no comment=defconf name=bridge
/interface list
add comment=defconf name=WAN
add comment=defconf name=LAN
/interface wireless security-profiles
set [ find default=yes ] supplicant-identity=MikroTik
/ip pool
add name=dhcp ranges=192.168.88.20-192.168.88.254
/ip dhcp-server
add address-pool=dhcp interface=bridge lease-time=10m name=defconf
/port
set 0 name=serial0
/interface bridge port
add bridge=bridge comment=defconf interface=ether2
add bridge=bridge comment=defconf interface=ether3
add bridge=bridge comment=defconf interface=ether4
add bridge=bridge comment=defconf interface=ether5
add bridge=bridge comment=defconf interface=ether6
add bridge=bridge comment=defconf interface=ether7
add bridge=bridge comment=defconf interface=ether8
add bridge=bridge comment=defconf interface=sfp1
/ip neighbor discovery-settings
set discover-interface-list=LAN
/interface list member
add comment=defconf interface=bridge list=LAN
add comment=defconf interface=ether1 list=WAN
/ip address
add address=192.168.88.1/24 comment=defconf interface=bridge network=\
    192.168.88.0
/ip dhcp-client
add comment=defconf interface=ether1
/ip dhcp-server lease
add address=192.168.88.2 client-id=1:10:5b:ad:1:a6:28 mac-address=\
    10:5B:AD:01:A6:28 server=defconf
/ip dhcp-server network
add address=192.168.88.0/24 comment=defconf dns-server=192.168.88.1 gateway=\
    192.168.88.1
/ip dns
set allow-remote-requests=yes
/ip dns static
add address=192.168.88.1 comment=defconf name=router.lan
/ip firewall filter
add action=accept chain=input comment=\
    "defconf: accept established,related,untracked" connection-state=\
    established,related,untracked
add action=drop chain=input comment="defconf: drop invalid" connection-state=\
    invalid
add action=accept chain=input comment="defconf: accept ICMP" protocol=icmp
add action=accept chain=input comment=\
    "defconf: accept to local loopback (for CAPsMAN)" dst-address=127.0.0.1
add action=drop chain=input comment="defconf: drop all not coming from LAN" \
    in-interface-list=!LAN
add action=accept chain=forward comment="defconf: accept in ipsec policy" \
    ipsec-policy=in,ipsec
add action=accept chain=forward comment="defconf: accept out ipsec policy" \
    ipsec-policy=out,ipsec
add action=fasttrack-connection chain=forward comment="defconf: fasttrack" \
    connection-state=established,related hw-offload=yes
add action=accept chain=forward comment=\
    "defconf: accept established,related, untracked" connection-state=\
    established,related,untracked
add action=drop chain=forward comment="defconf: drop invalid" \
    connection-state=invalid
add action=drop chain=forward comment=\
    "defconf: drop all from WAN not DSTNATed" connection-nat-state=!dstnat \
    connection-state=new in-interface-list=WAN
/ip firewall nat
add action=masquerade chain=srcnat comment="defconf: masquerade" \
    ipsec-policy=out,none out-interface-list=WAN
add action=dst-nat chain=dstnat comment=http8000 dst-address=0.0.0.0 \
    dst-port=8000 in-interface-list=all protocol=tcp to-addresses=\
    192.168.88.253 to-ports=8000
/ipv6 firewall address-list
add address=::/128 comment="defconf: unspecified address" list=bad_ipv6
add address=::1/128 comment="defconf: lo" list=bad_ipv6
add address=fec0::/10 comment="defconf: site-local" list=bad_ipv6
add address=::ffff:0.0.0.0/96 comment="defconf: ipv4-mapped" list=bad_ipv6
add address=::/96 comment="defconf: ipv4 compat" list=bad_ipv6
add address=100::/64 comment="defconf: discard only " list=bad_ipv6
add address=2001:db8::/32 comment="defconf: documentation" list=bad_ipv6
add address=2001:10::/28 comment="defconf: ORCHID" list=bad_ipv6
add address=3ffe::/16 comment="defconf: 6bone" list=bad_ipv6
/ipv6 firewall filter
add action=accept chain=input comment=\
    "defconf: accept established,related,untracked" connection-state=\
    established,related,untracked
add action=drop chain=input comment="defconf: drop invalid" connection-state=\
    invalid
add action=accept chain=input comment="defconf: accept ICMPv6" protocol=\
    icmpv6
add action=accept chain=input comment="defconf: accept UDP traceroute" port=\
    33434-33534 protocol=udp
add action=accept chain=input comment=\
    "defconf: accept DHCPv6-Client prefix delegation." dst-port=546 protocol=\
    udp src-address=fe80::/10
add action=accept chain=input comment="defconf: accept IKE" dst-port=500,4500 \
    protocol=udp
add action=accept chain=input comment="defconf: accept ipsec AH" protocol=\
    ipsec-ah
add action=accept chain=input comment="defconf: accept ipsec ESP" protocol=\
    ipsec-esp
add action=accept chain=input comment=\
    "defconf: accept all that matches ipsec policy" ipsec-policy=in,ipsec
add action=drop chain=input comment=\
    "defconf: drop everything else not coming from LAN" in-interface-list=\
    !LAN
add action=accept chain=forward comment=\
    "defconf: accept established,related,untracked" connection-state=\
    established,related,untracked
add action=drop chain=forward comment="defconf: drop invalid" \
    connection-state=invalid
add action=drop chain=forward comment=\
    "defconf: drop packets with bad src ipv6" src-address-list=bad_ipv6
add action=drop chain=forward comment=\
    "defconf: drop packets with bad dst ipv6" dst-address-list=bad_ipv6
add action=drop chain=forward comment="defconf: rfc4890 drop hop-limit=1" \
    hop-limit=equal:1 protocol=icmpv6
add action=accept chain=forward comment="defconf: accept ICMPv6" protocol=\
    icmpv6
add action=accept chain=forward comment="defconf: accept HIP" protocol=139
add action=accept chain=forward comment="defconf: accept IKE" dst-port=\
    500,4500 protocol=udp
add action=accept chain=forward comment="defconf: accept ipsec AH" protocol=\
    ipsec-ah
add action=accept chain=forward comment="defconf: accept ipsec ESP" protocol=\
    ipsec-esp
add action=accept chain=forward comment=\
    "defconf: accept all that matches ipsec policy" ipsec-policy=in,ipsec
add action=drop chain=forward comment=\
    "defconf: drop everything else not coming from LAN" in-interface-list=\
    !LAN
/system clock
set time-zone-name=Europe/Bucharest
/system note
set show-at-login=no
/system routerboard settings
set enter-setup-on=delete-key
/tool mac-server
set allowed-interface-list=LAN
/tool mac-server mac-winbox
set allowed-interface-list=LAN
```
