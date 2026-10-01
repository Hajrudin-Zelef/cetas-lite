---
id: collect-261001-cisco/cisco/c-en-us-td-docs-routers-ios-config-17-x-sec-vpn-b-security-vpn-m-sec-flex-spoke-e0593c3a-3
title: "c-en-us-td-docs-routers-ios-config-17-x-sec-vpn-b-security-vpn-m-sec-flex-spoke--e0593c3a"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-cisco/c-en-us-td-docs-routers-ios-config-17-x-sec-vpn-b-security-vpn-m-sec-flex-spoke--e0593c3a.md
source_anchor: ""
source_lines: [166, 361]
sha256: 5ddc8af3fd80e4b3aef31c6d6c204cef88339eb7c629e2decae09ec1e444c500
---

# c-en-us-td-docs-routers-ios-config-17-x-sec-vpn-b-security-vpn-m-sec-flex-spoke--e0593c3a

 set ikev2-profile default
!
interface Tunnel0
 ip address negotiated
 ip nhrp network-id 1
 ip nhrp shortcut virtual-template 1
 ip nhrp redirect
 tunnel source Ethernet0/0
 tunnel destination 10.0.0.100
 tunnel protection ipsec profile default
!
interface Ethernet0/0
 ip address 10.0.0.120 255.255.255.0
!
interface Ethernet1/0
 ip address 192.168.120.1 255.255.255.0
!
interface Virtual-Template1 type tunnel
 ip unnumbered Tunnel0
 ip nhrp network-id 1
 ip nhrp shortcut virtual-template 1
 ip nhrp redirect
 tunnel protection ipsec profile default
!
ip access-list standard flex-route
 permit 192.168.120.0 0.0.0.255
The following example shows how to configure FlexVPN spoke to spoke with dynamic routing, using BGP on the FlexVPN server (with dynamic neighbor discovery) and the FlexVPN client. The following is the configuration on the FlexVPN server:
hostname hub
!
crypto ikev2 authorization policy default
 pool flex-pool
 def-domain cisco.com
 route set interface
!
crypto ikev2 profile default
 match identity remote fqdn domain cisco.com
 identity local fqdn hub.cisco.com
 authentication local rsa-sig
 authentication remote rsa-sig
 pki trustpoint CA
 aaa authorization group cert list default default
 virtual-template 1
!
crypto ipsec profile default
 set ikev2-profile default
!
interface Loopback0
 ip address 172.16.1.1 255.255.255.255
!
interface Ethernet0/0
 ip address 10.0.0.100 255.255.255.0
!
interface Virtual-Template1 type tunnel
 ip unnumbered Loopback0
 ip nhrp network-id 1
 ip nhrp redirect
 tunnel protection ipsec profile default
!
ip local pool flex-pool 172.16.0.1 172.16.0.254
!
router bgp 65100
 bgp router-id 10.0.0.100
 bgp log-neighbor-changes
 bgp listen range 172.16.0.0/24 peer-group spokes
 neighbor spokes peer-group
 neighbor spokes remote-as 65100
 neighbor spokes transport connection-mode passive
 neighbor spokes update-source Loopback0
 !
 address-family ipv4
  neighbor spokes activate
  neighbor spokes default-originate
  neighbor spokes prefix-list no-default in
 exit-address-family
!
ip prefix-list no-default seq 5 deny 0.0.0.0/0
ip prefix-list no-default seq 10 permit 0.0.0.0/0 le 32
hostname spoke1
!
crypto ikev2 authorization policy default
 route set interface
!
crypto ikev2 profile default
 match identity remote fqdn domain cisco.com
 identity local fqdn spoke1.cisco.com
 authentication local rsa-sig
 authentication remote rsa-sig
 pki trustpoint CA
 aaa authorization group cert list default default
 virtual-template 1
!
crypto ipsec profile default
 set ikev2-profile default
!
interface Tunnel0
 ip address negotiated
 ip nhrp network-id 1
 ip nhrp shortcut virtual-template 1
 ip nhrp redirect
 tunnel source Ethernet0/0
 tunnel destination 10.0.0.100
 tunnel protection ipsec profile default
!
interface Ethernet0/0
 ip address 10.0.0.110 255.255.255.0
!
interface Ethernet1/0
 ip address 192.168.110.1 255.255.255.0
!
interface Virtual-Template1 type tunnel
 ip unnumbered Tunnel0
 ip nhrp network-id 1
 ip nhrp shortcut virtual-template 1
 ip nhrp redirect
 tunnel protection ipsec profile default
!
router bgp 65100
 bgp router-id 10.0.0.110
 bgp log-neighbor-changes
 neighbor hubs peer-group
 neighbor hubs remote-as 65100
 neighbor hubs update-source Tunnel0
 neighbor 172.16.1.1 peer-group hubs
 !
 address-family ipv4
  network 192.168.110.0
  neighbor 172.16.1.1 activate
 exit-address-family
hostname spoke2
!
crypto ikev2 authorization policy default
 route set interface
 route set access-list flex-route
!
crypto ikev2 profile default
 match identity remote fqdn domain cisco.com
 identity local fqdn spoke2.cisco.com
 authentication local rsa-sig
 authentication remote rsa-sig
 pki trustpoint CA
 aaa authorization group cert list default default
 virtual-template 1
!
crypto ipsec profile default
 set ikev2-profile default
!
interface Tunnel0
 ip address negotiated
 ip nhrp network-id 1
 ip nhrp shortcut virtual-template 1
 ip nhrp redirect
 tunnel source Ethernet0/0
 tunnel destination 10.0.0.100
 tunnel protection ipsec profile default
!
interface Ethernet0/0
 ip address 10.0.0.120 255.255.255.0
!
interface Ethernet1/0
 ip address 192.168.120.1 255.255.255.0
!
interface Virtual-Template1 type tunnel
 ip unnumbered Tunnel0
 ip nhrp network-id 1
 ip nhrp shortcut virtual-template 1
 ip nhrp redirect
 tunnel protection ipsec profile default
!
router bgp 65100
 bgp router-id 10.0.0.120
 bgp log-neighbor-changes
 neighbor hubs peer-group
 neighbor hubs remote-as 65100
 neighbor hubs update-source Tunnel0
 neighbor 172.16.1.1 peer-group hubs
 !
 address-family ipv4
  network 192.168.120.0
  neighbor 172.16.1.1 activate
 exit-address-family
| Related Topic | Document Title | 
|---|---|
| Cisco IOS commands | Cisco IOS Master Command List, All Releases | 
| Security commands |  | 
| Shortcut Switching Enhancements | Shortcut Switching Enhancements for NHRP in DMVPN Networks | 
| Description | Link | 
|---|---|
| The Cisco Support and Documentation website provides online resources to download documentation, software, and tools. Use these resources to install and configure the software and to troubleshoot and resolve technical issues with Cisco products and technologies. Access to most tools on the Cisco Support and Documentation website requires a Cisco.com user ID and password. | http://www.cisco.com/cisco/web/support/index.html | 
The following table provides release information about the feature or features described in this module. This table lists only the software release that introduced support for a given feature in a given software release train. Unless noted otherwise, subsequent releases of that software release train also support that feature.
Use Cisco Feature Navigator to find information about platform support and Cisco software image support. To access Cisco Feature Navigator, go to www.cisco.com/go/cfn. An account on Cisco.com is not required.
| Table 1. Feature Information for                                     		  FlexVPN Spoke to Spoke |  |  | 
|---|---|---|
| Feature Name | Releases | Feature Information | 
|---|---|---|
| FlexVPN Spoke to Spoke |  | The FlexVPN Spoke to Spoke feature enables a FlexVPN client to establish a direct crypto channel with another FlexVPN client. The feature leverages VTIs, IKEv2, and NHRP to build spoke to spoke connections. In Cisco IOS Release 15.2(2)T, this feature was introduced. The following commands were introduced or modified: ip unnumbered loopback0, tunnel source, tunnel mode gre ip, nhrp network-id, ip nhrp redirect, ip nhrp shortcut. | 
| FlexVPN IPv6 Direct Spoke to Spoke |  | The FlexVPN IPv6 Direct Spoke to Spoke feature supports the use of IPv6 addresses for FlexVPN spokes. The support for IPv6 addresses provides support for IPv6 over IPv4, IPv4 over IPv6, and IPv6 over IPv6 transports. The following commands were introduced or modified: ipv6 nhrp shortcut. |
