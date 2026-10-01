---
id: collect-261001-general-networking/general-networking/wireless-design-and-configure-configuration-guides-encryption-and-authentication-0c3f4597-1
title: "wireless-design-and-configure-configuration-guides-encryption-and-authentication-0c3f4597"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-general-networking/wireless-design-and-configure-configuration-guides-encryption-and-authentication-0c3f4597.md
source_anchor: ""
source_lines: [1, 92]
sha256: ae39402291d39c81d1ace86238867c8624dfc38f53ab1f20850a39305b71b0a5
---

# wireless-design-and-configure-configuration-guides-encryption-and-authentication-0c3f4597

WPA3 Encryption and Configuration Guide
Click 日本語 for Japanese
WPA3 Overview
WPA3 is the latest and third iteration of the Wi-Fi Protected Access (WPA) standard, developed by the Wi-Fi Alliance, and serves as the successor to WPA2. The WPA standard was originally created by the Wi-Fi Alliance Security Technical Task Group, chaired by Cisco’s Stephen Orr, with the goal of standardizing wireless security.
WPA3 introduces advanced features for enterprise, personal, and open networks by enhancing cryptographic strength and enabling a more secure authentication process across all WPA3-supported devices.
It is designed to:
- Strengthen wireless security
- Simplify secure connectivity for users
- Provide strong protection even with weak passwords
Enhance security for public and open Wi-Fi networks
The WPA3 Enterprise form extends the solid foundation provided by WPA2 Enterprise by making it mandatory to use Protected Management Frames (PMF) on all connections. This security feature protects against such dangerous attacks as Denial of Service (DoS), honeypots, and eavesdropping.
Supported WPA3 Modes
● WPA3-Enterprise, for 802.1X security networks. This leverages IEEE 802.1X with SHA-256 as the Authentication and Key Management (AKM).
● WPA3-Personal, which uses the Simultaneous Authentication of Equals (SAE) method for personal security networks. There are two sub-methods to derive the Password Element in SAE:
o Hunting and Pecking (HnP)
o Hash-to-Element (H2E)
Note: Wi-Fi 6E (6 GHz) and Wi-Fi 7 requires Hash-to-Element as mandatory, as HnP is prone to brute force dictionary attacks.
● WPA3 Transition Mode (WPA2+WPA3 security-based WLANs for both personal and enterprise).
● Opportunistic Wireless Encryption (OWE) for open security networks.
WPA3 Requirements for 6 GHz operations and Wi-Fi 7
Wi-Fi Alliance mandated WPA3 for 6 GHz band and Wi-Fi 7 to ensure modern security and protect from vulnerabilities and provide a secure foundation for new device ecosystems.
WPA3 Requirements in Wi-Fi 6E (6 GHz) :
· WPA3 is mandatory for all Wi-Fi 6E devices operating in 6 GHz band.
· WPA3-Personal (SAE) with H2E for home/personal use.
· WPA3-Enterprise (802.1X, with optional 192-bit security suite) for enterprise deployments.
· Enhanced Open (OWE) for open networks requiring encryption without passwords
· Protected Management Frames (PMF) is mandatory in 6 GHz.
· WPA2 is not permitted in 6 GHz operation.
Note: As per the WPA3 v3.4 specifications (Section 11.2), Enhanced Open transition mode is not supported with 6 GHz.
There are no new specific ciphers or algorithm requirements for WPA3-Enterprise, apart from 802.11w/Protected Management Frame (PMF) enforcement. Many vendors, including Cisco, consider 802.1X-SHA256 or "FT + 802.1X" (which actually is 802.1X with SHA256 and Fast Transition on top) only to be WPA3 compliant and plain 802.1X (which uses SHA1) is considered part of WPA2, therefore not fit/supported for 6 GHz.
WPA3 Requirements in Wi-Fi 7:
· WPA3 is mandatory for all Wi-Fi 7 devices for features like Multi Link Operation and 802.11be data rates.
· WPA3-Personal (SAE) with GCMP256 as Cipher and SAE-EXT-KEY or the FT equivalent of it FT-SAE-EXT-KEY as AKMs.
· WPA3-Enterprise with AES (CCMP128) and 802.1X-SHA256 or the FT equivalent of it FT+802.1X (which still uses SHA256, though it’s not explicit in naming) as AKM.
Note: Cipher requirement of GCMP256 is required for WPA3-Enterprise. However, it’s not strictly enforced in the Access Point and Wireless clients.
· Enhanced Open (OWE) with GCMP256 as Cipher for open networks requiring encryption without passwords
· Protected Management Frames (PMF) is mandatory.
· Beacon Protection is mandatory.
Note: Similar to Wi-Fi 6E, Enhanced Transition Mode is not supported for Wi-Fi 7 operation in 6 GHz band. It is recommended to configure a pure OWE only WLAN for Wi-Fi 7 operation in 6 GHz band.
Note: Because Wi-Fi 7 is still a recent certification at the time of this writing, with an as early as possible release, many vendors did not enforce all these security requirements from the beginning.
The table below provides the security requirements for different Wi-Fi standards.
Note: AKMs and Cipher highlighted in “red” are mandatory for Wi-Fi 7.
Note: Standard requires GCMP256 as cipher for WPA3 Enterprise with 802.1x-SHA256, but most clients in the market today are capable of Wi-Fi 7 functionality with AES (CCMP128)
Migration Considerations
In many existing networks, WLANs continue to operate on Wi-Fi 6 or earlier standards, secured with WPA2 or older protocols. This is primarily due to legacy Access Point (AP) infrastructure—such as Wave-1 APs—that do not support WPA3, as well as a significant number of client devices lacking WPA3 capability.
With the introduction of Wi-Fi 6E (6 GHz) and Wi-Fi 7, the Wi-Fi Alliance has mandated WPA3 as a requirement for operating in the 6 GHz spectrum and for enabling the full feature set of Wi-Fi 7. This creates practical challenges when migrating to modern AP platforms while simultaneously upgrading all WLANs to WPA3.
Organizations have three practical design and deployment options to address these challenges:
1. Migrate all WLANs to WPA3/Enhanced Open
- 
    Provides a fully secure WLAN environment.
- 
    However, it poses challenges in maintaining coexistence with legacy WLANs and clients.
- 
    Best suited if: 
  - 
        All APs support WPA3 (i.e., Wave-2 or newer).
  - 
        No legacy clients remain that lack WPA3 support.
- 
        
2. Redesign or introduce new WLANs with WPA3/Enhanced Open
- 
    Retain existing WLANs to continue supporting legacy clients.
- 
    Deploy new WLANs using WPA3/Enhanced Open to: 
  - 
        Meet Wi-Fi 6E and Wi-Fi 7 requirements.
  - 
        Support modern clients with WPA3 capability.
- 
        
- 
    This hybrid model offers flexibility but requires managing two sets of SSIDs, which increases operational overhead.
3. Use Transition Modes to support multiple security standards
- 
    Convert WPA2 WLANs to WPA3 Transition Mode (also known as WPA2+WPA3 Mixed Mode).
- 
    In this mode, APs broadcast both WPA2 and WPA3 capabilities: 
  - 
        WPA2-only clients can continue connecting.
  - 
        WPA3-capable clients connect using WPA3.
- 
        
- 
    This approach: 
  - 
        Preserves existing SSID names.
  - 
        Provides a smooth migration path with minimal disruption for legacy devices.
- 
        
