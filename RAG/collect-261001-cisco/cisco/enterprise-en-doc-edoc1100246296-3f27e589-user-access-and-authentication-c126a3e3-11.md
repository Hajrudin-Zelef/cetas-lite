---
id: collect-261001-cisco/cisco/enterprise-en-doc-edoc1100246296-3f27e589-user-access-and-authentication-c126a3e3-11
title: "Configure DeviceA to generate a local key pair."
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-cisco/enterprise-en-doc-edoc1100246296-3f27e589-user-access-and-authentication-c126a3e3.md
source_anchor: ""
source_lines: [1611, 1743]
sha256: 503c48d22972ef6e1b262eace0068b55e37b8f2867ea2747e400e6e491d5edad
---

# Configure DeviceA to generate a local key pair.

[DeviceA-Vlanif101] ip address 10.23.101.1 24
[DeviceA-Vlanif101] dhcp select interface
[DeviceA-Vlanif101] quit
[DeviceA] ip route-static 10.23.200.0 255.255.255.0 10.23.101.2
# Create an AP group to which the APs with the same configuration can be added.
[DeviceA] wlan
[DeviceA-wlan] ap-group name ap-group1
[DeviceA-wlan-ap-group-ap-group1] quit
# Create a regulatory domain profile, configure the WAC country code in the profile, and bind the profile to the AP group.
[DeviceA-wlan] regulatory-domain-profile name domain1
[DeviceA-wlan-regulate-domain-domain1] country-code cn
Warning: Modifying the country code will clear the channel and power configurations of radios, and requires the APs to be restarted if they run V200R019C10 or earlier. Continue? [Y/N]:y
[DeviceA-wlan-regulate-domain-domain1] quit
[DeviceA-wlan] ap-group name ap-group1
[DeviceA-wlan-ap-group-ap-group1] regulatory-domain-profile domain1
Warning: This configuration change will clear the channel and power configurations of radios, and may restart APs. Continue?[Y/N]:y
[DeviceA-wlan-ap-group-ap-group1] quit
[DeviceA-wlan] quit
# Configure the WAC's source interface.
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
After the peer device goes online, run the undo capwap dtls no-auth enable command to disable the function of establishing CAPWAP DTLS sessions in none authentication mode, thereby preventing unauthorized devices from accessing the network.
The default AP authentication mode is MAC address authentication. If the default settings have not been changed, you do not need to run the ap auth-mode mac-auth command.
[DeviceA] wlan
[DeviceA-wlan] ap auth-mode mac-auth
[DeviceA-wlan] ap-id 0 ap-mac 00e0-fc12-3456
[DeviceA-wlan-ap-0] ap-name area_1
[DeviceA-wlan-ap-0] ap-group ap-group1
Warning: This operation may cause AP reset. If the country code changes, it will clear channel, power and antenna gain configuration s of the radio, Whether to continue? [Y/N]:y
[DeviceA-wlan-ap-0] quit
[DeviceA-wlan] quit
# After the AP goes online, disable the function of establishing CAPWAP DTLS sessions in non-authentication mode.
[DeviceA] undo capwap dtls no-auth enable
Ensure that the RADIUS server address, port number, and shared key are correctly configured and are the same as those on the RADIUS server.
# Configure a RADIUS server template.
[DeviceA] radius-server template radius_huawei
[DeviceA-radius-radius_huawei] radius-server authentication 10.23.200.1 1812
[DeviceA-radius-radius_huawei] radius-server accounting 10.23.200.1 1813
[DeviceA-radius-radius_huawei] radius-server shared-key cipher YsHsjx_202206mc@1
[DeviceA-radius-radius_huawei] quit
# Configure an authentication scheme that uses RADIUS authentication.
[DeviceA] aaa
[DeviceA-aaa] authentication-scheme scheme1
[DeviceA-aaa-authen-scheme1] authentication-mode radius
[DeviceA-aaa-authen-scheme1] quit
# Configure an accounting scheme that uses RADIUS accounting.
# Create the authentication domain example.com and bind an authentication scheme, an accounting scheme, and a RADIUS server template to the authentication domain.
[DeviceA-aaa] domain example.com
[DeviceA-aaa-domain-example.com] authentication-scheme scheme1
[DeviceA-aaa-domain-example.com] accounting-scheme scheme2
[DeviceA-aaa-domain-example.com] radius-server radius_huawei
[DeviceA-aaa-domain-example.com] quit
[DeviceA-aaa] quit
[DeviceA] authentication-profile name p1
[DeviceA-authentication-profile-p1] mac-access-profile m1
[DeviceA-authentication-profile-p1] access-domain example.com force
[DeviceA-authentication-profile-p1] quit
# Create the security profile wlan-security and configure a security policy in the profile.
[DeviceA] wlan
[DeviceA-wlan] security-profile name wlan-security
[DeviceA-wlan-sec-prof-wlan-security] security open
[DeviceA-wlan-sec-prof-wlan-security] quit
# Create the SSID profile wlan-ssid and set the SSID name to wlan-net.
[DeviceA-wlan] ssid-profile name wlan-ssid
[DeviceA-wlan-ssid-prof-wlan-ssid] ssid wlan-net
[DeviceA-wlan-ssid-prof-wlan-ssid] quit
# Create the VAP profile wlan-vap, configure the service data forwarding mode and service VLAN, and bind the security profile, SSID profile, and authentication profile to the VAP profile.
[DeviceA-wlan] vap-profile name wlan-vap
[DeviceA-wlan-vap-prof-wlan-vap] forward-mode tunnel
[DeviceA-wlan-vap-prof-wlan-vap] service-vlan vlan-id 101
[DeviceA-wlan-vap-prof-wlan-vap] security-profile wlan-security
[DeviceA-wlan-vap-prof-wlan-vap] ssid-profile wlan-ssid
[DeviceA-wlan-vap-prof-wlan-vap] authentication-profile p1
[DeviceA-wlan-vap-prof-wlan-vap] quit
# Bind the VAP profile wlan-vap to the AP group and apply the profile to radios 0 and 1 of the APs.
[DeviceA-wlan] ap-group name ap-group1
[DeviceA-wlan-ap-group-ap-group1] vap-profile wlan-vap wlan 1 radio 0
[DeviceA-wlan-ap-group-ap-group1] vap-profile wlan-vap wlan 1 radio 1
[DeviceA-wlan-ap-group-ap-group1] quit
The automatic channel and power calibration functions are enabled by default. The manual channel and power configurations take effect only when the two functions are disabled. The channel and power configurations for the AP radios in this example are for reference only. In practice, configure the channel and power of AP radios based on the actual country codes of APs and network planning.
[DeviceA-wlan] ap-id 0
[DeviceA-wlan-ap-0] radio 0
[DeviceA-wlan-ap-0-radio-0] calibrate auto-channel-select disable
[DeviceA-wlan-ap-0-radio-0] calibrate auto-txpower-select disable
[DeviceA-wlan-ap-0-radio-0] channel 20mhz 6
Warning: This action may cause service interruption. Continue?[Y/N]y 
[DeviceA-wlan-ap-0-radio-0] eirp 127
[DeviceA-wlan-ap-0-radio-0] quit
[DeviceA-wlan-ap-0] radio 1
[DeviceA-wlan-ap-0-radio-1] calibrate auto-channel-select disable
[DeviceA-wlan-ap-0-radio-1] calibrate auto-txpower-select disable
[DeviceA-wlan-ap-0-radio-1] channel 80mhz 149
Warning: This action may cause service interruption. Continue?[Y/N]y 
[DeviceA-wlan-ap-0-radio-1] eirp 127
[DeviceA-wlan-ap-0-radio-1] quit
[DeviceA-wlan-ap-0] quit
[DeviceA-wlan] quit
#
sysname DeviceA
#
vlan batch 100 to 101
#
authentication-profile name p1
 mac-access-profile m1
 access-domain example.com force
#
mac-access-profile name m1
#
dhcp enable
#
radius-server template radius_huawei
 radius-server shared-key cipher %+%##!!!!!!!!!"!!!!"!!!!*!!!!Cd/`W03KjAwAqn64E<\TxGC_SOri<2BP\A+!!!!!2jp5!!!!!!B!!!!Oe2HMc->XMa#TLDZUaJBFJtm#XVj*E:S*|(N7`J1B:3QY!!!!!!!!!!!!!!!%+%# 
 radius-server authentication 10.23.200.1 1812 weight 80
 radius-server accounting 10.23.200.1 1813 weight 80
#
aaa
 authentication-scheme scheme1
  authentication-mode radius
 accounting-scheme scheme2
