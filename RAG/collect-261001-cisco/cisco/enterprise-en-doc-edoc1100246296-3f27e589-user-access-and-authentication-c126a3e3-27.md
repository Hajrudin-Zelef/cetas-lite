---
id: collect-261001-cisco/cisco/enterprise-en-doc-edoc1100246296-3f27e589-user-access-and-authentication-c126a3e3-27
title: "Configure DeviceA to generate a local key pair."
domain: cisco
role: reference
task: reference
actors: ["Huawei"]
dates: []
keywords: ["agent", "voice"]
source: docs/RAG/collect-261001-cisco/enterprise-en-doc-edoc1100246296-3f27e589-user-access-and-authentication-c126a3e3.md
source_anchor: ""
source_lines: [4760, 4915]
sha256: 680ef663ee7258880393fdfd68dd80ce53970df55c22b284d58f8aac13f6e797
---

# Configure DeviceA to generate a local key pair.

  accounting-scheme scheme2
  radius-server rd1
#
interface Vlanif100
 ip address 192.168.100.10 255.255.255.0
#
interface Vlanif200
 ip address 192.168.200.10 255.255.255.0
#
interface Vlanif300
 ip address 192.168.1.20 255.255.255.0
#
interface 10GE1/0/1
 port link-type hybrid
 port hybrid pvid vlan 100     
 port hybrid tagged vlan 200
 port hybrid untagged vlan 100    
 stp edged-port enable 
 authentication-profile p1
 lldp tlv-enable med-tlv network-policy voice-vlan vlan 200
 lldp mdn enable txrx
 voice-vlan 200 enable
 voice-vlan remark-mode mac-address
#
interface 10GE1/0/2
 port link-type access
 port default vlan 300
#
ip route-static 192.168.2.0 255.255.255.0 192.168.1.2
#
mac-access-profile name m1
#
voice-vlan mac-address 00e0-fc12-3456 mask ffff-ff00-0000
#
return
In Figure 3-141, to save investment costs, the customer requires that IP phones access the network through VoIP. The IP phones cannot obtain voice VLAN IDs and can only send untagged voice packets. The network needs to meet the following requirements:
| Item | Value | 
|---|---|
| Voice VLAN | VLAN 100 | 
| MAC address | 00e0-fcc7-0001 00e0-fc8f-0002 | 
| Address segment | 10.20.20.1/24 | 
| Authentication mode | MAC address authentication | 
| Item | Value | 
|---|---|
| VLAN and IP address used by SwitchA to communicate with SwitchB | VLAN 200 and 10.10.20.1/24 | 
| VLAN and IP address used by SwitchB to communicate with SwitchA | VLAN 200 and 10.10.20.2/24 | 
| IP address of SwitchA | 192.168.100.200 | 
| MAC access profile name | ipphone | 
| IP address of the RADIUS authentication and accounting server | 192.168.100.182 | 
| Port number of the RADIUS authentication server | 1812 | 
| Port number of the RADIUS accounting server | 1813 | 
| RADIUS shared key | Huawei2012 | 
# Create voice VLAN 100.
<HUAWEI> system-view 
[HUAWEI] sysname SwitchA 
[SwitchA] vlan batch 100
# Add interfaces to VLAN 100 in untagged mode.
[SwitchA] interface 10ge 1/0/1 
[SwitchA-10GE 1/0/1] port link-type hybrid 
[SwitchA-10GE 1/0/1] port hybrid untagged vlan 100  //Add the interface to voice VLAN 100 in untagged mode because packets sent by IP phones do not carry tags.
[SwitchA-10GE 1/0/1] quit 
[SwitchA] interface 10ge 1/0/2
[SwitchA-10GE1/0/2] port link-type hybrid 
[SwitchA-10GE1/0/2] port hybrid untagged vlan 100 
[SwitchA-10GE1/0/2] quit
[SwitchA] interface 10ge 1/0/1  
[SwitchA-10GE 1/0/1] voice-vlan 100 enable  //Enable the voice VLAN function on the interface.
[SwitchA-10GE 1/0/1] voice-vlan remark-mode mac-address  //Enable the interface to identify voice packets based on MAC addresses.
[SwitchA-10GE 1/0/1] port hybrid pvid vlan 100  //Set the PVID of the interface.
[SwitchA-10GE 1/0/1] quit 
[SwitchA] interface 10ge 1/0/2
[SwitchA-10GE1/0/2] voice-vlan 100 enable 
[SwitchA-10GE1/0/2] voice-vlan remark-mode mac-address 
[SwitchA-10GE1/0/2] port hybrid pvid vlan 100 
[SwitchA-10GE1/0/2] quit 
[SwitchA] voice-vlan mac-address 00e0-fcc7-0000 mask ffff-ffff-0000   
[SwitchA] voice-vlan mac-address 00e0-fc8f-0000 mask ffff-ffff-0000
# Configure the DHCP relay function on an interface.
[SwitchA] dhcp enable  //Enable DHCP globally.
[SwitchA] interface Vlanif 100 
[SwitchA-Vlanif100] ip address 10.20.20.1 255.255.255.0  //Configure an IP address for VLANIF 100.
[SwitchA-Vlanif100] dhcp select relay  //Enable the DHCP relay function on VLANIF 100.
[SwitchA-Vlanif100] dhcp relay server-ip 10.10.20.2  //Configure the DHCP server address on the DHCP relay agent.
[SwitchA-Vlanif100] quit
# Create VLANIF 200.
[SwitchA] vlan batch 200 
[SwitchA] interface Vlanif 200 
[SwitchA-Vlanif200] ip address 10.10.20.1 255.255.255.0  //Configure an IP address for VLANIF 200 for communication with SwitchB.
[SwitchA-Vlanif200] quit
# Add the uplink interface to VLAN 200.
[SwitchA] interface 10ge 1/0/3 
[SwitchA-10GE1/0/3] port link-type access 
[SwitchA-10GE1/0/3] port default vlan 200 
[SwitchA-10GE1/0/3] quit
# Configure a default static route.
[SwitchA] ip route-static 0.0.0.0 0.0.0.0 10.10.20.2  //The next-hop address of the route is the IP address of VLANIF 200 on SwitchB.
# Configure an address pool.
<HUAWEI> system-view 
[HUAWEI] sysname SwitchB 
[SwitchB] ip pool ip-phone  //Create an address pool to allocate IP addresses to IP phones.
[SwitchB-ip-pool-ip-phone] gateway-list 10.20.20.1  //Configure the gateway address of IP phones.
[SwitchB-ip-pool-ip-phone] network 10.20.20.0 mask 255.255.255.0  //Configure the network segment where IP addresses can be allocated in the address pool.
[SwitchB-ip-pool-ip-phone] quit
# Configure the DHCP server function.
[SwitchB] dhcp enable  //Enable DHCP globally. By default, DHCP is disabled.
[SwitchB] vlan batch 200 
[SwitchB] interface Vlanif 200  //Create VLANIF 200.
[SwitchB-Vlanif200] ip address 10.10.20.2 255.255.255.0  //Configure an IP address for the VLANIF interface.
[SwitchB-Vlanif200] dhcp select global  //Configure the interface to allocate IP addresses to IP phones from the global address pool.
[SwitchB-Vlanif200] quit
# Add the downlink interface to VLAN 200.
[SwitchB] interface 10ge 1/0/3 
[SwitchB-10GE1/0/3] port link-type access 
[SwitchB-10GE1/0/3] port default vlan 200 
[SwitchB-10GE1/0/3] quit
# Configure a return route.
[SwitchB] ip route-static 10.20.20.0 255.255.255.0 10.10.20.1
# Create and configure a RADIUS server template.
[SwitchA] radius-server template ipphone  //Create a RADIUS server template named ipphone.
[SwitchA-radius-ipphone] radius-server authentication 192.168.100.182 1812  //Configure the IP address and port number of the RADIUS authentication server.
[SwitchA-radius-ipphone] radius-server accounting 192.168.100.182 1813  //Configure the IP address and port number of the RADIUS accounting server.
[SwitchA-radius-ipphone] radius-server shared-key cipher YsHsjx_202206  //Configure the RADIUS shared key.
[SwitchA-radius-ipphone] quit
# Configure an authentication scheme.
[SwitchA] aaa 
[SwitchA-aaa] authentication-scheme radius  //Create an AAA authentication scheme named radius.
[SwitchA-aaa-authen-radius] authentication-mode radius  //Set the authentication mode to RADIUS.
[SwitchA-aaa-authen-radius] quit
# Create an AAA domain, and configure the RADIUS server template and authentication scheme.
[SwitchA-aaa] domain default  //Configure a domain named default.
[SwitchA-aaa-domain-default] authentication-scheme radius  //Bind the authentication scheme radius to the domain.
[SwitchA-aaa-domain-default] radius-server ipphone  //Bind the RADIUS server template ipphone to the domain.
[SwitchA-aaa-domain-default] quit 
[SwitchA-aaa] quit
# Configure an access profile.
[SwitchA] mac-access-profile name ipphone  //Create a MAC access profile named ipphone.
[SwitchA-mac-access-profile-ipphone] quit
# Configure an authentication profile.
[SwitchA] authentication-profile name ipphone  //Configure an authentication profile.
[SwitchA-authen-profile-ipphone] mac-access-profile ipphone  //Bind the MAC access profile to the authentication profile.
[SwitchA-authen-profile-ipphone] quit
# Apply the authentication profile to interfaces.
[SwitchA] interface 10ge 1/0/1  
[SwitchA-10GE 1/0/1] authentication-profile ipphone  
[SwitchA-10GE 1/0/1] quit 
[SwitchA] interface 10ge 1/0/2
[SwitchA-10GE1/0/2] authentication-profile ipphone 
[SwitchA-10GE1/0/2] quit
In the authorization settings, set the VLAN to be authorized to VLAN 100, authorize the attribute HW-Voice-Vlan (33), and set the attribute value to 1.
If the authorized VLAN is used together with the attribute HW-Voice-Vlan (33), the authorized VLAN is a voice VLAN.
[SwitchA] display access-user 
 ------------------------------------------------------------------------------  
 UserID Username     IP address       MAC            Status           
 ------------------------------------------------------------------------------  
 564   001bd4c71fa9  10.20.20.198     00e0-fcc7-1fa9 Success         
 565   0021a08f2fa8  10.20.20.199     00e0-fc8f-2fa8 Success          
