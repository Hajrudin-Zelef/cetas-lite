---
id: collect-261001-cisco/cisco/enterprise-en-doc-edoc1100246296-3f27e589-user-access-and-authentication-c126a3e3-29
title: "Configure DeviceA to generate a local key pair."
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-cisco/enterprise-en-doc-edoc1100246296-3f27e589-user-access-and-authentication-c126a3e3.md
source_anchor: ""
source_lines: [5110, 5236]
sha256: 8570e3e3b66fcafe3d840165682292a92433da3c0492cfb90fd706f168938253
---

# Configure DeviceA to generate a local key pair.

[Switch_A-authen-radius_huawei] authentication-mode radius
[Switch_A-authen-radius_huawei] quit
[Switch_A-aaa] accounting-scheme scheme1
[Switch_A-aaa-accounting-scheme1] accounting-mode radius
[Switch_A-aaa-accounting-scheme1] accounting realtime 15
[Switch_A-aaa-accounting-scheme1] quit
In this example, the device is connected to Agile Controller-Campus. The accounting function is not provided for accounting purposes, and is only used to maintain terminal online information through accounting packets.
The accounting realtime command is used to set the real-time accounting interval. A shorter real-time accounting interval requires higher performance of the device and RADIUS server. The real-time accounting interval needs to be set based on the user quantity.
| User Quantity | Real-Time Accounting Interval | 
|---|---|
| 1 to 99 | 3 minutes | 
| 100 to 499 | 6 minutes | 
| 500 to 999 | 12 minutes | 
| ≥ 1000 | ≥ 15 minutes | 
# Create the authentication domain huawei and bind the RADIUS authentication scheme radius_huawei, accounting scheme scheme1, and RADIUS server template radius_huawei to the domain.
[Switch_A-aaa] domain huawei
[Switch_A-aaa-domain-huawei] authentication-scheme radius_huawei
[Switch_A-aaa-domain-huawei] accounting-scheme scheme1
[Switch_A-aaa-domain-huawei] radius-server radius_huawei
[Switch_A-aaa-domain-huawei] quit
[Switch_A-aaa] quit
# Configure the 802.1X access profile d1.
By default, an 802.1X access profile uses EAP authentication. Ensure that the RADIUS server supports the EAP protocol; otherwise, it cannot process 802.1X authentication requests.
[Switch_A] dot1x-access-profile name d1
[Switch_A-dot1x-access-profile-d1] quit
# Configure the authentication profile p1, bind the 802.1X access profile d1 to the authentication profile, and configure the forcible authentication domain huawei for users using the authentication profile.
[Switch_A] authentication-profile name p1
[Switch_A-authen-profile-p1] dot1x-access-profile d1
[Switch_A-authen-profile-p1] access-domain huawei force
[Switch_A-authen-profile-p1] quit
# Bind the authentication profile p1 to 10GE 1/0/1 and enable 802.1X authentication on the interface.
[Switch_A] interface 10ge 1/0/1
[Switch_A-10GE1/0/1] authentication-profile p1
[Switch_A-10GE1/0/1] quit
# Check whether a user can pass RADIUS authentication. (The test user test and password YsHsjx_202206 have been configured on the RADIUS server.)
[Switch_A] test-aaa test YsHsjx_202206 radius-template radius_huawei 
Info: Account test succeeded.
# Create the 802.1X client profile huawei, enter the 802.1X client profile view, and set the 802.1X authentication mode to EAP-PEAP.
<HUAWEI> system-view
[HUAWEI] sysname WAC
[WAC] dot1x-client-profile name huawei
[WAC-dot1x-client-profile-huawei] eap-method eap-peap username huawei password cipher YsHsjx_202206
[WAC-dot1x-client-profile-huawei] quit
# Bind the created 802.1X client profile to an AP wired port profile so that the 802.1X client profile can take effect. Create the wired port profile wired-port1 and bind the 802.1X client profile huawei to it.
[WAC] wlan
[WAC-wlan-view] wired-port-profile name wired-port1
[WAC-wlan-wired-port-wired-port1] dot1x-client-profile huawei
[WAC-wlan-wired-port-wired-port1] quit
# Create an AP group to which the APs with the same configuration can be added. Bind the AP wired port profile wired-port1 to MultiGE0 of the AP in the AP group ap-group1.
[WAC-wlan-view] ap-group name ap-group1
[WAC-wlan-ap-group-ap-group1] wired-port-profile wired-port1 MultiGE 0
[WAC-wlan-ap-group-ap-group1] quit
# Configure the 802.1X client function on the AP.
MDCLI> diagnose wlan-nsc set-dot1x-client interface MultiGE0/0/0 peap-user-name huawei peap-password
Enter password:
Confirm password:
# Check the 802.1X client configuration on the AP.
MDCLI> diagnose
MDCLI> dot1x-client
MDCLI> query-dot1x-client-configuration profile-name multige0
  Profile name                            : multige0
  EAP Method                              : PEAP
  EAP Username                            : huawei
  EAP Password                            : ******
  SSL Policy Name                         : multige0
In this example, tunnel forwarding is used to forward service data. If direct forwarding is used, you are advised to configure port isolation on 10GE 0/0/1 that connects the AC to APs. If port isolation is not configured, unnecessary broadcast packets will be transmitted in VLANs or WLAN users on different APs will be able to communicate with each other at Layer 2.
[WAC] vlan batch 100 101 200
[WAC] interface 10ge 1/0/1
[WAC-10GE1/0/1] port link-type trunk
[WAC-10GE1/0/1] port trunk allow-pass vlan 100
[WAC-10GE1/0/1] quit
# Add the WAC's uplink interface 10GE 1/0/2 to VLAN 200.
[WAC] interface 10ge 1/0/2
[WAC-10GE1/0/2] port link-type trunk
[WAC-10GE1/0/2] port trunk allow-pass vlan 200
[WAC-10GE1/0/2] quit
[AC] interface vlanif 200
[AC-Vlanif200] ip address 10.23.200.2 24
[AC-Vlanif200] quit
# Configure the WAC as a DHCP server to assign IP addresses to APs from the IP address pool on VLANIF 100, and to assign IP addresses to STAs from the IP address pool on VLANIF 101.
[WAC] dhcp enable
[WAC] interface vlanif 100
[WAC-Vlanif100] ip address 10.23.100.1 24
[WAC-Vlanif100] dhcp select interface
[WAC-Vlanif100] dhcp server excluded-ip-address 10.23.100.2
[WAC-Vlanif100] quit
[WAC] interface vlanif 101
[WAC-Vlanif101] ip address 10.23.101.1 24
[WAC-Vlanif101] dhcp select interface
[WAC-Vlanif101] quit
[WAC] wlan
[WAC-wlan] regulatory-domain-profile name domain1
[WAC-wlan-regulate-domain-domain1] country-code cn
Warning: Modifying the country code will clear the channel and power configurations of radios, and requires the APs to be restarted if they run V200R019C10 or earlier. Continue? [Y/N]:y
[WAC-wlan-regulate-domain-domain1] quit
[WAC-wlan] ap-group name ap-group1
[WAC-wlan-ap-group-ap-group1] regulatory-domain-profile domain1
Warning: This configuration change will clear the channel and power configurations of radios, and may restart APs. Continue?[Y/N]:y
[WAC-wlan-ap-group-ap-group1] quit
[WAC-wlan] quit
[WAC] capwap dtls no-auth enable
Warning: This operation allows for device access in non-DTLS encryption mode even when DTLS is enabled and brings security risks. Af
ter the device goes online for the first time, disable this function to prevent security risks. Continue? [Y/N]:y
[WAC] capwap source interface vlanif 100
Set the DTLS PSK(contains 8-32 plain-text characters, or 128 or 148 cipher-text characters that must be a combination of at least tw
o of the following: lowercase letters a to z, uppercase letters A to Z, digits, and special characters):********
Confirm PSK:********
Set the user name for FIT APs(The value is a string of 4 to 31 characters, which can contain letters, underscores, and digits, and m
ust start with a letter):********
Set the password for FIT APs(plain-text password of 8-128 characters or cipher-text password of 128-268 characters that must be a co
mbination of at least three of the following: lowercase letters a to z, uppercase letters A to Z, digits, and special characters):******
Confirm password:********
Set the PSK of the global offline management VAP(plain-text password of 8-63 characters or cipher-text password of 128-188 character
s that must be a combination of at least two of the following: lowercase letters a to z, uppercase letters A to Z, digits, and speci
al characters):********
Confirm PSK:********
[WAC] wlan
[WAC-wlan] ap auth-mode mac-auth
[WAC-wlan] ap-id 0 ap-mac 00e0-fc12-3456
[WAC-wlan-ap-0] ap-name area_1
[WAC-wlan-ap-0] ap-group ap-group1
Warning: This operation may cause AP reset. If the country code changes, it will clear channel, power and antenna gain configuration s of the radio, Whether to continue? [Y/N]:y
[WAC-wlan-ap-0] quit
[WAC-wlan] quit
[WAC] undo capwap dtls no-auth enable
[WAC] radius-server template radius_huawei
[WAC-radius-radius_huawei] radius-server authentication 10.23.200.1 1812
