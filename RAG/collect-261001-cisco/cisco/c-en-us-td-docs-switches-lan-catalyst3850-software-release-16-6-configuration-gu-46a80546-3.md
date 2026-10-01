---
id: collect-261001-cisco/cisco/c-en-us-td-docs-switches-lan-catalyst3850-software-release-16-6-configuration-gu-46a80546-3
title: "c-en-us-td-docs-switches-lan-catalyst3850-software-release-16-6-configuration-gu-46a80546"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: ["voice"]
source: docs/RAG/collect-261001-cisco/c-en-us-td-docs-switches-lan-catalyst3850-software-release-16-6-configuration-gu-46a80546.md
source_anchor: ""
source_lines: [76, 115]
sha256: 5b2d9a4e8edd29ac10675f63a5d57a035486447cdedbab36a3396168afe59f3f
---

# c-en-us-td-docs-switches-lan-catalyst3850-software-release-16-6-configuration-gu-46a80546

                                       The no mab logging verbose global configuration command filters MAC authentication bypass (MAB) verbose messages
| Table 2. Authentication Manager                                           		  Commands and Earlier 802.1x Commands |  |  | 
|---|---|---|
| The authentication manager commands in Cisco IOS Release 12.2(50)SE or later | The equivalent 802.1x commands in Cisco IOS Release 12.2(46)SE and earlier | Description | 
|---|---|---|
| authentication control-direction {both \| in} | dot1x control-direction {both \| in} | Enable 802.1x authentication with the wake-on-LAN (WoL) feature, and configure the port control as unidirectional or bidirectional. | 
| authentication event | dot1x auth-fail vlan dot1x critical (interface configuration) dot1x guest-vlan6 | Enable the restricted VLAN on a port. Enable the inaccessible-authentication-bypass feature. Specify an active VLAN as an 802.1x guest VLAN. | 
| authentication fallback fallback-profile | dot1x fallback fallback-profile | Configure a port to use web authentication as a fallback method for clients that do not support 802.1x authentication. | 
| authentication host-mode [multi-auth \| multi-domain \| multi-host \| single-host] | dot1x host-mode {single-host \| multi-host \| multi-domain} | Allow a single host (client) or multiple hosts on an 802.1x-authorized port. | 
| authentication order | mab | Provides the flexibility to define the order of authentication methods to be used. | 
| authentication periodic | dot1x reauthentication | Enable periodic re-authentication of the client. | 
| authentication port-control {auto \| force-authorized \| force-un authorized} | dot1x port-control {auto \| force-authorized \| force-unauthorized} | Enable manual control of the authorization state of the port. | 
| authentication timer | dot1x timeout | Set the 802.1x timers. | 
| authentication violation {protect \| restrict \| shutdown} | dot1x violation-mode {shutdown \| restrict \| protect} | Configure the violation modes that occur when a new device connects to a port or when a new device connects to a port after the maximum number of devices are connected to that port. | 
Ports in Authorized and Unauthorized States
During 802.1x authentication, depending on the switch port state, the switch can grant a client access to the network. The port starts in the unauthorized state. While in this state, the port that is not configured as a voice VLAN port disallows all ingress and egress traffic except for 802.1x authentication, CDP, and STP packets. When a client is successfully authenticated, the port changes to the authorized state, allowing all traffic for the client to flow normally. If the port is configured as a voice VLAN port, the port allows VoIP traffic and 802.1x protocol packets before the client is successfully authenticated.
| Note | CDP bypass is not supported and may cause a port to go into err-disabled state. | 
If a client that does not support 802.1x authentication connects to an unauthorized 802.1x port, the switch requests the client’s identity. In this situation, the client does not respond to the request, the port remains in the unauthorized state, and the client is not granted access to the network.
In contrast, when an 802.1x-enabled client connects to a port that is not running the 802.1x standard, the client initiates the authentication process by sending the EAPOL-start frame. When no response is received, the client sends the request for a fixed number of times. Because no response is received, the client begins sending frames as if the port is in the authorized state.
You control the port authorization state by using the authentication port-control interface configuration command and these keywords:
-  
                                 		  
                                 force-authorized —disables 802.1x authentication and causes the port to change to the authorized state without any authentication exchange required. The port sends and receives normal traffic without 802.1x-based authentication of the client. This is the default setting.
-  
                                 		  
                                 force-unauthorized —causes the port to remain in the unauthorized state, ignoring all attempts by the client to authenticate. The switch cannot provide authentication services to the client through the port.
-  
                                 		  
                                 auto —enables 802.1x authentication and causes the port to begin in the unauthorized state, allowing only EAPOL frames to be sent and received through the port. The authentication process begins when the link state of the port changes from down to up or when an EAPOL-start frame is received. The switch requests the identity of the client and begins relaying authentication messages between the client and the authentication server. Each client attempting to access the network is uniquely identified by the switch by using the client MAC address.
If the client is successfully authenticated (receives an Accept frame from the authentication server), the port state changes to authorized, and all frames from the authenticated client are allowed through the port. If the authentication fails, the port remains in the unauthorized state, but authentication can be retried. If the authentication server cannot be reached, the switch can resend the request. If no response is received from the server after the specified number of attempts, authentication fails, and network access is not granted.
When a client logs off, it sends an EAPOL-logoff message, causing the switch port to change to the unauthorized state.
If the link state of a port changes from up to down, or if an EAPOL-logoff frame is received, the port returns to the unauthorized state.
Port-Based Authentication and Switch Stacks
If a switch is added to or removed from a switch stack, 802.1x authentication is not affected as long as the IP connectivity between the RADIUS server and the stack remains intact. This statement also applies if the stack's active switch is removed from the switch stack. Note that if the active switch fails, a stack member becomes the new active switch of the stack by using the election process, and the 802.1x authentication process continues as usual.
If IP connectivity to the RADIUS server is interrupted because the switch that was connected to the server is removed or fails, these events occur:
-  
                                 		  
                                 Ports that are already authenticated and that do not have periodic re-authentication enabled remain in the authenticated state. Communication with the RADIUS server is not required.
-  
                                 		  
