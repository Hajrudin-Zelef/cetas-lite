---
id: collect-260926-mikrotik/mikrotik/arborescence-des-pages-77-1
title: "Summary"
domain: mikrotik
role: reference
task: reference
actors: ["Microsoft"]
dates: []
keywords: ["parameters"]
source: docs/RAG/lot-mikrotik/RouterOS/arborescence-des-pages-77.md
source_anchor: ""
source_lines: [1, 87]
sha256: e1b197dc521a7155876372f4a2cd471af64de48b3a9d91565999209db8105697
---

# Summary

RADIUS, short for Remote Authentication Dial-In User Service, is a remote server that provides authentication and accounting facilities to various network appliances. RADIUS authentication and accounting allows the ISP or network administrator to manage PPP user access and accounting from one server throughout a large network. The MikroTik RouterOS has a RADIUS client that can authenticate router's local users, HotSpot, PPP, PPPoE, PPTP, L2TP, OVPN, SSTP, IPsec and ISDN connections. The attributes received from the RADIUS server override the ones set in the default profile, but if some parameters are not received they are taken from the respective default profile.

The RADIUS server database is consulted only if no matching user access record is found in the router's local database.

If RADIUS accounting is enabled, accounting information is also sent to the RADIUS server default for that service.

# RADIUS Client

| **Sub-menu:**`/radius` | 
|---|

This sub-menu allows adding and removing RADIUS clients.

The order of added items in this list is significant.

## Properties

| Property | Description | 
|---|---|
| **accounting-backup** (*yes \| no* ; Default:**no** ) | Whether the configuration is for the backup RADIUS server | 
| **accounting-port** (*integer [1..65535]* ; Default:**1813** ) | RADIUS server port used for accounting | 
| **address** (*IPv4/IPv6 address* ; Default:**0.0.0.0** ) | IPv4 or IPv6 address of RADIUS server. The following formats are accepted: - *ipv4* - *ipv4*`@`*vrf* - *ipv6* - *ipv6*`@`*vrf* | 
| **authentication-port** (*integer [1..65535]* ; Default:**1812** ) | RADIUS server port used for authentication. | 
| **called-id** (*string* ; Default: ) | Value depends on Point-to-Point protocol: PPPoE - service name, PPTP - server's IP address, L2TP - server's IP address. | 
| **certificate** (*string* ; Default: ) | Certificate file to use for communicating with RADIUS Server with RadSec enabled. | 
| **comment** (*string* ; Default: ) |  | 
| **disabled** (*yes \| no* ; Default:**no** ) |  | 
| **domain** (*string* ; Default: ) | Microsoft Windows domain of client passed to RADIUS servers that require domain validation. | 
| **protocol** (*radsec \| udp* ; Default:**udp** ) | Specifies the protocol to use when communicating with the RADIUS Server. | 
| **radsec-timeout**  (*time,* Default:**3300ms** ) | Timeout after which the request should be resent over RadSec protocol. | 
| **require-message-auth** (*no \| yes-for-request-resp* Default:**yes-for-request-resp** ) | Specifies if Message-Authenticator attributes are required. | 
| **realm** (*string* ; Default: ) | Explicitly stated realm (user domain), so the users do not have to provide proper ISP domain name in the user name. | 
| **secret** (*string* ; Default: )*sensitive* | The shared secret used to access the RADIUS server. | 
| **service** (*ppp\|login\|hotspot\|wireless\|dhcp\|ipsec\|dot1x* ; Default: ) | Router services that will use this RADIUS server:  | 
| **src-address** (*ipv4/ipv6 address* ; Default:**0.0.0.0** ) | Source IP/IPv6 address of the packets sent to the RADIUS server | 
| **timeout** (*time* ; Default:**1100ms** ) | Timeout after which the request should be resent. | 

When the RADIUS server is authenticating the user with CHAP, MS-CHAPv1, MS-CHAPv2, it is not using a shared secret, the secret is used only in the authentication reply, and the router (RADIUS client) verifies it. So if you have the wrong shared secret, the RADIUS server will accept a request, but the router won't accept the reply. You can see that with "*/radius monitor*" command, the "bad-replies" number should increase whenever somebody tries to connect.

If RadSec is enabled, make sure your RADIUS Server is using "**radsec**" as the shared secret, otherwise, the RADIUS Server will not be able to decrypt data correctly (unprintable characters). With RadSec RouterOS forces the shared secret to "radsec" regardless of what has been set manually. For more details see - RFC6614.

## Example

To set up a RADIUS Client for HotSpot and PPP services that will authenticate against a RADIUS Server (10.0.0.3), you need to do the following:

To set up a RADIUS Client with RadSec, you need to do the following:

Make sure the specified certificate is trusted and common-name verification on both ends will work properly.

Here is a trivial example of certificates that can be used for RADIUS RadSec configuration:

To view RADIUS Client statistics, you need to do the following:

Make sure you enable RADIUS authentication for the desired services:

# Connection Terminating from RADIUS

| **Sub-menu:**`/radius incoming` | 
|---|

This facility supports unsolicited messages sent from the RADIUS server. Unsolicited messages extend RADIUS protocol commands, that allow terminating a session that has already been connected from the RADIUS server. For this purpose, DM (Disconnect-Messages) is used. Disconnect messages cause a user session to be terminated immediately.

RouterOS doesn't support POD (Packet of Disconnect) the other RADIUS access request packet that performs a similar function as Disconnect Messages

## Properties

| Property | Description | 
|---|---|
| **accept** (*yes \| no* ; Default:**no** ) | Whether to accept unsolicited messages | 
| **port** (*integer* ; Default:**1700** ) | The port number to listen for the requests on | 
| **vrf** (*VRF name* ; default value:**main** ) | Set VRF on which service is listening for incoming connections | 

# Supported RADIUS Attributes

Here you can download the RADIUS reference dictionary, that includes all supported RADIUS attributes by MikroTik device. This file is designed for FreeRADIUS, but may also be used by other RADIUS servers. Note, it may conflict with the default configuration file of your RADIUS server, correct the configuration, not the dictionary, as no other attributes are supported by MikroTik RouterOS. There is also the RADIUS MikroTik specific dictionary that can be included in an existing dictionary to support MikroTik vendor-specific attributes.

Below you will find description about attributes and how they are used on MikroTik devices during communication with RADIUS.

## Definitions

- **PPPs** - PPP,PPTP, PPPoE
- **default configuration** -settings in default profile (for PPPs) or HotSpot server settings (for HotSpot)

## Access-Request packet

