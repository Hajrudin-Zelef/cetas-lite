---
id: collect-261001-general-networking/general-networking/wireless-design-and-configure-deployment-guides-mac-based-access-control-using-m-ced9a91e-1
title: "wireless-design-and-configure-deployment-guides-mac-based-access-control-using-m-ced9a91e"
domain: general-networking
role: reference
task: reference
actors: ["Microsoft"]
dates: []
keywords: []
source: docs/RAG/collect-261001-general-networking/wireless-design-and-configure-deployment-guides-mac-based-access-control-using-m-ced9a91e.md
source_anchor: ""
source_lines: [1, 54]
sha256: c407d9aded2b2f97aebdbf5d2f93e579b887b9ec0fc62535981fd919441d702f
---

# wireless-design-and-configure-deployment-guides-mac-based-access-control-using-m-ced9a91e

Configuring RADIUS Authentication with WPA2-Enterprise
Cisco Meraki MR access points (APs) offer a number of authentication methods for wireless association, including the use of external authentication servers to support WiFi Protected Access 2 - Enterprise (WPA2-Enterprise). This article outlines dashboard configuration to use a RADIUS server for WPA2-Enterprise authentication, RADIUS server requirements, and an example server configuration using Windows Network Policy Server (NPS).
For troubleshooting guidance, please refer to RADIUS Issue Resolution Guide.
Overview
WPA2-Enterprise with 802.1X authentication can be used to authenticate users or computers in a domain. The supplicant (wireless client) authenticates against the RADIUS server (authentication server) using an Extensible Authentication Protocol (EAP) method configured on the RADIUS server. The gateway access point (authenticator) sends authentication messages between the supplicant and authentication server. This means the RADIUS server is responsible for authenticating users.
Access points perform Extensible Authentication Protocol Over LAN (EAPOL) exchanges between the supplicant and convert these to RADIUS Access-Requests messages, which are sent to the RADIUS server's IP address and UDP port specified in dashboard. Gateway access points need to receive a RADIUS Access-Accept message from the RADIUS server in order to grant the supplicant access to the network.
To achieve the best performance, we recommend placing the RADIUS server and gateway access points within the same layer-2 broadcast domain to avoid firewall, routing, or authentication delays. Keep in mind the access point is not responsible for authenticating wireless clients and acts as an intermediary between clients and the RADIUS server.
The following image provides a detailed breakdown of the Protected Extensible Authentication Protocol (PEAP) with the Microsoft Challenge Handshake Authentication Protocol version 2 (MSCHAPv2) association process:
Supported RADIUS Attributes
When WPA2-Enterprise with 802.1X authentication is configured, the following attributes are present in the Access-Request messages sent from the Cisco Meraki access point to the customer's RADIUS server.
Note: Please refer to RFC 2865 for details on these attributes, additional notes for certain attributes are included below.
- User-Name
- NAS-IP-Address
- NAS-Port
- Called-Station-ID: Contains (1) the Meraki access point's MAC address (all caps, octets separated by hyphens) and (2) the Service Set Identifier (SSID) on which the wireless device is connecting. These 2 fields are separated by a colon. Example: "AA-BB-CC-DD-EE-FF:SSID_NAME".
Note: SSIDs broadcasted by repeater access points in a mesh deployment can't use NAS-IP-Address attribute because repeater access points do not have IP addresses assigned. You can use NAS-ID attribute instead, which by default carries NODE_MAC:VAP_NUM.
- Calling-Station-ID: Contains the MAC address of the wireless device (all caps, octets separated by hyphens). Example: "AA-BB-CC-DD-EE-FF".
- Framed-MTU
- NAS-Port-Type
- Connect-Info
- Meraki-Device-Name: Name of the Meraki device as configured in the dashboard
The following attributes are honored by Cisco Meraki when received in an Access-Accept message from the customer's RADIUS server to the Cisco Meraki access point:
- Tunnel-Private-Group-ID: Contains the VLAN ID that should be applied to a wireless user or device. (This can be configured to override VLAN settings that an administrator has configured for a particular SSID in the Cisco Meraki Cloud Controller.)
- Tunnel-Type: Specifies the tunneling protocol. Example: VLAN.
- Tunnel-Medium-Type: Sets the transport medium type used to create the tunnel. Example: 802 (which includes 802.11).
- Filter-Id / Reply-Message / Airespace-ACL-Name / Aruba-User-Role: Any of these attributes can be used to convey a policy that should be applied to a wireless user or device. (Please ensure that the attribute type matches what is configured on the Network-wide > Configure > Group policies page in the Cisco Meraki Cloud Controller, and the attribute value matches the name of the group policy configured on that page.)
RADIUS Configuration
The most common EAP configuration is PEAP with MSCHAPv2, which prompts users for credentials (either user or machine authentication).
Note: Certificate-based authentication using EAP-TLS is also supported by the Meraki platform, but is outside the scope of this document. For more information, refer to the documentation on RADIUS: WPA2-Enterprise With EAP-TLS.
RADIUS Server Requirements
There are many server options available for RADIUS, which should work with MR access points if configured correctly. Please refer to your RADIUS server documentation for specifics, but the key requirements for WPA2-Enterprise with Meraki are as follows:
- The server must host a certificate from a Certificate Authority (CA) trusted by clients on the network.
- All gateway access points broadcasting the WPA2-Enterprise SSID must be configured as RADIUS clients/authenticators on the server with a shared secret.
- The RADIUS server must have a user base to authenticate against.
Once the RADIUS server is configured, refer to the Dashboard Configuration section below for instructions on how to add your RADIUS server to dashboard.
User vs. Machine Authentication
The most common method of authentication with PEAP-MSCHAPv2 is user authentication, in which clients are prompted to enter their domain credentials. It is also possible to configure RADIUS for machine authentication, in which the computers themselves are authenticated against RADIUS, so the user doesn't need to provide any credentials to gain access. Typically, you can use EAP-TLS to configure machine authentication. However, some RADIUS server options make it simple to use PEAP-MSCHAPv2 to configure machine authentication (including Windows NPS, outlined in the example configuration below).
Note: "Machine Authentication" is not the same as MAC-based authentication, which is another configuration option in dashboard under Wireless > Configure > Access Control. Machine authentication, specifically, refers to devices authenticating against RADIUS.
Example RADIUS Configuration (Windows NPS + AD)
The following example configuration outlines how to set up Windows NPS as a RADIUS server, with Active Directory (AD) acting as a userbase:
- Add the NPS role to Windows Server.
- Add a trusted certificate to NPS.
- Add APs as RADIUS clients on the NPS server.
- Configure a policy in NPS to support PEAP-MSCHAPv2.
- (Optional for machine auth) Deploy PEAP-MSCHAPv2 wireless network settings to domain member computers using Group Policy.
Add the NPS Role to Windows Server
Microsoft's RADIUS server offering for Windows Server 2008 and later is their NPS. Please refer to the following two Microsoft documents for instructions on adding the NPS role to Windows Server, and registering the new NPS server in Active Directory (allowing it to use AD as its userbase):
Add a Trusted Certificate to NPS
A RADIUS server must host a certificate that allows both network clients and Meraki access points to validate the server's identity. There are three options for this certificate:
- Acquire a certificate from a trusted Certificate Authority
As long as the CA used is trusted by clients on the network, a certificate can be purchased and uploaded to NPS to verify server identity (required by clients). Common examples of trusted CAs include GoDaddy and VeriSign.
- Implement a Public Key Infrastructure (PKI) and generate a certificate (advanced)
A PKI can be used on the network to issue certificates trusted by clients on the network. We recommend you have a strong understanding of PKI for this option.
- Generate a self-signed certificate and turn off client server validation (insecure)
