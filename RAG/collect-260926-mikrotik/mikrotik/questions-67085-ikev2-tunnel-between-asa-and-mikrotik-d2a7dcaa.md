---
id: collect-260926-mikrotik/mikrotik/questions-67085-ikev2-tunnel-between-asa-and-mikrotik-d2a7dcaa
title: "questions-67085-ikev2-tunnel-between-asa-and-mikrotik-d2a7dcaa"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/lot-mikrotik/forum/ipsec/questions-67085-ikev2-tunnel-between-asa-and-mikrotik-d2a7dcaa.md
source_anchor: ""
source_lines: [1, 206]
sha256: 3c6d2ea25595b54458b0b0486f35ba4dd3db90d637607859d866e4702a2515c5
---

# questions-67085-ikev2-tunnel-between-asa-and-mikrotik-d2a7dcaa

Trying to move from pfSense to Mikrotik for an office router, and the only stumbling block is maintaining a site-to-site IPSEC tunnel between it and our Cisco ASA. The settings all look correct to me, and the tunnels show up on both sides (see note below) but no traffic passes between networks.
The only suspicious thing I can find is this message in the Cisco logs:
Apr  7 13:08:35 asa1.pofp.internal %ASA-4-750003: Local:9.8.7.6:500 Remote:2.3.4.5:500 Username:Unknown 
IKEv2 Negotiation aborted due to ERROR: Failed to receive the AUTH msg before the timer expired
There is no NAT involved here, and no firewalls between these devices.
I'm hesitant to mention it for fear of muddying the waters, but the tunnel has worked twice during my testing, with this same configuration. The success seemed to happen randomly while manually tearing down tunnels on both devices and may be related to the timing of the initiation on both sides? However, in both cases the tunnels stopped passing traffic after the P2 timeout.
Mikrotik config:
/ip ipsec profile
add dh-group=ecp521 dpd-interval=10s enc-algorithm=aes-256 hash-algorithm=sha512 name=asa-p1 nat-traversal=no
/ip ipsec peer
add address=9.8.7.6/32 exchange-mode=ike2 name=NOC port=500 profile=asa-p1 send-initial-contact=no
/ip ipsec proposal
set [ find default=yes ] disabled=yes
add auth-algorithms=sha512 enc-algorithms=aes-256-gcm lifetime=8h name=asa-p2 pfs-group=ecp521
/ip ipsec identity
add peer=NOC secret="*****"
/ip ipsec policy
set 0 disabled=yes
add dst-address=192.168.242.0/24 proposal=asa-p2 sa-dst-address=9.8.7.6 sa-src-address=0.0.0.0 src-address=192.168.243.0/24 tunnel=yes
/ip firewall nat
add chain=srcnat dst-address=192.168.242.0/24 src-address=192.168.243.0/24
add chain=srcnat dst-address=192.168.243.0/24 src-address=192.168.242.0/24
add action=masquerade chain=srcnat out-interface="WAN port"
/ip firewall filter
add action=accept chain=input comment="Allow established input traffic" connection-state=established,related
add action=accept chain=input comment=IPSEC dst-port=500 in-interface="WAN port" protocol=udp
add action=accept chain=input comment="IPSEC NAT-T" dst-port=4500 in-interface="WAN port" protocol=udp
add action=accept chain=input comment="IPSEC ESP" in-interface="WAN port" protocol=ipsec-esp
...
Cisco config
object network NOC-network
 subnet 192.168.242.0 255.255.255.0
object network Calgary-network
 subnet 192.168.243.0 255.255.255.0
crypto ipsec ikev2 ipsec-proposal AESGCM
 protocol esp encryption aes-gcm-256
 protocol esp integrity sha-512
crypto ipsec ikev2 sa-strength-enforcement
crypto ipsec security-association pmtu-aging infinite
crypto ikev2 policy 2
 encryption aes-gcm-256
 integrity null
 group 21 24
 prf sha512
 lifetime seconds 86400
crypto ikev2 policy 3
 encryption aes-256
 integrity sha512
 group 21 24
 prf sha512
 lifetime seconds 86400
crypto ikev2 enable OUTSIDE
group-policy GroupPolicy_IKEv2 internal
group-policy GroupPolicy_IKEv2 attributes
 vpn-tunnel-protocol ikev2 
tunnel-group 2.3.4.5 type ipsec-l2l
tunnel-group 2.3.4.5 general-attributes
 default-group-policy GroupPolicy_IKEv2
tunnel-group 2.3.4.5 ipsec-attributes
 ikev2 remote-authentication pre-shared-key *****
 ikev2 local-authentication pre-shared-key *****
access-list OUTSIDE_cryptomap_1 extended permit ip object NOC-network object Calgary-network 
nat (INSIDE,OUTSIDE) source static NOC-network NOC-network destination static Calgary-network Calgary-network no-proxy-arp route-lookup
crypto map OUTSIDE_map 2 match address OUTSIDE_cryptomap_1
crypto map OUTSIDE_map 2 set pfs group21
crypto map OUTSIDE_map 2 set peer 2.3.4.5 
crypto map OUTSIDE_map 2 set ikev2 ipsec-proposal AESGCM
crypto map OUTSIDE_map 2 set security-association lifetime kilobytes unlimited
crypto map OUTSIDE_map 2 set nat-t-disable
crypto map OUTSIDE_map interface OUTSIDE
For some reason the ASA shows two bi-directional tunnels up:
While the Mikrotik only sees one (it shows each direction as a separate entry, unlike the ASA.)
> ip ipsec installed-sa print 
Flags: H - hw-aead, A - AH, E - ESP 
 0  E spi=0x6FFE0E4 src-address=9.8.7.6 dst-address=2.3.4.5 state=mature enc-algorithm=aes-gcm enc-key-size=288 
      enc-key="2a217b491be5a5297a8a78759e940bc4677b59834630282a2a24baaf3198c6539cc435b0" add-lifetime=6h24m8s/8h10s replay=128 
 1  E spi=0xF315FE3C src-address=2.3.4.5 dst-address=9.8.7.6 state=mature enc-algorithm=aes-gcm enc-key-size=288 
      enc-key="405b00868a64c35521ccfa6feac97316d19220bb4b7b3346964ad0dd0415a54d3ccda8ca" add-lifetime=6h24m8s/8h10s replay=128 
Packet tracer output:
CORP-ASA1# packet-tracer input INSIDE tcp 192.168.242.100 1234 192.168.243.100$
Phase: 1
Type: ACCESS-LIST
Subtype: 
Result: ALLOW
Config:
Implicit Rule
Additional Information:
MAC Access list
Phase: 2
Type: ROUTE-LOOKUP
Subtype: Resolve Egress Interface
Result: ALLOW
Config:
Additional Information:
found next-hop 9.8.7.6 using egress ifc  OUTSIDE
Phase: 3
Type: UN-NAT
Subtype: static
Result: ALLOW
Config:
nat (INSIDE,OUTSIDE) source static NOC-network NOC-network destination static Calgary-network Calgary-network no-proxy-arp route-lookup
Additional Information:
NAT divert to egress interface OUTSIDE
Untranslate 192.168.243.100/1234 to 192.168.243.100/1234
Phase: 4
Type: ACCESS-LIST
Subtype: log
Result: ALLOW
Config:
access-group INSIDE_access_in in interface INSIDE
access-list INSIDE_access_in extended permit ip object-group DM_INLINE_NETWORK_4 any 
object-group network DM_INLINE_NETWORK_4
 network-object aaa:bbb:ccc:242::/64
 network-object 192.168.242.0 255.255.255.0
Additional Information:
Phase: 5
Type: CONN-SETTINGS
Subtype: 
Result: ALLOW
Config:
class-map class-default
 match any
policy-map global_policy
 class class-default
  set connection decrement-ttl
service-policy global_policy global
Additional Information:
Phase: 6
Type: NAT
Subtype: 
Result: ALLOW
Config:
nat (INSIDE,OUTSIDE) source static NOC-network NOC-network destination static Calgary-network Calgary-network no-proxy-arp route-lookup
Additional Information:
Static translate 192.168.242.100/1234 to 192.168.242.100/1234
Phase: 7
Type: NAT
Subtype: per-session
Result: ALLOW
Config:
Additional Information:
Phase: 8
Type: IP-OPTIONS
Subtype: 
Result: ALLOW
Config:       
Additional Information:
Phase: 9
Type: FOVER
Subtype: standby-update
Result: ALLOW
Config:
Additional Information:
Phase: 10
Type: FLOW-EXPORT
Subtype: 
Result: ALLOW
Config:
Additional Information:
Phase: 11
Type: VPN
Subtype: encrypt
Result: ALLOW
Config:
Additional Information:
Phase: 12     
Type: NAT
Subtype: rpf-check
Result: ALLOW
Config:
nat (INSIDE,OUTSIDE) source static NOC-network NOC-network destination static Calgary-network Calgary-network no-proxy-arp route-lookup
Additional Information:
Phase: 13
Type: VPN
Subtype: ipsec-tunnel-flow
Result: ALLOW
Config:
Additional Information:
Phase: 14
Type: NAT
Subtype: per-session
Result: ALLOW
Config:
Additional Information:
Phase: 15
Type: IP-OPTIONS
Subtype:      
Result: ALLOW
Config:
Additional Information:
Phase: 16
Type: FLOW-CREATION
Subtype: 
Result: ALLOW
Config:
Additional Information:
New flow created with id 2426669873, packet dispatched to next module
Result:
input-interface: INSIDE
input-status: up
input-line-status: up
output-interface: OUTSIDE
output-status: up
output-line-status: up
Action: allow
packet-tracer input <192.168.242.0 interface> tcp 192.168.242.100 1234 192.168.243.100 1234and add the output to your original post (with proper formatting).
