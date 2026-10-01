---
id: collect-261001-cisco/cisco/questions-6690-configuring-a-cisco-asa-5505-for-routing-0410ef3c
title: "questions-6690-configuring-a-cisco-asa-5505-for-routing-0410ef3c"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: ["grok", "parameters"]
source: docs/RAG/collect-261001-cisco/questions-6690-configuring-a-cisco-asa-5505-for-routing-0410ef3c.md
source_anchor: ""
source_lines: [1, 171]
sha256: a60794e4f49e23ef1248f52c5f4f1725e54fa2dd138bad65bc02c94429b8fd4b
---

# questions-6690-configuring-a-cisco-asa-5505-for-routing-0410ef3c

I'll start by pointig out that networking issues have always left me scratching my head. There is something about routing especially that I just haven't had that "Oh I get it" moment yet, so it's likely this is a very basic misconfiguration.
I am trying to set up a Cisco ASA 5505 to be connected with a public IP address on one interface, and to have the second interface connect to our internal network. Right now our internal network is on 192.168.1.0/24, and the public IP we have from the ISP is in the 125.x.x.x range.
I have set up 3 interfaces on the ASA,
- Inside IP: 192.168.1.3
Mask: 255.255.255.0
Ports 1-6
Sec Level 100
- Outside IP: 125.x.x.x
Mask: 255.0.0.0
Port 0
Sec Level 0
- Management IP: 10.0.0.1
Mask: 255.255.255.0
Port 7
Restricted to management
Restrict flow to Outside (had to do this for licensing reasons)
I have an access rule on the firewall that allows IP traffic from any source on the inside network to any less secure network (there is one for both IPv4 and 6).
With the internet-facing router connected to port 0, I connect a computer to port 1 with IP 192.168.1.20, Mask 255.255.255.0 and gateway 192.168.1.3. Using this PC I can ping 192.168.1.3 but no further - I cannot ping the public IP address or 8.8.8.8.
If I connect the same PC directly to the router and assign it the public IP address directly I can access the internet no problem. I believe that the issue here is that I need to add some static routes to show the path from inside<->outside, but I just do not grok them well enough to know how to structure them. Any help or suggestions would be very much appreciated.
Edit: See below the running config, with some redactions.
: Saved
:
ASA Version 8.2(1) 
!
hostname AUS-FW-01
domain-name =============.com.au
enable password ============= encrypted
passwd ============= encrypted
names
name 10.61.76.0 CNC-network
name 10.61.73.0 LAN-network
name 192.168.1.0 LEGACY-network
name 10.61.75.0 RANDD-network
name 10.61.74.0 WIRELESS-network
name 10.61.73.5 InternalGateway
name 192.168.1.200 DCServer
name 192.168.1.214 MUVS-TP description Trueform Print Server
!
interface Vlan1
 nameif inside
 security-level 100
 ip address 192.168.1.3 255.255.255.0 
!
interface Vlan2
 nameif outside
 security-level 0
 ip address 125.x.x.10 255.255.255.252 
!
interface Vlan12
 no forward interface Vlan2
 nameif Management
 security-level 0
 ip address 10.0.0.1 255.255.255.0 
 management-only
!
interface Ethernet0/0
 switchport access vlan 2
!
interface Ethernet0/1
!
interface Ethernet0/2
!
interface Ethernet0/3
!
interface Ethernet0/4
!
interface Ethernet0/5
!
interface Ethernet0/6
!
interface Ethernet0/7
 switchport access vlan 12
!
ftp mode passive
clock timezone EST 10
dns domain-lookup outside
dns server-group DefaultDNS
 name-server =============
 name-server =============
 domain-name =============.com.au
same-security-traffic permit intra-interface
access-list outside_access_in remark Email access
access-list outside_access_in extended permit tcp any LEGACY-network 255.255.255.0 eq smtp 
access-list outside_access_in remark ActiveSync Email Access
access-list outside_access_in extended permit tcp any LEGACY-network 255.255.255.0 eq https 
access-list outside_access_in remark Pronto Trueform Printing
access-list outside_access_in extended permit tcp any LEGACY-network 255.255.255.0 eq lpd 
pager lines 24
logging asdm informational
mtu inside 1500
mtu outside 1500
mtu Management 1500
icmp unreachable rate-limit 1 burst-size 1
no asdm history enable
arp timeout 14400
global (outside) 1 interface
nat (inside) 1 0.0.0.0 0.0.0.0
static (inside,outside) tcp interface smtp DCServer smtp netmask 255.255.255.255 
static (inside,outside) tcp interface https DCServer https netmask 255.255.255.255 
static (inside,outside) tcp interface lpd MUVS-TP lpd netmask 255.255.255.255 
access-group outside_access_in in interface outside
route inside LAN-network 255.255.255.0 10.61.79.2 1
route inside LEGACY-network 255.255.255.0 InternalGateway 1
timeout xlate 3:00:00
timeout conn 1:00:00 half-closed 0:10:00 udp 0:02:00 icmp 0:00:02
timeout sunrpc 0:10:00 h323 0:05:00 h225 1:00:00 mgcp 0:05:00 mgcp-pat 0:05:00
timeout sip 0:30:00 sip_media 0:02:00 sip-invite 0:03:00 sip-disconnect 0:02:00
timeout sip-provisional-media 0:02:00 uauth 0:05:00 absolute
timeout tcp-proxy-reassembly 0:01:00
dynamic-access-policy-record DfltAccessPolicy
http server enable
http 10.61.79.0 255.255.255.0 inside
http LAN-network 255.255.255.0 inside
http 10.0.0.0 255.255.255.0 Management
no snmp-server location
no snmp-server contact
snmp-server enable traps snmp authentication linkup linkdown coldstart
crypto ipsec security-association lifetime seconds 28800
crypto ipsec security-association lifetime kilobytes 4608000
telnet timeout 5
ssh timeout 5
console timeout 0
vpdn username ============= password ============= store-local
dhcpd auto_config outside
!
threat-detection basic-threat
threat-detection statistics port
threat-detection statistics protocol
threat-detection statistics access-list
no threat-detection statistics tcp-intercept
webvpn
!
class-map inspection_default
 match default-inspection-traffic
!
!
policy-map type inspect dns preset_dns_map
 parameters
  message-length maximum 512
policy-map global_policy
 class inspection_default
  inspect dns preset_dns_map 
  inspect ftp 
  inspect h323 h225 
  inspect h323 ras 
  inspect rsh 
  inspect rtsp 
  inspect esmtp 
  inspect sqlnet 
  inspect skinny  
  inspect sunrpc 
  inspect xdmcp 
  inspect sip  
  inspect netbios 
  inspect tftp 
!
service-policy global_policy global
prompt hostname context 
Cryptochecksum:=============
: end
asdm location LAN-network 255.255.255.0 inside
asdm location WIRELESS-network 255.255.255.0 inside
asdm location RANDD-network 255.255.255.0 inside
asdm location CNC-network 255.255.255.0 inside
asdm location LEGACY-network 255.255.255.0 inside
asdm location InternalGateway 255.255.255.255 inside
asdm location DCServer 255.255.255.255 inside
asdm location MUVS-TP 255.255.255.255 inside
no asdm history enable
PS: I wasn't sure if this was a better fit here or on SF. I have no problem with it being moved.
125.x.x.9. When she tested it and showed me it was working for me to signoff she did so by connecting her laptop to it and setting her IP to125.x.x.10, her mask to255.0.0.0and her default gateway to125.x.x.9. I have made the assumption those were the correct values based on her using them. Is it unlikely for a public IP to have such a wide network?
