---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/hc-en-us-articles-360015268353-configuring-a-radius-server-in-unifi-14e441a2
title: "hc-en-us-articles-360015268353-configuring-a-radius-server-in-unifi-14e441a2"
domain: unifi-ubiquiti
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-unifi-ubiquiti/hc-en-us-articles-360015268353-configuring-a-radius-server-in-unifi-14e441a2.md
source_anchor: ""
source_lines: [1, 64]
sha256: a11c9decfb91733733f55bbcfa1c290a6c3df8f7786c756d339f3dab045ee2d5
---

# hc-en-us-articles-360015268353-configuring-a-radius-server-in-unifi-14e441a2

Configuring a RADIUS Server in UniFi
UniFi Gateways come equipped with a built-in RADIUS server, which can be used with the 802.1X standard to provide secure authentication for VPNs and network access.
How to Enable the RADIUS Server
- Navigate to Settings > Networks and scroll down to the RADIUS Servers section.
- Select the Default RADIUS Server to access its Settings.
- Select Enable.
- Adjust the rest of its settings, including:
  - Secret: Pre-shared key provisioned to the authenticator device(s) and the RADIUS server. This provides authentication between the two, ensuring RADIUS message integrity.
  - Authentication Port: The port through which RADIUS authentication messages are sent and received by authenticator and RADIUS server devices.
  - Accounting Port: The port through which RADIUS accounting messages are sent and received by authenticator and RADIUS server devices.
  - Accounting Interim Interval: The time (in ms) in which a RADIUS access request packet is sent with an Acct-Status-Type attribute with the value "interim-update". This update requests the status of an active session. "Interim" records contain the current session duration and can provide data usage details.
Note: The Radius server collects client information sent by the authenticator that can be used for accounting and network activity reporting. This information is sent when the user logs on and logs off, these are usually called accounting requests. For more information on RADIUS accounting, see RFC2866.
How to Create Users in the Network Application
- Navigate to Settings > Networks and scroll down to the RADIUS Servers section.
- Select Local Credentials.
- Fill in the following details:
For information on authenticating devices based on MAC addresses, visit RADIUS-Based MAC Authentication and 802.1X.
How to Enable RADIUS Assigned VLAN
- Navigate to Settings > Networks and scroll down to the RADIUS Servers section and select your profile.
- Enable RADIUS Assigned VLAN Support for the desired network types and apply changes.
- Select Local Credentials.
- Assign your Users to specific VLANs.
  - For dynamic VLAN users, set the tunnel-type to 13 and the tunnel-medium-type to 6.
Note: If the user profile does not specify a VLAN, the client will fall back to the untagged VLAN.
How to Enable RADIUS over TLS (RADSEC)
Before you begin, please note that RADIUS over TLS is for WiFi only.
- Obtain a certificate bundle from your Passpoint provider or a third-party RADIUS server. This bundle should include a Client Certificate, Private Key, and CA Certificate(s).
- Go to Settings > Networks > RADIUS Servers and select the profile you want to configure.
- Toggle the TLS option to enable RADIUS over TLS.
- Upload the Client Certificate, Private Key, and CA Certificate(s) from your bundle. These certificates should be in PEM format.
Note: Omit the password field unless explicitly mandated by your Passpoint provider or RADIUS server.
How to Enable RADIUS Change of Authorizations (CoA) on a Wireless Network
- Navigate to Settings > WiFi and select your Wireless Network.
- Set your Security Protocol to WPA2 Enterprise or WPA3 Enterprise.
- Select your previously configured RADIUS Profile.
- Enable DAS/DAC (CoA).
Note: Radius CoA has the following requirements:
- RADIUS Accounting servers must be configured.
- A short interim update (e.g., 300 seconds) is recommended—check with your RADIUS provider for their recommendation.
- Port 3799 and greater must be open between the AP and RADIUS server for CoA information to be exchanged.
- 
Required RADIUS attributes:
  - Account (User-Name, Framed-IP-Address, Calling-Station-ID, etc)
  - Acct-Session-id
  - NAS-IP-Address
  - NAS-Identifier
What is 802.1X?
The 802.1X standard is an access control standard for authenticating devices on a network. 802.1X has four components:
- Supplicant: The device requesting access to the system services.
- Authenticator: The port or device that communicates with the Authentication Server before granting system access.
- Authentication Server: The external server (e.g., a RADIUS server) that performs the authentication, indicating whether the supplicant is authorized to access system services.
- Accounting Server: The (optional) external server (e.g., a RADIUS server) that records information about the RADIUS session, including client access credentials and connection time, during logon/logoff and periodically while connected.
Support for RADIUS over TLS (RADSEC) has been added to UniFi Network 8.4 and newer versions. This requires a Client Certificate, Private Key, and CA Certificate from a supported RADIUS server.
802.1X Authentication Process
A client device is authorized by a RADIUS server with 802.1X in the following process:
- The client device is prompted for credentials.
- The user enters credentials.
- The client device sends a data link layer request to an authenticator for network access.
- The authenticator sends a message called a "RADIUS Access Request" to the RADIUS server.
- The RADIUS server returns one of three responses to the authenticator:
  - Access-Reject: The user is denied network access.
  - Access-Challenge: The user must provide additional authentication, such as secondary password, token, PIN, or card. This message is also used in more complex authentication scenarios where a secure tunnel is established between the user’s device and RADIUS server.
  - Access-Accept: The user is granted access to the network.
- If Accounting is configured and an Interim Update Interval is set, the client will periodically send information to the accounting server until that client disconnects from the network.
