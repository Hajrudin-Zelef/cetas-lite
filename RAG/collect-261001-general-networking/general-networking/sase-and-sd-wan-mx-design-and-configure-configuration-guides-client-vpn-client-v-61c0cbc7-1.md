---
id: collect-261001-general-networking/general-networking/sase-and-sd-wan-mx-design-and-configure-configuration-guides-client-vpn-client-v-61c0cbc7-1
title: "sase-and-sd-wan-mx-design-and-configure-configuration-guides-client-vpn-client-v-61c0cbc7"
domain: general-networking
role: reference
task: reference
actors: ["Apple", "Google", "Microsoft"]
dates: []
keywords: ["training"]
source: docs/RAG/collect-261001-general-networking/sase-and-sd-wan-mx-design-and-configure-configuration-guides-client-vpn-client-v-61c0cbc7.md
source_anchor: ""
source_lines: [1, 106]
sha256: 805bc45f658f0623463e57a1ee39dd6595fb292f6aab0f5c9371ed4262b9eaf7
---

# sase-and-sd-wan-mx-design-and-configure-configuration-guides-client-vpn-client-v-61c0cbc7

Client VPN OS Configuration
This article outlines instructions to configure a client VPN connection on commonly used operating systems.
Learn more with these free online training courses on the Meraki Learning Hub:
A note on IKEv2
In the firmware release for MX 26.1.2 we've introduced IKEv2 support for Client VPN. For MX configuration details see this article
- Hostname: You must use the MX's Hostname. This can be found in the Dashboard under Security & SD-WAN >  Client VPN 
    
  - The host name should look like this: xyz-123abc123abc.dynamic-m.com
- Authentication: IKEv2 utilizes certificate-based authentication for the server and EAP-MSCHAPv2 for the user auth (RADIUS). There is no Pre-Shared Key (PSK) used for IKEv2 connections. At the time of this writing, only RADIUS authentication is supported for IKEv2
- Local ID: enter your VPN username
- Firmware: MX 26.1.2 or higher.
Android
Note: Android devices running Android 12 and above do not support Layer 2 Tunneling Protocol/Internet Protocol Security (L2TP/IPsec) VPNs. Devices with existing configurations will continue to work. L2TP Client VPN connection cannot be configured on new devices.
To check the Android version of a device, see Check & update your Android version in Google Support.
To configure an Android device to connect to the client VPN, see Connect to a virtual private network (VPN) on Android in Google Support.
The following VPN information is needed to complete the setup:
- Name: This can be anything you want to name the connection, for example, "Work VPN"
- Type: 
    
  - L2TP: Select L2TP/IPSEC PSK
  - IKEv2: Select IKEv2/IPSEC MSCHAPv2
- Server address: Enter the hostname (for example: abc-abcd.dynamic-m.com) or the active WAN IP (for example: a.b.c.d)
    
  - Hostname is required for IKEv2 and preferred for L2TP to improve reliability during WAN failover
  - This information is located in the Meraki dashboard under Security & SD-WAN > Monitor > Appliance status
- IPSEC Identifier: (IKEv2 only) Enter VPN username
- IPSec pre-shared key: (L2TP Only) Enter the pre-shared key that admin created in Security & SD-WAN > Configure > Client VPN
Chrome OS
To configure a Chrome OS device to connect to client VPN, see Set up virtual private networks (VPNs) in Google Support.
The following VPN information is needed to complete the setup:
- Service name: This can be anything you want to name this connection, for example, "Work VPN"
- Provider type: 
    
  - L2TP: Select L2TP/IPSEC PSK
  - IKEv2: Select IKEv2/IPSEC MSCHAPv2
- Server hostname (or Remote Identity): Enter the hostname (for example: abc-abcd.dynamic-m.com) or the active WAN IP (for example: a.b.c.d)
    
  - Hostname is required for IKEv2 and preferred for L2TP to improve reliability during WAN failover
  - This information is located in the Meraki dashboard under Security & SD-WAN > Monitor > Appliance status
- IPSEC Identifier (or Local Identity): (IKEv2 only) Enter VPN username
- Authentication type: (L2TP only) Select Pre-shared key
- Pre-shared key: (LT2P only) Enter the shared secret that admin created in Security & SD-WAN > Configure > Client VPN
- Username: Credentials for connecting to VPN—if using Meraki authentication, this will be an email address
- Password: Credentials for connecting to VPN
iOS
To configure an iOS device to connect to the client VPN, follow these steps:
- Navigate to Settings > General > VPN & Device Management > VPN > Add VPN Configuration
- Type: 
    
  - L2TP: Set to L2TP
  - IKEv2: Set to IKEv2
- Description: This can be anything you want to name this connection, for example, "Work VPN"
- Server: Enter the hostname (for example: abc-abcd.dynamic-m.com) or the active WAN IP (for example: a.b.c.d)
    
  - Hostname is required for IKEv2 and preferred for L2TP to improve reliability during WAN failover
  - This information is located in the Meraki dashboard under Security & SD-WAN > Monitor > Appliance status
- Remote ID: (IKEv2 only) Enter the hostname (for example: abc-abcd.dynamic-m.com)
- Local ID: (IKEv2 only) enter your VPN username
- Account: Enter the username
- Password: Enter if desired
    
  - If the password is left blank, it will need to be entered each time the device attempts to connect to the client VPN
- Secret: (L2TP Only) Enter the shared secret that admin created in Security & SD-WAN > Configure > Client VPN
- Ensure that Send All Traffic is set to on
- Save the configuration
To configure a macOS device to connect to client VPN, see Set up a VPN connection on Mac in Apple Support.
The following VPN information is needed:
- Display Name: This can be anything you want to name this connection, for example, "Work VPN"
- Server: Enter the hostname (for example: abc-abcd.dynamic-m.com) or the active WAN IP (for example: a.b.c.d)
    
  - Hostname is required for IKEv2 and preferred for L2TP to improve reliability during WAN failover
  - This information is located in the Meraki dashboard under Security & SD-WAN > Monitor > Appliance status
- Remote ID: (IKEv2 only) Enter the hostname (for example: abc-abcd.dynamic-m.com)
- Local ID: (IKEv2 only) enter your VPN username
- Account Name: Enter the account name of the user (based on AD, RADIUS, or Meraki cloud authentication)
- Password: User password (based on AD, RADIUS or Meraki cloud authentication)
- Machine Authentication > Shared Secret: (L2TP only) Enter the shared secret that admin created in Security & SD-WAN > Configure > Client VPN
Ensure that the MACs network adapter service order includes the VPN interface as the first item (in the list) otherwise all the traffic will not leave on the Client VPN tunnel. For more reference, see Change the order of the network services your Mac uses in Apple support.
The following authentication methods are supported:
- User authentication: Active Directory (AD), RADIUS*, or Meraki-hosted authentication
- Machine authentication: Pre-shared keys
When using Meraki-hosted authentication, the VPN account and username setting is the user email address entered in the Meraki dashboard.
To configure a Windows 10 or Windows 11 device to connect to client VPN, see Connect to a VPN in Windows in Microsoft Support page. Additional settings must be changed for Windows devices when trying to connect using IKEv2, as the default settings are not secure. Note that additional configuration changes are required for Windows devices connecting via IKEv2, as the default settings are not secure. Please see this article on the Microsoft support page for configuration details.
Below is an example of supported security settings for configuring a Windows device using IKEv2:
- Encryption Algorithm: AES256
- Integrity (Hash) Algorithm: SHA256
- Diffie-Hellman Group (Key Size): DH14
The following VPN information is needed to complete the setup:
- In the Settings app  on your Windows device, select Network & internet > VPN > Add VPN.
    
  - VPN provider: Set to Windows (built-in)
  - Connection name: This can be anything you want to name this connection, for example, "Work VPN"
  - Server name or address: Enter the hostname (for example: abc-abcd.dynamic-m.com) or the active WAN IP (for example: a.b.c.d)
        
    - Hostname is required for IKEv2 and preferred for L2TP to improve reliability during WAN failover
    - This information is located in the Meraki dashboard under Security & SD-WAN > Monitor > Appliance status
  - VPN type: 
        
    - LT2P: Select L2TP/IPsec with pre-shared key
    - IKEv2: Select IKE v2
  - User name and Password: optional
For L2TP:
For IKEv2:
After the VPN connection has been created, set the Authentication protocols:
- Choose the VPN connection and then select Advanced options > More VPN properties > Edit > Security Tab. 
    
