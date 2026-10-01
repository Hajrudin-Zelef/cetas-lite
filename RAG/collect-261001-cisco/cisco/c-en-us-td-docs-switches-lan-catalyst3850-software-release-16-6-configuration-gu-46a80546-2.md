---
id: collect-261001-cisco/cisco/c-en-us-td-docs-switches-lan-catalyst3850-software-release-16-6-configuration-gu-46a80546-2
title: "c-en-us-td-docs-switches-lan-catalyst3850-software-release-16-6-configuration-gu-46a80546"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: ["ethernet"]
source: docs/RAG/collect-261001-cisco/c-en-us-td-docs-switches-lan-catalyst3850-software-release-16-6-configuration-gu-46a80546.md
source_anchor: ""
source_lines: [37, 75]
sha256: 0e2e911d136c101bc2b62edfeb251c453d6675e688f8e7819972785c5473412f
---

# c-en-us-td-docs-switches-lan-catalyst3850-software-release-16-6-configuration-gu-46a80546

                                 You manually re-authenticate the client by entering the dot1x re-authenticate interface interface-id privileged EXEC command.
Port-Based Authentication Initiation and Message Exchange
During 802.1x authentication, the switch or the client can initiate authentication. If you enable authentication on a port by using the authentication port-control auto interface configuration command, the switch initiates authentication when the link state changes from down to up or periodically as long as the port remains up and unauthenticated. The switch sends an EAP-request/identity frame to the client to request its identity. Upon receipt of the frame, the client responds with an EAP-response/identity frame.
However, if during bootup, the client does not receive an EAP-request/identity frame from the switch, the client can initiate authentication by sending an EAPOL-start frame, which prompts the switch to request the client’s identity.
| Note | If 802.1x authentication is not enabled or supported on the network access device, any EAPOL frames from the client are dropped. If the client does not receive an EAP-request/identity frame after three attempts to start authentication, the client sends frames as if the port is in the authorized state. A port in the authorized state effectively means that the client has been successfully authenticated. | 
When the client supplies its identity, the switch begins its role as the intermediary, passing EAP frames between the client and the authentication server until authentication succeeds or fails. If the authentication succeeds, the switch port becomes authorized. If the authentication fails, authentication can be retried, the port might be assigned to a VLAN that provides limited services, or network access is not granted.
The specific exchange of EAP frames depends on the authentication method being used.
If 802.1x authentication times out while waiting for an EAPOL message exchange and MAC authentication bypass is enabled, the switch can authorize the client when the switch detects an Ethernet packet from the client. The switch uses the MAC address of the client as its identity and includes this information in the RADIUS-access/request frame that is sent to the RADIUS server. After the server sends the switch the RADIUS-access/accept frame (authorization is successful), the port becomes authorized. If authorization fails and a guest VLAN is specified, the switch assigns the port to the guest VLAN. If the switch detects an EAPOL packet while waiting for an Ethernet packet, the switch stops the MAC authentication bypass process and starts 802.1x authentication.
Authentication Manager for Port-Based Authentication
Port-Based Authentication Methods
| Table 1. 802.1x Features |  |  |  |  | 
|---|---|---|---|---|
| Authentication method | Mode |  |  |  | 
|---|---|---|---|---|
|  | Single host | Multiple host | MDA | Multiple Authentication | 
| 802.1x | VLAN assignment Per-user ACL Filter-ID attribute Downloadable ACL Redirect URL | VLAN assignment | VLAN assignment Per-user ACL Filter-Id attribute Downloadable ACL Redirect URL | VLAN assignment Per-user ACL Filter-Id attribute Downloadable ACL Redirect URL | 
| MAC authentication bypass | VLAN assignment Per-user ACL Filter-ID attribute Downloadable ACL Redirect URL | VLAN assignment | VLAN assignment Per-user ACL Filter-Id attribute Downloadable ACL Redirect URL | VLAN assignment Per-user ACL Filter-Id attribute Downloadable ACL Redirect URL | 
| Standalone web authentication | Proxy ACL, Filter-Id attribute, downloadable ACL |  |  |  | 
| NAC Layer 2 IP validation | Filter-Id attribute Downloadable ACL Redirect URL | Filter-Id attribute Downloadable ACL Redirect URL | Filter-Id attribute Downloadable ACL Redirect URL | Filter-Id attribute Downloadable ACL Redirect URL | 
| Web authentication as fallback method | Proxy ACL Filter-Id attribute Downloadable ACL | Proxy ACL Filter-Id attribute Downloadable ACL | Proxy ACL Filter-Id attribute Downloadable ACL | Proxy ACL Filter-Id attribute Downloadable ACL | 
Per-User ACLs and Filter-Ids
| Note | Using role-based ACLs as Filter-Id is not recommended. | 
More than one host can be authenticated on MDA-enabled and multiauth ports. The ACL policy applied for one host does not effect the traffic of another host. If only one host is authenticated on a multi-host port, and the other hosts gain network access without authentication, the ACL policy for the first host can be applied to the other connected hosts by specifying any in the source address.
Port-Based Authentication Manager CLI Commands
The authentication-manager interface-configuration commands control all the authentication methods, such as 802.1x, MAC authentication bypass, and web authentication. The authentication manager commands determine the priority and order of authentication methods applied to a connected host.
The authentication manager commands control generic authentication features, such as host-mode, violation mode, and the authentication timer. Generic authentication commands include the authentication host-mode , authentication violation , and authentication timer interface configuration commands.
802.1x-specific commands begin with the dot1x keyword. For example, the authentication port-control auto interface configuration command enables authentication on an interface.
To disable dot1x on a switch, remove the configuration globally by using the no dot1x system-auth-control command, and also remove it from all configured interfaces.
| Note | If 802.1x authentication is globally disabled, other authentication methods are still enabled on that port, such as web authentication. To re-enable dot1x on the switch, you must configure both the dot1x global and interface configurations. Incomplete configurations can cause high CPU utilization. | 
The authentication manager commands provide the same functionality as earlier 802.1x commands.
When filtering out verbose system messages generated by the authentication manager, the filtered content typically relates to authentication success. You can also filter verbose messages for 802.1x authentication and MAB authentication. There is a separate command for each authentication method:
-  
                                       			 
                                       The no authentication logging verbose global configuration command filters verbose messages from the authentication manager.
-  
                                       			 
                                       The no dot1x logging verbose global configuration command filters 802.1x authentication verbose messages.
-  
                                       			 
