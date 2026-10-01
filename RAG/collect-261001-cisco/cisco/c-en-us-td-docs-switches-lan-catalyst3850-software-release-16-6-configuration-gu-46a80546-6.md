---
id: collect-261001-cisco/cisco/c-en-us-td-docs-switches-lan-catalyst3850-software-release-16-6-configuration-gu-46a80546-6
title: "c-en-us-td-docs-switches-lan-catalyst3850-software-release-16-6-configuration-gu-46a80546"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: ["voice"]
source: docs/RAG/collect-261001-cisco/c-en-us-td-docs-switches-lan-catalyst3850-software-release-16-6-configuration-gu-46a80546.md
source_anchor: ""
source_lines: [223, 301]
sha256: c1487bf36e14f6046fa3bdd6611ab2d1046c8bc80bc79ec94fe696f720709674
---

# c-en-us-td-docs-switches-lan-catalyst3850-software-release-16-6-configuration-gu-46a80546

                                    STOP–sent when a session terminates
| Note | To view debug logs for RADIUS and AAA, use the show platform software trace message smd command. For more information, see the Tracing Commands section in Command Reference Guide, Cisco IOS XE Denali 16.1.1. | 
This table lists the AV pairs and when they are sent are sent by the switch.
| Table 3. Accounting AV Pairs |  |  |  |  | 
|---|---|---|---|---|
| Attribute Number | AV Pair Name | START | INTERIM | STOP | 
|---|---|---|---|---|
| Attribute[1] | User-Name | Always | Always | Always | 
| Attribute[4] | NAS-IP-Address | Always | Always | Always | 
| Attribute[5] | NAS-Port | Always | Always | Always | 
| Attribute[8] | Framed-IP-Address | Never | Sometimes3 | Sometimes | 
| Attribute[30] | Called-Station-ID | Always | Always | Always | 
| Attribute[31] | Calling-Station-ID | Always | Always | Always | 
| Attribute[40] | Acct-Status-Type | Always | Always | Always | 
| Attribute[41] | Acct-Delay-Time | Always | Always | Always | 
| Attribute[42] | Acct-Input-Octets | Never | Always | Always | 
| Attribute[43] | Acct-Output-Octets | Never | Always | Always | 
| Attribute[47] | Acct-Input-Packets | Never | Always | Always | 
| Attribute[48] | Acct-Output-Packets | Never | Always | Always | 
| Attribute[44] | Acct-Session-ID | Always | Always | Always | 
| Attribute[45] | Acct-Authentic | Always | Always | Always | 
| Attribute[46] | Acct-Session-Time | Never | Always | Always | 
| Attribute[49] | Acct-Terminate-Cause | Never | Never | Always | 
| Attribute[61] | NAS-Port-Type | Always | Always | Always | 
802.1x Readiness Check
The 802.1x readiness check monitors 802.1x activity on all the switch ports and displays information about the devices connected to the ports that support 802.1x. You can use this feature to determine if the devices connected to the switch ports are 802.1x-capable. You use an alternate authentication such as MAC authentication bypass or web authentication for the devices that do not support 802.1x functionality.
This feature only works if the supplicant on the client supports a query with the NOTIFY EAP notification packet. The client must respond within the 802.1x timeout value.
Switch-to-RADIUS-Server Communication
RADIUS security servers are identified by their hostname or IP address, hostname and specific UDP port numbers, or IP address and specific UDP port numbers. The combination of the IP address and UDP port number creates a unique identifier, which enables RADIUS requests to be sent to multiple UDP ports on a server at the same IP address. If two different host entries on the same RADIUS server are configured for the same service—for example, authentication—the second host entry configured acts as the fail-over backup to the first one. The RADIUS host entries are tried in the order that they were configured.
802.1x Authentication with VLAN Assignment
The switch supports 802.1x authentication with VLAN assignment. After successful 802.1x authentication of a port, the RADIUS server sends the VLAN assignment to configure the switch port. The RADIUS server database maintains the username-to-VLAN mappings, assigning the VLAN based on the username of the client connected to the switch port. You can use this feature to limit network access for certain users.
Voice device authentication is supported with multidomain host mode in Cisco IOS Release 12.2(37)SE. In Cisco IOS Release 12.2(40)SE and later, when a voice device is authorized and the RADIUS server returned an authorized VLAN, the voice VLAN on the port is configured to send and receive packets on the assigned voice VLAN. Voice VLAN assignment behaves the same as data VLAN assignment on multidomain authentication (MDA)-enabled ports.
When configured on the switch and the RADIUS server, 802.1x authentication with VLAN assignment has these characteristics:
- 
                                 		  
                                 If no VLAN is supplied by the RADIUS server or if 802.1x authentication is disabled, the port is configured in its access VLAN after successful authentication. Recall that an access VLAN is a VLAN assigned to an access port. All packets sent from or received on this port belong to this VLAN.
- 
                                 		  
                                 If 802.1x authentication is enabled but the VLAN information from the RADIUS server is not valid, authorization fails and configured VLAN remains in use. This prevents ports from appearing unexpectedly in an inappropriate VLAN because of a configuration error. Configuration errors could include specifying a VLAN for a routed port, a malformed VLAN ID, a nonexistent or internal (routed port) VLAN ID, an RSPAN VLAN, a shut down or suspended VLAN. In the case of a multidomain host port, configuration errors can also be due to an attempted assignment of a data VLAN that matches the configured or assigned voice VLAN ID (or the reverse).
- 
                                 		  
                                 If 802.1x authentication is enabled and all information from the RADIUS server is valid, the authorized device is placed in the specified VLAN after authentication.
- 
                                 		  
                                 If the multiple-hosts mode is enabled on an 802.1x port, all hosts are placed in the same VLAN (specified by the RADIUS server) as the first authenticated host.
- 
                                 		  
                                 Enabling port security does not impact the RADIUS server-assigned VLAN behavior.
- 
                                 		  
                                 If 802.1x authentication is disabled on the port, it is returned to the configured access VLAN and configured voice VLAN.
- 
                                 				
                                 If an 802.1x port is authenticated and put in the RADIUS server-assigned VLAN, any change to the port access VLAN configuration does not take effect. In the case of a multidomain host, the same applies to voice devices when the port is fully authorized with these exceptions: 
  - 
                                       						
                                       If the VLAN configuration change of one device results in matching the other device configured or assigned VLAN, then authorization of all devices on the port is terminated and multidomain host mode is disabled until a valid configuration is restored where data and voice device configured VLANs no longer match.
  - 
                                       						
                                       If a voice device is authorized and is using a downloaded voice VLAN, the removal of the voice VLAN configuration, or modifying the configuration value to dot1p or untagged results in voice device un-authorization and the disablement of multi-domain host mode.
- 
                                       						
                                       
When the port is in the force authorized, force unauthorized, unauthorized, or shutdown state, it is put into the configured access VLAN.
To configure VLAN assignment you need to perform these tasks:
- 
                                 				
                                 Enable AAA authorization by using the network keyword to allow interface configuration from the RADIUS server.
- 
                                 				
                                 Enable 802.1x authentication. (The VLAN assignment feature is automatically enabled when you configure 802.1x authentication on an access port).
- 
                                 				
                                 Assign vendor-specific tunnel attributes in the RADIUS server. The RADIUS server must return these attributes to the switch: 
  - 
                                       						
                                       [64] Tunnel-Type = VLAN
  - 
                                       						
