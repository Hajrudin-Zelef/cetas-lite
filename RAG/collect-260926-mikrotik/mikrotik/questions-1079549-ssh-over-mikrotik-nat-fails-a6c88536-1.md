---
id: collect-260926-mikrotik/mikrotik/questions-1079549-ssh-over-mikrotik-nat-fails-a6c88536-1
title: "oct/04/2021 23:17:07 by RouterOS 6.47.4"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/lot-mikrotik/forum/firewall-nat/questions-1079549-ssh-over-mikrotik-nat-fails-a6c88536.md
source_anchor: ""
source_lines: [1, 113]
sha256: dced72f9f3916c0a898b1b0efaf0d7f3e887f4cbe49e9f70ec1d11f24eaf946d
---

# oct/04/2021 23:17:07 by RouterOS 6.47.4

I have a server on my network which I want to expose to external SSH connections. I can ssh directly to the device from my network (e.g. ssh 192.168.88.162 works fine). I have a NAT rule set up so that connections to 17722 reroute to 22. However this fails:
PS C:\Users\Me> ssh -vvv -i .\.ssh\id_rsa -p 17722 me@160.119.XXX.XXX
OpenSSH_for_Windows_8.1p1, LibreSSL 3.0.2
debug1: Reading configuration data C:\\Users\\Me/.ssh/config
debug3: Failed to open file:C:/ProgramData/ssh/ssh_config error:2
debug2: resolve_canonicalize: hostname 160.119.XXX.XXX is address
debug2: ssh_connect_direct
debug1: Connecting to 160.119.XXX.XXX [160.119.XXX.XXX] port 17722.
debug3: finish_connect - ERROR: async io completed with error: 10060, io:00000222C310DC10
debug1: connect to address 160.119.XXX.XXX port 17722: Connection timed out
ssh: connect to host 160.119.XXX.XXX port 17722: Connection timed out
How can I get the ssh from the external IP work?
tcpdump on the server shows:
me@JanJansen:~ $ grep 64236 tcpdump
22:13:56.097727 IP 192.168.88.177.64236 > 192.168.88.162.ssh: Flags [S], seq 3490646443, win 64240, options [mss 1460,nop,wscale 8,nop,nop,sackOK], length 0
22:13:56.098213 IP 192.168.88.162.ssh > 192.168.88.177.64236: Flags [S.], seq 869880002, ack 3490646444, win 64240, options [mss 1460,nop,nop,sackOK,nop,wscale 6], length 0
22:13:57.105046 IP 192.168.88.177.64236 > 192.168.88.162.ssh: Flags [S], seq 3490646443, win 64240, options [mss 1460,nop,wscale 8,nop,nop,sackOK], length 0
22:13:57.105398 IP 192.168.88.162.ssh > 192.168.88.177.64236: Flags [S.], seq 869880002, ack 3490646444, win 64240, options [mss 1460,nop,nop,sackOK,nop,wscale 6], length 0
22:13:58.162258 IP 192.168.88.162.ssh > 192.168.88.177.64236: Flags [S.], seq 869880002, ack 3490646444, win 64240, options [mss 1460,nop,nop,sackOK,nop,wscale 6], length 0
22:13:59.117541 IP 192.168.88.177.64236 > 192.168.88.162.ssh: Flags [S], seq 3490646443, win 64240, options [mss 1460,nop,wscale 8,nop,nop,sackOK], length 0
22:13:59.117912 IP 192.168.88.162.ssh > 192.168.88.177.64236: Flags [S.], seq 869880002, ack 3490646444, win 64240, options [mss 1460,nop,nop,sackOK,nop,wscale 6], length 0
22:14:01.122237 IP 192.168.88.162.ssh > 192.168.88.177.64236: Flags [S.], seq 869880002, ack 3490646444, win 64240, options [mss 1460,nop,nop,sackOK,nop,wscale 6], length 0
Compared to a local ssh which has a lot of [P] flags, not [S], which is the biggest difference I see, but I have no idea how to act on this.
For what it's worth, I can see the NAT rules in effect. Here's the router config:
# oct/04/2021 23:17:07 by RouterOS 6.47.4
# software id = VBLW-UG4R
#
# model = 951Ui-2HnD
# serial number = B8710C65021A
/interface bridge
add admin-mac=48:8F:5A:79:92:71 auto-mac=no comment=defconf name=bridge
/interface wireless
set [ find default-name=wlan1 ] mode=ap-bridge ssid=MikroTik wireless-protocol=802.11
/interface pppoe-client
add add-default-route=yes disabled=no interface=ether5 name=pppoe-out1 use-peer-dns=yes user=gkruger@websquad.co.za
/interface list
add comment=defconf name=WAN
add comment=defconf name=LAN
/interface wireless security-profiles
set [ find default=yes ] supplicant-identity=MikroTik
/ip pool
add name=dhcp ranges=192.168.88.10-192.168.88.254
/ip dhcp-server
add address-pool=dhcp disabled=no interface=bridge lease-time=12h name=defconf
/interface bridge port
add bridge=bridge comment=defconf interface=ether2
add bridge=bridge comment=defconf interface=ether3
add bridge=bridge comment=defconf interface=ether4
add bridge=bridge comment=defconf interface=wlan1
/ip neighbor discovery-settings
set discover-interface-list=LAN
/interface list member
add comment=defconf interface=bridge list=LAN
add comment=defconf interface=ether1 list=WAN
add interface=pppoe-out1 list=WAN
/ip address
add address=192.168.88.1/24 comment=defconf interface=ether2 network=192.168.88.0
/ip arp
add address=192.168.88.162 comment=JanJansen interface=bridge mac-address=00:0F:13:39:20:33
add address=192.168.88.177 comment=Sarevok interface=bridge mac-address=40:8D:5C:58:C0:97
add address=192.168.88.202 interface=bridge mac-address=32:63:2A:49:58:D9
add address=192.168.88.101 interface=bridge mac-address=18:56:80:24:47:12
/ip cloud
set ddns-enabled=yes
/ip dhcp-client
add comment=defconf interface=ether1
/ip dhcp-server network
add address=192.168.88.0/24 comment=defconf gateway=192.168.88.1
/ip dns
set allow-remote-requests=yes
/ip dns static
add address=192.168.88.1 comment=defconf name=router.lan
/ip firewall filter
add action=accept chain=input comment="defconf: accept established,related,untracked" connection-state=established,related,untracked
add action=drop chain=input comment="defconf: drop invalid" connection-state=invalid
add action=accept chain=input comment="defconf: accept ICMP" protocol=icmp
add action=accept chain=input comment="defconf: accept to local loopback (for CAPsMAN)" dst-address=127.0.0.1
add action=drop chain=input comment="defconf: drop all not coming from LAN" in-interface-list=!LAN
add action=accept chain=forward comment="defconf: accept in ipsec policy" ipsec-policy=in,ipsec
add action=accept chain=forward comment="defconf: accept out ipsec policy" ipsec-policy=out,ipsec
add action=fasttrack-connection chain=forward comment="defconf: fasttrack" connection-state=established,related
add action=accept chain=forward comment="defconf: accept established,related, untracked" connection-state=established,related,untracked
add action=drop chain=forward comment="defconf: drop invalid" connection-state=invalid
add action=drop chain=forward comment="defconf: drop all from WAN not DSTNATed" connection-nat-state=!dstnat connection-state=new in-interface-list=WAN
add action=accept chain=forward comment="Outside SSH" dst-port=22 log=yes protocol=tcp
/ip firewall nat
add action=masquerade chain=srcnat comment="defconf: masquerade" ipsec-policy=out,none out-interface-list=WAN
add action=dst-nat chain=dstnat comment=JanJansen dst-port=17722 in-interface=bridge log=yes protocol=tcp to-addresses=192.168.88.162 to-ports=22
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
add address=::224.0.0.0/100 comment="defconf: other" list=bad_ipv6
add address=::127.0.0.0/104 comment="defconf: other" list=bad_ipv6
add address=::/104 comment="defconf: other" list=bad_ipv6
add address=::255.0.0.0/104 comment="defconf: other" list=bad_ipv6
/ipv6 firewall filter
add action=accept chain=input comment="defconf: accept established,related,untracked" connection-state=established,related,untracked
add action=drop chain=input comment="defconf: drop invalid" connection-state=invalid
add action=accept chain=input comment="defconf: accept ICMPv6" protocol=icmpv6
add action=accept chain=input comment="defconf: accept UDP traceroute" port=33434-33534 protocol=udp
add action=accept chain=input comment="defconf: accept DHCPv6-Client prefix delegation." dst-port=546 protocol=udp src-address=fe80::/10
add action=accept chain=input comment="defconf: accept IKE" dst-port=500,4500 protocol=udp
add action=accept chain=input comment="defconf: accept ipsec AH" protocol=ipsec-ah
add action=accept chain=input comment="defconf: accept ipsec ESP" protocol=ipsec-esp
add action=accept chain=input comment="defconf: accept all that matches ipsec policy" ipsec-policy=in,ipsec
add action=drop chain=input comment="defconf: drop everything else not coming from LAN" in-interface-list=!LAN
