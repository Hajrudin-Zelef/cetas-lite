---
id: collect-261001-cisco/cisco/enterprise-en-doc-edoc1100246296-3f27e589-user-access-and-authentication-c126a3e3-26
title: "Configure DeviceA to generate a local key pair."
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: ["voice"]
source: docs/RAG/collect-261001-cisco/enterprise-en-doc-edoc1100246296-3f27e589-user-access-and-authentication-c126a3e3.md
source_anchor: ""
source_lines: [4552, 4759]
sha256: 6ce89eb1575eb60aef8f80467bd22bcb44a35496bd0a1f729395dfa3140190ca
---

# Configure DeviceA to generate a local key pair.

 radius-server shared-key cipher %+%##!!!!!!!!!"!!!!"!!!!*!!!!SKvr${[Fs.3t@/5k|BENhEu>W(3\~XG!!D;!!!!!2jp5!!!!!!A!!!!3"pK8qv!}9M#(4$jGWvQF/R[CNe/+:W^jk8HUe&W%+%#
 radius-server authentication 10.5.1.3 1812 weight 80
 radius-server accounting 10.5.1.3 1813 weight 80
#
aaa
 authentication-scheme abc
  authentication-mode radius
 accounting-scheme acco1
  accounting-mode radius
  accounting realtime 15 
 domain isp
  authentication-scheme abc
  accounting-scheme acco1
  radius-server rd1
#
dhcp enable
#
interface Vlanif100
 ip address 10.1.1.1 255.255.255.0
 dhcp select interface
#
interface Vlanif200
 ip address 10.2.1.1 255.255.255.0
 dhcp select interface
#
interface Vlanif300
 ip address 10.3.1.1 255.255.255.0
#
interface 10GE1/0/1
 port link-type hybrid
 port hybrid pvid vlan 100     
 port hybrid tagged vlan 200
 port hybrid untagged vlan 100    
 stp edged-port enable 
 lldp mdn enable txrx 
 lldp tlv-enable med-tlv network-policy voice-vlan vlan 200 cos 6 dscp 60 
 voice-vlan 200 enable
 authentication-profile p1
#
interface 10GE1/0/2
 port link-type hybrid
 port hybrid pvid vlan 100 
 port hybrid tagged vlan 200    
 port hybrid untagged vlan 100 
 stp edged-port enable 
 lldp mdn enable txrx 
 lldp tlv-enable med-tlv network-policy voice-vlan vlan 200 cos 6 dscp 60 
 voice-vlan 200 enable
 authentication-profile p1
#
interface 10GE1/0/3
 port link-type access
 port default vlan 300 
#
ip route-static 0.0.0.0 0.0.0.0 10.3.1.2
#
dot1x-access-profile name d1
 dot1x timer client-timeout 30
#
mac-access-profile name m1
#
return
In Figure 3-139, to save investment costs, users require that IP phones and PCs connect to the network through VoIP. IP phones support LLDP and can obtain the voice VLAN through LLDP. The network needs to meet the following requirements:
IP phones connect to DeviceA through MAC address authentication and PCs connect to DeviceA through 802.1X authentication.
# Enable MAC address bypass authentication.
[DeviceA-authen-profile-p1] access-domain isp dot1x force
[DeviceA-authen-profile-p1] access-domain isp mac-authen force   
[DeviceA-authen-profile-p1] quit
After a phone starts, 802.1X authentication is triggered. When 802.1X authentication fails, MAC address authentication is triggered, and MAC address authentication of the phone is successful.
#
sysname DeviceA
#
vlan batch 100 200 300
#
authentication-profile name p1
 mac-access-profile m1
 dot1x-access-profile d1 
 authentication dot1x-mac-bypass
 access-domain isp dot1x force
 access-domain isp mac-authen force
#
radius-server template rd1
 radius-server shared-key cipher %+%##!!!!!!!!!"!!!!"!!!!*!!!!SKvr${[Fs.3t@/5k|BENhEu>W(3\~XG!!D;!!!!!2jp5!!!!!!A!!!!3"pK8qv!}9M#(4$jGWvQF/R[CNe/+:W^jk8HUe&W%+%#
 radius-server authentication 10.5.1.3 1812 weight 80
 radius-server accounting 10.5.1.3 1813 weight 80
#
aaa
 authentication-scheme abc
  authentication-mode radius
 accounting-scheme acco1
  accounting-mode radius
  accounting realtime 15 
 domain isp
  authentication-scheme abc
  accounting-scheme acco1
  radius-server rd1
#
dhcp enable
#
interface Vlanif100
 ip address 10.1.1.1 255.255.255.0
 dhcp select interface
#
interface Vlanif200
 ip address 10.2.1.1 255.255.255.0
 dhcp select interface
#
interface Vlanif300
 ip address 10.3.1.1 255.255.255.0
#
interface 10GE1/0/1
 port link-type hybrid
 port hybrid pvid vlan 100     
 port hybrid tagged vlan 200
 port hybrid untagged vlan 100    
 stp edged-port enable 
 lldp mdn enable txrx
 lldp tlv-enable med-tlv network-policy voice-vlan vlan 200 cos 6 dscp 60 
 voice-vlan 200 enable
 authentication-profile p1
#
interface 10GE1/0/2
 port link-type hybrid
 port hybrid pvid vlan 100 
 port hybrid tagged vlan 200    
 port hybrid untagged vlan 100 
 stp edged-port enable 
 lldp mdn enable txrx
 lldp tlv-enable med-tlv network-policy voice-vlan vlan 200 cos 6 dscp 60 
 voice-vlan 200 enable
 authentication-profile p1
#
interface 10GE1/0/3
 port link-type access
 port default vlan 300 
#
ip route-static 0.0.0.0 0.0.0.0 10.3.1.2
#
dot1x-access-profile name d1
 dot1x timer client-timeout 30
#
mac-access-profile name m1
#
return
By default, the device can trigger MAC address authentication for users upon receiving DHCP, ARP, DHCPv6, or ND packets. If a static IPv4 address is configured for a client, no DHCP or ARP packet is exchanged between the client and device. In this case, the function of triggering MAC address authentication through LLDP or CDP packets needs to be configured. During access authentication of an IP phone, LLDP or CDP negotiation packets sent from the IP phone do not carry VLAN tags. After receiving the packets, an access device adds data VLAN tags to the packets. Since the function of triggering MAC address authentication through any packet is configured, the IP phone is authenticated and goes online in the data VLAN.
In Figure 3-140, DeviceA is connected to a voice service device (IP phone) and a data service device (PC), and uses VLAN 200 to transmit voice packets and VLAN 100 to transmit data packets. The PC is connected to the IP phone, and the IP phone is connected to DeviceA. Users require high quality of the voice service; therefore, voice data flows must be transmitted with a high priority to ensure the call quality.
<HUAWEI> sysname DeviceA
[DeviceA] vlan batch 100 200 300
[DeviceA] lldp enable   
# Configure the PVID of 10GE 1/0/1 and the allowed data VLAN on 10GE 1/0/1.
[DeviceA] interface 10ge 1/0/1
[DeviceA-10GE 1/0/1] portswitch
[DeviceA-10GE 1/0/1] port link-type hybrid 
[DeviceA-10GE 1/0/1] port hybrid pvid vlan 100
[DeviceA-10GE 1/0/1] port hybrid untagged vlan 100
[DeviceA-10GE 1/0/1] quit
[DeviceA] interface vlanif 100
[DeviceA-Vlanif100] ip address 192.168.100.10 24
[DeviceA-Vlanif100] quit
[DeviceA] interface vlanif 200
[DeviceA-Vlanif200] ip address 192.168.200.10 24
[DeviceA-Vlanif200] quit
# Configure 10GE 1/0/2 connecting DeviceA to the RADIUS server as an access interface and add the interface to VLAN 300.
[DeviceA] interface 10ge 1/0/2
[DeviceA-10GE1/0/2] portswitch
[DeviceA-10GE1/0/2] port link-type access
[DeviceA-10GE1/0/2] port default vlan 300 
[DeviceA-10GE1/0/2] quit
[DeviceA] interface vlanif 300
[DeviceA-Vlanif300] ip address 192.168.1.20 24
[DeviceA-Vlanif300] quit
[DeviceA] voice-vlan mac-address 00e0-fc12-3456 mask ffff-ff00-0000   
[DeviceA] interface 10ge 1/0/1
[DeviceA-10GE 1/0/1] voice-vlan 200 enable  
[DeviceA-10GE 1/0/1] stp edged-port enable   
[DeviceA-10GE 1/0/1] port hybrid tagged vlan 200
[DeviceA-10GE 1/0/1] voice-vlan remark-mode mac-address  
[DeviceA-10GE 1/0/1] lldp tlv-enable med-tlv network-policy voice-vlan vlan 200   
[DeviceA-10GE 1/0/1] lldp mdn enable txrx   //For IP phones that use non-standard discovery protocols such as CDP, perform this step to ensure that the device can establish neighbor relationships with these IP phones.
[DeviceA-10GE 1/0/1] quit
[DeviceA] ip route-static 192.168.2.0 255.255.255.0 192.168.1.2
# Configure the MAC access profile m1 and enable the function of triggering MAC address authentication through any packet.
[DeviceA] mac-access-profile name m1
[DeviceA-mac-access-profile-m1] authentication trigger-condition dhcp arp dhcpv6 nd any-l2-packet
[DeviceA-mac-access-profile-m1] quit
The IP phone can go online and deliver clear, high-quality voice communication. The PC can also go online normally.
#
sysname DeviceA
#
vlan batch 100 200 300
#
authentication-profile name p1
 mac-access-profile m1
 access-domain example.com force
#
radius-server template rd1
 radius-server shared-key cipher %+%##!!!!!!!!!"!!!!"!!!!*!!!!SKvr${[Fs.3t@/5k|BENhEu>W(3\~XG!!D;!!!!!2jp5!!!!!!A!!!!3"pK8qv!}9M#(4$jGWvQF/R[CNe/+:W^jk8HUe&W%+%#
 radius-server authentication 192.168.2.30 1812 weight 80
 radius-server accounting 192.168.2.30 1813 weight 80
#
aaa
 authentication-scheme abc
  authentication-mode radius
 accounting-scheme scheme2
  accounting-mode radius
  accounting realtime 15
 domain example.com
  authentication-scheme abc
