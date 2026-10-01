---
id: collect-261001-fortinet/fortinet/docs-fortinet-com-v2-attachments-5aa37c8a-1a11-11e9-9685-f8bc1258b856-fortigate-ba9ee001-7
title: "docs-fortinet-com-v2-attachments-5aa37c8a-1a11-11e9-9685-f8bc1258b856-fortigate--ba9ee001"
domain: fortinet
role: reference
task: reference
actors: ["Apple"]
dates: []
keywords: []
source: docs/RAG/collect-261001-fortinet/docs-fortinet-com-v2-attachments-5aa37c8a-1a11-11e9-9685-f8bc1258b856-fortigate--ba9ee001.md
source_anchor: ""
source_lines: [1005, 1191]
sha256: 7ed0be05fba3e299a7ce3bcbc8df01b88b11f46962e2502b8078d98fea11c151
---

# docs-fortinet-com-v2-attachments-5aa37c8a-1a11-11e9-9685-f8bc1258b856-fortigate--ba9ee001

Example 1: Remote Sites with Different Subnets IPsec VPN in Transparent Mode
Example 1: Remote Sites with Different Subnets
This example provides a configuration example for IPsec VPN tunnels between three FortiGate in Transparent
Mode in different subnets, as well as some troubleshooting steps.
The expectation for this example is that PC1 will be able to communicate via the IPsec tunnels with PC2 and
PC3, which are in different subnets.
The requirements for this example are:
l Because both FortiGate are not in the same broadcast domain (separated by routers), the hosts on each side must
be on different subnets.
l FortiGate management IP addresses must be in the same subnet as the local hosts
l The default gateways (router1 ,router2, router3) for PC1 , PC2 and PC3 must be behind port2 in order for the
FortiGate to match the appropriate Encrypt firewall policy (port1 --> port2)
Configuration of FortiGate 1 (FGT1):
Only relevant parts of configuration are provided.
config system settings
set opmode transparent
set manageip 10.1.1.100/255.255.255.0
end
config router static
edit 1
32 Transparent Mode for FortiOS 5.6.3
Fortinet Technologies Inc.

IPsec VPN in Transparent Mode Example 1: Remote Sites with Different Subnets
set gateway 10.1.1.254
next
end
config firewall address
edit "10.1.1.0/24"
set subnet 10.1.1.0 255.255.255.0
next
edit "10.2.2.0/24"
set subnet 10.2.2.0 255.255.255.0
next
edit "10.3.3.0/24"
set subnet 10.3.3.0 255.255.255.0
next
end
config vpn ipsec phase1
edit "to_FGT2"
set proposal 3des-sha1 aes128-sha1 des-md5
set remote-gw 10.2.2.100
set psksecret fortinet
next
edit "to_FGT3"
set proposal 3des-sha1 aes128-sha1 des-md5
set remote-gw 10.3.3.100
set psksecret fortinet
next
end
config vpn ipsec phase2
edit "to_FGT2"
set keepalive enable
set phase1name "to_FGT2"
set proposal 3des-sha1 aes128-sha1
set dst-subnet 10.2.2.0 255.255.255.0
set src-subnet 10.1.1.0 255.255.255.0
next
edit "to_FGT3"
set keepalive enable
set phase1name "to_FGT3"
set proposal 3des-sha1 aes128-sha1
set dst-subnet 10.3.3.0 255.255.255.0
set src-subnet 10.1.1.0 255.255.255.0
next
end
config firewall policy
edit 1
set srcintf "port1"
set dstintf "port2"
set srcaddr "10.1.1.0/24"
set dstaddr "10.2.2.0/24"
set action ipsec
set schedule "always"
set service "ANY"
set inbound enable
set outbound enable
Transparent Mode for FortiOS 5.6.3
Fortinet Technologies Inc.
33

Example 1: Remote Sites with Different Subnets IPsec VPN in Transparent Mode
set vpntunnel "to_FGT2"
next
edit 3
set srcintf "port1"
set dstintf "port2"
set srcaddr "10.1.1.0/24"
set dstaddr "10.3.3.0/24"
set action ipsec
set schedule "always"
set service "ANY"
set inbound enable
set outbound enable
set vpntunnel "to_FGT3"
next
end
Configuration of FortiGate 2 (FGT2):
Only relevant parts of configuration are provided.
config system settings
set opmode transparent
set manageip 10.2.2.100/255.255.255.0
end
config router static
edit 1
set gateway 10.2.2.254
next
end
config firewall address
edit "10.1.1.0/24"
set subnet 10.1.1.0 255.255.255.0
next
edit "10.2.2.0/24"
set subnet 10.2.2.0 255.255.255.0
next
end
config vpn ipsec phase1
edit "to_FGT1"
set nattraversal disable
set proposal 3des-sha1 aes128-sha1 des-md5
set remote-gw 10.1.1.100
set psksecret fortinet
next
end
config vpn ipsec phase2
edit "to_FGT1"
set keepalive enable
set phase1name "to_FGT1"
set proposal 3des-sha1 aes128-sha1
set dst-subnet 10.1.1.0 255.255.255.0
set src-subnet 10.2.2.0 255.255.255.0
next
34 Transparent Mode for FortiOS 5.6.3
Fortinet Technologies Inc.

IPsec VPN in Transparent Mode Example 1: Remote Sites with Different Subnets
end
config firewall policy
edit 1
set srcintf "port1"
set dstintf "port2"
set srcaddr "10.2.2.0/24"
set dstaddr "10.1.1.0/24"
set action ipsec
set schedule "always"
set service "ANY"
set inbound enable
set outbound enable
set vpntunnel "to_FGT1"
next
end
Troubleshooting procedure
All steps given when PC1 pings PC2.
Verify if IPsec tunnels are up
FGT1 # diagnose vpn tunnel list
list all ipsec tunnel in vd 0
------------------------------------------------------
name=to_FGT2 ver=0 serial=1 10.1.1.100:0->10.2.2.100:0 lgwy=dyn tun=tunnel mode= auto
bound_if=0
proxyid_num=1 child_num=0 refcnt=7 ilast=0 olast=0
stat: rxp=0 txp=0 rxb=0 txb=0
dpd: mode=active on=1 idle=5000ms retry=3 count=0 seqno=1455
natt: mode=none draft=0 interval=0 remote_port=0
proxyid=to_FGT2 proto=0 sa=0 ref=1 auto_negotiate=0 serial=5
src: 10.1.1.0/255.255.255.0:0
dst: 10.2.2.0/255.255.255.0:0
The above tunnel is down (output given as example)!
FGT2 # diagnose vpn tunnel list
list all ipsec tunnel in vd 0
------------------------------------------------------
name=to_FGT1 10.2.2.100:0->10.1.1.100:0 lgwy=dyn tun=tunnel mode=auto bound_if=0
proxyid_num=1 child_num=0 refcnt=7 ilast=1 olast=1
stat: rxp=0 txp=0 rxb=0 txb=0
dpd: mode=active on=1 idle=5000ms retry=3 count=0 seqno=21 natt:
mode=none draft=0 interval=0 remote_port=0
proxyid=to_FGT1 proto=0 sa=1 ref=2 auto_negotiate=0 serial=1
src: 10.2.2.0/255.255.255.0:0
dst: 10.1.1.0/255.255.255.0:0
SA: ref=3 options=00000009 type=00 soft=0 mtu=1436 expire=1771 replaywin=0 seq no=1
life: type=01 bytes=0/0 timeout=1773/1800
dec: spi=c1a8e951 esp=3des key=24 9213fdf22b150e01abb3535d1a647044eebf772b92f2f7ee
ah=sha1 key=20 66a38bf99f0b2d234f64b5a05187995c4f56f6bb
enc: spi=322067b4 esp=3des key=24 720e5680329937fb3630b7ed70bd41bb3114d3c269ae8b61
Transparent Mode for FortiOS 5.6.3
Fortinet Technologies Inc.
35

