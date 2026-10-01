---
id: collect-261001-general-networking/general-networking/wireless-design-and-configure-configuration-guides-encryption-and-authentication-0c3f4597-2
title: "wireless-design-and-configure-configuration-guides-encryption-and-authentication-0c3f4597"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-general-networking/wireless-design-and-configure-configuration-guides-encryption-and-authentication-0c3f4597.md
source_anchor: ""
source_lines: [93, 175]
sha256: d3ce03c5a1c696368dd6d8e39f42b6c58359e2a2e0798aa384d191e9e1933c1c
---

# wireless-design-and-configure-configuration-guides-encryption-and-authentication-0c3f4597

Network preparation before enabling Wi-Fi 7
Below is a quick snapshot of how existing SSIDs can be transitioned to ensure Wi-Fi 6E or Wi-Fi 7 compliance.
| Use Case | WLAN Security Today | Migrate to Wi-Fi 7 compliant SSID | 
| Guest Access | Open | OWE | 
| Corporate SSID with Radius Auth | WPA2 | WPA3 (or) WPA3 Transition (Enterprise) | 
| IoT/Guest (PSK based) | WPA2 | WPA3 (or) WPA3 Transiton (Personal) | 
| Corporate Secure | SuiteB 192 Bit | SuiteB 192 Bit | 
For more information you can refer to the migration guide
Configuration Workflow
The WPA3 configuration workflow consists of two main steps:
- 
    Create a WLAN using WPA3-Enterprise, WPA3-Personal, or OWE (Opportunistic Wireless Encryption).
- 
    (Optional) Enable SSIDs for Wi-Fi 7, if the network includes Wi-Fi 7 access points. This step is required due to the additional security requirements introduced with Wi-Fi 7.
Note: Wi-Fi 6E and Wi-Fi 7 introduced MBSSID (Multiple SSID), which allows groups of four SSIDs to be advertised within a single Beacon or Probe Response frame in the 6 GHz band. The dashboard supports four MBSSID groups, each containing four SSIDs. For an access point to advertise a group’s Wi-Fi 7 capability, all SSIDs in that group must be Wi-Fi 7 compliant. If any SSID in the group is not compliant, the group’s capability is reduced to Wi-Fi 6E.
Note: Per-group SSID configuration is available starting in MR 32.1.4. In earlier releases (MR 31.1.x through MR 32.1.3), all SSIDs transmitted by the AP were required to be Wi-Fi 7 compliant. If even one SSID was not security-compliant with Wi-Fi 7, the AP’s capability was restricted to Wi-Fi 6E.
WPA3 Enterprise
WPA3-Enterprise builds upon the foundation of WPA2-Enterprise with the additional requirement of using Protected Management Frames on all WPA3 connections with 802.1X for user authentication with a RADIUS server. By default, WPA3 uses 128-bit encryption, but it also introduces an optionally configurable SuiteB-192 bit cryptographic strength encryption using GMCP-256, which gives additional protection to any network transmitting sensitive data. The WPA3-Enterprise is highly preferred and recommended to be used and commonly seen in enterprises, financial institutions, government, and other market sectors where network security is most critical.
WPA3 Enterprise has three modes of operation available on dashboard to meet the network requirements as needed. They are
1. WPA3 Only
2. WPA3 192-bit security
3. WPA3 Transition Mode
WPA3 Only
To have a WPA3 Enterprise only WLAN, follow these steps:
1. Navigate to Wireless --> Access control -->security
2. Select Enterprise with my Radius server
3. Set the WPA encryption selection as WPA3 Only.
4. If the network has Wi-Fi 7 APs, select GCMP256 in the Advanced WPA3 Settings.
5. Configure the Radius server.
WPA3 192-bit Security
This mode utilizes 192-bit security while still using the 802.1X standard to provide a secure wireless network for enterprise use. This provides a superior encryption method to better protect any kind of data. The security suite is aligned with the recommendations from the Commercial National Security Algorithm (CNSA) suite and is commonly placed in high-security Wi-Fi networks such as in government, defense, finance, and other industries.
WPA3 192-bit security will be exclusive for EAP-TLS, which will require certificates on both the supplicant and RADIUS server. Also, to use WPA3 192-bit enterprise, the RADIUS servers must use one of the permitted EAP ciphers:
- TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384
- TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384
- TLS_DHE_RSA_WITH_AES_256_GCM_SHA384
Note: SuiteB - 192 Bit SSID is compliant with Wi-Fi 7, so no changes are required within the SSID's configuration in Dashboard, but the SSID must be re-configured using higher number SSIDs.
Note: When Wi-Fi 7 is enabled for All SSIDs sub tab it should be noted that SuiteB - 192 Bit SSID will only be supported on SSIDs #13 - #15. This is to ensure that the correct MBSSID grouping is enabled for all types of SSIDs on the network. If a network has SuiteB - 192 Bit SSID configured on a lower number SSID, Dashboard will show an alert stating Wi-Fi 7 cannot be enabled because the following SSIDs do not meet the security requirements.
Note: If SuiteB - 192 Bit SSID is already configured between SSIDs #13 - #15. no changes are needed.
To have only a WPA3 192-bit security only WLAN, follow these steps:
1. Navigate to Wireless --> Access control -->security
2. Select Enterprise with my Radius server
3. Set the WPA encryption selection as WPA3 192-bit security from the drop down menu.
4. Configure the Radius server.
Enterprise WPA3 Transition Mode
WPA3 Transition mode for 802.1X enables clients to connect to a single SSID with dynamic encryption. This is done by using WPA2 for 2.4 GHz and 5 GHz, while using WPA3 for 6 GHz radio. This allows Wi-Fi 5, 6 and 6E clients to connect to the same broadcasting SSID configured for RADIUS-based authentication. With WPA3 Transition Mode, clients can roam between WPA2 enterprise and WPA3 enterprise SSIDs. When a client roams from a WPA2 to WPA3 SSID reauthentication will take place, but with minimal disruption to connectivity.
Note: This feature is available from MR 31.1.x and above firmware versions.
To have only a WPA3 Transition security WLAN, follow these steps:
1. Navigate to Wireless --> Access control--> security
2. Select Enterprise with my Radius server
3. Set the WPA encryption selection as WPA3 Transition security from the drop down menu.
4. If the network has Wi-Fi 7 APs, select GCMP256 in the Advanced WPA3 Settings.
5. Configure the Radius server.
Note: It’s recommended to set 802.11w to Enabled, which will allow WPA2 clients that do not have support for PMF to associate.
Note: GCMP256 cipher is needed to be enabled on the Access Point for Wi-Fi 7, but, in reality many clients in the market today can do a Wi-Fi 7 features like MLO, 11be rates with AES (CCMP128).
Note: With MR 31.1.7, Wi-Fi 6E Access Points configured with WPA3 Transition mode will enable the 6 GHz radio.
WPA3-Personal
WPA3-Personal uses 128-bit cryptographic-strength encryption with a password-based authentication method through SAE for user authentication purposes. In addition, unlike WPA2-Personal, WPA3-Personal heightens network security against offline dictionary attacks by limiting password guesses and requiring users to interact with a live network every time they do so. This requirement makes hacking into a network much more time-consuming and dissuades attempts at a brute force attack.
WPA3-Personal provides the following key advantages:
● Creates a shared secret that is different for each SAE authentication
● Protects against brute force “dictionary” attacks and passive attacks
● Provides forward secrecy
WPA3 Personal has two modes of operation available on dashboard to meet the network requirements as needed. They are
1. WPA3 Only
2. WPA3 Transition Mode
WPA3 Only
To have a Personal WPA3 only WLAN, follow these steps:
1. Navigate to Wireless > Configure > Access control > Security
2. Select Password option and enter the password.
3. Set the WPA encryption selection as WPA3 Only.
4. If the network has Wi-Fi 7 APs, expand the Advanced WPA3 settings for cipher and AKM
a. GCMP256 as Cipher
b. SAE and SAE-EXT-KEY as the AKM.
5. If the network has Wi-Fi 6/6E APs and no Wi-Fi 7 APs, expand the Advanced WPA3 settings for cipher and AKM
a. Select SAE as the AKM.
Note: Wi-Fi 6/6E standard does not enforce GCMP256 and SAE-EXT.
Personal WPA3 Transition Mode
To have Personal WPA3 Transition Mode WLAN, follow these steps:
1. Navigate to Wireless > Access control > security
2. Select Password option and enter the password.
3. Set the WPA encryption selection as WPA3 Transition Mode.
4. If the network has Wi-Fi 7 APs, expand the Advanced WPA3 settings for cipher and AKM
a. GCMP256 as Cipher
b. SAE and SAE-EXT-KEY as the AKM.
