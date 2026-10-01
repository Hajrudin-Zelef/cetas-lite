---
id: collect-261001-general-networking/general-networking/wireless-design-and-configure-configuration-guides-encryption-and-authentication-0c3f4597-3
title: "wireless-design-and-configure-configuration-guides-encryption-and-authentication-0c3f4597"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-general-networking/wireless-design-and-configure-configuration-guides-encryption-and-authentication-0c3f4597.md
source_anchor: ""
source_lines: [176, 229]
sha256: 4c066e9a65f74602c94a2c8e1e2823411efcbf90919f685d6025b0e96def4042
---

# wireless-design-and-configure-configuration-guides-encryption-and-authentication-0c3f4597

5. If the network has Wi-Fi 6/6E APs and no Wi-Fi 7 APs, then SAE as AKM alone would suffice. No need of GCMP-256 as the Cipher and SAE-EXT as the AKM.
a. Select SAE as the AKM.
Note: It’s recommended to set 802.11w to Enabled, which will allow WPA2 clients that do not have support for PMF to associate.
Enhanced Open (OWE - Opportunistic Wireless Encryption)
OWE is a security method paired with an open-security wireless network to provide it with encryption to protect the network from eavesdroppers. With OWE, the client and AP perform a Diffie-Hellman key exchange during the endpoint association packet exchange and use the resulting PMK to conduct the 4-way handshake. Being associated with open-security wireless networks, OWE can be used with regular open networks as well as those associated with captive portals.
OWE has two modes of operation available on dashboard to meet the network requirements as needed. They are
1. OWE
2. OWE Transition Mode
Note: OWE Transition Mode requires the network to be running MR 32+
The WPA3 Spec v3.4, Section 11.3, states that “The AP's BSS Configuration shall not allow Wi-Fi Enhanced Open Transition Mode (i.e., where the OWETransition Mode element is included in Beacons and Probe responses)” Hence OWE Transition is not valid with 6 GHz and Wi-Fi 7
OWE
To have a OWE only WLAN, follow these steps
1. Navigate to Wireless-->Configure -->Access Control -->Security
2. Select Opportunistic Wireless Encryption (OWE)
3. Select the WPA Encryption as WPA3 only.
4. If the network has Wi-Fi 7 APs, expand the Advanced WPA3 settings for cipher and AKM
a. GCMP256 as Cipher
OWE Transition Mode
The Opportunistic Wireless Encryption (OWE) transition mode enables OWE and non-OWE STAs to connect to the same SSID simultaneously. OWE transition mode enables a seamless transition from Open unencrypted WLANs to OWE WLANs without impacting the wireless connection.
- 
    Both the open WLAN and the OWE WLAN transmit beacon frames. Beacon and probe response frames from the OWE WLAN include the Wi-Fi Alliance vendor IE to encapsulate the BSSID and SSID of the open WLAN, and similarly, the open WLAN also includes for OWE WLAN.
- 
    An OWE STA only displays to the user in the list of available networks the SSID of the Open BSS of an OWE AP operating in OWE Transition Mode, and hides the OWE BSSID of that OWE AP.
- 
    WPA3-capable clients associate with the OWE SSID via the Open SSID.
- 
    Non-WPA3 clients associate directly with the Open SSID.
To have a OWE Transition WLAN, follow these steps.
1. Navigate to Wireless --> Configure--> Acces Control -->Security.
2. First, create a Open WLAN.
3. Create a OWE WLAN.
4. Set the WPA encryption selection as WPA3 Transition Mode.
5. Point the Open WLAN.
Limitations for OWE Transiton Mode
OWE Transition mode is not supported by Wi-Fi 7. OWE transition relies on a pair of SSIDs where one is an Open SSID. Open SSIDs are not compliant with Wi-Fi 7. This is not the case with PSK Transition and Enterprise Transition SSIDs.
|  | 2.4/5 GHz | 6 GHz | 
| Wi-Fi 7 | OWE Only | OWE Only | 
| Wi-Fi 6/6E | Open, OWE Transition and OWE | OWE Only | 
802.11be SSID Group Configuration
The Per SSID toggle for 802.11be (Wi-Fi 7) feature allows WiFi 7 access points to operate some SSIDs in 802.11ax mode while others operate in 802.11be mode. This is also called "ax/be mixed mode", "per-SSID 11be".
If the network has Wi-Fi 7 APs, then 802.11be has to be enabled per group. As stated earlier, all the SSIDs in a group of four has to be security compliant with Wi-Fi 7 requirements. The dashboard has four SSID groups.
Group 1 – SSID 1 to 4
Group 2 – SSID 5 to 8
Group 3 – SSID 9 to 12
Group 4 – SSID 13 to 15
Hence, the users are required to re-arrange the SSIDs that complies to the security within a group. As an example, all SSIDs that comply to Wi-Fi 7 requirement can be in Group 1, the SSIDs like WPA2 or Open that do not comply to Wi-Fi 7 can be in Group 2 and so on.
To enable 11be Per SSID Group,
- Navigate to Wireless > Configuration > Radio Settings > RF Profiles
- Edit the profile that is of interest.
- Navigate to 802.11be section in the General Tab.
- Toggle to Per SSID Group.
- Ensure all SSIDs within a group are Wi-Fi 7 compliant and the 80.11be knob is enabled.
- For 8011.be SSID groups that has non-compliant the SSIDs will be disabled.
Note: Please ensure proper security settings in a group. If security settings are changed after enabling, it impacts the entire MBSSID group. The MBSSID will be recomputed after the config change.
