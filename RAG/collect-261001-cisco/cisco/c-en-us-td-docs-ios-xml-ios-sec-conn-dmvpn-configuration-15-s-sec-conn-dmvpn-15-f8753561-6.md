---
id: collect-261001-cisco/cisco/c-en-us-td-docs-ios-xml-ios-sec-conn-dmvpn-configuration-15-s-sec-conn-dmvpn-15-f8753561-6
title: "c-en-us-td-docs-ios-xml-ios-sec-conn-dmvpn-configuration-15-s-sec-conn-dmvpn-15--f8753561"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-cisco/c-en-us-td-docs-ios-xml-ios-sec-conn-dmvpn-configuration-15-s-sec-conn-dmvpn-15--f8753561.md
source_anchor: ""
source_lines: [951, 1072]
sha256: bbfff46eef25b9e9cef67afe879e0ee29abf128df49eccea756762f53dc2f290
---

# c-en-us-td-docs-ios-xml-ios-sec-conn-dmvpn-configuration-15-s-sec-conn-dmvpn-15--f8753561

Device(config-if)# bandwidth 1000
Device(config-if)# ip address 10.0.0.11 255.255.255.0
Device(config-if)# ip mtu 1400
Device(config-if)# ip nhrp authentication test 
Device(config-if)# ip nhrp network-id 100000
Device(config-if)# ip nhrp nhs 10.0.0.1 nbma 2001:DB8:0:FFFF:1::1 multicast
Device(config-if)# vip nhrp shortcut
Device(config-if)# delay 1000
Device(config-if)# ipv6 address 2001:DB8:0:100::B/64
Device(config-if)# ipv6 mtu 1400
Device(config-if)# ipv6 nd ra mtu suppress
Device(config-if)# no ipv6 redirects
Device(config-if)# ipv6 eigrp 1
Device(config-if)# ipv6 nhrp authentication testv6
Device(config-if)# ipv6 nhrp network-id 100006
Device(config-if)# ipv6 nhrp nhs 2001:DB8:0:100::1 nbma 2001:DB8:0:FFFF:1::1 multicast
Device(config-if)# ipv6 nhrp shortcut
Device(config-if)# tunnel source Ethernet0/0
Device(config-if)# tunnel mode gre multipoint ipv6
Device(config-if)# tunnel key 100000
Device(config-if)# end
.
.
The following sample output is from the show dmvpn command, with the ipv6 and detail keywords, for the spoke:
Legend: Attrb --> S - Static, D - Dynamic, I - Incomplete
        N - NATed, L - Local, X - No Socket
        # Ent --> Number of NHRP entries with same NBMA peer
        NHS Status: E --> Expecting Replies, R --> Responding
        UpDn Time --> Up or Down Time for a Tunnel
==========================================================================
Interface Tunnel1 is up/up, Addr. is 10.0.0.1, VRF "" 
   Tunnel Src./Dest. addr: 192.169.2.10/MGRE, Tunnel VRF ""
   Protocol/Transport: "multi-GRE/IP", Protect "test_profile" 
IPv6 NHS: 2001::6 RE
Type:Spoke, Total NBMA Peers (v4/v6): 1
    1.Peer NBMA Address: 192.169.2.9
        Tunnel IPv6 Address: 2001::6
        IPv6 Target Network: 2001::/112
        # Ent: 2, Status: NHRP, UpDn Time: never, Cache Attrib: S
IPv6 NHS: 2001::6 RE
Type:Unknown, Total NBMA Peers (v4/v6): 1
    2.Peer NBMA Address: 192.169.2.9
        Tunnel IPv6 Address: FE80::1
        IPv6 Target Network: FE80::1/128
        # Ent: 0, Status: UP, UpDn Time: 00:00:24, Cache Attrib: D
Pending DMVPN Sessions:
Interface: Tunnel1
  IKE SA: local 192.169.2.10/500 remote 192.169.2.9/500 Active 
  Crypto Session Status: UP-ACTIVE     
  fvrf: (none), Phase1_id: 192.169.2.9
  IPSEC FLOW: permit 47 host 192.169.2.10 host 192.169.2.9 
        Active SAs: 2, origin: crypto map
   Outbound SPI : 0x6F75C431, transform : esp-aes esp-sha-hmac 
    Socket State: Open
 
	 Example: Configuring the NHRP Redirect and Shortcut Features on the Hub
Device(config)# interface tunnel 5
Device(config-if)# ipv6 address 2001:DB8:1:1::72/64
Device(config-if)# ipv6 nhrp redirect
 
		Device(config-if)# ipv6 nhrp shortcut
 
	 Example: Configuring NHRP on the Hub and Spoke
Hub
Device# show ipv6 nhrp
2001::4/128 via 2001::4
   Tunnel1 created 00:02:40, expire 00:00:47
   Type: dynamic, Flags: unique registered used 
   NBMA address: 192.169.2.10 
2001::5/128 via 2001::5
   Tunnel1 created 00:02:37, expire 00:00:47
   Type: dynamic, Flags: unique registered used 
   NBMA address: 192.169.2.11 
FE80::2/128 via 2001::4
   Tunnel1 created 00:02:40, expire 00:00:47
   Type: dynamic, Flags: unique registered used 
   NBMA address: 192.169.2.10 
FE80::3/128 via 2001::5
   Tunnel1 created 00:02:37, expire 00:00:47
   Type: dynamic, Flags: unique registered used 
   NBMA address: 192.169.2.11 
Spoke
Device# show ipv6 nhrp
2001::8/128
   Tunnel1 created 00:00:13, expire 00:02:51
   Type: incomplete, Flags: negative 
   Cache hits: 2
2001::/112 via 2001::6
   Tunnel1 created 00:01:16, never expire 
   Type: static, Flags: used 
   NBMA address: 192.169.2.9
FE80::1/128 via FE80::1
   Tunnel1 created 00:01:15, expire 00:00:43
   Type: dynamic, Flags: 
   NBMA address: 192.169.2.9 
Additional References
Related Documents
| Related Topic | Document Title | 
|---|---|
| IPv6 addressing and connectivity | IPv6 Configuration Guide | 
| Dynamic Multipoint VPN | Dynamic Multipoint VPN Configuration Guide | 
| Cisco IOS commands | Master Command List, All Releases | 
| IPv6 commands | IPv6 Command Reference | 
| Cisco IOS IPv6 features | IPv6 Feature Mapping | 
| Recommended cryptographic algorithms | Next Generation Encryption | 
Standards and RFCs
| Standard/RFC | Title | 
|---|---|
| RFCs for IPv6 | IPv6 RFcs | 
Technical Assistance
| Description | Link | 
|---|---|
| The Cisco Support and Documentation website provides online resources to download documentation, software, and tools. Use these resources to install and configure the software and to troubleshoot and resolve technical issues with Cisco products and technologies. Access to most tools on the Cisco Support and Documentation website requires a Cisco.com user ID and password. | http://www.cisco.com/cisco/web/support/index.html | 
Feature Information for IPv6 over DMVPN
The following table provides release information about the feature or features described in this module. This table lists only the software release that introduced support for a given feature in a given software release train. Unless noted otherwise, subsequent releases of that software release train also support that feature.
Use Cisco Feature Navigator to find information about platform support and Cisco software image support. To access Cisco Feature Navigator, go to www.cisco.com/go/cfn. An account on Cisco.com is not required.
| Table 1 Feature Information for IPv6 over DMVPN |  |  | 
|---|---|---|
| Feature Name | Releases | Feature Information | 
|---|---|---|
| IPv6 Transport for DMVPN | 15.3(1)S | The IPv6 transport for DMVPN feature builds IPv6 WAN-side capability into NHRP tunnels and the underlying IPsec encryption, and enables IPv6 to transport payloads on the Internet. The IPv6 transport for DMVPN feature is enabled by default. No new commands were introduced or modifed. | 
Back to Top
