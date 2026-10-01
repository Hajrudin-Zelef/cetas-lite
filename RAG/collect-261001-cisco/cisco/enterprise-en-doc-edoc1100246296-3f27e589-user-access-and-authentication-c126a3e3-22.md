---
id: collect-261001-cisco/cisco/enterprise-en-doc-edoc1100246296-3f27e589-user-access-and-authentication-c126a3e3-22
title: "Configure DeviceA to generate a local key pair."
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-cisco/enterprise-en-doc-edoc1100246296-3f27e589-user-access-and-authentication-c126a3e3.md
source_anchor: ""
source_lines: [3750, 3955]
sha256: c7e99851c32ec58bae3d218d5f4532bea06fa72706d9b1449d37b7616afa93be
---

# Configure DeviceA to generate a local key pair.

[DeviceA-Vlanif103] ip address 192.168.103.1 24
[DeviceA-Vlanif103] quit
[DeviceA] ip route-static 192.168.100.0 255.255.255.0 192.168.103.2
[DeviceA] pki realm abcd  
[DeviceA-pki-abcd] quit  
[DeviceA] ssl policy huawei 
[DeviceA-ssl-policy-huawei] pki-domain abcd
Warning: A preset certificate is loaded to the specified PKI domain. The current operation has security risks. Continue? [Y/N]: y
[DeviceA-ssl-policy-huawei] quit
[DeviceA] haca-server template HACA 
[DeviceA-haca-HACA] haca-server server-address 192.168.100.100 50301 ssl-policy huawei 
[DeviceA-haca-HACA] haca enable
[DeviceA-haca-HACA] quit
[DeviceA] url-template name urlTemplate_0
[DeviceA-url-template-urlTemplate_0] url https://192.168.100.100:19008/portal
[DeviceA-url-template-urlTemplate_0] url-parameter device-mac wac-mac redirect-url redirect-url ssid ssid user-ipaddress uaddress user-mac umac
[DeviceA-url-template-urlTemplate_0] quit
[DeviceA] web-auth-server Portal
[DeviceA-web-auth-server-Portal] server-ip 192.168.100.100 
[DeviceA-web-auth-server-Portal] port 50100 
[DeviceA-web-auth-server-Portal] url-template urlTemplate_0 
[DeviceA-web-auth-server-Portal] protocol haca
[DeviceA-web-auth-server-Portal] quit
[DeviceA] portal-access-profile name wlan-net-auth 
[DeviceA-portal-access-profile-wlan-net-auth] web-auth-server Portal
[DeviceA] aaa
[DeviceA-aaa] authentication-scheme wlan-net-auth
[DeviceA-aaa-authen-wlan-net-auth] authentication-mode haca
[DeviceA-aaa-authen-wlan-net-auth] quit
# Set the accounting mode to non-accounting.
[DeviceA-aaa] accounting-scheme wlan-net-auth 
[DeviceA-aaa-accounting-wlan-net-auth] accounting-mode none
[DeviceA-aaa-accounting-wlan-net-auth] quit
[DeviceA-aaa] domain wlan-net-auth
[DeviceA-aaa-domain-wlan-net-auth] authentication-scheme wlan-net-auth
[DeviceA-aaa-domain-wlan-net-auth] accounting-scheme wlan-net-auth
[DeviceA-aaa-domain-wlan-net-auth] haca-server HACA
[DeviceA-aaa-domain-wlan-net-auth] quit
[DeviceA-aaa] quit
[DeviceA]mac-access-profile name wlan-net-auth
[DeviceA-mac-access-profile-wlan-net-auth] quit
[DeviceA] authentication-profile name wlan-net-auth
[DeviceA-authen-profile-wlan-net-auth] mac-access-profile wlan-net-auth 
[DeviceA-authen-profile-wlan-net-auth] portal-access-profile wlan-net-auth 
[DeviceA-authen-profile-wlan-net-auth] free-rule-template default_free_rule 
[DeviceA-authen-profile-wlan-net-auth] authentication timer re-authen pre-authen 0 wlan-user 
[DeviceA-authen-profile-wlan-net-auth] access-domain wlan-net-auth
[DeviceA-authen-profile-wlan-net-auth] quit
# Create a regulatory domain profile, configure the WAC country code in the profile, and bind the profile to the created AP group.
[DeviceA] interface vlanif 100
[DeviceA-Vlanif100] ip address 10.23.100.1 24
[DeviceA-Vlanif100] quit
[DeviceA] capwap dtls no-auth enable
Warning: This operation allows for device access in non-DTLS encryption mode even when DTLS is enabled and brings security risks. Af
ter the device goes online for the first time, disable this function to prevent security risks. Continue? [Y/N]:y
[DeviceA] capwap source interface vlanif 100
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
[DeviceA] wlan
[DeviceA-wlan-view] security-profile name wlan-net-security
[DeviceA-wlan-sec-prof-wlan-net-security] security open
[DeviceA-wlan-sec-prof-wlan-net-security] quit
[DeviceA-wlan-view] ssid-profile name wlan-net-ssid
[DeviceA-wlan-ssid-prof-wlan-net-ssid] ssid wlan-net
[DeviceA-wlan-ssid-prof-wlan-net-ssid] quit
[DeviceA-wlan-view] vap-profile name wlan-net-vap
[DeviceA-wlan-vap-prof-wlan-net-vap] forward-mode tunnel
[DeviceA-wlan-vap-prof-wlan-net-vap] service-vlan vlan-id 101
[DeviceA-wlan-vap-prof-wlan-net-vap] security-profile wlan-net-security
[DeviceA-wlan-vap-prof-wlan-net-vap] ssid-profile wlan-net-ssid
[DeviceA-wlan-vap-prof-wlan-net-vap] authentication-profile wlan-net-auth
[DeviceA-wlan-vap-prof-wlan-net-vap] quit
[DeviceA-wlan-view] ap-group name ap-group1
[DeviceA-wlan-ap-group-ap-group1] vap-profile wlan-net-vap wlan 1 radio 0
[DeviceA-wlan-ap-group-ap-group1] vap-profile wlan-net-vap wlan 1 radio 1
[DeviceA-wlan-ap-group-ap-group1] quit
#
sysname DeviceB
#
vlan batch 102
#
interface 10GE1/0/1
 description to_AP1
 port link-type trunk
 port trunk pvid vlan 102
 undo port trunk allow-pass vlan 1
 port trunk allow-pass vlan 102
#
interface 10GE1/0/2
 description to_AP2
 port link-type trunk
 port trunk pvid vlan 102
 undo port trunk allow-pass vlan 1
 port trunk allow-pass vlan 102
#
interface 10GE1/0/3
 description to_WAC
 port link-type trunk
 undo port trunk allow-pass vlan 1
 port trunk allow-pass vlan 102
#
return
#
pki realm abcd
#
sysname DeviceA
#
ssl policy huawei
 pki-domain abcd
#
vlan batch 101 to 103
#
free-rule-template name default_free_rule 
 free-rule 1 destination ip 10.23.202.1 mask 255.255.255.255
#
authentication-profile name wlan-net-auth
 mac-access-profile wlan-net-auth 
 portal-access-profile wlan-net-auth 
 free-rule-template default_free_rule 
 authentication timer re-authen pre-authen 0 wlan-user 
 access-domain wlan-net-auth
#
mac-access-profile name wlan-net-auth
#
haca-server template HACA 
 haca-server server-address 192.168.100.100 50301 ssl-policy huawei
 haca enable
#
url-template name urlTemplate_0
 url https://192.168.100.100:19008/portal
 url-parameter device-mac wac-mac redirect-url redirect-url ssid ssid user-ipaddress uaddress user-mac umac
#
web-auth-server Portal
 server-ip 192.168.100.100 
 port 50100 
 url-template urlTemplate_0 
 protocol haca
#
portal-access-profile name wlan-net-auth 
 web-auth-server Portal
#
aaa
 authentication-scheme wlan-net-auth
  authentication-mode haca
 accounting-scheme wlan-net-auth 
  accounting-mode none
 domain wlan-net-auth
  authentication-scheme wlan-net-auth
  accounting-scheme wlan-net-auth
  haca-server HACA
#
dhcp enable
#
interface Vlanif101
 ip address 192.168.101.1 255.255.255.0
 dhcp select interface
#
interface Vlanif102
 ip address 192.168.102.1 255.255.255.0
 dhcp select interface
#
interface Vlanif103
 ip address 192.168.103.1 255.255.255.0
#
interface 10GE1/0/1
 description to_DeviceB
 port link-type trunk
 undo port trunk allow-pass vlan 1
 port trunk allow-pass vlan 102
#
interface 10GE1/0/2
 description to_HACA_Server
 port link-type trunk
 undo port trunk allow-pass vlan 1
 port trunk allow-pass vlan 103
#
capwap source interface vlanif 102
capwap dtls psk %+%##!!!!!!!!!"!!!!"!!!!*!!!!a4O4$.&PxRTGtHR(VaB4*(KhWT"C@K+HZ%@!!!!!2jp5!!!!!!;!!!!..AjEF}u<=!%e=F%Dz-Y[R--3yrC-*MVXf<!!!!!%+%#
#
wlan
 security-profile name wlan-net-security
  security open
 ssid-profile name wlan-net-ssid
  ssid wlan-net
 vap-profile name wlan-net-vap
  forward-mode tunnel
  service-vlan vlan-id 101
  ssid-profile wlan-net-ssid
  security-profile wlan-net-security
  authentication-profile wlan-net-auth
 regulatory-domain-profile name domain1
 ap-group name ap-group1
  regulatory-domain-profile domain1
  radio 0
   vap-profile wlan-net-vap wlan 1
  radio 1
