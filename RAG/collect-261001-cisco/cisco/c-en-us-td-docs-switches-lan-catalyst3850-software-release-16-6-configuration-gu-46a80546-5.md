---
id: collect-261001-cisco/cisco/c-en-us-td-docs-switches-lan-catalyst3850-software-release-16-6-configuration-gu-46a80546-5
title: "c-en-us-td-docs-switches-lan-catalyst3850-software-release-16-6-configuration-gu-46a80546"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: ["voice"]
source: docs/RAG/collect-261001-cisco/c-en-us-td-docs-switches-lan-catalyst3850-software-release-16-6-configuration-gu-46a80546.md
source_anchor: ""
source_lines: [173, 222]
sha256: b783aed9abc5f6b7c558ceeadf356bf9611803e00e74e673e0175a788d87c630
---

# c-en-us-td-docs-switches-lan-catalyst3850-software-release-16-6-configuration-gu-46a80546

                                       IPv6 control packets: In IPv6 deployments, Router Advertisements (RA) are processed by hosts that are not supposed to receive them. When a host from one VLAN receives RA from a different VLAN, the host assign incorrect IPv6 address to itself. Such a host is unable to get access to the network. The workaround is to enable the IPv6 first hop security so that the broadcast ICMPv6 packets are converted to unicast and sent out from multi-auth enabled ports.. The packet is replicated for each client in multi-auth port belonging to the VLAN and the destination MAC is set to an individual client. Ports having one VLAN, ICMPv6 packets broadcast normally.
- 
                                       		  
                                       IP multicast: Multicast traffic destined to a multicast group gets replicated for different VLANs if the hosts on those VLANs join the multicast group. When two hosts in different VLANs join a multicast group (on the same mutli-auth port), two copies of each multicast packet are sent out from that port.
MAC Move
When a MAC address is authenticated on one switch port, that address is not allowed on another authentication manager-enabled port of the switch. If the switch detects that same MAC address on another authentication manager-enabled port, the address is not allowed.
There are situations where a MAC address might need to move from one port to another on the same switch. For example, when there is another device (for example a hub or an IP phone) between an authenticated host and a switch port, you might want to disconnect the host from the device and connect it directly to another port on the same switch.
You can globally enable MAC move so the device is reauthenticated on the new port. When a host moves to a second port, the session on the first port is deleted, and the host is reauthenticated on the new port. MAC move is supported on all host modes. (The authenticated host can move to any port on the switch, no matter which host mode is enabled on the that port.) When a MAC address moves from one port to another, the switch terminates the authenticated session on the original port and initiates a new authentication sequence on the new port. The MAC move feature applies to both voice and data hosts.
| Note | In open authentication mode, a MAC address is immediately moved from the original port to the new port, with no requirement for authorization on the new port. | 
MAC Replace
The MAC replace feature can be configured to address the violation that occurs when a host attempts to connect to a port where another host was previously authenticated.
| Note | This feature does not apply to ports in multi-auth mode, because violations are not triggered in that mode. It does not apply to ports in multiple host mode, because in that mode, only the first host requires authentication. | 
If you configure the authentication violation interface configuration command with the replace keyword, the authentication process on a port in multi-domain mode is:
- 
                                 		  
                                 A new MAC address is received on a port with an existing authenticated MAC address.
- 
                                 		  
                                 The authentication manager replaces the MAC address of the current data host on the port with the new MAC address.
- 
                                 		  
                                 The authentication manager initiates the authentication process for the new MAC address.
- 
                                 		  
                                 If the authentication manager determines that the new host is a voice host, the original voice host is removed.
If a port is in open authentication mode, any new MAC address is immediately added to the MAC address table.
802.1x Accounting
The 802.1x standard defines how users are authorized and authenticated for network access but does not keep track of network usage. 802.1x accounting is disabled by default. You can enable 802.1x accounting to monitor this activity on 802.1x-enabled ports:
- 
                                 User successfully authenticates.
- 
                                 User logs off.
- 
                                 Link-down occurs.
- 
                                 Re-authentication successfully occurs.
- 
                                 Re-authentication fails.
The switch does not log 802.1x accounting information. Instead, it sends this information to the RADIUS server, which must be configured to log accounting messages.
802.1x Accounting Attribute-Value Pairs
The information sent to the RADIUS server is represented in the form of Attribute-Value (AV) pairs. These AV pairs provide data for different applications. (For example, a billing application might require information that is in the Acct-Input-Octets or the Acct-Output-Octets attributes of a RADIUS packet.)
AV pairs are automatically sent by a switch that is configured for 802.1x accounting. Three types of RADIUS accounting packets are sent by a switch:
-  
                                    			 
                                    START–sent when a new user session starts
-  
                                    			 
                                    INTERIM–sent during an existing session for updates
-  
                                    			 
