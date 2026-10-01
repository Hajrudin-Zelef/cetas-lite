---
id: collect-261001-cisco/cisco/c-en-us-td-docs-switches-lan-catalyst3850-software-release-16-6-configuration-gu-46a80546-14
title: "c-en-us-td-docs-switches-lan-catalyst3850-software-release-16-6-configuration-gu-46a80546"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: ["voice"]
source: docs/RAG/collect-261001-cisco/c-en-us-td-docs-switches-lan-catalyst3850-software-release-16-6-configuration-gu-46a80546.md
source_anchor: ""
source_lines: [635, 675]
sha256: 087ba4ab7e7082f0916d415dd088396c828a7b8b30b47f0002b6d1fbb1462113
---

# c-en-us-td-docs-switches-lan-catalyst3850-software-release-16-6-configuration-gu-46a80546

                                 MDA can use MAC authentication bypass as a fallback mechanism to allow the switch port to connect to devices that do not support IEEE 802.1x authentication.
-  
                                 		  
                                 When a data or a voice device is detected on a port, its MAC address is blocked until authorization succeeds. If the authorization fails, the MAC address remains blocked for 5 minutes.
-  
                                 		  
                                 If more than five devices are detected on the data VLAN or more than one voice device is detected on the voice VLAN while a port is unauthorized, the port is error disabled.
-  
                                 		  
                                 When a port host mode is changed from single- or multihost to multidomain mode, an authorized data device remains authorized on the port. However, a Cisco IP phone that has been allowed on the port voice VLAN is automatically removed and must be reauthenticated on that port.
-  
                                 		  
                                 Active fallback mechanisms such as guest VLAN and restricted VLAN remain configured after a port changes from single- or multihost mode to multidomain mode.
-  
                                 		  
                                 Switching a port host mode from multidomain to single- or multihost mode removes all authorized devices from the port.
-  
                                 		  
                                 If a data domain is authorized first and placed in the guest VLAN, non-IEEE 802.1x-capable voice devices need to tag their packets on the voice VLAN to trigger authentication.
-  
                                 		  
                                 We do not recommend per-user ACLs with an MDA-enabled port. An authorized device with a per-user ACL policy might impact traffic on both the voice and data VLANs of the port. If used, only one device on the port should enforce per-user ACLs.
802.1x Supplicant and Authenticator Switches with Network Edge Access Topology (NEAT)
The Network Edge Access Topology (NEAT) feature extends identity to areas outside the wiring closet (such as conference rooms). This allows any type of device to authenticate on the port.
-  
                                 		  
                                 802.1x switch supplicant: You can configure a switch to act as a supplicant to another switch by using the 802.1x supplicant feature. This configuration is helpful in a scenario, where, for example, a switch is outside a wiring closet and is connected to an upstream switch through a trunk port. A switch configured with the 802.1x switch supplicant feature authenticates with the upstream switch for secure connectivity. Once the supplicant switch authenticates successfully the port mode changes from access to trunk in an authenticator switch. In a supplicant switch you must manually configure trunk when enabling CISP.
-  
                                 		  
                                 If the access VLAN is configured on the authenticator switch, it becomes the native VLAN for the trunk port after successful authentication.
In the default state, when you connect a supplicant switch to an authenticator switch that has BPDU guard enabled, the authenticator port could be error-disabled if it receives a Spanning Tree Protocol (STP) bridge protocol data unit (BPDU) packets before the supplicant switch has authenticated. Beginning with Cisco IOS Release 15.0(1)SE, you can control traffic exiting the supplicant port during the authentication period. Entering the dot1x supplicant controlled transient global configuration command temporarily blocks the supplicant port during authentication to ensure that the authenticator port does not shut down before authentication completes. If authentication fails, the supplicant port opens. Entering the no dot1x supplicant controlled transient global configuration command opens the supplicant port during the authentication period. This is the default behavior.
We strongly recommend using the dot1x supplicant controlled transientcommand on a supplicant switch when BPDU guard is enabled on the authenticator switch port with the spanning-tree bpduguard enable interface configuration command.
| Note | If you globally enable BPDU guard on the authenticator switch by using the spanning-tree portfast bpduguard default global configuration command, entering the dot1x supplicant controlled transient command does not prevent the BPDU violation. | 
You can enable MDA or multiauth mode on the authenticator switch interface that connects to one more supplicant switches. Multihost mode is not supported on the authenticator switch interface.
When you reboot an authenticator switch with single-host mode enabled on the interface, the interface may move to err-disabled state before authentication. To recover from err-disabled state, flap the authenticator port to activate the interface again and initiate authentication.
Use the dot1x supplicant force-multicast global configuration command on the supplicant switch for Network Edge Access Topology (NEAT) to work in all host modes.
-  
                                 		  
                                 Host Authorization: Ensures that only traffic from authorized hosts (connecting to the switch with supplicant) is allowed on the network. The switches use Client Information Signalling Protocol (CISP) to send the MAC addresses connecting to the supplicant switch to the authenticator switch.
-  
                                 		  
