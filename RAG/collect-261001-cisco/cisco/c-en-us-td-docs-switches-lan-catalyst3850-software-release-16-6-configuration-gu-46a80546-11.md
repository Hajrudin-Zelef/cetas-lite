---
id: collect-261001-cisco/cisco/c-en-us-td-docs-switches-lan-catalyst3850-software-release-16-6-configuration-gu-46a80546-11
title: "c-en-us-td-docs-switches-lan-catalyst3850-software-release-16-6-configuration-gu-46a80546"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: ["distribution", "voice"]
source: docs/RAG/collect-261001-cisco/c-en-us-td-docs-switches-lan-catalyst3850-software-release-16-6-configuration-gu-46a80546.md
source_anchor: ""
source_lines: [464, 501]
sha256: 88aa63087bf00251b5f331a5afad0b352cf2d82cdb01a15ee3f373f3d3b76020
---

# c-en-us-td-docs-switches-lan-catalyst3850-software-release-16-6-configuration-gu-46a80546

                                    If the new active switch is elected, the link between the switch stack and RADIUS server might change, and the new stack immediately sends keepalive packets to update the status of the RADIUS servers. If the server status changes from dead to alive, the switch re-authenticates all switch ports in the critical-authentication state.
When a member is added to the stack, the stack's active switch sends the member the server status.
802.1x Critical Voice VLAN
When an IP phone connected to a port is authenticated by the Cisco Identity Services Engine (ISE), the phone is put into the voice domain. If the ISE is not reachable, the switch cannot determine if the device is a voice device. If the server is unavailable, the phone cannot access the voice network and therefore cannot operate.
For data traffic, you can configure inaccessible authentication bypass, or critical authentication, to allow traffic to pass through on the native VLAN when the server is not available. If the RADIUS authentication server is unavailable (down) and inaccessible authentication bypass is enabled, the switch grants the client access to the network and puts the port in the critical-authentication state in the RADIUS-configured or the user-specified access VLAN. When the switch cannot reach the configured RADIUS servers and new hosts cannot be authenticated, the switch connects those hosts to critical ports. A new host trying to connect to the critical port is moved to a user-specified access VLAN, the critical VLAN, and granted limited authentication.
| Note | Dynamic assignment of critical voice VLAN is not supported with nested service templates. It causes the device to switch between VLANs continuously in a loop. | 
You can enter the authentication event server dead action authorize voice interface configuration command to configure the critical voice VLAN feature. When the ISE does not respond, the port goes into critical authentication mode. When traffic coming from the host is tagged with the voice VLAN, the connected device (the phone) is put in the configured voice VLAN for the port. The IP phones learn the voice VLAN identification through Cisco Discovery Protocol (Cisco devices) or through LLDP or DHCP.
You can configure the voice VLAN for a port by entering the switchport voice vlan vlan-id interface configuration command.
This feature is supported in multidomain and multi-auth host modes. Although you can enter the command when the switch in single-host or multi-host mode, the command has no effect unless the device changes to multidomain or multi-auth host mode.
802.1x User Distribution
You can configure 802.1x user distribution to load-balance users with the same group name across multiple different VLANs.
The VLANs are either supplied by the RADIUS server or configured through the switch CLI under a VLAN group name.
- 
                                 Configure the RADIUS server to send more than one VLAN name for a user. The multiple VLAN names can be sent as part of the response to the user. The 802.1x user distribution tracks all the users in a particular VLAN and achieves load balancing by moving the authorized user to the least populated VLAN.
- 
                                 Configure the RADIUS server to send a VLAN group name for a user. The VLAN group name can be sent as part of the response to the user. You can search for the selected VLAN group name among the VLAN group names that you configured by using the switch CLI. If the VLAN group name is found, the corresponding VLANs under this VLAN group name are searched to find the least populated VLAN. Load balancing is achieved by moving the corresponding authorized user to that VLAN. 
 Note
 The RADIUS server can send the VLAN information in any combination of VLAN-IDs, VLAN names, or VLAN groups. 
802.1x User Distribution Configuration Guidelines
- 
                                    Confirm that at least one VLAN is mapped to the VLAN group.
- 
                                    You can map more than one VLAN to a VLAN group.
- 
                                    You can modify the VLAN group by adding or deleting a VLAN.
- 
                                    When you clear an existing VLAN from the VLAN group name, none of the authenticated ports in the VLAN are cleared, but the mappings are removed from the existing VLAN group.
- 
                                    If you clear the last VLAN from the VLAN group name, the VLAN group is cleared.
- 
                                    You can clear a VLAN group even when the active VLANs are mapped to the group. When you clear a VLAN group, none of the ports or users that are in the authenticated state in any VLAN within the group are cleared, but the VLAN mappings to the VLAN group are cleared.
IEEE 802.1x Authentication with Voice VLAN Ports
A voice VLAN port is a special access port associated with two VLAN identifiers:
- 
                                 		  
                                 VVID to carry voice traffic to and from the IP phone. The VVID is used to configure the IP phone connected to the port.
- 
                                 		  
