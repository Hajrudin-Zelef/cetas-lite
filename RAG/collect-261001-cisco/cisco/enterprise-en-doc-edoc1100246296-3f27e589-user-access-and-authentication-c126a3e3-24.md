---
id: collect-261001-cisco/cisco/enterprise-en-doc-edoc1100246296-3f27e589-user-access-and-authentication-c126a3e3-24
title: "Configure DeviceA to generate a local key pair."
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-cisco/enterprise-en-doc-edoc1100246296-3f27e589-user-access-and-authentication-c126a3e3.md
source_anchor: ""
source_lines: [4149, 4319]
sha256: 1b58655bf25d6d83d31fa3ae5321674500b6ea80dd8b06a96029c08f6c319303
---

# Configure DeviceA to generate a local key pair.

On the enterprise network shown in Figure 3-137, DeviceC, DeviceB, and DeviceA function as an access switch, an aggregation switch, and a core switch, respectively. DeviceA assigns IP addresses to users and APs. Wired users are connected to DeviceC and can access the enterprise network only after passing 802.1X or MAC address authentication. Wireless users are connected to the AP that is connected to DeviceC and can access the enterprise network only after passing MAC address-prioritized Portal authentication. Authentication is not required for the AP as the authentication server does not store AP information.
# On DeviceA, create VLANs and configure allowed VLANs on interfaces. (VLANIF 100 is the source interface of the AC, VLAN 200 is the wired service VLAN, VLAN 201 is the wireless service VLAN, and VLAN 300 is the uplink VLAN for communicating with servers.)
<HUAWEI> system-view
[HUAWEI] sysname DeviceA
[DeviceA] vlan batch 100 200 201 300
[DeviceA] interface eth-trunk 10
[DeviceA-Eth-Trunk10] port link-type trunk
[DeviceA-Eth-Trunk10] port trunk allow-pass vlan 100 200
[DeviceA-Eth-Trunk10] quit
[DeviceA] interface 10ge 1/0/1
[DeviceA-10GE1/0/1] eth-trunk 10
[DeviceA-10GE1/0/1] quit
[DeviceA] interface 10ge 1/0/2
[DeviceA-10GE1/0/2] eth-trunk 10
[DeviceA-10GE1/0/2] quit
[DeviceA] interface 10ge 1/0/3
[DeviceA-10GE1/0/3] port link-type trunk
[DeviceA-10GE1/0/3] port trunk allow-pass vlan 300
[DeviceA-10GE1/0/3] quit
[DeviceA] interface vlanif 100
[DeviceA-Vlanif100] ip address 10.1.1.1 24  
[DeviceA-Vlanif100] quit
[DeviceA] interface vlanif 200
[DeviceA-Vlanif200] ip address 10.2.1.1 24
[DeviceA-Vlanif200] quit
[DeviceA] interface vlanif 201
[DeviceA-Vlanif201] ip address 10.2.2.1 24
[DeviceA-Vlanif201] quit
[DeviceA] interface vlanif 300
[DeviceA-Vlanif300] ip address 10.3.1.1 24
[DeviceA-Vlanif300] quit
[DeviceA] ip route-static 10.5.1.0 255.255.255.0 10.3.1.2 
# On DeviceB, configure transparent transmission of 802.1X packets, create VLANs, and configure allowed VLANs on interfaces.
DeviceB is a Layer 2 switch. To ensure that 802.1X authentication can be performed for users, configure transparent transmission of 802.1X packets on DeviceB.
<HUAWEI> system-view
[HUAWEI] sysname DeviceB
[DeviceB] vlan batch 100 200
[DeviceB] l2protocol-tunnel user-defined-protocol 802.1X protocol-mac 0180-c200-0003 group-mac 0100-0000-0002
[DeviceB] interface eth-trunk 10
[DeviceB-Eth-Trunk10] l2protocol-tunnel user-defined-protocol 802.1X enable
[DeviceB-Eth-Trunk10] port link-type trunk
[DeviceB-Eth-Trunk10] port trunk allow-pass vlan 100 200
[DeviceB-Eth-Trunk10] quit
[DeviceB] interface 10ge 1/0/1
[DeviceB-10GE1/0/1] eth-trunk 10
[DeviceB-10GE1/0/1] quit
[DeviceB] interface 10ge 1/0/2
[DeviceB-10GE1/0/2] eth-trunk 10
[DeviceB-10GE1/0/2] quit
[DeviceB] interface 10ge 1/0/3
[DeviceB-10GE1/0/3] l2protocol-tunnel user-defined-protocol 802.1X enable
[DeviceB-10GE1/0/3] port link-type trunk
[DeviceB-10GE1/0/3] port trunk allow-pass vlan 100 200
[DeviceB-10GE1/0/3] quit
# On DeviceC, configure transparent transmission of 802.1X packets, create VLANs, and configure allowed VLANs on interfaces.
DeviceC is a Layer 2 switch. To ensure that 802.1X authentication can be performed for users, configure transparent transmission of 802.1X packets on DeviceC.
<HUAWEI> system-view
[HUAWEI] sysname DeviceC
[DeviceC] vlan batch 100 200
[DeviceC] l2protocol-tunnel user-defined-protocol 802.1X protocol-mac 0180-c200-0003 group-mac 0100-0000-0002
[DeviceC] interface 10ge 1/0/1
[DeviceC-10GE1/0/1] port link-type trunk
[DeviceC-10GE1/0/1] undo port trunk allow-pass vlan 1
[DeviceC-10GE1/0/1] port trunk pvid vlan 100
[DeviceC-10GE1/0/1] port trunk allow-pass vlan 100
[DeviceC-10GE1/0/1] quit
[DeviceC] interface 10ge 1/0/2
[DeviceC-10GE1/0/2] l2protocol-tunnel user-defined-protocol 802.1X enable
[DeviceC-10GE1/0/2] port link-type access
[DeviceC-10GE1/0/2] port default vlan 200
[DeviceC-10GE1/0/2] quit
[DeviceC] interface 10ge 1/0/3
[DeviceC-10GE1/0/3] l2protocol-tunnel user-defined-protocol 802.1X enable
[DeviceC-10GE1/0/3] port link-type trunk
[DeviceC-10GE1/0/3] port trunk allow-pass vlan 100 200
[DeviceC-10GE1/0/3] quit
Configure VLANIF 100, VLANIF 200, and VLANIF 201 to assign IP addresses to APs, wired users, and wireless users, respectively.
[DeviceA] dhcp enable
[DeviceA] interface vlanif 100
[DeviceA-Vlanif100] dhcp select interface
[DeviceA-Vlanif100] quit
[DeviceA] interface vlanif 200
[DeviceA-Vlanif200] dhcp select interface 
[DeviceA-Vlanif200] quit
[DeviceA] interface vlanif 201
[DeviceA-Vlanif201] dhcp select interface 
[DeviceA-Vlanif201] quit
[DeviceA] dot1x-access-profile name d1
[DeviceA-dot1x-access-profile-d1] dot1x authentication-method eap
[DeviceA-dot1x-access-profile-d1] quit
[DeviceA] web-auth-server abc
[DeviceA-web-auth-server-abc] server-source ip-address 10.3.1.1   
[DeviceA-web-auth-server-abc] server-ip 10.5.1.3
[DeviceA-web-auth-server-abc] port 50200
[DeviceA-web-auth-server-abc] url http://10.5.1.3:8080/portal
[DeviceA-web-auth-server-abc] shared-key cipher YsHsjx_202206
[DeviceA-web-auth-server-abc] quit
[DeviceA] authentication-profile name p1
[DeviceA-authen-profile-p1] mac-access-profile m1    
[DeviceA-authen-profile-p1] dot1x-access-profile d1
[DeviceA] authentication-profile name p2
[DeviceA-authen-profile-p2] mac-access-profile m1    
[DeviceA-authen-profile-p2] portal-access-profile web1
[DeviceA-authen-profile-p2] access-domain isp
# Enable MAC address bypass authentication and configure an authentication domain.
[DeviceA-authen-profile-p1] authentication dot1x-mac-bypass   
[DeviceA-authen-profile-p1] access-domain isp  
[DeviceA-authen-profile-p1] quit
[DeviceA] interface eth-trunk 10
[DeviceA-Eth-Trunk10] authentication-profile p1
[DeviceA-Eth-Trunk10] quit
[DeviceA] free-rule-template name default_free_rule
[DeviceA-free-rule-default_free_rule] free-rule 1 source vlan 100 
[DeviceA-free-rule-default_free_rule] quit
[DeviceA] authentication-profile name p1
[DeviceA-authen-profile-p1] free-rule-template default_free_rule
[DeviceA-authen-profile-p1] quit
[DeviceA] wlan
[DeviceA-wlan] security-profile name wlan-net-security
[DeviceA-wlan-sec-prof-wlan-net-security] security open
[DeviceA-wlan-sec-prof-wlan-net-security] quit
[DeviceA-wlan] ssid-profile name wlan-net-ssid
[DeviceA-wlan-ssid-prof-wlan-net-ssid] ssid wlan-net
[DeviceA-wlan-ssid-prof-wlan-net-ssid] quit
[DeviceA-wlan] vap-profile name wlan-net-vap
[DeviceA-wlan-vap-prof-wlan-net-vap] forward-mode tunnel
[DeviceA-wlan-vap-prof-wlan-net-vap] service-vlan vlan-id 201
[DeviceA-wlan-vap-prof-wlan-net-vap] security-profile wlan-net-security
[DeviceA-wlan-vap-prof-wlan-net-vap] ssid-profile wlan-net-ssid
[DeviceA-wlan-vap-prof-wlan-net-vap] authentication-profile p2
[DeviceA-wlan-vap-prof-wlan-net-vap] quit
[DeviceA-wlan] ap-group name ap-group1
[DeviceA-wlan-ap-group-ap-group1] vap-profile wlan-net-vap wlan 1 radio 0
[DeviceA-wlan-ap-group-ap-group1] vap-profile wlan-net-vap wlan 1 radio 1
[DeviceA-wlan-ap-group-ap-group1] vap-profile wlan-net-vap wlan 1 radio 2
[DeviceA-wlan-ap-group-ap-group1] quit
Users can access the network after being successfully authenticated.
After the users go online, you can run the display access-user command on the device to check online user information.
#
sysname DeviceA
#
vlan batch 100 200 201 300
#
dhcp enable
#
authentication-profile name p1
 mac-access-profile m1
 dot1x-access-profile d1
 authentication dot1x-mac-bypass
 access-domain isp
 free-rule-template default_free_rule
#
authentication-profile name p2
 mac-access-profile m1
 portal-access-profile web1
 access-domain isp
#
free-rule-template name default_free_rule
 free-rule 1 source vlan 100
#
web-auth-server abc
 server-source ip-address 10.3.1.1
 server-ip 10.5.1.3
 port 50200
 url http://10.5.1.3:8080/portal
 shared-key cipher %^%#-;k-</gjVJ=ZK&Ea)<WB(j1FD8HJOGq^@$Ly=\0Y%^%#
#
portal-access-profile name web1
 web-auth-server abc
#
radius-server template rd1
