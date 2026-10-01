---
id: collect-261001-cisco/cisco/c-en-us-td-docs-routers-ios-config-17-x-sec-vpn-b-security-vpn-m-sec-flex-spoke-e0593c3a-2
title: "c-en-us-td-docs-routers-ios-config-17-x-sec-vpn-b-security-vpn-m-sec-flex-spoke--e0593c3a"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: ["agent"]
source: docs/RAG/collect-261001-cisco/c-en-us-td-docs-routers-ios-config-17-x-sec-vpn-b-security-vpn-m-sec-flex-spoke--e0593c3a.md
source_anchor: ""
source_lines: [47, 165]
sha256: 9793a97971d4e3a0ea8260d8cd940115b45527e5a7b7c3367b6e0f0162be540c
---

# c-en-us-td-docs-routers-ios-config-17-x-sec-vpn-b-security-vpn-m-sec-flex-spoke--e0593c3a

| Step 4 | Do one of the following:  Example: Device(config-if)# ip nhrp shortcut 1Example: Device(config-if)# ipv6 nhrp shortcut 1 | Enables NHRP shortcuts on the FlexVPN client tunnel interface. This is necessary to establish spoke-to-spoke tunnels. The virtual-template number specified in this configuration and the virtual-template number specified in the Configuring the Virtual Tunnel Interface on the FlexVPN Spoke task must be same. | 
| Step 5 | exit Example: Device(config-if)# exit | Exits interface configuration mode and returns to global configuration mode. | 
|  | Command or Action | Purpose | 
|---|---|---|
| Step 1 | enable Example: Device> enable | Enables privileged EXEC mode.  | 
| Step 2 | configure terminal Example: Device# configure terminal | Enters global configuration mode. | 
| Step 3 | interface virtual-template number type tunnel Example: Device(config)# interface virtual-template 1 type tunnel | Creates a virtual template interface that can be configured and applied dynamically to create virtual access interfaces. | 
| Step 4 | ip unnumbered tunnel number Example: Device(config-if)# ip unnumbered tunnel 1 | Assigns the IPv4 address of the FlexVPN tunnel interface to the virtual tunnel interface. | 
| Step 5 | Do one of the following:  Example: Device(config-if)# ip nhrp network-id 1Example: Device(config-if)# ipv6 nhrp network-id 1 | Enables NHRP on the interface. | 
| Step 6 | Do one of the following:  Example: Device(config-if)# ip nhrp shortcut 1Example: Device(config-if)# ipv6 nhrp shortcut 1 | Enables NHRP shortcut switching on an interface. | 
| Step 7 | ip nhrp redirect [timeout seconds] Example: Device(config-if)# ip nhrp redirect | Enables NHRP redirects on the virtual tunnel interface. This is useful when networks move from one spoke to another.  | 
| Step 8 | exit Example: Device(config-if)# exit | Exits interface configuration mode and returns to global configuration mode. | 
| Note |  | 
Use the following commands to verify the FlexVPN spoke configuration.
| Step 1 | show crypto ikev2 client flexvpn Example: Device# show crypto ikev2 client flexvpn    Profile : flexblk   Current state:ACTIVE   Peer : 4001::2000:1   Source : Ethernet0/0   ivrf : IP DEFAULT   fvrf : IP DEFAULT   Backup group: None   Tunnel interface : Tunnel0   | 
| Step 2 | show ipv6 route Example: Device# show ipv6 route  IPv6 Routing Table - default - 15 entries Codes: C - Connected, L - Local, S - Static, U - Per-user Static route        B - BGP, HA - Home Agent, MR - Mobile Router, R - RIP        H - NHRP, I1 - ISIS L1, I2 - ISIS L2, IA - ISIS interarea        IS - ISIS summary, D - EIGRP, EX - EIGRP external, NM - NEMO        ND - ND Default, NDp - ND Prefix, DCE - Destination, NDr - Redirect        l - LISP        O - OSPF Intra, OI - OSPF Inter, OE1 - OSPF ext 1, OE2 - OSPF ext 2        ON1 - OSPF NSSA ext 1, ON2 - OSPF NSSA ext 2 C   3001::/112 [0/0]      via Tunnel0, directly connected S   3001::1/128 [2/0], tag 1      via 3001::1, Virtual-Access1 [Shortcut]      via Virtual-Access1, directly connected L   3001::2/128 [0/0]      via Tunnel0, receive S   3001::3/128 [2/0], tag 1      via Tunnel0, directly connected C   4001::2000:0/112 [0/0]      via Ethernet0/0, directly connected L   4001::2000:3/128 [0/0]      via Ethernet0/0, receive S   5001::/64 [2/0], tag 1      via Tunnel0, directly connected C   5001::2000:0/112 [0/0]      via Loopback0, directly connected L   5001::2000:1/128 [0/0]      via Loopback0, receive D   5001::3000:0/112 [90/28288000]      via FE80::A8BB:CCFF:FE01:F400, Tunnel0 D   5001::4000:0/112 [90/28288000]      via FE80::A8BB:CCFF:FE01:F400, Tunnel0 H   5001::4000:1/128 [250/1]      via 3001::1, Virtual-Access1 C   5001::5000:0/112 [0/0]      via Loopback1, directly connected L   5001::5000:1/128 [0/0]      via Loopback1, receive L   FF00::/8 [0/0]      via Null0, receive   | 
| Step 3 | show ipv6 nhrp Example: Device# show ipv6 nhrp  3001::1/128 via 3001::1    Virtual-Access1 created 00:01:52, expire 01:58:14    Type: dynamic, Flags: router implicit rib nho    NBMA address: 172.17.1.9     (Claimed NBMA address: 172.16.2.1) 5001::4000:1/128 via 3001::1    Virtual-Access1 created 00:00:56, expire 01:59:03    Type: dynamic, Flags: router rib    NBMA address: 172.17.1.9     (Claimed NBMA address: 172.16.2.1) 5001::5000:1/128 via 3001::2    Virtual-Access1 created 00:01:52, expire 01:58:14    Type: dynamic, Flags: router unique local    NBMA address: 172.17.2.10 Example: Device# show ipv6 nhrp  3001::1/128 via 3001::1    Virtual-Access1 created 00:01:52, expire 01:58:14    Type: dynamic, Flags: router implicit rib nho    NBMA address: 4001::2000:2 5001::4000:1/128 via 3001::1    Virtual-Access1 created 00:00:56, expire 01:59:03    Type: dynamic, Flags: router rib    NBMA address: 4001::2000:2 5001::5000:1/128 via 3001::2    Virtual-Access1 created 00:01:52, expire 01:58:14    Type: dynamic, Flags: router unique local    NBMA address: 4001::2000:3   | 
Here are few tips for troubleshooting FlexVPN spoke configuration:
| Problem | Troubleshooting Tips | 
|---|---|
| Spoke to hub connection is not created. | A connection may not be created due to the absence of virtual access interfaces created at the hub.  | 
| Spoke to spoke tunnel is not created. | Traffic must flow from spoke to spoke via the hub to initiate a spoke to spoke tunnel.  | 
The following example shows how to configure FlexVPN spoke to spoke with IKE-propagated static routing on the FlexVPN server and the FlexVPN client. The following is the configuration on the FlexVPN server:
hostname hub
!
crypto ikev2 authorization policy default
 pool flex-pool
 def-domain cisco.com
 route set interface
 route set access-list flex-route
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
ip access-list standard flex-route
 permit any
The following is the configuration on the first FlexVPN client:
hostname spoke1
!
crypto ikev2 authorization policy default
 route set interface
 route set access-list flex-route
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
ip access-list standard flex-route
 permit 192.168.110.0 0.0.0.255
The following is the configuration on the second FlexVPN client:
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
