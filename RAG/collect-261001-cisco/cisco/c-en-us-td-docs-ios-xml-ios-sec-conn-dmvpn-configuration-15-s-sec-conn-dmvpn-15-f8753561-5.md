---
id: collect-261001-cisco/cisco/c-en-us-td-docs-ios-xml-ios-sec-conn-dmvpn-configuration-15-s-sec-conn-dmvpn-15-f8753561-5
title: "c-en-us-td-docs-ios-xml-ios-sec-conn-dmvpn-configuration-15-s-sec-conn-dmvpn-15--f8753561"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: ["ethernet"]
source: docs/RAG/collect-261001-cisco/c-en-us-td-docs-ios-xml-ios-sec-conn-dmvpn-configuration-15-s-sec-conn-dmvpn-15--f8753561.md
source_anchor: ""
source_lines: [794, 950]
sha256: 64f7ca29ae3b709f3f97ad41296a7273a715cd93d3c721365a2ae0536604433f
---

# c-en-us-td-docs-ios-xml-ios-sec-conn-dmvpn-configuration-15-s-sec-conn-dmvpn-15--f8753561

			  
				debug 
				nhrp 
			  [cache | 
			 extension | 
			 packet | 
			 rate] 
		  
6.   
       
			  
				debug 
				nhrp 
				condition 
			 [interface 
			 tunnel 
			 number | 
			 peer {nbma {ipv4-address | 
			 fqdn-string | 
			 ipv6-address} | 
			 tunnel {ip-address | 
			 ipv6-address}} | 
			 vrf 
			 vrf-name] 
		  
7.   
       
			  
				debug 
				nhrp 
			  
			  
				error 
			  
		  
DETAILED STEPS
|  | Command or Action | Purpose | 
|---|---|---|
| Step 1 | enable Example: Device> enable | Enables privileged EXEC mode. | 
| Step 2 | clear  				dmvpn  				session  			  [interface  				tunnel  			 number \|  			 peer {ipv4-address \|  			 fqdn-string \|  			 ipv6-address} \|  			 vrf  			 vrf-name] [static] Example: Device# clear dmvpn session | Clears DMVPN sessions. | 
| Step 3 | clear  				ipv6  				nhrp  			  [ipv6-address \|  			 counters Example: Device# clear ipv6 nhrp | Clears all dynamic entries from the NHRP cache. | 
| Step 4 | debug  				dmvpn  			  {all \|  			 error \|  			 detail \|  			 packet} {all \|  			 debug-type} Example: Device# debug dmvpn | Displays debug DMVPN session information. | 
| Step 5 | debug  				nhrp  			  [cache \|  			 extension \|  			 packet \|  			 rate] Example: Device# debug nhrp ipv6 | Enables NHRP debugging. | 
| Step 6 | debug  				nhrp  				condition  			 [interface  			 tunnel  			 number \|  			 peer {nbma {ipv4-address \|  			 fqdn-string \|  			 ipv6-address} \|  			 tunnel {ip-address \|  			 ipv6-address}} \|  			 vrf  			 vrf-name] Example: Device# debug nhrp condition | Enables NHRP conditional debugging. | 
| Step 7 | debug  				nhrp  			   			   				error Example: Device# debug nhrp ipv6 error | Displays NHRP error-level debugging information. | 
Examples
The following sample output is from the debug nhrpcommand with the ipv6 keyword:
Device# debug nhrp ipv6
Aug  9 13:13:41.486: NHRP: Attempting to send packet via DEST
			- 2001:DB8:3c4d:0015:0000:0000:1a2f:3d2c/32
Aug  9 13:13:41.486: NHRP: Encapsulation succeeded.  
Aug  9 13:13:41.486: NHRP: Tunnel NBMA addr 11.11.11.99
Aug  9 13:13:41.486: NHRP: Send Registration Request via Tunnel0 vrf 0, packet size: 105
Aug  9 13:13:41.486: src: 2001:DB8:3c4d:0015:0000:0000:1a2f:3d2c/32, 
	         dst: 2001:DB8:3c4d:0015:0000:0000:1a2f:3d2c/32
Aug  9 13:13:41.486: NHRP: 105 bytes out Tunnel0
Aug  9 13:13:41.486: NHRP: Receive Registration Reply via Tunnel0 vrf 0, packet size: 125
Configuration Examples for IPv6 over DMVPN
Example: Configuring an IPsec Profile
Device(config)# crypto identity router1
 
		
Device(config)# crypto ipsec profile example1
Device(config-crypto-map)# set transform-set example-set
Device(config-crypto-map)# set identity router1
Device(config-crypto-map)# set security-association lifetime seconds 1800 
 
		Device(config-crypto-map)# set pfs group14 
 
	 Example: Configuring the Hub for DMVPN
Device# configure terminal
Device(config)# interface tunnel 5
 
Device(config-if)# ipv6 address 2001:DB8:1:1::72/64
Device(config-if)# ipv6 address fe80::2001 link-local
Device(config-if)# ipv6 mtu 1400 
Device(config-if)# ipv6 nhrp authentication examplexx
Device(config-if)# ipv6 nhrp map multicast dynamic
Device(config-if)# ipv6 nhrp network-id 99
Device(config-if)# tunnel source ethernet 0
Device(config-if)# tunnel mode gre multipoint
Device(config-if)# tunnel protection ipsec profile example_profile
Device(config-if)# bandwidth 1200
Device(config-if)# ipv6 nhrp holdtime 3600
The following sample output is from the show dmvpn command, with the ipv6 and detail keywords, for the hub:
Device# show dmvpn ipv6 detail
Legend: Attrb --> S - Static, D - Dynamic, I - Incomplete
        N - NATed, L - Local, X - No Socket
        # Ent --> Number of NHRP entries with same NBMA peer
        NHS Status: E --> Expecting Replies, R --> Responding
        UpDn Time --> Up or Down Time for a Tunnel
==========================================================================
Interface Tunnel1 is up/up, Addr. is 10.0.0.3, VRF "" 
   Tunnel Src./Dest. addr: 192.169.2.9/MGRE, Tunnel VRF ""
   Protocol/Transport: "multi-GRE/IP", Protect "test_profile" 
Type:Hub, Total NBMA Peers (v4/v6): 2
    1.Peer NBMA Address: 192.169.2.10
        Tunnel IPv6 Address: 2001::4
        IPv6 Target Network: 2001::4/128
        # Ent: 2, Status: UP, UpDn Time: 00:01:51, Cache Attrib: D
Type:Hub, Total NBMA Peers (v4/v6): 2
    2.Peer NBMA Address: 192.169.2.10
        Tunnel IPv6 Address: 2001::4
        IPv6 Target Network: FE80::2/128
        # Ent: 0, Status: UP, UpDn Time: 00:01:51, Cache Attrib: D
Type:Hub, Total NBMA Peers (v4/v6): 2
    3.Peer NBMA Address: 192.169.2.11
Tunnel IPv6 Address: 2001::5
        IPv6 Target Network: 2001::5/128
        # Ent: 2, Status: UP, UpDn Time: 00:26:38, Cache Attrib: D
Type:Hub, Total NBMA Peers (v4/v6): 2
    4.Peer NBMA Address: 192.169.2.11
        Tunnel IPv6 Address: 2001::5
        IPv6 Target Network: FE80::3/128
        # Ent: 0, Status: UP, UpDn Time: 00:26:38, Cache Attrib: D
Pending DMVPN Sessions:
Interface: Tunnel1
  IKE SA: local 192.169.2.9/500 remote 192.169.2.10/500 Active 
  Crypto Session Status: UP-ACTIVE     
  fvrf: (none), Phase1_id: 192.169.2.10
  IPSEC FLOW: permit 47 host 192.169.2.9 host 192.169.2.10 
        Active SAs: 2, origin: crypto map
   Outbound SPI : 0x BB0ED02, transform : esp-aes esp-sha-hmac 
    Socket State: Open
Interface: Tunnel1
  IKE SA: local 192.169.2.9/500 remote 192.169.2.11/500 Active 
  Crypto Session Status: UP-ACTIVE     
  fvrf: (none), Phase1_id: 192.169.2.11
  IPSEC FLOW: permit 47 host 192.169.2.9 host 192.169.2.11 
        Active SAs: 2, origin: crypto map
   Outbound SPI : 0xB79B277B, transform : esp-aes esp-sha-hmac 
    Socket State: Open
 
		
Example: Configuring the Spoke for DMVPN
Device# configure terminal
Device(config)# crypto ikev2 keyring DMVPN
Device(config)# peer DMVPN
Device(config)# address 0.0.0.0 0.0.0.0
Device(config)# pre-shared-key cisco123
Device(config)# peer DMVPNv6
Device(config)# address ::/0
Device(config)# pre-shared-key cisco123v6
Device(config)# crypto ikev2 profile DMVPN
Device(config)# match identity remote address 0.0.0.0
Device(config)# match identity remote address ::/0
Device(config)# authentication local pre-share
Device(config)# authentication remote pre-share
Device(config)# keyring DMVPN
Device(config)# dpd 30 5 on-demand
Device(config)# crypto ipsec transform-set DMVPN esp-aes esp-sha-hmac
Device(config)# mode transport
Device(config)# crypto ipsec profile DMVPN
Device(config)# set transform-set DMVPN
Device(config)# set ikev2-profile DMVPN
Device(config)# interface tunnel 5
 
