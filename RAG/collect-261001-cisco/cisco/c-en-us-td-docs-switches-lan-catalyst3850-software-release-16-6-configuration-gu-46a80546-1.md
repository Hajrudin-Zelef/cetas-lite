---
id: collect-261001-cisco/cisco/c-en-us-td-docs-switches-lan-catalyst3850-software-release-16-6-configuration-gu-46a80546-1
title: "c-en-us-td-docs-switches-lan-catalyst3850-software-release-16-6-configuration-gu-46a80546"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: ["voice"]
source: docs/RAG/collect-261001-cisco/c-en-us-td-docs-switches-lan-catalyst3850-software-release-16-6-configuration-gu-46a80546.md
source_anchor: ""
source_lines: [1, 36]
sha256: 3ec1deb569e43f9b26934e99d59f2ed7a5da41a9f8a4dc7ea1cd862d13acd996
---

# c-en-us-td-docs-switches-lan-catalyst3850-software-release-16-6-configuration-gu-46a80546

Information About 802.1x Port-Based Authentication
The 802.1x standard defines a client-server-based access control and authentication protocol that prevents unauthorized clients from connecting to a LAN through publicly accessible ports unless they are properly authenticated. The authentication server authenticates each client connected to a switch port before making available any services offered by the switch or the LAN.
| Note | TACACS is not supported with 802.1x authentication. | 
Until the client is authenticated, 802.1x access control allows only Extensible Authentication Protocol over LAN (EAPOL), Cisco Discovery Protocol (CDP), and Spanning Tree Protocol (STP) traffic through the port to which the client is connected. After authentication is successful, normal traffic can pass through the port.
The table shown below lists the maximum number of each client session supported on Catalyst 3850 and Catalyst 3650 switches:
| Client session | Maximum sessions supported | 
|---|---|
| Maximum dot1x or MAB client sessions | 2000 | 
| Maximum web-based authentication sessions | 2000 | 
| Maximum dot1x sessions with critical-auth VLAN enabled and server re-initialized | 2000 | 
| Maximum MAB sessions with various session features applied | 2000 | 
| Maximum dot1x sessions with service templates or session features applied | 2000 | 
Port-Based Authentication Process
To configure IEEE 802.1X port-based authentication, you must enable authentication, authorization, and accounting (AAA) and specify the authentication method list. A method list describes the sequence and authentication method to be queried to authenticate a user.
The AAA process begins with authentication. When 802.1x port-based authentication is enabled and the client supports 802.1x-compliant client software, these events occur:
-  
                                 		  
                                 If the client identity is valid and the 802.1x authentication succeeds, the switch grants the client access to the network.
-  
                                 		  
                                 If 802.1x authentication times out while waiting for an EAPOL message exchange and MAC authentication bypass is enabled, the switch can use the client MAC address for authorization. If the client MAC address is valid and the authorization succeeds, the switch grants the client access to the network. If the client MAC address is invalid and the authorization fails, the switch assigns the client to a guest VLAN that provides limited services if a guest VLAN is configured.
-  
                                 		  
                                 If the switch gets an invalid identity from an 802.1x-capable client and a restricted VLAN is specified, the switch can assign the client to a restricted VLAN that provides limited services.
-  
                                 		  
                                 If the RADIUS authentication server is unavailable (down) and inaccessible authentication bypass is enabled, the switch grants the client access to the network by putting the port in the critical-authentication state in the RADIUS-configured or the user-specified access VLAN. 
 Note
 Inaccessible authentication bypass is also referred to as critical authentication or the AAA fail policy. 
If Multi Domain Authentication (MDA) is enabled on a port, this flow can be used with some exceptions that are applicable to voice authorization.
The switch re-authenticates a client when one of these situations occurs:
-  
                                 		  
                                 Periodic re-authentication is enabled, and the re-authentication timer expires. You can configure the re-authentication timer to use a switch-specific value or to be based on values from the RADIUS server. After 802.1x authentication using a RADIUS server is configured, the switch uses timers based on the Session-Timeout RADIUS attribute (Attribute[27]) and the Termination-Action RADIUS attribute (Attribute [29]). The Session-Timeout RADIUS attribute (Attribute[27]) specifies the time after which re-authentication occurs. The range is 1 to 65535 seconds. The Termination-Action RADIUS attribute (Attribute [29]) specifies the action to take during re-authentication. The actions are Initialize and ReAuthenticate. When the Initialize action is set (the attribute value is DEFAULT), the 802.1x session ends, and connectivity is lost during re-authentication. When the ReAuthenticate action is set (the attribute value is RADIUS-Request), the session is not affected during re-authentication.
-  
                                 		  
