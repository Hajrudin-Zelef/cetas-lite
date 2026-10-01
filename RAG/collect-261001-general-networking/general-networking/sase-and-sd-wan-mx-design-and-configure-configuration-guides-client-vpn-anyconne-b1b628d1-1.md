---
id: collect-261001-general-networking/general-networking/sase-and-sd-wan-mx-design-and-configure-configuration-guides-client-vpn-anyconne-b1b628d1-1
title: "sase-and-sd-wan-mx-design-and-configure-configuration-guides-client-vpn-anyconne-b1b628d1"
domain: general-networking
role: reference
task: reference
actors: ["AWS"]
dates: []
keywords: ["aws", "datacenter", "training"]
source: docs/RAG/collect-261001-general-networking/sase-and-sd-wan-mx-design-and-configure-configuration-guides-client-vpn-anyconne-b1b628d1.md
source_anchor: ""
source_lines: [1, 40]
sha256: 49668f8631b7e101fb1941bbf6eacb0809393dabfec411fe3b2ceb9211a2be34
---

# sase-and-sd-wan-mx-design-and-configure-configuration-guides-client-vpn-anyconne-b1b628d1

AnyConnect on the MX Appliance
Click 日本語 for Japanese
Overview
A new AnyConnect Settings page in now available, and can be enabled on the dashboard by navigating to Organization > Configure > Early Access
The Cisco AnyConnect Secure Mobility Client consistently raises the bar by making the remote-access experience easy for end users. It helps enable a highly secure connectivity experience across a broad set of PC and mobile devices. This document provides information on the AnyConnect integration on Meraki appliances and instructions for configuring AnyConnect on the Meraki dashboard.
Learn more with these free online training courses on the Meraki Learning Hub:
Feature
The AnyConnect VPN server on the MX uses Transport Layer Security (TLS) & Datagram Transport Layer Security (DTLS) for tunneling and requires AnyConnect VPN client version 4.8 or higher on either Windows, macOS, Linux, or mobile devices to terminate remote access connections successfully. The AnyConnect client negotiates a tunnel with the AnyConnect server and gives you the ability to access resources or networks on or connected to the AnyConnect server (MX). Unlike the AnyConnect implementation on the Adaptive Security Appliance (ASA), with support for other features like host scan, web launch, etc, the MX security appliance supports Secure Socket Layer (SSL), VPN, and other AnyConnect modules that do not require additional configuration on the MX. For more details, see AnyConnect on ASA vs. MX.
The MX supports Layer 2 Tunneling Protocol (L2TP)/Internet Protocol Security (IPsec) Client VPN and AnyConnect VPN simultaneously.
AnyConnect can be used in place of L2TP/IPSec Client VPN configurations on operating systems that no longer support L2TP VPN services as it is a TLS & DTLS application based VPN.
AnyConnect is currently not supported on template bound networks.
AnyConnect VPN traffic can be routed through Site-to-Site VPN (both AutoVPN and Non-Meraki VPN).
Use Cases
AnyConnect can be used to securely connect remote users to Branch Offices, Datacenter or Public Cloud environments. Using AnyConnect with the Meraki MX Appliance for remote access can enable users secure and seamless connectivity between different locations. Remote users can connect to a Branch office and transverse the Secure Software Defined Wide Area Network (SD-WAN) AutoVPN tunnel to access recourses in the Amazon Web Services (AWS)/Azure, etc., or other locations within the SD-WAN fabric.
Caveats
There are certain caveats to keep in mind before enabling AnyConnect:
- 
    Supported MX models: MX600, MX450, MX400, MX250, MX105, MX100, MX95, MX85, MX84, MX75, MX68(W,CW), MX67(C,W), MX65(W)*, MX64(W)*, Z3(C), Z4(C), vMX
*MX65(W) and MX64(W) only supports AnyConnect when running on firmware 17.6+
Not supported: MX90, 80, 60, Z1 (The AnyConnect Settings page will not be visible on dashboard for these models)
- 
    Either NAT Exceptions (No NAT) on MX Security Appliances or AnyConnect can be enabled per WAN uplink
- 
    IPsec and AnyConnect share the same configured RADIUS and Active directory servers
- 
    AnyConnect does not currently support cellular uplink (integrated or USB modem)
Note:
Stateless high availability and WAN failover are supported with AnyConnect on the MX. This means, when HA or WAN failover occurs, active user sessions will be disconnected and users will need to reconnect to the new active WAN link or the new primary MX. If you have two uplinks, its recommended to use the unique DDNS names of WAN1 and WAN2 as primary and backup servers in the AnyConnect connection profile.
How to Enable AnyConnect on Your Dashboard
Having reviewed the caveats, upgrade your MX security appliance to the required firmware version.
- 
    To enable AnyConnect, ensure that your network firmware version meets the minimum requirement. For MX64(W) and MX65(W), the firmware version must be 17.6+, and for all other supported models, the latest MX-16 firmware. Upgrades can be managed by navigating to Organization > Monitor > Firmware upgrades. For more details on firmware upgrades see Managing Firmware Upgrades
Server Settings
To enable AnyConnect VPN, select Enabled from the Cisco Secure Client Settings radio button on the Security & SD-WAN > Configure > Client VPN > Cisco Secure Client Settings tab. The following AnyConnect VPN options can be configured:
| Hostname: This is used by Client VPN users to connect to the MX. This hostname is a Dynamic DNS (DDNS) host record that resolves to the Public IP address of the MX. The DDNS hostname is a prerequisite for publicly trusted certificate enrollment. You can change this hostname by following instructions on the Dynamic DNS (DDNS) article. For an alternative to DDNS enrolled certificates, see Custom hostname certificates section in this article.  Cisco Secure Client port: This specifies the port the AnyConnect server will accept and negotiate tunnels on.  Log-in banner: This specifies the message seen on the AnyConnect client when a user successfully authenticates. If configured, a connecting user must acknowledge the message before getting network access on the VPN. To disable the log-in banner simply leave the banner field blank. Profile update: This specifies the AnyConnect VPN configuration profile that gets pushed to the user on authentication. |  | 
| Cisco Secure Client VPN subnet: This specifies the address pool used for authenticated clients. IPv6 prefix (MX 18.104+): This specifies IPv6 prefix for AnyConnect to support IPv6 to both terminate a client VPN tunnel as well as IPv6 traffic inside the tunnel. More information can be found in IPv6 Support on MX Security & SD-WAN Platforms - VPN document in Configuring IPv6 for AnyConnect section. DNS nameservers: This specifies the Domain Name System (DNS) settings assigned to the client. DNS suffix: This specifies the default domain name or DNS suffix passed to the AnyConnect client to append to DNS queries that omit the domain field. This domain name only applies to tunnelled packets. Client routing: This is used to specify full or split-tunnel rules pushed to the AnyConnect client device. You can send all traffic through VPN, all traffic except traffic going to specific destinations, or only send traffic going to specific destinations. Dynamic client routing: This is used to specify full or split-tunnel rules pushed to the AnyConnect client device by hostname. For more details see Dynamic Client routing section in this article. |  | 
| Authentication Type: This is used to specify authentication with Meraki Cloud, SAML, RADIUS, or Active Directory. Certificate Authentication: This is used to configure the trusted Certificate Authority (CA) file that is used to authenticate client devices. This configuration is only required if you need to authenticate client devices with a certificate. Only certificates PEM format (*.pem) are supported at this time. RADIUS servers: This is used to specify the RADIUS host IP address, RADIUS port that the MX Security Appliance will use to communicate to the server, and the RADIUS shared secret. RADIUS message-authenticator verification: This is used to enable a RADIUS attribute that ensures integrity and authenticity of a message. Group policy with RADIUS filter-ID: This is used to enable dashboard group policy application using the filter passed by the RADIUS server. RADIUS timeout: This is used to modify the RADIUS time-out for two-factor authentication and authentication server failover. Default Group Policy: This is used to apply a default group policy to all connecting AnyConnect clients. For more details see Group Policies section in this article. |  | 
Server Certificates
The AnyConnect server on the MX uses TLS 1.2 for tunnel negotiation, and TLS 1.3 on MX 19.1+, hence it needs a server identity certificate. The MX supports three certificate options:
Option 1: Auto-generated certificate with DDNS hostname
