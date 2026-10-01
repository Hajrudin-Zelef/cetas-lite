---
id: collect-261001-cisco/cisco/enterprise-en-doc-edoc1100246296-3f27e589-user-access-and-authentication-c126a3e3-34
title: "Configure DeviceA to generate a local key pair."
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-cisco/enterprise-en-doc-edoc1100246296-3f27e589-user-access-and-authentication-c126a3e3.md
source_anchor: ""
source_lines: [6100, 6312]
sha256: 90e369c991a432ed01b88cc76848e536a4cb29f03983bcaf0cef306f1c89e849
---

# Configure the source address for communicating with the RADIUS or Portal server. When both M-LAG and NAC are configured, it is recommended that DeviceA and DeviceB have the same source address configured.
[DeviceA] url-template name url1 
[DeviceA-url-template-url1] url http://192.168.10.1:19008/portal 
[DeviceA-url-template-url1] url-parameter set device-ip 192.168.1.1
[DeviceA-url-template-url1] url-parameter device-ip ac-ip user-ipaddress uaddress user-mac umac 
[DeviceA-url-template-url1] quit
[DeviceA] web-auth-server server-source all-interface
[DeviceA] web-auth-server abc
[DeviceA-web-auth-server-abc] server-ip 192.168.10.1
[DeviceA-web-auth-server-abc] source-ip 192.168.1.1
[DeviceA-web-auth-server-abc] shared-key cipher YsHsjx_202206
[DeviceA-web-auth-server-abc] port 50200
[DeviceA-web-auth-server-abc] url-template url1
[DeviceA-web-auth-server-abc] quit
#
sysname DeviceA
#
dfs-group 1
 priority 150
 dual-active detection source ip 10.1.1.1 peer 10.1.1.2
 authentication-mode hmac-sha256 password %+%##!!!!!!!!!"!!!!"!!!!*!!!!C+tR0CW9x*eB&pWp`t),Azgw-h\o8#4LZPD!!!!!!!!!!!!!!!9!!!!>fwJ)I0E{=:%,*,XRhbH&t0MCy_8=7!!!!!!!!!!%+%#
 dfs-group state switchover disable
#
vlan batch 10 to 11
#
stp mode rstp 
stp v-stp enable
#
access-user m-lag enable
#
radius-server template rd1
 radius-server shared-key cipher %+%##!!!!!!!!!"!!!!"!!!!*!!!!Cd/`W03KjAwAqn64E<\TxGC_SOri<2BP\A+!!!!!2jp5!!!!!!B!!!!Oe2HMc->XMa#TLDZUaJBFJtm#XVj*E:S*|(N7`J1B:3QY!!!!!!!!!!!!!!!%+%# 
 radius-server authentication 192.168.10.1 1812 source ip-address 192.168.1.1 weight 80
 radius-server accounting 192.168.10.1 1813 source ip-address 192.168.1.1 weight 80
#
aaa
 authentication-scheme abc
  authentication-mode radius
 accounting-scheme scheme2
  accounting-mode radius
  accounting realtime 15
 domain example.com
  authentication-scheme abc
  accounting-scheme scheme2
  radius-server rd1
#
url-template name url1 
 url http://192.168.10.1:19008/portal 
 url-parameter set device-ip 192.168.1.1
 url-parameter device-ip ac-ip user-ipaddress uaddress user-mac umac 
#
web-auth-server server-source all-interface
web-auth-server abc
 server-ip 192.168.10.1
 source-ip 192.168.1.1
 port 50200
 shared-key cipher %^%#4~ZXE3]6@BXu;2;aw}hA{rSb,@"L@T#e{%6G1AiD%^%#
 url-template url1
#
portal-access-profile name portal1
 web-auth-server abc
#
authentication-profile name p1
 portal-access-profile portal1
 access-domain example.com force
#
interface MEth0/0/0
 ip address 10.1.1.1 255.255.255.0
#
interface Vlanif10
 ip address 192.168.1.1 255.255.255.0
 mac-address 0000-0000-0011
#
interface Vlanif11
 ip address 10.2.1.1 255.255.255.0 
 mac-address 0000-5e00-0101
#
interface Eth-Trunk0
 mode lacp-static
 peer-link 1
#
interface Eth-Trunk1
 port link-type access
 port default vlan 11
 mode lacp-static
 dfs-group 1 m-lag 1
 authentication-profile p1
#
interface Eth-Trunk2
 port link-type trunk
 port trunk allow-pass vlan 10 
 mode lacp-static 
 dfs-group 1 m-lag 2
#
interface 10GE1/0/1
 eth-trunk 2
#
interface 10GE1/0/2
 eth-trunk 1
#
interface 10GE1/0/3
 eth-trunk 0
#
interface 10GE1/0/4
 eth-trunk 0
#
interface 10GE1/0/5
 eth-trunk 1
#
interface 10GE1/0/6
 eth-trunk 2
#
return
#
sysname DeviceB
#
dfs-group 1
 priority 120
 dual-active detection source ip 10.1.1.2 peer 10.1.1.1
 authentication-mode hmac-sha256 password %+%##!!!!!!!!!"!!!!"!!!!*!!!!=I9f8>C{!P_bhB31@7r-=jrS8c|_"(Bn~#=!!!!!!!!!!!!!!!9!!!!kx-6@.tGA(wAt/IQXl6>[g{6YlOi9$!!!!!!!!!!%+%#
 dfs-group state switchover disable
#
vlan batch 10 to 11
#
stp mode rstp 
stp v-stp enable
#
access-user m-lag enable
#
radius-server template rd1
 radius-server shared-key cipher %+%##!!!!!!!!!"!!!!"!!!!*!!!!Cd/`W03KjAwAqn64E<\TxGC_SOri<2BP\A+!!!!!2jp5!!!!!!B!!!!Oe2HMc->XMa#TLDZUaJBFJtm#XVj*E:S*|(N7`J1B:3QY!!!!!!!!!!!!!!!%+%# 
 radius-server authentication 192.168.10.1 1812 source ip-address 192.168.1.1 weight 80
 radius-server accounting 192.168.10.1 1813 source ip-address 192.168.1.1 weight 80
#
aaa
 authentication-scheme abc
  authentication-mode radius
 accounting-scheme scheme2
  accounting-mode radius
  accounting realtime 15
 domain example.com
  authentication-scheme abc
  accounting-scheme scheme2
  radius-server rd1
#
url-template name url1 
 url http://192.168.10.1:19008/portal 
 url-parameter set device-ip 192.168.1.1
 url-parameter device-ip ac-ip user-ipaddress uaddress user-mac umac 
#
web-auth-server server-source all-interface
web-auth-server abc
 server-ip 192.168.10.1
 source-ip 192.168.1.1
 port 50200
 shared-key cipher %^%#4~ZXE3]6@BXu;2;aw}hA{rSb,@"L@T#e{%6G1AiD%^%#
 url-template url1
#
portal-access-profile name portal1
 web-auth-server abc
#
authentication-profile name p1
 portal-access-profile portal1
 access-domain example.com force
#
interface MEth0/0/0
 ip address 10.1.1.2 255.255.255.0
#
interface Vlanif10
 ip address 192.168.1.1.1 255.255.255.0
 mac-address 0000-0000-0011
#
interface Vlanif11 
 ip address 10.2.1.1 255.255.255.0 
 mac-address 0000-5e00-0101  
#
interface Eth-Trunk0
 mode lacp-static
 peer-link 1
#
interface Eth-Trunk1
 port link-type access
 port default vlan 11
 mode lacp-static
 dfs-group 1 m-lag 1
 authentication-profile p1
#
interface Eth-Trunk2
 port link-type trunk
 port trunk allow-pass vlan 10 
 mode lacp-static
 dfs-group 1 m-lag 2
#
interface 10GE1/0/1
 eth-trunk 2
#
interface 10GE1/0/2
 eth-trunk 1
#
interface 10GE1/0/3
 eth-trunk 0
#
interface 10GE1/0/4
 eth-trunk 0
#
interface 10GE1/0/5
 eth-trunk 1
#
interface 10GE1/0/6
 eth-trunk 2
#
return
Select the content with the mouse pointer to quickly report the problem.
